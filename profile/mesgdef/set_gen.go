// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// Set is a Set message.
type Set struct {
	Timestamp         time.Time // Timestamp of the set
	Duration          uint32    // Scale: 1000; Units: s;
	Repetitions       uint16    // # of repitions of the movement
	Weight            uint16    // Scale: 16; Units: kg; Amount of weight applied for the set
	SetType           typedef.SetType
	StartTime         time.Time                  // Start time of the set
	Category          []typedef.ExerciseCategory // Array: [N];
	CategorySubtype   []uint16                   // Array: [N]; Based on the associated category, see [category]_exercise_names
	WeightDisplayUnit typedef.FitBaseUnit
	MessageIndex      typedef.MessageIndex
	WktStepIndex      typedef.MessageIndex

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSet creates new Set struct based on given mesg.
// If mesg is nil, it will return Set with all fields being set to its corresponding invalid value.
func NewSet(mesg *proto.Message) *Set {
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

	return &Set{
		Timestamp:         datetime.ToTime(vals[254]),
		Duration:          typeconv.ToUint32[uint32](vals[0]),
		Repetitions:       typeconv.ToUint16[uint16](vals[3]),
		Weight:            typeconv.ToUint16[uint16](vals[4]),
		SetType:           typeconv.ToUint8[typedef.SetType](vals[5]),
		StartTime:         datetime.ToTime(vals[6]),
		Category:          typeconv.ToSliceUint16[typedef.ExerciseCategory](vals[7]),
		CategorySubtype:   typeconv.ToSliceUint16[uint16](vals[8]),
		WeightDisplayUnit: typeconv.ToUint16[typedef.FitBaseUnit](vals[9]),
		MessageIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[10]),
		WktStepIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[11]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Set into proto.Message.
func (m *Set) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumSet)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Duration
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Repetitions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Repetitions
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Weight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.Weight
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8[uint8](m.SetType) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = typeconv.ToUint8[uint8](m.SetType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if datetime.ToUint32(m.StartTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = datetime.ToUint32(m.StartTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToSliceUint16[uint16](m.Category) != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToSliceUint16[uint16](m.Category)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.CategorySubtype != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.CategorySubtype
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.WeightDisplayUnit) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToUint16[uint16](m.WeightDisplayUnit)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.WktStepIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = typeconv.ToUint16[uint16](m.WktStepIndex)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Set's valid fields.
func (m *Set) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if m.Duration != basetype.Uint32Invalid {
		size++
	}
	if m.Repetitions != basetype.Uint16Invalid {
		size++
	}
	if m.Weight != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint8[uint8](m.SetType) != basetype.Uint8Invalid {
		size++
	}
	if datetime.ToUint32(m.StartTime) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToSliceUint16[uint16](m.Category) != nil {
		size++
	}
	if m.CategorySubtype != nil {
		size++
	}
	if typeconv.ToUint16[uint16](m.WeightDisplayUnit) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.WktStepIndex) != basetype.Uint16Invalid {
		size++
	}
	return size
}

// SetTimestamp sets Set value.
//
// Timestamp of the set
func (m *Set) SetTimestamp(v time.Time) *Set {
	m.Timestamp = v
	return m
}

// SetDuration sets Set value.
//
// Scale: 1000; Units: s;
func (m *Set) SetDuration(v uint32) *Set {
	m.Duration = v
	return m
}

// SetRepetitions sets Set value.
//
// # of repitions of the movement
func (m *Set) SetRepetitions(v uint16) *Set {
	m.Repetitions = v
	return m
}

// SetWeight sets Set value.
//
// Scale: 16; Units: kg; Amount of weight applied for the set
func (m *Set) SetWeight(v uint16) *Set {
	m.Weight = v
	return m
}

// SetSetType sets Set value.
func (m *Set) SetSetType(v typedef.SetType) *Set {
	m.SetType = v
	return m
}

// SetStartTime sets Set value.
//
// Start time of the set
func (m *Set) SetStartTime(v time.Time) *Set {
	m.StartTime = v
	return m
}

// SetCategory sets Set value.
//
// Array: [N];
func (m *Set) SetCategory(v []typedef.ExerciseCategory) *Set {
	m.Category = v
	return m
}

// SetCategorySubtype sets Set value.
//
// Array: [N]; Based on the associated category, see [category]_exercise_names
func (m *Set) SetCategorySubtype(v []uint16) *Set {
	m.CategorySubtype = v
	return m
}

// SetWeightDisplayUnit sets Set value.
func (m *Set) SetWeightDisplayUnit(v typedef.FitBaseUnit) *Set {
	m.WeightDisplayUnit = v
	return m
}

// SetMessageIndex sets Set value.
func (m *Set) SetMessageIndex(v typedef.MessageIndex) *Set {
	m.MessageIndex = v
	return m
}

// SetWktStepIndex sets Set value.
func (m *Set) SetWktStepIndex(v typedef.MessageIndex) *Set {
	m.WktStepIndex = v
	return m
}

// SetDeveloperFields Set's DeveloperFields.
func (m *Set) SetDeveloperFields(developerFields ...proto.DeveloperField) *Set {
	m.DeveloperFields = developerFields
	return m
}
