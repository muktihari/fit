// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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
	"math"
	"time"
)

// HrvValue is a HrvValue message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HrvValue struct {
	Timestamp time.Time
	Value     uint16 // Scale: 128; Units: ms; 5 minute RMSSD

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrvValue creates new HrvValue struct based on given mesg.
// If mesg is nil, it will return HrvValue with all fields being set to its corresponding invalid value.
func NewHrvValue(mesg *proto.Message) *HrvValue {
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

	return &HrvValue{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		Value:     vals[0].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HrvValue into proto.Message. If options is nil, default options will be used.
func (m *HrvValue) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHrvValue}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.Value != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.Value)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HrvValue) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// ValueScaled return Value in its scaled value.
// If Value value is invalid, float64 invalid value will be returned.
//
// Scale: 128; Units: ms; 5 minute RMSSD
func (m *HrvValue) ValueScaled() float64 {
	if m.Value == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.Value, 128, 0)
}

// SetTimestamp sets Timestamp value.
func (m *HrvValue) SetTimestamp(v time.Time) *HrvValue {
	m.Timestamp = v
	return m
}

// SetValue sets Value value.
//
// Scale: 128; Units: ms; 5 minute RMSSD
func (m *HrvValue) SetValue(v uint16) *HrvValue {
	m.Value = v
	return m
}

// SetValueScaled is similar to SetValue except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 128; Units: ms; 5 minute RMSSD
func (m *HrvValue) SetValueScaled(v float64) *HrvValue {
	m.Value = uint16(scaleoffset.Discard(v, 128, 0))
	return m
}

// SetDeveloperFields HrvValue's DeveloperFields.
func (m *HrvValue) SetDeveloperFields(developerFields ...proto.DeveloperField) *HrvValue {
	m.DeveloperFields = developerFields
	return m
}
