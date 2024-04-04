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

// Totals is a Totals message.
type Totals struct {
	Timestamp    time.Time // Units: s
	TimerTime    uint32    // Units: s; Excludes pauses
	Distance     uint32    // Units: m
	Calories     uint32    // Units: kcal
	ElapsedTime  uint32    // Units: s; Includes pauses
	ActiveTime   uint32    // Units: s
	MessageIndex typedef.MessageIndex
	Sessions     uint16
	Sport        typedef.Sport
	SportIndex   uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTotals creates new Totals struct based on given mesg.
// If mesg is nil, it will return Totals with all fields being set to its corresponding invalid value.
func NewTotals(mesg *proto.Message) *Totals {
	vals := [255]proto.Value{}

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

	return &Totals{
		Timestamp:    datetime.ToTime(vals[253].Uint32()),
		TimerTime:    vals[0].Uint32(),
		Distance:     vals[1].Uint32(),
		Calories:     vals[2].Uint32(),
		ElapsedTime:  vals[4].Uint32(),
		ActiveTime:   vals[6].Uint32(),
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Sessions:     vals[5].Uint16(),
		Sport:        typedef.Sport(vals[3].Uint8()),
		SportIndex:   vals[9].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Totals into proto.Message. If options is nil, default options will be used.
func (m *Totals) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumTotals)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.TimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.TimerTime)
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.Distance)
		fields = append(fields, field)
	}
	if m.Calories != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.Calories)
		fields = append(fields, field)
	}
	if m.ElapsedTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.ElapsedTime)
		fields = append(fields, field)
	}
	if m.ActiveTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.ActiveTime)
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Sessions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.Sessions)
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if m.SportIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(m.SportIndex)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets Totals value.
//
// Units: s
func (m *Totals) SetTimestamp(v time.Time) *Totals {
	m.Timestamp = v
	return m
}

// SetTimerTime sets Totals value.
//
// Units: s; Excludes pauses
func (m *Totals) SetTimerTime(v uint32) *Totals {
	m.TimerTime = v
	return m
}

// SetDistance sets Totals value.
//
// Units: m
func (m *Totals) SetDistance(v uint32) *Totals {
	m.Distance = v
	return m
}

// SetCalories sets Totals value.
//
// Units: kcal
func (m *Totals) SetCalories(v uint32) *Totals {
	m.Calories = v
	return m
}

// SetElapsedTime sets Totals value.
//
// Units: s; Includes pauses
func (m *Totals) SetElapsedTime(v uint32) *Totals {
	m.ElapsedTime = v
	return m
}

// SetActiveTime sets Totals value.
//
// Units: s
func (m *Totals) SetActiveTime(v uint32) *Totals {
	m.ActiveTime = v
	return m
}

// SetMessageIndex sets Totals value.
func (m *Totals) SetMessageIndex(v typedef.MessageIndex) *Totals {
	m.MessageIndex = v
	return m
}

// SetSessions sets Totals value.
func (m *Totals) SetSessions(v uint16) *Totals {
	m.Sessions = v
	return m
}

// SetSport sets Totals value.
func (m *Totals) SetSport(v typedef.Sport) *Totals {
	m.Sport = v
	return m
}

// SetSportIndex sets Totals value.
func (m *Totals) SetSportIndex(v uint8) *Totals {
	m.SportIndex = v
	return m
}

// SetDeveloperFields Totals's DeveloperFields.
func (m *Totals) SetDeveloperFields(developerFields ...proto.DeveloperField) *Totals {
	m.DeveloperFields = developerFields
	return m
}
