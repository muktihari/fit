// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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

	vals := [254]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &MonitoringHrData{
		Timestamp:                  typeconv.ToUint32[typedef.DateTime](vals[253]),
		RestingHeartRate:           typeconv.ToUint8[uint8](vals[0]),
		CurrentDayRestingHeartRate: typeconv.ToUint8[uint8](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts MonitoringHrData into proto.Message.
func (m *MonitoringHrData) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumMonitoringHrData)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = typeconv.ToUint32[uint32](m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.RestingHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.CurrentDayRestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.CurrentDayRestingHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of MonitoringHrData's valid fields.
func (m *MonitoringHrData) size() byte {
	var size byte
	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.CurrentDayRestingHeartRate != basetype.Uint8Invalid {
		size++
	}
	return size
}
