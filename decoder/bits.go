// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/proto"
)

const (
	bit    = 8
	maxBit = 32
	size   = maxBit / bit
)

// bitsFromValue convert value into 32-bits unsigned integer.
//
// Profile.xlsx (on Bits header's comment) says: Current implementation only supports Bits value of max 32.
func bitsFromValue(value proto.Value) (bits uint32, ok bool) {
	switch value.Type() {
	case proto.TypeInt8:
		return uint32(value.Int8()), true
	case proto.TypeUint8:
		return uint32(value.Uint8()), true
	case proto.TypeInt16:
		return uint32(value.Int16()), true
	case proto.TypeUint16:
		return uint32(value.Uint16()), true
	case proto.TypeInt32:
		return uint32(value.Int32()), true
	case proto.TypeUint32:
		return value.Uint32(), true
	case proto.TypeInt64:
		return uint32(value.Int64()), true
	case proto.TypeUint64:
		return uint32(value.Uint64()), true
	case proto.TypeFloat32:
		return uint32(value.Float32()), true
	case proto.TypeFloat64:
		return uint32(value.Float64()), true
	case proto.TypeSliceUint8:
		val := value.SliceUint8()
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
func valueFromBits(bits uint32, baseType basetype.BaseType) proto.Value {
	switch baseType {
	case basetype.Sint8:
		return proto.Int8(int8(bits))
	case basetype.Enum, basetype.Uint8, basetype.Uint8z:
		return proto.Uint8(uint8(bits))
	case basetype.Sint16:
		return proto.Int16(int16(bits))
	case basetype.Uint16, basetype.Uint16z:
		return proto.Uint16(uint16(bits))
	case basetype.Sint32:
		return proto.Int32(int32(bits))
	case basetype.Uint32, basetype.Uint32z:
		return proto.Uint32(uint32(bits))
	case basetype.Float32:
		return proto.Float32(float32(bits))
	case basetype.Float64:
		return proto.Float64(float64(bits))
	case basetype.Sint64:
		return proto.Int64(int64(bits))
	case basetype.Uint64, basetype.Uint64z:
		return proto.Uint64(uint64(bits))
	}
	return proto.Value{}
}
