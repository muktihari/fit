// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Spo2MeasurementType byte

const (
	Spo2MeasurementTypeOffWrist        Spo2MeasurementType = 0
	Spo2MeasurementTypeSpotCheck       Spo2MeasurementType = 1
	Spo2MeasurementTypeContinuousCheck Spo2MeasurementType = 2
	Spo2MeasurementTypePeriodic        Spo2MeasurementType = 3
	Spo2MeasurementTypeInvalid         Spo2MeasurementType = 0xFF // INVALID
)

var spo2measurementtypetostrs = map[Spo2MeasurementType]string{
	Spo2MeasurementTypeOffWrist:        "off_wrist",
	Spo2MeasurementTypeSpotCheck:       "spot_check",
	Spo2MeasurementTypeContinuousCheck: "continuous_check",
	Spo2MeasurementTypePeriodic:        "periodic",
	Spo2MeasurementTypeInvalid:         "invalid",
}

func (s Spo2MeasurementType) String() string {
	val, ok := spo2measurementtypetostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtospo2measurementtype = func() map[string]Spo2MeasurementType {
	m := make(map[string]Spo2MeasurementType)
	for t, str := range spo2measurementtypetostrs {
		m[str] = Spo2MeasurementType(t)
	}
	return m
}()

// FromString parse string into Spo2MeasurementType constant it's represent, return Spo2MeasurementTypeInvalid if not found.
func Spo2MeasurementTypeFromString(s string) Spo2MeasurementType {
	val, ok := strtospo2measurementtype[s]
	if !ok {
		return strtospo2measurementtype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSpo2MeasurementType() []Spo2MeasurementType {
	vs := make([]Spo2MeasurementType, 0, len(spo2measurementtypetostrs))
	for i := range spo2measurementtypetostrs {
		vs = append(vs, Spo2MeasurementType(i))
	}
	return vs
}
