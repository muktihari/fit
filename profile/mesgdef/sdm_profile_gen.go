// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

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

// SdmProfile is a SdmProfile message.
type SdmProfile struct {
	MessageIndex      typedef.MessageIndex
	Enabled           bool
	SdmAntId          uint16
	SdmCalFactor      uint16 // Scale: 10; Units: %;
	Odometer          uint32 // Scale: 100; Units: m;
	SpeedSource       bool   // Use footpod for speed source instead of GPS
	SdmAntIdTransType uint8
	OdometerRollover  uint8 // Rollover counter that can be used to extend the odometer

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSdmProfile creates new SdmProfile struct based on given mesg. If mesg is nil or mesg.Num is not equal to SdmProfile mesg number, it will return nil.
func NewSdmProfile(mesg proto.Message) *SdmProfile {
	if mesg.Num != typedef.MesgNumSdmProfile {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid,  /* MessageIndex */
		0:   false,                   /* Enabled */
		1:   basetype.Uint16zInvalid, /* SdmAntId */
		2:   basetype.Uint16Invalid,  /* SdmCalFactor */
		3:   basetype.Uint32Invalid,  /* Odometer */
		4:   false,                   /* SpeedSource */
		5:   basetype.Uint8zInvalid,  /* SdmAntIdTransType */
		7:   basetype.Uint8Invalid,   /* OdometerRollover */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &SdmProfile{
		MessageIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Enabled:           typeconv.ToBool[bool](vals[0]),
		SdmAntId:          typeconv.ToUint16z[uint16](vals[1]),
		SdmCalFactor:      typeconv.ToUint16[uint16](vals[2]),
		Odometer:          typeconv.ToUint32[uint32](vals[3]),
		SpeedSource:       typeconv.ToBool[bool](vals[4]),
		SdmAntIdTransType: typeconv.ToUint8z[uint8](vals[5]),
		OdometerRollover:  typeconv.ToUint8[uint8](vals[7]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SdmProfile mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSdmProfile)
func (m SdmProfile) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSdmProfile {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.Enabled,
		1:   m.SdmAntId,
		2:   m.SdmCalFactor,
		3:   m.Odometer,
		4:   m.SpeedSource,
		5:   m.SdmAntIdTransType,
		7:   m.OdometerRollover,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
