// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fitcsv

import (
	"bufio"
	"bytes"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	_ decoder.MesgDefListener = &FitToCsvConv{}
	_ decoder.MesgListener    = &FitToCsvConv{}
)

// FitToCsvConv is an implementation for listeners that receive message events and convert them into CSV records.
type FitToCsvConv struct {
	w io.Writer // A writer to write the complete csv data.

	// Temporary writers that can be used to write the csv data (without header). Default: immem.
	// Since we can only write the correct header after all data is retrieved.
	inmem              *bytes.Buffer // Temporary writer in memory
	ondisk             *os.File      // Temporary writer on disk
	ondiskBufferWriter *bufio.Writer // Ondisk buffer writer to reduce syscall on write

	buf *bytes.Buffer // Buffer for writing bytes
	err error         // Error occurred while receiving messages. It's up to the handler whether to stop or continue.

	options *options

	maxFields int // Since we write messages as they arrive, we don't know the exact number of headers we need, so we create an approximate size.

	developerDataIds  []*mesgdef.DeveloperDataId
	fieldDescriptions []*mesgdef.FieldDescription

	mesgc chan any      // This buffered event channel can accept either proto.Message or proto.MessageDefinition maintaining the order of arrival.
	done  chan struct{} // Tells that all messages have been completely processed.
}

type Option interface{ apply(o *options) }

type options struct {
	channelBufferSize         int
	useDisk                   bool // Write temporary data in disk instead of in memory.
	ondiskWriteBuffer         int  // Buffer on read/write in disk when useDisk is specified.
	printRawValue             bool // Print scaled value as is in the binary form (sint, uint, etc.) than in its representation.
	printUnknownMesgNum       bool // Print 'unknown(68)' instead of 'unknown'.
	printOnlyValidValue       bool // Print invalid value e.g. 65535 for uint16.
	printSemicirclesInDegrees bool // Print latitude and longitude in degrees instead of semicircles.
}

func defaultOptions() *options {
	return &options{
		channelBufferSize:         1000, // Ensures the broadcaster isn't blocked even if 1000 events are generated; we should have caught up by then.
		useDisk:                   false,
		ondiskWriteBuffer:         4 << 10,
		printRawValue:             false,
		printUnknownMesgNum:       false,
		printOnlyValidValue:       false,
		printSemicirclesInDegrees: false,
	}
}

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

func WithChannelBufferSize(size int) Option {
	return fnApply(func(o *options) {
		if size > 0 {
			o.channelBufferSize = size
		}
	})
}

func WithPrintRawValue() Option {
	return fnApply(func(o *options) { o.printRawValue = true })
}

func WithPrintUnknownMesgNum() Option {
	return fnApply(func(o *options) { o.printUnknownMesgNum = true })
}

func WithUseDisk(writeBuffer int) Option {
	return fnApply(func(o *options) {
		o.useDisk = true
		if writeBuffer > 0 {
			o.ondiskWriteBuffer = writeBuffer
		}
	})
}

func WithPrintOnlyValidValue() Option {
	return fnApply(func(o *options) {
		o.printOnlyValidValue = true
	})
}

func WithPrintSemicirclesInDegrees() Option {
	return fnApply(func(o *options) {
		o.printSemicirclesInDegrees = true
	})
}

// NewConverter creates a new fit to csv converter.
// The caller must call Wait() to wait all events are received and finalizing the convert process.
func NewConverter(w io.Writer, opts ...Option) *FitToCsvConv {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	c := &FitToCsvConv{
		w:       w,
		buf:     new(bytes.Buffer),
		options: options,
		mesgc:   make(chan any, options.channelBufferSize),
		done:    make(chan struct{}),
	}

	if options.useDisk {
		c.ondisk, c.err = os.CreateTemp(".", "fitconv-fit-to-csv-temp-file")
		if c.err != nil {
			return c
		}
		c.ondiskBufferWriter = bufio.NewWriterSize(c.ondisk, options.ondiskWriteBuffer)
	} else {
		c.inmem = bytes.NewBuffer(nil)
	}

	go c.handleEvent() // spawn only once.

	return c
}

// Err returns any error that occur during processing events.
func (c *FitToCsvConv) Err() error { return c.err }

// OnMesgDef receive message definition from broadcaster
func (c *FitToCsvConv) OnMesgDef(mesgDef proto.MessageDefinition) { c.mesgc <- mesgDef }

// OnMesgDef receive message from broadcaster
func (c *FitToCsvConv) OnMesg(mesg proto.Message) { c.mesgc <- mesg }

