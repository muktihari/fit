// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// Length is a Length message.
type Length struct {
	Timestamp                  time.Time
	StartTime                  time.Time
	StrokeCount                []uint16 // Array: [N]; Units: counts; stroke_type enum used as the index
	ZoneCount                  []uint16 // Array: [N]; Units: counts; zone number used as the index
	TotalElapsedTime           uint32   // Scale: 1000; Units: s
	TotalTimerTime             uint32   // Scale: 1000; Units: s
	MessageIndex               typedef.MessageIndex
	TotalStrokes               uint16 // Units: strokes
	AvgSpeed                   uint16 // Scale: 1000; Units: m/s
	TotalCalories              uint16 // Units: kcal
	PlayerScore                uint16
	OpponentScore              uint16
	EnhancedAvgRespirationRate uint16 // Scale: 100; Units: Breaths/min
	EnhancedMaxRespirationRate uint16 // Scale: 100; Units: Breaths/min
	Event                      typedef.Event
	EventType                  typedef.EventType
	SwimStroke                 typedef.SwimStroke // Units: swim_stroke
	AvgSwimmingCadence         uint8              // Units: strokes/min
	EventGroup                 uint8
	LengthType                 typedef.LengthType
	AvgRespirationRate         uint8
	MaxRespirationRate         uint8

	IsExpandedFields [24]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewLength creates new Length struct based on given mesg.
// If mesg is nil, it will return Length with all fields being set to its corresponding invalid value.
func NewLength(mesg *proto.Message) *Length {
	vals := [255]proto.Value{}
	isExpandedFields := [24]bool{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < byte(len(isExpandedFields)) {
				isExpandedFields[mesg.Fields[i].Num] = mesg.Fields[i].IsExpandedField
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Length{
		Timestamp:                  datetime.ToTime(vals[253].Uint32()),
		StartTime:                  datetime.ToTime(vals[2].Uint32()),
		StrokeCount:                vals[20].SliceUint16(),
		ZoneCount:                  vals[21].SliceUint16(),
		TotalElapsedTime:           vals[3].Uint32(),
		TotalTimerTime:             vals[4].Uint32(),
		MessageIndex:               typedef.MessageIndex(vals[254].Uint16()),
		TotalStrokes:               vals[5].Uint16(),
		AvgSpeed:                   vals[6].Uint16(),
		TotalCalories:              vals[11].Uint16(),
		PlayerScore:                vals[18].Uint16(),
		OpponentScore:              vals[19].Uint16(),
		EnhancedAvgRespirationRate: vals[22].Uint16(),
		EnhancedMaxRespirationRate: vals[23].Uint16(),
		Event:                      typedef.Event(vals[0].Uint8()),
		EventType:                  typedef.EventType(vals[1].Uint8()),
		SwimStroke:                 typedef.SwimStroke(vals[7].Uint8()),
		AvgSwimmingCadence:         vals[9].Uint8(),
		EventGroup:                 vals[10].Uint8(),
		LengthType:                 typedef.LengthType(vals[12].Uint8()),
		AvgRespirationRate:         vals[24].Uint8(),
		MaxRespirationRate:         vals[25].Uint8(),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Length into proto.Message. If options is nil, default options will be used.
func (m *Length) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumLength}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.StartTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(datetime.ToUint32(m.StartTime))
		fields = append(fields, field)
	}
	if m.StrokeCount != nil {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = proto.SliceUint16(m.StrokeCount)
		fields = append(fields, field)
	}
	if m.ZoneCount != nil {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.SliceUint16(m.ZoneCount)
		fields = append(fields, field)
	}
	if m.TotalElapsedTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.TotalElapsedTime)
		fields = append(fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.TotalTimerTime)
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.TotalStrokes != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.TotalStrokes)
		fields = append(fields, field)
	}
	if m.AvgSpeed != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint16(m.AvgSpeed)
		fields = append(fields, field)
	}
	if m.TotalCalories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint16(m.TotalCalories)
		fields = append(fields, field)
	}
	if m.PlayerScore != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = proto.Uint16(m.PlayerScore)
		fields = append(fields, field)
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.Uint16(m.OpponentScore)
		fields = append(fields, field)
	}
	if m.EnhancedAvgRespirationRate != basetype.Uint16Invalid && ((m.IsExpandedFields[22] && options.IncludeExpandedFields) || !m.IsExpandedFields[22]) {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint16(m.EnhancedAvgRespirationRate)
		field.IsExpandedField = m.IsExpandedFields[22]
		fields = append(fields, field)
	}
	if m.EnhancedMaxRespirationRate != basetype.Uint16Invalid && ((m.IsExpandedFields[23] && options.IncludeExpandedFields) || !m.IsExpandedFields[23]) {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Uint16(m.EnhancedMaxRespirationRate)
		field.IsExpandedField = m.IsExpandedFields[23]
		fields = append(fields, field)
	}
	if byte(m.Event) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Event))
		fields = append(fields, field)
	}
	if byte(m.EventType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.EventType))
		fields = append(fields, field)
	}
	if byte(m.SwimStroke) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.SwimStroke))
		fields = append(fields, field)
	}
	if m.AvgSwimmingCadence != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(m.AvgSwimmingCadence)
		fields = append(fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(m.EventGroup)
		fields = append(fields, field)
	}
	if byte(m.LengthType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(byte(m.LengthType))
		fields = append(fields, field)
	}
	if m.AvgRespirationRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = proto.Uint8(m.AvgRespirationRate)
		fields = append(fields, field)
	}
	if m.MaxRespirationRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint8(m.MaxRespirationRate)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TotalElapsedTimeScaled return TotalElapsedTime in its scaled value [Scale: 1000; Units: s].
