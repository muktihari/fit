// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CalfRaiseExerciseName uint16

const (
	CalfRaiseExerciseName3WayCalfRaise                      CalfRaiseExerciseName = 0
	CalfRaiseExerciseName3WayWeightedCalfRaise              CalfRaiseExerciseName = 1
	CalfRaiseExerciseName3WaySingleLegCalfRaise             CalfRaiseExerciseName = 2
	CalfRaiseExerciseName3WayWeightedSingleLegCalfRaise     CalfRaiseExerciseName = 3
	CalfRaiseExerciseNameDonkeyCalfRaise                    CalfRaiseExerciseName = 4
	CalfRaiseExerciseNameWeightedDonkeyCalfRaise            CalfRaiseExerciseName = 5
	CalfRaiseExerciseNameSeatedCalfRaise                    CalfRaiseExerciseName = 6
	CalfRaiseExerciseNameWeightedSeatedCalfRaise            CalfRaiseExerciseName = 7
	CalfRaiseExerciseNameSeatedDumbbellToeRaise             CalfRaiseExerciseName = 8
	CalfRaiseExerciseNameSingleLegBentKneeCalfRaise         CalfRaiseExerciseName = 9
	CalfRaiseExerciseNameWeightedSingleLegBentKneeCalfRaise CalfRaiseExerciseName = 10
	CalfRaiseExerciseNameSingleLegDeclinePushUp             CalfRaiseExerciseName = 11
	CalfRaiseExerciseNameSingleLegDonkeyCalfRaise           CalfRaiseExerciseName = 12
	CalfRaiseExerciseNameWeightedSingleLegDonkeyCalfRaise   CalfRaiseExerciseName = 13
	CalfRaiseExerciseNameSingleLegHipRaiseWithKneeHold      CalfRaiseExerciseName = 14
	CalfRaiseExerciseNameSingleLegStandingCalfRaise         CalfRaiseExerciseName = 15
	CalfRaiseExerciseNameSingleLegStandingDumbbellCalfRaise CalfRaiseExerciseName = 16
	CalfRaiseExerciseNameStandingBarbellCalfRaise           CalfRaiseExerciseName = 17
	CalfRaiseExerciseNameStandingCalfRaise                  CalfRaiseExerciseName = 18
	CalfRaiseExerciseNameWeightedStandingCalfRaise          CalfRaiseExerciseName = 19
	CalfRaiseExerciseNameStandingDumbbellCalfRaise          CalfRaiseExerciseName = 20
	CalfRaiseExerciseNameInvalid                            CalfRaiseExerciseName = 0xFFFF
)

func (c CalfRaiseExerciseName) Uint16() uint16 { return uint16(c) }

func (c CalfRaiseExerciseName) String() string {
	switch c {
	case CalfRaiseExerciseName3WayCalfRaise:
		return "3_way_calf_raise"
	case CalfRaiseExerciseName3WayWeightedCalfRaise:
		return "3_way_weighted_calf_raise"
	case CalfRaiseExerciseName3WaySingleLegCalfRaise:
		return "3_way_single_leg_calf_raise"
	case CalfRaiseExerciseName3WayWeightedSingleLegCalfRaise:
		return "3_way_weighted_single_leg_calf_raise"
	case CalfRaiseExerciseNameDonkeyCalfRaise:
		return "donkey_calf_raise"
	case CalfRaiseExerciseNameWeightedDonkeyCalfRaise:
		return "weighted_donkey_calf_raise"
	case CalfRaiseExerciseNameSeatedCalfRaise:
		return "seated_calf_raise"
	case CalfRaiseExerciseNameWeightedSeatedCalfRaise:
		return "weighted_seated_calf_raise"
	case CalfRaiseExerciseNameSeatedDumbbellToeRaise:
		return "seated_dumbbell_toe_raise"
	case CalfRaiseExerciseNameSingleLegBentKneeCalfRaise:
		return "single_leg_bent_knee_calf_raise"
	case CalfRaiseExerciseNameWeightedSingleLegBentKneeCalfRaise:
		return "weighted_single_leg_bent_knee_calf_raise"
	case CalfRaiseExerciseNameSingleLegDeclinePushUp:
		return "single_leg_decline_push_up"
	case CalfRaiseExerciseNameSingleLegDonkeyCalfRaise:
		return "single_leg_donkey_calf_raise"
	case CalfRaiseExerciseNameWeightedSingleLegDonkeyCalfRaise:
		return "weighted_single_leg_donkey_calf_raise"
	case CalfRaiseExerciseNameSingleLegHipRaiseWithKneeHold:
		return "single_leg_hip_raise_with_knee_hold"
	case CalfRaiseExerciseNameSingleLegStandingCalfRaise:
		return "single_leg_standing_calf_raise"
	case CalfRaiseExerciseNameSingleLegStandingDumbbellCalfRaise:
		return "single_leg_standing_dumbbell_calf_raise"
	case CalfRaiseExerciseNameStandingBarbellCalfRaise:
		return "standing_barbell_calf_raise"
	case CalfRaiseExerciseNameStandingCalfRaise:
		return "standing_calf_raise"
	case CalfRaiseExerciseNameWeightedStandingCalfRaise:
		return "weighted_standing_calf_raise"
	case CalfRaiseExerciseNameStandingDumbbellCalfRaise:
		return "standing_dumbbell_calf_raise"
	default:
		return "CalfRaiseExerciseNameInvalid(" + strconv.FormatUint(uint64(c), 10) + ")"
	}
}

