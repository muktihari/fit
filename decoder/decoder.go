// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	// ErrNotFITFile will be returned if the first byte of every FIT sequence does not match
	// with FIT FileHeader's Size specification (either 12 or 14), byte 8-12 are not ".FIT",
	// or byte 4-8 are all zero (FileHeader's DataSize == 0).
	ErrNotFITFile = errorString("not a FIT file")

	// ErrCRCChecksumMismatch will be returned if the CRC checksum does not match with
	// the CRC in the file, whether during FileHeader or messages integrity checks.
	ErrCRCChecksumMismatch = errorString("crc checksum mismatch")

	// ErrMesgDefMissing will be returned if message definition for the incoming message data is missing.
	ErrMesgDefMissing = errorString("message definition missing") // NOTE: Kept exported since it's used by RawDecoder

	errInvalidBaseType = errorString("invalid basetype")
	errMissingFileId   = errorString("missing file_id")
)

// Decoder is FIT file decoder. See New() for details.
type Decoder struct {
	readBuffer  *readBuffer // read from io.Reader with buffer without extra copying.
	n           int64       // n is a read bytes counter, always moving forward, do not reset (except on full reset).
	factory     Factory
	accumulator *Accumulator
	crc16       hash.Hash16
	err         error // Any error occurs during process.

	fieldsArray          [255]proto.Field
	developerFieldsArray [255]proto.DeveloperField

	options options

	once           sync.Once // It is used to invoke decodeFileHeader exactly once. Must be reassigned on init/reset.
	cur            uint32    // The current byte position relative to bytes of the messages, reset on next chained FIT file.
	timestamp      uint32    // Active timestamp
	lastTimeOffset byte      // Last time offset

	// FIT File Representation
	fileHeader proto.FileHeader
	messages   []proto.Message
	crc        uint16

	// FileId Message is a special message that must present in a FIT file.
	fileId *mesgdef.FileId

	// Message Definition Lookup
	localMessageDefinitions      [proto.LocalMesgNumMask + 1]*proto.MessageDefinition // message definition for upcoming message data
	localMessageDefinitionsArray [proto.LocalMesgNumMask + 1]proto.MessageDefinition  // PERF: backing array for message definition

	// Developer Data Lookup
	developerDataIndexes []uint8
	fieldDescriptions    []*mesgdef.FieldDescription
}

