// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SupportedExdScreenLayouts uint32 // Base: uint32z

const (
	SupportedExdScreenLayoutsFullScreen                SupportedExdScreenLayouts = 0x00000001
	SupportedExdScreenLayoutsHalfVertical              SupportedExdScreenLayouts = 0x00000002
	SupportedExdScreenLayoutsHalfHorizontal            SupportedExdScreenLayouts = 0x00000004
	SupportedExdScreenLayoutsHalfVerticalRightSplit    SupportedExdScreenLayouts = 0x00000008
	SupportedExdScreenLayoutsHalfHorizontalBottomSplit SupportedExdScreenLayouts = 0x00000010
	SupportedExdScreenLayoutsFullQuarterSplit          SupportedExdScreenLayouts = 0x00000020
	SupportedExdScreenLayoutsHalfVerticalLeftSplit     SupportedExdScreenLayouts = 0x00000040
	SupportedExdScreenLayoutsHalfHorizontalTopSplit    SupportedExdScreenLayouts = 0x00000080
	SupportedExdScreenLayoutsInvalid                   SupportedExdScreenLayouts = 0x0
)

func (s SupportedExdScreenLayouts) Uint32() uint32 { return uint32(s) }

func (s SupportedExdScreenLayouts) String() string {
	switch s {
	case SupportedExdScreenLayoutsFullScreen:
		return "full_screen"
	case SupportedExdScreenLayoutsHalfVertical:
		return "half_vertical"
	case SupportedExdScreenLayoutsHalfHorizontal:
		return "half_horizontal"
	case SupportedExdScreenLayoutsHalfVerticalRightSplit:
		return "half_vertical_right_split"
	case SupportedExdScreenLayoutsHalfHorizontalBottomSplit:
		return "half_horizontal_bottom_split"
	case SupportedExdScreenLayoutsFullQuarterSplit:
		return "full_quarter_split"
	case SupportedExdScreenLayoutsHalfVerticalLeftSplit:
		return "half_vertical_left_split"
	case SupportedExdScreenLayoutsHalfHorizontalTopSplit:
		return "half_horizontal_top_split"
	default:
		return "SupportedExdScreenLayoutsInvalid(" + strconv.FormatUint(uint64(s), 10) + ")"
	}
}

// FromString parse string into SupportedExdScreenLayouts constant it's represent, return SupportedExdScreenLayoutsInvalid if not found.
func SupportedExdScreenLayoutsFromString(s string) SupportedExdScreenLayouts {
	switch s {
	case "full_screen":
		return SupportedExdScreenLayoutsFullScreen
	case "half_vertical":
		return SupportedExdScreenLayoutsHalfVertical
	case "half_horizontal":
		return SupportedExdScreenLayoutsHalfHorizontal
	case "half_vertical_right_split":
		return SupportedExdScreenLayoutsHalfVerticalRightSplit
	case "half_horizontal_bottom_split":
		return SupportedExdScreenLayoutsHalfHorizontalBottomSplit
	case "full_quarter_split":
		return SupportedExdScreenLayoutsFullQuarterSplit
	case "half_vertical_left_split":
		return SupportedExdScreenLayoutsHalfVerticalLeftSplit
	case "half_horizontal_top_split":
		return SupportedExdScreenLayoutsHalfHorizontalTopSplit
	default:
		return SupportedExdScreenLayoutsInvalid
	}
}

// List returns all constants.
func ListSupportedExdScreenLayouts() []SupportedExdScreenLayouts {
	return []SupportedExdScreenLayouts{
		SupportedExdScreenLayoutsFullScreen,
		SupportedExdScreenLayoutsHalfVertical,
		SupportedExdScreenLayoutsHalfHorizontal,
		SupportedExdScreenLayoutsHalfVerticalRightSplit,
		SupportedExdScreenLayoutsHalfHorizontalBottomSplit,
		SupportedExdScreenLayoutsFullQuarterSplit,
		SupportedExdScreenLayoutsHalfVerticalLeftSplit,
		SupportedExdScreenLayoutsHalfHorizontalTopSplit,
	}
}
