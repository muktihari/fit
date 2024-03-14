// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DisplayPower byte

const (
	DisplayPowerWatts      DisplayPower = 0
	DisplayPowerPercentFtp DisplayPower = 1
	DisplayPowerInvalid    DisplayPower = 0xFF
)

func (d DisplayPower) String() string {
	switch d {
	case DisplayPowerWatts:
		return "watts"
	case DisplayPowerPercentFtp:
		return "percent_ftp"
	default:
		return "DisplayPowerInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DisplayPower constant it's represent, return DisplayPowerInvalid if not found.
func DisplayPowerFromString(s string) DisplayPower {
	switch s {
	case "watts":
		return DisplayPowerWatts
	case "percent_ftp":
		return DisplayPowerPercentFtp
	default:
		return DisplayPowerInvalid
	}
}

// List returns all constants.
func ListDisplayPower() []DisplayPower {
	return []DisplayPower{
		DisplayPowerWatts,
		DisplayPowerPercentFtp,
	}
}
