// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// WeatherAlert is a WeatherAlert message.
type WeatherAlert struct {
	Timestamp  time.Time
	ReportId   string                    // Unique identifier from GCS report ID string, length is 12
	IssueTime  time.Time                 // Time alert was issued
	ExpireTime time.Time                 // Time alert expires
	Severity   typedef.WeatherSeverity   // Warning, Watch, Advisory, Statement
	Type       typedef.WeatherSevereType // Tornado, Severe Thunderstorm, etc.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWeatherAlert creates new WeatherAlert struct based on given mesg.
// If mesg is nil, it will return WeatherAlert with all fields being set to its corresponding invalid value.
func NewWeatherAlert(mesg *proto.Message) *WeatherAlert {
	vals := [254]any{}

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

	return &WeatherAlert{
		Timestamp:  datetime.ToTime(vals[253]),
		ReportId:   typeconv.ToString[string](vals[0]),
		IssueTime:  datetime.ToTime(vals[1]),
		ExpireTime: datetime.ToTime(vals[2]),
		Severity:   typeconv.ToEnum[typedef.WeatherSeverity](vals[3]),
		Type:       typeconv.ToEnum[typedef.WeatherSevereType](vals[4]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WeatherAlert into proto.Message.
func (m *WeatherAlert) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumWeatherAlert)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ReportId != basetype.StringInvalid && m.ReportId != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ReportId
		mesg.Fields = append(mesg.Fields, field)
	}
	if datetime.ToUint32(m.IssueTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = datetime.ToUint32(m.IssueTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if datetime.ToUint32(m.ExpireTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = datetime.ToUint32(m.ExpireTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Severity) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = typeconv.ToEnum[byte](m.Severity)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = typeconv.ToEnum[byte](m.Type)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of WeatherAlert's valid fields.
func (m *WeatherAlert) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if m.ReportId != basetype.StringInvalid && m.ReportId != "" {
		size++
	}
	if datetime.ToUint32(m.IssueTime) != basetype.Uint32Invalid {
		size++
	}
	if datetime.ToUint32(m.ExpireTime) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Severity) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Type) != basetype.EnumInvalid {
		size++
	}
	return size
}

// SetTimestamp sets WeatherAlert value.
func (m *WeatherAlert) SetTimestamp(v time.Time) *WeatherAlert {
	m.Timestamp = v
	return m
}

// SetReportId sets WeatherAlert value.
//
// Unique identifier from GCS report ID string, length is 12
func (m *WeatherAlert) SetReportId(v string) *WeatherAlert {
	m.ReportId = v
	return m
}

// SetIssueTime sets WeatherAlert value.
//
// Time alert was issued
func (m *WeatherAlert) SetIssueTime(v time.Time) *WeatherAlert {
	m.IssueTime = v
	return m
}

// SetExpireTime sets WeatherAlert value.
//
// Time alert expires
func (m *WeatherAlert) SetExpireTime(v time.Time) *WeatherAlert {
	m.ExpireTime = v
	return m
}

// SetSeverity sets WeatherAlert value.
//
// Warning, Watch, Advisory, Statement
func (m *WeatherAlert) SetSeverity(v typedef.WeatherSeverity) *WeatherAlert {
	m.Severity = v
	return m
}

// SetType sets WeatherAlert value.
//
// Tornado, Severe Thunderstorm, etc.
func (m *WeatherAlert) SetType(v typedef.WeatherSevereType) *WeatherAlert {
	m.Type = v
	return m
}

// SetDeveloperFields WeatherAlert's DeveloperFields.
func (m *WeatherAlert) SetDeveloperFields(developerFields ...proto.DeveloperField) *WeatherAlert {
	m.DeveloperFields = developerFields
	return m
}
