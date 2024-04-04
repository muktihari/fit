// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LateralRaiseExerciseName uint16

const (
	LateralRaiseExerciseName45DegreeCableExternalRotation         LateralRaiseExerciseName = 0
	LateralRaiseExerciseNameAlternatingLateralRaiseWithStaticHold LateralRaiseExerciseName = 1
	LateralRaiseExerciseNameBarMuscleUp                           LateralRaiseExerciseName = 2
	LateralRaiseExerciseNameBentOverLateralRaise                  LateralRaiseExerciseName = 3
	LateralRaiseExerciseNameCableDiagonalRaise                    LateralRaiseExerciseName = 4
	LateralRaiseExerciseNameCableFrontRaise                       LateralRaiseExerciseName = 5
	LateralRaiseExerciseNameCalorieRow                            LateralRaiseExerciseName = 6
	LateralRaiseExerciseNameComboShoulderRaise                    LateralRaiseExerciseName = 7
	LateralRaiseExerciseNameDumbbellDiagonalRaise                 LateralRaiseExerciseName = 8
	LateralRaiseExerciseNameDumbbellVRaise                        LateralRaiseExerciseName = 9
	LateralRaiseExerciseNameFrontRaise                            LateralRaiseExerciseName = 10
	LateralRaiseExerciseNameLeaningDumbbellLateralRaise           LateralRaiseExerciseName = 11
	LateralRaiseExerciseNameLyingDumbbellRaise                    LateralRaiseExerciseName = 12
	LateralRaiseExerciseNameMuscleUp                              LateralRaiseExerciseName = 13
	LateralRaiseExerciseNameOneArmCableLateralRaise               LateralRaiseExerciseName = 14
	LateralRaiseExerciseNameOverhandGripRearLateralRaise          LateralRaiseExerciseName = 15
	LateralRaiseExerciseNamePlateRaises                           LateralRaiseExerciseName = 16
	LateralRaiseExerciseNameRingDip                               LateralRaiseExerciseName = 17
	LateralRaiseExerciseNameWeightedRingDip                       LateralRaiseExerciseName = 18
	LateralRaiseExerciseNameRingMuscleUp                          LateralRaiseExerciseName = 19
	LateralRaiseExerciseNameWeightedRingMuscleUp                  LateralRaiseExerciseName = 20
	LateralRaiseExerciseNameRopeClimb                             LateralRaiseExerciseName = 21
	LateralRaiseExerciseNameWeightedRopeClimb                     LateralRaiseExerciseName = 22
	LateralRaiseExerciseNameScaption                              LateralRaiseExerciseName = 23
	LateralRaiseExerciseNameSeatedLateralRaise                    LateralRaiseExerciseName = 24
	LateralRaiseExerciseNameSeatedRearLateralRaise                LateralRaiseExerciseName = 25
	LateralRaiseExerciseNameSideLyingLateralRaise                 LateralRaiseExerciseName = 26
	LateralRaiseExerciseNameStandingLift                          LateralRaiseExerciseName = 27
	LateralRaiseExerciseNameSuspendedRow                          LateralRaiseExerciseName = 28
	LateralRaiseExerciseNameUnderhandGripRearLateralRaise         LateralRaiseExerciseName = 29
	LateralRaiseExerciseNameWallSlide                             LateralRaiseExerciseName = 30
	LateralRaiseExerciseNameWeightedWallSlide                     LateralRaiseExerciseName = 31
	LateralRaiseExerciseNameArmCircles                            LateralRaiseExerciseName = 32
	LateralRaiseExerciseNameShavingTheHead                        LateralRaiseExerciseName = 33
	LateralRaiseExerciseNameInvalid                               LateralRaiseExerciseName = 0xFFFF
)

func (l LateralRaiseExerciseName) Uint16() uint16 { return uint16(l) }

