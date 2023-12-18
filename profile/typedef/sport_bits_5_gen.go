// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits5 uint8

const (
	SportBits5WaterSkiing SportBits5 = 0x01
	SportBits5Kayaking    SportBits5 = 0x02
	SportBits5Rafting     SportBits5 = 0x04
	SportBits5Windsurfing SportBits5 = 0x08
	SportBits5Kitesurfing SportBits5 = 0x10
	SportBits5Tactical    SportBits5 = 0x20
	SportBits5Jumpmaster  SportBits5 = 0x40
	SportBits5Boxing      SportBits5 = 0x80
	SportBits5Invalid     SportBits5 = 0x0
)

func (s SportBits5) String() string {
	switch s {
	case SportBits5WaterSkiing:
		return "water_skiing"
	case SportBits5Kayaking:
		return "kayaking"
	case SportBits5Rafting:
		return "rafting"
	case SportBits5Windsurfing:
		return "windsurfing"
	case SportBits5Kitesurfing:
		return "kitesurfing"
	case SportBits5Tactical:
		return "tactical"
	case SportBits5Jumpmaster:
		return "jumpmaster"
	case SportBits5Boxing:
		return "boxing"
	default:
		return "SportBits5Invalid(" + strconv.FormatUint(uint64(s), 10) + ")"
	}
}

// FromString parse string into SportBits5 constant it's represent, return SportBits5Invalid if not found.
func SportBits5FromString(s string) SportBits5 {
	switch s {
	case "water_skiing":
		return SportBits5WaterSkiing
	case "kayaking":
		return SportBits5Kayaking
	case "rafting":
		return SportBits5Rafting
	case "windsurfing":
		return SportBits5Windsurfing
	case "kitesurfing":
		return SportBits5Kitesurfing
	case "tactical":
		return SportBits5Tactical
	case "jumpmaster":
		return SportBits5Jumpmaster
	case "boxing":
		return SportBits5Boxing
	default:
		return SportBits5Invalid
	}
}

// List returns all constants.
func ListSportBits5() []SportBits5 {
	return []SportBits5{
		SportBits5WaterSkiing,
		SportBits5Kayaking,
		SportBits5Rafting,
		SportBits5Windsurfing,
		SportBits5Kitesurfing,
		SportBits5Tactical,
		SportBits5Jumpmaster,
		SportBits5Boxing,
	}
}
