// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

func TestUnmarshalValue(t *testing.T) {
	tt := []struct {
		value       proto.Value
		baseType    basetype.BaseType
		profileType profile.ProfileType
		isArray     bool
		expected    proto.Value // if nil, expected = value
		err         error
	}{
		{value: proto.Uint8(1), baseType: basetype.Enum, profileType: profile.Enum},
		{value: proto.Uint8(2), baseType: basetype.Byte, profileType: profile.Byte},
		{value: proto.Int8(3), baseType: basetype.Sint8, profileType: profile.Sint8},
		{value: proto.Int8(-3), baseType: basetype.Sint8, profileType: profile.Sint8},
		{value: proto.Uint8(4), baseType: basetype.Uint8, profileType: profile.Uint8},
		{value: proto.Uint8(5), baseType: basetype.Uint8z, profileType: profile.Uint8z},
		{value: proto.Int16(6), baseType: basetype.Sint16, profileType: profile.Sint16},
		{value: proto.Int16(-6), baseType: basetype.Sint16, profileType: profile.Sint16},
		{value: proto.Uint16(7), baseType: basetype.Uint16, profileType: profile.Uint16},
		{value: proto.Uint16(8), baseType: basetype.Uint16z, profileType: profile.Uint16z},
		{value: proto.Int32(9), baseType: basetype.Sint32, profileType: profile.Sint32},
		{value: proto.Int32(-9), baseType: basetype.Sint32, profileType: profile.Sint32},
		{value: proto.Uint32(10), baseType: basetype.Uint32, profileType: profile.Uint32},
		{value: proto.Uint32(11), baseType: basetype.Uint32z, profileType: profile.Uint32z},
		{value: proto.Int64(12), baseType: basetype.Sint64, profileType: profile.Sint64},
		{value: proto.Int64(-12), baseType: basetype.Sint64, profileType: profile.Sint64},
		{value: proto.Uint64(13), baseType: basetype.Uint64, profileType: profile.Uint64},
		{value: proto.Uint64(14), baseType: basetype.Uint64z, profileType: profile.Uint64z},
		{value: proto.Float32(15.1), baseType: basetype.Float32, profileType: profile.Float32},
		{value: proto.Float32(-15.1), baseType: basetype.Float32, profileType: profile.Float32},
		{value: proto.Float64(15.1), baseType: basetype.Float64, profileType: profile.Float64},
		{value: proto.Float64(-15.1), baseType: basetype.Float64, profileType: profile.Float64},
		{value: proto.String("FIT SDK"), baseType: basetype.String, profileType: profile.String},
		{value: proto.String(""), baseType: basetype.String, profileType: profile.String},
		{value: proto.SliceUint8([]byte{1, 1}), baseType: basetype.Enum, profileType: profile.Enum, isArray: true},
		{value: proto.SliceUint8([]byte{1, 2}), baseType: basetype.Byte, profileType: profile.Byte, isArray: true},
		{value: proto.SliceInt8([]int8{1, 3}), baseType: basetype.Sint8, profileType: profile.Sint8, isArray: true},
		{value: proto.SliceInt8([]int8{1, -3}), baseType: basetype.Sint8, profileType: profile.Sint8, isArray: true},
		{value: proto.SliceUint8([]uint8{1, 4}), baseType: basetype.Uint8, profileType: profile.Uint8, isArray: true},
		{value: proto.SliceUint8([]uint8{1, 5}), baseType: basetype.Uint8z, profileType: profile.Uint8z, isArray: true},
		{value: proto.SliceInt16([]int16{1, 6}), baseType: basetype.Sint16, profileType: profile.Sint16, isArray: true},
		{value: proto.SliceInt16([]int16{1, -6}), baseType: basetype.Sint16, profileType: profile.Sint16, isArray: true},
		{value: proto.SliceUint16([]uint16{1, 7}), baseType: basetype.Uint16, profileType: profile.Uint16, isArray: true},
		{value: proto.SliceUint16([]uint16{1, 8}), baseType: basetype.Uint16z, profileType: profile.Uint16z, isArray: true},
		{value: proto.SliceInt32([]int32{1, 9}), baseType: basetype.Sint32, profileType: profile.Sint32, isArray: true},
		{value: proto.SliceInt32([]int32{1, -9}), baseType: basetype.Sint32, profileType: profile.Sint32, isArray: true},
		{value: proto.SliceUint32([]uint32{1, 1}), baseType: basetype.Uint32, profileType: profile.Uint32, isArray: true},
		{value: proto.SliceUint32([]uint32{1, 1}), baseType: basetype.Uint32z, profileType: profile.Uint32z, isArray: true},
		{value: proto.SliceInt64([]int64{1, 1}), baseType: basetype.Sint64, profileType: profile.Sint64, isArray: true},
		{value: proto.SliceInt64([]int64{1, -2}), baseType: basetype.Sint64, profileType: profile.Sint64, isArray: true},
		{value: proto.SliceUint64([]uint64{1, 1}), baseType: basetype.Uint64, profileType: profile.Uint64, isArray: true},
		{value: proto.SliceUint64([]uint64{1, 1}), baseType: basetype.Uint64z, profileType: profile.Uint64z, isArray: true},
		{value: proto.SliceFloat32([]float32{1, 1.1}), baseType: basetype.Float32, profileType: profile.Float32, isArray: true},
		{value: proto.SliceFloat32([]float32{1, -5.1}), baseType: basetype.Float32, profileType: profile.Float32, isArray: true},
		{value: proto.SliceFloat64([]float64{1, 1.1}), baseType: basetype.Float64, profileType: profile.Float64, isArray: true},
		{value: proto.SliceFloat64([]float64{1, -5.1}), baseType: basetype.Float64, profileType: profile.Float64, isArray: true},
		{value: proto.SliceString([]string{"a", "b"}), baseType: basetype.String, profileType: profile.String, isArray: true},
		{
			value:    proto.SliceUint8(stringsToBytes([]string{"mobile_app_version", "\x00", "\x00"}...)),
			expected: proto.SliceString([]string{"mobile_app_version"}),
			baseType: basetype.String,
			isArray:  true,
		},
		{value: proto.Int8(0), baseType: basetype.FromString("invalid"), err: proto.ErrTypeNotSupported},
		{value: proto.Bool(typedef.BoolTrue), baseType: basetype.Enum, profileType: profile.Bool},
		{value: proto.SliceBool([]typedef.Bool{typedef.BoolFalse, typedef.BoolTrue, typedef.BoolFalse}), baseType: basetype.Enum, profileType: profile.Bool, isArray: true},
	}

	for i, tc := range tt {
		for arch := byte(0); arch <= 1; arch++ {
			t.Run(fmt.Sprintf("[%d] %T(%v)", i, tc.value.Any(), tc.value.Any()), func(t *testing.T) {
				b, err := tc.value.MarshalAppend(nil, arch)
				if err != nil {
					t.Fatalf("marshal failed: %v", err)
				}

				v, err := proto.UnmarshalValue(b, arch, tc.baseType, tc.profileType, tc.isArray)
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
}

func stringsToBytes(vals ...string) []byte {
	b := []byte{}
	for i := range vals {
		b = append(b, vals[i]...)
	}
	return b
}

func TestUnmarshalValueSliceAlloc(t *testing.T) {
	tt := []struct {
		value       proto.Value
		profileType profile.ProfileType
	}{
		{value: proto.SliceBool(make([]typedef.Bool, 256)), profileType: profile.Bool},
		{value: proto.SliceUint8(make([]uint8, 256)), profileType: profile.Uint8},
		{value: proto.SliceInt8(make([]int8, 256)), profileType: profile.Sint8},
		{value: proto.SliceUint8(make([]uint8, 256)), profileType: profile.Uint8},
		{value: proto.SliceInt16(make([]int16, 256)), profileType: profile.Sint16},
		{value: proto.SliceUint16(make([]uint16, 256)), profileType: profile.Uint16},
		{value: proto.SliceInt32(make([]int32, 256)), profileType: profile.Sint32},
		{value: proto.SliceUint32(make([]uint32, 256)), profileType: profile.Uint32},
		{value: proto.SliceInt64(make([]int64, 256)), profileType: profile.Sint64},
		{value: proto.SliceUint64(make([]uint64, 256)), profileType: profile.Uint64},
		{value: proto.SliceFloat32(make([]float32, 256)), profileType: profile.Float32},
		{value: proto.SliceFloat64(make([]float64, 256)), profileType: profile.Float64},
		{value: proto.SliceString(make([]string, 256)), profileType: profile.String},
	}

	for i, tc := range tt {
		for arch := byte(0); arch < 2; arch++ {
			t.Run(fmt.Sprintf("[%d] arch: %d: %v", i, arch, tc.profileType.String()), func(t *testing.T) {
				b, _ := tc.value.MarshalAppend(nil, arch)
				alloc := testing.AllocsPerRun(10, func() {
					_, _ = proto.UnmarshalValue(b, arch, tc.profileType.BaseType(), tc.profileType, true)
				})
				if alloc > 1 {
					t.Fatalf("expected alloc: 1, got: %g", alloc)
				}
			})
		}
	}
}

func BenchmarkUnmarshalValue(b *testing.B) {
	b.StopTimer()
	v := proto.Uint32(100)
	buf, _ := v.MarshalAppend(nil, 0)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = proto.UnmarshalValue(buf, 0, basetype.Uint32, profile.Uint32, false)
	}
}