// Factory defines a contract that any Factory containing these method can be used by the Decoder.
type Factory interface {
	// CreateField creates new field based on defined messages in the factory.
	// If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

type options struct {
	factory               Factory
	logWriter             io.Writer
	mesgListeners         []MesgListener    // Each listener will received every decoded message.
	mesgDefListeners      []MesgDefListener // Each listener will received every decoded message definition.
	readBufferSize        int
	shouldChecksum        bool
	broadcastOnly         bool
	shouldExpandComponent bool
	broadcastMesgCopy     bool
}

func defaultOptions() options {
	return options{
		factory:               factory.StandardFactory(),
		logWriter:             nil,
		readBufferSize:        defaultReadBufferSize,
		shouldChecksum:        true,
		broadcastOnly:         false,
		shouldExpandComponent: true,
		broadcastMesgCopy:     false,
	}
}

// Option is Decoder's option.
type Option func(o *options)

// WithFactory sets custom factory.
func WithFactory(factory Factory) Option {
	return func(o *options) {
		if factory != nil {
			o.factory = factory
		}
	}
}

// WithMesgListener adds listeners to the listener pool, where each listener will receive
// every message as soon as it is decoded. The listeners will be appended not replaced.
// If users need to reset use Reset().
func WithMesgListener(listeners ...MesgListener) Option {
	return func(o *options) {
		o.mesgListeners = append(o.mesgListeners, listeners...)
	}
}

// WithMesgDefListener adds listeners to the listener pool, where each listener will receive
// every message definition as soon as it is decoded. The listeners will be appended not replaced.
// If users need to reset use Reset().
func WithMesgDefListener(listeners ...MesgDefListener) Option {
	return func(o *options) {
		o.mesgDefListeners = append(o.mesgDefListeners, listeners...)
	}
}

// WithBroadcastOnly directs the Decoder to only broadcast the messages without retaining them, reducing memory usage when
// it's not going to be used anyway. This option is intended to be used with WithMesgListener and
// When this option is specified, the Decode will return a FIT with empty messages.
func WithBroadcastOnly() Option {
	return func(o *options) { o.broadcastOnly = true }
}

// WithBroadcastMesgCopy directs the Decoder to copy the mesg before passing it to listeners
// (it was the default behavior on version <= v0.14.0).
func WithBroadcastMesgCopy() Option {
	return func(o *options) { o.broadcastMesgCopy = true }
}

// WithIgnoreChecksum directs the Decoder to not checking data integrity (CRC Checksum).
func WithIgnoreChecksum() Option {
	return func(o *options) { o.shouldChecksum = false }
}

// WithNoComponentExpansion directs the Decoder to not expand the components.
func WithNoComponentExpansion() Option {
	return func(o *options) { o.shouldExpandComponent = false }
}

// WithLogWriter specifies where the log messages will be written to. By default, the Decoder do not write any log if
// log writer is not specified. The Decoder will only write log messages when it encountered a bad encoded FIT file such as:
//   - Field Definition's Size (or Developer Field Definition's Size) is zero.
//   - Field Definition's Size (or Developer Field Definition's Size) is less than basetype's Size.
//     e.g. Size 1 bytes but having basetype uint32 (4 bytes).
//   - Field Definition's Size is more than basetype's Size but field.Array is false.
//   - Encountering a Developer Field without prior DeveloperDataId or FieldDescription Message.
func WithLogWriter(w io.Writer) Option {
	return func(o *options) { o.logWriter = w }
}

// WithReadBufferSize directs the Decoder to use this buffer size for reading from io.Reader instead of default 4096.
func WithReadBufferSize(size int) Option {
	return func(o *options) { o.readBufferSize = size }
}

// New returns a FIT File Decoder to decode given r.
//
// The FIT protocol allows for multiple FIT files to be chained together in a single FIT file.
// Each FIT file in the chain must be a properly formatted FIT file (header, data records, CRC).
//
// To decode a chained FIT file containing multiple FIT data, invoke Decode() or DecodeWithContext()
// method multiple times. For convenience, we can wrap it with the Next() method as follows (optional):
//
//	for dec.Next() {
//	   fit, err := dec.Decode()
//	}
//
// Note: Decoder already implements efficient io.Reader buffering, so there's no need to wrap 'r' using *bufio.Reader;
// doing so will only reduce performance.
func New(r io.Reader, opts ...Option) *Decoder {
	d := &Decoder{
		readBuffer:  new(readBuffer),
		accumulator: NewAccumulator(),
		crc16:       crc16.New(nil),
	}
	d.Reset(r, opts...)
	return d
}

// Reset resets the Decoder to read its input from r, clear any error and
// reset previous options to default options so any options needs to be inputed again.
// It is similar to New() but it retains the underlying storage for use by
// future decode to reduce memory allocs.
func (d *Decoder) Reset(r io.Reader, opts ...Option) {
	d.reset()
	d.n = 0 // Must reset bytes counter since it's a full reset.

	// Reuse listeners' slices
	for i := range d.options.mesgListeners {
		d.options.mesgListeners[i] = nil // avoid memory leaks
	}
	mesgListeners := d.options.mesgListeners[:0]
	for i := range d.options.mesgDefListeners {
		d.options.mesgDefListeners[i] = nil // avoid memory leaks
	}
	mesgDefListeners := d.options.mesgDefListeners[:0]

	d.options = defaultOptions()
	d.options.mesgListeners = mesgListeners
	d.options.mesgDefListeners = mesgDefListeners
	for i := range opts {
		opts[i](&d.options)
	}

	d.readBuffer.Reset(r, d.options.readBufferSize)
	d.factory = d.options.factory
}

func (d *Decoder) reset() {
	d.accumulator.Reset()
	d.crc16.Reset()
	d.once = sync.Once{}
	d.cur = 0
	d.timestamp = 0
	d.lastTimeOffset = 0
	d.err = nil
	d.fileHeader = proto.FileHeader{}
	d.messages = nil
	d.crc = 0
	d.fileId = nil
	d.developerDataIndexes = d.developerDataIndexes[:0]
	d.fieldDescriptions = d.fieldDescriptions[:0]
}

// releaseTemporaryObjects releases objects that being created on a single decode process
// by stops referencing those objects so it can be released on next GC cycle.
func (d *Decoder) releaseTemporaryObjects() {
	d.localMessageDefinitions = [proto.LocalMesgNumMask + 1]*proto.MessageDefinition{}
	d.fieldsArray = [255]proto.Field{}
	d.developerFieldsArray = [255]proto.DeveloperField{}
	d.messages = nil
	for i := range d.fieldDescriptions {
		d.fieldDescriptions[i] = nil
	}
}

// CheckIntegrity checks all FIT sequences of given reader are valid determined by these following checks:
//  1. Has valid FileHeader's size and bytes 8–11 of the FileHeader is “.FIT”
//  2. FileHeader's DataSize > 0
//  3. CRC checksum of messages should match with File's CRC value.
//
// It returns the number of sequences completed and any error encountered. The number of sequences completed can help recovering
// valid FIT sequences in a chained FIT that contains invalid or corrupted data.
//
// After invoking this method, the underlying reader should be reset afterward as the reader has been fully read.
// If the underlying reader implements io.Seeker, we can do reader.Seek(0, io.SeekStart).
func (d *Decoder) CheckIntegrity() (seq int, err error) {
	if d.err != nil {
		return 0, d.err
	}

	shouldChecksum := d.options.shouldChecksum
	d.options.shouldChecksum = true // Must checksum

	for {
		// Check File Header Integrity
		pos := d.n
		if err = d.decodeFileHeaderOnce(); err != nil {
			if pos != 0 && pos == d.n && err == io.EOF {
				// When EOF error occurs exactly after a sequence has been completed,
				// make the error as nil, it means we have reached the desirable EOF.
				err = nil
			}
			break
		}
		// Read bytes acquired by messages to calculate crc checksum of its contents
		if err = d.discardMessages(); err != nil {
			break
		}
		if err = d.decodeCRC(); err != nil {
			break
		}
		d.once = sync.Once{}
		d.cur = 0
		seq++
	}

	if err != nil { // When there is an error, wrap it with informative message before return.
		err = fmt.Errorf("byte pos: %d: %w", d.n, err)
	}

	// Reset used variables so that the decoder can be reused by the same reader.
	d.reset()
	d.n = 0 // Must reset bytes counter
	d.options.shouldChecksum = shouldChecksum

	return seq, err
}

// discardMessages efficiently discards bytes used by messages.
func (d *Decoder) discardMessages() (err error) {
	for d.cur < d.fileHeader.DataSize {
		size := int(d.fileHeader.DataSize - d.cur)
		if size > reservedbuf {
			size = reservedbuf
		}
		if _, err = d.readN(size); err != nil { // Discard bytes
			return err
		}
	}
	return nil
}

// PeekFileHeader decodes only up to FileHeader (first 12-14 bytes) without decoding the whole reader.
//
// If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.
func (d *Decoder) PeekFileHeader() (*proto.FileHeader, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	return &d.fileHeader, nil
}

// PeekFileId decodes only up to FileId message without decoding the whole reader.
// FileId message should be the first message of any FIT file, otherwise return an error.
//
// If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.
func (d *Decoder) PeekFileId() (*mesgdef.FileId, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	for len(d.messages) == 0 {
		if d.err = d.decodeMessage(); d.err != nil {
			return nil, d.err
		}
	}
	if d.fileId == nil {
		mesg := &d.messages[0]
		return nil, fmt.Errorf("expect file_id as first mesg, got: %s(%d): %w",
			mesg.Num, mesg.Num, errMissingFileId)
	}
	return d.fileId, nil
}

// Next checks whether next bytes are still a valid FIT File sequence. Return false when invalid or reach EOF.
func (d *Decoder) Next() bool {
	if d.err != nil {
		return false
	}
	if d.n == 0 {
		return true
	}
	return d.decodeFileHeaderOnce() == nil
}

// Decode method decodes `r` into FIT data. One invocation will produce one valid FIT data or
// an error if it occurs. To decode a chained FIT file containing multiple FIT data, invoke this
// method multiple times. For convenience, we can wrap it with the Next() method as follows (optional):
//
//	for dec.Next() {
//	     fit, err := dec.Decode()
//	     if err != nil {
//	         return err
//	     }
//	}
func (d *Decoder) Decode() (*proto.FIT, error) {
	if d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	defer d.releaseTemporaryObjects()
	if d.err = d.decodeMessages(); d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeCRC(); d.err != nil {
		return nil, d.err
	}
	fit := &proto.FIT{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}
	d.reset()
	return fit, nil
}

// Discard discards a single FIT file sequence and returns any error encountered. This method directs the Decoder to
// point to the byte sequence of the next valid FIT file sequence, discarding the current FIT file sequence.
//
// Example: - A chained FIT file consist of Activity, Course, Workout and Settings. And we only want to decode Course.
//
//	for dec.Next() {
//		fileId, err := dec.PeekFileId()
//		if err != nil {
//			return err
//		}
//		if fileId.Type != typedef.FileCourse {
//		    if err := dec.Discard(); err != nil {
//		    	return err
//		    }
//		    continue
//		}
//		fit, err := dec.Decode()
//		if err != nil {
//			return err
//		}
//	 }
func (d *Decoder) Discard() error {
	if d.err != nil {
		return d.err
	}

	optionsShouldChecksum := d.options.shouldChecksum
	d.options.shouldChecksum = false
	defer func() { d.options.shouldChecksum = optionsShouldChecksum }()

	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return d.err
	}
	if d.err = d.discardMessages(); d.err != nil {
		return d.err
	}
	if _, d.err = d.readN(2); d.err != nil { // Discard File CRC
		return d.err
	}
	d.reset()
	return d.err
}

