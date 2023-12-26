// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"testing"

	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
)

func TestCollect(t *testing.T) {
	type value struct {
		mesgNum      typedef.MesgNum
		destFieldNum byte
		value        uint32
		expected     uint32
	}

	tt := []struct {
		name          string
		collects      []value
		toAccumulates []value
	}{
		{
			name: "collect new values",
			collects: []value{
				{mesgNum: mesgnum.Record, destFieldNum: fieldnum.RecordAccumulatedPower, value: 5},
			},
			toAccumulates: []value{
				{
					mesgNum:      mesgnum.Record,
					destFieldNum: fieldnum.RecordAccumulatedPower,
					value:        11,
					expected:     11,
				},
				{
					mesgNum:      mesgnum.Record,
					destFieldNum: fieldnum.RecordTotalCycles,
					value:        70,
					expected:     70,
				},
			},
		},
		{
			name: "collect replace existing values",
			collects: []value{
				{mesgNum: mesgnum.Record, destFieldNum: fieldnum.RecordAccumulatedPower, value: 5},
				{mesgNum: mesgnum.Record, destFieldNum: fieldnum.RecordAccumulatedPower, value: 5},
			},
			toAccumulates: []value{
				{
					mesgNum:      mesgnum.Record,
					destFieldNum: fieldnum.RecordAccumulatedPower,
					value:        11,
					expected:     11,
				},
				{
					mesgNum:      mesgnum.Record,
					destFieldNum: fieldnum.RecordTotalCycles,
					value:        70,
					expected:     70,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			accumu := NewAccumulator()
			for i := range tc.collects {
				val := &tc.collects[i]
				accumu.Collect(val.mesgNum, val.destFieldNum, val.value)
			}
			for i := range tc.toAccumulates {
				val := &tc.toAccumulates[i]
				accumulatedValue := accumu.Accumulate(val.mesgNum, val.destFieldNum, val.value, 16)
				if accumulatedValue != tc.toAccumulates[i].expected {
					t.Fatalf("expected: %d, got: %d", tc.toAccumulates[i].expected, accumulatedValue)
				}
			}
		})
	}
}

func TestReset(t *testing.T) {
	accumu := NewAccumulator()
	accumu.Collect(mesgnum.Record, fieldnum.RecordSpeed, 1000)

	if len(accumu.AccumulatedValues) != 1 {
		t.Fatalf("expected AccumulatedValues is 1, got: %d", len(accumu.AccumulatedValues))
	}

	accumu.Reset()

	if len(accumu.AccumulatedValues) != 0 {
		t.Fatalf("expected AccumulatedValues is 0 after reset, got: %d", len(accumu.AccumulatedValues))
	}
}
