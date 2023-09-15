// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type DefinedBool bool
type DefinedInt8 int8
type DefinedUint8 uint8
type DefinedInt16 int16
type DefinedUint16 uint16
type DefinedInt32 int32
type DefinedUint32 uint32
type DefinedInt64 int64
type DefinedUint64 uint64
type DefinedFloat32 float32
type DefinedFloat64 float64
type DefinedString string
type DefinedAny any

type privateDefinedFloat64 float64

func TestMarshal(t *testing.T) {
	var i16 int16 = 10
	var nilptri16 *int16

	tt := []struct {
		value any
		err   error
	}{
		// // primitive-types
		{value: false},
		{value: true},
		{value: int8(-19)},
		{value: uint8(129)},
		{value: int16(1429)},
		{value: int16(-429)},
		{value: uint16(9929)},
		{value: int32(819293429)},
		{value: int32(-8979123)},
		{value: uint32(9929)},
		{value: int64(819293429)},
		{value: int64(-8979123)},
		{value: uint64(9929)},
		{value: float32(819293429.192321)},
		{value: float32(-8979123.546734)},
		{value: float64(8192934298908979.192321)},
		{value: float64(-897912398989898.546734)},
		{value: "Fit SDK"},
		{value: ""},
		{value: []bool{true, false}},
		{value: []byte{1, 2}},
		{value: []uint8{1, 2}},
		{value: []int8{-19}},
		{value: []uint8{129}},
		{value: []int16{1429}},
		{value: []int16{-429}},
		{value: []uint16{9929}},
		{value: []int32{819293429}},
		{value: []int32{-8979123}},
		{value: []uint32{9929}},
		{value: []string{"not supported"}, err: ErrTypeNotSupported},
		{value: []int64{819293429}},
		{value: []int64{-8979123}},
		{value: []uint64{9929}},
		{value: []float32{819293429.192321}},
		{value: []float32{-8979123.546734}},
		{value: []float64{8192934298908979.192321}},
		{value: []float64{-897912398989898.546734}},
		{value: []any{-897912398989898.546734}, err: ErrTypeNotSupported},

		// Defined Types
		{value: []DefinedInt8{DefinedInt8(1), DefinedInt8(2)}},

		// Types genenerated by fitgen:
		{value: File(1)},
		{value: MesgNum(29)},
		{value: Checksum(1)},
		{value: FileFlags(1)},

		// User defined type marshaled using reflection:
		{value: DefinedBool(true)},
		{value: DefinedInt8(123)},
		{value: DefinedUint8(123)},
		{value: DefinedInt16(123)},
		{value: DefinedUint16(123)},
		{value: DefinedInt32(123)},
		{value: DefinedUint32(123)},
		{value: DefinedInt64(123)},
		{value: DefinedUint64(123)},
		{value: DefinedFloat32(123)},
		{value: DefinedFloat64(123)},
		{value: DefinedString("Fit SDK")},
		{value: DefinedString("")},

		// Unsupported types:
		{value: nil, err: ErrNilValue},
		{value: complex128(1), err: ErrTypeNotSupported},
		{value: nilptri16, err: ErrTypeNotSupported},
		{value: &i16},
		{value: []int16{10, 10}},
		{value: []*int16{&i16}},
		{value: []*int16{nilptri16}, err: ErrTypeNotSupported},
		{value: []*int16{nil}, err: ErrTypeNotSupported},
		{value: []DefinedAny{int16(129)}},
		{value: []DefinedAny{func() {}}, err: ErrTypeNotSupported},
		{value: []DefinedAny{DefinedAny(nil)}, err: ErrTypeNotSupported},
		{value: []DefinedAny{[]DefinedAny{}}, err: ErrTypeNotSupported},
		{value: []privateDefinedFloat64{8.234}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v))", tc.value, tc.value), func(t *testing.T) {
			b, err := Marshal(tc.value, binary.LittleEndian)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}

			buf := new(bytes.Buffer)
			if err := marshalWithReflectionForTest(buf, tc.value); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(b, buf.Bytes()); diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

// using little-endian
func marshalWithReflectionForTest(w io.Writer, value any) error {
	if value == nil {
		return fmt.Errorf("can't interface '%T': %w", value, ErrTypeNotSupported)
	}
	rv := reflect.Indirect(reflect.ValueOf(value))
	if !rv.IsValid() {
		return fmt.Errorf("can't interface '%T': %w", value, ErrTypeNotSupported)
	}

	if rv.Type().Kind() == reflect.Slice || rv.Type().Kind() == reflect.String {
		for i := 0; i < rv.Len(); i++ {
			if !rv.Index(i).CanInterface() || rv.Index(i).Interface() == nil {
				continue
			}
			if err := marshalWithReflectionForTest(w, rv.Index(i).Interface()); err != nil {
				return err
			}
		}
		if rv.Type().Kind() == reflect.String {
			w.Write([]byte{0x00}) // utf-8 null terminator
		}
		return nil
	}

	return binary.Write(w, binary.LittleEndian, rv.Interface())
}
