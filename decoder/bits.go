// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/profile/basetype"
)

const (
	bit    = 8
	maxBit = 32
	size   = maxBit / bit
)

// bitsFromValue convert value into 32-bits unsigned integer.
//
// Profile.xlsx (on Bits header's comment) says: Current implementation only supports Bits value of max 32.
func bitsFromValue(value any) (bits uint32, ok bool) {
	switch val := value.(type) {
	case int8:
		return uint32(val), true
	case uint8:
		return uint32(val), true
	case int16:
		return uint32(val), true
	case uint16:
		return uint32(val), true
	case int32:
		return uint32(val), true
	case uint32:
		return uint32(val), true
	case int64:
		return uint32(val), true
	case uint64:
		return uint32(val), true
	case float32:
		return uint32(val), true
	case float64:
		return uint32(val), true
	case []byte:
		if len(val) > size {
			return 0, false
		}
		for i := range val {
			if val[i] == basetype.ByteInvalid { // all values must be valid
				return 0, false
			}
			bits |= uint32(val[i]) << (i * bit) // little-endian
		}
		return bits, true
	}
	return 0, false
}

// valueFromBits cast back bits into it's original value.
func valueFromBits(bits uint32, baseType basetype.BaseType) any {
	switch baseType {
	case basetype.Sint8:
		return int8(bits)
	case basetype.Uint8, basetype.Uint8z:
		return uint8(bits)
	case basetype.Sint16:
		return int16(bits)
	case basetype.Uint16, basetype.Uint16z:
		return uint16(bits)
	case basetype.Sint32:
		return int32(bits)
	case basetype.Uint32, basetype.Uint32z:
		return uint32(bits)
	case basetype.Float32:
		return float32(bits)
	case basetype.Float64:
		return float64(bits)
	case basetype.Sint64:
		return int64(bits)
	case basetype.Uint64, basetype.Uint64z:
		return uint64(bits)
	}
	return baseType.Invalid()
}
