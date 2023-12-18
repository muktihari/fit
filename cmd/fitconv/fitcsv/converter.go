// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fitcsv

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/listener"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	_ listener.MesgDefListener = &FitToCsvConv{}
	_ listener.MesgListener    = &FitToCsvConv{}
)

// FitToCsvConv is an implementation for listeners that receive message events and convert them into CSV records.
type FitToCsvConv struct {
	w io.Writer // A writer to write the complete csv data.

	// Temporary writers that can be used to write the csv data (without header). Default: immem.
	// Since we can only write the correct header after all data is retrieved.
	inmem  *bytes.Buffer // Temporary writer in memory
	ondisk *os.File      // Temporary writer on disk

	buf *bytes.Buffer // Buffer for writing bytes
	err error         // Error occurred while receiving messages. It's up to the handler whether to stop or continue.

	options *options

	maxFields int // Since we write messages as they arrive, we don't know the exact number of headers we need, so we create an approximate size.

	developerDataIdMessages []proto.Message
	fieldDesciptionMessages []proto.Message

	eventch chan any      // This buffered event channel can accept either proto.Message or proto.MessageDefinition maintaining the order of arrival.
	done    chan struct{} // Tells that all messages have been completely processed.
}

type Option interface{ apply(o *options) }

type options struct {
	channelBufferSize int
	rawValue          bool // Print scaled value as is in the binary form (sint, uint, etc.) than in its representation.
	unknownNumber     bool // Print 'unknown(68)' instead of 'unknown'.
	useDisk           bool // Write temporary data in disk instead of in memory.
}

