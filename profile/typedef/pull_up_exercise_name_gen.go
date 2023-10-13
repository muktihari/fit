// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type PullUpExerciseName uint16

const (
	PullUpExerciseNameBandedPullUps                    PullUpExerciseName = 0
	PullUpExerciseName30DegreeLatPulldown              PullUpExerciseName = 1
	PullUpExerciseNameBandAssistedChinUp               PullUpExerciseName = 2
	PullUpExerciseNameCloseGripChinUp                  PullUpExerciseName = 3
	PullUpExerciseNameWeightedCloseGripChinUp          PullUpExerciseName = 4
	PullUpExerciseNameCloseGripLatPulldown             PullUpExerciseName = 5
	PullUpExerciseNameCrossoverChinUp                  PullUpExerciseName = 6
	PullUpExerciseNameWeightedCrossoverChinUp          PullUpExerciseName = 7
	PullUpExerciseNameEzBarPullover                    PullUpExerciseName = 8
	PullUpExerciseNameHangingHurdle                    PullUpExerciseName = 9
	PullUpExerciseNameWeightedHangingHurdle            PullUpExerciseName = 10
	PullUpExerciseNameKneelingLatPulldown              PullUpExerciseName = 11
	PullUpExerciseNameKneelingUnderhandGripLatPulldown PullUpExerciseName = 12
	PullUpExerciseNameLatPulldown                      PullUpExerciseName = 13
	PullUpExerciseNameMixedGripChinUp                  PullUpExerciseName = 14
	PullUpExerciseNameWeightedMixedGripChinUp          PullUpExerciseName = 15
	PullUpExerciseNameMixedGripPullUp                  PullUpExerciseName = 16
	PullUpExerciseNameWeightedMixedGripPullUp          PullUpExerciseName = 17
	PullUpExerciseNameReverseGripPulldown              PullUpExerciseName = 18
	PullUpExerciseNameStandingCablePullover            PullUpExerciseName = 19
	PullUpExerciseNameStraightArmPulldown              PullUpExerciseName = 20
	PullUpExerciseNameSwissBallEzBarPullover           PullUpExerciseName = 21
	PullUpExerciseNameTowelPullUp                      PullUpExerciseName = 22
	PullUpExerciseNameWeightedTowelPullUp              PullUpExerciseName = 23
	PullUpExerciseNameWeightedPullUp                   PullUpExerciseName = 24
	PullUpExerciseNameWideGripLatPulldown              PullUpExerciseName = 25
	PullUpExerciseNameWideGripPullUp                   PullUpExerciseName = 26
	PullUpExerciseNameWeightedWideGripPullUp           PullUpExerciseName = 27
	PullUpExerciseNameBurpeePullUp                     PullUpExerciseName = 28
	PullUpExerciseNameWeightedBurpeePullUp             PullUpExerciseName = 29
	PullUpExerciseNameJumpingPullUps                   PullUpExerciseName = 30
	PullUpExerciseNameWeightedJumpingPullUps           PullUpExerciseName = 31
	PullUpExerciseNameKippingPullUp                    PullUpExerciseName = 32
	PullUpExerciseNameWeightedKippingPullUp            PullUpExerciseName = 33
	PullUpExerciseNameLPullUp                          PullUpExerciseName = 34
	PullUpExerciseNameWeightedLPullUp                  PullUpExerciseName = 35
	PullUpExerciseNameSuspendedChinUp                  PullUpExerciseName = 36
	PullUpExerciseNameWeightedSuspendedChinUp          PullUpExerciseName = 37
	PullUpExerciseNamePullUp                           PullUpExerciseName = 38
	PullUpExerciseNameInvalid                          PullUpExerciseName = 0xFFFF // INVALID
)

