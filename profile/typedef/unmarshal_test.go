// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef_test

import (
	"encoding/binary"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
)

func TestUnmarshal(t *testing.T) {
	tt := []struct {
		value    any
		ref      basetype.BaseType
		isArray  bool
		expected any // if nil, expected = value
		err      error
	}{
		{value: byte(1), ref: basetype.Enum},
		{value: byte(2), ref: basetype.Byte},
		{value: int8(3), ref: basetype.Sint8},
		{value: int8(-3), ref: basetype.Sint8},
		{value: uint8(4), ref: basetype.Uint8},
		{value: uint8(5), ref: basetype.Uint8z},
		{value: int16(6), ref: basetype.Sint16},
		{value: int16(-6), ref: basetype.Sint16},
		{value: uint16(7), ref: basetype.Uint16},
		{value: uint16(8), ref: basetype.Uint16z},
		{value: int32(9), ref: basetype.Sint32},
		{value: int32(-9), ref: basetype.Sint32},
		{value: uint32(10), ref: basetype.Uint32},
		{value: uint32(11), ref: basetype.Uint32z},
		{value: int64(12), ref: basetype.Sint64},
		{value: int64(-12), ref: basetype.Sint64},
		{value: uint64(13), ref: basetype.Uint64},
		{value: uint64(14), ref: basetype.Uint64z},
		{value: float32(15.1), ref: basetype.Float32},
		{value: float32(-15.1), ref: basetype.Float32},
		{value: float64(15.1), ref: basetype.Float64},
		{value: float64(-15.1), ref: basetype.Float64},
		{value: string("Fit SDK"), ref: basetype.String},
		{value: string(""), ref: basetype.String},
		{value: []byte{1, 1}, ref: basetype.Enum, isArray: true},
		{value: []byte{1, 2}, ref: basetype.Byte, isArray: true},
		{value: []int8{1, 3}, ref: basetype.Sint8, isArray: true},
		{value: []int8{1, -3}, ref: basetype.Sint8, isArray: true},
		{value: []uint8{1, 4}, ref: basetype.Uint8, isArray: true},
		{value: []uint8{1, 5}, ref: basetype.Uint8z, isArray: true},
		{value: []int16{1, 6}, ref: basetype.Sint16, isArray: true},
		{value: []int16{1, -6}, ref: basetype.Sint16, isArray: true},
		{value: []uint16{1, 7}, ref: basetype.Uint16, isArray: true},
		{value: []uint16{1, 8}, ref: basetype.Uint16z, isArray: true},
		{value: []int32{1, 9}, ref: basetype.Sint32, isArray: true},
		{value: []int32{1, -9}, ref: basetype.Sint32, isArray: true},
		{value: []uint32{1, 1}, ref: basetype.Uint32, isArray: true},
		{value: []uint32{1, 1}, ref: basetype.Uint32z, isArray: true},
		{value: []int64{1, 1}, ref: basetype.Sint64, isArray: true},
		{value: []int64{1, -2}, ref: basetype.Sint64, isArray: true},
		{value: []uint64{1, 1}, ref: basetype.Uint64, isArray: true},
		{value: []uint64{1, 1}, ref: basetype.Uint64z, isArray: true},
		{value: []float32{1, 1.1}, ref: basetype.Float32, isArray: true},
		{value: []float32{1, -5.1}, ref: basetype.Float32, isArray: true},
		{value: []float64{1, 1.1}, ref: basetype.Float64, isArray: true},
		{value: []float64{1, -5.1}, ref: basetype.Float64, isArray: true},
		{value: []string{"a", "b"}, ref: basetype.String, isArray: true},
		{
			value:    stringsToBytes([]string{"mobile_app_version", "\x00", "\x00"}...),
			expected: []string{"mobile_app_version"},
			ref:      basetype.String,
			isArray:  true,
		},
		{value: int8(0), ref: basetype.FromString("invalid"), err: typedef.ErrTypeNotSupported},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v)", tc.value, tc.value), func(t *testing.T) {
			b, err := typedef.Marshal(tc.value, binary.LittleEndian)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}

			v, err := typedef.Unmarshal(b, binary.LittleEndian, tc.ref, tc.isArray)
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Fatalf("expected err: %v, got: %v", tc.err, err)
				}
				return
			}

			if tc.expected == nil {
				tc.expected = tc.value
			}
			if diff := cmp.Diff(v, tc.expected); diff != "" {
				t.Fatal(diff)
			}

			// Extra check for bytes, the value should be copied
			if in, ok := tc.value.([]byte); ok {
				out, ok := v.([]byte)
				if !ok {
					return
				}

				in[0] = 255
				out[0] = 100

				if in[0] == out[0] {
					t.Fatalf("slice of bytes should not be referenced")
				}
			}
		})
	}
}

func stringsToBytes(vals ...string) []byte {
	b := []byte{}
	for i := range vals {
		b = append(b, vals[i]...)
	}
	return b
}
