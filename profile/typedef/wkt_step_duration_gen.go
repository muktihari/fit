// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WktStepDuration byte

const (
	WktStepDurationTime                               WktStepDuration = 0
	WktStepDurationDistance                           WktStepDuration = 1
	WktStepDurationHrLessThan                         WktStepDuration = 2
	WktStepDurationHrGreaterThan                      WktStepDuration = 3
	WktStepDurationCalories                           WktStepDuration = 4
	WktStepDurationOpen                               WktStepDuration = 5
	WktStepDurationRepeatUntilStepsCmplt              WktStepDuration = 6
	WktStepDurationRepeatUntilTime                    WktStepDuration = 7
	WktStepDurationRepeatUntilDistance                WktStepDuration = 8
	WktStepDurationRepeatUntilCalories                WktStepDuration = 9
	WktStepDurationRepeatUntilHrLessThan              WktStepDuration = 10
	WktStepDurationRepeatUntilHrGreaterThan           WktStepDuration = 11
	WktStepDurationRepeatUntilPowerLessThan           WktStepDuration = 12
	WktStepDurationRepeatUntilPowerGreaterThan        WktStepDuration = 13
	WktStepDurationPowerLessThan                      WktStepDuration = 14
	WktStepDurationPowerGreaterThan                   WktStepDuration = 15
	WktStepDurationTrainingPeaksTss                   WktStepDuration = 16
	WktStepDurationRepeatUntilPowerLastLapLessThan    WktStepDuration = 17
	WktStepDurationRepeatUntilMaxPowerLastLapLessThan WktStepDuration = 18
	WktStepDurationPower3SLessThan                    WktStepDuration = 19
	WktStepDurationPower10SLessThan                   WktStepDuration = 20
	WktStepDurationPower30SLessThan                   WktStepDuration = 21
	WktStepDurationPower3SGreaterThan                 WktStepDuration = 22
	WktStepDurationPower10SGreaterThan                WktStepDuration = 23
	WktStepDurationPower30SGreaterThan                WktStepDuration = 24
	WktStepDurationPowerLapLessThan                   WktStepDuration = 25
	WktStepDurationPowerLapGreaterThan                WktStepDuration = 26
	WktStepDurationRepeatUntilTrainingPeaksTss        WktStepDuration = 27
	WktStepDurationRepetitionTime                     WktStepDuration = 28
	WktStepDurationReps                               WktStepDuration = 29
	WktStepDurationTimeOnly                           WktStepDuration = 31
	WktStepDurationInvalid                            WktStepDuration = 0xFF // INVALID
)

var wktstepdurationtostrs = map[WktStepDuration]string{
	WktStepDurationTime:                               "time",
	WktStepDurationDistance:                           "distance",
	WktStepDurationHrLessThan:                         "hr_less_than",
	WktStepDurationHrGreaterThan:                      "hr_greater_than",
	WktStepDurationCalories:                           "calories",
	WktStepDurationOpen:                               "open",
	WktStepDurationRepeatUntilStepsCmplt:              "repeat_until_steps_cmplt",
	WktStepDurationRepeatUntilTime:                    "repeat_until_time",
	WktStepDurationRepeatUntilDistance:                "repeat_until_distance",
	WktStepDurationRepeatUntilCalories:                "repeat_until_calories",
	WktStepDurationRepeatUntilHrLessThan:              "repeat_until_hr_less_than",
	WktStepDurationRepeatUntilHrGreaterThan:           "repeat_until_hr_greater_than",
	WktStepDurationRepeatUntilPowerLessThan:           "repeat_until_power_less_than",
	WktStepDurationRepeatUntilPowerGreaterThan:        "repeat_until_power_greater_than",
	WktStepDurationPowerLessThan:                      "power_less_than",
	WktStepDurationPowerGreaterThan:                   "power_greater_than",
	WktStepDurationTrainingPeaksTss:                   "training_peaks_tss",
	WktStepDurationRepeatUntilPowerLastLapLessThan:    "repeat_until_power_last_lap_less_than",
	WktStepDurationRepeatUntilMaxPowerLastLapLessThan: "repeat_until_max_power_last_lap_less_than",
	WktStepDurationPower3SLessThan:                    "power_3s_less_than",
	WktStepDurationPower10SLessThan:                   "power_10s_less_than",
	WktStepDurationPower30SLessThan:                   "power_30s_less_than",
	WktStepDurationPower3SGreaterThan:                 "power_3s_greater_than",
	WktStepDurationPower10SGreaterThan:                "power_10s_greater_than",
	WktStepDurationPower30SGreaterThan:                "power_30s_greater_than",
	WktStepDurationPowerLapLessThan:                   "power_lap_less_than",
	WktStepDurationPowerLapGreaterThan:                "power_lap_greater_than",
	WktStepDurationRepeatUntilTrainingPeaksTss:        "repeat_until_training_peaks_tss",
	WktStepDurationRepetitionTime:                     "repetition_time",
	WktStepDurationReps:                               "reps",
	WktStepDurationTimeOnly:                           "time_only",
	WktStepDurationInvalid:                            "invalid",
}

func (w WktStepDuration) String() string {
	val, ok := wktstepdurationtostrs[w]
	if !ok {
		return strconv.Itoa(int(w))
	}
	return val
}

var strtowktstepduration = func() map[string]WktStepDuration {
	m := make(map[string]WktStepDuration)
	for t, str := range wktstepdurationtostrs {
		m[str] = WktStepDuration(t)
	}
	return m
}()

// FromString parse string into WktStepDuration constant it's represent, return WktStepDurationInvalid if not found.
func WktStepDurationFromString(s string) WktStepDuration {
	val, ok := strtowktstepduration[s]
	if !ok {
		return strtowktstepduration["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWktStepDuration() []WktStepDuration {
	vs := make([]WktStepDuration, 0, len(wktstepdurationtostrs))
	for i := range wktstepdurationtostrs {
		vs = append(vs, WktStepDuration(i))
	}
	return vs
}
