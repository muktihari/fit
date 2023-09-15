// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basetype

import (
	"math"
	"reflect"
)

// BaseTypeNumMask used to get the index/order of the constants (start from 0, Enum).
// Example: (Sint16 & BaseTypeNumMask) -> 3.
const BaseTypeNumMask = 0x1F

// BaseType is the base of all types used in Fit.
type BaseType byte

const (
	Enum    BaseType = 0x00
	Sint8   BaseType = 0x01 // 2’s complement format
	Uint8   BaseType = 0x02
	Sint16  BaseType = 0x83 // 2’s complement format
	Uint16  BaseType = 0x84
	Sint32  BaseType = 0x85 // 2’s complement format
	Uint32  BaseType = 0x86
	String  BaseType = 0x07 // Null terminated string encoded in UTF-8 format: 0x00
	Float32 BaseType = 0x88
	Float64 BaseType = 0x89
	Uint8z  BaseType = 0x0A
	Uint16z BaseType = 0x8B
	Uint32z BaseType = 0x8C
	Byte    BaseType = 0x0D // Array of bytes. Field is invalid if all bytes are invalid.
	Sint64  BaseType = 0x8E // 2’s complement format
	Uint64  BaseType = 0x8F
	Uint64z BaseType = 0x90
)

const (
	EnumInvalid    byte   = math.MaxUint8  // 0xFF
	Sint8Invalid   int8   = math.MaxInt8   // 0x7F
	Uint8Invalid   uint8  = math.MaxUint8  // 0xFF
	Sint16Invalid  int16  = math.MaxInt16  // 0x7FFF
	Uint16Invalid  uint16 = math.MaxUint16 // 0xFFFF
	Sint32Invalid  int32  = math.MaxInt32  // 0x7FFFFFFF
	Uint32Invalid  uint32 = math.MaxUint32 // 0xFFFFFFFF
	StringInvalid  string = "\x00"         // 0x00. Same as string([]byte{0x00}), 0x00 is utf-8 null-terminated string.
	Float32Invalid uint32 = math.MaxUint32 // 0xFFFFFFFF. math.Float32frombits(0xFFFFFFFF) produces float64 NaN which is uncomparable. Can only check using math.IsNaN or convert it in Integer form like this.
	Float64Invalid uint64 = math.MaxUint64 // 0xFFFFFFFFFFFFFFFF. math.Float64frombits(0xFFFFFFFFFFFFFFFF) produces float32 NaN which is uncomparable. Can only check using math.IsNaN or convert it in Integer form like this.
	Uint8zInvalid  uint8  = 0              // 0x00
	Uint16zInvalid uint16 = 0              // 0x0000
	Uint32zInvalid uint32 = 0              // 0x00000000
	ByteInvalid    byte   = math.MaxUint8  // 0xFF
	Sint64Invalid  int64  = math.MaxInt64  // 0x7FFFFFFFFFFFFFFF
	Uint64Invalid  uint64 = math.MaxUint64 // 0xFFFFFFFFFFFFFFFF
	Uint64zInvalid uint64 = 0              // 0x0000000000000000
)

type typestruct struct {
	base          BaseType
	stringer      string
	size          byte
	endianAbility byte
	invalid       any
	goType        string
	isInteger     bool
	kind          reflect.Kind
}

