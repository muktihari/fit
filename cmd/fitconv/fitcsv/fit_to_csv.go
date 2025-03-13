// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fitcsv

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	_ decoder.MesgDefListener = &FITToCSVConv{}
	_ decoder.MesgListener    = &FITToCSVConv{}
)

// FITToCSVConv is an implementation for listeners that receive message events and convert them into CSV records.
type FITToCSVConv struct {
	iter int
	w    io.Writer // A writer to write the complete csv data.

	// Temporary writers that can be used to write the csv data (without header). Default: immem.
	// Since we can only write the correct header after all data is retrieved.
	inmem              *bytes.Buffer // Temporary writer in memory
	ondisk             *os.File      // Temporary writer on disk
	ondiskBufferWriter *bufio.Writer // Ondisk buffer writer to reduce syscall on write

	buf *bytes.Buffer // Buffer for writing bytes
	err error         // Error occurred while receiving messages. It's up to the handler whether to stop or continue.

	options *options

	maxFields int // Since we write messages as they arrive, we don't know the exact number of headers we need, so we create an approximate size.

	fieldDescriptions []*mesgdef.FieldDescription // Key: uint16(DeveloperDataIndex) << 8 | uint16(FieldDefinitionNumber)

	mesgc chan any      // This buffered event channel can accept either proto.Message or proto.MessageDefinition maintaining the order of arrival.
	done  chan struct{} // Tells that all messages have been completely processed.
}

type options struct {
	channelBufferSize         int
	ondiskWriteBuffer         int  // Buffer on read/write in disk when useDisk is specified.
	useDisk                   bool // Write temporary data in disk instead of in memory.
	printRawValue             bool // Print scaled value as is in the binary form (sint, uint, etc.) than in its representation.
	verbose                   bool // Print 'unknown(68)' instead of 'unknown'.
	printOnlyValidValue       bool // Print invalid value e.g. 65535 for uint16.
	printGPSPositionInDegrees bool // Print latitude and longitude in degrees instead of semicircles.
	trimTrailingCommas        bool
}

func defaultOptions() *options {
	return &options{
		channelBufferSize:         1000, // Ensures the broadcaster isn't blocked even if 1000 events are generated; we should have caught up by then.
		useDisk:                   false,
		ondiskWriteBuffer:         4 << 10,
		printRawValue:             false,
		verbose:                   false,
		printOnlyValidValue:       false,
		printGPSPositionInDegrees: false,
		trimTrailingCommas:        false,
	}
}

// Option is FITToCSVConv's option.
type Option func(o *options)

func WithChannelBufferSize(size int) Option {
	return func(o *options) {
		if size > 0 {
			o.channelBufferSize = size
		}
	}
}

func WithPrintRawValue() Option {
	return func(o *options) { o.printRawValue = true }
}

func WithPrintUnknownMesgNum() Option {
	return func(o *options) { o.verbose = true }
}

func WithUseDisk(writeBuffer int) Option {
	return func(o *options) {
		o.useDisk = true
		if writeBuffer > 0 {
			o.ondiskWriteBuffer = writeBuffer
		}
	}
}

func WithPrintOnlyValidValue() Option {
	return func(o *options) {
		o.printOnlyValidValue = true
	}
}

func WithPrintGPSPositionInDegrees() Option {
	return func(o *options) {
		o.printGPSPositionInDegrees = true
	}
}

func WithTrimTrailingCommas() Option {
	return func(o *options) {
		o.trimTrailingCommas = true
	}
}

// NewFITToCSVConv creates a new FIT to CSV converter.
// The caller must call Wait() to wait all events are received and finalizing the convert process.
func NewFITToCSVConv(w io.Writer, opts ...Option) *FITToCSVConv {
	options := defaultOptions()
	for i := range opts {
		opts[i](options)
	}

	c := &FITToCSVConv{
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
		c.inmem = bytes.NewBuffer(make([]byte, 0, 1000<<10))
	}

	go c.handleEvent() // spawn only once.

	return c
}

