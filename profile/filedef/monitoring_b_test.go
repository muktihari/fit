// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/proto"
)

func TestNewMonitoringB(t *testing.T) {
	f := filedef.NewMonitoringB()
	fileId := *mesgdef.NewFileId(nil)
	fileId.Type = typedef.FileMonitoringB
	if diff := cmp.Diff(f.FileId, fileId); diff != "" {
		t.Fatal(diff)
	}
}

func newMonitoringBMessageForTest(now time.Time) []proto.Message {
	mesgsB := newMonitoringAMessageForTest(now)
	ftype := mesgsB[0].FieldByNum(fieldnum.FileIdType)
	ftype.Value = proto.Uint8(uint8(typedef.FileMonitoringB))
	return mesgsB
}

func TestMonitoringBCorrectness(t *testing.T) {
	mesgsB := newMonitoringBMessageForTest(time.Now())
	ftype := mesgsB[0].FieldByNum(fieldnum.FileIdType)
	ftype.Value = proto.Uint8(uint8(typedef.FileMonitoringB))

	monitoringB := filedef.NewMonitoringB(mesgsB...)
	if monitoringB.FileId.Type != typedef.FileMonitoringB {
		t.Fatalf("expected: %v, got: %v", typedef.FileMonitoringB, monitoringB.FileId.Type)
	}

	fit := monitoringB.ToFIT(nil)

	histogramExpected := map[typedef.MesgNum]int{}
	for i := range mesgsB {
		histogramExpected[mesgsB[i].Num]++
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
