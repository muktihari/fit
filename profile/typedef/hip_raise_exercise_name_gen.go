// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type HipRaiseExerciseName uint16

const (
	HipRaiseExerciseNameBarbellHipThrustOnFloor                         HipRaiseExerciseName = 0
	HipRaiseExerciseNameBarbellHipThrustWithBench                       HipRaiseExerciseName = 1
	HipRaiseExerciseNameBentKneeSwissBallReverseHipRaise                HipRaiseExerciseName = 2
	HipRaiseExerciseNameWeightedBentKneeSwissBallReverseHipRaise        HipRaiseExerciseName = 3
	HipRaiseExerciseNameBridgeWithLegExtension                          HipRaiseExerciseName = 4
	HipRaiseExerciseNameWeightedBridgeWithLegExtension                  HipRaiseExerciseName = 5
	HipRaiseExerciseNameClamBridge                                      HipRaiseExerciseName = 6
	HipRaiseExerciseNameFrontKickTabletop                               HipRaiseExerciseName = 7
	HipRaiseExerciseNameWeightedFrontKickTabletop                       HipRaiseExerciseName = 8
	HipRaiseExerciseNameHipExtensionAndCross                            HipRaiseExerciseName = 9
	HipRaiseExerciseNameWeightedHipExtensionAndCross                    HipRaiseExerciseName = 10
	HipRaiseExerciseNameHipRaise                                        HipRaiseExerciseName = 11
	HipRaiseExerciseNameWeightedHipRaise                                HipRaiseExerciseName = 12
	HipRaiseExerciseNameHipRaiseWithFeetOnSwissBall                     HipRaiseExerciseName = 13
	HipRaiseExerciseNameWeightedHipRaiseWithFeetOnSwissBall             HipRaiseExerciseName = 14
	HipRaiseExerciseNameHipRaiseWithHeadOnBosuBall                      HipRaiseExerciseName = 15
	HipRaiseExerciseNameWeightedHipRaiseWithHeadOnBosuBall              HipRaiseExerciseName = 16
	HipRaiseExerciseNameHipRaiseWithHeadOnSwissBall                     HipRaiseExerciseName = 17
	HipRaiseExerciseNameWeightedHipRaiseWithHeadOnSwissBall             HipRaiseExerciseName = 18
	HipRaiseExerciseNameHipRaiseWithKneeSqueeze                         HipRaiseExerciseName = 19
	HipRaiseExerciseNameWeightedHipRaiseWithKneeSqueeze                 HipRaiseExerciseName = 20
	HipRaiseExerciseNameInclineRearLegExtension                         HipRaiseExerciseName = 21
	HipRaiseExerciseNameWeightedInclineRearLegExtension                 HipRaiseExerciseName = 22
	HipRaiseExerciseNameKettlebellSwing                                 HipRaiseExerciseName = 23
	HipRaiseExerciseNameMarchingHipRaise                                HipRaiseExerciseName = 24
	HipRaiseExerciseNameWeightedMarchingHipRaise                        HipRaiseExerciseName = 25
	HipRaiseExerciseNameMarchingHipRaiseWithFeetOnASwissBall            HipRaiseExerciseName = 26
	HipRaiseExerciseNameWeightedMarchingHipRaiseWithFeetOnASwissBall    HipRaiseExerciseName = 27
	HipRaiseExerciseNameReverseHipRaise                                 HipRaiseExerciseName = 28
	HipRaiseExerciseNameWeightedReverseHipRaise                         HipRaiseExerciseName = 29
	HipRaiseExerciseNameSingleLegHipRaise                               HipRaiseExerciseName = 30
	HipRaiseExerciseNameWeightedSingleLegHipRaise                       HipRaiseExerciseName = 31
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBench                HipRaiseExerciseName = 32
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBench        HipRaiseExerciseName = 33
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBosuBall             HipRaiseExerciseName = 34
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBosuBall     HipRaiseExerciseName = 35
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnFoamRoller           HipRaiseExerciseName = 36
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnFoamRoller   HipRaiseExerciseName = 37
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnMedicineBall         HipRaiseExerciseName = 38
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnMedicineBall HipRaiseExerciseName = 39
	HipRaiseExerciseNameSingleLegHipRaiseWithHeadOnBosuBall             HipRaiseExerciseName = 40
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithHeadOnBosuBall     HipRaiseExerciseName = 41
	HipRaiseExerciseNameWeightedClamBridge                              HipRaiseExerciseName = 42
	HipRaiseExerciseNameSingleLegSwissBallHipRaiseAndLegCurl            HipRaiseExerciseName = 43
	HipRaiseExerciseNameClams                                           HipRaiseExerciseName = 44
	HipRaiseExerciseNameInnerThighCircles                               HipRaiseExerciseName = 45 // Deprecated do not use
	HipRaiseExerciseNameInnerThighSideLift                              HipRaiseExerciseName = 46 // Deprecated do not use
	HipRaiseExerciseNameLegCircles                                      HipRaiseExerciseName = 47
	HipRaiseExerciseNameLegLift                                         HipRaiseExerciseName = 48
	HipRaiseExerciseNameLegLiftInExternalRotation                       HipRaiseExerciseName = 49
	HipRaiseExerciseNameInvalid                                         HipRaiseExerciseName = 0xFFFF // INVALID
)

