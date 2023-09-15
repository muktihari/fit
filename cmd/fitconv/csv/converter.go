// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"bytes"
	"io"
	"math"
	"strconv"
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/listener"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

const (
	defaultHeaderSize        = 40   // Header size estimation
	defaultChannelBufferSize = 1000 // Ensures the broadcaster isn't blocked event if 1000 events are generated; we should have caught up by then.
)

var (
	_ listener.MesgDefListener = &Conv{}
	_ listener.MesgListener    = &Conv{}
)

// Conv is an implementation for listeners that receive message events and convert them into CSV records.
type Conv struct {
	w    io.Writer     // A writer to write to.
	buf  *bytes.Buffer // Buffer for writing bytes
	once sync.Once     // Do something exactly once. Our use case is printing the header on the first received message definition.
	err  error         // Error occurred while receiving messages. It's up to the handler whether to stop or continue.

	headerSize       int  // Since we write messages as they arrive, we don't know the exact number of headers we need, so we create an approximate size.
	debugUseRawValue bool // Print scaled value as is in the binary form (sint, uint, etc.) than in its representation.

	developerDataIdMessages []proto.Message
	fieldDesciptionMessages []proto.Message

	eventch chan any      // This buffered event channel can accept either proto.Message or proto.MessageDefinition maintaining the order of arrival.
	done    chan struct{} // Tells that all messages have been completely processed.
}

// NewConverter is a shorthand for NewConvWithOptions(w, defaultHeaderSize, defaultChannelBufferSize).
func NewConverter(w io.Writer) *Conv {
	return NewConverterWithOptions(w, defaultHeaderSize, defaultChannelBufferSize)
}

// NewConverterWithOptions creates a new csv converter *Conv with customizable header and buffer sizes.
// The header size is estimated since we write it as events arrive, and precision isn't required.
// The channel buffer size is used to pool incoming events, as we can't predict the event rate or processing speed,
// ensuring the broadcaster is not blocked.
//
// The caller must call CloseAndWait() after the broadcasting of events is complete to ensure all buffered events are processed.
func NewConverterWithOptions(w io.Writer, headerSize, channelBufferSize int) *Conv {
	conv := &Conv{
		w:                w,
		buf:              new(bytes.Buffer),
		headerSize:       headerSize,
		debugUseRawValue: false, // Useful for debugging.
		eventch:          make(chan any, channelBufferSize),
		done:             make(chan struct{}),
	}
	go conv.handleEvent() // spawn only once.
	return conv
}

// Err returns any error that occur during processing events.
func (c *Conv) Err() error { return c.err }

// OnMesgDef receive message definition from broadcaster
func (c *Conv) OnMesgDef(mesgDef proto.MessageDefinition) { c.eventch <- mesgDef }

// OnMesgDef receive message from broadcaster
func (c *Conv) OnMesg(mesg proto.Message) { c.eventch <- mesg }

