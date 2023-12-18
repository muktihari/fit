// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WatchfaceMode byte

const (
	WatchfaceModeDigital   WatchfaceMode = 0
	WatchfaceModeAnalog    WatchfaceMode = 1
	WatchfaceModeConnectIq WatchfaceMode = 2
	WatchfaceModeDisabled  WatchfaceMode = 3
	WatchfaceModeInvalid   WatchfaceMode = 0xFF
)

func (w WatchfaceMode) String() string {
	switch w {
	case WatchfaceModeDigital:
		return "digital"
	case WatchfaceModeAnalog:
		return "analog"
	case WatchfaceModeConnectIq:
		return "connect_iq"
	case WatchfaceModeDisabled:
		return "disabled"
	default:
		return "WatchfaceModeInvalid(" + strconv.Itoa(int(w)) + ")"
	}
}

// FromString parse string into WatchfaceMode constant it's represent, return WatchfaceModeInvalid if not found.
func WatchfaceModeFromString(s string) WatchfaceMode {
	switch s {
	case "digital":
		return WatchfaceModeDigital
	case "analog":
		return WatchfaceModeAnalog
	case "connect_iq":
		return WatchfaceModeConnectIq
	case "disabled":
		return WatchfaceModeDisabled
	default:
		return WatchfaceModeInvalid
	}
}

// List returns all constants.
func ListWatchfaceMode() []WatchfaceMode {
	return []WatchfaceMode{
		WatchfaceModeDigital,
		WatchfaceModeAnalog,
		WatchfaceModeConnectIq,
		WatchfaceModeDisabled,
	}
}
