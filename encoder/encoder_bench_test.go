package encoder_test

import (
	"io"
	"testing"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
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
	fit := new(proto.FIT)
	fit.Messages = make([]proto.Message, 0, recodSize)
	fit.Messages = append(fit.Messages,
		factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
			fieldnum.FileIdType:         typedef.FileActivity,
			fieldnum.FileIdManufacturer: typedef.ManufacturerBryton,
			fieldnum.FileIdProductName:  "1901",
			fieldnum.FileIdNumber:       uint16(0),
			fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
			fieldnum.FileIdSerialNumber: uint32(5122),
		}),
		factory.CreateMesg(mesgnum.Sport).WithFieldValues(map[byte]any{
			fieldnum.SportSport:    typedef.SportCycling,
			fieldnum.SportSubSport: typedef.SubSportRoad,
		}),
		factory.CreateMesg(mesgnum.Activity).WithFieldValues(map[byte]any{
			fieldnum.ActivityTimestamp:      datetime.ToUint32(now),
			fieldnum.ActivityType:           typedef.ActivityTypeCycling,
			fieldnum.ActivityTotalTimerTime: uint32(30877.0 * 1000),
			fieldnum.ActivityNumSessions:    uint16(1),
			fieldnum.ActivityEvent:          typedef.EventActivity,
		}),
		factory.CreateMesg(mesgnum.Session).WithFieldValues(map[byte]any{
			fieldnum.SessionTimestamp:        datetime.ToUint32(now),
			fieldnum.SessionStartTime:        datetime.ToUint32(now),
			fieldnum.SessionTotalElapsedTime: uint32(30877.0 * 1000),
			fieldnum.SessionTotalDistance:    uint32(32172.05 * 100),
			fieldnum.SessionSport:            typedef.SportCycling,
			fieldnum.SessionSubSport:         typedef.SubSportRoad,
			fieldnum.SessionTotalMovingTime:  uint32(22079.0 * 1000),
			fieldnum.SessionTotalCalories:    uint16(12824),
			fieldnum.SessionAvgSpeed:         uint16(5.98 * 1000),
			fieldnum.SessionMaxSpeed:         uint16(13.05 * 1000),
			fieldnum.SessionMaxAltitude:      uint16((504.0 + 500) * 5),
			fieldnum.SessionTotalAscent:      uint16(909),
			fieldnum.SessionTotalDescent:     uint16(901),
			fieldnum.SessionSwcLat:           int32(0),
			fieldnum.SessionSwcLong:          int32(0),
			fieldnum.SessionNecLat:           int32(0),
			fieldnum.SessionNecLong:          int32(0),
		}),
	)

	for i := 0; i < recodSize-len(fit.Messages); i++ {
		now = now.Add(time.Second) // only time is moving forward
		if i%100 == 0 {            // add event every 100 message
			fit.Messages = append(fit.Messages, factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStop)),
			))
			now = now.Add(10 * time.Second) // gap
			fit.Messages = append(fit.Messages, factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStart)),
			))
			now = now.Add(time.Second) // gap
		}

		record := factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp:    datetime.ToUint32(now),
			fieldnum.RecordPositionLat:  int32(-90481372),
			fieldnum.RecordPositionLong: int32(1323227263),
			fieldnum.RecordSpeed:        uint16(8.33 * 1000),
			fieldnum.RecordDistance:     uint32(405.81 * 100),
			fieldnum.RecordHeartRate:    uint8(110),
			fieldnum.RecordCadence:      uint8(85),
			fieldnum.RecordAltitude:     uint16((166.0 + 500.0) * 5.0),
			fieldnum.RecordTemperature:  int8(32),
		})

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
		enc := encoder.New(io.Discard, encoder.WithNormalHeader(15))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(io.Discard, encoder.WithNormalHeader(15))
		}
	})
	b.Run("compressed timestamp header", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(io.Discard, encoder.WithCompressedTimestampHeader())
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(io.Discard, encoder.WithCompressedTimestampHeader())
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
		enc := encoder.New(DiscardAt, encoder.WithNormalHeader(15))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(DiscardAt, encoder.WithNormalHeader(15))
		}
	})
	b.Run("compressed timestamp header", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(DiscardAt, encoder.WithCompressedTimestampHeader())
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.Reset(DiscardAt, encoder.WithCompressedTimestampHeader())
		}
	})
}

func BenchmarkReset(b *testing.B) {
	mv := encoder.NewMessageValidator()
	b.Run("benchmark New()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = encoder.New(nil, encoder.WithMessageValidator(mv))
		}
	})
	b.Run("benchmark Reset()", func(b *testing.B) {
		b.StopTimer()
		enc := encoder.New(nil)
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			enc.Reset(nil, encoder.WithMessageValidator(mv))
		}
	})
}
