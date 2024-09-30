// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// SplitSummary is a SplitSummary message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SplitSummary struct {
	TotalTimerTime  uint32 // Scale: 1000; Units: s
	TotalDistance   uint32 // Scale: 100; Units: m
	AvgSpeed        uint32 // Scale: 1000; Units: m/s
	MaxSpeed        uint32 // Scale: 1000; Units: m/s
	AvgVertSpeed    int32  // Scale: 1000; Units: m/s
	TotalCalories   uint32 // Units: kcal
	TotalMovingTime uint32 // Scale: 1000; Units: s
	MessageIndex    typedef.MessageIndex
	NumSplits       uint16
	TotalAscent     uint16 // Units: m
	TotalDescent    uint16 // Units: m
	SplitType       typedef.SplitType
	AvgHeartRate    uint8 // Units: bpm
	MaxHeartRate    uint8 // Units: bpm

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSplitSummary creates new SplitSummary struct based on given mesg.
// If mesg is nil, it will return SplitSummary with all fields being set to its corresponding invalid value.
func NewSplitSummary(mesg *proto.Message) *SplitSummary {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &SplitSummary{
		MessageIndex:    typedef.MessageIndex(vals[254].Uint16()),
		SplitType:       typedef.SplitType(vals[0].Uint8()),
		NumSplits:       vals[3].Uint16(),
		TotalTimerTime:  vals[4].Uint32(),
		TotalDistance:   vals[5].Uint32(),
		AvgSpeed:        vals[6].Uint32(),
		MaxSpeed:        vals[7].Uint32(),
		TotalAscent:     vals[8].Uint16(),
		TotalDescent:    vals[9].Uint16(),
		AvgHeartRate:    vals[10].Uint8(),
		MaxHeartRate:    vals[11].Uint8(),
		AvgVertSpeed:    vals[12].Int32(),
		TotalCalories:   vals[13].Uint32(),
		TotalMovingTime: vals[77].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SplitSummary into proto.Message. If options is nil, default options will be used.
func (m *SplitSummary) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSplitSummary}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.SplitType != typedef.SplitTypeInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.SplitType))
		fields = append(fields, field)
	}
	if m.NumSplits != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.NumSplits)
		fields = append(fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.TotalTimerTime)
		fields = append(fields, field)
	}
	if m.TotalDistance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.TotalDistance)
		fields = append(fields, field)
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.AvgSpeed)
		fields = append(fields, field)
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint32(m.MaxSpeed)
		fields = append(fields, field)
	}
	if m.TotalAscent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint16(m.TotalAscent)
		fields = append(fields, field)
	}
	if m.TotalDescent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(m.TotalDescent)
		fields = append(fields, field)
	}
	if m.AvgHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(m.AvgHeartRate)
		fields = append(fields, field)
	}
	if m.MaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.MaxHeartRate)
		fields = append(fields, field)
	}
	if m.AvgVertSpeed != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Int32(m.AvgVertSpeed)
		fields = append(fields, field)
	}
	if m.TotalCalories != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint32(m.TotalCalories)
		fields = append(fields, field)
	}
	if m.TotalMovingTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 77)
		field.Value = proto.Uint32(m.TotalMovingTime)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TotalTimerTimeScaled return TotalTimerTime in its scaled value.
// If TotalTimerTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *SplitSummary) TotalTimerTimeScaled() float64 {
	if m.TotalTimerTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalTimerTime)/1000 - 0
}

// TotalDistanceScaled return TotalDistance in its scaled value.
// If TotalDistance value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *SplitSummary) TotalDistanceScaled() float64 {
	if m.TotalDistance == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalDistance)/100 - 0
}

// AvgSpeedScaled return AvgSpeed in its scaled value.
// If AvgSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) AvgSpeedScaled() float64 {
	if m.AvgSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.AvgSpeed)/1000 - 0
}

// MaxSpeedScaled return MaxSpeed in its scaled value.
// If MaxSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) MaxSpeedScaled() float64 {
	if m.MaxSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.MaxSpeed)/1000 - 0
}

// AvgVertSpeedScaled return AvgVertSpeed in its scaled value.
// If AvgVertSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) AvgVertSpeedScaled() float64 {
	if m.AvgVertSpeed == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.AvgVertSpeed)/1000 - 0
}

// TotalMovingTimeScaled return TotalMovingTime in its scaled value.
// If TotalMovingTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *SplitSummary) TotalMovingTimeScaled() float64 {
	if m.TotalMovingTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalMovingTime)/1000 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *SplitSummary) SetMessageIndex(v typedef.MessageIndex) *SplitSummary {
	m.MessageIndex = v
	return m
}

