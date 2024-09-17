package combiner

import (
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// accumulator is different than decoder's Accumulator, this works by just
// add the incoming value with the previous value, returning the result.
type accumulator struct{ values []value }

// Collect collects the value, if it already exist,
// update the existing value with this new value.
func (a *accumulator) Collect(mesgNum typedef.MesgNum, fieldNum byte, val proto.Value) {
	for i := range a.values {
		v := &a.values[i]
		if v.mesgNum == mesgNum && v.fieldNum == fieldNum {
			v.value = val
			v.last = val
			return
		}
	}
	a.values = append(a.values, value{mesgNum: mesgNum, fieldNum: fieldNum, value: val, last: val})
}

// Accumulate accumulates val with the previously collected value before.
// If not found, it will be collected instead, returning the result.
func (a *accumulator) Accumulate(mesgNum typedef.MesgNum, fieldNum byte, val proto.Value) proto.Value {
	for i := range a.values {
		v := &a.values[i]
		if v.mesgNum == mesgNum && v.fieldNum == fieldNum {
			v.last = sum(val, v.value) // val must be first argument in case we are handling slices
			return v.last
		}
	}
	a.values = append(a.values, value{mesgNum: mesgNum, fieldNum: fieldNum, value: val, last: val})
	return val
}

// SequenceCompleted marks the sequence as completed, last value of
// the previous sequence now becomes the new value.
func (a *accumulator) SequenceCompleted() {
	for i := range a.values {
		v := &a.values[i]
		v.value = v.last
	}
}

type value struct {
	mesgNum  typedef.MesgNum
	fieldNum byte
	value    proto.Value
	last     proto.Value
}

// sum sums v1 and v2, returning the result.
func sum(v1, v2 proto.Value) proto.Value {
	switch v1.Type() {
	case proto.TypeInt8:
		return proto.Int8(v1.Int8() + v2.Int8())
	case proto.TypeUint8:
		return proto.Uint8(v1.Uint8() + v2.Uint8())
	case proto.TypeInt16:
		return proto.Int16(v1.Int16() + v2.Int16())
	case proto.TypeUint16:
		return proto.Uint16(v1.Uint16() + v2.Uint16())
	case proto.TypeInt32:
		return proto.Int32(v1.Int32() + v2.Int32())
	case proto.TypeUint32:
		return proto.Uint32(v1.Uint32() + v2.Uint32())
	case proto.TypeInt64:
		return proto.Int64(v1.Int64() + v2.Int64())
	case proto.TypeUint64:
		return proto.Uint64(v1.Uint64() + v2.Uint64())
	case proto.TypeFloat32:
		return proto.Float32(v1.Float32() + v2.Float32())
	case proto.TypeFloat64:
		return proto.Float64(v1.Float64() + v2.Float64())
	case proto.TypeSliceInt8:
		return proto.SliceInt8(sumslice(v1.SliceInt8(), v2.SliceInt8()))
	case proto.TypeSliceUint8:
		return proto.SliceUint8(sumslice(v1.SliceUint8(), v2.SliceUint8()))
	case proto.TypeSliceInt16:
		return proto.SliceInt16(sumslice(v1.SliceInt16(), v2.SliceInt16()))
	case proto.TypeSliceUint16:
		return proto.SliceUint16(sumslice(v1.SliceUint16(), v2.SliceUint16()))
	case proto.TypeSliceInt32:
		return proto.SliceInt32(sumslice(v1.SliceInt32(), v2.SliceInt32()))
	case proto.TypeSliceUint32:
		return proto.SliceUint32(sumslice(v1.SliceUint32(), v2.SliceUint32()))
	case proto.TypeSliceInt64:
		return proto.SliceInt64(sumslice(v1.SliceInt64(), v2.SliceInt64()))
	case proto.TypeSliceUint64:
		return proto.SliceUint64(sumslice(v1.SliceUint64(), v2.SliceUint64()))
	case proto.TypeSliceFloat32:
		return proto.SliceFloat32(sumslice(v1.SliceFloat32(), v2.SliceFloat32()))
	case proto.TypeSliceFloat64:
		return proto.SliceFloat64(sumslice(v1.SliceFloat64(), v2.SliceFloat64()))
	}
	return v1
}

// sumslice sums slice values.
func sumslice[S []E, E scaleoffset.Numeric](v1, v2 S) S {
	if len(v1) >= len(v2) {
		for i := range v1 {
			if i == len(v2) {
				break
			}
			v1[i] += v2[i]
		}
		return v1
	}
	for i := range v2 {
		if i == len(v1) {
			v1 = append(v1, v2[i:]...)
			break
		}
		v1[i] += v2[i]
	}
	return v1
}
