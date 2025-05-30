// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
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
	Quality     []uint8  // Array: [N]; 1 = high confidence. 0 = low confidence. N/A when gap = 1
	Gap         []uint8  // Array: [N]; 1 = gap (time represents ms gap length). 0 = BBI data
	TimestampMs uint16   // Units: ms; Millisecond resolution of the timestamp

	state [1]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewRawBbi creates new RawBbi struct based on given mesg.
// If mesg is nil, it will return RawBbi with all fields being set to its corresponding invalid value.
func NewRawBbi(mesg *proto.Message) *RawBbi {
	m := new(RawBbi)
	m.Reset(mesg)
	return m
}

// Reset resets all RawBbi's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *RawBbi) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		state           [1]uint8
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		var n int
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				n++
			}
		}
		unknownFields = make([]proto.Field, 0, n)
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 5 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			if mesg.Fields[i].Num < 254 {
				vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
			}
		}
		developerFields = mesg.DeveloperFields
	}

	*m = RawBbi{
		Timestamp:   datetime.ToTime(vals[253].Uint32()),
		TimestampMs: vals[0].Uint16(),
		Data:        vals[1].SliceUint16(),
		Time:        vals[2].SliceUint16(),
		Quality:     vals[3].SliceUint8(),
		Gap:         vals[4].SliceUint8(),

		state: state,

		UnknownFields:   unknownFields,
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

	fields := make([]proto.Field, 0, 6)
	mesg := proto.Message{Num: typedef.MesgNumRawBbi}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
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
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.Quality != nil {
		if expanded := m.IsExpandedField(3); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 3)
			field.Value = proto.SliceUint8(m.Quality)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.Gap != nil {
		if expanded := m.IsExpandedField(4); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 4)
			field.Value = proto.SliceUint8(m.Gap)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}

	n := len(fields)
	mesg.Fields = make([]proto.Field, n+len(m.UnknownFields))
	copy(mesg.Fields[:n], fields)
	copy(mesg.Fields[n:], m.UnknownFields)

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
// Units: ms; Millisecond resolution of the timestamp
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
// Array: [N]; 1 = high confidence. 0 = low confidence. N/A when gap = 1
func (m *RawBbi) SetQuality(v []uint8) *RawBbi {
	m.Quality = v
	return m
}

// SetGap sets Gap value.
//
// Array: [N]; 1 = gap (time represents ms gap length). 0 = BBI data
func (m *RawBbi) SetGap(v []uint8) *RawBbi {
	m.Gap = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *RawBbi) SetUnknownFields(unknownFields ...proto.Field) *RawBbi {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
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
