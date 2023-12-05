// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type OlympicLiftExerciseName uint16

const (
	OlympicLiftExerciseNameBarbellHangPowerClean      OlympicLiftExerciseName = 0
	OlympicLiftExerciseNameBarbellHangSquatClean      OlympicLiftExerciseName = 1
	OlympicLiftExerciseNameBarbellPowerClean          OlympicLiftExerciseName = 2
	OlympicLiftExerciseNameBarbellPowerSnatch         OlympicLiftExerciseName = 3
	OlympicLiftExerciseNameBarbellSquatClean          OlympicLiftExerciseName = 4
	OlympicLiftExerciseNameCleanAndJerk               OlympicLiftExerciseName = 5
	OlympicLiftExerciseNameBarbellHangPowerSnatch     OlympicLiftExerciseName = 6
	OlympicLiftExerciseNameBarbellHangPull            OlympicLiftExerciseName = 7
	OlympicLiftExerciseNameBarbellHighPull            OlympicLiftExerciseName = 8
	OlympicLiftExerciseNameBarbellSnatch              OlympicLiftExerciseName = 9
	OlympicLiftExerciseNameBarbellSplitJerk           OlympicLiftExerciseName = 10
	OlympicLiftExerciseNameClean                      OlympicLiftExerciseName = 11
	OlympicLiftExerciseNameDumbbellClean              OlympicLiftExerciseName = 12
	OlympicLiftExerciseNameDumbbellHangPull           OlympicLiftExerciseName = 13
	OlympicLiftExerciseNameOneHandDumbbellSplitSnatch OlympicLiftExerciseName = 14
	OlympicLiftExerciseNamePushJerk                   OlympicLiftExerciseName = 15
	OlympicLiftExerciseNameSingleArmDumbbellSnatch    OlympicLiftExerciseName = 16
	OlympicLiftExerciseNameSingleArmHangSnatch        OlympicLiftExerciseName = 17
	OlympicLiftExerciseNameSingleArmKettlebellSnatch  OlympicLiftExerciseName = 18
	OlympicLiftExerciseNameSplitJerk                  OlympicLiftExerciseName = 19
	OlympicLiftExerciseNameSquatCleanAndJerk          OlympicLiftExerciseName = 20
	OlympicLiftExerciseNameInvalid                    OlympicLiftExerciseName = 0xFFFF // INVALID
)

var olympicliftexercisenametostrs = map[OlympicLiftExerciseName]string{
	OlympicLiftExerciseNameBarbellHangPowerClean:      "barbell_hang_power_clean",
	OlympicLiftExerciseNameBarbellHangSquatClean:      "barbell_hang_squat_clean",
	OlympicLiftExerciseNameBarbellPowerClean:          "barbell_power_clean",
	OlympicLiftExerciseNameBarbellPowerSnatch:         "barbell_power_snatch",
	OlympicLiftExerciseNameBarbellSquatClean:          "barbell_squat_clean",
	OlympicLiftExerciseNameCleanAndJerk:               "clean_and_jerk",
	OlympicLiftExerciseNameBarbellHangPowerSnatch:     "barbell_hang_power_snatch",
	OlympicLiftExerciseNameBarbellHangPull:            "barbell_hang_pull",
	OlympicLiftExerciseNameBarbellHighPull:            "barbell_high_pull",
	OlympicLiftExerciseNameBarbellSnatch:              "barbell_snatch",
	OlympicLiftExerciseNameBarbellSplitJerk:           "barbell_split_jerk",
	OlympicLiftExerciseNameClean:                      "clean",
	OlympicLiftExerciseNameDumbbellClean:              "dumbbell_clean",
	OlympicLiftExerciseNameDumbbellHangPull:           "dumbbell_hang_pull",
	OlympicLiftExerciseNameOneHandDumbbellSplitSnatch: "one_hand_dumbbell_split_snatch",
	OlympicLiftExerciseNamePushJerk:                   "push_jerk",
	OlympicLiftExerciseNameSingleArmDumbbellSnatch:    "single_arm_dumbbell_snatch",
	OlympicLiftExerciseNameSingleArmHangSnatch:        "single_arm_hang_snatch",
	OlympicLiftExerciseNameSingleArmKettlebellSnatch:  "single_arm_kettlebell_snatch",
	OlympicLiftExerciseNameSplitJerk:                  "split_jerk",
	OlympicLiftExerciseNameSquatCleanAndJerk:          "squat_clean_and_jerk",
	OlympicLiftExerciseNameInvalid:                    "invalid",
}

func (o OlympicLiftExerciseName) String() string {
	val, ok := olympicliftexercisenametostrs[o]
	if !ok {
		return strconv.FormatUint(uint64(o), 10)
	}
	return val
}

var strtoolympicliftexercisename = func() map[string]OlympicLiftExerciseName {
	m := make(map[string]OlympicLiftExerciseName)
	for t, str := range olympicliftexercisenametostrs {
		m[str] = OlympicLiftExerciseName(t)
	}
	return m
}()

// FromString parse string into OlympicLiftExerciseName constant it's represent, return OlympicLiftExerciseNameInvalid if not found.
func OlympicLiftExerciseNameFromString(s string) OlympicLiftExerciseName {
	val, ok := strtoolympicliftexercisename[s]
	if !ok {
		return strtoolympicliftexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListOlympicLiftExerciseName() []OlympicLiftExerciseName {
	vs := make([]OlympicLiftExerciseName, 0, len(olympicliftexercisenametostrs))
	for i := range olympicliftexercisenametostrs {
		vs = append(vs, OlympicLiftExerciseName(i))
	}
	return vs
}
