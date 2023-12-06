// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TapSensitivity byte

const (
	TapSensitivityHigh    TapSensitivity = 0
	TapSensitivityMedium  TapSensitivity = 1
	TapSensitivityLow     TapSensitivity = 2
	TapSensitivityInvalid TapSensitivity = 0xFF // INVALID
)

var tapsensitivitytostrs = map[TapSensitivity]string{
	TapSensitivityHigh:    "high",
	TapSensitivityMedium:  "medium",
	TapSensitivityLow:     "low",
	TapSensitivityInvalid: "invalid",
}

func (t TapSensitivity) String() string {
	val, ok := tapsensitivitytostrs[t]
	if !ok {
		return strconv.Itoa(int(t))
	}
	return val
}

var strtotapsensitivity = func() map[string]TapSensitivity {
	m := make(map[string]TapSensitivity)
	for t, str := range tapsensitivitytostrs {
		m[str] = TapSensitivity(t)
	}
	return m
}()

// FromString parse string into TapSensitivity constant it's represent, return TapSensitivityInvalid if not found.
func TapSensitivityFromString(s string) TapSensitivity {
	val, ok := strtotapsensitivity[s]
	if !ok {
		return strtotapsensitivity["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListTapSensitivity() []TapSensitivity {
	vs := make([]TapSensitivity, 0, len(tapsensitivitytostrs))
	for i := range tapsensitivitytostrs {
		vs = append(vs, TapSensitivity(i))
	}
	return vs
}
