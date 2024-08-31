// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semicircles_test

import (
	"fmt"
	"testing"

	"github.com/muktihari/fit/kit/semicircles"
)

func TestToDegrees(t *testing.T) {
	tt := []struct {
		semicircles int32
		degrees     float64
	}{
		{semicircles: -91900524, degrees: -7.703012935817242},
		{semicircles: 1319585919, degrees: 110.6064139958471},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.semicircles), func(t *testing.T) {
			result := semicircles.ToDegrees(tc.semicircles)
			if result != tc.degrees {
				t.Fatalf("expected: %v, got: %v", tc.degrees, result)
			}
		})
	}
}

func TestToSemicircles(t *testing.T) {
	tt := []struct {
		degrees     float64
		semicircles int32
	}{
		{degrees: -7.703012935817242, semicircles: -91900524},
		{degrees: 110.6064139958471, semicircles: 1319585919},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%#v", tc.degrees), func(t *testing.T) {
			result := semicircles.ToSemicircles(tc.degrees)
			if result != tc.semicircles {
				t.Fatalf("expected: %v, got: %v", tc.semicircles, result)
			}
		})
	}
}
