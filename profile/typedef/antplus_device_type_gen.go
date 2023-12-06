// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AntplusDeviceType uint8

const (
	AntplusDeviceTypeAntfs                   AntplusDeviceType = 1
	AntplusDeviceTypeBikePower               AntplusDeviceType = 11
	AntplusDeviceTypeEnvironmentSensorLegacy AntplusDeviceType = 12
	AntplusDeviceTypeMultiSportSpeedDistance AntplusDeviceType = 15
	AntplusDeviceTypeControl                 AntplusDeviceType = 16
	AntplusDeviceTypeFitnessEquipment        AntplusDeviceType = 17
	AntplusDeviceTypeBloodPressure           AntplusDeviceType = 18
	AntplusDeviceTypeGeocacheNode            AntplusDeviceType = 19
	AntplusDeviceTypeLightElectricVehicle    AntplusDeviceType = 20
	AntplusDeviceTypeEnvSensor               AntplusDeviceType = 25
	AntplusDeviceTypeRacquet                 AntplusDeviceType = 26
	AntplusDeviceTypeControlHub              AntplusDeviceType = 27
	AntplusDeviceTypeMuscleOxygen            AntplusDeviceType = 31
	AntplusDeviceTypeShifting                AntplusDeviceType = 34
	AntplusDeviceTypeBikeLightMain           AntplusDeviceType = 35
	AntplusDeviceTypeBikeLightShared         AntplusDeviceType = 36
	AntplusDeviceTypeExd                     AntplusDeviceType = 38
	AntplusDeviceTypeBikeRadar               AntplusDeviceType = 40
	AntplusDeviceTypeBikeAero                AntplusDeviceType = 46
	AntplusDeviceTypeWeightScale             AntplusDeviceType = 119
	AntplusDeviceTypeHeartRate               AntplusDeviceType = 120
	AntplusDeviceTypeBikeSpeedCadence        AntplusDeviceType = 121
	AntplusDeviceTypeBikeCadence             AntplusDeviceType = 122
	AntplusDeviceTypeBikeSpeed               AntplusDeviceType = 123
	AntplusDeviceTypeStrideSpeedDistance     AntplusDeviceType = 124
	AntplusDeviceTypeInvalid                 AntplusDeviceType = 0xFF // INVALID
)

var antplusdevicetypetostrs = map[AntplusDeviceType]string{
	AntplusDeviceTypeAntfs:                   "antfs",
	AntplusDeviceTypeBikePower:               "bike_power",
	AntplusDeviceTypeEnvironmentSensorLegacy: "environment_sensor_legacy",
	AntplusDeviceTypeMultiSportSpeedDistance: "multi_sport_speed_distance",
	AntplusDeviceTypeControl:                 "control",
	AntplusDeviceTypeFitnessEquipment:        "fitness_equipment",
	AntplusDeviceTypeBloodPressure:           "blood_pressure",
	AntplusDeviceTypeGeocacheNode:            "geocache_node",
	AntplusDeviceTypeLightElectricVehicle:    "light_electric_vehicle",
	AntplusDeviceTypeEnvSensor:               "env_sensor",
	AntplusDeviceTypeRacquet:                 "racquet",
	AntplusDeviceTypeControlHub:              "control_hub",
	AntplusDeviceTypeMuscleOxygen:            "muscle_oxygen",
	AntplusDeviceTypeShifting:                "shifting",
	AntplusDeviceTypeBikeLightMain:           "bike_light_main",
	AntplusDeviceTypeBikeLightShared:         "bike_light_shared",
	AntplusDeviceTypeExd:                     "exd",
	AntplusDeviceTypeBikeRadar:               "bike_radar",
	AntplusDeviceTypeBikeAero:                "bike_aero",
	AntplusDeviceTypeWeightScale:             "weight_scale",
	AntplusDeviceTypeHeartRate:               "heart_rate",
	AntplusDeviceTypeBikeSpeedCadence:        "bike_speed_cadence",
	AntplusDeviceTypeBikeCadence:             "bike_cadence",
	AntplusDeviceTypeBikeSpeed:               "bike_speed",
	AntplusDeviceTypeStrideSpeedDistance:     "stride_speed_distance",
	AntplusDeviceTypeInvalid:                 "invalid",
}

func (a AntplusDeviceType) String() string {
	val, ok := antplusdevicetypetostrs[a]
	if !ok {
		return strconv.FormatUint(uint64(a), 10)
	}
	return val
}

var strtoantplusdevicetype = func() map[string]AntplusDeviceType {
	m := make(map[string]AntplusDeviceType)
	for t, str := range antplusdevicetypetostrs {
		m[str] = AntplusDeviceType(t)
	}
	return m
}()

// FromString parse string into AntplusDeviceType constant it's represent, return AntplusDeviceTypeInvalid if not found.
func AntplusDeviceTypeFromString(s string) AntplusDeviceType {
	val, ok := strtoantplusdevicetype[s]
	if !ok {
		return strtoantplusdevicetype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListAntplusDeviceType() []AntplusDeviceType {
	vs := make([]AntplusDeviceType, 0, len(antplusdevicetypetostrs))
	for i := range antplusdevicetypetostrs {
		vs = append(vs, AntplusDeviceType(i))
	}
	return vs
}
