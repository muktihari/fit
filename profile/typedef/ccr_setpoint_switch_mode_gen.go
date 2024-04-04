// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CcrSetpointSwitchMode byte

const (
	CcrSetpointSwitchModeManual    CcrSetpointSwitchMode = 0 // User switches setpoints manually
	CcrSetpointSwitchModeAutomatic CcrSetpointSwitchMode = 1 // Switch automatically based on depth
	CcrSetpointSwitchModeInvalid   CcrSetpointSwitchMode = 0xFF
)

func (c CcrSetpointSwitchMode) Byte() byte { return byte(c) }

func (c CcrSetpointSwitchMode) String() string {
	switch c {
	case CcrSetpointSwitchModeManual:
		return "manual"
	case CcrSetpointSwitchModeAutomatic:
		return "automatic"
	default:
		return "CcrSetpointSwitchModeInvalid(" + strconv.Itoa(int(c)) + ")"
	}
}

// FromString parse string into CcrSetpointSwitchMode constant it's represent, return CcrSetpointSwitchModeInvalid if not found.
func CcrSetpointSwitchModeFromString(s string) CcrSetpointSwitchMode {
	switch s {
	case "manual":
		return CcrSetpointSwitchModeManual
	case "automatic":
		return CcrSetpointSwitchModeAutomatic
	default:
		return CcrSetpointSwitchModeInvalid
	}
}

// List returns all constants.
func ListCcrSetpointSwitchMode() []CcrSetpointSwitchMode {
	return []CcrSetpointSwitchMode{
		CcrSetpointSwitchModeManual,
		CcrSetpointSwitchModeAutomatic,
	}
}
