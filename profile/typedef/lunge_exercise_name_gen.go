// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LungeExerciseName uint16

const (
	LungeExerciseNameOverheadLunge                                 LungeExerciseName = 0
	LungeExerciseNameLungeMatrix                                   LungeExerciseName = 1
	LungeExerciseNameWeightedLungeMatrix                           LungeExerciseName = 2
	LungeExerciseNameAlternatingBarbellForwardLunge                LungeExerciseName = 3
	LungeExerciseNameAlternatingDumbbellLungeWithReach             LungeExerciseName = 4
	LungeExerciseNameBackFootElevatedDumbbellSplitSquat            LungeExerciseName = 5
	LungeExerciseNameBarbellBoxLunge                               LungeExerciseName = 6
	LungeExerciseNameBarbellBulgarianSplitSquat                    LungeExerciseName = 7
	LungeExerciseNameBarbellCrossoverLunge                         LungeExerciseName = 8
	LungeExerciseNameBarbellFrontSplitSquat                        LungeExerciseName = 9
	LungeExerciseNameBarbellLunge                                  LungeExerciseName = 10
	LungeExerciseNameBarbellReverseLunge                           LungeExerciseName = 11
	LungeExerciseNameBarbellSideLunge                              LungeExerciseName = 12
	LungeExerciseNameBarbellSplitSquat                             LungeExerciseName = 13
	LungeExerciseNameCoreControlRearLunge                          LungeExerciseName = 14
	LungeExerciseNameDiagonalLunge                                 LungeExerciseName = 15
	LungeExerciseNameDropLunge                                     LungeExerciseName = 16
	LungeExerciseNameDumbbellBoxLunge                              LungeExerciseName = 17
	LungeExerciseNameDumbbellBulgarianSplitSquat                   LungeExerciseName = 18
	LungeExerciseNameDumbbellCrossoverLunge                        LungeExerciseName = 19
	LungeExerciseNameDumbbellDiagonalLunge                         LungeExerciseName = 20
	LungeExerciseNameDumbbellLunge                                 LungeExerciseName = 21
	LungeExerciseNameDumbbellLungeAndRotation                      LungeExerciseName = 22
	LungeExerciseNameDumbbellOverheadBulgarianSplitSquat           LungeExerciseName = 23
	LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress        LungeExerciseName = 24
	LungeExerciseNameDumbbellSideLunge                             LungeExerciseName = 25
	LungeExerciseNameElevatedFrontFootBarbellSplitSquat            LungeExerciseName = 26
	LungeExerciseNameFrontFootElevatedDumbbellSplitSquat           LungeExerciseName = 27
	LungeExerciseNameGunslingerLunge                               LungeExerciseName = 28
	LungeExerciseNameLawnmowerLunge                                LungeExerciseName = 29
	LungeExerciseNameLowLungeWithIsometricAdduction                LungeExerciseName = 30
	LungeExerciseNameLowSideToSideLunge                            LungeExerciseName = 31
	LungeExerciseNameLunge                                         LungeExerciseName = 32
	LungeExerciseNameWeightedLunge                                 LungeExerciseName = 33
	LungeExerciseNameLungeWithArmReach                             LungeExerciseName = 34
	LungeExerciseNameLungeWithDiagonalReach                        LungeExerciseName = 35
	LungeExerciseNameLungeWithSideBend                             LungeExerciseName = 36
	LungeExerciseNameOffsetDumbbellLunge                           LungeExerciseName = 37
	LungeExerciseNameOffsetDumbbellReverseLunge                    LungeExerciseName = 38
	LungeExerciseNameOverheadBulgarianSplitSquat                   LungeExerciseName = 39
	LungeExerciseNameOverheadDumbbellReverseLunge                  LungeExerciseName = 40
	LungeExerciseNameOverheadDumbbellSplitSquat                    LungeExerciseName = 41
	LungeExerciseNameOverheadLungeWithRotation                     LungeExerciseName = 42
	LungeExerciseNameReverseBarbellBoxLunge                        LungeExerciseName = 43
	LungeExerciseNameReverseBoxLunge                               LungeExerciseName = 44
	LungeExerciseNameReverseDumbbellBoxLunge                       LungeExerciseName = 45
	LungeExerciseNameReverseDumbbellCrossoverLunge                 LungeExerciseName = 46
	LungeExerciseNameReverseDumbbellDiagonalLunge                  LungeExerciseName = 47
	LungeExerciseNameReverseLungeWithReachBack                     LungeExerciseName = 48
	LungeExerciseNameWeightedReverseLungeWithReachBack             LungeExerciseName = 49
	LungeExerciseNameReverseLungeWithTwistAndOverheadReach         LungeExerciseName = 50
	LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach LungeExerciseName = 51
	LungeExerciseNameReverseSlidingBoxLunge                        LungeExerciseName = 52
	LungeExerciseNameWeightedReverseSlidingBoxLunge                LungeExerciseName = 53
	LungeExerciseNameReverseSlidingLunge                           LungeExerciseName = 54
	LungeExerciseNameWeightedReverseSlidingLunge                   LungeExerciseName = 55
	LungeExerciseNameRunnersLungeToBalance                         LungeExerciseName = 56
	LungeExerciseNameWeightedRunnersLungeToBalance                 LungeExerciseName = 57
	LungeExerciseNameShiftingSideLunge                             LungeExerciseName = 58
	LungeExerciseNameSideAndCrossoverLunge                         LungeExerciseName = 59
	LungeExerciseNameWeightedSideAndCrossoverLunge                 LungeExerciseName = 60
	LungeExerciseNameSideLunge                                     LungeExerciseName = 61
	LungeExerciseNameWeightedSideLunge                             LungeExerciseName = 62
	LungeExerciseNameSideLungeAndPress                             LungeExerciseName = 63
	LungeExerciseNameSideLungeJumpOff                              LungeExerciseName = 64
	LungeExerciseNameSideLungeSweep                                LungeExerciseName = 65
	LungeExerciseNameWeightedSideLungeSweep                        LungeExerciseName = 66
	LungeExerciseNameSideLungeToCrossoverTap                       LungeExerciseName = 67
	LungeExerciseNameWeightedSideLungeToCrossoverTap               LungeExerciseName = 68
	LungeExerciseNameSideToSideLungeChops                          LungeExerciseName = 69
	LungeExerciseNameWeightedSideToSideLungeChops                  LungeExerciseName = 70
	LungeExerciseNameSiffJumpLunge                                 LungeExerciseName = 71
	LungeExerciseNameWeightedSiffJumpLunge                         LungeExerciseName = 72
	LungeExerciseNameSingleArmReverseLungeAndPress                 LungeExerciseName = 73
	LungeExerciseNameSlidingLateralLunge                           LungeExerciseName = 74
	LungeExerciseNameWeightedSlidingLateralLunge                   LungeExerciseName = 75
	LungeExerciseNameWalkingBarbellLunge                           LungeExerciseName = 76
	LungeExerciseNameWalkingDumbbellLunge                          LungeExerciseName = 77
	LungeExerciseNameWalkingLunge                                  LungeExerciseName = 78
	LungeExerciseNameWeightedWalkingLunge                          LungeExerciseName = 79
	LungeExerciseNameWideGripOverheadBarbellSplitSquat             LungeExerciseName = 80
	LungeExerciseNameInvalid                                       LungeExerciseName = 0xFFFF // INVALID
)

