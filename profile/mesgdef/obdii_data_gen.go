// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// ObdiiData is a ObdiiData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ObdiiData struct {
	Timestamp        time.Time // Units: s; Timestamp message was output
	TimeOffset       []uint16  // Array: [N]; Units: ms; Offset of PID reading [i] from start_timestamp+start_timestamp_ms. Readings may span across seconds.
	RawData          []byte    // Array: [N]; Raw parameter data
	PidDataSize      []uint8   // Array: [N]; Optional, data size of PID[i]. If not specified refer to SAE J1979.
	SystemTime       []uint32  // Array: [N]; System time associated with sample expressed in ms, can be used instead of time_offset. There will be a system_time value for each raw_data element. For multibyte pids the system_time is repeated.
	StartTimestamp   time.Time // Timestamp of first sample recorded in the message. Used with time_offset to generate time of each sample
	TimestampMs      uint16    // Units: ms; Fractional part of timestamp, added to timestamp
	StartTimestampMs uint16    // Units: ms; Fractional part of start_timestamp
	Pid              byte      // Parameter ID

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewObdiiData creates new ObdiiData struct based on given mesg.
// If mesg is nil, it will return ObdiiData with all fields being set to its corresponding invalid value.
func NewObdiiData(mesg *proto.Message) *ObdiiData {
	m := new(ObdiiData)
	m.Reset(mesg)
	return m
}

// Reset resets all ObdiiData's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *ObdiiData) Reset(mesg *proto.Message) {
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

	*m = ObdiiData{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		TimestampMs:      vals[0].Uint16(),
		TimeOffset:       vals[1].SliceUint16(),
		Pid:              vals[2].Uint8(),
		RawData:          vals[3].SliceUint8(),
		PidDataSize:      vals[4].SliceUint8(),
		SystemTime:       vals[5].SliceUint32(),
		StartTimestamp:   datetime.ToTime(vals[6].Uint32()),
		StartTimestampMs: vals[7].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts ObdiiData into proto.Message. If options is nil, default options will be used.
func (m *ObdiiData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fields := make([]proto.Field, 0, 9)
	mesg := proto.Message{Num: typedef.MesgNumObdiiData}

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
	if m.TimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint16(m.TimeOffset)
		fields = append(fields, field)
	}
	if m.Pid != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.Pid)
		fields = append(fields, field)
	}
	if m.RawData != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceUint8(m.RawData)
		fields = append(fields, field)
	}
	if m.PidDataSize != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint8(m.PidDataSize)
		fields = append(fields, field)
	}
	if m.SystemTime != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceUint32(m.SystemTime)
		fields = append(fields, field)
	}
	if !m.StartTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(uint32(m.StartTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.StartTimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint16(m.StartTimestampMs)
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
func (m *ObdiiData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// StartTimestampUint32 returns StartTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *ObdiiData) StartTimestampUint32() uint32 { return datetime.ToUint32(m.StartTimestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s; Timestamp message was output
func (m *ObdiiData) SetTimestamp(v time.Time) *ObdiiData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Fractional part of timestamp, added to timestamp
func (m *ObdiiData) SetTimestampMs(v uint16) *ObdiiData {
	m.TimestampMs = v
	return m
}

// SetTimeOffset sets TimeOffset value.
//
// Array: [N]; Units: ms; Offset of PID reading [i] from start_timestamp+start_timestamp_ms. Readings may span across seconds.
func (m *ObdiiData) SetTimeOffset(v []uint16) *ObdiiData {
	m.TimeOffset = v
	return m
}

// SetPid sets Pid value.
//
// Parameter ID
func (m *ObdiiData) SetPid(v byte) *ObdiiData {
	m.Pid = v
	return m
}

// SetRawData sets RawData value.
//
// Array: [N]; Raw parameter data
func (m *ObdiiData) SetRawData(v []byte) *ObdiiData {
	m.RawData = v
	return m
}

// SetPidDataSize sets PidDataSize value.
//
// Array: [N]; Optional, data size of PID[i]. If not specified refer to SAE J1979.
func (m *ObdiiData) SetPidDataSize(v []uint8) *ObdiiData {
	m.PidDataSize = v
	return m
}

// SetSystemTime sets SystemTime value.
//
// Array: [N]; System time associated with sample expressed in ms, can be used instead of time_offset. There will be a system_time value for each raw_data element. For multibyte pids the system_time is repeated.
func (m *ObdiiData) SetSystemTime(v []uint32) *ObdiiData {
	m.SystemTime = v
	return m
}

// SetStartTimestamp sets StartTimestamp value.
//
// Timestamp of first sample recorded in the message. Used with time_offset to generate time of each sample
func (m *ObdiiData) SetStartTimestamp(v time.Time) *ObdiiData {
	m.StartTimestamp = v
	return m
}

// SetStartTimestampMs sets StartTimestampMs value.
//
// Units: ms; Fractional part of start_timestamp
func (m *ObdiiData) SetStartTimestampMs(v uint16) *ObdiiData {
	m.StartTimestampMs = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *ObdiiData) SetUnknownFields(unknownFields ...proto.Field) *ObdiiData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *ObdiiData) SetDeveloperFields(developerFields ...proto.DeveloperField) *ObdiiData {
	m.DeveloperFields = developerFields
	return m
}
