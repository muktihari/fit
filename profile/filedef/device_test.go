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

func newDeviceMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileDevice)),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		),
		factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
			factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
			factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
		),
		factory.CreateMesgOnly(mesgnum.Software).WithFields(
			factory.CreateField(mesgnum.Software, fieldnum.SoftwareMessageIndex).WithValue(uint16(typedef.MessageIndexReserved)),
		),
		factory.CreateMesgOnly(mesgnum.Capabilities).WithFields(
			factory.CreateField(mesgnum.Capabilities, fieldnum.CapabilitiesSports).WithValue([]uint8{
				uint8(typedef.SportBits0Basketball),
				uint8(typedef.SportBits1AmericanFootball),
				uint8(typedef.SportBits2Paddling),
			}),
		),
		factory.CreateMesgOnly(mesgnum.FileCapabilities).WithFields(
			factory.CreateField(mesgnum.FileCapabilities, fieldnum.FileCapabilitiesType).WithValue(uint8(typedef.FileActivity)),
		),
		factory.CreateMesgOnly(mesgnum.MesgCapabilities).WithFields(
			factory.CreateField(mesgnum.MesgCapabilities, fieldnum.MesgCapabilitiesFile).WithValue(uint8(typedef.FileActivity)),
		),
		factory.CreateMesgOnly(mesgnum.FieldCapabilities).WithFields(
			factory.CreateField(mesgnum.FieldCapabilities, fieldnum.FieldCapabilitiesFile).WithValue(uint8(typedef.FileActivity)),
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

func TestDeviceCorrectness(t *testing.T) {
	mesgs := newDeviceMessageForTest(time.Now())

	device := filedef.NewDevice(mesgs...)
	if device.FileId.Type != typedef.FileDevice {
		t.Fatalf("expected: %v, got: %v", typedef.FileDevice, device.FileId.Type)
	}

	fit := device.ToFIT(nil) // use standard factory

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
