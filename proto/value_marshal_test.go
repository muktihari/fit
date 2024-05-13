// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

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

func TestValueMarshalAppend(t *testing.T) {
	tt := []struct {
		b     *[]byte
		value Value
		err   error
	}{
		{b: kit.Ptr([]byte{}), value: Bool(false)},
		{b: kit.Ptr([]byte{}), value: Bool(true)},
		{b: kit.Ptr([]byte{}), value: Int8(-19)},
		{b: kit.Ptr([]byte{}), value: Uint8(129)},
		{b: kit.Ptr([]byte{}), value: Int16(1429)},
		{b: kit.Ptr([]byte{}), value: Int16(-429)},
		{b: kit.Ptr([]byte{}), value: Uint16(9929)},
		{b: kit.Ptr([]byte{}), value: Int32(819293429)},
		{b: kit.Ptr([]byte{}), value: Int32(-8979123)},
		{b: kit.Ptr([]byte{}), value: Uint32(9929)},
		{b: kit.Ptr([]byte{}), value: Int64(819293429)},
		{b: kit.Ptr([]byte{}), value: Int64(-8979123)},
		{b: kit.Ptr([]byte{}), value: Uint64(9929)},
		{b: kit.Ptr([]byte{}), value: Float32(819293429.192321)},
		{b: kit.Ptr([]byte{}), value: Float32(-8979123.546734)},
		{b: kit.Ptr([]byte{}), value: Float64(8192934298908979.192321)},
		{b: kit.Ptr([]byte{}), value: Float64(-897912398989898.546734)},
		{b: kit.Ptr([]byte{}), value: String("FIT SDK")},
		{b: kit.Ptr([]byte{}), value: String("")},
		{b: kit.Ptr([]byte{}), value: SliceBool([]bool{true, false})},
		{b: kit.Ptr([]byte{}), value: SliceUint8([]byte{1, 2})},
		{b: kit.Ptr([]byte{}), value: SliceUint8([]uint8{1, 2})},
		{b: kit.Ptr([]byte{}), value: SliceInt8([]int8{-19})},
		{b: kit.Ptr([]byte{}), value: SliceUint8([]uint8{129})},
		{b: kit.Ptr([]byte{}), value: SliceInt16([]int16{1429})},
		{b: kit.Ptr([]byte{}), value: SliceInt16([]int16{-429})},
		{b: kit.Ptr([]byte{}), value: SliceUint16([]uint16{9929})},
		{b: kit.Ptr([]byte{}), value: SliceInt32([]int32{819293429})},
		{b: kit.Ptr([]byte{}), value: SliceInt32([]int32{-8979123})},
		{b: kit.Ptr([]byte{}), value: SliceUint32([]uint32{9929})},
		{b: kit.Ptr([]byte{}), value: SliceString([]string{"supported"})},
		{b: kit.Ptr([]byte{}), value: SliceString([]string{})},
		{b: kit.Ptr([]byte{}), value: SliceString([]string{""})},
		{b: kit.Ptr([]byte{}), value: SliceString([]string{"\x00"})},
		{b: kit.Ptr([]byte{}), value: SliceString([]string{"\x00", "\x00"})},
		{b: kit.Ptr([]byte{}), value: SliceString([]string{string([]byte{'\x00'})})},
		{b: kit.Ptr([]byte{}), value: SliceInt64([]int64{819293429})},
		{b: kit.Ptr([]byte{}), value: SliceInt64([]int64{-8979123})},
		{b: kit.Ptr([]byte{}), value: SliceUint64([]uint64{9929})},
		{b: kit.Ptr([]byte{}), value: SliceFloat32([]float32{819293429.192321})},
		{b: kit.Ptr([]byte{}), value: SliceFloat32([]float32{-8979123.546734})},
		{b: kit.Ptr([]byte{}), value: SliceFloat64([]float64{8192934298908979.192321})},
		{b: kit.Ptr([]byte{}), value: SliceFloat64([]float64{-897912398989898.546734})},
		{b: kit.Ptr([]byte{}), value: Value{}, err: ErrTypeNotSupported},
	}

	for i, tc := range tt {
		for arch := byte(0); arch <= 1; arch++ {
			t.Run(fmt.Sprintf("[%d] %T(%v))", i, tc.value.Any(), tc.value.Any()), func(t *testing.T) {
				arr := pool.Get().(*[MaxBytesPerMessage]byte)
				defer pool.Put(arr)
				b := arr[:0]

				var err error
				*tc.b, err = tc.value.MarshalAppend(b, arch)
				if !errors.Is(err, tc.err) {
					t.Fatalf("expected err: %v, got: %v", tc.err, err)
				}
				if err != nil {
					return
				}

				buf := new(bytes.Buffer)
				if err := marshalValueWithReflectionForTest(buf, tc.value, arch); err != nil {
					t.Fatalf("marshalWithReflectionForTest: %v", err)
				}

				if len(*tc.b) == 0 && len(buf.Bytes()) == 0 {
					return
				}

				if diff := cmp.Diff(*tc.b, buf.Bytes()); diff != "" {
					fmt.Printf("value: %v, b: %v, buf: %v\n", tc.value.Any(), *tc.b, buf.Bytes())
					t.Fatal(diff)
				}

			})
		}
	}
}

func marshalValueWithReflectionForTest(w io.Writer, value Value, arch byte) error {
	if value.Type() == TypeInvalid {
		return fmt.Errorf("can't interface '%T': %w", value, ErrTypeNotSupported)
	}
	rv := reflect.Indirect(reflect.ValueOf(value.Any()))
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
			iface := rv.Index(i).Interface()
			val := Any(iface)
			if err := marshalValueWithReflectionForTest(w, val, arch); err != nil {
				return err
			}
		}
		if rv.Len() == 0 && rv.Type() == reflect.TypeOf([]string{}) {
			w.Write([]byte{0})
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
		if arch == 0 {
			return binary.Write(w, binary.LittleEndian, b)
		} else {
			return binary.Write(w, binary.BigEndian, b)
		}
	}
	if arch == 0 {
		return binary.Write(w, binary.LittleEndian, rv.Interface())
	}
	return binary.Write(w, binary.BigEndian, rv.Interface())
}

func BenchmarkValueMarshalAppend(b *testing.B) {
	b.StopTimer()
	value := SliceUint16(make([]uint16, 256/2))
	buf := make([]byte, 0, 256)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = value.MarshalAppend(buf, littleEndian)
	}
}
