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

// ChronoShotData is a ChronoShotData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ChronoShotData struct {
	Timestamp time.Time
	ShotSpeed uint32 // Scale: 1000; Units: m/s
	ShotNum   uint16

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewChronoShotData creates new ChronoShotData struct based on given mesg.
// If mesg is nil, it will return ChronoShotData with all fields being set to its corresponding invalid value.
func NewChronoShotData(mesg *proto.Message) *ChronoShotData {
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

	return &ChronoShotData{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		ShotSpeed: vals[0].Uint32(),
		ShotNum:   vals[1].Uint16(),

		UnknownFields:   unknownFields,
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

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumChronoShotData}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
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
func (m *ChronoShotData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// ShotSpeedScaled return ShotSpeed in its scaled value.
// If ShotSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotData) ShotSpeedScaled() float64 {
	if m.ShotSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.ShotSpeed)/1000 - 0
}

// SetTimestamp sets Timestamp value.
func (m *ChronoShotData) SetTimestamp(v time.Time) *ChronoShotData {
	m.Timestamp = v
	return m
}

// SetShotSpeed sets ShotSpeed value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotData) SetShotSpeed(v uint32) *ChronoShotData {
	m.ShotSpeed = v
	return m
}

// SetShotSpeedScaled is similar to SetShotSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotData) SetShotSpeedScaled(v float64) *ChronoShotData {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.ShotSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.ShotSpeed = uint32(unscaled)
	return m
}

// SetShotNum sets ShotNum value.
func (m *ChronoShotData) SetShotNum(v uint16) *ChronoShotData {
	m.ShotNum = v
	return m
}

// SetDeveloperFields ChronoShotData's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *ChronoShotData) SetUnknownFields(unknownFields ...proto.Field) *ChronoShotData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields ChronoShotData's DeveloperFields.
func (m *ChronoShotData) SetDeveloperFields(developerFields ...proto.DeveloperField) *ChronoShotData {
	m.DeveloperFields = developerFields
	return m
}
