// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/byteorder"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/listener"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	// Integrity errors
	ErrNotAFitFile         = errors.New("not a fit file")
	ErrDataSizeZero        = errors.New("data size zero")
	ErrCRCChecksumMismatch = errors.New("crc checksum mismatch")

	// Message-field related errors
	ErrMesgDefMissing         = errors.New("message definition missing")
	ErrFieldValueTypeMismatch = errors.New("field value type mismatch")
	ErrByteSizeMismatch       = errors.New("byte size mismath")
)

// Decoder is Fit file decoder. See New() for details.
type Decoder struct {
	r           io.Reader
	factory     Factory
	accumulator *Accumulator
	crc16       hash.Hash16

	// The maximum n field definition in a mesg is 255 and we need 3 byte per field. 255 * 3 = 765.
	// So we will never exceed 765.
	bytesArray [255 * 3]byte

	// The maximum n field in a mesg is 255, with this backing array, we don't have to worry about component expansions
	// that require allocating additional fields triggering runtime.mallocgc.
	fieldsArray [256]proto.Field

	options *options

	decodeHeaderOnce  func() error // The func to decode header exactly once, return the error of the first invocation if any. Initialized on New().
	n                 int64        // The n read bytes counter, always moving forward, do not reset (except on full reset).
	cur               uint32       // The current byte position relative to bytes of the messages, reset on next chained Fit file.
	timestamp         uint32       // Active timestamp
	lastTimeOffset    byte         // Last time offset
	sequenceCompleted bool         // True after a decode is completed. Reset to false on Next().
	err               error        // Any error occurs during process.

	// Fit File Representation
	fileHeader proto.FileHeader
	messages   []proto.Message
	crc        uint16

	// FileId Message is a special message that must present in a Fit file.
	fileId *mesgdef.FileId

	// Message Definition Lookup
	localMessageDefinitions [proto.LocalMesgNumMask + 1]*proto.MessageDefinition // message definition for upcoming message data

	// Developer Data Lookup
	developerDataIds  []*mesgdef.DeveloperDataId
	fieldDescriptions []*mesgdef.FieldDescription

	// Listeners
	mesgListeners    []listener.MesgListener    // Each listener will received every decoded message.
	mesgDefListeners []listener.MesgDefListener // Each listener will received every decoded message definition.
}

