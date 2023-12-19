// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// AntTx is a AntTx message.
type AntTx struct {
	Timestamp           time.Time // Units: s;
	FractionalTimestamp uint16    // Scale: 32768; Units: s;
	MesgId              byte
	MesgData            []byte // Array: [N];
	ChannelNumber       uint8
	Data                []byte // Array: [N];

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAntTx creates new AntTx struct based on given mesg.
// If mesg is nil, it will return AntTx with all fields being set to its corresponding invalid value.
func NewAntTx(mesg *proto.Message) *AntTx {
	vals := [254]any{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &AntTx{
		Timestamp:           datetime.ToTime(vals[253]),
		FractionalTimestamp: typeconv.ToUint16[uint16](vals[0]),
		MesgId:              typeconv.ToByte[byte](vals[1]),
		MesgData:            typeconv.ToSliceByte[byte](vals[2]),
		ChannelNumber:       typeconv.ToUint8[uint8](vals[3]),
		Data:                typeconv.ToSliceByte[byte](vals[4]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts AntTx into proto.Message.
func (m *AntTx) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumAntTx)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.FractionalTimestamp
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MesgId != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.MesgId
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MesgData != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.MesgData
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ChannelNumber != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.ChannelNumber
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Data != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.Data
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of AntTx's valid fields.
func (m *AntTx) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		size++
	}
	if m.MesgId != basetype.ByteInvalid {
		size++
	}
	if m.MesgData != nil {
		size++
	}
	if m.ChannelNumber != basetype.Uint8Invalid {
		size++
	}
	if m.Data != nil {
		size++
	}
	return size
}

// SetTimestamp sets AntTx value.
//
// Units: s;
func (m *AntTx) SetTimestamp(v time.Time) *AntTx {
	m.Timestamp = v
	return m
}

// SetFractionalTimestamp sets AntTx value.
//
// Scale: 32768; Units: s;
func (m *AntTx) SetFractionalTimestamp(v uint16) *AntTx {
	m.FractionalTimestamp = v
	return m
}

// SetMesgId sets AntTx value.
func (m *AntTx) SetMesgId(v byte) *AntTx {
	m.MesgId = v
	return m
}

// SetMesgData sets AntTx value.
//
// Array: [N];
func (m *AntTx) SetMesgData(v []byte) *AntTx {
	m.MesgData = v
	return m
}

// SetChannelNumber sets AntTx value.
func (m *AntTx) SetChannelNumber(v uint8) *AntTx {
	m.ChannelNumber = v
	return m
}

// SetData sets AntTx value.
//
// Array: [N];
func (m *AntTx) SetData(v []byte) *AntTx {
	m.Data = v
	return m
}

// SetDeveloperFields AntTx's DeveloperFields.
func (m *AntTx) SetDeveloperFields(developerFields ...proto.DeveloperField) *AntTx {
	m.DeveloperFields = developerFields
	return m
}
