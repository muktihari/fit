// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// TankSummary is a TankSummary message.
type TankSummary struct {
	Timestamp     typedef.DateTime // Units: s;
	Sensor        typedef.AntChannelId
	StartPressure uint16 // Scale: 100; Units: bar;
	EndPressure   uint16 // Scale: 100; Units: bar;
	VolumeUsed    uint32 // Scale: 100; Units: L;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTankSummary creates new TankSummary struct based on given mesg. If mesg is nil or mesg.Num is not equal to TankSummary mesg number, it will return nil.
func NewTankSummary(mesg proto.Message) *TankSummary {
	if mesg.Num != typedef.MesgNumTankSummary {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid,  /* Timestamp */
		0:   basetype.Uint32zInvalid, /* Sensor */
		1:   basetype.Uint16Invalid,  /* StartPressure */
		2:   basetype.Uint16Invalid,  /* EndPressure */
		3:   basetype.Uint32Invalid,  /* VolumeUsed */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &TankSummary{
		Timestamp:     typeconv.ToUint32[typedef.DateTime](vals[253]),
		Sensor:        typeconv.ToUint32z[typedef.AntChannelId](vals[0]),
		StartPressure: typeconv.ToUint16[uint16](vals[1]),
		EndPressure:   typeconv.ToUint16[uint16](vals[2]),
		VolumeUsed:    typeconv.ToUint32[uint32](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to TankSummary mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumTankSummary)
func (m TankSummary) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumTankSummary {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.Sensor,
		1:   m.StartPressure,
		2:   m.EndPressure,
		3:   m.VolumeUsed,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
