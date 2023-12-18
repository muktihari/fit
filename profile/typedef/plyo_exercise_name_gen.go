// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type PlyoExerciseName uint16

const (
	PlyoExerciseNameAlternatingJumpLunge                  PlyoExerciseName = 0
	PlyoExerciseNameWeightedAlternatingJumpLunge          PlyoExerciseName = 1
	PlyoExerciseNameBarbellJumpSquat                      PlyoExerciseName = 2
	PlyoExerciseNameBodyWeightJumpSquat                   PlyoExerciseName = 3
	PlyoExerciseNameWeightedJumpSquat                     PlyoExerciseName = 4
	PlyoExerciseNameCrossKneeStrike                       PlyoExerciseName = 5
	PlyoExerciseNameWeightedCrossKneeStrike               PlyoExerciseName = 6
	PlyoExerciseNameDepthJump                             PlyoExerciseName = 7
	PlyoExerciseNameWeightedDepthJump                     PlyoExerciseName = 8
	PlyoExerciseNameDumbbellJumpSquat                     PlyoExerciseName = 9
	PlyoExerciseNameDumbbellSplitJump                     PlyoExerciseName = 10
	PlyoExerciseNameFrontKneeStrike                       PlyoExerciseName = 11
	PlyoExerciseNameWeightedFrontKneeStrike               PlyoExerciseName = 12
	PlyoExerciseNameHighBoxJump                           PlyoExerciseName = 13
	PlyoExerciseNameWeightedHighBoxJump                   PlyoExerciseName = 14
	PlyoExerciseNameIsometricExplosiveBodyWeightJumpSquat PlyoExerciseName = 15
	PlyoExerciseNameWeightedIsometricExplosiveJumpSquat   PlyoExerciseName = 16
	PlyoExerciseNameLateralLeapAndHop                     PlyoExerciseName = 17
	PlyoExerciseNameWeightedLateralLeapAndHop             PlyoExerciseName = 18
	PlyoExerciseNameLateralPlyoSquats                     PlyoExerciseName = 19
	PlyoExerciseNameWeightedLateralPlyoSquats             PlyoExerciseName = 20
	PlyoExerciseNameLateralSlide                          PlyoExerciseName = 21
	PlyoExerciseNameWeightedLateralSlide                  PlyoExerciseName = 22
	PlyoExerciseNameMedicineBallOverheadThrows            PlyoExerciseName = 23
	PlyoExerciseNameMedicineBallSideThrow                 PlyoExerciseName = 24
	PlyoExerciseNameMedicineBallSlam                      PlyoExerciseName = 25
	PlyoExerciseNameSideToSideMedicineBallThrows          PlyoExerciseName = 26
	PlyoExerciseNameSideToSideShuffleJump                 PlyoExerciseName = 27
	PlyoExerciseNameWeightedSideToSideShuffleJump         PlyoExerciseName = 28
	PlyoExerciseNameSquatJumpOntoBox                      PlyoExerciseName = 29
	PlyoExerciseNameWeightedSquatJumpOntoBox              PlyoExerciseName = 30
	PlyoExerciseNameSquatJumpsInAndOut                    PlyoExerciseName = 31
	PlyoExerciseNameWeightedSquatJumpsInAndOut            PlyoExerciseName = 32
	PlyoExerciseNameInvalid                               PlyoExerciseName = 0xFFFF
)