// FromString parse string into CalfRaiseExerciseName constant it's represent, return CalfRaiseExerciseNameInvalid if not found.
func CalfRaiseExerciseNameFromString(s string) CalfRaiseExerciseName {
	switch s {
	case "3_way_calf_raise":
		return CalfRaiseExerciseName3WayCalfRaise
	case "3_way_weighted_calf_raise":
		return CalfRaiseExerciseName3WayWeightedCalfRaise
	case "3_way_single_leg_calf_raise":
		return CalfRaiseExerciseName3WaySingleLegCalfRaise
	case "3_way_weighted_single_leg_calf_raise":
		return CalfRaiseExerciseName3WayWeightedSingleLegCalfRaise
	case "donkey_calf_raise":
		return CalfRaiseExerciseNameDonkeyCalfRaise
	case "weighted_donkey_calf_raise":
		return CalfRaiseExerciseNameWeightedDonkeyCalfRaise
	case "seated_calf_raise":
		return CalfRaiseExerciseNameSeatedCalfRaise
	case "weighted_seated_calf_raise":
		return CalfRaiseExerciseNameWeightedSeatedCalfRaise
	case "seated_dumbbell_toe_raise":
		return CalfRaiseExerciseNameSeatedDumbbellToeRaise
	case "single_leg_bent_knee_calf_raise":
		return CalfRaiseExerciseNameSingleLegBentKneeCalfRaise
	case "weighted_single_leg_bent_knee_calf_raise":
		return CalfRaiseExerciseNameWeightedSingleLegBentKneeCalfRaise
	case "single_leg_decline_push_up":
		return CalfRaiseExerciseNameSingleLegDeclinePushUp
	case "single_leg_donkey_calf_raise":
		return CalfRaiseExerciseNameSingleLegDonkeyCalfRaise
	case "weighted_single_leg_donkey_calf_raise":
		return CalfRaiseExerciseNameWeightedSingleLegDonkeyCalfRaise
	case "single_leg_hip_raise_with_knee_hold":
		return CalfRaiseExerciseNameSingleLegHipRaiseWithKneeHold
	case "single_leg_standing_calf_raise":
		return CalfRaiseExerciseNameSingleLegStandingCalfRaise
	case "single_leg_standing_dumbbell_calf_raise":
		return CalfRaiseExerciseNameSingleLegStandingDumbbellCalfRaise
	case "standing_barbell_calf_raise":
		return CalfRaiseExerciseNameStandingBarbellCalfRaise
	case "standing_calf_raise":
		return CalfRaiseExerciseNameStandingCalfRaise
	case "weighted_standing_calf_raise":
		return CalfRaiseExerciseNameWeightedStandingCalfRaise
	case "standing_dumbbell_calf_raise":
		return CalfRaiseExerciseNameStandingDumbbellCalfRaise
	default:
		return CalfRaiseExerciseNameInvalid
	}
}

// List returns all constants.
func ListCalfRaiseExerciseName() []CalfRaiseExerciseName {
	return []CalfRaiseExerciseName{
		CalfRaiseExerciseName3WayCalfRaise,
		CalfRaiseExerciseName3WayWeightedCalfRaise,
		CalfRaiseExerciseName3WaySingleLegCalfRaise,
		CalfRaiseExerciseName3WayWeightedSingleLegCalfRaise,
		CalfRaiseExerciseNameDonkeyCalfRaise,
		CalfRaiseExerciseNameWeightedDonkeyCalfRaise,
		CalfRaiseExerciseNameSeatedCalfRaise,
		CalfRaiseExerciseNameWeightedSeatedCalfRaise,
		CalfRaiseExerciseNameSeatedDumbbellToeRaise,
		CalfRaiseExerciseNameSingleLegBentKneeCalfRaise,
		CalfRaiseExerciseNameWeightedSingleLegBentKneeCalfRaise,
		CalfRaiseExerciseNameSingleLegDeclinePushUp,
		CalfRaiseExerciseNameSingleLegDonkeyCalfRaise,
		CalfRaiseExerciseNameWeightedSingleLegDonkeyCalfRaise,
		CalfRaiseExerciseNameSingleLegHipRaiseWithKneeHold,
		CalfRaiseExerciseNameSingleLegStandingCalfRaise,
		CalfRaiseExerciseNameSingleLegStandingDumbbellCalfRaise,
		CalfRaiseExerciseNameStandingBarbellCalfRaise,
		CalfRaiseExerciseNameStandingCalfRaise,
		CalfRaiseExerciseNameWeightedStandingCalfRaise,
		CalfRaiseExerciseNameStandingDumbbellCalfRaise,
	}
}
