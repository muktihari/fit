// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadBufferReset(t *testing.T) {
	tt := []struct {
		name         string
		r            io.Reader
		size         int
		expectedSize int
	}{
		{
			name:         "default",
			r:            fnReaderOK,
			size:         4096,
			expectedSize: 4096,
		},
		{
			name:         "less than minimal size",
			r:            fnReaderOK,
			size:         reservedbuf - 1,
			expectedSize: reservedbuf,
		},
		{
			name:         "more than maximum size",
			r:            fnReaderOK,
			size:         maxReadBufferSize + 1,
			expectedSize: maxReadBufferSize,
		},
		{
			name:         "nil buffer",
			r:            nil,
			size:         4096,
			expectedSize: 4096,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			b := &readBuffer{
				r:    fnReader(func(b []byte) (n int, err error) { return n, err }),
				cur:  123,
				last: 456,
			}
			b.Reset(tc.r, tc.size)

			if diff := cmp.Diff(b.r, tc.r, cmp.Transformer("r", func(r io.Reader) uintptr {
				return reflect.ValueOf(r).Pointer()
			})); diff != "" {
				t.Fatal(diff)
			}

			if b.cur != 0 && b.last != 0 {
				t.Fatalf("cur and last should be zero after reset, got: %d and %d", b.cur, b.last)
			}

			size := len(b.buf) - reservedbuf
			if size != tc.expectedSize {
				t.Fatalf("expected size: %d, got: %d", tc.expectedSize, size)
			}
		})
	}
}

func TestReadBufferReadN(t *testing.T) {
	tt := []struct {
		name   string
		size   int
		r      io.Reader
		testFn func(buf *readBuffer) error
	}{
		{
			name: "reader has exactly 4096 bytes",
			size: defaultReadBufferSize,
			r: func() io.Reader {
				buf := make([]byte, 4096)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					n = copy(b, buf[cur:])
					cur += n
					return
				})
			}(),
			testFn: func(buf *readBuffer) error {
				_, err := buf.ReadN(96)
				if err != nil {
					return err
				}
				_, err = buf.ReadN(4000)
				if err != nil {
					return err
				}
				_, err = buf.ReadN(1)
				if err == io.EOF {
					return nil
				}
				return err
			},
		},
		{
			name: "reader has 14_096 bytes",
			size: defaultReadBufferSize,
			r: func() io.Reader {
				buf := make([]byte, 14_096)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					n = copy(b, buf[cur:])
					cur += n
					return
				})
			}(),
			testFn: func(buf *readBuffer) error {
				_, err := buf.ReadN(96)
				if err != nil {
					return err
				}
				_, err = buf.ReadN(4000)
				if err != nil {
					return err
				}
				for i := 0; i < 10_000; i++ {
					_, err = buf.ReadN(1)
					if err != nil {
						return fmt.Errorf("[%d]: %w", i, err)
					}
				}
				_, err = buf.ReadN(1)
				if err == io.EOF {
					return nil
				}
				return err
			},
		},
		{
			name: "when remaining is zero, cur and last should same",
			size: defaultReadBufferSize,
			r: func() io.Reader {
				buf := make([]byte, 4096)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					n = copy(b, buf[cur:])
					cur += n
					return
				})
			}(),
			testFn: func(buf *readBuffer) error {
				_, err := buf.ReadN(4096)
				if !errors.Is(err, nil) {
					return fmt.Errorf("expected nil, got: %v", err)
				}

				// try read when remaining is now zero and reader should return EOF.
				_, err = buf.ReadN(1)
				if !errors.Is(err, io.EOF) {
					return fmt.Errorf("expected EOF, got: %v", err)
				}
				if buf.cur != buf.last {
					return fmt.Errorf("expected cur (%d) == last (%d), got: false",
						buf.cur, buf.last)
				}

				// reset r to simulate changing r or if r is an io.ReadSeeker that has been seeked to the beginning.
				buf.r = bytes.NewBuffer(make([]byte, 4096))
				_, err = buf.ReadN(4096)
				if !errors.Is(err, nil) {
					return fmt.Errorf("expected nil, got: %v", err)
				}
				_, err = buf.ReadN(1)
				if !errors.Is(err, io.EOF) {
					return fmt.Errorf("expected EOF, got: %v", err)
				}

				return nil
			},
		},
		{
			name: "test read 765 when buffer is set to minimum",
			size: 16, // 16 was the previous minBufSize, it now will be changed to 765
			r:    fnReaderOK,
			testFn: func(buf *readBuffer) error {
				_, err := buf.ReadN(reservedbuf)
				return err
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			b := new(readBuffer)
			b.Reset(tc.r, 4096)
			if err := tc.testFn(b); err != nil {
				t.Fatalf("expected nil, got: %v", err)
			}
		})
	}
}
