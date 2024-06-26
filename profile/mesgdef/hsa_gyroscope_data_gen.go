// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// HsaGyroscopeData is a HsaGyroscopeData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
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

	return &HsaGyroscopeData{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		TimestampMs:      vals[0].Uint16(),
		SamplingInterval: vals[1].Uint16(),
		GyroX:            vals[2].SliceInt16(),
		GyroY:            vals[3].SliceInt16(),
		GyroZ:            vals[4].SliceInt16(),
		Timestamp32K:     vals[5].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaGyroscopeData into proto.Message. If options is nil, default options will be used.
func (m *HsaGyroscopeData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHsaGyroscopeData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.SamplingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.SamplingInterval)
		fields = append(fields, field)
	}
	if m.GyroX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceInt16(m.GyroX)
		fields = append(fields, field)
	}
	if m.GyroY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceInt16(m.GyroY)
		fields = append(fields, field)
	}
	if m.GyroZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceInt16(m.GyroZ)
		fields = append(fields, field)
	}
	if m.Timestamp32K != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.Timestamp32K)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaGyroscopeData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// GyroXScaled return GyroX in its scaled value.
// If GyroX value is invalid, nil will be returned.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) GyroXScaled() []float64 {
	if m.GyroX == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.GyroX, 28.57143, 0)
}

// GyroYScaled return GyroY in its scaled value.
// If GyroY value is invalid, nil will be returned.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) GyroYScaled() []float64 {
	if m.GyroY == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.GyroY, 28.57143, 0)
}

// GyroZScaled return GyroZ in its scaled value.
// If GyroZ value is invalid, nil will be returned.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) GyroZScaled() []float64 {
	if m.GyroZ == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.GyroZ, 28.57143, 0)
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *HsaGyroscopeData) SetTimestamp(v time.Time) *HsaGyroscopeData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond resolution of the timestamp
func (m *HsaGyroscopeData) SetTimestampMs(v uint16) *HsaGyroscopeData {
	m.TimestampMs = v
	return m
}

// SetSamplingInterval sets SamplingInterval value.
//
// Units: 1/32768 s; Sampling Interval in 32 kHz timescale
func (m *HsaGyroscopeData) SetSamplingInterval(v uint16) *HsaGyroscopeData {
	m.SamplingInterval = v
	return m
}

// SetGyroX sets GyroX value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) SetGyroX(v []int16) *HsaGyroscopeData {
	m.GyroX = v
	return m
}

// SetGyroXScaled is similar to SetGyroX except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) SetGyroXScaled(vs []float64) *HsaGyroscopeData {
	m.GyroX = scaleoffset.DiscardSlice[int16](vs, 28.57143, 0)
	return m
}

// SetGyroY sets GyroY value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) SetGyroY(v []int16) *HsaGyroscopeData {
	m.GyroY = v
	return m
}

// SetGyroYScaled is similar to SetGyroY except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) SetGyroYScaled(vs []float64) *HsaGyroscopeData {
	m.GyroY = scaleoffset.DiscardSlice[int16](vs, 28.57143, 0)
	return m
}

// SetGyroZ sets GyroZ value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) SetGyroZ(v []int16) *HsaGyroscopeData {
	m.GyroZ = v
	return m
}

// SetGyroZScaled is similar to SetGyroZ except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) SetGyroZScaled(vs []float64) *HsaGyroscopeData {
	m.GyroZ = scaleoffset.DiscardSlice[int16](vs, 28.57143, 0)
	return m
}

// SetTimestamp32K sets Timestamp32K value.
//
// Units: 1/32768 s; 32 kHz timestamp
func (m *HsaGyroscopeData) SetTimestamp32K(v uint32) *HsaGyroscopeData {
	m.Timestamp32K = v
	return m
}

// SetDeveloperFields HsaGyroscopeData's DeveloperFields.
func (m *HsaGyroscopeData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaGyroscopeData {
	m.DeveloperFields = developerFields
	return m
}
