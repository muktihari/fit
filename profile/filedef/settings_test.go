// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"fmt"
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

func newSettingsMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileSettings)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.UserProfile).WithFields(
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(29)),
		),
		factory.CreateMesgOnly(mesgnum.HrmProfile).WithFields(
			factory.CreateField(mesgnum.HrmProfile, fieldnum.HrmProfileEnabled).WithValue(true),
		),
		factory.CreateMesgOnly(mesgnum.SdmProfile).WithFields(
			factory.CreateField(mesgnum.SdmProfile, fieldnum.SdmProfileEnabled).WithValue(true),
		),
		factory.CreateMesgOnly(mesgnum.BikeProfile).WithFields(
			factory.CreateField(mesgnum.BikeProfile, fieldnum.BikeProfileEnabled).WithValue(true),
		),
		factory.CreateMesgOnly(mesgnum.DeviceSettings).WithFields(
			factory.CreateField(mesgnum.DeviceSettings, fieldnum.DeviceSettingsBacklightMode).WithValue(uint8(typedef.BacklightModeAutoBrightness)),
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

func TestSettingsCorrectness(t *testing.T) {
	mesgs := newSettingsMessageForTest(time.Now())

	settings := filedef.NewSettings(mesgs...)
	if settings.FileId.Type != typedef.FileSettings {
		t.Fatalf("expected: %v, got: %v", typedef.FileSettings, settings.FileId.Type)
	}

	fit := settings.ToFit(nil) // use standard factory

	// ignore fields order, make the order asc, as long as the data is equal, we consider equal.
	sortFields(mesgs)
	sortFields(fit.Messages)

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
