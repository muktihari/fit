// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

const RecordSize = 100_000

func createBigActivityFile(ctx context.Context) error {
	f, err := os.OpenFile(filepath.Join(testdata, "big_activity.fit"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	now := datetime.ToTime(uint32(1062766519))
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

	n := RecordSize - len(fit.Messages)
	for i := 0; i < n; i++ {
		now = now.Add(time.Second) // only time is moving forward
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
		fit.Messages = append(fit.Messages, record)
	}

	enc := encoder.New(f)
	if err := enc.EncodeWithContext(ctx, fit); err != nil {
		return err
	}

	return nil
}
