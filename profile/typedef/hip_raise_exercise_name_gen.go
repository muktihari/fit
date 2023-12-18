// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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
	HipRaiseExerciseNameInvalid                                         HipRaiseExerciseName = 0xFFFF
)

func (h HipRaiseExerciseName) String() string {
	switch h {
	case HipRaiseExerciseNameBarbellHipThrustOnFloor:
		return "barbell_hip_thrust_on_floor"
	case HipRaiseExerciseNameBarbellHipThrustWithBench:
		return "barbell_hip_thrust_with_bench"
	case HipRaiseExerciseNameBentKneeSwissBallReverseHipRaise:
		return "bent_knee_swiss_ball_reverse_hip_raise"
	case HipRaiseExerciseNameWeightedBentKneeSwissBallReverseHipRaise:
		return "weighted_bent_knee_swiss_ball_reverse_hip_raise"
	case HipRaiseExerciseNameBridgeWithLegExtension:
		return "bridge_with_leg_extension"
	case HipRaiseExerciseNameWeightedBridgeWithLegExtension:
		return "weighted_bridge_with_leg_extension"
	case HipRaiseExerciseNameClamBridge:
		return "clam_bridge"
	case HipRaiseExerciseNameFrontKickTabletop:
		return "front_kick_tabletop"
	case HipRaiseExerciseNameWeightedFrontKickTabletop:
		return "weighted_front_kick_tabletop"
	case HipRaiseExerciseNameHipExtensionAndCross:
		return "hip_extension_and_cross"
	case HipRaiseExerciseNameWeightedHipExtensionAndCross:
		return "weighted_hip_extension_and_cross"
	case HipRaiseExerciseNameHipRaise:
		return "hip_raise"
	case HipRaiseExerciseNameWeightedHipRaise:
		return "weighted_hip_raise"
	case HipRaiseExerciseNameHipRaiseWithFeetOnSwissBall:
		return "hip_raise_with_feet_on_swiss_ball"
	case HipRaiseExerciseNameWeightedHipRaiseWithFeetOnSwissBall:
		return "weighted_hip_raise_with_feet_on_swiss_ball"
	case HipRaiseExerciseNameHipRaiseWithHeadOnBosuBall:
		return "hip_raise_with_head_on_bosu_ball"
	case HipRaiseExerciseNameWeightedHipRaiseWithHeadOnBosuBall:
		return "weighted_hip_raise_with_head_on_bosu_ball"
	case HipRaiseExerciseNameHipRaiseWithHeadOnSwissBall:
		return "hip_raise_with_head_on_swiss_ball"
	case HipRaiseExerciseNameWeightedHipRaiseWithHeadOnSwissBall:
		return "weighted_hip_raise_with_head_on_swiss_ball"
	case HipRaiseExerciseNameHipRaiseWithKneeSqueeze:
		return "hip_raise_with_knee_squeeze"
	case HipRaiseExerciseNameWeightedHipRaiseWithKneeSqueeze:
		return "weighted_hip_raise_with_knee_squeeze"
	case HipRaiseExerciseNameInclineRearLegExtension:
		return "incline_rear_leg_extension"
	case HipRaiseExerciseNameWeightedInclineRearLegExtension:
		return "weighted_incline_rear_leg_extension"
	case HipRaiseExerciseNameKettlebellSwing:
		return "kettlebell_swing"
	case HipRaiseExerciseNameMarchingHipRaise:
		return "marching_hip_raise"
	case HipRaiseExerciseNameWeightedMarchingHipRaise:
		return "weighted_marching_hip_raise"
	case HipRaiseExerciseNameMarchingHipRaiseWithFeetOnASwissBall:
		return "marching_hip_raise_with_feet_on_a_swiss_ball"
	case HipRaiseExerciseNameWeightedMarchingHipRaiseWithFeetOnASwissBall:
		return "weighted_marching_hip_raise_with_feet_on_a_swiss_ball"
	case HipRaiseExerciseNameReverseHipRaise:
		return "reverse_hip_raise"
	case HipRaiseExerciseNameWeightedReverseHipRaise:
		return "weighted_reverse_hip_raise"
	case HipRaiseExerciseNameSingleLegHipRaise:
		return "single_leg_hip_raise"
	case HipRaiseExerciseNameWeightedSingleLegHipRaise:
		return "weighted_single_leg_hip_raise"
	case HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBench:
		return "single_leg_hip_raise_with_foot_on_bench"
	case HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBench:
		return "weighted_single_leg_hip_raise_with_foot_on_bench"
	case HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBosuBall:
		return "single_leg_hip_raise_with_foot_on_bosu_ball"
	case HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBosuBall:
		return "weighted_single_leg_hip_raise_with_foot_on_bosu_ball"
	case HipRaiseExerciseNameSingleLegHipRaiseWithFootOnFoamRoller:
		return "single_leg_hip_raise_with_foot_on_foam_roller"
	case HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnFoamRoller:
		return "weighted_single_leg_hip_raise_with_foot_on_foam_roller"
	case HipRaiseExerciseNameSingleLegHipRaiseWithFootOnMedicineBall:
		return "single_leg_hip_raise_with_foot_on_medicine_ball"
	case HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnMedicineBall:
		return "weighted_single_leg_hip_raise_with_foot_on_medicine_ball"
	case HipRaiseExerciseNameSingleLegHipRaiseWithHeadOnBosuBall:
		return "single_leg_hip_raise_with_head_on_bosu_ball"
	case HipRaiseExerciseNameWeightedSingleLegHipRaiseWithHeadOnBosuBall:
		return "weighted_single_leg_hip_raise_with_head_on_bosu_ball"
	case HipRaiseExerciseNameWeightedClamBridge:
		return "weighted_clam_bridge"
	case HipRaiseExerciseNameSingleLegSwissBallHipRaiseAndLegCurl:
		return "single_leg_swiss_ball_hip_raise_and_leg_curl"
	case HipRaiseExerciseNameClams:
		return "clams"
	case HipRaiseExerciseNameInnerThighCircles:
		return "inner_thigh_circles"
	case HipRaiseExerciseNameInnerThighSideLift:
		return "inner_thigh_side_lift"
	case HipRaiseExerciseNameLegCircles:
		return "leg_circles"
	case HipRaiseExerciseNameLegLift:
		return "leg_lift"
	case HipRaiseExerciseNameLegLiftInExternalRotation:
		return "leg_lift_in_external_rotation"
	default:
		return "HipRaiseExerciseNameInvalid(" + strconv.FormatUint(uint64(h), 10) + ")"
	}
}

