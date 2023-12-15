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

// Hr is a Hr message.
type Hr struct {
	Timestamp           typedef.DateTime
	FractionalTimestamp uint16   // Scale: 32768; Units: s;
	Time256             uint8    // Scale: 256; Units: s;
	FilteredBpm         []uint8  // Array: [N]; Units: bpm;
	EventTimestamp      []uint32 // Scale: 1024; Array: [N]; Units: s;
	EventTimestamp12    []byte   // Scale: 1024; Array: [N]; Units: s;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHr creates new Hr struct based on given mesg. If mesg is nil or mesg.Num is not equal to Hr mesg number, it will return nil.
func NewHr(mesg proto.Message) *Hr {
	if mesg.Num != typedef.MesgNumHr {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* FractionalTimestamp */
		1:   nil, /* Time256 */
		6:   nil, /* FilteredBpm */
		9:   nil, /* EventTimestamp */
		10:  nil, /* EventTimestamp12 */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &Hr{
		Timestamp:           typeconv.ToUint32[typedef.DateTime](vals[253]),
		FractionalTimestamp: typeconv.ToUint16[uint16](vals[0]),
		Time256:             typeconv.ToUint8[uint8](vals[1]),
		FilteredBpm:         typeconv.ToSliceUint8[uint8](vals[6]),
		EventTimestamp:      typeconv.ToSliceUint32[uint32](vals[9]),
		EventTimestamp12:    typeconv.ToSliceByte[byte](vals[10]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Hr mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumHr)
func (m *Hr) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumHr {
		return
	}

	vals := [...]any{
		253: typeconv.ToUint32[uint32](m.Timestamp),
		0:   m.FractionalTimestamp,
		1:   m.Time256,
		6:   m.FilteredBpm,
		9:   m.EventTimestamp,
		10:  m.EventTimestamp12,
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
