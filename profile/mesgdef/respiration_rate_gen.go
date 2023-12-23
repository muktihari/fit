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

// RespirationRate is a RespirationRate message.
type RespirationRate struct {
	Timestamp       time.Time
	RespirationRate int16 // Scale: 100; Units: breaths/min; Breaths * 100 /min, -300 indicates invalid, -200 indicates large motion, -100 indicates off wrist

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewRespirationRate creates new RespirationRate struct based on given mesg.
// If mesg is nil, it will return RespirationRate with all fields being set to its corresponding invalid value.
func NewRespirationRate(mesg *proto.Message) *RespirationRate {
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

	return &RespirationRate{
		Timestamp:       datetime.ToTime(vals[253]),
		RespirationRate: typeconv.ToSint16[int16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts RespirationRate into proto.Message.
func (m *RespirationRate) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumRespirationRate)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.RespirationRate != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.RespirationRate
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets RespirationRate value.
func (m *RespirationRate) SetTimestamp(v time.Time) *RespirationRate {
	m.Timestamp = v
	return m
}

// SetRespirationRate sets RespirationRate value.
//
// Scale: 100; Units: breaths/min; Breaths * 100 /min, -300 indicates invalid, -200 indicates large motion, -100 indicates off wrist
func (m *RespirationRate) SetRespirationRate(v int16) *RespirationRate {
	m.RespirationRate = v
	return m
}

// SetDeveloperFields RespirationRate's DeveloperFields.
func (m *RespirationRate) SetDeveloperFields(developerFields ...proto.DeveloperField) *RespirationRate {
	m.DeveloperFields = developerFields
	return m
}
