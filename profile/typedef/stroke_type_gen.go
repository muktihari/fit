// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type StrokeType byte

const (
	StrokeTypeNoEvent  StrokeType = 0
	StrokeTypeOther    StrokeType = 1 // stroke was detected but cannot be identified
	StrokeTypeServe    StrokeType = 2
	StrokeTypeForehand StrokeType = 3
	StrokeTypeBackhand StrokeType = 4
	StrokeTypeSmash    StrokeType = 5
	StrokeTypeInvalid  StrokeType = 0xFF // INVALID
)

var stroketypetostrs = map[StrokeType]string{
	StrokeTypeNoEvent:  "no_event",
	StrokeTypeOther:    "other",
	StrokeTypeServe:    "serve",
	StrokeTypeForehand: "forehand",
	StrokeTypeBackhand: "backhand",
	StrokeTypeSmash:    "smash",
	StrokeTypeInvalid:  "invalid",
}

func (s StrokeType) String() string {
	val, ok := stroketypetostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtostroketype = func() map[string]StrokeType {
	m := make(map[string]StrokeType)
	for t, str := range stroketypetostrs {
		m[str] = StrokeType(t)
	}
	return m
}()

// FromString parse string into StrokeType constant it's represent, return StrokeTypeInvalid if not found.
func StrokeTypeFromString(s string) StrokeType {
	val, ok := strtostroketype[s]
	if !ok {
		return strtostroketype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListStrokeType() []StrokeType {
	vs := make([]StrokeType, 0, len(stroketypetostrs))
	for i := range stroketypetostrs {
		vs = append(vs, StrokeType(i))
	}
	return vs
}
