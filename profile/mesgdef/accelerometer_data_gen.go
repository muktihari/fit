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

// AccelerometerData is a AccelerometerData message.
type AccelerometerData struct {
	Timestamp                  time.Time // Units: s; Whole second part of the timestamp
	TimestampMs                uint16    // Units: ms; Millisecond part of the timestamp.
	SampleTimeOffset           []uint16  // Array: [N]; Units: ms; Each time in the array describes the time at which the accelerometer sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in accel_x and accel_y and accel_z
	AccelX                     []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	AccelY                     []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	AccelZ                     []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	CalibratedAccelX           []float32 // Array: [N]; Units: g; Calibrated accel reading
	CalibratedAccelY           []float32 // Array: [N]; Units: g; Calibrated accel reading
	CalibratedAccelZ           []float32 // Array: [N]; Units: g; Calibrated accel reading
	CompressedCalibratedAccelX []int16   // Array: [N]; Units: mG; Calibrated accel reading
	CompressedCalibratedAccelY []int16   // Array: [N]; Units: mG; Calibrated accel reading
	CompressedCalibratedAccelZ []int16   // Array: [N]; Units: mG; Calibrated accel reading

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAccelerometerData creates new AccelerometerData struct based on given mesg.
// If mesg is nil, it will return AccelerometerData with all fields being set to its corresponding invalid value.
func NewAccelerometerData(mesg *proto.Message) *AccelerometerData {
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

	return &AccelerometerData{
		Timestamp:                  datetime.ToTime(vals[253]),
		TimestampMs:                typeconv.ToUint16[uint16](vals[0]),
		SampleTimeOffset:           typeconv.ToSliceUint16[uint16](vals[1]),
		AccelX:                     typeconv.ToSliceUint16[uint16](vals[2]),
		AccelY:                     typeconv.ToSliceUint16[uint16](vals[3]),
		AccelZ:                     typeconv.ToSliceUint16[uint16](vals[4]),
		CalibratedAccelX:           typeconv.ToSliceFloat32[float32](vals[5]),
		CalibratedAccelY:           typeconv.ToSliceFloat32[float32](vals[6]),
		CalibratedAccelZ:           typeconv.ToSliceFloat32[float32](vals[7]),
		CompressedCalibratedAccelX: typeconv.ToSliceSint16[int16](vals[8]),
		CompressedCalibratedAccelY: typeconv.ToSliceSint16[int16](vals[9]),
		CompressedCalibratedAccelZ: typeconv.ToSliceSint16[int16](vals[10]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts AccelerometerData into proto.Message.
func (m *AccelerometerData) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumAccelerometerData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TimestampMs
		fields = append(fields, field)
	}
	if m.SampleTimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.SampleTimeOffset
		fields = append(fields, field)
	}
	if m.AccelX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.AccelX
		fields = append(fields, field)
	}
	if m.AccelY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.AccelY
		fields = append(fields, field)
	}
	if m.AccelZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.AccelZ
		fields = append(fields, field)
	}
	if m.CalibratedAccelX != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.CalibratedAccelX
		fields = append(fields, field)
	}
	if m.CalibratedAccelY != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.CalibratedAccelY
		fields = append(fields, field)
	}
	if m.CalibratedAccelZ != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.CalibratedAccelZ
		fields = append(fields, field)
	}
	if m.CompressedCalibratedAccelX != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.CompressedCalibratedAccelX
		fields = append(fields, field)
	}
	if m.CompressedCalibratedAccelY != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.CompressedCalibratedAccelY
		fields = append(fields, field)
	}
	if m.CompressedCalibratedAccelZ != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.CompressedCalibratedAccelZ
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets AccelerometerData value.
//
// Units: s; Whole second part of the timestamp
func (m *AccelerometerData) SetTimestamp(v time.Time) *AccelerometerData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets AccelerometerData value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *AccelerometerData) SetTimestampMs(v uint16) *AccelerometerData {
	m.TimestampMs = v
	return m
}

// SetSampleTimeOffset sets AccelerometerData value.
//
// Array: [N]; Units: ms; Each time in the array describes the time at which the accelerometer sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in accel_x and accel_y and accel_z
func (m *AccelerometerData) SetSampleTimeOffset(v []uint16) *AccelerometerData {
	m.SampleTimeOffset = v
	return m
}

// SetAccelX sets AccelerometerData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *AccelerometerData) SetAccelX(v []uint16) *AccelerometerData {
	m.AccelX = v
	return m
}

// SetAccelY sets AccelerometerData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *AccelerometerData) SetAccelY(v []uint16) *AccelerometerData {
	m.AccelY = v
	return m
}

// SetAccelZ sets AccelerometerData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *AccelerometerData) SetAccelZ(v []uint16) *AccelerometerData {
	m.AccelZ = v
	return m
}

// SetCalibratedAccelX sets AccelerometerData value.
//
// Array: [N]; Units: g; Calibrated accel reading
func (m *AccelerometerData) SetCalibratedAccelX(v []float32) *AccelerometerData {
	m.CalibratedAccelX = v
	return m
}

// SetCalibratedAccelY sets AccelerometerData value.
//
// Array: [N]; Units: g; Calibrated accel reading
func (m *AccelerometerData) SetCalibratedAccelY(v []float32) *AccelerometerData {
	m.CalibratedAccelY = v
	return m
}

// SetCalibratedAccelZ sets AccelerometerData value.
//
// Array: [N]; Units: g; Calibrated accel reading
func (m *AccelerometerData) SetCalibratedAccelZ(v []float32) *AccelerometerData {
	m.CalibratedAccelZ = v
	return m
}

// SetCompressedCalibratedAccelX sets AccelerometerData value.
//
// Array: [N]; Units: mG; Calibrated accel reading
func (m *AccelerometerData) SetCompressedCalibratedAccelX(v []int16) *AccelerometerData {
	m.CompressedCalibratedAccelX = v
	return m
}

// SetCompressedCalibratedAccelY sets AccelerometerData value.
//
// Array: [N]; Units: mG; Calibrated accel reading
func (m *AccelerometerData) SetCompressedCalibratedAccelY(v []int16) *AccelerometerData {
	m.CompressedCalibratedAccelY = v
	return m
}

// SetCompressedCalibratedAccelZ sets AccelerometerData value.
//
// Array: [N]; Units: mG; Calibrated accel reading
func (m *AccelerometerData) SetCompressedCalibratedAccelZ(v []int16) *AccelerometerData {
	m.CompressedCalibratedAccelZ = v
	return m
}

// SetDeveloperFields AccelerometerData's DeveloperFields.
func (m *AccelerometerData) SetDeveloperFields(developerFields ...proto.DeveloperField) *AccelerometerData {
	m.DeveloperFields = developerFields
	return m
}