// handleEvent processes events from a buffered channel.
// It should not be concurrently spawned multiple times, as it relies on maintaining event order.
func (c *FitToCsvConv) handleEvent() {
	for event := range c.mesgc {
		switch mesg := event.(type) {
		case proto.MessageDefinition:
			c.writeMesgDef(mesg)
		case proto.Message:
			switch mesg.Num {
			case mesgnum.DeveloperDataId:
				c.developerDataIds = append(c.developerDataIds, mesgdef.NewDeveloperDataId(&mesg))
			case mesgnum.FieldDescription:
				c.fieldDescriptions = append(c.fieldDescriptions, mesgdef.NewFieldDescription(&mesg))
			}
			c.writeMesg(mesg)
		}
	}
	close(c.done)
}

// Wait closes the buffered channel and wait until all event handling is completed and finalize the data.
func (c *FitToCsvConv) Wait() {
	close(c.mesgc)
	<-c.done
	c.finalize()
}

func (c *FitToCsvConv) removeTemporaryFile() {
	if c.ondisk == nil {
		return
	}
	name := c.ondisk.Name()
	c.ondisk.Close()
	os.Remove(name)
}

func (c *FitToCsvConv) finalize() {
	defer c.removeTemporaryFile()

	c.printHeader()
	if c.err != nil {
		return
	}

	if c.options.useDisk {
		c.err = c.ondiskBufferWriter.Flush() // flush remaining buffer
		if c.err != nil {
			return
		}
		_, c.err = c.ondisk.Seek(0, io.SeekStart)
		if c.err != nil {
			return
		}
		_, c.err = io.Copy(c.w, c.ondisk) // do not wrap with bufio.Reader
	} else {
		_, c.err = io.Copy(c.w, c.inmem)
	}
}

func (c *FitToCsvConv) printHeader() {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Type,Local Number,Message")
	for i := 0; i < c.maxFields*3; i += 3 {
		num := strconv.Itoa((i / 3) + 1)
		c.buf.WriteString(",Field ")
		c.buf.WriteString(num)

		c.buf.WriteString(",Value ")
		c.buf.WriteString(num)

		c.buf.WriteString(",Units ")
		c.buf.WriteString(num)
	}
	c.buf.WriteByte('\n') // line break

	_, c.err = c.w.Write(c.buf.Bytes())
	c.buf.Reset()
}

func formatUnknown(num int) string {
	return "unknown(" + strconv.Itoa(num) + ")"
}

func (c *FitToCsvConv) writeMesgDef(mesgDef proto.MessageDefinition) {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Definition,")
	c.buf.WriteString(strconv.Itoa(int(proto.LocalMesgNum(mesgDef.Header))))
	c.buf.WriteByte(',')

	mesgName := mesgDef.MesgNum.String()
	if strings.HasPrefix(mesgName, "MesgNumInvalid") {
		mesgName = factory.NameUnknown
		if c.options.printUnknownMesgNum {
			mesgName = formatUnknown(int(mesgDef.MesgNum))
		}
	}
	c.buf.WriteString(mesgName)
	c.buf.WriteByte(',')

	for i := range mesgDef.FieldDefinitions {
		fieldDef := mesgDef.FieldDefinitions[i]
		field := factory.CreateField(mesgDef.MesgNum, fieldDef.Num)

		name := field.Name
		if c.options.printUnknownMesgNum && name == factory.NameUnknown {
			name = formatUnknown(int(field.Num))
		}

		c.buf.WriteString(name)
		c.buf.WriteByte(',')

		c.buf.WriteString(strconv.Itoa(int(fieldDef.Size / fieldDef.BaseType.Size())))
		c.buf.WriteByte(',')

		// empty column
		c.buf.WriteByte(',')
	}

	for i := range mesgDef.DeveloperFieldDefinitions {
		devFieldDef := &mesgDef.DeveloperFieldDefinitions[i]

		c.buf.WriteString(c.devFieldName(devFieldDef))
		c.buf.WriteByte(',')

		c.buf.WriteString(strconv.Itoa(int(devFieldDef.Size)))
		c.buf.WriteByte(',')

		// empty column
		c.buf.WriteByte(',')
	}

	size := len(mesgDef.FieldDefinitions) + len(mesgDef.DeveloperFieldDefinitions)
	if size > c.maxFields {
		c.maxFields = size
	}

	c.buf.WriteByte('\n') // line break

	if c.options.useDisk {
		_, c.err = c.ondiskBufferWriter.Write(c.buf.Bytes())
	} else {
		_, c.err = c.inmem.Write(c.buf.Bytes())
	}
	c.buf.Reset()
}

