// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SegmentLapStatus byte

const (
	SegmentLapStatusEnd     SegmentLapStatus = 0
	SegmentLapStatusFail    SegmentLapStatus = 1
	SegmentLapStatusInvalid SegmentLapStatus = 0xFF
)

func (s SegmentLapStatus) Byte() byte { return byte(s) }

func (s SegmentLapStatus) String() string {
	switch s {
	case SegmentLapStatusEnd:
		return "end"
	case SegmentLapStatusFail:
		return "fail"
	default:
		return "SegmentLapStatusInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SegmentLapStatus constant it's represent, return SegmentLapStatusInvalid if not found.
func SegmentLapStatusFromString(s string) SegmentLapStatus {
	switch s {
	case "end":
		return SegmentLapStatusEnd
	case "fail":
		return SegmentLapStatusFail
	default:
		return SegmentLapStatusInvalid
	}
}

// List returns all constants.
func ListSegmentLapStatus() []SegmentLapStatus {
	return []SegmentLapStatus{
		SegmentLapStatusEnd,
		SegmentLapStatusFail,
	}
}
