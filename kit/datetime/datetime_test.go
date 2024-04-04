// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datetime_test

import (
	"testing"
	"time"

	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

func TestToTime(t *testing.T) {
	tt := []struct {
		name string
		u32  any
		time time.Time
	}{
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  proto.Uint32(uint32(1009835528)),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  uint32(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  typedef.DateTime(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  typedef.LocalDateTime(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "nil",
			u32:  nil,
			time: time.Time{},
		},
		{
			name: "struct{}{}",
			u32:  struct{}{},
			time: time.Time{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := datetime.ToTime(tc.u32)
			if !res.Equal(tc.time) {
				t.Fatalf("expected time: %s, got: %s", tc.time, res)
			}
		})
	}
}

func TestToLocalTime(t *testing.T) {
	locJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatal(err)
	}

	tt := []struct {
		name   string
		value  any
		result time.Time
	}{
		{
			name:   "utc to utc+7",
			value:  uint32(1029622579), // Tue, 16 Aug 2022 22:16:19 GMT
			result: time.Date(2022, 8, 17, 05, 16, 19, 0, locJakarta),
		},
		{
			name:   "unsupported tipe",
			value:  1029622579, // int int
			result: time.Time{},
		},
		{
			name:   "value is already in time.Time",
			value:  time.Date(2022, 8, 17, 05, 16, 19, 0, time.UTC),
			result: time.Date(2022, 8, 17, 12, 16, 19, 0, locJakarta),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_time := datetime.ToLocalTime(tc.value, 7)
			if _time.Format(time.RFC3339) != tc.result.Format(time.RFC3339) { // can't use time equal since the loc object is different
				t.Fatalf("expected: %s, got: %s", tc.result.Format(time.RFC3339), _time.Format(time.RFC3339))
			}
		})
	}
}

// 2022-08-17T05:16:19+07:00, got: 2022-08-17T12:16:19+07:00
func TestTzOffsetHours(t *testing.T) {
	tt := []struct {
		name          string
		localDateTime time.Time
		t             time.Time
		tzOffsetHours int
	}{
		{
			name:          "",
			localDateTime: datetime.ToTime(uint32(1029647779)), // Wed, 17 Aug 2022 05:16:19 GMT
			t:             datetime.ToTime(uint32(1029622579)), // Tue, 16 Aug 2022 22:16:19 GMT
			tzOffsetHours: 7,                                   // actual gap 7 hours
		},
		{
			name:          "",
			localDateTime: datetime.ToTime(uint32(1029648779)), // Wed, 17 Aug 2022 05:32:59 GMT
			t:             datetime.ToTime(uint32(1029622579)), // Tue, 16 Aug 2022 22:16:19 GMT
			tzOffsetHours: 7,                                   // actual gap 7.278 hours
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tz := datetime.TzOffsetHours(tc.localDateTime, tc.t)
			if tz != tc.tzOffsetHours {
				t.Fatalf("expected: %d, got: %d", tc.tzOffsetHours, tz)
			}
		})
	}
}

func TestTzOffsetHoursFromUint32(t *testing.T) {
	tt := []struct {
		name          string
		localDateTime uint32
		t             uint32
		tzOffsetHours int
	}{
		{
			name:          "",
			localDateTime: 1029647779, // Wed, 17 Aug 2022 05:16:19 GMT
			t:             1029622579, // Tue, 16 Aug 2022 22:16:19 GMT
			tzOffsetHours: 7,          // actual gap 7 hours
		},
		{
			name:          "",
			localDateTime: 1029648779, // Wed, 17 Aug 2022 05:32:59 GMT
			t:             1029622579, // Tue, 16 Aug 2022 22:16:19 GMT
			tzOffsetHours: 7,          // actual gap 7.278 hours
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tz := datetime.TzOffsetHoursFromUint32(tc.localDateTime, tc.t)
			if tz != tc.tzOffsetHours {
				t.Fatalf("expected: %d, got: %d", tc.tzOffsetHours, tz)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	tt := []struct {
		name string
		time time.Time
		u32  uint32
	}{
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
			u32:  1009835528,
		},
		{
			name: time.Time{}.Format(time.RFC3339),
			time: time.Time{},
			u32:  basetype.Uint32Invalid,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := datetime.ToUint32(tc.time)
			if res != tc.u32 {
				t.Fatalf("expected time: %d, got: %d", tc.u32, res)
			}
		})
	}
}

func TestEpoch(t *testing.T) {
	ep := datetime.Epoch()
	expected := time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)
	if !ep.Equal(expected) {
		t.Fatalf("expected: %s, got: %s", expected, ep)
	}
}
