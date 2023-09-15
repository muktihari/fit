// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"reflect"
)

// Len returns how many element of slice val. It val is not a slice, return 1.
func Len(val any) byte {
	switch vs := val.(type) { // Fast Path
	case int8, uint8, int16, uint16, int32, uint32, float32, float64, int64, uint64:
		return 1
	case []int8:
		return byte(len(vs))
	case []uint8:
		return byte(len(vs))
	case []int16:
		return byte(len(vs))
	case []uint16:
		return byte(len(vs))
	case []int32:
		return byte(len(vs))
	case []uint32:
		return byte(len(vs))
	case string:
		return byte(len(vs))
	case []float32:
		return byte(len(vs))
	case []float64:
		return byte(len(vs))
	case []int64:
		return byte(len(vs))
	case []uint64:
		return byte(len(vs))
	}

	// Fallback to reflection
	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Slice {
		return 1
	}

	return byte(rv.Len())
}
