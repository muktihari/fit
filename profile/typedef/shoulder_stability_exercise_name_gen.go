// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ShoulderStabilityExerciseName uint16

const (
	ShoulderStabilityExerciseName90DegreeCableExternalRotation          ShoulderStabilityExerciseName = 0
	ShoulderStabilityExerciseNameBandExternalRotation                   ShoulderStabilityExerciseName = 1
	ShoulderStabilityExerciseNameBandInternalRotation                   ShoulderStabilityExerciseName = 2
	ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation ShoulderStabilityExerciseName = 3
	ShoulderStabilityExerciseNameCableExternalRotation                  ShoulderStabilityExerciseName = 4
	ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation   ShoulderStabilityExerciseName = 5
	ShoulderStabilityExerciseNameFloorIRaise                            ShoulderStabilityExerciseName = 6
	ShoulderStabilityExerciseNameWeightedFloorIRaise                    ShoulderStabilityExerciseName = 7
	ShoulderStabilityExerciseNameFloorTRaise                            ShoulderStabilityExerciseName = 8
	ShoulderStabilityExerciseNameWeightedFloorTRaise                    ShoulderStabilityExerciseName = 9
	ShoulderStabilityExerciseNameFloorYRaise                            ShoulderStabilityExerciseName = 10
	ShoulderStabilityExerciseNameWeightedFloorYRaise                    ShoulderStabilityExerciseName = 11
	ShoulderStabilityExerciseNameInclineIRaise                          ShoulderStabilityExerciseName = 12
	ShoulderStabilityExerciseNameWeightedInclineIRaise                  ShoulderStabilityExerciseName = 13
	ShoulderStabilityExerciseNameInclineLRaise                          ShoulderStabilityExerciseName = 14
	ShoulderStabilityExerciseNameWeightedInclineLRaise                  ShoulderStabilityExerciseName = 15
	ShoulderStabilityExerciseNameInclineTRaise                          ShoulderStabilityExerciseName = 16
	ShoulderStabilityExerciseNameWeightedInclineTRaise                  ShoulderStabilityExerciseName = 17
	ShoulderStabilityExerciseNameInclineWRaise                          ShoulderStabilityExerciseName = 18
	ShoulderStabilityExerciseNameWeightedInclineWRaise                  ShoulderStabilityExerciseName = 19
	ShoulderStabilityExerciseNameInclineYRaise                          ShoulderStabilityExerciseName = 20
	ShoulderStabilityExerciseNameWeightedInclineYRaise                  ShoulderStabilityExerciseName = 21
	ShoulderStabilityExerciseNameLyingExternalRotation                  ShoulderStabilityExerciseName = 22
	ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation         ShoulderStabilityExerciseName = 23
	ShoulderStabilityExerciseNameStandingLRaise                         ShoulderStabilityExerciseName = 24
	ShoulderStabilityExerciseNameSwissBallIRaise                        ShoulderStabilityExerciseName = 25
	ShoulderStabilityExerciseNameWeightedSwissBallIRaise                ShoulderStabilityExerciseName = 26
	ShoulderStabilityExerciseNameSwissBallTRaise                        ShoulderStabilityExerciseName = 27
	ShoulderStabilityExerciseNameWeightedSwissBallTRaise                ShoulderStabilityExerciseName = 28
	ShoulderStabilityExerciseNameSwissBallWRaise                        ShoulderStabilityExerciseName = 29
	ShoulderStabilityExerciseNameWeightedSwissBallWRaise                ShoulderStabilityExerciseName = 30
	ShoulderStabilityExerciseNameSwissBallYRaise                        ShoulderStabilityExerciseName = 31
	ShoulderStabilityExerciseNameWeightedSwissBallYRaise                ShoulderStabilityExerciseName = 32
	ShoulderStabilityExerciseNameInvalid                                ShoulderStabilityExerciseName = 0xFFFF // INVALID
)

