// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type GasConsumptionRateType byte

const (
	GasConsumptionRateTypePressureSac GasConsumptionRateType = 0    // Pressure-based Surface Air Consumption
	GasConsumptionRateTypeVolumeSac   GasConsumptionRateType = 1    // Volumetric Surface Air Consumption
	GasConsumptionRateTypeRmv         GasConsumptionRateType = 2    // Respiratory Minute Volume
	GasConsumptionRateTypeInvalid     GasConsumptionRateType = 0xFF // INVALID
)

var gasconsumptionratetypetostrs = map[GasConsumptionRateType]string{
	GasConsumptionRateTypePressureSac: "pressure_sac",
	GasConsumptionRateTypeVolumeSac:   "volume_sac",
	GasConsumptionRateTypeRmv:         "rmv",
	GasConsumptionRateTypeInvalid:     "invalid",
}

func (g GasConsumptionRateType) String() string {
	val, ok := gasconsumptionratetypetostrs[g]
	if !ok {
		return strconv.Itoa(int(g))
	}
	return val
}

var strtogasconsumptionratetype = func() map[string]GasConsumptionRateType {
	m := make(map[string]GasConsumptionRateType)
	for t, str := range gasconsumptionratetypetostrs {
		m[str] = GasConsumptionRateType(t)
	}
	return m
}()

// FromString parse string into GasConsumptionRateType constant it's represent, return GasConsumptionRateTypeInvalid if not found.
func GasConsumptionRateTypeFromString(s string) GasConsumptionRateType {
	val, ok := strtogasconsumptionratetype[s]
	if !ok {
		return strtogasconsumptionratetype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListGasConsumptionRateType() []GasConsumptionRateType {
	vs := make([]GasConsumptionRateType, 0, len(gasconsumptionratetypetostrs))
	for i := range gasconsumptionratetypetostrs {
		vs = append(vs, GasConsumptionRateType(i))
	}
	return vs
}
