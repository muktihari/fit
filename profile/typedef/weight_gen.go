// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Weight uint16

const (
	WeightCalculating Weight = 0xFFFE
	WeightInvalid     Weight = 0xFFFF
)

func (w Weight) Uint16() uint16 { return uint16(w) }

func (w Weight) String() string {
	switch w {
	case WeightCalculating:
		return "calculating"
	default:
		return "WeightInvalid(" + strconv.FormatUint(uint64(w), 10) + ")"
	}
}

// FromString parse string into Weight constant it's represent, return WeightInvalid if not found.
func WeightFromString(s string) Weight {
	switch s {
	case "calculating":
		return WeightCalculating
	default:
		return WeightInvalid
	}
}

// List returns all constants.
func ListWeight() []Weight {
	return []Weight{
		WeightCalculating,
	}
}