// FromString parse string into HipRaiseExerciseName constant it's represent, return HipRaiseExerciseNameInvalid if not found.
func HipRaiseExerciseNameFromString(s string) HipRaiseExerciseName {
	switch s {
	case "barbell_hip_thrust_on_floor":
		return HipRaiseExerciseNameBarbellHipThrustOnFloor
	case "barbell_hip_thrust_with_bench":
		return HipRaiseExerciseNameBarbellHipThrustWithBench
	case "bent_knee_swiss_ball_reverse_hip_raise":
		return HipRaiseExerciseNameBentKneeSwissBallReverseHipRaise
	case "weighted_bent_knee_swiss_ball_reverse_hip_raise":
		return HipRaiseExerciseNameWeightedBentKneeSwissBallReverseHipRaise
	case "bridge_with_leg_extension":
		return HipRaiseExerciseNameBridgeWithLegExtension
	case "weighted_bridge_with_leg_extension":
		return HipRaiseExerciseNameWeightedBridgeWithLegExtension
	case "clam_bridge":
		return HipRaiseExerciseNameClamBridge
	case "front_kick_tabletop":
		return HipRaiseExerciseNameFrontKickTabletop
	case "weighted_front_kick_tabletop":
		return HipRaiseExerciseNameWeightedFrontKickTabletop
	case "hip_extension_and_cross":
		return HipRaiseExerciseNameHipExtensionAndCross
	case "weighted_hip_extension_and_cross":
		return HipRaiseExerciseNameWeightedHipExtensionAndCross
	case "hip_raise":
		return HipRaiseExerciseNameHipRaise
	case "weighted_hip_raise":
		return HipRaiseExerciseNameWeightedHipRaise
	case "hip_raise_with_feet_on_swiss_ball":
		return HipRaiseExerciseNameHipRaiseWithFeetOnSwissBall
	case "weighted_hip_raise_with_feet_on_swiss_ball":
		return HipRaiseExerciseNameWeightedHipRaiseWithFeetOnSwissBall
	case "hip_raise_with_head_on_bosu_ball":
		return HipRaiseExerciseNameHipRaiseWithHeadOnBosuBall
	case "weighted_hip_raise_with_head_on_bosu_ball":
		return HipRaiseExerciseNameWeightedHipRaiseWithHeadOnBosuBall
	case "hip_raise_with_head_on_swiss_ball":
		return HipRaiseExerciseNameHipRaiseWithHeadOnSwissBall
	case "weighted_hip_raise_with_head_on_swiss_ball":
		return HipRaiseExerciseNameWeightedHipRaiseWithHeadOnSwissBall
	case "hip_raise_with_knee_squeeze":
		return HipRaiseExerciseNameHipRaiseWithKneeSqueeze
	case "weighted_hip_raise_with_knee_squeeze":
		return HipRaiseExerciseNameWeightedHipRaiseWithKneeSqueeze
	case "incline_rear_leg_extension":
		return HipRaiseExerciseNameInclineRearLegExtension
	case "weighted_incline_rear_leg_extension":
		return HipRaiseExerciseNameWeightedInclineRearLegExtension
	case "kettlebell_swing":
		return HipRaiseExerciseNameKettlebellSwing
	case "marching_hip_raise":
		return HipRaiseExerciseNameMarchingHipRaise
	case "weighted_marching_hip_raise":
		return HipRaiseExerciseNameWeightedMarchingHipRaise
	case "marching_hip_raise_with_feet_on_a_swiss_ball":
		return HipRaiseExerciseNameMarchingHipRaiseWithFeetOnASwissBall
	case "weighted_marching_hip_raise_with_feet_on_a_swiss_ball":
		return HipRaiseExerciseNameWeightedMarchingHipRaiseWithFeetOnASwissBall
	case "reverse_hip_raise":
		return HipRaiseExerciseNameReverseHipRaise
	case "weighted_reverse_hip_raise":
		return HipRaiseExerciseNameWeightedReverseHipRaise
	case "single_leg_hip_raise":
		return HipRaiseExerciseNameSingleLegHipRaise
	case "weighted_single_leg_hip_raise":
		return HipRaiseExerciseNameWeightedSingleLegHipRaise
	case "single_leg_hip_raise_with_foot_on_bench":
		return HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBench
	case "weighted_single_leg_hip_raise_with_foot_on_bench":
		return HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBench
	case "single_leg_hip_raise_with_foot_on_bosu_ball":
		return HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBosuBall
	case "weighted_single_leg_hip_raise_with_foot_on_bosu_ball":
		return HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBosuBall
	case "single_leg_hip_raise_with_foot_on_foam_roller":
		return HipRaiseExerciseNameSingleLegHipRaiseWithFootOnFoamRoller
	case "weighted_single_leg_hip_raise_with_foot_on_foam_roller":
		return HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnFoamRoller
	case "single_leg_hip_raise_with_foot_on_medicine_ball":
		return HipRaiseExerciseNameSingleLegHipRaiseWithFootOnMedicineBall
	case "weighted_single_leg_hip_raise_with_foot_on_medicine_ball":
		return HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnMedicineBall
	case "single_leg_hip_raise_with_head_on_bosu_ball":
		return HipRaiseExerciseNameSingleLegHipRaiseWithHeadOnBosuBall
	case "weighted_single_leg_hip_raise_with_head_on_bosu_ball":
		return HipRaiseExerciseNameWeightedSingleLegHipRaiseWithHeadOnBosuBall
	case "weighted_clam_bridge":
		return HipRaiseExerciseNameWeightedClamBridge
	case "single_leg_swiss_ball_hip_raise_and_leg_curl":
		return HipRaiseExerciseNameSingleLegSwissBallHipRaiseAndLegCurl
	case "clams":
		return HipRaiseExerciseNameClams
	case "inner_thigh_circles":
		return HipRaiseExerciseNameInnerThighCircles
	case "inner_thigh_side_lift":
		return HipRaiseExerciseNameInnerThighSideLift
	case "leg_circles":
		return HipRaiseExerciseNameLegCircles
	case "leg_lift":
		return HipRaiseExerciseNameLegLift
	case "leg_lift_in_external_rotation":
		return HipRaiseExerciseNameLegLiftInExternalRotation
	default:
		return HipRaiseExerciseNameInvalid
	}
}

