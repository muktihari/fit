// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ExdDataUnits byte

const (
	ExdDataUnitsNoUnits                        ExdDataUnits = 0
	ExdDataUnitsLaps                           ExdDataUnits = 1
	ExdDataUnitsMilesPerHour                   ExdDataUnits = 2
	ExdDataUnitsKilometersPerHour              ExdDataUnits = 3
	ExdDataUnitsFeetPerHour                    ExdDataUnits = 4
	ExdDataUnitsMetersPerHour                  ExdDataUnits = 5
	ExdDataUnitsDegreesCelsius                 ExdDataUnits = 6
	ExdDataUnitsDegreesFarenheit               ExdDataUnits = 7
	ExdDataUnitsZone                           ExdDataUnits = 8
	ExdDataUnitsGear                           ExdDataUnits = 9
	ExdDataUnitsRpm                            ExdDataUnits = 10
	ExdDataUnitsBpm                            ExdDataUnits = 11
	ExdDataUnitsDegrees                        ExdDataUnits = 12
	ExdDataUnitsMillimeters                    ExdDataUnits = 13
	ExdDataUnitsMeters                         ExdDataUnits = 14
	ExdDataUnitsKilometers                     ExdDataUnits = 15
	ExdDataUnitsFeet                           ExdDataUnits = 16
	ExdDataUnitsYards                          ExdDataUnits = 17
	ExdDataUnitsKilofeet                       ExdDataUnits = 18
	ExdDataUnitsMiles                          ExdDataUnits = 19
	ExdDataUnitsTime                           ExdDataUnits = 20
	ExdDataUnitsEnumTurnType                   ExdDataUnits = 21
	ExdDataUnitsPercent                        ExdDataUnits = 22
	ExdDataUnitsWatts                          ExdDataUnits = 23
	ExdDataUnitsWattsPerKilogram               ExdDataUnits = 24
	ExdDataUnitsEnumBatteryStatus              ExdDataUnits = 25
	ExdDataUnitsEnumBikeLightBeamAngleMode     ExdDataUnits = 26
	ExdDataUnitsEnumBikeLightBatteryStatus     ExdDataUnits = 27
	ExdDataUnitsEnumBikeLightNetworkConfigType ExdDataUnits = 28
	ExdDataUnitsLights                         ExdDataUnits = 29
	ExdDataUnitsSeconds                        ExdDataUnits = 30
	ExdDataUnitsMinutes                        ExdDataUnits = 31
	ExdDataUnitsHours                          ExdDataUnits = 32
	ExdDataUnitsCalories                       ExdDataUnits = 33
	ExdDataUnitsKilojoules                     ExdDataUnits = 34
	ExdDataUnitsMilliseconds                   ExdDataUnits = 35
	ExdDataUnitsSecondPerMile                  ExdDataUnits = 36
	ExdDataUnitsSecondPerKilometer             ExdDataUnits = 37
	ExdDataUnitsCentimeter                     ExdDataUnits = 38
	ExdDataUnitsEnumCoursePoint                ExdDataUnits = 39
	ExdDataUnitsBradians                       ExdDataUnits = 40
	ExdDataUnitsEnumSport                      ExdDataUnits = 41
	ExdDataUnitsInchesHg                       ExdDataUnits = 42
	ExdDataUnitsMmHg                           ExdDataUnits = 43
	ExdDataUnitsMbars                          ExdDataUnits = 44
	ExdDataUnitsHectoPascals                   ExdDataUnits = 45
	ExdDataUnitsFeetPerMin                     ExdDataUnits = 46
	ExdDataUnitsMetersPerMin                   ExdDataUnits = 47
	ExdDataUnitsMetersPerSec                   ExdDataUnits = 48
	ExdDataUnitsEightCardinal                  ExdDataUnits = 49
	ExdDataUnitsInvalid                        ExdDataUnits = 0xFF // INVALID
)

