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

// ChronoShotData is a ChronoShotData message.
type ChronoShotData struct {
	Timestamp time.Time
	ShotSpeed uint32 // Scale: 1000; Units: m/s
	ShotNum   uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewChronoShotData creates new ChronoShotData struct based on given mesg.
// If mesg is nil, it will return ChronoShotData with all fields being set to its corresponding invalid value.
func NewChronoShotData(mesg *proto.Message) *ChronoShotData {
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

	return &ChronoShotData{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		ShotSpeed: vals[0].Uint32(),
		ShotNum:   vals[1].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ChronoShotData into proto.Message. If options is nil, default options will be used.
func (m *ChronoShotData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumChronoShotData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.ShotSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.ShotSpeed)
		fields = append(fields, field)
	}
	if m.ShotNum != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.ShotNum)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// ShotSpeedScaled return ShotSpeed in its scaled value [Scale: 1000; Units: m/s].
//
// If ShotSpeed value is invalid, float64 invalid value will be returned.
func (m *ChronoShotData) ShotSpeedScaled() float64 {
	if m.ShotSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.ShotSpeed, 1000, 0)
}

// SetTimestamp sets ChronoShotData value.
func (m *ChronoShotData) SetTimestamp(v time.Time) *ChronoShotData {
	m.Timestamp = v
	return m
}

// SetShotSpeed sets ChronoShotData value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotData) SetShotSpeed(v uint32) *ChronoShotData {
	m.ShotSpeed = v
	return m
}

// SetShotNum sets ChronoShotData value.
func (m *ChronoShotData) SetShotNum(v uint16) *ChronoShotData {
	m.ShotNum = v
	return m
}

// SetDeveloperFields ChronoShotData's DeveloperFields.
func (m *ChronoShotData) SetDeveloperFields(developerFields ...proto.DeveloperField) *ChronoShotData {
	m.DeveloperFields = developerFields
	return m
}
