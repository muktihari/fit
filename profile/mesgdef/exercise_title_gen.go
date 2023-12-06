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

// ExerciseTitle is a ExerciseTitle message.
type ExerciseTitle struct {
	MessageIndex     typedef.MessageIndex
	ExerciseCategory typedef.ExerciseCategory
	ExerciseName     uint16
	WktStepName      string

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewExerciseTitle creates new ExerciseTitle struct based on given mesg. If mesg is nil or mesg.Num is not equal to ExerciseTitle mesg number, it will return nil.
func NewExerciseTitle(mesg proto.Message) *ExerciseTitle {
	if mesg.Num != typedef.MesgNumExerciseTitle {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		0:   basetype.Uint16Invalid, /* ExerciseCategory */
		1:   basetype.Uint16Invalid, /* ExerciseName */
		2:   nil,                    /* WktStepName */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &ExerciseTitle{
		MessageIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		ExerciseCategory: typeconv.ToUint16[typedef.ExerciseCategory](vals[0]),
		ExerciseName:     typeconv.ToUint16[uint16](vals[1]),
		WktStepName:      typeconv.ToString[string](vals[2]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to ExerciseTitle mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumExerciseTitle)
func (m ExerciseTitle) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumExerciseTitle {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.ExerciseCategory,
		1:   m.ExerciseName,
		2:   m.WktStepName,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
