// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Capabilities is a Capabilities message.
type Capabilities struct {
	Languages             []uint8              // Array: [N]; Use language_bits_x types where x is index of array.
	Sports                []typedef.SportBits0 // Array: [N]; Use sport_bits_x types where x is index of array.
	WorkoutsSupported     typedef.WorkoutCapabilities
	ConnectivitySupported typedef.ConnectivityCapabilities

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewCapabilities creates new Capabilities struct based on given mesg.
// If mesg is nil, it will return Capabilities with all fields being set to its corresponding invalid value.
func NewCapabilities(mesg *proto.Message) *Capabilities {
	vals := [24]any{}

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

	return &Capabilities{
		Languages:             typeconv.ToSliceUint8z[uint8](vals[0]),
		Sports:                typeconv.ToSliceUint8z[typedef.SportBits0](vals[1]),
		WorkoutsSupported:     typeconv.ToUint32z[typedef.WorkoutCapabilities](vals[21]),
		ConnectivitySupported: typeconv.ToUint32z[typedef.ConnectivityCapabilities](vals[23]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Capabilities into proto.Message.
func (m *Capabilities) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumCapabilities)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToSliceUint8z[uint8](m.Languages) != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToSliceUint8z[uint8](m.Languages)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToSliceUint8z[uint8](m.Sports) != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToSliceUint8z[uint8](m.Sports)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32z[uint32](m.WorkoutsSupported) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = typeconv.ToUint32z[uint32](m.WorkoutsSupported)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32z[uint32](m.ConnectivitySupported) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = typeconv.ToUint32z[uint32](m.ConnectivitySupported)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Capabilities's valid fields.
func (m *Capabilities) size() byte {
	var size byte
	if typeconv.ToSliceUint8z[uint8](m.Languages) != nil {
		size++
	}
	if typeconv.ToSliceUint8z[uint8](m.Sports) != nil {
		size++
	}
	if typeconv.ToUint32z[uint32](m.WorkoutsSupported) != basetype.Uint32zInvalid {
		size++
	}
	if typeconv.ToUint32z[uint32](m.ConnectivitySupported) != basetype.Uint32zInvalid {
		size++
	}
	return size
}

// SetLanguages sets Capabilities value.
//
// Array: [N]; Use language_bits_x types where x is index of array.
func (m *Capabilities) SetLanguages(v []uint8) *Capabilities {
	m.Languages = v
	return m
}

// SetSports sets Capabilities value.
//
// Array: [N]; Use sport_bits_x types where x is index of array.
func (m *Capabilities) SetSports(v []typedef.SportBits0) *Capabilities {
	m.Sports = v
	return m
}

// SetWorkoutsSupported sets Capabilities value.
func (m *Capabilities) SetWorkoutsSupported(v typedef.WorkoutCapabilities) *Capabilities {
	m.WorkoutsSupported = v
	return m
}

// SetConnectivitySupported sets Capabilities value.
func (m *Capabilities) SetConnectivitySupported(v typedef.ConnectivityCapabilities) *Capabilities {
	m.ConnectivitySupported = v
	return m
}

// SetDeveloperFields Capabilities's DeveloperFields.
func (m *Capabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *Capabilities {
	m.DeveloperFields = developerFields
	return m
}
