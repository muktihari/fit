// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// MetZone is a MetZone message.
type MetZone struct {
	MessageIndex typedef.MessageIndex
	Calories     uint16 // Scale: 10; Units: kcal / min
	HighBpm      uint8
	FatCalories  uint8 // Scale: 10; Units: kcal / min

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMetZone creates new MetZone struct based on given mesg.
// If mesg is nil, it will return MetZone with all fields being set to its corresponding invalid value.
func NewMetZone(mesg *proto.Message) *MetZone {
	vals := [255]any{}

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

	return &MetZone{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Calories:     typeconv.ToUint16[uint16](vals[2]),
		HighBpm:      typeconv.ToUint8[uint8](vals[1]),
		FatCalories:  typeconv.ToUint8[uint8](vals[3]),

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

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumMetZone)

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if m.Calories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Calories
		fields = append(fields, field)
	}
	if m.HighBpm != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.HighBpm
		fields = append(fields, field)
	}
	if m.FatCalories != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.FatCalories
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// CaloriesScaled return Calories in its scaled value [Scale: 10; Units: kcal / min].
//
// If Calories value is invalid, float64 invalid value will be returned.
func (m *MetZone) CaloriesScaled() float64 {
	if m.Calories == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Calories, 10, 0)
}

// FatCaloriesScaled return FatCalories in its scaled value [Scale: 10; Units: kcal / min].
//
// If FatCalories value is invalid, float64 invalid value will be returned.
func (m *MetZone) FatCaloriesScaled() float64 {
	if m.FatCalories == basetype.Uint8Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.FatCalories, 10, 0)
}

// SetMessageIndex sets MetZone value.
func (m *MetZone) SetMessageIndex(v typedef.MessageIndex) *MetZone {
	m.MessageIndex = v
	return m
}

// SetCalories sets MetZone value.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) SetCalories(v uint16) *MetZone {
	m.Calories = v
	return m
}

// SetHighBpm sets MetZone value.
func (m *MetZone) SetHighBpm(v uint8) *MetZone {
	m.HighBpm = v
	return m
}

// SetFatCalories sets MetZone value.
//
// Scale: 10; Units: kcal / min
func (m *MetZone) SetFatCalories(v uint8) *MetZone {
	m.FatCalories = v
	return m
}

// SetDeveloperFields MetZone's DeveloperFields.
func (m *MetZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *MetZone {
	m.DeveloperFields = developerFields
	return m
}