// decodeFileHeaderOnce invokes decodeFileHeader exactly once.
func (d *Decoder) decodeFileHeaderOnce() error {
	d.once.Do(func() { d.err = d.decodeFileHeader() })
	return d.err
}

// decodeFileHeader is only invoked through decodeFileHeaderOnce.
func (d *Decoder) decodeFileHeader() error {
	b, err := d.readBuffer.ReadN(1)
	if err != nil {
		return err
	}
	d.n += 1
	size := b[0]

	if size != 12 && size != 14 { // current spec is either 12 or 14
		return fmt.Errorf("file header size [%d] is invalid: %w", size, ErrNotFITFile)
	}
	_, _ = d.crc16.Write(b)

	rem := int(size - 1)
	b, err = d.readBuffer.ReadN(rem)
	if err != nil {
		return err
	}
	d.n += int64(rem)

	// PERF: Neither string(b[7:11]) nor assigning proto.DataTypeFIT constant to a variable escape to the heap.
	if string(b[7:11]) != proto.DataTypeFIT {
		return ErrNotFITFile
	}

	d.fileHeader = proto.FileHeader{
		Size:            size,
		ProtocolVersion: proto.Version(b[0]),
		ProfileVersion:  binary.LittleEndian.Uint16(b[1:3]),
		DataSize:        binary.LittleEndian.Uint32(b[3:7]),
		DataType:        proto.DataTypeFIT,
	}

	if err := proto.Validate(d.fileHeader.ProtocolVersion); err != nil {
		return err
	}

	if d.fileHeader.DataSize == 0 {
		return fmt.Errorf("invalid data size: %w", ErrNotFITFile)
	}

	if size == 14 {
		d.fileHeader.CRC = binary.LittleEndian.Uint16(b[11:13])
	}

	if d.fileHeader.CRC == 0x0000 || !d.options.shouldChecksum { // do not need to check file header's crc integrity.
		d.crc16.Reset()
		return nil
	}

	_, _ = d.crc16.Write(b[:len(b)-2])
	if d.crc16.Sum16() != d.fileHeader.CRC { // check file header integrity
		return fmt.Errorf("expected file header's crc: %d, got: %d: %w", d.fileHeader.CRC, d.crc16.Sum16(), ErrCRCChecksumMismatch)
	}

	d.crc16.Reset() // this hash will be re-used for calculating data integrity.

	return nil
}

