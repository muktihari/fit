// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// RawBbi is a RawBbi message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type RawBbi struct {
	Timestamp   time.Time
	Data        []uint16 // Array: [N]; 1 bit for gap indicator, 1 bit for quality indicator, and 14 bits for Beat-to-Beat interval values in whole-integer millisecond resolution
	Time        []uint16 // Array: [N]; Units: ms; Array of millisecond times between beats
	Quality     []uint8  // Array: [N]
	Gap         []uint8  // Array: [N]
	TimestampMs uint16   // Units: ms; ms since last overnight_raw_bbi message

	state [1]uint8 // Used for tracking expanded fields.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewRawBbi creates new RawBbi struct based on given mesg.
// If mesg is nil, it will return RawBbi with all fields being set to its corresponding invalid value.
func NewRawBbi(mesg *proto.Message) *RawBbi {
	vals := [254]proto.Value{}

	var state [1]uint8
	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < 5 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &RawBbi{
		Timestamp:   datetime.ToTime(vals[253].Uint32()),
		TimestampMs: vals[0].Uint16(),
		Data:        vals[1].SliceUint16(),
		Time:        vals[2].SliceUint16(),
		Quality:     vals[3].SliceUint8(),
		Gap:         vals[4].SliceUint8(),

		state: state,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts RawBbi into proto.Message. If options is nil, default options will be used.
func (m *RawBbi) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumRawBbi}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.Data != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint16(m.Data)
		fields = append(fields, field)
	}
	if m.Time != nil {
		if expanded := m.IsExpandedField(2); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 2)
			field.Value = proto.SliceUint16(m.Time)
			field.IsExpandedField = m.IsExpandedField(2)
			fields = append(fields, field)
		}
	}
	if m.Quality != nil {
		if expanded := m.IsExpandedField(3); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 3)
			field.Value = proto.SliceUint8(m.Quality)
			field.IsExpandedField = m.IsExpandedField(3)
			fields = append(fields, field)
		}
	}
	if m.Gap != nil {
		if expanded := m.IsExpandedField(4); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 4)
			field.Value = proto.SliceUint8(m.Gap)
			field.IsExpandedField = m.IsExpandedField(4)
			fields = append(fields, field)
		}
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *RawBbi) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
func (m *RawBbi) SetTimestamp(v time.Time) *RawBbi {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; ms since last overnight_raw_bbi message
func (m *RawBbi) SetTimestampMs(v uint16) *RawBbi {
	m.TimestampMs = v
	return m
}

// SetData sets Data value.
//
// Array: [N]; 1 bit for gap indicator, 1 bit for quality indicator, and 14 bits for Beat-to-Beat interval values in whole-integer millisecond resolution
func (m *RawBbi) SetData(v []uint16) *RawBbi {
	m.Data = v
	return m
}

// SetTime sets Time value.
//
// Array: [N]; Units: ms; Array of millisecond times between beats
func (m *RawBbi) SetTime(v []uint16) *RawBbi {
	m.Time = v
	return m
}

// SetQuality sets Quality value.
//
// Array: [N]
func (m *RawBbi) SetQuality(v []uint8) *RawBbi {
	m.Quality = v
	return m
}

// SetGap sets Gap value.
//
// Array: [N]
func (m *RawBbi) SetGap(v []uint8) *RawBbi {
	m.Gap = v
	return m
}

// SetDeveloperFields RawBbi's DeveloperFields.
func (m *RawBbi) SetDeveloperFields(developerFields ...proto.DeveloperField) *RawBbi {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 2, 3, 4.
func (m *RawBbi) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 2, 3, 4:
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
// a component expansion. Eligible for field number: 2, 3, 4.
func (m *RawBbi) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 5 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}