func (p PlyoExerciseName) String() string {
	switch p {
	case PlyoExerciseNameAlternatingJumpLunge:
		return "alternating_jump_lunge"
	case PlyoExerciseNameWeightedAlternatingJumpLunge:
		return "weighted_alternating_jump_lunge"
	case PlyoExerciseNameBarbellJumpSquat:
		return "barbell_jump_squat"
	case PlyoExerciseNameBodyWeightJumpSquat:
		return "body_weight_jump_squat"
	case PlyoExerciseNameWeightedJumpSquat:
		return "weighted_jump_squat"
	case PlyoExerciseNameCrossKneeStrike:
		return "cross_knee_strike"
	case PlyoExerciseNameWeightedCrossKneeStrike:
		return "weighted_cross_knee_strike"
	case PlyoExerciseNameDepthJump:
		return "depth_jump"
	case PlyoExerciseNameWeightedDepthJump:
		return "weighted_depth_jump"
	case PlyoExerciseNameDumbbellJumpSquat:
		return "dumbbell_jump_squat"
	case PlyoExerciseNameDumbbellSplitJump:
		return "dumbbell_split_jump"
	case PlyoExerciseNameFrontKneeStrike:
		return "front_knee_strike"
	case PlyoExerciseNameWeightedFrontKneeStrike:
		return "weighted_front_knee_strike"
	case PlyoExerciseNameHighBoxJump:
		return "high_box_jump"
	case PlyoExerciseNameWeightedHighBoxJump:
		return "weighted_high_box_jump"
	case PlyoExerciseNameIsometricExplosiveBodyWeightJumpSquat:
		return "isometric_explosive_body_weight_jump_squat"
	case PlyoExerciseNameWeightedIsometricExplosiveJumpSquat:
		return "weighted_isometric_explosive_jump_squat"
	case PlyoExerciseNameLateralLeapAndHop:
		return "lateral_leap_and_hop"
	case PlyoExerciseNameWeightedLateralLeapAndHop:
		return "weighted_lateral_leap_and_hop"
	case PlyoExerciseNameLateralPlyoSquats:
		return "lateral_plyo_squats"
	case PlyoExerciseNameWeightedLateralPlyoSquats:
		return "weighted_lateral_plyo_squats"
	case PlyoExerciseNameLateralSlide:
		return "lateral_slide"
	case PlyoExerciseNameWeightedLateralSlide:
		return "weighted_lateral_slide"
	case PlyoExerciseNameMedicineBallOverheadThrows:
		return "medicine_ball_overhead_throws"
	case PlyoExerciseNameMedicineBallSideThrow:
		return "medicine_ball_side_throw"
	case PlyoExerciseNameMedicineBallSlam:
		return "medicine_ball_slam"
	case PlyoExerciseNameSideToSideMedicineBallThrows:
		return "side_to_side_medicine_ball_throws"
	case PlyoExerciseNameSideToSideShuffleJump:
		return "side_to_side_shuffle_jump"
	case PlyoExerciseNameWeightedSideToSideShuffleJump:
		return "weighted_side_to_side_shuffle_jump"
	case PlyoExerciseNameSquatJumpOntoBox:
		return "squat_jump_onto_box"
	case PlyoExerciseNameWeightedSquatJumpOntoBox:
		return "weighted_squat_jump_onto_box"
	case PlyoExerciseNameSquatJumpsInAndOut:
		return "squat_jumps_in_and_out"
	case PlyoExerciseNameWeightedSquatJumpsInAndOut:
		return "weighted_squat_jumps_in_and_out"
	default:
		return "PlyoExerciseNameInvalid(" + strconv.FormatUint(uint64(p), 10) + ")"
	}
}

