// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BenchPressExerciseName uint16

const (
	BenchPressExerciseNameAlternatingDumbbellChestPressOnSwissBall BenchPressExerciseName = 0
	BenchPressExerciseNameBarbellBenchPress                        BenchPressExerciseName = 1
	BenchPressExerciseNameBarbellBoardBenchPress                   BenchPressExerciseName = 2
	BenchPressExerciseNameBarbellFloorPress                        BenchPressExerciseName = 3
	BenchPressExerciseNameCloseGripBarbellBenchPress               BenchPressExerciseName = 4
	BenchPressExerciseNameDeclineDumbbellBenchPress                BenchPressExerciseName = 5
	BenchPressExerciseNameDumbbellBenchPress                       BenchPressExerciseName = 6
	BenchPressExerciseNameDumbbellFloorPress                       BenchPressExerciseName = 7
	BenchPressExerciseNameInclineBarbellBenchPress                 BenchPressExerciseName = 8
	BenchPressExerciseNameInclineDumbbellBenchPress                BenchPressExerciseName = 9
	BenchPressExerciseNameInclineSmithMachineBenchPress            BenchPressExerciseName = 10
	BenchPressExerciseNameIsometricBarbellBenchPress               BenchPressExerciseName = 11
	BenchPressExerciseNameKettlebellChestPress                     BenchPressExerciseName = 12
	BenchPressExerciseNameNeutralGripDumbbellBenchPress            BenchPressExerciseName = 13
	BenchPressExerciseNameNeutralGripDumbbellInclineBenchPress     BenchPressExerciseName = 14
	BenchPressExerciseNameOneArmFloorPress                         BenchPressExerciseName = 15
	BenchPressExerciseNameWeightedOneArmFloorPress                 BenchPressExerciseName = 16
	BenchPressExerciseNamePartialLockout                           BenchPressExerciseName = 17
	BenchPressExerciseNameReverseGripBarbellBenchPress             BenchPressExerciseName = 18
	BenchPressExerciseNameReverseGripInclineBenchPress             BenchPressExerciseName = 19
	BenchPressExerciseNameSingleArmCableChestPress                 BenchPressExerciseName = 20
	BenchPressExerciseNameSingleArmDumbbellBenchPress              BenchPressExerciseName = 21
	BenchPressExerciseNameSmithMachineBenchPress                   BenchPressExerciseName = 22
	BenchPressExerciseNameSwissBallDumbbellChestPress              BenchPressExerciseName = 23
	BenchPressExerciseNameTripleStopBarbellBenchPress              BenchPressExerciseName = 24
	BenchPressExerciseNameWideGripBarbellBenchPress                BenchPressExerciseName = 25
	BenchPressExerciseNameAlternatingDumbbellChestPress            BenchPressExerciseName = 26
	BenchPressExerciseNameInvalid                                  BenchPressExerciseName = 0xFFFF // INVALID
)

var benchpressexercisenametostrs = map[BenchPressExerciseName]string{
	BenchPressExerciseNameAlternatingDumbbellChestPressOnSwissBall: "alternating_dumbbell_chest_press_on_swiss_ball",
	BenchPressExerciseNameBarbellBenchPress:                        "barbell_bench_press",
	BenchPressExerciseNameBarbellBoardBenchPress:                   "barbell_board_bench_press",
	BenchPressExerciseNameBarbellFloorPress:                        "barbell_floor_press",
	BenchPressExerciseNameCloseGripBarbellBenchPress:               "close_grip_barbell_bench_press",
	BenchPressExerciseNameDeclineDumbbellBenchPress:                "decline_dumbbell_bench_press",
	BenchPressExerciseNameDumbbellBenchPress:                       "dumbbell_bench_press",
	BenchPressExerciseNameDumbbellFloorPress:                       "dumbbell_floor_press",
	BenchPressExerciseNameInclineBarbellBenchPress:                 "incline_barbell_bench_press",
	BenchPressExerciseNameInclineDumbbellBenchPress:                "incline_dumbbell_bench_press",
	BenchPressExerciseNameInclineSmithMachineBenchPress:            "incline_smith_machine_bench_press",
	BenchPressExerciseNameIsometricBarbellBenchPress:               "isometric_barbell_bench_press",
	BenchPressExerciseNameKettlebellChestPress:                     "kettlebell_chest_press",
	BenchPressExerciseNameNeutralGripDumbbellBenchPress:            "neutral_grip_dumbbell_bench_press",
	BenchPressExerciseNameNeutralGripDumbbellInclineBenchPress:     "neutral_grip_dumbbell_incline_bench_press",
	BenchPressExerciseNameOneArmFloorPress:                         "one_arm_floor_press",
	BenchPressExerciseNameWeightedOneArmFloorPress:                 "weighted_one_arm_floor_press",
	BenchPressExerciseNamePartialLockout:                           "partial_lockout",
	BenchPressExerciseNameReverseGripBarbellBenchPress:             "reverse_grip_barbell_bench_press",
	BenchPressExerciseNameReverseGripInclineBenchPress:             "reverse_grip_incline_bench_press",
	BenchPressExerciseNameSingleArmCableChestPress:                 "single_arm_cable_chest_press",
	BenchPressExerciseNameSingleArmDumbbellBenchPress:              "single_arm_dumbbell_bench_press",
	BenchPressExerciseNameSmithMachineBenchPress:                   "smith_machine_bench_press",
	BenchPressExerciseNameSwissBallDumbbellChestPress:              "swiss_ball_dumbbell_chest_press",
	BenchPressExerciseNameTripleStopBarbellBenchPress:              "triple_stop_barbell_bench_press",
	BenchPressExerciseNameWideGripBarbellBenchPress:                "wide_grip_barbell_bench_press",
	BenchPressExerciseNameAlternatingDumbbellChestPress:            "alternating_dumbbell_chest_press",
	BenchPressExerciseNameInvalid:                                  "invalid",
}

func (b BenchPressExerciseName) String() string {
	val, ok := benchpressexercisenametostrs[b]
	if !ok {
		return strconv.FormatUint(uint64(b), 10)
	}
	return val
}

var strtobenchpressexercisename = func() map[string]BenchPressExerciseName {
	m := make(map[string]BenchPressExerciseName)
	for t, str := range benchpressexercisenametostrs {
		m[str] = BenchPressExerciseName(t)
	}
	return m
}()

// FromString parse string into BenchPressExerciseName constant it's represent, return BenchPressExerciseNameInvalid if not found.
func BenchPressExerciseNameFromString(s string) BenchPressExerciseName {
	val, ok := strtobenchpressexercisename[s]
	if !ok {
		return strtobenchpressexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListBenchPressExerciseName() []BenchPressExerciseName {
	vs := make([]BenchPressExerciseName, 0, len(benchpressexercisenametostrs))
	for i := range benchpressexercisenametostrs {
		vs = append(vs, BenchPressExerciseName(i))
	}
	return vs
}
