// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fitcsv

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// format formats given val into string
//
//	custom formating:
//	- float64: 0 -> 0.0, 1,2300 -> 1.23, 1,2340001 -> 1,2340001
//
//	slice formatting: every value element will be separated by '|' pipe symbol.
//	- []uint8{1,2,3} -> "1|2|3"
func format(val proto.Value) string {
	switch val.Type() { // fast path
	case proto.TypeBool:
		return strconv.FormatUint(uint64(val.Bool()), 10)
	case proto.TypeInt8:
		return strconv.FormatInt(int64(val.Int8()), 10)
	case proto.TypeUint8:
		return strconv.FormatUint(uint64(val.Uint8()), 10)
	case proto.TypeInt16:
		return strconv.FormatInt(int64(val.Int16()), 10)
	case proto.TypeUint16:
		return strconv.FormatUint(uint64(val.Uint16()), 10)
	case proto.TypeInt32:
		return strconv.FormatInt(int64(val.Int32()), 10)
	case proto.TypeUint32:
		return strconv.FormatUint(uint64(val.Uint32()), 10)
	case proto.TypeInt64:
		return strconv.FormatInt(int64(val.Uint64()), 10)
	case proto.TypeUint64:
		return strconv.FormatUint(uint64(val.Uint64()), 10)
	case proto.TypeFloat32:
		value := val.Float32()
		if value == float32(int64(value)) {
			return strconv.FormatFloat(float64(value), 'f', 1, 64)
		}
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case proto.TypeFloat64:
		value := val.Float64()
		if value == float64(int64(value)) {
			return strconv.FormatFloat(value, 'f', 1, 64)
		}
		return strconv.FormatFloat(value, 'g', -1, 64)
	case proto.TypeString:
		s := strings.Map(func(r rune) rune {
			if !unicode.IsPrint(r) || r == '"' {
				return -1
			}
			return r
		}, val.String())
		return s
	case proto.TypeSliceBool:
		return concat(val.SliceBool(), func(v typedef.Bool) string {
			return strconv.FormatUint(uint64(v), 10)
		})
	case proto.TypeSliceInt8:
		return concat(val.SliceInt8(), func(v int8) string {
			return strconv.FormatInt(int64(v), 10)
		})
	case proto.TypeSliceUint8:
		return concat(val.SliceUint8(), func(v uint8) string {
			return strconv.FormatUint(uint64(v), 10)
		})
	case proto.TypeSliceInt16:
		return concat(val.SliceInt16(), func(v int16) string {
			return strconv.FormatInt(int64(v), 10)
		})
	case proto.TypeSliceUint16:
		return concat(val.SliceUint16(), func(v uint16) string {
			return strconv.FormatUint(uint64(v), 10)
		})
	case proto.TypeSliceInt32:
		return concat(val.SliceInt32(), func(v int32) string {
			return strconv.FormatInt(int64(v), 10)
		})
	case proto.TypeSliceUint32:
		return concat(val.SliceUint32(), func(v uint32) string {
			return strconv.FormatUint(uint64(v), 10)
		})
	case proto.TypeSliceInt64:
		return concat(val.SliceInt64(), func(v int64) string {
			return strconv.FormatInt(v, 10)
		})
	case proto.TypeSliceUint64:
		return concat(val.SliceUint64(), func(v uint64) string {
			return strconv.FormatUint(v, 10)
		})
	case proto.TypeSliceFloat32:
		return concat(val.SliceFloat32(), func(v float32) string {
			if v == float32(int32(v)) {
				return strconv.FormatFloat(float64(v), 'f', 1, 64)
			}
			return strconv.FormatFloat(float64(v), 'g', -1, 64)
		})
	case proto.TypeSliceFloat64:
		return concat(val.SliceFloat64(), func(v float64) string {
			if v == float64(int64(v)) {
				return strconv.FormatFloat(v, 'f', 1, 64)
			}
			return strconv.FormatFloat(v, 'g', -1, 64)
		})
	case proto.TypeSliceString:
		vs := val.SliceString()
		for i := range vs {
			vs[i] = strings.Map(func(r rune) rune {
				if !unicode.IsPrint(r) || r == '"' {
					return -1
				}
				return r
			}, vs[i])
		}
		return concat(vs, func(v string) string {
			return v
		})
	default:
		return fmt.Sprintf("%v", val)
	}
}

func concat[S []E, E any](s S, formatFn func(v E) string) string {
	var buf strings.Builder
	for i := range s {
		buf.WriteString(formatFn(s[i]))
		if i < len(s)-1 {
			buf.WriteByte('|')
		}
	}
	return buf.String()
}
