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

// VideoFrame is a VideoFrame message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type VideoFrame struct {
	Timestamp   time.Time // Units: s; Whole second part of the timestamp
	FrameNumber uint32    // Number of the frame that the timestamp and timestamp_ms correlate to
	TimestampMs uint16    // Units: ms; Millisecond part of the timestamp.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideoFrame creates new VideoFrame struct based on given mesg.
// If mesg is nil, it will return VideoFrame with all fields being set to its corresponding invalid value.
func NewVideoFrame(mesg *proto.Message) *VideoFrame {
	vals := [254]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &VideoFrame{
		Timestamp:   datetime.ToTime(vals[253].Uint32()),
		TimestampMs: vals[0].Uint16(),
		FrameNumber: vals[1].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts VideoFrame into proto.Message. If options is nil, default options will be used.
func (m *VideoFrame) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumVideoFrame}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.FrameNumber != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.FrameNumber)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *VideoFrame) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s; Whole second part of the timestamp
func (m *VideoFrame) SetTimestamp(v time.Time) *VideoFrame {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *VideoFrame) SetTimestampMs(v uint16) *VideoFrame {
	m.TimestampMs = v
	return m
}

// SetFrameNumber sets FrameNumber value.
//
// Number of the frame that the timestamp and timestamp_ms correlate to
func (m *VideoFrame) SetFrameNumber(v uint32) *VideoFrame {
	m.FrameNumber = v
	return m
}

// SetDeveloperFields VideoFrame's DeveloperFields.
func (m *VideoFrame) SetDeveloperFields(developerFields ...proto.DeveloperField) *VideoFrame {
	m.DeveloperFields = developerFields
	return m
}
