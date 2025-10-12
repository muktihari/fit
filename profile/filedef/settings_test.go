// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestNewSettings(t *testing.T) {
	f := filedef.NewSettings()
	fileId := *mesgdef.NewFileId(nil)
	fileId.Type = typedef.FileSettings
	if diff := cmp.Diff(f.FileId, fileId); diff != "" {
		t.Fatal(diff)
	}
}

func newSettingsMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileSettings)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.DeveloperDataId, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.FieldDescription, Fields: []proto.Field{
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.UserProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(29)),
		}},
		{Num: mesgnum.HrmProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.HrmProfile, fieldnum.HrmProfileEnabled).WithValue(true),
		}},
		{Num: mesgnum.SdmProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.SdmProfile, fieldnum.SdmProfileEnabled).WithValue(true),
		}},
		{Num: mesgnum.BikeProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.BikeProfile, fieldnum.BikeProfileEnabled).WithValue(true),
		}},
		{Num: mesgnum.DeviceSettings, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeviceSettings, fieldnum.DeviceSettingsBacklightMode).WithValue(uint8(typedef.BacklightModeAutoBrightness)),
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

func TestSettingsCorrectness(t *testing.T) {
	mesgs := newSettingsMessageForTest(time.Now())

	settings := filedef.NewSettings(mesgs...)
	if settings.FileId.Type != typedef.FileSettings {
		t.Fatalf("expected: %v, got: %v", typedef.FileSettings, settings.FileId.Type)
	}

	fit := settings.ToFIT(nil) // use standard factory

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

	for k, expectedCount := range histogramExpected {
		if resultCount := histogramResult[k]; expectedCount != resultCount {
			t.Errorf("expected message count: %d, got: %d", expectedCount, resultCount)
		}
	}
}
