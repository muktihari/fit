// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DisplayHeart byte

const (
	DisplayHeartBpm     DisplayHeart = 0
	DisplayHeartMax     DisplayHeart = 1
	DisplayHeartReserve DisplayHeart = 2
	DisplayHeartInvalid DisplayHeart = 0xFF
)

func (d DisplayHeart) String() string {
	switch d {
	case DisplayHeartBpm:
		return "bpm"
	case DisplayHeartMax:
		return "max"
	case DisplayHeartReserve:
		return "reserve"
	default:
		return "DisplayHeartInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DisplayHeart constant it's represent, return DisplayHeartInvalid if not found.
func DisplayHeartFromString(s string) DisplayHeart {
	switch s {
	case "bpm":
		return DisplayHeartBpm
	case "max":
		return DisplayHeartMax
	case "reserve":
		return DisplayHeartReserve
	default:
		return DisplayHeartInvalid
	}
}

// List returns all constants.
func ListDisplayHeart() []DisplayHeart {
	return []DisplayHeart{
		DisplayHeartBpm,
		DisplayHeartMax,
		DisplayHeartReserve,
	}
}
