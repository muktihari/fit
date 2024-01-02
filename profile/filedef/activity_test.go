// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func createFieldComparer() cmp.Option {
	return cmp.Comparer(func(field1, field2 proto.Field) bool {
		// Compare float in integer form. since when f is NaN, f != f.
		switch f1 := field1.Value.(type) {
		case float32:
			f2, ok := field2.Value.(float32)
			if !ok {
				break
			}
			return math.Float32bits(f1) == math.Float32bits(f2)
		case float64:
			f2, ok := field2.Value.(float64)
			if !ok {
				break
			}
			return math.Float64bits(f1) == math.Float64bits(f2)
		}
		return cmp.Diff(field1, field2) == ""
	})
}

// incrementSecond will increment v for 1 second to similate incoming message in 1 second sampling.
func incrementSecond(v *time.Time) time.Time {
	*v = v.Add(time.Second)
	return *v
}

func newActivityMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivity)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.DeviceInfo).WithFields(
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
		),
		factory.CreateMesgOnly(mesgnum.UserProfile).WithFields(
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileFriendlyName).WithValue("Mary Jane"),
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(21)),
		),
		factory.CreateMesgOnly(mesgnum.Event).WithFields(
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
			factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStart)),
		),
		factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Event).WithFields(
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Event).WithFields(
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Lap).WithFields(
			factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(datetime.ToUint32(now)), // intentionally using same timestamp as last messag)e
		),
		factory.CreateMesgOnly(mesgnum.Session).WithFields(
			factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Activity).WithFields(
			factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		// Unordered optional Messages
		factory.CreateMesgOnly(mesgnum.Length).WithFields(
			factory.CreateField(mesgnum.Length, fieldnum.LengthAvgSpeed).WithValue(uint16(1000)),
		),
		factory.CreateMesgOnly(mesgnum.SegmentLap).WithFields(
			factory.CreateField(mesgnum.SegmentLap, fieldnum.SegmentLapAvgCadence).WithValue(uint8(100)),
		),
		factory.CreateMesgOnly(mesgnum.ZonesTarget).WithFields(
			factory.CreateField(mesgnum.ZonesTarget, fieldnum.ZonesTargetMaxHeartRate).WithValue(uint8(190)),
		),
		factory.CreateMesgOnly(mesgnum.Workout).WithFields(
			factory.CreateField(mesgnum.Workout, fieldnum.WorkoutSport).WithValue(uint8(typedef.SportCycling)),
		),
		factory.CreateMesgOnly(mesgnum.WorkoutStep).WithFields(
			factory.CreateField(mesgnum.WorkoutStep, fieldnum.WorkoutStepIntensity).WithValue(uint8(typedef.IntensityActive)),
		),
		factory.CreateMesgOnly(mesgnum.Hr).WithFields(
			factory.CreateField(mesgnum.Hr, fieldnum.HrTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Hrv).WithFields(
			factory.CreateField(mesgnum.Hrv, fieldnum.HrvTime).WithValue([]uint16{uint16(1000)}),
		),
		// Unrelated messages
		factory.CreateMesgOnly(mesgnum.BarometerData).WithFields(
			factory.CreateField(mesgnum.BarometerData, fieldnum.BarometerDataTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.CoursePoint).WithFields(
			factory.CreateField(mesgnum.CoursePoint, fieldnum.CoursePointTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
	}
}

func TestActivityCorrectness(t *testing.T) {
	mesgs := newActivityMessageForTest(time.Now())

	activity := filedef.NewActivity(mesgs...)
	if activity.FileId.Type != typedef.FileActivity {
		t.Fatalf("expected: %v, got: %v", typedef.FileActivity, activity.FileId.Type)
	}

	fit := activity.ToFit(nil) // use standard factory

	if diff := cmp.Diff(mesgs, fit.Messages, createFieldComparer()); diff != "" {
		fmt.Println("messages order:")
		for i := range fit.Messages {
			mesg := fit.Messages[i]
			fmt.Printf("%d: %s\n", mesg.Num, mesg.Num)
		}
		fmt.Println("")
		t.Fatal(diff)
	}
}
