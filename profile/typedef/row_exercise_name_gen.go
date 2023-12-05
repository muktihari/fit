// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type RowExerciseName uint16

const (
	RowExerciseNameBarbellStraightLegDeadliftToRow            RowExerciseName = 0
	RowExerciseNameCableRowStanding                           RowExerciseName = 1
	RowExerciseNameDumbbellRow                                RowExerciseName = 2
	RowExerciseNameElevatedFeetInvertedRow                    RowExerciseName = 3
	RowExerciseNameWeightedElevatedFeetInvertedRow            RowExerciseName = 4
	RowExerciseNameFacePull                                   RowExerciseName = 5
	RowExerciseNameFacePullWithExternalRotation               RowExerciseName = 6
	RowExerciseNameInvertedRowWithFeetOnSwissBall             RowExerciseName = 7
	RowExerciseNameWeightedInvertedRowWithFeetOnSwissBall     RowExerciseName = 8
	RowExerciseNameKettlebellRow                              RowExerciseName = 9
	RowExerciseNameModifiedInvertedRow                        RowExerciseName = 10
	RowExerciseNameWeightedModifiedInvertedRow                RowExerciseName = 11
	RowExerciseNameNeutralGripAlternatingDumbbellRow          RowExerciseName = 12
	RowExerciseNameOneArmBentOverRow                          RowExerciseName = 13
	RowExerciseNameOneLeggedDumbbellRow                       RowExerciseName = 14
	RowExerciseNameRenegadeRow                                RowExerciseName = 15
	RowExerciseNameReverseGripBarbellRow                      RowExerciseName = 16
	RowExerciseNameRopeHandleCableRow                         RowExerciseName = 17
	RowExerciseNameSeatedCableRow                             RowExerciseName = 18
	RowExerciseNameSeatedDumbbellRow                          RowExerciseName = 19
	RowExerciseNameSingleArmCableRow                          RowExerciseName = 20
	RowExerciseNameSingleArmCableRowAndRotation               RowExerciseName = 21
	RowExerciseNameSingleArmInvertedRow                       RowExerciseName = 22
	RowExerciseNameWeightedSingleArmInvertedRow               RowExerciseName = 23
	RowExerciseNameSingleArmNeutralGripDumbbellRow            RowExerciseName = 24
	RowExerciseNameSingleArmNeutralGripDumbbellRowAndRotation RowExerciseName = 25
	RowExerciseNameSuspendedInvertedRow                       RowExerciseName = 26
	RowExerciseNameWeightedSuspendedInvertedRow               RowExerciseName = 27
	RowExerciseNameTBarRow                                    RowExerciseName = 28
	RowExerciseNameTowelGripInvertedRow                       RowExerciseName = 29
	RowExerciseNameWeightedTowelGripInvertedRow               RowExerciseName = 30
	RowExerciseNameUnderhandGripCableRow                      RowExerciseName = 31
	RowExerciseNameVGripCableRow                              RowExerciseName = 32
	RowExerciseNameWideGripSeatedCableRow                     RowExerciseName = 33
	RowExerciseNameInvalid                                    RowExerciseName = 0xFFFF // INVALID
)

var rowexercisenametostrs = map[RowExerciseName]string{
	RowExerciseNameBarbellStraightLegDeadliftToRow:            "barbell_straight_leg_deadlift_to_row",
	RowExerciseNameCableRowStanding:                           "cable_row_standing",
	RowExerciseNameDumbbellRow:                                "dumbbell_row",
	RowExerciseNameElevatedFeetInvertedRow:                    "elevated_feet_inverted_row",
	RowExerciseNameWeightedElevatedFeetInvertedRow:            "weighted_elevated_feet_inverted_row",
	RowExerciseNameFacePull:                                   "face_pull",
	RowExerciseNameFacePullWithExternalRotation:               "face_pull_with_external_rotation",
	RowExerciseNameInvertedRowWithFeetOnSwissBall:             "inverted_row_with_feet_on_swiss_ball",
	RowExerciseNameWeightedInvertedRowWithFeetOnSwissBall:     "weighted_inverted_row_with_feet_on_swiss_ball",
	RowExerciseNameKettlebellRow:                              "kettlebell_row",
	RowExerciseNameModifiedInvertedRow:                        "modified_inverted_row",
	RowExerciseNameWeightedModifiedInvertedRow:                "weighted_modified_inverted_row",
	RowExerciseNameNeutralGripAlternatingDumbbellRow:          "neutral_grip_alternating_dumbbell_row",
	RowExerciseNameOneArmBentOverRow:                          "one_arm_bent_over_row",
	RowExerciseNameOneLeggedDumbbellRow:                       "one_legged_dumbbell_row",
	RowExerciseNameRenegadeRow:                                "renegade_row",
	RowExerciseNameReverseGripBarbellRow:                      "reverse_grip_barbell_row",
	RowExerciseNameRopeHandleCableRow:                         "rope_handle_cable_row",
	RowExerciseNameSeatedCableRow:                             "seated_cable_row",
	RowExerciseNameSeatedDumbbellRow:                          "seated_dumbbell_row",
	RowExerciseNameSingleArmCableRow:                          "single_arm_cable_row",
	RowExerciseNameSingleArmCableRowAndRotation:               "single_arm_cable_row_and_rotation",
	RowExerciseNameSingleArmInvertedRow:                       "single_arm_inverted_row",
	RowExerciseNameWeightedSingleArmInvertedRow:               "weighted_single_arm_inverted_row",
	RowExerciseNameSingleArmNeutralGripDumbbellRow:            "single_arm_neutral_grip_dumbbell_row",
	RowExerciseNameSingleArmNeutralGripDumbbellRowAndRotation: "single_arm_neutral_grip_dumbbell_row_and_rotation",
	RowExerciseNameSuspendedInvertedRow:                       "suspended_inverted_row",
	RowExerciseNameWeightedSuspendedInvertedRow:               "weighted_suspended_inverted_row",
	RowExerciseNameTBarRow:                                    "t_bar_row",
	RowExerciseNameTowelGripInvertedRow:                       "towel_grip_inverted_row",
	RowExerciseNameWeightedTowelGripInvertedRow:               "weighted_towel_grip_inverted_row",
	RowExerciseNameUnderhandGripCableRow:                      "underhand_grip_cable_row",
	RowExerciseNameVGripCableRow:                              "v_grip_cable_row",
	RowExerciseNameWideGripSeatedCableRow:                     "wide_grip_seated_cable_row",
	RowExerciseNameInvalid:                                    "invalid",
}

func (r RowExerciseName) String() string {
	val, ok := rowexercisenametostrs[r]
	if !ok {
		return strconv.FormatUint(uint64(r), 10)
	}
	return val
}

var strtorowexercisename = func() map[string]RowExerciseName {
	m := make(map[string]RowExerciseName)
	for t, str := range rowexercisenametostrs {
		m[str] = RowExerciseName(t)
	}
	return m
}()

// FromString parse string into RowExerciseName constant it's represent, return RowExerciseNameInvalid if not found.
func RowExerciseNameFromString(s string) RowExerciseName {
	val, ok := strtorowexercisename[s]
	if !ok {
		return strtorowexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListRowExerciseName() []RowExerciseName {
	vs := make([]RowExerciseName, 0, len(rowexercisenametostrs))
	for i := range rowexercisenametostrs {
		vs = append(vs, RowExerciseName(i))
	}
	return vs
}
