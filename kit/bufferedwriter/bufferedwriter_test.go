// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufferedwriter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestNew(t *testing.T) {
	bw := New(nil)
	_, ok := bw.(*bufio.Writer)
	if !ok {
		t.Fatalf("expected *bufio.Writer, got: %T", bw)
	}
}

type writerAt struct {
	io.Writer
	io.WriterAt
}
type writeSeeker struct {
	io.Writer
	io.Seeker
}

func TestNewSize(t *testing.T) {
	tt := []struct {
		w            io.Writer
		size         int
		expectedSize int
	}{
		{w: (io.Writer)(nil), size: 0, expectedSize: 4 << 10},
		{w: (io.Writer)(nil), size: 1 << 10, expectedSize: 1 << 10},
		{w: (*writerAt)(nil), size: 8 << 10, expectedSize: 8 << 10}, // 8 KB
		{w: (*writeSeeker)(nil), size: 8 << 10, expectedSize: 8 << 10},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %T", i, tc.w), func(t *testing.T) {
			bw := NewSize(tc.w, tc.size)
			var size int
			switch w := bw.(type) {
			case *bufio.Writer:
				size = w.Size()
			case *WriterAt:
				size = w.Size()
			case *WriteSeeker:
				size = w.Size()
			}
			if size != tc.expectedSize {
				t.Fatalf("expected size: %d, got: %d", tc.expectedSize, size)
			}
		})
	}
}

type fnWriter func(b []byte) (n int, err error)

func (f fnWriter) Write(b []byte) (n int, err error) { return f(b) }

type fnWriterAt func(p []byte, offset int64) (n int, err error)

func (f fnWriterAt) WriteAt(p []byte, offset int64) (n int, err error) { return f(p, offset) }

type fnSeeker func(offset int64, whence int) (n int64, err error)

func (f fnSeeker) Seek(offset int64, whence int) (n int64, err error) { return f(offset, whence) }

var (
	fnWriteOK    = fnWriter(func(b []byte) (n int, err error) { return len(b), nil })
	fnWriteAtOK  = fnWriterAt(func(p []byte, offset int64) (n int, err error) { return len(p), nil })
	fnSeekerOK   = fnSeeker(func(offset int64, whence int) (n int64, err error) { return offset, nil })
	fnWriteErr   = fnWriter(func(b []byte) (n int, err error) { return 0, io.EOF })
	fnWriteAtErr = fnWriterAt(func(p []byte, offset int64) (n int, err error) { return 0, io.EOF })
	fnSeekerErr  = fnSeeker(func(offset int64, whence int) (n int64, err error) { return 0, io.ErrUnexpectedEOF })
)

func TestWriterAt(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
		errs []error
	}{
		{
			name: "writerAt happy path",
			w: writerAt{
				Writer:   fnWriteOK,
				WriterAt: fnWriteAtOK,
			},
			errs: []error{nil, nil, nil},
		},
		{
			name: "writerAt writeAt error",
			w: func() io.Writer {
				return writerAt{
					Writer:   fnWriteOK,
					WriterAt: fnWriteAtErr,
				}
			}(),
			errs: []error{nil, io.EOF, nil},
		},
		{
			name: "writerAt flush error",
			w: func() io.Writer {
				return writerAt{
					Writer:   fnWriteErr,
					WriterAt: fnWriteAtOK,
				}
			}(),
			errs: []error{nil, io.EOF, io.EOF},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			wa := New(tc.w).(*WriterAt)
			_, err := wa.Write([]byte{1, 2, 3, 4})
			if !errors.Is(err, tc.errs[0]) {
				t.Fatalf("expected write err: %v, got: %v", tc.errs[0], err)
			}
			_, err = wa.WriteAt([]byte{0, 0}, 0)
			if !errors.Is(err, tc.errs[1]) {
				t.Fatalf("expected writeAt err: %v, got: %v", tc.errs[1], err)
			}
			err = wa.Flush()
			if !errors.Is(err, tc.errs[2]) {
				t.Fatalf("expected flush err: %v, got: %v", tc.errs[2], err)
			}
		})
	}
}

func TestWriteSeeker(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
		errs []error
	}{
		{
			name: "writeSeeker happy path",
			w: writeSeeker{
				Writer: fnWriteOK,
				Seeker: fnSeekerOK,
			},
			errs: []error{nil, nil, nil},
		},
		{
			name: "writeSeeker seek error",
			w: writeSeeker{
				Writer: fnWriteOK,
				Seeker: fnSeekerErr,
			},
			errs: []error{nil, io.ErrUnexpectedEOF, nil},
		},
		{
			name: "writeSeeker flush error",
			w: writeSeeker{
				Writer: fnWriteErr,
				Seeker: fnSeekerOK,
			},
			errs: []error{nil, io.EOF, io.EOF},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			ws := New(tc.w).(*WriteSeeker)
			_, err := ws.Write([]byte{1, 2, 3, 4})
			if !errors.Is(err, tc.errs[0]) {
				t.Fatalf("expected write err: %v, got: %v", tc.errs[0], err)
			}
			_, err = ws.Seek(0, io.SeekStart)
			if !errors.Is(err, tc.errs[1]) {
				t.Fatalf("expected seek err: %v, got: %v", tc.errs[1], err)
			}
			err = ws.Flush()
			if !errors.Is(err, tc.errs[2]) {
				t.Fatalf("expected flush err: %v, got: %v", tc.errs[2], err)
			}
		})
	}
}

type mockWriterAt struct {
	io.Writer
	io.WriterAt
}

type mockWriteSeeker struct {
	io.Writer
	io.Seeker
}

func TestTypeAssertion(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
		ok   bool
	}{
		{name: "writerAt", w: New(mockWriterAt{fnWriteOK, fnWriteAtOK}), ok: true},
		{name: "writeSeeker", w: New(mockWriteSeeker{fnWriteOK, fnSeekerOK}), ok: true},
		{name: "writer only", w: New(fnWriteOK), ok: true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			switch wr := tc.w.(type) {
			case io.WriterAt, io.WriteSeeker:
				if tc.ok != true {
					t.Fatalf("w: %T: expected: %v, got: %v", wr, tc.ok, true)
				}
			case *bufio.Writer:
				if tc.ok != true {
					t.Fatalf("w: %T: expected: %v, got: %v", wr, tc.ok, true)
				}
			default:
				if tc.ok != false {
					t.Fatalf("w: %T: expected: %v, got: %v", tc.w, tc.ok, false)
				}
			}
		})
	}
}
