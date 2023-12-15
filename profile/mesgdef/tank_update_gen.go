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

// TankUpdate is a TankUpdate message.
type TankUpdate struct {
	Timestamp typedef.DateTime // Units: s;
	Sensor    typedef.AntChannelId
	Pressure  uint16 // Scale: 100; Units: bar;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTankUpdate creates new TankUpdate struct based on given mesg. If mesg is nil or mesg.Num is not equal to TankUpdate mesg number, it will return nil.
func NewTankUpdate(mesg proto.Message) *TankUpdate {
	if mesg.Num != typedef.MesgNumTankUpdate {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* Sensor */
		1:   nil, /* Pressure */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &TankUpdate{
		Timestamp: typeconv.ToUint32[typedef.DateTime](vals[253]),
		Sensor:    typeconv.ToUint32z[typedef.AntChannelId](vals[0]),
		Pressure:  typeconv.ToUint16[uint16](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to TankUpdate mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumTankUpdate)
func (m *TankUpdate) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumTankUpdate {
		return
	}

	vals := [...]any{
		253: typeconv.ToUint32[uint32](m.Timestamp),
		0:   typeconv.ToUint32z[uint32](m.Sensor),
		1:   m.Pressure,
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
