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

// TimeInZone is a TimeInZone message.
type TimeInZone struct {
	Timestamp                typedef.DateTime // Units: s;
	ReferenceMesg            typedef.MesgNum
	ReferenceIndex           typedef.MessageIndex
	TimeInHrZone             []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInSpeedZone          []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInCadenceZone        []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInPowerZone          []uint32 // Scale: 1000; Array: [N]; Units: s;
	HrZoneHighBoundary       []uint8  // Array: [N]; Units: bpm;
	SpeedZoneHighBoundary    []uint16 // Scale: 1000; Array: [N]; Units: m/s;
	CadenceZoneHighBondary   []uint8  // Array: [N]; Units: rpm;
	PowerZoneHighBoundary    []uint16 // Array: [N]; Units: watts;
	HrCalcType               typedef.HrZoneCalc
	MaxHeartRate             uint8
	RestingHeartRate         uint8
	ThresholdHeartRate       uint8
	PwrCalcType              typedef.PwrZoneCalc
	FunctionalThresholdPower uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTimeInZone creates new TimeInZone struct based on given mesg. If mesg is nil or mesg.Num is not equal to TimeInZone mesg number, it will return nil.
func NewTimeInZone(mesg proto.Message) *TimeInZone {
	if mesg.Num != typedef.MesgNumTimeInZone {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* ReferenceMesg */
		1:   basetype.Uint16Invalid, /* ReferenceIndex */
		2:   nil,                    /* TimeInHrZone */
		3:   nil,                    /* TimeInSpeedZone */
		4:   nil,                    /* TimeInCadenceZone */
		5:   nil,                    /* TimeInPowerZone */
		6:   nil,                    /* HrZoneHighBoundary */
		7:   nil,                    /* SpeedZoneHighBoundary */
		8:   nil,                    /* CadenceZoneHighBondary */
		9:   nil,                    /* PowerZoneHighBoundary */
		10:  basetype.EnumInvalid,   /* HrCalcType */
		11:  basetype.Uint8Invalid,  /* MaxHeartRate */
		12:  basetype.Uint8Invalid,  /* RestingHeartRate */
		13:  basetype.Uint8Invalid,  /* ThresholdHeartRate */
		14:  basetype.EnumInvalid,   /* PwrCalcType */
		15:  basetype.Uint16Invalid, /* FunctionalThresholdPower */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &TimeInZone{
		Timestamp:                typeconv.ToUint32[typedef.DateTime](vals[253]),
		ReferenceMesg:            typeconv.ToUint16[typedef.MesgNum](vals[0]),
		ReferenceIndex:           typeconv.ToUint16[typedef.MessageIndex](vals[1]),
		TimeInHrZone:             typeconv.ToSliceUint32[uint32](vals[2]),
		TimeInSpeedZone:          typeconv.ToSliceUint32[uint32](vals[3]),
		TimeInCadenceZone:        typeconv.ToSliceUint32[uint32](vals[4]),
		TimeInPowerZone:          typeconv.ToSliceUint32[uint32](vals[5]),
		HrZoneHighBoundary:       typeconv.ToSliceUint8[uint8](vals[6]),
		SpeedZoneHighBoundary:    typeconv.ToSliceUint16[uint16](vals[7]),
		CadenceZoneHighBondary:   typeconv.ToSliceUint8[uint8](vals[8]),
		PowerZoneHighBoundary:    typeconv.ToSliceUint16[uint16](vals[9]),
		HrCalcType:               typeconv.ToEnum[typedef.HrZoneCalc](vals[10]),
		MaxHeartRate:             typeconv.ToUint8[uint8](vals[11]),
		RestingHeartRate:         typeconv.ToUint8[uint8](vals[12]),
		ThresholdHeartRate:       typeconv.ToUint8[uint8](vals[13]),
		PwrCalcType:              typeconv.ToEnum[typedef.PwrZoneCalc](vals[14]),
		FunctionalThresholdPower: typeconv.ToUint16[uint16](vals[15]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to TimeInZone mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumTimeInZone)
func (m TimeInZone) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumTimeInZone {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.ReferenceMesg,
		1:   m.ReferenceIndex,
		2:   m.TimeInHrZone,
		3:   m.TimeInSpeedZone,
		4:   m.TimeInCadenceZone,
		5:   m.TimeInPowerZone,
		6:   m.HrZoneHighBoundary,
		7:   m.SpeedZoneHighBoundary,
		8:   m.CadenceZoneHighBondary,
		9:   m.PowerZoneHighBoundary,
		10:  m.HrCalcType,
		11:  m.MaxHeartRate,
		12:  m.RestingHeartRate,
		13:  m.ThresholdHeartRate,
		14:  m.PwrCalcType,
		15:  m.FunctionalThresholdPower,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