// Factory defines a contract that any Factory containing these method can be used by the Decoder.
type Factory interface {
	// CreateMesgOnly create new message with Fields and DeveloperFields are being nil. If not found, it returns new message with "unknown" name.
	CreateMesgOnly(mesgNum typedef.MesgNum) proto.Message
	// CreateField create new field based on defined messages in the factory. If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

type options struct {
	factory               Factory
	mesgListeners         []listener.MesgListener
	mesgDefListeners      []listener.MesgDefListener
	shouldChecksum        bool
	broadcastOnly         bool
	shouldExpandComponent bool
}

func defaultOptions() *options {
	return &options{
		factory:               factory.StandardFactory(),
		shouldChecksum:        true,
		broadcastOnly:         false,
		shouldExpandComponent: true,
	}
}

type Option interface{ apply(o *options) }

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

// WithFactory sets custom factory.
func WithFactory(factory Factory) Option {
	return fnApply(func(o *options) {
		if factory != nil {
			o.factory = factory
		}
	})
}

// WithMesgListener adds listeners to the listener pool, where each listener is broadcasted every message.
// The listeners will be appended not replaced. If users need to reset use Reset().
func WithMesgListener(listeners ...listener.MesgListener) Option {
	return fnApply(func(o *options) {
		o.mesgListeners = append(o.mesgListeners, listeners...)
	})
}

// WithMesgDefListener adds listeners to the listener pool, where each listener is broadcasted every message definition.
// The listeners will be appended not replaced. If users need to reset use Reset().
func WithMesgDefListener(listeners ...listener.MesgDefListener) Option {
	return fnApply(func(o *options) {
		o.mesgDefListeners = append(o.mesgDefListeners, listeners...)
	})
}

// WithBroadcastOnly directs the Decoder to only broadcast the messages without retaining them, reducing memory usage when
// it's not going to be used anyway. This option is intended to be used with WithMesgListener and WithMesgDefListener.
// When this option is specified, the Decode will return a fit with empty messages.
func WithBroadcastOnly() Option {
	return fnApply(func(o *options) { o.broadcastOnly = true })
}

// WithIgnoreChecksum directs the Decoder to not checking data integrity (CRC Checksum).
func WithIgnoreChecksum() Option {
	return fnApply(func(o *options) { o.shouldChecksum = false })
}

// WithNoComponentExpansion directs the Decoder to not expand the components.
func WithNoComponentExpansion() Option {
	return fnApply(func(o *options) { o.shouldExpandComponent = false })
}

// New returns a FIT File Decoder to decode given r.
//
// The FIT protocol allows for multiple FIT files to be chained together in a single FIT file.
// Each FIT file in the chain must be a properly formatted FIT file (header, data records, CRC).
//
// To decode chained FIT files, use Next() to check if r hasn't reach EOF and next bytes are still a valid FIT sequences.
//
//	for dec.Next() {
//	   fit, err := dec.Decode()
//	}
//
// Note: We encourage wrapping r into a buffered reader such as bufio.NewReader(r),
// decode process requires byte by byte reading and having frequent read on non-buffered reader might impact performance,
// especially if it involves syscall such as reading a file.
func New(r io.Reader, opts ...Option) *Decoder {
	options := defaultOptions()
	for i := range opts {
		opts[i].apply(options)
	}

	d := &Decoder{
		r:                       r,
		options:                 options,
		factory:                 options.factory,
		accumulator:             NewAccumulator(),
		crc16:                   crc16.New(crc16.MakeFitTable()),
		localMessageDefinitions: [proto.LocalMesgNumMask + 1]*proto.MessageDefinition{},
		mesgListeners:           options.mesgListeners,
		mesgDefListeners:        options.mesgDefListeners,
		developerDataIds:        make([]*mesgdef.DeveloperDataId, 0),
		fieldDescriptions:       make([]*mesgdef.FieldDescription, 0),
	}
	d.initDecodeHeaderOnce()

	return d
}

// initDecodeHeaderOnce initializes decodeHeaderOnce() to decode the header exactly once, if error occur
// during the first invocation, the error will be returned everytime decodeHeaderOnce() is invoked.
func (d *Decoder) initDecodeHeaderOnce() {
	var once = sync.Once{}
	d.decodeHeaderOnce = func() error {
		once.Do(func() { d.err = d.decodeHeader() })
		return d.err
	}
}

// PeekFileId decodes only up to FileId message without decoding the whole reader.
// FileId message should be the first message of any Fit file, otherwise return an error.
//
// After this method is invoked, Decode picks up where this left then continue decoding next messages instead of starting from zero.
// This method is idempotent and can be invoked even after Decode has been invoked.
func (d *Decoder) PeekFileId() (fileId *mesgdef.FileId, err error) {
	if d.err != nil {
		return nil, d.err
	}
	defer func() { d.err = err }()
	if err = d.decodeHeaderOnce(); err != nil {
		return
	}
	for d.fileId == nil {
		if err = d.decodeMessage(); err != nil {
			return
		}
	}
	return d.fileId, nil
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
		// Check Header Integrity
		pos := d.n
		if err = d.decodeHeaderOnce(); err != nil {
			if pos != 0 && pos == d.n && err == io.EOF {
				// When EOF error occurs exactly after a sequence has been completed,
				// make the error as nil, it means we have reached the desirable EOF.
				err = nil
			}
			break
		}
		// Read bytes acquired by messages to calculate crc checksum of its contents
		for d.cur < d.fileHeader.DataSize {
			size := d.fileHeader.DataSize - d.cur
			if arraySize := uint32(len(d.bytesArray)); size > arraySize {
				size = arraySize
			}
			if err = d.read(d.bytesArray[:size]); err != nil { // Discard bytes
				break
			}
		}
		if err = d.decodeCRC(); err != nil {
			break
		}
		// Check crc checksum of messages should match with file's crc.
		if d.crc16.Sum16() != d.crc {
			err = ErrCRCChecksumMismatch
			break
		}
		d.initDecodeHeaderOnce()
		d.crc16.Reset()
		d.cur = 0
		seq++
	}

	if err != nil { // When there is an error, wrap it with informative message before return.
		err = fmt.Errorf("byte pos: %d: %w", d.n, err)
	}

	// Reset used variables so that the decoder can be reused by the same reader.
	d.reset()
	d.n = 0 // Must reset bytes counter
	d.err = err
	d.options.shouldChecksum = shouldChecksum

	return seq, err
}

