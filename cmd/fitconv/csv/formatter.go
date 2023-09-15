// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"fmt"
	"strconv"
)

// format formats given val into string
//
//	custom formating:
//	  - float64: 0 -> 0.0, 1,2300 -> 1.23, 1,2340001 -> 1,2340001
func format(val any) string {
	switch v := val.(type) { // fast path
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case int64:
		return strconv.FormatInt(int64(v), 10)
	case uint64:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		if v == float32(int64(v)) {
			return strconv.FormatFloat(float64(v), 'f', 1, 64)
		}
		return strconv.FormatFloat(float64(v), 'g', -1, 64)
	case float64:
		if v == float64(int64(v)) {
			return strconv.FormatFloat(v, 'f', 1, 64)
		}
		return strconv.FormatFloat(v, 'g', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", val)
	}
}
