// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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

func newSportMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileSport)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.ZonesTarget).WithFields(
			factory.CreateField(mesgnum.ZonesTarget, fieldnum.ZonesTargetMaxHeartRate).WithValue(uint8(190)),
		),
		factory.CreateMesgOnly(mesgnum.Sport).WithFields(
			factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(uint8(typedef.SportAmericanFootball)),
		),
		factory.CreateMesgOnly(mesgnum.HrZone).WithFields(
			factory.CreateField(mesgnum.HrZone, fieldnum.HrZoneHighBpm).WithValue(uint8(177)),
		),
		factory.CreateMesgOnly(mesgnum.PowerZone).WithFields(
			factory.CreateField(mesgnum.PowerZone, fieldnum.PowerZoneHighValue).WithValue(uint16(200)),
		),
		factory.CreateMesgOnly(mesgnum.MetZone).WithFields(
			factory.CreateField(mesgnum.MetZone, fieldnum.MetZoneHighBpm).WithValue(uint8(178)),
		),
		factory.CreateMesgOnly(mesgnum.SpeedZone).WithFields(
			factory.CreateField(mesgnum.SpeedZone, fieldnum.SpeedZoneHighValue).WithValue(uint16(10000)),
		),
		factory.CreateMesgOnly(mesgnum.CadenceZone).WithFields(
			factory.CreateField(mesgnum.CadenceZone, fieldnum.CadenceZoneHighValue).WithValue(uint8(100)),
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

func TestSportCorrectness(t *testing.T) {
	mesgs := newSportMessageForTest(time.Now())

	sport := filedef.NewSport(mesgs...)
	if sport.FileId.Type != typedef.FileSport {
		t.Fatalf("expected: %v, got: %v", typedef.FileSport, sport.FileId.Type)
	}

	fit := sport.ToFit(nil) // use standard factory

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
