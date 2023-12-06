// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BikeLightNetworkConfigType byte

const (
	BikeLightNetworkConfigTypeAuto           BikeLightNetworkConfigType = 0
	BikeLightNetworkConfigTypeIndividual     BikeLightNetworkConfigType = 4
	BikeLightNetworkConfigTypeHighVisibility BikeLightNetworkConfigType = 5
	BikeLightNetworkConfigTypeTrail          BikeLightNetworkConfigType = 6
	BikeLightNetworkConfigTypeInvalid        BikeLightNetworkConfigType = 0xFF // INVALID
)

var bikelightnetworkconfigtypetostrs = map[BikeLightNetworkConfigType]string{
	BikeLightNetworkConfigTypeAuto:           "auto",
	BikeLightNetworkConfigTypeIndividual:     "individual",
	BikeLightNetworkConfigTypeHighVisibility: "high_visibility",
	BikeLightNetworkConfigTypeTrail:          "trail",
	BikeLightNetworkConfigTypeInvalid:        "invalid",
}

func (b BikeLightNetworkConfigType) String() string {
	val, ok := bikelightnetworkconfigtypetostrs[b]
	if !ok {
		return strconv.Itoa(int(b))
	}
	return val
}

var strtobikelightnetworkconfigtype = func() map[string]BikeLightNetworkConfigType {
	m := make(map[string]BikeLightNetworkConfigType)
	for t, str := range bikelightnetworkconfigtypetostrs {
		m[str] = BikeLightNetworkConfigType(t)
	}
	return m
}()

// FromString parse string into BikeLightNetworkConfigType constant it's represent, return BikeLightNetworkConfigTypeInvalid if not found.
func BikeLightNetworkConfigTypeFromString(s string) BikeLightNetworkConfigType {
	val, ok := strtobikelightnetworkconfigtype[s]
	if !ok {
		return strtobikelightnetworkconfigtype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListBikeLightNetworkConfigType() []BikeLightNetworkConfigType {
	vs := make([]BikeLightNetworkConfigType, 0, len(bikelightnetworkconfigtypetostrs))
	for i := range bikelightnetworkconfigtypetostrs {
		vs = append(vs, BikeLightNetworkConfigType(i))
	}
	return vs
}
