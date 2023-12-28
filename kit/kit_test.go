// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kit_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit"
)

func TestPtr(t *testing.T) {
	val := float64(10)
	ptr := kit.Ptr(val)

	if diff := cmp.Diff(val, *ptr); diff != "" {
		t.Fatal(diff)
	}
}
