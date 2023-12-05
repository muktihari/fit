// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

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
	AttitudeValidityInvalid                AttitudeValidity = 0xFFFF // INVALID
)

var attitudevaliditytostrs = map[AttitudeValidity]string{
	AttitudeValidityTrackAngleHeadingValid: "track_angle_heading_valid",
	AttitudeValidityPitchValid:             "pitch_valid",
	AttitudeValidityRollValid:              "roll_valid",
	AttitudeValidityLateralBodyAccelValid:  "lateral_body_accel_valid",
	AttitudeValidityNormalBodyAccelValid:   "normal_body_accel_valid",
	AttitudeValidityTurnRateValid:          "turn_rate_valid",
	AttitudeValidityHwFail:                 "hw_fail",
	AttitudeValidityMagInvalid:             "mag_invalid",
	AttitudeValidityNoGps:                  "no_gps",
	AttitudeValidityGpsInvalid:             "gps_invalid",
	AttitudeValiditySolutionCoasting:       "solution_coasting",
	AttitudeValidityTrueTrackAngle:         "true_track_angle",
	AttitudeValidityMagneticHeading:        "magnetic_heading",
	AttitudeValidityInvalid:                "invalid",
}

func (a AttitudeValidity) String() string {
	val, ok := attitudevaliditytostrs[a]
	if !ok {
		return strconv.FormatUint(uint64(a), 10)
	}
	return val
}

var strtoattitudevalidity = func() map[string]AttitudeValidity {
	m := make(map[string]AttitudeValidity)
	for t, str := range attitudevaliditytostrs {
		m[str] = AttitudeValidity(t)
	}
	return m
}()

// FromString parse string into AttitudeValidity constant it's represent, return AttitudeValidityInvalid if not found.
func AttitudeValidityFromString(s string) AttitudeValidity {
	val, ok := strtoattitudevalidity[s]
	if !ok {
		return strtoattitudevalidity["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListAttitudeValidity() []AttitudeValidity {
	vs := make([]AttitudeValidity, 0, len(attitudevaliditytostrs))
	for i := range attitudevaliditytostrs {
		vs = append(vs, AttitudeValidity(i))
	}
	return vs
}
