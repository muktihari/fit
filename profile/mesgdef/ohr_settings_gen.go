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

// OhrSettings is a OhrSettings message.
type OhrSettings struct {
	Timestamp time.Time // Units: s
	Enabled   typedef.Switch

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewOhrSettings creates new OhrSettings struct based on given mesg.
// If mesg is nil, it will return OhrSettings with all fields being set to its corresponding invalid value.
func NewOhrSettings(mesg *proto.Message) *OhrSettings {
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

	return &OhrSettings{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		Enabled:   typedef.Switch(vals[0].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts OhrSettings into proto.Message. If options is nil, default options will be used.
func (m *OhrSettings) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumOhrSettings)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if byte(m.Enabled) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Enabled))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets OhrSettings value.
//
// Units: s
func (m *OhrSettings) SetTimestamp(v time.Time) *OhrSettings {
	m.Timestamp = v
	return m
}

// SetEnabled sets OhrSettings value.
func (m *OhrSettings) SetEnabled(v typedef.Switch) *OhrSettings {
	m.Enabled = v
	return m
}

// SetDeveloperFields OhrSettings's DeveloperFields.
func (m *OhrSettings) SetDeveloperFields(developerFields ...proto.DeveloperField) *OhrSettings {
	m.DeveloperFields = developerFields
	return m
}
