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
	"math"
	"time"
)

// TankSummary is a TankSummary message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type TankSummary struct {
	Timestamp     time.Time            // Units: s
	Sensor        typedef.AntChannelId // Base: uint32z
	VolumeUsed    uint32               // Scale: 100; Units: L
	StartPressure uint16               // Scale: 100; Units: bar
	EndPressure   uint16               // Scale: 100; Units: bar

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewTankSummary creates new TankSummary struct based on given mesg.
// If mesg is nil, it will return TankSummary with all fields being set to its corresponding invalid value.
func NewTankSummary(mesg *proto.Message) *TankSummary {
	vals := [254]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &TankSummary{
		Timestamp:     datetime.ToTime(vals[253].Uint32()),
		Sensor:        typedef.AntChannelId(vals[0].Uint32z()),
		StartPressure: vals[1].Uint16(),
		EndPressure:   vals[2].Uint16(),
		VolumeUsed:    vals[3].Uint32(),

		UnknownFields:   unknownFields,
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

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumTankSummary}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Sensor != typedef.AntChannelIdInvalid {
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

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

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
	return float64(m.StartPressure)/100 - 0
}

// EndPressureScaled return EndPressure in its scaled value.
// If EndPressure value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: bar
func (m *TankSummary) EndPressureScaled() float64 {
	if m.EndPressure == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.EndPressure)/100 - 0
}

// VolumeUsedScaled return VolumeUsed in its scaled value.
// If VolumeUsed value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: L
func (m *TankSummary) VolumeUsedScaled() float64 {
	if m.VolumeUsed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.VolumeUsed)/100 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *TankSummary) SetTimestamp(v time.Time) *TankSummary {
	m.Timestamp = v
	return m
}

// SetSensor sets Sensor value.
//
// Base: uint32z
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
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.StartPressure = uint16(basetype.Uint16Invalid)
		return m
	}
	m.StartPressure = uint16(unscaled)
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
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.EndPressure = uint16(basetype.Uint16Invalid)
		return m
	}
	m.EndPressure = uint16(unscaled)
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
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.VolumeUsed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.VolumeUsed = uint32(unscaled)
	return m
}

// SetDeveloperFields TankSummary's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *TankSummary) SetUnknownFields(unknownFields ...proto.Field) *TankSummary {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields TankSummary's DeveloperFields.
func (m *TankSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *TankSummary {
	m.DeveloperFields = developerFields
	return m
}
