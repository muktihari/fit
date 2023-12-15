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

// MaxMetData is a MaxMetData message.
type MaxMetData struct {
	UpdateTime     typedef.DateTime // Time maxMET and vo2 were calculated
	Vo2Max         uint16           // Scale: 10; Units: mL/kg/min;
	Sport          typedef.Sport
	SubSport       typedef.SubSport
	MaxMetCategory typedef.MaxMetCategory
	CalibratedData bool                          // Indicates if calibrated data was used in the calculation
	HrSource       typedef.MaxMetHeartRateSource // Indicates if the estimate was obtained using a chest strap or wrist heart rate
	SpeedSource    typedef.MaxMetSpeedSource     // Indidcates if the estimate was obtained using onboard GPS or connected GPS

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMaxMetData creates new MaxMetData struct based on given mesg. If mesg is nil or mesg.Num is not equal to MaxMetData mesg number, it will return nil.
func NewMaxMetData(mesg proto.Message) *MaxMetData {
	if mesg.Num != typedef.MesgNumMaxMetData {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		0:  nil, /* UpdateTime */
		2:  nil, /* Vo2Max */
		5:  nil, /* Sport */
		6:  nil, /* SubSport */
		8:  nil, /* MaxMetCategory */
		9:  nil, /* CalibratedData */
		12: nil, /* HrSource */
		13: nil, /* SpeedSource */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &MaxMetData{
		UpdateTime:     typeconv.ToUint32[typedef.DateTime](vals[0]),
		Vo2Max:         typeconv.ToUint16[uint16](vals[2]),
		Sport:          typeconv.ToEnum[typedef.Sport](vals[5]),
		SubSport:       typeconv.ToEnum[typedef.SubSport](vals[6]),
		MaxMetCategory: typeconv.ToEnum[typedef.MaxMetCategory](vals[8]),
		CalibratedData: typeconv.ToBool[bool](vals[9]),
		HrSource:       typeconv.ToEnum[typedef.MaxMetHeartRateSource](vals[12]),
		SpeedSource:    typeconv.ToEnum[typedef.MaxMetSpeedSource](vals[13]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to MaxMetData mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumMaxMetData)
func (m *MaxMetData) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumMaxMetData {
		return
	}

	vals := [...]any{
		0:  typeconv.ToUint32[uint32](m.UpdateTime),
		2:  m.Vo2Max,
		5:  typeconv.ToEnum[byte](m.Sport),
		6:  typeconv.ToEnum[byte](m.SubSport),
		8:  typeconv.ToEnum[byte](m.MaxMetCategory),
		9:  m.CalibratedData,
		12: typeconv.ToEnum[byte](m.HrSource),
		13: typeconv.ToEnum[byte](m.SpeedSource),
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
