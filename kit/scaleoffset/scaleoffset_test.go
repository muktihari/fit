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
	"github.com/muktihari/fit/proto"
)

func TestApplyValue(t *testing.T) {
	tt := []struct {
		value  any
		scale  float64
		offset float64
		result proto.Value
	}{
		{value: proto.Uint16(37304), scale: 5, offset: 500, result: proto.Float64(6960.8)},
		{value: proto.Int8(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Uint8(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Int16(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Uint16(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Int32(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Uint32(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Int64(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Uint64(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Float32(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.Float64(10), scale: 5, offset: 1, result: proto.Float64(1)},
		{value: proto.SliceUint16([]uint16{37304}), scale: 5, offset: 500, result: proto.SliceFloat64([]float64{6960.8})},
		{value: proto.SliceInt8([]int8{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceUint8([]uint8{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceInt16([]int16{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceUint16([]uint16{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceInt32([]int32{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceUint32([]uint32{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceInt64([]int64{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceUint64([]uint64{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceFloat32([]float32{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceFloat64([]float64{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.SliceInt8([]int8{10}), scale: 5, offset: 1, result: proto.SliceFloat64([]float64{1})},
		{value: proto.String("fit"), scale: 1, offset: 0, result: proto.String("fit")},
		{value: proto.String("fit"), scale: 255, offset: 255, result: proto.String("fit")},
		{value: uint16(37304), scale: 5, offset: 500, result: proto.Float64(6960.8)},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("[%d] %#v", i, tc.value)
		value, ok := tc.value.(proto.Value)
		if ok {
			name = fmt.Sprintf("[%d] %#v", i, value.Any())
		}
		t.Run(name, func(t *testing.T) {
			result := scaleoffset.ApplyValue(tc.value, tc.scale, tc.offset)
			if diff := cmp.Diff(result, tc.result,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

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
		{value: proto.Uint8(10), scale: 5, offset: 1, result: float64(1)},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("[%d] %#v", i, tc.value)
		value, ok := tc.value.(proto.Value)
		if ok {
			name = fmt.Sprintf("[%d] %#v", i, value.Any())
		}
		t.Run(name, func(t *testing.T) {
			result := scaleoffset.ApplyAny(tc.value, tc.scale, tc.offset)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDiscardValue(t *testing.T) {
	tt := []struct {
		name     string
		value    any
		scale    float64
		offset   float64
		baseType basetype.BaseType
		result   proto.Value
	}{
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint16, result: proto.Uint16(37304)},
		{value: proto.Float64(5), scale: 2, offset: 2, baseType: basetype.Byte, result: proto.Uint8(14)},
		{value: proto.Float64(5), scale: 2, offset: 2, baseType: basetype.Sint8, result: proto.Int8(14)},
		{value: proto.Float64(5), scale: 2, offset: 2, baseType: basetype.Uint8, result: proto.Uint8(14)},
		{value: proto.Float64(5), scale: 2, offset: 2, baseType: basetype.Uint8z, result: proto.Uint8(14)},
		{value: proto.Float64(5), scale: 2, offset: 2, baseType: basetype.Sint16, result: proto.Int16(14)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint16z, result: proto.Uint16(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Sint32, result: proto.Int32(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint32, result: proto.Uint32(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint32z, result: proto.Uint32(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Sint64, result: proto.Int64(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint64, result: proto.Uint64(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint64z, result: proto.Uint64(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Float32, result: proto.Float32(37304)},
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Float64, result: proto.Float64(37304)},
		{value: proto.Float64(6960.8), scale: 1, offset: 0, baseType: basetype.Float64, result: proto.Float64(6960.8)},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Uint16, result: proto.SliceUint16([]uint16{37304})},
		{value: proto.SliceFloat64([]float64{5}), scale: 2, offset: 2, baseType: basetype.Byte, result: proto.SliceUint8([]byte{14})},
		{value: proto.SliceFloat64([]float64{5}), scale: 2, offset: 2, baseType: basetype.Sint8, result: proto.SliceInt8([]int8{14})},
		{value: proto.SliceFloat64([]float64{5}), scale: 2, offset: 2, baseType: basetype.Uint8, result: proto.SliceUint8([]uint8{14})},
		{value: proto.SliceFloat64([]float64{5}), scale: 2, offset: 2, baseType: basetype.Uint8z, result: proto.SliceUint8([]uint8{14})},
		{value: proto.SliceFloat64([]float64{5}), scale: 2, offset: 2, baseType: basetype.Sint16, result: proto.SliceInt16([]int16{14})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Uint16z, result: proto.SliceUint16([]uint16{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Sint32, result: proto.SliceInt32([]int32{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Uint32, result: proto.SliceUint32([]uint32{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Uint32z, result: proto.SliceUint32([]uint32{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Sint64, result: proto.SliceInt64([]int64{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Uint64, result: proto.SliceUint64([]uint64{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Uint64z, result: proto.SliceUint64([]uint64{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Float32, result: proto.SliceFloat32([]float32{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 5, offset: 500, baseType: basetype.Float64, result: proto.SliceFloat64([]float64{37304})},
		{value: proto.SliceFloat64([]float64{6960.8}), scale: 1, offset: 0, baseType: basetype.Float64, result: proto.SliceFloat64([]float64{6960.8})},
		{value: proto.String("fit"), scale: 5, offset: 500, baseType: basetype.Float64, result: proto.String("fit")},
		{value: float64(6960.8), scale: 5, offset: 500, baseType: basetype.Uint16z, result: proto.Uint16(37304)},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("[%d] %#v", i, tc.value)
		value, ok := tc.value.(proto.Value)
		if ok {
			name = fmt.Sprintf("[%d] %#v", i, value.Any())
		}
		t.Run(name, func(t *testing.T) {
			result := scaleoffset.DiscardValue(tc.value, tc.baseType, tc.scale, tc.offset)
			if diff := cmp.Diff(result, tc.result,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
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
		{value: proto.Float64(6960.8), scale: 5, offset: 500, baseType: basetype.Sint64, result: int64(37304)},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("[%d] %#v", i, tc.value)
		value, ok := tc.value.(proto.Value)
		if ok {
			name = fmt.Sprintf("[%d] %#v", i, value.Any())
		}
		t.Run(name, func(t *testing.T) {
			result := scaleoffset.DiscardAny(tc.value, tc.baseType, tc.scale, tc.offset)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
