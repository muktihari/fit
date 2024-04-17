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

// HsaHeartRateData is a HsaHeartRateData message.
type HsaHeartRateData struct {
	Timestamp          time.Time // Units: s
	HeartRate          []uint8   // Array: [N]; Units: bpm; Beats / min
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds
	Status             uint8     // Status of measurements in buffer - 0 indicates SEARCHING 1 indicates LOCKED

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaHeartRateData creates new HsaHeartRateData struct based on given mesg.
// If mesg is nil, it will return HsaHeartRateData with all fields being set to its corresponding invalid value.
func NewHsaHeartRateData(mesg *proto.Message) *HsaHeartRateData {
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

	return &HsaHeartRateData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		HeartRate:          vals[2].SliceUint8(),
		ProcessingInterval: vals[0].Uint16(),
		Status:             vals[1].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaHeartRateData into proto.Message. If options is nil, default options will be used.
func (m *HsaHeartRateData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHsaHeartRateData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.HeartRate != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint8(m.HeartRate)
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ProcessingInterval)
		fields = append(fields, field)
	}
	if m.Status != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.Status)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaHeartRateData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets HsaHeartRateData value.
//
// Units: s
func (m *HsaHeartRateData) SetTimestamp(v time.Time) *HsaHeartRateData {
	m.Timestamp = v
	return m
}

// SetHeartRate sets HsaHeartRateData value.
//
// Array: [N]; Units: bpm; Beats / min
func (m *HsaHeartRateData) SetHeartRate(v []uint8) *HsaHeartRateData {
	m.HeartRate = v
	return m
}

// SetProcessingInterval sets HsaHeartRateData value.
//
// Units: s; Processing interval length in seconds
func (m *HsaHeartRateData) SetProcessingInterval(v uint16) *HsaHeartRateData {
	m.ProcessingInterval = v
	return m
}

// SetStatus sets HsaHeartRateData value.
//
// Status of measurements in buffer - 0 indicates SEARCHING 1 indicates LOCKED
func (m *HsaHeartRateData) SetStatus(v uint8) *HsaHeartRateData {
	m.Status = v
	return m
}

// SetDeveloperFields HsaHeartRateData's DeveloperFields.
func (m *HsaHeartRateData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaHeartRateData {
	m.DeveloperFields = developerFields
	return m
}
