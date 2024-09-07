// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reducer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/carto/rdp"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	cd                = filepath.Dir(filename)
	testdata          = filepath.Join(cd, "..", "..", "..", "testdata")
	fromGarminForums  = filepath.Join(testdata, "from_garmin_forums")
)

func TestReduceRDP(t *testing.T) {
	f, err := os.Open(filepath.Join(fromGarminForums, "triathlon_summary_last.fit"))
	if err != nil {
		t.Fatalf("could not open file: %v", err)
	}
	defer f.Close()

	dec := decoder.New(f, decoder.WithNoComponentExpansion())
	fit, err := dec.Decode()
	if err != nil {
		t.Fatalf("could not decode: %v", err)
	}

	prev := float64(len(fit.Messages))
	if err := Reduce(fit, WithRDP(0.001)); err != nil {
		t.Fatalf("could not reduce: %v", err)
	}
	current := float64(len(fit.Messages))

	// NOTE: Creating proper test for this requires dedication; this can be done later.
	// For now, let's just have this guard that will fail our test when we make incorrect implementation.
	// This threshold is just an estimation, at least we don't get lower value than this.
	// If we do, then implementation may be incorrect.
	const threshold = 80.0
	if percent := current / prev * 100; percent < threshold || percent > 100 {
		t.Fatalf("expected size after reduce >%g%% && <= 100%%, got: %g%%", threshold, percent)
	}
}

func TestFindFragment(t *testing.T) {
	tt := []struct {
		name      string
		records   []int
		points    []rdp.Point
		fragments []int
	}{
		{
			name:      "valid middle",
			records:   []int{1, 2, 3, 4, 5, 6},
			points:    []rdp.Point{{Index: 2}, {Index: 4}, {Index: 5}},
			fragments: []int{1, 3, 6},
		},
		{
			name:      "valid middle-end",
			records:   []int{1, 2, 3, 4, 5, 6},
			points:    []rdp.Point{{Index: 2}, {Index: 3}, {Index: 6}},
			fragments: []int{1, 4, 5},
		},
		{
			name:      "valid begin-middle",
			records:   []int{1, 2, 3, 4, 5, 6},
			points:    []rdp.Point{{Index: 1}, {Index: 3}, {Index: 5}},
			fragments: []int{2, 4, 6},
		},
		{
			name:      "valid begin-middle-end",
			records:   []int{2, 3, 4, 5, 6},
			points:    []rdp.Point{{Index: 2}, {Index: 3}, {Index: 6}},
			fragments: []int{4, 5},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			fragments := findFragments(tc.records, tc.points)
			if diff := cmp.Diff(fragments, tc.fragments); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDefragment(t *testing.T) {
	makeSeq := func(n uint16) []proto.Message {
		ms := make([]proto.Message, n)
		for i := uint16(0); i < n; i++ {
			ms[i].Num = typedef.MesgNum(i)
		}
		return ms
	}

	tt := []struct {
		name      string
		mesgs     []proto.Message
		fragments []int
		expect    []proto.Message
	}{
		{
			name:      "",
			mesgs:     makeSeq(5), // 0, 1, 2, 3, 4
			fragments: []int{0, 2},
			expect:    []proto.Message{{Num: 1}, {Num: 3}, {Num: 4}},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			mesgs := defragment(tc.mesgs, tc.fragments)
			if diff := cmp.Diff(mesgs, tc.expect,
				cmp.Transformer("printAsInteger", func(v typedef.MesgNum) uint16 { return uint16(v) }),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
