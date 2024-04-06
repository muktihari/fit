// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"math"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/muktihari/fit/profile/basetype"
)

// stringptr is an identifier distinguish a string pointer to a []byte pointer which both are *byte.
type stringptr *byte

// Type is value's type
type Type byte

const (
	TypeInvalid Type = iota
	TypeBool
	TypeInt8
	TypeUint8
	TypeInt16
	TypeUint16
	TypeInt32
	TypeUint32
	TypeInt64
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeString
	TypeSliceBool
	TypeSliceInt8
	TypeSliceUint8
	TypeSliceInt16
	TypeSliceUint16
	TypeSliceInt32
	TypeSliceUint32
	TypeSliceInt64
	TypeSliceUint64
	TypeSliceFloat32
	TypeSliceFloat64
	TypeSliceString
)

var typeStrings = [...]string{
	"Invalid",
	"Bool",
	"Int8",
	"Uint8",
	"Int16",
	"Uint16",
	"Int32",
	"Uint32",
	"Int64",
	"Uint64",
	"Float32",
	"Float64",
	"String",
	"SliceBool",
	"SliceInt8",
	"SliceUint8",
	"SliceInt16",
	"SliceUint16",
	"SliceInt32",
	"SliceUint32",
	"SliceInt64",
	"SliceUint64",
	"SliceFloat32",
	"SliceFloat64",
	"SliceString",
}

func (t Type) String() string {
	if t < Type(len(typeStrings)) {
		return typeStrings[t]
	}
	return "proto.TypeInvalid(" + strconv.Itoa(int(t)) + ")"
}

// Value is a zero alloc implementation value that hold any FIT protocol value
// (value of primitive-types or slice of primitive-types).
//
// To compare two Values of not known type, compare the results of the Any method.
// Using == on two Values is disallowed.
type Value struct {
	_ [0]func() // disallow ==
	// num holds a numeric value when it's single value, and hold slice's len when it's a slice value.
	num uint64
	// any holds a Type when it's a single value, and hold a pointer to the slice when it's slice value.
	//
	// This implementation takes advantage of compiler interface optimization:
	// - Compiler adds global [256]byte array called 'staticbytes' to every binary.
	// - So putting single-byte value into an interface{} do not allocate.
	//
	// ref: https://commaok.xyz/post/interface-allocs
	any any
}

// Return the underlying type the Value holds.
func (v Value) Type() Type {
	if typ, ok := v.any.(Type); ok {
		return typ
	}
	switch v.any.(type) {
	case stringptr:
		return TypeString
	case *bool:
		return TypeSliceBool
	case *int8:
		return TypeSliceInt8
	case *uint8:
		return TypeSliceUint8
	case *int16:
		return TypeSliceInt16
	case *uint16:
		return TypeSliceUint16
	case *int32:
		return TypeSliceInt32
	case *uint32:
		return TypeSliceUint32
	case *int64:
		return TypeSliceInt64
	case *uint64:
		return TypeSliceUint64
	case *float32:
		return TypeSliceFloat32
	case *float64:
		return TypeSliceFloat64
	case *string:
		return TypeSliceString
	}

	return TypeInvalid
}

// Int8 returns Value as int8, if it's not a valid int8 value, it returns basetype.Sint8Invalid (0x7F).
func (v Value) Int8() int8 {
	if v.any != TypeInt8 {
		return basetype.Sint8Invalid
	}
	return int8(v.num)
}

// Bool returns Value as bool, if it's not a valid bool value, it returns false.
func (v Value) Bool() bool {
	if v.any != TypeBool || v.num != 1 {
		return false
	}
	return true
}

// Uint8 returns Value as uint8, if it's not a valid uint8 value, it returns basetype.Uint8Invalid (0xFF).
func (v Value) Uint8() uint8 {
	if v.any != TypeUint8 {
		return basetype.Uint8Invalid
	}
	return uint8(v.num)
}

// Uint8z returns Value as uint8, if it's not a valid uint8 value, it returns basetype.Uint8zInvalid (0).
func (v Value) Uint8z() uint8 {
	if v.any != TypeUint8 {
		return basetype.Uint8zInvalid
	}
	return uint8(v.num)
}

