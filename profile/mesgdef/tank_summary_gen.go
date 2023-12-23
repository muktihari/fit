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

// TankSummary is a TankSummary message.
type TankSummary struct {
	Timestamp     time.Time // Units: s;
	Sensor        typedef.AntChannelId
	StartPressure uint16 // Scale: 100; Units: bar;
	EndPressure   uint16 // Scale: 100; Units: bar;
	VolumeUsed    uint32 // Scale: 100; Units: L;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTankSummary creates new TankSummary struct based on given mesg.
// If mesg is nil, it will return TankSummary with all fields being set to its corresponding invalid value.
func NewTankSummary(mesg *proto.Message) *TankSummary {
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

	return &TankSummary{
		Timestamp:     datetime.ToTime(vals[253]),
		Sensor:        typeconv.ToUint32z[typedef.AntChannelId](vals[0]),
		StartPressure: typeconv.ToUint16[uint16](vals[1]),
		EndPressure:   typeconv.ToUint16[uint16](vals[2]),
		VolumeUsed:    typeconv.ToUint32[uint32](vals[3]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts TankSummary into proto.Message.
func (m *TankSummary) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumTankSummary)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if typeconv.ToUint32z[uint32](m.Sensor) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToUint32z[uint32](m.Sensor)
		fields = append(fields, field)
	}
	if m.StartPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.StartPressure
		fields = append(fields, field)
	}
	if m.EndPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.EndPressure
		fields = append(fields, field)
	}
	if m.VolumeUsed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.VolumeUsed
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets TankSummary value.
//
// Units: s;
func (m *TankSummary) SetTimestamp(v time.Time) *TankSummary {
	m.Timestamp = v
	return m
}

// SetSensor sets TankSummary value.
func (m *TankSummary) SetSensor(v typedef.AntChannelId) *TankSummary {
	m.Sensor = v
	return m
}

// SetStartPressure sets TankSummary value.
//
// Scale: 100; Units: bar;
func (m *TankSummary) SetStartPressure(v uint16) *TankSummary {
	m.StartPressure = v
	return m
}

// SetEndPressure sets TankSummary value.
//
// Scale: 100; Units: bar;
func (m *TankSummary) SetEndPressure(v uint16) *TankSummary {
	m.EndPressure = v
	return m
}

// SetVolumeUsed sets TankSummary value.
//
// Scale: 100; Units: L;
func (m *TankSummary) SetVolumeUsed(v uint32) *TankSummary {
	m.VolumeUsed = v
	return m
}

// SetDeveloperFields TankSummary's DeveloperFields.
func (m *TankSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *TankSummary {
	m.DeveloperFields = developerFields
	return m
}
