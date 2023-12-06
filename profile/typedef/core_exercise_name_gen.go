// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CoreExerciseName uint16

const (
	CoreExerciseNameAbsJabs                          CoreExerciseName = 0
	CoreExerciseNameWeightedAbsJabs                  CoreExerciseName = 1
	CoreExerciseNameAlternatingPlateReach            CoreExerciseName = 2
	CoreExerciseNameBarbellRollout                   CoreExerciseName = 3
	CoreExerciseNameWeightedBarbellRollout           CoreExerciseName = 4
	CoreExerciseNameBodyBarObliqueTwist              CoreExerciseName = 5
	CoreExerciseNameCableCorePress                   CoreExerciseName = 6
	CoreExerciseNameCableSideBend                    CoreExerciseName = 7
	CoreExerciseNameSideBend                         CoreExerciseName = 8
	CoreExerciseNameWeightedSideBend                 CoreExerciseName = 9
	CoreExerciseNameCrescentCircle                   CoreExerciseName = 10
	CoreExerciseNameWeightedCrescentCircle           CoreExerciseName = 11
	CoreExerciseNameCyclingRussianTwist              CoreExerciseName = 12
	CoreExerciseNameWeightedCyclingRussianTwist      CoreExerciseName = 13
	CoreExerciseNameElevatedFeetRussianTwist         CoreExerciseName = 14
	CoreExerciseNameWeightedElevatedFeetRussianTwist CoreExerciseName = 15
	CoreExerciseNameHalfTurkishGetUp                 CoreExerciseName = 16
	CoreExerciseNameKettlebellWindmill               CoreExerciseName = 17
	CoreExerciseNameKneelingAbWheel                  CoreExerciseName = 18
	CoreExerciseNameWeightedKneelingAbWheel          CoreExerciseName = 19
	CoreExerciseNameModifiedFrontLever               CoreExerciseName = 20
	CoreExerciseNameOpenKneeTucks                    CoreExerciseName = 21
	CoreExerciseNameWeightedOpenKneeTucks            CoreExerciseName = 22
	CoreExerciseNameSideAbsLegLift                   CoreExerciseName = 23
	CoreExerciseNameWeightedSideAbsLegLift           CoreExerciseName = 24
	CoreExerciseNameSwissBallJackknife               CoreExerciseName = 25
	CoreExerciseNameWeightedSwissBallJackknife       CoreExerciseName = 26
	CoreExerciseNameSwissBallPike                    CoreExerciseName = 27
	CoreExerciseNameWeightedSwissBallPike            CoreExerciseName = 28
	CoreExerciseNameSwissBallRollout                 CoreExerciseName = 29
	CoreExerciseNameWeightedSwissBallRollout         CoreExerciseName = 30
	CoreExerciseNameTriangleHipPress                 CoreExerciseName = 31
	CoreExerciseNameWeightedTriangleHipPress         CoreExerciseName = 32
	CoreExerciseNameTrxSuspendedJackknife            CoreExerciseName = 33
	CoreExerciseNameWeightedTrxSuspendedJackknife    CoreExerciseName = 34
	CoreExerciseNameUBoat                            CoreExerciseName = 35
	CoreExerciseNameWeightedUBoat                    CoreExerciseName = 36
	CoreExerciseNameWindmillSwitches                 CoreExerciseName = 37
	CoreExerciseNameWeightedWindmillSwitches         CoreExerciseName = 38
	CoreExerciseNameAlternatingSlideOut              CoreExerciseName = 39
	CoreExerciseNameWeightedAlternatingSlideOut      CoreExerciseName = 40
	CoreExerciseNameGhdBackExtensions                CoreExerciseName = 41
	CoreExerciseNameWeightedGhdBackExtensions        CoreExerciseName = 42
	CoreExerciseNameOverheadWalk                     CoreExerciseName = 43
	CoreExerciseNameInchworm                         CoreExerciseName = 44
	CoreExerciseNameWeightedModifiedFrontLever       CoreExerciseName = 45
	CoreExerciseNameRussianTwist                     CoreExerciseName = 46
	CoreExerciseNameAbdominalLegRotations            CoreExerciseName = 47 // Deprecated do not use
	CoreExerciseNameArmAndLegExtensionOnKnees        CoreExerciseName = 48
	CoreExerciseNameBicycle                          CoreExerciseName = 49
	CoreExerciseNameBicepCurlWithLegExtension        CoreExerciseName = 50
	CoreExerciseNameCatCow                           CoreExerciseName = 51
	CoreExerciseNameCorkscrew                        CoreExerciseName = 52
	CoreExerciseNameCrissCross                       CoreExerciseName = 53
	CoreExerciseNameCrissCrossWithBall               CoreExerciseName = 54 // Deprecated do not use
	CoreExerciseNameDoubleLegStretch                 CoreExerciseName = 55
	CoreExerciseNameKneeFolds                        CoreExerciseName = 56
	CoreExerciseNameLowerLift                        CoreExerciseName = 57
	CoreExerciseNameNeckPull                         CoreExerciseName = 58
	CoreExerciseNamePelvicClocks                     CoreExerciseName = 59
	CoreExerciseNameRollOver                         CoreExerciseName = 60
	CoreExerciseNameRollUp                           CoreExerciseName = 61
	CoreExerciseNameRolling                          CoreExerciseName = 62
	CoreExerciseNameRowing1                          CoreExerciseName = 63
	CoreExerciseNameRowing2                          CoreExerciseName = 64
	CoreExerciseNameScissors                         CoreExerciseName = 65
	CoreExerciseNameSingleLegCircles                 CoreExerciseName = 66
	CoreExerciseNameSingleLegStretch                 CoreExerciseName = 67
	CoreExerciseNameSnakeTwist1And2                  CoreExerciseName = 68 // Deprecated do not use
	CoreExerciseNameSwan                             CoreExerciseName = 69
	CoreExerciseNameSwimming                         CoreExerciseName = 70
	CoreExerciseNameTeaser                           CoreExerciseName = 71
	CoreExerciseNameTheHundred                       CoreExerciseName = 72
	CoreExerciseNameInvalid                          CoreExerciseName = 0xFFFF // INVALID
)

