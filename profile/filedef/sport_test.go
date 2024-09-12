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
)

func newSportMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileSport)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.DeveloperDataId, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.FieldDescription, Fields: []proto.Field{
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.ZonesTarget, Fields: []proto.Field{
			factory.CreateField(mesgnum.ZonesTarget, fieldnum.ZonesTargetMaxHeartRate).WithValue(uint8(190)),
		}},
		{Num: mesgnum.Sport, Fields: []proto.Field{
			factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(uint8(typedef.SportAmericanFootball)),
		}},
		{Num: mesgnum.HrZone, Fields: []proto.Field{
			factory.CreateField(mesgnum.HrZone, fieldnum.HrZoneHighBpm).WithValue(uint8(177)),
		}},
		{Num: mesgnum.PowerZone, Fields: []proto.Field{
			factory.CreateField(mesgnum.PowerZone, fieldnum.PowerZoneHighValue).WithValue(uint16(200)),
		}},
		{Num: mesgnum.MetZone, Fields: []proto.Field{
			factory.CreateField(mesgnum.MetZone, fieldnum.MetZoneHighBpm).WithValue(uint8(178)),
		}},
		{Num: mesgnum.SpeedZone, Fields: []proto.Field{
			factory.CreateField(mesgnum.SpeedZone, fieldnum.SpeedZoneHighValue).WithValue(uint16(10000)),
		}},
		{Num: mesgnum.CadenceZone, Fields: []proto.Field{
			factory.CreateField(mesgnum.CadenceZone, fieldnum.CadenceZoneHighValue).WithValue(uint8(100)),
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

func TestSportCorrectness(t *testing.T) {
	mesgs := newSportMessageForTest(time.Now())

	sport := filedef.NewSport(mesgs...)
	if sport.FileId.Type != typedef.FileSport {
		t.Fatalf("expected: %v, got: %v", typedef.FileSport, sport.FileId.Type)
	}

	fit := sport.ToFIT(nil) // use standard factory

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
