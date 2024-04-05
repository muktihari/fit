// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SplitType byte

const (
	SplitTypeAscentSplit      SplitType = 1
	SplitTypeDescentSplit     SplitType = 2
	SplitTypeIntervalActive   SplitType = 3
	SplitTypeIntervalRest     SplitType = 4
	SplitTypeIntervalWarmup   SplitType = 5
	SplitTypeIntervalCooldown SplitType = 6
	SplitTypeIntervalRecovery SplitType = 7
	SplitTypeIntervalOther    SplitType = 8
	SplitTypeClimbActive      SplitType = 9
	SplitTypeClimbRest        SplitType = 10
	SplitTypeSurfActive       SplitType = 11
	SplitTypeRunActive        SplitType = 12
	SplitTypeRunRest          SplitType = 13
	SplitTypeWorkoutRound     SplitType = 14
	SplitTypeRwdRun           SplitType = 17 // run/walk detection running
	SplitTypeRwdWalk          SplitType = 18 // run/walk detection walking
	SplitTypeWindsurfActive   SplitType = 21
	SplitTypeRwdStand         SplitType = 22 // run/walk detection standing
	SplitTypeTransition       SplitType = 23 // Marks the time going from ascent_split to descent_split/used in backcountry ski
	SplitTypeSkiLiftSplit     SplitType = 28
	SplitTypeSkiRunSplit      SplitType = 29
	SplitTypeInvalid          SplitType = 0xFF
)

func (s SplitType) Byte() byte { return byte(s) }

func (s SplitType) String() string {
	switch s {
	case SplitTypeAscentSplit:
		return "ascent_split"
	case SplitTypeDescentSplit:
		return "descent_split"
	case SplitTypeIntervalActive:
		return "interval_active"
	case SplitTypeIntervalRest:
		return "interval_rest"
	case SplitTypeIntervalWarmup:
		return "interval_warmup"
	case SplitTypeIntervalCooldown:
		return "interval_cooldown"
	case SplitTypeIntervalRecovery:
		return "interval_recovery"
	case SplitTypeIntervalOther:
		return "interval_other"
	case SplitTypeClimbActive:
		return "climb_active"
	case SplitTypeClimbRest:
		return "climb_rest"
	case SplitTypeSurfActive:
		return "surf_active"
	case SplitTypeRunActive:
		return "run_active"
	case SplitTypeRunRest:
		return "run_rest"
	case SplitTypeWorkoutRound:
		return "workout_round"
	case SplitTypeRwdRun:
		return "rwd_run"
	case SplitTypeRwdWalk:
		return "rwd_walk"
	case SplitTypeWindsurfActive:
		return "windsurf_active"
	case SplitTypeRwdStand:
		return "rwd_stand"
	case SplitTypeTransition:
		return "transition"
	case SplitTypeSkiLiftSplit:
		return "ski_lift_split"
	case SplitTypeSkiRunSplit:
		return "ski_run_split"
	default:
		return "SplitTypeInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SplitType constant it's represent, return SplitTypeInvalid if not found.
func SplitTypeFromString(s string) SplitType {
	switch s {
	case "ascent_split":
		return SplitTypeAscentSplit
	case "descent_split":
		return SplitTypeDescentSplit
	case "interval_active":
		return SplitTypeIntervalActive
	case "interval_rest":
		return SplitTypeIntervalRest
	case "interval_warmup":
		return SplitTypeIntervalWarmup
	case "interval_cooldown":
		return SplitTypeIntervalCooldown
	case "interval_recovery":
		return SplitTypeIntervalRecovery
	case "interval_other":
		return SplitTypeIntervalOther
	case "climb_active":
		return SplitTypeClimbActive
	case "climb_rest":
		return SplitTypeClimbRest
	case "surf_active":
		return SplitTypeSurfActive
	case "run_active":
		return SplitTypeRunActive
	case "run_rest":
		return SplitTypeRunRest
	case "workout_round":
		return SplitTypeWorkoutRound
	case "rwd_run":
		return SplitTypeRwdRun
	case "rwd_walk":
		return SplitTypeRwdWalk
	case "windsurf_active":
		return SplitTypeWindsurfActive
	case "rwd_stand":
		return SplitTypeRwdStand
	case "transition":
		return SplitTypeTransition
	case "ski_lift_split":
		return SplitTypeSkiLiftSplit
	case "ski_run_split":
		return SplitTypeSkiRunSplit
	default:
		return SplitTypeInvalid
	}
}

// List returns all constants.
func ListSplitType() []SplitType {
	return []SplitType{
		SplitTypeAscentSplit,
		SplitTypeDescentSplit,
		SplitTypeIntervalActive,
		SplitTypeIntervalRest,
		SplitTypeIntervalWarmup,
		SplitTypeIntervalCooldown,
		SplitTypeIntervalRecovery,
		SplitTypeIntervalOther,
		SplitTypeClimbActive,
		SplitTypeClimbRest,
		SplitTypeSurfActive,
		SplitTypeRunActive,
		SplitTypeRunRest,
		SplitTypeWorkoutRound,
		SplitTypeRwdRun,
		SplitTypeRwdWalk,
		SplitTypeWindsurfActive,
		SplitTypeRwdStand,
		SplitTypeTransition,
		SplitTypeSkiLiftSplit,
		SplitTypeSkiRunSplit,
	}
}
