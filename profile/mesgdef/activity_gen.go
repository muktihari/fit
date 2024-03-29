// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// Activity is a Activity message.
type Activity struct {
	Timestamp      time.Time
	LocalTimestamp time.Time // timestamp epoch expressed in local time, used to convert activity timestamps to local time
	TotalTimerTime uint32    // Scale: 1000; Units: s; Exclude pauses
	NumSessions    uint16
	Type           typedef.Activity
	Event          typedef.Event
	EventType      typedef.EventType
	EventGroup     uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewActivity creates new Activity struct based on given mesg.
// If mesg is nil, it will return Activity with all fields being set to its corresponding invalid value.
func NewActivity(mesg *proto.Message) *Activity {
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

	return &Activity{
		Timestamp:      datetime.ToTime(vals[253]),
		LocalTimestamp: datetime.ToTime(vals[5]),
		TotalTimerTime: typeconv.ToUint32[uint32](vals[0]),
		NumSessions:    typeconv.ToUint16[uint16](vals[1]),
		Type:           typeconv.ToEnum[typedef.Activity](vals[2]),
		Event:          typeconv.ToEnum[typedef.Event](vals[3]),
		EventType:      typeconv.ToEnum[typedef.EventType](vals[4]),
		EventGroup:     typeconv.ToUint8[uint8](vals[6]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Activity into proto.Message.
func (m *Activity) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumActivity)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.LocalTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = datetime.ToUint32(m.LocalTimestamp)
		fields = append(fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TotalTimerTime
		fields = append(fields, field)
	}
	if m.NumSessions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.NumSessions
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = byte(m.Type)
		fields = append(fields, field)
	}
	if byte(m.Event) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = byte(m.Event)
		fields = append(fields, field)
	}
	if byte(m.EventType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = byte(m.EventType)
		fields = append(fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.EventGroup
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TotalTimerTimeScaled return TotalTimerTime in its scaled value [Scale: 1000; Units: s; Exclude pauses].
//
// If TotalTimerTime value is invalid, float64 invalid value will be returned.
func (m *Activity) TotalTimerTimeScaled() float64 {
	if m.TotalTimerTime == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.TotalTimerTime, 1000, 0)
}

// SetTimestamp sets Activity value.
func (m *Activity) SetTimestamp(v time.Time) *Activity {
	m.Timestamp = v
	return m
}

// SetLocalTimestamp sets Activity value.
//
// timestamp epoch expressed in local time, used to convert activity timestamps to local time
func (m *Activity) SetLocalTimestamp(v time.Time) *Activity {
	m.LocalTimestamp = v
	return m
}

// SetTotalTimerTime sets Activity value.
//
// Scale: 1000; Units: s; Exclude pauses
func (m *Activity) SetTotalTimerTime(v uint32) *Activity {
	m.TotalTimerTime = v
	return m
}

// SetNumSessions sets Activity value.
func (m *Activity) SetNumSessions(v uint16) *Activity {
	m.NumSessions = v
	return m
}

// SetType sets Activity value.
func (m *Activity) SetType(v typedef.Activity) *Activity {
	m.Type = v
	return m
}

// SetEvent sets Activity value.
func (m *Activity) SetEvent(v typedef.Event) *Activity {
	m.Event = v
	return m
}

// SetEventType sets Activity value.
func (m *Activity) SetEventType(v typedef.EventType) *Activity {
	m.EventType = v
	return m
}

// SetEventGroup sets Activity value.
func (m *Activity) SetEventGroup(v uint8) *Activity {
	m.EventGroup = v
	return m
}

// SetDeveloperFields Activity's DeveloperFields.
func (m *Activity) SetDeveloperFields(developerFields ...proto.DeveloperField) *Activity {
	m.DeveloperFields = developerFields
	return m
}
