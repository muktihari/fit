// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ChopExerciseName uint16

const (
	ChopExerciseNameCablePullThrough                   ChopExerciseName = 0
	ChopExerciseNameCableRotationalLift                ChopExerciseName = 1
	ChopExerciseNameCableWoodchop                      ChopExerciseName = 2
	ChopExerciseNameCrossChopToKnee                    ChopExerciseName = 3
	ChopExerciseNameWeightedCrossChopToKnee            ChopExerciseName = 4
	ChopExerciseNameDumbbellChop                       ChopExerciseName = 5
	ChopExerciseNameHalfKneelingRotation               ChopExerciseName = 6
	ChopExerciseNameWeightedHalfKneelingRotation       ChopExerciseName = 7
	ChopExerciseNameHalfKneelingRotationalChop         ChopExerciseName = 8
	ChopExerciseNameHalfKneelingRotationalReverseChop  ChopExerciseName = 9
	ChopExerciseNameHalfKneelingStabilityChop          ChopExerciseName = 10
	ChopExerciseNameHalfKneelingStabilityReverseChop   ChopExerciseName = 11
	ChopExerciseNameKneelingRotationalChop             ChopExerciseName = 12
	ChopExerciseNameKneelingRotationalReverseChop      ChopExerciseName = 13
	ChopExerciseNameKneelingStabilityChop              ChopExerciseName = 14
	ChopExerciseNameKneelingWoodchopper                ChopExerciseName = 15
	ChopExerciseNameMedicineBallWoodChops              ChopExerciseName = 16
	ChopExerciseNamePowerSquatChops                    ChopExerciseName = 17
	ChopExerciseNameWeightedPowerSquatChops            ChopExerciseName = 18
	ChopExerciseNameStandingRotationalChop             ChopExerciseName = 19
	ChopExerciseNameStandingSplitRotationalChop        ChopExerciseName = 20
	ChopExerciseNameStandingSplitRotationalReverseChop ChopExerciseName = 21
	ChopExerciseNameStandingStabilityReverseChop       ChopExerciseName = 22
	ChopExerciseNameInvalid                            ChopExerciseName = 0xFFFF // INVALID
)

var chopexercisenametostrs = map[ChopExerciseName]string{
	ChopExerciseNameCablePullThrough:                   "cable_pull_through",
	ChopExerciseNameCableRotationalLift:                "cable_rotational_lift",
	ChopExerciseNameCableWoodchop:                      "cable_woodchop",
	ChopExerciseNameCrossChopToKnee:                    "cross_chop_to_knee",
	ChopExerciseNameWeightedCrossChopToKnee:            "weighted_cross_chop_to_knee",
	ChopExerciseNameDumbbellChop:                       "dumbbell_chop",
	ChopExerciseNameHalfKneelingRotation:               "half_kneeling_rotation",
	ChopExerciseNameWeightedHalfKneelingRotation:       "weighted_half_kneeling_rotation",
	ChopExerciseNameHalfKneelingRotationalChop:         "half_kneeling_rotational_chop",
	ChopExerciseNameHalfKneelingRotationalReverseChop:  "half_kneeling_rotational_reverse_chop",
	ChopExerciseNameHalfKneelingStabilityChop:          "half_kneeling_stability_chop",
	ChopExerciseNameHalfKneelingStabilityReverseChop:   "half_kneeling_stability_reverse_chop",
	ChopExerciseNameKneelingRotationalChop:             "kneeling_rotational_chop",
	ChopExerciseNameKneelingRotationalReverseChop:      "kneeling_rotational_reverse_chop",
	ChopExerciseNameKneelingStabilityChop:              "kneeling_stability_chop",
	ChopExerciseNameKneelingWoodchopper:                "kneeling_woodchopper",
	ChopExerciseNameMedicineBallWoodChops:              "medicine_ball_wood_chops",
	ChopExerciseNamePowerSquatChops:                    "power_squat_chops",
	ChopExerciseNameWeightedPowerSquatChops:            "weighted_power_squat_chops",
	ChopExerciseNameStandingRotationalChop:             "standing_rotational_chop",
	ChopExerciseNameStandingSplitRotationalChop:        "standing_split_rotational_chop",
	ChopExerciseNameStandingSplitRotationalReverseChop: "standing_split_rotational_reverse_chop",
	ChopExerciseNameStandingStabilityReverseChop:       "standing_stability_reverse_chop",
	ChopExerciseNameInvalid:                            "invalid",
}

func (c ChopExerciseName) String() string {
	val, ok := chopexercisenametostrs[c]
	if !ok {
		return strconv.FormatUint(uint64(c), 10)
	}
	return val
}

var strtochopexercisename = func() map[string]ChopExerciseName {
	m := make(map[string]ChopExerciseName)
	for t, str := range chopexercisenametostrs {
		m[str] = ChopExerciseName(t)
	}
	return m
}()

// FromString parse string into ChopExerciseName constant it's represent, return ChopExerciseNameInvalid if not found.
func ChopExerciseNameFromString(s string) ChopExerciseName {
	val, ok := strtochopexercisename[s]
	if !ok {
		return strtochopexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListChopExerciseName() []ChopExerciseName {
	vs := make([]ChopExerciseName, 0, len(chopexercisenametostrs))
	for i := range chopexercisenametostrs {
		vs = append(vs, ChopExerciseName(i))
	}
	return vs
}
