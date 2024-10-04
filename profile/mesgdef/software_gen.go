// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// Software is a Software message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Software struct {
	PartNumber   string
	MessageIndex typedef.MessageIndex
	Version      uint16 // Scale: 100

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSoftware creates new Software struct based on given mesg.
// If mesg is nil, it will return Software with all fields being set to its corresponding invalid value.
func NewSoftware(mesg *proto.Message) *Software {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Software{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Version:      vals[3].Uint16(),
		PartNumber:   vals[5].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Software into proto.Message. If options is nil, default options will be used.
func (m *Software) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSoftware}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Version != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.Version)
		fields = append(fields, field)
	}
	if m.PartNumber != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.String(m.PartNumber)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// VersionScaled return Version in its scaled value.
// If Version value is invalid, float64 invalid value will be returned.
//
// Scale: 100
func (m *Software) VersionScaled() float64 {
	if m.Version == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Version)/100 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *Software) SetMessageIndex(v typedef.MessageIndex) *Software {
	m.MessageIndex = v
	return m
}

// SetVersion sets Version value.
//
// Scale: 100
func (m *Software) SetVersion(v uint16) *Software {
	m.Version = v
	return m
}

// SetVersionScaled is similar to SetVersion except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100
func (m *Software) SetVersionScaled(v float64) *Software {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Version = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Version = uint16(unscaled)
	return m
}

// SetPartNumber sets PartNumber value.
func (m *Software) SetPartNumber(v string) *Software {
	m.PartNumber = v
	return m
}

// SetUnknownFields Software's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Software) SetUnknownFields(unknownFields ...proto.Field) *Software {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Software's DeveloperFields.
func (m *Software) SetDeveloperFields(developerFields ...proto.DeveloperField) *Software {
	m.DeveloperFields = developerFields
	return m
}
