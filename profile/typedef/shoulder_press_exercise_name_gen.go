// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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
	ShoulderPressExerciseNameInvalid                                  ShoulderPressExerciseName = 0xFFFF
)

func (s ShoulderPressExerciseName) String() string {
	switch s {
	case ShoulderPressExerciseNameAlternatingDumbbellShoulderPress:
		return "alternating_dumbbell_shoulder_press"
	case ShoulderPressExerciseNameArnoldPress:
		return "arnold_press"
	case ShoulderPressExerciseNameBarbellFrontSquatToPushPress:
		return "barbell_front_squat_to_push_press"
	case ShoulderPressExerciseNameBarbellPushPress:
		return "barbell_push_press"
	case ShoulderPressExerciseNameBarbellShoulderPress:
		return "barbell_shoulder_press"
	case ShoulderPressExerciseNameDeadCurlPress:
		return "dead_curl_press"
	case ShoulderPressExerciseNameDumbbellAlternatingShoulderPressAndTwist:
		return "dumbbell_alternating_shoulder_press_and_twist"
	case ShoulderPressExerciseNameDumbbellHammerCurlToLungeToPress:
		return "dumbbell_hammer_curl_to_lunge_to_press"
	case ShoulderPressExerciseNameDumbbellPushPress:
		return "dumbbell_push_press"
	case ShoulderPressExerciseNameFloorInvertedShoulderPress:
		return "floor_inverted_shoulder_press"
	case ShoulderPressExerciseNameWeightedFloorInvertedShoulderPress:
		return "weighted_floor_inverted_shoulder_press"
	case ShoulderPressExerciseNameInvertedShoulderPress:
		return "inverted_shoulder_press"
	case ShoulderPressExerciseNameWeightedInvertedShoulderPress:
		return "weighted_inverted_shoulder_press"
	case ShoulderPressExerciseNameOneArmPushPress:
		return "one_arm_push_press"
	case ShoulderPressExerciseNameOverheadBarbellPress:
		return "overhead_barbell_press"
	case ShoulderPressExerciseNameOverheadDumbbellPress:
		return "overhead_dumbbell_press"
	case ShoulderPressExerciseNameSeatedBarbellShoulderPress:
		return "seated_barbell_shoulder_press"
	case ShoulderPressExerciseNameSeatedDumbbellShoulderPress:
		return "seated_dumbbell_shoulder_press"
	case ShoulderPressExerciseNameSingleArmDumbbellShoulderPress:
		return "single_arm_dumbbell_shoulder_press"
	case ShoulderPressExerciseNameSingleArmStepUpAndPress:
		return "single_arm_step_up_and_press"
	case ShoulderPressExerciseNameSmithMachineOverheadPress:
		return "smith_machine_overhead_press"
	case ShoulderPressExerciseNameSplitStanceHammerCurlToPress:
		return "split_stance_hammer_curl_to_press"
	case ShoulderPressExerciseNameSwissBallDumbbellShoulderPress:
		return "swiss_ball_dumbbell_shoulder_press"
	case ShoulderPressExerciseNameWeightPlateFrontRaise:
		return "weight_plate_front_raise"
	default:
		return "ShoulderPressExerciseNameInvalid(" + strconv.FormatUint(uint64(s), 10) + ")"
	}
}

