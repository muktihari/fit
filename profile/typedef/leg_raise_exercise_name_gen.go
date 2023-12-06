// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LegRaiseExerciseName uint16

const (
	LegRaiseExerciseNameHangingKneeRaise                   LegRaiseExerciseName = 0
	LegRaiseExerciseNameHangingLegRaise                    LegRaiseExerciseName = 1
	LegRaiseExerciseNameWeightedHangingLegRaise            LegRaiseExerciseName = 2
	LegRaiseExerciseNameHangingSingleLegRaise              LegRaiseExerciseName = 3
	LegRaiseExerciseNameWeightedHangingSingleLegRaise      LegRaiseExerciseName = 4
	LegRaiseExerciseNameKettlebellLegRaises                LegRaiseExerciseName = 5
	LegRaiseExerciseNameLegLoweringDrill                   LegRaiseExerciseName = 6
	LegRaiseExerciseNameWeightedLegLoweringDrill           LegRaiseExerciseName = 7
	LegRaiseExerciseNameLyingStraightLegRaise              LegRaiseExerciseName = 8
	LegRaiseExerciseNameWeightedLyingStraightLegRaise      LegRaiseExerciseName = 9
	LegRaiseExerciseNameMedicineBallLegDrops               LegRaiseExerciseName = 10
	LegRaiseExerciseNameQuadrupedLegRaise                  LegRaiseExerciseName = 11
	LegRaiseExerciseNameWeightedQuadrupedLegRaise          LegRaiseExerciseName = 12
	LegRaiseExerciseNameReverseLegRaise                    LegRaiseExerciseName = 13
	LegRaiseExerciseNameWeightedReverseLegRaise            LegRaiseExerciseName = 14
	LegRaiseExerciseNameReverseLegRaiseOnSwissBall         LegRaiseExerciseName = 15
	LegRaiseExerciseNameWeightedReverseLegRaiseOnSwissBall LegRaiseExerciseName = 16
	LegRaiseExerciseNameSingleLegLoweringDrill             LegRaiseExerciseName = 17
	LegRaiseExerciseNameWeightedSingleLegLoweringDrill     LegRaiseExerciseName = 18
	LegRaiseExerciseNameWeightedHangingKneeRaise           LegRaiseExerciseName = 19
	LegRaiseExerciseNameLateralStepover                    LegRaiseExerciseName = 20
	LegRaiseExerciseNameWeightedLateralStepover            LegRaiseExerciseName = 21
	LegRaiseExerciseNameInvalid                            LegRaiseExerciseName = 0xFFFF // INVALID
)

var legraiseexercisenametostrs = map[LegRaiseExerciseName]string{
	LegRaiseExerciseNameHangingKneeRaise:                   "hanging_knee_raise",
	LegRaiseExerciseNameHangingLegRaise:                    "hanging_leg_raise",
	LegRaiseExerciseNameWeightedHangingLegRaise:            "weighted_hanging_leg_raise",
	LegRaiseExerciseNameHangingSingleLegRaise:              "hanging_single_leg_raise",
	LegRaiseExerciseNameWeightedHangingSingleLegRaise:      "weighted_hanging_single_leg_raise",
	LegRaiseExerciseNameKettlebellLegRaises:                "kettlebell_leg_raises",
	LegRaiseExerciseNameLegLoweringDrill:                   "leg_lowering_drill",
	LegRaiseExerciseNameWeightedLegLoweringDrill:           "weighted_leg_lowering_drill",
	LegRaiseExerciseNameLyingStraightLegRaise:              "lying_straight_leg_raise",
	LegRaiseExerciseNameWeightedLyingStraightLegRaise:      "weighted_lying_straight_leg_raise",
	LegRaiseExerciseNameMedicineBallLegDrops:               "medicine_ball_leg_drops",
	LegRaiseExerciseNameQuadrupedLegRaise:                  "quadruped_leg_raise",
	LegRaiseExerciseNameWeightedQuadrupedLegRaise:          "weighted_quadruped_leg_raise",
	LegRaiseExerciseNameReverseLegRaise:                    "reverse_leg_raise",
	LegRaiseExerciseNameWeightedReverseLegRaise:            "weighted_reverse_leg_raise",
	LegRaiseExerciseNameReverseLegRaiseOnSwissBall:         "reverse_leg_raise_on_swiss_ball",
	LegRaiseExerciseNameWeightedReverseLegRaiseOnSwissBall: "weighted_reverse_leg_raise_on_swiss_ball",
	LegRaiseExerciseNameSingleLegLoweringDrill:             "single_leg_lowering_drill",
	LegRaiseExerciseNameWeightedSingleLegLoweringDrill:     "weighted_single_leg_lowering_drill",
	LegRaiseExerciseNameWeightedHangingKneeRaise:           "weighted_hanging_knee_raise",
	LegRaiseExerciseNameLateralStepover:                    "lateral_stepover",
	LegRaiseExerciseNameWeightedLateralStepover:            "weighted_lateral_stepover",
	LegRaiseExerciseNameInvalid:                            "invalid",
}

func (l LegRaiseExerciseName) String() string {
	val, ok := legraiseexercisenametostrs[l]
	if !ok {
		return strconv.FormatUint(uint64(l), 10)
	}
	return val
}

var strtolegraiseexercisename = func() map[string]LegRaiseExerciseName {
	m := make(map[string]LegRaiseExerciseName)
	for t, str := range legraiseexercisenametostrs {
		m[str] = LegRaiseExerciseName(t)
	}
	return m
}()

// FromString parse string into LegRaiseExerciseName constant it's represent, return LegRaiseExerciseNameInvalid if not found.
func LegRaiseExerciseNameFromString(s string) LegRaiseExerciseName {
	val, ok := strtolegraiseexercisename[s]
	if !ok {
		return strtolegraiseexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLegRaiseExerciseName() []LegRaiseExerciseName {
	vs := make([]LegRaiseExerciseName, 0, len(legraiseexercisenametostrs))
	for i := range legraiseexercisenametostrs {
		vs = append(vs, LegRaiseExerciseName(i))
	}
	return vs
}
