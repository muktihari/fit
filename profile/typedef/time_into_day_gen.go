// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TimeIntoDay uint32

const (
	TimeIntoDayInvalid TimeIntoDay = 0xFFFFFFFF // INVALID
)

var timeintodaytostrs = map[TimeIntoDay]string{
	TimeIntoDayInvalid: "invalid",
}

func (t TimeIntoDay) String() string {
	val, ok := timeintodaytostrs[t]
	if !ok {
		return strconv.FormatUint(uint64(t), 10)
	}
	return val
}

var strtotimeintoday = func() map[string]TimeIntoDay {
	m := make(map[string]TimeIntoDay)
	for t, str := range timeintodaytostrs {
		m[str] = TimeIntoDay(t)
	}
	return m
}()

// FromString parse string into TimeIntoDay constant it's represent, return TimeIntoDayInvalid if not found.
func TimeIntoDayFromString(s string) TimeIntoDay {
	val, ok := strtotimeintoday[s]
	if !ok {
		return strtotimeintoday["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListTimeIntoDay() []TimeIntoDay {
	vs := make([]TimeIntoDay, 0, len(timeintodaytostrs))
	for i := range timeintodaytostrs {
		vs = append(vs, TimeIntoDay(i))
	}
	return vs
}