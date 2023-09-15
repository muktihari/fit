// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutil_test

import (
	"testing"

	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
)

func TestTrimRepeatedChar(t *testing.T) {
	tt := []struct {
		s string
		r string
	}{
		{s: "a,,b,,,c", r: "a,b,c"},
		{s: ",,a,,b,,,c", r: ",a,b,c"},
	}
	for _, tc := range tt {
		t.Run(tc.s, func(t *testing.T) {
			r := strutil.TrimRepeatedChar(tc.s, ',')
			if r != tc.r {
				t.Fatalf("expected: %s, got: %s", tc.r, r)
			}
		})
	}
}
