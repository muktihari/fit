// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef_test

import (
	"fmt"
	"testing"

	"github.com/muktihari/fit/profile/typedef"
)

type Uint8 uint8

func TestLen(t *testing.T) {
	tt := []struct {
		value any
		size  byte
	}{
		{value: int8(10), size: 1},
		{value: uint8(10), size: 1},
		{value: int16(10), size: 1},
		{value: uint16(10), size: 1},
		{value: int32(10), size: 1},
		{value: uint32(10), size: 1},
		{value: float32(10), size: 1},
		{value: float64(10), size: 1},
		{value: int64(10), size: 1},
		{value: uint64(10), size: 1},
		{value: []int8{10, 9, 8, 7}, size: 4},
		{value: []uint8{10, 9, 8, 7}, size: 4},
		{value: []int16{10, 9, 8, 7}, size: 4},
		{value: []uint16{10, 9, 8, 7}, size: 4},
		{value: []int32{10, 9, 8, 7}, size: 4},
		{value: []uint32{10, 9, 8, 7}, size: 4},
		{value: string("fit sdk"), size: 7},
		{value: []float32{10, 9, 8, 7}, size: 4},
		{value: []float64{10, 9, 8, 7}, size: 4},
		{value: []int64{10, 9, 8, 7}, size: 4},
		{value: []uint64{10, 9, 8, 7}, size: 4},
		{value: Uint8(7), size: 1},
		{value: []Uint8{10, 9, 8, 7}, size: 4},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v)", tc.value, tc.value), func(t *testing.T) {
			size := typedef.Len(tc.value)
			if size != tc.size {
				t.Fatalf("expected: %d, got: %d", tc.size, size)
			}
		})
	}
}
