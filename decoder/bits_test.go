// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"

	"github.com/muktihari/fit/proto"
)

func TestMakeBits(t *testing.T) {
	tt := []struct {
		value    proto.Value
		expected bits
		ok       bool
	}{
		{
			value:    proto.Int8(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint8(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Int16(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint16(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Int32(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint32(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Int64(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint64(10),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Float32(10.5),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Float64(12.9),
			expected: bits{[32]uint64{12}}, ok: true,
		},
		{
			value:    proto.SliceInt8([]int8{10}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value: proto.SliceUint8(func() []uint8 {
				var b []uint8
				b = binary.LittleEndian.AppendUint64(b, 10)
				b = binary.LittleEndian.AppendUint64(b, 15)
				return b
			}()),
			expected: bits{[32]uint64{10, 15}}, ok: true,
		},
		{
			value:    proto.SliceInt16([]int16{10}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value: proto.SliceUint16([]uint16{10, 25, 55, 11, 12, 13, 14, 15}),
			ok:    true,
			expected: bits{[32]uint64{
				func() uint64 {
					var b []uint8
					b = binary.LittleEndian.AppendUint16(b, 10)
					b = binary.LittleEndian.AppendUint16(b, 25)
					b = binary.LittleEndian.AppendUint16(b, 55)
					b = binary.LittleEndian.AppendUint16(b, 11)
					return binary.LittleEndian.Uint64(b)
				}(),
				func() uint64 {
					var b []uint8
					b = binary.LittleEndian.AppendUint16(b, 12)
					b = binary.LittleEndian.AppendUint16(b, 13)
					b = binary.LittleEndian.AppendUint16(b, 14)
					b = binary.LittleEndian.AppendUint16(b, 15)
					return binary.LittleEndian.Uint64(b)
				}(),
			}},
		},
		{
			value:    proto.SliceInt32([]int32{10}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceUint32([]uint32{10}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceInt64([]int64{10}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceUint64([]uint64{10}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceFloat32([]float32{10.5}),
			expected: bits{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceFloat64([]float64{12.9}),
			expected: bits{[32]uint64{12}}, ok: true,
		},
		{
			value:    proto.String("invalid"),
			expected: bits{}, ok: false,
		},
		{
			value:    proto.Value{},
			expected: bits{}, ok: false,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.value.Type()), func(t *testing.T) {
			vbits, ok := makeBits(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected ok: %t, got: %t", tc.ok, ok)
			}
			if vbits != tc.expected {
				t.Fatalf("expected bits:\n%v,\n got:\n%v", tc.expected, vbits)
			}
		})
	}
}

func TestBitsPull(t *testing.T) {
	type pull struct {
		bits  byte
		value uint32
		ok    bool
		vbits bits
	}
	tt := []struct {
		name  string
		vbits bits
		pulls []pull
	}{
		{
			name:  "single value one pull",
			vbits: bits{store: [32]uint64{20}},
			pulls: []pull{{bits: 8, value: 20, ok: true, vbits: bits{store: [32]uint64{0}}}},
		},
		{
			name:  "single value multiple pull",
			vbits: bits{store: [32]uint64{math.MaxUint16}},
			pulls: []pull{
				{bits: 8, value: 255, ok: true, vbits: bits{store: [32]uint64{255}}},
				{bits: 8, value: 255, ok: true, vbits: bits{store: [32]uint64{0}}},
			},
		},
		{
			name:  "slice value one pull",
			vbits: bits{store: [32]uint64{math.MaxUint64, math.MaxUint16}},
			pulls: []pull{
				{bits: 8, value: 255, ok: true, vbits: bits{store: [32]uint64{math.MaxUint64, 255}}},
			},
		},
		{
			name:  "slice value multiple pull",
			vbits: bits{store: [32]uint64{math.MaxUint64, math.MaxUint16}},
			pulls: []pull{
				{bits: 8, value: 255, ok: true, vbits: bits{store: [32]uint64{math.MaxUint64, 255}}},
				{bits: 8, value: 255, ok: true, vbits: bits{store: [32]uint64{math.MaxUint64}}},
			},
		},
		{
			name:  "single value one pull store is zero",
			vbits: bits{store: [32]uint64{0}},
			pulls: []pull{{bits: 8, value: 0, ok: false, vbits: bits{store: [32]uint64{0}}}},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			for _, p := range tc.pulls {
				u32, ok := tc.vbits.Pull(p.bits)
				if ok != p.ok {
					t.Fatalf("expected ok: %t, got: %t", p.ok, ok)
				}
				if u32 != p.value {
					t.Fatalf("expected value: %t, got: %t", p.ok, ok)
				}
				if tc.vbits != p.vbits {
					t.Fatalf("expected bits:\n%v,\n got:\n%v", tc.vbits, p.vbits)
				}
			}
		})
	}
}