func (d *Decoder) decodeMessages() (err error) {
	for d.cur < d.fileHeader.DataSize {
		if err = d.decodeMessage(); err != nil {
			return fmt.Errorf("decodeMessage [byte pos: %d]: %w", d.n, err)
		}
	}
	return nil
}

func (d *Decoder) decodeMessage() error {
	b, err := d.readN(1)
	if err != nil {
		return err
	}
	header := b[0]

	// NOTE: Compressed Timestamp Header Bit 5-6 is the local message type.
	// Bit 6 overlap with MesgDefinitionMask; It's a message definition only if Bit 7 is zero.
	if (header & (proto.MesgCompressedHeaderMask | proto.MesgDefinitionMask)) == proto.MesgDefinitionMask {
		return d.decodeMessageDefinition(header)
	}
	return d.decodeMessageData(header)
}

func (d *Decoder) decodeMessageDefinition(header byte) error {
	b, err := d.readN(5)
	if err != nil {
		return err
	}

	localMesgNum := header & proto.LocalMesgNumMask
	mesgDef := d.localMessageDefinitions[localMesgNum]
	if mesgDef == nil {
		// PERF: Use backing array to avoid object creation. On init, allocate slices
		// with max cap for more deterministic performance by avoiding re-allocation.
		mesgDef = &d.localMessageDefinitionsArray[localMesgNum]
		if mesgDef.FieldDefinitions == nil && mesgDef.DeveloperFieldDefinitions == nil {
			mesgDef.FieldDefinitions = make([]proto.FieldDefinition, 0, 255)
			mesgDef.DeveloperFieldDefinitions = make([]proto.DeveloperFieldDefinition, 0, 255)
		}
	}

	mesgDef.Header = header
	mesgDef.Reserved = b[0]
	mesgDef.Architecture = b[1]
	if mesgDef.Architecture == proto.LittleEndian {
		mesgDef.MesgNum = typedef.MesgNum(binary.LittleEndian.Uint16(b[2:4]))
	} else {
		mesgDef.MesgNum = typedef.MesgNum(binary.BigEndian.Uint16(b[2:4]))
	}

	n := int(b[4])
	b, err = d.readN(n * 3) // 3 byte per field
	if err != nil {
		return err
	}

	mesgDef.FieldDefinitions = mesgDef.FieldDefinitions[:0]
	var baseType basetype.BaseType
	for ; len(b) >= 3; b = b[3:] {
		baseType = basetype.BaseType(b[2])
		if !baseType.Valid() {
			return fmt.Errorf("message definition number: %s(%d): fields[%d].BaseType: %s: %w",
				mesgDef.MesgNum, mesgDef.MesgNum, len(mesgDef.FieldDefinitions), baseType, errInvalidBaseType)
		}
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions,
			proto.FieldDefinition{
				Num:      b[0],
				Size:     b[1],
				BaseType: baseType,
			})
	}

	mesgDef.DeveloperFieldDefinitions = mesgDef.DeveloperFieldDefinitions[:0]
	if (header & proto.DevDataMask) == proto.DevDataMask {
		b, err = d.readN(1)
		if err != nil {
			return err
		}

		n = int(b[0])
		b, err = d.readN(n * 3) // 3 byte per field
		if err != nil {
			return err
		}

		for ; len(b) >= 3; b = b[3:] {
			mesgDef.DeveloperFieldDefinitions = append(mesgDef.DeveloperFieldDefinitions,
				proto.DeveloperFieldDefinition{
					Num:                b[0],
					Size:               b[1],
					DeveloperDataIndex: b[2],
				})
		}
	}

	d.localMessageDefinitions[localMesgNum] = mesgDef

	if len(d.options.mesgDefListeners) > 0 {
		// Clone since we don't have control of the object lifecycle outside Decoder.
		mesgDef := *mesgDef
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions[:0:0], mesgDef.FieldDefinitions...)
		mesgDef.DeveloperFieldDefinitions = append(mesgDef.DeveloperFieldDefinitions[:0:0], mesgDef.DeveloperFieldDefinitions...)
		for i := range d.options.mesgDefListeners {
			d.options.mesgDefListeners[i].OnMesgDef(mesgDef) // blocking or non-blocking depends on listeners' implementation.
		}
	}

	return nil
}

