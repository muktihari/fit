// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scaleoffset

import (
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

// Apply applies scale and offset on value.
func Apply[T Numeric](value T, scale, offset float64) float64 {
	return float64(value)/scale - offset
}

// Apply applies scale and offset on slice values.
func ApplySlice[S []E, E Numeric](values S, scale, offset float64) []float64 {
	vals := make([]float64, len(values))
	for i := range values {
		vals[i] = Apply(values[i], scale, offset)
	}
	return vals
}

// ApplyValue applies scale and offset when possible, otherwise return original value.
// This function can only accept primitive-types such as int8, []int8, uint32, []uint32, etc in the form of proto.Value.
func ApplyValue(value proto.Value, scale, offset float64) proto.Value {
	if scale == 1 && offset == 0 {
		return value
	}
	switch value.Type() {
	case proto.TypeInt8:
		return proto.Float64(Apply(value.Int8(), scale, offset))
	case proto.TypeUint8:
		return proto.Float64(Apply(value.Uint8(), scale, offset))
	case proto.TypeInt16:
		return proto.Float64(Apply(value.Int16(), scale, offset))
	case proto.TypeUint16:
		return proto.Float64(Apply(value.Uint16(), scale, offset))
	case proto.TypeInt32:
		return proto.Float64(Apply(value.Int32(), scale, offset))
	case proto.TypeUint32:
		return proto.Float64(Apply(value.Uint32(), scale, offset))
	case proto.TypeInt64:
		return proto.Float64(Apply(value.Int64(), scale, offset))
	case proto.TypeUint64:
		return proto.Float64(Apply(value.Uint64(), scale, offset))
	case proto.TypeFloat32:
		return proto.Float64(Apply(value.Float32(), scale, offset))
	case proto.TypeFloat64:
		return proto.Float64(Apply(value.Float64(), scale, offset))
	case proto.TypeSliceInt8:
		return proto.SliceFloat64(ApplySlice(value.SliceInt8(), scale, offset))
	case proto.TypeSliceUint8:
		return proto.SliceFloat64(ApplySlice(value.SliceUint8(), scale, offset))
	case proto.TypeSliceInt16:
		return proto.SliceFloat64(ApplySlice(value.SliceInt16(), scale, offset))
	case proto.TypeSliceUint16:
		return proto.SliceFloat64(ApplySlice(value.SliceUint16(), scale, offset))
	case proto.TypeSliceInt32:
		return proto.SliceFloat64(ApplySlice(value.SliceInt32(), scale, offset))
	case proto.TypeSliceUint32:
		return proto.SliceFloat64(ApplySlice(value.SliceUint32(), scale, offset))
	case proto.TypeSliceInt64:
		return proto.SliceFloat64(ApplySlice(value.SliceInt64(), scale, offset))
	case proto.TypeSliceUint64:
		return proto.SliceFloat64(ApplySlice(value.SliceUint64(), scale, offset))
	case proto.TypeSliceFloat32:
		return proto.SliceFloat64(ApplySlice(value.SliceFloat32(), scale, offset))
	case proto.TypeSliceFloat64:
		return proto.SliceFloat64(ApplySlice(value.SliceFloat64(), scale, offset))
	}
	return value // not supported, return original value
}

// Discard discards applied scale and offset on value.
func Discard(value, scale, offset float64) float64 {
	if scale == 1 && offset == 0 {
		return value
	}
	return (value + offset) * scale
}

// DiscardSlice discards applied scale and offset on slice values.
func DiscardSlice[T Numeric](values []float64, scale, offset float64) []T {
	vals := make([]T, len(values))
	for i := range values {
		if scale == 1 && offset == 0 {
			vals[i] = T(values[i])
			continue
		}
		vals[i] = T((values[i] + offset) * scale)
	}
	return vals
}

func DiscardValue(value proto.Value, baseType basetype.BaseType, scale, offset float64) proto.Value {
	switch value.Type() {
	case proto.TypeFloat64:
		dv := Discard(value.Float64(), scale, offset)
		switch baseType {
		case basetype.Sint8:
			return proto.Int8(int8(dv))
		case basetype.Byte, basetype.Uint8, basetype.Uint8z:
			return proto.Uint8(uint8(dv))
		case basetype.Sint16:
			return proto.Int16(int16(dv))
		case basetype.Uint16, basetype.Uint16z:
			return proto.Uint16(uint16(dv))
		case basetype.Sint32:
			return proto.Int32(int32(dv))
		case basetype.Uint32, basetype.Uint32z:
			return proto.Uint32(uint32(dv))
		case basetype.Float32:
			return proto.Float32(float32(dv))
		case basetype.Float64:
			return proto.Float64(float64(dv))
		case basetype.Sint64:
			return proto.Int64(int64(dv))
		case basetype.Uint64, basetype.Uint64z:
			return proto.Uint64(uint64(dv))
		}
	case proto.TypeSliceFloat64:
		switch baseType {
		case basetype.Byte, basetype.Uint8, basetype.Uint8z:
			return proto.SliceUint8(DiscardSlice[uint8](value.SliceFloat64(), scale, offset))
		case basetype.Sint8:
			return proto.SliceInt8(DiscardSlice[int8](value.SliceFloat64(), scale, offset))
		case basetype.Sint16:
			return proto.SliceInt16(DiscardSlice[int16](value.SliceFloat64(), scale, offset))
		case basetype.Uint16, basetype.Uint16z:
			return proto.SliceUint16(DiscardSlice[uint16](value.SliceFloat64(), scale, offset))
		case basetype.Sint32:
			return proto.SliceInt32(DiscardSlice[int32](value.SliceFloat64(), scale, offset))
		case basetype.Uint32, basetype.Uint32z:
			return proto.SliceUint32(DiscardSlice[uint32](value.SliceFloat64(), scale, offset))
		case basetype.Float32:
			return proto.SliceFloat32(DiscardSlice[float32](value.SliceFloat64(), scale, offset))
		case basetype.Float64:
			return proto.SliceFloat64(DiscardSlice[float64](value.SliceFloat64(), scale, offset))
		case basetype.Sint64:
			return proto.SliceInt64(DiscardSlice[int64](value.SliceFloat64(), scale, offset))
		case basetype.Uint64, basetype.Uint64z:
			return proto.SliceUint64(DiscardSlice[uint64](value.SliceFloat64(), scale, offset))
		}
	}
	return value
}
