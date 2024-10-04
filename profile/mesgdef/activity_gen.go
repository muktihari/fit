// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// Activity is a Activity message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Activity struct {
	Timestamp      time.Time
	LocalTimestamp time.Time // timestamp epoch expressed in local time, used to convert activity timestamps to local time
	TotalTimerTime uint32    // Scale: 1000; Units: s; Exclude pauses
	NumSessions    uint16
	Type           typedef.Activity
	Event          typedef.Event
	EventType      typedef.EventType
	EventGroup     uint8

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewActivity creates new Activity struct based on given mesg.
// If mesg is nil, it will return Activity with all fields being set to its corresponding invalid value.
func NewActivity(mesg *proto.Message) *Activity {
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
		unknownFields = sliceutil.Clone(unknownFields)
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Activity{
		Timestamp:      datetime.ToTime(vals[253].Uint32()),
		TotalTimerTime: vals[0].Uint32(),
		NumSessions:    vals[1].Uint16(),
		Type:           typedef.Activity(vals[2].Uint8()),
		Event:          typedef.Event(vals[3].Uint8()),
		EventType:      typedef.EventType(vals[4].Uint8()),
		LocalTimestamp: datetime.ToTime(vals[5].Uint32()),
		EventGroup:     vals[6].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Activity into proto.Message. If options is nil, default options will be used.
func (m *Activity) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumActivity}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.TotalTimerTime)
		fields = append(fields, field)
	}
	if m.NumSessions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.NumSessions)
		fields = append(fields, field)
	}
	if m.Type != typedef.ActivityInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if m.Event != typedef.EventInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Event))
		fields = append(fields, field)
	}
	if m.EventType != typedef.EventTypeInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.EventType))
		fields = append(fields, field)
	}
	if !m.LocalTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(uint32(m.LocalTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.EventGroup)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Activity) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// LocalTimestampUint32 returns LocalTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Activity) LocalTimestampUint32() uint32 { return datetime.ToUint32(m.LocalTimestamp) }

// TotalTimerTimeScaled return TotalTimerTime in its scaled value.
// If TotalTimerTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s; Exclude pauses
func (m *Activity) TotalTimerTimeScaled() float64 {
	if m.TotalTimerTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalTimerTime)/1000 - 0
}

// SetTimestamp sets Timestamp value.
func (m *Activity) SetTimestamp(v time.Time) *Activity {
	m.Timestamp = v
	return m
}

// SetTotalTimerTime sets TotalTimerTime value.
//
// Scale: 1000; Units: s; Exclude pauses
func (m *Activity) SetTotalTimerTime(v uint32) *Activity {
	m.TotalTimerTime = v
	return m
}

// SetTotalTimerTimeScaled is similar to SetTotalTimerTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s; Exclude pauses
func (m *Activity) SetTotalTimerTimeScaled(v float64) *Activity {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalTimerTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalTimerTime = uint32(unscaled)
	return m
}

// SetNumSessions sets NumSessions value.
func (m *Activity) SetNumSessions(v uint16) *Activity {
	m.NumSessions = v
	return m
}

// SetType sets Type value.
func (m *Activity) SetType(v typedef.Activity) *Activity {
	m.Type = v
	return m
}

// SetEvent sets Event value.
func (m *Activity) SetEvent(v typedef.Event) *Activity {
	m.Event = v
	return m
}

// SetEventType sets EventType value.
func (m *Activity) SetEventType(v typedef.EventType) *Activity {
	m.EventType = v
	return m
}

// SetLocalTimestamp sets LocalTimestamp value.
//
// timestamp epoch expressed in local time, used to convert activity timestamps to local time
func (m *Activity) SetLocalTimestamp(v time.Time) *Activity {
	m.LocalTimestamp = v
	return m
}

// SetEventGroup sets EventGroup value.
func (m *Activity) SetEventGroup(v uint8) *Activity {
	m.EventGroup = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Activity) SetUnknownFields(unknownFields ...proto.Field) *Activity {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *Activity) SetDeveloperFields(developerFields ...proto.DeveloperField) *Activity {
	m.DeveloperFields = developerFields
	return m
}
