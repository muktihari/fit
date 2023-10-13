// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WarmUpExerciseName uint16

const (
	WarmUpExerciseNameQuadrupedRocking            WarmUpExerciseName = 0
	WarmUpExerciseNameNeckTilts                   WarmUpExerciseName = 1
	WarmUpExerciseNameAnkleCircles                WarmUpExerciseName = 2
	WarmUpExerciseNameAnkleDorsiflexionWithBand   WarmUpExerciseName = 3
	WarmUpExerciseNameAnkleInternalRotation       WarmUpExerciseName = 4
	WarmUpExerciseNameArmCircles                  WarmUpExerciseName = 5
	WarmUpExerciseNameBentOverReachToSky          WarmUpExerciseName = 6
	WarmUpExerciseNameCatCamel                    WarmUpExerciseName = 7
	WarmUpExerciseNameElbowToFootLunge            WarmUpExerciseName = 8
	WarmUpExerciseNameForwardAndBackwardLegSwings WarmUpExerciseName = 9
	WarmUpExerciseNameGroiners                    WarmUpExerciseName = 10
	WarmUpExerciseNameInvertedHamstringStretch    WarmUpExerciseName = 11
	WarmUpExerciseNameLateralDuckUnder            WarmUpExerciseName = 12
	WarmUpExerciseNameNeckRotations               WarmUpExerciseName = 13
	WarmUpExerciseNameOppositeArmAndLegBalance    WarmUpExerciseName = 14
	WarmUpExerciseNameReachRollAndLift            WarmUpExerciseName = 15
	WarmUpExerciseNameScorpion                    WarmUpExerciseName = 16 // Deprecated do not use
	WarmUpExerciseNameShoulderCircles             WarmUpExerciseName = 17
	WarmUpExerciseNameSideToSideLegSwings         WarmUpExerciseName = 18
	WarmUpExerciseNameSleeperStretch              WarmUpExerciseName = 19
	WarmUpExerciseNameSlideOut                    WarmUpExerciseName = 20
	WarmUpExerciseNameSwissBallHipCrossover       WarmUpExerciseName = 21
	WarmUpExerciseNameSwissBallReachRollAndLift   WarmUpExerciseName = 22
	WarmUpExerciseNameSwissBallWindshieldWipers   WarmUpExerciseName = 23
	WarmUpExerciseNameThoracicRotation            WarmUpExerciseName = 24
	WarmUpExerciseNameWalkingHighKicks            WarmUpExerciseName = 25
	WarmUpExerciseNameWalkingHighKnees            WarmUpExerciseName = 26
	WarmUpExerciseNameWalkingKneeHugs             WarmUpExerciseName = 27
	WarmUpExerciseNameWalkingLegCradles           WarmUpExerciseName = 28
	WarmUpExerciseNameWalkout                     WarmUpExerciseName = 29
	WarmUpExerciseNameWalkoutFromPushUpPosition   WarmUpExerciseName = 30
	WarmUpExerciseNameInvalid                     WarmUpExerciseName = 0xFFFF // INVALID
)

var warmupexercisenametostrs = map[WarmUpExerciseName]string{
	WarmUpExerciseNameQuadrupedRocking:            "quadruped_rocking",
	WarmUpExerciseNameNeckTilts:                   "neck_tilts",
	WarmUpExerciseNameAnkleCircles:                "ankle_circles",
	WarmUpExerciseNameAnkleDorsiflexionWithBand:   "ankle_dorsiflexion_with_band",
	WarmUpExerciseNameAnkleInternalRotation:       "ankle_internal_rotation",
	WarmUpExerciseNameArmCircles:                  "arm_circles",
	WarmUpExerciseNameBentOverReachToSky:          "bent_over_reach_to_sky",
	WarmUpExerciseNameCatCamel:                    "cat_camel",
	WarmUpExerciseNameElbowToFootLunge:            "elbow_to_foot_lunge",
	WarmUpExerciseNameForwardAndBackwardLegSwings: "forward_and_backward_leg_swings",
	WarmUpExerciseNameGroiners:                    "groiners",
	WarmUpExerciseNameInvertedHamstringStretch:    "inverted_hamstring_stretch",
	WarmUpExerciseNameLateralDuckUnder:            "lateral_duck_under",
	WarmUpExerciseNameNeckRotations:               "neck_rotations",
	WarmUpExerciseNameOppositeArmAndLegBalance:    "opposite_arm_and_leg_balance",
	WarmUpExerciseNameReachRollAndLift:            "reach_roll_and_lift",
	WarmUpExerciseNameScorpion:                    "scorpion",
	WarmUpExerciseNameShoulderCircles:             "shoulder_circles",
	WarmUpExerciseNameSideToSideLegSwings:         "side_to_side_leg_swings",
	WarmUpExerciseNameSleeperStretch:              "sleeper_stretch",
	WarmUpExerciseNameSlideOut:                    "slide_out",
	WarmUpExerciseNameSwissBallHipCrossover:       "swiss_ball_hip_crossover",
	WarmUpExerciseNameSwissBallReachRollAndLift:   "swiss_ball_reach_roll_and_lift",
	WarmUpExerciseNameSwissBallWindshieldWipers:   "swiss_ball_windshield_wipers",
	WarmUpExerciseNameThoracicRotation:            "thoracic_rotation",
	WarmUpExerciseNameWalkingHighKicks:            "walking_high_kicks",
	WarmUpExerciseNameWalkingHighKnees:            "walking_high_knees",
	WarmUpExerciseNameWalkingKneeHugs:             "walking_knee_hugs",
	WarmUpExerciseNameWalkingLegCradles:           "walking_leg_cradles",
	WarmUpExerciseNameWalkout:                     "walkout",
	WarmUpExerciseNameWalkoutFromPushUpPosition:   "walkout_from_push_up_position",
	WarmUpExerciseNameInvalid:                     "invalid",
}

func (w WarmUpExerciseName) String() string {
	val, ok := warmupexercisenametostrs[w]
	if !ok {
		return strconv.FormatUint(uint64(w), 10)
	}
	return val
}

var strtowarmupexercisename = func() map[string]WarmUpExerciseName {
	m := make(map[string]WarmUpExerciseName)
	for t, str := range warmupexercisenametostrs {
		m[str] = WarmUpExerciseName(t)
	}
	return m
}()

// FromString parse string into WarmUpExerciseName constant it's represent, return WarmUpExerciseNameInvalid if not found.
func WarmUpExerciseNameFromString(s string) WarmUpExerciseName {
	val, ok := strtowarmupexercisename[s]
	if !ok {
		return strtowarmupexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWarmUpExerciseName() []WarmUpExerciseName {
	vs := make([]WarmUpExerciseName, 0, len(warmupexercisenametostrs))
	for i := range warmupexercisenametostrs {
		vs = append(vs, WarmUpExerciseName(i))
	}
	return vs
}