// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// ClimbPro is a ClimbPro message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ClimbPro struct {
	Timestamp     time.Time // Units: s
	PositionLat   int32     // Units: semicircles
	PositionLong  int32     // Units: semicircles
	CurrentDist   float32   // Units: m
	ClimbNumber   uint16
	ClimbProEvent typedef.ClimbProEvent
	ClimbCategory uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewClimbPro creates new ClimbPro struct based on given mesg.
// If mesg is nil, it will return ClimbPro with all fields being set to its corresponding invalid value.
func NewClimbPro(mesg *proto.Message) *ClimbPro {
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

	return &ClimbPro{
		Timestamp:     datetime.ToTime(vals[253].Uint32()),
		PositionLat:   vals[0].Int32(),
		PositionLong:  vals[1].Int32(),
		ClimbProEvent: typedef.ClimbProEvent(vals[2].Uint8()),
		ClimbNumber:   vals[3].Uint16(),
		ClimbCategory: vals[4].Uint8(),
		CurrentDist:   vals[5].Float32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ClimbPro into proto.Message. If options is nil, default options will be used.
func (m *ClimbPro) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumClimbPro}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Int32(m.PositionLat)
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Int32(m.PositionLong)
		fields = append(fields, field)
	}
	if byte(m.ClimbProEvent) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.ClimbProEvent))
		fields = append(fields, field)
	}
	if m.ClimbNumber != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.ClimbNumber)
		fields = append(fields, field)
	}
	if m.ClimbCategory != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(m.ClimbCategory)
		fields = append(fields, field)
	}
	if math.Float32bits(m.CurrentDist) != basetype.Float32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Float32(m.CurrentDist)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *ClimbPro) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// PositionLatDegrees returns PositionLat in degrees instead of semicircles.
// If PositionLat value is invalid, float64 invalid value will be returned.
func (m *ClimbPro) PositionLatDegrees() float64 {
	if m.PositionLat == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return semicircles.ToDegrees(m.PositionLat)
}

// PositionLongDegrees returns PositionLong in degrees instead of semicircles.
// If PositionLong value is invalid, float64 invalid value will be returned.
func (m *ClimbPro) PositionLongDegrees() float64 {
	if m.PositionLong == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return semicircles.ToDegrees(m.PositionLong)
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *ClimbPro) SetTimestamp(v time.Time) *ClimbPro {
	m.Timestamp = v
	return m
}

// SetPositionLat sets PositionLat value.
//
// Units: semicircles
func (m *ClimbPro) SetPositionLat(v int32) *ClimbPro {
	m.PositionLat = v
	return m
}

// SetPositionLatDegrees is similar to SetPositionLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *ClimbPro) SetPositionLatDegrees(degrees float64) *ClimbPro {
	m.PositionLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetPositionLong sets PositionLong value.
//
// Units: semicircles
func (m *ClimbPro) SetPositionLong(v int32) *ClimbPro {
	m.PositionLong = v
	return m
}

// SetPositionLongDegrees is similar to SetPositionLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *ClimbPro) SetPositionLongDegrees(degrees float64) *ClimbPro {
	m.PositionLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetClimbProEvent sets ClimbProEvent value.
func (m *ClimbPro) SetClimbProEvent(v typedef.ClimbProEvent) *ClimbPro {
	m.ClimbProEvent = v
	return m
}

// SetClimbNumber sets ClimbNumber value.
func (m *ClimbPro) SetClimbNumber(v uint16) *ClimbPro {
	m.ClimbNumber = v
	return m
}

// SetClimbCategory sets ClimbCategory value.
func (m *ClimbPro) SetClimbCategory(v uint8) *ClimbPro {
	m.ClimbCategory = v
	return m
}

// SetCurrentDist sets CurrentDist value.
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
