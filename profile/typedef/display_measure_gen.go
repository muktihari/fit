// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DisplayMeasure byte

const (
	DisplayMeasureMetric   DisplayMeasure = 0
	DisplayMeasureStatute  DisplayMeasure = 1
	DisplayMeasureNautical DisplayMeasure = 2
	DisplayMeasureInvalid  DisplayMeasure = 0xFF // INVALID
)

var displaymeasuretostrs = map[DisplayMeasure]string{
	DisplayMeasureMetric:   "metric",
	DisplayMeasureStatute:  "statute",
	DisplayMeasureNautical: "nautical",
	DisplayMeasureInvalid:  "invalid",
}

func (d DisplayMeasure) String() string {
	val, ok := displaymeasuretostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodisplaymeasure = func() map[string]DisplayMeasure {
	m := make(map[string]DisplayMeasure)
	for t, str := range displaymeasuretostrs {
		m[str] = DisplayMeasure(t)
	}
	return m
}()

// FromString parse string into DisplayMeasure constant it's represent, return DisplayMeasureInvalid if not found.
func DisplayMeasureFromString(s string) DisplayMeasure {
	val, ok := strtodisplaymeasure[s]
	if !ok {
		return strtodisplaymeasure["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDisplayMeasure() []DisplayMeasure {
	vs := make([]DisplayMeasure, 0, len(displaymeasuretostrs))
	for i := range displaymeasuretostrs {
		vs = append(vs, DisplayMeasure(i))
	}
	return vs
}