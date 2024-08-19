// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"testing"
)

func TestTrimRightZero(t *testing.T) {
	tt := []struct {
		str      string
		expected string
	}{
		{str: "", expected: ""},
		{str: "\x00", expected: ""},
		{str: "Open Water", expected: "Open Water"},
		{str: "Open Water\x00", expected: "Open Water"},
		{str: "Open Water\x00\x00", expected: "Open Water"},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			res := trimRightZero([]byte(tc.str))
			if string(res) != tc.expected {
				t.Fatalf("expected: %s, got: %s", tc.expected, res)
			}
		})
	}
}

func BenchmarkTrimRightZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = trimRightZero([]byte(""))
		_ = trimRightZero([]byte("\x00"))
		_ = trimRightZero([]byte("Open Water"))
		_ = trimRightZero([]byte("Open Water\x00"))
		_ = trimRightZero([]byte("Open Water\x00\x00"))
	}
}

func TestUTF8String(t *testing.T) {
	tt := []struct {
		in  []byte
		out string
	}{
		{in: []byte("0000000000000�0000000"), out: "0000000000000�0000000"},
		{in: []byte("0000000000000\xe80000000"), out: "0000000000000�0000000"},
	}

	for _, tc := range tt {
		out := utf8String(tc.in)
		if out != tc.out {
			t.Fatalf("expected: %q, got: %q", tc.out, out)
		}
	}
}

func BenchmarkUTF8String(b *testing.B) {
	b.Run("valid utf8 string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = utf8String([]byte("0000000000000�0000000"))
		}
	})
	b.Run("invalid utf8 string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = utf8String([]byte("0000000000000\xe80000000"))
		}
	})
}
