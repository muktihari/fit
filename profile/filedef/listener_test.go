// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"math"
	"reflect"
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

func createFloat32Comparer() cmp.Option {
	return cmp.Comparer(func(f1, f2 float32) bool {
		// Compare float in integer form. since when f is NaN, f != f.
		return math.Float32bits(f1) == math.Float32bits(f2)
	})
}

func createFloat64Comparer() cmp.Option {
	return cmp.Comparer(func(f1, f2 float64) bool {
		// Compare float in integer form. since when f is NaN, f != f.
		return math.Float64bits(f1) == math.Float64bits(f2)
	})
}

type dummyFile struct{}

func (dummyFile) Add(mesg proto.Message)             {}
func (dummyFile) ToFIT(o *mesgdef.Options) proto.FIT { return proto.FIT{} }

var _ filedef.File = (*dummyFile)(nil)

func TestListenerForSingleFitFile(t *testing.T) {
	type table struct {
		name    string
		options []filedef.Option
		mesgs   []proto.Message
		result  filedef.File
	}

	now := time.Now()
	tt := []table{
		{
			name:   "default listener for activity",
			mesgs:  newActivityMessageForTest(now),
			result: filedef.NewActivity(newActivityMessageForTest(now)...),
		},
		{
			name:   "default listener for activity summary",
			mesgs:  newActivitySummaryMessageForTest(now),
			result: filedef.NewActivitySummary(newActivitySummaryMessageForTest(now)...),
		},
		{
			name:   "default listener for blood pressure",
			mesgs:  newBloodPressureMessageForTest(now),
			result: filedef.NewBloodPressure(newBloodPressureMessageForTest(now)...),
		},
		{
			name:   "default listener for course",
			mesgs:  newCourseMessageForTest(now),
			result: filedef.NewCourse(newCourseMessageForTest(now)...),
		},
		{
			name:   "default listener for device",
			mesgs:  newDeviceMessageForTest(now),
			result: filedef.NewDevice(newDeviceMessageForTest(now)...),
		},
		{
			name:   "default listener for goals",
			mesgs:  newGoalsMessageForTest(now),
			result: filedef.NewGoals(newGoalsMessageForTest(now)...),
		},
		{
			name:   "default listener for monitoring A",
			mesgs:  newMonitoringAMessageForTest(now),
			result: filedef.NewMonitoringA(newMonitoringAMessageForTest(now)...),
		},
		{
			name:   "default listener for monitoring B",
			mesgs:  newMonitoringBMessageForTest(now),
			result: filedef.NewMonitoringB(newMonitoringBMessageForTest(now)...),
		},
		{
			name:   "default listener for monitoring daily",
			mesgs:  newMonitoringDailyMessageForTest(now),
			result: filedef.NewMonitoringDaily(newMonitoringDailyMessageForTest(now)...),
		},
		{
			name:   "default listener for schedules",
			mesgs:  newSchedulesMessageForTest(now),
			result: filedef.NewSchedules(newSchedulesMessageForTest(now)...),
		},
		{
			name:   "default listener for segment",
			mesgs:  newSegmentMessageForTest(now),
			result: filedef.NewSegment(newSegmentMessageForTest(now)...),
		},
		{
			name:   "default listener for segment list",
			mesgs:  newSegmentListMessageForTest(now),
			result: filedef.NewSegmentList(newSegmentListMessageForTest(now)...),
		},
		{
			name:   "default listener for settings",
			mesgs:  newSettingsMessageForTest(now),
			result: filedef.NewSettings(newSettingsMessageForTest(now)...),
		},
		{
			name:   "default listener for sport",
			mesgs:  newSportMessageForTest(now),
			result: filedef.NewSport(newSportMessageForTest(now)...),
		},
		{
			name:   "default listener for totals",
			mesgs:  newTotalsMessageForTest(now),
			result: filedef.NewTotals(newTotalsMessageForTest(now)...),
		},
		{
			name:   "default listener for weight",
			mesgs:  newWeightMessageForTest(now),
			result: filedef.NewWeight(newWeightMessageForTest(now)...),
		},
		{
			name:   "default listener for workout",
			mesgs:  newWorkoutMessageForTest(now),
			result: filedef.NewWorkout(newWorkoutMessageForTest(now)...),
		},
		{
			name:   "replace activity with dummy file; PredefinedFileSet",
			mesgs:  newActivityMessageForTest(now),
			result: new(dummyFile),
			options: func() []filedef.Option {
				sets := filedef.PredefinedFileSet()
				sets[typedef.FileActivity] = func() filedef.File { return new(dummyFile) }
				return []filedef.Option{filedef.WithFileSets(sets)}
			}(),
		},
		{
			name:   "replace activity with dummy file; WithFileFunc",
			mesgs:  newActivityMessageForTest(now),
			result: new(dummyFile),
			options: []filedef.Option{
				filedef.WithFileFunc(typedef.FileActivity, func() filedef.File { return new(dummyFile) }),
			},
		},
		{
			name: "listener for not specified fileset, course",
			options: []filedef.Option{
				filedef.WithFileSets(map[typedef.File]func() filedef.File{
					typedef.FileActivity: func() filedef.File { return filedef.NewActivity() },
				}),
				filedef.WithChannelBuffer(100),
			},
			mesgs:  newWorkoutMessageForTest(now),
			result: nil,
		},
		func() table {
			mesgs := newActivityMessageForTest(now)
			mesgs = append(mesgs,
				proto.Message{Num: mesgnum.Record,
					Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(incrementSecond(&now))),
					},
					DeveloperFields: []proto.DeveloperField{
						{
							DeveloperDataIndex: 0,
							Num:                0,
							Value:              proto.Uint8(100),
						},
					}},
			)
			return table{
				name:   "default listener for activity containing developer fields",
				mesgs:  mesgs,
				result: filedef.NewActivity(mesgs...),
			}
		}(),
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lis := filedef.NewListener(tc.options...)
			defer lis.Close()

			for _, mesg := range tc.mesgs {
				lis.OnMesg(mesg)
			}

			result := lis.File()

			if diff := cmp.Diff(tc.result, result,
				createFloat32Comparer(),
				createFloat64Comparer(),
				valueTransformer(),
				cmp.Exporter(func(t reflect.Type) bool { return true }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestListenerForChainedFitFile(t *testing.T) {
	now := time.Now()

	// Simulate a chained FIT file that have 3 File different FIT file type.
	activityMesgs := newActivityMessageForTest(now)
	courseMesgs := newCourseMessageForTest(now)
	workoutMesgs := newWorkoutMessageForTest(now)

	expectedResult := []filedef.File{
		filedef.NewActivity(activityMesgs...),
		filedef.NewCourse(courseMesgs...),
		filedef.NewWorkout(workoutMesgs...),
	}

	lis := filedef.NewListener()
	defer lis.Close()

	result := make([]filedef.File, 0, len(expectedResult))

	for _, mesg := range activityMesgs {
		lis.OnMesg(mesg)
	}
	result = append(result, lis.File())

	for _, mesg := range courseMesgs {
		lis.OnMesg(mesg)
	}
	result = append(result, lis.File())

	for _, mesg := range workoutMesgs {
		lis.OnMesg(mesg)
	}
	result = append(result, lis.File())

	if diff := cmp.Diff(expectedResult, result,
		createFloat32Comparer(),
		createFloat64Comparer(),
		valueTransformer(),
		cmp.Exporter(func(t reflect.Type) bool { return true }),
	); diff != "" {
		t.Fatal(diff)
	}
}

func TestClose(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("expected not panic, got: %v", err)
		}
	}()

	l := filedef.NewListener()
	l.Close()
	l.Close() // already closed, should not panic
}

func TestReset(t *testing.T) {
	l := filedef.NewListener()
	l.Reset(filedef.WithChannelBuffer(256))
}
