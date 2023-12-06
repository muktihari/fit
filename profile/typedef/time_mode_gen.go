// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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
	TimeModeInvalid           TimeMode = 0xFF // INVALID
)

var timemodetostrs = map[TimeMode]string{
	TimeModeHour12:            "hour12",
	TimeModeHour24:            "hour24",
	TimeModeMilitary:          "military",
	TimeModeHour12WithSeconds: "hour_12_with_seconds",
	TimeModeHour24WithSeconds: "hour_24_with_seconds",
	TimeModeUtc:               "utc",
	TimeModeInvalid:           "invalid",
}

func (t TimeMode) String() string {
	val, ok := timemodetostrs[t]
	if !ok {
		return strconv.Itoa(int(t))
	}
	return val
}

var strtotimemode = func() map[string]TimeMode {
	m := make(map[string]TimeMode)
	for t, str := range timemodetostrs {
		m[str] = TimeMode(t)
	}
	return m
}()

// FromString parse string into TimeMode constant it's represent, return TimeModeInvalid if not found.
func TimeModeFromString(s string) TimeMode {
	val, ok := strtotimemode[s]
	if !ok {
		return strtotimemode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListTimeMode() []TimeMode {
	vs := make([]TimeMode, 0, len(timemodetostrs))
	for i := range timemodetostrs {
		vs = append(vs, TimeMode(i))
	}
	return vs
}
