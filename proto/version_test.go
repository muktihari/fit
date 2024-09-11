// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"errors"
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
		ok      bool
	}{
		{
			major:   1,
			minor:   1,
			version: proto.Version((1 << 4) | 1),
			ok:      true,
		},
		{
			major:   2,
			minor:   1,
			version: proto.Version((1 << 4) | 1),
			ok:      false,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d, %d", tc.major, tc.minor), func(t *testing.T) {
			version, ok := proto.CreateVersion(tc.major, tc.minor)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
			if !ok {
				return
			}
			if diff := cmp.Diff(version, tc.version); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestValidateVersion(t *testing.T) {
	tt := []struct {
		version proto.Version
		err     error
	}{
		{version: 32, err: nil},
		{version: 64, err: proto.ErrProtocolVersionNotSupported},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d", tc.version), func(t *testing.T) {
			err := proto.Validate(tc.version)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}

		})
	}
}

func TestVersionMajorMinor(t *testing.T) {
	major := proto.VersionMajor(32)
	if major != 2 {
		t.Fatalf("expected major: 2, got: %d", major)
	}

	minor := proto.VersionMinor(7)
	if minor != 7 {
		t.Fatalf("expected minor: 7, got: %d", minor)
	}
}
