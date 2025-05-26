// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestAccumulatorCollect(t *testing.T) {
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

func TestAccumulatorCollectValue(t *testing.T) {
	tt := []struct {
		val               proto.Value
		accumulatedValues []value
	}{
		{val: proto.Int8(1), accumulatedValues: []value{{last: 1, value: 1}}},
		{val: proto.Uint8(2), accumulatedValues: []value{{last: 2, value: 2}}},
		{val: proto.Int16(3), accumulatedValues: []value{{last: 3, value: 3}}},
		{val: proto.Uint16(4), accumulatedValues: []value{{last: 4, value: 4}}},
		{val: proto.Int32(5), accumulatedValues: []value{{last: 5, value: 5}}},
		{val: proto.Uint32(6), accumulatedValues: []value{{last: 6, value: 6}}},
		{val: proto.Int64(7), accumulatedValues: []value{{last: 7, value: 7}}},
		{val: proto.Uint64(8), accumulatedValues: []value{{last: 8, value: 8}}},
		{val: proto.Float32(9), accumulatedValues: []value{{last: 9, value: 9}}},
		{val: proto.Float64(10), accumulatedValues: []value{{last: 10, value: 10}}},
		{val: proto.SliceInt8([]int8{1, 2}), accumulatedValues: []value{{last: 2, value: 2}}},
		{val: proto.SliceUint8([]uint8{2, 3}), accumulatedValues: []value{{last: 3, value: 3}}},
		{val: proto.SliceInt16([]int16{3, 4}), accumulatedValues: []value{{last: 4, value: 4}}},
		{val: proto.SliceUint16([]uint16{4, 5}), accumulatedValues: []value{{last: 5, value: 5}}},
		{val: proto.SliceInt32([]int32{5, 6}), accumulatedValues: []value{{last: 6, value: 6}}},
		{val: proto.SliceUint32([]uint32{6, 7}), accumulatedValues: []value{{last: 7, value: 7}}},
		{val: proto.SliceInt64([]int64{7, 8}), accumulatedValues: []value{{last: 8, value: 8}}},
		{val: proto.SliceUint64([]uint64{8, 9}), accumulatedValues: []value{{last: 9, value: 9}}},
		{val: proto.SliceFloat32([]float32{9, 10}), accumulatedValues: []value{{last: 10, value: 10}}},
		{val: proto.SliceFloat64([]float64{10, 11}), accumulatedValues: []value{{last: 11, value: 11}}},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %v", i, tc.val.Any()), func(t *testing.T) {
			accumu := NewAccumulator()
			accumu.CollectValue(0, 0, tc.val)
			if diff := cmp.Diff(accumu.values, tc.accumulatedValues,
				cmp.AllowUnexported(value{}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestAccumulatorReset(t *testing.T) {
	accumu := NewAccumulator()
	accumu.Collect(mesgnum.Record, fieldnum.RecordSpeed, 1000)

	if len(accumu.values) != 1 {
		t.Fatalf("expected AccumulatedValues is 1, got: %d", len(accumu.values))
	}

	accumu.Reset()

	if len(accumu.values) != 0 {
		t.Fatalf("expected AccumulatedValues is 0 after reset, got: %d", len(accumu.values))
	}
}
