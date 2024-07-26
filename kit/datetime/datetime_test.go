// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datetime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
)

func TestToTime(t *testing.T) {
	tt := []struct {
		name string
		u32  uint32
		time time.Time
	}{
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  uint32(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  uint32(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  uint32(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "Thu, 30 Dec 2021 21:52:08 GMT",
			u32:  uint32(1009835528),
			time: time.Date(2021, time.December, 30, 21, 52, 8, 00, time.UTC),
		},
		{
			name: "nil",
			u32:  basetype.Uint32Invalid,
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
		value  time.Time
		result time.Time
	}{
		{
			name:   "unsupported tipe",
			value:  time.Time{},
			result: time.Time{},
		},
		{
			name:   "value is already in time.Time",
			value:  time.Date(2022, 8, 17, 05, 16, 19, 0, time.UTC),
			result: time.Date(2022, 8, 17, 12, 16, 19, 0, locJakarta),
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
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
		{
			name: "before FIT's epoch",
			time: time.Date(1945, time.November, 10, 0, 0, 0, 0, time.UTC),
			u32:  basetype.Uint32Invalid,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := datetime.ToUint32(tc.time)
			if res != tc.u32 {
				t.Fatalf("expected time in uint32: %d, got: %d", tc.u32, res)
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

func BenchmarkToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = datetime.ToTime(uint32(1009835528))
	}
}

func BenchmarkToLocalTime(b *testing.B) {
	t := datetime.ToTime(uint32(1009835528))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		datetime.ToLocalTime(t, 8)
	}
}
