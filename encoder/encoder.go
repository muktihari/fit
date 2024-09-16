// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	ErrNilWriter     = errorString("nil writer")
	ErrEmptyMessages = errorString("empty messages")
	ErrMissingFileId = errorString("missing file_id mesg")

	ErrWriterAtOrWriteSeekerIsExpected = errorString("io.WriterAt or io.WriteSeeker is expected")
)

// HeaderOption is header option for encoding message's Header.
type HeaderOption byte

const (
	// HeaderOptionNormal is the default header option. This option allow us to use local message type 0-15
	// as the maximum number of allowed message definition interleave.
	HeaderOptionNormal HeaderOption = 0

	// HeaderOptionCompressedTimestamp optimizes file size by compressing timestamp's field in a message into
	// its message header. Saves 7 bytes per message when its timestamp is compressed: 3 bytes for field definition
	// and 4 bytes for the uint32 timestamp value. When this option is selected, only local messages type 0-3 is available.
	HeaderOptionCompressedTimestamp HeaderOption = 1
)

// Encoder is FIT file encoder. See New() for details.
type Encoder struct {
	w                 io.Writer   // A writer to write the encoded FIT bytes.
	n                 int64       // Total bytes written to w, will keep counting for every Encode invocation.
	lastFileHeaderPos int64       // The byte position of the last header.
	crc16             hash.Hash16 // Calculate the CRC-16 checksum for ensuring header and message integrity.

	options           options          // Encoder's options.
	protocolValidator *proto.Validator // Validates message's properties should match the targeted protocol version requirements.
	localMesgNumLRU   *lru             // LRU cache for writing local message definition

	dataSize uint32 // Data size of messages in bytes for a single FIT file.

	// This timestamp reference serves as current active timestamp when 'headerOptionCompressedTimestamp' is specified.
	// The first timestamp value is retrieved from the first message containing a valid timestamp field,
	// and will change every rollover event occurrence.
	timestampReference uint32

	mesgDef proto.MessageDefinition // Temporary message definition to reduce alloc.

	// Dynamic-sized buffer for encoding, starting at 1536 bytes (see PR #415 and #416 for details).
	// It starts small but grows as needed and may only grow when using Message's MarshalAppend.
	buf []byte
}

type options struct {
	messageValidator MessageValidator
	writeBufferSize  int
	protocolVersion  proto.Version
	endianness       byte
	headerOption     HeaderOption
	localMessageType byte
}

func defaultOptions() options {
	return options{
		endianness:      proto.LittleEndian,
		headerOption:    HeaderOptionNormal,
		writeBufferSize: defaultWriteBufferSize,
	}
}

// Option is Encoder's option.
type Option func(o *options)

// WithProtocolVersion directs the Encoder to use specific ProtocolVersion for the entire encoding.
// By default, Encoder will use ProtocolVersion in FileHeader for each FIT file, if it's not specified,
// it will use proto.V1. This option overrides the FileHeader's ProtocolVersion and forces all FIT
// files to use this ProtocolVersion during encoding.
//
// NOTE: If the given protocolVersion is not supported, the Protocol Version will not be changed.
// Please validate using proto.Validate when putting user-defined Protocol Version to check
// whether it is supported or not. Or just use predefined Protocol Version constants such as
// proto.V1, proto.V2, etc, which the validity is ensured.
func WithProtocolVersion(protocolVersion proto.Version) Option {
	return func(o *options) {
		if proto.Validate(protocolVersion) == nil {
			o.protocolVersion = protocolVersion
		}
	}
}

// WithMessageValidator directs the Encoder to use this message validator instead of the default one.
func WithMessageValidator(validator MessageValidator) Option {
	return func(o *options) {
		if validator != nil {
			o.messageValidator = validator
		}
	}
}

// WithBigEndian directs the Encoder to encode values in Big-Endian bytes order (default: Little-Endian).
func WithBigEndian() Option {
	return func(o *options) { o.endianness = proto.BigEndian }
}

// WithHeaderOption directs the Encoder to use this option instead of default HeaderOptionNormal with local message type zero.
//
//   - If HeaderOptionNormal is selected, valid local message type value is 0-15; invalid values will be treated as 15.
//
//   - If HeaderOptionCompressedTimestamp is selected, valid local message type value is 0-3; invalid values will be treated as 3.
//
//     Saves 7 bytes per message when its timestamp is compressed: 3 bytes for field definition
//     and 4 bytes for the uint32 timestamp value.
//
//   - Otherwise, no change will be made and the Encoder will use default values.
//
// NOTE: To minimize the required RAM for decoding, it's recommended to use a minimal number of local message type.
// For instance, embedded devices may only support decoding data from local message type 0. Additionally,
// multiple local message types should be avoided in file types like settings, where messages of the same type
// can be grouped together.
func WithHeaderOption(headerOption HeaderOption, localMessageType byte) Option {
	return func(o *options) {
		switch headerOption {
		case HeaderOptionNormal:
			if localMessageType > 15 {
				localMessageType = 15
			}
		case HeaderOptionCompressedTimestamp:
			if localMessageType > 3 {
				localMessageType = 3
			}
		default:
			return
		}
		o.headerOption = headerOption
		o.localMessageType = localMessageType
	}
}