var shoulderstabilityexercisenametostrs = map[ShoulderStabilityExerciseName]string{
	ShoulderStabilityExerciseName90DegreeCableExternalRotation:          "90_degree_cable_external_rotation",
	ShoulderStabilityExerciseNameBandExternalRotation:                   "band_external_rotation",
	ShoulderStabilityExerciseNameBandInternalRotation:                   "band_internal_rotation",
	ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation: "bent_arm_lateral_raise_and_external_rotation",
	ShoulderStabilityExerciseNameCableExternalRotation:                  "cable_external_rotation",
	ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation:   "dumbbell_face_pull_with_external_rotation",
	ShoulderStabilityExerciseNameFloorIRaise:                            "floor_i_raise",
	ShoulderStabilityExerciseNameWeightedFloorIRaise:                    "weighted_floor_i_raise",
	ShoulderStabilityExerciseNameFloorTRaise:                            "floor_t_raise",
	ShoulderStabilityExerciseNameWeightedFloorTRaise:                    "weighted_floor_t_raise",
	ShoulderStabilityExerciseNameFloorYRaise:                            "floor_y_raise",
	ShoulderStabilityExerciseNameWeightedFloorYRaise:                    "weighted_floor_y_raise",
	ShoulderStabilityExerciseNameInclineIRaise:                          "incline_i_raise",
	ShoulderStabilityExerciseNameWeightedInclineIRaise:                  "weighted_incline_i_raise",
	ShoulderStabilityExerciseNameInclineLRaise:                          "incline_l_raise",
	ShoulderStabilityExerciseNameWeightedInclineLRaise:                  "weighted_incline_l_raise",
	ShoulderStabilityExerciseNameInclineTRaise:                          "incline_t_raise",
	ShoulderStabilityExerciseNameWeightedInclineTRaise:                  "weighted_incline_t_raise",
	ShoulderStabilityExerciseNameInclineWRaise:                          "incline_w_raise",
	ShoulderStabilityExerciseNameWeightedInclineWRaise:                  "weighted_incline_w_raise",
	ShoulderStabilityExerciseNameInclineYRaise:                          "incline_y_raise",
	ShoulderStabilityExerciseNameWeightedInclineYRaise:                  "weighted_incline_y_raise",
	ShoulderStabilityExerciseNameLyingExternalRotation:                  "lying_external_rotation",
	ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation:         "seated_dumbbell_external_rotation",
	ShoulderStabilityExerciseNameStandingLRaise:                         "standing_l_raise",
	ShoulderStabilityExerciseNameSwissBallIRaise:                        "swiss_ball_i_raise",
	ShoulderStabilityExerciseNameWeightedSwissBallIRaise:                "weighted_swiss_ball_i_raise",
	ShoulderStabilityExerciseNameSwissBallTRaise:                        "swiss_ball_t_raise",
	ShoulderStabilityExerciseNameWeightedSwissBallTRaise:                "weighted_swiss_ball_t_raise",
	ShoulderStabilityExerciseNameSwissBallWRaise:                        "swiss_ball_w_raise",
	ShoulderStabilityExerciseNameWeightedSwissBallWRaise:                "weighted_swiss_ball_w_raise",
	ShoulderStabilityExerciseNameSwissBallYRaise:                        "swiss_ball_y_raise",
	ShoulderStabilityExerciseNameWeightedSwissBallYRaise:                "weighted_swiss_ball_y_raise",
	ShoulderStabilityExerciseNameInvalid:                                "invalid",
}

func (s ShoulderStabilityExerciseName) String() string {
	val, ok := shoulderstabilityexercisenametostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtoshoulderstabilityexercisename = func() map[string]ShoulderStabilityExerciseName {
	m := make(map[string]ShoulderStabilityExerciseName)
	for t, str := range shoulderstabilityexercisenametostrs {
		m[str] = ShoulderStabilityExerciseName(t)
	}
	return m
}()

// FromString parse string into ShoulderStabilityExerciseName constant it's represent, return ShoulderStabilityExerciseNameInvalid if not found.
func ShoulderStabilityExerciseNameFromString(s string) ShoulderStabilityExerciseName {
	val, ok := strtoshoulderstabilityexercisename[s]
	if !ok {
		return strtoshoulderstabilityexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListShoulderStabilityExerciseName() []ShoulderStabilityExerciseName {
	vs := make([]ShoulderStabilityExerciseName, 0, len(shoulderstabilityexercisenametostrs))
	for i := range shoulderstabilityexercisenametostrs {
		vs = append(vs, ShoulderStabilityExerciseName(i))
	}
	return vs
}
