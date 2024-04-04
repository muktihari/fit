// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
)

func TestTypeString(t *testing.T) {
	for i, typStr := range typeStrings {
		t.Run(typStr, func(t *testing.T) {
			if str := Type(i).String(); str != typStr {
				t.Fatalf("expected: %s, got: %s", typStr, str)
			}
		})
	}
	invalid := fmt.Sprintf("proto.TypeInvalid(%d)", len(typeStrings)+1)
	if str := Type(len(typeStrings) + 1).String(); str != invalid {
		t.Fatalf("expected: %s, got: %s", invalid, str)
	}
}

func TestBool(t *testing.T) {
	input := bool(true)
	t.Run("valid", func(t *testing.T) {
		v := Bool(input)
		if v.Bool() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Bool())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Bool() != false {
			t.Fatalf("expected: %v, got: %v", false, v.Bool())
		}
	})
}

func TestInt8(t *testing.T) {
	input := int8(100)
	t.Run("valid", func(t *testing.T) {
		v := Int8(input)
		if v.Int8() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Int8())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Int8() != basetype.Sint8Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Sint8Invalid, v.Int8())
		}
	})
}

func TestUint8(t *testing.T) {
	input := uint8(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint8(input)
		if v.Uint8() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint8())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint8() != basetype.Uint8Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint8Invalid, v.Uint8())
		}
	})
}

func TestUint8z(t *testing.T) {
	input := uint8(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint8(input)
		if v.Uint8z() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint8z())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint8z() != basetype.Uint8zInvalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint8zInvalid, v.Uint8())
		}
	})
}

func TestInt16(t *testing.T) {
	input := int16(100)
	t.Run("valid", func(t *testing.T) {
		v := Int16(input)
		if v.Int16() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Int16())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Int16() != basetype.Sint16Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Sint16Invalid, v.Int16())
		}
	})
}

func TestUint16(t *testing.T) {
	input := uint16(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint16(input)
		if v.Uint16() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint16())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint16() != basetype.Uint16Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint16Invalid, v.Uint16())
		}
	})
}

func TestUint16z(t *testing.T) {
	input := uint16(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint16(input)
		if v.Uint16z() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint16z())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint16z() != basetype.Uint16zInvalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint16zInvalid, v.Uint16())
		}
	})
}

func TestInt32(t *testing.T) {
	input := int32(100)
	t.Run("valid", func(t *testing.T) {
		v := Int32(input)
		if v.Int32() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Int32())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Int32() != basetype.Sint32Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Sint32Invalid, v.Int32())
		}
	})
}

func TestUint32(t *testing.T) {
	input := uint32(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint32(input)
		if v.Uint32() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint32())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint32() != basetype.Uint32Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint32Invalid, v.Uint32())
		}
	})
}

func TestUint32z(t *testing.T) {
	input := uint32(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint32(input)
		if v.Uint32z() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint32z())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint32z() != basetype.Uint32zInvalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint32zInvalid, v.Uint32())
		}
	})
}

func TestInt64(t *testing.T) {
	input := int64(100)
	t.Run("valid", func(t *testing.T) {
		v := Int64(input)
		if v.Int64() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Int64())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Int64() != basetype.Sint64Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Sint64Invalid, v.Int64())
		}
	})
}

func TestUint64(t *testing.T) {
	input := uint64(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint64(input)
		if v.Uint64() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint64())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint64() != basetype.Uint64Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint64Invalid, v.Uint64())
		}
	})
}

func TestUint64z(t *testing.T) {
	input := uint64(100)
	t.Run("valid", func(t *testing.T) {
		v := Uint64(input)
		if v.Uint64z() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Uint64z())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.Uint64z() != basetype.Uint64zInvalid {
			t.Fatalf("expected: %v, got: %v", basetype.Uint64zInvalid, v.Uint64())
		}
	})
}

func TestFloat32(t *testing.T) {
	input := float32(100)
	t.Run("valid", func(t *testing.T) {
		v := Float32(input)
		if v.Float32() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Float32())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if math.Float32bits(v.Float32()) != basetype.Float32Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Float32Invalid, v.Float32())
		}
	})
}

