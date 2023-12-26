package decoder

import (
	"fmt"
	"testing"

	"github.com/muktihari/fit/profile/basetype"
)

func TestBitsFromValue(t *testing.T) {
	tt := []struct {
		value    any
		expected uint32
		ok       bool
	}{
		{value: int8(10), expected: 10, ok: true},
		{value: uint8(10), expected: 10, ok: true},
		{value: int16(10), expected: 10, ok: true},
		{value: uint16(10), expected: 10, ok: true},
		{value: int32(10), expected: 10, ok: true},
		{value: uint32(10), expected: 10, ok: true},
		{value: int64(10), expected: 10, ok: true},
		{value: uint64(10), expected: 10, ok: true},
		{value: float32(10), expected: 10, ok: true},
		{value: float64(10), expected: 10, ok: true},
		{value: []byte{1, 1, 1}, expected: 1<<0 | 1<<8 | 1<<16, ok: true},
		{value: []byte{1, 255, 1}, expected: 0, ok: false},
		{value: make([]byte, 33), expected: 0, ok: false},
		{value: "string value", expected: 0, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v (%T)", tc.value, tc.value), func(t *testing.T) {
			res, ok := bitsFromValue(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected ok: %t, got: %t", tc.ok, ok)
			}
			if res != tc.expected {
				t.Fatalf("expected: %d, got: %d", tc.expected, res)
			}
		})
	}
}

func TestValueFromBits(t *testing.T) {
	tt := []struct {
		sbits    uint32
		basetype basetype.BaseType
		value    any
		ok       bool
	}{
		{sbits: 10, basetype: basetype.Sint8, value: int8(10), ok: true},
		{sbits: 10, basetype: basetype.Uint8, value: uint8(10), ok: true},
		{sbits: 10, basetype: basetype.Sint16, value: int16(10), ok: true},
		{sbits: 10, basetype: basetype.Uint16, value: uint16(10), ok: true},
		{sbits: 10, basetype: basetype.Sint32, value: int32(10), ok: true},
		{sbits: 10, basetype: basetype.Uint32, value: uint32(10), ok: true},
		{sbits: 10, basetype: basetype.Sint64, value: int64(10), ok: true},
		{sbits: 10, basetype: basetype.Uint64, value: uint64(10), ok: true},
		{sbits: 10, basetype: basetype.Float32, value: float32(10), ok: true},
		{sbits: 10, basetype: basetype.Float64, value: float64(10), ok: true},
		{sbits: 10, basetype: basetype.String, value: basetype.StringInvalid, ok: false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s %v (%T)", tc.basetype, tc.value, tc.value), func(t *testing.T) {
			res := valueFromBits(tc.sbits, tc.basetype)
			if res != tc.value {
				t.Fatalf("expected: %v, got: %v", tc.value, res)
			}
		})
	}
}
