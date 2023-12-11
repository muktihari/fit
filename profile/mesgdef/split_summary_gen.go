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

// SplitSummary is a SplitSummary message.
type SplitSummary struct {
	MessageIndex    typedef.MessageIndex
	SplitType       typedef.SplitType
	NumSplits       uint16
	TotalTimerTime  uint32 // Scale: 1000; Units: s;
	TotalDistance   uint32 // Scale: 100; Units: m;
	AvgSpeed        uint32 // Scale: 1000; Units: m/s;
	MaxSpeed        uint32 // Scale: 1000; Units: m/s;
	TotalAscent     uint16 // Units: m;
	TotalDescent    uint16 // Units: m;
	AvgHeartRate    uint8  // Units: bpm;
	MaxHeartRate    uint8  // Units: bpm;
	AvgVertSpeed    int32  // Scale: 1000; Units: m/s;
	TotalCalories   uint32 // Units: kcal;
	TotalMovingTime uint32 // Scale: 1000; Units: s;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSplitSummary creates new SplitSummary struct based on given mesg. If mesg is nil or mesg.Num is not equal to SplitSummary mesg number, it will return nil.
func NewSplitSummary(mesg proto.Message) *SplitSummary {
	if mesg.Num != typedef.MesgNumSplitSummary {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		254: nil, /* MessageIndex */
		0:   nil, /* SplitType */
		3:   nil, /* NumSplits */
		4:   nil, /* TotalTimerTime */
		5:   nil, /* TotalDistance */
		6:   nil, /* AvgSpeed */
		7:   nil, /* MaxSpeed */
		8:   nil, /* TotalAscent */
		9:   nil, /* TotalDescent */
		10:  nil, /* AvgHeartRate */
		11:  nil, /* MaxHeartRate */
		12:  nil, /* AvgVertSpeed */
		13:  nil, /* TotalCalories */
		77:  nil, /* TotalMovingTime */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &SplitSummary{
		MessageIndex:    typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		SplitType:       typeconv.ToEnum[typedef.SplitType](vals[0]),
		NumSplits:       typeconv.ToUint16[uint16](vals[3]),
		TotalTimerTime:  typeconv.ToUint32[uint32](vals[4]),
		TotalDistance:   typeconv.ToUint32[uint32](vals[5]),
		AvgSpeed:        typeconv.ToUint32[uint32](vals[6]),
		MaxSpeed:        typeconv.ToUint32[uint32](vals[7]),
		TotalAscent:     typeconv.ToUint16[uint16](vals[8]),
		TotalDescent:    typeconv.ToUint16[uint16](vals[9]),
		AvgHeartRate:    typeconv.ToUint8[uint8](vals[10]),
		MaxHeartRate:    typeconv.ToUint8[uint8](vals[11]),
		AvgVertSpeed:    typeconv.ToSint32[int32](vals[12]),
		TotalCalories:   typeconv.ToUint32[uint32](vals[13]),
		TotalMovingTime: typeconv.ToUint32[uint32](vals[77]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SplitSummary mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSplitSummary)
func (m SplitSummary) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSplitSummary {
		return
	}

	vals := [...]any{
		254: m.MessageIndex,
		0:   m.SplitType,
		3:   m.NumSplits,
		4:   m.TotalTimerTime,
		5:   m.TotalDistance,
		6:   m.AvgSpeed,
		7:   m.MaxSpeed,
		8:   m.TotalAscent,
		9:   m.TotalDescent,
		10:  m.AvgHeartRate,
		11:  m.MaxHeartRate,
		12:  m.AvgVertSpeed,
		13:  m.TotalCalories,
		77:  m.TotalMovingTime,
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