// Int16 returns Value as int16, if it's not a valid int16 value, it returns basetype.Sint16Invalid (0x7FFF).
func (v Value) Int16() int16 {
	if v.any != TypeInt16 {
		return basetype.Sint16Invalid
	}
	return int16(v.num)
}

// Uint16 returns Value as uint16, if it's not a valid uint16 value, it returns basetype.Uint16Invalid (0xFFFF).
func (v Value) Uint16() uint16 {
	if v.any != TypeUint16 {
		return basetype.Uint16Invalid
	}
	return uint16(v.num)
}

// Uint16z returns Value as uint16, if it's not a valid uint16 value, it returns basetype.Uint16zInvalid (0).
func (v Value) Uint16z() uint16 {
	if v.any != TypeUint16 {
		return basetype.Uint16zInvalid
	}
	return uint16(v.num)
}

// Int32 returns Value as int32, if it's not a valid int32 value, it returns basetype.Sint32Invalid (0x7FFFFFFF).
func (v Value) Int32() int32 {
	if v.any != TypeInt32 {
		return basetype.Sint32Invalid
	}
	return int32(v.num)
}

// Uint32 returns Value as uint32, if it's not a valid uint32 value, it returns basetype.Uint32Invalid (0xFFFFFFFF).
func (v Value) Uint32() uint32 {
	if v.any != TypeUint32 {
		return basetype.Uint32Invalid
	}
	return uint32(v.num)
}

// Uint32z returns Value as uint32, if it's not a valid uint32 value, it returns basetype.Uint32zInvalid (0).
func (v Value) Uint32z() uint32 {
	if v.any != TypeUint32 {
		return basetype.Uint32zInvalid
	}
	return uint32(v.num)
}

// Int64 returns Value as int64, if it's not a valid int64 value, it returns basetype.Sint64Invalid (0x7FFFFFFFFFFFFFFF).
func (v Value) Int64() int64 {
	if v.any != TypeInt64 {
		return basetype.Sint64Invalid
	}
	return int64(v.num)
}

// Uint64 returns Value as uint64, if it's not a valid uint64 value, it returns basetype.Uint64Invalid (0xFFFFFFFFFFFFFFFF).
func (v Value) Uint64() uint64 {
	if v.any != TypeUint64 {
		return basetype.Uint64Invalid
	}
	return v.num
}

// Uint64z returns Value as uint64, if it's not a valid uint64 value, it returns basetype.Uint64Invalid (0).
func (v Value) Uint64z() uint64 {
	if v.any != TypeUint64 {
		return basetype.Uint64zInvalid
	}
	return uint64(v.num)
}

// Float32 returns Value as float32, if it's not a valid float32 value, it returns basetype.Float32Invalid (0xFFFFFFFF) in float32 value.
func (v Value) Float32() float32 {
	if v.any != TypeFloat32 {
		return math.Float32frombits(basetype.Float32Invalid)
	}
	return math.Float32frombits(uint32(v.num))
}

