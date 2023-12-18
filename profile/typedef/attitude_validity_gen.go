// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AttitudeValidity uint16

const (
	AttitudeValidityTrackAngleHeadingValid AttitudeValidity = 0x0001
	AttitudeValidityPitchValid             AttitudeValidity = 0x0002
	AttitudeValidityRollValid              AttitudeValidity = 0x0004
	AttitudeValidityLateralBodyAccelValid  AttitudeValidity = 0x0008
	AttitudeValidityNormalBodyAccelValid   AttitudeValidity = 0x0010
	AttitudeValidityTurnRateValid          AttitudeValidity = 0x0020
	AttitudeValidityHwFail                 AttitudeValidity = 0x0040
	AttitudeValidityMagInvalid             AttitudeValidity = 0x0080
	AttitudeValidityNoGps                  AttitudeValidity = 0x0100
	AttitudeValidityGpsInvalid             AttitudeValidity = 0x0200
	AttitudeValiditySolutionCoasting       AttitudeValidity = 0x0400
	AttitudeValidityTrueTrackAngle         AttitudeValidity = 0x0800
	AttitudeValidityMagneticHeading        AttitudeValidity = 0x1000
	AttitudeValidityInvalid                AttitudeValidity = 0xFFFF
)

func (a AttitudeValidity) String() string {
	switch a {
	case AttitudeValidityTrackAngleHeadingValid:
		return "track_angle_heading_valid"
	case AttitudeValidityPitchValid:
		return "pitch_valid"
	case AttitudeValidityRollValid:
		return "roll_valid"
	case AttitudeValidityLateralBodyAccelValid:
		return "lateral_body_accel_valid"
	case AttitudeValidityNormalBodyAccelValid:
		return "normal_body_accel_valid"
	case AttitudeValidityTurnRateValid:
		return "turn_rate_valid"
	case AttitudeValidityHwFail:
		return "hw_fail"
	case AttitudeValidityMagInvalid:
		return "mag_invalid"
	case AttitudeValidityNoGps:
		return "no_gps"
	case AttitudeValidityGpsInvalid:
		return "gps_invalid"
	case AttitudeValiditySolutionCoasting:
		return "solution_coasting"
	case AttitudeValidityTrueTrackAngle:
		return "true_track_angle"
	case AttitudeValidityMagneticHeading:
		return "magnetic_heading"
	default:
		return "AttitudeValidityInvalid(" + strconv.FormatUint(uint64(a), 10) + ")"
	}
}

// FromString parse string into AttitudeValidity constant it's represent, return AttitudeValidityInvalid if not found.
func AttitudeValidityFromString(s string) AttitudeValidity {
	switch s {
	case "track_angle_heading_valid":
		return AttitudeValidityTrackAngleHeadingValid
	case "pitch_valid":
		return AttitudeValidityPitchValid
	case "roll_valid":
		return AttitudeValidityRollValid
	case "lateral_body_accel_valid":
		return AttitudeValidityLateralBodyAccelValid
	case "normal_body_accel_valid":
		return AttitudeValidityNormalBodyAccelValid
	case "turn_rate_valid":
		return AttitudeValidityTurnRateValid
	case "hw_fail":
		return AttitudeValidityHwFail
	case "mag_invalid":
		return AttitudeValidityMagInvalid
	case "no_gps":
		return AttitudeValidityNoGps
	case "gps_invalid":
		return AttitudeValidityGpsInvalid
	case "solution_coasting":
		return AttitudeValiditySolutionCoasting
	case "true_track_angle":
		return AttitudeValidityTrueTrackAngle
	case "magnetic_heading":
		return AttitudeValidityMagneticHeading
	default:
		return AttitudeValidityInvalid
	}
}

// List returns all constants.
func ListAttitudeValidity() []AttitudeValidity {
	return []AttitudeValidity{
		AttitudeValidityTrackAngleHeadingValid,
		AttitudeValidityPitchValid,
		AttitudeValidityRollValid,
		AttitudeValidityLateralBodyAccelValid,
		AttitudeValidityNormalBodyAccelValid,
		AttitudeValidityTurnRateValid,
		AttitudeValidityHwFail,
		AttitudeValidityMagInvalid,
		AttitudeValidityNoGps,
		AttitudeValidityGpsInvalid,
		AttitudeValiditySolutionCoasting,
		AttitudeValidityTrueTrackAngle,
		AttitudeValidityMagneticHeading,
	}
}
