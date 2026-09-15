// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package combiner

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"slices"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
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
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(100)),
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
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
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
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(100)),
					}},
					{Num: mesgnum.SplitSummary, Fields: []proto.Field{
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunActive),
						factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(1)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 20),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
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
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(100)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 2),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(200)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 10),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(300)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(300)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(400)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(400)),
				}},
				{Num: mesgnum.Sport, Fields: []proto.Field{
					factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(typedef.SportCycling),
					factory.CreateField(mesgnum.Sport, fieldnum.SportSubSport).WithValue(typedef.SubSportGeneric),
					factory.CreateField(mesgnum.Sport, fieldnum.SportName).WithValue("Cycling"),
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
					factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(400)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(4000)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(6000 + 7000)), // gap 7000
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(6000)),          // exclude pauses and gap
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
		{
			name: "cycling + cycling + running",
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
					{Num: mesgnum.Sport, Fields: []proto.Field{
						factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Sport, fieldnum.SportSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Sport, fieldnum.SportName).WithValue("Cycling"),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
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
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
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
				{Messages: []proto.Message{
					{Num: mesgnum.FileId, Fields: []proto.Field{
						factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now) + 100),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 100),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(100)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 200),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportRunning),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportTrail),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Running"),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 200),
						factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now) + 100),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(2000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(3000)),
					}},
					{Num: mesgnum.Activity, Fields: []proto.Field{
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityAutoMultiSport),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now) + 120),
						factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp).WithValue(datetime.ToUint32(now.Add(7*time.Hour)) + 120),
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
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 100),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(500)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 200),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(600)),
				}},
				{Num: mesgnum.Sport, Fields: []proto.Field{
					factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(typedef.SportCycling),
					factory.CreateField(mesgnum.Sport, fieldnum.SportSubSport).WithValue(typedef.SubSportGeneric),
					factory.CreateField(mesgnum.Sport, fieldnum.SportName).WithValue("Cycling"),
				}},
				{Num: mesgnum.Sport, Fields: []proto.Field{
					factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(typedef.SportRunning),
					factory.CreateField(mesgnum.Sport, fieldnum.SportSubSport).WithValue(typedef.SubSportTrail),
					factory.CreateField(mesgnum.Sport, fieldnum.SportName).WithValue("Running"),
				}},
				{Num: mesgnum.SplitSummary, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, proto.FieldNumTimestamp).WithValue(datetime.ToUint32(now) + 200), // Additional
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunActive),
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(2)),
				}},
				{Num: mesgnum.SplitSummary, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, proto.FieldNumTimestamp).WithValue(datetime.ToUint32(now) + 200), // Additional
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummarySplitType).WithValue(typedef.SplitTypeRunRest),
					factory.CreateField(mesgnum.SplitSummary, fieldnum.SplitSummaryNumSplits).WithValue(uint16(1)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 200),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(400)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(4000)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(6000 + 7000)), // gap 7000
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(6000)),          // exclude pauses and gap
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportRunning),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportTrail),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Running"),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 200),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now) + 100),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(200)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(2000)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(3000)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(3000)),
				}},
				{Num: mesgnum.Activity, Fields: []proto.Field{
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityAutoMultiSport),
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime).WithValue(uint32(199000)), // 19s
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now) + 200),
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp).WithValue(
						datetime.ToUint32(now.Add(7*time.Hour)) + 200,
					),
					factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(2)),
				}},
			}},
		},
		{
			name: "total training effect",
			fits: []*proto.FIT{
				{Messages: []proto.Message{
					{Num: mesgnum.FileId, Fields: []proto.Field{
						factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 1),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(100)),
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(100)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 2),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 2),
						factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(2000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTrainingEffect).WithValue(uint8(20)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalAnaerobicTrainingEffect).WithValue(uint8(5)),
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
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(100)),
					}},
					{Num: mesgnum.Record, Fields: []proto.Field{
						factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 20),
						factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(200)),
					}},
					{Num: mesgnum.Session, Fields: []proto.Field{
						factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
						factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 20),
						factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now) + 10),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(200)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(2000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(3000)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTrainingEffect).WithValue(uint8(30)),
						factory.CreateField(mesgnum.Session, fieldnum.SessionTotalAnaerobicTrainingEffect).WithValue(uint8(10)),
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
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(100)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 2),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(200)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(200)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 10),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(300)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(300)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(400)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordAccumulatedPower).WithValue(uint32(400)),
				}},
				{Num: mesgnum.Sport, Fields: []proto.Field{
					factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(typedef.SportCycling),
					factory.CreateField(mesgnum.Sport, fieldnum.SportSubSport).WithValue(typedef.SubSportGeneric),
					factory.CreateField(mesgnum.Sport, fieldnum.SportName).WithValue("Cycling"),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportGeneric),
					factory.CreateField(mesgnum.Session, fieldnum.SessionSportProfileName).WithValue("Cycling"),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now) + 20),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(400)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(4000)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(6000 + 7000)), // gap 7000
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(6000)),          // exclude pauses and gap
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTrainingEffect).WithValue(uint8(30)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalAnaerobicTrainingEffect).WithValue(uint8(10)),
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

