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

// Activity is a common file type that most wearable device or cycling computer uses to record activities.
//
// Please note since we group the same mesgdef types in slices, we lose the arrival order of the messages.
// But for messages that have timestamp, we can reconstruct the messages by timestamp order.
//
// ref: https://developer.garmin.com/fit/file-types/activity/
type Activity struct {
	FileId mesgdef.FileId // must have mesg

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	// Required Messages
	Activity *mesgdef.Activity
	Sessions []*mesgdef.Session
	Laps     []*mesgdef.Lap
	Records  []*mesgdef.Record

	// Optional Messages
	UserProfile  *mesgdef.UserProfile
	DeviceInfos  []*mesgdef.DeviceInfo
	Events       []*mesgdef.Event
	Lengths      []*mesgdef.Length
	SegmentLap   []*mesgdef.SegmentLap
	ZonesTargets []*mesgdef.ZonesTarget
	Workouts     []*mesgdef.Workout
	WorkoutSteps []*mesgdef.WorkoutStep
	HRs          []*mesgdef.Hr
	HRVs         []*mesgdef.Hrv

	// Messages not related to Activity
	UnrelatedMessages []proto.Message
}

var _ File = &Activity{}

func NewActivity(mesgs ...proto.Message) *Activity {
	f := &Activity{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Activity) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(mesg))
	case mesgnum.Activity:
		f.Activity = mesgdef.NewActivity(mesg)
	case mesgnum.Session:
		f.Sessions = append(f.Sessions, mesgdef.NewSession(mesg))
	case mesgnum.Lap:
		f.Laps = append(f.Laps, mesgdef.NewLap(mesg))
	case mesgnum.Record:
		f.Records = append(f.Records, mesgdef.NewRecord(mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(mesg))
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(mesg)
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(mesg))
	case mesgnum.Length:
		f.Lengths = append(f.Lengths, mesgdef.NewLength(mesg))
	case mesgnum.SegmentLap:
		f.SegmentLap = append(f.SegmentLap, mesgdef.NewSegmentLap(mesg))
	case mesgnum.ZonesTarget:
		f.ZonesTargets = append(f.ZonesTargets, mesgdef.NewZonesTarget(mesg))
	case mesgnum.Workout:
		f.Workouts = append(f.Workouts, mesgdef.NewWorkout(mesg))
	case mesgnum.WorkoutStep:
		f.WorkoutSteps = append(f.WorkoutSteps, mesgdef.NewWorkoutStep(mesg))
	case mesgnum.Hr:
		f.HRs = append(f.HRs, mesgdef.NewHr(mesg))
	case mesgnum.Hrv:
		f.HRVs = append(f.HRVs, mesgdef.NewHrv(mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Activity) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 3 // non slice fields

	size += len(f.Sessions) + len(f.Laps) + len(f.Records) + len(f.DeviceInfos) +
		len(f.Events) + len(f.Lengths) + len(f.SegmentLap) + len(f.ZonesTargets) +
		len(f.Workouts) + len(f.WorkoutSteps) + len(f.HRs) + len(f.HRVs) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, fac, mesgnum.DeviceInfo, f.DeviceInfos)

	if f.UserProfile != nil {
		fit.Messages = append(fit.Messages, f.UserProfile.ToMesg(fac))
	}

	if f.Activity != nil {
		fit.Messages = append(fit.Messages, f.Activity.ToMesg(fac))
	}

	ToMesgs(&fit.Messages, fac, mesgnum.Session, f.Sessions)
	ToMesgs(&fit.Messages, fac, mesgnum.Lap, f.Laps)
	ToMesgs(&fit.Messages, fac, mesgnum.Record, f.Records)
	ToMesgs(&fit.Messages, fac, mesgnum.Event, f.Events)
	ToMesgs(&fit.Messages, fac, mesgnum.Length, f.Lengths)
	ToMesgs(&fit.Messages, fac, mesgnum.SegmentLap, f.SegmentLap)
	ToMesgs(&fit.Messages, fac, mesgnum.ZonesTarget, f.ZonesTargets)
	ToMesgs(&fit.Messages, fac, mesgnum.Workout, f.Workouts)
	ToMesgs(&fit.Messages, fac, mesgnum.WorkoutStep, f.WorkoutSteps)
	ToMesgs(&fit.Messages, fac, mesgnum.Hr, f.HRs)
	ToMesgs(&fit.Messages, fac, mesgnum.Hrv, f.HRVs)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
