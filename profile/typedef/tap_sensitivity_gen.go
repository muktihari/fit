// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TapSensitivity byte

const (
	TapSensitivityHigh    TapSensitivity = 0
	TapSensitivityMedium  TapSensitivity = 1
	TapSensitivityLow     TapSensitivity = 2
	TapSensitivityInvalid TapSensitivity = 0xFF
)

func (t TapSensitivity) Byte() byte { return byte(t) }

func (t TapSensitivity) String() string {
	switch t {
	case TapSensitivityHigh:
		return "high"
	case TapSensitivityMedium:
		return "medium"
	case TapSensitivityLow:
		return "low"
	default:
		return "TapSensitivityInvalid(" + strconv.Itoa(int(t)) + ")"
	}
}

// FromString parse string into TapSensitivity constant it's represent, return TapSensitivityInvalid if not found.
func TapSensitivityFromString(s string) TapSensitivity {
	switch s {
	case "high":
		return TapSensitivityHigh
	case "medium":
		return TapSensitivityMedium
	case "low":
		return TapSensitivityLow
	default:
		return TapSensitivityInvalid
	}
}

// List returns all constants.
func ListTapSensitivity() []TapSensitivity {
	return []TapSensitivity{
		TapSensitivityHigh,
		TapSensitivityMedium,
		TapSensitivityLow,
	}
}
