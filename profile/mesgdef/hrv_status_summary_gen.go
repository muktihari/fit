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

// HrvStatusSummary is a HrvStatusSummary message.
type HrvStatusSummary struct {
	Timestamp             typedef.DateTime
	WeeklyAverage         uint16 // Scale: 128; Units: ms; 7 day RMSSD average over sleep
	LastNightAverage      uint16 // Scale: 128; Units: ms; Last night RMSSD average over sleep
	LastNight5MinHigh     uint16 // Scale: 128; Units: ms; 5 minute high RMSSD value over sleep
	BaselineLowUpper      uint16 // Scale: 128; Units: ms; 3 week baseline, upper boundary of low HRV status
	BaselineBalancedLower uint16 // Scale: 128; Units: ms; 3 week baseline, lower boundary of balanced HRV status
	BaselineBalancedUpper uint16 // Scale: 128; Units: ms; 3 week baseline, upper boundary of balanced HRV status
	Status                typedef.HrvStatus

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrvStatusSummary creates new HrvStatusSummary struct based on given mesg. If mesg is nil or mesg.Num is not equal to HrvStatusSummary mesg number, it will return nil.
func NewHrvStatusSummary(mesg proto.Message) *HrvStatusSummary {
	if mesg.Num != typedef.MesgNumHrvStatusSummary {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* WeeklyAverage */
		1:   basetype.Uint16Invalid, /* LastNightAverage */
		2:   basetype.Uint16Invalid, /* LastNight5MinHigh */
		3:   basetype.Uint16Invalid, /* BaselineLowUpper */
		4:   basetype.Uint16Invalid, /* BaselineBalancedLower */
		5:   basetype.Uint16Invalid, /* BaselineBalancedUpper */
		6:   basetype.EnumInvalid,   /* Status */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &HrvStatusSummary{
		Timestamp:             typeconv.ToUint32[typedef.DateTime](vals[253]),
		WeeklyAverage:         typeconv.ToUint16[uint16](vals[0]),
		LastNightAverage:      typeconv.ToUint16[uint16](vals[1]),
		LastNight5MinHigh:     typeconv.ToUint16[uint16](vals[2]),
		BaselineLowUpper:      typeconv.ToUint16[uint16](vals[3]),
		BaselineBalancedLower: typeconv.ToUint16[uint16](vals[4]),
		BaselineBalancedUpper: typeconv.ToUint16[uint16](vals[5]),
		Status:                typeconv.ToEnum[typedef.HrvStatus](vals[6]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to HrvStatusSummary mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumHrvStatusSummary)
func (m HrvStatusSummary) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumHrvStatusSummary {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.WeeklyAverage,
		1:   m.LastNightAverage,
		2:   m.LastNight5MinHigh,
		3:   m.BaselineLowUpper,
		4:   m.BaselineBalancedLower,
		5:   m.BaselineBalancedUpper,
		6:   m.Status,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
