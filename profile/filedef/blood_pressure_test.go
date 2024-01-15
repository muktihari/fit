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

func newBloodPressureMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileBloodPressure)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.UserProfile).WithFields(
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileAge).WithValue(uint8(27)),
		),
		factory.CreateMesgOnly(mesgnum.BloodPressure).WithFields(
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureSystolicPressure).WithValue(uint16(110)),
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureDiastolicPressure).WithValue(uint16(80)),
			factory.CreateField(mesgnum.BloodPressure, fieldnum.BloodPressureHeartRate).WithValue(uint8(100)),
		),
		factory.CreateMesgOnly(mesgnum.DeviceInfo).WithFields(
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
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

func TestBloodPressureCorrectness(t *testing.T) {
	mesgs := newBloodPressureMessageForTest(time.Now())

	bloodPressure := filedef.NewBloodPressure(mesgs...)
	if bloodPressure.FileId.Type != typedef.FileBloodPressure {
		t.Fatalf("expected: %v, got: %v", typedef.FileBloodPressure, bloodPressure.FileId.Type)
	}

	fit := bloodPressure.ToFit(nil) // use standard factory

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
