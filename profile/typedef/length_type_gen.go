// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LengthType byte

const (
	LengthTypeIdle    LengthType = 0    // Rest period. Length with no strokes
	LengthTypeActive  LengthType = 1    // Length with strokes.
	LengthTypeInvalid LengthType = 0xFF // INVALID
)

var lengthtypetostrs = map[LengthType]string{
	LengthTypeIdle:    "idle",
	LengthTypeActive:  "active",
	LengthTypeInvalid: "invalid",
}

func (l LengthType) String() string {
	val, ok := lengthtypetostrs[l]
	if !ok {
		return strconv.Itoa(int(l))
	}
	return val
}

var strtolengthtype = func() map[string]LengthType {
	m := make(map[string]LengthType)
	for t, str := range lengthtypetostrs {
		m[str] = LengthType(t)
	}
	return m
}()

// FromString parse string into LengthType constant it's represent, return LengthTypeInvalid if not found.
func LengthTypeFromString(s string) LengthType {
	val, ok := strtolengthtype[s]
	if !ok {
		return strtolengthtype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLengthType() []LengthType {
	vs := make([]LengthType, 0, len(lengthtypetostrs))
	for i := range lengthtypetostrs {
		vs = append(vs, LengthType(i))
	}
	return vs
}
