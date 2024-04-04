// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"
	"testing"
)

func TestIsValueEqualTo(t *testing.T) {
	tt := []struct {
		field Field
		value int64
		eq    bool
	}{
		{
			field: Field{Value: Uint8(89)},
			value: 89,
			eq:    true,
		},
		{
			field: Field{Value: String("fit")},
			value: 89,
			eq:    false,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v, %t", tc.value, tc.eq), func(t *testing.T) {
			if eq := tc.field.isValueEqualTo(tc.value); eq != tc.eq {
				t.Fatalf("expected: %t, got: %t", tc.eq, eq)
			}
		})
	}
}

func TestConvertToInt64(t *testing.T) {
	tt := []struct {
		value    Value
		expected int64
		ok       bool
	}{
		{value: Value{}, expected: 0, ok: false},
		{value: Int8(10), expected: 10, ok: true},
		{value: Uint8(10), expected: 10, ok: true},
		{value: Int16(10), expected: 10, ok: true},
		{value: Uint16(10), expected: 10, ok: true},
		{value: Int32(10), expected: 10, ok: true},
		{value: Uint32(10), expected: 10, ok: true},
		{value: Int64(10), expected: 10, ok: true},
		{value: Uint64(10), expected: 10, ok: true},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %v", i, tc.value.Any()), func(t *testing.T) {
			val, ok := convertToInt64(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected: %v, got: %v", tc.ok, ok)
			}
			if val != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, val)
			}
		})
	}
}
