// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	ErrNilWriter     = errors.New("nil writer")
	ErrEmptyMessages = errors.New("empty messages")
	ErrMissingFileId = errors.New("missing file_id mesg")

	ErrWriterAtOrWriteSeekerIsExpected = errors.New("io.WriterAt or io.WriteSeeker is expected")
)

const (
	RolloverEvent = 32 // The 5-bit time offset rolls over every 32 seconds. When an incoming time offset is less than previous time offset, rollover event has occurred.
)

// headerOption is header option.
type headerOption byte

const (
	littleEndian      = 0
	bigEndian         = 1
	defaultEndianness = littleEndian

	// headerOptionNormal is the default header option.
	// This option has two sub-option to select from:
	//   1. LocalMessageTypeZero [Default]
	// 		Optimized for all devices. It only use LocalMesgNum 0.
	//   2. MultipleLocalMessageType
	//      Using multiple local message types optimizes file size by avoiding the need to interleave different message typedef.
	//      The number of multiple local message type can be specified between 0-15.
	headerOptionNormal headerOption = 0

	// Optimized to reduce file size way further by compressing timestamp's field in a message into its message header.
	// When this enabled, LocalMesgNum 0 is automatically used since the 5 lsb is used for the timestamp.
	headerOptionCompressedTimestamp headerOption = 1
)

// Encoder is Fit file encoder. See New() for details.
type Encoder struct {
	w             io.Writer // A writer to write the bytes encoded fit.
	n             int64     // Total bytes written to w, will keep counting for every Encode invocation.
	lastHeaderPos int64     // The byte position of the last header.
	options       *options  // Encoder's options.

	messageValidator  MessageValidator // Validates all messages before encoding, ensuring the resulting Fit file is correct.
	protocolValidator *proto.Validator // Validates message's properties should match the targeted protocol version requirements.

	defaultFileHeader proto.FileHeader // Default header to encode when not specified.

	dataSize uint32      // Data size of messages in bytes for a single Fit file.
	crc16    hash.Hash16 // Calculate the CRC-16 checksum for ensuring header and message integrity.

	localMesgNumLRU *lru // LRU cache for writing local message definition

	// This timestamp reference is retrieved from the first message containing a valid timestamp field.
	// It is initialized only if the 'compressedTimestamp' option is applied and reset when decoding is completed.
	// This reference only serves as a flag and might be helpful for debugging purposes.
	timestampReference uint32

	mesgDef    *proto.MessageDefinition // Temporary message definition to reduce alloc.
	mesgDefBuf *bytes.Buffer            // Temporary message definition bytes buffer to reduce alloc (Only used for LRU).

	// A wrapper that act as a multi writer for writing message definitions and messages to w and crc16 simultaneously.
	// We don't use io.MultiWriter since we need to change w on every encode message.
	wrapWriterAndCrc16 *wrapWriterAndCrc16
}

type options struct {
	protocolVersion          proto.Version
	messageValidator         MessageValidator
	endianness               byte
	headerOption             headerOption
	multipleLocalMessageType byte
}

func defaultOptions() *options {
	return &options{
		endianness:       defaultEndianness,
		protocolVersion:  proto.V1,
		messageValidator: NewMessageValidator(),
		headerOption:     headerOptionNormal,
	}
}

type Option interface{ apply(o *options) }

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

