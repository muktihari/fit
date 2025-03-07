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

// AntTx is a AntTx message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type AntTx struct {
	Timestamp           time.Time // Units: s
	MesgData            []byte    // Array: [N]
	Data                []byte    // Array: [N]
	FractionalTimestamp uint16    // Scale: 32768; Units: s
	MesgId              byte
	ChannelNumber       uint8

	state [1]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewAntTx creates new AntTx struct based on given mesg.
// If mesg is nil, it will return AntTx with all fields being set to its corresponding invalid value.
func NewAntTx(mesg *proto.Message) *AntTx {
	m := new(AntTx)
	m.Reset(mesg)
	return m
}

// Reset resets all AntTx's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *AntTx) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		state           [1]uint8
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 5 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = AntTx{
		Timestamp:           datetime.ToTime(vals[253].Uint32()),
		FractionalTimestamp: vals[0].Uint16(),
		MesgId:              vals[1].Uint8(),
		MesgData:            vals[2].SliceUint8(),
		ChannelNumber:       vals[3].Uint8(),
		Data:                vals[4].SliceUint8(),

		state: state,

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts AntTx into proto.Message. If options is nil, default options will be used.
func (m *AntTx) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumAntTx}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.FractionalTimestamp)
		fields = append(fields, field)
	}
	if m.MesgId != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.MesgId)
		fields = append(fields, field)
	}
	if m.MesgData != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint8(m.MesgData)
		fields = append(fields, field)
	}
	if m.ChannelNumber != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(3); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 3)
			field.Value = proto.Uint8(m.ChannelNumber)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.Data != nil {
		if expanded := m.IsExpandedField(4); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 4)
			field.Value = proto.SliceUint8(m.Data)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *AntTx) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// FractionalTimestampScaled return FractionalTimestamp in its scaled value.
// If FractionalTimestamp value is invalid, float64 invalid value will be returned.
//
// Scale: 32768; Units: s
func (m *AntTx) FractionalTimestampScaled() float64 {
	if m.FractionalTimestamp == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.FractionalTimestamp)/32768 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *AntTx) SetTimestamp(v time.Time) *AntTx {
	m.Timestamp = v
	return m
}

// SetFractionalTimestamp sets FractionalTimestamp value.
//
// Scale: 32768; Units: s
func (m *AntTx) SetFractionalTimestamp(v uint16) *AntTx {
	m.FractionalTimestamp = v
	return m
}

// SetFractionalTimestampScaled is similar to SetFractionalTimestamp except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 32768; Units: s
func (m *AntTx) SetFractionalTimestampScaled(v float64) *AntTx {
	unscaled := (v + 0) * 32768
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.FractionalTimestamp = uint16(basetype.Uint16Invalid)
		return m
	}
	m.FractionalTimestamp = uint16(unscaled)
	return m
}

// SetMesgId sets MesgId value.
func (m *AntTx) SetMesgId(v byte) *AntTx {
	m.MesgId = v
	return m
}

// SetMesgData sets MesgData value.
//
// Array: [N]
func (m *AntTx) SetMesgData(v []byte) *AntTx {
	m.MesgData = v
	return m
}

// SetChannelNumber sets ChannelNumber value.
func (m *AntTx) SetChannelNumber(v uint8) *AntTx {
	m.ChannelNumber = v
	return m
}

// SetData sets Data value.
//
// Array: [N]
func (m *AntTx) SetData(v []byte) *AntTx {
	m.Data = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *AntTx) SetUnknownFields(unknownFields ...proto.Field) *AntTx {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *AntTx) SetDeveloperFields(developerFields ...proto.DeveloperField) *AntTx {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 3, 4.
func (m *AntTx) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 3, 4:
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
// a component expansion. Eligible for field number: 3, 4.
func (m *AntTx) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 5 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}
