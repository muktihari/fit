// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

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

// BarometerData is a BarometerData message.
type BarometerData struct {
	Timestamp        typedef.DateTime // Units: s; Whole second part of the timestamp
	TimestampMs      uint16           // Units: ms; Millisecond part of the timestamp.
	SampleTimeOffset []uint16         // Array: [N]; Units: ms; Each time in the array describes the time at which the barometer sample with the corrosponding index was taken. The samples may span across seconds. Array size must match the number of samples in baro_cal
	BaroPres         []uint32         // Array: [N]; Units: Pa; These are the raw ADC reading. The samples may span across seconds. A conversion will need to be done on this data once read.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewBarometerData creates new BarometerData struct based on given mesg. If mesg is nil or mesg.Num is not equal to BarometerData mesg number, it will return nil.
func NewBarometerData(mesg proto.Message) *BarometerData {
	if mesg.Num != typedef.MesgNumBarometerData {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* TimestampMs */
		1:   nil,                    /* SampleTimeOffset */
		2:   nil,                    /* BaroPres */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &BarometerData{
		Timestamp:        typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		SampleTimeOffset: typeconv.ToSliceUint16[uint16](vals[1]),
		BaroPres:         typeconv.ToSliceUint32[uint32](vals[2]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to BarometerData mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumBarometerData)
func (m BarometerData) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumBarometerData {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.SampleTimeOffset,
		2:   m.BaroPres,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
