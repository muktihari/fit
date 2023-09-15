// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv

import (
	"github.com/muktihari/fit/profile/basetype"
)

// NumericToInt64 convert any numeric (~int | ~float) value of val to int64, if val is non-numeric value return false.
func NumericToInt64(val any) (int64, bool) {
	switch v := val.(type) {
	case int8:
		return int64(v), true
	case uint8:
		return int64(v), true
	case int16:
		return int64(v), true
	case uint16:
		return int64(v), true
	case int32:
		return int64(v), true
	case uint32:
		return int64(v), true
	case int64:
		return v, true
	case uint64:
		return int64(v), true
	case float32:
		return int64(v), true
	case float64:
		return int64(v), true
	}
	return 0, false
}

// IntegerToInt64 converts any integer value of val to int64, if val is non-integer value return false.
func IntegerToInt64(val any) (int64, bool) {
	switch v := val.(type) {
	case int8:
		return int64(v), true
	case uint8:
		return int64(v), true
	case int16:
		return int64(v), true
	case uint16:
		return int64(v), true
	case int32:
		return int64(v), true
	case uint32:
		return int64(v), true
	case int64:
		return v, true
	case uint64:
		return int64(v), true
	}
	return 0, false
}

// FloatToInt64 converts any floating-point value of val to int64, if val is non-floating-point value return false.
func FloatToInt64(val any) (int64, bool) {
	switch v := val.(type) {
	case float32:
		return int64(v), true
	case float64:
		return int64(v), true
	}
	return 0, false
}

func Int64ToNumber(val int64, baseType basetype.BaseType) any {
	switch baseType {
	case basetype.Sint8:
		return int8(val)
	case basetype.Uint8, basetype.Uint8z:
		return uint8(val)
	case basetype.Sint16:
		return int16(val)
	case basetype.Uint16, basetype.Uint16z:
		return uint16(val)
	case basetype.Sint32:
		return int32(val)
	case basetype.Uint32, basetype.Uint32z:
		return uint32(val)
	case basetype.Float32:
		return float32(val)
	case basetype.Float64:
		return float64(val)
	case basetype.Sint64:
		return int64(val)
	case basetype.Uint64, basetype.Uint64z:
		return uint64(val)
	}

	return baseType.Invalid()
}
