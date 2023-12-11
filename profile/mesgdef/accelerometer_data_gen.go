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

// AccelerometerData is a AccelerometerData message.
type AccelerometerData struct {
	Timestamp                  typedef.DateTime // Units: s; Whole second part of the timestamp
	TimestampMs                uint16           // Units: ms; Millisecond part of the timestamp.
	SampleTimeOffset           []uint16         // Array: [N]; Units: ms; Each time in the array describes the time at which the accelerometer sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in accel_x and accel_y and accel_z
	AccelX                     []uint16         // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	AccelY                     []uint16         // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	AccelZ                     []uint16         // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	CalibratedAccelX           []float32        // Array: [N]; Units: g; Calibrated accel reading
	CalibratedAccelY           []float32        // Array: [N]; Units: g; Calibrated accel reading
	CalibratedAccelZ           []float32        // Array: [N]; Units: g; Calibrated accel reading
	CompressedCalibratedAccelX []int16          // Array: [N]; Units: mG; Calibrated accel reading
	CompressedCalibratedAccelY []int16          // Array: [N]; Units: mG; Calibrated accel reading
	CompressedCalibratedAccelZ []int16          // Array: [N]; Units: mG; Calibrated accel reading

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAccelerometerData creates new AccelerometerData struct based on given mesg. If mesg is nil or mesg.Num is not equal to AccelerometerData mesg number, it will return nil.
func NewAccelerometerData(mesg proto.Message) *AccelerometerData {
	if mesg.Num != typedef.MesgNumAccelerometerData {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* TimestampMs */
		1:   nil, /* SampleTimeOffset */
		2:   nil, /* AccelX */
		3:   nil, /* AccelY */
		4:   nil, /* AccelZ */
		5:   nil, /* CalibratedAccelX */
		6:   nil, /* CalibratedAccelY */
		7:   nil, /* CalibratedAccelZ */
		8:   nil, /* CompressedCalibratedAccelX */
		9:   nil, /* CompressedCalibratedAccelY */
		10:  nil, /* CompressedCalibratedAccelZ */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &AccelerometerData{
		Timestamp:                  typeconv.ToUint32[typedef.DateTime](vals[253]),
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

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to AccelerometerData mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumAccelerometerData)
func (m AccelerometerData) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumAccelerometerData {
		return
	}

	vals := [...]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.SampleTimeOffset,
		2:   m.AccelX,
		3:   m.AccelY,
		4:   m.AccelZ,
		5:   m.CalibratedAccelX,
		6:   m.CalibratedAccelY,
		7:   m.CalibratedAccelZ,
		8:   m.CompressedCalibratedAccelX,
		9:   m.CompressedCalibratedAccelY,
		10:  m.CompressedCalibratedAccelZ,
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
