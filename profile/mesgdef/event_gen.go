// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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

	vals := [254]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
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

// ToMesg converts Event into proto.Message.
func (m *Event) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumEvent)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = typeconv.ToUint32[uint32](m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Event) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.Event)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.EventType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToEnum[byte](m.EventType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Data16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Data16
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Data != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Data
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.EventGroup
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Score != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.Score
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.OpponentScore
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.FrontGearNum) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToUint8z[uint8](m.FrontGearNum)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.FrontGear) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = typeconv.ToUint8z[uint8](m.FrontGear)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.RearGearNum) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = typeconv.ToUint8z[uint8](m.RearGearNum)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.RearGear) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = typeconv.ToUint8z[uint8](m.RearGear)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8[uint8](m.DeviceIndex) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = typeconv.ToUint8[uint8](m.DeviceIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.ActivityType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = typeconv.ToEnum[byte](m.ActivityType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32[uint32](m.StartTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = typeconv.ToUint32[uint32](m.StartTimestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.RadarThreatLevelMax) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = typeconv.ToEnum[byte](m.RadarThreatLevelMax)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.RadarThreatCount != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = m.RadarThreatCount
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.RadarThreatAvgApproachSpeed != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.RadarThreatAvgApproachSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.RadarThreatMaxApproachSpeed != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = m.RadarThreatMaxApproachSpeed
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Event's valid fields.
func (m *Event) size() byte {
	var size byte
	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Event) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.EventType) != basetype.EnumInvalid {
		size++
	}
	if m.Data16 != basetype.Uint16Invalid {
		size++
	}
	if m.Data != basetype.Uint32Invalid {
		size++
	}
	if m.EventGroup != basetype.Uint8Invalid {
		size++
	}
	if m.Score != basetype.Uint16Invalid {
		size++
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.FrontGearNum) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.FrontGear) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.RearGearNum) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.RearGear) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8[uint8](m.DeviceIndex) != basetype.Uint8Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.ActivityType) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToUint32[uint32](m.StartTimestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.RadarThreatLevelMax) != basetype.EnumInvalid {
		size++
	}
	if m.RadarThreatCount != basetype.Uint8Invalid {
		size++
	}
	if m.RadarThreatAvgApproachSpeed != basetype.Uint8Invalid {
		size++
	}
	if m.RadarThreatMaxApproachSpeed != basetype.Uint8Invalid {
		size++
	}
	return size
}