func defaultOptions() *options {
	return &options{
		channelBufferSize: 1000, // Ensures the broadcaster isn't blocked even if 1000 events are generated; we should have caught up by then.
		rawValue:          false,
		unknownNumber:     false,
		useDisk:           false,
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

func WithRawValue() Option {
	return fnApply(func(o *options) { o.rawValue = true })
}

func WithUnknownNumber() Option {
	return fnApply(func(o *options) { o.unknownNumber = true })
}

func WithUseDisk() Option {
	return fnApply(func(o *options) { o.useDisk = true })
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
		inmem:   bytes.NewBuffer(nil),
		buf:     new(bytes.Buffer),
		options: options,
		eventch: make(chan any, options.channelBufferSize),
		done:    make(chan struct{}),
	}

	if options.useDisk {
		c.ondisk, c.err = os.CreateTemp(".", "fitconv-fit-to-csv-temp-file")
	}

	go c.handleEvent() // spawn only once.

	return c
}

// Err returns any error that occur during processing events.
func (c *FitToCsvConv) Err() error { return c.err }

// OnMesgDef receive message definition from broadcaster
func (c *FitToCsvConv) OnMesgDef(mesgDef proto.MessageDefinition) { c.eventch <- mesgDef }

// OnMesgDef receive message from broadcaster
func (c *FitToCsvConv) OnMesg(mesg proto.Message) { c.eventch <- mesg }

// handleEvent processes events from a buffered channel.
// It should not be concurrently spawned multiple times, as it relies on maintaining event order.
func (c *FitToCsvConv) handleEvent() {
	for event := range c.eventch {
		switch data := event.(type) {
		case proto.MessageDefinition:
			c.writeMesgDef(data)
		case proto.Message:
			switch data.Num {
			case mesgnum.DeveloperDataId:
				c.developerDataIdMessages = append(c.developerDataIdMessages, data)
			case mesgnum.FieldDescription:
				c.fieldDesciptionMessages = append(c.fieldDesciptionMessages, data)
			}
			c.writeMesg(data)
		}
	}
	close(c.done)
}

// Wait closes the buffered channel and wait until all event handling is completed and finalize the data.
func (c *FitToCsvConv) Wait() {
	close(c.eventch)
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
		_, c.err = c.ondisk.Seek(0, io.SeekStart)
		if c.err != nil {
			return
		}
		_, c.err = io.Copy(c.w, c.ondisk)
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
		c.buf.WriteString(",Field " + num)
		c.buf.WriteString(",Value " + num)
		c.buf.WriteString(",Units " + num)
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
	c.buf.WriteString(strconv.Itoa(int(proto.LocalMesgNum(mesgDef.Header))) + ",")
	mesgName := mesgDef.MesgNum.String()
	if strings.Contains(strings.ToLower(mesgName), "invalid") {
		mesgName = factory.NameUnknown
		if c.options.unknownNumber {
			mesgName = formatUnknown(int(mesgDef.MesgNum))
		}
	}
	c.buf.WriteString(mesgName)
	c.buf.WriteRune(',')

	for i := range mesgDef.FieldDefinitions {
		fieldDef := mesgDef.FieldDefinitions[i]
		field := factory.CreateField(mesgDef.MesgNum, fieldDef.Num)
		name := field.Name
		if c.options.unknownNumber && name == factory.NameUnknown {
			name = formatUnknown(int(field.Num))
		}
		c.buf.WriteString(name + ",")
		c.buf.WriteString(strconv.Itoa(int(fieldDef.Size/fieldDef.BaseType.Size())) + ",")
		c.buf.WriteString(",")
	}

	for i := range mesgDef.DeveloperFieldDefinitions {
		devFieldDef := &mesgDef.DeveloperFieldDefinitions[i]
		c.buf.WriteString(c.devFieldName(devFieldDef) + ",")
		c.buf.WriteString(strconv.Itoa(int(devFieldDef.Size)) + ",")
		c.buf.WriteString(",")
	}

	size := len(mesgDef.FieldDefinitions) + len(mesgDef.DeveloperFieldDefinitions)
	if size > c.maxFields {
		c.maxFields = size
	}

	c.buf.WriteByte('\n') // line break

	if c.options.useDisk {
		_, c.err = c.ondisk.Write(c.buf.Bytes())
	} else {
		_, c.err = c.inmem.Write(c.buf.Bytes())
	}
	c.buf.Reset()
}

func (c *FitToCsvConv) devFieldName(devFieldDef *proto.DeveloperFieldDefinition) string {
	for i := range c.fieldDesciptionMessages {
		fieldDescMesg := &c.fieldDesciptionMessages[i]
		devDataIndex := fieldDescMesg.FieldByNum(fieldnum.FieldDescriptionDeveloperDataIndex)
		if devDataIndex == nil {
			continue
		}
		fieldDefNum := fieldDescMesg.FieldByNum(fieldnum.FieldDescriptionFieldDefinitionNumber)
		if fieldDefNum == nil {
			continue
		}

		if typeconv.ToByte[byte](devDataIndex.Value) == devFieldDef.DeveloperDataIndex &&
			typeconv.ToByte[byte](fieldDefNum.Value) == devFieldDef.Num {
			fieldName := fieldDescMesg.FieldByNum(fieldnum.FieldDescriptionFieldName)
			if fieldName == nil {
				break
			}
			return typeconv.ToString[string](fieldName.Value)
		}
	}

	if c.options.unknownNumber {
		return formatUnknown(int(devFieldDef.Num))
	}

	return factory.NameUnknown
}

func (c *FitToCsvConv) writeMesg(mesg proto.Message) {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Data,")
	c.buf.WriteString(strconv.Itoa(int(proto.LocalMesgNum(mesg.Header))) + ",")
	mesgName := mesg.Num.String()
	if strings.Contains(strings.ToLower(mesgName), "invalid") {
		mesgName = factory.NameUnknown
		if c.options.unknownNumber {
			mesgName = formatUnknown(int(mesg.Num))
		}
	}
	c.buf.WriteString(mesgName)
	c.buf.WriteRune(',')

	var fieldCounter int
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		name, units := field.Name, field.Units
		if c.options.unknownNumber && field.Name == factory.NameUnknown {
			name = formatUnknown(int(field.Num))
		}
		if subField, ok := field.SubFieldSubtitution(&mesg); ok {
			name, units = subField.Name, subField.Units
		}

		if vals, ok := sliceAny(field.Value); ok { // array
			c.buf.WriteString(name)
			c.buf.WriteByte(',')

			c.buf.WriteByte('"')
			for i := range vals {
				value := vals[i]
				if !c.options.rawValue {
					value = scaleoffset.ApplyAny(vals[i], field.Scale, field.Offset)
				}
				c.buf.WriteString(format(value))
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

		c.buf.WriteString(name)
		c.buf.WriteByte(',')

		value := field.Value
		if !c.options.rawValue {
			value = scaleoffset.ApplyAny(field.Value, field.Scale, field.Offset)
		}
		c.buf.WriteByte('"')
		c.buf.WriteString(format(value))
		c.buf.WriteString("\",")

		c.buf.WriteString(units)
		c.buf.WriteByte(',')
		fieldCounter++
	}

	for i := range mesg.DeveloperFields {
		devField := &mesg.DeveloperFields[i]

		c.buf.WriteString(devField.Name)
		c.buf.WriteByte(',')
		c.buf.WriteString(format(devField.Value))
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
		_, c.err = c.ondisk.Write(c.buf.Bytes())
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
	}

	return nil, false
}
