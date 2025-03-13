// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"testing"
	"time"

	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func newBloodPressureMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileBloodPressure)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.DeveloperDataId, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.FieldDescription, Fields: []proto.Field{
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.UserProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(27)),
		}},
		{Num: mesgnum.BloodPressure, Fields: []proto.Field{
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureSystolicPressure).WithValue(uint16(110)),
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureDiastolicPressure).WithValue(uint16(80)),
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureHeartRate).WithValue(uint8(100)),
		}},
		{Num: mesgnum.DeviceInfo, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		// Unrelated messages
		{Num: mesgnum.BarometerData, Fields: []proto.Field{
			factory.CreateField(mesgnum.BarometerData, fieldnum.BarometerDataTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
		{Num: mesgnum.CoursePoint, Fields: []proto.Field{
			factory.CreateField(mesgnum.CoursePoint, fieldnum.CoursePointTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		}},
	}
}

func TestBloodPressureCorrectness(t *testing.T) {
	mesgs := newBloodPressureMessageForTest(time.Now())

	bloodPressure := filedef.NewBloodPressure(mesgs...)
	if bloodPressure.FileId.Type != typedef.FileBloodPressure {
		t.Fatalf("expected: %v, got: %v", typedef.FileBloodPressure, bloodPressure.FileId.Type)
	}

	fit := bloodPressure.ToFIT(nil) // use standard factory

	histogramExpected := map[typedef.MesgNum]int{}
	for i := range mesgs {
		histogramExpected[mesgs[i].Num]++
	}

	histogramResult := map[typedef.MesgNum]int{}
	for i := range fit.Messages {
		histogramResult[fit.Messages[i].Num]++
	}

	if len(histogramExpected) != len(histogramResult) {
		t.Fatalf("expected len: %d, got: %d", len(histogramExpected), len(histogramResult))
	}

	for key, expectedCount := range histogramExpected {
		if resultCount := histogramResult[key]; expectedCount != resultCount {
			t.Errorf("expected message count: %d, got: %d", expectedCount, resultCount)
		}
	}
}
