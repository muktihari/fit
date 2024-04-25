// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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

func newWorkoutMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileWorkout)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.Workout).WithFields(
			factory.CreateField(mesgnum.Workout, fieldnum.WorkoutSport).WithValue(uint8(typedef.SportSwimming)),
		),
		factory.CreateMesgOnly(mesgnum.WorkoutStep).WithFields(
			factory.CreateField(mesgnum.WorkoutStep, fieldnum.WorkoutStepEquipment).WithValue(uint8(typedef.WorkoutEquipmentSwimFins)),
		),
		factory.CreateMesgOnly(mesgnum.WorkoutStep).WithFields(
			factory.CreateField(mesgnum.WorkoutStep, fieldnum.WorkoutStepEquipment).WithValue(uint8(typedef.WorkoutEquipmentSwimSnorkel)),
		),
		// Unrelated messages
		factory.CreateMesgOnly(mesgnum.CoursePoint).WithFields(
			factory.CreateField(mesgnum.CoursePoint, fieldnum.CoursePointTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.BarometerData).WithFields(
			factory.CreateField(mesgnum.BarometerData, fieldnum.BarometerDataTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
	}
}

func TestWorkoutCorrectness(t *testing.T) {
	mesgs := newWorkoutMessageForTest(time.Now())

	workout := filedef.NewWorkout(mesgs...)
	if workout.FileId.Type != typedef.FileWorkout {
		t.Fatalf("expected: %v, got: %v", typedef.FileActivity, workout.FileId.Type)
	}

	fit := workout.ToFIT(nil)

	// ignore fields order, make the order asc, as long as the data is equal, we consider equal.
	sortFields(mesgs)
	sortFields(fit.Messages)

	if diff := cmp.Diff(mesgs, fit.Messages, valueTransformer()); diff != "" {
		fmt.Println("messages order:")
		for i := range fit.Messages {
			mesg := fit.Messages[i]
			fmt.Printf("%d: %s\n", mesg.Num, mesg.Num)
		}
		fmt.Println("")
		t.Fatal(diff)
	}

	// Edit unrelated message, should not change the resulting messages.
	mesgs[len(mesgs)-1].Fields[0].Value = proto.Uint32(datetime.ToUint32(time.Now()))
	if diff := cmp.Diff(mesgs, fit.Messages, valueTransformer()); diff == "" {
		t.Fatalf("the modification reflect on the resulting messages")
	}
}
