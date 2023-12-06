// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WaterType byte

const (
	WaterTypeFresh   WaterType = 0
	WaterTypeSalt    WaterType = 1
	WaterTypeEn13319 WaterType = 2
	WaterTypeCustom  WaterType = 3
	WaterTypeInvalid WaterType = 0xFF // INVALID
)

var watertypetostrs = map[WaterType]string{
	WaterTypeFresh:   "fresh",
	WaterTypeSalt:    "salt",
	WaterTypeEn13319: "en13319",
	WaterTypeCustom:  "custom",
	WaterTypeInvalid: "invalid",
}

func (w WaterType) String() string {
	val, ok := watertypetostrs[w]
	if !ok {
		return strconv.Itoa(int(w))
	}
	return val
}

var strtowatertype = func() map[string]WaterType {
	m := make(map[string]WaterType)
	for t, str := range watertypetostrs {
		m[str] = WaterType(t)
	}
	return m
}()

// FromString parse string into WaterType constant it's represent, return WaterTypeInvalid if not found.
func WaterTypeFromString(s string) WaterType {
	val, ok := strtowatertype[s]
	if !ok {
		return strtowatertype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWaterType() []WaterType {
	vs := make([]WaterType, 0, len(watertypetostrs))
	for i := range watertypetostrs {
		vs = append(vs, WaterType(i))
	}
	return vs
}
