// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ActivityLevel byte

const (
	ActivityLevelLow     ActivityLevel = 0
	ActivityLevelMedium  ActivityLevel = 1
	ActivityLevelHigh    ActivityLevel = 2
	ActivityLevelInvalid ActivityLevel = 0xFF
)

func (a ActivityLevel) Byte() byte { return byte(a) }

func (a ActivityLevel) String() string {
	switch a {
	case ActivityLevelLow:
		return "low"
	case ActivityLevelMedium:
		return "medium"
	case ActivityLevelHigh:
		return "high"
	default:
		return "ActivityLevelInvalid(" + strconv.Itoa(int(a)) + ")"
	}
}

// FromString parse string into ActivityLevel constant it's represent, return ActivityLevelInvalid if not found.
func ActivityLevelFromString(s string) ActivityLevel {
	switch s {
	case "low":
		return ActivityLevelLow
	case "medium":
		return ActivityLevelMedium
	case "high":
		return ActivityLevelHigh
	default:
		return ActivityLevelInvalid
	}
}

// List returns all constants.
func ListActivityLevel() []ActivityLevel {
	return []ActivityLevel{
		ActivityLevelLow,
		ActivityLevelMedium,
		ActivityLevelHigh,
	}
}
