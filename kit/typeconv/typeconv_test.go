// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeconv_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/typeconv"
)

func TestIntegerToInt64(t *testing.T) {
	tt := []struct {
		value  any
		result int64
		ok     bool
	}{
		{value: nil, result: 0, ok: false},
		{value: int8(10), result: 10, ok: true},
		{value: uint8(10), result: 10, ok: true},
		{value: int16(10), result: 10, ok: true},
		{value: uint16(10), result: 10, ok: true},
		{value: int32(10), result: 10, ok: true},
		{value: uint32(10), result: 10, ok: true},
		{value: float32(10.1), result: 0, ok: false},
		{value: float64(10.2), result: 0, ok: false},
		{value: int64(10), result: 10, ok: true},
		{value: uint64(10), result: 10, ok: true},
		{value: string("fit"), result: 0, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%#v)", tc.value, tc.value), func(t *testing.T) {
			result, ok := typeconv.IntegerToInt64(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
