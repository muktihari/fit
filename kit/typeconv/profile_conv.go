// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv

import (
	"reflect"
)

// ToBool converts any boolean value in the form of ~bool or ~byte to T bool.
func ToBool[T ~bool](val any) T {
	if val == nil {
		return false
	}

	switch v := val.(type) {
	case T:
		return v
	case bool:
		return T(v)
	case byte:
		var bv T
		if v == 1 {
			bv = true
		}
		return bv
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Bool:
		return T(rv.Bool())
	case reflect.Uint8:
		value := rv.Uint()
		var bv T
		if value == 1 {
			bv = true
		}
		return bv
	}

	return false
}

func ToSliceBool[T ~bool](val any) []T {
	if val == nil {
		return nil
	}

	switch vals := val.(type) {
	case []T:
		return vals
	case []bool:
		b := make([]T, 0, len(vals))
		for _, value := range vals {
			b = append(b, T(value))
		}
		return b
	case []uint8:
		b := make([]T, 0, len(vals))
		for _, value := range vals {
			b = append(b, ToBool[T](value))
		}
		return b
	}

	rv := reflect.ValueOf(val)

	if rv.Kind() != reflect.Slice {
		return nil
	}

	if rv.Len() == 0 {
		return nil
	}

	switch rv.Index(0).Kind() {
	case reflect.Bool:
		vs := make([]T, 0, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			vs = append(vs, T(rv.Index(i).Bool()))
		}
		return vs
	case reflect.Uint8:
		vs := make([]T, 0, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			var vb T
			if rv.Index(i).Uint() == 1 {
				vb = true
			}
			vs = append(vs, vb)
		}
		return vs
	}

	return nil
}
