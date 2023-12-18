// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ExerciseCategory uint16

const (
	ExerciseCategoryBenchPress        ExerciseCategory = 0
	ExerciseCategoryCalfRaise         ExerciseCategory = 1
	ExerciseCategoryCardio            ExerciseCategory = 2
	ExerciseCategoryCarry             ExerciseCategory = 3
	ExerciseCategoryChop              ExerciseCategory = 4
	ExerciseCategoryCore              ExerciseCategory = 5
	ExerciseCategoryCrunch            ExerciseCategory = 6
	ExerciseCategoryCurl              ExerciseCategory = 7
	ExerciseCategoryDeadlift          ExerciseCategory = 8
	ExerciseCategoryFlye              ExerciseCategory = 9
	ExerciseCategoryHipRaise          ExerciseCategory = 10
	ExerciseCategoryHipStability      ExerciseCategory = 11
	ExerciseCategoryHipSwing          ExerciseCategory = 12
	ExerciseCategoryHyperextension    ExerciseCategory = 13
	ExerciseCategoryLateralRaise      ExerciseCategory = 14
	ExerciseCategoryLegCurl           ExerciseCategory = 15
	ExerciseCategoryLegRaise          ExerciseCategory = 16
	ExerciseCategoryLunge             ExerciseCategory = 17
	ExerciseCategoryOlympicLift       ExerciseCategory = 18
	ExerciseCategoryPlank             ExerciseCategory = 19
	ExerciseCategoryPlyo              ExerciseCategory = 20
	ExerciseCategoryPullUp            ExerciseCategory = 21
	ExerciseCategoryPushUp            ExerciseCategory = 22
	ExerciseCategoryRow               ExerciseCategory = 23
	ExerciseCategoryShoulderPress     ExerciseCategory = 24
	ExerciseCategoryShoulderStability ExerciseCategory = 25
	ExerciseCategoryShrug             ExerciseCategory = 26
	ExerciseCategorySitUp             ExerciseCategory = 27
	ExerciseCategorySquat             ExerciseCategory = 28
	ExerciseCategoryTotalBody         ExerciseCategory = 29
	ExerciseCategoryTricepsExtension  ExerciseCategory = 30
	ExerciseCategoryWarmUp            ExerciseCategory = 31
	ExerciseCategoryRun               ExerciseCategory = 32
	ExerciseCategoryUnknown           ExerciseCategory = 65534
	ExerciseCategoryInvalid           ExerciseCategory = 0xFFFF
)

func (e ExerciseCategory) String() string {
	switch e {
	case ExerciseCategoryBenchPress:
		return "bench_press"
	case ExerciseCategoryCalfRaise:
		return "calf_raise"
	case ExerciseCategoryCardio:
		return "cardio"
	case ExerciseCategoryCarry:
		return "carry"
	case ExerciseCategoryChop:
		return "chop"
	case ExerciseCategoryCore:
		return "core"
	case ExerciseCategoryCrunch:
		return "crunch"
	case ExerciseCategoryCurl:
		return "curl"
	case ExerciseCategoryDeadlift:
		return "deadlift"
	case ExerciseCategoryFlye:
		return "flye"
	case ExerciseCategoryHipRaise:
		return "hip_raise"
	case ExerciseCategoryHipStability:
		return "hip_stability"
	case ExerciseCategoryHipSwing:
		return "hip_swing"
	case ExerciseCategoryHyperextension:
		return "hyperextension"
	case ExerciseCategoryLateralRaise:
		return "lateral_raise"
	case ExerciseCategoryLegCurl:
		return "leg_curl"
	case ExerciseCategoryLegRaise:
		return "leg_raise"
	case ExerciseCategoryLunge:
		return "lunge"
	case ExerciseCategoryOlympicLift:
		return "olympic_lift"
	case ExerciseCategoryPlank:
		return "plank"
	case ExerciseCategoryPlyo:
		return "plyo"
	case ExerciseCategoryPullUp:
		return "pull_up"
	case ExerciseCategoryPushUp:
		return "push_up"
	case ExerciseCategoryRow:
		return "row"
	case ExerciseCategoryShoulderPress:
		return "shoulder_press"
	case ExerciseCategoryShoulderStability:
		return "shoulder_stability"
	case ExerciseCategoryShrug:
		return "shrug"
	case ExerciseCategorySitUp:
		return "sit_up"
	case ExerciseCategorySquat:
		return "squat"
	case ExerciseCategoryTotalBody:
		return "total_body"
	case ExerciseCategoryTricepsExtension:
		return "triceps_extension"
	case ExerciseCategoryWarmUp:
		return "warm_up"
	case ExerciseCategoryRun:
		return "run"
	case ExerciseCategoryUnknown:
		return "unknown"
	default:
		return "ExerciseCategoryInvalid(" + strconv.FormatUint(uint64(e), 10) + ")"
	}
}

