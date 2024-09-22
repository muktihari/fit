// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"math"
	"testing"
	"time"

	"slices"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
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
		// Messages have no timestamp field:
		0: {Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivity)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		1: {Num: mesgnum.DeveloperDataId, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		}},
		2: {Num: mesgnum.FieldDescription, Fields: []proto.Field{
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue([]string{"Heart Rate"}),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
		}},
		3: {Num: mesgnum.UserProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileFriendlyName).WithValue("Mary Jane"),
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(21)),
		}},
		4: {Num: mesgnum.ZonesTarget, Fields: []proto.Field{
			factory.CreateField(mesgnum.ZonesTarget, fieldnum.ZonesTargetMaxHeartRate).WithValue(uint8(190)),
		}},
		5: {Num: mesgnum.Workout, Fields: []proto.Field{
			factory.CreateField(mesgnum.Workout, fieldnum.WorkoutSport).WithValue(uint8(typedef.SportCycling)),
		}},
		6: {Num: mesgnum.WorkoutStep, Fields: []proto.Field{
			factory.CreateField(mesgnum.WorkoutStep, fieldnum.WorkoutStepIntensity).WithValue(uint8(typedef.IntensityActive)),
		}},
		7: {Num: mesgnum.Hrv, Fields: []proto.Field{
			factory.CreateField(mesgnum.Hrv, fieldnum.HrvTime).WithValue([]uint16{uint16(1000)}),
		}},
		8: {Num: mesgnum.Split, Fields: []proto.Field{
			factory.CreateField(mesgnum.Split, fieldnum.SplitTotalDistance).WithValue(uint32(10000)),
		}},
		9: {Num: mesgnum.SplitSummary, Fields: []proto.Field{
			factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeAscentSplit.Byte()),
		}},
		10: {Num: mesgnum.Sport, Fields: []proto.Field{
			factory.CreateField(mesgnum.Sport, fieldnum.SportName).WithValue("cycling"),
		}},
		// Messages have timestamp field:
		11: {Num: mesgnum.DeviceInfo, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
		}},
		12: {Num: mesgnum.Event, Fields: []proto.Field{
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
			factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStart)),
		}},
		13: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		14: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		15: {Num: mesgnum.Event, Fields: []proto.Field{
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		16: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		17: {Num: mesgnum.Event, Fields: []proto.Field{
			// Intentionally using same timestamp as last message.
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
		}},
		18: {Num: mesgnum.Lap, Fields: []proto.Field{
			factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		19: {Num: mesgnum.Session, Fields: []proto.Field{
			factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		20: {Num: mesgnum.Activity, Fields: []proto.Field{
			factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		21: {Num: mesgnum.Length, Fields: []proto.Field{
			factory.CreateField(mesgnum.Length, fieldnum.LengthTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.Length, fieldnum.LengthAvgSpeed).WithValue(uint16(1000)),
		}},
		22: {Num: mesgnum.SegmentLap, Fields: []proto.Field{
			factory.CreateField(mesgnum.SegmentLap, fieldnum.SegmentLapTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.SegmentLap, fieldnum.SegmentLapAvgCadence).WithValue(uint8(100)),
		}},
		23: {Num: mesgnum.Hr, Fields: []proto.Field{
			factory.CreateField(mesgnum.Hr, fieldnum.HrTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		// Unrelated messages
		24: {Num: mesgnum.BarometerData, Fields: []proto.Field{
			factory.CreateField(mesgnum.BarometerData, fieldnum.BarometerDataTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		25: {Num: mesgnum.BikeProfile, Fields: []proto.Field{ // have no timestamp, should be sorted at top.
			factory.CreateField(mesgnum.BikeProfile, fieldnum.BikeProfileBikeWeight).WithValue(uint16(100)),
		}},
		// Special case:
		// 1. CoursePoint's Timestamp Num is 1
		// 2. Set's Timestamp Num is 254
		26: {Num: mesgnum.CoursePoint, Fields: []proto.Field{
			factory.CreateField(mesgnum.CoursePoint, fieldnum.CoursePointTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		27: {Num: mesgnum.Set, Fields: []proto.Field{
			factory.CreateField(mesgnum.Set, fieldnum.SetTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		// Optional messages:
		28: {Num: mesgnum.GpsMetadata, Fields: []proto.Field{
			factory.CreateField(mesgnum.GpsMetadata, fieldnum.GpsMetadataTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		29: {Num: mesgnum.TimeInZone, Fields: []proto.Field{
			factory.CreateField(mesgnum.TimeInZone, fieldnum.TimeInZoneTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
	}

	ordered = []proto.Message{
		0:  mesgs[0],
		1:  mesgs[1],
		2:  mesgs[2],
		3:  mesgs[3],
		4:  mesgs[4],
		5:  mesgs[5],
		6:  mesgs[6],
		7:  mesgs[7],
		8:  mesgs[8],
		9:  mesgs[9],
		10: mesgs[10],
		11: mesgs[25],
		12: mesgs[11],
		13: mesgs[12],
		14: mesgs[13],
		15: mesgs[14],
		16: mesgs[15],
		17: mesgs[16],
		18: mesgs[17],
		19: mesgs[18],
		20: mesgs[19],
		21: mesgs[20],
		22: mesgs[21],
		23: mesgs[22],
		24: mesgs[23],
		25: mesgs[24],
		26: mesgs[26],
		27: mesgs[27],
		28: mesgs[28],
		29: mesgs[29],
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
