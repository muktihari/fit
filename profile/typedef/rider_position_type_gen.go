// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type RiderPositionType byte

const (
	RiderPositionTypeSeated               RiderPositionType = 0
	RiderPositionTypeStanding             RiderPositionType = 1
	RiderPositionTypeTransitionToSeated   RiderPositionType = 2
	RiderPositionTypeTransitionToStanding RiderPositionType = 3
	RiderPositionTypeInvalid              RiderPositionType = 0xFF
)

func (r RiderPositionType) Byte() byte { return byte(r) }

func (r RiderPositionType) String() string {
	switch r {
	case RiderPositionTypeSeated:
		return "seated"
	case RiderPositionTypeStanding:
		return "standing"
	case RiderPositionTypeTransitionToSeated:
		return "transition_to_seated"
	case RiderPositionTypeTransitionToStanding:
		return "transition_to_standing"
	default:
		return "RiderPositionTypeInvalid(" + strconv.Itoa(int(r)) + ")"
	}
}

// FromString parse string into RiderPositionType constant it's represent, return RiderPositionTypeInvalid if not found.
func RiderPositionTypeFromString(s string) RiderPositionType {
	switch s {
	case "seated":
		return RiderPositionTypeSeated
	case "standing":
		return RiderPositionTypeStanding
	case "transition_to_seated":
		return RiderPositionTypeTransitionToSeated
	case "transition_to_standing":
		return RiderPositionTypeTransitionToStanding
	default:
		return RiderPositionTypeInvalid
	}
}

// List returns all constants.
func ListRiderPositionType() []RiderPositionType {
	return []RiderPositionType{
		RiderPositionTypeSeated,
		RiderPositionTypeStanding,
		RiderPositionTypeTransitionToSeated,
		RiderPositionTypeTransitionToStanding,
	}
}
