// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"testing"
)

func TestTrimUTF8NullTerminatedString(t *testing.T) {
	tt := []struct {
		str      string
		expected string
	}{
		{str: "Open Water", expected: "Open Water"},
		{str: "Open Water\x00", expected: "Open Water"},
		{str: "Open Water\x00\x00", expected: "Open Water"},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			res := trimUTF8NullTerminatedString([]byte(tc.str))
			if string(res) != tc.expected {
				t.Fatalf("expected: %s, got: %s", tc.expected, res)
			}
		})
	}

}
