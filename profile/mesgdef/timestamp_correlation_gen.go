// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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
	"time"
)

// TimestampCorrelation is a TimestampCorrelation message.
type TimestampCorrelation struct {
	Timestamp                 time.Time // Units: s; Whole second part of UTC timestamp at the time the system timestamp was recorded.
	SystemTimestamp           time.Time // Units: s; Whole second part of the system timestamp
	LocalTimestamp            time.Time // Units: s; timestamp epoch expressed in local time used to convert timestamps to local time
	FractionalTimestamp       uint16    // Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.
	FractionalSystemTimestamp uint16    // Scale: 32768; Units: s; Fractional part of the system timestamp
	TimestampMs               uint16    // Units: ms; Millisecond part of the UTC timestamp at the time the system timestamp was recorded.
	SystemTimestampMs         uint16    // Units: ms; Millisecond part of the system timestamp

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTimestampCorrelation creates new TimestampCorrelation struct based on given mesg.
// If mesg is nil, it will return TimestampCorrelation with all fields being set to its corresponding invalid value.
func NewTimestampCorrelation(mesg *proto.Message) *TimestampCorrelation {
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

	return &TimestampCorrelation{
		Timestamp:                 datetime.ToTime(vals[253].Uint32()),
		SystemTimestamp:           datetime.ToTime(vals[1].Uint32()),
		LocalTimestamp:            datetime.ToTime(vals[3].Uint32()),
		FractionalTimestamp:       vals[0].Uint16(),
		FractionalSystemTimestamp: vals[2].Uint16(),
		TimestampMs:               vals[4].Uint16(),
		SystemTimestampMs:         vals[5].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts TimestampCorrelation into proto.Message. If options is nil, default options will be used.
func (m *TimestampCorrelation) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumTimestampCorrelation)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.SystemTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(datetime.ToUint32(m.SystemTimestamp))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.LocalTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(datetime.ToUint32(m.LocalTimestamp))
		fields = append(fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.FractionalTimestamp)
		fields = append(fields, field)
	}
	if m.FractionalSystemTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.FractionalSystemTimestamp)
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.SystemTimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.SystemTimestampMs)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// FractionalTimestampScaled return FractionalTimestamp in its scaled value [Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.].
//
// If FractionalTimestamp value is invalid, float64 invalid value will be returned.
func (m *TimestampCorrelation) FractionalTimestampScaled() float64 {
	if m.FractionalTimestamp == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.FractionalTimestamp, 32768, 0)
}

// FractionalSystemTimestampScaled return FractionalSystemTimestamp in its scaled value [Scale: 32768; Units: s; Fractional part of the system timestamp].
//
// If FractionalSystemTimestamp value is invalid, float64 invalid value will be returned.
func (m *TimestampCorrelation) FractionalSystemTimestampScaled() float64 {
	if m.FractionalSystemTimestamp == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.FractionalSystemTimestamp, 32768, 0)
}

// SetTimestamp sets TimestampCorrelation value.
//
// Units: s; Whole second part of UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetTimestamp(v time.Time) *TimestampCorrelation {
	m.Timestamp = v
	return m
}

// SetSystemTimestamp sets TimestampCorrelation value.
//
// Units: s; Whole second part of the system timestamp
func (m *TimestampCorrelation) SetSystemTimestamp(v time.Time) *TimestampCorrelation {
	m.SystemTimestamp = v
	return m
}

// SetLocalTimestamp sets TimestampCorrelation value.
//
// Units: s; timestamp epoch expressed in local time used to convert timestamps to local time
func (m *TimestampCorrelation) SetLocalTimestamp(v time.Time) *TimestampCorrelation {
	m.LocalTimestamp = v
	return m
}

// SetFractionalTimestamp sets TimestampCorrelation value.
//
// Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetFractionalTimestamp(v uint16) *TimestampCorrelation {
	m.FractionalTimestamp = v
	return m
}

// SetFractionalSystemTimestamp sets TimestampCorrelation value.
//
// Scale: 32768; Units: s; Fractional part of the system timestamp
func (m *TimestampCorrelation) SetFractionalSystemTimestamp(v uint16) *TimestampCorrelation {
	m.FractionalSystemTimestamp = v
	return m
}

// SetTimestampMs sets TimestampCorrelation value.
//
// Units: ms; Millisecond part of the UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetTimestampMs(v uint16) *TimestampCorrelation {
	m.TimestampMs = v
	return m
}

// SetSystemTimestampMs sets TimestampCorrelation value.
//
// Units: ms; Millisecond part of the system timestamp
func (m *TimestampCorrelation) SetSystemTimestampMs(v uint16) *TimestampCorrelation {
	m.SystemTimestampMs = v
	return m
}

// SetDeveloperFields TimestampCorrelation's DeveloperFields.
func (m *TimestampCorrelation) SetDeveloperFields(developerFields ...proto.DeveloperField) *TimestampCorrelation {
	m.DeveloperFields = developerFields
	return m
}
