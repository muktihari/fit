// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type MaxMetHeartRateSource byte

const (
	MaxMetHeartRateSourceWhr     MaxMetHeartRateSource = 0 // Wrist Heart Rate Monitor
	MaxMetHeartRateSourceHrm     MaxMetHeartRateSource = 1 // Chest Strap Heart Rate Monitor
	MaxMetHeartRateSourceInvalid MaxMetHeartRateSource = 0xFF
)

func (m MaxMetHeartRateSource) Byte() byte { return byte(m) }

func (m MaxMetHeartRateSource) String() string {
	switch m {
	case MaxMetHeartRateSourceWhr:
		return "whr"
	case MaxMetHeartRateSourceHrm:
		return "hrm"
	default:
		return "MaxMetHeartRateSourceInvalid(" + strconv.Itoa(int(m)) + ")"
	}
}

// FromString parse string into MaxMetHeartRateSource constant it's represent, return MaxMetHeartRateSourceInvalid if not found.
func MaxMetHeartRateSourceFromString(s string) MaxMetHeartRateSource {
	switch s {
	case "whr":
		return MaxMetHeartRateSourceWhr
	case "hrm":
		return MaxMetHeartRateSourceHrm
	default:
		return MaxMetHeartRateSourceInvalid
	}
}

// List returns all constants.
func ListMaxMetHeartRateSource() []MaxMetHeartRateSource {
	return []MaxMetHeartRateSource{
		MaxMetHeartRateSourceWhr,
		MaxMetHeartRateSourceHrm,
	}
}
