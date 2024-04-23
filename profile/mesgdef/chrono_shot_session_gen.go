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

// ChronoShotSession is a ChronoShotSession message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ChronoShotSession struct {
	Timestamp      time.Time
	MinSpeed       uint32 // Scale: 1000; Units: m/s
	MaxSpeed       uint32 // Scale: 1000; Units: m/s
	AvgSpeed       uint32 // Scale: 1000; Units: m/s
	GrainWeight    uint32 // Scale: 10; Units: gr
	ShotCount      uint16
	ProjectileType typedef.ProjectileType

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewChronoShotSession creates new ChronoShotSession struct based on given mesg.
// If mesg is nil, it will return ChronoShotSession with all fields being set to its corresponding invalid value.
func NewChronoShotSession(mesg *proto.Message) *ChronoShotSession {
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

	return &ChronoShotSession{
		Timestamp:      datetime.ToTime(vals[253].Uint32()),
		MinSpeed:       vals[0].Uint32(),
		MaxSpeed:       vals[1].Uint32(),
		AvgSpeed:       vals[2].Uint32(),
		ShotCount:      vals[3].Uint16(),
		ProjectileType: typedef.ProjectileType(vals[4].Uint8()),
		GrainWeight:    vals[5].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ChronoShotSession into proto.Message. If options is nil, default options will be used.
func (m *ChronoShotSession) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumChronoShotSession}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.MinSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.MinSpeed)
		fields = append(fields, field)
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.MaxSpeed)
		fields = append(fields, field)
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.AvgSpeed)
		fields = append(fields, field)
	}
	if m.ShotCount != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.ShotCount)
		fields = append(fields, field)
	}
	if byte(m.ProjectileType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.ProjectileType))
		fields = append(fields, field)
	}
	if m.GrainWeight != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.GrainWeight)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *ChronoShotSession) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// MinSpeedScaled return MinSpeed in its scaled value [Scale: 1000; Units: m/s].
//
// If MinSpeed value is invalid, float64 invalid value will be returned.
func (m *ChronoShotSession) MinSpeedScaled() float64 {
	if m.MinSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.MinSpeed, 1000, 0)
}

// MaxSpeedScaled return MaxSpeed in its scaled value [Scale: 1000; Units: m/s].
//
// If MaxSpeed value is invalid, float64 invalid value will be returned.
func (m *ChronoShotSession) MaxSpeedScaled() float64 {
	if m.MaxSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.MaxSpeed, 1000, 0)
}

// AvgSpeedScaled return AvgSpeed in its scaled value [Scale: 1000; Units: m/s].
//
// If AvgSpeed value is invalid, float64 invalid value will be returned.
func (m *ChronoShotSession) AvgSpeedScaled() float64 {
	if m.AvgSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgSpeed, 1000, 0)
}

// GrainWeightScaled return GrainWeight in its scaled value [Scale: 10; Units: gr].
//
// If GrainWeight value is invalid, float64 invalid value will be returned.
func (m *ChronoShotSession) GrainWeightScaled() float64 {
	if m.GrainWeight == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.GrainWeight, 10, 0)
}

// SetTimestamp sets ChronoShotSession value.
func (m *ChronoShotSession) SetTimestamp(v time.Time) *ChronoShotSession {
	m.Timestamp = v
	return m
}

// SetMinSpeed sets ChronoShotSession value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetMinSpeed(v uint32) *ChronoShotSession {
	m.MinSpeed = v
	return m
}

// SetMaxSpeed sets ChronoShotSession value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetMaxSpeed(v uint32) *ChronoShotSession {
	m.MaxSpeed = v
	return m
}

// SetAvgSpeed sets ChronoShotSession value.
//
// Scale: 1000; Units: m/s
func (m *ChronoShotSession) SetAvgSpeed(v uint32) *ChronoShotSession {
	m.AvgSpeed = v
	return m
}

// SetShotCount sets ChronoShotSession value.
func (m *ChronoShotSession) SetShotCount(v uint16) *ChronoShotSession {
	m.ShotCount = v
	return m
}

// SetProjectileType sets ChronoShotSession value.
func (m *ChronoShotSession) SetProjectileType(v typedef.ProjectileType) *ChronoShotSession {
	m.ProjectileType = v
	return m
}

// SetGrainWeight sets ChronoShotSession value.
//
// Scale: 10; Units: gr
func (m *ChronoShotSession) SetGrainWeight(v uint32) *ChronoShotSession {
	m.GrainWeight = v
	return m
}

// SetDeveloperFields ChronoShotSession's DeveloperFields.
func (m *ChronoShotSession) SetDeveloperFields(developerFields ...proto.DeveloperField) *ChronoShotSession {
	m.DeveloperFields = developerFields
	return m
}