// FromString parse string into PlyoExerciseName constant it's represent, return PlyoExerciseNameInvalid if not found.
func PlyoExerciseNameFromString(s string) PlyoExerciseName {
	switch s {
	case "alternating_jump_lunge":
		return PlyoExerciseNameAlternatingJumpLunge
	case "weighted_alternating_jump_lunge":
		return PlyoExerciseNameWeightedAlternatingJumpLunge
	case "barbell_jump_squat":
		return PlyoExerciseNameBarbellJumpSquat
	case "body_weight_jump_squat":
		return PlyoExerciseNameBodyWeightJumpSquat
	case "weighted_jump_squat":
		return PlyoExerciseNameWeightedJumpSquat
	case "cross_knee_strike":
		return PlyoExerciseNameCrossKneeStrike
	case "weighted_cross_knee_strike":
		return PlyoExerciseNameWeightedCrossKneeStrike
	case "depth_jump":
		return PlyoExerciseNameDepthJump
	case "weighted_depth_jump":
		return PlyoExerciseNameWeightedDepthJump
	case "dumbbell_jump_squat":
		return PlyoExerciseNameDumbbellJumpSquat
	case "dumbbell_split_jump":
		return PlyoExerciseNameDumbbellSplitJump
	case "front_knee_strike":
		return PlyoExerciseNameFrontKneeStrike
	case "weighted_front_knee_strike":
		return PlyoExerciseNameWeightedFrontKneeStrike
	case "high_box_jump":
		return PlyoExerciseNameHighBoxJump
	case "weighted_high_box_jump":
		return PlyoExerciseNameWeightedHighBoxJump
	case "isometric_explosive_body_weight_jump_squat":
		return PlyoExerciseNameIsometricExplosiveBodyWeightJumpSquat
	case "weighted_isometric_explosive_jump_squat":
		return PlyoExerciseNameWeightedIsometricExplosiveJumpSquat
	case "lateral_leap_and_hop":
		return PlyoExerciseNameLateralLeapAndHop
	case "weighted_lateral_leap_and_hop":
		return PlyoExerciseNameWeightedLateralLeapAndHop
	case "lateral_plyo_squats":
		return PlyoExerciseNameLateralPlyoSquats
	case "weighted_lateral_plyo_squats":
		return PlyoExerciseNameWeightedLateralPlyoSquats
	case "lateral_slide":
		return PlyoExerciseNameLateralSlide
	case "weighted_lateral_slide":
		return PlyoExerciseNameWeightedLateralSlide
	case "medicine_ball_overhead_throws":
		return PlyoExerciseNameMedicineBallOverheadThrows
	case "medicine_ball_side_throw":
		return PlyoExerciseNameMedicineBallSideThrow
	case "medicine_ball_slam":
		return PlyoExerciseNameMedicineBallSlam
	case "side_to_side_medicine_ball_throws":
		return PlyoExerciseNameSideToSideMedicineBallThrows
	case "side_to_side_shuffle_jump":
		return PlyoExerciseNameSideToSideShuffleJump
	case "weighted_side_to_side_shuffle_jump":
		return PlyoExerciseNameWeightedSideToSideShuffleJump
	case "squat_jump_onto_box":
		return PlyoExerciseNameSquatJumpOntoBox
	case "weighted_squat_jump_onto_box":
		return PlyoExerciseNameWeightedSquatJumpOntoBox
	case "squat_jumps_in_and_out":
		return PlyoExerciseNameSquatJumpsInAndOut
	case "weighted_squat_jumps_in_and_out":
		return PlyoExerciseNameWeightedSquatJumpsInAndOut
	default:
		return PlyoExerciseNameInvalid
	}
}

// List returns all constants.
func ListPlyoExerciseName() []PlyoExerciseName {
	return []PlyoExerciseName{
		PlyoExerciseNameAlternatingJumpLunge,
		PlyoExerciseNameWeightedAlternatingJumpLunge,
		PlyoExerciseNameBarbellJumpSquat,
		PlyoExerciseNameBodyWeightJumpSquat,
		PlyoExerciseNameWeightedJumpSquat,
		PlyoExerciseNameCrossKneeStrike,
		PlyoExerciseNameWeightedCrossKneeStrike,
		PlyoExerciseNameDepthJump,
		PlyoExerciseNameWeightedDepthJump,
		PlyoExerciseNameDumbbellJumpSquat,
		PlyoExerciseNameDumbbellSplitJump,
		PlyoExerciseNameFrontKneeStrike,
		PlyoExerciseNameWeightedFrontKneeStrike,
		PlyoExerciseNameHighBoxJump,
		PlyoExerciseNameWeightedHighBoxJump,
		PlyoExerciseNameIsometricExplosiveBodyWeightJumpSquat,
		PlyoExerciseNameWeightedIsometricExplosiveJumpSquat,
		PlyoExerciseNameLateralLeapAndHop,
		PlyoExerciseNameWeightedLateralLeapAndHop,
		PlyoExerciseNameLateralPlyoSquats,
		PlyoExerciseNameWeightedLateralPlyoSquats,
		PlyoExerciseNameLateralSlide,
		PlyoExerciseNameWeightedLateralSlide,
		PlyoExerciseNameMedicineBallOverheadThrows,
		PlyoExerciseNameMedicineBallSideThrow,
		PlyoExerciseNameMedicineBallSlam,
		PlyoExerciseNameSideToSideMedicineBallThrows,
		PlyoExerciseNameSideToSideShuffleJump,
		PlyoExerciseNameWeightedSideToSideShuffleJump,
		PlyoExerciseNameSquatJumpOntoBox,
		PlyoExerciseNameWeightedSquatJumpOntoBox,
		PlyoExerciseNameSquatJumpsInAndOut,
		PlyoExerciseNameWeightedSquatJumpsInAndOut,
	}
}
