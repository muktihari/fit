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

// NmeaSentence is a NmeaSentence message.
type NmeaSentence struct {
	Timestamp   typedef.DateTime // Units: s; Timestamp message was output
	TimestampMs uint16           // Units: ms; Fractional part of timestamp, added to timestamp
	Sentence    string           // NMEA sentence

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewNmeaSentence creates new NmeaSentence struct based on given mesg. If mesg is nil or mesg.Num is not equal to NmeaSentence mesg number, it will return nil.
func NewNmeaSentence(mesg proto.Message) *NmeaSentence {
	if mesg.Num != typedef.MesgNumNmeaSentence {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* TimestampMs */
		1:   nil, /* Sentence */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &NmeaSentence{
		Timestamp:   typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs: typeconv.ToUint16[uint16](vals[0]),
		Sentence:    typeconv.ToString[string](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to NmeaSentence mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumNmeaSentence)
func (m NmeaSentence) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumNmeaSentence {
		return
	}

	vals := [...]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.Sentence,
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
