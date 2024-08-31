// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bufio"
	"errors"
	"io"
	"reflect"
	"testing"
)

func TestNewWriteBuffer(t *testing.T) {
	tt := []struct {
		name     string
		size     int
		w        io.Writer
		expected io.Writer
	}{
		{name: "nil", size: -1, w: nil, expected: nil},
		{name: "writerAt", size: 4096, w: mockWriterAt{}, expected: &writerAt{}},
		{name: "writerAt", size: 4096, w: mockWriteSeeker{}, expected: &writeSeeker{}},
		{name: "writer", size: 4096, w: fnWriteOK, expected: &bufio.Writer{}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			wr := newWriteBuffer(tc.w, tc.size)
			if reflect.TypeOf(wr) != reflect.TypeOf(tc.expected) {
				t.Fatalf("expected type: %T, got: %T", tc.expected, wr)
			}
		})
	}
}

func TestWriterAt(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
		errs [2]error
	}{
		{
			name: "happy flow",
			w:    mockWriterAt{Writer: fnWriteOK, WriterAt: fnWriteAtOK},
			errs: [2]error{nil, nil},
		},
		{
			name: "flush returns error",
			w:    mockWriterAt{Writer: fnWriteErr, WriterAt: fnWriteAtOK},
			errs: [2]error{nil, io.EOF},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := newWriteBuffer(tc.w, 4096).(*writerAt)
			_, err := w.Write([]byte{0, 0})
			if !errors.Is(err, tc.errs[0]) {
				t.Fatalf("expected error: %v, got: %v", tc.errs[0], err)
			}
			_, err = w.WriteAt([]byte{0, 0}, 0)
			if !errors.Is(err, tc.errs[1]) {
				t.Fatalf("expected error: %v, got: %v", tc.errs[1], err)
			}
		})
	}
}

func TestWriteSeeker(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
		errs [2]error
	}{
		{
			name: "happy flow",
			w:    mockWriteSeeker{Writer: fnWriteOK, Seeker: fnSeekOK},
			errs: [2]error{nil, nil},
		},
		{
			name: "flush returns error",
			w:    mockWriteSeeker{Writer: fnWriteErr, Seeker: fnSeekOK},
			errs: [2]error{nil, io.EOF},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := newWriteBuffer(tc.w, 4096).(*writeSeeker)
			_, err := w.Write([]byte{0, 0})
			if !errors.Is(err, tc.errs[0]) {
				t.Fatalf("expected error: %v, got: %v", tc.errs[0], err)
			}
			_, err = w.Seek(0, io.SeekStart)
			if !errors.Is(err, tc.errs[1]) {
				t.Fatalf("expected error: %v, got: %v", tc.errs[1], err)
			}
		})
	}
}
