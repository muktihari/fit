// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type HrvStatus byte

const (
	HrvStatusNone       HrvStatus = 0
	HrvStatusPoor       HrvStatus = 1
	HrvStatusLow        HrvStatus = 2
	HrvStatusUnbalanced HrvStatus = 3
	HrvStatusBalanced   HrvStatus = 4
	HrvStatusInvalid    HrvStatus = 0xFF
)

func (h HrvStatus) String() string {
	switch h {
	case HrvStatusNone:
		return "none"
	case HrvStatusPoor:
		return "poor"
	case HrvStatusLow:
		return "low"
	case HrvStatusUnbalanced:
		return "unbalanced"
	case HrvStatusBalanced:
		return "balanced"
	default:
		return "HrvStatusInvalid(" + strconv.Itoa(int(h)) + ")"
	}
}

// FromString parse string into HrvStatus constant it's represent, return HrvStatusInvalid if not found.
func HrvStatusFromString(s string) HrvStatus {
	switch s {
	case "none":
		return HrvStatusNone
	case "poor":
		return HrvStatusPoor
	case "low":
		return HrvStatusLow
	case "unbalanced":
		return HrvStatusUnbalanced
	case "balanced":
		return HrvStatusBalanced
	default:
		return HrvStatusInvalid
	}
}

// List returns all constants.
func ListHrvStatus() []HrvStatus {
	return []HrvStatus{
		HrvStatusNone,
		HrvStatusPoor,
		HrvStatusLow,
		HrvStatusUnbalanced,
		HrvStatusBalanced,
	}
}
