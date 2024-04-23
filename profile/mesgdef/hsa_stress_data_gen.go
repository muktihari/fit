// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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

// HsaStressData is a HsaStressData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaStressData struct {
	Timestamp          time.Time
	StressLevel        []int8 // Array: [N]; Units: s; Stress Level ( 0 - 100 ) -300 indicates invalid -200 indicates large motion -100 indicates off wrist
	ProcessingInterval uint16 // Units: s; Processing interval length in seconds

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaStressData creates new HsaStressData struct based on given mesg.
// If mesg is nil, it will return HsaStressData with all fields being set to its corresponding invalid value.
func NewHsaStressData(mesg *proto.Message) *HsaStressData {
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

	return &HsaStressData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		ProcessingInterval: vals[0].Uint16(),
		StressLevel:        vals[1].SliceInt8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaStressData into proto.Message. If options is nil, default options will be used.
func (m *HsaStressData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHsaStressData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ProcessingInterval)
		fields = append(fields, field)
	}
	if m.StressLevel != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceInt8(m.StressLevel)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaStressData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets HsaStressData value.
func (m *HsaStressData) SetTimestamp(v time.Time) *HsaStressData {
	m.Timestamp = v
	return m
}

// SetProcessingInterval sets HsaStressData value.
//
// Units: s; Processing interval length in seconds
func (m *HsaStressData) SetProcessingInterval(v uint16) *HsaStressData {
	m.ProcessingInterval = v
	return m
}

// SetStressLevel sets HsaStressData value.
//
// Array: [N]; Units: s; Stress Level ( 0 - 100 ) -300 indicates invalid -200 indicates large motion -100 indicates off wrist
func (m *HsaStressData) SetStressLevel(v []int8) *HsaStressData {
	m.StressLevel = v
	return m
}

// SetDeveloperFields HsaStressData's DeveloperFields.
func (m *HsaStressData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaStressData {
	m.DeveloperFields = developerFields
	return m
}