var exddataunitstostrs = map[ExdDataUnits]string{
	ExdDataUnitsNoUnits:                        "no_units",
	ExdDataUnitsLaps:                           "laps",
	ExdDataUnitsMilesPerHour:                   "miles_per_hour",
	ExdDataUnitsKilometersPerHour:              "kilometers_per_hour",
	ExdDataUnitsFeetPerHour:                    "feet_per_hour",
	ExdDataUnitsMetersPerHour:                  "meters_per_hour",
	ExdDataUnitsDegreesCelsius:                 "degrees_celsius",
	ExdDataUnitsDegreesFarenheit:               "degrees_farenheit",
	ExdDataUnitsZone:                           "zone",
	ExdDataUnitsGear:                           "gear",
	ExdDataUnitsRpm:                            "rpm",
	ExdDataUnitsBpm:                            "bpm",
	ExdDataUnitsDegrees:                        "degrees",
	ExdDataUnitsMillimeters:                    "millimeters",
	ExdDataUnitsMeters:                         "meters",
	ExdDataUnitsKilometers:                     "kilometers",
	ExdDataUnitsFeet:                           "feet",
	ExdDataUnitsYards:                          "yards",
	ExdDataUnitsKilofeet:                       "kilofeet",
	ExdDataUnitsMiles:                          "miles",
	ExdDataUnitsTime:                           "time",
	ExdDataUnitsEnumTurnType:                   "enum_turn_type",
	ExdDataUnitsPercent:                        "percent",
	ExdDataUnitsWatts:                          "watts",
	ExdDataUnitsWattsPerKilogram:               "watts_per_kilogram",
	ExdDataUnitsEnumBatteryStatus:              "enum_battery_status",
	ExdDataUnitsEnumBikeLightBeamAngleMode:     "enum_bike_light_beam_angle_mode",
	ExdDataUnitsEnumBikeLightBatteryStatus:     "enum_bike_light_battery_status",
	ExdDataUnitsEnumBikeLightNetworkConfigType: "enum_bike_light_network_config_type",
	ExdDataUnitsLights:                         "lights",
	ExdDataUnitsSeconds:                        "seconds",
	ExdDataUnitsMinutes:                        "minutes",
	ExdDataUnitsHours:                          "hours",
	ExdDataUnitsCalories:                       "calories",
	ExdDataUnitsKilojoules:                     "kilojoules",
	ExdDataUnitsMilliseconds:                   "milliseconds",
	ExdDataUnitsSecondPerMile:                  "second_per_mile",
	ExdDataUnitsSecondPerKilometer:             "second_per_kilometer",
	ExdDataUnitsCentimeter:                     "centimeter",
	ExdDataUnitsEnumCoursePoint:                "enum_course_point",
	ExdDataUnitsBradians:                       "bradians",
	ExdDataUnitsEnumSport:                      "enum_sport",
	ExdDataUnitsInchesHg:                       "inches_hg",
	ExdDataUnitsMmHg:                           "mm_hg",
	ExdDataUnitsMbars:                          "mbars",
	ExdDataUnitsHectoPascals:                   "hecto_pascals",
	ExdDataUnitsFeetPerMin:                     "feet_per_min",
	ExdDataUnitsMetersPerMin:                   "meters_per_min",
	ExdDataUnitsMetersPerSec:                   "meters_per_sec",
	ExdDataUnitsEightCardinal:                  "eight_cardinal",
	ExdDataUnitsInvalid:                        "invalid",
}

func (e ExdDataUnits) String() string {
	val, ok := exddataunitstostrs[e]
	if !ok {
		return strconv.Itoa(int(e))
	}
	return val
}

var strtoexddataunits = func() map[string]ExdDataUnits {
	m := make(map[string]ExdDataUnits)
	for t, str := range exddataunitstostrs {
		m[str] = ExdDataUnits(t)
	}
	return m
}()

// FromString parse string into ExdDataUnits constant it's represent, return ExdDataUnitsInvalid if not found.
func ExdDataUnitsFromString(s string) ExdDataUnits {
	val, ok := strtoexddataunits[s]
	if !ok {
		return strtoexddataunits["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListExdDataUnits() []ExdDataUnits {
	vs := make([]ExdDataUnits, 0, len(exddataunitstostrs))
	for i := range exddataunitstostrs {
		vs = append(vs, ExdDataUnits(i))
	}
	return vs
}
