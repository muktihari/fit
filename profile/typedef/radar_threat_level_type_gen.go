// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type RadarThreatLevelType byte

const (
	RadarThreatLevelTypeThreatUnknown         RadarThreatLevelType = 0
	RadarThreatLevelTypeThreatNone            RadarThreatLevelType = 1
	RadarThreatLevelTypeThreatApproaching     RadarThreatLevelType = 2
	RadarThreatLevelTypeThreatApproachingFast RadarThreatLevelType = 3
	RadarThreatLevelTypeInvalid               RadarThreatLevelType = 0xFF // INVALID
)

var radarthreatleveltypetostrs = map[RadarThreatLevelType]string{
	RadarThreatLevelTypeThreatUnknown:         "threat_unknown",
	RadarThreatLevelTypeThreatNone:            "threat_none",
	RadarThreatLevelTypeThreatApproaching:     "threat_approaching",
	RadarThreatLevelTypeThreatApproachingFast: "threat_approaching_fast",
	RadarThreatLevelTypeInvalid:               "invalid",
}

func (r RadarThreatLevelType) String() string {
	val, ok := radarthreatleveltypetostrs[r]
	if !ok {
		return strconv.Itoa(int(r))
	}
	return val
}

var strtoradarthreatleveltype = func() map[string]RadarThreatLevelType {
	m := make(map[string]RadarThreatLevelType)
	for t, str := range radarthreatleveltypetostrs {
		m[str] = RadarThreatLevelType(t)
	}
	return m
}()

// FromString parse string into RadarThreatLevelType constant it's represent, return RadarThreatLevelTypeInvalid if not found.
func RadarThreatLevelTypeFromString(s string) RadarThreatLevelType {
	val, ok := strtoradarthreatleveltype[s]
	if !ok {
		return strtoradarthreatleveltype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListRadarThreatLevelType() []RadarThreatLevelType {
	vs := make([]RadarThreatLevelType, 0, len(radarthreatleveltypetostrs))
	for i := range radarthreatleveltypetostrs {
		vs = append(vs, RadarThreatLevelType(i))
	}
	return vs
}
