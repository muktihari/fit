// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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
	SplitTypeInvalid          SplitType = 0xFF // INVALID
)

var splittypetostrs = map[SplitType]string{
	SplitTypeAscentSplit:      "ascent_split",
	SplitTypeDescentSplit:     "descent_split",
	SplitTypeIntervalActive:   "interval_active",
	SplitTypeIntervalRest:     "interval_rest",
	SplitTypeIntervalWarmup:   "interval_warmup",
	SplitTypeIntervalCooldown: "interval_cooldown",
	SplitTypeIntervalRecovery: "interval_recovery",
	SplitTypeIntervalOther:    "interval_other",
	SplitTypeClimbActive:      "climb_active",
	SplitTypeClimbRest:        "climb_rest",
	SplitTypeSurfActive:       "surf_active",
	SplitTypeRunActive:        "run_active",
	SplitTypeRunRest:          "run_rest",
	SplitTypeWorkoutRound:     "workout_round",
	SplitTypeRwdRun:           "rwd_run",
	SplitTypeRwdWalk:          "rwd_walk",
	SplitTypeWindsurfActive:   "windsurf_active",
	SplitTypeRwdStand:         "rwd_stand",
	SplitTypeTransition:       "transition",
	SplitTypeSkiLiftSplit:     "ski_lift_split",
	SplitTypeSkiRunSplit:      "ski_run_split",
	SplitTypeInvalid:          "invalid",
}

func (s SplitType) String() string {
	val, ok := splittypetostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosplittype = func() map[string]SplitType {
	m := make(map[string]SplitType)
	for t, str := range splittypetostrs {
		m[str] = SplitType(t)
	}
	return m
}()

// FromString parse string into SplitType constant it's represent, return SplitTypeInvalid if not found.
func SplitTypeFromString(s string) SplitType {
	val, ok := strtosplittype[s]
	if !ok {
		return strtosplittype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSplitType() []SplitType {
	vs := make([]SplitType, 0, len(splittypetostrs))
	for i := range splittypetostrs {
		vs = append(vs, SplitType(i))
	}
	return vs
}
