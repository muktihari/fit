// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/muktihari/fit/profile/basetype"
)

// Unmarshal unmarshals b into a proto.Value.
// The caller should ensure that the length of the given b matches its corresponding base type's size, otherwise it might panic.
func Unmarshal(b []byte, bo binary.ByteOrder, ref basetype.BaseType, isArray bool) (Value, error) {
	switch ref {
	case basetype.Sint8:
		if isArray {
			vs := make([]int8, 0, len(b))
			for i := 0; i < len(b); i++ {
				vs = append(vs, int8(b[i]))
			}
			return SliceInt8(vs), nil
		}
		return Int8(int8(b[0])), nil
	case basetype.Enum, basetype.Byte,
		basetype.Uint8, basetype.Uint8z:
		if isArray {
			vals := make([]byte, len(b))
			copy(vals, b)
			return SliceUint8(vals), nil
		}
		return Uint8(b[0]), nil
	case basetype.Sint16:
		if isArray {
			const n = 2
			vs := make([]int16, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, int16(bo.Uint16(b[i:i+n])))
			}
			return SliceInt16(vs), nil
		}
		return Int16(int16(bo.Uint16(b))), nil
	case basetype.Uint16, basetype.Uint16z:
		if isArray {
			const n = 2
			vs := make([]uint16, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, bo.Uint16(b[i:i+n]))
			}
			return SliceUint16(vs), nil
		}
		return Uint16(bo.Uint16(b)), nil
	case basetype.Sint32:
		if isArray {
			const n = 4
			vs := make([]int32, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, int32(bo.Uint32(b[i:i+n])))
			}
			return SliceInt32(vs), nil
		}
		return Int32(int32(bo.Uint32(b))), nil
	case basetype.Uint32, basetype.Uint32z:
		if isArray {
			const n = 4
			vs := make([]uint32, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, bo.Uint32(b[i:i+n]))
			}
			return SliceUint32(vs), nil
		}
		return Uint32(bo.Uint32(b)), nil
	case basetype.Sint64:
		if isArray {
			const n = 8
			vs := make([]int64, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, int64(bo.Uint64(b[i:i+n])))
			}
			return SliceInt64(vs), nil
		}
		return Int64(int64(bo.Uint64(b))), nil
	case basetype.Uint64, basetype.Uint64z:
		if isArray {
			const n = 8
			vs := make([]uint64, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, bo.Uint64(b[i:i+n]))
			}
			return SliceUint64(vs), nil
		}
		return Uint64(bo.Uint64(b)), nil
	case basetype.Float32:
		if isArray {
			const n = 4
			vs := make([]float32, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, math.Float32frombits(bo.Uint32(b[i:i+n])))
			}
			return SliceFloat32(vs), nil
		}
		return Float32(math.Float32frombits(bo.Uint32(b))), nil
	case basetype.Float64:
		if isArray {
			const n = 8
			vs := make([]float64, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, math.Float64frombits(bo.Uint64(b[i:i+n])))
			}
			return SliceFloat64(vs), nil
		}
		return Float64(math.Float64frombits(bo.Uint64(b))), nil
	case basetype.String:
		if isArray {
			var size byte
			last := 0
			for i := range b {
				if b[i] == '\x00' {
					if last != i { // only if not an invalid string
						size++
					}
					last = i + 1
				}
			}
			last = 0
			vs := make([]string, 0, size)
			for i := range b {
				if b[i] == '\x00' {
					if last != i { // only if not an invalid string
						vs = append(vs, string(b[last:i]))
					}
					last = i + 1
				}
			}
			return SliceString(vs), nil
		}
		b = trimUTF8NullTerminatedString(b)
		return String(string(b)), nil
	}

	return Value{}, fmt.Errorf("type %s(%d) is not supported: %w", ref, ref, ErrTypeNotSupported)
}

// Note: The size may be a multiple of the underlying FIT Base Type size indicating the field contains multiple elements represented as an array.
func size(lenbytes int, typesize byte) byte {
	return byte(lenbytes % int(typesize))
}

func trimUTF8NullTerminatedString(b []byte) []byte {
	pos := -1
	for i := len(b); i > 0; i-- {
		if b[i-1] != '\x00' {
			break
		}
		pos = i - 1
	}
	if pos < 0 {
		return b
	}
	return b[:pos]
}
