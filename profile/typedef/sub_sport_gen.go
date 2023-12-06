// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SubSport byte

const (
	SubSportGeneric              SubSport = 0
	SubSportTreadmill            SubSport = 1  // Run/Fitness Equipment
	SubSportStreet               SubSport = 2  // Run
	SubSportTrail                SubSport = 3  // Run
	SubSportTrack                SubSport = 4  // Run
	SubSportSpin                 SubSport = 5  // Cycling
	SubSportIndoorCycling        SubSport = 6  // Cycling/Fitness Equipment
	SubSportRoad                 SubSport = 7  // Cycling
	SubSportMountain             SubSport = 8  // Cycling
	SubSportDownhill             SubSport = 9  // Cycling
	SubSportRecumbent            SubSport = 10 // Cycling
	SubSportCyclocross           SubSport = 11 // Cycling
	SubSportHandCycling          SubSport = 12 // Cycling
	SubSportTrackCycling         SubSport = 13 // Cycling
	SubSportIndoorRowing         SubSport = 14 // Fitness Equipment
	SubSportElliptical           SubSport = 15 // Fitness Equipment
	SubSportStairClimbing        SubSport = 16 // Fitness Equipment
	SubSportLapSwimming          SubSport = 17 // Swimming
	SubSportOpenWater            SubSport = 18 // Swimming
	SubSportFlexibilityTraining  SubSport = 19 // Training
	SubSportStrengthTraining     SubSport = 20 // Training
	SubSportWarmUp               SubSport = 21 // Tennis
	SubSportMatch                SubSport = 22 // Tennis
	SubSportExercise             SubSport = 23 // Tennis
	SubSportChallenge            SubSport = 24
	SubSportIndoorSkiing         SubSport = 25 // Fitness Equipment
	SubSportCardioTraining       SubSport = 26 // Training
	SubSportIndoorWalking        SubSport = 27 // Walking/Fitness Equipment
	SubSportEBikeFitness         SubSport = 28 // E-Biking
	SubSportBmx                  SubSport = 29 // Cycling
	SubSportCasualWalking        SubSport = 30 // Walking
	SubSportSpeedWalking         SubSport = 31 // Walking
	SubSportBikeToRunTransition  SubSport = 32 // Transition
	SubSportRunToBikeTransition  SubSport = 33 // Transition
	SubSportSwimToBikeTransition SubSport = 34 // Transition
	SubSportAtv                  SubSport = 35 // Motorcycling
	SubSportMotocross            SubSport = 36 // Motorcycling
	SubSportBackcountry          SubSport = 37 // Alpine Skiing/Snowboarding
	SubSportResort               SubSport = 38 // Alpine Skiing/Snowboarding
	SubSportRcDrone              SubSport = 39 // Flying
	SubSportWingsuit             SubSport = 40 // Flying
	SubSportWhitewater           SubSport = 41 // Kayaking/Rafting
	SubSportSkateSkiing          SubSport = 42 // Cross Country Skiing
	SubSportYoga                 SubSport = 43 // Training
	SubSportPilates              SubSport = 44 // Fitness Equipment
	SubSportIndoorRunning        SubSport = 45 // Run
	SubSportGravelCycling        SubSport = 46 // Cycling
	SubSportEBikeMountain        SubSport = 47 // Cycling
	SubSportCommuting            SubSport = 48 // Cycling
	SubSportMixedSurface         SubSport = 49 // Cycling
	SubSportNavigate             SubSport = 50
	SubSportTrackMe              SubSport = 51
	SubSportMap                  SubSport = 52
	SubSportSingleGasDiving      SubSport = 53 // Diving
	SubSportMultiGasDiving       SubSport = 54 // Diving
	SubSportGaugeDiving          SubSport = 55 // Diving
	SubSportApneaDiving          SubSport = 56 // Diving
	SubSportApneaHunting         SubSport = 57 // Diving
	SubSportVirtualActivity      SubSport = 58
	SubSportObstacle             SubSport = 59 // Used for events where participants run, crawl through mud, climb over walls, etc.
	SubSportBreathing            SubSport = 62
	SubSportSailRace             SubSport = 65  // Sailing
	SubSportUltra                SubSport = 67  // Ultramarathon
	SubSportIndoorClimbing       SubSport = 68  // Climbing
	SubSportBouldering           SubSport = 69  // Climbing
	SubSportHiit                 SubSport = 70  // High Intensity Interval Training
	SubSportAmrap                SubSport = 73  // HIIT
	SubSportEmom                 SubSport = 74  // HIIT
	SubSportTabata               SubSport = 75  // HIIT
	SubSportPickleball           SubSport = 84  // Racket
	SubSportPadel                SubSport = 85  // Racket
	SubSportFlyCanopy            SubSport = 110 // Flying
	SubSportFlyParaglide         SubSport = 111 // Flying
	SubSportFlyParamotor         SubSport = 112 // Flying
	SubSportFlyPressurized       SubSport = 113 // Flying
	SubSportFlyNavigate          SubSport = 114 // Flying
	SubSportFlyTimer             SubSport = 115 // Flying
	SubSportFlyAltimeter         SubSport = 116 // Flying
	SubSportFlyWx                SubSport = 117 // Flying
	SubSportFlyVfr               SubSport = 118 // Flying
	SubSportFlyIfr               SubSport = 119 // Flying
	SubSportAll                  SubSport = 254
	SubSportInvalid              SubSport = 0xFF // INVALID
)

