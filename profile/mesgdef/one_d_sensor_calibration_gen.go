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

// OneDSensorCalibration is a OneDSensorCalibration message.
type OneDSensorCalibration struct {
	Timestamp          time.Time          // Units: s; Whole second part of the timestamp
	SensorType         typedef.SensorType // Indicates which sensor the calibration is for
	CalibrationFactor  uint32             // Calibration factor used to convert from raw ADC value to degrees, g, etc.
	CalibrationDivisor uint32             // Units: counts; Calibration factor divisor
	LevelShift         uint32             // Level shift value used to shift the ADC value back into range
	OffsetCal          int32              // Internal Calibration factor

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewOneDSensorCalibration creates new OneDSensorCalibration struct based on given mesg.
// If mesg is nil, it will return OneDSensorCalibration with all fields being set to its corresponding invalid value.
func NewOneDSensorCalibration(mesg *proto.Message) *OneDSensorCalibration {
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

	return &OneDSensorCalibration{
		Timestamp:          datetime.ToTime(vals[253]),
		SensorType:         typeconv.ToEnum[typedef.SensorType](vals[0]),
		CalibrationFactor:  typeconv.ToUint32[uint32](vals[1]),
		CalibrationDivisor: typeconv.ToUint32[uint32](vals[2]),
		LevelShift:         typeconv.ToUint32[uint32](vals[3]),
		OffsetCal:          typeconv.ToSint32[int32](vals[4]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts OneDSensorCalibration into proto.Message.
func (m *OneDSensorCalibration) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumOneDSensorCalibration)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if byte(m.SensorType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.SensorType)
		fields = append(fields, field)
	}
	if m.CalibrationFactor != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.CalibrationFactor
		fields = append(fields, field)
	}
	if m.CalibrationDivisor != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.CalibrationDivisor
		fields = append(fields, field)
	}
	if m.LevelShift != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.LevelShift
		fields = append(fields, field)
	}
	if m.OffsetCal != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.OffsetCal
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets OneDSensorCalibration value.
//
// Units: s; Whole second part of the timestamp
func (m *OneDSensorCalibration) SetTimestamp(v time.Time) *OneDSensorCalibration {
	m.Timestamp = v
	return m
}

// SetSensorType sets OneDSensorCalibration value.
//
// Indicates which sensor the calibration is for
func (m *OneDSensorCalibration) SetSensorType(v typedef.SensorType) *OneDSensorCalibration {
	m.SensorType = v
	return m
}

// SetCalibrationFactor sets OneDSensorCalibration value.
//
// Calibration factor used to convert from raw ADC value to degrees, g, etc.
func (m *OneDSensorCalibration) SetCalibrationFactor(v uint32) *OneDSensorCalibration {
	m.CalibrationFactor = v
	return m
}

// SetCalibrationDivisor sets OneDSensorCalibration value.
//
// Units: counts; Calibration factor divisor
func (m *OneDSensorCalibration) SetCalibrationDivisor(v uint32) *OneDSensorCalibration {
	m.CalibrationDivisor = v
	return m
}

// SetLevelShift sets OneDSensorCalibration value.
//
// Level shift value used to shift the ADC value back into range
func (m *OneDSensorCalibration) SetLevelShift(v uint32) *OneDSensorCalibration {
	m.LevelShift = v
	return m
}

// SetOffsetCal sets OneDSensorCalibration value.
//
// Internal Calibration factor
func (m *OneDSensorCalibration) SetOffsetCal(v int32) *OneDSensorCalibration {
	m.OffsetCal = v
	return m
}

// SetDeveloperFields OneDSensorCalibration's DeveloperFields.
func (m *OneDSensorCalibration) SetDeveloperFields(developerFields ...proto.DeveloperField) *OneDSensorCalibration {
	m.DeveloperFields = developerFields
	return m
}
