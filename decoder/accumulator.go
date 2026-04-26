// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// accumulator is value accumulator.
type accumulator struct {
	values []accuValue // use slice over map since len(values) is relatively small
}

// Collect collects value, it will either append the value when not exist or replace existing one.
func (a *accumulator) Collect(mesgNum typedef.MesgNum, fieldNum byte, val proto.Value) {
	switch val.Type() {
	case proto.TypeInt8:
		a.collect(mesgNum, fieldNum, uint64(val.Int8()))
	case proto.TypeUint8:
		a.collect(mesgNum, fieldNum, uint64(val.Uint8()))
	case proto.TypeInt16:
		a.collect(mesgNum, fieldNum, uint64(val.Int16()))
	case proto.TypeUint16:
		a.collect(mesgNum, fieldNum, uint64(val.Uint16()))
	case proto.TypeInt32:
		a.collect(mesgNum, fieldNum, uint64(val.Int32()))
	case proto.TypeUint32:
		a.collect(mesgNum, fieldNum, uint64(val.Uint32()))
	case proto.TypeInt64:
		a.collect(mesgNum, fieldNum, uint64(val.Int64()))
	case proto.TypeUint64:
		a.collect(mesgNum, fieldNum, uint64(val.Uint64()))
	case proto.TypeFloat32:
		a.collect(mesgNum, fieldNum, uint64(val.Float32()))
	case proto.TypeFloat64:
		a.collect(mesgNum, fieldNum, uint64(val.Float64()))
	case proto.TypeSliceInt8:
		vals := val.SliceInt8()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceUint8:
		vals := val.SliceUint8()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceInt16:
		vals := val.SliceInt16()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceUint16:
		vals := val.SliceUint16()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceInt32:
		vals := val.SliceInt32()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceUint32:
		vals := val.SliceUint32()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceInt64:
		vals := val.SliceInt64()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceUint64:
		vals := val.SliceUint64()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceFloat32:
		vals := val.SliceFloat32()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	case proto.TypeSliceFloat64:
		vals := val.SliceFloat64()
		if n := len(vals); n > 0 {
			a.collect(mesgNum, fieldNum, uint64(vals[n-1]))
		}
	}
}

// collect collects uint64 value, it will either append the value when not exist or replace existing one.
func (a *accumulator) collect(mesgNum typedef.MesgNum, fieldNum byte, val uint64) {
	for i := range a.values {
		av := &a.values[i]
		if av.mesgNum == mesgNum && av.fieldNum == fieldNum {
			av.value = val
			av.last = val
			return
		}
	}
	a.values = append(a.values, accuValue{
		mesgNum:  mesgNum,
		fieldNum: fieldNum,
		value:    val,
		last:     val,
	})
}

// Accumulate calculates the accumulated value and update it accordingly. If targeted value
// does not exist, it will be collected and the original value will be returned.
func (a *accumulator) Accumulate(mesgNum typedef.MesgNum, destFieldNum byte, val uint64, bits byte) uint64 {
	for i := range a.values {
		av := &a.values[i]
		if av.mesgNum == mesgNum && av.fieldNum == destFieldNum {
			var mask uint64 = (1 << bits) - 1
			av.value += (val - av.last) & mask
			av.last = val
			return av.value
		}
	}
	a.values = append(a.values, accuValue{
		mesgNum:  mesgNum,
		fieldNum: destFieldNum,
		value:    val,
		last:     val,
	})
	return val
}

// Reset resets the accumulator. It retains the underlying storage for use by
// future use to reduce memory allocs.
func (a *accumulator) Reset() { a.values = a.values[:0] }

type accuValue struct {
	mesgNum  typedef.MesgNum
	fieldNum byte
	last     uint64
	value    uint64
}