func (d *Decoder) decodeMessageData(header byte) (err error) {
	localMesgNum := header
	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask {
		localMesgNum = (header & proto.CompressedLocalMesgNumMask) >> proto.CompressedBitShift
	}
	mesgDef := d.localMessageDefinitions[localMesgNum&proto.LocalMesgNumMask] // bounds check eliminated due to the mask
	if mesgDef == nil {
		return ErrMesgDefMissing
	}

	mesg := proto.Message{Num: mesgDef.MesgNum}
	mesg.Header = header
	mesg.Fields = d.fieldsArray[:0]

	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask { // Compressed Timestamp Message Data
		timeOffset := header & proto.CompressedTimeMask
		d.timestamp += uint32((timeOffset - d.lastTimeOffset) & proto.CompressedTimeMask)
		d.lastTimeOffset = timeOffset

		timestampField := d.factory.CreateField(mesgDef.MesgNum, proto.FieldNumTimestamp)
		if timestampField.Name == factory.NameUnknown {
			timestampField.BaseType = basetype.Uint32
			timestampField.Type = profile.DateTime
		}
		timestampField.Value = proto.Uint32(d.timestamp)

		mesg.Fields = append(mesg.Fields, timestampField) // add timestamp field
	}

	if err = d.decodeFields(mesgDef, &mesg); err != nil {
		return err
	}

	// FileId Message
	if d.fileId == nil && mesg.Num == mesgnum.FileId {
		d.fileId = mesgdef.NewFileId(&mesg)
	}

	// Prerequisites for decoding developer fields
	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		// These messages must occur before any related field description messages are written to the proto.
		d.developerDataIndexes = append(d.developerDataIndexes,
			mesg.FieldValueByNum(fieldnum.DeveloperDataIdDeveloperDataIndex).Uint8())
	case mesgnum.FieldDescription:
		// These messages must occur in the file before any related developer data is written to the proto.
		d.fieldDescriptions = append(d.fieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	}

	if len(mesgDef.DeveloperFieldDefinitions) != 0 {
		mesg.DeveloperFields = d.developerFieldsArray[:0]
		if err = d.decodeDeveloperFields(mesgDef, &mesg); err != nil {
			return err
		}
	}

	if !d.options.broadcastOnly || d.options.broadcastMesgCopy {
		mesg.Fields = append(mesg.Fields[:0:0], mesg.Fields...)
		mesg.DeveloperFields = append(mesg.DeveloperFields[:0:0], mesg.DeveloperFields...)
	}

	if !d.options.broadcastOnly {
		d.messages = append(d.messages, mesg)
	}

	for i := range d.options.mesgListeners {
		d.options.mesgListeners[i].OnMesg(mesg) // blocking or non-blocking depends on listeners' implementation.
	}

	return nil
}

