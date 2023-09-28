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

// ZonesTarget is a ZonesTarget message.
type ZonesTarget struct {
	MaxHeartRate             uint8
	ThresholdHeartRate       uint8
	FunctionalThresholdPower uint16
	HrCalcType               typedef.HrZoneCalc
	PwrCalcType              typedef.PwrZoneCalc

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewZonesTarget creates new ZonesTarget struct based on given mesg. If mesg is nil or mesg.Num is not equal to ZonesTarget mesg number, it will return nil.
func NewZonesTarget(mesg proto.Message) *ZonesTarget {
	if mesg.Num != typedef.MesgNumZonesTarget {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		1: basetype.Uint8Invalid,  /* MaxHeartRate */
		2: basetype.Uint8Invalid,  /* ThresholdHeartRate */
		3: basetype.Uint16Invalid, /* FunctionalThresholdPower */
		5: basetype.EnumInvalid,   /* HrCalcType */
		7: basetype.EnumInvalid,   /* PwrCalcType */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &ZonesTarget{
		MaxHeartRate:             typeconv.ToUint8[uint8](vals[1]),
		ThresholdHeartRate:       typeconv.ToUint8[uint8](vals[2]),
		FunctionalThresholdPower: typeconv.ToUint16[uint16](vals[3]),
		HrCalcType:               typeconv.ToEnum[typedef.HrZoneCalc](vals[5]),
		PwrCalcType:              typeconv.ToEnum[typedef.PwrZoneCalc](vals[7]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to ZonesTarget mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumZonesTarget)
func (m ZonesTarget) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumZonesTarget {
		return
	}

	vals := [256]any{
		1: m.MaxHeartRate,
		2: m.ThresholdHeartRate,
		3: m.FunctionalThresholdPower,
		5: m.HrCalcType,
		7: m.PwrCalcType,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
