// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// OneDSensorCalibration is a OneDSensorCalibration message.
type OneDSensorCalibration struct {
	Timestamp          typedef.DateTime   // Units: s; Whole second part of the timestamp
	SensorType         typedef.SensorType // Indicates which sensor the calibration is for
	CalibrationFactor  uint32             // Calibration factor used to convert from raw ADC value to degrees, g, etc.
	CalibrationDivisor uint32             // Units: counts; Calibration factor divisor
	LevelShift         uint32             // Level shift value used to shift the ADC value back into range
	OffsetCal          int32              // Internal Calibration factor

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewOneDSensorCalibration creates new OneDSensorCalibration struct based on given mesg. If mesg is nil or mesg.Num is not equal to OneDSensorCalibration mesg number, it will return nil.
func NewOneDSensorCalibration(mesg proto.Message) *OneDSensorCalibration {
	if mesg.Num != typedef.MesgNumOneDSensorCalibration {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* SensorType */
		1:   nil, /* CalibrationFactor */
		2:   nil, /* CalibrationDivisor */
		3:   nil, /* LevelShift */
		4:   nil, /* OffsetCal */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &OneDSensorCalibration{
		Timestamp:          typeconv.ToUint32[typedef.DateTime](vals[253]),
		SensorType:         typeconv.ToEnum[typedef.SensorType](vals[0]),
		CalibrationFactor:  typeconv.ToUint32[uint32](vals[1]),
		CalibrationDivisor: typeconv.ToUint32[uint32](vals[2]),
		LevelShift:         typeconv.ToUint32[uint32](vals[3]),
		OffsetCal:          typeconv.ToSint32[int32](vals[4]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to OneDSensorCalibration mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumOneDSensorCalibration)
func (m OneDSensorCalibration) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumOneDSensorCalibration {
		return
	}

	vals := [...]any{
		253: m.Timestamp,
		0:   m.SensorType,
		1:   m.CalibrationFactor,
		2:   m.CalibrationDivisor,
		3:   m.LevelShift,
		4:   m.OffsetCal,
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		field.Value = vals[field.Num]
	}

	mesg.DeveloperFields = m.DeveloperFields
}
