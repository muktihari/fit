// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type WorkoutFile struct {
	FileId       *mesgdef.FileId
	Workout      *mesgdef.Workout
	WorkoutSteps []*mesgdef.WorkoutStep

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription
}

func NewWorkoutFile(mesgs ...proto.Message) (f *WorkoutFile, ok bool) {
	f = &WorkoutFile{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	if f.FileId == nil || f.FileId.Type != typedef.FileCourse {
		return
	}

	return f, true
}

func (f *WorkoutFile) Add(mesg proto.Message) {
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
	}
}
