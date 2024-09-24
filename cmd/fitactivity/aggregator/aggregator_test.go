// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aggregator_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/cmd/fitactivity/aggregator"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

func BenchmarkAggregate(b *testing.B) {
	s1 := mesgdef.NewSession(nil).SetTotalCalories(10)
	s2 := mesgdef.NewSession(nil).SetTotalCalories(20)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		aggregator.Aggregate(s1, s2)
	}
}

func TestAggregate(t *testing.T) {
	tt := []struct {
		name string
		dst  any
		src  any
		exp  any
	}{
		{
			name: "Total",
			dst:  mesgdef.NewSession(nil).SetTotalCalories(10),
			src:  mesgdef.NewSession(nil).SetTotalCalories(15),
			exp:  mesgdef.NewSession(nil).SetTotalCalories(25),
		},
		{
			name: "Max",
			dst:  mesgdef.NewSession(nil).SetMaxAltitude(1000),
			src:  mesgdef.NewSession(nil).SetMaxAltitude(2000),
			exp:  mesgdef.NewSession(nil).SetMaxAltitude(2000),
		},
		{
			name: "EnhancedMax",
			dst:  mesgdef.NewSession(nil).SetEnhancedMaxAltitude(1000),
			src:  mesgdef.NewSession(nil).SetEnhancedMaxAltitude(2000),
			exp:  mesgdef.NewSession(nil).SetEnhancedMaxAltitude(2000),
		},
		{
			name: "Min",
			dst:  mesgdef.NewSession(nil).SetMinHeartRate(60),
			src:  mesgdef.NewSession(nil).SetMinHeartRate(80),
			exp:  mesgdef.NewSession(nil).SetMinHeartRate(60),
		},
		{
			name: "EnhancedMin",
			dst:  mesgdef.NewSession(nil).SetEnhancedMinAltitude(60),
			src:  mesgdef.NewSession(nil).SetEnhancedMinAltitude(80),
			exp:  mesgdef.NewSession(nil).SetEnhancedMinAltitude(60),
		},
		{
			name: "Avg",
			dst:  mesgdef.NewSession(nil).SetAvgTemperature(20),
			src:  mesgdef.NewSession(nil).SetAvgTemperature(22),
			exp:  mesgdef.NewSession(nil).SetAvgTemperature(21),
		},
		{
			name: "EnhancedAvg",
			dst:  mesgdef.NewSession(nil).SetEnhancedAvgAltitude(60),
			src:  mesgdef.NewSession(nil).SetEnhancedAvgAltitude(80),
			exp:  mesgdef.NewSession(nil).SetEnhancedAvgAltitude(70),
		},
		{
			name: "Fill dst invalid",
			dst:  mesgdef.NewSession(nil),
			src:  mesgdef.NewSession(nil).SetSport(typedef.SportCycling),
			exp:  mesgdef.NewSession(nil).SetSport(typedef.SportCycling),
		},
		{
			name: "Fill all valid",
			dst:  mesgdef.NewSession(nil).SetSubSport(typedef.SubSportRoad),
			src:  mesgdef.NewSession(nil).SetSubSport(typedef.SubSportAll),
			exp:  mesgdef.NewSession(nil).SetSubSport(typedef.SubSportRoad),
		},
		{
			name: "Fill timestamp dst not zero",
			dst:  mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(2)),
			src:  mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(5)),
			exp:  mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(2)),
		},
		{
			name: "Fill timestamp dst zero",
			dst:  mesgdef.NewSession(nil),
			src:  mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(5)),
			exp:  mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(5)),
		},
		{
			name: "Total struct pointer embedding",
			dst: &struct {
				*mesgdef.Session
				Total uint8
			}{
				Session: mesgdef.NewSession(nil).SetTotalTimerTime(10),
				Total:   1,
			},
			src: &struct {
				*mesgdef.Session
				Total uint8
			}{
				Session: mesgdef.NewSession(nil).SetTotalTimerTime(10),
				Total:   2,
			},
			exp: &struct {
				*mesgdef.Session
				Total uint8
			}{
				Session: mesgdef.NewSession(nil).SetTotalTimerTime(20),
				Total:   3,
			},
		},
		{
			name: "Total struct pointer field",
			dst: &struct {
				S     *mesgdef.Session
				Total uint8
			}{
				S:     mesgdef.NewSession(nil),
				Total: 1,
			},
			src: &struct {
				S     *mesgdef.Session
				Total uint8
			}{
				S:     mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(1)),
				Total: 2,
			},
			exp: &struct {
				S     *mesgdef.Session
				Total uint8
			}{
				S:     mesgdef.NewSession(nil).SetTimestamp(datetime.Epoch().Add(1)),
				Total: 3,
			},
		},
		{
			name: "unknown fields",
			dst: mesgdef.NewSession(nil).SetUnknownFields(
				proto.Field{FieldBase: &proto.FieldBase{Num: 1}, Value: proto.Uint8(1)},
				proto.Field{FieldBase: &proto.FieldBase{Num: 2}, Value: proto.Uint8(2)},
			),
			src: mesgdef.NewSession(nil).SetUnknownFields(
				proto.Field{FieldBase: &proto.FieldBase{Num: 2}, Value: proto.Uint8(22)},
				proto.Field{FieldBase: &proto.FieldBase{Num: 3}, Value: proto.Uint8(3)},
			),
			exp: mesgdef.NewSession(nil).SetUnknownFields(
				proto.Field{FieldBase: &proto.FieldBase{Num: 1}, Value: proto.Uint8(1)},
				proto.Field{FieldBase: &proto.FieldBase{Num: 2}, Value: proto.Uint8(2)},
				proto.Field{FieldBase: &proto.FieldBase{Num: 3}, Value: proto.Uint8(3)},
			),
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			aggregator.Aggregate(tc.dst, tc.src)

			if diff := cmp.Diff(tc.dst, tc.exp,
				cmp.Exporter(func(t reflect.Type) bool { return true }),
				cmp.Transformer("float32", func(f32 float32) uint32 {
					return math.Float32bits(f32)
				}),
				cmp.Transformer("float64", func(f64 float64) uint64 {
					return math.Float64bits(f64)
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestTotal(t *testing.T) {
	tt := []struct {
		name string
		dst  any
		src  any
		exp  any
	}{
		{name: "Total int8 all valid",
			dst: &struct{ Total int8 }{Total: 10},
			src: &struct{ Total int8 }{Total: 10},
			exp: &struct{ Total int8 }{Total: 20}},
		{name: "Total uint8 all valid",
			dst: &struct{ Total uint8 }{Total: 10},
			src: &struct{ Total uint8 }{Total: 10},
			exp: &struct{ Total uint8 }{Total: 20}},
		{name: "Total int16 all valid",
			dst: &struct{ Total int16 }{Total: 10},
			src: &struct{ Total int16 }{Total: 10},
			exp: &struct{ Total int16 }{Total: 20}},
		{name: "Total uint16 all valid",
			dst: &struct{ Total uint16 }{Total: 10},
			src: &struct{ Total uint16 }{Total: 10},
			exp: &struct{ Total uint16 }{Total: 20}},
		{name: "Total int32 all valid",
			dst: &struct{ Total int32 }{Total: 10},
			src: &struct{ Total int32 }{Total: 10},
			exp: &struct{ Total int32 }{Total: 20}},
		{name: "Total uint32 all valid",
			dst: &struct{ Total uint32 }{Total: 10},
			src: &struct{ Total uint32 }{Total: 10},
			exp: &struct{ Total uint32 }{Total: 20}},
		{name: "Total int64 all valid",
			dst: &struct{ Total int64 }{Total: 10},
			src: &struct{ Total int64 }{Total: 10},
			exp: &struct{ Total int64 }{Total: 20}},
		{name: "Total uint64 all valid",
			dst: &struct{ Total uint64 }{Total: 10},
			src: &struct{ Total uint64 }{Total: 10},
			exp: &struct{ Total uint64 }{Total: 20}},
		{name: "Total float32 all valid",
			dst: &struct{ Total float32 }{Total: 10},
			src: &struct{ Total float32 }{Total: 10.5},
			exp: &struct{ Total float32 }{Total: 20.5}},
		{name: "Total float64 all valid",
			dst: &struct{ Total float64 }{Total: 10.5},
			src: &struct{ Total float64 }{Total: 10},
			exp: &struct{ Total float64 }{Total: 20.5}},
		{name: "Total []uint8 same size",
			dst: &struct{ Total []uint8 }{Total: []uint8{1, 2, 3}},
			src: &struct{ Total []uint8 }{Total: []uint8{1, 2, 3}},
			exp: &struct{ Total []uint8 }{Total: []uint8{2, 4, 6}}},

		{name: "Total int8 dst invalid",
			dst: &struct{ Total int8 }{Total: basetype.Sint8Invalid},
			src: &struct{ Total int8 }{Total: 10},
			exp: &struct{ Total int8 }{Total: 10}},
		{name: "Total uint8 dst invalid",
			dst: &struct{ Total uint8 }{Total: basetype.Uint8Invalid},
			src: &struct{ Total uint8 }{Total: 10},
			exp: &struct{ Total uint8 }{Total: 10}},
		{name: "Total int16 dst invalid",
			dst: &struct{ Total int16 }{Total: basetype.Sint16Invalid},
			src: &struct{ Total int16 }{Total: 10},
			exp: &struct{ Total int16 }{Total: 10}},
		{name: "Total uint16 dst invalid",
			dst: &struct{ Total uint16 }{Total: basetype.Uint16Invalid},
			src: &struct{ Total uint16 }{Total: 10},
			exp: &struct{ Total uint16 }{Total: 10}},
		{name: "Total int32 dst invalid",
			dst: &struct{ Total int32 }{Total: basetype.Sint32Invalid},
			src: &struct{ Total int32 }{Total: 10},
			exp: &struct{ Total int32 }{Total: 10}},
		{name: "Total uint32 dst invalid",
			dst: &struct{ Total uint32 }{Total: basetype.Uint32Invalid},
			src: &struct{ Total uint32 }{Total: 10},
			exp: &struct{ Total uint32 }{Total: 10}},
		{name: "Total int64 dst invalid",
			dst: &struct{ Total int64 }{Total: basetype.Sint64Invalid},
			src: &struct{ Total int64 }{Total: 10},
			exp: &struct{ Total int64 }{Total: 10}},
		{name: "Total uint64 dst invalid",
			dst: &struct{ Total uint64 }{Total: basetype.Uint64Invalid},
			src: &struct{ Total uint64 }{Total: 10},
			exp: &struct{ Total uint64 }{Total: 10}},
		{name: "Total float32 dst invalid",
			dst: &struct{ Total float32 }{Total: math.Float32frombits(basetype.Float32Invalid)},
			src: &struct{ Total float32 }{Total: 10},
			exp: &struct{ Total float32 }{Total: 10}},
		{name: "Total float64 dst invalid",
			dst: &struct{ Total float64 }{Total: math.Float64frombits(basetype.Float64Invalid)},
			src: &struct{ Total float64 }{Total: 10},
			exp: &struct{ Total float64 }{Total: 10}},
		{name: "Total []uint8 dst less",
			dst: &struct{ Total []uint8 }{Total: []uint8{1, 2}},
			src: &struct{ Total []uint8 }{Total: []uint8{1, 2, 3}},
			exp: &struct{ Total []uint8 }{Total: []uint8{2, 4, 3}}},

		{name: "Total int8 src invalid",
			dst: &struct{ Total int8 }{Total: 10},
			src: &struct{ Total int8 }{Total: basetype.Sint8Invalid},
			exp: &struct{ Total int8 }{Total: 10}},
		{name: "Total uint8 src invalid",
			dst: &struct{ Total uint8 }{Total: 10},
			src: &struct{ Total uint8 }{Total: basetype.Uint8Invalid},
			exp: &struct{ Total uint8 }{Total: 10}},
		{name: "Total int16 src invalid",
			dst: &struct{ Total int16 }{Total: 10},
			src: &struct{ Total int16 }{Total: basetype.Sint16Invalid},
			exp: &struct{ Total int16 }{Total: 10}},
		{name: "Total uint16 src invalid",
			dst: &struct{ Total uint16 }{Total: 10},
			src: &struct{ Total uint16 }{Total: basetype.Uint16Invalid},
			exp: &struct{ Total uint16 }{Total: 10}},
		{name: "Total int32 src invalid",
			dst: &struct{ Total int32 }{Total: 10},
			src: &struct{ Total int32 }{Total: basetype.Sint32Invalid},
			exp: &struct{ Total int32 }{Total: 10}},
		{name: "Total uint32 src invalid",
			dst: &struct{ Total uint32 }{Total: 10},
			src: &struct{ Total uint32 }{Total: basetype.Uint32Invalid},
			exp: &struct{ Total uint32 }{Total: 10}},
		{name: "Total int64 src invalid",
			dst: &struct{ Total int64 }{Total: 10},
			src: &struct{ Total int64 }{Total: basetype.Sint64Invalid},
			exp: &struct{ Total int64 }{Total: 10}},
		{name: "Total uint64 src invalid",
			dst: &struct{ Total uint64 }{Total: 10},
			src: &struct{ Total uint64 }{Total: basetype.Uint64Invalid},
			exp: &struct{ Total uint64 }{Total: 10}},
		{name: "Total float32 src invalid",
			dst: &struct{ Total float32 }{Total: 10},
			src: &struct{ Total float32 }{Total: math.Float32frombits(basetype.Float32Invalid)},
			exp: &struct{ Total float32 }{Total: 10}},
		{name: "Total float64 src invalid",
			dst: &struct{ Total float64 }{Total: 10},
			src: &struct{ Total float64 }{Total: math.Float64frombits(basetype.Float64Invalid)},
			exp: &struct{ Total float64 }{Total: 10}},
		{name: "Total []uint8 src less",
			dst: &struct{ Total []uint8 }{Total: []uint8{1, 2, 3}},
			src: &struct{ Total []uint8 }{Total: []uint8{1, 2}},
			exp: &struct{ Total []uint8 }{Total: []uint8{2, 4, 3}}},

		{name: "Total [2]uint8 array",
			dst: &struct{ Total [2]uint8 }{Total: [2]uint8{1, 2}},
			src: &struct{ Total [2]uint8 }{Total: [2]uint8{1, 2}},
			exp: &struct{ Total [2]uint8 }{Total: [2]uint8{2, 4}}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			aggregator.Aggregate(tc.dst, tc.src)

			if diff := cmp.Diff(tc.dst, tc.exp,
				cmp.Exporter(func(t reflect.Type) bool { return true }),
				cmp.Transformer("float32", func(f32 float32) uint32 {
					return math.Float32bits(f32)
				}),
				cmp.Transformer("float64", func(f64 float64) uint64 {
					return math.Float64bits(f64)
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tt := []struct {
		name string
		dst  any
		src  any
		exp  any
	}{
		{name: "Max int8 all valid",
			dst: &struct{ Max int8 }{Max: 10},
			src: &struct{ Max int8 }{Max: 20},
			exp: &struct{ Max int8 }{Max: 20}},
		{name: "Max uint8 all valid",
			dst: &struct{ Max uint8 }{Max: 10},
			src: &struct{ Max uint8 }{Max: 20},
			exp: &struct{ Max uint8 }{Max: 20}},
		{name: "Max int16 all valid",
			dst: &struct{ Max int16 }{Max: 10},
			src: &struct{ Max int16 }{Max: 20},
			exp: &struct{ Max int16 }{Max: 20}},
		{name: "Max uint16 all valid",
			dst: &struct{ Max uint16 }{Max: 10},
			src: &struct{ Max uint16 }{Max: 20},
			exp: &struct{ Max uint16 }{Max: 20}},
		{name: "Max int32 all valid",
			dst: &struct{ Max int32 }{Max: 10},
			src: &struct{ Max int32 }{Max: 20},
			exp: &struct{ Max int32 }{Max: 20}},
		{name: "Max uint32 all valid",
			dst: &struct{ Max uint32 }{Max: 10},
			src: &struct{ Max uint32 }{Max: 20},
			exp: &struct{ Max uint32 }{Max: 20}},
		{name: "Max int64 all valid",
			dst: &struct{ Max int64 }{Max: 10},
			src: &struct{ Max int64 }{Max: 20},
			exp: &struct{ Max int64 }{Max: 20}},
		{name: "Max uint64 all valid",
			dst: &struct{ Max uint64 }{Max: 10},
			src: &struct{ Max uint64 }{Max: 20},
			exp: &struct{ Max uint64 }{Max: 20}},
		{name: "Max float32 all valid",
			dst: &struct{ Max float32 }{Max: 10},
			src: &struct{ Max float32 }{Max: 20},
			exp: &struct{ Max float32 }{Max: 20}},
		{name: "Max float64 all valid",
			dst: &struct{ Max float64 }{Max: 10},
			src: &struct{ Max float64 }{Max: 20},
			exp: &struct{ Max float64 }{Max: 20}},
		{name: "Max []uint8 same size",
			dst: &struct{ Max []uint8 }{Max: []uint8{1, 2, 3}},
			src: &struct{ Max []uint8 }{Max: []uint8{2, 1, 3}},
			exp: &struct{ Max []uint8 }{Max: []uint8{2, 2, 3}}},

		{name: "Max int8 dst invalid",
			dst: &struct{ Max int8 }{Max: basetype.Sint8Invalid},
			src: &struct{ Max int8 }{Max: 20},
			exp: &struct{ Max int8 }{Max: 20}},
		{name: "Max uint8 dst invalid",
			dst: &struct{ Max uint8 }{Max: basetype.Uint8Invalid},
			src: &struct{ Max uint8 }{Max: 20},
			exp: &struct{ Max uint8 }{Max: 20}},
		{name: "Max int16 dst invalid",
			dst: &struct{ Max int16 }{Max: basetype.Sint16Invalid},
			src: &struct{ Max int16 }{Max: 20},
			exp: &struct{ Max int16 }{Max: 20}},
		{name: "Max uint16 dst invalid",
			dst: &struct{ Max uint16 }{Max: basetype.Uint16Invalid},
			src: &struct{ Max uint16 }{Max: 20},
			exp: &struct{ Max uint16 }{Max: 20}},
		{name: "Max int32 dst invalid",
			dst: &struct{ Max int32 }{Max: basetype.Sint32Invalid},
			src: &struct{ Max int32 }{Max: 20},
			exp: &struct{ Max int32 }{Max: 20}},
		{name: "Max uint32 dst invalid",
			dst: &struct{ Max uint32 }{Max: basetype.Uint32Invalid},
			src: &struct{ Max uint32 }{Max: 20},
			exp: &struct{ Max uint32 }{Max: 20}},
		{name: "Max int64 dst invalid",
			dst: &struct{ Max int64 }{Max: basetype.Sint64Invalid},
			src: &struct{ Max int64 }{Max: 20},
			exp: &struct{ Max int64 }{Max: 20}},
		{name: "Max uint64 dst invalid",
			dst: &struct{ Max uint64 }{Max: basetype.Uint64Invalid},
			src: &struct{ Max uint64 }{Max: 20},
			exp: &struct{ Max uint64 }{Max: 20}},
		{name: "Max float32 dst invalid",
			dst: &struct{ Max float32 }{Max: math.Float32frombits(basetype.Float32Invalid)},
			src: &struct{ Max float32 }{Max: 20},
			exp: &struct{ Max float32 }{Max: 20}},
		{name: "Max float64 dst invalid",
			dst: &struct{ Max float64 }{Max: math.Float64frombits(basetype.Float64Invalid)},
			src: &struct{ Max float64 }{Max: 20},
			exp: &struct{ Max float64 }{Max: 20}},
		{name: "Max []uint8 dst less",
			dst: &struct{ Max []uint8 }{Max: []uint8{1, 2}},
			src: &struct{ Max []uint8 }{Max: []uint8{2, 1, 3}},
			exp: &struct{ Max []uint8 }{Max: []uint8{2, 2, 3}}},

		{name: "Max int8 src invalid",
			dst: &struct{ Max int8 }{Max: 10},
			src: &struct{ Max int8 }{Max: basetype.Sint8Invalid},
			exp: &struct{ Max int8 }{Max: 10}},
		{name: "Max uint8 src invalid",
			dst: &struct{ Max uint8 }{Max: 10},
			src: &struct{ Max uint8 }{Max: basetype.Uint8Invalid},
			exp: &struct{ Max uint8 }{Max: 10}},
		{name: "Max int16 src invalid",
			dst: &struct{ Max int16 }{Max: 10},
			src: &struct{ Max int16 }{Max: basetype.Sint16Invalid},
			exp: &struct{ Max int16 }{Max: 10}},
		{name: "Max uint16 src invalid",
			dst: &struct{ Max uint16 }{Max: 10},
			src: &struct{ Max uint16 }{Max: basetype.Uint16Invalid},
			exp: &struct{ Max uint16 }{Max: 10}},
		{name: "Max int32 src invalid",
			dst: &struct{ Max int32 }{Max: 10},
			src: &struct{ Max int32 }{Max: basetype.Sint32Invalid},
			exp: &struct{ Max int32 }{Max: 10}},
		{name: "Max uint32 src invalid",
			dst: &struct{ Max uint32 }{Max: 10},
			src: &struct{ Max uint32 }{Max: basetype.Uint32Invalid},
			exp: &struct{ Max uint32 }{Max: 10}},
		{name: "Max int64 src invalid",
			dst: &struct{ Max int64 }{Max: 10},
			src: &struct{ Max int64 }{Max: basetype.Sint64Invalid},
			exp: &struct{ Max int64 }{Max: 10}},
		{name: "Max uint64 src invalid",
			dst: &struct{ Max uint64 }{Max: 10},
			src: &struct{ Max uint64 }{Max: basetype.Uint64Invalid},
			exp: &struct{ Max uint64 }{Max: 10}},
		{name: "Max float32 src invalid",
			dst: &struct{ Max float32 }{Max: 10},
			src: &struct{ Max float32 }{Max: math.Float32frombits(basetype.Float32Invalid)},
			exp: &struct{ Max float32 }{Max: 10}},
		{name: "Max float64 src invalid",
			dst: &struct{ Max float64 }{Max: 10},
			src: &struct{ Max float64 }{Max: math.Float64frombits(basetype.Float64Invalid)},
			exp: &struct{ Max float64 }{Max: 10}},
		{name: "Max []uint8 src less",
			dst: &struct{ Max []uint8 }{Max: []uint8{1, 2, 3}},
			src: &struct{ Max []uint8 }{Max: []uint8{2, 1}},
			exp: &struct{ Max []uint8 }{Max: []uint8{2, 2, 3}}},

		{name: "Max [2]uint8 array",
			dst: &struct{ Max [2]uint8 }{Max: [2]uint8{1, 2}},
			src: &struct{ Max [2]uint8 }{Max: [2]uint8{2, 1}},
			exp: &struct{ Max [2]uint8 }{Max: [2]uint8{2, 2}}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			aggregator.Aggregate(tc.dst, tc.src)

			if diff := cmp.Diff(tc.dst, tc.exp,
				cmp.Exporter(func(t reflect.Type) bool { return true }),
				cmp.Transformer("float32", func(f32 float32) uint32 {
					return math.Float32bits(f32)
				}),
				cmp.Transformer("float64", func(f64 float64) uint64 {
					return math.Float64bits(f64)
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tt := []struct {
		name string
		dst  any
		src  any
		exp  any
	}{
		{name: "Min int8 all valid",
			dst: &struct{ Min int8 }{Min: 20},
			src: &struct{ Min int8 }{Min: 10},
			exp: &struct{ Min int8 }{Min: 10}},
		{name: "Min uint8 all valid",
			dst: &struct{ Min uint8 }{Min: 20},
			src: &struct{ Min uint8 }{Min: 10},
			exp: &struct{ Min uint8 }{Min: 10}},
		{name: "Min int16 all valid",
			dst: &struct{ Min int16 }{Min: 20},
			src: &struct{ Min int16 }{Min: 10},
			exp: &struct{ Min int16 }{Min: 10}},
		{name: "Min uint16 all valid",
			dst: &struct{ Min uint16 }{Min: 20},
			src: &struct{ Min uint16 }{Min: 10},
			exp: &struct{ Min uint16 }{Min: 10}},
		{name: "Min int32 all valid",
			dst: &struct{ Min int32 }{Min: 20},
			src: &struct{ Min int32 }{Min: 10},
			exp: &struct{ Min int32 }{Min: 10}},
		{name: "Min uint32 all valid",
			dst: &struct{ Min uint32 }{Min: 20},
			src: &struct{ Min uint32 }{Min: 10},
			exp: &struct{ Min uint32 }{Min: 10}},
		{name: "Min int64 all valid",
			dst: &struct{ Min int64 }{Min: 20},
			src: &struct{ Min int64 }{Min: 10},
			exp: &struct{ Min int64 }{Min: 10}},
		{name: "Min uint64 all valid",
			dst: &struct{ Min uint64 }{Min: 20},
			src: &struct{ Min uint64 }{Min: 10},
			exp: &struct{ Min uint64 }{Min: 10}},
		{name: "Min float32 all valid",
			dst: &struct{ Min float32 }{Min: 20},
			src: &struct{ Min float32 }{Min: 10},
			exp: &struct{ Min float32 }{Min: 10}},
		{name: "Min float64 all valid",
			dst: &struct{ Min float64 }{Min: 20},
			src: &struct{ Min float64 }{Min: 10},
			exp: &struct{ Min float64 }{Min: 10}},
		{name: "Min []uint8 same size",
			dst: &struct{ Min []uint8 }{Min: []uint8{1, 2, 3}},
			src: &struct{ Min []uint8 }{Min: []uint8{2, 1, 3}},
			exp: &struct{ Min []uint8 }{Min: []uint8{1, 1, 3}}},

		{name: "Min int8 dst invalid",
			dst: &struct{ Min int8 }{Min: basetype.Sint8Invalid},
			src: &struct{ Min int8 }{Min: 20},
			exp: &struct{ Min int8 }{Min: 20}},
		{name: "Min uint8 dst invalid",
			dst: &struct{ Min uint8 }{Min: basetype.Uint8Invalid},
			src: &struct{ Min uint8 }{Min: 20},
			exp: &struct{ Min uint8 }{Min: 20}},
		{name: "Min int16 dst invalid",
			dst: &struct{ Min int16 }{Min: basetype.Sint16Invalid},
			src: &struct{ Min int16 }{Min: 20},
			exp: &struct{ Min int16 }{Min: 20}},
		{name: "Min uint16 dst invalid",
			dst: &struct{ Min uint16 }{Min: basetype.Uint16Invalid},
			src: &struct{ Min uint16 }{Min: 20},
			exp: &struct{ Min uint16 }{Min: 20}},
		{name: "Min int32 dst invalid",
			dst: &struct{ Min int32 }{Min: basetype.Sint32Invalid},
			src: &struct{ Min int32 }{Min: 20},
			exp: &struct{ Min int32 }{Min: 20}},
		{name: "Min uint32 dst invalid",
			dst: &struct{ Min uint32 }{Min: basetype.Uint32Invalid},
			src: &struct{ Min uint32 }{Min: 20},
			exp: &struct{ Min uint32 }{Min: 20}},
		{name: "Min int64 dst invalid",
			dst: &struct{ Min int64 }{Min: basetype.Sint64Invalid},
			src: &struct{ Min int64 }{Min: 20},
			exp: &struct{ Min int64 }{Min: 20}},
		{name: "Min uint64 dst invalid",
			dst: &struct{ Min uint64 }{Min: basetype.Uint64Invalid},
			src: &struct{ Min uint64 }{Min: 20},
			exp: &struct{ Min uint64 }{Min: 20}},
		{name: "Min float32 dst invalid",
			dst: &struct{ Min float32 }{Min: math.Float32frombits(basetype.Float32Invalid)},
			src: &struct{ Min float32 }{Min: 20},
			exp: &struct{ Min float32 }{Min: 20}},
		{name: "Min float64 dst invalid",
			dst: &struct{ Min float64 }{Min: math.Float64frombits(basetype.Float64Invalid)},
			src: &struct{ Min float64 }{Min: 20},
			exp: &struct{ Min float64 }{Min: 20}},
		{name: "Min []uint8 dst less",
			dst: &struct{ Min []uint8 }{Min: []uint8{1, 2}},
			src: &struct{ Min []uint8 }{Min: []uint8{2, 1, 3}},
			exp: &struct{ Min []uint8 }{Min: []uint8{1, 1, 3}}},

		{name: "Min int8 src invalid",
			dst: &struct{ Min int8 }{Min: 10},
			src: &struct{ Min int8 }{Min: basetype.Sint8Invalid},
			exp: &struct{ Min int8 }{Min: 10}},
		{name: "Min uint8 src invalid",
			dst: &struct{ Min uint8 }{Min: 10},
			src: &struct{ Min uint8 }{Min: basetype.Uint8Invalid},
			exp: &struct{ Min uint8 }{Min: 10}},
		{name: "Min int16 src invalid",
			dst: &struct{ Min int16 }{Min: 10},
			src: &struct{ Min int16 }{Min: basetype.Sint16Invalid},
			exp: &struct{ Min int16 }{Min: 10}},
		{name: "Min uint16 src invalid",
			dst: &struct{ Min uint16 }{Min: 10},
			src: &struct{ Min uint16 }{Min: basetype.Uint16Invalid},
			exp: &struct{ Min uint16 }{Min: 10}},
		{name: "Min int32 src invalid",
			dst: &struct{ Min int32 }{Min: 10},
			src: &struct{ Min int32 }{Min: basetype.Sint32Invalid},
			exp: &struct{ Min int32 }{Min: 10}},
		{name: "Min uint32 src invalid",
			dst: &struct{ Min uint32 }{Min: 10},
			src: &struct{ Min uint32 }{Min: basetype.Uint32Invalid},
			exp: &struct{ Min uint32 }{Min: 10}},
		{name: "Min int64 src invalid",
			dst: &struct{ Min int64 }{Min: 10},
			src: &struct{ Min int64 }{Min: basetype.Sint64Invalid},
			exp: &struct{ Min int64 }{Min: 10}},
		{name: "Min uint64 src invalid",
			dst: &struct{ Min uint64 }{Min: 10},
			src: &struct{ Min uint64 }{Min: basetype.Uint64Invalid},
			exp: &struct{ Min uint64 }{Min: 10}},
		{name: "Min float32 src invalid",
			dst: &struct{ Min float32 }{Min: 10},
			src: &struct{ Min float32 }{Min: math.Float32frombits(basetype.Float32Invalid)},
			exp: &struct{ Min float32 }{Min: 10}},
		{name: "Min float64 src invalid",
			dst: &struct{ Min float64 }{Min: 10},
			src: &struct{ Min float64 }{Min: math.Float64frombits(basetype.Float64Invalid)},
			exp: &struct{ Min float64 }{Min: 10}},
		{name: "Min []uint8 src less",
			dst: &struct{ Min []uint8 }{Min: []uint8{1, 2, 3}},
			src: &struct{ Min []uint8 }{Min: []uint8{2, 1}},
			exp: &struct{ Min []uint8 }{Min: []uint8{1, 1, 3}}},

		{name: "Min [2]uint8 array",
			dst: &struct{ Min [2]uint8 }{Min: [2]uint8{1, 2}},
			src: &struct{ Min [2]uint8 }{Min: [2]uint8{2, 1}},
			exp: &struct{ Min [2]uint8 }{Min: [2]uint8{1, 1}}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			aggregator.Aggregate(tc.dst, tc.src)

			if diff := cmp.Diff(tc.dst, tc.exp,
				cmp.Exporter(func(t reflect.Type) bool { return true }),
				cmp.Transformer("float32", func(f32 float32) uint32 {
					return math.Float32bits(f32)
				}),
				cmp.Transformer("float64", func(f64 float64) uint64 {
					return math.Float64bits(f64)
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestAvg(t *testing.T) {
	tt := []struct {
		name string
		dst  any
		src  any
		exp  any
	}{
		{name: "Avg int8 all valid",
			dst: &struct{ Avg int8 }{Avg: 20},
			src: &struct{ Avg int8 }{Avg: 10},
			exp: &struct{ Avg int8 }{Avg: 15}},
		{name: "Avg uint8 all valid",
			dst: &struct{ Avg uint8 }{Avg: 20},
			src: &struct{ Avg uint8 }{Avg: 10},
			exp: &struct{ Avg uint8 }{Avg: 15}},
		{name: "Avg int16 all valid",
			dst: &struct{ Avg int16 }{Avg: 20},
			src: &struct{ Avg int16 }{Avg: 10},
			exp: &struct{ Avg int16 }{Avg: 15}},
		{name: "Avg uint16 all valid",
			dst: &struct{ Avg uint16 }{Avg: 20},
			src: &struct{ Avg uint16 }{Avg: 10},
			exp: &struct{ Avg uint16 }{Avg: 15}},
		{name: "Avg int32 all valid",
			dst: &struct{ Avg int32 }{Avg: 20},
			src: &struct{ Avg int32 }{Avg: 10},
			exp: &struct{ Avg int32 }{Avg: 15}},
		{name: "Avg uint32 all valid",
			dst: &struct{ Avg uint32 }{Avg: 20},
			src: &struct{ Avg uint32 }{Avg: 10},
			exp: &struct{ Avg uint32 }{Avg: 15}},
		{name: "Avg int64 all valid",
			dst: &struct{ Avg int64 }{Avg: 20},
			src: &struct{ Avg int64 }{Avg: 10},
			exp: &struct{ Avg int64 }{Avg: 15}},
		{name: "Avg uint64 all valid",
			dst: &struct{ Avg uint64 }{Avg: 20},
			src: &struct{ Avg uint64 }{Avg: 10},
			exp: &struct{ Avg uint64 }{Avg: 15}},
		{name: "Avg float32 all valid",
			dst: &struct{ Avg float32 }{Avg: 20},
			src: &struct{ Avg float32 }{Avg: 11},
			exp: &struct{ Avg float32 }{Avg: 15.5}},
		{name: "Avg float64 all valid",
			dst: &struct{ Avg float64 }{Avg: 20},
			src: &struct{ Avg float64 }{Avg: 11},
			exp: &struct{ Avg float64 }{Avg: 15.5}},
		{name: "Avg []uint8 same size",
			dst: &struct{ Avg []uint8 }{Avg: []uint8{10, 5, 30}},
			src: &struct{ Avg []uint8 }{Avg: []uint8{20, 5, 50}},
			exp: &struct{ Avg []uint8 }{Avg: []uint8{15, 5, 40}}},

		{name: "Avg int8 dst invalid",
			dst: &struct{ Avg int8 }{Avg: basetype.Sint8Invalid},
			src: &struct{ Avg int8 }{Avg: 20},
			exp: &struct{ Avg int8 }{Avg: 20}},
		{name: "Avg uint8 dst invalid",
			dst: &struct{ Avg uint8 }{Avg: basetype.Uint8Invalid},
			src: &struct{ Avg uint8 }{Avg: 20},
			exp: &struct{ Avg uint8 }{Avg: 20}},
		{name: "Avg int16 dst invalid",
			dst: &struct{ Avg int16 }{Avg: basetype.Sint16Invalid},
			src: &struct{ Avg int16 }{Avg: 20},
			exp: &struct{ Avg int16 }{Avg: 20}},
		{name: "Avg uint16 dst invalid",
			dst: &struct{ Avg uint16 }{Avg: basetype.Uint16Invalid},
			src: &struct{ Avg uint16 }{Avg: 20},
			exp: &struct{ Avg uint16 }{Avg: 20}},
		{name: "Avg int32 dst invalid",
			dst: &struct{ Avg int32 }{Avg: basetype.Sint32Invalid},
			src: &struct{ Avg int32 }{Avg: 20},
			exp: &struct{ Avg int32 }{Avg: 20}},
		{name: "Avg uint32 dst invalid",
			dst: &struct{ Avg uint32 }{Avg: basetype.Uint32Invalid},
			src: &struct{ Avg uint32 }{Avg: 20},
			exp: &struct{ Avg uint32 }{Avg: 20}},
		{name: "Avg int64 dst invalid",
			dst: &struct{ Avg int64 }{Avg: basetype.Sint64Invalid},
			src: &struct{ Avg int64 }{Avg: 20},
			exp: &struct{ Avg int64 }{Avg: 20}},
		{name: "Avg uint64 dst invalid",
			dst: &struct{ Avg uint64 }{Avg: basetype.Uint64Invalid},
			src: &struct{ Avg uint64 }{Avg: 20},
			exp: &struct{ Avg uint64 }{Avg: 20}},
		{name: "Avg float32 dst invalid",
			dst: &struct{ Avg float32 }{Avg: math.Float32frombits(basetype.Float32Invalid)},
			src: &struct{ Avg float32 }{Avg: 20},
			exp: &struct{ Avg float32 }{Avg: 20}},
		{name: "Avg float64 dst invalid",
			dst: &struct{ Avg float64 }{Avg: math.Float64frombits(basetype.Float64Invalid)},
			src: &struct{ Avg float64 }{Avg: 20},
			exp: &struct{ Avg float64 }{Avg: 20}},
		{name: "Avg []uint8 dst less",
			dst: &struct{ Avg []uint8 }{Avg: []uint8{10, 5}},
			src: &struct{ Avg []uint8 }{Avg: []uint8{20, 5, 50}},
			exp: &struct{ Avg []uint8 }{Avg: []uint8{15, 5, 50}}},

		{name: "Avg int8 src invalid",
			dst: &struct{ Avg int8 }{Avg: 10},
			src: &struct{ Avg int8 }{Avg: basetype.Sint8Invalid},
			exp: &struct{ Avg int8 }{Avg: 10}},
		{name: "Avg uint8 src invalid",
			dst: &struct{ Avg uint8 }{Avg: 10},
			src: &struct{ Avg uint8 }{Avg: basetype.Uint8Invalid},
			exp: &struct{ Avg uint8 }{Avg: 10}},
		{name: "Avg int16 src invalid",
			dst: &struct{ Avg int16 }{Avg: 10},
			src: &struct{ Avg int16 }{Avg: basetype.Sint16Invalid},
			exp: &struct{ Avg int16 }{Avg: 10}},
		{name: "Avg uint16 src invalid",
			dst: &struct{ Avg uint16 }{Avg: 10},
			src: &struct{ Avg uint16 }{Avg: basetype.Uint16Invalid},
			exp: &struct{ Avg uint16 }{Avg: 10}},
		{name: "Avg int32 src invalid",
			dst: &struct{ Avg int32 }{Avg: 10},
			src: &struct{ Avg int32 }{Avg: basetype.Sint32Invalid},
			exp: &struct{ Avg int32 }{Avg: 10}},
		{name: "Avg uint32 src invalid",
			dst: &struct{ Avg uint32 }{Avg: 10},
			src: &struct{ Avg uint32 }{Avg: basetype.Uint32Invalid},
			exp: &struct{ Avg uint32 }{Avg: 10}},
		{name: "Avg int64 src invalid",
			dst: &struct{ Avg int64 }{Avg: 10},
			src: &struct{ Avg int64 }{Avg: basetype.Sint64Invalid},
			exp: &struct{ Avg int64 }{Avg: 10}},
		{name: "Avg uint64 src invalid",
			dst: &struct{ Avg uint64 }{Avg: 10},
			src: &struct{ Avg uint64 }{Avg: basetype.Uint64Invalid},
			exp: &struct{ Avg uint64 }{Avg: 10}},
		{name: "Avg float32 src invalid",
			dst: &struct{ Avg float32 }{Avg: 10},
			src: &struct{ Avg float32 }{Avg: math.Float32frombits(basetype.Float32Invalid)},
			exp: &struct{ Avg float32 }{Avg: 10}},
		{name: "Avg float64 src invalid",
			dst: &struct{ Avg float64 }{Avg: 10},
			src: &struct{ Avg float64 }{Avg: math.Float64frombits(basetype.Float64Invalid)},
			exp: &struct{ Avg float64 }{Avg: 10}},
		{name: "Avg []uint8 src less",
			dst: &struct{ Avg []uint8 }{Avg: []uint8{10, 5, 30}},
			src: &struct{ Avg []uint8 }{Avg: []uint8{20, 5}},
			exp: &struct{ Avg []uint8 }{Avg: []uint8{15, 5, 30}}},

		{name: "Avg [2]uint8 array",
			dst: &struct{ Avg [2]uint8 }{Avg: [2]uint8{10, 42}},
			src: &struct{ Avg [2]uint8 }{Avg: [2]uint8{20, 24}},
			exp: &struct{ Avg [2]uint8 }{Avg: [2]uint8{15, 33}}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			aggregator.Aggregate(tc.dst, tc.src)

			if diff := cmp.Diff(tc.dst, tc.exp,
				cmp.Exporter(func(t reflect.Type) bool { return true }),
				cmp.Transformer("float32", func(f32 float32) uint32 {
					return math.Float32bits(f32)
				}),
				cmp.Transformer("float64", func(f64 float64) uint64 {
					return math.Float64bits(f64)
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestFill(t *testing.T) {
	tt := []struct {
		name string
		dst  any
		src  any
		exp  any
	}{
		{name: "ShouldFill int8 all valid",
			dst: &struct{ ShouldFill int8 }{ShouldFill: 20},
			src: &struct{ ShouldFill int8 }{ShouldFill: 10},
			exp: &struct{ ShouldFill int8 }{ShouldFill: 20}},
		{name: "ShouldFill uint8 all valid",
			dst: &struct{ ShouldFill uint8 }{ShouldFill: 20},
			src: &struct{ ShouldFill uint8 }{ShouldFill: 10},
			exp: &struct{ ShouldFill uint8 }{ShouldFill: 20}},
		{name: "ShouldFill int16 all valid",
			dst: &struct{ ShouldFill int16 }{ShouldFill: 20},
			src: &struct{ ShouldFill int16 }{ShouldFill: 10},
			exp: &struct{ ShouldFill int16 }{ShouldFill: 20}},
		{name: "ShouldFill uint16 all valid",
			dst: &struct{ ShouldFill uint16 }{ShouldFill: 20},
			src: &struct{ ShouldFill uint16 }{ShouldFill: 10},
			exp: &struct{ ShouldFill uint16 }{ShouldFill: 20}},
		{name: "ShouldFill int32 all valid",
			dst: &struct{ ShouldFill int32 }{ShouldFill: 20},
			src: &struct{ ShouldFill int32 }{ShouldFill: 10},
			exp: &struct{ ShouldFill int32 }{ShouldFill: 20}},
		{name: "ShouldFill uint32 all valid",
			dst: &struct{ ShouldFill uint32 }{ShouldFill: 20},
			src: &struct{ ShouldFill uint32 }{ShouldFill: 10},
			exp: &struct{ ShouldFill uint32 }{ShouldFill: 20}},
		{name: "ShouldFill int64 all valid",
			dst: &struct{ ShouldFill int64 }{ShouldFill: 20},
			src: &struct{ ShouldFill int64 }{ShouldFill: 10},
			exp: &struct{ ShouldFill int64 }{ShouldFill: 20}},
		{name: "ShouldFill uint64 all valid",
			dst: &struct{ ShouldFill uint64 }{ShouldFill: 20},
			src: &struct{ ShouldFill uint64 }{ShouldFill: 10},
			exp: &struct{ ShouldFill uint64 }{ShouldFill: 20}},
		{name: "ShouldFill float32 all valid",
			dst: &struct{ ShouldFill float32 }{ShouldFill: 20},
			src: &struct{ ShouldFill float32 }{ShouldFill: 10},
			exp: &struct{ ShouldFill float32 }{ShouldFill: 20}},
		{name: "ShouldFill float64 all valid",
			dst: &struct{ ShouldFill float64 }{ShouldFill: 20},
			src: &struct{ ShouldFill float64 }{ShouldFill: 10},
			exp: &struct{ ShouldFill float64 }{ShouldFill: 20}},
		{name: "ShouldFill string all valid",
			dst: &struct{ ShouldFill string }{ShouldFill: "abc"},
			src: &struct{ ShouldFill string }{ShouldFill: "def"},
			exp: &struct{ ShouldFill string }{ShouldFill: "abc"}},
		{name: "ShouldFill bool all valid",
			dst: &struct{ ShouldFill bool }{ShouldFill: true},
			src: &struct{ ShouldFill bool }{ShouldFill: true},
			exp: &struct{ ShouldFill bool }{ShouldFill: true}},
		{name: "ShouldFill []uint8 same size",
			dst: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{1, 1, basetype.Uint8Invalid}},
			src: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{3, 3, 3}},
			exp: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{1, 1, 3}}},

		{name: "ShouldFill int8 dst invalid",
			dst: &struct{ ShouldFill int8 }{ShouldFill: basetype.Sint8Invalid},
			src: &struct{ ShouldFill int8 }{ShouldFill: 20},
			exp: &struct{ ShouldFill int8 }{ShouldFill: 20}},
		{name: "ShouldFill uint8 dst invalid",
			dst: &struct{ ShouldFill uint8 }{ShouldFill: basetype.Uint8Invalid},
			src: &struct{ ShouldFill uint8 }{ShouldFill: 20},
			exp: &struct{ ShouldFill uint8 }{ShouldFill: 20}},
		{name: "ShouldFill int16 dst invalid",
			dst: &struct{ ShouldFill int16 }{ShouldFill: basetype.Sint16Invalid},
			src: &struct{ ShouldFill int16 }{ShouldFill: 20},
			exp: &struct{ ShouldFill int16 }{ShouldFill: 20}},
		{name: "ShouldFill uint16 dst invalid",
			dst: &struct{ ShouldFill uint16 }{ShouldFill: basetype.Uint16Invalid},
			src: &struct{ ShouldFill uint16 }{ShouldFill: 20},
			exp: &struct{ ShouldFill uint16 }{ShouldFill: 20}},
		{name: "ShouldFill int32 dst invalid",
			dst: &struct{ ShouldFill int32 }{ShouldFill: basetype.Sint32Invalid},
			src: &struct{ ShouldFill int32 }{ShouldFill: 20},
			exp: &struct{ ShouldFill int32 }{ShouldFill: 20}},
		{name: "ShouldFill uint32 dst invalid",
			dst: &struct{ ShouldFill uint32 }{ShouldFill: basetype.Uint32Invalid},
			src: &struct{ ShouldFill uint32 }{ShouldFill: 20},
			exp: &struct{ ShouldFill uint32 }{ShouldFill: 20}},
		{name: "ShouldFill int64 dst invalid",
			dst: &struct{ ShouldFill int64 }{ShouldFill: basetype.Sint64Invalid},
			src: &struct{ ShouldFill int64 }{ShouldFill: 20},
			exp: &struct{ ShouldFill int64 }{ShouldFill: 20}},
		{name: "ShouldFill uint64 dst invalid",
			dst: &struct{ ShouldFill uint64 }{ShouldFill: basetype.Uint64Invalid},
			src: &struct{ ShouldFill uint64 }{ShouldFill: 20},
			exp: &struct{ ShouldFill uint64 }{ShouldFill: 20}},
		{name: "ShouldFill float32 dst invalid",
			dst: &struct{ ShouldFill float32 }{ShouldFill: math.Float32frombits(basetype.Float32Invalid)},
			src: &struct{ ShouldFill float32 }{ShouldFill: 20},
			exp: &struct{ ShouldFill float32 }{ShouldFill: 20}},
		{name: "ShouldFill float64 dst invalid",
			dst: &struct{ ShouldFill float64 }{ShouldFill: math.Float64frombits(basetype.Float64Invalid)},
			src: &struct{ ShouldFill float64 }{ShouldFill: 20},
			exp: &struct{ ShouldFill float64 }{ShouldFill: 20}},
		{name: "ShouldFill string dst invalid",
			dst: &struct{ ShouldFill string }{ShouldFill: ""},
			src: &struct{ ShouldFill string }{ShouldFill: "def"},
			exp: &struct{ ShouldFill string }{ShouldFill: "def"}},
		{name: "ShouldFill bool dst invalid",
			dst: &struct{ ShouldFill bool }{ShouldFill: false},
			src: &struct{ ShouldFill bool }{ShouldFill: true},
			exp: &struct{ ShouldFill bool }{ShouldFill: true}},
		{name: "ShouldFill []uint8 dst less",
			dst: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{1, basetype.Uint8Invalid}},
			src: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{3, 3, 3}},
			exp: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{1, 3, 3}}},

		{name: "ShouldFill int8 src invalid",
			dst: &struct{ ShouldFill int8 }{ShouldFill: 10},
			src: &struct{ ShouldFill int8 }{ShouldFill: basetype.Sint8Invalid},
			exp: &struct{ ShouldFill int8 }{ShouldFill: 10}},
		{name: "ShouldFill uint8 src invalid",
			dst: &struct{ ShouldFill uint8 }{ShouldFill: 10},
			src: &struct{ ShouldFill uint8 }{ShouldFill: basetype.Uint8Invalid},
			exp: &struct{ ShouldFill uint8 }{ShouldFill: 10}},
		{name: "ShouldFill int16 src invalid",
			dst: &struct{ ShouldFill int16 }{ShouldFill: 10},
			src: &struct{ ShouldFill int16 }{ShouldFill: basetype.Sint16Invalid},
			exp: &struct{ ShouldFill int16 }{ShouldFill: 10}},
		{name: "ShouldFill uint16 src invalid",
			dst: &struct{ ShouldFill uint16 }{ShouldFill: 10},
			src: &struct{ ShouldFill uint16 }{ShouldFill: basetype.Uint16Invalid},
			exp: &struct{ ShouldFill uint16 }{ShouldFill: 10}},
		{name: "ShouldFill int32 src invalid",
			dst: &struct{ ShouldFill int32 }{ShouldFill: 10},
			src: &struct{ ShouldFill int32 }{ShouldFill: basetype.Sint32Invalid},
			exp: &struct{ ShouldFill int32 }{ShouldFill: 10}},
		{name: "ShouldFill uint32 src invalid",
			dst: &struct{ ShouldFill uint32 }{ShouldFill: 10},
			src: &struct{ ShouldFill uint32 }{ShouldFill: basetype.Uint32Invalid},
			exp: &struct{ ShouldFill uint32 }{ShouldFill: 10}},
		{name: "ShouldFill int64 src invalid",
			dst: &struct{ ShouldFill int64 }{ShouldFill: 10},
			src: &struct{ ShouldFill int64 }{ShouldFill: basetype.Sint64Invalid},
			exp: &struct{ ShouldFill int64 }{ShouldFill: 10}},
		{name: "ShouldFill uint64 src invalid",
			dst: &struct{ ShouldFill uint64 }{ShouldFill: 10},
			src: &struct{ ShouldFill uint64 }{ShouldFill: basetype.Uint64Invalid},
			exp: &struct{ ShouldFill uint64 }{ShouldFill: 10}},
		{name: "ShouldFill float32 src invalid",
			dst: &struct{ ShouldFill float32 }{ShouldFill: 10},
			src: &struct{ ShouldFill float32 }{ShouldFill: math.Float32frombits(basetype.Float32Invalid)},
			exp: &struct{ ShouldFill float32 }{ShouldFill: 10}},
		{name: "ShouldFill float64 src invalid",
			dst: &struct{ ShouldFill float64 }{ShouldFill: 10},
			src: &struct{ ShouldFill float64 }{ShouldFill: math.Float64frombits(basetype.Float64Invalid)},
			exp: &struct{ ShouldFill float64 }{ShouldFill: 10}},
		{name: "ShouldFill string src invalid",
			dst: &struct{ ShouldFill string }{ShouldFill: "abc"},
			src: &struct{ ShouldFill string }{ShouldFill: ""},
			exp: &struct{ ShouldFill string }{ShouldFill: "abc"}},
		{name: "ShouldFill bool src invalid",
			dst: &struct{ ShouldFill bool }{ShouldFill: true},
			src: &struct{ ShouldFill bool }{ShouldFill: false},
			exp: &struct{ ShouldFill bool }{ShouldFill: true}},
		{name: "ShouldFill []uint8 src less",
			dst: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{basetype.Uint8Invalid, 1, 1}},
			src: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{3, 3}},
			exp: &struct{ ShouldFill []uint8 }{ShouldFill: []uint8{3, 1, 1}}},

		{name: "ShouldFill [2]uint8 array",
			dst: &struct{ ShouldFill [2]uint8 }{ShouldFill: [2]uint8{basetype.Uint8Invalid, 1}},
			src: &struct{ ShouldFill [2]uint8 }{ShouldFill: [2]uint8{2, 2}},
			exp: &struct{ ShouldFill [2]uint8 }{ShouldFill: [2]uint8{2, 1}}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			aggregator.Aggregate(tc.dst, tc.src)

			if diff := cmp.Diff(tc.dst, tc.exp,
				cmp.Exporter(func(t reflect.Type) bool { return true }),
				cmp.Transformer("float32", func(f32 float32) uint32 {
					return math.Float32bits(f32)
				}),
				cmp.Transformer("float64", func(f64 float64) uint64 {
					return math.Float64bits(f64)
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
