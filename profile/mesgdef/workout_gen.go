// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// Workout is a Workout message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Workout struct {
	WktName        string
	Capabilities   typedef.WorkoutCapabilities
	MessageIndex   typedef.MessageIndex
	NumValidSteps  uint16 // number of valid steps
	PoolLength     uint16 // Scale: 100; Units: m
	Sport          typedef.Sport
	SubSport       typedef.SubSport
	PoolLengthUnit typedef.DisplayMeasure

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWorkout creates new Workout struct based on given mesg.
// If mesg is nil, it will return Workout with all fields being set to its corresponding invalid value.
func NewWorkout(mesg *proto.Message) *Workout {
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

	return &Workout{
		MessageIndex:   typedef.MessageIndex(vals[254].Uint16()),
		Sport:          typedef.Sport(vals[4].Uint8()),
		Capabilities:   typedef.WorkoutCapabilities(vals[5].Uint32z()),
		NumValidSteps:  vals[6].Uint16(),
		WktName:        vals[8].String(),
		SubSport:       typedef.SubSport(vals[11].Uint8()),
		PoolLength:     vals[14].Uint16(),
		PoolLengthUnit: typedef.DisplayMeasure(vals[15].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Workout into proto.Message. If options is nil, default options will be used.
func (m *Workout) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumWorkout}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if uint32(m.Capabilities) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(uint32(m.Capabilities))
		fields = append(fields, field)
	}
	if m.NumValidSteps != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint16(m.NumValidSteps)
		fields = append(fields, field)
	}
	if m.WktName != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.String(m.WktName)
		fields = append(fields, field)
	}
	if byte(m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(byte(m.SubSport))
		fields = append(fields, field)
	}
	if m.PoolLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint16(m.PoolLength)
		fields = append(fields, field)
	}
	if byte(m.PoolLengthUnit) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint8(byte(m.PoolLengthUnit))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// PoolLengthScaled return PoolLength in its scaled value.
// If PoolLength value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *Workout) PoolLengthScaled() float64 {
	if m.PoolLength == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.PoolLength)/100 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *Workout) SetMessageIndex(v typedef.MessageIndex) *Workout {
	m.MessageIndex = v
	return m
}

// SetSport sets Sport value.
func (m *Workout) SetSport(v typedef.Sport) *Workout {
	m.Sport = v
	return m
}

// SetCapabilities sets Capabilities value.
func (m *Workout) SetCapabilities(v typedef.WorkoutCapabilities) *Workout {
	m.Capabilities = v
	return m
}

// SetNumValidSteps sets NumValidSteps value.
//
// number of valid steps
func (m *Workout) SetNumValidSteps(v uint16) *Workout {
	m.NumValidSteps = v
	return m
}

// SetWktName sets WktName value.
func (m *Workout) SetWktName(v string) *Workout {
	m.WktName = v
	return m
}

// SetSubSport sets SubSport value.
func (m *Workout) SetSubSport(v typedef.SubSport) *Workout {
	m.SubSport = v
	return m
}

// SetPoolLength sets PoolLength value.
//
// Scale: 100; Units: m
func (m *Workout) SetPoolLength(v uint16) *Workout {
	m.PoolLength = v
	return m
}

// SetPoolLengthScaled is similar to SetPoolLength except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *Workout) SetPoolLengthScaled(v float64) *Workout {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.PoolLength = uint16(basetype.Uint16Invalid)
		return m
	}
	m.PoolLength = uint16(unscaled)
	return m
}

// SetPoolLengthUnit sets PoolLengthUnit value.
func (m *Workout) SetPoolLengthUnit(v typedef.DisplayMeasure) *Workout {
	m.PoolLengthUnit = v
	return m
}

// SetDeveloperFields Workout's DeveloperFields.
func (m *Workout) SetDeveloperFields(developerFields ...proto.DeveloperField) *Workout {
	m.DeveloperFields = developerFields
	return m
}
