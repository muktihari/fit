// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TimeMode byte

const (
	TimeModeHour12            TimeMode = 0
	TimeModeHour24            TimeMode = 1 // Does not use a leading zero and has a colon
	TimeModeMilitary          TimeMode = 2 // Uses a leading zero and does not have a colon
	TimeModeHour12WithSeconds TimeMode = 3
	TimeModeHour24WithSeconds TimeMode = 4
	TimeModeUtc               TimeMode = 5
	TimeModeInvalid           TimeMode = 0xFF
)

func (t TimeMode) Byte() byte { return byte(t) }

func (t TimeMode) String() string {
	switch t {
	case TimeModeHour12:
		return "hour12"
	case TimeModeHour24:
		return "hour24"
	case TimeModeMilitary:
		return "military"
	case TimeModeHour12WithSeconds:
		return "hour_12_with_seconds"
	case TimeModeHour24WithSeconds:
		return "hour_24_with_seconds"
	case TimeModeUtc:
		return "utc"
	default:
		return "TimeModeInvalid(" + strconv.Itoa(int(t)) + ")"
	}
}

// FromString parse string into TimeMode constant it's represent, return TimeModeInvalid if not found.
func TimeModeFromString(s string) TimeMode {
	switch s {
	case "hour12":
		return TimeModeHour12
	case "hour24":
		return TimeModeHour24
	case "military":
		return TimeModeMilitary
	case "hour_12_with_seconds":
		return TimeModeHour12WithSeconds
	case "hour_24_with_seconds":
		return TimeModeHour24WithSeconds
	case "utc":
		return TimeModeUtc
	default:
		return TimeModeInvalid
	}
}

// List returns all constants.
func ListTimeMode() []TimeMode {
	return []TimeMode{
		TimeModeHour12,
		TimeModeHour24,
		TimeModeMilitary,
		TimeModeHour12WithSeconds,
		TimeModeHour24WithSeconds,
		TimeModeUtc,
	}
}
