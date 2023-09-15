// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
)

func TestNumericToInt64(t *testing.T) {
	tt := []struct {
		value  any
		result int64
		ok     bool
	}{
		{value: nil, result: 0, ok: false},
		{value: int8(10), result: 10, ok: true},
		{value: uint8(10), result: 10, ok: true},
		{value: int16(10), result: 10, ok: true},
		{value: uint16(10), result: 10, ok: true},
		{value: int32(10), result: 10, ok: true},
		{value: uint32(10), result: 10, ok: true},
		{value: float32(10.1), result: 10, ok: true},
		{value: float64(10.2), result: 10, ok: true},
		{value: int64(10), result: 10, ok: true},
		{value: uint64(10), result: 10, ok: true},
		{value: string("fit"), result: 0, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result, ok := typeconv.NumericToInt64(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestIntegerToInt64(t *testing.T) {
	tt := []struct {
		value  any
		result int64
		ok     bool
	}{
		{value: nil, result: 0, ok: false},
		{value: int8(10), result: 10, ok: true},
		{value: uint8(10), result: 10, ok: true},
		{value: int16(10), result: 10, ok: true},
		{value: uint16(10), result: 10, ok: true},
		{value: int32(10), result: 10, ok: true},
		{value: uint32(10), result: 10, ok: true},
		{value: float32(10.1), result: 0, ok: false},
		{value: float64(10.2), result: 0, ok: false},
		{value: int64(10), result: 10, ok: true},
		{value: uint64(10), result: 10, ok: true},
		{value: string("fit"), result: 0, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result, ok := typeconv.IntegerToInt64(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestFloatToInt64(t *testing.T) {
	tt := []struct {
		value  any
		result int64
		ok     bool
	}{
		{value: nil, result: 0, ok: false},
		{value: uint32(10), result: 0, ok: false},
		{value: float32(10.1), result: 10, ok: true},
		{value: float64(11.2), result: 11, ok: true},
		{value: string("fit"), result: 0, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result, ok := typeconv.FloatToInt64(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestInt64ToNumber(t *testing.T) {
	tt := []struct {
		value    int64
		baseType basetype.BaseType
		result   any
	}{
		{value: 10, baseType: basetype.Sint8, result: int8(10)},
		{value: 10, baseType: basetype.Uint8, result: uint8(10)},
		{value: 10, baseType: basetype.Uint8z, result: uint8(10)},
		{value: 10, baseType: basetype.Sint16, result: int16(10)},
		{value: 10, baseType: basetype.Uint16, result: uint16(10)},
		{value: 10, baseType: basetype.Sint32, result: int32(10)},
		{value: 10, baseType: basetype.Uint32, result: uint32(10)},
		{value: 10, baseType: basetype.Float32, result: float32(10)},
		{value: 10, baseType: basetype.Float64, result: float64(10)},
		{value: 10, baseType: basetype.Sint64, result: int64(10)},
		{value: 10, baseType: basetype.Uint64, result: uint64(10)},
		{value: 10, baseType: basetype.Enum, result: basetype.Enum.Invalid()},
		{value: 10, baseType: basetype.String, result: basetype.String.Invalid()},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.Int64ToNumber(tc.value, tc.baseType)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
