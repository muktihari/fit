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

// ExerciseTitle is a ExerciseTitle message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
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

	return &ExerciseTitle{
		MessageIndex:     typedef.MessageIndex(vals[254].Uint16()),
		ExerciseCategory: typedef.ExerciseCategory(vals[0].Uint16()),
		ExerciseName:     vals[1].Uint16(),
		WktStepName:      vals[2].SliceString(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ExerciseTitle into proto.Message. If options is nil, default options will be used.
func (m *ExerciseTitle) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumExerciseTitle}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if uint16(m.ExerciseCategory) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.ExerciseCategory))
		fields = append(fields, field)
	}
	if m.ExerciseName != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.ExerciseName)
		fields = append(fields, field)
	}
	if m.WktStepName != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceString(m.WktStepName)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
func (m *ExerciseTitle) SetMessageIndex(v typedef.MessageIndex) *ExerciseTitle {
	m.MessageIndex = v
	return m
}

// SetExerciseCategory sets ExerciseCategory value.
func (m *ExerciseTitle) SetExerciseCategory(v typedef.ExerciseCategory) *ExerciseTitle {
	m.ExerciseCategory = v
	return m
}

// SetExerciseName sets ExerciseName value.
func (m *ExerciseTitle) SetExerciseName(v uint16) *ExerciseTitle {
	m.ExerciseName = v
	return m
}

// SetWktStepName sets WktStepName value.
//
// Array: [N]
func (m *ExerciseTitle) SetWktStepName(v []string) *ExerciseTitle {
	m.WktStepName = v
	return m
}

// SetDeveloperFields ExerciseTitle's DeveloperFields.
func (m *ExerciseTitle) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExerciseTitle {
	m.DeveloperFields = developerFields
	return m
}
