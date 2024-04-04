// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TricepsExtensionExerciseName uint16

const (
	TricepsExtensionExerciseNameBenchDip                                     TricepsExtensionExerciseName = 0
	TricepsExtensionExerciseNameWeightedBenchDip                             TricepsExtensionExerciseName = 1
	TricepsExtensionExerciseNameBodyWeightDip                                TricepsExtensionExerciseName = 2
	TricepsExtensionExerciseNameCableKickback                                TricepsExtensionExerciseName = 3
	TricepsExtensionExerciseNameCableLyingTricepsExtension                   TricepsExtensionExerciseName = 4
	TricepsExtensionExerciseNameCableOverheadTricepsExtension                TricepsExtensionExerciseName = 5
	TricepsExtensionExerciseNameDumbbellKickback                             TricepsExtensionExerciseName = 6
	TricepsExtensionExerciseNameDumbbellLyingTricepsExtension                TricepsExtensionExerciseName = 7
	TricepsExtensionExerciseNameEzBarOverheadTricepsExtension                TricepsExtensionExerciseName = 8
	TricepsExtensionExerciseNameInclineDip                                   TricepsExtensionExerciseName = 9
	TricepsExtensionExerciseNameWeightedInclineDip                           TricepsExtensionExerciseName = 10
	TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension            TricepsExtensionExerciseName = 11
	TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension             TricepsExtensionExerciseName = 12
	TricepsExtensionExerciseNameLyingEzBarTricepsExtension                   TricepsExtensionExerciseName = 13
	TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress   TricepsExtensionExerciseName = 14
	TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension             TricepsExtensionExerciseName = 15
	TricepsExtensionExerciseNameRecliningTricepsPress                        TricepsExtensionExerciseName = 16
	TricepsExtensionExerciseNameReverseGripPressdown                         TricepsExtensionExerciseName = 17
	TricepsExtensionExerciseNameReverseGripTricepsPressdown                  TricepsExtensionExerciseName = 18
	TricepsExtensionExerciseNameRopePressdown                                TricepsExtensionExerciseName = 19
	TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension        TricepsExtensionExerciseName = 20
	TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension       TricepsExtensionExerciseName = 21
	TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension          TricepsExtensionExerciseName = 22
	TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension     TricepsExtensionExerciseName = 23
	TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension    TricepsExtensionExerciseName = 24
	TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension TricepsExtensionExerciseName = 25
	TricepsExtensionExerciseNameSingleLegBenchDipAndKick                     TricepsExtensionExerciseName = 26
	TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick             TricepsExtensionExerciseName = 27
	TricepsExtensionExerciseNameSingleLegDip                                 TricepsExtensionExerciseName = 28
	TricepsExtensionExerciseNameWeightedSingleLegDip                         TricepsExtensionExerciseName = 29
	TricepsExtensionExerciseNameStaticLyingTricepsExtension                  TricepsExtensionExerciseName = 30
	TricepsExtensionExerciseNameSuspendedDip                                 TricepsExtensionExerciseName = 31
	TricepsExtensionExerciseNameWeightedSuspendedDip                         TricepsExtensionExerciseName = 32
	TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension       TricepsExtensionExerciseName = 33
	TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension          TricepsExtensionExerciseName = 34
	TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension       TricepsExtensionExerciseName = 35
	TricepsExtensionExerciseNameTabletopDip                                  TricepsExtensionExerciseName = 36
	TricepsExtensionExerciseNameWeightedTabletopDip                          TricepsExtensionExerciseName = 37
	TricepsExtensionExerciseNameTricepsExtensionOnFloor                      TricepsExtensionExerciseName = 38
	TricepsExtensionExerciseNameTricepsPressdown                             TricepsExtensionExerciseName = 39
	TricepsExtensionExerciseNameWeightedDip                                  TricepsExtensionExerciseName = 40
	TricepsExtensionExerciseNameInvalid                                      TricepsExtensionExerciseName = 0xFFFF
)

func (t TricepsExtensionExerciseName) Uint16() uint16 { return uint16(t) }

