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

// VideoClip is a VideoClip message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type VideoClip struct {
	StartTimestamp   time.Time
	EndTimestamp     time.Time
	ClipStart        uint32 // Units: ms; Start of clip in video time
	ClipEnd          uint32 // Units: ms; End of clip in video time
	ClipNumber       uint16
	StartTimestampMs uint16
	EndTimestampMs   uint16

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewVideoClip creates new VideoClip struct based on given mesg.
// If mesg is nil, it will return VideoClip with all fields being set to its corresponding invalid value.
func NewVideoClip(mesg *proto.Message) *VideoClip {
	m := new(VideoClip)
	m.Reset(mesg)
	return m
}

// Reset resets all VideoClip's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *VideoClip) Reset(mesg *proto.Message) {
	var (
		vals            [8]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 7 || mesg.Fields[i].Name == factory.NameUnknown {
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

	*m = VideoClip{
		ClipNumber:       vals[0].Uint16(),
		StartTimestamp:   datetime.ToTime(vals[1].Uint32()),
		StartTimestampMs: vals[2].Uint16(),
		EndTimestamp:     datetime.ToTime(vals[3].Uint32()),
		EndTimestampMs:   vals[4].Uint16(),
		ClipStart:        vals[6].Uint32(),
		ClipEnd:          vals[7].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts VideoClip into proto.Message. If options is nil, default options will be used.
func (m *VideoClip) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumVideoClip}

	if m.ClipNumber != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ClipNumber)
		fields = append(fields, field)
	}
	if !m.StartTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(uint32(m.StartTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.StartTimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.StartTimestampMs)
		fields = append(fields, field)
	}
	if !m.EndTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(uint32(m.EndTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.EndTimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.EndTimestampMs)
		fields = append(fields, field)
	}
	if m.ClipStart != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.ClipStart)
		fields = append(fields, field)
	}
	if m.ClipEnd != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint32(m.ClipEnd)
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

// StartTimestampUint32 returns StartTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *VideoClip) StartTimestampUint32() uint32 { return datetime.ToUint32(m.StartTimestamp) }

// EndTimestampUint32 returns EndTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *VideoClip) EndTimestampUint32() uint32 { return datetime.ToUint32(m.EndTimestamp) }

// SetClipNumber sets ClipNumber value.
func (m *VideoClip) SetClipNumber(v uint16) *VideoClip {
	m.ClipNumber = v
	return m
}

// SetStartTimestamp sets StartTimestamp value.
func (m *VideoClip) SetStartTimestamp(v time.Time) *VideoClip {
	m.StartTimestamp = v
	return m
}

// SetStartTimestampMs sets StartTimestampMs value.
func (m *VideoClip) SetStartTimestampMs(v uint16) *VideoClip {
	m.StartTimestampMs = v
	return m
}

// SetEndTimestamp sets EndTimestamp value.
func (m *VideoClip) SetEndTimestamp(v time.Time) *VideoClip {
	m.EndTimestamp = v
	return m
}

// SetEndTimestampMs sets EndTimestampMs value.
func (m *VideoClip) SetEndTimestampMs(v uint16) *VideoClip {
	m.EndTimestampMs = v
	return m
}

// SetClipStart sets ClipStart value.
//
// Units: ms; Start of clip in video time
func (m *VideoClip) SetClipStart(v uint32) *VideoClip {
	m.ClipStart = v
	return m
}

// SetClipEnd sets ClipEnd value.
//
// Units: ms; End of clip in video time
func (m *VideoClip) SetClipEnd(v uint32) *VideoClip {
	m.ClipEnd = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *VideoClip) SetUnknownFields(unknownFields ...proto.Field) *VideoClip {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *VideoClip) SetDeveloperFields(developerFields ...proto.DeveloperField) *VideoClip {
	m.DeveloperFields = developerFields
	return m
}
