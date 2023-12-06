// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// DiveApneaAlarm is a DiveApneaAlarm message.
type DiveApneaAlarm struct {
	MessageIndex     typedef.MessageIndex  // Index of the alarm
	Depth            uint32                // Scale: 1000; Units: m; Depth setting (m) for depth type alarms
	Time             int32                 // Units: s; Time setting (s) for time type alarms
	Enabled          bool                  // Enablement flag
	AlarmType        typedef.DiveAlarmType // Alarm type setting
	Sound            typedef.Tone          // Tone and Vibe setting for the alarm.
	DiveTypes        []typedef.SubSport    // Array: [N]; Dive types the alarm will trigger on
	Id               uint32                // Alarm ID
	PopupEnabled     bool                  // Show a visible pop-up for this alarm
	TriggerOnDescent bool                  // Trigger the alarm on descent
	TriggerOnAscent  bool                  // Trigger the alarm on ascent
	Repeating        bool                  // Repeat alarm each time threshold is crossed?
	Speed            int32                 // Scale: 1000; Units: mps; Ascent/descent rate (mps) setting for speed type alarms

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDiveApneaAlarm creates new DiveApneaAlarm struct based on given mesg. If mesg is nil or mesg.Num is not equal to DiveApneaAlarm mesg number, it will return nil.
func NewDiveApneaAlarm(mesg proto.Message) *DiveApneaAlarm {
	if mesg.Num != typedef.MesgNumDiveApneaAlarm {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		0:   basetype.Uint32Invalid, /* Depth */
		1:   basetype.Sint32Invalid, /* Time */
		2:   false,                  /* Enabled */
		3:   basetype.EnumInvalid,   /* AlarmType */
		4:   basetype.EnumInvalid,   /* Sound */
		5:   nil,                    /* DiveTypes */
		6:   basetype.Uint32Invalid, /* Id */
		7:   false,                  /* PopupEnabled */
		8:   false,                  /* TriggerOnDescent */
		9:   false,                  /* TriggerOnAscent */
		10:  false,                  /* Repeating */
		11:  basetype.Sint32Invalid, /* Speed */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &DiveApneaAlarm{
		MessageIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Depth:            typeconv.ToUint32[uint32](vals[0]),
		Time:             typeconv.ToSint32[int32](vals[1]),
		Enabled:          typeconv.ToBool[bool](vals[2]),
		AlarmType:        typeconv.ToEnum[typedef.DiveAlarmType](vals[3]),
		Sound:            typeconv.ToEnum[typedef.Tone](vals[4]),
		DiveTypes:        typeconv.ToSliceEnum[typedef.SubSport](vals[5]),
		Id:               typeconv.ToUint32[uint32](vals[6]),
		PopupEnabled:     typeconv.ToBool[bool](vals[7]),
		TriggerOnDescent: typeconv.ToBool[bool](vals[8]),
		TriggerOnAscent:  typeconv.ToBool[bool](vals[9]),
		Repeating:        typeconv.ToBool[bool](vals[10]),
		Speed:            typeconv.ToSint32[int32](vals[11]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DiveApneaAlarm mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDiveApneaAlarm)
func (m DiveApneaAlarm) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDiveApneaAlarm {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.Depth,
		1:   m.Time,
		2:   m.Enabled,
		3:   m.AlarmType,
		4:   m.Sound,
		5:   m.DiveTypes,
		6:   m.Id,
		7:   m.PopupEnabled,
		8:   m.TriggerOnDescent,
		9:   m.TriggerOnAscent,
		10:  m.Repeating,
		11:  m.Speed,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