func (t TricepsExtensionExerciseName) String() string {
	switch t {
	case TricepsExtensionExerciseNameBenchDip:
		return "bench_dip"
	case TricepsExtensionExerciseNameWeightedBenchDip:
		return "weighted_bench_dip"
	case TricepsExtensionExerciseNameBodyWeightDip:
		return "body_weight_dip"
	case TricepsExtensionExerciseNameCableKickback:
		return "cable_kickback"
	case TricepsExtensionExerciseNameCableLyingTricepsExtension:
		return "cable_lying_triceps_extension"
	case TricepsExtensionExerciseNameCableOverheadTricepsExtension:
		return "cable_overhead_triceps_extension"
	case TricepsExtensionExerciseNameDumbbellKickback:
		return "dumbbell_kickback"
	case TricepsExtensionExerciseNameDumbbellLyingTricepsExtension:
		return "dumbbell_lying_triceps_extension"
	case TricepsExtensionExerciseNameEzBarOverheadTricepsExtension:
		return "ez_bar_overhead_triceps_extension"
	case TricepsExtensionExerciseNameInclineDip:
		return "incline_dip"
	case TricepsExtensionExerciseNameWeightedInclineDip:
		return "weighted_incline_dip"
	case TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension:
		return "incline_ez_bar_lying_triceps_extension"
	case TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension:
		return "lying_dumbbell_pullover_to_extension"
	case TricepsExtensionExerciseNameLyingEzBarTricepsExtension:
		return "lying_ez_bar_triceps_extension"
	case TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress:
		return "lying_triceps_extension_to_close_grip_bench_press"
	case TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension:
		return "overhead_dumbbell_triceps_extension"
	case TricepsExtensionExerciseNameRecliningTricepsPress:
		return "reclining_triceps_press"
	case TricepsExtensionExerciseNameReverseGripPressdown:
		return "reverse_grip_pressdown"
	case TricepsExtensionExerciseNameReverseGripTricepsPressdown:
		return "reverse_grip_triceps_pressdown"
	case TricepsExtensionExerciseNameRopePressdown:
		return "rope_pressdown"
	case TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension:
		return "seated_barbell_overhead_triceps_extension"
	case TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension:
		return "seated_dumbbell_overhead_triceps_extension"
	case TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension:
		return "seated_ez_bar_overhead_triceps_extension"
	case TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension:
		return "seated_single_arm_overhead_dumbbell_extension"
	case TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension:
		return "single_arm_dumbbell_overhead_triceps_extension"
	case TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension:
		return "single_dumbbell_seated_overhead_triceps_extension"
	case TricepsExtensionExerciseNameSingleLegBenchDipAndKick:
		return "single_leg_bench_dip_and_kick"
	case TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick:
		return "weighted_single_leg_bench_dip_and_kick"
	case TricepsExtensionExerciseNameSingleLegDip:
		return "single_leg_dip"
	case TricepsExtensionExerciseNameWeightedSingleLegDip:
		return "weighted_single_leg_dip"
	case TricepsExtensionExerciseNameStaticLyingTricepsExtension:
		return "static_lying_triceps_extension"
	case TricepsExtensionExerciseNameSuspendedDip:
		return "suspended_dip"
	case TricepsExtensionExerciseNameWeightedSuspendedDip:
		return "weighted_suspended_dip"
	case TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension:
		return "swiss_ball_dumbbell_lying_triceps_extension"
	case TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension:
		return "swiss_ball_ez_bar_lying_triceps_extension"
	case TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension:
		return "swiss_ball_ez_bar_overhead_triceps_extension"
	case TricepsExtensionExerciseNameTabletopDip:
		return "tabletop_dip"
	case TricepsExtensionExerciseNameWeightedTabletopDip:
		return "weighted_tabletop_dip"
	case TricepsExtensionExerciseNameTricepsExtensionOnFloor:
		return "triceps_extension_on_floor"
	case TricepsExtensionExerciseNameTricepsPressdown:
		return "triceps_pressdown"
	case TricepsExtensionExerciseNameWeightedDip:
		return "weighted_dip"
	default:
		return "TricepsExtensionExerciseNameInvalid(" + strconv.FormatUint(uint64(t), 10) + ")"
	}
}

