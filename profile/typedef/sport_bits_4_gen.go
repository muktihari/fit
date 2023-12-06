// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits4 uint8

const (
	SportBits4Sailing               SportBits4 = 0x01
	SportBits4IceSkating            SportBits4 = 0x02
	SportBits4SkyDiving             SportBits4 = 0x04
	SportBits4Snowshoeing           SportBits4 = 0x08
	SportBits4Snowmobiling          SportBits4 = 0x10
	SportBits4StandUpPaddleboarding SportBits4 = 0x20
	SportBits4Surfing               SportBits4 = 0x40
	SportBits4Wakeboarding          SportBits4 = 0x80
	SportBits4Invalid               SportBits4 = 0x0 // INVALID
)

var sportbits4tostrs = map[SportBits4]string{
	SportBits4Sailing:               "sailing",
	SportBits4IceSkating:            "ice_skating",
	SportBits4SkyDiving:             "sky_diving",
	SportBits4Snowshoeing:           "snowshoeing",
	SportBits4Snowmobiling:          "snowmobiling",
	SportBits4StandUpPaddleboarding: "stand_up_paddleboarding",
	SportBits4Surfing:               "surfing",
	SportBits4Wakeboarding:          "wakeboarding",
	SportBits4Invalid:               "invalid",
}

func (s SportBits4) String() string {
	val, ok := sportbits4tostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosportbits4 = func() map[string]SportBits4 {
	m := make(map[string]SportBits4)
	for t, str := range sportbits4tostrs {
		m[str] = SportBits4(t)
	}
	return m
}()

// FromString parse string into SportBits4 constant it's represent, return SportBits4Invalid if not found.
func SportBits4FromString(s string) SportBits4 {
	val, ok := strtosportbits4[s]
	if !ok {
		return strtosportbits4["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportBits4() []SportBits4 {
	vs := make([]SportBits4, 0, len(sportbits4tostrs))
	for i := range sportbits4tostrs {
		vs = append(vs, SportBits4(i))
	}
	return vs
}
