// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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
	DateModeInvalid  DateMode = 0xFF
)

func (d DateMode) Byte() byte { return byte(d) }

func (d DateMode) String() string {
	switch d {
	case DateModeDayMonth:
		return "day_month"
	case DateModeMonthDay:
		return "month_day"
	default:
		return "DateModeInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DateMode constant it's represent, return DateModeInvalid if not found.
func DateModeFromString(s string) DateMode {
	switch s {
	case "day_month":
		return DateModeDayMonth
	case "month_day":
		return DateModeMonthDay
	default:
		return DateModeInvalid
	}
}

// List returns all constants.
func ListDateMode() []DateMode {
	return []DateMode{
		DateModeDayMonth,
		DateModeMonthDay,
	}
}
