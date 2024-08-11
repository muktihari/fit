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
	"time"
)

// StressLevel is a StressLevel message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type StressLevel struct {
	StressLevelTime  time.Time // Units: s; Time stress score was calculated
	StressLevelValue int16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewStressLevel creates new StressLevel struct based on given mesg.
// If mesg is nil, it will return StressLevel with all fields being set to its corresponding invalid value.
func NewStressLevel(mesg *proto.Message) *StressLevel {
	vals := [2]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 1 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &StressLevel{
		StressLevelValue: vals[0].Int16(),
		StressLevelTime:  datetime.ToTime(vals[1].Uint32()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts StressLevel into proto.Message. If options is nil, default options will be used.
func (m *StressLevel) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumStressLevel}

	if m.StressLevelValue != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Int16(m.StressLevelValue)
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.StressLevelTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(datetime.ToUint32(m.StressLevelTime))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// StressLevelTimeUint32 returns StressLevelTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *StressLevel) StressLevelTimeUint32() uint32 { return datetime.ToUint32(m.StressLevelTime) }

// SetStressLevelValue sets StressLevelValue value.
func (m *StressLevel) SetStressLevelValue(v int16) *StressLevel {
	m.StressLevelValue = v
	return m
}

// SetStressLevelTime sets StressLevelTime value.
//
// Units: s; Time stress score was calculated
func (m *StressLevel) SetStressLevelTime(v time.Time) *StressLevel {
	m.StressLevelTime = v
	return m
}

// SetDeveloperFields StressLevel's DeveloperFields.
func (m *StressLevel) SetDeveloperFields(developerFields ...proto.DeveloperField) *StressLevel {
	m.DeveloperFields = developerFields
	return m
}