var basetypes = [...]typestruct{
	{
		base:          Enum,
		stringer:      "enum",
		size:          1,
		endianAbility: 0,
		invalid:       EnumInvalid,
		goType:        "byte", /* treat it as byte */
		isInteger:     true,
		kind:          reflect.Uint8,
	},
	{
		base:          Sint8,
		stringer:      "sint8",
		size:          1,
		endianAbility: 0,
		invalid:       Sint8Invalid,
		goType:        "int8",
		isInteger:     true,
		kind:          reflect.Int8,
	},
	{
		base:          Uint8,
		stringer:      "uint8",
		size:          1,
		endianAbility: 0,
		invalid:       Uint8Invalid,
		goType:        "uint8",
		isInteger:     true,
		kind:          reflect.Uint8,
	},
	{
		base:          Sint16,
		stringer:      "sint16",
		size:          2,
		endianAbility: 1,
		invalid:       Sint16Invalid,
		goType:        "int16",
		isInteger:     true,
		kind:          reflect.Int16,
	},
	{
		base:          Uint16,
		stringer:      "uint16",
		size:          2,
		endianAbility: 1,
		invalid:       Uint16Invalid,
		goType:        "uint16",
		isInteger:     true,
		kind:          reflect.Uint16,
	},
	{
		base:          Sint32,
		stringer:      "sint32",
		size:          4,
		endianAbility: 1,
		invalid:       Sint32Invalid,
		goType:        "int32",
		isInteger:     true,
		kind:          reflect.Int32,
	},
	{
		base:          Uint32,
		stringer:      "uint32",
		size:          4,
		endianAbility: 1,
		invalid:       Uint32Invalid,
		goType:        "uint32",
		isInteger:     true,
		kind:          reflect.Uint32,
	},
	{
		base:          String,
		stringer:      "string",
		size:          1,
		endianAbility: 0,
		invalid:       StringInvalid,
		goType:        "string",
		isInteger:     false,
		kind:          reflect.String,
	},
	{
		base:          Float32,
		stringer:      "float32",
		size:          4,
		endianAbility: 1,
		invalid:       math.Float32frombits(Float32Invalid),
		goType:        "float32",
		isInteger:     false,
		kind:          reflect.Float32,
	},
	{
		base:          Float64,
		stringer:      "float64",
		size:          8,
		endianAbility: 1,
		invalid:       math.Float64frombits(Float64Invalid),
		goType:        "float64",
		isInteger:     false,
		kind:          reflect.Float64,
	},
	{
		base:          Uint8z,
		stringer:      "uint8z",
		size:          1,
		endianAbility: 0,
		invalid:       Uint8zInvalid,
		goType:        "uint8",
		isInteger:     true,
		kind:          reflect.Uint8,
	},
	{
		base:          Uint16z,
		stringer:      "uint16z",
		size:          2,
		endianAbility: 1,
		invalid:       Uint16zInvalid,
		goType:        "uint16",
		isInteger:     true,
		kind:          reflect.Uint16,
	},
	{
		base:          Uint32z,
		stringer:      "uint32z",
		size:          4,
		endianAbility: 1,
		invalid:       Uint32zInvalid,
		goType:        "uint32",
		isInteger:     true,
		kind:          reflect.Uint32,
	},
	{
		base:          Byte,
		stringer:      "byte",
		size:          1,
		endianAbility: 0,
		invalid:       ByteInvalid,
		goType:        "byte",
		isInteger:     true,
		kind:          reflect.Uint8,
	},
	{
		base:          Sint64,
		stringer:      "sint64",
		size:          8,
		endianAbility: 1,
		invalid:       Sint64Invalid,
		goType:        "int64",
		isInteger:     true,
		kind:          reflect.Int64,
	},
	{
		base:          Uint64,
		stringer:      "uint64",
		size:          8,
		endianAbility: 1,
		invalid:       Uint64Invalid,
		goType:        "uint64",
		isInteger:     true,
		kind:          reflect.Uint64,
	},
	{
		base:          Uint64z,
		stringer:      "uint64z",
		size:          8,
		endianAbility: 1,
		invalid:       Uint64zInvalid,
		goType:        "uint64",
		isInteger:     true,
		kind:          reflect.Uint64,
	},
}

// FromString convert given s into BaseType, if not valid 255 will be returned.
func FromString(s string) BaseType {
	v, ok := stringmap[s]
	if !ok {
		return 255 // invalid, highest value of byte.
	}
	return v
}

// List returns all constants.
func List() []BaseType {
	// no need to be reserved in memory, retrieve from existing variable only when needed.
	vs := make([]BaseType, 0, len(basetypes))
	for _, v := range basetypes {
		vs = append(vs, BaseType(v.base))
	}
	return vs
}

// String returns string representation of t.
func (t BaseType) String() string {
	if !valid(t) {
		return "invalid"
	}
	return basetypes[t&BaseTypeNumMask].stringer
}

var stringmap = func() map[string]BaseType {
	m := make(map[string]BaseType, len(basetypes))
	for key, val := range basetypes {
		m[val.stringer] = BaseType(key)
	}
	return m
}()

// Size returns how many bytes the value need in binary form.
func (t BaseType) Size() byte {
	if !valid(t) {
		return 0
	}
	return basetypes[t&BaseTypeNumMask].size
}

// GoType returns go equivalent type in string.
func (t BaseType) GoType() string {
	if !valid(t) {
		return "invalid"
	}
	return basetypes[t&BaseTypeNumMask].goType
}

func (t BaseType) Kind() reflect.Kind {
	if !valid(t) {
		return reflect.Invalid
	}
	return basetypes[t&BaseTypeNumMask].kind
}

// EndianAbility return whether t have endianess.
func (t BaseType) EndianAbility() byte {
	if !valid(t) {
		return 0
	}
	return basetypes[t&BaseTypeNumMask].endianAbility
}

func (t BaseType) IsInteger() bool {
	if !valid(t) {
		return false
	}
	return basetypes[t&BaseTypeNumMask].isInteger
}

// Invalid returns invalid value of t. e.g. Byte is 255 (its highest value).
func (t BaseType) Invalid() any {
	if !valid(t) {
		return "invalid"
	}
	return basetypes[t&BaseTypeNumMask].invalid
}

func valid(t BaseType) bool { return t&BaseTypeNumMask < BaseType(len(basetypes)) }
