// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BikeLightBeamAngleMode uint8

const (
	BikeLightBeamAngleModeManual  BikeLightBeamAngleMode = 0
	BikeLightBeamAngleModeAuto    BikeLightBeamAngleMode = 1
	BikeLightBeamAngleModeInvalid BikeLightBeamAngleMode = 0xFF // INVALID
)

var bikelightbeamanglemodetostrs = map[BikeLightBeamAngleMode]string{
	BikeLightBeamAngleModeManual:  "manual",
	BikeLightBeamAngleModeAuto:    "auto",
	BikeLightBeamAngleModeInvalid: "invalid",
}

func (b BikeLightBeamAngleMode) String() string {
	val, ok := bikelightbeamanglemodetostrs[b]
	if !ok {
		return strconv.FormatUint(uint64(b), 10)
	}
	return val
}

var strtobikelightbeamanglemode = func() map[string]BikeLightBeamAngleMode {
	m := make(map[string]BikeLightBeamAngleMode)
	for t, str := range bikelightbeamanglemodetostrs {
		m[str] = BikeLightBeamAngleMode(t)
	}
	return m
}()

// FromString parse string into BikeLightBeamAngleMode constant it's represent, return BikeLightBeamAngleModeInvalid if not found.
func BikeLightBeamAngleModeFromString(s string) BikeLightBeamAngleMode {
	val, ok := strtobikelightbeamanglemode[s]
	if !ok {
		return strtobikelightbeamanglemode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListBikeLightBeamAngleMode() []BikeLightBeamAngleMode {
	vs := make([]BikeLightBeamAngleMode, 0, len(bikelightbeamanglemodetostrs))
	for i := range bikelightbeamanglemodetostrs {
		vs = append(vs, BikeLightBeamAngleMode(i))
	}
	return vs
}
