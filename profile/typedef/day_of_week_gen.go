// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DayOfWeek byte

const (
	DayOfWeekSunday    DayOfWeek = 0
	DayOfWeekMonday    DayOfWeek = 1
	DayOfWeekTuesday   DayOfWeek = 2
	DayOfWeekWednesday DayOfWeek = 3
	DayOfWeekThursday  DayOfWeek = 4
	DayOfWeekFriday    DayOfWeek = 5
	DayOfWeekSaturday  DayOfWeek = 6
	DayOfWeekInvalid   DayOfWeek = 0xFF // INVALID
)

var dayofweektostrs = map[DayOfWeek]string{
	DayOfWeekSunday:    "sunday",
	DayOfWeekMonday:    "monday",
	DayOfWeekTuesday:   "tuesday",
	DayOfWeekWednesday: "wednesday",
	DayOfWeekThursday:  "thursday",
	DayOfWeekFriday:    "friday",
	DayOfWeekSaturday:  "saturday",
	DayOfWeekInvalid:   "invalid",
}

func (d DayOfWeek) String() string {
	val, ok := dayofweektostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodayofweek = func() map[string]DayOfWeek {
	m := make(map[string]DayOfWeek)
	for t, str := range dayofweektostrs {
		m[str] = DayOfWeek(t)
	}
	return m
}()

// FromString parse string into DayOfWeek constant it's represent, return DayOfWeekInvalid if not found.
func DayOfWeekFromString(s string) DayOfWeek {
	val, ok := strtodayofweek[s]
	if !ok {
		return strtodayofweek["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDayOfWeek() []DayOfWeek {
	vs := make([]DayOfWeek, 0, len(dayofweektostrs))
	for i := range dayofweektostrs {
		vs = append(vs, DayOfWeek(i))
	}
	return vs
}
