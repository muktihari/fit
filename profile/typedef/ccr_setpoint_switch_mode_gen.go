// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CcrSetpointSwitchMode byte

const (
	CcrSetpointSwitchModeManual    CcrSetpointSwitchMode = 0    // User switches setpoints manually
	CcrSetpointSwitchModeAutomatic CcrSetpointSwitchMode = 1    // Switch automatically based on depth
	CcrSetpointSwitchModeInvalid   CcrSetpointSwitchMode = 0xFF // INVALID
)

var ccrsetpointswitchmodetostrs = map[CcrSetpointSwitchMode]string{
	CcrSetpointSwitchModeManual:    "manual",
	CcrSetpointSwitchModeAutomatic: "automatic",
	CcrSetpointSwitchModeInvalid:   "invalid",
}

func (c CcrSetpointSwitchMode) String() string {
	val, ok := ccrsetpointswitchmodetostrs[c]
	if !ok {
		return strconv.Itoa(int(c))
	}
	return val
}

var strtoccrsetpointswitchmode = func() map[string]CcrSetpointSwitchMode {
	m := make(map[string]CcrSetpointSwitchMode)
	for t, str := range ccrsetpointswitchmodetostrs {
		m[str] = CcrSetpointSwitchMode(t)
	}
	return m
}()

// FromString parse string into CcrSetpointSwitchMode constant it's represent, return CcrSetpointSwitchModeInvalid if not found.
func CcrSetpointSwitchModeFromString(s string) CcrSetpointSwitchMode {
	val, ok := strtoccrsetpointswitchmode[s]
	if !ok {
		return strtoccrsetpointswitchmode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListCcrSetpointSwitchMode() []CcrSetpointSwitchMode {
	vs := make([]CcrSetpointSwitchMode, 0, len(ccrsetpointswitchmodetostrs))
	for i := range ccrsetpointswitchmodetostrs {
		vs = append(vs, CcrSetpointSwitchMode(i))
	}
	return vs
}