// WithWriteBufferSize directs the Encoder to use this buffer size for writing to io.Writer instead of default 4096.
// When size <= 0, the Encoder will write directly to io.Writer without buffering.
func WithWriteBufferSize(size int) Option {
	return func(o *options) { o.writeBufferSize = size }
}

// New returns a FIT File Encoder to encode FIT data to given w.
//
// # Encoding Strategy
//
// Since an invalid FileHeader means an invalid FIT file, we need to ensure that the FileHeader is correct,
// specifically the FileHeader's DataSize (size of messages in bytes) and FileHeader's CRC checksum should
// be correct after everything is written.
//
// There are two strategies to achieve that and it depends on what kind of [io.Writer] is provided:
//
//   - [io.WriteSeeker] or [io.WriterAt]: Encoder will update the FileHeader's DataSize and CRC after
//     the encoding process is completed since we can write at a specific byte position, making it more
//     ideal and efficient. If the given [io.Writer] implements both of them, it will be treated as [io.WriteSeeker].
//   - [io.Writer]: Encoder needs to iterate through the messages once to calculate the FileHeader's DataSize
//     and CRC by writing to [io.Discard], then re-iterate through the messages again for the actual writing.
//
// Caveats:
//
//   - If [io.Writer] is an [*os.File] opened with O_APPEND, the behavior of the Encoder is not specified.
//     If you intend to append the encoding result to an existing file to create a chained FIT file,
//     open the file without O_APPEND then file.Seek(0, io.SeekEnd) before putting the file into the Encoder.
//   - When using [io.WriterAt], the given [io.Writer] is expected to be empty (brand new), otherwise, the resulting
//     data will be corrupted.
//
// Note: Encoder already implements efficient io.Writer buffering, so there's no need to wrap 'w' with a buffer;
// doing so will only reduce performance. If you don't want the Encoder to buffer the writing, please direct the
// Encoder to do so using WithWriteBufferSize(0).
func New(w io.Writer, opts ...Option) *Encoder {
	e := &Encoder{
		crc16:             crc16.New(nil),
		protocolValidator: new(proto.Validator),
		localMesgNumLRU:   new(lru),
		buf:               make([]byte, 0, 1536),
		mesgDef: proto.MessageDefinition{
			FieldDefinitions:          make([]proto.FieldDefinition, 0, 255),
			DeveloperFieldDefinitions: make([]proto.DeveloperFieldDefinition, 0, 255),
		},
	}
	e.Reset(w, opts...)
	return e
}

// Reset resets the Encoder to write its output to w and reset previous options to
// default options so any options needs to be inputed again. It is similar to New()
// but it retains the underlying storage for use by future encode to reduce memory allocs.
func (e *Encoder) Reset(w io.Writer, opts ...Option) {
	e.n = 0
	e.lastFileHeaderPos = 0

	e.options = defaultOptions()
	for i := range opts {
		opts[i](&e.options)
	}

	e.w = newWriteBuffer(w, e.options.writeBufferSize)

	if e.options.messageValidator == nil {
		e.options.messageValidator = NewMessageValidator()
	}

	e.reset()

	e.localMesgNumLRU.ResetWithNewSize(e.options.localMessageType + 1)
}

// reset resets the encoder's data that is being used for encoding,
// allowing the encoder to be reused for writing the subsequent chained FIT file.
func (e *Encoder) reset() {
	e.options.messageValidator.Reset()
	e.crc16.Reset()
	e.localMesgNumLRU.Reset()
	e.dataSize = 0
	e.timestampReference = 0
}

