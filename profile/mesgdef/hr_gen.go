// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// Hr is a Hr message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Hr struct {
	Timestamp           time.Time
	FilteredBpm         []uint8  // Array: [N]; Units: bpm
	EventTimestamp      []uint32 // Array: [N]; Scale: 1024; Units: s
	EventTimestamp12    []byte   // Array: [N]; Units: s
	FractionalTimestamp uint16   // Scale: 32768; Units: s
	Time256             uint8    // Scale: 256; Units: s

	state [2]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewHr creates new Hr struct based on given mesg.
// If mesg is nil, it will return Hr with all fields being set to its corresponding invalid value.
func NewHr(mesg *proto.Message) *Hr {
	vals := [254]proto.Value{}

	var state [2]uint8
	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 10 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Hr{
		Timestamp:           datetime.ToTime(vals[253].Uint32()),
		FractionalTimestamp: vals[0].Uint16(),
		Time256:             vals[1].Uint8(),
		FilteredBpm:         vals[6].SliceUint8(),
		EventTimestamp:      vals[9].SliceUint32(),
		EventTimestamp12:    vals[10].SliceUint8(),

		state: state,

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Hr into proto.Message. If options is nil, default options will be used.
func (m *Hr) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHr}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		if expanded := m.IsExpandedField(0); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 0)
			field.Value = proto.Uint16(m.FractionalTimestamp)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.Time256 != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.Time256)
		fields = append(fields, field)
	}
	if m.FilteredBpm != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceUint8(m.FilteredBpm)
		fields = append(fields, field)
	}
	if m.EventTimestamp != nil {
		if expanded := m.IsExpandedField(9); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 9)
			field.Value = proto.SliceUint32(m.EventTimestamp)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.EventTimestamp12 != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.SliceUint8(m.EventTimestamp12)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Hr) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// FractionalTimestampScaled return FractionalTimestamp in its scaled value.
// If FractionalTimestamp value is invalid, float64 invalid value will be returned.
//
// Scale: 32768; Units: s
func (m *Hr) FractionalTimestampScaled() float64 {
	if m.FractionalTimestamp == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.FractionalTimestamp)/32768 - 0
}

// Time256Scaled return Time256 in its scaled value.
// If Time256 value is invalid, float64 invalid value will be returned.
//
// Scale: 256; Units: s
func (m *Hr) Time256Scaled() float64 {
	if m.Time256 == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Time256)/256 - 0
}

// EventTimestampScaled return EventTimestamp in its scaled value.
// If EventTimestamp value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1024; Units: s
func (m *Hr) EventTimestampScaled() []float64 {
	if m.EventTimestamp == nil {
		return nil
	}
	var vals = make([]float64, len(m.EventTimestamp))
	for i := range m.EventTimestamp {
		if m.EventTimestamp[i] == basetype.Uint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.EventTimestamp[i])/1024 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
func (m *Hr) SetTimestamp(v time.Time) *Hr {
	m.Timestamp = v
	return m
}

// SetFractionalTimestamp sets FractionalTimestamp value.
//
// Scale: 32768; Units: s
func (m *Hr) SetFractionalTimestamp(v uint16) *Hr {
	m.FractionalTimestamp = v
	return m
}

// SetFractionalTimestampScaled is similar to SetFractionalTimestamp except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 32768; Units: s
func (m *Hr) SetFractionalTimestampScaled(v float64) *Hr {
	unscaled := (v + 0) * 32768
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.FractionalTimestamp = uint16(basetype.Uint16Invalid)
		return m
	}
	m.FractionalTimestamp = uint16(unscaled)
	return m
}

// SetTime256 sets Time256 value.
//
// Scale: 256; Units: s
func (m *Hr) SetTime256(v uint8) *Hr {
	m.Time256 = v
	return m
}

// SetTime256Scaled is similar to SetTime256 except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 256; Units: s
func (m *Hr) SetTime256Scaled(v float64) *Hr {
	unscaled := (v + 0) * 256
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.Time256 = uint8(basetype.Uint8Invalid)
		return m
	}
	m.Time256 = uint8(unscaled)
	return m
}

// SetFilteredBpm sets FilteredBpm value.
//
// Array: [N]; Units: bpm
func (m *Hr) SetFilteredBpm(v []uint8) *Hr {
	m.FilteredBpm = v
	return m
}

// SetEventTimestamp sets EventTimestamp value.
//
// Array: [N]; Scale: 1024; Units: s
func (m *Hr) SetEventTimestamp(v []uint32) *Hr {
	m.EventTimestamp = v
	return m
}

// SetEventTimestampScaled is similar to SetEventTimestamp except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1024; Units: s
func (m *Hr) SetEventTimestampScaled(vs []float64) *Hr {
	if vs == nil {
		m.EventTimestamp = nil
		return m
	}
	m.EventTimestamp = make([]uint32, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1024
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
			m.EventTimestamp[i] = uint32(basetype.Uint32Invalid)
			continue
		}
		m.EventTimestamp[i] = uint32(unscaled)
	}
	return m
}

// SetEventTimestamp12 sets EventTimestamp12 value.
//
// Array: [N]; Units: s
func (m *Hr) SetEventTimestamp12(v []byte) *Hr {
	m.EventTimestamp12 = v
	return m
}

// SetDeveloperFields Hr's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Hr) SetUnknownFields(unknownFields ...proto.Field) *Hr {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Hr's DeveloperFields.
func (m *Hr) SetDeveloperFields(developerFields ...proto.DeveloperField) *Hr {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 0, 9.
func (m *Hr) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 0, 9:
	default:
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	m.state[pos] &^= bit
	if flag {
		m.state[pos] |= bit
	}
	return true
}

// IsExpandedField checks whether given fieldNum is a field generated through
// a component expansion. Eligible for field number: 0, 9.
func (m *Hr) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 10 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}