func (d *Decoder) decodeFields(mesgDef *proto.MessageDefinition, mesg *proto.Message) error {
	for i := range mesgDef.FieldDefinitions {
		fieldDef := &mesgDef.FieldDefinitions[i]

		// We enforce field.Array for string type to match the value defined in factory for all non-unknown fields.
		var overrideStringArray bool
		field := d.factory.CreateField(mesgDef.MesgNum, fieldDef.Num)
		if field.Name == factory.NameUnknown {
			// Assign fieldDef's type for unknown field so later we can encode it as per its original value.
			field.BaseType = fieldDef.BaseType
			field.Type = profile.ProfileTypeFromBaseType(field.BaseType)
			// Check if the size corresponds to an array.
			field.Array = fieldDef.Size > field.BaseType.Size() && fieldDef.Size%field.BaseType.Size() == 0
			// Fallback to FIT Protocol's string rule: decoder will determine it by counting the utf8 null-terminated string.
			overrideStringArray = field.BaseType == basetype.String
		}

		var (
			baseType   = field.BaseType
			profilType = field.Type
			array      = field.Array
		)

		// Gracefully handle poorly encoded FIT file.
		if fieldDef.Size == 0 {
			d.logField(mesg, fieldDef, "Size is zero. Skip")
			continue
		} else if fieldDef.Size < baseType.Size() {
			baseType = basetype.Byte
			profilType = profile.Byte
			array = fieldDef.Size > baseType.Size() && fieldDef.Size&baseType.Size() == 0
			d.logField(mesg, fieldDef, "Size is less than expected. Fallback: decode as byte(s) and convert the value")
		} else if fieldDef.Size > baseType.Size() && !field.Array && baseType != basetype.String {
			d.logField(mesg, fieldDef, "field.Array is false. Fallback: retrieve first array's value only")
		}

		val, err := d.readValue(fieldDef.Size, mesgDef.Architecture, baseType, profilType, array, overrideStringArray)
		if err != nil {
			return err
		}

		if baseType != field.BaseType { // Convert value
			bitVal, _ := bitsFromValue(val)
			val = valueFromBits(bitVal, field.BaseType)
		}

		if field.Num == proto.FieldNumTimestamp && val.Type() == proto.TypeUint32 {
			timestamp := val.Uint32()
			d.timestamp = timestamp
			d.lastTimeOffset = byte(timestamp & proto.CompressedTimeMask)
		}

		field.Value = val

		if d.options.shouldExpandComponent && field.Accumulate {
			if val, ok := bitsFromValue(field.Value); ok {
				d.accumulator.Collect(mesg.Num, field.Num, val) // Collect the field values to be used in component expansion.
			}
		}

		mesg.Fields = append(mesg.Fields, field)
	}

	if !d.options.shouldExpandComponent {
		return nil
	}

	// Now that all fields has been decoded, we need to expand all components and accumulate the accumulable values.
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if subField := field.SubFieldSubtitution(mesg); subField != nil {
			// Expand sub-field components as the main field components
			d.expandComponents(mesg, field, subField.Components)
			continue
		}
		// No sub-field can interpret as main field, expand main field components
		d.expandComponents(mesg, field, field.Components)
	}

	return nil
}

