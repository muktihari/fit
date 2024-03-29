// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// HsaWristTemperatureData is a HsaWristTemperatureData message.
type HsaWristTemperatureData struct {
	Timestamp          time.Time // Units: s
	Value              []uint16  // Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaWristTemperatureData creates new HsaWristTemperatureData struct based on given mesg.
// If mesg is nil, it will return HsaWristTemperatureData with all fields being set to its corresponding invalid value.
func NewHsaWristTemperatureData(mesg *proto.Message) *HsaWristTemperatureData {
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

	return &HsaWristTemperatureData{
		Timestamp:          datetime.ToTime(vals[253]),
		Value:              typeconv.ToSliceUint16[uint16](vals[1]),
		ProcessingInterval: typeconv.ToUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaWristTemperatureData into proto.Message.
func (m *HsaWristTemperatureData) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumHsaWristTemperatureData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.Value != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Value
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ProcessingInterval
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// ValueScaled return Value in its scaled value [Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading].
//
// If Value value is invalid, nil will be returned.
func (m *HsaWristTemperatureData) ValueScaled() []float64 {
	if m.Value == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.Value, 1000, 0)
}

// SetTimestamp sets HsaWristTemperatureData value.
//
// Units: s
func (m *HsaWristTemperatureData) SetTimestamp(v time.Time) *HsaWristTemperatureData {
	m.Timestamp = v
	return m
}

// SetValue sets HsaWristTemperatureData value.
//
// Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading
func (m *HsaWristTemperatureData) SetValue(v []uint16) *HsaWristTemperatureData {
	m.Value = v
	return m
}

// SetProcessingInterval sets HsaWristTemperatureData value.
//
// Units: s; Processing interval length in seconds
func (m *HsaWristTemperatureData) SetProcessingInterval(v uint16) *HsaWristTemperatureData {
	m.ProcessingInterval = v
	return m
}

// SetDeveloperFields HsaWristTemperatureData's DeveloperFields.
func (m *HsaWristTemperatureData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaWristTemperatureData {
	m.DeveloperFields = developerFields
	return m
}