func (c *FitToCsvConv) devFieldName(devFieldDef *proto.DeveloperFieldDefinition) string {
	var fieldDescription *mesgdef.FieldDescription
	for i := range c.fieldDescriptions {
		fieldDef := c.fieldDescriptions[i]
		if fieldDef.DeveloperDataIndex != devFieldDef.DeveloperDataIndex {
			continue
		}
		if fieldDef.FieldDefinitionNumber != devFieldDef.Num {
			continue
		}
		fieldDescription = fieldDef
		break
	}
	if fieldDescription != nil {
		return strings.Join(fieldDescription.FieldName, "|")
	}
	if c.options.printUnknownMesgNum {
		return formatUnknown(int(devFieldDef.Num))
	}
	return factory.NameUnknown
}

func (c *FitToCsvConv) writeMesg(mesg proto.Message) {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Data,")
	c.buf.WriteString(strconv.Itoa(int(proto.LocalMesgNum(mesg.Header))))
	c.buf.WriteByte(',')
	mesgName := mesg.Num.String()
	if strings.HasPrefix(mesgName, "MesgNumInvalid") {
		mesgName = factory.NameUnknown
		if c.options.printUnknownMesgNum {
			mesgName = formatUnknown(int(mesg.Num))
		}
	}
	c.buf.WriteString(mesgName)
	c.buf.WriteByte(',')

	var fieldCounter int
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		name, units := field.Name, field.Units
		if c.options.printUnknownMesgNum && field.Name == factory.NameUnknown {
			name = formatUnknown(int(field.Num))
		}
		if subField := field.SubFieldSubtitution(&mesg); subField != nil {
			name, units = subField.Name, subField.Units
		}

		var value proto.Value
		if !c.options.printRawValue {
			value = scaleoffset.ApplyValue(field.Value, field.Scale, field.Offset)
		}

		if vals, ok := sliceAny(value.Any()); ok { // array
			c.buf.WriteString(name)
			c.buf.WriteByte(',')

			c.buf.WriteByte('"')

			for i := range vals {
				c.buf.WriteString(format(vals[i]))
				if i < len(vals)-1 {
					c.buf.WriteByte('|')
				}
			}
			c.buf.WriteString("\",")

			c.buf.WriteString(units)
			c.buf.WriteByte(',')
			fieldCounter++
			continue
		}

		if c.options.printOnlyValidValue && !isValueValid(field) {
			continue
		}

		c.buf.WriteString(name)
		c.buf.WriteByte(',')

		scaledValue := field.Value
		if !c.options.printRawValue {
			scaledValue = scaleoffset.ApplyValue(field.Value, field.Scale, field.Offset)
		}

		if c.options.printSemicirclesInDegrees && field.Units == "semicircles" {
			scaledValue = proto.Float64(semicircles.ToDegrees(scaledValue.Int32()))
			units = "degrees"
		}

		c.buf.WriteByte('"')
		c.buf.WriteString(format(scaledValue.Any()))
		c.buf.WriteString("\",")

		c.buf.WriteString(units)
		c.buf.WriteByte(',')
		fieldCounter++
	}

	for i := range mesg.DeveloperFields {
		devField := &mesg.DeveloperFields[i]

		c.buf.WriteString(devField.Name)
		c.buf.WriteByte(',')

		if vals, ok := sliceAny(devField.Value.Any()); ok { // array
			for i := range vals {
				c.buf.WriteString(format(vals[i]))
				if i < len(vals)-1 {
					c.buf.WriteByte('|')
				}
			}
		} else {
			c.buf.WriteString(format(devField.Value.Any()))
		}

		c.buf.WriteByte(',')
		c.buf.WriteString(devField.Units)
		c.buf.WriteByte(',')
		fieldCounter++
	}

	if fieldCounter > c.maxFields {
		c.maxFields = fieldCounter
	}

	c.buf.WriteByte('\n') // line break

	if c.options.useDisk {
		_, c.err = c.ondiskBufferWriter.Write(c.buf.Bytes())
	} else {
		_, c.err = c.inmem.Write(c.buf.Bytes())
	}
	c.buf.Reset()
}

func sliceAny(val any) (vals []any, isSlice bool) {
	isSlice = true
	switch vs := val.(type) {
	case []int8:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []uint8:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []int16:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []uint16:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []int32:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []uint32:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []int64:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []uint64:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []float32:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []float64:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	case []string:
		for i := range vs {
			vals = append(vals, vs[i])
		}
		return
	}

	return nil, false
}

func isValueValid(field *proto.Field) bool {
	switch field.BaseType {
	case basetype.Float32:
		f32 := field.Value.Float32()
		if math.Float32bits(f32) == basetype.Float32Invalid {
			return false
		}
	case basetype.Float64:
		f64 := field.Value.Float64()
		if math.Float64bits(f64) == basetype.Float64Invalid {
			return false
		}
	default:
		if field.Value.Type() == proto.TypeInvalid {
			return false
		}
	}
	return true
}
