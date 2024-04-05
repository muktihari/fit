// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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

// HrvStatusSummary is a HrvStatusSummary message.
type HrvStatusSummary struct {
	Timestamp             time.Time
	WeeklyAverage         uint16 // Scale: 128; Units: ms; 7 day RMSSD average over sleep
	LastNightAverage      uint16 // Scale: 128; Units: ms; Last night RMSSD average over sleep
	LastNight5MinHigh     uint16 // Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
	BaselineLowUpper      uint16 // Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
	BaselineBalancedLower uint16 // Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
	BaselineBalancedUpper uint16 // Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
	Status                typedef.HrvStatus

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrvStatusSummary creates new HrvStatusSummary struct based on given mesg.
// If mesg is nil, it will return HrvStatusSummary with all fields being set to its corresponding invalid value.
func NewHrvStatusSummary(mesg *proto.Message) *HrvStatusSummary {
	vals := [254]proto.Value{}

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

	return &HrvStatusSummary{
		Timestamp:             datetime.ToTime(vals[253].Uint32()),
		WeeklyAverage:         vals[0].Uint16(),
		LastNightAverage:      vals[1].Uint16(),
		LastNight5MinHigh:     vals[2].Uint16(),
		BaselineLowUpper:      vals[3].Uint16(),
		BaselineBalancedLower: vals[4].Uint16(),
		BaselineBalancedUpper: vals[5].Uint16(),
		Status:                typedef.HrvStatus(vals[6].Uint8()),

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

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHrvStatusSummary}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
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
	if byte(m.Status) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(byte(m.Status))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// WeeklyAverageScaled return WeeklyAverage in its scaled value [Scale: 128; Units: ms; 7 day RMSSD average over sleep].
//
// If WeeklyAverage value is invalid, float64 invalid value will be returned.
func (m *HrvStatusSummary) WeeklyAverageScaled() float64 {
	if m.WeeklyAverage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.WeeklyAverage, 128, 0)
}

// LastNightAverageScaled return LastNightAverage in its scaled value [Scale: 128; Units: ms; Last night RMSSD average over sleep].
//
// If LastNightAverage value is invalid, float64 invalid value will be returned.
func (m *HrvStatusSummary) LastNightAverageScaled() float64 {
	if m.LastNightAverage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.LastNightAverage, 128, 0)
}

// LastNight5MinHighScaled return LastNight5MinHigh in its scaled value [Scale: 128; Units: ms; 5 minute high RMSSD value over sleep].
//
// If LastNight5MinHigh value is invalid, float64 invalid value will be returned.
func (m *HrvStatusSummary) LastNight5MinHighScaled() float64 {
	if m.LastNight5MinHigh == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.LastNight5MinHigh, 128, 0)
}

// BaselineLowUpperScaled return BaselineLowUpper in its scaled value [Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status].
//
// If BaselineLowUpper value is invalid, float64 invalid value will be returned.
func (m *HrvStatusSummary) BaselineLowUpperScaled() float64 {
	if m.BaselineLowUpper == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.BaselineLowUpper, 128, 0)
}

// BaselineBalancedLowerScaled return BaselineBalancedLower in its scaled value [Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status].
//
// If BaselineBalancedLower value is invalid, float64 invalid value will be returned.
func (m *HrvStatusSummary) BaselineBalancedLowerScaled() float64 {
	if m.BaselineBalancedLower == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.BaselineBalancedLower, 128, 0)
}

// BaselineBalancedUpperScaled return BaselineBalancedUpper in its scaled value [Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status].
//
// If BaselineBalancedUpper value is invalid, float64 invalid value will be returned.
func (m *HrvStatusSummary) BaselineBalancedUpperScaled() float64 {
	if m.BaselineBalancedUpper == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.BaselineBalancedUpper, 128, 0)
}

// SetTimestamp sets HrvStatusSummary value.
func (m *HrvStatusSummary) SetTimestamp(v time.Time) *HrvStatusSummary {
	m.Timestamp = v
	return m
}

// SetWeeklyAverage sets HrvStatusSummary value.
//
// Scale: 128; Units: ms; 7 day RMSSD average over sleep
func (m *HrvStatusSummary) SetWeeklyAverage(v uint16) *HrvStatusSummary {
	m.WeeklyAverage = v
	return m
}

// SetLastNightAverage sets HrvStatusSummary value.
//
// Scale: 128; Units: ms; Last night RMSSD average over sleep
func (m *HrvStatusSummary) SetLastNightAverage(v uint16) *HrvStatusSummary {
	m.LastNightAverage = v
	return m
}

// SetLastNight5MinHigh sets HrvStatusSummary value.
//
// Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
func (m *HrvStatusSummary) SetLastNight5MinHigh(v uint16) *HrvStatusSummary {
	m.LastNight5MinHigh = v
	return m
}

// SetBaselineLowUpper sets HrvStatusSummary value.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
func (m *HrvStatusSummary) SetBaselineLowUpper(v uint16) *HrvStatusSummary {
	m.BaselineLowUpper = v
	return m
}

// SetBaselineBalancedLower sets HrvStatusSummary value.
//
// Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
func (m *HrvStatusSummary) SetBaselineBalancedLower(v uint16) *HrvStatusSummary {
	m.BaselineBalancedLower = v
	return m
}

// SetBaselineBalancedUpper sets HrvStatusSummary value.
//
// Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
func (m *HrvStatusSummary) SetBaselineBalancedUpper(v uint16) *HrvStatusSummary {
	m.BaselineBalancedUpper = v
	return m
}

// SetStatus sets HrvStatusSummary value.
func (m *HrvStatusSummary) SetStatus(v typedef.HrvStatus) *HrvStatusSummary {
	m.Status = v
	return m
}

// SetDeveloperFields HrvStatusSummary's DeveloperFields.
func (m *HrvStatusSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *HrvStatusSummary {
	m.DeveloperFields = developerFields
	return m
}
