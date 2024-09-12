// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/profile/typedef"
)

// Accumulator is value accumulator.
type Accumulator struct {
	values []value // use slice over map since len(values) is relatively small
}

// NewAccumulator creates new accumulator.
func NewAccumulator() *Accumulator {
	return &Accumulator{}
}

// Collect collects value, it will either append the value when not exist or replace existing one.
func (a *Accumulator) Collect(mesgNum typedef.MesgNum, destFieldNum byte, val uint32) {
	for i := range a.values {
		av := &a.values[i]
		if av.mesgNum == mesgNum && av.fieldNum == destFieldNum {
			av.value = val
			av.last = val
			return
		}
	}
	a.values = append(a.values, value{
		mesgNum:  mesgNum,
		fieldNum: destFieldNum,
		value:    val,
		last:     val,
	})
}

// Accumulate calculates the accumulated value and update accordingly. It returns the original value
// when the corresponding value does not exist.
func (a *Accumulator) Accumulate(mesgNum typedef.MesgNum, destFieldNum byte, val uint32, bits byte) uint32 {
	for i := range a.values {
		av := &a.values[i]
		if av.mesgNum == mesgNum && av.fieldNum == destFieldNum {
			return av.accumulate(val, bits)
		}
	}
	return val
}

// Reset resets the accumulator. Tt retains the underlying storage for use by
// future use to reduce memory allocs.
func (a *Accumulator) Reset() { a.values = a.values[:0] }

type value struct {
	mesgNum  typedef.MesgNum
	fieldNum byte
	last     uint32
	value    uint32
}

func (a *value) accumulate(val uint32, bits byte) uint32 {
	var mask uint32 = (1 << bits) - 1
	a.value += (val - a.last) & mask
	a.last = val
	return a.value
}
