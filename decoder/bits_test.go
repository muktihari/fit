// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/proto"
)

func TestMakeBitValue(t *testing.T) {
	tt := []struct {
		value    proto.Value
		expected bitvalue
		ok       bool
	}{
		{
			value:    proto.Int8(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint8(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Int16(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint16(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Int32(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint32(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Int64(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Uint64(10),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Float32(10.5),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.Float64(12.9),
			expected: bitvalue{[32]uint64{12}}, ok: true,
		},
		{
			value:    proto.SliceInt8([]int8{10}),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value: proto.SliceUint8(func() []uint8 {
				var b []uint8
				b = binary.LittleEndian.AppendUint64(b, 10)
				b = binary.LittleEndian.AppendUint64(b, 15)
				return b
			}()),
			expected: bitvalue{[32]uint64{10, 15}}, ok: true,
		},
		{
			value:    proto.SliceInt16([]int16{10}),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value: proto.SliceUint16([]uint16{10, 25, 55, 11, 12, 13, 14, 15}),
			ok:    true,
			expected: bitvalue{[32]uint64{
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
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceUint32([]uint32{10}),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceInt64([]int64{10}),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceUint64([]uint64{10}),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceFloat32([]float32{10.5}),
			expected: bitvalue{[32]uint64{10}}, ok: true,
		},
		{
			value:    proto.SliceFloat64([]float64{12.9}),
			expected: bitvalue{[32]uint64{12}}, ok: true,
		},
		{
			value:    proto.String("invalid"),
			expected: bitvalue{[32]uint64{31: math.MaxUint64}}, ok: false,
		},
		{
			value:    proto.Value{},
			expected: bitvalue{[32]uint64{31: math.MaxUint64}}, ok: false,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.value.Type()), func(t *testing.T) {
			bitVal, ok := makeBitValue(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected ok: %t, got: %t", tc.ok, ok)
			}
			if bitVal != tc.expected {
				t.Fatalf("expected bitVal: %v, got: %v", tc.expected, bitVal)
			}
		})
	}
}

func TestMakeBitFromUint32(t *testing.T) {
	tt := []struct {
		u32      uint32
		expected bitvalue
	}{
		{u32: 20, expected: bitvalue{store: [32]uint64{20}}},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.u32), func(t *testing.T) {
			bitVal := makeBitValueFromUint32(tc.u32)
			if bitVal != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, bitVal)
			}
		})
	}
}

func TestBitValueAsUint32(t *testing.T) {
	tt := []struct {
		bitVal   bitvalue
		expected uint32
	}{
		{bitVal: bitvalue{store: [32]uint64{500}}, expected: 500},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.bitVal.store[0]), func(t *testing.T) {
			u32 := tc.bitVal.AsUint32()
			if u32 != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, u32)
			}
		})
	}
}

