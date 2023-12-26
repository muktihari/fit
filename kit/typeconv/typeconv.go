// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv

// IntegerToInt64 converts any integer value of val to int64, if val is non-integer value return false.
func IntegerToInt64(val any) (int64, bool) {
	switch v := val.(type) {
	case int8:
		return int64(v), true
	case uint8:
		return int64(v), true
	case int16:
		return int64(v), true
	case uint16:
		return int64(v), true
	case int32:
		return int64(v), true
	case uint32:
		return int64(v), true
	case int64:
		return v, true
	case uint64:
		return int64(v), true
	}
	return 0, false
}
