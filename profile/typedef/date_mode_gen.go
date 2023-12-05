// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DateMode byte

const (
	DateModeDayMonth DateMode = 0
	DateModeMonthDay DateMode = 1
	DateModeInvalid  DateMode = 0xFF // INVALID
)

var datemodetostrs = map[DateMode]string{
	DateModeDayMonth: "day_month",
	DateModeMonthDay: "month_day",
	DateModeInvalid:  "invalid",
}

func (d DateMode) String() string {
	val, ok := datemodetostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodatemode = func() map[string]DateMode {
	m := make(map[string]DateMode)
	for t, str := range datemodetostrs {
		m[str] = DateMode(t)
	}
	return m
}()

// FromString parse string into DateMode constant it's represent, return DateModeInvalid if not found.
func DateModeFromString(s string) DateMode {
	val, ok := strtodatemode[s]
	if !ok {
		return strtodatemode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDateMode() []DateMode {
	vs := make([]DateMode, 0, len(datemodetostrs))
	for i := range datemodetostrs {
		vs = append(vs, DateMode(i))
	}
	return vs
}
