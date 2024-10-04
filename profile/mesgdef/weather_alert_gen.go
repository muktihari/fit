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
	"time"
)

// WeatherAlert is a WeatherAlert message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type WeatherAlert struct {
	Timestamp  time.Time
	ReportId   string                    // Unique identifier from GCS report ID string, length is 12
	IssueTime  time.Time                 // Time alert was issued
	ExpireTime time.Time                 // Time alert expires
	Severity   typedef.WeatherSeverity   // Warning, Watch, Advisory, Statement
	Type       typedef.WeatherSevereType // Tornado, Severe Thunderstorm, etc.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewWeatherAlert creates new WeatherAlert struct based on given mesg.
// If mesg is nil, it will return WeatherAlert with all fields being set to its corresponding invalid value.
func NewWeatherAlert(mesg *proto.Message) *WeatherAlert {
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
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &WeatherAlert{
		Timestamp:  datetime.ToTime(vals[253].Uint32()),
		ReportId:   vals[0].String(),
		IssueTime:  datetime.ToTime(vals[1].Uint32()),
		ExpireTime: datetime.ToTime(vals[2].Uint32()),
		Severity:   typedef.WeatherSeverity(vals[3].Uint8()),
		Type:       typedef.WeatherSevereType(vals[4].Uint8()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts WeatherAlert into proto.Message. If options is nil, default options will be used.
func (m *WeatherAlert) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumWeatherAlert}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.ReportId != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.ReportId)
		fields = append(fields, field)
	}
	if !m.IssueTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(uint32(m.IssueTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if !m.ExpireTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(uint32(m.ExpireTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Severity != typedef.WeatherSeverityInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Severity))
		fields = append(fields, field)
	}
	if m.Type != typedef.WeatherSevereTypeInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *WeatherAlert) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// IssueTimeUint32 returns IssueTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *WeatherAlert) IssueTimeUint32() uint32 { return datetime.ToUint32(m.IssueTime) }

// ExpireTimeUint32 returns ExpireTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *WeatherAlert) ExpireTimeUint32() uint32 { return datetime.ToUint32(m.ExpireTime) }

// SetTimestamp sets Timestamp value.
func (m *WeatherAlert) SetTimestamp(v time.Time) *WeatherAlert {
	m.Timestamp = v
	return m
}

// SetReportId sets ReportId value.
//
// Unique identifier from GCS report ID string, length is 12
func (m *WeatherAlert) SetReportId(v string) *WeatherAlert {
	m.ReportId = v
	return m
}

// SetIssueTime sets IssueTime value.
//
// Time alert was issued
func (m *WeatherAlert) SetIssueTime(v time.Time) *WeatherAlert {
	m.IssueTime = v
	return m
}

// SetExpireTime sets ExpireTime value.
//
// Time alert expires
func (m *WeatherAlert) SetExpireTime(v time.Time) *WeatherAlert {
	m.ExpireTime = v
	return m
}

// SetSeverity sets Severity value.
//
// Warning, Watch, Advisory, Statement
func (m *WeatherAlert) SetSeverity(v typedef.WeatherSeverity) *WeatherAlert {
	m.Severity = v
	return m
}

// SetType sets Type value.
//
// Tornado, Severe Thunderstorm, etc.
func (m *WeatherAlert) SetType(v typedef.WeatherSevereType) *WeatherAlert {
	m.Type = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *WeatherAlert) SetUnknownFields(unknownFields ...proto.Field) *WeatherAlert {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *WeatherAlert) SetDeveloperFields(developerFields ...proto.DeveloperField) *WeatherAlert {
	m.DeveloperFields = developerFields
	return m
}
