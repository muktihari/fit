// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits6 uint8

const (
	SportBits6FloorClimbing SportBits6 = 0x01
	SportBits6Invalid       SportBits6 = 0x0 // INVALID
)

var sportbits6tostrs = map[SportBits6]string{
	SportBits6FloorClimbing: "floor_climbing",
	SportBits6Invalid:       "invalid",
}

func (s SportBits6) String() string {
	val, ok := sportbits6tostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosportbits6 = func() map[string]SportBits6 {
	m := make(map[string]SportBits6)
	for t, str := range sportbits6tostrs {
		m[str] = SportBits6(t)
	}
	return m
}()

// FromString parse string into SportBits6 constant it's represent, return SportBits6Invalid if not found.
func SportBits6FromString(s string) SportBits6 {
	val, ok := strtosportbits6[s]
	if !ok {
		return strtosportbits6["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportBits6() []SportBits6 {
	vs := make([]SportBits6, 0, len(sportbits6tostrs))
	for i := range sportbits6tostrs {
		vs = append(vs, SportBits6(i))
	}
	return vs
}
