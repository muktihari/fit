// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/bufferedwriter"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

const (
	RecordSize = 200_000
)

func createBigActivityFile(ctx context.Context) error {
	f, err := os.OpenFile(filepath.Join(testdata, "big_activity.fit"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	now := time.Now()
	fit := new(proto.Fit)
	fit.Messages = make([]proto.Message, 0, RecordSize)
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

	for i := 0; i < RecordSize-len(fit.Messages); i++ {
		now = now.Add(time.Second) // only time is moving forward
		fit.Messages = append(fit.Messages, factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp:    datetime.ToUint32(now),
			fieldnum.RecordPositionLat:  int32(-90481372),
			fieldnum.RecordPositionLong: int32(1323227263),
			fieldnum.RecordSpeed:        uint16(8.33 * 1000),
			fieldnum.RecordDistance:     uint32(405.81 * 100),
			fieldnum.RecordHeartRate:    uint8(110),
			fieldnum.RecordCadence:      uint8(85),
			fieldnum.RecordAltitude:     uint16((166.0 + 500.0) * 5.0),
			fieldnum.RecordTemperature:  int8(32),
		}))
	}

	bw := bufferedwriter.NewSize(f, 1000<<10)
	defer bw.Flush()

	enc := encoder.New(bw)
	if err := enc.Encode(context.Background(), fit); err != nil {
		return err
	}

	return nil
}
