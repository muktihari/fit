// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

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

// Set is a Set message.
type Set struct {
	Timestamp         typedef.DateTime // Timestamp of the set
	Duration          uint32           // Scale: 1000; Units: s;
	Repetitions       uint16           // # of repitions of the movement
	Weight            uint16           // Scale: 16; Units: kg; Amount of weight applied for the set
	SetType           typedef.SetType
	StartTime         typedef.DateTime           // Start time of the set
	Category          []typedef.ExerciseCategory // Array: [N];
	CategorySubtype   []uint16                   // Array: [N]; Based on the associated category, see [category]_exercise_names
	WeightDisplayUnit typedef.FitBaseUnit
	MessageIndex      typedef.MessageIndex
	WktStepIndex      typedef.MessageIndex

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSet creates new Set struct based on given mesg. If mesg is nil or mesg.Num is not equal to Set mesg number, it will return nil.
func NewSet(mesg proto.Message) *Set {
	if mesg.Num != typedef.MesgNumSet {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint32Invalid, /* Duration */
		3:   basetype.Uint16Invalid, /* Repetitions */
		4:   basetype.Uint16Invalid, /* Weight */
		5:   basetype.Uint8Invalid,  /* SetType */
		6:   basetype.Uint32Invalid, /* StartTime */
		7:   nil,                    /* Category */
		8:   nil,                    /* CategorySubtype */
		9:   basetype.Uint16Invalid, /* WeightDisplayUnit */
		10:  basetype.Uint16Invalid, /* MessageIndex */
		11:  basetype.Uint16Invalid, /* WktStepIndex */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Set{
		Timestamp:         typeconv.ToUint32[typedef.DateTime](vals[254]),
		Duration:          typeconv.ToUint32[uint32](vals[0]),
		Repetitions:       typeconv.ToUint16[uint16](vals[3]),
		Weight:            typeconv.ToUint16[uint16](vals[4]),
		SetType:           typeconv.ToUint8[typedef.SetType](vals[5]),
		StartTime:         typeconv.ToUint32[typedef.DateTime](vals[6]),
		Category:          typeconv.ToSliceUint16[typedef.ExerciseCategory](vals[7]),
		CategorySubtype:   typeconv.ToSliceUint16[uint16](vals[8]),
		WeightDisplayUnit: typeconv.ToUint16[typedef.FitBaseUnit](vals[9]),
		MessageIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[10]),
		WktStepIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[11]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Set mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSet)
func (m Set) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSet {
		return
	}

	vals := [256]any{
		254: m.Timestamp,
		0:   m.Duration,
		3:   m.Repetitions,
		4:   m.Weight,
		5:   m.SetType,
		6:   m.StartTime,
		7:   m.Category,
		8:   m.CategorySubtype,
		9:   m.WeightDisplayUnit,
		10:  m.MessageIndex,
		11:  m.WktStepIndex,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}