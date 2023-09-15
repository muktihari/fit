// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scaleoffset

import (
	"github.com/muktihari/fit/profile/basetype"
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
	vals := make([]float64, 0, len(values))
	for i := 0; i < len(values); i++ {
		vals = append(vals, Apply[E](values[i], scale, offset))
	}
	return vals
}

// ApplyAny applies scale and offset when possible, otherwise return original value.
// This function can only accept primitive-types value such as int8, []int8, uint32, []uint32, etc.
func ApplyAny(value any, scale, offset float64) any {
	if scale == 1 && offset == 0 {
		return value
	}

	switch val := value.(type) {
	case int8:
		return Apply(val, scale, offset)
	case []int8:
		return ApplySlice(val, scale, offset)
	case uint8:
		return Apply(val, scale, offset)
	case []uint8:
		return ApplySlice(val, scale, offset)
	case int16:
		return Apply(val, scale, offset)
	case []int16:
		return ApplySlice(val, scale, offset)
	case uint16:
		return Apply(val, scale, offset)
	case []uint16:
		return ApplySlice(val, scale, offset)
	case int32:
		return Apply(val, scale, offset)
	case []int32:
		return ApplySlice(val, scale, offset)
	case uint32:
		return Apply(val, scale, offset)
	case []uint32:
		return ApplySlice(val, scale, offset)
	case int64:
		return Apply(val, scale, offset)
	case []int64:
		return ApplySlice(val, scale, offset)
	case uint64:
		return Apply(val, scale, offset)
	case []uint64:
		return ApplySlice(val, scale, offset)
	case float32:
		return Apply(val, scale, offset)
	case []float32:
		return ApplySlice(val, scale, offset)
	case float64:
		return Apply(val, scale, offset)
	case []float64:
		return ApplySlice(val, scale, offset)
	default:
		return val // not supported, return original value
	}
}

// Discard discards applied scale and offset on value.
func Discard(value, scale, offset float64) float64 {
	if scale == 1 && offset == 0 {
		return value
	}
	return (value + offset) * scale
}

// DiscardSlice discards applied scale and offset on slice values.
func DiscardSlice[E Numeric](values []float64, scale, offset float64) []E {
	vals := make([]E, 0, len(values))
	for i := 0; i < len(values); i++ {
		if scale == 1 && offset == 0 {
			vals = append(vals, E(values[i]))
			continue
		}
		vals = append(vals, E((values[i]+offset)*scale))
	}
	return vals
}

// DiscardAny restores scaled value in the form of float64 or []float64 to its basetype's form.
func DiscardAny(value any, baseType basetype.BaseType, scale, offset float64) any {
	switch val := value.(type) {
	case float64: // a scaled value will always in float64 form.
		dv := Discard(val, scale, offset)

		switch baseType {
		case basetype.Sint8:
			return int8(dv)
		case basetype.Byte, basetype.Uint8, basetype.Uint8z:
			return uint8(dv)
		case basetype.Sint16:
			return int16(dv)
		case basetype.Uint16, basetype.Uint16z:
			return uint16(dv)
		case basetype.Sint32:
			return int32(dv)
		case basetype.Uint32, basetype.Uint32z:
			return uint32(dv)
		case basetype.Float32:
			return float32(dv)
		case basetype.Float64:
			return float64(dv)
		case basetype.Sint64:
			return int64(dv)
		case basetype.Uint64, basetype.Uint64z:
			return uint64(dv)
		}
	case []float64: // array of scaled values will always in []float64 form.
		switch baseType {
		case basetype.Byte, basetype.Uint8, basetype.Uint8z:
			return DiscardSlice[uint8](val, scale, offset)
		case basetype.Sint8:
			return DiscardSlice[int8](val, scale, offset)
		case basetype.Sint16:
			return DiscardSlice[int16](val, scale, offset)
		case basetype.Uint16, basetype.Uint16z:
			return DiscardSlice[uint16](val, scale, offset)
		case basetype.Sint32:
			return DiscardSlice[int32](val, scale, offset)
		case basetype.Uint32, basetype.Uint32z:
			return DiscardSlice[uint32](val, scale, offset)
		case basetype.Float32:
			return DiscardSlice[float32](val, scale, offset)
		case basetype.Float64:
			return DiscardSlice[float64](val, scale, offset)
		case basetype.Sint64:
			return DiscardSlice[int64](val, scale, offset)
		case basetype.Uint64, basetype.Uint64z:
			return DiscardSlice[uint64](val, scale, offset)
		}
	}

	return value // not supported, return original value
}
