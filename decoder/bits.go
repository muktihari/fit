// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"math"

	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/proto"
)

// bitvalue holds proto.Value in its integer form, enabling us to do bitwise operation over it.
// This is used for component expansion as Field's Value requiring expansion can hold up to
// 255 byte (2040 bits) data, this is obviously way more bits than Go's primitive value can handle.
//
// In Profile.xlsx (at least v21.141), the biggest value for component expansion is "raw_bbi"
// message for having 240 bits on "data" and the second is "hr" message for having 120 bits
// on "event_timestamp_12".
type bitvalue struct {
	// NOTE: We use array to avoid memory allocation, it's simple to maintain and it has more
	// deterministic performance. Max value to hold is 2040 bits, so the value of the last
	// index will always less than math.MaxUint64. We use the last index to determine the
	// validity of this stuct, if last index is math.MaxUint64, it's an invalid bitval.
	store [32]uint64
}

// makeBitValue creates bitvalue from proto.Value.
func makeBitValue(value proto.Value) (v bitvalue, ok bool) {
	switch value.Type() {
	case proto.TypeInt8:
		return bitvalue{store: [32]uint64{uint64(value.Int8())}}, true
	case proto.TypeUint8:
		return bitvalue{store: [32]uint64{uint64(value.Uint8())}}, true
	case proto.TypeInt16:
		return bitvalue{store: [32]uint64{uint64(value.Int16())}}, true
	case proto.TypeUint16:
		return bitvalue{store: [32]uint64{uint64(value.Uint16())}}, true
	case proto.TypeInt32:
		return bitvalue{store: [32]uint64{uint64(value.Int32())}}, true
	case proto.TypeUint32:
		return bitvalue{store: [32]uint64{uint64(value.Uint32())}}, true
	case proto.TypeInt64:
		return bitvalue{store: [32]uint64{uint64(value.Int64())}}, true
	case proto.TypeUint64:
		return bitvalue{store: [32]uint64{value.Uint64()}}, true
	case proto.TypeFloat32:
		return bitvalue{store: [32]uint64{uint64(value.Float32())}}, true
	case proto.TypeFloat64:
		return bitvalue{store: [32]uint64{uint64(value.Float64())}}, true
	case proto.TypeSliceInt8:
		return bitvalue{store: makeStoreFromSlice(value.SliceInt8(), 1)}, true
	case proto.TypeSliceUint8:
		return bitvalue{store: makeStoreFromSlice(value.SliceUint8(), 1)}, true
	case proto.TypeSliceInt16:
		return bitvalue{store: makeStoreFromSlice(value.SliceInt16(), 2)}, true
	case proto.TypeSliceUint16:
		return bitvalue{store: makeStoreFromSlice(value.SliceUint16(), 2)}, true
	case proto.TypeSliceInt32:
		return bitvalue{store: makeStoreFromSlice(value.SliceInt32(), 4)}, true
	case proto.TypeSliceUint32:
		return bitvalue{store: makeStoreFromSlice(value.SliceUint32(), 4)}, true
	case proto.TypeSliceInt64:
		return bitvalue{store: makeStoreFromSlice(value.SliceInt64(), 8)}, true
	case proto.TypeSliceUint64:
		return bitvalue{store: makeStoreFromSlice(value.SliceUint64(), 8)}, true
	case proto.TypeSliceFloat32:
		return bitvalue{store: makeStoreFromSlice(value.SliceFloat32(), 4)}, true
	case proto.TypeSliceFloat64:
		return bitvalue{store: makeStoreFromSlice(value.SliceFloat64(), 8)}, true
	}
	return bitvalue{store: [32]uint64{31: math.MaxUint64}}, false
}

type numeric interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// makeStoreFromSlice creates value store from given s (slice of supported numeric type).
func makeStoreFromSlice[S []E, E numeric](s S, bitsize uint8) (store [32]uint64) {
	var index uint8
	value, pos := uint64(0), uint8(0)
	for {
		if len(s) == 0 {
			store[index] = value
			break
		}
		if pos == 8 {
			store[index] = value
			value, pos = 0, 0
			index++
		}
		value |= uint64(s[0]) << (pos * 8)
		pos += bitsize
		s = s[1:]
	}
	return store
}

// makeBitValueFromUint32 make bitval from uint32.
func makeBitValueFromUint32(v uint32) bitvalue {
	return bitvalue{store: [32]uint64{uint64(v)}}
}