// discardMessages efficiently discards bytes used by messages.
func (d *Decoder) discardMessages() (err error) {
	arraySize := uint32(len(d.bytesArray))
	for d.cur < d.fileHeader.DataSize {
		size := d.fileHeader.DataSize - d.cur
		if size > arraySize {
			size = arraySize
		}
		if err = d.read(d.bytesArray[:size]); err != nil { // Discard bytes
			return err
		}
	}
	return nil
}

// Discard discards a single FIT file sequence and returns any error encountered. This method directs the Decoder to
// point to the byte sequence of the next valid FIT file sequence, discarding the current FIT file sequence.
//
// Example: - A chained FIT file consist of Activity, Course, Workout and Settings. And we only want to decode Course.
//
//		for dec.Next() {
//			fileId, err := dec.PeekFileId()
//			if err != nil {
//				return err
//			}
//			if fileId.Type != typedef.FileCourse {
//				if err := dec.Discard(); err != nil {
//			    	return err
//			    }
//				continue
//			}
//			fit, err := dec.Decode()
//			if err != nil {
//				return err
//	     	}
//		 }
func (d *Decoder) Discard() error {
	if d.err != nil {
		return d.err
	}

	optionsShouldChecksum := d.options.shouldChecksum
	d.options.shouldChecksum = false
	defer func() { d.options.shouldChecksum = optionsShouldChecksum }()

	if d.err = d.decodeHeaderOnce(); d.err != nil {
		return d.err
	}
	if d.err = d.discardMessages(); d.err != nil {
		return d.err
	}
	if d.err = d.read(d.bytesArray[:2]); d.err != nil { // Discard File CRC
		return d.err
	}
	d.sequenceCompleted = true
	return d.err
}

// Next checks whether next bytes are still a valid Fit File sequence. Return false when invalid or reach EOF.
func (d *Decoder) Next() bool {
	if d.err != nil {
		return false
	}

	if !d.sequenceCompleted {
		return true
	}

	d.reset() // reset values for the next chained Fit file

	// err is saved in the func, any exported will call this func anyway.
	return d.decodeHeaderOnce() == nil
}

func (d *Decoder) reset() {
	for i := range d.localMessageDefinitions {
		d.localMessageDefinitions[i] = nil
	}

	d.accumulator.Reset()
	d.crc16.Reset()
	d.fileHeader = proto.FileHeader{}
	if !d.options.broadcastOnly {
		d.messages = nil // Must create new.
	}
	d.fileId = nil
	d.developerDataIds = d.developerDataIds[:0]
	d.fieldDescriptions = d.fieldDescriptions[:0]
	d.crc = 0
	d.cur = 0
	d.timestamp = 0
	d.lastTimeOffset = 0
	d.sequenceCompleted = false

	d.initDecodeHeaderOnce() // reset to enable invocation.
}

// Reset resets the Decoder to read its input from r, clear any error and
// reset previous options to default options so any options needs to be inputed again.
// It is similar to New() but it retains the underlying storage for use by
// future decode to reduce memory allocs (except messages need to be re-allocated).
func (d *Decoder) Reset(r io.Reader, opts ...Option) {
	d.reset()
	d.n = 0 // Must reset bytes counter since it's a full reset.
	d.err = nil
	d.r = r

	d.options = defaultOptions()
	for i := range opts {
		opts[i].apply(d.options)
	}

	d.factory = d.options.factory
	d.mesgListeners = d.options.mesgListeners
	d.mesgDefListeners = d.options.mesgDefListeners
}

