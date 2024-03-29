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
	"github.com/muktihari/fit/kit"
)

type test_bool bool
type test_int8 int8
type test_uint8 uint8
type test_int16 int16
type test_uint16 uint16
type test_int32 int32
type test_uint32 uint32
type test_int64 int64
type test_uint64 uint64
type test_float32 float32
type test_float64 float64
type test_string string
type test_any any

func TestMarshal(t *testing.T) {
	tt := []struct {
		value any
		err   error
	}{
		{value: float32(819293429.192321), err: nil},
		{value: nil, err: ErrNilValue},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v))", tc.value, tc.value), func(t *testing.T) {
			b, err := Marshal(tc.value, binary.LittleEndian)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err %s nil, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			buf := new(bytes.Buffer)
			err = marshalWithReflectionForTest(buf, tc.value)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(b, buf.Bytes()); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMarshalTo(t *testing.T) {
	var i16 int16 = 10
	var nilptri16 *int16

	tt := []struct {
		b     *[]byte
		value any
		err   error
	}{
		// // primitive-types
		{b: kit.Ptr([]byte{}), value: false},
		{b: kit.Ptr([]byte{}), value: true},
		{b: kit.Ptr([]byte{}), value: int8(-19)},
		{b: kit.Ptr([]byte{}), value: uint8(129)},
		{b: kit.Ptr([]byte{}), value: int16(1429)},
		{b: kit.Ptr([]byte{}), value: int16(-429)},
		{b: kit.Ptr([]byte{}), value: uint16(9929)},
		{b: kit.Ptr([]byte{}), value: int32(819293429)},
		{b: kit.Ptr([]byte{}), value: int32(-8979123)},
		{b: kit.Ptr([]byte{}), value: uint32(9929)},
		{b: kit.Ptr([]byte{}), value: int64(819293429)},
		{b: kit.Ptr([]byte{}), value: int64(-8979123)},
		{b: kit.Ptr([]byte{}), value: uint64(9929)},
		{b: kit.Ptr([]byte{}), value: float32(819293429.192321)},
		{b: kit.Ptr([]byte{}), value: float32(-8979123.546734)},
		{b: kit.Ptr([]byte{}), value: float64(8192934298908979.192321)},
		{b: kit.Ptr([]byte{}), value: float64(-897912398989898.546734)},
		{b: kit.Ptr([]byte{}), value: "Fit SDK"},
		{b: kit.Ptr([]byte{}), value: ""},
		{b: kit.Ptr([]byte{}), value: []bool{true, false}},
		{b: kit.Ptr([]byte{}), value: []byte{1, 2}},
		{b: kit.Ptr([]byte{}), value: []uint8{1, 2}},
		{b: kit.Ptr([]byte{}), value: []int8{-19}},
		{b: kit.Ptr([]byte{}), value: []uint8{129}},
		{b: kit.Ptr([]byte{}), value: []int16{1429}},
		{b: kit.Ptr([]byte{}), value: []int16{-429}},
		{b: kit.Ptr([]byte{}), value: []uint16{9929}},
		{b: kit.Ptr([]byte{}), value: []int32{819293429}},
		{b: kit.Ptr([]byte{}), value: []int32{-8979123}},
		{b: kit.Ptr([]byte{}), value: []uint32{9929}},
		{b: kit.Ptr([]byte{}), value: []string{"supported"}},
		{b: kit.Ptr([]byte{}), value: []string{""}},
		{b: kit.Ptr([]byte{}), value: []string{"\x00"}},
		{b: kit.Ptr([]byte{}), value: []string{"\x00", "\x00"}},
		{b: kit.Ptr([]byte{}), value: []string{string([]byte{'\x00'})}},
		{b: kit.Ptr([]byte{}), value: []int64{819293429}},
		{b: kit.Ptr([]byte{}), value: []int64{-8979123}},
		{b: kit.Ptr([]byte{}), value: []uint64{9929}},
		{b: kit.Ptr([]byte{}), value: []float32{819293429.192321}},
		{b: kit.Ptr([]byte{}), value: []float32{-8979123.546734}},
		{b: kit.Ptr([]byte{}), value: []float64{8192934298908979.192321}},
		{b: kit.Ptr([]byte{}), value: []float64{-897912398989898.546734}},
		{b: kit.Ptr([]byte{}), value: []any{-897912398989898.546734}, err: ErrTypeNotSupported},

		// test_ Types
		{b: kit.Ptr([]byte{}), value: []test_int8{test_int8(1), test_int8(2)}},

		// Types genenerated by fitgen:
		{b: kit.Ptr([]byte{}), value: File(1)},
		{b: kit.Ptr([]byte{}), value: MesgNum(29)},
		{b: kit.Ptr([]byte{}), value: Checksum(1)},
		{b: kit.Ptr([]byte{}), value: FileFlags(1)},

		// User test_ type marshaled using reflection:
		{b: kit.Ptr([]byte{}), value: test_bool(true)},
		{b: kit.Ptr([]byte{}), value: test_int8(123)},
		{b: kit.Ptr([]byte{}), value: test_uint8(123)},
		{b: kit.Ptr([]byte{}), value: test_int16(123)},
		{b: kit.Ptr([]byte{}), value: test_uint16(123)},
		{b: kit.Ptr([]byte{}), value: test_int32(123)},
		{b: kit.Ptr([]byte{}), value: test_uint32(123)},
		{b: kit.Ptr([]byte{}), value: test_int64(123)},
		{b: kit.Ptr([]byte{}), value: test_uint64(123)},
		{b: kit.Ptr([]byte{}), value: test_float32(123)},
		{b: kit.Ptr([]byte{}), value: test_float64(123)},
		{b: kit.Ptr([]byte{}), value: test_string("Fit SDK")},
		{b: kit.Ptr([]byte{}), value: test_string("")},

		// Unsupported types:
		{b: kit.Ptr([]byte{}), value: nil, err: ErrNilValue},
		{b: kit.Ptr([]byte{}), value: complex128(1), err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: nilptri16, err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: &i16},
		{b: kit.Ptr([]byte{}), value: []int16{10, 10}},
		{b: kit.Ptr([]byte{}), value: []*int16{&i16}},
		{b: kit.Ptr([]byte{}), value: []*int16{nilptri16}, err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: []*int16{nil}, err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: []test_any{int16(129)}},
		{b: kit.Ptr([]byte{}), value: []test_any{func() {}}, err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: []test_any{test_any(nil)}, err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: []test_any{[]test_any{}}, err: ErrTypeNotSupported},
		{b: kit.Ptr([]byte{}), value: []test_float64{8.234}},

		{b: nil, value: uint8(129), err: ErrNilDest},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%T(%v))", tc.value, tc.value), func(t *testing.T) {
			err := MarshalTo(tc.b, tc.value, binary.LittleEndian)
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

			if len(*tc.b) == 0 && len(buf.Bytes()) == 0 {
				return
			}

			if diff := cmp.Diff(*tc.b, buf.Bytes()); diff != "" {
				fmt.Printf("value: %v, b: %v, buf: %v\n", tc.value, *tc.b, buf.Bytes())
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

	if rv.Type().Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			if !rv.Index(i).CanInterface() || rv.Index(i).Interface() == nil {
				continue
			}
			if rv.Index(i).Kind() == reflect.String {
				value := rv.Index(i).String()
				if len(value) == 0 {
					rv.Index(i).SetString("\x00")
				}
			}
			if err := marshalWithReflectionForTest(w, rv.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil
	}

	if rv.Kind() == reflect.String {
		b := []byte(rv.String())
		if len(b) == 0 {
			b = []byte{0x0}
		} else if b[len(b)-1] != '\x00' {
			b = append([]byte(b), '\x00')
		}
		return binary.Write(w, binary.LittleEndian, b)
	}

	return binary.Write(w, binary.LittleEndian, rv.Interface())
}