func (d *Decoder) expandComponents(mesg *proto.Message, containingField *proto.Field, components []proto.Component) {
	if !containingField.Value.Valid(containingField.BaseType) {
		return
	}

	if len(components) == 0 {
		return
	}

	bitVal, ok := bitsFromValue(containingField.Value)
	if !ok {
		return
	}

	// PERF: Reuse a single variable 'componentField' instead of declaring it inside the loop to prevent
	// the Compiler's escape analysis from moving it to the heap.
	var componentField proto.Field
	for i := range components {
		component := &components[i]

		componentField = d.factory.CreateField(mesg.Num, component.FieldNum)
		componentField.IsExpandedField = true

		if component.Accumulate {
			bitVal = d.accumulator.Accumulate(mesg.Num, component.FieldNum, bitVal, component.Bits)
		}

		var val = bitVal
		if len(components) > 1 {
			if bitVal == 0 {
				break // no more bits to shift
			}
			var mask uint32 = (1 << component.Bits) - 1 // e.g. (1 << 8) - 1     = 255
			val = val & mask                            // e.g. 0x27010E08 & 255 = 0x08
			bitVal = bitVal >> component.Bits           // e.g. 0x27010E08 >> 8  = 0x27010E
		}

		componentScaled := scaleoffset.Apply(val, component.Scale, component.Offset)
		val = uint32(scaleoffset.Discard(componentScaled, componentField.Scale, componentField.Offset))
		componentField.Value = valueFromBits(val, componentField.BaseType)

		mesg.Fields = append(mesg.Fields, componentField)

		// The destination field (componentField) can itself contain components requiring expansion.
		// e.g. compressed_speed_distance -> (speed, distance), speed -> enhanced_speed.
		if subField := componentField.SubFieldSubtitution(mesg); subField != nil {
			d.expandComponents(mesg, &componentField, subField.Components)
		} else {
			d.expandComponents(mesg, &componentField, componentField.Components)
		}
	}
}

func (d *Decoder) decodeDeveloperFields(mesgDef *proto.MessageDefinition, mesg *proto.Message) error {
	for i := range mesgDef.DeveloperFieldDefinitions {
		devFieldDef := &mesgDef.DeveloperFieldDefinitions[i]

		var ok bool
		for _, developerDataIndex := range d.developerDataIndexes {
			if developerDataIndex == devFieldDef.DeveloperDataIndex {
				ok = true
				break
			}
		}

		if !ok {
			// NOTE: Currently, we allow missing DeveloperDataId message,
			// we only use FieldDescription messages to decode developer data.
			d.log("mesg.Num: %d, developerFields[%d].Num: %d: missing developer data id with developer data index '%d'",
				mesg.Num, i, devFieldDef.Num, devFieldDef.DeveloperDataIndex)
		}

		// Find the FieldDescription that refers to this DeveloperField.
		// The combination of the Developer Data Index and Field Definition Number
		// create a unique id for each Field Description.
		var fieldDesc *mesgdef.FieldDescription
		for _, f := range d.fieldDescriptions {
			if f.DeveloperDataIndex != devFieldDef.DeveloperDataIndex {
				continue
			}
			if f.FieldDefinitionNumber != devFieldDef.Num {
				continue
			}
			fieldDesc = f
			break
		}

		if fieldDesc == nil {
			d.log("mesg.Num: %d, developerFields[%d].Num: %d: Can't interpret developer field, "+
				"no field description mesg found. Just read acquired bytes (%d) and move forward. [byte pos: %d]\n",
				mesg.Num, i, devFieldDef.Num, devFieldDef.Size, d.n)
			if _, err := d.readN(int(devFieldDef.Size)); err != nil {
				return fmt.Errorf("no field description found, unable to read acquired bytes: %w", err)
			}
			continue
		}

		if !fieldDesc.FitBaseTypeId.Valid() {
			return fmt.Errorf("fieldDescription.FitBaseTypeId: %s: %w",
				fieldDesc.FitBaseTypeId, errInvalidBaseType)
		}

		var isArray bool
		typeSize := fieldDesc.FitBaseTypeId.Size()
		if devFieldDef.Size > typeSize && devFieldDef.Size%typeSize == 0 {
			isArray = true
		}

		baseType := fieldDesc.FitBaseTypeId
		profileType := profile.ProfileTypeFromBaseType(baseType)

		// Gracefully handle poorly encoded FIT file.
		if devFieldDef.Size == 0 {
			d.logDeveloperField(mesg, devFieldDef, fieldDesc.FitBaseTypeId, "Size is zero. Skip")
			continue
		} else if devFieldDef.Size < fieldDesc.FitBaseTypeId.Size() {
			baseType = basetype.Byte
			profileType = profile.Byte
			isArray = devFieldDef.Size > baseType.Size() && devFieldDef.Size&baseType.Size() == 0
			d.logDeveloperField(mesg, devFieldDef, fieldDesc.FitBaseTypeId,
				"Size is less than expected. Fallback: decode as byte(s) and convert the value")
		}

		// NOTE: It seems there is no standard on utilizing Array field to handle []string in developer fields.
		// Discussion: https://forums.garmin.com/developer/fit-sdk/f/discussion/355554/how-to-determine-developer-field-s-value-type-is-a-string-or-string
		overrideStringArray := fieldDesc.FitBaseTypeId == basetype.String
		val, err := d.readValue(devFieldDef.Size, mesgDef.Architecture, baseType, profileType, isArray, overrideStringArray)
		if err != nil {
			return err
		}

		if baseType != fieldDesc.FitBaseTypeId { // Convert value
			bitVal, _ := bitsFromValue(val)
			val = valueFromBits(bitVal, fieldDesc.FitBaseTypeId)
		}

		// NOTE: Decoder will not attempt to validate native data when both NativeMesgNum and NativeFieldNum are valid.
		// Users need to handle this themselves due to the limited context available.
		mesg.DeveloperFields = append(mesg.DeveloperFields,
			proto.DeveloperField{
				Num:                devFieldDef.Num,
				DeveloperDataIndex: devFieldDef.DeveloperDataIndex,
				Value:              val,
			})
	}
	return nil
}

