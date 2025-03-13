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

func newTotalsMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileTotals)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.DeveloperDataId, Fields: []proto.Field{
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.FieldDescription, Fields: []proto.Field{
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		}},
		{Num: mesgnum.Totals, Fields: []proto.Field{
			factory.CreateField(mesgnum.Totals, fieldnum.TotalsSport).WithValue(uint8(typedef.SportSoccer)),
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

func TestTotalsCorrectness(t *testing.T) {
	mesgs := newTotalsMessageForTest(time.Now())

	totals := filedef.NewTotals(mesgs...)
	if totals.FileId.Type != typedef.FileTotals {
		t.Fatalf("expected: %v, got: %v", typedef.FileTotals, totals.FileId.Type)
	}

	fit := totals.ToFIT(nil) // use standard factory

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
