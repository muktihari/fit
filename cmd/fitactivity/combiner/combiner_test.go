// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package combiner

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

func TestCombine(t *testing.T) {
	now := time.Now()
	tt := []struct {
		name     string
		fits     []*proto.FIT
		expected *proto.FIT
		err      error
	}{
		{
			name: "cycling + cycling",
			fits: []*proto.FIT{
				{Messages: []proto.Message{
					{Num: mesgnum.FileId, Fields: []proto.Field{
						factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 1),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(100)),
					}},
					{Num: mesgnum.SplitSummary, Fields: []proto.Field{
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunActive),
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(1)),
					}},
					{Num: mesgnum.SplitSummary, Fields: []proto.Field{
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunRest),
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(1)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 2),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 2),
						factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(2000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(3000)),
					}},
					{Num: mesgnum.Activity, Fields: []proto.Field{
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityAutoMultiSport),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now) + 2),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp).WithValue(datetime.ToUint32(now.Add(7*time.Hour)) + 2),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
					}},
				}},
				{Messages: []proto.Message{
					{Num: mesgnum.FileId, Fields: []proto.Field{
						factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now) + 10),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 10),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(100)),
					}},
					{Num: mesgnum.SplitSummary, Fields: []proto.Field{
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunActive),
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(1)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 20),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 20),
						factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now) + 10),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(2000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(3000)),
					}},
					{Num: mesgnum.Activity, Fields: []proto.Field{
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityAutoMultiSport),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now) + 12),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp).WithValue(datetime.ToUint32(now.Add(7*time.Hour)) + 12),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
					}},
				}},
			},
			expected: &proto.FIT{Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 1),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(100)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 2),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 10),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(300)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(400)),
				}},
				{Num: mesgnum.SplitSummary, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, proto.FieldNumTimestamp).WithValue(datetime.ToUint32(now) + 20), // Additional
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunActive),
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(2)),
				}},
				{Num: mesgnum.SplitSummary, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, proto.FieldNumTimestamp).WithValue(datetime.ToUint32(now) + 20), // Additional
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunRest),
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(1)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(400)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(4000)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(6000 + 7000)), // gap 7000
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(6000 + 7000)),   // gap 7000
				}},
				{Num: mesgnum.Activity, Fields: []proto.Field{
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityAutoMultiSport),
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime).WithValue(uint32(19000)), // 19s
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp).WithValue(
						datetime.ToUint32(now.Add(7*time.Hour)) + 20,
					),
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
				}},
			}},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			fit, err := Combine(tc.fits)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(fit, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) interface{} { return v.Any() }),
				cmp.Transformer("Fields", func(fields []proto.Field) []proto.Field {
					slices.SortStableFunc(fields, func(f1, f2 proto.Field) int {
						if f1.Num < f2.Num {
							return -1
						} else if f1.Num > f2.Num {
							return 1
						}
						return 0
					})
					return fields
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestGetLastDistanceOrZero(t *testing.T) {
	now := time.Now()
	tt := []struct {
		name  string
		mesgs []proto.Message
		dist  uint32
	}{
		{
			name: "happy flow",
			mesgs: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now))),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now) + 1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(datetime.ToUint32(now) + 1)),
				}},
			},
			dist: 2,
		},
		{
			name: "return zero",
			mesgs: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now))),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now) + 1)),
				}},
			},
			dist: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := getLastDistanceOrZero(tc.mesgs)
			if d != tc.dist {
				t.Fatalf("expected: %d, got: %d", tc.dist, d)
			}
		})
	}
}

func TestGetFirstTimestamp(t *testing.T) {
	now := time.Now()
	tt := []struct {
		name      string
		mesgs     []proto.Message
		timestamp uint32
	}{
		{
			name: "happy flow",
			mesgs: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now))),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now) + 1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(datetime.ToUint32(now) + 1)),
				}},
			},
			timestamp: datetime.ToUint32(now),
		},
		{
			name: "return basetype.Uint32Invalid",
			mesgs: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
			},
			timestamp: basetype.Uint32Invalid,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := getFirstTimestamp(tc.mesgs)
			if d != tc.timestamp {
				t.Fatalf("expected: %d, got: %d", tc.timestamp, d)
			}
		})
	}
}

func TestGetLastTimestamp(t *testing.T) {
	now := time.Now()
	tt := []struct {
		name      string
		mesgs     []proto.Message
		timestamp uint32
	}{
		{
			name: "happy flow",
			mesgs: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now))),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(datetime.ToUint32(now) + 1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(datetime.ToUint32(now) + 2)),
				}},
			},
			timestamp: datetime.ToUint32(now) + 2,
		},
		{
			name: "return basetype.Uint32Invalid",
			mesgs: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
			},
			timestamp: basetype.Uint32Invalid,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := getLastTimestamp(tc.mesgs)
			if d != tc.timestamp {
				t.Fatalf("expected: %d, got: %d", tc.timestamp, d)
			}
		})
	}
}

func TestGetTimezone(t *testing.T) {
	now := time.Now()
	tt := []struct {
		name       string
		activities []*mesgdef.Activity
		timezone   int
	}{
		{
			name: "happy flow",
			activities: []*mesgdef.Activity{
				mesgdef.NewActivity(nil).
					SetTimestamp(now).
					SetLocalTimestamp(now.Add(7 * time.Hour)),
				mesgdef.NewActivity(nil),
			},
			timezone: 7,
		},
		{
			name: "return basetype.Uint32Invalid",
			activities: []*mesgdef.Activity{
				mesgdef.NewActivity(nil),
				mesgdef.NewActivity(nil),
			},
			timezone: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tz := getTimezone(tc.activities)
			if tz != tc.timezone {
				t.Fatalf("expected: %d, got: %d", tc.timezone, tz)
			}
		})
	}
}
