// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef_test

import (
	"fmt"
	"testing"

	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
)

type test_uint8 uint8
type test_string string

func TestLen(t *testing.T) {
	tt := []struct {
		value       any
		baseType    basetype.BaseType
		sizeInBytes int
	}{
		{value: int8(10), sizeInBytes: 1, baseType: basetype.Sint8},
		{value: uint8(10), sizeInBytes: 1, baseType: basetype.Uint8},
		{value: int16(10), sizeInBytes: 2, baseType: basetype.Sint16},
		{value: uint16(10), sizeInBytes: 2, baseType: basetype.Uint16},
		{value: int32(10), sizeInBytes: 4, baseType: basetype.Sint32},
		{value: uint32(10), sizeInBytes: 4, baseType: basetype.Uint32},
		{value: float32(10), sizeInBytes: 4, baseType: basetype.Float32},
		{value: float64(10), sizeInBytes: 8, baseType: basetype.Float64},
		{value: int64(10), sizeInBytes: 8, baseType: basetype.Sint64},
		{value: uint64(10), sizeInBytes: 8, baseType: basetype.Uint64},
		{value: []int8{10, 9, 8, 7}, sizeInBytes: 4, baseType: basetype.Sint8},
		{value: []uint8{10, 9, 8, 7}, sizeInBytes: 4, baseType: basetype.Uint8},
		{value: []int16{10, 9, 8, 7}, sizeInBytes: 4 * 2, baseType: basetype.Uint16},
		{value: []uint16{10, 9, 8, 7}, sizeInBytes: 4 * 2, baseType: basetype.Uint16},
		{value: []int32{10, 9, 8, 7}, sizeInBytes: 4 * 4, baseType: basetype.Sint32},
		{value: []uint32{10, 9, 8, 7}, sizeInBytes: 4 * 4, baseType: basetype.Uint32},
		{value: string(""), sizeInBytes: 1, baseType: basetype.String},
		{value: string("\x00"), sizeInBytes: 1, baseType: basetype.String},
		{value: string("fit sdk"), sizeInBytes: 8, baseType: basetype.String},
		{value: string("fit sdk\x00"), sizeInBytes: 8, baseType: basetype.String},
		{value: []string{"fit sdk"}, sizeInBytes: 8, baseType: basetype.String},
		{value: []string{}, sizeInBytes: 1, baseType: basetype.String},
		{value: []string{""}, sizeInBytes: 1, baseType: basetype.String},
		{value: []string{"\x00"}, sizeInBytes: 1, baseType: basetype.String},
		{value: []string{"\x00\x00\x00"}, sizeInBytes: 3, baseType: basetype.String},
		{value: []string{"\x00", "\x00", "\x00"}, sizeInBytes: 3, baseType: basetype.String},
		{value: []string{"fit sdk", "a"}, sizeInBytes: 10, baseType: basetype.String},
		{value: []string{"fit sdk\x00", "a\x00"}, sizeInBytes: 10, baseType: basetype.String},
		{value: []float32{10, 9, 8, 7}, sizeInBytes: 4 * 4, baseType: basetype.Float32},
		{value: []float64{10, 9, 8, 7}, sizeInBytes: 4 * 8, baseType: basetype.Float64},
		{value: []int64{10, 9, 8, 7}, sizeInBytes: 4 * 8, baseType: basetype.Sint64},
		{value: []uint64{10, 9, 8, 7}, sizeInBytes: 4 * 8, baseType: basetype.Uint64},
		{value: test_uint8(7), sizeInBytes: 1, baseType: basetype.Uint8},
		{value: []test_uint8{10, 9, 8, 7}, sizeInBytes: 4, baseType: basetype.Uint8},
		{value: test_string(""), sizeInBytes: 1, baseType: basetype.String},
		{value: test_string("\x00"), sizeInBytes: 1, baseType: basetype.String},
		{value: test_string("abc"), sizeInBytes: 4, baseType: basetype.String},
		{value: []test_string{}, sizeInBytes: 1, baseType: basetype.String},
		{value: []test_string{""}, sizeInBytes: 1, baseType: basetype.String},
		{value: []test_string{"\x00"}, sizeInBytes: 1, baseType: basetype.String},
		{value: []test_string{"abc"}, sizeInBytes: 4, baseType: basetype.String},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v)", tc.value, tc.value), func(t *testing.T) {
			size := typedef.Sizeof(tc.value, tc.baseType)
			if size != tc.sizeInBytes {
				t.Fatalf("expected: %d, got: %d", tc.sizeInBytes, size)
			}
		})
	}
}
