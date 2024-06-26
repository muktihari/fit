// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/profile/typedef"
)

type Accumulator struct {
	AccumulatedValues []AccumulatedValue // use slice over map since len(values) is relatively small
}

func NewAccumulator() *Accumulator {
	return &Accumulator{} // No need to make AccumulatedValues as it will be created on append anyway.
}

func (a *Accumulator) Collect(mesgNum typedef.MesgNum, destFieldNum byte, value uint32) {
	for i := range a.AccumulatedValues {
		field := &a.AccumulatedValues[i]
		if field.MesgNum == mesgNum && field.DestFieldNum == destFieldNum {
			field.Value = value
			field.Last = value
			return
		}
	}
	a.AccumulatedValues = append(a.AccumulatedValues, AccumulatedValue{
		MesgNum:      mesgNum,
		DestFieldNum: destFieldNum,
		Value:        value,
		Last:         value,
	})
}

func (a *Accumulator) Accumulate(mesgNum typedef.MesgNum, destFieldNum byte, value uint32, bits byte) uint32 {
	for i := range a.AccumulatedValues {
		av := &a.AccumulatedValues[i]
		if av.MesgNum == mesgNum && av.DestFieldNum == destFieldNum {
			return av.Accumulate(value, bits)
		}
	}
	return value
}

func (a *Accumulator) Reset() { a.AccumulatedValues = a.AccumulatedValues[:0] }

type AccumulatedValue struct {
	MesgNum      typedef.MesgNum
	DestFieldNum byte
	Last         uint32
	Value        uint32
}

func (a *AccumulatedValue) Accumulate(value uint32, bits byte) uint32 {
	var mask uint32 = (1 << bits) - 1
	a.Value += (value - a.Last) & mask
	a.Last = value
	return a.Value
}