var lungeexercisenametostrs = map[LungeExerciseName]string{
	LungeExerciseNameOverheadLunge:                                 "overhead_lunge",
	LungeExerciseNameLungeMatrix:                                   "lunge_matrix",
	LungeExerciseNameWeightedLungeMatrix:                           "weighted_lunge_matrix",
	LungeExerciseNameAlternatingBarbellForwardLunge:                "alternating_barbell_forward_lunge",
	LungeExerciseNameAlternatingDumbbellLungeWithReach:             "alternating_dumbbell_lunge_with_reach",
	LungeExerciseNameBackFootElevatedDumbbellSplitSquat:            "back_foot_elevated_dumbbell_split_squat",
	LungeExerciseNameBarbellBoxLunge:                               "barbell_box_lunge",
	LungeExerciseNameBarbellBulgarianSplitSquat:                    "barbell_bulgarian_split_squat",
	LungeExerciseNameBarbellCrossoverLunge:                         "barbell_crossover_lunge",
	LungeExerciseNameBarbellFrontSplitSquat:                        "barbell_front_split_squat",
	LungeExerciseNameBarbellLunge:                                  "barbell_lunge",
	LungeExerciseNameBarbellReverseLunge:                           "barbell_reverse_lunge",
	LungeExerciseNameBarbellSideLunge:                              "barbell_side_lunge",
	LungeExerciseNameBarbellSplitSquat:                             "barbell_split_squat",
	LungeExerciseNameCoreControlRearLunge:                          "core_control_rear_lunge",
	LungeExerciseNameDiagonalLunge:                                 "diagonal_lunge",
	LungeExerciseNameDropLunge:                                     "drop_lunge",
	LungeExerciseNameDumbbellBoxLunge:                              "dumbbell_box_lunge",
	LungeExerciseNameDumbbellBulgarianSplitSquat:                   "dumbbell_bulgarian_split_squat",
	LungeExerciseNameDumbbellCrossoverLunge:                        "dumbbell_crossover_lunge",
	LungeExerciseNameDumbbellDiagonalLunge:                         "dumbbell_diagonal_lunge",
	LungeExerciseNameDumbbellLunge:                                 "dumbbell_lunge",
	LungeExerciseNameDumbbellLungeAndRotation:                      "dumbbell_lunge_and_rotation",
	LungeExerciseNameDumbbellOverheadBulgarianSplitSquat:           "dumbbell_overhead_bulgarian_split_squat",
	LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress:        "dumbbell_reverse_lunge_to_high_knee_and_press",
	LungeExerciseNameDumbbellSideLunge:                             "dumbbell_side_lunge",
	LungeExerciseNameElevatedFrontFootBarbellSplitSquat:            "elevated_front_foot_barbell_split_squat",
	LungeExerciseNameFrontFootElevatedDumbbellSplitSquat:           "front_foot_elevated_dumbbell_split_squat",
	LungeExerciseNameGunslingerLunge:                               "gunslinger_lunge",
	LungeExerciseNameLawnmowerLunge:                                "lawnmower_lunge",
	LungeExerciseNameLowLungeWithIsometricAdduction:                "low_lunge_with_isometric_adduction",
	LungeExerciseNameLowSideToSideLunge:                            "low_side_to_side_lunge",
	LungeExerciseNameLunge:                                         "lunge",
	LungeExerciseNameWeightedLunge:                                 "weighted_lunge",
	LungeExerciseNameLungeWithArmReach:                             "lunge_with_arm_reach",
	LungeExerciseNameLungeWithDiagonalReach:                        "lunge_with_diagonal_reach",
	LungeExerciseNameLungeWithSideBend:                             "lunge_with_side_bend",
	LungeExerciseNameOffsetDumbbellLunge:                           "offset_dumbbell_lunge",
	LungeExerciseNameOffsetDumbbellReverseLunge:                    "offset_dumbbell_reverse_lunge",
	LungeExerciseNameOverheadBulgarianSplitSquat:                   "overhead_bulgarian_split_squat",
	LungeExerciseNameOverheadDumbbellReverseLunge:                  "overhead_dumbbell_reverse_lunge",
	LungeExerciseNameOverheadDumbbellSplitSquat:                    "overhead_dumbbell_split_squat",
	LungeExerciseNameOverheadLungeWithRotation:                     "overhead_lunge_with_rotation",
	LungeExerciseNameReverseBarbellBoxLunge:                        "reverse_barbell_box_lunge",
	LungeExerciseNameReverseBoxLunge:                               "reverse_box_lunge",
	LungeExerciseNameReverseDumbbellBoxLunge:                       "reverse_dumbbell_box_lunge",
	LungeExerciseNameReverseDumbbellCrossoverLunge:                 "reverse_dumbbell_crossover_lunge",
	LungeExerciseNameReverseDumbbellDiagonalLunge:                  "reverse_dumbbell_diagonal_lunge",
	LungeExerciseNameReverseLungeWithReachBack:                     "reverse_lunge_with_reach_back",
	LungeExerciseNameWeightedReverseLungeWithReachBack:             "weighted_reverse_lunge_with_reach_back",
	LungeExerciseNameReverseLungeWithTwistAndOverheadReach:         "reverse_lunge_with_twist_and_overhead_reach",
	LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach: "weighted_reverse_lunge_with_twist_and_overhead_reach",
	LungeExerciseNameReverseSlidingBoxLunge:                        "reverse_sliding_box_lunge",
	LungeExerciseNameWeightedReverseSlidingBoxLunge:                "weighted_reverse_sliding_box_lunge",
	LungeExerciseNameReverseSlidingLunge:                           "reverse_sliding_lunge",
	LungeExerciseNameWeightedReverseSlidingLunge:                   "weighted_reverse_sliding_lunge",
	LungeExerciseNameRunnersLungeToBalance:                         "runners_lunge_to_balance",
	LungeExerciseNameWeightedRunnersLungeToBalance:                 "weighted_runners_lunge_to_balance",
	LungeExerciseNameShiftingSideLunge:                             "shifting_side_lunge",
	LungeExerciseNameSideAndCrossoverLunge:                         "side_and_crossover_lunge",
	LungeExerciseNameWeightedSideAndCrossoverLunge:                 "weighted_side_and_crossover_lunge",
	LungeExerciseNameSideLunge:                                     "side_lunge",
	LungeExerciseNameWeightedSideLunge:                             "weighted_side_lunge",
	LungeExerciseNameSideLungeAndPress:                             "side_lunge_and_press",
	LungeExerciseNameSideLungeJumpOff:                              "side_lunge_jump_off",
	LungeExerciseNameSideLungeSweep:                                "side_lunge_sweep",
	LungeExerciseNameWeightedSideLungeSweep:                        "weighted_side_lunge_sweep",
	LungeExerciseNameSideLungeToCrossoverTap:                       "side_lunge_to_crossover_tap",
	LungeExerciseNameWeightedSideLungeToCrossoverTap:               "weighted_side_lunge_to_crossover_tap",
	LungeExerciseNameSideToSideLungeChops:                          "side_to_side_lunge_chops",
	LungeExerciseNameWeightedSideToSideLungeChops:                  "weighted_side_to_side_lunge_chops",
	LungeExerciseNameSiffJumpLunge:                                 "siff_jump_lunge",
	LungeExerciseNameWeightedSiffJumpLunge:                         "weighted_siff_jump_lunge",
	LungeExerciseNameSingleArmReverseLungeAndPress:                 "single_arm_reverse_lunge_and_press",
	LungeExerciseNameSlidingLateralLunge:                           "sliding_lateral_lunge",
	LungeExerciseNameWeightedSlidingLateralLunge:                   "weighted_sliding_lateral_lunge",
	LungeExerciseNameWalkingBarbellLunge:                           "walking_barbell_lunge",
	LungeExerciseNameWalkingDumbbellLunge:                          "walking_dumbbell_lunge",
	LungeExerciseNameWalkingLunge:                                  "walking_lunge",
	LungeExerciseNameWeightedWalkingLunge:                          "weighted_walking_lunge",
	LungeExerciseNameWideGripOverheadBarbellSplitSquat:             "wide_grip_overhead_barbell_split_squat",
	LungeExerciseNameInvalid:                                       "invalid",
}

func (l LungeExerciseName) String() string {
	val, ok := lungeexercisenametostrs[l]
	if !ok {
		return strconv.FormatUint(uint64(l), 10)
	}
	return val
}

var strtolungeexercisename = func() map[string]LungeExerciseName {
	m := make(map[string]LungeExerciseName)
	for t, str := range lungeexercisenametostrs {
		m[str] = LungeExerciseName(t)
	}
	return m
}()

// FromString parse string into LungeExerciseName constant it's represent, return LungeExerciseNameInvalid if not found.
func LungeExerciseNameFromString(s string) LungeExerciseName {
	val, ok := strtolungeexercisename[s]
	if !ok {
		return strtolungeexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLungeExerciseName() []LungeExerciseName {
	vs := make([]LungeExerciseName, 0, len(lungeexercisenametostrs))
	for i := range lungeexercisenametostrs {
		vs = append(vs, LungeExerciseName(i))
	}
	return vs
}
