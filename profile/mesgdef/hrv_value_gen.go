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

// HrvValue is a HrvValue message.
type HrvValue struct {
	Timestamp time.Time
	Value     uint16 // Scale: 128; Units: ms; 5 minute RMSSD

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrvValue creates new HrvValue struct based on given mesg.
// If mesg is nil, it will return HrvValue with all fields being set to its corresponding invalid value.
func NewHrvValue(mesg *proto.Message) *HrvValue {
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

	return &HrvValue{
		Timestamp: datetime.ToTime(vals[253]),
		Value:     typeconv.ToUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HrvValue into proto.Message.
func (m *HrvValue) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumHrvValue)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Value != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Value
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of HrvValue's valid fields.
func (m *HrvValue) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if m.Value != basetype.Uint16Invalid {
		size++
	}
	return size
}

// SetTimestamp sets HrvValue value.
func (m *HrvValue) SetTimestamp(v time.Time) *HrvValue {
	m.Timestamp = v
	return m
}

// SetValue sets HrvValue value.
//
// Scale: 128; Units: ms; 5 minute RMSSD
func (m *HrvValue) SetValue(v uint16) *HrvValue {
	m.Value = v
	return m
}

// SetDeveloperFields HrvValue's DeveloperFields.
func (m *HrvValue) SetDeveloperFields(developerFields ...proto.DeveloperField) *HrvValue {
	m.DeveloperFields = developerFields
	return m
}
