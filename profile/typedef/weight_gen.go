// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Weight uint16

const (
	WeightCalculating Weight = 0xFFFE
	WeightInvalid     Weight = 0xFFFF // INVALID
)

var weighttostrs = map[Weight]string{
	WeightCalculating: "calculating",
	WeightInvalid:     "invalid",
}

func (w Weight) String() string {
	val, ok := weighttostrs[w]
	if !ok {
		return strconv.FormatUint(uint64(w), 10)
	}
	return val
}

var strtoweight = func() map[string]Weight {
	m := make(map[string]Weight)
	for t, str := range weighttostrs {
		m[str] = Weight(t)
	}
	return m
}()

// FromString parse string into Weight constant it's represent, return WeightInvalid if not found.
func WeightFromString(s string) Weight {
	val, ok := strtoweight[s]
	if !ok {
		return strtoweight["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWeight() []Weight {
	vs := make([]Weight, 0, len(weighttostrs))
	for i := range weighttostrs {
		vs = append(vs, Weight(i))
	}
	return vs
}