func TestFloat64(t *testing.T) {
	input := float64(100)
	t.Run("valid", func(t *testing.T) {
		v := Float64(input)
		if v.Float64() != input {
			t.Fatalf("expected: %v, got: %v", input, v.Float64())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if math.Float64bits(v.Float64()) != basetype.Float64Invalid {
			t.Fatalf("expected: %v, got: %v", basetype.Float64Invalid, v.Float64())
		}
	})
}

func TestString(t *testing.T) {
	input := "fit"
	t.Run("valid", func(t *testing.T) {
		v := String(input)
		if v.String() != input {
			t.Fatalf("expected: %v, got: %v", input, v.String())
		}
	})
	t.Run("invalid", func(t *testing.T) {
		v := Value{}
		if v.String() != basetype.StringInvalid {
			t.Fatalf("expected: %v, got: %v", basetype.StringInvalid, v.String())
		}
	})
}

func TestSliceBool(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []bool{true, false}
		value := SliceBool(slice)
		result := value.SliceBool()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []test_bool{true, false}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]bool)(unsafe.Pointer(&slice))), len(slice))
		value := SliceBool(slice)
		result := value.SliceBool()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceBool()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceInt8(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []int8{1, 2}
		value := SliceInt8(slice)
		result := value.SliceInt8()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []int8{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]int8)(unsafe.Pointer(&slice))), len(slice))
		value := SliceInt8(slice)
		result := value.SliceInt8()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceInt8()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceUint8(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []uint8{1, 2}
		value := SliceUint8(slice)
		result := value.SliceUint8()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []uint8{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]uint8)(unsafe.Pointer(&slice))), len(slice))
		value := SliceUint8(slice)
		result := value.SliceUint8()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceUint8()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceInt16(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []int16{1, 2}
		value := SliceInt16(slice)
		result := value.SliceInt16()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []int16{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]int16)(unsafe.Pointer(&slice))), len(slice))
		value := SliceInt16(slice)
		result := value.SliceInt16()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceInt16()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceUint16(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []uint16{1, 2}
		value := SliceUint16(slice)
		result := value.SliceUint16()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []uint16{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]uint16)(unsafe.Pointer(&slice))), len(slice))
		value := SliceUint16(slice)
		result := value.SliceUint16()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceUint16()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceInt32(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []int32{1, 2}
		value := SliceInt32(slice)
		result := value.SliceInt32()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []int32{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]int32)(unsafe.Pointer(&slice))), len(slice))
		value := SliceInt32(slice)
		result := value.SliceInt32()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceInt32()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceUint32(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []uint32{1, 2}
		value := SliceUint32(slice)
		result := value.SliceUint32()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []uint32{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]uint32)(unsafe.Pointer(&slice))), len(slice))
		value := SliceUint32(slice)
		result := value.SliceUint32()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceUint32()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceInt64(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []int64{1, 2}
		value := SliceInt64(slice)
		result := value.SliceInt64()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []int64{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]int64)(unsafe.Pointer(&slice))), len(slice))
		value := SliceInt64(slice)
		result := value.SliceInt64()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceInt64()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceUint64(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []uint64{1, 2}
		value := SliceUint64(slice)
		result := value.SliceUint64()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []uint64{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]uint64)(unsafe.Pointer(&slice))), len(slice))
		value := SliceUint64(slice)
		result := value.SliceUint64()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceUint64()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceFloat32(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []float32{1, 2}
		value := SliceFloat32(slice)
		result := value.SliceFloat32()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []float32{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]float32)(unsafe.Pointer(&slice))), len(slice))
		value := SliceFloat32(slice)
		result := value.SliceFloat32()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceFloat32()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceFloat64(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []float64{1, 2}
		value := SliceFloat64(slice)
		result := value.SliceFloat64()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []float64{1, 2}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]float64)(unsafe.Pointer(&slice))), len(slice))
		value := SliceFloat64(slice)
		result := value.SliceFloat64()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceFloat64()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