// Decode method decodes `r` into Fit data. One invocation will produce one valid Fit data or an error if it occurs.
// To decode a chained Fit file that contains more than one Fit data, this decode method should be invoked
// multiple times. It is recommended to wrap it with the Next() method when you are uncertain if it's a chained fit file.
//
//	for dec.Next() {
//	     fit, err := dec.Decode()
//	     if err != nil {
//	         return err
//	     }
//	}
func (d *Decoder) Decode() (fit *proto.Fit, err error) {
	if d.err != nil {
		return nil, d.err
	}
	defer func() { d.err = err }()
	if err = d.decodeHeaderOnce(); err != nil {
		return nil, err
	}
	if err = d.decodeMessages(); err != nil {
		return nil, err
	}
	if err = d.decodeCRC(); err != nil {
		return nil, err
	}
	if d.options.shouldChecksum && d.crc16.Sum16() != d.crc { // check data integrity
		err = fmt.Errorf("expected crc %d, got: %d: %w", d.crc, d.crc16.Sum16(), ErrCRCChecksumMismatch)
		return nil, err
	}
	d.sequenceCompleted = true
	return &proto.Fit{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}, nil
}

func (d *Decoder) decodeHeader() error {
	b := d.bytesArray[:1]
	n, err := io.ReadFull(d.r, b)
	d.n += int64(n)
	if err != nil {
		return err
	}

	if b[0] != 12 && b[0] != 14 { // current spec is either 12 or 14
		return fmt.Errorf("header size [%d] is invalid: %w", b[0], ErrNotAFitFile)
	}

	size := b[0]
	b = d.bytesArray[1:size]
	n, err = io.ReadFull(d.r, b)
	d.n += int64(n)
	if err != nil {
		return err
	}

	d.fileHeader = proto.FileHeader{
		Size:            size,
		ProtocolVersion: b[0],
		ProfileVersion:  binary.LittleEndian.Uint16(b[1:3]),
		DataSize:        binary.LittleEndian.Uint32(b[3:7]),
	}

	// PERF: Neither string(b[7:11]) nor assigning proto.DataTypeFIT constant to a variable escape to the heap.
	if string(b[7:11]) != proto.DataTypeFIT {
		return ErrNotAFitFile
	}
	d.fileHeader.DataType = proto.DataTypeFIT

	if err := proto.Validate(d.fileHeader.ProtocolVersion); err != nil {
		return err
	}

	if d.fileHeader.DataSize == 0 {
		return ErrDataSizeZero
	}

	if size == 14 {
		d.fileHeader.CRC = binary.LittleEndian.Uint16(b[11:13])
	}

	if d.fileHeader.CRC == 0x0000 || !d.options.shouldChecksum { // do not need to check header's crc integrity.
		return nil
	}

	_, _ = d.crc16.Write(d.bytesArray[:12])
	if d.crc16.Sum16() != d.fileHeader.CRC { // check header integrity
		return fmt.Errorf("expected header's crc: %d, got: %d: %w", d.fileHeader.CRC, d.crc16.Sum16(), ErrCRCChecksumMismatch)
	}

	d.crc16.Reset() // this hash will be re-used for calculating data integrity.

	return nil
}

func (d *Decoder) decodeMessages() error {
	for d.cur < d.fileHeader.DataSize {
		if err := d.decodeMessage(); err != nil {
			return fmt.Errorf("decodeMessage [byte pos: %d]: %w", d.n, err)
		}
	}
	return nil
}

func (d *Decoder) decodeMessage() error {
	header, err := d.readByte()
	if err != nil {
		return err
	}

	if (header & proto.MesgDefinitionMask) == proto.MesgDefinitionMask {
		return d.decodeMessageDefinition(header)
	}

	return d.decodeMessageData(header)
}

