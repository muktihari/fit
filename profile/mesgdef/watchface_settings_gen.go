// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// WatchfaceSettings is a WatchfaceSettings message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type WatchfaceSettings struct {
	MessageIndex typedef.MessageIndex
	Mode         typedef.WatchfaceMode
	Layout       byte

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWatchfaceSettings creates new WatchfaceSettings struct based on given mesg.
// If mesg is nil, it will return WatchfaceSettings with all fields being set to its corresponding invalid value.
func NewWatchfaceSettings(mesg *proto.Message) *WatchfaceSettings {
	vals := [255]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &WatchfaceSettings{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Mode:         typedef.WatchfaceMode(vals[0].Uint8()),
		Layout:       vals[1].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WatchfaceSettings into proto.Message. If options is nil, default options will be used.
func (m *WatchfaceSettings) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumWatchfaceSettings}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if byte(m.Mode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Mode))
		fields = append(fields, field)
	}
	if m.Layout != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.Layout)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetLayout returns Dynamic Field interpretation of Layout. Otherwise, returns the original value of Layout.
//
// Based on m.Mode:
//   - name: "digital_layout", value: typedef.DigitalWatchfaceLayout(m.Layout)
//   - name: "analog_layout", value: typedef.AnalogWatchfaceLayout(m.Layout)
//
// Otherwise:
//   - name: "layout", value: m.Layout
func (m *WatchfaceSettings) GetLayout() (name string, value any) {
	switch m.Mode {
	case typedef.WatchfaceModeDigital:
		return "digital_layout", typedef.DigitalWatchfaceLayout(m.Layout)
	case typedef.WatchfaceModeAnalog:
		return "analog_layout", typedef.AnalogWatchfaceLayout(m.Layout)
	}
	return "layout", m.Layout
}

// SetMessageIndex sets MessageIndex value.
func (m *WatchfaceSettings) SetMessageIndex(v typedef.MessageIndex) *WatchfaceSettings {
	m.MessageIndex = v
	return m
}

// SetMode sets Mode value.
func (m *WatchfaceSettings) SetMode(v typedef.WatchfaceMode) *WatchfaceSettings {
	m.Mode = v
	return m
}

// SetLayout sets Layout value.
func (m *WatchfaceSettings) SetLayout(v byte) *WatchfaceSettings {
	m.Layout = v
	return m
}

// SetDeveloperFields WatchfaceSettings's DeveloperFields.
func (m *WatchfaceSettings) SetDeveloperFields(developerFields ...proto.DeveloperField) *WatchfaceSettings {
	m.DeveloperFields = developerFields
	return m
}
