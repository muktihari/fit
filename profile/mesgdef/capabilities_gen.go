// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"unsafe"
)

// Capabilities is a Capabilities message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Capabilities struct {
	Languages             []uint8                          // Base: uint8z; Array: [N]; Use language_bits_x types where x is index of array.
	Sports                []typedef.SportBits0             // Base: uint8z; Array: [N]; Use sport_bits_x types where x is index of array.
	WorkoutsSupported     typedef.WorkoutCapabilities      // Base: uint32z
	ConnectivitySupported typedef.ConnectivityCapabilities // Base: uint32z

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewCapabilities creates new Capabilities struct based on given mesg.
// If mesg is nil, it will return Capabilities with all fields being set to its corresponding invalid value.
func NewCapabilities(mesg *proto.Message) *Capabilities {
	m := new(Capabilities)
	m.Reset(mesg)
	return m
}

// Reset resets all Capabilities's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *Capabilities) Reset(mesg *proto.Message) {
	var (
		vals            [24]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		var n int
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				n++
			}
		}
		unknownFields = make([]proto.Field, 0, n)
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 24 {
				vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
			}
		}
		developerFields = mesg.DeveloperFields
	}

	*m = Capabilities{
		Languages: vals[0].SliceUint8(),
		Sports: func() []typedef.SportBits0 {
			sliceValue := vals[1].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.SportBits0)(ptr), len(sliceValue))
		}(),
		WorkoutsSupported:     typedef.WorkoutCapabilities(vals[21].Uint32z()),
		ConnectivitySupported: typedef.ConnectivityCapabilities(vals[23].Uint32z()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Capabilities into proto.Message. If options is nil, default options will be used.
func (m *Capabilities) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fields := make([]proto.Field, 0, 4)
	mesg := proto.Message{Num: typedef.MesgNumCapabilities}

	if m.Languages != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.SliceUint8(m.Languages)
		fields = append(fields, field)
	}
	if m.Sports != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint8(m.Sports)
		fields = append(fields, field)
	}
	if m.WorkoutsSupported != typedef.WorkoutCapabilitiesInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint32(uint32(m.WorkoutsSupported))
		fields = append(fields, field)
	}
	if m.ConnectivitySupported != typedef.ConnectivityCapabilitiesInvalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Uint32(uint32(m.ConnectivitySupported))
		fields = append(fields, field)
	}

	n := len(fields)
	mesg.Fields = make([]proto.Field, n+len(m.UnknownFields))
	copy(mesg.Fields[:n], fields)
	copy(mesg.Fields[n:], m.UnknownFields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetLanguages sets Languages value.
//
// Base: uint8z; Array: [N]; Use language_bits_x types where x is index of array.
func (m *Capabilities) SetLanguages(v []uint8) *Capabilities {
	m.Languages = v
	return m
}

// SetSports sets Sports value.
//
// Base: uint8z; Array: [N]; Use sport_bits_x types where x is index of array.
func (m *Capabilities) SetSports(v []typedef.SportBits0) *Capabilities {
	m.Sports = v
	return m
}

// SetWorkoutsSupported sets WorkoutsSupported value.
//
// Base: uint32z
func (m *Capabilities) SetWorkoutsSupported(v typedef.WorkoutCapabilities) *Capabilities {
	m.WorkoutsSupported = v
	return m
}

// SetConnectivitySupported sets ConnectivitySupported value.
//
// Base: uint32z
func (m *Capabilities) SetConnectivitySupported(v typedef.ConnectivityCapabilities) *Capabilities {
	m.ConnectivitySupported = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Capabilities) SetUnknownFields(unknownFields ...proto.Field) *Capabilities {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *Capabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *Capabilities {
	m.DeveloperFields = developerFields
	return m
}
