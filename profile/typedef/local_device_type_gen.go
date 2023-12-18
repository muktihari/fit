// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LocalDeviceType uint8

const (
	LocalDeviceTypeGps           LocalDeviceType = 0  // Onboard gps receiver
	LocalDeviceTypeGlonass       LocalDeviceType = 1  // Onboard glonass receiver
	LocalDeviceTypeGpsGlonass    LocalDeviceType = 2  // Onboard gps glonass receiver
	LocalDeviceTypeAccelerometer LocalDeviceType = 3  // Onboard sensor
	LocalDeviceTypeBarometer     LocalDeviceType = 4  // Onboard sensor
	LocalDeviceTypeTemperature   LocalDeviceType = 5  // Onboard sensor
	LocalDeviceTypeWhr           LocalDeviceType = 10 // Onboard wrist HR sensor
	LocalDeviceTypeSensorHub     LocalDeviceType = 12 // Onboard software package
	LocalDeviceTypeInvalid       LocalDeviceType = 0xFF
)

func (l LocalDeviceType) String() string {
	switch l {
	case LocalDeviceTypeGps:
		return "gps"
	case LocalDeviceTypeGlonass:
		return "glonass"
	case LocalDeviceTypeGpsGlonass:
		return "gps_glonass"
	case LocalDeviceTypeAccelerometer:
		return "accelerometer"
	case LocalDeviceTypeBarometer:
		return "barometer"
	case LocalDeviceTypeTemperature:
		return "temperature"
	case LocalDeviceTypeWhr:
		return "whr"
	case LocalDeviceTypeSensorHub:
		return "sensor_hub"
	default:
		return "LocalDeviceTypeInvalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LocalDeviceType constant it's represent, return LocalDeviceTypeInvalid if not found.
func LocalDeviceTypeFromString(s string) LocalDeviceType {
	switch s {
	case "gps":
		return LocalDeviceTypeGps
	case "glonass":
		return LocalDeviceTypeGlonass
	case "gps_glonass":
		return LocalDeviceTypeGpsGlonass
	case "accelerometer":
		return LocalDeviceTypeAccelerometer
	case "barometer":
		return LocalDeviceTypeBarometer
	case "temperature":
		return LocalDeviceTypeTemperature
	case "whr":
		return LocalDeviceTypeWhr
	case "sensor_hub":
		return LocalDeviceTypeSensorHub
	default:
		return LocalDeviceTypeInvalid
	}
}

// List returns all constants.
func ListLocalDeviceType() []LocalDeviceType {
	return []LocalDeviceType{
		LocalDeviceTypeGps,
		LocalDeviceTypeGlonass,
		LocalDeviceTypeGpsGlonass,
		LocalDeviceTypeAccelerometer,
		LocalDeviceTypeBarometer,
		LocalDeviceTypeTemperature,
		LocalDeviceTypeWhr,
		LocalDeviceTypeSensorHub,
	}
}
