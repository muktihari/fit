// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SourceType byte

const (
	SourceTypeAnt                SourceType = 0    // External device connected with ANT
	SourceTypeAntplus            SourceType = 1    // External device connected with ANT+
	SourceTypeBluetooth          SourceType = 2    // External device connected with BT
	SourceTypeBluetoothLowEnergy SourceType = 3    // External device connected with BLE
	SourceTypeWifi               SourceType = 4    // External device connected with Wifi
	SourceTypeLocal              SourceType = 5    // Onboard device
	SourceTypeInvalid            SourceType = 0xFF // INVALID
)

var sourcetypetostrs = map[SourceType]string{
	SourceTypeAnt:                "ant",
	SourceTypeAntplus:            "antplus",
	SourceTypeBluetooth:          "bluetooth",
	SourceTypeBluetoothLowEnergy: "bluetooth_low_energy",
	SourceTypeWifi:               "wifi",
	SourceTypeLocal:              "local",
	SourceTypeInvalid:            "invalid",
}

func (s SourceType) String() string {
	val, ok := sourcetypetostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosourcetype = func() map[string]SourceType {
	m := make(map[string]SourceType)
	for t, str := range sourcetypetostrs {
		m[str] = SourceType(t)
	}
	return m
}()

// FromString parse string into SourceType constant it's represent, return SourceTypeInvalid if not found.
func SourceTypeFromString(s string) SourceType {
	val, ok := strtosourcetype[s]
	if !ok {
		return strtosourcetype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSourceType() []SourceType {
	vs := make([]SourceType, 0, len(sourcetypetostrs))
	for i := range sourcetypetostrs {
		vs = append(vs, SourceType(i))
	}
	return vs
}
