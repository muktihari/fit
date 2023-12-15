// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// WeatherAlert is a WeatherAlert message.
type WeatherAlert struct {
	Timestamp  typedef.DateTime
	ReportId   string                    // Unique identifier from GCS report ID string, length is 12
	IssueTime  typedef.DateTime          // Time alert was issued
	ExpireTime typedef.DateTime          // Time alert expires
	Severity   typedef.WeatherSeverity   // Warning, Watch, Advisory, Statement
	Type       typedef.WeatherSevereType // Tornado, Severe Thunderstorm, etc.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWeatherAlert creates new WeatherAlert struct based on given mesg. If mesg is nil or mesg.Num is not equal to WeatherAlert mesg number, it will return nil.
func NewWeatherAlert(mesg proto.Message) *WeatherAlert {
	if mesg.Num != typedef.MesgNumWeatherAlert {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* ReportId */
		1:   nil, /* IssueTime */
		2:   nil, /* ExpireTime */
		3:   nil, /* Severity */
		4:   nil, /* Type */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &WeatherAlert{
		Timestamp:  typeconv.ToUint32[typedef.DateTime](vals[253]),
		ReportId:   typeconv.ToString[string](vals[0]),
		IssueTime:  typeconv.ToUint32[typedef.DateTime](vals[1]),
		ExpireTime: typeconv.ToUint32[typedef.DateTime](vals[2]),
		Severity:   typeconv.ToEnum[typedef.WeatherSeverity](vals[3]),
		Type:       typeconv.ToEnum[typedef.WeatherSevereType](vals[4]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to WeatherAlert mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumWeatherAlert)
func (m *WeatherAlert) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumWeatherAlert {
		return
	}

	vals := [...]any{
		253: typeconv.ToUint32[uint32](m.Timestamp),
		0:   m.ReportId,
		1:   typeconv.ToUint32[uint32](m.IssueTime),
		2:   typeconv.ToUint32[uint32](m.ExpireTime),
		3:   typeconv.ToEnum[byte](m.Severity),
		4:   typeconv.ToEnum[byte](m.Type),
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		field.Value = vals[field.Num]
	}

	mesg.DeveloperFields = m.DeveloperFields
}