// Float64 returns Value as float64, if it's not a valid float64 value, it returns basetype.Float64Invalid (0xFFFFFFFFFFFFFFFF) in float64 value.
func (v Value) Float64() float64 {
	if v.any != TypeFloat64 {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return math.Float64frombits(v.num)
}

// String returns Value as string, if it's not a valid string value, it returns basetype.StringInvalid.
// This should not be treated as a Go's String method, use Any() if you want to print the underlying value.
func (v Value) String() string {
	ptr, ok := v.any.(stringptr)
	if !ok {
		return basetype.StringInvalid
	}
	return unsafe.String(ptr, v.num)
}

// SliceBool returns Value as []bool, if it's not a valid []bool value, it returns nil.
func (v Value) SliceBool() []bool {
	ptr, ok := v.any.(*bool)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceInt8 returns Value as []int8, if it's not a valid []int8 value, it returns nil.
func (v Value) SliceInt8() []int8 {
	ptr, ok := v.any.(*int8)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceUint8 returns Value as []uint8, if it's not a valid []uint8 value, it returns nil.
func (v Value) SliceUint8() []uint8 {
	ptr, ok := v.any.(*uint8)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceInt16 returns Value as []int16, if it's not a valid []int16 value, it returns nil.
func (v Value) SliceInt16() []int16 {
	ptr, ok := v.any.(*int16)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceUint16 returns Value as []uint16, if it's not a valid []uint16 value, it returns nil.
func (v Value) SliceUint16() []uint16 {
	ptr, ok := v.any.(*uint16)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceInt32 returns Value as []int32, if it's not a valid []int32 value, it returns nil.
func (v Value) SliceInt32() []int32 {
	ptr, ok := v.any.(*int32)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceUint32 returns Value as []uint32, if it's not a valid []uint32 value, it returns nil.
func (v Value) SliceUint32() []uint32 {
	ptr, ok := v.any.(*uint32)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceInt64 returns Value as []int64, if it's not a valid []int64 value, it returns nil.
func (v Value) SliceInt64() []int64 {
	ptr, ok := v.any.(*int64)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceUint64 returns Value as []uint64, if it's not a valid []uint64 value, it returns nil.
func (v Value) SliceUint64() []uint64 {
	ptr, ok := v.any.(*uint64)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceFloat32 returns Value as []float32, if it's not a valid []float32 value, it returns nil.
func (v Value) SliceFloat32() []float32 {
	ptr, ok := v.any.(*float32)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceFloat64 returns Value as []float64, if it's not a valid []float64 value, it returns nil.
func (v Value) SliceFloat64() []float64 {
	ptr, ok := v.any.(*float64)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// SliceString returns Value as []string, if it's not a valid []string value, it returns nil.
func (v Value) SliceString() []string {
	ptr, ok := v.any.(*string)
	if !ok {
		return nil
	}
	return unsafe.Slice(ptr, v.num)
}

// Any returns Value's underlying value.
func (v Value) Any() any {
	switch v.Type() {
	case TypeBool:
		return v.Bool()
	case TypeInt8:
		return v.Int8()
	case TypeUint8:
		return v.Uint8()
	case TypeInt16:
		return v.Int16()
	case TypeUint16:
		return v.Uint16()
	case TypeInt32:
		return v.Int32()
	case TypeUint32:
		return v.Uint32()
	case TypeInt64:
		return v.Int64()
	case TypeUint64:
		return v.Uint64()
	case TypeFloat32:
		return v.Float32()
	case TypeFloat64:
		return v.Float64()
	case TypeString:
		return v.String()
	case TypeSliceBool:
		return v.SliceBool()
	case TypeSliceInt8:
		return v.SliceInt8()
	case TypeSliceUint8:
		return v.SliceUint8()
	case TypeSliceInt16:
		return v.SliceInt16()
	case TypeSliceUint16:
		return v.SliceUint16()
	case TypeSliceInt32:
		return v.SliceInt32()
	case TypeSliceUint32:
		return v.SliceUint32()
	case TypeSliceInt64:
		return v.SliceInt64()
	case TypeSliceUint64:
		return v.SliceUint64()
	case TypeSliceFloat32:
		return v.SliceFloat32()
	case TypeSliceFloat64:
		return v.SliceFloat64()
	case TypeSliceString:
		return v.SliceString()
	}
	return nil
}

// Bool converts bool as Value.
func Bool(v bool) Value {
	var num uint64
	if v {
		num = 1
	}
	return Value{num: num, any: TypeBool}
}

// Int8 converts int8 as Value.
func Int8(v int8) Value {
	return Value{num: uint64(v), any: TypeInt8}
}

// Uint8 converts uint8 as Value.
func Uint8(v uint8) Value {
	return Value{num: uint64(v), any: TypeUint8}
}

// Int16 converts int16 as Value.
func Int16(v int16) Value {
	return Value{num: uint64(v), any: TypeInt16}
}

// Uint16 converts uint16 as Value.
func Uint16(v uint16) Value {
	return Value{num: uint64(v), any: TypeUint16}
}

// Int32 converts int32 as Value.
func Int32(v int32) Value {
	return Value{num: uint64(v), any: TypeInt32}
}

// Uint32 converts uint32 as Value.
func Uint32(v uint32) Value {
	return Value{num: uint64(v), any: TypeUint32}
}

// Int64 converts int64 as Value.
func Int64(v int64) Value {
	return Value{num: uint64(v), any: TypeInt64}
}

// Uint64 converts uint64 as Value.
func Uint64(v uint64) Value {
	return Value{num: v, any: TypeUint64}
}

// Float32 converts float32 as Value.
func Float32(v float32) Value {
	return Value{num: uint64(math.Float32bits(v)), any: TypeFloat32}
}

// Float64 converts float64 as Value.
func Float64(v float64) Value {
	return Value{num: math.Float64bits(v), any: TypeFloat64}
}

// String converts string as Value.
func String(v string) Value {
	return Value{num: uint64(len(v)), any: stringptr(unsafe.StringData(v))}
}

// HACK: The use of *(*[]ArbitraryType)(unsafe.Pointer(&slice) below should be safe (in unsafe world) since we only use it to
// temporarily cast the type to make unsafe.SliceData return the pointer as *ArbitraryType. The actual slice is handled by
// unsafe.SliceData so we don't lose the data.

// SliceBool converts []bool as Value.
func SliceBool[S []E, E ~bool](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]bool)(unsafe.Pointer(&s)))}
}

// SliceInt8 converts []int8 as Value.
func SliceInt8[S []E, E ~int8](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]int8)(unsafe.Pointer(&s)))}
}

// SliceUint8 converts []uint8 as Value.
func SliceUint8[S []E, E ~uint8](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]uint8)(unsafe.Pointer(&s)))}
}

