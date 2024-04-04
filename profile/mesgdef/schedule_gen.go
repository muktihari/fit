// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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

// Schedule is a Schedule message.
type Schedule struct {
	TimeCreated   time.Time // Corresponds to file_id of scheduled workout / course.
	ScheduledTime time.Time
	SerialNumber  uint32               // Corresponds to file_id of scheduled workout / course.
	Manufacturer  typedef.Manufacturer // Corresponds to file_id of scheduled workout / course.
	Product       uint16               // Corresponds to file_id of scheduled workout / course.
	Type          typedef.Schedule
	Completed     bool // TRUE if this activity has been started

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSchedule creates new Schedule struct based on given mesg.
// If mesg is nil, it will return Schedule with all fields being set to its corresponding invalid value.
func NewSchedule(mesg *proto.Message) *Schedule {
	vals := [7]proto.Value{}

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

	return &Schedule{
		TimeCreated:   datetime.ToTime(vals[3].Uint32()),
		ScheduledTime: datetime.ToTime(vals[6].Uint32()),
		SerialNumber:  vals[2].Uint32z(),
		Manufacturer:  typedef.Manufacturer(vals[0].Uint16()),
		Product:       vals[1].Uint16(),
		Type:          typedef.Schedule(vals[5].Uint8()),
		Completed:     vals[4].Bool(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Schedule into proto.Message. If options is nil, default options will be used.
func (m *Schedule) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumSchedule}

	if datetime.ToUint32(m.TimeCreated) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(datetime.ToUint32(m.TimeCreated))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.ScheduledTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(datetime.ToUint32(m.ScheduledTime))
		fields = append(fields, field)
	}
	if uint32(m.SerialNumber) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.SerialNumber)
		fields = append(fields, field)
	}
	if uint16(m.Manufacturer) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.Manufacturer))
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.Product)
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if m.Completed != false {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Bool(m.Completed)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimeCreated sets Schedule value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetTimeCreated(v time.Time) *Schedule {
	m.TimeCreated = v
	return m
}

// SetScheduledTime sets Schedule value.
func (m *Schedule) SetScheduledTime(v time.Time) *Schedule {
	m.ScheduledTime = v
	return m
}

// SetSerialNumber sets Schedule value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetSerialNumber(v uint32) *Schedule {
	m.SerialNumber = v
	return m
}

// SetManufacturer sets Schedule value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetManufacturer(v typedef.Manufacturer) *Schedule {
	m.Manufacturer = v
	return m
}

// SetProduct sets Schedule value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetProduct(v uint16) *Schedule {
	m.Product = v
	return m
}

// SetType sets Schedule value.
func (m *Schedule) SetType(v typedef.Schedule) *Schedule {
	m.Type = v
	return m
}

// SetCompleted sets Schedule value.
//
// TRUE if this activity has been started
func (m *Schedule) SetCompleted(v bool) *Schedule {
	m.Completed = v
	return m
}

// SetDeveloperFields Schedule's DeveloperFields.
func (m *Schedule) SetDeveloperFields(developerFields ...proto.DeveloperField) *Schedule {
	m.DeveloperFields = developerFields
	return m
}
