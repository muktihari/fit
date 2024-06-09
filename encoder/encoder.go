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

// headerOption is header option.
type headerOption byte

const (
	littleEndian = 0
	bigEndian    = 1

	// headerOptionNormal is the default header option.
	// This option has two sub-option to select from:
	//   1. LocalMessageTypeZero [Default]
	// 		Optimized for all devices. It only use LocalMesgNum 0.
	//   2. MultipleLocalMessageTypes
	//      Using multiple local message types optimizes file size by avoiding the need to interleave different
	//      message definition. The number of multiple local message type can be specified between 0-15.
	headerOptionNormal headerOption = 0

	// Optimized to reduce file size way further by compressing timestamp's field in a message into its message header.
	// When this enabled, LocalMesgNum 0 is automatically used since the 5 lsb is used for the timestamp.
	headerOptionCompressedTimestamp headerOption = 1
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

	mesgDef    proto.MessageDefinition        // Temporary message definition to reduce alloc.
	bytesArray [proto.MaxBytesPerMessage]byte // General purpose array for encoding process.
}

type options struct {
	messageValidator         MessageValidator
	protocolVersion          proto.Version
	writeBufferSize          int
	endianness               byte
	headerOption             headerOption
	multipleLocalMessageType byte
}

func defaultOptions() options {
	return options{
		protocolVersion: proto.V1,
		endianness:      littleEndian,
		headerOption:    headerOptionNormal,
		writeBufferSize: defaultWriteBuffer,
	}
}

type Option interface{ apply(o *options) }

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

// WithProtocolVersion directs the Encoder to use specific Protocol Version (default: proto.V1).
// If the given protocolVersion is not supported, the Protocol Version will not be changed.
// Please validate using proto.Validate when putting user-defined Protocol Version to check
// whether it is supported or not. Or just use predefined Protocol Version constants such as
// proto.V1, proto.V2, etc, which the validity is ensured.
func WithProtocolVersion(protocolVersion proto.Version) Option {
	return fnApply(func(o *options) {
		if proto.Validate(byte(protocolVersion)) == nil {
			o.protocolVersion = protocolVersion
		}
	})
}

// WithMessageValidator directs the Encoder to use this message validator instead of the default one.
func WithMessageValidator(validator MessageValidator) Option {
	return fnApply(func(o *options) {
		if validator != nil {
			o.messageValidator = validator
		}
	})
}

// WithBigEndian directs the Encoder to encode values in Big-Endian bytes order (default: Little-Endian).
func WithBigEndian() Option {
	return fnApply(func(o *options) { o.endianness = bigEndian })
}

// WithCompressedTimestampHeader directs the Encoder to compress timestamp in header to reduce file size.
// Saves 7 bytes per message: 3 bytes for field definition and 4 bytes for the uint32 timestamp value.
func WithCompressedTimestampHeader() Option {
	return fnApply(func(o *options) { o.headerOption = headerOptionCompressedTimestamp })
}

// WithNormalHeader directs the Encoder to use NormalHeader for encoding the message using multiple local message types.
// By default, the Encoder uses local message type 0. This option allows users to specify values between 0-15 (while
// entering zero is equivalent to using the default option, nothing is changed). Using multiple local message types
// optimizes file size by avoiding the need to interleave different message definition.
//
// Note: To minimize the required RAM for decoding, it's recommended to use a minimal number of local message types.
// For instance, embedded devices may only support decoding data from local message type 0. Additionally,
// multiple local message types should be avoided in file types like settings, where messages of the same type
// can be grouped together.
func WithNormalHeader(multipleLocalMessageType byte) Option {
	if multipleLocalMessageType > proto.LocalMesgNumMask {
		multipleLocalMessageType = proto.LocalMesgNumMask
	}
	return fnApply(func(o *options) {
		o.headerOption = headerOptionNormal
		o.multipleLocalMessageType = multipleLocalMessageType
	})
}

