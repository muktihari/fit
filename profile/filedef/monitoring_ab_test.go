// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"testing"
	"time"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

func newMonitoringAMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileMonitoringA)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.MonitoringInfo).WithFields(
			factory.CreateField(mesgnum.MonitoringInfo, fieldnum.MonitoringInfoActivityType).WithValue([]uint8{
				uint8(typedef.ActivityTypeCycling),
				uint8(typedef.ActivityTypeRunning),
			}),
		),
		factory.CreateMesgOnly(mesgnum.Monitoring).WithFields(
			factory.CreateField(mesgnum.Monitoring, fieldnum.MonitoringActivityType).WithValue(uint8(typedef.ActivityTypeCycling)),
		),
		factory.CreateMesgOnly(mesgnum.Monitoring).WithFields(
			factory.CreateField(mesgnum.Monitoring, fieldnum.MonitoringActivityType).WithValue(uint8(typedef.ActivityTypeRunning)),
		),
		factory.CreateMesgOnly(mesgnum.DeviceInfo).WithFields(
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoBatteryStatus).WithValue(uint8(typedef.BatteryStatusGood)),
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

func newMonitoringBMessageForTest(now time.Time) []proto.Message {
	mesgsB := slices.Clone(newMonitoringAMessageForTest(now))
	ftype := mesgsB[0].FieldByNum(fieldnum.FileIdType)
	ftype.Value = proto.Uint8(uint8(typedef.FileMonitoringB))
	return mesgsB
}

func TestMonitoringABCorrectness(t *testing.T) {
	mesgsA := newMonitoringAMessageForTest(time.Now())

	monitoringA := filedef.NewMonitoringAB(mesgsA...)
	if monitoringA.FileId.Type != typedef.FileMonitoringA {
		t.Fatalf("expected: %v, got: %v", typedef.FileActivity, monitoringA.FileId.Type)
	}

	fit := monitoringA.ToFIT(nil) // use standard factory

	compare(t, mesgsA, fit.Messages)

	mesgsB := newMonitoringBMessageForTest(time.Now())
	ftype := mesgsB[0].FieldByNum(fieldnum.FileIdType)
	ftype.Value = proto.Uint8(uint8(typedef.FileMonitoringB))

	monitoringB := filedef.NewMonitoringAB(mesgsB...)
	if monitoringB.FileId.Type != typedef.FileMonitoringB {
		t.Fatalf("expected: %v, got: %v", typedef.FileMonitoringB, monitoringA.FileId.Type)
	}

	fit = monitoringB.ToFIT(nil) // use standard factory

	compare(t, mesgsB, fit.Messages)
}

func compare(t *testing.T, expected, result []proto.Message) {
	histogramExpected := map[typedef.MesgNum]int{}
	for i := range expected {
		histogramExpected[expected[i].Num]++
	}

	histogramResult := map[typedef.MesgNum]int{}
	for i := range result {
		histogramResult[result[i].Num]++
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
