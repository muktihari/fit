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

// Workout is a Workout message.
type Workout struct {
	MessageIndex   typedef.MessageIndex
	Sport          typedef.Sport
	Capabilities   typedef.WorkoutCapabilities
	NumValidSteps  uint16 // number of valid steps
	WktName        string
	SubSport       typedef.SubSport
	PoolLength     uint16 // Scale: 100; Units: m;
	PoolLengthUnit typedef.DisplayMeasure

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWorkout creates new Workout struct based on given mesg.
// If mesg is nil, it will return Workout with all fields being set to its corresponding invalid value.
func NewWorkout(mesg *proto.Message) *Workout {
	vals := [255]any{}

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

	return &Workout{
		MessageIndex:   typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Sport:          typeconv.ToEnum[typedef.Sport](vals[4]),
		Capabilities:   typeconv.ToUint32z[typedef.WorkoutCapabilities](vals[5]),
		NumValidSteps:  typeconv.ToUint16[uint16](vals[6]),
		WktName:        typeconv.ToString[string](vals[8]),
		SubSport:       typeconv.ToEnum[typedef.SubSport](vals[11]),
		PoolLength:     typeconv.ToUint16[uint16](vals[14]),
		PoolLengthUnit: typeconv.ToEnum[typedef.DisplayMeasure](vals[15]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Workout into proto.Message.
func (m *Workout) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumWorkout)

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = typeconv.ToEnum[byte](m.Sport)
		fields = append(fields, field)
	}
	if typeconv.ToUint32z[uint32](m.Capabilities) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = typeconv.ToUint32z[uint32](m.Capabilities)
		fields = append(fields, field)
	}
	if m.NumValidSteps != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.NumValidSteps
		fields = append(fields, field)
	}
	if m.WktName != basetype.StringInvalid && m.WktName != "" {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.WktName
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = typeconv.ToEnum[byte](m.SubSport)
		fields = append(fields, field)
	}
	if m.PoolLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.PoolLength
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.PoolLengthUnit) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = typeconv.ToEnum[byte](m.PoolLengthUnit)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets Workout value.
func (m *Workout) SetMessageIndex(v typedef.MessageIndex) *Workout {
	m.MessageIndex = v
	return m
}

// SetSport sets Workout value.
func (m *Workout) SetSport(v typedef.Sport) *Workout {
	m.Sport = v
	return m
}

// SetCapabilities sets Workout value.
func (m *Workout) SetCapabilities(v typedef.WorkoutCapabilities) *Workout {
	m.Capabilities = v
	return m
}

// SetNumValidSteps sets Workout value.
//
// number of valid steps
func (m *Workout) SetNumValidSteps(v uint16) *Workout {
	m.NumValidSteps = v
	return m
}

// SetWktName sets Workout value.
func (m *Workout) SetWktName(v string) *Workout {
	m.WktName = v
	return m
}

// SetSubSport sets Workout value.
func (m *Workout) SetSubSport(v typedef.SubSport) *Workout {
	m.SubSport = v
	return m
}

// SetPoolLength sets Workout value.
//
// Scale: 100; Units: m;
func (m *Workout) SetPoolLength(v uint16) *Workout {
	m.PoolLength = v
	return m
}

// SetPoolLengthUnit sets Workout value.
func (m *Workout) SetPoolLengthUnit(v typedef.DisplayMeasure) *Workout {
	m.PoolLengthUnit = v
	return m
}

// SetDeveloperFields Workout's DeveloperFields.
func (m *Workout) SetDeveloperFields(developerFields ...proto.DeveloperField) *Workout {
	m.DeveloperFields = developerFields
	return m
}
