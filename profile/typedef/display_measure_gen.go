// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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
	DisplayMeasureInvalid  DisplayMeasure = 0xFF
)

func (d DisplayMeasure) Byte() byte { return byte(d) }

func (d DisplayMeasure) String() string {
	switch d {
	case DisplayMeasureMetric:
		return "metric"
	case DisplayMeasureStatute:
		return "statute"
	case DisplayMeasureNautical:
		return "nautical"
	default:
		return "DisplayMeasureInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DisplayMeasure constant it's represent, return DisplayMeasureInvalid if not found.
func DisplayMeasureFromString(s string) DisplayMeasure {
	switch s {
	case "metric":
		return DisplayMeasureMetric
	case "statute":
		return DisplayMeasureStatute
	case "nautical":
		return DisplayMeasureNautical
	default:
		return DisplayMeasureInvalid
	}
}

// List returns all constants.
func ListDisplayMeasure() []DisplayMeasure {
	return []DisplayMeasure{
		DisplayMeasureMetric,
		DisplayMeasureStatute,
		DisplayMeasureNautical,
	}
}
