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

// MetZone is a MetZone message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type MetZone struct {
	MessageIndex typedef.MessageIndex
	Calories     uint16 // Scale: 10; Units: kcal / min
	HighBpm      uint8
	FatCalories  uint8 // Scale: 10; Units: kcal / min

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewMetZone creates new MetZone struct based on given mesg.
// If mesg is nil, it will return MetZone with all fields being set to its corresponding invalid value.
func NewMetZone(mesg *proto.Message) *MetZone {
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
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &MetZone{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		HighBpm:      vals[1].Uint8(),
		Calories:     vals[2].Uint16(),
		FatCalories:  vals[3].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts MetZone into proto.Message. If options is nil, default options will be used.
func (m *MetZone) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumMetZone}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.HighBpm != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.HighBpm)
		fields = append(fields, field)
	}
	if m.Calories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.Calories)
		fields = append(fields, field)
	}
	if m.FatCalories != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.FatCalories)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// CaloriesScaled return Calories in its scaled value.
// If Calories value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) CaloriesScaled() float64 {
	if m.Calories == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Calories)/10 - 0
}

// FatCaloriesScaled return FatCalories in its scaled value.
// If FatCalories value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) FatCaloriesScaled() float64 {
	if m.FatCalories == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.FatCalories)/10 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *MetZone) SetMessageIndex(v typedef.MessageIndex) *MetZone {
	m.MessageIndex = v
	return m
}

// SetHighBpm sets HighBpm value.
func (m *MetZone) SetHighBpm(v uint8) *MetZone {
	m.HighBpm = v
	return m
}

// SetCalories sets Calories value.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) SetCalories(v uint16) *MetZone {
	m.Calories = v
	return m
}

// SetCaloriesScaled is similar to SetCalories except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) SetCaloriesScaled(v float64) *MetZone {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Calories = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Calories = uint16(unscaled)
	return m
}

// SetFatCalories sets FatCalories value.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) SetFatCalories(v uint8) *MetZone {
	m.FatCalories = v
	return m
}

// SetFatCaloriesScaled is similar to SetFatCalories except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) SetFatCaloriesScaled(v float64) *MetZone {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.FatCalories = uint8(basetype.Uint8Invalid)
		return m
	}
	m.FatCalories = uint8(unscaled)
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *MetZone) SetUnknownFields(unknownFields ...proto.Field) *MetZone {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *MetZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *MetZone {
	m.DeveloperFields = developerFields
	return m
}
