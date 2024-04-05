// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc16_test

import (
	"testing"

	"github.com/muktihari/fit/kit/hash/crc16"
)

func TestCRC16(t *testing.T) {
	b := []byte{14, 32, 84, 8, 214, 204, 9, 0, 46, 70, 73, 84} // example from some fit header.
	crc := uint16(12856)                                       // should match this checksum.

	c16 := crc16.New(crc16.MakeFitTable())

	if c16.BlockSize() != 1 {
		t.Fatalf("blocksize mismatch")
	}

	if c16.Size() != 2 {
		t.Fatalf("size mismatch")
	}

	_, _ = c16.Write(b)
	if c16.Sum16() != crc {
		t.Fatalf("expected: %v, got: %v", crc, c16.Sum16())
	}

	sum := c16.Sum([]byte{10})
	expect := []byte{10, 50, 56}
	for i := range sum {
		if sum[i] != expect[i] {
			t.Fatalf("expected: %v, got: %v", expect[i], sum[i])
		}
	}

	c16.Reset()

	if c16.Sum16() != 0 {
		t.Fatalf("expected 0 after reset, got: %v", c16.Sum16())
	}
}
