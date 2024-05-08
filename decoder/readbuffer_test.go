// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewReadBuffer(t *testing.T) {
	tt := []struct {
		name   string
		size   int
		lenBuf int
	}{
		{name: "default", size: defaultReadBufferSize, lenBuf: reservedbuf + defaultReadBufferSize},
		{name: "less than minReadBufferSize", size: 8, lenBuf: reservedbuf + minReadBufferSize},
		{name: "8192", size: 8192, lenBuf: reservedbuf + 8192},
		{name: "more than maxReadBufferSize", size: math.MaxInt, lenBuf: reservedbuf + maxReadBufferSize},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			b := newReadBuffer(nil, tc.size)
			if b.cur != reservedbuf {
				t.Fatalf("expected cur: %d, got: %d", reservedbuf, b.cur)
			}
			if b.last != reservedbuf {
				t.Fatalf("expected last: %d, got: %d", reservedbuf, b.last)
			}
			if len(b.buf) != tc.lenBuf {
				t.Fatalf("expected len(buf): %d, got: %d", tc.lenBuf, len(b.buf))
			}
		})
	}
}

func TestReadBufferReadN(t *testing.T) {
	tt := []struct {
		name   string
		r      io.Reader
		cur    int
		last   int
		testFn func(buf *readBuffer) error
	}{
		{
			name: "reader has exactly 4096 bytes",
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
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			b := newReadBuffer(tc.r, 4096)
			if err := tc.testFn(b); err != nil {
				t.Fatalf("expected nil, got: %v", err)
			}
		})
	}
}

func TestReadBufferResetAndSize(t *testing.T) {
	r := io.Reader(fnReaderOK)
	b := newReadBuffer(nil, 4096)

	b.Reset(r, 4096*2)
	if b.Size() != 4096*2 {
		t.Fatalf("expected: %d, got: %d", 4096*2, b.Size())
	}
	if diff := cmp.Diff(r, b.r, cmp.Transformer("r", func(r io.Reader) uintptr {
		return reflect.ValueOf(r).Pointer()
	})); diff != "" {
		t.Fatal(diff)
	}

	// Revert back
	b.Reset(nil, 4096)
	if diff := cmp.Diff(nil, b.r, cmp.Transformer("r", func(r io.Reader) uintptr {
		return reflect.ValueOf(r).Pointer()
	})); diff != "" {
		t.Fatal(diff)
	}
	if b.Size() != 4096 {
		t.Fatalf("expected: %d, got: %d", 4096, b.Size())
	}
}