func (l LateralRaiseExerciseName) String() string {
	switch l {
	case LateralRaiseExerciseName45DegreeCableExternalRotation:
		return "45_degree_cable_external_rotation"
	case LateralRaiseExerciseNameAlternatingLateralRaiseWithStaticHold:
		return "alternating_lateral_raise_with_static_hold"
	case LateralRaiseExerciseNameBarMuscleUp:
		return "bar_muscle_up"
	case LateralRaiseExerciseNameBentOverLateralRaise:
		return "bent_over_lateral_raise"
	case LateralRaiseExerciseNameCableDiagonalRaise:
		return "cable_diagonal_raise"
	case LateralRaiseExerciseNameCableFrontRaise:
		return "cable_front_raise"
	case LateralRaiseExerciseNameCalorieRow:
		return "calorie_row"
	case LateralRaiseExerciseNameComboShoulderRaise:
		return "combo_shoulder_raise"
	case LateralRaiseExerciseNameDumbbellDiagonalRaise:
		return "dumbbell_diagonal_raise"
	case LateralRaiseExerciseNameDumbbellVRaise:
		return "dumbbell_v_raise"
	case LateralRaiseExerciseNameFrontRaise:
		return "front_raise"
	case LateralRaiseExerciseNameLeaningDumbbellLateralRaise:
		return "leaning_dumbbell_lateral_raise"
	case LateralRaiseExerciseNameLyingDumbbellRaise:
		return "lying_dumbbell_raise"
	case LateralRaiseExerciseNameMuscleUp:
		return "muscle_up"
	case LateralRaiseExerciseNameOneArmCableLateralRaise:
		return "one_arm_cable_lateral_raise"
	case LateralRaiseExerciseNameOverhandGripRearLateralRaise:
		return "overhand_grip_rear_lateral_raise"
	case LateralRaiseExerciseNamePlateRaises:
		return "plate_raises"
	case LateralRaiseExerciseNameRingDip:
		return "ring_dip"
	case LateralRaiseExerciseNameWeightedRingDip:
		return "weighted_ring_dip"
	case LateralRaiseExerciseNameRingMuscleUp:
		return "ring_muscle_up"
	case LateralRaiseExerciseNameWeightedRingMuscleUp:
		return "weighted_ring_muscle_up"
	case LateralRaiseExerciseNameRopeClimb:
		return "rope_climb"
	case LateralRaiseExerciseNameWeightedRopeClimb:
		return "weighted_rope_climb"
	case LateralRaiseExerciseNameScaption:
		return "scaption"
	case LateralRaiseExerciseNameSeatedLateralRaise:
		return "seated_lateral_raise"
	case LateralRaiseExerciseNameSeatedRearLateralRaise:
		return "seated_rear_lateral_raise"
	case LateralRaiseExerciseNameSideLyingLateralRaise:
		return "side_lying_lateral_raise"
	case LateralRaiseExerciseNameStandingLift:
		return "standing_lift"
	case LateralRaiseExerciseNameSuspendedRow:
		return "suspended_row"
	case LateralRaiseExerciseNameUnderhandGripRearLateralRaise:
		return "underhand_grip_rear_lateral_raise"
	case LateralRaiseExerciseNameWallSlide:
		return "wall_slide"
	case LateralRaiseExerciseNameWeightedWallSlide:
		return "weighted_wall_slide"
	case LateralRaiseExerciseNameArmCircles:
		return "arm_circles"
	case LateralRaiseExerciseNameShavingTheHead:
		return "shaving_the_head"
	default:
		return "LateralRaiseExerciseNameInvalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LateralRaiseExerciseName constant it's represent, return LateralRaiseExerciseNameInvalid if not found.
func LateralRaiseExerciseNameFromString(s string) LateralRaiseExerciseName {
	switch s {
	case "45_degree_cable_external_rotation":
		return LateralRaiseExerciseName45DegreeCableExternalRotation
	case "alternating_lateral_raise_with_static_hold":
		return LateralRaiseExerciseNameAlternatingLateralRaiseWithStaticHold
	case "bar_muscle_up":
		return LateralRaiseExerciseNameBarMuscleUp
	case "bent_over_lateral_raise":
		return LateralRaiseExerciseNameBentOverLateralRaise
	case "cable_diagonal_raise":
		return LateralRaiseExerciseNameCableDiagonalRaise
	case "cable_front_raise":
		return LateralRaiseExerciseNameCableFrontRaise
	case "calorie_row":
		return LateralRaiseExerciseNameCalorieRow
	case "combo_shoulder_raise":
		return LateralRaiseExerciseNameComboShoulderRaise
	case "dumbbell_diagonal_raise":
		return LateralRaiseExerciseNameDumbbellDiagonalRaise
	case "dumbbell_v_raise":
		return LateralRaiseExerciseNameDumbbellVRaise
	case "front_raise":
		return LateralRaiseExerciseNameFrontRaise
	case "leaning_dumbbell_lateral_raise":
		return LateralRaiseExerciseNameLeaningDumbbellLateralRaise
	case "lying_dumbbell_raise":
		return LateralRaiseExerciseNameLyingDumbbellRaise
	case "muscle_up":
		return LateralRaiseExerciseNameMuscleUp
	case "one_arm_cable_lateral_raise":
		return LateralRaiseExerciseNameOneArmCableLateralRaise
	case "overhand_grip_rear_lateral_raise":
		return LateralRaiseExerciseNameOverhandGripRearLateralRaise
	case "plate_raises":
		return LateralRaiseExerciseNamePlateRaises
	case "ring_dip":
		return LateralRaiseExerciseNameRingDip
	case "weighted_ring_dip":
		return LateralRaiseExerciseNameWeightedRingDip
	case "ring_muscle_up":
		return LateralRaiseExerciseNameRingMuscleUp
	case "weighted_ring_muscle_up":
		return LateralRaiseExerciseNameWeightedRingMuscleUp
	case "rope_climb":
		return LateralRaiseExerciseNameRopeClimb
	case "weighted_rope_climb":
		return LateralRaiseExerciseNameWeightedRopeClimb
	case "scaption":
		return LateralRaiseExerciseNameScaption
	case "seated_lateral_raise":
		return LateralRaiseExerciseNameSeatedLateralRaise
	case "seated_rear_lateral_raise":
		return LateralRaiseExerciseNameSeatedRearLateralRaise
	case "side_lying_lateral_raise":
		return LateralRaiseExerciseNameSideLyingLateralRaise
	case "standing_lift":
		return LateralRaiseExerciseNameStandingLift
	case "suspended_row":
		return LateralRaiseExerciseNameSuspendedRow
	case "underhand_grip_rear_lateral_raise":
		return LateralRaiseExerciseNameUnderhandGripRearLateralRaise
	case "wall_slide":
		return LateralRaiseExerciseNameWallSlide
	case "weighted_wall_slide":
		return LateralRaiseExerciseNameWeightedWallSlide
	case "arm_circles":
		return LateralRaiseExerciseNameArmCircles
	case "shaving_the_head":
		return LateralRaiseExerciseNameShavingTheHead
	default:
		return LateralRaiseExerciseNameInvalid
	}
}

// List returns all constants.
func ListLateralRaiseExerciseName() []LateralRaiseExerciseName {
	return []LateralRaiseExerciseName{
		LateralRaiseExerciseName45DegreeCableExternalRotation,
		LateralRaiseExerciseNameAlternatingLateralRaiseWithStaticHold,
		LateralRaiseExerciseNameBarMuscleUp,
		LateralRaiseExerciseNameBentOverLateralRaise,
		LateralRaiseExerciseNameCableDiagonalRaise,
		LateralRaiseExerciseNameCableFrontRaise,
		LateralRaiseExerciseNameCalorieRow,
		LateralRaiseExerciseNameComboShoulderRaise,
		LateralRaiseExerciseNameDumbbellDiagonalRaise,
		LateralRaiseExerciseNameDumbbellVRaise,
		LateralRaiseExerciseNameFrontRaise,
		LateralRaiseExerciseNameLeaningDumbbellLateralRaise,
		LateralRaiseExerciseNameLyingDumbbellRaise,
		LateralRaiseExerciseNameMuscleUp,
		LateralRaiseExerciseNameOneArmCableLateralRaise,
		LateralRaiseExerciseNameOverhandGripRearLateralRaise,
		LateralRaiseExerciseNamePlateRaises,
		LateralRaiseExerciseNameRingDip,
		LateralRaiseExerciseNameWeightedRingDip,
		LateralRaiseExerciseNameRingMuscleUp,
		LateralRaiseExerciseNameWeightedRingMuscleUp,
		LateralRaiseExerciseNameRopeClimb,
		LateralRaiseExerciseNameWeightedRopeClimb,
		LateralRaiseExerciseNameScaption,
		LateralRaiseExerciseNameSeatedLateralRaise,
		LateralRaiseExerciseNameSeatedRearLateralRaise,
		LateralRaiseExerciseNameSideLyingLateralRaise,
		LateralRaiseExerciseNameStandingLift,
		LateralRaiseExerciseNameSuspendedRow,
		LateralRaiseExerciseNameUnderhandGripRearLateralRaise,
		LateralRaiseExerciseNameWallSlide,
		LateralRaiseExerciseNameWeightedWallSlide,
		LateralRaiseExerciseNameArmCircles,
		LateralRaiseExerciseNameShavingTheHead,
	}
}
