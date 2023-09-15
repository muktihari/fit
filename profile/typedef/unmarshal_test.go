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
		value any
		ref   basetype.BaseType
		err   error
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
		{value: []byte{1, 1}, ref: basetype.Enum},
		{value: []byte{1, 2}, ref: basetype.Byte},
		{value: []int8{1, 3}, ref: basetype.Sint8},
		{value: []int8{1, -3}, ref: basetype.Sint8},
		{value: []uint8{1, 4}, ref: basetype.Uint8},
		{value: []uint8{1, 5}, ref: basetype.Uint8z},
		{value: []int16{1, 6}, ref: basetype.Sint16},
		{value: []int16{1, -6}, ref: basetype.Sint16},
		{value: []uint16{1, 7}, ref: basetype.Uint16},
		{value: []uint16{1, 8}, ref: basetype.Uint16z},
		{value: []int32{1, 9}, ref: basetype.Sint32},
		{value: []int32{1, -9}, ref: basetype.Sint32},
		{value: []uint32{1, 1}, ref: basetype.Uint32},
		{value: []uint32{1, 1}, ref: basetype.Uint32z},
		{value: []int64{1, 1}, ref: basetype.Sint64},
		{value: []int64{1, -2}, ref: basetype.Sint64},
		{value: []uint64{1, 1}, ref: basetype.Uint64},
		{value: []uint64{1, 1}, ref: basetype.Uint64z},
		{value: []float32{1, 1.1}, ref: basetype.Float32},
		{value: []float32{1, -5.1}, ref: basetype.Float32},
		{value: []float64{1, 1.1}, ref: basetype.Float64},
		{value: []float64{1, -5.1}, ref: basetype.Float64},
		{value: int8(0), ref: basetype.FromString("invalid"), err: typedef.ErrTypeNotSupported},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v)", tc.value, tc.value), func(t *testing.T) {
			b, err := typedef.Marshal(tc.value, binary.LittleEndian)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}

			v, err := typedef.Unmarshal(b, binary.LittleEndian, tc.ref)
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Fatalf("expected err: %v, got: %v", tc.err, err)
				}
				return
			}
			if diff := cmp.Diff(v, tc.value); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