var pullupexercisenametostrs = map[PullUpExerciseName]string{
	PullUpExerciseNameBandedPullUps:                    "banded_pull_ups",
	PullUpExerciseName30DegreeLatPulldown:              "30_degree_lat_pulldown",
	PullUpExerciseNameBandAssistedChinUp:               "band_assisted_chin_up",
	PullUpExerciseNameCloseGripChinUp:                  "close_grip_chin_up",
	PullUpExerciseNameWeightedCloseGripChinUp:          "weighted_close_grip_chin_up",
	PullUpExerciseNameCloseGripLatPulldown:             "close_grip_lat_pulldown",
	PullUpExerciseNameCrossoverChinUp:                  "crossover_chin_up",
	PullUpExerciseNameWeightedCrossoverChinUp:          "weighted_crossover_chin_up",
	PullUpExerciseNameEzBarPullover:                    "ez_bar_pullover",
	PullUpExerciseNameHangingHurdle:                    "hanging_hurdle",
	PullUpExerciseNameWeightedHangingHurdle:            "weighted_hanging_hurdle",
	PullUpExerciseNameKneelingLatPulldown:              "kneeling_lat_pulldown",
	PullUpExerciseNameKneelingUnderhandGripLatPulldown: "kneeling_underhand_grip_lat_pulldown",
	PullUpExerciseNameLatPulldown:                      "lat_pulldown",
	PullUpExerciseNameMixedGripChinUp:                  "mixed_grip_chin_up",
	PullUpExerciseNameWeightedMixedGripChinUp:          "weighted_mixed_grip_chin_up",
	PullUpExerciseNameMixedGripPullUp:                  "mixed_grip_pull_up",
	PullUpExerciseNameWeightedMixedGripPullUp:          "weighted_mixed_grip_pull_up",
	PullUpExerciseNameReverseGripPulldown:              "reverse_grip_pulldown",
	PullUpExerciseNameStandingCablePullover:            "standing_cable_pullover",
	PullUpExerciseNameStraightArmPulldown:              "straight_arm_pulldown",
	PullUpExerciseNameSwissBallEzBarPullover:           "swiss_ball_ez_bar_pullover",
	PullUpExerciseNameTowelPullUp:                      "towel_pull_up",
	PullUpExerciseNameWeightedTowelPullUp:              "weighted_towel_pull_up",
	PullUpExerciseNameWeightedPullUp:                   "weighted_pull_up",
	PullUpExerciseNameWideGripLatPulldown:              "wide_grip_lat_pulldown",
	PullUpExerciseNameWideGripPullUp:                   "wide_grip_pull_up",
	PullUpExerciseNameWeightedWideGripPullUp:           "weighted_wide_grip_pull_up",
	PullUpExerciseNameBurpeePullUp:                     "burpee_pull_up",
	PullUpExerciseNameWeightedBurpeePullUp:             "weighted_burpee_pull_up",
	PullUpExerciseNameJumpingPullUps:                   "jumping_pull_ups",
	PullUpExerciseNameWeightedJumpingPullUps:           "weighted_jumping_pull_ups",
	PullUpExerciseNameKippingPullUp:                    "kipping_pull_up",
	PullUpExerciseNameWeightedKippingPullUp:            "weighted_kipping_pull_up",
	PullUpExerciseNameLPullUp:                          "l_pull_up",
	PullUpExerciseNameWeightedLPullUp:                  "weighted_l_pull_up",
	PullUpExerciseNameSuspendedChinUp:                  "suspended_chin_up",
	PullUpExerciseNameWeightedSuspendedChinUp:          "weighted_suspended_chin_up",
	PullUpExerciseNamePullUp:                           "pull_up",
	PullUpExerciseNameInvalid:                          "invalid",
}

func (p PullUpExerciseName) String() string {
	val, ok := pullupexercisenametostrs[p]
	if !ok {
		return strconv.FormatUint(uint64(p), 10)
	}
	return val
}

var strtopullupexercisename = func() map[string]PullUpExerciseName {
	m := make(map[string]PullUpExerciseName)
	for t, str := range pullupexercisenametostrs {
		m[str] = PullUpExerciseName(t)
	}
	return m
}()

// FromString parse string into PullUpExerciseName constant it's represent, return PullUpExerciseNameInvalid if not found.
func PullUpExerciseNameFromString(s string) PullUpExerciseName {
	val, ok := strtopullupexercisename[s]
	if !ok {
		return strtopullupexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListPullUpExerciseName() []PullUpExerciseName {
	vs := make([]PullUpExerciseName, 0, len(pullupexercisenametostrs))
	for i := range pullupexercisenametostrs {
		vs = append(vs, PullUpExerciseName(i))
	}
	return vs
}