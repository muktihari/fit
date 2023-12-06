// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AnalogWatchfaceLayout byte

const (
	AnalogWatchfaceLayoutMinimal     AnalogWatchfaceLayout = 0
	AnalogWatchfaceLayoutTraditional AnalogWatchfaceLayout = 1
	AnalogWatchfaceLayoutModern      AnalogWatchfaceLayout = 2
	AnalogWatchfaceLayoutInvalid     AnalogWatchfaceLayout = 0xFF // INVALID
)

var analogwatchfacelayouttostrs = map[AnalogWatchfaceLayout]string{
	AnalogWatchfaceLayoutMinimal:     "minimal",
	AnalogWatchfaceLayoutTraditional: "traditional",
	AnalogWatchfaceLayoutModern:      "modern",
	AnalogWatchfaceLayoutInvalid:     "invalid",
}

func (a AnalogWatchfaceLayout) String() string {
	val, ok := analogwatchfacelayouttostrs[a]
	if !ok {
		return strconv.Itoa(int(a))
	}
	return val
}

var strtoanalogwatchfacelayout = func() map[string]AnalogWatchfaceLayout {
	m := make(map[string]AnalogWatchfaceLayout)
	for t, str := range analogwatchfacelayouttostrs {
		m[str] = AnalogWatchfaceLayout(t)
	}
	return m
}()

// FromString parse string into AnalogWatchfaceLayout constant it's represent, return AnalogWatchfaceLayoutInvalid if not found.
func AnalogWatchfaceLayoutFromString(s string) AnalogWatchfaceLayout {
	val, ok := strtoanalogwatchfacelayout[s]
	if !ok {
		return strtoanalogwatchfacelayout["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListAnalogWatchfaceLayout() []AnalogWatchfaceLayout {
	vs := make([]AnalogWatchfaceLayout, 0, len(analogwatchfacelayouttostrs))
	for i := range analogwatchfacelayouttostrs {
		vs = append(vs, AnalogWatchfaceLayout(i))
	}
	return vs
}
