// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WorkoutPower uint32

const (
	WorkoutPowerWattsOffset WorkoutPower = 1000
	WorkoutPowerInvalid     WorkoutPower = 0xFFFFFFFF // INVALID
)

var workoutpowertostrs = map[WorkoutPower]string{
	WorkoutPowerWattsOffset: "watts_offset",
	WorkoutPowerInvalid:     "invalid",
}

func (w WorkoutPower) String() string {
	val, ok := workoutpowertostrs[w]
	if !ok {
		return strconv.FormatUint(uint64(w), 10)
	}
	return val
}

var strtoworkoutpower = func() map[string]WorkoutPower {
	m := make(map[string]WorkoutPower)
	for t, str := range workoutpowertostrs {
		m[str] = WorkoutPower(t)
	}
	return m
}()

// FromString parse string into WorkoutPower constant it's represent, return WorkoutPowerInvalid if not found.
func WorkoutPowerFromString(s string) WorkoutPower {
	val, ok := strtoworkoutpower[s]
	if !ok {
		return strtoworkoutpower["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWorkoutPower() []WorkoutPower {
	vs := make([]WorkoutPower, 0, len(workoutpowertostrs))
	for i := range workoutpowertostrs {
		vs = append(vs, WorkoutPower(i))
	}
	return vs
}
