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

// ObdiiData is a ObdiiData message.
type ObdiiData struct {
	Timestamp        typedef.DateTime // Units: s; Timestamp message was output
	TimestampMs      uint16           // Units: ms; Fractional part of timestamp, added to timestamp
	TimeOffset       []uint16         // Array: [N]; Units: ms; Offset of PID reading [i] from start_timestamp+start_timestamp_ms. Readings may span accross seconds.
	Pid              byte             // Parameter ID
	RawData          []byte           // Array: [N]; Raw parameter data
	PidDataSize      []uint8          // Array: [N]; Optional, data size of PID[i]. If not specified refer to SAE J1979.
	SystemTime       []uint32         // Array: [N]; System time associated with sample expressed in ms, can be used instead of time_offset. There will be a system_time value for each raw_data element. For multibyte pids the system_time is repeated.
	StartTimestamp   typedef.DateTime // Timestamp of first sample recorded in the message. Used with time_offset to generate time of each sample
	StartTimestampMs uint16           // Units: ms; Fractional part of start_timestamp

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewObdiiData creates new ObdiiData struct based on given mesg. If mesg is nil or mesg.Num is not equal to ObdiiData mesg number, it will return nil.
func NewObdiiData(mesg proto.Message) *ObdiiData {
	if mesg.Num != typedef.MesgNumObdiiData {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* TimestampMs */
		1:   nil, /* TimeOffset */
		2:   nil, /* Pid */
		3:   nil, /* RawData */
		4:   nil, /* PidDataSize */
		5:   nil, /* SystemTime */
		6:   nil, /* StartTimestamp */
		7:   nil, /* StartTimestampMs */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &ObdiiData{
		Timestamp:        typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		TimeOffset:       typeconv.ToSliceUint16[uint16](vals[1]),
		Pid:              typeconv.ToByte[byte](vals[2]),
		RawData:          typeconv.ToSliceByte[byte](vals[3]),
		PidDataSize:      typeconv.ToSliceUint8[uint8](vals[4]),
		SystemTime:       typeconv.ToSliceUint32[uint32](vals[5]),
		StartTimestamp:   typeconv.ToUint32[typedef.DateTime](vals[6]),
		StartTimestampMs: typeconv.ToUint16[uint16](vals[7]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to ObdiiData mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumObdiiData)
func (m ObdiiData) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumObdiiData {
		return
	}

	vals := [...]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.TimeOffset,
		2:   m.Pid,
		3:   m.RawData,
		4:   m.PidDataSize,
		5:   m.SystemTime,
		6:   m.StartTimestamp,
		7:   m.StartTimestampMs,
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
