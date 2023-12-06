// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SwimStroke byte

const (
	SwimStrokeFreestyle    SwimStroke = 0
	SwimStrokeBackstroke   SwimStroke = 1
	SwimStrokeBreaststroke SwimStroke = 2
	SwimStrokeButterfly    SwimStroke = 3
	SwimStrokeDrill        SwimStroke = 4
	SwimStrokeMixed        SwimStroke = 5
	SwimStrokeIm           SwimStroke = 6    // IM is a mixed interval containing the same number of lengths for each of: Butterfly, Backstroke, Breaststroke, Freestyle, swam in that order.
	SwimStrokeInvalid      SwimStroke = 0xFF // INVALID
)

var swimstroketostrs = map[SwimStroke]string{
	SwimStrokeFreestyle:    "freestyle",
	SwimStrokeBackstroke:   "backstroke",
	SwimStrokeBreaststroke: "breaststroke",
	SwimStrokeButterfly:    "butterfly",
	SwimStrokeDrill:        "drill",
	SwimStrokeMixed:        "mixed",
	SwimStrokeIm:           "im",
	SwimStrokeInvalid:      "invalid",
}

func (s SwimStroke) String() string {
	val, ok := swimstroketostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtoswimstroke = func() map[string]SwimStroke {
	m := make(map[string]SwimStroke)
	for t, str := range swimstroketostrs {
		m[str] = SwimStroke(t)
	}
	return m
}()

// FromString parse string into SwimStroke constant it's represent, return SwimStrokeInvalid if not found.
func SwimStrokeFromString(s string) SwimStroke {
	val, ok := strtoswimstroke[s]
	if !ok {
		return strtoswimstroke["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSwimStroke() []SwimStroke {
	vs := make([]SwimStroke, 0, len(swimstroketostrs))
	for i := range swimstroketostrs {
		vs = append(vs, SwimStroke(i))
	}
	return vs
}