var subsporttostrs = map[SubSport]string{
	SubSportGeneric:              "generic",
	SubSportTreadmill:            "treadmill",
	SubSportStreet:               "street",
	SubSportTrail:                "trail",
	SubSportTrack:                "track",
	SubSportSpin:                 "spin",
	SubSportIndoorCycling:        "indoor_cycling",
	SubSportRoad:                 "road",
	SubSportMountain:             "mountain",
	SubSportDownhill:             "downhill",
	SubSportRecumbent:            "recumbent",
	SubSportCyclocross:           "cyclocross",
	SubSportHandCycling:          "hand_cycling",
	SubSportTrackCycling:         "track_cycling",
	SubSportIndoorRowing:         "indoor_rowing",
	SubSportElliptical:           "elliptical",
	SubSportStairClimbing:        "stair_climbing",
	SubSportLapSwimming:          "lap_swimming",
	SubSportOpenWater:            "open_water",
	SubSportFlexibilityTraining:  "flexibility_training",
	SubSportStrengthTraining:     "strength_training",
	SubSportWarmUp:               "warm_up",
	SubSportMatch:                "match",
	SubSportExercise:             "exercise",
	SubSportChallenge:            "challenge",
	SubSportIndoorSkiing:         "indoor_skiing",
	SubSportCardioTraining:       "cardio_training",
	SubSportIndoorWalking:        "indoor_walking",
	SubSportEBikeFitness:         "e_bike_fitness",
	SubSportBmx:                  "bmx",
	SubSportCasualWalking:        "casual_walking",
	SubSportSpeedWalking:         "speed_walking",
	SubSportBikeToRunTransition:  "bike_to_run_transition",
	SubSportRunToBikeTransition:  "run_to_bike_transition",
	SubSportSwimToBikeTransition: "swim_to_bike_transition",
	SubSportAtv:                  "atv",
	SubSportMotocross:            "motocross",
	SubSportBackcountry:          "backcountry",
	SubSportResort:               "resort",
	SubSportRcDrone:              "rc_drone",
	SubSportWingsuit:             "wingsuit",
	SubSportWhitewater:           "whitewater",
	SubSportSkateSkiing:          "skate_skiing",
	SubSportYoga:                 "yoga",
	SubSportPilates:              "pilates",
	SubSportIndoorRunning:        "indoor_running",
	SubSportGravelCycling:        "gravel_cycling",
	SubSportEBikeMountain:        "e_bike_mountain",
	SubSportCommuting:            "commuting",
	SubSportMixedSurface:         "mixed_surface",
	SubSportNavigate:             "navigate",
	SubSportTrackMe:              "track_me",
	SubSportMap:                  "map",
	SubSportSingleGasDiving:      "single_gas_diving",
	SubSportMultiGasDiving:       "multi_gas_diving",
	SubSportGaugeDiving:          "gauge_diving",
	SubSportApneaDiving:          "apnea_diving",
	SubSportApneaHunting:         "apnea_hunting",
	SubSportVirtualActivity:      "virtual_activity",
	SubSportObstacle:             "obstacle",
	SubSportBreathing:            "breathing",
	SubSportSailRace:             "sail_race",
	SubSportUltra:                "ultra",
	SubSportIndoorClimbing:       "indoor_climbing",
	SubSportBouldering:           "bouldering",
	SubSportHiit:                 "hiit",
	SubSportAmrap:                "amrap",
	SubSportEmom:                 "emom",
	SubSportTabata:               "tabata",
	SubSportPickleball:           "pickleball",
	SubSportPadel:                "padel",
	SubSportFlyCanopy:            "fly_canopy",
	SubSportFlyParaglide:         "fly_paraglide",
	SubSportFlyParamotor:         "fly_paramotor",
	SubSportFlyPressurized:       "fly_pressurized",
	SubSportFlyNavigate:          "fly_navigate",
	SubSportFlyTimer:             "fly_timer",
	SubSportFlyAltimeter:         "fly_altimeter",
	SubSportFlyWx:                "fly_wx",
	SubSportFlyVfr:               "fly_vfr",
	SubSportFlyIfr:               "fly_ifr",
	SubSportAll:                  "all",
	SubSportInvalid:              "invalid",
}

func (s SubSport) String() string {
	val, ok := subsporttostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosubsport = func() map[string]SubSport {
	m := make(map[string]SubSport)
	for t, str := range subsporttostrs {
		m[str] = SubSport(t)
	}
	return m
}()

// FromString parse string into SubSport constant it's represent, return SubSportInvalid if not found.
func SubSportFromString(s string) SubSport {
	val, ok := strtosubsport[s]
	if !ok {
		return strtosubsport["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSubSport() []SubSport {
	vs := make([]SubSport, 0, len(subsporttostrs))
	for i := range subsporttostrs {
		vs = append(vs, SubSport(i))
	}
	return vs
}
