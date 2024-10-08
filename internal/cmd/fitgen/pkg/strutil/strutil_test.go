// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strutil_test

import (
	"fmt"
	"testing"

	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
)

func BenchmarkToTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strutil.ToTitle("avg_speed")
	}
}

func TestToTitle(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		{
			input:  "avg_speed",
			output: "AvgSpeed",
		},
		{
			input:  "avg speed",
			output: "AvgSpeed",
		},
		{
			input:  "timestamp_32k",
			output: "Timestamp32K",
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.input), func(t *testing.T) {
			out := strutil.ToTitle(tc.input)
			if out != tc.output {
				t.Fatalf("expected: %q, got: %q", tc.output, out)
			}
		})
	}
}

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