func TestSliceString(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		slice := []string{"fit", "sdk"}
		value := SliceString(slice)
		result := value.SliceString()
		if diff := cmp.Diff(slice, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("correct custom type", func(t *testing.T) {
		slice := []test_string{"fit", "sdk"}
		expected := unsafe.Slice(unsafe.SliceData(*(*[]string)(unsafe.Pointer(&slice))), len(slice))
		value := SliceString(slice)
		result := value.SliceString()
		if len(slice) != len(result) { // compare result to original slice to ensure the cast is work as intended
			t.Fatalf("expected len: %d, got: %d", len(slice), len(result))
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Fatal(diff)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		value := Value{}
		result := value.SliceString()
		if result != nil {
			t.Fatalf("expected nil, got: %v", result)
		}
	})
}

type test_bool bool
type test_int8 int8
type test_uint8 uint8
type test_int16 int16
type test_uint16 uint16
type test_int32 int32
type test_uint32 uint32
type test_int64 int64
type test_uint64 uint64
type test_float32 float32
type test_float64 float64
type test_string string

func TestAny(t *testing.T) {
	tt := []struct {
		value    any
		expected any
	}{
		{value: int(0), expected: nil},
		{value: uint(0), expected: nil},
		{value: []any{}, expected: nil},
		{value: Bool(true), expected: true},
		{value: bool(true), expected: bool(true)},
		{value: bool(false), expected: bool(false)},
		{value: int8(10), expected: int8(10)},
		{value: uint8(10), expected: uint8(10)},
		{value: int16(10), expected: int16(10)},
		{value: uint16(10), expected: uint16(10)},
		{value: int32(10), expected: int32(10)},
		{value: uint32(10), expected: uint32(10)},
		{value: int64(10), expected: int64(10)},
		{value: uint64(10), expected: uint64(10)},
		{value: float32(10), expected: float32(10)},
		{value: float64(10), expected: float64(10)},
		{value: string("fit"), expected: string("fit")},
		{value: []bool{true, false}, expected: []bool{true, false}},
		{value: []int8{1, 2}, expected: []int8{1, 2}},
		{value: []uint8{1, 2}, expected: []uint8{1, 2}},
		{value: []int16{1, 2}, expected: []int16{1, 2}},
		{value: []uint16{1, 2}, expected: []uint16{1, 2}},
		{value: []int32{1, 2}, expected: []int32{1, 2}},
		{value: []uint32{1, 2}, expected: []uint32{1, 2}},
		{value: []int64{1, 2}, expected: []int64{1, 2}},
		{value: []uint64{1, 2}, expected: []uint64{1, 2}},
		{value: []float32{1, 2}, expected: []float32{1, 2}},
		{value: []float64{1, 2}, expected: []float64{1, 2}},
		{value: []string{"fit", "sdk"}, expected: []string{"fit", "sdk"}},
		{value: test_bool(true), expected: bool(true)},
		{value: test_bool(false), expected: bool(false)},
		{value: test_int8(10), expected: int8(10)},
		{value: test_uint8(10), expected: uint8(10)},
		{value: test_int16(10), expected: int16(10)},
		{value: test_uint16(10), expected: uint16(10)},
		{value: test_int32(10), expected: int32(10)},
		{value: test_uint32(10), expected: uint32(10)},
		{value: test_int64(10), expected: int64(10)},
		{value: test_uint64(10), expected: uint64(10)},
		{value: test_float32(10), expected: float32(10)},
		{value: test_float64(10), expected: float64(10)},
		{value: test_string("fit"), expected: string("fit")},
		{value: []test_int8{1, 2}, expected: []int8{1, 2}},
		{value: []test_uint8{1, 2}, expected: []uint8{1, 2}},
		{value: []test_int16{1, 2}, expected: []int16{1, 2}},
		{value: []test_uint16{1, 2}, expected: []uint16{1, 2}},
		{value: []test_int32{1, 2}, expected: []int32{1, 2}},
		{value: []test_uint32{1, 2}, expected: []uint32{1, 2}},
		{value: []test_int64{1, 2}, expected: []int64{1, 2}},
		{value: []test_uint64{1, 2}, expected: []uint64{1, 2}},
		{value: []test_float32{1, 2}, expected: []float32{1, 2}},
		{value: []test_float64{1, 2}, expected: []float64{1, 2}},
		{value: []test_string{"fit", "sdk"}, expected: []string{"fit", "sdk"}},
		{value: []test_string{}, expected: nil},
		{value: []struct{}{}, expected: nil},
		{value: []struct{}{{}}, expected: nil},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %T(%v)", i, tc.value, tc.value), func(t *testing.T) {
			value := Any(tc.value)
			if diff := cmp.Diff(value.Any(), tc.expected,
				cmp.Transformer("Value", func(v Value) any { return v.Any() }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tt := []struct {
		value       Value
		baseType    basetype.BaseType
		sizeInBytes int
	}{
		{value: Value{}, sizeInBytes: 0, baseType: basetype.Sint8},
		{value: Int8(10), sizeInBytes: 1, baseType: basetype.Sint8},
		{value: Uint8(10), sizeInBytes: 1, baseType: basetype.Uint8},
		{value: Int16(10), sizeInBytes: 2, baseType: basetype.Sint16},
		{value: Uint16(10), sizeInBytes: 2, baseType: basetype.Uint16},
		{value: Int32(10), sizeInBytes: 4, baseType: basetype.Sint32},
		{value: Uint32(10), sizeInBytes: 4, baseType: basetype.Uint32},
		{value: Float32(10), sizeInBytes: 4, baseType: basetype.Float32},
		{value: Float64(10), sizeInBytes: 8, baseType: basetype.Float64},
		{value: Int64(10), sizeInBytes: 8, baseType: basetype.Sint64},
		{value: Uint64(10), sizeInBytes: 8, baseType: basetype.Uint64},
		{value: SliceInt8([]int8{10, 9, 8, 7}), sizeInBytes: 4, baseType: basetype.Sint8},
		{value: SliceUint8([]uint8{10, 9, 8, 7}), sizeInBytes: 4, baseType: basetype.Uint8},
		{value: SliceInt16([]int16{10, 9, 8, 7}), sizeInBytes: 4 * 2, baseType: basetype.Uint16},
		{value: SliceUint16([]uint16{10, 9, 8, 7}), sizeInBytes: 4 * 2, baseType: basetype.Uint16},
		{value: SliceInt32([]int32{10, 9, 8, 7}), sizeInBytes: 4 * 4, baseType: basetype.Sint32},
		{value: SliceUint32([]uint32{10, 9, 8, 7}), sizeInBytes: 4 * 4, baseType: basetype.Uint32},
		{value: String(""), sizeInBytes: 1, baseType: basetype.String},
		{value: String("\x00"), sizeInBytes: 1, baseType: basetype.String},
		{value: String("fit sdk"), sizeInBytes: 8, baseType: basetype.String},
		{value: String("fit sdk\x00"), sizeInBytes: 8, baseType: basetype.String},
		{value: SliceString([]string{"fit sdk"}), sizeInBytes: 8, baseType: basetype.String},
		{value: SliceString([]string{}), sizeInBytes: 1, baseType: basetype.String},
		{value: SliceString([]string{""}), sizeInBytes: 1, baseType: basetype.String},
		{value: SliceString([]string{"\x00"}), sizeInBytes: 1, baseType: basetype.String},
		{value: SliceString([]string{"\x00\x00\x00"}), sizeInBytes: 3, baseType: basetype.String},
		{value: SliceString([]string{"\x00", "\x00", "\x00"}), sizeInBytes: 3, baseType: basetype.String},
		{value: SliceString([]string{"fit sdk", "a"}), sizeInBytes: 10, baseType: basetype.String},
		{value: SliceString([]string{"fit sdk\x00", "a\x00"}), sizeInBytes: 10, baseType: basetype.String},
		{value: SliceFloat32([]float32{10, 9, 8, 7}), sizeInBytes: 4 * 4, baseType: basetype.Float32},
		{value: SliceFloat64([]float64{10, 9, 8, 7}), sizeInBytes: 4 * 8, baseType: basetype.Float64},
		{value: SliceInt64([]int64{10, 9, 8, 7}), sizeInBytes: 4 * 8, baseType: basetype.Sint64},
		{value: SliceUint64([]uint64{10, 9, 8, 7}), sizeInBytes: 4 * 8, baseType: basetype.Uint64},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v)", tc.value, tc.value.Any()), func(t *testing.T) {
			size := Sizeof(tc.value, tc.baseType)
			if size != tc.sizeInBytes {
				t.Fatalf("expected: %d, got: %d", tc.sizeInBytes, size)
			}
		})
	}
}

func BenchmarkSliceBool(b *testing.B) {
	s := []bool{true, false}

	b.Run("[]bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := SliceBool(s)
			_ = v.SliceBool()
		}
	})
}
