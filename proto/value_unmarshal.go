// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"encoding/binary"
	"fmt"
	"math"
	"sync"
	"unicode/utf8"

	"github.com/muktihari/fit/profile/basetype"
)

// UnmarshalValue unmarshals b into a proto.Value.
// The caller should ensure that the len(b) matches its corresponding base type's size, otherwise it might panic.
func UnmarshalValue(b []byte, arch byte, ref basetype.BaseType, isArray bool) (Value, error) {
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
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, int16(binary.LittleEndian.Uint16(b[:n])))
				} else {
					vs = append(vs, int16(binary.BigEndian.Uint16(b[:n])))
				}
			}
			return SliceInt16(vs), nil
		}
		if arch == littleEndian {
			return Int16(int16(binary.LittleEndian.Uint16(b))), nil
		}
		return Int16(int16(binary.BigEndian.Uint16(b))), nil
	case basetype.Uint16, basetype.Uint16z:
		if isArray {
			const n = 2
			vs := make([]uint16, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, binary.LittleEndian.Uint16(b[:n]))
				} else {
					vs = append(vs, binary.BigEndian.Uint16(b[:n]))
				}
			}
			return SliceUint16(vs), nil
		}
		if arch == littleEndian {
			return Uint16(binary.LittleEndian.Uint16(b)), nil
		}
		return Uint16(binary.BigEndian.Uint16(b)), nil
	case basetype.Sint32:
		if isArray {
			const n = 4
			vs := make([]int32, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, int32(binary.LittleEndian.Uint32(b[:n])))
				} else {
					vs = append(vs, int32(binary.BigEndian.Uint32(b[:n])))
				}
			}
			return SliceInt32(vs), nil
		}
		if arch == littleEndian {
			return Int32(int32(binary.LittleEndian.Uint32(b))), nil
		}
		return Int32(int32(binary.BigEndian.Uint32(b))), nil
	case basetype.Uint32, basetype.Uint32z:
		if isArray {
			const n = 4
			vs := make([]uint32, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, binary.LittleEndian.Uint32(b[:n]))
				} else {
					vs = append(vs, binary.BigEndian.Uint32(b[:n]))
				}
			}
			return SliceUint32(vs), nil
		}
		if arch == littleEndian {
			return Uint32(binary.LittleEndian.Uint32(b)), nil
		}
		return Uint32(binary.BigEndian.Uint32(b)), nil
	case basetype.Sint64:
		if isArray {
			const n = 8
			vs := make([]int64, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, int64(binary.LittleEndian.Uint64(b[:n])))
				} else {
					vs = append(vs, int64(binary.BigEndian.Uint64(b[:n])))
				}
			}
			return SliceInt64(vs), nil
		}
		if arch == littleEndian {
			return Int64(int64(binary.LittleEndian.Uint64(b))), nil
		}
		return Int64(int64(binary.BigEndian.Uint64(b))), nil
	case basetype.Uint64, basetype.Uint64z:
		if isArray {
			const n = 8
			vs := make([]uint64, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, binary.LittleEndian.Uint64(b[:n]))
				} else {
					vs = append(vs, binary.BigEndian.Uint64(b[:n]))
				}
			}
			return SliceUint64(vs), nil
		}
		if arch == littleEndian {
			return Uint64(binary.LittleEndian.Uint64(b)), nil
		}
		return Uint64(binary.BigEndian.Uint64(b)), nil
	case basetype.Float32:
		if isArray {
			const n = 4
			vs := make([]float32, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, math.Float32frombits(binary.LittleEndian.Uint32(b[:n])))
				} else {
					vs = append(vs, math.Float32frombits(binary.BigEndian.Uint32(b[:n])))
				}
			}
			return SliceFloat32(vs), nil
		}
		if arch == littleEndian {
			return Float32(math.Float32frombits(binary.LittleEndian.Uint32(b))), nil
		}
		return Float32(math.Float32frombits(binary.BigEndian.Uint32(b))), nil
	case basetype.Float64:
		if isArray {
			const n = 8
			vs := make([]float64, 0, size(len(b), n))
			for ; len(b) >= n; b = b[n:] {
				if arch == littleEndian {
					vs = append(vs, math.Float64frombits(binary.LittleEndian.Uint64(b[:n])))
				} else {
					vs = append(vs, math.Float64frombits(binary.BigEndian.Uint64(b[:n])))
				}
			}
			return SliceFloat64(vs), nil
		}
		if arch == littleEndian {
			return Float64(math.Float64frombits(binary.LittleEndian.Uint64(b))), nil
		}
		return Float64(math.Float64frombits(binary.BigEndian.Uint64(b))), nil
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
						vs = append(vs, utf8String(b[last:i]))
					}
					last = i + 1
				}
			}
			return SliceString(vs), nil
		}
		b = trimUTF8NullTerminatedString(b)
		return String(utf8String(b)), nil
	}

	return Value{}, fmt.Errorf("type %s(%d) is not supported: %w", ref, ref, ErrTypeNotSupported)
}

// Note: The size may be a multiple of the underlying FIT Base Type size indicating the field contains multiple elements represented as an array.
func size(lenbytes int, typesize byte) byte {
	return byte(lenbytes % int(typesize))
}

// trimUTF8NullTerminatedString trims all utf8 null-terminated string including the paddings.
func trimUTF8NullTerminatedString(b []byte) []byte {
	pos := len(b)
	for i := pos - 1; i >= 0; i-- {
		if b[i] != '\x00' {
			return b[:i+1]
		}
		pos = i
	}
	return b[:pos]
}

// smallpool is an [256]byte array pool.
var smallpool = sync.Pool{New: func() any { return new([256]byte) }}

// utf8String converts b into a valid utf8 string.
// Any invalid utf8 character will be converted into utf8.RuneError.
func utf8String(b []byte) string {
	if utf8.Valid(b) { // Fast path
		return string(b)
	}
	arr := smallpool.Get().(*[256]byte)
	defer smallpool.Put(arr)
	buf := arr[:0]
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		buf = utf8.AppendRune(buf, r)
		b = b[size:]
	}
	return string(buf)
}