// FromString parse string into TricepsExtensionExerciseName constant it's represent, return TricepsExtensionExerciseNameInvalid if not found.
func TricepsExtensionExerciseNameFromString(s string) TricepsExtensionExerciseName {
	switch s {
	case "bench_dip":
		return TricepsExtensionExerciseNameBenchDip
	case "weighted_bench_dip":
		return TricepsExtensionExerciseNameWeightedBenchDip
	case "body_weight_dip":
		return TricepsExtensionExerciseNameBodyWeightDip
	case "cable_kickback":
		return TricepsExtensionExerciseNameCableKickback
	case "cable_lying_triceps_extension":
		return TricepsExtensionExerciseNameCableLyingTricepsExtension
	case "cable_overhead_triceps_extension":
		return TricepsExtensionExerciseNameCableOverheadTricepsExtension
	case "dumbbell_kickback":
		return TricepsExtensionExerciseNameDumbbellKickback
	case "dumbbell_lying_triceps_extension":
		return TricepsExtensionExerciseNameDumbbellLyingTricepsExtension
	case "ez_bar_overhead_triceps_extension":
		return TricepsExtensionExerciseNameEzBarOverheadTricepsExtension
	case "incline_dip":
		return TricepsExtensionExerciseNameInclineDip
	case "weighted_incline_dip":
		return TricepsExtensionExerciseNameWeightedInclineDip
	case "incline_ez_bar_lying_triceps_extension":
		return TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension
	case "lying_dumbbell_pullover_to_extension":
		return TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension
	case "lying_ez_bar_triceps_extension":
		return TricepsExtensionExerciseNameLyingEzBarTricepsExtension
	case "lying_triceps_extension_to_close_grip_bench_press":
		return TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress
	case "overhead_dumbbell_triceps_extension":
		return TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension
	case "reclining_triceps_press":
		return TricepsExtensionExerciseNameRecliningTricepsPress
	case "reverse_grip_pressdown":
		return TricepsExtensionExerciseNameReverseGripPressdown
	case "reverse_grip_triceps_pressdown":
		return TricepsExtensionExerciseNameReverseGripTricepsPressdown
	case "rope_pressdown":
		return TricepsExtensionExerciseNameRopePressdown
	case "seated_barbell_overhead_triceps_extension":
		return TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension
	case "seated_dumbbell_overhead_triceps_extension":
		return TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension
	case "seated_ez_bar_overhead_triceps_extension":
		return TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension
	case "seated_single_arm_overhead_dumbbell_extension":
		return TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension
	case "single_arm_dumbbell_overhead_triceps_extension":
		return TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension
	case "single_dumbbell_seated_overhead_triceps_extension":
		return TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension
	case "single_leg_bench_dip_and_kick":
		return TricepsExtensionExerciseNameSingleLegBenchDipAndKick
	case "weighted_single_leg_bench_dip_and_kick":
		return TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick
	case "single_leg_dip":
		return TricepsExtensionExerciseNameSingleLegDip
	case "weighted_single_leg_dip":
		return TricepsExtensionExerciseNameWeightedSingleLegDip
	case "static_lying_triceps_extension":
		return TricepsExtensionExerciseNameStaticLyingTricepsExtension
	case "suspended_dip":
		return TricepsExtensionExerciseNameSuspendedDip
	case "weighted_suspended_dip":
		return TricepsExtensionExerciseNameWeightedSuspendedDip
	case "swiss_ball_dumbbell_lying_triceps_extension":
		return TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension
	case "swiss_ball_ez_bar_lying_triceps_extension":
		return TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension
	case "swiss_ball_ez_bar_overhead_triceps_extension":
		return TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension
	case "tabletop_dip":
		return TricepsExtensionExerciseNameTabletopDip
	case "weighted_tabletop_dip":
		return TricepsExtensionExerciseNameWeightedTabletopDip
	case "triceps_extension_on_floor":
		return TricepsExtensionExerciseNameTricepsExtensionOnFloor
	case "triceps_pressdown":
		return TricepsExtensionExerciseNameTricepsPressdown
	case "weighted_dip":
		return TricepsExtensionExerciseNameWeightedDip
	default:
		return TricepsExtensionExerciseNameInvalid
	}
}

// List returns all constants.
func ListTricepsExtensionExerciseName() []TricepsExtensionExerciseName {
	return []TricepsExtensionExerciseName{
		TricepsExtensionExerciseNameBenchDip,
		TricepsExtensionExerciseNameWeightedBenchDip,
		TricepsExtensionExerciseNameBodyWeightDip,
		TricepsExtensionExerciseNameCableKickback,
		TricepsExtensionExerciseNameCableLyingTricepsExtension,
		TricepsExtensionExerciseNameCableOverheadTricepsExtension,
		TricepsExtensionExerciseNameDumbbellKickback,
		TricepsExtensionExerciseNameDumbbellLyingTricepsExtension,
		TricepsExtensionExerciseNameEzBarOverheadTricepsExtension,
		TricepsExtensionExerciseNameInclineDip,
		TricepsExtensionExerciseNameWeightedInclineDip,
		TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension,
		TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension,
		TricepsExtensionExerciseNameLyingEzBarTricepsExtension,
		TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress,
		TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension,
		TricepsExtensionExerciseNameRecliningTricepsPress,
		TricepsExtensionExerciseNameReverseGripPressdown,
		TricepsExtensionExerciseNameReverseGripTricepsPressdown,
		TricepsExtensionExerciseNameRopePressdown,
		TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension,
		TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension,
		TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension,
		TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension,
		TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension,
		TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension,
		TricepsExtensionExerciseNameSingleLegBenchDipAndKick,
		TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick,
		TricepsExtensionExerciseNameSingleLegDip,
		TricepsExtensionExerciseNameWeightedSingleLegDip,
		TricepsExtensionExerciseNameStaticLyingTricepsExtension,
		TricepsExtensionExerciseNameSuspendedDip,
		TricepsExtensionExerciseNameWeightedSuspendedDip,
		TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension,
		TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension,
		TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension,
		TricepsExtensionExerciseNameTabletopDip,
		TricepsExtensionExerciseNameWeightedTabletopDip,
		TricepsExtensionExerciseNameTricepsExtensionOnFloor,
		TricepsExtensionExerciseNameTricepsPressdown,
		TricepsExtensionExerciseNameWeightedDip,
	}
}