//
// If TotalElapsedTime value is invalid, float64 invalid value will be returned.
func (m *Length) TotalElapsedTimeScaled() float64 {
	if m.TotalElapsedTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.TotalElapsedTime, 1000, 0)
}

// TotalTimerTimeScaled return TotalTimerTime in its scaled value [Scale: 1000; Units: s].
//
// If TotalTimerTime value is invalid, float64 invalid value will be returned.
func (m *Length) TotalTimerTimeScaled() float64 {
	if m.TotalTimerTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.TotalTimerTime, 1000, 0)
}

// AvgSpeedScaled return AvgSpeed in its scaled value [Scale: 1000; Units: m/s].
//
// If AvgSpeed value is invalid, float64 invalid value will be returned.
func (m *Length) AvgSpeedScaled() float64 {
	if m.AvgSpeed == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgSpeed, 1000, 0)
}

// EnhancedAvgRespirationRateScaled return EnhancedAvgRespirationRate in its scaled value [Scale: 100; Units: Breaths/min].
//
// If EnhancedAvgRespirationRate value is invalid, float64 invalid value will be returned.
func (m *Length) EnhancedAvgRespirationRateScaled() float64 {
	if m.EnhancedAvgRespirationRate == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.EnhancedAvgRespirationRate, 100, 0)
}

// EnhancedMaxRespirationRateScaled return EnhancedMaxRespirationRate in its scaled value [Scale: 100; Units: Breaths/min].
//
// If EnhancedMaxRespirationRate value is invalid, float64 invalid value will be returned.
func (m *Length) EnhancedMaxRespirationRateScaled() float64 {
	if m.EnhancedMaxRespirationRate == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.EnhancedMaxRespirationRate, 100, 0)
}

// SetTimestamp sets Length value.
func (m *Length) SetTimestamp(v time.Time) *Length {
	m.Timestamp = v
	return m
}

// SetStartTime sets Length value.
func (m *Length) SetStartTime(v time.Time) *Length {
	m.StartTime = v
	return m
}

