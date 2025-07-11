// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DeadliftExerciseName uint16

const (
	DeadliftExerciseNameBarbellDeadlift                       DeadliftExerciseName = 0
	DeadliftExerciseNameBarbellStraightLegDeadlift            DeadliftExerciseName = 1
	DeadliftExerciseNameDumbbellDeadlift                      DeadliftExerciseName = 2
	DeadliftExerciseNameDumbbellSingleLegDeadliftToRow        DeadliftExerciseName = 3
	DeadliftExerciseNameDumbbellStraightLegDeadlift           DeadliftExerciseName = 4
	DeadliftExerciseNameKettlebellFloorToShelf                DeadliftExerciseName = 5
	DeadliftExerciseNameOneArmOneLegDeadlift                  DeadliftExerciseName = 6
	DeadliftExerciseNameRackPull                              DeadliftExerciseName = 7
	DeadliftExerciseNameRotationalDumbbellStraightLegDeadlift DeadliftExerciseName = 8
	DeadliftExerciseNameSingleArmDeadlift                     DeadliftExerciseName = 9
	DeadliftExerciseNameSingleLegBarbellDeadlift              DeadliftExerciseName = 10
	DeadliftExerciseNameSingleLegBarbellStraightLegDeadlift   DeadliftExerciseName = 11
	DeadliftExerciseNameSingleLegDeadliftWithBarbell          DeadliftExerciseName = 12
	DeadliftExerciseNameSingleLegRdlCircuit                   DeadliftExerciseName = 13
	DeadliftExerciseNameSingleLegRomanianDeadliftWithDumbbell DeadliftExerciseName = 14
	DeadliftExerciseNameSumoDeadlift                          DeadliftExerciseName = 15
	DeadliftExerciseNameSumoDeadliftHighPull                  DeadliftExerciseName = 16
	DeadliftExerciseNameTrapBarDeadlift                       DeadliftExerciseName = 17
	DeadliftExerciseNameWideGripBarbellDeadlift               DeadliftExerciseName = 18
	DeadliftExerciseNameKettlebellDeadlift                    DeadliftExerciseName = 20
	DeadliftExerciseNameKettlebellSumoDeadlift                DeadliftExerciseName = 21
	DeadliftExerciseNameRomanianDeadlift                      DeadliftExerciseName = 23
	DeadliftExerciseNameSingleLegRomanianDeadliftCircuit      DeadliftExerciseName = 24
	DeadliftExerciseNameStraightLegDeadlift                   DeadliftExerciseName = 25
	DeadliftExerciseNameInvalid                               DeadliftExerciseName = 0xFFFF
)

func (d DeadliftExerciseName) Uint16() uint16 { return uint16(d) }

func (d DeadliftExerciseName) String() string {
	switch d {
	case DeadliftExerciseNameBarbellDeadlift:
		return "barbell_deadlift"
	case DeadliftExerciseNameBarbellStraightLegDeadlift:
		return "barbell_straight_leg_deadlift"
	case DeadliftExerciseNameDumbbellDeadlift:
		return "dumbbell_deadlift"
	case DeadliftExerciseNameDumbbellSingleLegDeadliftToRow:
		return "dumbbell_single_leg_deadlift_to_row"
	case DeadliftExerciseNameDumbbellStraightLegDeadlift:
		return "dumbbell_straight_leg_deadlift"
	case DeadliftExerciseNameKettlebellFloorToShelf:
		return "kettlebell_floor_to_shelf"
	case DeadliftExerciseNameOneArmOneLegDeadlift:
		return "one_arm_one_leg_deadlift"
	case DeadliftExerciseNameRackPull:
		return "rack_pull"
	case DeadliftExerciseNameRotationalDumbbellStraightLegDeadlift:
		return "rotational_dumbbell_straight_leg_deadlift"
	case DeadliftExerciseNameSingleArmDeadlift:
		return "single_arm_deadlift"
	case DeadliftExerciseNameSingleLegBarbellDeadlift:
		return "single_leg_barbell_deadlift"
	case DeadliftExerciseNameSingleLegBarbellStraightLegDeadlift:
		return "single_leg_barbell_straight_leg_deadlift"
	case DeadliftExerciseNameSingleLegDeadliftWithBarbell:
		return "single_leg_deadlift_with_barbell"
	case DeadliftExerciseNameSingleLegRdlCircuit:
		return "single_leg_rdl_circuit"
	case DeadliftExerciseNameSingleLegRomanianDeadliftWithDumbbell:
		return "single_leg_romanian_deadlift_with_dumbbell"
	case DeadliftExerciseNameSumoDeadlift:
		return "sumo_deadlift"
	case DeadliftExerciseNameSumoDeadliftHighPull:
		return "sumo_deadlift_high_pull"
	case DeadliftExerciseNameTrapBarDeadlift:
		return "trap_bar_deadlift"
	case DeadliftExerciseNameWideGripBarbellDeadlift:
		return "wide_grip_barbell_deadlift"
	case DeadliftExerciseNameKettlebellDeadlift:
		return "kettlebell_deadlift"
	case DeadliftExerciseNameKettlebellSumoDeadlift:
		return "kettlebell_sumo_deadlift"
	case DeadliftExerciseNameRomanianDeadlift:
		return "romanian_deadlift"
	case DeadliftExerciseNameSingleLegRomanianDeadliftCircuit:
		return "single_leg_romanian_deadlift_circuit"
	case DeadliftExerciseNameStraightLegDeadlift:
		return "straight_leg_deadlift"
	default:
		return "DeadliftExerciseNameInvalid(" + strconv.FormatUint(uint64(d), 10) + ")"
	}
}

