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

// Course is a common file type used as points of courses to assist with on- and off-road navigation,
// to provide turn by turn directions, or with virtual training applications to simulate real-world activities.
//
// Please note since we group the same mesgdef types in slices, we lose the arrival order of the messages.
// But for messages that have timestamp, we can reconstruct the messages by timestamp order.
//
// ref: https://developer.garmin.com/fit/file-types/course/
type Course struct {
	FileId mesgdef.FileId // must have mesg

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	// Required Messages
	Course  *mesgdef.Course
	Lap     *mesgdef.Lap
	Records []*mesgdef.Record
	Events  []*mesgdef.Event

	// Optional Messages
	CoursePoints []*mesgdef.CoursePoint

	// Messages not related to Course
	UnrelatedMessages []proto.Message
}

var _ File = &Course{}

func NewCourse(mesgs ...proto.Message) *Course {
	f := &Course{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Course) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Course:
		f.Course = mesgdef.NewCourse(&mesg)
	case mesgnum.Lap:
		f.Lap = mesgdef.NewLap(&mesg)
	case mesgnum.Record:
		f.Records = append(f.Records, mesgdef.NewRecord(&mesg))
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(&mesg))
	case mesgnum.CoursePoint:
		f.CoursePoints = append(f.CoursePoints, mesgdef.NewCoursePoint(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Course) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	size := 3 /* non slice fields */

	size += len(f.Records) + len(f.Events) + len(f.CoursePoints) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.Course != nil {
		fit.Messages = append(fit.Messages, f.Course.ToMesg(fac))
	}

	if f.Lap != nil {
		fit.Messages = append(fit.Messages, f.Lap.ToMesg(fac))
	}

	ToMesgs(&fit.Messages, fac, mesgnum.Record, f.Records)
	ToMesgs(&fit.Messages, fac, mesgnum.Event, f.Events)
	ToMesgs(&fit.Messages, fac, mesgnum.CoursePoint, f.CoursePoints)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
