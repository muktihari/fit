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

// MonitoringHrData is a MonitoringHrData message.
type MonitoringHrData struct {
	Timestamp                  typedef.DateTime // Units: s; Must align to logging interval, for example, time must be 00:00:00 for daily log.
	RestingHeartRate           uint8            // Units: bpm; 7-day rolling average
	CurrentDayRestingHeartRate uint8            // Units: bpm; RHR for today only. (Feeds into 7-day average)

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMonitoringHrData creates new MonitoringHrData struct based on given mesg. If mesg is nil or mesg.Num is not equal to MonitoringHrData mesg number, it will return nil.
func NewMonitoringHrData(mesg proto.Message) *MonitoringHrData {
	if mesg.Num != typedef.MesgNumMonitoringHrData {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint8Invalid,  /* RestingHeartRate */
		1:   basetype.Uint8Invalid,  /* CurrentDayRestingHeartRate */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &MonitoringHrData{
		Timestamp:                  typeconv.ToUint32[typedef.DateTime](vals[253]),
		RestingHeartRate:           typeconv.ToUint8[uint8](vals[0]),
		CurrentDayRestingHeartRate: typeconv.ToUint8[uint8](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to MonitoringHrData mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumMonitoringHrData)
func (m MonitoringHrData) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumMonitoringHrData {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.RestingHeartRate,
		1:   m.CurrentDayRestingHeartRate,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
