// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// MagnetometerData is a MagnetometerData message.
type MagnetometerData struct {
	Timestamp        time.Time // Units: s; Whole second part of the timestamp
	SampleTimeOffset []uint16  // Array: [N]; Units: ms; Each time in the array describes the time at which the compass sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in cmps_x and cmps_y and cmps_z
	MagX             []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	MagY             []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	MagZ             []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	CalibratedMagX   []float32 // Array: [N]; Units: G; Calibrated Magnetometer reading
	CalibratedMagY   []float32 // Array: [N]; Units: G; Calibrated Magnetometer reading
	CalibratedMagZ   []float32 // Array: [N]; Units: G; Calibrated Magnetometer reading
	TimestampMs      uint16    // Units: ms; Millisecond part of the timestamp.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMagnetometerData creates new MagnetometerData struct based on given mesg.
// If mesg is nil, it will return MagnetometerData with all fields being set to its corresponding invalid value.
func NewMagnetometerData(mesg *proto.Message) *MagnetometerData {
	vals := [254]proto.Value{}

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

	return &MagnetometerData{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		SampleTimeOffset: vals[1].SliceUint16(),
		MagX:             vals[2].SliceUint16(),
		MagY:             vals[3].SliceUint16(),
		MagZ:             vals[4].SliceUint16(),
		CalibratedMagX:   vals[5].SliceFloat32(),
		CalibratedMagY:   vals[6].SliceFloat32(),
		CalibratedMagZ:   vals[7].SliceFloat32(),
		TimestampMs:      vals[0].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts MagnetometerData into proto.Message. If options is nil, default options will be used.
func (m *MagnetometerData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumMagnetometerData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.SampleTimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint16(m.SampleTimeOffset)
		fields = append(fields, field)
	}
	if m.MagX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint16(m.MagX)
		fields = append(fields, field)
	}
	if m.MagY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceUint16(m.MagY)
		fields = append(fields, field)
	}
	if m.MagZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint16(m.MagZ)
		fields = append(fields, field)
	}
	if m.CalibratedMagX != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceFloat32(m.CalibratedMagX)
		fields = append(fields, field)
	}
	if m.CalibratedMagY != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceFloat32(m.CalibratedMagY)
		fields = append(fields, field)
	}
	if m.CalibratedMagZ != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceFloat32(m.CalibratedMagZ)
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *MagnetometerData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets MagnetometerData value.
//
// Units: s; Whole second part of the timestamp
func (m *MagnetometerData) SetTimestamp(v time.Time) *MagnetometerData {
	m.Timestamp = v
	return m
}

// SetSampleTimeOffset sets MagnetometerData value.
//
// Array: [N]; Units: ms; Each time in the array describes the time at which the compass sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in cmps_x and cmps_y and cmps_z
func (m *MagnetometerData) SetSampleTimeOffset(v []uint16) *MagnetometerData {
	m.SampleTimeOffset = v
	return m
}

// SetMagX sets MagnetometerData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *MagnetometerData) SetMagX(v []uint16) *MagnetometerData {
	m.MagX = v
	return m
}

// SetMagY sets MagnetometerData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *MagnetometerData) SetMagY(v []uint16) *MagnetometerData {
	m.MagY = v
	return m
}

// SetMagZ sets MagnetometerData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *MagnetometerData) SetMagZ(v []uint16) *MagnetometerData {
	m.MagZ = v
	return m
}

// SetCalibratedMagX sets MagnetometerData value.
//
// Array: [N]; Units: G; Calibrated Magnetometer reading
func (m *MagnetometerData) SetCalibratedMagX(v []float32) *MagnetometerData {
	m.CalibratedMagX = v
	return m
}

// SetCalibratedMagY sets MagnetometerData value.
//
// Array: [N]; Units: G; Calibrated Magnetometer reading
func (m *MagnetometerData) SetCalibratedMagY(v []float32) *MagnetometerData {
	m.CalibratedMagY = v
	return m
}

// SetCalibratedMagZ sets MagnetometerData value.
//
// Array: [N]; Units: G; Calibrated Magnetometer reading
func (m *MagnetometerData) SetCalibratedMagZ(v []float32) *MagnetometerData {
	m.CalibratedMagZ = v
	return m
}

// SetTimestampMs sets MagnetometerData value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *MagnetometerData) SetTimestampMs(v uint16) *MagnetometerData {
	m.TimestampMs = v
	return m
}

// SetDeveloperFields MagnetometerData's DeveloperFields.
func (m *MagnetometerData) SetDeveloperFields(developerFields ...proto.DeveloperField) *MagnetometerData {
	m.DeveloperFields = developerFields
	return m
}
