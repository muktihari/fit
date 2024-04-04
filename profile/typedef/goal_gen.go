// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Goal byte

const (
	GoalTime          Goal = 0
	GoalDistance      Goal = 1
	GoalCalories      Goal = 2
	GoalFrequency     Goal = 3
	GoalSteps         Goal = 4
	GoalAscent        Goal = 5
	GoalActiveMinutes Goal = 6
	GoalInvalid       Goal = 0xFF
)

func (g Goal) Byte() byte { return byte(g) }

func (g Goal) String() string {
	switch g {
	case GoalTime:
		return "time"
	case GoalDistance:
		return "distance"
	case GoalCalories:
		return "calories"
	case GoalFrequency:
		return "frequency"
	case GoalSteps:
		return "steps"
	case GoalAscent:
		return "ascent"
	case GoalActiveMinutes:
		return "active_minutes"
	default:
		return "GoalInvalid(" + strconv.Itoa(int(g)) + ")"
	}
}

// FromString parse string into Goal constant it's represent, return GoalInvalid if not found.
func GoalFromString(s string) Goal {
	switch s {
	case "time":
		return GoalTime
	case "distance":
		return GoalDistance
	case "calories":
		return GoalCalories
	case "frequency":
		return GoalFrequency
	case "steps":
		return GoalSteps
	case "ascent":
		return GoalAscent
	case "active_minutes":
		return GoalActiveMinutes
	default:
		return GoalInvalid
	}
}

// List returns all constants.
func ListGoal() []Goal {
	return []Goal{
		GoalTime,
		GoalDistance,
		GoalCalories,
		GoalFrequency,
		GoalSteps,
		GoalAscent,
		GoalActiveMinutes,
	}
}
