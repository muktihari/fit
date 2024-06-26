// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SegmentSelectionType byte

const (
	SegmentSelectionTypeStarred   SegmentSelectionType = 0
	SegmentSelectionTypeSuggested SegmentSelectionType = 1
	SegmentSelectionTypeInvalid   SegmentSelectionType = 0xFF
)

func (s SegmentSelectionType) Byte() byte { return byte(s) }

func (s SegmentSelectionType) String() string {
	switch s {
	case SegmentSelectionTypeStarred:
		return "starred"
	case SegmentSelectionTypeSuggested:
		return "suggested"
	default:
		return "SegmentSelectionTypeInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SegmentSelectionType constant it's represent, return SegmentSelectionTypeInvalid if not found.
func SegmentSelectionTypeFromString(s string) SegmentSelectionType {
	switch s {
	case "starred":
		return SegmentSelectionTypeStarred
	case "suggested":
		return SegmentSelectionTypeSuggested
	default:
		return SegmentSelectionTypeInvalid
	}
}

// List returns all constants.
func ListSegmentSelectionType() []SegmentSelectionType {
	return []SegmentSelectionType{
		SegmentSelectionTypeStarred,
		SegmentSelectionTypeSuggested,
	}
}
