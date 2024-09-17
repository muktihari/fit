package combiner

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

func TestSum(t *testing.T) {
	tt := []struct {
		v1 proto.Value
		v2 proto.Value
		r  proto.Value
	}{
		{v1: proto.Int8(1), v2: proto.Int8(2), r: proto.Int8(3)},
		{v1: proto.Uint8(1), v2: proto.Uint8(2), r: proto.Uint8(3)},
		{v1: proto.Int16(1), v2: proto.Int16(2), r: proto.Int16(3)},
		{v1: proto.Uint16(1), v2: proto.Uint16(2), r: proto.Uint16(3)},
		{v1: proto.Int32(1), v2: proto.Int32(2), r: proto.Int32(3)},
		{v1: proto.Uint32(1), v2: proto.Uint32(2), r: proto.Uint32(3)},
		{v1: proto.Int64(1), v2: proto.Int64(2), r: proto.Int64(3)},
		{v1: proto.Uint64(1), v2: proto.Uint64(2), r: proto.Uint64(3)},
		{v1: proto.Float32(1), v2: proto.Float32(2), r: proto.Float32(3)},
		{v1: proto.Float64(1), v2: proto.Float64(2), r: proto.Float64(3)},
		{v1: proto.SliceInt8([]int8{1}), v2: proto.SliceInt8([]int8{1}), r: proto.SliceInt8([]int8{2})},
		{v1: proto.SliceUint8([]uint8{1}), v2: proto.SliceUint8([]uint8{1}), r: proto.SliceUint8([]uint8{2})},
		{v1: proto.SliceInt16([]int16{1}), v2: proto.SliceInt16([]int16{1}), r: proto.SliceInt16([]int16{2})},
		{v1: proto.SliceUint16([]uint16{1}), v2: proto.SliceUint16([]uint16{1}), r: proto.SliceUint16([]uint16{2})},
		{v1: proto.SliceInt32([]int32{1}), v2: proto.SliceInt32([]int32{1}), r: proto.SliceInt32([]int32{2})},
		{v1: proto.SliceUint32([]uint32{1}), v2: proto.SliceUint32([]uint32{1}), r: proto.SliceUint32([]uint32{2})},
		{v1: proto.SliceInt64([]int64{1}), v2: proto.SliceInt64([]int64{1}), r: proto.SliceInt64([]int64{2})},
		{v1: proto.SliceUint64([]uint64{1}), v2: proto.SliceUint64([]uint64{1}), r: proto.SliceUint64([]uint64{2})},
		{v1: proto.SliceFloat32([]float32{1}), v2: proto.SliceFloat32([]float32{1}), r: proto.SliceFloat32([]float32{2})},
		{v1: proto.SliceFloat64([]float64{1}), v2: proto.SliceFloat64([]float64{1}), r: proto.SliceFloat64([]float64{2})},
		{v1: proto.Bool(typedef.BoolTrue), v2: proto.Bool(typedef.BoolFalse), r: proto.Bool(typedef.BoolTrue)},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.v1.Type()), func(t *testing.T) {
			r := sum(tc.v1, tc.v2)
			if diff := cmp.Diff(r.Any(), tc.r.Any()); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	tt := []struct {
		name string
		v1   proto.Value
		v2   proto.Value
		r    []int8
	}{
		{
			name: "same len",
			v1:   proto.SliceInt8([]int8{10, 10}),
			v2:   proto.SliceInt8([]int8{10, 10}),
			r:    []int8{20, 20},
		},
		{
			name: "len v1 > v2",
			v1:   proto.SliceInt8([]int8{10, 10, 10}),
			v2:   proto.SliceInt8([]int8{10, 10}),
			r:    []int8{20, 20, 10},
		},
		{
			name: "len v1 < v2",
			v1:   proto.SliceInt8([]int8{10, 10}),
			v2:   proto.SliceInt8([]int8{10, 10, 5}),
			r:    []int8{20, 20, 5},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			r := sumslice(tc.v1.SliceInt8(), tc.v2.SliceInt8())
			if diff := cmp.Diff(r, tc.r); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