// FromString parse string into ShoulderPressExerciseName constant it's represent, return ShoulderPressExerciseNameInvalid if not found.
func ShoulderPressExerciseNameFromString(s string) ShoulderPressExerciseName {
	switch s {
	case "alternating_dumbbell_shoulder_press":
		return ShoulderPressExerciseNameAlternatingDumbbellShoulderPress
	case "arnold_press":
		return ShoulderPressExerciseNameArnoldPress
	case "barbell_front_squat_to_push_press":
		return ShoulderPressExerciseNameBarbellFrontSquatToPushPress
	case "barbell_push_press":
		return ShoulderPressExerciseNameBarbellPushPress
	case "barbell_shoulder_press":
		return ShoulderPressExerciseNameBarbellShoulderPress
	case "dead_curl_press":
		return ShoulderPressExerciseNameDeadCurlPress
	case "dumbbell_alternating_shoulder_press_and_twist":
		return ShoulderPressExerciseNameDumbbellAlternatingShoulderPressAndTwist
	case "dumbbell_hammer_curl_to_lunge_to_press":
		return ShoulderPressExerciseNameDumbbellHammerCurlToLungeToPress
	case "dumbbell_push_press":
		return ShoulderPressExerciseNameDumbbellPushPress
	case "floor_inverted_shoulder_press":
		return ShoulderPressExerciseNameFloorInvertedShoulderPress
	case "weighted_floor_inverted_shoulder_press":
		return ShoulderPressExerciseNameWeightedFloorInvertedShoulderPress
	case "inverted_shoulder_press":
		return ShoulderPressExerciseNameInvertedShoulderPress
	case "weighted_inverted_shoulder_press":
		return ShoulderPressExerciseNameWeightedInvertedShoulderPress
	case "one_arm_push_press":
		return ShoulderPressExerciseNameOneArmPushPress
	case "overhead_barbell_press":
		return ShoulderPressExerciseNameOverheadBarbellPress
	case "overhead_dumbbell_press":
		return ShoulderPressExerciseNameOverheadDumbbellPress
	case "seated_barbell_shoulder_press":
		return ShoulderPressExerciseNameSeatedBarbellShoulderPress
	case "seated_dumbbell_shoulder_press":
		return ShoulderPressExerciseNameSeatedDumbbellShoulderPress
	case "single_arm_dumbbell_shoulder_press":
		return ShoulderPressExerciseNameSingleArmDumbbellShoulderPress
	case "single_arm_step_up_and_press":
		return ShoulderPressExerciseNameSingleArmStepUpAndPress
	case "smith_machine_overhead_press":
		return ShoulderPressExerciseNameSmithMachineOverheadPress
	case "split_stance_hammer_curl_to_press":
		return ShoulderPressExerciseNameSplitStanceHammerCurlToPress
	case "swiss_ball_dumbbell_shoulder_press":
		return ShoulderPressExerciseNameSwissBallDumbbellShoulderPress
	case "weight_plate_front_raise":
		return ShoulderPressExerciseNameWeightPlateFrontRaise
	default:
		return ShoulderPressExerciseNameInvalid
	}
}

// List returns all constants.
func ListShoulderPressExerciseName() []ShoulderPressExerciseName {
	return []ShoulderPressExerciseName{
		ShoulderPressExerciseNameAlternatingDumbbellShoulderPress,
		ShoulderPressExerciseNameArnoldPress,
		ShoulderPressExerciseNameBarbellFrontSquatToPushPress,
		ShoulderPressExerciseNameBarbellPushPress,
		ShoulderPressExerciseNameBarbellShoulderPress,
		ShoulderPressExerciseNameDeadCurlPress,
		ShoulderPressExerciseNameDumbbellAlternatingShoulderPressAndTwist,
		ShoulderPressExerciseNameDumbbellHammerCurlToLungeToPress,
		ShoulderPressExerciseNameDumbbellPushPress,
		ShoulderPressExerciseNameFloorInvertedShoulderPress,
		ShoulderPressExerciseNameWeightedFloorInvertedShoulderPress,
		ShoulderPressExerciseNameInvertedShoulderPress,
		ShoulderPressExerciseNameWeightedInvertedShoulderPress,
		ShoulderPressExerciseNameOneArmPushPress,
		ShoulderPressExerciseNameOverheadBarbellPress,
		ShoulderPressExerciseNameOverheadDumbbellPress,
		ShoulderPressExerciseNameSeatedBarbellShoulderPress,
		ShoulderPressExerciseNameSeatedDumbbellShoulderPress,
		ShoulderPressExerciseNameSingleArmDumbbellShoulderPress,
		ShoulderPressExerciseNameSingleArmStepUpAndPress,
		ShoulderPressExerciseNameSmithMachineOverheadPress,
		ShoulderPressExerciseNameSplitStanceHammerCurlToPress,
		ShoulderPressExerciseNameSwissBallDumbbellShoulderPress,
		ShoulderPressExerciseNameWeightPlateFrontRaise,
	}
}