func (d *Decoder) decodeMessageDefinition(header byte) error {
	b := d.bytesArray[:5]
	if err := d.read(b); err != nil {
		return err
	}

	// PERF: Reuse existing object when possible, as it is intended solely for temporary use.
	localMesgNum := header & proto.LocalMesgNumMask
	mesgDef := d.localMessageDefinitions[localMesgNum]
	if mesgDef == nil {
		mesgDef = &proto.MessageDefinition{}
	}

	mesgDef.Header = header
	mesgDef.Reserved = b[0]
	mesgDef.Architecture = b[1]
	mesgDef.MesgNum = typedef.MesgNum(byteorder.Select(b[1]).Uint16(b[2:4]))

	n := b[4]
	b = d.bytesArray[:uint16(n)*3] // 3 byte per field
	if err := d.read(b); err != nil {
		return err
	}

	mesgDef.FieldDefinitions = mesgDef.FieldDefinitions[:0]
	if byte(cap(mesgDef.FieldDefinitions)) < n { // PERF: Only alloc when necessary
		mesgDef.FieldDefinitions = make([]proto.FieldDefinition, 0, n)
	}

	for i := 0; i < len(b); i += 3 {
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions, proto.FieldDefinition{
			Num:      b[i],
			Size:     b[i+1],
			BaseType: basetype.BaseType(b[i+2]),
		})
	}

	mesgDef.DeveloperFieldDefinitions = mesgDef.DeveloperFieldDefinitions[:0]
	if (header & proto.DevDataMask) == proto.DevDataMask {
		n, err := d.readByte()
		if err != nil {
			return err
		}

		b := d.bytesArray[:uint16(n)*3] // 3 byte per field
		if err := d.read(b); err != nil {
			return err
		}

		if byte(cap(mesgDef.DeveloperFieldDefinitions)) < n { // PERF: Only alloc when necessary
			mesgDef.DeveloperFieldDefinitions = make([]proto.DeveloperFieldDefinition, 0, n)
		}

		for i := 0; i < len(b); i += 3 {
			mesgDef.DeveloperFieldDefinitions = append(mesgDef.DeveloperFieldDefinitions, proto.DeveloperFieldDefinition{
				Num:                b[i],
				Size:               b[i+1],
				DeveloperDataIndex: b[i+2],
			})
		}
	}

	d.localMessageDefinitions[localMesgNum] = mesgDef

	if len(d.mesgDefListeners) > 0 {
		mesgDef := mesgDef.Clone() // Clone since we don't have control of the object lifecycle outside Decoder.
		for _, mesgDefListener := range d.mesgDefListeners {
			mesgDefListener.OnMesgDef(mesgDef) // blocking or non-blocking depends on listeners' implementation.
		}
	}

	return nil
}

