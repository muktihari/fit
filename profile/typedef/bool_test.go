// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"fmt"
	"testing"
)

func TestBoolString(t *testing.T) {
	tt := []struct {
		name     string
		value    Bool
		expected string
	}{
		{name: "true", value: BoolTrue, expected: "true"},
		{name: "true", value: BoolFalse, expected: "false"},
		{name: "invalid", value: BoolInvalid, expected: "BoolInvalid(255)"},
		{name: "2", value: Bool(2), expected: "BoolInvalid(2)"},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			v := tc.value.String()
			if v != tc.expected {
				t.Fatalf("expected: %q, got: %q", tc.expected, v)
			}
		})
	}
}

func TestBoolFromString(t *testing.T) {
	tt := []struct {
		name     string
		value    string
		expected Bool
	}{
		{name: "true", value: "true", expected: BoolTrue},
		{name: "false", value: "false", expected: BoolFalse},
		{name: "invalid", value: "invalid", expected: BoolInvalid},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			v := BoolFromString(tc.value)
			if v != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, v)
			}
		})
	}
}

func TestBoolFromBool(t *testing.T) {
	tt := []struct {
		name     string
		value    bool
		expected Bool
	}{
		{name: "true", value: true, expected: BoolTrue},
		{name: "false", value: false, expected: BoolFalse},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			v := BoolFromBool(tc.value)
			if v != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, v)
			}
		})
	}
}

func TestBoolUint8(t *testing.T) {
	v := BoolTrue
	if expected := uint8(v); expected != v.Uint8() {
		t.Fatalf("expected: %v, got: %v", expected, v.Uint8())
	}
}