// SetSplitType sets SplitType value.
func (m *SplitSummary) SetSplitType(v typedef.SplitType) *SplitSummary {
	m.SplitType = v
	return m
}

// SetNumSplits sets NumSplits value.
func (m *SplitSummary) SetNumSplits(v uint16) *SplitSummary {
	m.NumSplits = v
	return m
}

// SetTotalTimerTime sets TotalTimerTime value.
//
// Scale: 1000; Units: s
func (m *SplitSummary) SetTotalTimerTime(v uint32) *SplitSummary {
	m.TotalTimerTime = v
	return m
}

// SetTotalTimerTimeScaled is similar to SetTotalTimerTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *SplitSummary) SetTotalTimerTimeScaled(v float64) *SplitSummary {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalTimerTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalTimerTime = uint32(unscaled)
	return m
}

// SetTotalDistance sets TotalDistance value.
//
// Scale: 100; Units: m
func (m *SplitSummary) SetTotalDistance(v uint32) *SplitSummary {
	m.TotalDistance = v
	return m
}

// SetTotalDistanceScaled is similar to SetTotalDistance except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *SplitSummary) SetTotalDistanceScaled(v float64) *SplitSummary {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalDistance = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalDistance = uint32(unscaled)
	return m
}

// SetAvgSpeed sets AvgSpeed value.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) SetAvgSpeed(v uint32) *SplitSummary {
	m.AvgSpeed = v
	return m
}

// SetAvgSpeedScaled is similar to SetAvgSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) SetAvgSpeedScaled(v float64) *SplitSummary {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.AvgSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.AvgSpeed = uint32(unscaled)
	return m
}

// SetMaxSpeed sets MaxSpeed value.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) SetMaxSpeed(v uint32) *SplitSummary {
	m.MaxSpeed = v
	return m
}

// SetMaxSpeedScaled is similar to SetMaxSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) SetMaxSpeedScaled(v float64) *SplitSummary {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.MaxSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.MaxSpeed = uint32(unscaled)
	return m
}

// SetTotalAscent sets TotalAscent value.
//
// Units: m
func (m *SplitSummary) SetTotalAscent(v uint16) *SplitSummary {
	m.TotalAscent = v
	return m
}

// SetTotalDescent sets TotalDescent value.
//
// Units: m
func (m *SplitSummary) SetTotalDescent(v uint16) *SplitSummary {
	m.TotalDescent = v
	return m
}

// SetAvgHeartRate sets AvgHeartRate value.
//
// Units: bpm
func (m *SplitSummary) SetAvgHeartRate(v uint8) *SplitSummary {
	m.AvgHeartRate = v
	return m
}

// SetMaxHeartRate sets MaxHeartRate value.
//
// Units: bpm
func (m *SplitSummary) SetMaxHeartRate(v uint8) *SplitSummary {
	m.MaxHeartRate = v
	return m
}

// SetAvgVertSpeed sets AvgVertSpeed value.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) SetAvgVertSpeed(v int32) *SplitSummary {
	m.AvgVertSpeed = v
	return m
}

// SetAvgVertSpeedScaled is similar to SetAvgVertSpeed except it accepts a scaled value.
// This method automatically converts the given value to its int32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *SplitSummary) SetAvgVertSpeedScaled(v float64) *SplitSummary {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint32Invalid) {
		m.AvgVertSpeed = int32(basetype.Sint32Invalid)
		return m
	}
	m.AvgVertSpeed = int32(unscaled)
	return m
}

// SetTotalCalories sets TotalCalories value.
//
// Units: kcal
func (m *SplitSummary) SetTotalCalories(v uint32) *SplitSummary {
	m.TotalCalories = v
	return m
}

// SetTotalMovingTime sets TotalMovingTime value.
//
// Scale: 1000; Units: s
func (m *SplitSummary) SetTotalMovingTime(v uint32) *SplitSummary {
	m.TotalMovingTime = v
	return m
}

// SetTotalMovingTimeScaled is similar to SetTotalMovingTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *SplitSummary) SetTotalMovingTimeScaled(v float64) *SplitSummary {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalMovingTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalMovingTime = uint32(unscaled)
	return m
}

// SetDeveloperFields SplitSummary's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SplitSummary) SetUnknownFields(unknownFields ...proto.Field) *SplitSummary {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields SplitSummary's DeveloperFields.
func (m *SplitSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *SplitSummary {
	m.DeveloperFields = developerFields
	return m
}