// List returns all constants.
func ListHipRaiseExerciseName() []HipRaiseExerciseName {
	return []HipRaiseExerciseName{
		HipRaiseExerciseNameBarbellHipThrustOnFloor,
		HipRaiseExerciseNameBarbellHipThrustWithBench,
		HipRaiseExerciseNameBentKneeSwissBallReverseHipRaise,
		HipRaiseExerciseNameWeightedBentKneeSwissBallReverseHipRaise,
		HipRaiseExerciseNameBridgeWithLegExtension,
		HipRaiseExerciseNameWeightedBridgeWithLegExtension,
		HipRaiseExerciseNameClamBridge,
		HipRaiseExerciseNameFrontKickTabletop,
		HipRaiseExerciseNameWeightedFrontKickTabletop,
		HipRaiseExerciseNameHipExtensionAndCross,
		HipRaiseExerciseNameWeightedHipExtensionAndCross,
		HipRaiseExerciseNameHipRaise,
		HipRaiseExerciseNameWeightedHipRaise,
		HipRaiseExerciseNameHipRaiseWithFeetOnSwissBall,
		HipRaiseExerciseNameWeightedHipRaiseWithFeetOnSwissBall,
		HipRaiseExerciseNameHipRaiseWithHeadOnBosuBall,
		HipRaiseExerciseNameWeightedHipRaiseWithHeadOnBosuBall,
		HipRaiseExerciseNameHipRaiseWithHeadOnSwissBall,
		HipRaiseExerciseNameWeightedHipRaiseWithHeadOnSwissBall,
		HipRaiseExerciseNameHipRaiseWithKneeSqueeze,
		HipRaiseExerciseNameWeightedHipRaiseWithKneeSqueeze,
		HipRaiseExerciseNameInclineRearLegExtension,
		HipRaiseExerciseNameWeightedInclineRearLegExtension,
		HipRaiseExerciseNameKettlebellSwing,
		HipRaiseExerciseNameMarchingHipRaise,
		HipRaiseExerciseNameWeightedMarchingHipRaise,
		HipRaiseExerciseNameMarchingHipRaiseWithFeetOnASwissBall,
		HipRaiseExerciseNameWeightedMarchingHipRaiseWithFeetOnASwissBall,
		HipRaiseExerciseNameReverseHipRaise,
		HipRaiseExerciseNameWeightedReverseHipRaise,
		HipRaiseExerciseNameSingleLegHipRaise,
		HipRaiseExerciseNameWeightedSingleLegHipRaise,
		HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBench,
		HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBench,
		HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBosuBall,
		HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBosuBall,
		HipRaiseExerciseNameSingleLegHipRaiseWithFootOnFoamRoller,
		HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnFoamRoller,
		HipRaiseExerciseNameSingleLegHipRaiseWithFootOnMedicineBall,
		HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnMedicineBall,
		HipRaiseExerciseNameSingleLegHipRaiseWithHeadOnBosuBall,
		HipRaiseExerciseNameWeightedSingleLegHipRaiseWithHeadOnBosuBall,
		HipRaiseExerciseNameWeightedClamBridge,
		HipRaiseExerciseNameSingleLegSwissBallHipRaiseAndLegCurl,
		HipRaiseExerciseNameClams,
		HipRaiseExerciseNameInnerThighCircles,
		HipRaiseExerciseNameInnerThighSideLift,
		HipRaiseExerciseNameLegCircles,
		HipRaiseExerciseNameLegLift,
		HipRaiseExerciseNameLegLiftInExternalRotation,
	}
}
