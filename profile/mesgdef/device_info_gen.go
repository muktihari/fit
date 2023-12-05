// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

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

// DeviceInfo is a DeviceInfo message.
type DeviceInfo struct {
	Timestamp           typedef.DateTime // Units: s;
	DeviceIndex         typedef.DeviceIndex
	DeviceType          uint8
	Manufacturer        typedef.Manufacturer
	SerialNumber        uint32
	Product             uint16
	SoftwareVersion     uint16 // Scale: 100;
	HardwareVersion     uint8
	CumOperatingTime    uint32 // Units: s; Reset by new battery or charge.
	BatteryVoltage      uint16 // Scale: 256; Units: V;
	BatteryStatus       typedef.BatteryStatus
	SensorPosition      typedef.BodyLocation // Indicates the location of the sensor
	Descriptor          string               // Used to describe the sensor or location
	AntTransmissionType uint8
	AntDeviceNumber     uint16
	AntNetwork          typedef.AntNetwork
	SourceType          typedef.SourceType
	ProductName         string // Optional free form string to indicate the devices name or model
	BatteryLevel        uint8  // Units: %;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDeviceInfo creates new DeviceInfo struct based on given mesg. If mesg is nil or mesg.Num is not equal to DeviceInfo mesg number, it will return nil.
func NewDeviceInfo(mesg proto.Message) *DeviceInfo {
	if mesg.Num != typedef.MesgNumDeviceInfo {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid,  /* Timestamp */
		0:   basetype.Uint8Invalid,   /* DeviceIndex */
		1:   basetype.Uint8Invalid,   /* DeviceType */
		2:   basetype.Uint16Invalid,  /* Manufacturer */
		3:   basetype.Uint32zInvalid, /* SerialNumber */
		4:   basetype.Uint16Invalid,  /* Product */
		5:   basetype.Uint16Invalid,  /* SoftwareVersion */
		6:   basetype.Uint8Invalid,   /* HardwareVersion */
		7:   basetype.Uint32Invalid,  /* CumOperatingTime */
		10:  basetype.Uint16Invalid,  /* BatteryVoltage */
		11:  basetype.Uint8Invalid,   /* BatteryStatus */
		18:  basetype.EnumInvalid,    /* SensorPosition */
		19:  basetype.StringInvalid,  /* Descriptor */
		20:  basetype.Uint8zInvalid,  /* AntTransmissionType */
		21:  basetype.Uint16zInvalid, /* AntDeviceNumber */
		22:  basetype.EnumInvalid,    /* AntNetwork */
		25:  basetype.EnumInvalid,    /* SourceType */
		27:  basetype.StringInvalid,  /* ProductName */
		32:  basetype.Uint8Invalid,   /* BatteryLevel */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &DeviceInfo{
		Timestamp:           typeconv.ToUint32[typedef.DateTime](vals[253]),
		DeviceIndex:         typeconv.ToUint8[typedef.DeviceIndex](vals[0]),
		DeviceType:          typeconv.ToUint8[uint8](vals[1]),
		Manufacturer:        typeconv.ToUint16[typedef.Manufacturer](vals[2]),
		SerialNumber:        typeconv.ToUint32z[uint32](vals[3]),
		Product:             typeconv.ToUint16[uint16](vals[4]),
		SoftwareVersion:     typeconv.ToUint16[uint16](vals[5]),
		HardwareVersion:     typeconv.ToUint8[uint8](vals[6]),
		CumOperatingTime:    typeconv.ToUint32[uint32](vals[7]),
		BatteryVoltage:      typeconv.ToUint16[uint16](vals[10]),
		BatteryStatus:       typeconv.ToUint8[typedef.BatteryStatus](vals[11]),
		SensorPosition:      typeconv.ToEnum[typedef.BodyLocation](vals[18]),
		Descriptor:          typeconv.ToString[string](vals[19]),
		AntTransmissionType: typeconv.ToUint8z[uint8](vals[20]),
		AntDeviceNumber:     typeconv.ToUint16z[uint16](vals[21]),
		AntNetwork:          typeconv.ToEnum[typedef.AntNetwork](vals[22]),
		SourceType:          typeconv.ToEnum[typedef.SourceType](vals[25]),
		ProductName:         typeconv.ToString[string](vals[27]),
		BatteryLevel:        typeconv.ToUint8[uint8](vals[32]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DeviceInfo mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDeviceInfo)
func (m DeviceInfo) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDeviceInfo {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.DeviceIndex,
		1:   m.DeviceType,
		2:   m.Manufacturer,
		3:   m.SerialNumber,
		4:   m.Product,
		5:   m.SoftwareVersion,
		6:   m.HardwareVersion,
		7:   m.CumOperatingTime,
		10:  m.BatteryVoltage,
		11:  m.BatteryStatus,
		18:  m.SensorPosition,
		19:  m.Descriptor,
		20:  m.AntTransmissionType,
		21:  m.AntDeviceNumber,
		22:  m.AntNetwork,
		25:  m.SourceType,
		27:  m.ProductName,
		32:  m.BatteryLevel,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
