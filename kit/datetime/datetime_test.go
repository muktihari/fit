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
