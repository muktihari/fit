// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ShoulderPressExerciseName uint16

const (
	ShoulderPressExerciseNameAlternatingDumbbellShoulderPress         ShoulderPressExerciseName = 0
	ShoulderPressExerciseNameArnoldPress                              ShoulderPressExerciseName = 1
	ShoulderPressExerciseNameBarbellFrontSquatToPushPress             ShoulderPressExerciseName = 2
	ShoulderPressExerciseNameBarbellPushPress                         ShoulderPressExerciseName = 3
	ShoulderPressExerciseNameBarbellShoulderPress                     ShoulderPressExerciseName = 4
	ShoulderPressExerciseNameDeadCurlPress                            ShoulderPressExerciseName = 5
	ShoulderPressExerciseNameDumbbellAlternatingShoulderPressAndTwist ShoulderPressExerciseName = 6
	ShoulderPressExerciseNameDumbbellHammerCurlToLungeToPress         ShoulderPressExerciseName = 7
	ShoulderPressExerciseNameDumbbellPushPress                        ShoulderPressExerciseName = 8
	ShoulderPressExerciseNameFloorInvertedShoulderPress               ShoulderPressExerciseName = 9
	ShoulderPressExerciseNameWeightedFloorInvertedShoulderPress       ShoulderPressExerciseName = 10
	ShoulderPressExerciseNameInvertedShoulderPress                    ShoulderPressExerciseName = 11
	ShoulderPressExerciseNameWeightedInvertedShoulderPress            ShoulderPressExerciseName = 12
	ShoulderPressExerciseNameOneArmPushPress                          ShoulderPressExerciseName = 13
	ShoulderPressExerciseNameOverheadBarbellPress                     ShoulderPressExerciseName = 14
	ShoulderPressExerciseNameOverheadDumbbellPress                    ShoulderPressExerciseName = 15
	ShoulderPressExerciseNameSeatedBarbellShoulderPress               ShoulderPressExerciseName = 16
	ShoulderPressExerciseNameSeatedDumbbellShoulderPress              ShoulderPressExerciseName = 17
	ShoulderPressExerciseNameSingleArmDumbbellShoulderPress           ShoulderPressExerciseName = 18
	ShoulderPressExerciseNameSingleArmStepUpAndPress                  ShoulderPressExerciseName = 19
	ShoulderPressExerciseNameSmithMachineOverheadPress                ShoulderPressExerciseName = 20
	ShoulderPressExerciseNameSplitStanceHammerCurlToPress             ShoulderPressExerciseName = 21
	ShoulderPressExerciseNameSwissBallDumbbellShoulderPress           ShoulderPressExerciseName = 22
	ShoulderPressExerciseNameWeightPlateFrontRaise                    ShoulderPressExerciseName = 23
	ShoulderPressExerciseNameInvalid                                  ShoulderPressExerciseName = 0xFFFF // INVALID
)

var shoulderpressexercisenametostrs = map[ShoulderPressExerciseName]string{
	ShoulderPressExerciseNameAlternatingDumbbellShoulderPress:         "alternating_dumbbell_shoulder_press",
	ShoulderPressExerciseNameArnoldPress:                              "arnold_press",
	ShoulderPressExerciseNameBarbellFrontSquatToPushPress:             "barbell_front_squat_to_push_press",
	ShoulderPressExerciseNameBarbellPushPress:                         "barbell_push_press",
	ShoulderPressExerciseNameBarbellShoulderPress:                     "barbell_shoulder_press",
	ShoulderPressExerciseNameDeadCurlPress:                            "dead_curl_press",
	ShoulderPressExerciseNameDumbbellAlternatingShoulderPressAndTwist: "dumbbell_alternating_shoulder_press_and_twist",
	ShoulderPressExerciseNameDumbbellHammerCurlToLungeToPress:         "dumbbell_hammer_curl_to_lunge_to_press",
	ShoulderPressExerciseNameDumbbellPushPress:                        "dumbbell_push_press",
	ShoulderPressExerciseNameFloorInvertedShoulderPress:               "floor_inverted_shoulder_press",
	ShoulderPressExerciseNameWeightedFloorInvertedShoulderPress:       "weighted_floor_inverted_shoulder_press",
	ShoulderPressExerciseNameInvertedShoulderPress:                    "inverted_shoulder_press",
	ShoulderPressExerciseNameWeightedInvertedShoulderPress:            "weighted_inverted_shoulder_press",
	ShoulderPressExerciseNameOneArmPushPress:                          "one_arm_push_press",
	ShoulderPressExerciseNameOverheadBarbellPress:                     "overhead_barbell_press",
	ShoulderPressExerciseNameOverheadDumbbellPress:                    "overhead_dumbbell_press",
	ShoulderPressExerciseNameSeatedBarbellShoulderPress:               "seated_barbell_shoulder_press",
	ShoulderPressExerciseNameSeatedDumbbellShoulderPress:              "seated_dumbbell_shoulder_press",
	ShoulderPressExerciseNameSingleArmDumbbellShoulderPress:           "single_arm_dumbbell_shoulder_press",
	ShoulderPressExerciseNameSingleArmStepUpAndPress:                  "single_arm_step_up_and_press",
	ShoulderPressExerciseNameSmithMachineOverheadPress:                "smith_machine_overhead_press",
	ShoulderPressExerciseNameSplitStanceHammerCurlToPress:             "split_stance_hammer_curl_to_press",
	ShoulderPressExerciseNameSwissBallDumbbellShoulderPress:           "swiss_ball_dumbbell_shoulder_press",
	ShoulderPressExerciseNameWeightPlateFrontRaise:                    "weight_plate_front_raise",
	ShoulderPressExerciseNameInvalid:                                  "invalid",
}

func (s ShoulderPressExerciseName) String() string {
	val, ok := shoulderpressexercisenametostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtoshoulderpressexercisename = func() map[string]ShoulderPressExerciseName {
	m := make(map[string]ShoulderPressExerciseName)
	for t, str := range shoulderpressexercisenametostrs {
		m[str] = ShoulderPressExerciseName(t)
	}
	return m
}()

// FromString parse string into ShoulderPressExerciseName constant it's represent, return ShoulderPressExerciseNameInvalid if not found.
func ShoulderPressExerciseNameFromString(s string) ShoulderPressExerciseName {
	val, ok := strtoshoulderpressexercisename[s]
	if !ok {
		return strtoshoulderpressexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListShoulderPressExerciseName() []ShoulderPressExerciseName {
	vs := make([]ShoulderPressExerciseName, 0, len(shoulderpressexercisenametostrs))
	for i := range shoulderpressexercisenametostrs {
		vs = append(vs, ShoulderPressExerciseName(i))
	}
	return vs
}
