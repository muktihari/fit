// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DisplayHeart byte

const (
	DisplayHeartBpm     DisplayHeart = 0
	DisplayHeartMax     DisplayHeart = 1
	DisplayHeartReserve DisplayHeart = 2
	DisplayHeartInvalid DisplayHeart = 0xFF // INVALID
)

var displayhearttostrs = map[DisplayHeart]string{
	DisplayHeartBpm:     "bpm",
	DisplayHeartMax:     "max",
	DisplayHeartReserve: "reserve",
	DisplayHeartInvalid: "invalid",
}

func (d DisplayHeart) String() string {
	val, ok := displayhearttostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodisplayheart = func() map[string]DisplayHeart {
	m := make(map[string]DisplayHeart)
	for t, str := range displayhearttostrs {
		m[str] = DisplayHeart(t)
	}
	return m
}()

// FromString parse string into DisplayHeart constant it's represent, return DisplayHeartInvalid if not found.
func DisplayHeartFromString(s string) DisplayHeart {
	val, ok := strtodisplayheart[s]
	if !ok {
		return strtodisplayheart["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDisplayHeart() []DisplayHeart {
	vs := make([]DisplayHeart, 0, len(displayhearttostrs))
	for i := range displayhearttostrs {
		vs = append(vs, DisplayHeart(i))
	}
	return vs
}