// FromString parse string into ExerciseCategory constant it's represent, return ExerciseCategoryInvalid if not found.
func ExerciseCategoryFromString(s string) ExerciseCategory {
	switch s {
	case "bench_press":
		return ExerciseCategoryBenchPress
	case "calf_raise":
		return ExerciseCategoryCalfRaise
	case "cardio":
		return ExerciseCategoryCardio
	case "carry":
		return ExerciseCategoryCarry
	case "chop":
		return ExerciseCategoryChop
	case "core":
		return ExerciseCategoryCore
	case "crunch":
		return ExerciseCategoryCrunch
	case "curl":
		return ExerciseCategoryCurl
	case "deadlift":
		return ExerciseCategoryDeadlift
	case "flye":
		return ExerciseCategoryFlye
	case "hip_raise":
		return ExerciseCategoryHipRaise
	case "hip_stability":
		return ExerciseCategoryHipStability
	case "hip_swing":
		return ExerciseCategoryHipSwing
	case "hyperextension":
		return ExerciseCategoryHyperextension
	case "lateral_raise":
		return ExerciseCategoryLateralRaise
	case "leg_curl":
		return ExerciseCategoryLegCurl
	case "leg_raise":
		return ExerciseCategoryLegRaise
	case "lunge":
		return ExerciseCategoryLunge
	case "olympic_lift":
		return ExerciseCategoryOlympicLift
	case "plank":
		return ExerciseCategoryPlank
	case "plyo":
		return ExerciseCategoryPlyo
	case "pull_up":
		return ExerciseCategoryPullUp
	case "push_up":
		return ExerciseCategoryPushUp
	case "row":
		return ExerciseCategoryRow
	case "shoulder_press":
		return ExerciseCategoryShoulderPress
	case "shoulder_stability":
		return ExerciseCategoryShoulderStability
	case "shrug":
		return ExerciseCategoryShrug
	case "sit_up":
		return ExerciseCategorySitUp
	case "squat":
		return ExerciseCategorySquat
	case "total_body":
		return ExerciseCategoryTotalBody
	case "triceps_extension":
		return ExerciseCategoryTricepsExtension
	case "warm_up":
		return ExerciseCategoryWarmUp
	case "run":
		return ExerciseCategoryRun
	case "unknown":
		return ExerciseCategoryUnknown
	default:
		return ExerciseCategoryInvalid
	}
}

// List returns all constants.
func ListExerciseCategory() []ExerciseCategory {
	return []ExerciseCategory{
		ExerciseCategoryBenchPress,
		ExerciseCategoryCalfRaise,
		ExerciseCategoryCardio,
		ExerciseCategoryCarry,
		ExerciseCategoryChop,
		ExerciseCategoryCore,
		ExerciseCategoryCrunch,
		ExerciseCategoryCurl,
		ExerciseCategoryDeadlift,
		ExerciseCategoryFlye,
		ExerciseCategoryHipRaise,
		ExerciseCategoryHipStability,
		ExerciseCategoryHipSwing,
		ExerciseCategoryHyperextension,
		ExerciseCategoryLateralRaise,
		ExerciseCategoryLegCurl,
		ExerciseCategoryLegRaise,
		ExerciseCategoryLunge,
		ExerciseCategoryOlympicLift,
		ExerciseCategoryPlank,
		ExerciseCategoryPlyo,
		ExerciseCategoryPullUp,
		ExerciseCategoryPushUp,
		ExerciseCategoryRow,
		ExerciseCategoryShoulderPress,
		ExerciseCategoryShoulderStability,
		ExerciseCategoryShrug,
		ExerciseCategorySitUp,
		ExerciseCategorySquat,
		ExerciseCategoryTotalBody,
		ExerciseCategoryTricepsExtension,
		ExerciseCategoryWarmUp,
		ExerciseCategoryRun,
		ExerciseCategoryUnknown,
	}
}
