// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datetime_test

import (
	"testing"
	"time"

	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/typedef"
)

func TestToTime(t *testing.T) {
	tt := []struct {
		name string
		u32  any
		time time.Time
	}{
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

func Test(t *testing.T) {
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

func TestTzOffsetHours(t *testing.T) {
	tt := []struct {
		name          string
		localDateTime typedef.LocalDateTime
		t             typedef.DateTime
		tzOffsetHours int
	}{
		{
			name:          "",
			localDateTime: 1029647779, // Wed, 17 Aug 2022 05:16:19 GMT
			t:             1029622579,
			tzOffsetHours: 7,
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
