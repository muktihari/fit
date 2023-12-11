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

// Schedule is a Schedule message.
type Schedule struct {
	Manufacturer  typedef.Manufacturer // Corresponds to file_id of scheduled workout / course.
	Product       uint16               // Corresponds to file_id of scheduled workout / course.
	SerialNumber  uint32               // Corresponds to file_id of scheduled workout / course.
	TimeCreated   typedef.DateTime     // Corresponds to file_id of scheduled workout / course.
	Completed     bool                 // TRUE if this activity has been started
	Type          typedef.Schedule
	ScheduledTime typedef.LocalDateTime

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSchedule creates new Schedule struct based on given mesg. If mesg is nil or mesg.Num is not equal to Schedule mesg number, it will return nil.
func NewSchedule(mesg proto.Message) *Schedule {
	if mesg.Num != typedef.MesgNumSchedule {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		0: nil, /* Manufacturer */
		1: nil, /* Product */
		2: nil, /* SerialNumber */
		3: nil, /* TimeCreated */
		4: nil, /* Completed */
		5: nil, /* Type */
		6: nil, /* ScheduledTime */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &Schedule{
		Manufacturer:  typeconv.ToUint16[typedef.Manufacturer](vals[0]),
		Product:       typeconv.ToUint16[uint16](vals[1]),
		SerialNumber:  typeconv.ToUint32z[uint32](vals[2]),
		TimeCreated:   typeconv.ToUint32[typedef.DateTime](vals[3]),
		Completed:     typeconv.ToBool[bool](vals[4]),
		Type:          typeconv.ToEnum[typedef.Schedule](vals[5]),
		ScheduledTime: typeconv.ToUint32[typedef.LocalDateTime](vals[6]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Schedule mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSchedule)
func (m Schedule) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSchedule {
		return
	}

	vals := [...]any{
		0: m.Manufacturer,
		1: m.Product,
		2: m.SerialNumber,
		3: m.TimeCreated,
		4: m.Completed,
		5: m.Type,
		6: m.ScheduledTime,
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
