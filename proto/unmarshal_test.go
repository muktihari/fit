// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"encoding/binary"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/proto"
)

func TestUnmarshal(t *testing.T) {
	tt := []struct {
		value    proto.Value
		ref      basetype.BaseType
		isArray  bool
		expected proto.Value // if nil, expected = value
		err      error
	}{
		{value: proto.Uint8(1), ref: basetype.Enum},
		{value: proto.Uint8(2), ref: basetype.Byte},
		{value: proto.Int8(3), ref: basetype.Sint8},
		{value: proto.Int8(-3), ref: basetype.Sint8},
		{value: proto.Uint8(4), ref: basetype.Uint8},
		{value: proto.Uint8(5), ref: basetype.Uint8z},
		{value: proto.Int16(6), ref: basetype.Sint16},
		{value: proto.Int16(-6), ref: basetype.Sint16},
		{value: proto.Uint16(7), ref: basetype.Uint16},
		{value: proto.Uint16(8), ref: basetype.Uint16z},
		{value: proto.Int32(9), ref: basetype.Sint32},
		{value: proto.Int32(-9), ref: basetype.Sint32},
		{value: proto.Uint32(10), ref: basetype.Uint32},
		{value: proto.Uint32(11), ref: basetype.Uint32z},
		{value: proto.Int64(12), ref: basetype.Sint64},
		{value: proto.Int64(-12), ref: basetype.Sint64},
		{value: proto.Uint64(13), ref: basetype.Uint64},
		{value: proto.Uint64(14), ref: basetype.Uint64z},
		{value: proto.Float32(15.1), ref: basetype.Float32},
		{value: proto.Float32(-15.1), ref: basetype.Float32},
		{value: proto.Float64(15.1), ref: basetype.Float64},
		{value: proto.Float64(-15.1), ref: basetype.Float64},
		{value: proto.String("Fit SDK"), ref: basetype.String},
		{value: proto.String(""), ref: basetype.String},
		{value: proto.SliceUint8([]byte{1, 1}), ref: basetype.Enum, isArray: true},
		{value: proto.SliceUint8([]byte{1, 2}), ref: basetype.Byte, isArray: true},
		{value: proto.SliceInt8([]int8{1, 3}), ref: basetype.Sint8, isArray: true},
		{value: proto.SliceInt8([]int8{1, -3}), ref: basetype.Sint8, isArray: true},
		{value: proto.SliceUint8([]uint8{1, 4}), ref: basetype.Uint8, isArray: true},
		{value: proto.SliceUint8([]uint8{1, 5}), ref: basetype.Uint8z, isArray: true},
		{value: proto.SliceInt16([]int16{1, 6}), ref: basetype.Sint16, isArray: true},
		{value: proto.SliceInt16([]int16{1, -6}), ref: basetype.Sint16, isArray: true},
		{value: proto.SliceUint16([]uint16{1, 7}), ref: basetype.Uint16, isArray: true},
		{value: proto.SliceUint16([]uint16{1, 8}), ref: basetype.Uint16z, isArray: true},
		{value: proto.SliceInt32([]int32{1, 9}), ref: basetype.Sint32, isArray: true},
		{value: proto.SliceInt32([]int32{1, -9}), ref: basetype.Sint32, isArray: true},
		{value: proto.SliceUint32([]uint32{1, 1}), ref: basetype.Uint32, isArray: true},
		{value: proto.SliceUint32([]uint32{1, 1}), ref: basetype.Uint32z, isArray: true},
		{value: proto.SliceInt64([]int64{1, 1}), ref: basetype.Sint64, isArray: true},
		{value: proto.SliceInt64([]int64{1, -2}), ref: basetype.Sint64, isArray: true},
		{value: proto.SliceUint64([]uint64{1, 1}), ref: basetype.Uint64, isArray: true},
		{value: proto.SliceUint64([]uint64{1, 1}), ref: basetype.Uint64z, isArray: true},
		{value: proto.SliceFloat32([]float32{1, 1.1}), ref: basetype.Float32, isArray: true},
		{value: proto.SliceFloat32([]float32{1, -5.1}), ref: basetype.Float32, isArray: true},
		{value: proto.SliceFloat64([]float64{1, 1.1}), ref: basetype.Float64, isArray: true},
		{value: proto.SliceFloat64([]float64{1, -5.1}), ref: basetype.Float64, isArray: true},
		{value: proto.SliceString([]string{"a", "b"}), ref: basetype.String, isArray: true},
		{
			value:    proto.SliceUint8(stringsToBytes([]string{"mobile_app_version", "\x00", "\x00"}...)),
			expected: proto.SliceString([]string{"mobile_app_version"}),
			ref:      basetype.String,
			isArray:  true,
		},
		{value: proto.Int8(0), ref: basetype.FromString("invalid"), err: proto.ErrTypeNotSupported},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %T(%v)", i, tc.value.Any(), tc.value.Any()), func(t *testing.T) {
			b, err := proto.Marshal(tc.value, binary.LittleEndian)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}

			v, err := proto.Unmarshal(b, binary.LittleEndian, tc.ref, tc.isArray)
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Fatalf("expected err: %v, got: %v", tc.err, err)
				}
				return
			}

			if tc.expected.Type() == proto.TypeInvalid {
				tc.expected = tc.value
			}
			if diff := cmp.Diff(v, tc.expected,
				cmp.Transformer("Value", func(val proto.Value) any {
					return val.Any()
				}),
			); diff != "" {
				t.Fatal(diff)
			}

			// Extra check for bytes, the value should be copied
			if in := tc.value.SliceUint8(); in != nil {
				out := v.SliceUint8()
				if out == nil {
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
