// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DeviceIndex uint8

const (
	DeviceIndexCreator DeviceIndex = 0    // Creator of the file is always device index 0.
	DeviceIndexInvalid DeviceIndex = 0xFF // INVALID
)

var deviceindextostrs = map[DeviceIndex]string{
	DeviceIndexCreator: "creator",
	DeviceIndexInvalid: "invalid",
}

func (d DeviceIndex) String() string {
	val, ok := deviceindextostrs[d]
	if !ok {
		return strconv.FormatUint(uint64(d), 10)
	}
	return val
}

var strtodeviceindex = func() map[string]DeviceIndex {
	m := make(map[string]DeviceIndex)
	for t, str := range deviceindextostrs {
		m[str] = DeviceIndex(t)
	}
	return m
}()

// FromString parse string into DeviceIndex constant it's represent, return DeviceIndexInvalid if not found.
func DeviceIndexFromString(s string) DeviceIndex {
	val, ok := strtodeviceindex[s]
	if !ok {
		return strtodeviceindex["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDeviceIndex() []DeviceIndex {
	vs := make([]DeviceIndex, 0, len(deviceindextostrs))
	for i := range deviceindextostrs {
		vs = append(vs, DeviceIndex(i))
	}
	return vs
}