func (d *Decoder) decodeCRC() error {
	b, err := d.readBuffer.ReadN(2)
	if err != nil {
		return err
	}
	d.n += 2
	d.crc = binary.LittleEndian.Uint16(b)
	if d.options.shouldChecksum && d.crc16.Sum16() != d.crc { // check data integrity
		return fmt.Errorf("expected crc %d, got: %d: %w", d.crc, d.crc16.Sum16(), ErrCRCChecksumMismatch)
	}
	d.crc16.Reset()
	return nil
}

func (d *Decoder) readN(n int) ([]byte, error) {
	b, err := d.readBuffer.ReadN(n)
	if err != nil {
		return nil, err
	}
	d.n, d.cur = d.n+int64(n), d.cur+uint32(n)
	if d.options.shouldChecksum {
		_, _ = d.crc16.Write(b)
	}
	return b, nil
}

// readValue reads message value bytes from reader and convert it into its corresponding type. Size should not be zero.
func (d *Decoder) readValue(size byte, arch byte, baseType basetype.BaseType, profileType profile.ProfileType, isArray, overrideStringArray bool) (val proto.Value, err error) {
	b, err := d.readN(int(size))
	if err != nil {
		return val, err
	}
	if overrideStringArray && baseType == basetype.String {
		isArray = strcount(b) > 1
	}
	return proto.UnmarshalValue(b, arch, baseType, profileType, isArray)
}

// log logs only if logWriter is not nil.
func (d *Decoder) log(format string, args ...any) {
	if d.options.logWriter == nil {
		return
	}
	fmt.Fprintf(d.options.logWriter, format, args...)
}

const logFieldTemplate = "mesg.Num: %q, %s.Num: %d, size: %d, type: %q (size: %d). %s. [bytes pos: %d]\n"

// logField logs field related issues only if logWriter is not nil.
func (d *Decoder) logField(m *proto.Message, fd *proto.FieldDefinition, msg string) {
	d.log(logFieldTemplate, m.Num, "field", fd.Num, fd.Size, fd.BaseType, fd.BaseType.Size(), msg, d.n)
}

// logDeveloperField logs developerField related issues only if logWriter is not nil.
func (d *Decoder) logDeveloperField(m *proto.Message, dfd *proto.DeveloperFieldDefinition, bt basetype.BaseType, msg string) {
	d.log(logFieldTemplate, m.Num, "developerField", dfd.Num, dfd.Size, bt, bt.Size(), msg, d.n)
}

// DecodeWithContext is similar to Decode but with respect to context propagation.
func (d *Decoder) DecodeWithContext(ctx context.Context) (*proto.FIT, error) {
	if d.err != nil {
		return nil, d.err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if d.err = checkContext(ctx); d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeFileHeaderOnce(); d.err != nil {
		return nil, d.err
	}
	defer d.releaseTemporaryObjects()
	if d.err = d.decodeMessagesWithContext(ctx); d.err != nil {
		return nil, d.err
	}
	if d.err = checkContext(ctx); d.err != nil {
		return nil, d.err
	}
	if d.err = d.decodeCRC(); d.err != nil {
		return nil, d.err
	}
	fit := &proto.FIT{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}
	d.reset()
	return fit, nil
}

func checkContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (d *Decoder) decodeMessagesWithContext(ctx context.Context) (err error) {
	for d.cur < d.fileHeader.DataSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err = d.decodeMessage(); err != nil {
				return fmt.Errorf("decodeMessage [byte pos: %d]: %w", d.n, err)
			}
		}
	}
	return nil
}

// strcount counts how many valid string in b.
// This should align with the logic in proto.UnmarshalValue.
func strcount(b []byte) (size byte) {
	last := 0
	for i := range b {
		if b[i] == '\x00' {
			if last != i { // only if not an invalid string
				size++
			}
			last = i + 1
		}
	}
	return size
}