// AsUint32 casts bitvalue into uint32, this will not update the bitval's value store.
// This is used mainly for accumulating value, all accumulated value type is <= uint32.
func (v *bitvalue) AsUint32() uint32 { return uint32(v.store[0] & math.MaxUint32) }

// ToValue converts bitval into proto.Value. This will not evaluate the whole
// bitvalue's value store as it's not necessary. bitval is only used for component
// expansion that is only using 32 bits per component. Additionally it's also used
// for converting []byte value into its representative numeric value. If bitvalue
// is invalid or baseType target is invalid, proto.Value{} will be returned.
func (v *bitvalue) ToValue(baseType basetype.BaseType) proto.Value {
	if v.store[31] == math.MaxUint64 {
		return proto.Value{}
	}
	switch baseType {
	case basetype.Sint8:
		return proto.Int8(int8(v.store[0] & math.MaxUint8))
	case basetype.Enum, basetype.Uint8, basetype.Uint8z:
		return proto.Uint8(uint8(v.store[0] & math.MaxUint8))
	case basetype.Sint16:
		return proto.Int16(int16(v.store[0] & math.MaxUint16))
	case basetype.Uint16, basetype.Uint16z:
		return proto.Uint16(uint16(v.store[0] & math.MaxUint16))
	case basetype.Sint32:
		return proto.Int32(int32(v.store[0] & math.MaxUint32))
	case basetype.Uint32, basetype.Uint32z:
		return proto.Uint32(uint32(v.store[0] & math.MaxUint32))
	case basetype.Sint64:
		return proto.Int64(int64(v.store[0]))
	case basetype.Uint64, basetype.Uint64z:
		return proto.Uint64(v.store[0])
	case basetype.Float32:
		return proto.Float32(float32(v.store[0] & math.MaxUint32))
	case basetype.Float64:
		return proto.Float64(float64(v.store[0]))
	}
	return proto.Value{}
}

// PullByMask retrieves a value of the specified bit size from the storage value and
// the storage value will be updated accordingly. If one of these conditions is met,
// zero and false will be returned:
//   - bitvalue is invalid
//   - value store run out bits (reach zero)
//   - given bits > 32
func (v *bitvalue) PullByMask(bits byte) (val uint32, ok bool) {
	if v.store[0] == 0 || bits > 32 {
		return 0, false
	}

	mask := uint64(1)<<bits - 1     // e.g. (1 << 8) - 1     = 255
	val = uint32(v.store[0] & mask) // e.g. 0x27010E08 & 255 = 0x08
	v.store[0] >>= bits             // e.g. 0x27010E08 >> 8  = 0x27010E

	for i := 1; i < len(v.store); i++ {
		if v.store[i] == 0 {
			break
		}
		v.store[i-1] = v.store[i-1]<<bits | v.store[i]&mask
		v.store[i] = v.store[i] >> bits
	}

	return val, true
}

// valueAppend appends elem into slice. Elem must be the proper element of
// slice's element otherwise, unexpected behavior.
func valueAppend(slice proto.Value, elem proto.Value) proto.Value {
	switch elem.Type() {
	case proto.TypeInt8:
		return proto.SliceInt8(append(slice.SliceInt8(), elem.Int8()))
	case proto.TypeUint8:
		return proto.SliceUint8(append(slice.SliceUint8(), elem.Uint8()))
	case proto.TypeInt16:
		return proto.SliceInt16(append(slice.SliceInt16(), elem.Int16()))
	case proto.TypeUint16:
		return proto.SliceUint16(append(slice.SliceUint16(), elem.Uint16()))
	case proto.TypeInt32:
		return proto.SliceInt32(append(slice.SliceInt32(), elem.Int32()))
	case proto.TypeUint32:
		return proto.SliceUint32(append(slice.SliceUint32(), elem.Uint32()))
	case proto.TypeInt64:
		return proto.SliceInt64(append(slice.SliceInt64(), elem.Int64()))
	case proto.TypeUint64:
		return proto.SliceUint64(append(slice.SliceUint64(), elem.Uint64()))
	case proto.TypeFloat32:
		return proto.SliceFloat32(append(slice.SliceFloat32(), elem.Float32()))
	case proto.TypeFloat64:
		return proto.SliceFloat64(append(slice.SliceFloat64(), elem.Float64()))
	}
	return slice
}
