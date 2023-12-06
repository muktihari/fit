// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

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

// MonitoringInfo is a MonitoringInfo message.
type MonitoringInfo struct {
	Timestamp            typedef.DateTime       // Units: s;
	LocalTimestamp       typedef.LocalDateTime  // Units: s; Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
	ActivityType         []typedef.ActivityType // Array: [N];
	CyclesToDistance     []uint16               // Scale: 5000; Array: [N]; Units: m/cycle; Indexed by activity_type
	CyclesToCalories     []uint16               // Scale: 5000; Array: [N]; Units: kcal/cycle; Indexed by activity_type
	RestingMetabolicRate uint16                 // Units: kcal / day;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMonitoringInfo creates new MonitoringInfo struct based on given mesg. If mesg is nil or mesg.Num is not equal to MonitoringInfo mesg number, it will return nil.
func NewMonitoringInfo(mesg proto.Message) *MonitoringInfo {
	if mesg.Num != typedef.MesgNumMonitoringInfo {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint32Invalid, /* LocalTimestamp */
		1:   nil,                    /* ActivityType */
		3:   nil,                    /* CyclesToDistance */
		4:   nil,                    /* CyclesToCalories */
		5:   basetype.Uint16Invalid, /* RestingMetabolicRate */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &MonitoringInfo{
		Timestamp:            typeconv.ToUint32[typedef.DateTime](vals[253]),
		LocalTimestamp:       typeconv.ToUint32[typedef.LocalDateTime](vals[0]),
		ActivityType:         typeconv.ToSliceEnum[typedef.ActivityType](vals[1]),
		CyclesToDistance:     typeconv.ToSliceUint16[uint16](vals[3]),
		CyclesToCalories:     typeconv.ToSliceUint16[uint16](vals[4]),
		RestingMetabolicRate: typeconv.ToUint16[uint16](vals[5]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to MonitoringInfo mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumMonitoringInfo)
func (m MonitoringInfo) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumMonitoringInfo {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.LocalTimestamp,
		1:   m.ActivityType,
		3:   m.CyclesToDistance,
		4:   m.CyclesToCalories,
		5:   m.RestingMetabolicRate,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
