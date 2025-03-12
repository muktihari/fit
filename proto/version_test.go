// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/proto"
)

func TestCreateVersion(t *testing.T) {
	tt := []struct {
		major   byte
		minor   byte
		version proto.Version
	}{
		{
			major:   1,
			minor:   1,
			version: proto.Version((1 << 4) | 1),
		},
		{
			major:   2,
			minor:   1,
			version: proto.Version((2 << 4) | 1),
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d, %d", tc.major, tc.minor), func(t *testing.T) {
			v := proto.CreateVersion(tc.major, tc.minor)
			if diff := cmp.Diff(v, tc.version); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestVersionMajorMinor(t *testing.T) {
	major := proto.Version(32).Major()
	if major != 2 {
		t.Fatalf("expected major: 2, got: %d", major)
	}

	minor := proto.Version(7).Minor()
	if minor != 7 {
		t.Fatalf("expected minor: 7, got: %d", minor)
	}
}
