// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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
	UserProfile    *mesgdef.UserProfile
	DeviceInfos    []*mesgdef.DeviceInfo // required fields: timestamp
	Events         []*mesgdef.Event
	Lengths        []*mesgdef.Length // required fields: timestamp, event, event_type
	SegmentLaps    []*mesgdef.SegmentLap
	ZonesTargets   []*mesgdef.ZonesTarget
	Workouts       []*mesgdef.Workout
	WorkoutSteps   []*mesgdef.WorkoutStep
	HRs            []*mesgdef.Hr
	HRVs           []*mesgdef.Hrv // required fields: time
	GpsMetadatas   []*mesgdef.GpsMetadata
	TimeInZones    []*mesgdef.TimeInZone
	Splits         []*mesgdef.Split
	SplitSummaries []*mesgdef.SplitSummary // entries must be unique within each split_type

	// Messages not related to Activity
	UnrelatedMessages []proto.Message
}

var _ File = (*Activity)(nil)

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
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(&mesg)
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(&mesg))
	case mesgnum.Length:
		f.Lengths = append(f.Lengths, mesgdef.NewLength(&mesg))
	case mesgnum.SegmentLap:
		f.SegmentLaps = append(f.SegmentLaps, mesgdef.NewSegmentLap(&mesg))
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
	case mesgnum.GpsMetadata:
		f.GpsMetadatas = append(f.GpsMetadatas, mesgdef.NewGpsMetadata(&mesg))
	case mesgnum.TimeInZone:
		f.TimeInZones = append(f.TimeInZones, mesgdef.NewTimeInZone(&mesg))
	case mesgnum.Split:
		f.Splits = append(f.Splits, mesgdef.NewSplit(&mesg))
	case mesgnum.SplitSummary:
		f.SplitSummaries = append(f.SplitSummaries, mesgdef.NewSplitSummary(&mesg))
	default:
		mesg.Fields = append(mesg.Fields[:0:0], mesg.Fields...)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Activity to proto.FIT. If options is nil, default options will be used.
func (f *Activity) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 3 // non slice fields

	size += len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.Sessions) +
		len(f.Laps) + len(f.Records) + len(f.DeviceInfos) + len(f.Events) +
		len(f.Lengths) + len(f.SegmentLaps) + len(f.ZonesTargets) + len(f.Workouts) +
		len(f.WorkoutSteps) + len(f.HRs) + len(f.HRVs) + len(f.GpsMetadatas) +
		len(f.TimeInZones) + len(f.Splits) + len(f.SplitSummaries) + len(f.UnrelatedMessages)

	fit := proto.FIT{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	var sortStartPos = 1 + len(f.DeveloperDataIds) + len(f.FieldDescriptions)
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	for i := range f.DeveloperDataIds {
		fit.Messages = append(fit.Messages, f.DeveloperDataIds[i].ToMesg(options))
	}
	for i := range f.FieldDescriptions {
		fit.Messages = append(fit.Messages, f.FieldDescriptions[i].ToMesg(options))
	}
	for i := range f.DeviceInfos {
		fit.Messages = append(fit.Messages, f.DeviceInfos[i].ToMesg(options))
	}
	if f.UserProfile != nil {
		fit.Messages = append(fit.Messages, f.UserProfile.ToMesg(options))
	}
	if f.Activity != nil {
		fit.Messages = append(fit.Messages, f.Activity.ToMesg(options))
	}
	for i := range f.Sessions {
		fit.Messages = append(fit.Messages, f.Sessions[i].ToMesg(options))
	}
	for i := range f.Laps {
		fit.Messages = append(fit.Messages, f.Laps[i].ToMesg(options))
	}
	for i := range f.Records {
		fit.Messages = append(fit.Messages, f.Records[i].ToMesg(options))
	}
	for i := range f.Events {
		fit.Messages = append(fit.Messages, f.Events[i].ToMesg(options))
	}
	for i := range f.Lengths {
		fit.Messages = append(fit.Messages, f.Lengths[i].ToMesg(options))
	}
	for i := range f.SegmentLaps {
		fit.Messages = append(fit.Messages, f.SegmentLaps[i].ToMesg(options))
	}
	for i := range f.ZonesTargets {
		fit.Messages = append(fit.Messages, f.ZonesTargets[i].ToMesg(options))
	}
	for i := range f.Workouts {
		fit.Messages = append(fit.Messages, f.Workouts[i].ToMesg(options))
	}
	for i := range f.WorkoutSteps {
		fit.Messages = append(fit.Messages, f.WorkoutSteps[i].ToMesg(options))
	}
	for i := range f.HRs {
		fit.Messages = append(fit.Messages, f.HRs[i].ToMesg(options))
	}
	for i := range f.HRVs {
		fit.Messages = append(fit.Messages, f.HRVs[i].ToMesg(options))
	}
	for i := range f.GpsMetadatas {
		fit.Messages = append(fit.Messages, f.GpsMetadatas[i].ToMesg(options))
	}
	for i := range f.TimeInZones {
		fit.Messages = append(fit.Messages, f.TimeInZones[i].ToMesg(options))
	}
	for i := range f.Splits {
		fit.Messages = append(fit.Messages, f.Splits[i].ToMesg(options))
	}
	for i := range f.SplitSummaries {
		fit.Messages = append(fit.Messages, f.SplitSummaries[i].ToMesg(options))
	}

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages[sortStartPos:])

	return fit
}
