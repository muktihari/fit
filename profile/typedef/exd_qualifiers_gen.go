// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ExdQualifiers byte

const (
	ExdQualifiersNoQualifier              ExdQualifiers = 0
	ExdQualifiersInstantaneous            ExdQualifiers = 1
	ExdQualifiersAverage                  ExdQualifiers = 2
	ExdQualifiersLap                      ExdQualifiers = 3
	ExdQualifiersMaximum                  ExdQualifiers = 4
	ExdQualifiersMaximumAverage           ExdQualifiers = 5
	ExdQualifiersMaximumLap               ExdQualifiers = 6
	ExdQualifiersLastLap                  ExdQualifiers = 7
	ExdQualifiersAverageLap               ExdQualifiers = 8
	ExdQualifiersToDestination            ExdQualifiers = 9
	ExdQualifiersToGo                     ExdQualifiers = 10
	ExdQualifiersToNext                   ExdQualifiers = 11
	ExdQualifiersNextCoursePoint          ExdQualifiers = 12
	ExdQualifiersTotal                    ExdQualifiers = 13
	ExdQualifiersThreeSecondAverage       ExdQualifiers = 14
	ExdQualifiersTenSecondAverage         ExdQualifiers = 15
	ExdQualifiersThirtySecondAverage      ExdQualifiers = 16
	ExdQualifiersPercentMaximum           ExdQualifiers = 17
	ExdQualifiersPercentMaximumAverage    ExdQualifiers = 18
	ExdQualifiersLapPercentMaximum        ExdQualifiers = 19
	ExdQualifiersElapsed                  ExdQualifiers = 20
	ExdQualifiersSunrise                  ExdQualifiers = 21
	ExdQualifiersSunset                   ExdQualifiers = 22
	ExdQualifiersComparedToVirtualPartner ExdQualifiers = 23
	ExdQualifiersMaximum24H               ExdQualifiers = 24
	ExdQualifiersMinimum24H               ExdQualifiers = 25
	ExdQualifiersMinimum                  ExdQualifiers = 26
	ExdQualifiersFirst                    ExdQualifiers = 27
	ExdQualifiersSecond                   ExdQualifiers = 28
	ExdQualifiersThird                    ExdQualifiers = 29
	ExdQualifiersShifter                  ExdQualifiers = 30
	ExdQualifiersLastSport                ExdQualifiers = 31
	ExdQualifiersMoving                   ExdQualifiers = 32
	ExdQualifiersStopped                  ExdQualifiers = 33
	ExdQualifiersEstimatedTotal           ExdQualifiers = 34
	ExdQualifiersZone9                    ExdQualifiers = 242
	ExdQualifiersZone8                    ExdQualifiers = 243
	ExdQualifiersZone7                    ExdQualifiers = 244
	ExdQualifiersZone6                    ExdQualifiers = 245
	ExdQualifiersZone5                    ExdQualifiers = 246
	ExdQualifiersZone4                    ExdQualifiers = 247
	ExdQualifiersZone3                    ExdQualifiers = 248
	ExdQualifiersZone2                    ExdQualifiers = 249
	ExdQualifiersZone1                    ExdQualifiers = 250
	ExdQualifiersInvalid                  ExdQualifiers = 0xFF // INVALID
)

var exdqualifierstostrs = map[ExdQualifiers]string{
	ExdQualifiersNoQualifier:              "no_qualifier",
	ExdQualifiersInstantaneous:            "instantaneous",
	ExdQualifiersAverage:                  "average",
	ExdQualifiersLap:                      "lap",
	ExdQualifiersMaximum:                  "maximum",
	ExdQualifiersMaximumAverage:           "maximum_average",
	ExdQualifiersMaximumLap:               "maximum_lap",
	ExdQualifiersLastLap:                  "last_lap",
	ExdQualifiersAverageLap:               "average_lap",
	ExdQualifiersToDestination:            "to_destination",
	ExdQualifiersToGo:                     "to_go",
	ExdQualifiersToNext:                   "to_next",
	ExdQualifiersNextCoursePoint:          "next_course_point",
	ExdQualifiersTotal:                    "total",
	ExdQualifiersThreeSecondAverage:       "three_second_average",
	ExdQualifiersTenSecondAverage:         "ten_second_average",
	ExdQualifiersThirtySecondAverage:      "thirty_second_average",
	ExdQualifiersPercentMaximum:           "percent_maximum",
	ExdQualifiersPercentMaximumAverage:    "percent_maximum_average",
	ExdQualifiersLapPercentMaximum:        "lap_percent_maximum",
	ExdQualifiersElapsed:                  "elapsed",
	ExdQualifiersSunrise:                  "sunrise",
	ExdQualifiersSunset:                   "sunset",
	ExdQualifiersComparedToVirtualPartner: "compared_to_virtual_partner",
	ExdQualifiersMaximum24H:               "maximum_24h",
	ExdQualifiersMinimum24H:               "minimum_24h",
	ExdQualifiersMinimum:                  "minimum",
	ExdQualifiersFirst:                    "first",
	ExdQualifiersSecond:                   "second",
	ExdQualifiersThird:                    "third",
	ExdQualifiersShifter:                  "shifter",
	ExdQualifiersLastSport:                "last_sport",
	ExdQualifiersMoving:                   "moving",
	ExdQualifiersStopped:                  "stopped",
	ExdQualifiersEstimatedTotal:           "estimated_total",
	ExdQualifiersZone9:                    "zone_9",
	ExdQualifiersZone8:                    "zone_8",
	ExdQualifiersZone7:                    "zone_7",
	ExdQualifiersZone6:                    "zone_6",
	ExdQualifiersZone5:                    "zone_5",
	ExdQualifiersZone4:                    "zone_4",
	ExdQualifiersZone3:                    "zone_3",
	ExdQualifiersZone2:                    "zone_2",
	ExdQualifiersZone1:                    "zone_1",
	ExdQualifiersInvalid:                  "invalid",
}

func (e ExdQualifiers) String() string {
	val, ok := exdqualifierstostrs[e]
	if !ok {
		return strconv.Itoa(int(e))
	}
	return val
}

var strtoexdqualifiers = func() map[string]ExdQualifiers {
	m := make(map[string]ExdQualifiers)
	for t, str := range exdqualifierstostrs {
		m[str] = ExdQualifiers(t)
	}
	return m
}()

// FromString parse string into ExdQualifiers constant it's represent, return ExdQualifiersInvalid if not found.
func ExdQualifiersFromString(s string) ExdQualifiers {
	val, ok := strtoexdqualifiers[s]
	if !ok {
		return strtoexdqualifiers["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListExdQualifiers() []ExdQualifiers {
	vs := make([]ExdQualifiers, 0, len(exdqualifierstostrs))
	for i := range exdqualifierstostrs {
		vs = append(vs, ExdQualifiers(i))
	}
	return vs
}
