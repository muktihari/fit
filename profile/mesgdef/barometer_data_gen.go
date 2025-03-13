// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// BarometerData is a BarometerData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type BarometerData struct {
	Timestamp        time.Time // Units: s; Whole second part of the timestamp
	SampleTimeOffset []uint16  // Array: [N]; Units: ms; Each time in the array describes the time at which the barometer sample with the corresponding index was taken. The samples may span across seconds. Array size must match the number of samples in baro_cal
	BaroPres         []uint32  // Array: [N]; Units: Pa; These are the raw ADC reading. The samples may span across seconds. A conversion will need to be done on this data once read.
	TimestampMs      uint16    // Units: ms; Millisecond part of the timestamp.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewBarometerData creates new BarometerData struct based on given mesg.
// If mesg is nil, it will return BarometerData with all fields being set to its corresponding invalid value.
func NewBarometerData(mesg *proto.Message) *BarometerData {
	m := new(BarometerData)
	m.Reset(mesg)
	return m
}

// Reset resets all BarometerData's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *BarometerData) Reset(mesg *proto.Message) {
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

	*m = BarometerData{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		TimestampMs:      vals[0].Uint16(),
		SampleTimeOffset: vals[1].SliceUint16(),
		BaroPres:         vals[2].SliceUint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts BarometerData into proto.Message. If options is nil, default options will be used.
func (m *BarometerData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumBarometerData}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.SampleTimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint16(m.SampleTimeOffset)
		fields = append(fields, field)
	}
	if m.BaroPres != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint32(m.BaroPres)
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
func (m *BarometerData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s; Whole second part of the timestamp
func (m *BarometerData) SetTimestamp(v time.Time) *BarometerData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *BarometerData) SetTimestampMs(v uint16) *BarometerData {
	m.TimestampMs = v
	return m
}

// SetSampleTimeOffset sets SampleTimeOffset value.
//
// Array: [N]; Units: ms; Each time in the array describes the time at which the barometer sample with the corresponding index was taken. The samples may span across seconds. Array size must match the number of samples in baro_cal
func (m *BarometerData) SetSampleTimeOffset(v []uint16) *BarometerData {
	m.SampleTimeOffset = v
	return m
}

// SetBaroPres sets BaroPres value.
//
// Array: [N]; Units: Pa; These are the raw ADC reading. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *BarometerData) SetBaroPres(v []uint32) *BarometerData {
	m.BaroPres = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *BarometerData) SetUnknownFields(unknownFields ...proto.Field) *BarometerData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *BarometerData) SetDeveloperFields(developerFields ...proto.DeveloperField) *BarometerData {
	m.DeveloperFields = developerFields
	return m
}
