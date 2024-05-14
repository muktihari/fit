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
	"math"
	"time"
)

// TankSummary is a TankSummary message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type TankSummary struct {
	Timestamp     time.Time // Units: s
	Sensor        typedef.AntChannelId
	VolumeUsed    uint32 // Scale: 100; Units: L
	StartPressure uint16 // Scale: 100; Units: bar
	EndPressure   uint16 // Scale: 100; Units: bar

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTankSummary creates new TankSummary struct based on given mesg.
// If mesg is nil, it will return TankSummary with all fields being set to its corresponding invalid value.
func NewTankSummary(mesg *proto.Message) *TankSummary {
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

	return &TankSummary{
		Timestamp:     datetime.ToTime(vals[253].Uint32()),
		Sensor:        typedef.AntChannelId(vals[0].Uint32z()),
		StartPressure: vals[1].Uint16(),
		EndPressure:   vals[2].Uint16(),
		VolumeUsed:    vals[3].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts TankSummary into proto.Message. If options is nil, default options will be used.
func (m *TankSummary) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumTankSummary}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if uint32(m.Sensor) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(uint32(m.Sensor))
		fields = append(fields, field)
	}
	if m.StartPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.StartPressure)
		fields = append(fields, field)
	}
	if m.EndPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.EndPressure)
		fields = append(fields, field)
	}
	if m.VolumeUsed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.VolumeUsed)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TankSummary) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// StartPressureScaled return StartPressure in its scaled value.
// If StartPressure value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: bar
func (m *TankSummary) StartPressureScaled() float64 {
	if m.StartPressure == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.StartPressure, 100, 0)
}

// EndPressureScaled return EndPressure in its scaled value.
// If EndPressure value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: bar
func (m *TankSummary) EndPressureScaled() float64 {
	if m.EndPressure == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.EndPressure, 100, 0)
}

// VolumeUsedScaled return VolumeUsed in its scaled value.
// If VolumeUsed value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: L
func (m *TankSummary) VolumeUsedScaled() float64 {
	if m.VolumeUsed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.VolumeUsed, 100, 0)
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *TankSummary) SetTimestamp(v time.Time) *TankSummary {
	m.Timestamp = v
	return m
}

// SetSensor sets Sensor value.
func (m *TankSummary) SetSensor(v typedef.AntChannelId) *TankSummary {
	m.Sensor = v
	return m
}

// SetStartPressure sets StartPressure value.
//
// Scale: 100; Units: bar
func (m *TankSummary) SetStartPressure(v uint16) *TankSummary {
	m.StartPressure = v
	return m
}

// SetStartPressureScaled is similar to SetStartPressure except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: bar
func (m *TankSummary) SetStartPressureScaled(v float64) *TankSummary {
	m.StartPressure = uint16(scaleoffset.Discard(v, 100, 0))
	return m
}

// SetEndPressure sets EndPressure value.
//
// Scale: 100; Units: bar
func (m *TankSummary) SetEndPressure(v uint16) *TankSummary {
	m.EndPressure = v
	return m
}

// SetEndPressureScaled is similar to SetEndPressure except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: bar
func (m *TankSummary) SetEndPressureScaled(v float64) *TankSummary {
	m.EndPressure = uint16(scaleoffset.Discard(v, 100, 0))
	return m
}

// SetVolumeUsed sets VolumeUsed value.
//
// Scale: 100; Units: L
func (m *TankSummary) SetVolumeUsed(v uint32) *TankSummary {
	m.VolumeUsed = v
	return m
}

// SetVolumeUsedScaled is similar to SetVolumeUsed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: L
func (m *TankSummary) SetVolumeUsedScaled(v float64) *TankSummary {
	m.VolumeUsed = uint32(scaleoffset.Discard(v, 100, 0))
	return m
}

// SetDeveloperFields TankSummary's DeveloperFields.
func (m *TankSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *TankSummary {
	m.DeveloperFields = developerFields
	return m
}