// Err returns any error that occur during processing events.
func (c *FITToCSVConv) Err() error { return c.err }

// OnMesgDef receive message definition from broadcaster
func (c *FITToCSVConv) OnMesgDef(mesgDef proto.MessageDefinition) { c.mesgc <- mesgDef }

// OnMesgDef receive message from broadcaster
func (c *FITToCSVConv) OnMesg(mesg proto.Message) { c.mesgc <- mesg }

// handleEvent processes events from a buffered channel.
// It should not be concurrently spawned multiple times, as it relies on maintaining event order.
func (c *FITToCSVConv) handleEvent() {
	for event := range c.mesgc {
		switch mesg := event.(type) {
		case proto.MessageDefinition:
			c.iter++
			c.writeMesgDef(mesg)
		case proto.Message:
			switch mesg.Num {
			case mesgnum.FieldDescription:
				c.fieldDescriptions = append(c.fieldDescriptions, mesgdef.NewFieldDescription(&mesg))
			}
			c.writeMesg(mesg)
		}
	}
	close(c.done)
}

// Wait closes the buffered channel and wait until all event handling is completed and finalize the data.
func (c *FITToCSVConv) Wait() {
	close(c.mesgc)
	<-c.done
	c.finalize()
}

func (c *FITToCSVConv) removeTemporaryFile() {
	if c.ondisk == nil {
		return
	}
	name := c.ondisk.Name()
	c.ondisk.Close()
	os.Remove(name)
}

func (c *FITToCSVConv) finalize() {
	defer c.removeTemporaryFile()

	c.printHeader()
	if c.err != nil {
		return
	}

	maxCommaCount := 2 + (c.maxFields * 3)

	if c.options.useDisk {
		c.err = c.ondiskBufferWriter.Flush() // flush remaining buffer
		if c.err != nil {
			return
		}
		_, c.err = c.ondisk.Seek(0, io.SeekStart)
		if c.err != nil {
			return
		}
		c.err = c.copy(c.w, c.ondisk, maxCommaCount) // do not wrap with bufio.Reader
	} else {
		c.err = c.copy(c.w, c.inmem, maxCommaCount)
	}
}