func TestCombineMultisportSessions(t *testing.T) {
	type sessionSummary struct {
		Sport          typedef.Sport
		TotalTimerTime uint32
	}
	tt := []struct {
		name     string
		sports   [][]typedef.Sport
		expected []sessionSummary
	}{
		{
			name: "different boundary sports",
			sports: [][]typedef.Sport{
				{typedef.SportSwimming, typedef.SportCycling},
				{typedef.SportRunning, typedef.SportWalking},
			},
			expected: []sessionSummary{
				{typedef.SportSwimming, 60000}, {typedef.SportCycling, 60000},
				{typedef.SportRunning, 60000}, {typedef.SportWalking, 60000},
			},
		},
		{
			name: "matching boundary after different sports",
			sports: [][]typedef.Sport{
				{typedef.SportSwimming},
				{typedef.SportCycling, typedef.SportRunning},
				{typedef.SportRunning, typedef.SportWalking},
			},
			expected: []sessionSummary{
				{typedef.SportSwimming, 60000}, {typedef.SportCycling, 60000},
				{typedef.SportRunning, 120000}, {typedef.SportWalking, 60000},
			},
		},
		{
			name: "multiple different boundaries",
			sports: [][]typedef.Sport{
				{typedef.SportSwimming},
				{typedef.SportCycling, typedef.SportRunning},
				{typedef.SportWalking, typedef.SportHiking},
			},
			expected: []sessionSummary{
				{typedef.SportSwimming, 60000}, {typedef.SportCycling, 60000},
				{typedef.SportRunning, 60000}, {typedef.SportWalking, 60000},
				{typedef.SportHiking, 60000},
			},
		},
		{
			name: "matching boundary sports",
			sports: [][]typedef.Sport{
				{typedef.SportSwimming, typedef.SportCycling},
				{typedef.SportCycling, typedef.SportRunning},
			},
			expected: []sessionSummary{
				{typedef.SportSwimming, 60000}, {typedef.SportCycling, 120000},
				{typedef.SportRunning, 60000},
			},
		},
		{
			name:   "single session after different sport",
			sports: [][]typedef.Sport{{typedef.SportCycling}, {typedef.SportRunning}},
			expected: []sessionSummary{
				{typedef.SportCycling, 60000}, {typedef.SportRunning, 60000},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
			var fits []*proto.FIT
			for _, sports := range tc.sports {
				fit := &proto.FIT{Messages: []proto.Message{
					mesgdef.NewFileId(nil).SetType(typedef.FileActivity).SetTimeCreated(start).ToMesg(nil),
				}}
				for _, sport := range sports {
					end := start.Add(time.Minute)
					fit.Messages = append(fit.Messages,
						mesgdef.NewRecord(nil).SetTimestamp(start).ToMesg(nil),
						mesgdef.NewRecord(nil).SetTimestamp(end).ToMesg(nil),
						mesgdef.NewSession(nil).SetSport(sport).SetStartTime(start).SetTimestamp(end).
							SetTotalElapsedTime(60000).SetTotalTimerTime(60000).ToMesg(nil),
					)
					start = end
				}
				fits = append(fits, fit)
			}

			fit, err := Combine(fits)
			if err != nil {
				t.Fatal(err)
			}
			var sessions []sessionSummary
			var activity *mesgdef.Activity
			for _, mesg := range fit.Messages {
				switch mesg.Num {
				case mesgnum.Session:
					session := mesgdef.NewSession(&mesg)
					sessions = append(sessions, sessionSummary{session.Sport, session.TotalTimerTime})
				case mesgnum.Activity:
					activity = mesgdef.NewActivity(&mesg)
				}
			}
			if diff := cmp.Diff(tc.expected, sessions); diff != "" {
				t.Errorf("sessions (-want +got):\n%s", diff)
			}
			if activity == nil {
				t.Fatal("missing activity summary")
			}
			if want := uint16(len(tc.expected)); activity.NumSessions != want {
				t.Errorf("activity.NumSessions = %d, want %d", activity.NumSessions, want)
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
