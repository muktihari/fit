// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Workout is a file contains instructions for performing a structured activity.
//
// ref: https://developer.garmin.com/fit/file-types/workout/
type Workout struct {
	FileId       *mesgdef.FileId
	Workout      *mesgdef.Workout
	WorkoutSteps []*mesgdef.WorkoutStep

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	// Messages not related to Workout
	UnrelatedMessages []proto.Message
}

var _ File = &Workout{}

func NewWorkout(mesgs ...proto.Message) *Workout {
	f := &Workout{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Workout) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = mesgdef.NewFileId(mesg)
	case mesgnum.Workout:
		f.Workout = mesgdef.NewWorkout(mesg)
	case mesgnum.WorkoutStep:
		f.WorkoutSteps = append(f.WorkoutSteps, mesgdef.NewWorkoutStep(mesg))
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Workout) ToFit(fac Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	size := 2 /* non slice fields */

	size += len(f.WorkoutSteps) + len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	if f.FileId != nil {
		mesg := fac.CreateMesg(mesgnum.FileId)
		f.FileId.PutMessage(&mesg)
		fit.Messages = append(fit.Messages, mesg)
	}

	PutMessages(fac, &fit.Messages, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	PutMessages(fac, &fit.Messages, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.Workout != nil {
		mesg := fac.CreateMesg(mesgnum.Workout)
		f.FileId.PutMessage(&mesg)
		fit.Messages = append(fit.Messages, mesg)
	}

	PutMessages(fac, &fit.Messages, mesgnum.WorkoutStep, f.WorkoutSteps)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
