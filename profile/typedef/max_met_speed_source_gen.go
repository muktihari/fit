// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type MaxMetSpeedSource byte

const (
	MaxMetSpeedSourceOnboardGps   MaxMetSpeedSource = 0
	MaxMetSpeedSourceConnectedGps MaxMetSpeedSource = 1
	MaxMetSpeedSourceCadence      MaxMetSpeedSource = 2
	MaxMetSpeedSourceInvalid      MaxMetSpeedSource = 0xFF
)

func (m MaxMetSpeedSource) Byte() byte { return byte(m) }

func (m MaxMetSpeedSource) String() string {
	switch m {
	case MaxMetSpeedSourceOnboardGps:
		return "onboard_gps"
	case MaxMetSpeedSourceConnectedGps:
		return "connected_gps"
	case MaxMetSpeedSourceCadence:
		return "cadence"
	default:
		return "MaxMetSpeedSourceInvalid(" + strconv.Itoa(int(m)) + ")"
	}
}

// FromString parse string into MaxMetSpeedSource constant it's represent, return MaxMetSpeedSourceInvalid if not found.
func MaxMetSpeedSourceFromString(s string) MaxMetSpeedSource {
	switch s {
	case "onboard_gps":
		return MaxMetSpeedSourceOnboardGps
	case "connected_gps":
		return MaxMetSpeedSourceConnectedGps
	case "cadence":
		return MaxMetSpeedSourceCadence
	default:
		return MaxMetSpeedSourceInvalid
	}
}

// List returns all constants.
func ListMaxMetSpeedSource() []MaxMetSpeedSource {
	return []MaxMetSpeedSource{
		MaxMetSpeedSourceOnboardGps,
		MaxMetSpeedSourceConnectedGps,
		MaxMetSpeedSourceCadence,
	}
}
