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

// Event is a Event message.
type Event struct {
	Timestamp                   typedef.DateTime // Units: s;
	Event                       typedef.Event
	EventType                   typedef.EventType
	Data16                      uint16
	Data                        uint32
	EventGroup                  uint8
	Score                       uint16 // Do not populate directly. Autogenerated by decoder for sport_point subfield components
	OpponentScore               uint16 // Do not populate directly. Autogenerated by decoder for sport_point subfield components
	FrontGearNum                uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Front gear number. 1 is innermost.
	FrontGear                   uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of front teeth.
	RearGearNum                 uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Rear gear number. 1 is innermost.
	RearGear                    uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of rear teeth.
	DeviceIndex                 typedef.DeviceIndex
	ActivityType                typedef.ActivityType         // Activity Type associated with an auto_activity_detect event
	StartTimestamp              typedef.DateTime             // Units: s; Timestamp of when the event started
	RadarThreatLevelMax         typedef.RadarThreatLevelType // Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
	RadarThreatCount            uint8                        // Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
	RadarThreatAvgApproachSpeed uint8                        // Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
	RadarThreatMaxApproachSpeed uint8                        // Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewEvent creates new Event struct based on given mesg. If mesg is nil or mesg.Num is not equal to Event mesg number, it will return nil.
func NewEvent(mesg proto.Message) *Event {
	if mesg.Num != typedef.MesgNumEvent {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.EnumInvalid,   /* Event */
		1:   basetype.EnumInvalid,   /* EventType */
		2:   basetype.Uint16Invalid, /* Data16 */
		3:   basetype.Uint32Invalid, /* Data */
		4:   basetype.Uint8Invalid,  /* EventGroup */
		7:   basetype.Uint16Invalid, /* Score */
		8:   basetype.Uint16Invalid, /* OpponentScore */
		9:   basetype.Uint8zInvalid, /* FrontGearNum */
		10:  basetype.Uint8zInvalid, /* FrontGear */
		11:  basetype.Uint8zInvalid, /* RearGearNum */
		12:  basetype.Uint8zInvalid, /* RearGear */
		13:  basetype.Uint8Invalid,  /* DeviceIndex */
		14:  basetype.EnumInvalid,   /* ActivityType */
		15:  basetype.Uint32Invalid, /* StartTimestamp */
		21:  basetype.EnumInvalid,   /* RadarThreatLevelMax */
		22:  basetype.Uint8Invalid,  /* RadarThreatCount */
		23:  basetype.Uint8Invalid,  /* RadarThreatAvgApproachSpeed */
		24:  basetype.Uint8Invalid,  /* RadarThreatMaxApproachSpeed */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Event{
		Timestamp:                   typeconv.ToUint32[typedef.DateTime](vals[253]),
		Event:                       typeconv.ToEnum[typedef.Event](vals[0]),
		EventType:                   typeconv.ToEnum[typedef.EventType](vals[1]),
		Data16:                      typeconv.ToUint16[uint16](vals[2]),
		Data:                        typeconv.ToUint32[uint32](vals[3]),
		EventGroup:                  typeconv.ToUint8[uint8](vals[4]),
		Score:                       typeconv.ToUint16[uint16](vals[7]),
		OpponentScore:               typeconv.ToUint16[uint16](vals[8]),
		FrontGearNum:                typeconv.ToUint8z[uint8](vals[9]),
		FrontGear:                   typeconv.ToUint8z[uint8](vals[10]),
		RearGearNum:                 typeconv.ToUint8z[uint8](vals[11]),
		RearGear:                    typeconv.ToUint8z[uint8](vals[12]),
		DeviceIndex:                 typeconv.ToUint8[typedef.DeviceIndex](vals[13]),
		ActivityType:                typeconv.ToEnum[typedef.ActivityType](vals[14]),
		StartTimestamp:              typeconv.ToUint32[typedef.DateTime](vals[15]),
		RadarThreatLevelMax:         typeconv.ToEnum[typedef.RadarThreatLevelType](vals[21]),
		RadarThreatCount:            typeconv.ToUint8[uint8](vals[22]),
		RadarThreatAvgApproachSpeed: typeconv.ToUint8[uint8](vals[23]),
		RadarThreatMaxApproachSpeed: typeconv.ToUint8[uint8](vals[24]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Event mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumEvent)
func (m Event) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumEvent {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.Event,
		1:   m.EventType,
		2:   m.Data16,
		3:   m.Data,
		4:   m.EventGroup,
		7:   m.Score,
		8:   m.OpponentScore,
		9:   m.FrontGearNum,
		10:  m.FrontGear,
		11:  m.RearGearNum,
		12:  m.RearGear,
		13:  m.DeviceIndex,
		14:  m.ActivityType,
		15:  m.StartTimestamp,
		21:  m.RadarThreatLevelMax,
		22:  m.RadarThreatCount,
		23:  m.RadarThreatAvgApproachSpeed,
		24:  m.RadarThreatMaxApproachSpeed,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
