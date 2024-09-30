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

// HrvStatusSummary is a HrvStatusSummary message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HrvStatusSummary struct {
	Timestamp             time.Time
	WeeklyAverage         uint16 // Scale: 128; Units: ms; 7 day RMSSD average over sleep
	LastNightAverage      uint16 // Scale: 128; Units: ms; Last night RMSSD average over sleep
	LastNight5MinHigh     uint16 // Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
	BaselineLowUpper      uint16 // Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
	BaselineBalancedLower uint16 // Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
	BaselineBalancedUpper uint16 // Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
	Status                typedef.HrvStatus

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewHrvStatusSummary creates new HrvStatusSummary struct based on given mesg.
// If mesg is nil, it will return HrvStatusSummary with all fields being set to its corresponding invalid value.
func NewHrvStatusSummary(mesg *proto.Message) *HrvStatusSummary {
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
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &HrvStatusSummary{
		Timestamp:             datetime.ToTime(vals[253].Uint32()),
		WeeklyAverage:         vals[0].Uint16(),
		LastNightAverage:      vals[1].Uint16(),
		LastNight5MinHigh:     vals[2].Uint16(),
		BaselineLowUpper:      vals[3].Uint16(),
		BaselineBalancedLower: vals[4].Uint16(),
		BaselineBalancedUpper: vals[5].Uint16(),
		Status:                typedef.HrvStatus(vals[6].Uint8()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts HrvStatusSummary into proto.Message. If options is nil, default options will be used.
func (m *HrvStatusSummary) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHrvStatusSummary}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.WeeklyAverage != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.WeeklyAverage)
		fields = append(fields, field)
	}
	if m.LastNightAverage != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.LastNightAverage)
		fields = append(fields, field)
	}
	if m.LastNight5MinHigh != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.LastNight5MinHigh)
		fields = append(fields, field)
	}
	if m.BaselineLowUpper != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.BaselineLowUpper)
		fields = append(fields, field)
	}
	if m.BaselineBalancedLower != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.BaselineBalancedLower)
		fields = append(fields, field)
	}
	if m.BaselineBalancedUpper != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.BaselineBalancedUpper)
		fields = append(fields, field)
	}
	if m.Status != typedef.HrvStatusInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(byte(m.Status))
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

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HrvStatusSummary) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// WeeklyAverageScaled return WeeklyAverage in its scaled value.
// If WeeklyAverage value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; 7 day RMSSD average over sleep
func (m *HrvStatusSummary) WeeklyAverageScaled() float64 {
	if m.WeeklyAverage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.WeeklyAverage)/128 - 0
}

// LastNightAverageScaled return LastNightAverage in its scaled value.
// If LastNightAverage value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; Last night RMSSD average over sleep
func (m *HrvStatusSummary) LastNightAverageScaled() float64 {
	if m.LastNightAverage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.LastNightAverage)/128 - 0
}

// LastNight5MinHighScaled return LastNight5MinHigh in its scaled value.
// If LastNight5MinHigh value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
func (m *HrvStatusSummary) LastNight5MinHighScaled() float64 {
	if m.LastNight5MinHigh == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.LastNight5MinHigh)/128 - 0
}

// BaselineLowUpperScaled return BaselineLowUpper in its scaled value.
// If BaselineLowUpper value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
func (m *HrvStatusSummary) BaselineLowUpperScaled() float64 {
	if m.BaselineLowUpper == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BaselineLowUpper)/128 - 0
}

// BaselineBalancedLowerScaled return BaselineBalancedLower in its scaled value.
// If BaselineBalancedLower value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
func (m *HrvStatusSummary) BaselineBalancedLowerScaled() float64 {
	if m.BaselineBalancedLower == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BaselineBalancedLower)/128 - 0
}

// BaselineBalancedUpperScaled return BaselineBalancedUpper in its scaled value.
// If BaselineBalancedUpper value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
func (m *HrvStatusSummary) BaselineBalancedUpperScaled() float64 {
	if m.BaselineBalancedUpper == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BaselineBalancedUpper)/128 - 0
}

// SetTimestamp sets Timestamp value.
func (m *HrvStatusSummary) SetTimestamp(v time.Time) *HrvStatusSummary {
	m.Timestamp = v
	return m
}

