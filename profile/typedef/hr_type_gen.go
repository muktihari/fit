// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type HrType byte

const (
	HrTypeNormal    HrType = 0
	HrTypeIrregular HrType = 1
	HrTypeInvalid   HrType = 0xFF
)

func (h HrType) Byte() byte { return byte(h) }

func (h HrType) String() string {
	switch h {
	case HrTypeNormal:
		return "normal"
	case HrTypeIrregular:
		return "irregular"
	default:
		return "HrTypeInvalid(" + strconv.Itoa(int(h)) + ")"
	}
}

// FromString parse string into HrType constant it's represent, return HrTypeInvalid if not found.
func HrTypeFromString(s string) HrType {
	switch s {
	case "normal":
		return HrTypeNormal
	case "irregular":
		return HrTypeIrregular
	default:
		return HrTypeInvalid
	}
}

// List returns all constants.
func ListHrType() []HrType {
	return []HrType{
		HrTypeNormal,
		HrTypeIrregular,
	}
}
