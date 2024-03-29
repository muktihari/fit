// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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

// ExerciseTitle is a ExerciseTitle message.
type ExerciseTitle struct {
	WktStepName      []string // Array: [N]
	MessageIndex     typedef.MessageIndex
	ExerciseCategory typedef.ExerciseCategory
	ExerciseName     uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewExerciseTitle creates new ExerciseTitle struct based on given mesg.
// If mesg is nil, it will return ExerciseTitle with all fields being set to its corresponding invalid value.
func NewExerciseTitle(mesg *proto.Message) *ExerciseTitle {
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

	return &ExerciseTitle{
		WktStepName:      typeconv.ToSliceString[string](vals[2]),
		MessageIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		ExerciseCategory: typeconv.ToUint16[typedef.ExerciseCategory](vals[0]),
		ExerciseName:     typeconv.ToUint16[uint16](vals[1]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ExerciseTitle into proto.Message.
func (m *ExerciseTitle) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumExerciseTitle)

	if m.WktStepName != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.WktStepName
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if uint16(m.ExerciseCategory) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = uint16(m.ExerciseCategory)
		fields = append(fields, field)
	}
	if m.ExerciseName != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.ExerciseName
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetWktStepName sets ExerciseTitle value.
//
// Array: [N]
func (m *ExerciseTitle) SetWktStepName(v []string) *ExerciseTitle {
	m.WktStepName = v
	return m
}

// SetMessageIndex sets ExerciseTitle value.
func (m *ExerciseTitle) SetMessageIndex(v typedef.MessageIndex) *ExerciseTitle {
	m.MessageIndex = v
	return m
}

// SetExerciseCategory sets ExerciseTitle value.
func (m *ExerciseTitle) SetExerciseCategory(v typedef.ExerciseCategory) *ExerciseTitle {
	m.ExerciseCategory = v
	return m
}

// SetExerciseName sets ExerciseTitle value.
func (m *ExerciseTitle) SetExerciseName(v uint16) *ExerciseTitle {
	m.ExerciseName = v
	return m
}

// SetDeveloperFields ExerciseTitle's DeveloperFields.
func (m *ExerciseTitle) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExerciseTitle {
	m.DeveloperFields = developerFields
	return m
}
