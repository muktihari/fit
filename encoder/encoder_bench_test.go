// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder_test

import (
	"io"
	"testing"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type discardAt struct{}

var _ io.Writer = discardAt{}
var _ io.WriterAt = discardAt{}

func (discardAt) Write(p []byte) (int, error) {
	return len(p), nil
}

func (discardAt) WriteAt(p []byte, off int64) (n int, err error) {
	return len(p), nil
}

var DiscardAt = discardAt{}

func createFitForBenchmark(recodSize int) *proto.FIT {
	now := time.Now()
	fit := &proto.FIT{
		Messages: []proto.Message{
			{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerBryton),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdProduct).WithValue(uint16(1901)),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("Rider 420"),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdNumber).WithValue(uint16(0)),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdSerialNumber).WithValue(uint32(5122)),
			}},
			{Num: mesgnum.Sport, Fields: []proto.Field{
				factory.CreateField(mesgnum.Sport, fieldnum.SportSport).WithValue(typedef.SportCycling),
				factory.CreateField(mesgnum.Sport, fieldnum.SportSubSport).WithValue(typedef.SubSportRoad),
			}},
			{Num: mesgnum.Activity, Fields: []proto.Field{
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityTypeCycling),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime).WithValue(uint32(30877.0 * 1000)),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityEvent).WithValue(typedef.EventActivity),
			}},
			{Num: mesgnum.Session, Fields: []proto.Field{
				factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionStartTime).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionTotalElapsedTime).WithValue(uint32(30877.0 * 1000)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(32172.05 * 100)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionSport).WithValue(typedef.SportCycling),
				factory.CreateField(mesgnum.Session, fieldnum.SessionSubSport).WithValue(typedef.SubSportRoad),
				factory.CreateField(mesgnum.Session, fieldnum.SessionTotalMovingTime).WithValue(uint32(22079.0 * 1000)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionTotalCalories).WithValue(uint16(12824)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgSpeed).WithValue(uint16(5.98 * 1000)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionMaxSpeed).WithValue(uint16(13.05 * 1000)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionMaxAltitude).WithValue(uint16((504.0 + 500) * 5)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionTotalAscent).WithValue(uint16(909)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDescent).WithValue(uint16(901)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionSwcLat).WithValue(int32(0)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionSwcLong).WithValue(int32(0)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionNecLat).WithValue(int32(0)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionNecLong).WithValue(int32(0)),
			}},
		},
	}

	for i := 0; i < recodSize-len(fit.Messages); i++ {
		now = now.Add(time.Second) // only time is moving forward
		if i%100 == 0 {            // add event every 100 message
			fit.Messages = append(fit.Messages, proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStop)),
			}})
			now = now.Add(10 * time.Second) // gap
			fit.Messages = append(fit.Messages, proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStart)),
			}})
			now = now.Add(time.Second) // gap
		}

		record := proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLat).WithValue(int32(-90481372)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordPositionLong).WithValue(int32(1323227263)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(8.33 * 1000)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(405.81 * 100)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(110)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(85)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordAltitude).WithValue(uint16((166.0 + 500.0) * 5.0)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordTemperature).WithValue(int8(32)),
		}}

		if i%200 == 0 { // assume every 200 record hr sensor is not sending any data
			record.RemoveFieldByNum(fieldnum.RecordHeartRate)
		}

		fit.Messages = append(fit.Messages, record)
	}

	return fit
}

func BenchmarkEncode(b *testing.B) {
	b.StopTimer()
	fit := createFitForBenchmark(100_000)
	b.StartTimer()

	b.Run("normal header zero", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(io.Discard)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(io.Discard)
		}
	})
	b.Run("normal header 15", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(io.Discard, encoder.WithHeaderOption(encoder.HeaderOptionNormal, 15))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(io.Discard, encoder.WithHeaderOption(encoder.HeaderOptionNormal, 15))
		}
	})
	b.Run("compressed timestamp header", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(io.Discard, encoder.WithHeaderOption(encoder.HeaderOptionCompressedTimestamp, 0))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(io.Discard, encoder.WithHeaderOption(encoder.HeaderOptionCompressedTimestamp, 0))
		}
	})
}

func BenchmarkEncodeWriterAt(b *testing.B) {
	b.StopTimer()
	fit := createFitForBenchmark(100_000)
	b.StartTimer()

	b.Run("normal header zero", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(DiscardAt)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(DiscardAt)
		}
	})
	b.Run("normal header 15", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(DiscardAt, encoder.WithHeaderOption(encoder.HeaderOptionNormal, 15))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(DiscardAt, encoder.WithHeaderOption(encoder.HeaderOptionNormal, 15))
		}
	})
	b.Run("compressed timestamp header", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(DiscardAt, encoder.WithHeaderOption(encoder.HeaderOptionCompressedTimestamp, 0))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(DiscardAt, encoder.WithHeaderOption(encoder.HeaderOptionCompressedTimestamp, 0))
		}
	})
}

func BenchmarkReset(b *testing.B) {
	b.Run("benchmark New()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = encoder.New(nil)
		}
	})
	b.Run("benchmark Reset()", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(nil)
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			enc.Reset(nil)
		}
	})
}
