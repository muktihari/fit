// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// AntRx is a AntRx message.
type AntRx struct {
	Timestamp           time.Time // Units: s
	MesgData            []byte    // Array: [N]
	Data                []byte    // Array: [N]
	FractionalTimestamp uint16    // Scale: 32768; Units: s
	MesgId              byte
	ChannelNumber       uint8

	IsExpandedFields [5]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAntRx creates new AntRx struct based on given mesg.
// If mesg is nil, it will return AntRx with all fields being set to its corresponding invalid value.
func NewAntRx(mesg *proto.Message) *AntRx {
	vals := [254]proto.Value{}
	isExpandedFields := [5]bool{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < byte(len(isExpandedFields)) {
				isExpandedFields[mesg.Fields[i].Num] = mesg.Fields[i].IsExpandedField
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &AntRx{
		Timestamp:           datetime.ToTime(vals[253].Uint32()),
		MesgData:            vals[2].SliceUint8(),
		Data:                vals[4].SliceUint8(),
		FractionalTimestamp: vals[0].Uint16(),
		MesgId:              vals[1].Uint8(),
		ChannelNumber:       vals[3].Uint8(),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts AntRx into proto.Message. If options is nil, default options will be used.
func (m *AntRx) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumAntRx}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.MesgData != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint8(m.MesgData)
		fields = append(fields, field)
	}
	if m.Data != nil && ((m.IsExpandedFields[4] && options.IncludeExpandedFields) || !m.IsExpandedFields[4]) {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint8(m.Data)
		field.IsExpandedField = m.IsExpandedFields[4]
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
	if m.ChannelNumber != basetype.Uint8Invalid && ((m.IsExpandedFields[3] && options.IncludeExpandedFields) || !m.IsExpandedFields[3]) {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.ChannelNumber)
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *AntRx) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// FractionalTimestampScaled return FractionalTimestamp in its scaled value [Scale: 32768; Units: s].
//
// If FractionalTimestamp value is invalid, float64 invalid value will be returned.
func (m *AntRx) FractionalTimestampScaled() float64 {
	if m.FractionalTimestamp == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.FractionalTimestamp, 32768, 0)
}

// SetTimestamp sets AntRx value.
//
// Units: s
func (m *AntRx) SetTimestamp(v time.Time) *AntRx {
	m.Timestamp = v
	return m
}

// SetMesgData sets AntRx value.
//
// Array: [N]
func (m *AntRx) SetMesgData(v []byte) *AntRx {
	m.MesgData = v
	return m
}

// SetData sets AntRx value.
//
// Array: [N]
func (m *AntRx) SetData(v []byte) *AntRx {
	m.Data = v
	return m
}

// SetFractionalTimestamp sets AntRx value.
//
// Scale: 32768; Units: s
func (m *AntRx) SetFractionalTimestamp(v uint16) *AntRx {
	m.FractionalTimestamp = v
	return m
}

// SetMesgId sets AntRx value.
func (m *AntRx) SetMesgId(v byte) *AntRx {
	m.MesgId = v
	return m
}

// SetChannelNumber sets AntRx value.
func (m *AntRx) SetChannelNumber(v uint8) *AntRx {
	m.ChannelNumber = v
	return m
}

// SetDeveloperFields AntRx's DeveloperFields.
func (m *AntRx) SetDeveloperFields(developerFields ...proto.DeveloperField) *AntRx {
	m.DeveloperFields = developerFields
	return m
}
