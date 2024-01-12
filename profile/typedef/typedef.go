// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"reflect"

	"github.com/muktihari/fit/profile/basetype"
)

// Sizeof returns the size of val in bytes.
func Sizeof(val any, baseType basetype.BaseType) int {
	return lenOf(val) * int(baseType.Size())
}

func lenOf(val any) int {
	switch vs := val.(type) { // Fast Path
	case int8, uint8, int16, uint16, int32, uint32, float32, float64, int64, uint64:
		return 1
	case string:
		if len(vs) == 0 {
			return 1 // utf-8 null terminated string
		}
		if l := len(vs); l > 0 && vs[l-1] == '\x00' {
			return l
		}
		return len(vs) + 1
	case []string:
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
	case []int8:
		return len(vs)
	case []uint8:
		return len(vs)
	case []int16:
		return len(vs)
	case []uint16:
		return len(vs)
	case []int32:
		return len(vs)
	case []uint32:
		return len(vs)
	case []float32:
		return len(vs)
	case []float64:
		return len(vs)
	case []int64:
		return len(vs)
	case []uint64:
		return len(vs)
	}

	// Fallback to reflection
	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.String:
		val := rv.String()
		if len(val) == 0 {
			return 1 // utf-8 null terminated string
		}
		if l := len(val); l > 0 && val[l-1] == '\x00' {
			return l
		}
		return len(val) + 1
	case reflect.Slice:
		if rv.Type().Elem().Kind() == reflect.String {
			var size int
			for i := 0; i < rv.Len(); i++ {
				val := rv.Index(i).String()
				if len(val) == 0 {
					size += 1 // utf-8 null terminated string
					continue
				}
				if l := len(val); l > 0 && val[l-1] == '\x00' {
					size += l
					continue
				}
				size += len(val) + 1
			}
			if size == 0 {
				return 1 // utf-8 null terminated string
			}
			return size
		}
		return rv.Len()
	default:
		return 1
	}
}
