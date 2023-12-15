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

// DeviceAuxBatteryInfo is a DeviceAuxBatteryInfo message.
type DeviceAuxBatteryInfo struct {
	Timestamp         typedef.DateTime
	DeviceIndex       typedef.DeviceIndex
	BatteryVoltage    uint16 // Scale: 256; Units: V;
	BatteryStatus     typedef.BatteryStatus
	BatteryIdentifier uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDeviceAuxBatteryInfo creates new DeviceAuxBatteryInfo struct based on given mesg. If mesg is nil or mesg.Num is not equal to DeviceAuxBatteryInfo mesg number, it will return nil.
func NewDeviceAuxBatteryInfo(mesg proto.Message) *DeviceAuxBatteryInfo {
	if mesg.Num != typedef.MesgNumDeviceAuxBatteryInfo {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* DeviceIndex */
		1:   nil, /* BatteryVoltage */
		2:   nil, /* BatteryStatus */
		3:   nil, /* BatteryIdentifier */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &DeviceAuxBatteryInfo{
		Timestamp:         typeconv.ToUint32[typedef.DateTime](vals[253]),
		DeviceIndex:       typeconv.ToUint8[typedef.DeviceIndex](vals[0]),
		BatteryVoltage:    typeconv.ToUint16[uint16](vals[1]),
		BatteryStatus:     typeconv.ToUint8[typedef.BatteryStatus](vals[2]),
		BatteryIdentifier: typeconv.ToUint8[uint8](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DeviceAuxBatteryInfo mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDeviceAuxBatteryInfo)
func (m *DeviceAuxBatteryInfo) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDeviceAuxBatteryInfo {
		return
	}

	vals := [...]any{
		253: typeconv.ToUint32[uint32](m.Timestamp),
		0:   typeconv.ToUint8[uint8](m.DeviceIndex),
		1:   m.BatteryVoltage,
		2:   typeconv.ToUint8[uint8](m.BatteryStatus),
		3:   m.BatteryIdentifier,
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
