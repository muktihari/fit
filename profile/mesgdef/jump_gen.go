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

// Jump is a Jump message.
type Jump struct {
	Timestamp     typedef.DateTime // Units: s;
	Distance      float32          // Units: m;
	Height        float32          // Units: m;
	Rotations     uint8
	HangTime      float32 // Units: s;
	Score         float32 // A score for a jump calculated based on hang time, rotations, and distance.
	PositionLat   int32   // Units: semicircles;
	PositionLong  int32   // Units: semicircles;
	Speed         uint16  // Scale: 1000; Units: m/s;
	EnhancedSpeed uint32  // Scale: 1000; Units: m/s;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewJump creates new Jump struct based on given mesg. If mesg is nil or mesg.Num is not equal to Jump mesg number, it will return nil.
func NewJump(mesg proto.Message) *Jump {
	if mesg.Num != typedef.MesgNumJump {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* Distance */
		1:   nil, /* Height */
		2:   nil, /* Rotations */
		3:   nil, /* HangTime */
		4:   nil, /* Score */
		5:   nil, /* PositionLat */
		6:   nil, /* PositionLong */
		7:   nil, /* Speed */
		8:   nil, /* EnhancedSpeed */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &Jump{
		Timestamp:     typeconv.ToUint32[typedef.DateTime](vals[253]),
		Distance:      typeconv.ToFloat32[float32](vals[0]),
		Height:        typeconv.ToFloat32[float32](vals[1]),
		Rotations:     typeconv.ToUint8[uint8](vals[2]),
		HangTime:      typeconv.ToFloat32[float32](vals[3]),
		Score:         typeconv.ToFloat32[float32](vals[4]),
		PositionLat:   typeconv.ToSint32[int32](vals[5]),
		PositionLong:  typeconv.ToSint32[int32](vals[6]),
		Speed:         typeconv.ToUint16[uint16](vals[7]),
		EnhancedSpeed: typeconv.ToUint32[uint32](vals[8]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Jump mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumJump)
func (m *Jump) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumJump {
		return
	}

	vals := [...]any{
		253: typeconv.ToUint32[uint32](m.Timestamp),
		0:   m.Distance,
		1:   m.Height,
		2:   m.Rotations,
		3:   m.HangTime,
		4:   m.Score,
		5:   m.PositionLat,
		6:   m.PositionLong,
		7:   m.Speed,
		8:   m.EnhancedSpeed,
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
