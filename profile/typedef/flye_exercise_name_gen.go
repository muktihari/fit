// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type FlyeExerciseName uint16

const (
	FlyeExerciseNameCableCrossover                    FlyeExerciseName = 0
	FlyeExerciseNameDeclineDumbbellFlye               FlyeExerciseName = 1
	FlyeExerciseNameDumbbellFlye                      FlyeExerciseName = 2
	FlyeExerciseNameInclineDumbbellFlye               FlyeExerciseName = 3
	FlyeExerciseNameKettlebellFlye                    FlyeExerciseName = 4
	FlyeExerciseNameKneelingRearFlye                  FlyeExerciseName = 5
	FlyeExerciseNameSingleArmStandingCableReverseFlye FlyeExerciseName = 6
	FlyeExerciseNameSwissBallDumbbellFlye             FlyeExerciseName = 7
	FlyeExerciseNameArmRotations                      FlyeExerciseName = 8
	FlyeExerciseNameHugATree                          FlyeExerciseName = 9
	FlyeExerciseNameInvalid                           FlyeExerciseName = 0xFFFF // INVALID
)

var flyeexercisenametostrs = map[FlyeExerciseName]string{
	FlyeExerciseNameCableCrossover:                    "cable_crossover",
	FlyeExerciseNameDeclineDumbbellFlye:               "decline_dumbbell_flye",
	FlyeExerciseNameDumbbellFlye:                      "dumbbell_flye",
	FlyeExerciseNameInclineDumbbellFlye:               "incline_dumbbell_flye",
	FlyeExerciseNameKettlebellFlye:                    "kettlebell_flye",
	FlyeExerciseNameKneelingRearFlye:                  "kneeling_rear_flye",
	FlyeExerciseNameSingleArmStandingCableReverseFlye: "single_arm_standing_cable_reverse_flye",
	FlyeExerciseNameSwissBallDumbbellFlye:             "swiss_ball_dumbbell_flye",
	FlyeExerciseNameArmRotations:                      "arm_rotations",
	FlyeExerciseNameHugATree:                          "hug_a_tree",
	FlyeExerciseNameInvalid:                           "invalid",
}

func (f FlyeExerciseName) String() string {
	val, ok := flyeexercisenametostrs[f]
	if !ok {
		return strconv.FormatUint(uint64(f), 10)
	}
	return val
}

var strtoflyeexercisename = func() map[string]FlyeExerciseName {
	m := make(map[string]FlyeExerciseName)
	for t, str := range flyeexercisenametostrs {
		m[str] = FlyeExerciseName(t)
	}
	return m
}()

// FromString parse string into FlyeExerciseName constant it's represent, return FlyeExerciseNameInvalid if not found.
func FlyeExerciseNameFromString(s string) FlyeExerciseName {
	val, ok := strtoflyeexercisename[s]
	if !ok {
		return strtoflyeexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListFlyeExerciseName() []FlyeExerciseName {
	vs := make([]FlyeExerciseName, 0, len(flyeexercisenametostrs))
	for i := range flyeexercisenametostrs {
		vs = append(vs, FlyeExerciseName(i))
	}
	return vs
}
