// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
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
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number, time_created

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	// Required Messages
	Activity *mesgdef.Activity  // required fields: timestamp, num_sessions, type, event, event_type
	Sessions []*mesgdef.Session // required fields: timestamp, start_time, total_elapsed_time, sport, event, event_type
	Laps     []*mesgdef.Lap     // required fields: timestamp, event, event_type
	Records  []*mesgdef.Record  // required fields: timestamp

	// Optional Messages
	UserProfile  *mesgdef.UserProfile
	DeviceInfos  []*mesgdef.DeviceInfo // required fields: timestamp
	Events       []*mesgdef.Event
	Lengths      []*mesgdef.Length // required fields: timestamp, event, event_type
	SegmentLap   []*mesgdef.SegmentLap
	ZonesTargets []*mesgdef.ZonesTarget
	Workouts     []*mesgdef.Workout
	WorkoutSteps []*mesgdef.WorkoutStep
	HRs          []*mesgdef.Hr
	HRVs         []*mesgdef.Hrv // required fields: time

	// Messages not related to Activity
	UnrelatedMessages []proto.Message
}

var _ File = &Activity{}

// NewActivity creates new Activity File.
func NewActivity(mesgs ...proto.Message) *Activity {
	f := &Activity{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

// Add adds mesg to the Activity.
func (f *Activity) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Activity:
		f.Activity = mesgdef.NewActivity(&mesg)
	case mesgnum.Session:
		f.Sessions = append(f.Sessions, mesgdef.NewSession(&mesg))
	case mesgnum.Lap:
		f.Laps = append(f.Laps, mesgdef.NewLap(&mesg))
	case mesgnum.Record:
		f.Records = append(f.Records, mesgdef.NewRecord(&mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(&mesg)
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(&mesg))
	case mesgnum.Length:
		f.Lengths = append(f.Lengths, mesgdef.NewLength(&mesg))
	case mesgnum.SegmentLap:
		f.SegmentLap = append(f.SegmentLap, mesgdef.NewSegmentLap(&mesg))
	case mesgnum.ZonesTarget:
		f.ZonesTargets = append(f.ZonesTargets, mesgdef.NewZonesTarget(&mesg))
	case mesgnum.Workout:
		f.Workouts = append(f.Workouts, mesgdef.NewWorkout(&mesg))
	case mesgnum.WorkoutStep:
		f.WorkoutSteps = append(f.WorkoutSteps, mesgdef.NewWorkoutStep(&mesg))
	case mesgnum.Hr:
		f.HRs = append(f.HRs, mesgdef.NewHr(&mesg))
	case mesgnum.Hrv:
		f.HRVs = append(f.HRVs, mesgdef.NewHrv(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFit converts Activity to proto.Fit. If options is nil, default options will be used.
func (f *Activity) ToFit(options *mesgdef.Options) proto.FIT {
	var size = 3 // non slice fields

	size += len(f.Sessions) + len(f.Laps) + len(f.Records) + len(f.DeviceInfos) +
		len(f.Events) + len(f.Lengths) + len(f.SegmentLap) + len(f.ZonesTargets) +
		len(f.Workouts) + len(f.WorkoutSteps) + len(f.HRs) + len(f.HRVs) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.FIT{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	ToMesgs(&fit.Messages, options, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, options, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, options, mesgnum.DeviceInfo, f.DeviceInfos)

	if f.UserProfile != nil {
		fit.Messages = append(fit.Messages, f.UserProfile.ToMesg(options))
	}

	if f.Activity != nil {
		fit.Messages = append(fit.Messages, f.Activity.ToMesg(options))
	}

	ToMesgs(&fit.Messages, options, mesgnum.Session, f.Sessions)
	ToMesgs(&fit.Messages, options, mesgnum.Lap, f.Laps)
	ToMesgs(&fit.Messages, options, mesgnum.Record, f.Records)
	ToMesgs(&fit.Messages, options, mesgnum.Event, f.Events)
	ToMesgs(&fit.Messages, options, mesgnum.Length, f.Lengths)
	ToMesgs(&fit.Messages, options, mesgnum.SegmentLap, f.SegmentLap)
	ToMesgs(&fit.Messages, options, mesgnum.ZonesTarget, f.ZonesTargets)
	ToMesgs(&fit.Messages, options, mesgnum.Workout, f.Workouts)
	ToMesgs(&fit.Messages, options, mesgnum.WorkoutStep, f.WorkoutSteps)
	ToMesgs(&fit.Messages, options, mesgnum.Hr, f.HRs)
	ToMesgs(&fit.Messages, options, mesgnum.Hrv, f.HRVs)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
