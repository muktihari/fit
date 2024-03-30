// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datetime

import (
	"strconv"
	"strings"
	"time"

	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
)

var epoch = time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)

// Epoch return fit epoch (31 Dec 1989 00:00:000 UTC) as time.Time
func Epoch() time.Time { return epoch }

// ToTime return new time based on given v.
func ToTime(value any) time.Time {
	if value == nil {
		return time.Time{}
	}

	var val uint32 = basetype.Uint32Invalid

	switch v := value.(type) {
	case uint32:
		val = v
	case typedef.DateTime:
		val = uint32(v)
	case typedef.LocalDateTime:
		val = uint32(v)
	}

	if val == basetype.Uint32Invalid {
		return time.Time{}
	}

	return epoch.Add(time.Duration(val) * time.Second)
}

// ToLocalTime returns time in local time zone by specifying the time zone offset hours (+7 for GMT+7).
func ToLocalTime(value any, tzOffsetHours int) time.Time {
	t, ok := value.(time.Time)
	if !ok {
		t = ToTime(value)
	}

	if t.IsZero() {
		return t
	}

	zone := new(strings.Builder)
	zone.WriteString("UTC")
	if tzOffsetHours > 0 {
		zone.WriteRune('+')
	}

	zone.WriteString(strconv.Itoa(tzOffsetHours)) // e.g. zone name -> UTC+7, UTC-7, etc...
	loc := time.FixedZone(zone.String(), tzOffsetHours*60*60)

	return t.In(loc)
}

// TzOffsetHours calculates time zone offset.
//
// formula ilustration: (activity.LocalTimestamp - activity.Timestamp) / 3600
func TzOffsetHours(localDateTime, dateTime time.Time) int {
	return int(localDateTime.Sub(dateTime).Seconds()) / 3600
}

// TzOffsetHoursFromUint32 is similar to TzOffsetHours but it took uint32 as parameters.
func TzOffsetHoursFromUint32(localDateTime, dateTime uint32) int {
	return int(localDateTime-dateTime) / 3600
}

// Convert t into uint32 fit representative time value.
func ToUint32(t time.Time) uint32 {
	if t.IsZero() {
		return basetype.Uint32Invalid
	}
	return uint32(t.Sub(epoch).Seconds())
}
