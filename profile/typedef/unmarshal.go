// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/muktihari/fit/profile/basetype"
)

// Unmarshal unmarshals b into a primitive-typed value.
// The caller should ensure that the length of the given b matches its corresponding base type's size, otherwise it might panic.
func Unmarshal(b []byte, bo binary.ByteOrder, ref basetype.BaseType, isArray bool) (any, error) {
	switch ref {
	case basetype.Enum, basetype.Byte:
		if isArray {
			return b, nil
		}
		return b[0], nil
	case basetype.Sint8:
		if isArray {
			vs := make([]int8, 0, len(b))
			for i := 0; i < len(b); i++ {
				vs = append(vs, int8(b[i]))
			}
			return vs, nil
		}
		return int8(b[0]), nil
	case basetype.Uint8, basetype.Uint8z:
		if isArray {
			return b, nil
		}
		return b[0], nil
	case basetype.Sint16:
		if isArray {
			const n = 2
			vs := make([]int16, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, int16(bo.Uint16(b[i:i+n])))
			}
			return vs, nil
		}
		return int16(bo.Uint16(b)), nil
	case basetype.Uint16, basetype.Uint16z:
		if isArray {
			const n = 2
			vs := make([]uint16, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, bo.Uint16(b[i:i+n]))
			}
			return vs, nil
		}
		return bo.Uint16(b), nil
	case basetype.Sint32:
		if isArray {
			const n = 4
			vs := make([]int32, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, int32(bo.Uint32(b[i:i+n])))
			}
			return vs, nil
		}
		return int32(bo.Uint32(b)), nil
	case basetype.Uint32, basetype.Uint32z:
		if isArray {
			const n = 4
			vs := make([]uint32, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, bo.Uint32(b[i:i+n]))
			}
			return vs, nil
		}
		return bo.Uint32(b), nil
	case basetype.Sint64:
		if isArray {
			const n = 8
			vs := make([]int64, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, int64(bo.Uint64(b[i:i+n])))
			}
			return vs, nil
		}
		return int64(bo.Uint64(b)), nil
	case basetype.Uint64, basetype.Uint64z:
		if isArray {
			const n = 8
			vs := make([]uint64, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, bo.Uint64(b[i:i+n]))
			}
			return vs, nil
		}
		return bo.Uint64(b), nil
	case basetype.Float32:
		if isArray {
			const n = 4
			vs := make([]float32, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, math.Float32frombits(bo.Uint32(b[i:i+n])))
			}
			return vs, nil
		}
		return math.Float32frombits(bo.Uint32(b)), nil
	case basetype.Float64:
		if isArray {
			const n = 8
			vs := make([]float64, 0, size(len(b), n))
			for i := 0; i < len(b); i += n {
				vs = append(vs, math.Float64frombits(bo.Uint64(b[i:i+n])))
			}
			return vs, nil
		}
		return math.Float64frombits(bo.Uint64(b)), nil
	case basetype.String:
		if isArray {
			var size byte
			for i := range b {
				if b[i] == '\x00' {
					size++
				}
			}
			last := 0
			vs := make([]string, 0, size)
			for i := range b {
				if b[i] == '\x00' {
					vs = append(vs, string(b[last:i]))
					last = i + i
				}
			}
			return vs, nil
		}

		if len(b) != 0 && b[len(b)-1] == '\x00' {
			b = b[:len(b)-1] // trim utf-8 null-terminated string
		}
		return string(b), nil
	}

	return nil, fmt.Errorf("type %s(%d) is not supported: %w", ref, ref, ErrTypeNotSupported)
}

// Note: The size may be a multiple of the underlying FIT Base Type size indicating the field contains multiple elements represented as an array.
func size(lenbytes int, typesize byte) byte {
	return byte(lenbytes % int(typesize))
}
