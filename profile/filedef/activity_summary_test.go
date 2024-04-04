// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
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

func newActivitySummaryMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivitySummary)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.Lap).WithFields(
			factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(datetime.ToUint32(now)), // intentionally using same timestamp as last messag)e
		),
		factory.CreateMesgOnly(mesgnum.Session).WithFields(
			factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
		),
		factory.CreateMesgOnly(mesgnum.Activity).WithFields(
			factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
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

func isMessageOrdered(in, out []proto.Message, t *testing.T) bool {
	var ordered = true
	for i := range in {
		if in[i].Num != out[i].Num {
			ordered = false
			t.Logf("mesg order[%d]: expected: %s, got: %s", i, out[i].Num, in[i].Num)
			continue
		}
		for j := range in[i].Fields {
			num1 := in[i].Fields[j].Num
			num2 := out[i].Fields[j].Num
			if num1 != num2 {
				ordered = false
				t.Logf("field order[%d][%d]: expected: %d, got: %d", i, j, num1, num2)
			}
		}
	}
	return ordered
}

func TestActivitySummaryCorrectness(t *testing.T) {
	mesgs := newActivitySummaryMessageForTest(time.Now())

	activitySummary := filedef.NewActivitySummary(mesgs...)
	if activitySummary.FileId.Type != typedef.FileActivitySummary {
		t.Fatalf("expected: %v, got: %v", typedef.FileActivitySummary, activitySummary.FileId.Type)
	}

	fit := activitySummary.ToFit(nil) // use standard factory

	// ignore fields order, make the order asc, as long as the data is equal, we consider equal.
	sortFields(mesgs)
	sortFields(fit.Messages)

	if !isMessageOrdered(mesgs, fit.Messages, t) {
		t.Fatalf("messages order mismatch")
	}

	if diff := cmp.Diff(mesgs, fit.Messages, valueTransformer()); diff != "" {
		t.Fatal(diff)
	}
}