// SliceInt16 converts []int16 as Value.
func SliceInt16[S []E, E ~int16](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]int16)(unsafe.Pointer(&s)))}
}

// SliceUint16 converts []uint16 as Value.
func SliceUint16[S []E, E ~uint16](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]uint16)(unsafe.Pointer(&s)))}
}

// SliceInt32 converts []int32 as Value.
func SliceInt32[S []E, E ~int32](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]int32)(unsafe.Pointer(&s)))}
}

// SliceUint32 converts []uint32 as Value.
func SliceUint32[S []E, E ~uint32](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]uint32)(unsafe.Pointer(&s)))}
}

// SliceInt64 converts []int64 as Value.
func SliceInt64[S []E, E ~int64](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]int64)(unsafe.Pointer(&s)))}
}

// SliceUint64 converts []uint64 as Value.
func SliceUint64[S []E, E ~uint64](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]uint64)(unsafe.Pointer(&s)))}
}

// SliceFloat32 converts []float32 as Value.
func SliceFloat32[S []E, E ~float32](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]float32)(unsafe.Pointer(&s)))}
}

// SliceFloat64 converts []float64 as Value.
func SliceFloat64[S []E, E ~float64](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]float64)(unsafe.Pointer(&s)))}
}

// SliceString converts []string as Value.
func SliceString[S []E, E ~string](s S) Value {
	return Value{num: uint64(len(s)), any: unsafe.SliceData(*(*[]string)(unsafe.Pointer(&s)))}
}