func TestBitValueToValue(t *testing.T) {
	tt := []struct {
		bitVal   bitvalue
		baseType basetype.BaseType
		expected proto.Value
	}{
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Sint8, expected: proto.Int8(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Uint8, expected: proto.Uint8(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Sint16, expected: proto.Int16(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Uint16, expected: proto.Uint16(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Sint32, expected: proto.Int32(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Uint32, expected: proto.Uint32(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Sint64, expected: proto.Int64(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Uint64, expected: proto.Uint64(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Float32, expected: proto.Float32(40)},
		{bitVal: bitvalue{store: [32]uint64{40}}, baseType: basetype.Float64, expected: proto.Float64(40)},
		{bitVal: bitvalue{store: [32]uint64{10}}, baseType: basetype.String, expected: proto.Value{}},
		{bitVal: bitvalue{store: [32]uint64{31: math.MaxUint64}}, baseType: basetype.Float64, expected: proto.Value{}},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.expected.Type()), func(t *testing.T) {
			v := tc.bitVal.ToValue(tc.baseType)
			if diff := cmp.Diff(v, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) any { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestBitValuePullByMask(t *testing.T) {
	type pull struct {
		bits   byte
		value  uint32
		ok     bool
		bitVal bitvalue
	}
	tt := []struct {
		name   string
		bitVal bitvalue
		pulls  []pull
	}{
		{
			name:   "single value one pull",
			bitVal: bitvalue{store: [32]uint64{20}},
			pulls:  []pull{{bits: 8, value: 20, ok: true, bitVal: bitvalue{store: [32]uint64{0}}}},
		},
		{
			name:   "single value multiple pull",
			bitVal: bitvalue{store: [32]uint64{math.MaxUint16}},
			pulls: []pull{
				{bits: 8, value: 255, ok: true, bitVal: bitvalue{store: [32]uint64{255}}},
				{bits: 8, value: 255, ok: true, bitVal: bitvalue{store: [32]uint64{0}}},
			},
		},
		{
			name:   "slice value one pull",
			bitVal: bitvalue{store: [32]uint64{math.MaxUint16, math.MaxUint16}},
			pulls: []pull{
				{bits: 8, value: 255, ok: true, bitVal: bitvalue{store: [32]uint64{math.MaxUint16, 255}}},
			},
		},
		{
			name:   "slice value multiple pull",
			bitVal: bitvalue{store: [32]uint64{math.MaxUint16, math.MaxUint16}},
			pulls: []pull{
				{bits: 8, value: 255, ok: true, bitVal: bitvalue{store: [32]uint64{math.MaxUint16, 255}}},
				{bits: 8, value: 255, ok: true, bitVal: bitvalue{store: [32]uint64{math.MaxUint16}}},
			},
		},
		{
			name:   "single value one pull bits > 32 (64)",
			bitVal: bitvalue{store: [32]uint64{20}},
			pulls:  []pull{{bits: 64, value: 0, ok: false, bitVal: bitvalue{store: [32]uint64{20}}}},
		},
		{
			name:   "single value one pull store is zero",
			bitVal: bitvalue{store: [32]uint64{0}},
			pulls:  []pull{{bits: 8, value: 0, ok: false, bitVal: bitvalue{store: [32]uint64{0}}}},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			for _, p := range tc.pulls {
				u32, ok := tc.bitVal.PullByMask(p.bits)
				if ok != p.ok {
					t.Fatalf("expected ok: %t, got: %t", p.ok, ok)
				}
				if u32 != p.value {
					t.Fatalf("expected value: %t, got: %t", p.ok, ok)
				}
				if tc.bitVal != p.bitVal {
					t.Fatalf("expected bitVal: %v, got: %v", tc.bitVal, tc.bitVal)
				}
			}
		})
	}
}

func TestValueAppend(t *testing.T) {
	tt := []struct {
		slice    proto.Value
		elem     proto.Value
		expected proto.Value
	}{
		{
			slice:    proto.SliceInt8([]int8{10}),
			elem:     proto.Int8(11),
			expected: proto.SliceInt8([]int8{10, 11}),
		},
		{
			slice:    proto.SliceUint8([]uint8{10}),
			elem:     proto.Uint8(11),
			expected: proto.SliceUint8([]uint8{10, 11}),
		},
		{
			slice:    proto.SliceInt16([]int16{10}),
			elem:     proto.Int16(11),
			expected: proto.SliceInt16([]int16{10, 11}),
		},
		{
			slice:    proto.SliceUint16([]uint16{10}),
			elem:     proto.Uint16(11),
			expected: proto.SliceUint16([]uint16{10, 11}),
		},
		{
			slice:    proto.SliceInt32([]int32{10}),
			elem:     proto.Int32(11),
			expected: proto.SliceInt32([]int32{10, 11}),
		},
		{
			slice:    proto.SliceUint32([]uint32{10}),
			elem:     proto.Uint32(11),
			expected: proto.SliceUint32([]uint32{10, 11}),
		},
		{
			slice:    proto.SliceInt64([]int64{10}),
			elem:     proto.Int64(11),
			expected: proto.SliceInt64([]int64{10, 11}),
		},
		{
			slice:    proto.SliceUint64([]uint64{10}),
			elem:     proto.Uint64(11),
			expected: proto.SliceUint64([]uint64{10, 11}),
		},
		{
			slice:    proto.SliceFloat32([]float32{10}),
			elem:     proto.Float32(11),
			expected: proto.SliceFloat32([]float32{10, 11}),
		},
		{
			slice:    proto.SliceFloat64([]float64{10}),
			elem:     proto.Float64(11),
			expected: proto.SliceFloat64([]float64{10, 11}),
		},
		{
			slice:    proto.SliceString([]string{"invalid"}),
			elem:     proto.String("qwerty"),
			expected: proto.SliceString([]string{"invalid"}),
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.slice.Type()), func(t *testing.T) {
			val := valueAppend(tc.slice, tc.elem)
			if diff := cmp.Diff(val, tc.expected,
				cmp.Transformer("Value", func(v proto.Value) any { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func BenchmarkMakeBitValueToValueRoundTrip(b *testing.B) {
	b.Run("to float64", func(b *testing.B) {
		const v = 1080122531
		val := proto.Float64(v)
		for i := 0; i < b.N; i++ {
			bitVal, _ := makeBitValue(val)
			if bitVal.ToValue(basetype.Float64).Float64() != val.Float64() {
				b.Fatalf("expected: %T: %v, got: %T: %v", v, v, bitVal, bitVal)
			}
		}
	})
	b.Run("to bytes", func(b *testing.B) {
		var buf []uint8
		buf = binary.LittleEndian.AppendUint64(buf, 10)
		buf = binary.LittleEndian.AppendUint64(buf, 11)
		val := proto.SliceUint8(buf)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bitVal, _ := makeBitValue(val)
			_ = bitVal.ToValue(basetype.Byte).SliceUint8()
		}
	})
}
