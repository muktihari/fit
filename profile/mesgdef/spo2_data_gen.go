// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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

// Spo2Data is a Spo2Data message.
type Spo2Data struct {
	Timestamp         time.Time // Units: s
	ReadingSpo2       uint8     // Units: percent
	ReadingConfidence uint8
	Mode              typedef.Spo2MeasurementType // Mode when data was captured

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSpo2Data creates new Spo2Data struct based on given mesg.
// If mesg is nil, it will return Spo2Data with all fields being set to its corresponding invalid value.
func NewSpo2Data(mesg *proto.Message) *Spo2Data {
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

	return &Spo2Data{
		Timestamp:         datetime.ToTime(vals[253]),
		ReadingSpo2:       typeconv.ToUint8[uint8](vals[0]),
		ReadingConfidence: typeconv.ToUint8[uint8](vals[1]),
		Mode:              typeconv.ToEnum[typedef.Spo2MeasurementType](vals[2]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Spo2Data into proto.Message.
func (m *Spo2Data) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSpo2Data)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.ReadingSpo2 != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ReadingSpo2
		fields = append(fields, field)
	}
	if m.ReadingConfidence != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.ReadingConfidence
		fields = append(fields, field)
	}
	if byte(m.Mode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = byte(m.Mode)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets Spo2Data value.
//
// Units: s
func (m *Spo2Data) SetTimestamp(v time.Time) *Spo2Data {
	m.Timestamp = v
	return m
}

// SetReadingSpo2 sets Spo2Data value.
//
// Units: percent
func (m *Spo2Data) SetReadingSpo2(v uint8) *Spo2Data {
	m.ReadingSpo2 = v
	return m
}

// SetReadingConfidence sets Spo2Data value.
func (m *Spo2Data) SetReadingConfidence(v uint8) *Spo2Data {
	m.ReadingConfidence = v
	return m
}

// SetMode sets Spo2Data value.
//
// Mode when data was captured
func (m *Spo2Data) SetMode(v typedef.Spo2MeasurementType) *Spo2Data {
	m.Mode = v
	return m
}

// SetDeveloperFields Spo2Data's DeveloperFields.
func (m *Spo2Data) SetDeveloperFields(developerFields ...proto.DeveloperField) *Spo2Data {
	m.DeveloperFields = developerFields
	return m
}
