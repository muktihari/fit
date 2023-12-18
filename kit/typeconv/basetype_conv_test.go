// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
)

type Int8 int8
type Uint8 uint8
type Int16 int16
type Uint16 uint16
type Int32 int32
type Uint32 uint32
type String string
type Float32 float32
type Float64 float64
type Int64 int64
type Uint64 uint64

func TestToEnum(t *testing.T) {
	value := any(uint8(10))
	v := typeconv.ToEnum[byte](value)
	if diff := cmp.Diff(v, value); diff != "" {
		t.Fatal(diff)
	}
}

func TestToByte(t *testing.T) {
	value := any(byte(10))
	v := typeconv.ToByte[byte](value)
	if diff := cmp.Diff(v, value); diff != "" {
		t.Fatal(diff)
	}
}

func TestToSint8(t *testing.T) {
	tt := []struct {
		value  any
		result int8
	}{
		{value: nil, result: basetype.Sint8Invalid},
		{value: int8(10), result: 10},
		{value: Int8(10), result: 10},
		{value: int16(10), result: basetype.Sint8Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToSint8[int8](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Int8", func(t *testing.T) {
				v := typeconv.ToSint8[Int8](tc.value)
				if diff := cmp.Diff(v, Int8(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint8(t *testing.T) {
	tt := []struct {
		value  any
		result uint8
	}{
		{value: nil, result: basetype.Uint8Invalid},
		{value: uint8(10), result: 10},
		{value: Uint8(10), result: 10},
		{value: int16(10), result: basetype.Uint8Invalid},
		{value: false, result: 0},
		{value: true, result: 1},
		{value: Bool(true), result: 1},
		{value: Bool(false), result: 0},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint8[uint8](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint8", func(t *testing.T) {
				v := typeconv.ToUint8[Uint8](tc.value)
				if diff := cmp.Diff(v, Uint8(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint8z(t *testing.T) {
	tt := []struct {
		value  any
		result uint8
	}{
		{value: nil, result: basetype.Uint8zInvalid},
		{value: uint8(10), result: 10},
		{value: Uint8(10), result: 10},
		{value: int16(10), result: basetype.Uint8zInvalid},
		{value: false, result: 0},
		{value: true, result: 1},
		{value: Bool(true), result: 1},
		{value: Bool(false), result: 0},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint8z[uint8](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint8", func(t *testing.T) {
				v := typeconv.ToUint8z[Uint8](tc.value)
				if diff := cmp.Diff(v, Uint8(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToSint16(t *testing.T) {
	tt := []struct {
		value  any
		result int16
	}{
		{value: nil, result: basetype.Sint16Invalid},
		{value: int16(10), result: 10},
		{value: Int16(10), result: 10},
		{value: int8(10), result: basetype.Sint16Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToSint16[int16](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Int16", func(t *testing.T) {
				v := typeconv.ToSint16[Int16](tc.value)
				if diff := cmp.Diff(v, Int16(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint16(t *testing.T) {
	tt := []struct {
		value  any
		result uint16
	}{
		{value: nil, result: basetype.Uint16Invalid},
		{value: uint16(10), result: 10},
		{value: Uint16(10), result: 10},
		{value: int8(10), result: basetype.Uint16Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint16[uint16](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint16", func(t *testing.T) {
				v := typeconv.ToUint16[Uint16](tc.value)
				if diff := cmp.Diff(v, Uint16(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint16z(t *testing.T) {
	tt := []struct {
		value  any
		result uint16
	}{
		{value: nil, result: basetype.Uint16zInvalid},
		{value: uint16(10), result: 10},
		{value: Uint16(10), result: 10},
		{value: int8(10), result: basetype.Uint16zInvalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint16z[uint16](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint16", func(t *testing.T) {
				v := typeconv.ToUint16z[Uint16](tc.value)
				if diff := cmp.Diff(v, Uint16(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToSint32(t *testing.T) {
	tt := []struct {
		value  any
		result int32
	}{
		{value: nil, result: basetype.Sint32Invalid},
		{value: int32(10), result: 10},
		{value: Int32(10), result: 10},
		{value: int8(10), result: basetype.Sint32Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToSint32[int32](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Int16", func(t *testing.T) {
				v := typeconv.ToSint32[Int32](tc.value)
				if diff := cmp.Diff(v, Int32(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint32(t *testing.T) {
	tt := []struct {
		value  any
		result uint32
	}{
		{value: nil, result: basetype.Uint32Invalid},
		{value: uint32(10), result: 10},
		{value: float32(math.Float32frombits(100)), result: 100},
		{value: Uint32(10), result: 10},
		{value: int8(10), result: basetype.Uint32Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint32[uint32](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint32", func(t *testing.T) {
				v := typeconv.ToUint32[Uint32](tc.value)
				if diff := cmp.Diff(v, Uint32(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint32z(t *testing.T) {
	tt := []struct {
		value  any
		result uint32
	}{
		{value: nil, result: basetype.Uint32zInvalid},
		{value: uint32(10), result: 10},
		{value: Uint32(10), result: 10},
		{value: int8(10), result: basetype.Uint32zInvalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint32z[uint32](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint32", func(t *testing.T) {
				v := typeconv.ToUint32z[Uint32](tc.value)
				if diff := cmp.Diff(v, Uint32(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToString(t *testing.T) {
	tt := []struct {
		value  any
		result string
	}{
		{value: nil, result: ""},
		{value: string("abc"), result: "abc"},
		{value: String("abc"), result: "abc"},
		{value: int8(10), result: ""},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToString[string](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint32", func(t *testing.T) {
				v := typeconv.ToString[String](tc.value)
				if diff := cmp.Diff(v, String(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToFloat32(t *testing.T) {
	tt := []struct {
		value  any
		result float32
	}{
		{value: nil, result: math.Float32frombits(basetype.Float32Invalid)},
		{value: float32(10.2), result: 10.2},
		{value: Float32(10.2), result: 10.2},
		{value: uint32(1093140480), result: 10.5},
		{value: Uint32(1093140480), result: 10.5},
		{value: int8(10), result: math.Float32frombits(basetype.Float32Invalid)},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToFloat32[float32](tc.value)
			if math.IsNaN(float64(v)) && math.IsNaN(float64(tc.result)) {
				return
			}

			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Float32", func(t *testing.T) {
				v := typeconv.ToFloat32[Float32](tc.value)
				if diff := cmp.Diff(v, Float32(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToFloat64(t *testing.T) {
	tt := []struct {
		value  any
		result float64
	}{
		{value: nil, result: math.Float64frombits(basetype.Float64Invalid)},
		{value: float64(10.2), result: 10.2},
		{value: Float64(10.2), result: 10.2},
		{value: uint64(4622100592565682176), result: 10.5},
		{value: Uint64(4622100592565682176), result: 10.5},
		{value: int8(10), result: math.Float64frombits(basetype.Float64Invalid)},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToFloat64[float64](tc.value)
			if math.IsNaN(float64(v)) && math.IsNaN(float64(tc.result)) {
				return
			}

			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Float64", func(t *testing.T) {
				v := typeconv.ToFloat64[Float64](tc.value)
				if diff := cmp.Diff(v, Float64(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToSint64(t *testing.T) {
	tt := []struct {
		value  any
		result int64
	}{
		{value: nil, result: basetype.Sint64Invalid},
		{value: int64(10), result: 10},
		{value: Int64(10), result: 10},
		{value: int8(10), result: basetype.Sint64Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToSint64[int64](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Int16", func(t *testing.T) {
				v := typeconv.ToSint64[Int64](tc.value)
				if diff := cmp.Diff(v, Int64(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint64(t *testing.T) {
	tt := []struct {
		value  any
		result uint64
	}{
		{value: nil, result: basetype.Uint64Invalid},
		{value: uint64(10), result: 10},
		{value: float64(math.Float64frombits(100)), result: 100},
		{value: Uint64(10), result: 10},
		{value: int8(10), result: basetype.Uint64Invalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint64[uint64](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
			t.Run("Uint64", func(t *testing.T) {
				v := typeconv.ToUint64[Uint64](tc.value)
				if diff := cmp.Diff(v, Uint64(tc.result)); diff != "" {
					t.Fatal(diff)
				}
			})
		})
	}
}

func TestToUint64z(t *testing.T) {
	tt := []struct {
		value  any
		result uint64
	}{
		{value: nil, result: basetype.Uint64zInvalid},
		{value: uint64(10), result: 10},
		{value: Uint64(10), result: 10},
		{value: int8(10), result: basetype.Uint64zInvalid},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			v := typeconv.ToUint64z[uint64](tc.value)
			if diff := cmp.Diff(v, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("Uint64", func(t *testing.T) {
			v := typeconv.ToUint64z[Uint64](tc.value)
			if diff := cmp.Diff(v, Uint64(tc.result)); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceEnum(t *testing.T) {
	value := []uint8{0, 1, 2}
	result := typeconv.ToSliceEnum[byte](value)
	if diff := cmp.Diff(value, result); diff != "" {
		t.Fatal(diff)
	}
}

func TestToSliceSint8(t *testing.T) {
	tt := []struct {
		value  any
		result []int8
	}{
		{value: nil, result: nil},
		{value: []int8{1, 2}, result: []int8{1, 2}},
		{value: []Int8{1, 2}, result: []int8{1, 2}},
		{value: int8(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceSint8[int8](tc.value)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}

		})

		t.Run("[]Int8", func(t *testing.T) {
			result := typeconv.ToSliceSint8[Int8](tc.value)
			var expected []Int8
			for _, v := range tc.result {
				expected = append(expected, Int8(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceUint8(t *testing.T) {
	tt := []struct {
		value  any
		result []uint8
	}{
		{value: nil, result: nil},
		{value: []uint8{1, 2}, result: []uint8{1, 2}},
		{value: []Uint8{1, 2}, result: []uint8{1, 2}},
		{value: []Uint8{}, result: nil},
		{value: uint8(0), result: nil},
		{value: []float32{0}, result: nil},
		{value: []bool{true, false}, result: []uint8{1, 0}},
		{value: []Bool{true, false}, result: []uint8{1, 0}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceUint8[uint8](tc.value)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Uint8", func(t *testing.T) {
			result := typeconv.ToSliceUint8[Uint8](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Uint8
			for _, v := range tc.result {
				expected = append(expected, Uint8(v))
			}
			if tc.result == nil {
				expected = nil
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceSint16(t *testing.T) {
	tt := []struct {
		value  any
		result []int16
	}{
		{value: nil, result: nil},
		{value: []int16{1, 2}, result: []int16{1, 2}},
		{value: []Int16{1, 2}, result: []int16{1, 2}},
		{value: []int16{}, result: nil},
		{value: []Int16{}, result: nil},
		{value: int16(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceSint16[int16](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Int16", func(t *testing.T) {
			result := typeconv.ToSliceSint16[Int16](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Int16
			for _, v := range tc.result {
				expected = append(expected, Int16(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceUint16(t *testing.T) {
	tt := []struct {
		value  any
		result []uint16
	}{
		{value: nil, result: nil},
		{value: []uint16{1, 2}, result: []uint16{1, 2}},
		{value: []Uint16{1, 2}, result: []uint16{1, 2}},
		{value: []uint16{}, result: nil},
		{value: []Uint16{}, result: nil},
		{value: uint16(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceUint16[uint16](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]uInt16", func(t *testing.T) {
			result := typeconv.ToSliceUint16[Uint16](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Uint16
			for _, v := range tc.result {
				expected = append(expected, Uint16(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceSint32(t *testing.T) {
	tt := []struct {
		value  any
		result []int32
	}{
		{value: nil, result: nil},
		{value: []int32{1, 2}, result: []int32{1, 2}},
		{value: []Int32{1, 2}, result: []int32{1, 2}},
		{value: []int32{}, result: nil},
		{value: []Int32{}, result: nil},
		{value: int32(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceSint32[int32](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Int32", func(t *testing.T) {
			result := typeconv.ToSliceSint32[Int32](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Int32
			for _, v := range tc.result {
				expected = append(expected, Int32(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceUint32(t *testing.T) {
	tt := []struct {
		value  any
		result []uint32
	}{
		{value: nil, result: nil},
		{value: []uint32{1, 2}, result: []uint32{1, 2}},
		{value: []Uint32{1, 2}, result: []uint32{1, 2}},
		{value: []uint32{}, result: nil},
		{value: []Uint32{}, result: nil},
		{value: uint32(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceUint32[uint32](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]uInt32", func(t *testing.T) {
			result := typeconv.ToSliceUint32[Uint32](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Uint32
			for _, v := range tc.result {
				expected = append(expected, Uint32(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceString(t *testing.T) {
	tt := []struct {
		value  any
		result []string
	}{
		{value: nil, result: nil},
		{value: []string{"a", "b"}, result: []string{"a", "b"}},
		{value: string(""), result: nil},
		{value: []String{"a", "b"}, result: []string{"a", "b"}},
		{value: []String{}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceString[string](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Float32", func(t *testing.T) {
			result := typeconv.ToSliceString[String](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []String
			for _, v := range tc.result {
				expected = append(expected, String(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceFloat32(t *testing.T) {
	tt := []struct {
		value  any
		result []float32
	}{
		{value: nil, result: nil},
		{value: []float32{1, 2}, result: []float32{1, 2}},
		{value: []Float32{1, 2}, result: []float32{1, 2}},
		{value: []Float32{}, result: nil},
		{value: []float32{}, result: nil},
		{value: float32(0), result: nil},
		{value: []int8{0}, result: nil},
		{value: []uint32{1093140480}, result: []float32{10.5}},
		{value: []Uint32{1093140480}, result: []float32{10.5}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceFloat32[float32](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Float32", func(t *testing.T) {
			result := typeconv.ToSliceFloat32[Float32](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Float32
			for _, v := range tc.result {
				expected = append(expected, Float32(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceFloat64(t *testing.T) {
	tt := []struct {
		value  any
		result []float64
	}{
		{value: nil, result: nil},
		{value: []float64{1, 2}, result: []float64{1, 2}},
		{value: []Float64{1, 2}, result: []float64{1, 2}},
		{value: []Float64{}, result: nil},
		{value: []float64{}, result: nil},
		{value: float64(0), result: nil},
		{value: []int8{0}, result: nil},
		{value: []uint64{4622100592565682176}, result: []float64{10.5}},
		{value: []Uint64{4622100592565682176}, result: []float64{10.5}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceFloat64[float64](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Float64", func(t *testing.T) {
			result := typeconv.ToSliceFloat64[Float64](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Float64
			for _, v := range tc.result {
				expected = append(expected, Float64(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceUint8z(t *testing.T) {
	value := []uint8{10}
	result := typeconv.ToSliceUint8z[uint8](value)
	if diff := cmp.Diff(result, value); diff != "" {
		t.Fatal(diff)
	}
}

func TestToSliceUint16z(t *testing.T) {
	value := []uint16{10}
	result := typeconv.ToSliceUint16z[uint16](value)
	if diff := cmp.Diff(result, value); diff != "" {
		t.Fatal(diff)
	}
}

func TestToSliceByte(t *testing.T) {
	value := []uint8{10}
	result := typeconv.ToSliceByte[uint8](value)
	if diff := cmp.Diff(result, value); diff != "" {
		t.Fatal(diff)
	}
}

func TestToSliceUint32z(t *testing.T) {
	value := []uint32{10}
	result := typeconv.ToSliceUint32z[uint32](value)
	if diff := cmp.Diff(result, value); diff != "" {
		t.Fatal(diff)
	}
}

func TestToSliceSint64(t *testing.T) {
	tt := []struct {
		value  any
		result []int64
	}{
		{value: nil, result: nil},
		{value: []int64{1, 2}, result: []int64{1, 2}},
		{value: []Int64{1, 2}, result: []int64{1, 2}},
		{value: []int64{}, result: nil},
		{value: []Int64{}, result: nil},
		{value: int64(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceSint64[int64](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]Int64", func(t *testing.T) {
			result := typeconv.ToSliceSint64[Int64](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Int64
			for _, v := range tc.result {
				expected = append(expected, Int64(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceUint64(t *testing.T) {
	tt := []struct {
		value  any
		result []uint64
	}{
		{value: nil, result: nil},
		{value: []uint64{1, 2}, result: []uint64{1, 2}},
		{value: []Uint64{1, 2}, result: []uint64{1, 2}},
		{value: []uint64{}, result: nil},
		{value: []Uint64{}, result: nil},
		{value: uint64(0), result: nil},
		{value: []float32{0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result := typeconv.ToSliceUint64[uint64](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})

		t.Run("[]uInt64", func(t *testing.T) {
			result := typeconv.ToSliceUint64[Uint64](tc.value)
			if len(result) == 0 && len(tc.result) == 0 {
				return
			}
			var expected []Uint64
			for _, v := range tc.result {
				expected = append(expected, Uint64(v))
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestToSliceUint64z(t *testing.T) {
	value := []uint64{10}
	result := typeconv.ToSliceUint64z[uint64](value)
	if diff := cmp.Diff(result, value); diff != "" {
		t.Fatal(diff)
	}
}
