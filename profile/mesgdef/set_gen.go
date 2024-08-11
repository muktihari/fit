// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
	"unsafe"
)

// Set is a Set message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Set struct {
	Timestamp         time.Time                  // Timestamp of the set
	StartTime         time.Time                  // Start time of the set
	Category          []typedef.ExerciseCategory // Array: [N]
	CategorySubtype   []uint16                   // Array: [N]; Based on the associated category, see [category]_exercise_names
	Duration          uint32                     // Scale: 1000; Units: s
	Repetitions       uint16                     // # of repitions of the movement
	Weight            uint16                     // Scale: 16; Units: kg; Amount of weight applied for the set
	WeightDisplayUnit typedef.FitBaseUnit
	MessageIndex      typedef.MessageIndex
	WktStepIndex      typedef.MessageIndex
	SetType           typedef.SetType

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSet creates new Set struct based on given mesg.
// If mesg is nil, it will return Set with all fields being set to its corresponding invalid value.
func NewSet(mesg *proto.Message) *Set {
	vals := [255]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Set{
		Timestamp:   datetime.ToTime(vals[254].Uint32()),
		Duration:    vals[0].Uint32(),
		Repetitions: vals[3].Uint16(),
		Weight:      vals[4].Uint16(),
		SetType:     typedef.SetType(vals[5].Uint8()),
		StartTime:   datetime.ToTime(vals[6].Uint32()),
		Category: func() []typedef.ExerciseCategory {
			sliceValue := vals[7].SliceUint16()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.ExerciseCategory)(ptr), len(sliceValue))
		}(),
		CategorySubtype:   vals[8].SliceUint16(),
		WeightDisplayUnit: typedef.FitBaseUnit(vals[9].Uint16()),
		MessageIndex:      typedef.MessageIndex(vals[10].Uint16()),
		WktStepIndex:      typedef.MessageIndex(vals[11].Uint16()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Set into proto.Message. If options is nil, default options will be used.
func (m *Set) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumSet}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.Duration)
		fields = append(fields, field)
	}
	if m.Repetitions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.Repetitions)
		fields = append(fields, field)
	}
	if m.Weight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.Weight)
		fields = append(fields, field)
	}
	if uint8(m.SetType) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(uint8(m.SetType))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.StartTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(datetime.ToUint32(m.StartTime))
		fields = append(fields, field)
	}
	if m.Category != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceUint16(m.Category)
		fields = append(fields, field)
	}
	if m.CategorySubtype != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceUint16(m.CategorySubtype)
		fields = append(fields, field)
	}
	if uint16(m.WeightDisplayUnit) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(uint16(m.WeightDisplayUnit))
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if uint16(m.WktStepIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint16(uint16(m.WktStepIndex))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Set) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// StartTimeUint32 returns StartTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Set) StartTimeUint32() uint32 { return datetime.ToUint32(m.StartTime) }

// DurationScaled return Duration in its scaled value.
// If Duration value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *Set) DurationScaled() float64 {
	if m.Duration == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Duration)/1000 - 0
}

// WeightScaled return Weight in its scaled value.
// If Weight value is invalid, float64 invalid value will be returned.
//
// Scale: 16; Units: kg; Amount of weight applied for the set
func (m *Set) WeightScaled() float64 {
	if m.Weight == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Weight)/16 - 0
}

// SetTimestamp sets Timestamp value.
//
// Timestamp of the set
func (m *Set) SetTimestamp(v time.Time) *Set {
	m.Timestamp = v
	return m
}

// SetDuration sets Duration value.
//
// Scale: 1000; Units: s
func (m *Set) SetDuration(v uint32) *Set {
	m.Duration = v
	return m
}

// SetDurationScaled is similar to SetDuration except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *Set) SetDurationScaled(v float64) *Set {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Duration = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Duration = uint32(unscaled)
	return m
}

// SetRepetitions sets Repetitions value.
//
// # of repitions of the movement
func (m *Set) SetRepetitions(v uint16) *Set {
	m.Repetitions = v
	return m
}

// SetWeight sets Weight value.
//
// Scale: 16; Units: kg; Amount of weight applied for the set
func (m *Set) SetWeight(v uint16) *Set {
	m.Weight = v
	return m
}

// SetWeightScaled is similar to SetWeight except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 16; Units: kg; Amount of weight applied for the set
func (m *Set) SetWeightScaled(v float64) *Set {
	unscaled := (v + 0) * 16
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Weight = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Weight = uint16(unscaled)
	return m
}

// SetSetType sets SetType value.
func (m *Set) SetSetType(v typedef.SetType) *Set {
	m.SetType = v
	return m
}

// SetStartTime sets StartTime value.
//
// Start time of the set
func (m *Set) SetStartTime(v time.Time) *Set {
	m.StartTime = v
	return m
}

// SetCategory sets Category value.
//
// Array: [N]
func (m *Set) SetCategory(v []typedef.ExerciseCategory) *Set {
	m.Category = v
	return m
}

// SetCategorySubtype sets CategorySubtype value.
//
// Array: [N]; Based on the associated category, see [category]_exercise_names
func (m *Set) SetCategorySubtype(v []uint16) *Set {
	m.CategorySubtype = v
	return m
}

// SetWeightDisplayUnit sets WeightDisplayUnit value.
func (m *Set) SetWeightDisplayUnit(v typedef.FitBaseUnit) *Set {
	m.WeightDisplayUnit = v
	return m
}

// SetMessageIndex sets MessageIndex value.
func (m *Set) SetMessageIndex(v typedef.MessageIndex) *Set {
	m.MessageIndex = v
	return m
}

// SetWktStepIndex sets WktStepIndex value.
func (m *Set) SetWktStepIndex(v typedef.MessageIndex) *Set {
	m.WktStepIndex = v
	return m
}

// SetDeveloperFields Set's DeveloperFields.
func (m *Set) SetDeveloperFields(developerFields ...proto.DeveloperField) *Set {
	m.DeveloperFields = developerFields
	return m
}
