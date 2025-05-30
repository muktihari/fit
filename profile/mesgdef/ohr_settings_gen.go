// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// OhrSettings is a OhrSettings message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type OhrSettings struct {
	Timestamp time.Time // Units: s
	Enabled   typedef.Switch

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewOhrSettings creates new OhrSettings struct based on given mesg.
// If mesg is nil, it will return OhrSettings with all fields being set to its corresponding invalid value.
func NewOhrSettings(mesg *proto.Message) *OhrSettings {
	m := new(OhrSettings)
	m.Reset(mesg)
	return m
}

// Reset resets all OhrSettings's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *OhrSettings) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		var n int
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				n++
			}
		}
		unknownFields = make([]proto.Field, 0, n)
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 254 {
				vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
			}
		}
		developerFields = mesg.DeveloperFields
	}

	*m = OhrSettings{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		Enabled:   typedef.Switch(vals[0].Uint8()),

		UnknownFields:   unknownFields,
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

	fields := make([]proto.Field, 0, 2)
	mesg := proto.Message{Num: typedef.MesgNumOhrSettings}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Enabled != typedef.SwitchInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Enabled))
		fields = append(fields, field)
	}

	n := len(fields)
	mesg.Fields = make([]proto.Field, n+len(m.UnknownFields))
	copy(mesg.Fields[:n], fields)
	copy(mesg.Fields[n:], m.UnknownFields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *OhrSettings) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *OhrSettings) SetTimestamp(v time.Time) *OhrSettings {
	m.Timestamp = v
	return m
}

// SetEnabled sets Enabled value.
func (m *OhrSettings) SetEnabled(v typedef.Switch) *OhrSettings {
	m.Enabled = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *OhrSettings) SetUnknownFields(unknownFields ...proto.Field) *OhrSettings {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *OhrSettings) SetDeveloperFields(developerFields ...proto.DeveloperField) *OhrSettings {
	m.DeveloperFields = developerFields
	return m
}