var coreexercisenametostrs = map[CoreExerciseName]string{
	CoreExerciseNameAbsJabs:                          "abs_jabs",
	CoreExerciseNameWeightedAbsJabs:                  "weighted_abs_jabs",
	CoreExerciseNameAlternatingPlateReach:            "alternating_plate_reach",
	CoreExerciseNameBarbellRollout:                   "barbell_rollout",
	CoreExerciseNameWeightedBarbellRollout:           "weighted_barbell_rollout",
	CoreExerciseNameBodyBarObliqueTwist:              "body_bar_oblique_twist",
	CoreExerciseNameCableCorePress:                   "cable_core_press",
	CoreExerciseNameCableSideBend:                    "cable_side_bend",
	CoreExerciseNameSideBend:                         "side_bend",
	CoreExerciseNameWeightedSideBend:                 "weighted_side_bend",
	CoreExerciseNameCrescentCircle:                   "crescent_circle",
	CoreExerciseNameWeightedCrescentCircle:           "weighted_crescent_circle",
	CoreExerciseNameCyclingRussianTwist:              "cycling_russian_twist",
	CoreExerciseNameWeightedCyclingRussianTwist:      "weighted_cycling_russian_twist",
	CoreExerciseNameElevatedFeetRussianTwist:         "elevated_feet_russian_twist",
	CoreExerciseNameWeightedElevatedFeetRussianTwist: "weighted_elevated_feet_russian_twist",
	CoreExerciseNameHalfTurkishGetUp:                 "half_turkish_get_up",
	CoreExerciseNameKettlebellWindmill:               "kettlebell_windmill",
	CoreExerciseNameKneelingAbWheel:                  "kneeling_ab_wheel",
	CoreExerciseNameWeightedKneelingAbWheel:          "weighted_kneeling_ab_wheel",
	CoreExerciseNameModifiedFrontLever:               "modified_front_lever",
	CoreExerciseNameOpenKneeTucks:                    "open_knee_tucks",
	CoreExerciseNameWeightedOpenKneeTucks:            "weighted_open_knee_tucks",
	CoreExerciseNameSideAbsLegLift:                   "side_abs_leg_lift",
	CoreExerciseNameWeightedSideAbsLegLift:           "weighted_side_abs_leg_lift",
	CoreExerciseNameSwissBallJackknife:               "swiss_ball_jackknife",
	CoreExerciseNameWeightedSwissBallJackknife:       "weighted_swiss_ball_jackknife",
	CoreExerciseNameSwissBallPike:                    "swiss_ball_pike",
	CoreExerciseNameWeightedSwissBallPike:            "weighted_swiss_ball_pike",
	CoreExerciseNameSwissBallRollout:                 "swiss_ball_rollout",
	CoreExerciseNameWeightedSwissBallRollout:         "weighted_swiss_ball_rollout",
	CoreExerciseNameTriangleHipPress:                 "triangle_hip_press",
	CoreExerciseNameWeightedTriangleHipPress:         "weighted_triangle_hip_press",
	CoreExerciseNameTrxSuspendedJackknife:            "trx_suspended_jackknife",
	CoreExerciseNameWeightedTrxSuspendedJackknife:    "weighted_trx_suspended_jackknife",
	CoreExerciseNameUBoat:                            "u_boat",
	CoreExerciseNameWeightedUBoat:                    "weighted_u_boat",
	CoreExerciseNameWindmillSwitches:                 "windmill_switches",
	CoreExerciseNameWeightedWindmillSwitches:         "weighted_windmill_switches",
	CoreExerciseNameAlternatingSlideOut:              "alternating_slide_out",
	CoreExerciseNameWeightedAlternatingSlideOut:      "weighted_alternating_slide_out",
	CoreExerciseNameGhdBackExtensions:                "ghd_back_extensions",
	CoreExerciseNameWeightedGhdBackExtensions:        "weighted_ghd_back_extensions",
	CoreExerciseNameOverheadWalk:                     "overhead_walk",
	CoreExerciseNameInchworm:                         "inchworm",
	CoreExerciseNameWeightedModifiedFrontLever:       "weighted_modified_front_lever",
	CoreExerciseNameRussianTwist:                     "russian_twist",
	CoreExerciseNameAbdominalLegRotations:            "abdominal_leg_rotations",
	CoreExerciseNameArmAndLegExtensionOnKnees:        "arm_and_leg_extension_on_knees",
	CoreExerciseNameBicycle:                          "bicycle",
	CoreExerciseNameBicepCurlWithLegExtension:        "bicep_curl_with_leg_extension",
	CoreExerciseNameCatCow:                           "cat_cow",
	CoreExerciseNameCorkscrew:                        "corkscrew",
	CoreExerciseNameCrissCross:                       "criss_cross",
	CoreExerciseNameCrissCrossWithBall:               "criss_cross_with_ball",
	CoreExerciseNameDoubleLegStretch:                 "double_leg_stretch",
	CoreExerciseNameKneeFolds:                        "knee_folds",
	CoreExerciseNameLowerLift:                        "lower_lift",
	CoreExerciseNameNeckPull:                         "neck_pull",
	CoreExerciseNamePelvicClocks:                     "pelvic_clocks",
	CoreExerciseNameRollOver:                         "roll_over",
	CoreExerciseNameRollUp:                           "roll_up",
	CoreExerciseNameRolling:                          "rolling",
	CoreExerciseNameRowing1:                          "rowing_1",
	CoreExerciseNameRowing2:                          "rowing_2",
	CoreExerciseNameScissors:                         "scissors",
	CoreExerciseNameSingleLegCircles:                 "single_leg_circles",
	CoreExerciseNameSingleLegStretch:                 "single_leg_stretch",
	CoreExerciseNameSnakeTwist1And2:                  "snake_twist_1_and_2",
	CoreExerciseNameSwan:                             "swan",
	CoreExerciseNameSwimming:                         "swimming",
	CoreExerciseNameTeaser:                           "teaser",
	CoreExerciseNameTheHundred:                       "the_hundred",
	CoreExerciseNameInvalid:                          "invalid",
}

func (c CoreExerciseName) String() string {
	val, ok := coreexercisenametostrs[c]
	if !ok {
		return strconv.FormatUint(uint64(c), 10)
	}
	return val
}

var strtocoreexercisename = func() map[string]CoreExerciseName {
	m := make(map[string]CoreExerciseName)
	for t, str := range coreexercisenametostrs {
		m[str] = CoreExerciseName(t)
	}
	return m
}()

// FromString parse string into CoreExerciseName constant it's represent, return CoreExerciseNameInvalid if not found.
func CoreExerciseNameFromString(s string) CoreExerciseName {
	val, ok := strtocoreexercisename[s]
	if !ok {
		return strtocoreexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListCoreExerciseName() []CoreExerciseName {
	vs := make([]CoreExerciseName, 0, len(coreexercisenametostrs))
	for i := range coreexercisenametostrs {
		vs = append(vs, CoreExerciseName(i))
	}
	return vs
}
