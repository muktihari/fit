// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basetype_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
)

func TestFromStringAndString(t *testing.T) {
	tt := []struct {
		s        string
		baseType basetype.BaseType
	}{
		{s: "enum", baseType: basetype.Enum},
		{s: "sint8", baseType: basetype.Sint8},
		{s: "uint8", baseType: basetype.Uint8},
		{s: "sint16", baseType: basetype.Sint16},
		{s: "uint16", baseType: basetype.Uint16},
		{s: "sint32", baseType: basetype.Sint32},
		{s: "uint32", baseType: basetype.Uint32},
		{s: "string", baseType: basetype.String},
		{s: "float32", baseType: basetype.Float32},
		{s: "float64", baseType: basetype.Float64},
		{s: "uint8z", baseType: basetype.Uint8z},
		{s: "uint16z", baseType: basetype.Uint16z},
		{s: "uint32z", baseType: basetype.Uint32z},
		{s: "byte", baseType: basetype.Byte},
		{s: "sint64", baseType: basetype.Sint64},
		{s: "uint64", baseType: basetype.Uint64},
		{s: "uint64z", baseType: basetype.Uint64z},
		{s: "invalid", baseType: basetype.BaseType(255)},
	}

	t.Run("FromString", func(t *testing.T) {
		for _, tc := range tt {
			t.Run(tc.s, func(t *testing.T) {
				baseType := basetype.FromString(tc.s)
				if baseType != tc.baseType {
					t.Fatalf("expected: %d, got: %d", tc.baseType, baseType)
				}
			})
		}
	})

	t.Run("String", func(t *testing.T) {
		for _, tc := range tt {
			t.Run(tc.s, func(t *testing.T) {
				s := tc.baseType.String()
				if s != tc.s {
					t.Fatalf("expected: %s, got: %s", tc.s, s)
				}
			})
		}
	})
}

func TestList(t *testing.T) {
	l := basetype.List()
	expected := []basetype.BaseType{
		basetype.Enum,
		basetype.Sint8,
		basetype.Uint8,
		basetype.Sint16,
		basetype.Uint16,
		basetype.Sint32,
		basetype.Uint32,
		basetype.String,
		basetype.Float32,
		basetype.Float64,
		basetype.Uint8z,
		basetype.Uint16z,
		basetype.Uint32z,
		basetype.Byte,
		basetype.Sint64,
		basetype.Uint64,
		basetype.Uint64z,
	}

	if diff := cmp.Diff(l, expected); diff != "" {
		t.Fatal(diff)
	}
}

func TestSize(t *testing.T) {
	tt := []struct {
		baseType basetype.BaseType
		size     byte
	}{
		{baseType: basetype.Enum, size: 1},
		{baseType: basetype.Byte, size: 1},
		{baseType: basetype.Sint8, size: 1},
		{baseType: basetype.Uint8, size: 1},
		{baseType: basetype.Uint8z, size: 1},
		{baseType: basetype.String, size: 1},
		{baseType: basetype.Sint16, size: 2},
		{baseType: basetype.Uint16, size: 2},
		{baseType: basetype.Uint16z, size: 2},
		{baseType: basetype.Sint32, size: 4},
		{baseType: basetype.Uint32, size: 4},
		{baseType: basetype.Uint32z, size: 4},
		{baseType: basetype.Float32, size: 4},
		{baseType: basetype.Sint64, size: 8},
		{baseType: basetype.Uint64, size: 8},
		{baseType: basetype.Uint64z, size: 8},
		{baseType: basetype.Float64, size: 8},
		{baseType: 255, size: 0},
	}
	for _, tc := range tt {
		t.Run(tc.baseType.String(), func(t *testing.T) {
			size := tc.baseType.Size()
			if size != tc.size {
				t.Fatalf("expected: %d, got: %d", tc.size, size)
			}
		})
	}
}

func TestGoType(t *testing.T) {
	tt := []struct {
		baseType basetype.BaseType
		goType   string
	}{
		{baseType: basetype.Enum, goType: "byte"},
		{baseType: basetype.Byte, goType: "byte"},
		{baseType: basetype.Sint8, goType: "int8"},
		{baseType: basetype.Uint8, goType: "uint8"},
		{baseType: basetype.Uint8z, goType: "uint8"},
		{baseType: basetype.String, goType: "string"},
		{baseType: basetype.Sint16, goType: "int16"},
		{baseType: basetype.Uint16, goType: "uint16"},
		{baseType: basetype.Uint16z, goType: "uint16"},
		{baseType: basetype.Sint32, goType: "int32"},
		{baseType: basetype.Uint32, goType: "uint32"},
		{baseType: basetype.Uint32z, goType: "uint32"},
		{baseType: basetype.Float32, goType: "float32"},
		{baseType: basetype.Sint64, goType: "int64"},
		{baseType: basetype.Uint64, goType: "uint64"},
		{baseType: basetype.Uint64z, goType: "uint64"},
		{baseType: basetype.Float64, goType: "float64"},
		{baseType: 255, goType: "invalid"},
	}
	for _, tc := range tt {
		t.Run(tc.baseType.String(), func(t *testing.T) {
			size := tc.baseType.GoType()
			if size != tc.goType {
				t.Fatalf("expected: %s, got: %s", tc.goType, size)
			}
		})
	}
}