func (d *Decoder) decodeMessageData(header byte) error {
	localMesgNum := proto.LocalMesgNum(header)
	mesgDef := d.localMessageDefinitions[localMesgNum]
	if mesgDef == nil {
		return ErrMesgDefMissing
	}

	mesg := d.factory.CreateMesgOnly(mesgDef.MesgNum)
	mesg.Header = header
	mesg.Reserved = mesgDef.Reserved
	mesg.Architecture = mesgDef.Architecture
	mesg.Fields = d.fieldsArray[:0]

	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask { // Compressed Timestamp Message Data
		timeOffset := header & proto.CompressedTimeMask
		d.timestamp += uint32((timeOffset - d.lastTimeOffset) & proto.CompressedTimeMask)
		d.lastTimeOffset = timeOffset

		timestampField := d.factory.CreateField(mesgDef.MesgNum, proto.FieldNumTimestamp)
		timestampField.Value = d.timestamp

		mesg.Fields = append(mesg.Fields, timestampField) // add timestamp field
	}

	if err := d.decodeFields(mesgDef, &mesg); err != nil {
		return err
	}

	mesg.Fields = make([]proto.Field, len(mesg.Fields))
	copy(mesg.Fields, d.fieldsArray[:])

	// FileId Message
	if d.fileId == nil && mesg.Num == mesgnum.FileId {
		d.fileId = mesgdef.NewFileId(&mesg)
	}

	// Prerequisites for decoding developer fields
	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		// These messages must occur before any related field description messages are written to the proto.
		d.developerDataIds = append(d.developerDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		// These messages must occur in the file before any related developer data is written to the proto.
		d.fieldDescriptions = append(d.fieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	}

	if len(mesgDef.DeveloperFieldDefinitions) != 0 {
		if err := d.decodeDeveloperFields(mesgDef, &mesg); err != nil {
			return err
		}
	}

	if !d.options.broadcastOnly {
		d.messages = append(d.messages, mesg)
	}

	for _, mesgListener := range d.mesgListeners {
		mesgListener.OnMesg(mesg) // blocking or non-blocking depends on listeners' implementation.
	}

	return nil
}

func (d *Decoder) decodeFields(mesgDef *proto.MessageDefinition, mesg *proto.Message) error {
	for i := range mesgDef.FieldDefinitions {
		fieldDef := &mesgDef.FieldDefinitions[i]

		field := d.factory.CreateField(mesgDef.MesgNum, fieldDef.Num)
		if field.Name == factory.NameUnknown {
			// Assign fieldDef's type for unknown field so later we can encode it as per its original value.
			field.Type = profile.ProfileTypeFromBaseType(fieldDef.BaseType)
			field.BaseType = fieldDef.BaseType
			// Check if the size corresponds to an array.
			field.Array = fieldDef.Size > fieldDef.BaseType.Size() && fieldDef.Size%fieldDef.BaseType.Size() == 0
		}

		val, err := d.readValue(fieldDef.Size, fieldDef.BaseType, field.Array, mesgDef.Architecture)
		if err != nil {
			return err
		}

		if field.Num == proto.FieldNumTimestamp {
			timestamp, ok := val.(uint32)
			if !ok {
				// This can only happen when:
				// 1. Profile.xlsx contain typo from official release or user add manufacturer specific message but specifying wrong type.
				// 2. User register the message in the factory but using different type.
				return fmt.Errorf("timestamp should be uint32, got: %T: %w", val, ErrFieldValueTypeMismatch)
			}
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
	if len(components) == 0 {
		return
	}

	bitVal, ok := bitsFromValue(containingField.Value)
	if !ok {
		return
	}

	// PERF: Reuse a single variable 'componentField' instead of declaring it inside the loop to prevent
	// the Compiler's escape analysis from moving it to the heap, which could occur due to the risk of
	// stack overflow caused by repeatedly creating the variable.
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
	mesg.DeveloperFields = make([]proto.DeveloperField, 0, len(mesgDef.DeveloperFieldDefinitions))
	for i := range mesgDef.DeveloperFieldDefinitions {
		devFieldDef := &mesgDef.DeveloperFieldDefinitions[i]

		// Find the FieldDescription that refers to this DeveloperField.
		// The combination of the Developer Data Index and Field Definition Number
		// create a unique id for each Field Description.
		var fieldDescription *mesgdef.FieldDescription
		for i := range d.fieldDescriptions {
			if d.fieldDescriptions[i].DeveloperDataIndex != devFieldDef.DeveloperDataIndex {
				continue
			}
			if d.fieldDescriptions[i].FieldDefinitionNumber != devFieldDef.Num {
				continue
			}
			fieldDescription = d.fieldDescriptions[i]
			break
		}

		if fieldDescription == nil {
			// Can't interpret this DeveloperField, no FieldDescription found. Just read acquired bytes and move forward.
			if err := d.read(d.bytesArray[:devFieldDef.Size]); err != nil {
				return fmt.Errorf("no field description found, unable to read acquired bytes: %w", err)
			}
			continue
		}

		developerField := proto.DeveloperField{
			Num:                devFieldDef.Num,
			DeveloperDataIndex: devFieldDef.DeveloperDataIndex,
			Size:               devFieldDef.Size,
			NativeMesgNum:      fieldDescription.NativeMesgNum,
			NativeFieldNum:     fieldDescription.NativeFieldNum,
			BaseType:           fieldDescription.FitBaseTypeId,
		}

		developerField.Name = strings.Join(fieldDescription.FieldName, "|")
		developerField.Units = strings.Join(fieldDescription.Units, "|")

		// TODO: We still don't know how []string should be handled in the developer field.
		// For the Field, we have "Array" (bool) for determining if the value is an array.
		// However, we could not find any reference on how to use the DeveloperField's Array (uint8).
		//
		// For example:
		// - "suuntoplus_plugin_owner_id" is []string, 1 - 10 strings, 1 - 64 characters each. (ref: https://apizone.suunto.com/fit-description).
		//   but still it does not specify DeveloperField's Array.
		//
		// Until we discover a better implementation, let's determine it by using multiplication of the size.
		// Consequently, all strings will be treated as arrays since its size is 1.
		var isArray bool
		if devFieldDef.Size > developerField.BaseType.Size() && devFieldDef.Size%developerField.BaseType.Size() == 0 {
			isArray = true
		}

		val, err := d.readValue(developerField.Size, developerField.BaseType, isArray, mesgDef.Architecture)
		if err != nil {
			return err
		}

		developerField.Value = val

		mesg.DeveloperFields = append(mesg.DeveloperFields, developerField)
	}
	return nil
}

func (d *Decoder) decodeCRC() error {
	b := d.bytesArray[:2]
	n, err := io.ReadFull(d.r, b)
	d.n += int64(n)
	if err != nil {
		return err
	}
	d.crc = binary.LittleEndian.Uint16(b)
	return nil
}

// read is used for reading messages from reader as this will mark the position and also calculate the crc as we read.
func (d *Decoder) read(b []byte) error {
	n, err := io.ReadFull(d.r, b)
	d.n, d.cur = d.n+int64(n), d.cur+uint32(n)
	if err != nil {
		return err
	}
	if d.options.shouldChecksum {
		_, _ = d.crc16.Write(b)
	}
	return nil
}

// readByte is shorthand for read([1]byte).
func (d *Decoder) readByte() (byte, error) {
	b := d.bytesArray[:1]
	err := d.read(b)
	return b[0], err
}

// readValue reads message value bytes from reader and convert it into its corresponding type.
func (d *Decoder) readValue(size byte, baseType basetype.BaseType, isArray bool, arch byte) (any, error) {
	b := d.bytesArray[:size]
	if err := d.read(b); err != nil {
		return nil, err
	}

	val, err := typedef.Unmarshal(b, byteorder.Select(arch), baseType, isArray)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// DecodeWithContext is similar to Decode but with respect to context propagation.
func (d *Decoder) DecodeWithContext(ctx context.Context) (fit *proto.Fit, err error) {
	if d.err != nil {
		return nil, d.err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	defer func() { d.err = err }()
	if err = checkContext(ctx); err != nil {
		return nil, err
	}
	if err = d.decodeHeaderOnce(); err != nil {
		return nil, err
	}
	if err = d.decodeMessagesWithContext(ctx); err != nil {
		return nil, err
	}
	if err = checkContext(ctx); err != nil {
		return nil, err
	}
	if err = d.decodeCRC(); err != nil {
		return nil, err
	}
	if d.options.shouldChecksum && d.crc16.Sum16() != d.crc { // check data integrity
		err = fmt.Errorf("expected crc %d, got: %d: %w", d.crc, d.crc16.Sum16(), ErrCRCChecksumMismatch)
		return nil, err
	}
	d.sequenceCompleted = true
	return &proto.Fit{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}, nil
}

func checkContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (d *Decoder) decodeMessagesWithContext(ctx context.Context) error {
	for d.cur < d.fileHeader.DataSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := d.decodeMessage(); err != nil {
				return fmt.Errorf("decodeMessage [byte pos: %d]: %w", d.n, err)
			}
		}
	}
	return nil
}