// FromString parse string into DeadliftExerciseName constant it's represent, return DeadliftExerciseNameInvalid if not found.
func DeadliftExerciseNameFromString(s string) DeadliftExerciseName {
	switch s {
	case "barbell_deadlift":
		return DeadliftExerciseNameBarbellDeadlift
	case "barbell_straight_leg_deadlift":
		return DeadliftExerciseNameBarbellStraightLegDeadlift
	case "dumbbell_deadlift":
		return DeadliftExerciseNameDumbbellDeadlift
	case "dumbbell_single_leg_deadlift_to_row":
		return DeadliftExerciseNameDumbbellSingleLegDeadliftToRow
	case "dumbbell_straight_leg_deadlift":
		return DeadliftExerciseNameDumbbellStraightLegDeadlift
	case "kettlebell_floor_to_shelf":
		return DeadliftExerciseNameKettlebellFloorToShelf
	case "one_arm_one_leg_deadlift":
		return DeadliftExerciseNameOneArmOneLegDeadlift
	case "rack_pull":
		return DeadliftExerciseNameRackPull
	case "rotational_dumbbell_straight_leg_deadlift":
		return DeadliftExerciseNameRotationalDumbbellStraightLegDeadlift
	case "single_arm_deadlift":
		return DeadliftExerciseNameSingleArmDeadlift
	case "single_leg_barbell_deadlift":
		return DeadliftExerciseNameSingleLegBarbellDeadlift
	case "single_leg_barbell_straight_leg_deadlift":
		return DeadliftExerciseNameSingleLegBarbellStraightLegDeadlift
	case "single_leg_deadlift_with_barbell":
		return DeadliftExerciseNameSingleLegDeadliftWithBarbell
	case "single_leg_rdl_circuit":
		return DeadliftExerciseNameSingleLegRdlCircuit
	case "single_leg_romanian_deadlift_with_dumbbell":
		return DeadliftExerciseNameSingleLegRomanianDeadliftWithDumbbell
	case "sumo_deadlift":
		return DeadliftExerciseNameSumoDeadlift
	case "sumo_deadlift_high_pull":
		return DeadliftExerciseNameSumoDeadliftHighPull
	case "trap_bar_deadlift":
		return DeadliftExerciseNameTrapBarDeadlift
	case "wide_grip_barbell_deadlift":
		return DeadliftExerciseNameWideGripBarbellDeadlift
	case "kettlebell_deadlift":
		return DeadliftExerciseNameKettlebellDeadlift
	case "kettlebell_sumo_deadlift":
		return DeadliftExerciseNameKettlebellSumoDeadlift
	case "romanian_deadlift":
		return DeadliftExerciseNameRomanianDeadlift
	case "single_leg_romanian_deadlift_circuit":
		return DeadliftExerciseNameSingleLegRomanianDeadliftCircuit
	case "straight_leg_deadlift":
		return DeadliftExerciseNameStraightLegDeadlift
	default:
		return DeadliftExerciseNameInvalid
	}
}

// List returns all constants.
func ListDeadliftExerciseName() []DeadliftExerciseName {
	return []DeadliftExerciseName{
		DeadliftExerciseNameBarbellDeadlift,
		DeadliftExerciseNameBarbellStraightLegDeadlift,
		DeadliftExerciseNameDumbbellDeadlift,
		DeadliftExerciseNameDumbbellSingleLegDeadliftToRow,
		DeadliftExerciseNameDumbbellStraightLegDeadlift,
		DeadliftExerciseNameKettlebellFloorToShelf,
		DeadliftExerciseNameOneArmOneLegDeadlift,
		DeadliftExerciseNameRackPull,
		DeadliftExerciseNameRotationalDumbbellStraightLegDeadlift,
		DeadliftExerciseNameSingleArmDeadlift,
		DeadliftExerciseNameSingleLegBarbellDeadlift,
		DeadliftExerciseNameSingleLegBarbellStraightLegDeadlift,
		DeadliftExerciseNameSingleLegDeadliftWithBarbell,
		DeadliftExerciseNameSingleLegRdlCircuit,
		DeadliftExerciseNameSingleLegRomanianDeadliftWithDumbbell,
		DeadliftExerciseNameSumoDeadlift,
		DeadliftExerciseNameSumoDeadliftHighPull,
		DeadliftExerciseNameTrapBarDeadlift,
		DeadliftExerciseNameWideGripBarbellDeadlift,
		DeadliftExerciseNameKettlebellDeadlift,
		DeadliftExerciseNameKettlebellSumoDeadlift,
		DeadliftExerciseNameRomanianDeadlift,
		DeadliftExerciseNameSingleLegRomanianDeadliftCircuit,
		DeadliftExerciseNameStraightLegDeadlift,
	}
}
