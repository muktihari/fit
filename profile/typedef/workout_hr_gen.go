// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WorkoutHr uint32

const (
	WorkoutHrBpmOffset WorkoutHr = 100
	WorkoutHrInvalid   WorkoutHr = 0xFFFFFFFF
)

func (w WorkoutHr) Uint32() uint32 { return uint32(w) }

func (w WorkoutHr) String() string {
	switch w {
	case WorkoutHrBpmOffset:
		return "bpm_offset"
	default:
		return "WorkoutHrInvalid(" + strconv.FormatUint(uint64(w), 10) + ")"
	}
}

// FromString parse string into WorkoutHr constant it's represent, return WorkoutHrInvalid if not found.
func WorkoutHrFromString(s string) WorkoutHr {
	switch s {
	case "bpm_offset":
		return WorkoutHrBpmOffset
	default:
		return WorkoutHrInvalid
	}
}

// List returns all constants.
func ListWorkoutHr() []WorkoutHr {
	return []WorkoutHr{
		WorkoutHrBpmOffset,
	}
}
