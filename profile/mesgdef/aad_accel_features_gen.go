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
	"math"
	"time"
)

// AadAccelFeatures is a AadAccelFeatures message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type AadAccelFeatures struct {
	Timestamp          time.Time
	EnergyTotal        uint32 // Total accelerometer energy in the interval
	Time               uint16 // Units: s; Time interval length in seconds
	ZeroCrossCnt       uint16 // Count of zero crossings
	TimeAboveThreshold uint16 // Scale: 25; Units: s; Total accelerometer time above threshold in the interval
	Instance           uint8  // Instance ID of zero crossing algorithm

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewAadAccelFeatures creates new AadAccelFeatures struct based on given mesg.
// If mesg is nil, it will return AadAccelFeatures with all fields being set to its corresponding invalid value.
func NewAadAccelFeatures(mesg *proto.Message) *AadAccelFeatures {
	m := new(AadAccelFeatures)
	m.Reset(mesg)
	return m
}

// Reset resets all AadAccelFeatures's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *AadAccelFeatures) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

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

	*m = AadAccelFeatures{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		Time:               vals[0].Uint16(),
		EnergyTotal:        vals[1].Uint32(),
		ZeroCrossCnt:       vals[2].Uint16(),
		Instance:           vals[3].Uint8(),
		TimeAboveThreshold: vals[4].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts AadAccelFeatures into proto.Message. If options is nil, default options will be used.
func (m *AadAccelFeatures) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumAadAccelFeatures}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Time != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.Time)
		fields = append(fields, field)
	}
	if m.EnergyTotal != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.EnergyTotal)
		fields = append(fields, field)
	}
	if m.ZeroCrossCnt != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.ZeroCrossCnt)
		fields = append(fields, field)
	}
	if m.Instance != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.Instance)
		fields = append(fields, field)
	}
	if m.TimeAboveThreshold != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.TimeAboveThreshold)
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
func (m *AadAccelFeatures) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// TimeAboveThresholdScaled return TimeAboveThreshold in its scaled value.
// If TimeAboveThreshold value is invalid, float64 invalid value will be returned.
//
// Scale: 25; Units: s; Total accelerometer time above threshold in the interval
func (m *AadAccelFeatures) TimeAboveThresholdScaled() float64 {
	if m.TimeAboveThreshold == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TimeAboveThreshold)/25 - 0
}

// SetTimestamp sets Timestamp value.
func (m *AadAccelFeatures) SetTimestamp(v time.Time) *AadAccelFeatures {
	m.Timestamp = v
	return m
}

// SetTime sets Time value.
//
// Units: s; Time interval length in seconds
func (m *AadAccelFeatures) SetTime(v uint16) *AadAccelFeatures {
	m.Time = v
	return m
}

// SetEnergyTotal sets EnergyTotal value.
//
// Total accelerometer energy in the interval
func (m *AadAccelFeatures) SetEnergyTotal(v uint32) *AadAccelFeatures {
	m.EnergyTotal = v
	return m
}

// SetZeroCrossCnt sets ZeroCrossCnt value.
//
// Count of zero crossings
func (m *AadAccelFeatures) SetZeroCrossCnt(v uint16) *AadAccelFeatures {
	m.ZeroCrossCnt = v
	return m
}

// SetInstance sets Instance value.
//
// Instance ID of zero crossing algorithm
func (m *AadAccelFeatures) SetInstance(v uint8) *AadAccelFeatures {
	m.Instance = v
	return m
}

// SetTimeAboveThreshold sets TimeAboveThreshold value.
//
// Scale: 25; Units: s; Total accelerometer time above threshold in the interval
func (m *AadAccelFeatures) SetTimeAboveThreshold(v uint16) *AadAccelFeatures {
	m.TimeAboveThreshold = v
	return m
}

// SetTimeAboveThresholdScaled is similar to SetTimeAboveThreshold except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 25; Units: s; Total accelerometer time above threshold in the interval
func (m *AadAccelFeatures) SetTimeAboveThresholdScaled(v float64) *AadAccelFeatures {
	unscaled := (v + 0) * 25
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.TimeAboveThreshold = uint16(basetype.Uint16Invalid)
		return m
	}
	m.TimeAboveThreshold = uint16(unscaled)
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *AadAccelFeatures) SetUnknownFields(unknownFields ...proto.Field) *AadAccelFeatures {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *AadAccelFeatures) SetDeveloperFields(developerFields ...proto.DeveloperField) *AadAccelFeatures {
	m.DeveloperFields = developerFields
	return m
}
