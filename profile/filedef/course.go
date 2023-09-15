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

type CourseFile struct {
	FileId  *mesgdef.FileId
	Course  *mesgdef.Course
	Lap     *mesgdef.Lap
	Records []*mesgdef.Record
	Events  []*mesgdef.Event

	// Optional Messages
	CoursePoints []*mesgdef.CoursePoint

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription
}

func NewCourseFile(mesgs ...proto.Message) (f *CourseFile, ok bool) {
	f = &CourseFile{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	if f.FileId == nil || f.FileId.Type != typedef.FileCourse {
		return
	}

	return f, true
}

func (f *CourseFile) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = mesgdef.NewFileId(mesg)
	case mesgnum.Course:
		f.Course = mesgdef.NewCourse(mesg)
	case mesgnum.Lap:
		f.Lap = mesgdef.NewLap(mesg)
	case mesgnum.Record:
		f.Records = append(f.Records, mesgdef.NewRecord(mesg))
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(mesg))
	case mesgnum.CoursePoint:
		f.CoursePoints = append(f.CoursePoints, mesgdef.NewCoursePoint(mesg))
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(mesg))
	}
}
