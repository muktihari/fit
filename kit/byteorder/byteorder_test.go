// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package byteorder_test

import (
	"encoding/binary"
	"testing"

	"github.com/muktihari/fit/kit/byteorder"
)

func TestSelect(t *testing.T) {
	if byteorder.Select(0) != binary.LittleEndian {
		t.Fatalf("expected little endian")
	}
	if byteorder.Select(1) != binary.BigEndian {
		t.Fatalf("expected big endian")
	}
}