func TestEndianAbility(t *testing.T) {
	tt := []struct {
		baseType  basetype.BaseType
		endianess byte
	}{
		{baseType: basetype.Sint8, endianess: 0},
		{baseType: basetype.Enum, endianess: 0},
		{baseType: basetype.Byte, endianess: 0},
		{baseType: basetype.Uint8, endianess: 0},
		{baseType: basetype.Uint8z, endianess: 0},
		{baseType: basetype.String, endianess: 0},
		{baseType: basetype.Sint16, endianess: 1},
		{baseType: basetype.Uint16, endianess: 1},
		{baseType: basetype.Uint16z, endianess: 1},
		{baseType: basetype.Sint32, endianess: 1},
		{baseType: basetype.Uint32, endianess: 1},
		{baseType: basetype.Uint32z, endianess: 1},
		{baseType: basetype.Float32, endianess: 1},
		{baseType: basetype.Sint64, endianess: 1},
		{baseType: basetype.Uint64, endianess: 1},
		{baseType: basetype.Uint64z, endianess: 1},
		{baseType: basetype.Float64, endianess: 1},
		{baseType: 255, endianess: 0},
	}
	for _, tc := range tt {
		t.Run(tc.baseType.String(), func(t *testing.T) {
			endianess := tc.baseType.EndianAbility()
			if endianess != tc.endianess {
				t.Fatalf("expected: %d, got: %d", tc.endianess, endianess)
			}
		})
	}
}

func TestIsInteger(t *testing.T) {
	tt := []struct {
		baseType  basetype.BaseType
		isInteger bool
	}{
		{baseType: basetype.Sint8, isInteger: true},
		{baseType: basetype.Enum, isInteger: true},
		{baseType: basetype.Byte, isInteger: true},
		{baseType: basetype.Uint8, isInteger: true},
		{baseType: basetype.Uint8z, isInteger: true},
		{baseType: basetype.String, isInteger: false},
		{baseType: basetype.Sint16, isInteger: true},
		{baseType: basetype.Uint16, isInteger: true},
		{baseType: basetype.Uint16z, isInteger: true},
		{baseType: basetype.Sint32, isInteger: true},
		{baseType: basetype.Uint32, isInteger: true},
		{baseType: basetype.Uint32z, isInteger: true},
		{baseType: basetype.Float32, isInteger: false},
		{baseType: basetype.Sint64, isInteger: true},
		{baseType: basetype.Uint64, isInteger: true},
		{baseType: basetype.Uint64z, isInteger: true},
		{baseType: basetype.Float64, isInteger: false},
		{baseType: 255, isInteger: false},
	}
	for _, tc := range tt {
		t.Run(tc.baseType.String(), func(t *testing.T) {
			isInteger := tc.baseType.IsInteger()
			if isInteger != tc.isInteger {
				t.Fatalf("expected: %t, got: %t", tc.isInteger, isInteger)
			}
		})
	}
}

func TestInvalid(t *testing.T) {
	tt := []struct {
		baseType basetype.BaseType
		invalid  any
	}{
		{baseType: basetype.Sint8, invalid: basetype.Sint8Invalid},
		{baseType: basetype.Enum, invalid: basetype.EnumInvalid},
		{baseType: basetype.Byte, invalid: basetype.ByteInvalid},
		{baseType: basetype.Uint8, invalid: basetype.Uint8Invalid},
		{baseType: basetype.Uint8z, invalid: basetype.Uint8zInvalid},
		{baseType: basetype.String, invalid: basetype.StringInvalid},
		{baseType: basetype.Sint16, invalid: basetype.Sint16Invalid},
		{baseType: basetype.Uint16, invalid: basetype.Uint16Invalid},
		{baseType: basetype.Uint16z, invalid: basetype.Uint16zInvalid},
		{baseType: basetype.Sint32, invalid: basetype.Sint32Invalid},
		{baseType: basetype.Uint32, invalid: basetype.Uint32Invalid},
		{baseType: basetype.Uint32z, invalid: basetype.Uint32zInvalid},
		{baseType: basetype.Float32, invalid: math.Float32frombits(basetype.Float32Invalid)},
		{baseType: basetype.Sint64, invalid: basetype.Sint64Invalid},
		{baseType: basetype.Uint64, invalid: basetype.Uint64Invalid},
		{baseType: basetype.Uint64z, invalid: basetype.Uint64zInvalid},
		{baseType: basetype.Float64, invalid: math.Float64frombits(basetype.Float64Invalid)},
		{baseType: 255, invalid: "invalid"},
	}
	for _, tc := range tt {
		t.Run(tc.baseType.String(), func(t *testing.T) {
			invalid := tc.baseType.Invalid()
			switch tc.baseType {
			case basetype.Float32:
				f32 := invalid.(float32)
				if u32 := math.Float32bits(f32); u32 != basetype.Uint32Invalid {
					t.Fatalf("expected: %d, got: %d", basetype.Uint32Invalid, u32)
				}
			case basetype.Float64:
				f64 := invalid.(float64)
				if u64 := math.Float64bits(f64); u64 != basetype.Uint64Invalid {
					t.Fatalf("expected: %d, got: %d", basetype.Uint64Invalid, u64)
				}
			default:
				if invalid != tc.invalid {
					t.Fatalf("expected: %t, got: %t", tc.invalid, invalid)
				}
			}
		})
	}
}
