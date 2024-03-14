// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

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

// HsaGyroscopeData is a HsaGyroscopeData message.
type HsaGyroscopeData struct {
	Timestamp        time.Time // Units: s
	GyroX            []int16   // Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
	GyroY            []int16   // Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
	GyroZ            []int16   // Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
	Timestamp32K     uint32    // Units: 1/32768 s; 32 kHz timestamp
	TimestampMs      uint16    // Units: ms; Millisecond resolution of the timestamp
	SamplingInterval uint16    // Units: 1/32768 s; Sampling Interval in 32 kHz timescale

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaGyroscopeData creates new HsaGyroscopeData struct based on given mesg.
// If mesg is nil, it will return HsaGyroscopeData with all fields being set to its corresponding invalid value.
func NewHsaGyroscopeData(mesg *proto.Message) *HsaGyroscopeData {
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

	return &HsaGyroscopeData{
		Timestamp:        datetime.ToTime(vals[253]),
		GyroX:            typeconv.ToSliceSint16[int16](vals[2]),
		GyroY:            typeconv.ToSliceSint16[int16](vals[3]),
		GyroZ:            typeconv.ToSliceSint16[int16](vals[4]),
		Timestamp32K:     typeconv.ToUint32[uint32](vals[5]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		SamplingInterval: typeconv.ToUint16[uint16](vals[1]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaGyroscopeData into proto.Message.
func (m *HsaGyroscopeData) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumHsaGyroscopeData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.GyroX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.GyroX
		fields = append(fields, field)
	}
	if m.GyroY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.GyroY
		fields = append(fields, field)
	}
	if m.GyroZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.GyroZ
		fields = append(fields, field)
	}
	if m.Timestamp32K != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.Timestamp32K
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TimestampMs
		fields = append(fields, field)
	}
	if m.SamplingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.SamplingInterval
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GyroXScaled return GyroX in its scaled value [Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement].
//
// If GyroX value is invalid, nil will be returned.
func (m *HsaGyroscopeData) GyroXScaled() []float64 {
	if m.GyroX == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.GyroX, 28.57143, 0)
}

// GyroYScaled return GyroY in its scaled value [Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement].
//
// If GyroY value is invalid, nil will be returned.
func (m *HsaGyroscopeData) GyroYScaled() []float64 {
	if m.GyroY == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.GyroY, 28.57143, 0)
}

// GyroZScaled return GyroZ in its scaled value [Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement].
//
// If GyroZ value is invalid, nil will be returned.
func (m *HsaGyroscopeData) GyroZScaled() []float64 {
	if m.GyroZ == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.GyroZ, 28.57143, 0)
}

// SetTimestamp sets HsaGyroscopeData value.
//
// Units: s
func (m *HsaGyroscopeData) SetTimestamp(v time.Time) *HsaGyroscopeData {
	m.Timestamp = v
	return m
}

// SetGyroX sets HsaGyroscopeData value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) SetGyroX(v []int16) *HsaGyroscopeData {
	m.GyroX = v
	return m
}

// SetGyroY sets HsaGyroscopeData value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) SetGyroY(v []int16) *HsaGyroscopeData {
	m.GyroY = v
	return m
}

// SetGyroZ sets HsaGyroscopeData value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) SetGyroZ(v []int16) *HsaGyroscopeData {
	m.GyroZ = v
	return m
}

// SetTimestamp32K sets HsaGyroscopeData value.
//
// Units: 1/32768 s; 32 kHz timestamp
func (m *HsaGyroscopeData) SetTimestamp32K(v uint32) *HsaGyroscopeData {
	m.Timestamp32K = v
	return m
}

// SetTimestampMs sets HsaGyroscopeData value.
//
// Units: ms; Millisecond resolution of the timestamp
func (m *HsaGyroscopeData) SetTimestampMs(v uint16) *HsaGyroscopeData {
	m.TimestampMs = v
	return m
}

// SetSamplingInterval sets HsaGyroscopeData value.
//
// Units: 1/32768 s; Sampling Interval in 32 kHz timescale
func (m *HsaGyroscopeData) SetSamplingInterval(v uint16) *HsaGyroscopeData {
	m.SamplingInterval = v
	return m
}

// SetDeveloperFields HsaGyroscopeData's DeveloperFields.
func (m *HsaGyroscopeData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaGyroscopeData {
	m.DeveloperFields = developerFields
	return m
}
