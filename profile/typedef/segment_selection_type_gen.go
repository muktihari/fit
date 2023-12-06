// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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
	SegmentSelectionTypeInvalid   SegmentSelectionType = 0xFF // INVALID
)

var segmentselectiontypetostrs = map[SegmentSelectionType]string{
	SegmentSelectionTypeStarred:   "starred",
	SegmentSelectionTypeSuggested: "suggested",
	SegmentSelectionTypeInvalid:   "invalid",
}

func (s SegmentSelectionType) String() string {
	val, ok := segmentselectiontypetostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosegmentselectiontype = func() map[string]SegmentSelectionType {
	m := make(map[string]SegmentSelectionType)
	for t, str := range segmentselectiontypetostrs {
		m[str] = SegmentSelectionType(t)
	}
	return m
}()

// FromString parse string into SegmentSelectionType constant it's represent, return SegmentSelectionTypeInvalid if not found.
func SegmentSelectionTypeFromString(s string) SegmentSelectionType {
	val, ok := strtosegmentselectiontype[s]
	if !ok {
		return strtosegmentselectiontype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSegmentSelectionType() []SegmentSelectionType {
	vs := make([]SegmentSelectionType, 0, len(segmentselectiontypetostrs))
	for i := range segmentselectiontypetostrs {
		vs = append(vs, SegmentSelectionType(i))
	}
	return vs
}
