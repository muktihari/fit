// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scaleoffset_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
)

func TestApplyAny(t *testing.T) {
	tt := []struct {
		value  any
		scale  float64
		offset float64
		result any
	}{
		{value: uint16(37304), scale: 5, offset: 500, result: float64(6960.8)},
		{value: int8(10), scale: 5, offset: 1, result: float64(1)},
		{value: uint8(10), scale: 5, offset: 1, result: float64(1)},
		{value: int16(10), scale: 5, offset: 1, result: float64(1)},
		{value: uint16(10), scale: 5, offset: 1, result: float64(1)},
		{value: int32(10), scale: 5, offset: 1, result: float64(1)},
		{value: uint32(10), scale: 5, offset: 1, result: float64(1)},
		{value: int64(10), scale: 5, offset: 1, result: float64(1)},
		{value: uint64(10), scale: 5, offset: 1, result: float64(1)},
		{value: float32(10), scale: 5, offset: 1, result: float64(1)},
		{value: float64(10), scale: 5, offset: 1, result: float64(1)},
		{value: []uint16{37304}, scale: 5, offset: 500, result: []float64{6960.8}},
		{value: []int8{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []uint8{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []int16{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []uint16{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []int32{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []uint32{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []int64{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []uint64{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []float32{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []float64{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: []int8{10}, scale: 5, offset: 1, result: []float64{1}},
		{value: string("fit"), scale: 1, offset: 0, result: "fit"},
		{value: string("fit"), scale: 255, offset: 255, result: "fit"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.value), func(t *testing.T) {
			result := scaleoffset.ApplyAny(tc.value, tc.scale, tc.offset)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDiscardAny(t *testing.T) {
	tt := []struct {
		name     string
		value    any
		scale    float64
		offset   float64
		baseType basetype.BaseType
		result   any
	}{
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint16, result: uint16(37304)},
		{value: float64(5), scale: 2, offset: 2, baseType: basetype.Byte, result: byte(14)},
		{value: float64(5), scale: 2, offset: 2, baseType: basetype.Sint8, result: int8(14)},
		{value: float64(5), scale: 2, offset: 2, baseType: basetype.Uint8, result: uint8(14)},
		{value: float64(5), scale: 2, offset: 2, baseType: basetype.Uint8z, result: uint8(14)},
		{value: float64(5), scale: 2, offset: 2, baseType: basetype.Sint16, result: int16(14)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint16z, result: uint16(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Sint32, result: int32(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint32, result: uint32(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint32z, result: uint32(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Sint64, result: int64(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint64, result: uint64(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint64z, result: uint64(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Float32, result: float32(37304)},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Float64, result: float64(37304)},
		{value: float64(6960.8), scale: 1, offset: 0, baseType: basetype.Float64, result: float64(6960.8)},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Uint16, result: []uint16{37304}},
		{value: []float64{5}, scale: 2, offset: 2, baseType: basetype.Byte, result: []byte{14}},
		{value: []float64{5}, scale: 2, offset: 2, baseType: basetype.Sint8, result: []int8{14}},
		{value: []float64{5}, scale: 2, offset: 2, baseType: basetype.Uint8, result: []uint8{14}},
		{value: []float64{5}, scale: 2, offset: 2, baseType: basetype.Uint8z, result: []uint8{14}},
		{value: []float64{5}, scale: 2, offset: 2, baseType: basetype.Sint16, result: []int16{14}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Uint16z, result: []uint16{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Sint32, result: []int32{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Uint32, result: []uint32{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Uint32z, result: []uint32{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Sint64, result: []int64{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Uint64, result: []uint64{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Uint64z, result: []uint64{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Float32, result: []float32{37304}},
		{value: []float64{6960.8}, scale: 5, offset: 500, baseType: basetype.Float64, result: []float64{37304}},
		{value: []float64{6960.8}, scale: 1, offset: 0, baseType: basetype.Float64, result: []float64{6960.8}},
		{value: string("fit"), scale: 5, offset: 500, baseType: basetype.Float64, result: string("fit")},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.value), func(t *testing.T) {
			result := scaleoffset.DiscardAny(tc.value, tc.baseType, tc.scale, tc.offset)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