var hipraiseexercisenametostrs = map[HipRaiseExerciseName]string{
	HipRaiseExerciseNameBarbellHipThrustOnFloor:                         "barbell_hip_thrust_on_floor",
	HipRaiseExerciseNameBarbellHipThrustWithBench:                       "barbell_hip_thrust_with_bench",
	HipRaiseExerciseNameBentKneeSwissBallReverseHipRaise:                "bent_knee_swiss_ball_reverse_hip_raise",
	HipRaiseExerciseNameWeightedBentKneeSwissBallReverseHipRaise:        "weighted_bent_knee_swiss_ball_reverse_hip_raise",
	HipRaiseExerciseNameBridgeWithLegExtension:                          "bridge_with_leg_extension",
	HipRaiseExerciseNameWeightedBridgeWithLegExtension:                  "weighted_bridge_with_leg_extension",
	HipRaiseExerciseNameClamBridge:                                      "clam_bridge",
	HipRaiseExerciseNameFrontKickTabletop:                               "front_kick_tabletop",
	HipRaiseExerciseNameWeightedFrontKickTabletop:                       "weighted_front_kick_tabletop",
	HipRaiseExerciseNameHipExtensionAndCross:                            "hip_extension_and_cross",
	HipRaiseExerciseNameWeightedHipExtensionAndCross:                    "weighted_hip_extension_and_cross",
	HipRaiseExerciseNameHipRaise:                                        "hip_raise",
	HipRaiseExerciseNameWeightedHipRaise:                                "weighted_hip_raise",
	HipRaiseExerciseNameHipRaiseWithFeetOnSwissBall:                     "hip_raise_with_feet_on_swiss_ball",
	HipRaiseExerciseNameWeightedHipRaiseWithFeetOnSwissBall:             "weighted_hip_raise_with_feet_on_swiss_ball",
	HipRaiseExerciseNameHipRaiseWithHeadOnBosuBall:                      "hip_raise_with_head_on_bosu_ball",
	HipRaiseExerciseNameWeightedHipRaiseWithHeadOnBosuBall:              "weighted_hip_raise_with_head_on_bosu_ball",
	HipRaiseExerciseNameHipRaiseWithHeadOnSwissBall:                     "hip_raise_with_head_on_swiss_ball",
	HipRaiseExerciseNameWeightedHipRaiseWithHeadOnSwissBall:             "weighted_hip_raise_with_head_on_swiss_ball",
	HipRaiseExerciseNameHipRaiseWithKneeSqueeze:                         "hip_raise_with_knee_squeeze",
	HipRaiseExerciseNameWeightedHipRaiseWithKneeSqueeze:                 "weighted_hip_raise_with_knee_squeeze",
	HipRaiseExerciseNameInclineRearLegExtension:                         "incline_rear_leg_extension",
	HipRaiseExerciseNameWeightedInclineRearLegExtension:                 "weighted_incline_rear_leg_extension",
	HipRaiseExerciseNameKettlebellSwing:                                 "kettlebell_swing",
	HipRaiseExerciseNameMarchingHipRaise:                                "marching_hip_raise",
	HipRaiseExerciseNameWeightedMarchingHipRaise:                        "weighted_marching_hip_raise",
	HipRaiseExerciseNameMarchingHipRaiseWithFeetOnASwissBall:            "marching_hip_raise_with_feet_on_a_swiss_ball",
	HipRaiseExerciseNameWeightedMarchingHipRaiseWithFeetOnASwissBall:    "weighted_marching_hip_raise_with_feet_on_a_swiss_ball",
	HipRaiseExerciseNameReverseHipRaise:                                 "reverse_hip_raise",
	HipRaiseExerciseNameWeightedReverseHipRaise:                         "weighted_reverse_hip_raise",
	HipRaiseExerciseNameSingleLegHipRaise:                               "single_leg_hip_raise",
	HipRaiseExerciseNameWeightedSingleLegHipRaise:                       "weighted_single_leg_hip_raise",
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBench:                "single_leg_hip_raise_with_foot_on_bench",
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBench:        "weighted_single_leg_hip_raise_with_foot_on_bench",
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBosuBall:             "single_leg_hip_raise_with_foot_on_bosu_ball",
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBosuBall:     "weighted_single_leg_hip_raise_with_foot_on_bosu_ball",
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnFoamRoller:           "single_leg_hip_raise_with_foot_on_foam_roller",
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnFoamRoller:   "weighted_single_leg_hip_raise_with_foot_on_foam_roller",
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnMedicineBall:         "single_leg_hip_raise_with_foot_on_medicine_ball",
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnMedicineBall: "weighted_single_leg_hip_raise_with_foot_on_medicine_ball",
	HipRaiseExerciseNameSingleLegHipRaiseWithHeadOnBosuBall:             "single_leg_hip_raise_with_head_on_bosu_ball",
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithHeadOnBosuBall:     "weighted_single_leg_hip_raise_with_head_on_bosu_ball",
	HipRaiseExerciseNameWeightedClamBridge:                              "weighted_clam_bridge",
	HipRaiseExerciseNameSingleLegSwissBallHipRaiseAndLegCurl:            "single_leg_swiss_ball_hip_raise_and_leg_curl",
	HipRaiseExerciseNameClams:                                           "clams",
	HipRaiseExerciseNameInnerThighCircles:                               "inner_thigh_circles",
	HipRaiseExerciseNameInnerThighSideLift:                              "inner_thigh_side_lift",
	HipRaiseExerciseNameLegCircles:                                      "leg_circles",
	HipRaiseExerciseNameLegLift:                                         "leg_lift",
	HipRaiseExerciseNameLegLiftInExternalRotation:                       "leg_lift_in_external_rotation",
	HipRaiseExerciseNameInvalid:                                         "invalid",
}

func (h HipRaiseExerciseName) String() string {
	val, ok := hipraiseexercisenametostrs[h]
	if !ok {
		return strconv.FormatUint(uint64(h), 10)
	}
	return val
}

var strtohipraiseexercisename = func() map[string]HipRaiseExerciseName {
	m := make(map[string]HipRaiseExerciseName)
	for t, str := range hipraiseexercisenametostrs {
		m[str] = HipRaiseExerciseName(t)
	}
	return m
}()

// FromString parse string into HipRaiseExerciseName constant it's represent, return HipRaiseExerciseNameInvalid if not found.
func HipRaiseExerciseNameFromString(s string) HipRaiseExerciseName {
	val, ok := strtohipraiseexercisename[s]
	if !ok {
		return strtohipraiseexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListHipRaiseExerciseName() []HipRaiseExerciseName {
	vs := make([]HipRaiseExerciseName, 0, len(hipraiseexercisenametostrs))
	for i := range hipraiseexercisenametostrs {
		vs = append(vs, HipRaiseExerciseName(i))
	}
	return vs
}