// Encode encodes FIT into the dest writer. Only FIT's Messages is required, while FileHeader and CRC will be
// filled automatically by the Encoder. However, we allow for custom FileHeader, such as when a user intentionally
// specifies a Size as 12 (legacy), a custom ProfileVersion or a custom ProtocolVersion.
// In these cases, those values will be encoded as-is, irrespective of the current SDK profile version.
//
// Multiple FIT files can be chained together into a single FIT file by calling Encode for each FIT data.
//
//	for _, fit := range fits {
//	   err := enc.Encode(fit)
//	}
//
// Encode chooses which strategy to use for encoding the data based on given writer.
func (e *Encoder) Encode(fit *proto.FIT) (err error) {
	switch e.w.(type) {
	case io.WriterAt, io.WriteSeeker:
		err = e.encodeWithDirectUpdateStrategy(fit)
	case io.Writer:
		err = e.encodeWithEarlyCheckStrategy(fit)
	default:
		err = ErrNilWriter
	}
	e.reset()
	if err != nil {
		return
	}
	if f, ok := e.w.(flusher); ok {
		return f.Flush()
	}
	return
}

// encodeWithDirectUpdateStrategy encodes all data to file, after completing,
// it updates the actual size of the messages that being written to the proto.
func (e *Encoder) encodeWithDirectUpdateStrategy(fit *proto.FIT) error {
	if err := e.encodeFileHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessages(fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}
	if err := e.updateFileHeader(&fit.FileHeader); err != nil {
		return err
	}
	return nil
}

