package concealer

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestConceal(t *testing.T) {
	tt := []struct {
		name           string
		mesgs          []proto.Message
		startThreshold uint32
		endThreshold   uint32
		expected       []proto.Message
	}{
		{
			name: "conceal overlap (full conceal), 4 records 2 laps, conceal up to 3rd record, 1st lap fully concealed, 2nd lap partially, session changed",
			mesgs: []proto.Message{
				0: {Num: mesgnum.FileId},
				1: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				2: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				3: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(1)),
				}},
				4: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				5: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				6: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				7: {Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(4)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLong).WithValue(int32(3)),
				}},
				8: {Num: mesgnum.Activity},
			},
			startThreshold: 3,
			endThreshold:   3,
			expected: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(4)),
				}},
				{Num: mesgnum.Activity},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			Conceal(tc.mesgs, tc.startThreshold, tc.endThreshold)
			if diff := cmp.Diff(tc.mesgs, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) interface{} { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestConcealStartPosition(t *testing.T) {
	tt := []struct {
		name           string
		mesgs          []proto.Message
		lapIndices     []int
		sessionIndices []int
		threshold      uint32
		expectedIndex  int
		expected       []proto.Message
	}{
		{name: "threshold zero", threshold: 0},
		{
			name: "4 records 2 laps, conceal up to 3rd record, 1st lap fully concealed, 2nd lap partially, session changed",
			mesgs: []proto.Message{
				0: {Num: mesgnum.FileId},
				1: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				2: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				3: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(1)),
				}},
				4: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				5: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				6: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				7: {Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(4)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLong).WithValue(int32(3)),
				}},
				8: {Num: mesgnum.Activity},
			},
			lapIndices:     []int{3, 6},
			sessionIndices: []int{7},
			threshold:      3,
			expectedIndex:  5,
			expected: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					// GPS Positions are fully concealed.
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(3)),  // New Start Lat
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(3)), // New Start Long
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(4)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLong).WithValue(int32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLong).WithValue(int32(3)),
				}},
				{Num: mesgnum.Activity},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			lastConcealStartIndex := concealStartPosition(tc.mesgs, tc.lapIndices, tc.sessionIndices, tc.threshold)
			if lastConcealStartIndex != tc.expectedIndex {
				t.Fatalf("expected last conceal start's record index is: %d, got: %d",
					tc.expectedIndex, lastConcealStartIndex)
			}
			if diff := cmp.Diff(tc.mesgs, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) interface{} { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestUpdateStartPosition(t *testing.T) {
	tt := []struct {
		name        string
		mesgs       []proto.Message
		indices     []int
		fieldNums   placeholder
		recordIndex int
		expected    []proto.Message
	}{
		{
			name: "laps: 4 records 2 laps, conceal up to 3rd record, 1st lap fully concealed, 2nd lap partially",
			mesgs: []proto.Message{
				0: {Num: mesgnum.FileId},
				1: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				2: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				3: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(1)),
				}},
				4: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				5: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				6: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				7: {Num: mesgnum.Activity},
			},
			indices:     []int{3, 6},
			fieldNums:   lapPlaceholder,
			recordIndex: 5,
			expected: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					// GPS Positions are fully concealed.
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(3)),  // New Start Lat
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(3)), // New Start Long
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				{Num: mesgnum.Activity},
			},
		},
		{
			name: "1 laps, 2 records all without gps",
			mesgs: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Activity},
			},
			indices:     []int{3},
			fieldNums:   lapPlaceholder,
			recordIndex: 2,
			expected: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapTotalTimerTime).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Activity},
			},
		},
	}

	for i, tc := range tt {
		if i != len(tt)-1 {
			continue
		}
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			updateStartPosition(tc.mesgs, tc.indices, tc.fieldNums, tc.recordIndex)
			if diff := cmp.Diff(tc.mesgs, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) interface{} { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestConcealEndPosition(t *testing.T) {
	tt := []struct {
		name           string
		mesgs          []proto.Message
		lapIndices     []int
		sessionIndices []int
		threshold      uint32
		expected       []proto.Message
	}{
		{name: "threshold zero", threshold: 0},
		{
			name: "laps: 4 records 2 laps, conceal up to 3rd record, 1st lap fully concealed, 2nd lap partially, session changed",
			mesgs: []proto.Message{
				0: {Num: mesgnum.FileId},
				1: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				2: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				3: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(1)),
				}},
				4: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				5: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				6: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				7: {Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(4)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLong).WithValue(int32(3)),
				}},
				8: {Num: mesgnum.Activity},
			},
			lapIndices:     []int{3, 6},
			sessionIndices: []int{7},
			threshold:      3,
			expected: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(0)),  // New Start Lat
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(0)), // New Start Long
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Session, Fields: []proto.Field{
					factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionTotalTimerTime).WithValue(uint32(4)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLong).WithValue(int32(0)),
				}},
				{Num: mesgnum.Activity},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			concealEndPosition(tc.mesgs, tc.lapIndices, tc.sessionIndices, 0, tc.threshold)
			if diff := cmp.Diff(tc.mesgs, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) interface{} { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestUpdateEndPosition(t *testing.T) {
	tt := []struct {
		name             string
		mesgs            []proto.Message
		indices          []int
		fieldNums        placeholder
		startRecordIndex int
		recordIndex      int
		expected         []proto.Message
	}{
		{
			name: "laps: 4 records 2 laps, conceal up to 3rd record, 1st lap fully concealed, 2nd lap partially",
			mesgs: []proto.Message{
				0: {Num: mesgnum.FileId},
				1: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				2: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				3: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(1)),
				}},
				4: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				5: {Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				6: {Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(2)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(3)),
				}},
				7: {Num: mesgnum.Activity},
			},
			indices:          []int{3, 6},
			fieldNums:        lapPlaceholder,
			startRecordIndex: 0,
			recordIndex:      1,
			expected: []proto.Message{
				{Num: mesgnum.FileId},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(0)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(1)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(1)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLat).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartPositionLong).WithValue(int32(0)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLat).WithValue(int32(0)),  // New Start Lat
					factory.CreateField(mesgnum.Lap, fieldnum.LapEndPositionLong).WithValue(int32(0)), // New Start Long
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(2)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(2)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(3)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(3)),
				}},
				{Num: mesgnum.Lap, Fields: []proto.Field{
					factory.CreateField(mesgnum.Lap, fieldnum.LapTimestamp).WithValue(uint32(3)),
					factory.CreateField(mesgnum.Lap, fieldnum.LapStartTime).WithValue(uint32(2)),
				}},
				{Num: mesgnum.Activity},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			updateEndPosition(tc.mesgs, tc.indices, tc.fieldNums, tc.startRecordIndex, tc.recordIndex)
			if diff := cmp.Diff(tc.mesgs, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) interface{} { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