// SetWeeklyAverage sets WeeklyAverage value.
//
// Scale: 128; Units: ms; 7 day RMSSD average over sleep
func (m *HrvStatusSummary) SetWeeklyAverage(v uint16) *HrvStatusSummary {
	m.WeeklyAverage = v
	return m
}

// SetWeeklyAverageScaled is similar to SetWeeklyAverage except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; 7 day RMSSD average over sleep
func (m *HrvStatusSummary) SetWeeklyAverageScaled(v float64) *HrvStatusSummary {
	unscaled := (v + 0) * 128
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.WeeklyAverage = uint16(basetype.Uint16Invalid)
		return m
	}
	m.WeeklyAverage = uint16(unscaled)
	return m
}

// SetLastNightAverage sets LastNightAverage value.
//
// Scale: 128; Units: ms; Last night RMSSD average over sleep
func (m *HrvStatusSummary) SetLastNightAverage(v uint16) *HrvStatusSummary {
	m.LastNightAverage = v
	return m
}

// SetLastNightAverageScaled is similar to SetLastNightAverage except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; Last night RMSSD average over sleep
func (m *HrvStatusSummary) SetLastNightAverageScaled(v float64) *HrvStatusSummary {
	unscaled := (v + 0) * 128
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.LastNightAverage = uint16(basetype.Uint16Invalid)
		return m
	}
	m.LastNightAverage = uint16(unscaled)
	return m
}

// SetLastNight5MinHigh sets LastNight5MinHigh value.
//
// Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
func (m *HrvStatusSummary) SetLastNight5MinHigh(v uint16) *HrvStatusSummary {
	m.LastNight5MinHigh = v
	return m
}

// SetLastNight5MinHighScaled is similar to SetLastNight5MinHigh except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
func (m *HrvStatusSummary) SetLastNight5MinHighScaled(v float64) *HrvStatusSummary {
	unscaled := (v + 0) * 128
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.LastNight5MinHigh = uint16(basetype.Uint16Invalid)
		return m
	}
	m.LastNight5MinHigh = uint16(unscaled)
	return m
}

// SetBaselineLowUpper sets BaselineLowUpper value.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
func (m *HrvStatusSummary) SetBaselineLowUpper(v uint16) *HrvStatusSummary {
	m.BaselineLowUpper = v
	return m
}

// SetBaselineLowUpperScaled is similar to SetBaselineLowUpper except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
func (m *HrvStatusSummary) SetBaselineLowUpperScaled(v float64) *HrvStatusSummary {
	unscaled := (v + 0) * 128
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BaselineLowUpper = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BaselineLowUpper = uint16(unscaled)
	return m
}

// SetBaselineBalancedLower sets BaselineBalancedLower value.
//
// Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
func (m *HrvStatusSummary) SetBaselineBalancedLower(v uint16) *HrvStatusSummary {
	m.BaselineBalancedLower = v
	return m
}

// SetBaselineBalancedLowerScaled is similar to SetBaselineBalancedLower except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
func (m *HrvStatusSummary) SetBaselineBalancedLowerScaled(v float64) *HrvStatusSummary {
	unscaled := (v + 0) * 128
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BaselineBalancedLower = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BaselineBalancedLower = uint16(unscaled)
	return m
}

// SetBaselineBalancedUpper sets BaselineBalancedUpper value.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
func (m *HrvStatusSummary) SetBaselineBalancedUpper(v uint16) *HrvStatusSummary {
	m.BaselineBalancedUpper = v
	return m
}

// SetBaselineBalancedUpperScaled is similar to SetBaselineBalancedUpper except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
func (m *HrvStatusSummary) SetBaselineBalancedUpperScaled(v float64) *HrvStatusSummary {
	unscaled := (v + 0) * 128
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BaselineBalancedUpper = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BaselineBalancedUpper = uint16(unscaled)
	return m
}

// SetStatus sets Status value.
func (m *HrvStatusSummary) SetStatus(v typedef.HrvStatus) *HrvStatusSummary {
	m.Status = v
	return m
}

// SetDeveloperFields HrvStatusSummary's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *HrvStatusSummary) SetUnknownFields(unknownFields ...proto.Field) *HrvStatusSummary {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields HrvStatusSummary's DeveloperFields.
func (m *HrvStatusSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *HrvStatusSummary {
	m.DeveloperFields = developerFields
	return m
}