// encodeWithEarlyCheckStrategy does early calculation of the size of the messages
// that will be written and then do the encoding process.
func (e *Encoder) encodeWithEarlyCheckStrategy(fit *proto.FIT) error {
	if err := e.calculateDataSize(fit); err != nil {
		return err
	}
	if err := e.encodeFileHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessages(fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) encodeFileHeader(header *proto.FileHeader) error {
	e.lastFileHeaderPos = e.n

	if header.Size != 12 { // allow legacy
		header.Size = 14
	}
	if header.ProfileVersion == 0 { // only change when zero to allow custom profile version
		header.ProfileVersion = profile.Version
	}

	if e.options.protocolVersion != 0 { // Override regardless the value in FileHeader.
		header.ProtocolVersion = e.options.protocolVersion
	} else if header.ProtocolVersion == 0 { // Default when not specified in FileHeader.
		header.ProtocolVersion = proto.V1
	}
	e.protocolValidator.SetProtocolVersion(header.ProtocolVersion)

	header.DataType = proto.DataTypeFIT
	header.CRC = 0 // recalculated

	b, _ := header.MarshalAppend(e.buf[:0])

	if header.Size != 14 {
		n, err := e.w.Write(b[:header.Size])
		e.n += int64(n)
		return err
	}

	_, _ = e.crc16.Write(b[:12])
	binary.LittleEndian.PutUint16(b[12:14], e.crc16.Sum16())
	header.CRC = e.crc16.Sum16()

	e.crc16.Reset() // this hash will be re-used for calculating data integrity.

	n, err := e.w.Write(b)
	e.n += int64(n)

	return err
}

const errInternal = errorString("encoder internal error")

// updateFileHeader updates the FileHeader if the DataSize is changed.
// The caller MUST ensure that e.w is either an io.WriterAt or an io.WriteSeeker.
func (e *Encoder) updateFileHeader(header *proto.FileHeader) (err error) {
	if header.DataSize == e.dataSize {
		return nil
	}

	header.DataSize = e.dataSize

	b, _ := header.MarshalAppend(e.buf[:0])

	if header.Size == 14 {
		_, _ = e.crc16.Write(b[:12]) // recalculate CRC Checksum since FileHeader is changed.
		header.CRC = e.crc16.Sum16() // update crc in FileHeader
		binary.LittleEndian.PutUint16(b[12:14], e.crc16.Sum16())
		e.crc16.Reset()
	}

	switch w := e.w.(type) {
	case io.WriteSeeker:
		// Ensure that we only write in our own data.
		size := e.n - e.lastFileHeaderPos
		_, err = w.Seek(-size, io.SeekCurrent)
		if err != nil {
			return err
		}
		var n int
		n, err = w.Write(b)
		if err != nil {
			return err
		}
		_, err = w.Seek(size-int64(n), io.SeekCurrent)
		return err
	case io.WriterAt:
		// This might write at the wrong offset if the writer was not
		// empty before being assigned to the Encoder. We only track
		// our own data and do not have context beyond it.
		_, err = w.WriteAt(b, e.lastFileHeaderPos)
		return err
	default:
		return errInternal // should not reach here except we code wrong implementation.
	}
}

// calculateDataSize calculates total data size of the messages by counting bytes written to io.Discard.
func (e *Encoder) calculateDataSize(fit *proto.FIT) error {
	n := e.n
	w := e.w

	e.w = io.Discard

	if err := e.encodeMessages(fit.Messages); err != nil {
		return fmt.Errorf("calculate data size: %w", err)
	}

	fit.FileHeader.DataSize = e.dataSize // update FileHeader's DataSize of the actual messages size
	e.reset()

	e.n = n
	e.w = w

	return nil
}

func (e *Encoder) encodeMessages(messages []proto.Message) error {
	if len(messages) == 0 {
		return ErrEmptyMessages
	}

	if messages[0].Num != mesgnum.FileId {
		return ErrMissingFileId
	}

	for i := range messages {
		mesg := &messages[i]
		if err := e.encodeMessage(mesg); err != nil {
			return fmt.Errorf("encode failed: at byte pos: %d, message index: %d, num: %d (%s): %w",
				e.n, i, mesg.Num, mesg.Num.String(), err)
		}
	}

	return nil
}

// encodeMessage marshals and encodes message definition and its message into w.
func (e *Encoder) encodeMessage(mesg *proto.Message) (err error) {
	mesg.Header = proto.MesgNormalHeaderMask

	if err = e.options.messageValidator.Validate(mesg); err != nil {
		return fmt.Errorf("message validation failed: %w", err)
	}

	var compressed bool
	if e.options.headerOption == HeaderOptionCompressedTimestamp {
		if e.w == io.Discard {
			// NOTE: Only for calculating data size (Early Check Strategy)
			var timestampField proto.Field
			var i int
			for i = range mesg.Fields {
				if mesg.Fields[i].Num == proto.FieldNumTimestamp {
					timestampField = mesg.Fields[i]
					break
				}
			}
			prevLen := len(mesg.Fields)
			compressed = e.compressTimestampIntoHeader(mesg)
			if compressed {
				defer func() { // Revert: put timestamp field back at original index
					mesg.Fields = mesg.Fields[:prevLen]
					copy(mesg.Fields[i+1:], mesg.Fields[i:])
					mesg.Fields[i] = timestampField
				}()
			}
		} else {
			compressed = e.compressTimestampIntoHeader(mesg)
		}
	}

	mesgDef := e.newMessageDefinition(mesg)
	if err := e.protocolValidator.ValidateMessageDefinition(mesgDef); err != nil {
		return err
	}

	b, _ := mesgDef.MarshalAppend(e.buf[:0])
	localMesgNum, isNewMesgDef := e.localMesgNumLRU.Put(b) // This might alloc memory since we need to copy the item.

	b[0] |= localMesgNum // Update the message definition header.
	if compressed {
		mesg.Header |= (localMesgNum << proto.CompressedBitShift)
	} else {
		mesg.Header |= localMesgNum
	}

	var n int
	if isNewMesgDef {
		n, err = e.w.Write(b)
		e.n, e.dataSize = e.n+int64(n), e.dataSize+uint32(n)
		if err != nil {
			return fmt.Errorf("write message definition failed: %w", err)
		}
		_, _ = e.crc16.Write(b)
	}

	// At this point, e.buf may grow. Re-assign e.buf in case slice has grown.
	e.buf, err = mesg.MarshalAppend(e.buf[:0], mesgDef.Architecture)
	if err != nil {
		return fmt.Errorf("marshal mesg failed: %w", err)
	}

	n, err = e.w.Write(e.buf)
	e.n, e.dataSize = e.n+int64(n), e.dataSize+uint32(n)
	if err != nil {
		return fmt.Errorf("write message failed: %w", err)
	}
	_, _ = e.crc16.Write(e.buf)

	return nil
}

func (e *Encoder) compressTimestampIntoHeader(mesg *proto.Message) (ok bool) {
	timestamp := mesg.FieldValueByNum(proto.FieldNumTimestamp).Uint32()
	if timestamp == basetype.Uint32Invalid {
		return false // not supported
	}

	if timestamp < uint32(typedef.DateTimeMin) {
		return false
	}

	// The 5-bit time offset rolls over every 32 seconds, it is necessary that the difference
	// between timestamp and timestamp reference be measured less than 32 seconds apart.
	if (timestamp - e.timestampReference) > proto.CompressedTimeMask {
		e.timestampReference = timestamp
		return false // Rollover event occurs, keep it as it is.
	}

	timeOffset := byte(timestamp & proto.CompressedTimeMask)
	mesg.Header = proto.MesgCompressedHeaderMask | timeOffset
	mesg.RemoveFieldByNum(proto.FieldNumTimestamp)
	return true
}

func (e *Encoder) newMessageDefinition(mesg *proto.Message) *proto.MessageDefinition {
	e.mesgDef.Header = proto.MesgDefinitionMask
	e.mesgDef.Reserved = 0
	e.mesgDef.Architecture = e.options.endianness
	e.mesgDef.MesgNum = mesg.Num
	e.mesgDef.FieldDefinitions = e.mesgDef.FieldDefinitions[:0]
	e.mesgDef.DeveloperFieldDefinitions = e.mesgDef.DeveloperFieldDefinitions[:0]

	for i := range mesg.Fields {
		e.mesgDef.FieldDefinitions = append(e.mesgDef.FieldDefinitions, proto.FieldDefinition{
			Num:      mesg.Fields[i].Num,
			Size:     byte(proto.Sizeof(mesg.Fields[i].Value)),
			BaseType: mesg.Fields[i].BaseType,
		})
	}

	if len(mesg.DeveloperFields) == 0 {
		return &e.mesgDef
	}

	e.mesgDef.Header |= proto.DevDataMask
	for i := range mesg.DeveloperFields {
		e.mesgDef.DeveloperFieldDefinitions = append(e.mesgDef.DeveloperFieldDefinitions, proto.DeveloperFieldDefinition{
			Num:                mesg.DeveloperFields[i].Num,
			Size:               byte(proto.Sizeof(mesg.DeveloperFields[i].Value)),
			DeveloperDataIndex: mesg.DeveloperFields[i].DeveloperDataIndex,
		})
	}

	return &e.mesgDef
}

func (e *Encoder) encodeCRC() error {
	b := binary.LittleEndian.AppendUint16(e.buf[:0], e.crc16.Sum16())

	n, err := e.w.Write(b)
	e.n += int64(n)
	if err != nil {
		return fmt.Errorf("could not write crc: %w", err)
	}

	e.crc16.Reset()

	return nil
}

// EncodeWithContext is similar to Encode but with respect to context propagation.
func (e *Encoder) EncodeWithContext(ctx context.Context, fit *proto.FIT) (err error) {
	switch e.w.(type) {
	case io.WriterAt, io.WriteSeeker:
		err = e.encodeWithDirectUpdateStrategyWithContext(ctx, fit)
	case io.Writer:
		err = e.encodeWithEarlyCheckStrategyWithContext(ctx, fit)
	default:
		err = ErrNilWriter
	}
	e.reset()
	if err != nil {
		return
	}
	if f, ok := e.w.(flusher); ok {
		return f.Flush()
	}
	return
}

func (e *Encoder) encodeWithDirectUpdateStrategyWithContext(ctx context.Context, fit *proto.FIT) error {
	if err := e.encodeFileHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessagesWithContext(ctx, fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}
	if err := e.updateFileHeader(&fit.FileHeader); err != nil {
		return err
	}
	return nil
}

func (e *Encoder) calculateDataSizeWithContext(ctx context.Context, fit *proto.FIT) error {
	n := e.n
	w := e.w

	e.w = io.Discard

	if err := e.encodeMessagesWithContext(ctx, fit.Messages); err != nil {
		return fmt.Errorf("calculate data size: %w", err)
	}

	fit.FileHeader.DataSize = e.dataSize // update FileHeader's DataSize of the actual messages size
	e.reset()

	e.n = n
	e.w = w

	return nil
}

func (e *Encoder) encodeWithEarlyCheckStrategyWithContext(ctx context.Context, fit *proto.FIT) error {
	if err := e.calculateDataSizeWithContext(ctx, fit); err != nil {
		return err
	}
	if err := e.encodeFileHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessagesWithContext(ctx, fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}
	return nil
}

func (e *Encoder) encodeMessagesWithContext(ctx context.Context, messages []proto.Message) error {
	if len(messages) == 0 {
		return ErrEmptyMessages
	}

	if messages[0].Num != mesgnum.FileId {
		return ErrMissingFileId
	}

	for i := range messages {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		mesg := &messages[i]
		if err := e.encodeMessage(mesg); err != nil {
			return fmt.Errorf("encode failed: at byte pos: %d, message index: %d, num: %d (%s): %w",
				e.n, i, mesg.Num, mesg.Num.String(), err)
		}
	}

	return nil
}

// StreamEncoder turns this Encoder into StreamEncoder to encode per message basis or in streaming fashion.
// It returns an error if the Encoder's Writer does not implement io.WriterAt or io.WriteSeeker.
// After invoking this method, it is recommended not to use the Encoder to avoid undefined behavior.
func (e *Encoder) StreamEncoder() (*StreamEncoder, error) {
	switch e.w.(type) {
	case io.WriterAt, io.WriteSeeker:
		return &StreamEncoder{enc: e}, nil
	default:
		return nil, fmt.Errorf("could not convert encoder into stream encoder: %w",
			ErrWriterAtOrWriteSeekerIsExpected)
	}
}
