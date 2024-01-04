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

// ClimbPro is a ClimbPro message.
type ClimbPro struct {
	Timestamp     time.Time // Units: s
	PositionLat   int32     // Units: semicircles
	PositionLong  int32     // Units: semicircles
	ClimbProEvent typedef.ClimbProEvent
	ClimbNumber   uint16
	ClimbCategory uint8
	CurrentDist   float32 // Units: m

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewClimbPro creates new ClimbPro struct based on given mesg.
// If mesg is nil, it will return ClimbPro with all fields being set to its corresponding invalid value.
func NewClimbPro(mesg *proto.Message) *ClimbPro {
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

	return &ClimbPro{
		Timestamp:     datetime.ToTime(vals[253]),
		PositionLat:   typeconv.ToSint32[int32](vals[0]),
		PositionLong:  typeconv.ToSint32[int32](vals[1]),
		ClimbProEvent: typeconv.ToEnum[typedef.ClimbProEvent](vals[2]),
		ClimbNumber:   typeconv.ToUint16[uint16](vals[3]),
		ClimbCategory: typeconv.ToUint8[uint8](vals[4]),
		CurrentDist:   typeconv.ToFloat32[float32](vals[5]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ClimbPro into proto.Message.
func (m *ClimbPro) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumClimbPro)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.PositionLat
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.PositionLong
		fields = append(fields, field)
	}
	if byte(m.ClimbProEvent) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = byte(m.ClimbProEvent)
		fields = append(fields, field)
	}
	if m.ClimbNumber != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.ClimbNumber
		fields = append(fields, field)
	}
	if m.ClimbCategory != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.ClimbCategory
		fields = append(fields, field)
	}
	if typeconv.ToUint32[uint32](m.CurrentDist) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.CurrentDist
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets ClimbPro value.
//
// Units: s
func (m *ClimbPro) SetTimestamp(v time.Time) *ClimbPro {
	m.Timestamp = v
	return m
}

// SetPositionLat sets ClimbPro value.
//
// Units: semicircles
func (m *ClimbPro) SetPositionLat(v int32) *ClimbPro {
	m.PositionLat = v
	return m
}

// SetPositionLong sets ClimbPro value.
//
// Units: semicircles
func (m *ClimbPro) SetPositionLong(v int32) *ClimbPro {
	m.PositionLong = v
	return m
}

// SetClimbProEvent sets ClimbPro value.
func (m *ClimbPro) SetClimbProEvent(v typedef.ClimbProEvent) *ClimbPro {
	m.ClimbProEvent = v
	return m
}

// SetClimbNumber sets ClimbPro value.
func (m *ClimbPro) SetClimbNumber(v uint16) *ClimbPro {
	m.ClimbNumber = v
	return m
}

// SetClimbCategory sets ClimbPro value.
func (m *ClimbPro) SetClimbCategory(v uint8) *ClimbPro {
	m.ClimbCategory = v
	return m
}

// SetCurrentDist sets ClimbPro value.
//
// Units: m
func (m *ClimbPro) SetCurrentDist(v float32) *ClimbPro {
	m.CurrentDist = v
	return m
}

// SetDeveloperFields ClimbPro's DeveloperFields.
func (m *ClimbPro) SetDeveloperFields(developerFields ...proto.DeveloperField) *ClimbPro {
	m.DeveloperFields = developerFields
	return m
}
