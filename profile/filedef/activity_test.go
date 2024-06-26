// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

func sortFields(mesgs []proto.Message) {
	for i := range mesgs {
		slices.SortStableFunc(mesgs[i].Fields, func(field1, field2 proto.Field) int {
			if field1.Num < field2.Num {
				return -1
			}
			if field1.Num > field2.Num {
				return 1
			}
			return 0
		})
	}
}

func valueTransformer() cmp.Option {
	return cmp.Transformer("Value", func(v proto.Value) any {
		switch v.Type() {
		case proto.TypeFloat32:
			return math.Float32bits(v.Float32())
		case proto.TypeFloat64:
			return math.Float64bits(v.Float64())
		}
		return v.Any()
	})
}

// incrementSecond will increment v for 1 second to similate incoming message in 1 second sampling.
func incrementSecond(v *time.Time) time.Time {
	*v = v.Add(time.Second)
	return *v
}

func newActivityMessageForTest(now time.Time) []proto.Message {
	mesgs, _ := newActivityMessagesWithExpectedOrder(now)
	return mesgs
}

func newActivityMessagesWithExpectedOrder(now time.Time) (mesgs []proto.Message, ordered []proto.Message) {
	mesgs = []proto.Message{
		0: factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivity)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		1: factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		2: factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue([]string{"Heart Rate"}),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
		),
		3: factory.CreateMesgOnly(mesgnum.DeviceInfo).WithFields(
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
		),
		4: factory.CreateMesgOnly(mesgnum.UserProfile).WithFields(
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileFriendlyName).WithValue("Mary Jane"),
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(21)),
		),
		5: factory.CreateMesgOnly(mesgnum.Event).WithFields(
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
			factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStart)),
		),
		6: factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		7: factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		8: factory.CreateMesgOnly(mesgnum.Event).WithFields(
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		9: factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		10: factory.CreateMesgOnly(mesgnum.Event).WithFields(
			// Intentionally using same timestamp as last message.
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
		),
		11: factory.CreateMesgOnly(mesgnum.Lap).WithFields(
			factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		12: factory.CreateMesgOnly(mesgnum.Session).WithFields(
			factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		13: factory.CreateMesgOnly(mesgnum.Activity).WithFields(
			factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		// Unordered optional Messages
		14: factory.CreateMesgOnly(mesgnum.Length).WithFields(
			factory.CreateField(mesgnum.Length, fieldnum.LengthAvgSpeed).WithValue(uint16(1000)),
		),
		15: factory.CreateMesgOnly(mesgnum.SegmentLap).WithFields(
			factory.CreateField(mesgnum.SegmentLap, fieldnum.SegmentLapAvgCadence).WithValue(uint8(100)),
		),
		16: factory.CreateMesgOnly(mesgnum.ZonesTarget).WithFields(
			factory.CreateField(mesgnum.ZonesTarget, fieldnum.ZonesTargetMaxHeartRate).WithValue(uint8(190)),
		),
		17: factory.CreateMesgOnly(mesgnum.Workout).WithFields(
			factory.CreateField(mesgnum.Workout, fieldnum.WorkoutSport).WithValue(uint8(typedef.SportCycling)),
		),
		18: factory.CreateMesgOnly(mesgnum.WorkoutStep).WithFields(
			factory.CreateField(mesgnum.WorkoutStep, fieldnum.WorkoutStepIntensity).WithValue(uint8(typedef.IntensityActive)),
		),
		19: factory.CreateMesgOnly(mesgnum.Hr).WithFields(
			factory.CreateField(mesgnum.Hr, fieldnum.HrTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		20: factory.CreateMesgOnly(mesgnum.Hrv).WithFields(
			factory.CreateField(mesgnum.Hrv, fieldnum.HrvTime).WithValue([]uint16{uint16(1000)}),
		),
		// Unrelated messages
		21: factory.CreateMesgOnly(mesgnum.BarometerData).WithFields(
			factory.CreateField(mesgnum.BarometerData, fieldnum.BarometerDataTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		// Special case:
		// 1. CoursePoint's Timestamp Num is 1
		// 2. Set's Timestamp Num is 254
		22: factory.CreateMesgOnly(mesgnum.CoursePoint).WithFields(
			factory.CreateField(mesgnum.CoursePoint, fieldnum.CoursePointTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		23: factory.CreateMesgOnly(mesgnum.Set).WithFields(
			factory.CreateField(mesgnum.Set, fieldnum.SetTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
	}

	ordered = []proto.Message{
		0:  mesgs[0],
		1:  mesgs[1],
		2:  mesgs[2],
		3:  mesgs[3],
		4:  mesgs[4],
		5:  mesgs[14],
		6:  mesgs[15],
		7:  mesgs[16],
		8:  mesgs[17],
		9:  mesgs[18],
		10: mesgs[20],
		11: mesgs[5],
		12: mesgs[6],
		13: mesgs[7],
		14: mesgs[8],
		15: mesgs[9],
		16: mesgs[10],
		17: mesgs[11],
		18: mesgs[12],
		19: mesgs[13],
		20: mesgs[19],
		21: mesgs[21],
		22: mesgs[22],
		23: mesgs[23],
	}
	return
}

func TestActivityCorrectness(t *testing.T) {
	mesgs, expected := newActivityMessagesWithExpectedOrder(time.Now())

	activity := filedef.NewActivity(mesgs...)
	if activity.FileId.Type != typedef.FileActivity {
		t.Fatalf("expected: %v, got: %v", typedef.FileActivity, activity.FileId.Type)
	}

	fit := activity.ToFIT(nil) // use standard factory

	if !isMessageOrdered(expected, fit.Messages, t) {
		t.Fatalf("messages order mismatch")
	}

	// ignore fields order, make the order asc, as long as the data is equal, we consider equal.
	sortFields(expected)
	sortFields(fit.Messages)

	if diff := cmp.Diff(expected, fit.Messages, valueTransformer()); diff != "" {
		t.Fatal(diff)
	}

	// Test if message is being referenced instead of copy:
	//  - Change any of message in the expected messages should not reflect on the resulting messages.
	expected[len(expected)-1].Fields[0].Value = proto.Uint32(datetime.ToUint32(time.Now()))
	if diff := cmp.Diff(expected, fit.Messages, valueTransformer()); diff == "" {
		t.Fatalf("the modification reflect on the resulting messages")
	}
}

func isMessageOrdered(expected, result []proto.Message, t *testing.T) bool {
	var ordered = true
	for i := range expected {
		if expected[i].Num != result[i].Num {
			ordered = false
			t.Logf("[%d]: expected: %s, got: %s, timestamps: [%v, %v]",
				i, expected[i].Num, result[i].Num,
				expected[i].FieldValueByNum(proto.FieldNumTimestamp).Any(),
				result[i].FieldValueByNum(proto.FieldNumTimestamp).Any())
			continue
		}

		t.Logf("[%d]: OK! (%s), timestamp: %v", expected[i].Num,
			expected[i].Num, expected[i].FieldValueByNum(proto.FieldNumTimestamp).Any())
	}
	return ordered
}
