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
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/byteorder"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/kit/typeconv"
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
	ErrCRCChecksumMismatch = errors.New("crc checksum mismatch")

	// Message-field related errors
	ErrMesgDefMissing         = errors.New("message definition missing")
	ErrFieldValueTypeMismatch = errors.New("field value type mismatch")
	ErrByteSizeMismatch       = errors.New("byte size mismath")
)

const (
	fieldNumTimestamp = 253 // Num for timestamp across all defined messages in the profile.

	// Buffer for component expansion to avoid runtime grow slice which more expensive than having buffered size at front.
	// The value 10 is from the current max components in factory (MesgNumHr -> event_timestamp_12).
	bufferSizeFields = 10
)

// Decoder is Fit file decoder. See New() for details.
type Decoder struct {
	r           io.Reader
	factory     Factory
	accumulator *Accumulator
	crc16       hash.Hash16

	options *options

	decodeHeaderOnce func() error // The func to decode header exactly once, return the error of the first invocation if any. Initialized on New().
	n                int64        // The n read bytes counter, always moving forward, do not reset
	cur              uint32       // The current byte position relative to bytes of the messages, reset on next chained Fit file.
	timestamp        uint32       // Active timestamp
	lastTimeOffset   byte         // Last time offset

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

// WithMesgListener adds a listener to the listener pool, where each listener is broadcasted every message.
func WithMesgListener(lis listener.MesgListener) Option {
	return fnApply(func(o *options) {
		if lis != nil {
			o.mesgListeners = append(o.mesgListeners, lis)
		}
	})
}

// WithMesgDefListener adds a listener to the listener pool, where each listener is broadcasted every message definition .
func WithMesgDefListener(lis listener.MesgDefListener) Option {
	return fnApply(func(o *options) {
		if lis != nil {
			o.mesgDefListeners = append(o.mesgDefListeners, lis)
		}
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
//	   fit, err := dec.Decode(context.Background())
//	}
//
// Note: We encourage wrapping r into a buffered reader such as bufio.NewReader(r),
// decode process requires byte by byte reading and having frequent read on non-buffered reader might impact performance,
// especially if it involves syscall such as reading a file.
func New(r io.Reader, opts ...Option) *Decoder {
	options := defaultOptions()
	for _, o := range opts {
		o.apply(options)
	}

	d := &Decoder{
		r:                       r,
		options:                 options,
		factory:                 options.factory,
		accumulator:             NewAccumulator(),
		crc16:                   crc16.New(crc16.MakeFitTable()),
		localMessageDefinitions: [proto.LocalMesgNumMask + 1]*proto.MessageDefinition{},
		messages:                make([]proto.Message, 0),
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
	var once, err = sync.Once{}, error(nil)
	var f = func() { err = d.decodeHeader() }
	d.decodeHeaderOnce = func() error {
		once.Do(f)
		return err
	}
}

// PeekFileId decodes only up to FileId message without decoding the whole reader.
// FileId message should be the first message of any Fit file, otherwise return an error.
//
// After this method is invoked, Decode picks up where this left then continue decoding next messages instead of starting from zero.
// This method is idempotent and can be invoked even after Decode has been invoked.
func (d *Decoder) PeekFileId() (fileId *mesgdef.FileId, err error) {
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

// Next checks whether next bytes are still a valid Fit File sequence. Return false when invalid or reach EOF.
func (d *Decoder) Next() bool {
	// reset values for the next chained Fit file
	d.accumulator = NewAccumulator()
	d.crc16.Reset()
	d.fileHeader = proto.FileHeader{}
	d.localMessageDefinitions = [proto.LocalMesgNumMask + 1]*proto.MessageDefinition{}
	d.messages = make([]proto.Message, 0)
	d.fileId = nil
	d.developerDataIds = make([]*mesgdef.DeveloperDataId, 0)
	d.fieldDescriptions = make([]*mesgdef.FieldDescription, 0)
	d.crc = 0
	d.cur = 0
	d.timestamp = 0
	d.lastTimeOffset = 0

	d.initDecodeHeaderOnce() // reset to enable invocation.

	// err is saved in the func, any exported will call this func anyway.
	return d.decodeHeaderOnce() == nil
}

// DecodeWithContext wraps Decode to respect context propagation.
func (d *Decoder) DecodeWithContext(ctx context.Context) (fit *proto.Fit, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	done := make(chan struct{})
	go func() {
		fit, err = d.Decode()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		return
	}
}

// Decode method decodes `r` into Fit data. One invocation will produce one valid Fit data or an error if it occurs.
// To decode a chained Fit file that contains more than one Fit data, this decode method should be invoked
// multiple times. It is recommended to wrap it with the Next() method when you are uncertain if it's a chained fit file.
//
//	for dec.Next() {
//	   fit, err := dec.Decode()
//	}
func (d *Decoder) Decode() (fit *proto.Fit, err error) {
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
		return nil, ErrCRCChecksumMismatch
	}
	return &proto.Fit{
		FileHeader: d.fileHeader,
		Messages:   d.messages,
		CRC:        d.crc,
	}, nil
}

func (d *Decoder) decodeHeader() error {
	b := make([]byte, 1)
	n, err := io.ReadFull(d.r, b)
	d.n += int64(n)
	if err != nil {
		return err
	}

	if b[0] != 12 && b[0] != 14 { // current spec is either 12 or 14
		return fmt.Errorf("header size [%d] is invalid: %w", b[0], ErrNotAFitFile)
	}

	size := b[0]
	b = make([]byte, size-1)
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
		DataType:        string(b[7:11]),
	}

	if err := proto.Validate(d.fileHeader.ProtocolVersion); err != nil {
		return err
	}

	if d.fileHeader.DataType != proto.DataTypeFIT {
		return ErrNotAFitFile
	}

	if size == 14 {
		d.fileHeader.CRC = binary.LittleEndian.Uint16(b[11:13])
	}

	if d.fileHeader.CRC == 0x0000 { // do not need to check header's crc integrity.
		return nil
	}

	_, _ = d.crc16.Write(append([]byte{size}, b[:11]...))

	if d.options.shouldChecksum && d.crc16.Sum16() != d.fileHeader.CRC { // check header integrity
		return ErrCRCChecksumMismatch
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
	b := make([]byte, 5)
	if err := d.read(b); err != nil {
		return err
	}

	mesgDef := proto.MessageDefinition{
		Header:       header,
		Reserved:     b[0],
		Architecture: b[1],
		MesgNum:      typedef.MesgNum(byteorder.Select(b[1]).Uint16(b[2:4])),
	}

	n := b[4]
	b = make([]byte, n*3) // 3 byte per field
	if err := d.read(b); err != nil {
		return err
	}

	mesgDef.FieldDefinitions = make([]proto.FieldDefinition, 0, n)
	for i := 0; i < len(b); i += 3 {
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions, proto.FieldDefinition{
			Num:      b[i],
			Size:     b[i+1],
			BaseType: basetype.BaseType(b[i+2]),
		})
	}

	if (header & proto.DevDataMask) == proto.DevDataMask {
		n, err := d.readByte()
		if err != nil {
			return err
		}

		b = make([]byte, n*3) // 3 byte per field
		if err := d.read(b); err != nil {
			return err
		}

		mesgDef.DeveloperFieldDefinitions = make([]proto.DeveloperFieldDefinition, 0, n)
		for i := 0; i < len(b); i += 3 {
			mesgDef.DeveloperFieldDefinitions = append(mesgDef.DeveloperFieldDefinitions, proto.DeveloperFieldDefinition{
				Num:                b[i],
				Size:               b[i+1],
				DeveloperDataIndex: b[i+2],
			})
		}
	}

	localMesgNum := header & proto.LocalMesgNumMask
	d.localMessageDefinitions[localMesgNum] = &mesgDef

	for _, mesgDefListener := range d.mesgDefListeners {
		mesgDefListener.OnMesgDef(mesgDef) // blocking or non-blocking depends on listeners' implementation.
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
	mesg.Reserved = mesgDef.Reserved

	if d.options.shouldExpandComponent {
		mesg.Fields = make([]proto.Field, 0, len(mesgDef.FieldDefinitions)+bufferSizeFields)
	} else {
		mesg.Fields = make([]proto.Field, 0, len(mesgDef.FieldDefinitions))
	}

	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask { // Compressed Timestamp Message Data
		timeOffset := header & proto.CompressedTimeMask
		d.timestamp += uint32((timeOffset - d.lastTimeOffset) & proto.CompressedTimeMask)
		d.lastTimeOffset = timeOffset

		timestampField := d.factory.CreateField(mesgDef.MesgNum, fieldNumTimestamp)
		timestampField.Value = d.timestamp

		mesg.Fields = append(mesg.Fields, timestampField) // add timestamp field
	}

	mesg.Header = header
	mesg.Architecture = mesgDef.Architecture

	if err := d.decodeFields(mesgDef, &mesg); err != nil {
		return err
	}

	// FileId Message
	if d.fileId == nil && mesg.Num == mesgnum.FileId {
		d.fileId = mesgdef.NewFileId(mesg)
	}

	// Prerequisites for decoding developer fields
	switch mesg.Num {
	case mesgnum.DeveloperDataId:
		// These messages must occur before any related field description messages are written to the proto.
		d.developerDataIds = append(d.developerDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		// These messages must occur in the file before any related developer data is written to the proto.
		d.fieldDescriptions = append(d.fieldDescriptions, mesgdef.NewFieldDescription(mesg))
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
			// Assign fieldDef's size and type for unknown field so later we can encode it as per its original value.
			field.Size = fieldDef.Size
			field.Type = profile.ProfileTypeFromString(fieldDef.BaseType.String())
		}

		val, err := d.readValue(fieldDef.Size, fieldDef.BaseType, mesgDef.Architecture)
		if err != nil {
			return err
		}

		if field.Num == fieldNumTimestamp {
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

		if isEligibleToAccumulate(field.Accumulate, field.Type.BaseType()) {
			if val, ok := typeconv.NumericToInt64(field.Value); ok {
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

		if subField, ok := field.SubFieldSubtitution(mesg); ok {
			// Expand sub-field components as the main field components
			d.expandComponents(mesg, field, subField.Components)
			continue
		}
		// No sub-field can interpret as main field, expand main field components
		d.expandComponents(mesg, field, field.Components)
	}

	return nil
}

func isEligibleToAccumulate(accumulate bool, baseType basetype.BaseType) bool {
	return accumulate && baseType != basetype.Enum && baseType != basetype.String
}

func (d *Decoder) expandComponents(mesg *proto.Message, containingField *proto.Field, components []proto.Component) {
	bitVal, ok := typeconv.NumericToInt64(containingField.Value)
	if !ok {
		return
	}

	for i := range components {
		component := &components[i]

		componentField := d.factory.CreateField(mesg.Num, component.FieldNum)
		componentField.IsExpandedField = true

		if isEligibleToAccumulate(component.Accumulate, componentField.Type.BaseType()) {
			bitVal = d.accumulator.Accumulate(mesg.Num, component.FieldNum, bitVal, component.Bits)
		}

		var val = bitVal
		if len(components) > 1 {
			if bitVal == 0 {
				break // no more bits to shift
			}
			var mask int64 = (1 << component.Bits) - 1 // e.g. (1 << 8) - 1     = 255
			val = val & mask                           // e.g. 0x27010E08 & 255 = 0x08
			bitVal = bitVal >> component.Bits          // e.g. 0x27010E08 >> 8  = 0x27010E
		}

		// QUESTION: Should we expand componentField.Components too (?)
		componentField.Value = typeconv.Int64ToNumber(val, componentField.Type.BaseType()) // cast back to original value
		mesg.Fields = append(mesg.Fields, componentField)
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
			continue // Can't interpret this DeveloperField, no FieldDescription found.
		}

		developerField := proto.DeveloperField{
			Num:                devFieldDef.Num,
			DeveloperDataIndex: devFieldDef.DeveloperDataIndex,
			Size:               devFieldDef.Size,
			NativeMesgNum:      fieldDescription.NativeMesgNum,
			NativeFieldNum:     fieldDescription.NativeFieldNum,
			Name:               fieldDescription.FieldName,
			Type:               fieldDescription.FitBaseTypeId,
			Units:              fieldDescription.Units,
		}

		val, err := d.readValue(developerField.Size, developerField.Type, mesgDef.Architecture)
		if err != nil {
			return err
		}

		developerField.Value = val

		mesg.DeveloperFields = append(mesg.DeveloperFields, developerField)
	}
	return nil
}

func (d *Decoder) decodeCRC() error {
	b := make([]byte, 2)
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
	_, _ = d.crc16.Write(b)
	return err
}

// readByte is shorthand for read([1]byte).
func (d *Decoder) readByte() (byte, error) {
	b := make([]byte, 1)
	if err := d.read(b); err != nil {
		return 0, err
	}
	return b[0], nil
}

// readValue reads message value bytes from reader and convert it into its corresponding type.
func (d *Decoder) readValue(size byte, baseType basetype.BaseType, arch byte) (any, error) {
	b := make([]byte, size)
	if err := d.read(b); err != nil {
		return nil, err
	}

	val, err := typedef.Unmarshal(b, byteorder.Select(arch), baseType)
	if err != nil {
		return nil, err
	}

	return val, nil
}