// copy reads src; for every read line, fill the missing comma before write into dest.
func (c *FITToCSVConv) copy(dest io.Writer, src io.Reader, maxCommaCount int) error {
	if c.options.trimTrailingCommas { //By default, we only populate a number of commas just right for field values.
		_, err := io.Copy(dest, src)
		return err
	}

	// Fill Padding Comma ','
	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		b := scanner.Bytes()
		_, err := dest.Write(b)
		if err != nil {
			return err
		}

		var count int
		var inQuote bool
		for _, c := range b {
			switch c {
			case '"':
				inQuote = !inQuote
			case ',':
				if !inQuote {
					count++
				}
			}
		}

		missing := maxCommaCount - count
		padding := bytes.Repeat([]byte{','}, missing)
		_, err = dest.Write(padding)
		if err != nil {
			return err
		}

		_, err = dest.Write([]byte{'\n'})
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *FITToCSVConv) printHeader() {
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

func (c *FITToCSVConv) writeMesgDef(mesgDef proto.MessageDefinition) {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Definition,")
	c.buf.WriteString(strconv.Itoa(int(proto.LocalMesgNum(mesgDef.Header))))
	c.buf.WriteByte(',')

	mesgName := mesgDef.MesgNum.String()
	if strings.HasPrefix(mesgName, "MesgNumInvalid") {
		mesgName = factory.NameUnknown
		if c.options.verbose {
			mesgName = formatUnknown(int(mesgDef.MesgNum))
		}
	}
	c.buf.WriteString(mesgName)
	c.buf.WriteByte(',')

	for i := range mesgDef.FieldDefinitions {
		fieldDef := mesgDef.FieldDefinitions[i]
		field := factory.CreateField(mesgDef.MesgNum, fieldDef.Num)

		name := field.Name
		if c.options.verbose && name == factory.NameUnknown {
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

		name := factory.NameUnknown
		fieldDesc := c.getFieldDescription(devFieldDef.DeveloperDataIndex, devFieldDef.Num)
		if fieldDesc != nil {
			name = strings.Join(fieldDesc.FieldName, "|")
		} else if c.options.verbose {
			name = formatUnknown(int(devFieldDef.Num))
		}
		c.buf.WriteString(name)
		c.buf.WriteByte(',')

		c.buf.WriteString(strconv.Itoa(int(devFieldDef.Size)))
		c.buf.WriteByte(',')

		// empty column
		c.buf.WriteByte(',')
	}

	c.buf.Truncate(c.buf.Len() - 1) // trim last ','

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

func (c *FITToCSVConv) writeMesg(mesg proto.Message) {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Data,")
	c.buf.WriteString(strconv.Itoa(int(proto.LocalMesgNum(mesg.Header))))
	c.buf.WriteByte(',')
	mesgName := mesg.Num.String()
	if strings.HasPrefix(mesgName, "MesgNumInvalid") {
		mesgName = factory.NameUnknown
		if c.options.verbose {
			mesgName = formatUnknown(int(mesg.Num))
		}
	}
	c.buf.WriteString(mesgName)
	c.buf.WriteByte(',')

	var fieldCounter int
	for i := range mesg.Fields {
		field := &mesg.Fields[i]

		value := field.Value
		if c.options.printOnlyValidValue && !field.Value.Valid(field.BaseType) {
			continue
		}

		name, units := field.Name, field.Units
		if c.options.verbose && field.Name == factory.NameUnknown {
			name = formatUnknown(int(field.Num))
			units = field.BaseType.String()
		}
		if subField := field.SubFieldSubtitution(&mesg); subField != nil {
			name, units = subField.Name, subField.Units
		}

		c.buf.WriteString(name)

		if !c.options.printRawValue {
			value = scaleoffset.ApplyValue(field.Value, field.Scale, field.Offset)
		}

		if c.options.printGPSPositionInDegrees && field.Units == "semicircles" {
			value = proto.Float64(semicircles.ToDegrees(value.Int32()))
			units = "degrees"
		}

		c.buf.WriteString(",\"")
		c.buf.WriteString(format(value))
		c.buf.WriteString("\",")

		c.buf.WriteString(units)
		c.buf.WriteByte(',')

		fieldCounter++
	}

	for i := range mesg.DeveloperFields {
		devField := &mesg.DeveloperFields[i]

		var units string
		name := factory.NameUnknown
		fieldDesc := c.getFieldDescription(devField.DeveloperDataIndex, devField.Num)
		if fieldDesc != nil {
			name = strings.Join(fieldDesc.FieldName, "|")
			units = strings.Join(fieldDesc.Units, "|")
		} else if c.options.verbose {
			name = formatUnknown(int(devField.Num))
		}

		c.buf.WriteString(name)

		c.buf.WriteString(",\"")
		c.buf.WriteString(format(devField.Value))
		c.buf.WriteString("\",")

		c.buf.WriteString(units)
		c.buf.WriteByte(',')

		fieldCounter++
	}

	if fieldCounter > c.maxFields {
		c.maxFields = fieldCounter
	}

	c.buf.Truncate(c.buf.Len() - 1) // trim last ','

	c.buf.WriteByte('\n') // line break

	if c.options.useDisk {
		_, c.err = c.ondiskBufferWriter.Write(c.buf.Bytes())
	} else {
		_, c.err = c.inmem.Write(c.buf.Bytes())
	}
	c.buf.Reset()
}

func (c *FITToCSVConv) getFieldDescription(developerDataIndex, fieldDefinitionNumber uint8) *mesgdef.FieldDescription {
	for _, fieldDesc := range c.fieldDescriptions {
		if fieldDesc.DeveloperDataIndex == developerDataIndex &&
			fieldDesc.FieldDefinitionNumber == fieldDefinitionNumber {
			return fieldDesc
		}
	}
	return nil
}
