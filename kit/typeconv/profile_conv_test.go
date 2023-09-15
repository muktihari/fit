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

type Bool bool

func TestToBool(t *testing.T) {
	tt := []struct {
		value  any
		result bool
	}{
		{value: nil, result: false},
		{value: uint8(1), result: true},
		{value: true, result: true},
		{value: Uint8(1), result: true},
		{value: Bool(true), result: true},
		{value: int(1), result: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.value), func(t *testing.T) {
			result := typeconv.ToBool[bool](tc.value)
			if result != tc.result {
				t.Fatalf("expected: %t, got: %t", tc.result, result)
			}
			result2 := typeconv.ToBool[Bool](tc.value)
			if result2 != Bool(tc.result) {
				t.Fatalf("expected: %t, got: %t", tc.result, result2)
			}
		})
	}
}

func TestToSliceBool(t *testing.T) {
	tt := []struct {
		value  any
		result []bool
	}{
		{value: nil, result: nil},
		{value: []bool{true, false}, result: []bool{true, false}},
		{value: []Bool{true, false}, result: []bool{true, false}},
		{value: []uint8{1, 0}, result: []bool{true, false}},
		{value: []Uint8{1, 0}, result: []bool{true, false}},
		{value: []Uint8{}, result: nil},
		{value: uint16(1), result: nil},
		{value: []uint16{1, 0}, result: nil},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("[]bool: %#v", tc.value), func(t *testing.T) {
			result := typeconv.ToSliceBool[bool](tc.value)
			if diff := cmp.Diff(result, tc.result); diff != "" {
				t.Fatalf(diff)
			}
		})
	}

	t.Run("[]Bool", func(t *testing.T) {
		value := []bool{true, false}
		expected := []Bool{true, false}

		result := typeconv.ToSliceBool[Bool](value)
		if diff := cmp.Diff(result, expected); diff != "" {
			t.Fatalf(diff)
		}
	})
}
