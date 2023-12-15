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

// AviationAttitude is a AviationAttitude message.
type AviationAttitude struct {
	Timestamp             typedef.DateTime           // Units: s; Timestamp message was output
	TimestampMs           uint16                     // Units: ms; Fractional part of timestamp, added to timestamp
	SystemTime            []uint32                   // Array: [N]; Units: ms; System time associated with sample expressed in ms.
	Pitch                 []int16                    // Scale: 10430.38; Array: [N]; Units: radians; Range -PI/2 to +PI/2
	Roll                  []int16                    // Scale: 10430.38; Array: [N]; Units: radians; Range -PI to +PI
	AccelLateral          []int16                    // Scale: 100; Array: [N]; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	AccelNormal           []int16                    // Scale: 100; Array: [N]; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	TurnRate              []int16                    // Scale: 1024; Array: [N]; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
	Stage                 []typedef.AttitudeStage    // Array: [N];
	AttitudeStageComplete []uint8                    // Array: [N]; Units: %; The percent complete of the current attitude stage. Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it. Range - 100
	Track                 []uint16                   // Scale: 10430.38; Array: [N]; Units: radians; Track Angle/Heading Range 0 - 2pi
	Validity              []typedef.AttitudeValidity // Array: [N];

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAviationAttitude creates new AviationAttitude struct based on given mesg. If mesg is nil or mesg.Num is not equal to AviationAttitude mesg number, it will return nil.
func NewAviationAttitude(mesg proto.Message) *AviationAttitude {
	if mesg.Num != typedef.MesgNumAviationAttitude {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* TimestampMs */
		1:   nil, /* SystemTime */
		2:   nil, /* Pitch */
		3:   nil, /* Roll */
		4:   nil, /* AccelLateral */
		5:   nil, /* AccelNormal */
		6:   nil, /* TurnRate */
		7:   nil, /* Stage */
		8:   nil, /* AttitudeStageComplete */
		9:   nil, /* Track */
		10:  nil, /* Validity */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &AviationAttitude{
		Timestamp:             typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs:           typeconv.ToUint16[uint16](vals[0]),
		SystemTime:            typeconv.ToSliceUint32[uint32](vals[1]),
		Pitch:                 typeconv.ToSliceSint16[int16](vals[2]),
		Roll:                  typeconv.ToSliceSint16[int16](vals[3]),
		AccelLateral:          typeconv.ToSliceSint16[int16](vals[4]),
		AccelNormal:           typeconv.ToSliceSint16[int16](vals[5]),
		TurnRate:              typeconv.ToSliceSint16[int16](vals[6]),
		Stage:                 typeconv.ToSliceEnum[typedef.AttitudeStage](vals[7]),
		AttitudeStageComplete: typeconv.ToSliceUint8[uint8](vals[8]),
		Track:                 typeconv.ToSliceUint16[uint16](vals[9]),
		Validity:              typeconv.ToSliceUint16[typedef.AttitudeValidity](vals[10]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to AviationAttitude mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumAviationAttitude)
func (m *AviationAttitude) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumAviationAttitude {
		return
	}

	vals := [...]any{
		253: typeconv.ToUint32[uint32](m.Timestamp),
		0:   m.TimestampMs,
		1:   m.SystemTime,
		2:   m.Pitch,
		3:   m.Roll,
		4:   m.AccelLateral,
		5:   m.AccelNormal,
		6:   m.TurnRate,
		7:   typeconv.ToSliceEnum[byte](m.Stage),
		8:   m.AttitudeStageComplete,
		9:   m.Track,
		10:  typeconv.ToSliceUint16[uint16](m.Validity),
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
