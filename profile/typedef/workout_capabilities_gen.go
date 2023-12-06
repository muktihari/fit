// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WorkoutCapabilities uint32

const (
	WorkoutCapabilitiesInterval         WorkoutCapabilities = 0x00000001
	WorkoutCapabilitiesCustom           WorkoutCapabilities = 0x00000002
	WorkoutCapabilitiesFitnessEquipment WorkoutCapabilities = 0x00000004
	WorkoutCapabilitiesFirstbeat        WorkoutCapabilities = 0x00000008
	WorkoutCapabilitiesNewLeaf          WorkoutCapabilities = 0x00000010
	WorkoutCapabilitiesTcx              WorkoutCapabilities = 0x00000020 // For backwards compatibility. Watch should add missing id fields then clear flag.
	WorkoutCapabilitiesSpeed            WorkoutCapabilities = 0x00000080 // Speed source required for workout step.
	WorkoutCapabilitiesHeartRate        WorkoutCapabilities = 0x00000100 // Heart rate source required for workout step.
	WorkoutCapabilitiesDistance         WorkoutCapabilities = 0x00000200 // Distance source required for workout step.
	WorkoutCapabilitiesCadence          WorkoutCapabilities = 0x00000400 // Cadence source required for workout step.
	WorkoutCapabilitiesPower            WorkoutCapabilities = 0x00000800 // Power source required for workout step.
	WorkoutCapabilitiesGrade            WorkoutCapabilities = 0x00001000 // Grade source required for workout step.
	WorkoutCapabilitiesResistance       WorkoutCapabilities = 0x00002000 // Resistance source required for workout step.
	WorkoutCapabilitiesProtected        WorkoutCapabilities = 0x00004000
	WorkoutCapabilitiesInvalid          WorkoutCapabilities = 0x0 // INVALID
)

var workoutcapabilitiestostrs = map[WorkoutCapabilities]string{
	WorkoutCapabilitiesInterval:         "interval",
	WorkoutCapabilitiesCustom:           "custom",
	WorkoutCapabilitiesFitnessEquipment: "fitness_equipment",
	WorkoutCapabilitiesFirstbeat:        "firstbeat",
	WorkoutCapabilitiesNewLeaf:          "new_leaf",
	WorkoutCapabilitiesTcx:              "tcx",
	WorkoutCapabilitiesSpeed:            "speed",
	WorkoutCapabilitiesHeartRate:        "heart_rate",
	WorkoutCapabilitiesDistance:         "distance",
	WorkoutCapabilitiesCadence:          "cadence",
	WorkoutCapabilitiesPower:            "power",
	WorkoutCapabilitiesGrade:            "grade",
	WorkoutCapabilitiesResistance:       "resistance",
	WorkoutCapabilitiesProtected:        "protected",
	WorkoutCapabilitiesInvalid:          "invalid",
}

func (w WorkoutCapabilities) String() string {
	val, ok := workoutcapabilitiestostrs[w]
	if !ok {
		return strconv.FormatUint(uint64(w), 10)
	}
	return val
}

var strtoworkoutcapabilities = func() map[string]WorkoutCapabilities {
	m := make(map[string]WorkoutCapabilities)
	for t, str := range workoutcapabilitiestostrs {
		m[str] = WorkoutCapabilities(t)
	}
	return m
}()

// FromString parse string into WorkoutCapabilities constant it's represent, return WorkoutCapabilitiesInvalid if not found.
func WorkoutCapabilitiesFromString(s string) WorkoutCapabilities {
	val, ok := strtoworkoutcapabilities[s]
	if !ok {
		return strtoworkoutcapabilities["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWorkoutCapabilities() []WorkoutCapabilities {
	vs := make([]WorkoutCapabilities, 0, len(workoutcapabilitiestostrs))
	for i := range workoutcapabilitiestostrs {
		vs = append(vs, WorkoutCapabilities(i))
	}
	return vs
}