// SetStrokeCount sets Length value.
//
// Array: [N]; Units: counts; stroke_type enum used as the index
func (m *Length) SetStrokeCount(v []uint16) *Length {
	m.StrokeCount = v
	return m
}

// SetZoneCount sets Length value.
//
// Array: [N]; Units: counts; zone number used as the index
func (m *Length) SetZoneCount(v []uint16) *Length {
	m.ZoneCount = v
	return m
}

// SetTotalElapsedTime sets Length value.
//
// Scale: 1000; Units: s
func (m *Length) SetTotalElapsedTime(v uint32) *Length {
	m.TotalElapsedTime = v
	return m
}

// SetTotalTimerTime sets Length value.
//
// Scale: 1000; Units: s
func (m *Length) SetTotalTimerTime(v uint32) *Length {
	m.TotalTimerTime = v
	return m
}

// SetMessageIndex sets Length value.
func (m *Length) SetMessageIndex(v typedef.MessageIndex) *Length {
	m.MessageIndex = v
	return m
}

// SetTotalStrokes sets Length value.
//
// Units: strokes
func (m *Length) SetTotalStrokes(v uint16) *Length {
	m.TotalStrokes = v
	return m
}

// SetAvgSpeed sets Length value.
//
// Scale: 1000; Units: m/s
func (m *Length) SetAvgSpeed(v uint16) *Length {
	m.AvgSpeed = v
	return m
}

// SetTotalCalories sets Length value.
//
// Units: kcal
func (m *Length) SetTotalCalories(v uint16) *Length {
	m.TotalCalories = v
	return m
}

// SetPlayerScore sets Length value.
func (m *Length) SetPlayerScore(v uint16) *Length {
	m.PlayerScore = v
	return m
}

// SetOpponentScore sets Length value.
func (m *Length) SetOpponentScore(v uint16) *Length {
	m.OpponentScore = v
	return m
}

// SetEnhancedAvgRespirationRate sets Length value.
//
// Scale: 100; Units: Breaths/min
func (m *Length) SetEnhancedAvgRespirationRate(v uint16) *Length {
	m.EnhancedAvgRespirationRate = v
	return m
}

// SetEnhancedMaxRespirationRate sets Length value.
//
// Scale: 100; Units: Breaths/min
func (m *Length) SetEnhancedMaxRespirationRate(v uint16) *Length {
	m.EnhancedMaxRespirationRate = v
	return m
}

// SetEvent sets Length value.
func (m *Length) SetEvent(v typedef.Event) *Length {
	m.Event = v
	return m
}

// SetEventType sets Length value.
func (m *Length) SetEventType(v typedef.EventType) *Length {
	m.EventType = v
	return m
}

// SetSwimStroke sets Length value.
//
// Units: swim_stroke
func (m *Length) SetSwimStroke(v typedef.SwimStroke) *Length {
	m.SwimStroke = v
	return m
}

// SetAvgSwimmingCadence sets Length value.
//
// Units: strokes/min
func (m *Length) SetAvgSwimmingCadence(v uint8) *Length {
	m.AvgSwimmingCadence = v
	return m
}

// SetEventGroup sets Length value.
func (m *Length) SetEventGroup(v uint8) *Length {
	m.EventGroup = v
	return m
}

// SetLengthType sets Length value.
func (m *Length) SetLengthType(v typedef.LengthType) *Length {
	m.LengthType = v
	return m
}

// SetAvgRespirationRate sets Length value.
func (m *Length) SetAvgRespirationRate(v uint8) *Length {
	m.AvgRespirationRate = v
	return m
}

// SetMaxRespirationRate sets Length value.
func (m *Length) SetMaxRespirationRate(v uint8) *Length {
	m.MaxRespirationRate = v
	return m
}

// SetDeveloperFields Length's DeveloperFields.
func (m *Length) SetDeveloperFields(developerFields ...proto.DeveloperField) *Length {
	m.DeveloperFields = developerFields
	return m
}
