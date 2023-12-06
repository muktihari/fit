// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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
	RiderPositionTypeInvalid              RiderPositionType = 0xFF // INVALID
)

var riderpositiontypetostrs = map[RiderPositionType]string{
	RiderPositionTypeSeated:               "seated",
	RiderPositionTypeStanding:             "standing",
	RiderPositionTypeTransitionToSeated:   "transition_to_seated",
	RiderPositionTypeTransitionToStanding: "transition_to_standing",
	RiderPositionTypeInvalid:              "invalid",
}

func (r RiderPositionType) String() string {
	val, ok := riderpositiontypetostrs[r]
	if !ok {
		return strconv.Itoa(int(r))
	}
	return val
}

var strtoriderpositiontype = func() map[string]RiderPositionType {
	m := make(map[string]RiderPositionType)
	for t, str := range riderpositiontypetostrs {
		m[str] = RiderPositionType(t)
	}
	return m
}()

// FromString parse string into RiderPositionType constant it's represent, return RiderPositionTypeInvalid if not found.
func RiderPositionTypeFromString(s string) RiderPositionType {
	val, ok := strtoriderpositiontype[s]
	if !ok {
		return strtoriderpositiontype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListRiderPositionType() []RiderPositionType {
	vs := make([]RiderPositionType, 0, len(riderpositiontypetostrs))
	for i := range riderpositiontypetostrs {
		vs = append(vs, RiderPositionType(i))
	}
	return vs
}
