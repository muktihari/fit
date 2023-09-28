// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

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

// Totals is a Totals message.
type Totals struct {
	MessageIndex typedef.MessageIndex
	Timestamp    typedef.DateTime // Units: s;
	TimerTime    uint32           // Units: s; Excludes pauses
	Distance     uint32           // Units: m;
	Calories     uint32           // Units: kcal;
	Sport        typedef.Sport
	ElapsedTime  uint32 // Units: s; Includes pauses
	Sessions     uint16
	ActiveTime   uint32 // Units: s;
	SportIndex   uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTotals creates new Totals struct based on given mesg. If mesg is nil or mesg.Num is not equal to Totals mesg number, it will return nil.
func NewTotals(mesg proto.Message) *Totals {
	if mesg.Num != typedef.MesgNumTotals {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint32Invalid, /* TimerTime */
		1:   basetype.Uint32Invalid, /* Distance */
		2:   basetype.Uint32Invalid, /* Calories */
		3:   basetype.EnumInvalid,   /* Sport */
		4:   basetype.Uint32Invalid, /* ElapsedTime */
		5:   basetype.Uint16Invalid, /* Sessions */
		6:   basetype.Uint32Invalid, /* ActiveTime */
		9:   basetype.Uint8Invalid,  /* SportIndex */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Totals{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Timestamp:    typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimerTime:    typeconv.ToUint32[uint32](vals[0]),
		Distance:     typeconv.ToUint32[uint32](vals[1]),
		Calories:     typeconv.ToUint32[uint32](vals[2]),
		Sport:        typeconv.ToEnum[typedef.Sport](vals[3]),
		ElapsedTime:  typeconv.ToUint32[uint32](vals[4]),
		Sessions:     typeconv.ToUint16[uint16](vals[5]),
		ActiveTime:   typeconv.ToUint32[uint32](vals[6]),
		SportIndex:   typeconv.ToUint8[uint8](vals[9]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Totals mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumTotals)
func (m Totals) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumTotals {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		253: m.Timestamp,
		0:   m.TimerTime,
		1:   m.Distance,
		2:   m.Calories,
		3:   m.Sport,
		4:   m.ElapsedTime,
		5:   m.Sessions,
		6:   m.ActiveTime,
		9:   m.SportIndex,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
