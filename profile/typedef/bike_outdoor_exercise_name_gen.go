// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BikeOutdoorExerciseName uint16

const (
	BikeOutdoorExerciseNameBike    BikeOutdoorExerciseName = 0
	BikeOutdoorExerciseNameInvalid BikeOutdoorExerciseName = 0xFFFF
)

func (b BikeOutdoorExerciseName) Uint16() uint16 { return uint16(b) }

func (b BikeOutdoorExerciseName) String() string {
	switch b {
	case BikeOutdoorExerciseNameBike:
		return "bike"
	default:
		return "BikeOutdoorExerciseNameInvalid(" + strconv.FormatUint(uint64(b), 10) + ")"
	}
}

// FromString parse string into BikeOutdoorExerciseName constant it's represent, return BikeOutdoorExerciseNameInvalid if not found.
func BikeOutdoorExerciseNameFromString(s string) BikeOutdoorExerciseName {
	switch s {
	case "bike":
		return BikeOutdoorExerciseNameBike
	default:
		return BikeOutdoorExerciseNameInvalid
	}
}

// List returns all constants.
func ListBikeOutdoorExerciseName() []BikeOutdoorExerciseName {
	return []BikeOutdoorExerciseName{
		BikeOutdoorExerciseNameBike,
	}
}
