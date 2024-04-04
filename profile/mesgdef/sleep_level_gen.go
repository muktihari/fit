// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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

// SleepLevel is a SleepLevel message.
type SleepLevel struct {
	Timestamp  time.Time // Units: s
	SleepLevel typedef.SleepLevel

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSleepLevel creates new SleepLevel struct based on given mesg.
// If mesg is nil, it will return SleepLevel with all fields being set to its corresponding invalid value.
func NewSleepLevel(mesg *proto.Message) *SleepLevel {
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

	return &SleepLevel{
		Timestamp:  datetime.ToTime(vals[253].Uint32()),
		SleepLevel: typedef.SleepLevel(vals[0].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SleepLevel into proto.Message. If options is nil, default options will be used.
func (m *SleepLevel) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSleepLevel)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if byte(m.SleepLevel) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.SleepLevel))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets SleepLevel value.
//
// Units: s
func (m *SleepLevel) SetTimestamp(v time.Time) *SleepLevel {
	m.Timestamp = v
	return m
}

// SetSleepLevel sets SleepLevel value.
func (m *SleepLevel) SetSleepLevel(v typedef.SleepLevel) *SleepLevel {
	m.SleepLevel = v
	return m
}

// SetDeveloperFields SleepLevel's DeveloperFields.
func (m *SleepLevel) SetDeveloperFields(developerFields ...proto.DeveloperField) *SleepLevel {
	m.DeveloperFields = developerFields
	return m
}
