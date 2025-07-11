// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LadderExerciseName uint16

const (
	LadderExerciseNameAgility LadderExerciseName = 0
	LadderExerciseNameSpeed   LadderExerciseName = 1
	LadderExerciseNameInvalid LadderExerciseName = 0xFFFF
)

func (l LadderExerciseName) Uint16() uint16 { return uint16(l) }

func (l LadderExerciseName) String() string {
	switch l {
	case LadderExerciseNameAgility:
		return "agility"
	case LadderExerciseNameSpeed:
		return "speed"
	default:
		return "LadderExerciseNameInvalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LadderExerciseName constant it's represent, return LadderExerciseNameInvalid if not found.
func LadderExerciseNameFromString(s string) LadderExerciseName {
	switch s {
	case "agility":
		return LadderExerciseNameAgility
	case "speed":
		return LadderExerciseNameSpeed
	default:
		return LadderExerciseNameInvalid
	}
}

// List returns all constants.
func ListLadderExerciseName() []LadderExerciseName {
	return []LadderExerciseName{
		LadderExerciseNameAgility,
		LadderExerciseNameSpeed,
	}
}
