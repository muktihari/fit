// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type Activity struct {
	FileId   *mesgdef.FileId
	Activity *mesgdef.Activity
	Sessions []*mesgdef.Session
	Laps     []*mesgdef.Lap
	Records  []*mesgdef.Record

	// Optional Messages
	DeviceInfo   *mesgdef.DeviceInfo
	UserProfile  *mesgdef.UserProfile
	Events       []*mesgdef.Event
	Lengths      []*mesgdef.Length
	SegmentLap   []*mesgdef.SegmentLap
	ZoneTargets  []*mesgdef.ZonesTarget
	Workouts     []*mesgdef.Workout
	WorkoutSteps []*mesgdef.WorkoutStep
	HRs          []*mesgdef.Hr
	HRVs         []*mesgdef.Hrv

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription
}

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
		f.FileId = mesgdef.NewFileId(mesg)
	case mesgnum.Activity:
		f.Activity = mesgdef.NewActivity(mesg)
	case mesgnum.Session:
		f.Sessions = append(f.Sessions, mesgdef.NewSession(mesg))
	case mesgnum.Lap:
		f.Laps = append(f.Laps, mesgdef.NewLap(mesg))
	case mesgnum.Record:
		f.Records = append(f.Records, mesgdef.NewRecord(mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfo = mesgdef.NewDeviceInfo(mesg)
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(mesg)
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(mesg))
	case mesgnum.Length:
		f.Lengths = append(f.Lengths, mesgdef.NewLength(mesg))
	case mesgnum.SegmentLap:
		f.SegmentLap = append(f.SegmentLap, mesgdef.NewSegmentLap(mesg))
	case mesgnum.ZonesTarget:
		f.ZoneTargets = append(f.ZoneTargets, mesgdef.NewZonesTarget(mesg))
	case mesgnum.Workout:
		f.Workouts = append(f.Workouts, mesgdef.NewWorkout(mesg))
	case mesgnum.Hr:
		f.HRs = append(f.HRs, mesgdef.NewHr(mesg))
	case mesgnum.Hrv:
		f.HRVs = append(f.HRVs, mesgdef.NewHrv(mesg))
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(mesg))
	}
}

func (f *Activity) ToFit(factory Factory) proto.Fit {
	var size = 4 // non slice fields

	size += len(f.Sessions) + len(f.Laps) + len(f.Records) +
		len(f.Events) + len(f.Lengths) + len(f.SegmentLap) + len(f.ZoneTargets) +
		len(f.Workouts) + len(f.WorkoutSteps) + len(f.HRs) + len(f.HRVs)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	if f.FileId != nil {
		mesg := factory.CreateMesg(mesgnum.FileId)
		f.FileId.PutMessage(&mesg)
		fit.Messages = append(fit.Messages, mesg)
	}

	PutMessages(factory, &fit.Messages, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	PutMessages(factory, &fit.Messages, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.Activity != nil {
		mesg := factory.CreateMesg(mesgnum.Activity)
		f.Activity.PutMessage(&mesg)
		fit.Messages = append(fit.Messages, mesg)
	}

	if f.DeviceInfo != nil {
		mesg := factory.CreateMesg(mesgnum.DeviceInfo)
		f.DeviceInfo.PutMessage(&mesg)
		fit.Messages = append(fit.Messages, mesg)
	}

	if f.UserProfile != nil {
		mesg := factory.CreateMesg(mesgnum.UserProfile)
		f.UserProfile.PutMessage(&mesg)
		fit.Messages = append(fit.Messages, mesg)
	}

	PutMessages(factory, &fit.Messages, mesgnum.Session, f.Sessions)
	PutMessages(factory, &fit.Messages, mesgnum.Lap, f.Laps)
	PutMessages(factory, &fit.Messages, mesgnum.Record, f.Records)
	PutMessages(factory, &fit.Messages, mesgnum.Event, f.Events)
	PutMessages(factory, &fit.Messages, mesgnum.Length, f.Lengths)
	PutMessages(factory, &fit.Messages, mesgnum.SegmentLap, f.SegmentLap)
	PutMessages(factory, &fit.Messages, mesgnum.ZonesTarget, f.ZoneTargets)
	PutMessages(factory, &fit.Messages, mesgnum.Workout, f.Workouts)
	PutMessages(factory, &fit.Messages, mesgnum.WorkoutStep, f.WorkoutSteps)
	PutMessages(factory, &fit.Messages, mesgnum.Hr, f.HRs)
	PutMessages(factory, &fit.Messages, mesgnum.Hrv, f.HRVs)

	return fit
}