// WithWriteBufferSize directs the Encoder to use this buffer size for writing to io.Writer instead of default 4096.
// When size <= 0, the Encoder will write directly to io.Writer without buffering.
func WithWriteBufferSize(size int) Option {
	return fnApply(func(o *options) { o.writeBufferSize = size })
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
//   - [io.WriterAt] or [io.WriteSeeker]: Encoder will update the FileHeader's DataSize and CRC after
//     the encoding process is completed since we can write at a specific byte position, making it more
//     ideal and efficient.
//   - [io.Writer]: Encoder needs to iterate through the messages once to calculate the FileHeader's DataSize
//     and CRC by writing to [io.Discard], then re-iterate through the messages again for the actual writing.
//
// Loading everything in memory and then writing it all later should preferably be avoided. While a FIT file
// is commonly small-sized, but by design, it can hold up to approximately 4GB. This is because the DataSize
// is of type uint32, and its maximum value is around that number. And also The FIT protocol allows for
// multiple FIT files to be chained together in a single FIT file. Each FIT file in the chain must be a properly
// formatted FIT file (FileHeader, Messages, CRC), making it more dynamic in size.
//
// Note: Encoder already implements efficient io.Writer buffering, so there's no need to wrap 'w' with a buffer;
// doing so will only reduce performance. If you don't want the Encoder to buffer the writing, please direct the
// Encoder to do so using WithWriteBufferSize(0).
func New(w io.Writer, opts ...Option) *Encoder {
	e := &Encoder{
		crc16:             crc16.New(nil),
		protocolValidator: new(proto.Validator),
		localMesgNumLRU:   new(lru),
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
		opts[i].apply(&e.options)
	}

	e.w = newWriteBuffer(w, e.options.writeBufferSize)

	if e.options.messageValidator == nil {
		e.options.messageValidator = NewMessageValidator()
	}

	e.reset()
	e.protocolValidator.SetProtocolVersion(e.options.protocolVersion)

	var lruSize byte = 1
	if e.options.headerOption == headerOptionNormal && e.options.multipleLocalMessageType > 0 {
		lruSize = e.options.multipleLocalMessageType + 1
	}
	e.localMesgNumLRU.ResetWithNewSize(lruSize)
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
// specifies a FileHeader's Size as 12 (legacy) or a custom FileHeader's ProfileVersion.
// In these cases, those two values will be encoded as-is, irrespective of the current SDK profile version.
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

	header.ProtocolVersion = byte(e.options.protocolVersion)
	header.DataType = proto.DataTypeFIT
	header.CRC = 0 // recalculated

	b, _ := header.MarshalAppend(e.bytesArray[:0])

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

// updateFileHeader updates the FileHeader if the DataSize is changed.
// The caller MUST ensure that e.w is either an io.WriterAt or an io.WriteSeeker.
func (e *Encoder) updateFileHeader(header *proto.FileHeader) error {
	if header.DataSize == e.dataSize {
		return nil
	}

	header.DataSize = e.dataSize

	b, _ := header.MarshalAppend(e.bytesArray[:0])

	if header.Size == 14 {
		_, _ = e.crc16.Write(b[:12]) // recalculate CRC Checksum since FileHeader is changed.
		header.CRC = e.crc16.Sum16() // update crc in FileHeader
		binary.LittleEndian.PutUint16(b[12:14], e.crc16.Sum16())
		e.crc16.Reset()
	}

	switch w := e.w.(type) {
	case io.WriterAt:
		_, err := w.WriteAt(b, e.lastFileHeaderPos)
		return err
	case io.WriteSeeker:
		_, err := w.Seek(e.lastFileHeaderPos, io.SeekStart)
		if err != nil {
			return err
		}
		_, err = w.Write(b)
		if err != nil {
			return err
		}
		_, err = w.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("encoder internal error") // should not reach here except we code wrong implementation.
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
func (e *Encoder) encodeMessage(mesg *proto.Message) error {
	mesg.Header = proto.MesgNormalHeaderMask
	mesg.Architecture = e.options.endianness

	if err := e.options.messageValidator.Validate(mesg); err != nil {
		return fmt.Errorf("message validation failed: %w", err)
	}

	if e.options.headerOption == headerOptionCompressedTimestamp {
		if e.w == io.Discard {
			// NOTE: Only for calculating data size (Early Check Strategy)
			var timestampField proto.Field
			if field := mesg.FieldByNum(proto.FieldNumTimestamp); field != nil {
				timestampField = *field
			}

			prevLen := len(mesg.Fields)
			e.compressTimestampIntoHeader(mesg)

			if prevLen != len(mesg.Fields) {
				defer func() { // Revert: put timestamp field back
					mesg.Fields = mesg.Fields[:prevLen]
					copy(mesg.Fields[1:], mesg.Fields[:len(mesg.Fields)])
					mesg.Fields[0] = timestampField
				}()
			}
		} else {
			e.compressTimestampIntoHeader(mesg)
		}
	}

	proto.CreateMessageDefinitionTo(&e.mesgDef, mesg)
	if err := e.protocolValidator.ValidateMessageDefinition(&e.mesgDef); err != nil {
		return err
	}

	b, _ := e.mesgDef.MarshalAppend(e.bytesArray[:0])
	localMesgNum, isNewMesgDef := e.localMesgNumLRU.Put(b) // This might alloc memory since we need to copy the item.
	if e.options.headerOption == headerOptionNormal {
		b[0] = (b[0] &^ proto.LocalMesgNumMask) | localMesgNum // Update the message definition header.
		mesg.Header = (mesg.Header &^ proto.LocalMesgNumMask) | localMesgNum
	}

	if isNewMesgDef {
		n, err := e.w.Write(b)
		e.n, e.dataSize = e.n+int64(n), e.dataSize+uint32(n)
		if err != nil {
			return fmt.Errorf("write message definition failed: %w", err)
		}
		_, _ = e.crc16.Write(b)
	}

	b, err := mesg.MarshalAppend(e.bytesArray[:0])
	if err != nil {
		return fmt.Errorf("marshal mesg failed: %w", err)
	}

	n, err := e.w.Write(b)
	e.n, e.dataSize = e.n+int64(n), e.dataSize+uint32(n)
	if err != nil {
		return fmt.Errorf("write message failed: %w", err)
	}
	_, _ = e.crc16.Write(b)

	return nil
}

func (e *Encoder) compressTimestampIntoHeader(mesg *proto.Message) {
	field := mesg.FieldByNum(proto.FieldNumTimestamp)
	if field == nil {
		return
	}

	if field.Value.Type() != proto.TypeUint32 {
		return // not supported
	}

	timestamp := field.Value.Uint32()
	if timestamp < uint32(typedef.DateTimeMin) {
		return
	}

	// The 5-bit time offset rolls over every 32 seconds, it is necessary that the difference
	// between timestamp and timestamp reference be measured less than 32 seconds apart.
	if (timestamp - e.timestampReference) > proto.CompressedTimeMask {
		e.timestampReference = timestamp
		return // Rollover event occurs, keep it as it is.
	}

	e.timestampReference = timestamp

	timeOffset := byte(timestamp & proto.CompressedTimeMask)
	mesg.Header |= proto.MesgCompressedHeaderMask | timeOffset
	mesg.RemoveFieldByNum(proto.FieldNumTimestamp)
}

func (e *Encoder) encodeCRC() error {
	b := e.bytesArray[:2]
	binary.LittleEndian.PutUint16(b, e.crc16.Sum16())

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
