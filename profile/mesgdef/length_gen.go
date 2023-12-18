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

// Length is a Length message.
type Length struct {
	MessageIndex               typedef.MessageIndex
	Timestamp                  typedef.DateTime
	Event                      typedef.Event
	EventType                  typedef.EventType
	StartTime                  typedef.DateTime
	TotalElapsedTime           uint32             // Scale: 1000; Units: s;
	TotalTimerTime             uint32             // Scale: 1000; Units: s;
	TotalStrokes               uint16             // Units: strokes;
	AvgSpeed                   uint16             // Scale: 1000; Units: m/s;
	SwimStroke                 typedef.SwimStroke // Units: swim_stroke;
	AvgSwimmingCadence         uint8              // Units: strokes/min;
	EventGroup                 uint8
	TotalCalories              uint16 // Units: kcal;
	LengthType                 typedef.LengthType
	PlayerScore                uint16
	OpponentScore              uint16
	StrokeCount                []uint16 // Array: [N]; Units: counts; stroke_type enum used as the index
	ZoneCount                  []uint16 // Array: [N]; Units: counts; zone number used as the index
	EnhancedAvgRespirationRate uint16   // Scale: 100; Units: Breaths/min;
	EnhancedMaxRespirationRate uint16   // Scale: 100; Units: Breaths/min;
	AvgRespirationRate         uint8
	MaxRespirationRate         uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewLength creates new Length struct based on given mesg. If mesg is nil or mesg.Num is not equal to Length mesg number, it will return nil.
func NewLength(mesg proto.Message) *Length {
	if mesg.Num != typedef.MesgNumLength {
		return nil
	}

	vals := [255]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &Length{
		MessageIndex:               typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Timestamp:                  typeconv.ToUint32[typedef.DateTime](vals[253]),
		Event:                      typeconv.ToEnum[typedef.Event](vals[0]),
		EventType:                  typeconv.ToEnum[typedef.EventType](vals[1]),
		StartTime:                  typeconv.ToUint32[typedef.DateTime](vals[2]),
		TotalElapsedTime:           typeconv.ToUint32[uint32](vals[3]),
		TotalTimerTime:             typeconv.ToUint32[uint32](vals[4]),
		TotalStrokes:               typeconv.ToUint16[uint16](vals[5]),
		AvgSpeed:                   typeconv.ToUint16[uint16](vals[6]),
		SwimStroke:                 typeconv.ToEnum[typedef.SwimStroke](vals[7]),
		AvgSwimmingCadence:         typeconv.ToUint8[uint8](vals[9]),
		EventGroup:                 typeconv.ToUint8[uint8](vals[10]),
		TotalCalories:              typeconv.ToUint16[uint16](vals[11]),
		LengthType:                 typeconv.ToEnum[typedef.LengthType](vals[12]),
		PlayerScore:                typeconv.ToUint16[uint16](vals[18]),
		OpponentScore:              typeconv.ToUint16[uint16](vals[19]),
		StrokeCount:                typeconv.ToSliceUint16[uint16](vals[20]),
		ZoneCount:                  typeconv.ToSliceUint16[uint16](vals[21]),
		EnhancedAvgRespirationRate: typeconv.ToUint16[uint16](vals[22]),
		EnhancedMaxRespirationRate: typeconv.ToUint16[uint16](vals[23]),
		AvgRespirationRate:         typeconv.ToUint8[uint8](vals[24]),
		MaxRespirationRate:         typeconv.ToUint8[uint8](vals[25]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts Length into proto.Message.
func (m *Length) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumLength)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
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
	if typeconv.ToUint32[uint32](m.StartTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = typeconv.ToUint32[uint32](m.StartTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalElapsedTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.TotalElapsedTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.TotalTimerTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalStrokes != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.TotalStrokes
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgSpeed != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.AvgSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.SwimStroke) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToEnum[byte](m.SwimStroke)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgSwimmingCadence != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.AvgSwimmingCadence
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.EventGroup
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalCalories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.TotalCalories
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.LengthType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = typeconv.ToEnum[byte](m.LengthType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.PlayerScore != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = m.PlayerScore
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = m.OpponentScore
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.StrokeCount != nil {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = m.StrokeCount
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ZoneCount != nil {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = m.ZoneCount
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EnhancedAvgRespirationRate != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = m.EnhancedAvgRespirationRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EnhancedMaxRespirationRate != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.EnhancedMaxRespirationRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgRespirationRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = m.AvgRespirationRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxRespirationRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = m.MaxRespirationRate
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Length's valid fields.
func (m *Length) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Event) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.EventType) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToUint32[uint32](m.StartTime) != basetype.Uint32Invalid {
		size++
	}
	if m.TotalElapsedTime != basetype.Uint32Invalid {
		size++
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		size++
	}
	if m.TotalStrokes != basetype.Uint16Invalid {
		size++
	}
	if m.AvgSpeed != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.SwimStroke) != basetype.EnumInvalid {
		size++
	}
	if m.AvgSwimmingCadence != basetype.Uint8Invalid {
		size++
	}
	if m.EventGroup != basetype.Uint8Invalid {
		size++
	}
	if m.TotalCalories != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.LengthType) != basetype.EnumInvalid {
		size++
	}
	if m.PlayerScore != basetype.Uint16Invalid {
		size++
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		size++
	}
	if m.StrokeCount != nil {
		size++
	}
	if m.ZoneCount != nil {
		size++
	}
	if m.EnhancedAvgRespirationRate != basetype.Uint16Invalid {
		size++
	}
	if m.EnhancedMaxRespirationRate != basetype.Uint16Invalid {
		size++
	}
	if m.AvgRespirationRate != basetype.Uint8Invalid {
		size++
	}
	if m.MaxRespirationRate != basetype.Uint8Invalid {
		size++
	}
	return size
}
