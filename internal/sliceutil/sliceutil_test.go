// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sliceutil

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/google/go-cmp/cmp"
)

func TestClone(t *testing.T) {
	tt := []struct {
		name     string
		values   []int
		expected []int
	}{
		{
			name:     "nil",
			values:   nil,
			expected: nil,
		},
		{
			name:     "has one value",
			values:   []int{0},
			expected: []int{0},
		},
		{
			name:     "has zero value",
			values:   []int{},
			expected: nil,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			s := Clone(tc.values)
			if tc.values == nil {
				if s != nil {
					t.Fatalf("expected nil")
				}
				return
			}
			if p1, p2 := unsafe.SliceData(s), unsafe.SliceData(tc.values); len(tc.values) == 0 && p1 == p2 {
				t.Fatalf("cloned values (%v) shares the same underlying array: %p == %p", s, p1, p2)
			}
			if tc.expected == nil {
				return
			}
			if diff := cmp.Diff(tc.values, tc.expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