// WithProtocolVersion directs the Encoder to use specific Protocol Version (default: proto.V1).
func WithProtocolVersion(protocolVersion proto.Version) Option {
	return fnApply(func(o *options) {
		o.protocolVersion = protocolVersion
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

// WithNormalHeader directs the Encoder to use NormalHeader for encoding the message using specified multiple local message typedef.
// By default, the Encoder uses local message type 0. This option allows users to specify values between 0-15 (while entering zero
// is equivalent to using the default option, nothing is changed). Using multiple local message types optimizes file size by avoiding
// the need to interleave different message typedef.
//
// Note: To minimize the required RAM for decoding, it's recommended to use a minimal number of local message types in a file.
// For instance, embedded devices may only support decoding data from local message type 0. Additionally, multiple local message types
// should be avoided in file types like settings, where messages of the same type can be grouped together.
func WithNormalHeader(multipleLocalMessageType byte) Option {
	if multipleLocalMessageType > proto.LocalMesgNumMask {
		multipleLocalMessageType = proto.LocalMesgNumMask
	}
	return fnApply(func(o *options) {
		o.multipleLocalMessageType = multipleLocalMessageType
		o.headerOption = headerOptionNormal
	})
}

// New returns a FIT File Encoder to encode FIT data to given w.
//
// # Encoding Strategy
//
// Since an invalid file header means an invalid Fit file, we need to ensure that the header is correct,
// specifically the header's DataSize (size of messages in bytes) and header's CRC checksum should be correct after everything is written.
//
// There are two strategies to achieve that and it depends on what kind of [io.Writer] is provided:
//   - [io.WriterAt] or [io.WriteSeeker]: Encoder can update the header's DataSize after
//     the encoding process is completed since we can write at a specific byte position, making it more ideal and efficient.
//   - [io.Writer]: Encoder needs to iterate through the messages once to calculate
//     the correct header's DataSize (calculating discarded bytes to [io.Discard]), then re-iterate through the messages again for the actual writing.
//
// Loading everything in memory and then writing it all later should preferably be avoided. While a Fit file is commonly small-sized, but by design, it can
// hold up to approximately 4GB. This is because the DataSize is of type uint32, and its maximum value is around that number.
// And also The FIT protocol allows for multiple FIT files to be chained together in a single FIT file. Each FIT file in the chain must be a properly
// formatted FIT file (header, data records, CRC), making it more dynamic in size.
//
// Note: We encourage wrapping w into a buffered writer such as bufferedwriter.New(w) that maintain io.WriteAt or io.WriteSeeker method (bufio.Writer does not).
// Encode process requires small bytes writing and having frequent write on non-buffered writer might impact performance,
// especially if it involves syscall such as writing a file. If you do wrap, don't forget to Flush() the buffered writer after encode is completed.
func New(w io.Writer, opts ...Option) *Encoder {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	var lruCapacity byte = 1
	if options.headerOption == headerOptionNormal && options.multipleLocalMessageType > 0 {
		lruCapacity = options.multipleLocalMessageType + 1
	}

	crc16 := crc16.New(crc16.MakeFitTable())
	return &Encoder{
		w:                 w,
		options:           options,
		crc16:             crc16,
		protocolValidator: proto.NewValidator(options.protocolVersion),
		messageValidator:  options.messageValidator,
		localMesgNumLRU:   newLRU(lruCapacity),
		defaultFileHeader: proto.FileHeader{
			Size:            proto.DefaultFileHeaderSize,
			ProtocolVersion: byte(options.protocolVersion),
			ProfileVersion:  profile.Version(),
			DataSize:        0, // calculated during encoding
			DataType:        proto.DataTypeFIT,
			CRC:             0, // calculated during encoding
		},
		mesgDef:            &proto.MessageDefinition{},
		mesgDefBuf:         bytes.NewBuffer(make([]byte, proto.MaxBytesPerMessageDefinition)),
		wrapWriterAndCrc16: &wrapWriterAndCrc16{writer: w, crc16: crc16},
	}
}

// Encode encodes fit into the dest writer. Encoder will do the following validations:
//  1. Calculating Header's CRC & DataSize and Fit's CRC: any mismatch calculation will be corrected and updated to the given fit structure.
//  2. Checking if fit.Messages are having all of its mesg definitions, return error if it's missing any mesg definition.
//
// Multiple Fit files can be chained together into a single Fit file by calling Encode for each fit data.
//
//	for _, fit := range fits {
//	   err := enc.Encode(fit)
//	}
//
// Encode chooses which strategy to use for encoding the data based on given writer and let the chosen strategy do the work.
func (e *Encoder) Encode(fit *proto.Fit) error {
	defer e.reset()

	// Encode Strategy
	switch e.w.(type) {
	case io.WriterAt, io.WriteSeeker:
		return e.encodeWithDirectUpdateStrategy(fit)
	case io.Writer:
		return e.encodeWithEarlyCheckStrategy(fit)
	}

	return ErrNilWriter
}

// encodeWithDirectUpdateStrategy encodes all data to file, after completing, it updates the actual size of the messages that being written to the proto.
func (e *Encoder) encodeWithDirectUpdateStrategy(fit *proto.Fit) error {
	if err := e.encodeHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessages(e.w, fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}
	if err := e.updateHeader(&fit.FileHeader); err != nil {
		return err
	}
	return nil
}

// encodeWithEarlyCheckStrategy does early calculation of the size of the messages that will be written and then do the encoding process.
func (e *Encoder) encodeWithEarlyCheckStrategy(fit *proto.Fit) error {
	if err := e.calculateDataSize(fit); err != nil {
		return err
	}
	if err := e.encodeHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessages(e.w, fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}

	return nil
}

// updateHeader updates the header since its content is changed, the header's CRC is recalculated before updating.
// The caller MUST ensure that e.w is either an io.WriterAt or an io.WriteSeeker.
func (e *Encoder) updateHeader(header *proto.FileHeader) error {
	if header.DataSize == e.dataSize {
		return nil
	}

	header.DataSize = e.dataSize

	b, _ := header.MarshalBinary()

	if header.Size >= 14 {
		_, _ = e.crc16.Write(b[:12]) // recalculate CRC Checksum since header is changed.
		header.CRC = e.crc16.Sum16() // update crc in header
		binary.LittleEndian.PutUint16(b[12:14], e.crc16.Sum16())
		e.crc16.Reset()
	}

	switch w := e.w.(type) {
	case io.WriterAt:
		_, err := w.WriteAt(b, e.lastHeaderPos)
		return err
	case io.WriteSeeker:
		_, err := w.Seek(e.lastHeaderPos, io.SeekStart)
		if err != nil {
			return err
		}
		_, err = w.Write(b)
		if err != nil {
			return err
		}

		// Restore the offset to the end since we have no control over the writer after the encoding process is completed.
		// In case the user accidentally writes to the writer, this ensures that the FIT File doesn't get corrupted.
		_, err = w.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("encoder internal error") // should not reach here except we code wrong implementation.
}

// calculateDataSize calculates total data size of the messages by counting bytes written to io.Discard.
func (e *Encoder) calculateDataSize(fit *proto.Fit) error {
	n := e.n

	if err := e.encodeMessages(io.Discard, fit.Messages); err != nil {
		return fmt.Errorf("calculate data size: %w", err)
	}

	if fit.FileHeader.Size == 0 {
		fit.FileHeader = e.defaultFileHeader
	}
	fit.FileHeader.DataSize = e.dataSize // update Header's DataSize of the actual messages size
	e.dataSize = 0
	e.n = n
	e.crc16.Reset()
	e.localMesgNumLRU.Reset()

	return nil
}

func (e *Encoder) encodeHeader(header *proto.FileHeader) error {
	e.lastHeaderPos = e.n

	if header.Size < 12 {
		*header = e.defaultFileHeader
	}
	header.ProtocolVersion = byte(e.options.protocolVersion)

	b, _ := header.MarshalBinary()

	if header.Size < 14 {
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

func (e *Encoder) encodeMessages(w io.Writer, messages []proto.Message) error {
	if len(messages) == 0 {
		return ErrEmptyMessages
	}

	if messages[0].Num != mesgnum.FileId {
		return ErrMissingFileId
	}

	for i := range messages {
		mesg := &messages[i]
		if err := e.encodeMessage(w, mesg); err != nil {
			return fmt.Errorf("encode failed: at byte pos: %d, message index: %d, num: %d (%s): %w",
				e.n, i, mesg.Num, mesg.Num.String(), err)
		}
	}

	return nil
}

// encodeMessage marshals and encodes message definition and its message into w.
func (e *Encoder) encodeMessage(w io.Writer, mesg *proto.Message) error {
	mesg.Header = proto.MesgNormalHeaderMask
	mesg.Architecture = e.options.endianness

	if err := e.messageValidator.Validate(mesg); err != nil {
		return fmt.Errorf("message validation failed: %w", err)
	}

	var isNewMesgDef bool
	e.mesgDefBuf.Reset()

	// Writing strategy based on the selected header option:
	switch e.options.headerOption {
	case headerOptionNormal:
		proto.CreateMessageDefinitionTo(e.mesgDef, mesg)
		if err := e.protocolValidator.ValidateMessageDefinition(e.mesgDef); err != nil {
			return err
		}

		var localMesgNum byte
		_, _ = e.mesgDef.WriteTo(e.mesgDefBuf)
		localMesgNum, isNewMesgDef = e.localMesgNumLRU.Put(e.mesgDefBuf.Bytes())
		e.mesgDef.Header = (e.mesgDef.Header &^ proto.LocalMesgNumMask) | localMesgNum // update the message definition header.

		mesg.Header = (mesg.Header &^ proto.LocalMesgNumMask) | localMesgNum
	case headerOptionCompressedTimestamp:
		e.compressTimestampIntoHeader(mesg)

		proto.CreateMessageDefinitionTo(e.mesgDef, mesg)
		if err := e.protocolValidator.ValidateMessageDefinition(e.mesgDef); err != nil {
			return err
		}

		_, _ = e.mesgDef.WriteTo(e.mesgDefBuf)
		_, isNewMesgDef = e.localMesgNumLRU.Put(e.mesgDefBuf.Bytes())
	}

	if isNewMesgDef {
		e.wrapWriterAndCrc16.writer = w // Change writer
		n, err := e.mesgDef.WriteTo(e.wrapWriterAndCrc16)
		e.dataSize += uint32(n)
		e.n += int64(n)
		if err != nil {
			return fmt.Errorf("write message failed: %w", err)
		}
	}

	e.wrapWriterAndCrc16.writer = w // Change writer
	n, err := mesg.WriteTo(e.wrapWriterAndCrc16)
	e.dataSize += uint32(n)
	e.n += int64(n)
	if err != nil {
		return fmt.Errorf("write message failed: %w", err)
	}

	return nil
}

func (e *Encoder) compressTimestampIntoHeader(mesg *proto.Message) {
	field := mesg.FieldByNum(proto.FieldNumTimestamp)
	if field == nil {
		return
	}

	var timestamp uint32
	switch val := field.Value.(type) {
	case uint32:
		timestamp = val
	case typedef.DateTime:
		timestamp = uint32(val)
	default:
		return // not supported
	}

	if timestamp < uint32(typedef.DateTimeMin) {
		return
	}

	if e.timestampReference == 0 || (timestamp-e.timestampReference) > RolloverEvent {
		// There should be at least one valid timestamp reference in a field prior to the use of
		// the compressed timestamp header. If the gap beetween new timestamp and last reference is more
		// than allowed Rollover Event (32 seconds), the new timestamp become new reference.
		e.timestampReference = timestamp
		return
	}

	e.timestampReference = timestamp

	timeOffset := byte(timestamp & proto.CompressedTimeMask)
	mesg.Header |= proto.MesgCompressedHeaderMask | timeOffset
	mesg.RemoveFieldByNum(proto.FieldNumTimestamp)
}

type wrapWriterAndCrc16 struct {
	writer io.Writer
	crc16  hash.Hash16
}

func (w *wrapWriterAndCrc16) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	if err != nil {
		return
	}
	_, _ = w.crc16.Write(p)
	return
}

func (e *Encoder) encodeCRC() error {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, e.crc16.Sum16())

	n, err := e.w.Write(b)
	if err != nil {
		return fmt.Errorf("could not write crc: %w", err)
	}

	e.n += int64(n)

	e.crc16.Reset()

	return nil
}

// reset resets the encoder's data that is being used for encoding, allowing the encoder to be reused for writing the subsequent chained FIT file.
func (e *Encoder) reset() {
	e.dataSize = 0
	e.crc16.Reset()
	e.localMesgNumLRU.Reset()
	e.timestampReference = 0
}

// EncodeWithContext is similiar to Encode but with respect to context propagation.
func (e *Encoder) EncodeWithContext(ctx context.Context, fit *proto.Fit) (err error) {
	defer e.reset()

	// Encode Strategy
	switch e.w.(type) {
	case io.WriterAt, io.WriteSeeker:
		return e.encodeWithDirectUpdateStrategyWithContext(ctx, fit)
	case io.Writer:
		return e.encodeWithEarlyCheckStrategyWithContext(ctx, fit)
	}

	return ErrNilWriter
}

func (e *Encoder) encodeWithDirectUpdateStrategyWithContext(ctx context.Context, fit *proto.Fit) error {
	if err := e.encodeHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessagesWithContext(ctx, e.w, fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}
	if err := e.updateHeader(&fit.FileHeader); err != nil {
		return err
	}
	return nil
}

func (e *Encoder) encodeWithEarlyCheckStrategyWithContext(ctx context.Context, fit *proto.Fit) error {
	if err := e.calculateDataSize(fit); err != nil {
		return err
	}
	if err := e.encodeHeader(&fit.FileHeader); err != nil {
		return err
	}
	if err := e.encodeMessagesWithContext(ctx, e.w, fit.Messages); err != nil {
		return err
	}
	fit.CRC = e.crc16.Sum16()
	if err := e.encodeCRC(); err != nil {
		return err
	}
	return nil
}

func (e *Encoder) encodeMessagesWithContext(ctx context.Context, w io.Writer, messages []proto.Message) error {
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
		if err := e.encodeMessage(w, mesg); err != nil {
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
