// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits3 uint8

const (
	SportBits3Driving         SportBits3 = 0x01
	SportBits3Golf            SportBits3 = 0x02
	SportBits3HangGliding     SportBits3 = 0x04
	SportBits3HorsebackRiding SportBits3 = 0x08
	SportBits3Hunting         SportBits3 = 0x10
	SportBits3Fishing         SportBits3 = 0x20
	SportBits3InlineSkating   SportBits3 = 0x40
	SportBits3RockClimbing    SportBits3 = 0x80
	SportBits3Invalid         SportBits3 = 0x0 // INVALID
)

var sportbits3tostrs = map[SportBits3]string{
	SportBits3Driving:         "driving",
	SportBits3Golf:            "golf",
	SportBits3HangGliding:     "hang_gliding",
	SportBits3HorsebackRiding: "horseback_riding",
	SportBits3Hunting:         "hunting",
	SportBits3Fishing:         "fishing",
	SportBits3InlineSkating:   "inline_skating",
	SportBits3RockClimbing:    "rock_climbing",
	SportBits3Invalid:         "invalid",
}

func (s SportBits3) String() string {
	val, ok := sportbits3tostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosportbits3 = func() map[string]SportBits3 {
	m := make(map[string]SportBits3)
	for t, str := range sportbits3tostrs {
		m[str] = SportBits3(t)
	}
	return m
}()

// FromString parse string into SportBits3 constant it's represent, return SportBits3Invalid if not found.
func SportBits3FromString(s string) SportBits3 {
	val, ok := strtosportbits3[s]
	if !ok {
		return strtosportbits3["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportBits3() []SportBits3 {
	vs := make([]SportBits3, 0, len(sportbits3tostrs))
	for i := range sportbits3tostrs {
		vs = append(vs, SportBits3(i))
	}
	return vs
}