// handleEvent processes events from a buffered channel.
// It should not be concurrently spawned multiple times, as it relies on maintaining event order.
func (c *Conv) handleEvent() {
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

// Wait closes the buffered channel and wait until all event handling is completed.
func (c *Conv) Wait() {
	close(c.eventch)
	<-c.done
}

func (c *Conv) printHeader() {
	c.buf.WriteString("Type,Local Number,Message")
	for i := 0; i < c.headerSize*3; i += 3 {
		num := strconv.Itoa((i / 3) + 1)
		c.buf.WriteString(",Field " + num)
		c.buf.WriteString(",Value " + num)
		c.buf.WriteString(",Units " + num)
	}
	c.buf.WriteByte('\n') // line break

	_, err := c.w.Write(c.buf.Bytes())
	c.err = err
	c.buf.Reset()
}

func (c *Conv) writeMesgDef(mesgDef proto.MessageDefinition) {
	if c.err != nil {
		return
	}

	c.once.Do(func() { c.printHeader() })

	c.buf.WriteString("Definition,")
	c.buf.WriteString(strconv.Itoa(int(mesgDef.LocalMesgNum)) + ",")
	c.buf.WriteString(mesgDef.MesgNum.String() + ",")

	for i := range mesgDef.FieldDefinitions {
		fieldDef := mesgDef.FieldDefinitions[i]
		field := factory.CreateField(mesgDef.MesgNum, fieldDef.Num)
		c.buf.WriteString(field.Name + ",")
		c.buf.WriteString(strconv.Itoa(int(fieldDef.Size/fieldDef.BaseType.Size())) + ",")
		c.buf.WriteString(",")
	}

	for i := range mesgDef.DeveloperFieldDefinitions {
		devFieldDef := &mesgDef.DeveloperFieldDefinitions[i]
		c.buf.WriteString(c.devFieldName(devFieldDef) + ",")
		c.buf.WriteString(strconv.Itoa(int(devFieldDef.Size)) + ",")
		c.buf.WriteString(",")
	}

	c.buf.WriteByte('\n') // line break

	_, err := c.w.Write(c.buf.Bytes())
	c.err = err
	c.buf.Reset()
}

func (c *Conv) devFieldName(devFieldDef *proto.DeveloperFieldDefinition) string {
	for i := range c.fieldDesciptionMessages {
		fieldDescMesg := &c.fieldDesciptionMessages[i]
		devDataIndex, ok := fieldDescMesg.FieldByNum(fieldnum.FieldDescriptionDeveloperDataIndex)
		if !ok {
			continue
		}
		fieldDefNum, ok := fieldDescMesg.FieldByNum(fieldnum.FieldDescriptionFieldDefinitionNumber)
		if !ok {
			continue
		}

		if typeconv.ToByte[byte](devDataIndex.Value) == devFieldDef.DeveloperDataIndex &&
			typeconv.ToByte[byte](fieldDefNum.Value) == devFieldDef.Num {
			fieldName, ok := fieldDescMesg.FieldByNum(fieldnum.FieldDescriptionFieldName)
			if !ok {
				break
			}
			return typeconv.ToString[string](fieldName.Value)
		}
	}

	return factory.NameUnknown
}

func (c *Conv) writeMesg(mesg proto.Message) {
	if c.err != nil {
		return
	}

	c.buf.WriteString("Data,")
	c.buf.WriteString(strconv.Itoa(int(mesg.LocalNum)) + ",")
	c.buf.WriteString(mesg.Num.String() + ",")

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		name, units := field.Name, field.Units
		if subField, ok := field.SubFieldSubtitution(&mesg); ok {
			name, units = subField.Name, subField.Units
		}

		if vals, ok := sliceAny(field.Value); ok { // array
			if vals == nil {
				continue // skip invalid values
			}

			c.buf.WriteString(name)
			c.buf.WriteByte(',')

			c.buf.WriteByte('"')
			for i := 0; i < len(vals); i++ {
				value := vals[i]
				if !c.debugUseRawValue {
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
			continue
		}

		if field.Value == field.Type.BaseType().Invalid() {
			continue
		}

		c.buf.WriteString(name)
		c.buf.WriteByte(',')

		value := field.Value
		if !c.debugUseRawValue {
			value = scaleoffset.ApplyAny(field.Value, field.Scale, field.Offset)
		}
		c.buf.WriteByte('"')
		c.buf.WriteString(format(value))
		c.buf.WriteString("\",")

		c.buf.WriteString(units)
		c.buf.WriteByte(',')
	}

	for i := range mesg.DeveloperFields {
		devField := &mesg.DeveloperFields[i]

		c.buf.WriteString(devField.Name)
		c.buf.WriteByte(',')
		c.buf.WriteString(format(devField.Value))
		c.buf.WriteByte(',')
		c.buf.WriteString(devField.Units)
		c.buf.WriteByte(',')
	}

	c.buf.WriteByte('\n') // line break

	_, err := c.w.Write(c.buf.Bytes())
	c.err = err
	c.buf.Reset()
}

func sliceAny(val any) (vals []any, isSlice bool) {
	isSlice = true
	switch vs := val.(type) {
	case []int8:
		for i := 0; i < len(vs); i++ {
			if vs[i] == basetype.Sint8.Invalid() {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []uint8:
		for i := 0; i < len(vs); i++ {
			if vs[i] == basetype.Uint8.Invalid() {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []int16:
		for i := 0; i < len(vs); i++ {
			if vs[i] == basetype.Sint16.Invalid() {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []uint16:
		for i := 0; i < len(vs); i++ {
			if vs[i] == basetype.Uint16.Invalid() {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []int32:
		for i := 0; i < len(vs); i++ {
			if vs[i] == basetype.Sint32.Invalid() {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []uint32:
		for i := 0; i < len(vs); i++ {
			if vs[i] == basetype.Uint32.Invalid() {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []float32:
		for i := 0; i < len(vs); i++ {
			if math.IsNaN(float64(vs[i])) {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	case []float64:
		for i := 0; i < len(vs); i++ {
			if math.IsNaN(float64(vs[i])) {
				continue
			}
			vals = append(vals, vs[i])
		}
		return
	}

	return nil, false
}