// Any converts any value into Value. If the given v is not a primitive-type value (or a slice of primitive-type)
// it will determine it using reflection, and if it's a non-primitive-type slice it will make 1 alloc.
// If v is not supported such as int, uint, []int, []uint, []any, slice with zero len, etc., Value with TypeInvalid is returned.
func Any(v any) Value {
	switch val := v.(type) { // Fast path
	case int, uint, []int, []uint, []any: // Fast return on invalid value
		return Value{any: TypeInvalid}
	case Value:
		return val
	case bool:
		return Bool(val)
	case int8:
		return Int8(val)
	case uint8:
		return Uint8(val)
	case int16:
		return Int16(val)
	case uint16:
		return Uint16(val)
	case int32:
		return Int32(val)
	case uint32:
		return Uint32(val)
	case int64:
		return Int64(val)
	case uint64:
		return Uint64(val)
	case float32:
		return Float32(val)
	case float64:
		return Float64(val)
	case string:
		return String(val)
	case []int8:
		return SliceInt8(val)
	case []uint8:
		return SliceUint8(val)
	case []int16:
		return SliceInt16(val)
	case []uint16:
		return SliceUint16(val)
	case []int32:
		return SliceInt32(val)
	case []uint32:
		return SliceUint32(val)
	case []int64:
		return SliceInt64(val)
	case []uint64:
		return SliceUint64(val)
	case []float32:
		return SliceFloat32(val)
	case []float64:
		return SliceFloat64(val)
	case []string:
		return SliceString(val)
	}

	// Fallback to reflection.
	rv := reflect.Indirect(reflect.ValueOf(v))
	switch rv.Kind() {
	case reflect.Bool:
		return Bool(rv.Bool())
	case reflect.Int8:
		return Int8(int8(rv.Int()))
	case reflect.Uint8:
		return Uint8(uint8(rv.Uint()))
	case reflect.Int16:
		return Int16(int16(rv.Int()))
	case reflect.Uint16:
		return Uint16(uint16(rv.Uint()))
	case reflect.Int32:
		return Int32(int32(rv.Int()))
	case reflect.Uint32:
		return Uint32(uint32(rv.Uint()))
	case reflect.Int64:
		return Int64(int64(rv.Int()))
	case reflect.Uint64:
		return Uint64(uint64(rv.Uint()))
	case reflect.Float32:
		return Float32(float32(rv.Float()))
	case reflect.Float64:
		return Float64(float64(rv.Float()))
	case reflect.String:
		return String(rv.String())
	case reflect.Slice: // Always alloc since it makes new slice.
		if rv.Len() == 0 {
			return Value{any: TypeInvalid}
		}
		switch rv.Index(0).Kind() {
		case reflect.Bool:
			var vals = make([]bool, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = rv.Index(i).Bool()
			}
			return SliceBool(vals)
		case reflect.Int8:
			var vals = make([]int8, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = int8(rv.Index(i).Int())
			}
			return SliceInt8(vals)
		case reflect.Uint8:
			var vals = make([]uint8, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = uint8(rv.Index(i).Uint())
			}
			return SliceUint8(vals)
		case reflect.Int16:
			var vals = make([]int16, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = int16(rv.Index(i).Int())
			}
			return SliceInt16(vals)
		case reflect.Uint16:
			var vals = make([]uint16, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = uint16(rv.Index(i).Uint())
			}
			return SliceUint16(vals)
		case reflect.Int32:
			var vals = make([]int32, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = int32(rv.Index(i).Int())
			}
			return SliceInt32(vals)
		case reflect.Uint32:
			var vals = make([]uint32, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = uint32(rv.Index(i).Uint())
			}
			return SliceUint32(vals)
		case reflect.Int64:
			var vals = make([]int64, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = int64(rv.Index(i).Int())
			}
			return SliceInt64(vals)
		case reflect.Uint64:
			var vals = make([]uint64, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = uint64(rv.Index(i).Uint())
			}
			return SliceUint64(vals)
		case reflect.Float32:
			var vals = make([]float32, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = float32(rv.Index(i).Float())
			}
			return SliceFloat32(vals)
		case reflect.Float64:
			var vals = make([]float64, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = float64(rv.Index(i).Float())
			}
			return SliceFloat64(vals)
		case reflect.String:
			var vals = make([]string, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vals[i] = string(rv.Index(i).String())
			}
			return SliceString(vals)
		}
	}

	return Value{any: TypeInvalid}
}

// Sizeof returns the size of val in bytes. For every string in Value, if the last index of the string is not '\x00', size += 1.
func Sizeof(val Value, baseType basetype.BaseType) int {
	return lenof(val) * int(baseType.Size())
}

func lenof(val Value) int {
	switch val.Type() { // Fast Path
	case TypeInvalid:
		return 0
	case TypeBool,
		TypeInt8,
		TypeUint8,
		TypeInt16,
		TypeUint16,
		TypeInt32,
		TypeUint32,
		TypeFloat32,
		TypeFloat64,
		TypeInt64,
		TypeUint64:
		return 1
	case TypeString:
		s := val.String()
		if len(s) == 0 {
			return 1 // utf-8 null terminated string
		}
		if l := len(s); l > 0 && s[l-1] == '\x00' {
			return l
		}
		return len(s) + 1
	case TypeSliceString:
		vs := val.SliceString()
		var size int
		for i := range vs {
			if len(vs[i]) == 0 {
				size += 1 // utf-8 null terminated string
				continue
			}
			if l := len(vs[i]); l > 0 && vs[i][l-1] == '\x00' {
				size += l
				continue
			}
			size += len(vs[i]) + 1
		}
		if size == 0 {
			return 1 // utf-8 null terminated string
		}
		return size
	}

	return int(val.num) // other slices
}
