// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"fmt"
	"testing"

	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/proto"
)

func TestBitsFromValue(t *testing.T) {
	tt := []struct {
		value    proto.Value
		expected uint32
		ok       bool
	}{
		{value: proto.Int8(10), expected: 10, ok: true},
		{value: proto.Uint8(10), expected: 10, ok: true},
		{value: proto.Int16(10), expected: 10, ok: true},
		{value: proto.Uint16(10), expected: 10, ok: true},
		{value: proto.Int32(10), expected: 10, ok: true},
		{value: proto.Uint32(10), expected: 10, ok: true},
		{value: proto.Int64(10), expected: 10, ok: true},
		{value: proto.Uint64(10), expected: 10, ok: true},
		{value: proto.Float32(10), expected: 10, ok: true},
		{value: proto.Float64(10), expected: 10, ok: true},
		{value: proto.SliceUint8([]byte{1, 1, 1}), expected: 1<<0 | 1<<8 | 1<<16, ok: true},
		{value: proto.SliceUint8([]byte{1, 255, 1}), expected: 0, ok: false},
		{value: proto.SliceUint8(make([]byte, 33)), expected: 0, ok: false},
		{value: proto.String("string value"), expected: 0, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v (%T)", tc.value, tc.value), func(t *testing.T) {
			res, ok := bitsFromValue(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected ok: %t, got: %t", tc.ok, ok)
			}
			if res != tc.expected {
				t.Fatalf("expected: %d, got: %d", tc.expected, res)
			}
		})
	}
}

func TestValueFromBits(t *testing.T) {
	tt := []struct {
		bitsVal  uint32
		basetype basetype.BaseType
		expected proto.Value
	}{
		{bitsVal: 10, basetype: basetype.Sint8, expected: proto.Int8(10)},
		{bitsVal: 10, basetype: basetype.Uint8, expected: proto.Uint8(10)},
		{bitsVal: 10, basetype: basetype.Sint16, expected: proto.Int16(10)},
		{bitsVal: 10, basetype: basetype.Uint16, expected: proto.Uint16(10)},
		{bitsVal: 10, basetype: basetype.Sint32, expected: proto.Int32(10)},
		{bitsVal: 10, basetype: basetype.Uint32, expected: proto.Uint32(10)},
		{bitsVal: 10, basetype: basetype.Sint64, expected: proto.Int64(10)},
		{bitsVal: 10, basetype: basetype.Uint64, expected: proto.Uint64(10)},
		{bitsVal: 10, basetype: basetype.Float32, expected: proto.Float32(10)},
		{bitsVal: 10, basetype: basetype.Float64, expected: proto.Float64(10)},
		{bitsVal: 10, basetype: basetype.String, expected: proto.Value{}},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s %v (%T)", i, tc.basetype, tc.expected.Any(), tc.expected.Any()), func(t *testing.T) {
			res := valueFromBits(tc.bitsVal, tc.basetype)
			if res.Any() != tc.expected.Any() {
				t.Fatalf("expected: %v, got: %v", tc.expected.Any(), res.Any())
			}
		})
	}
}

const v = 1080122531

func BenchmarkValueFromBits(b *testing.B) {
	val := proto.Float64(v)
	for i := 0; i < b.N; i++ {
		r := valueFromBits(v, basetype.Float64)
		if r.Float64() != val.Float64() {
			b.Fatalf("expected: %T: %v, got: %T: %v", v, v, r, r)
		}
	}
}
