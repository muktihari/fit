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

// GpsMetadata is a GpsMetadata message.
type GpsMetadata struct {
	Timestamp        typedef.DateTime // Units: s; Whole second part of the timestamp.
	TimestampMs      uint16           // Units: ms; Millisecond part of the timestamp.
	PositionLat      int32            // Units: semicircles;
	PositionLong     int32            // Units: semicircles;
	EnhancedAltitude uint32           // Scale: 5; Offset: 500; Units: m;
	EnhancedSpeed    uint32           // Scale: 1000; Units: m/s;
	Heading          uint16           // Scale: 100; Units: degrees;
	UtcTimestamp     typedef.DateTime // Units: s; Used to correlate UTC to system time if the timestamp of the message is in system time. This UTC time is derived from the GPS data.
	Velocity         []int16          // Scale: 100; Array: [3]; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewGpsMetadata creates new GpsMetadata struct based on given mesg. If mesg is nil or mesg.Num is not equal to GpsMetadata mesg number, it will return nil.
func NewGpsMetadata(mesg proto.Message) *GpsMetadata {
	if mesg.Num != typedef.MesgNumGpsMetadata {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* TimestampMs */
		1:   basetype.Sint32Invalid, /* PositionLat */
		2:   basetype.Sint32Invalid, /* PositionLong */
		3:   basetype.Uint32Invalid, /* EnhancedAltitude */
		4:   basetype.Uint32Invalid, /* EnhancedSpeed */
		5:   basetype.Uint16Invalid, /* Heading */
		6:   basetype.Uint32Invalid, /* UtcTimestamp */
		7:   nil,                    /* Velocity */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &GpsMetadata{
		Timestamp:        typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		PositionLat:      typeconv.ToSint32[int32](vals[1]),
		PositionLong:     typeconv.ToSint32[int32](vals[2]),
		EnhancedAltitude: typeconv.ToUint32[uint32](vals[3]),
		EnhancedSpeed:    typeconv.ToUint32[uint32](vals[4]),
		Heading:          typeconv.ToUint16[uint16](vals[5]),
		UtcTimestamp:     typeconv.ToUint32[typedef.DateTime](vals[6]),
		Velocity:         typeconv.ToSliceSint16[int16](vals[7]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to GpsMetadata mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumGpsMetadata)
func (m GpsMetadata) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumGpsMetadata {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.PositionLat,
		2:   m.PositionLong,
		3:   m.EnhancedAltitude,
		4:   m.EnhancedSpeed,
		5:   m.Heading,
		6:   m.UtcTimestamp,
		7:   m.Velocity,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
