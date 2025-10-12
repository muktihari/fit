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

func TestNewMonitoringA(t *testing.T) {
	a := filedef.NewMonitoringA()
	fileId := *mesgdef.NewFileId(nil)
	fileId.Type = typedef.FileMonitoringA
	if diff := cmp.Diff(a.FileId, fileId); diff != "" {
		t.Fatal(diff)
	}
}

func newMonitoringAMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileMonitoringA)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.DeveloperDataId, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.FieldDescription, Fields: []proto.Field{
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.MonitoringInfo, Fields: []proto.Field{
			factory.CreateField(mesgnum.MonitoringInfo, fieldnum.MonitoringInfoActivityType).WithValue([]uint8{
				uint8(typedef.ActivityTypeCycling),
				uint8(typedef.ActivityTypeRunning),
			}),
		}},
		{Num: mesgnum.Monitoring, Fields: []proto.Field{
			factory.CreateField(mesgnum.Monitoring, fieldnum.MonitoringActivityType).WithValue(uint8(typedef.ActivityTypeCycling)),
		}},
		{Num: mesgnum.Monitoring, Fields: []proto.Field{
			factory.CreateField(mesgnum.Monitoring, fieldnum.MonitoringActivityType).WithValue(uint8(typedef.ActivityTypeRunning)),
		}},
		{Num: mesgnum.DeviceInfo, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
			factory.CreateField(mesgnum.DeviceInfo, fieldnum.DeviceInfoBatteryStatus).WithValue(uint8(typedef.BatteryStatusGood)),
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

func TestMonitoringACorrectness(t *testing.T) {
	mesgsA := newMonitoringAMessageForTest(time.Now())

	monitoringA := filedef.NewMonitoringA(mesgsA...)
	if monitoringA.FileId.Type != typedef.FileMonitoringA {
		t.Fatalf("expected: %v, got: %v", typedef.FileActivity, monitoringA.FileId.Type)
	}

	fit := monitoringA.ToFIT(nil)

	histogramExpected := map[typedef.MesgNum]int{}
	for i := range mesgsA {
		histogramExpected[mesgsA[i].Num]++
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
