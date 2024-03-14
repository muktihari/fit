// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type HipSwingExerciseName uint16

const (
	HipSwingExerciseNameSingleArmKettlebellSwing HipSwingExerciseName = 0
	HipSwingExerciseNameSingleArmDumbbellSwing   HipSwingExerciseName = 1
	HipSwingExerciseNameStepOutSwing             HipSwingExerciseName = 2
	HipSwingExerciseNameInvalid                  HipSwingExerciseName = 0xFFFF
)

func (h HipSwingExerciseName) String() string {
	switch h {
	case HipSwingExerciseNameSingleArmKettlebellSwing:
		return "single_arm_kettlebell_swing"
	case HipSwingExerciseNameSingleArmDumbbellSwing:
		return "single_arm_dumbbell_swing"
	case HipSwingExerciseNameStepOutSwing:
		return "step_out_swing"
	default:
		return "HipSwingExerciseNameInvalid(" + strconv.FormatUint(uint64(h), 10) + ")"
	}
}

// FromString parse string into HipSwingExerciseName constant it's represent, return HipSwingExerciseNameInvalid if not found.
func HipSwingExerciseNameFromString(s string) HipSwingExerciseName {
	switch s {
	case "single_arm_kettlebell_swing":
		return HipSwingExerciseNameSingleArmKettlebellSwing
	case "single_arm_dumbbell_swing":
		return HipSwingExerciseNameSingleArmDumbbellSwing
	case "step_out_swing":
		return HipSwingExerciseNameStepOutSwing
	default:
		return HipSwingExerciseNameInvalid
	}
}

// List returns all constants.
func ListHipSwingExerciseName() []HipSwingExerciseName {
	return []HipSwingExerciseName{
		HipSwingExerciseNameSingleArmKettlebellSwing,
		HipSwingExerciseNameSingleArmDumbbellSwing,
		HipSwingExerciseNameStepOutSwing,
	}
}
