// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// DeviceAuxBatteryInfo is a DeviceAuxBatteryInfo message.
type DeviceAuxBatteryInfo struct {
	Timestamp         time.Time
	BatteryVoltage    uint16 // Scale: 256; Units: V
	DeviceIndex       typedef.DeviceIndex
	BatteryStatus     typedef.BatteryStatus
	BatteryIdentifier uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDeviceAuxBatteryInfo creates new DeviceAuxBatteryInfo struct based on given mesg.
// If mesg is nil, it will return DeviceAuxBatteryInfo with all fields being set to its corresponding invalid value.
func NewDeviceAuxBatteryInfo(mesg *proto.Message) *DeviceAuxBatteryInfo {
	vals := [254]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &DeviceAuxBatteryInfo{
		Timestamp:         datetime.ToTime(vals[253].Uint32()),
		BatteryVoltage:    vals[1].Uint16(),
		DeviceIndex:       typedef.DeviceIndex(vals[0].Uint8()),
		BatteryStatus:     typedef.BatteryStatus(vals[2].Uint8()),
		BatteryIdentifier: vals[3].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts DeviceAuxBatteryInfo into proto.Message. If options is nil, default options will be used.
func (m *DeviceAuxBatteryInfo) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumDeviceAuxBatteryInfo}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.BatteryVoltage != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.BatteryVoltage)
		fields = append(fields, field)
	}
	if uint8(m.DeviceIndex) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
	}
	if uint8(m.BatteryStatus) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(uint8(m.BatteryStatus))
		fields = append(fields, field)
	}
	if m.BatteryIdentifier != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.BatteryIdentifier)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// BatteryVoltageScaled return BatteryVoltage in its scaled value [Scale: 256; Units: V].
//
// If BatteryVoltage value is invalid, float64 invalid value will be returned.
func (m *DeviceAuxBatteryInfo) BatteryVoltageScaled() float64 {
	if m.BatteryVoltage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.BatteryVoltage, 256, 0)
}

// SetTimestamp sets DeviceAuxBatteryInfo value.
func (m *DeviceAuxBatteryInfo) SetTimestamp(v time.Time) *DeviceAuxBatteryInfo {
	m.Timestamp = v
	return m
}

// SetBatteryVoltage sets DeviceAuxBatteryInfo value.
//
// Scale: 256; Units: V
func (m *DeviceAuxBatteryInfo) SetBatteryVoltage(v uint16) *DeviceAuxBatteryInfo {
	m.BatteryVoltage = v
	return m
}

// SetDeviceIndex sets DeviceAuxBatteryInfo value.
func (m *DeviceAuxBatteryInfo) SetDeviceIndex(v typedef.DeviceIndex) *DeviceAuxBatteryInfo {
	m.DeviceIndex = v
	return m
}

// SetBatteryStatus sets DeviceAuxBatteryInfo value.
func (m *DeviceAuxBatteryInfo) SetBatteryStatus(v typedef.BatteryStatus) *DeviceAuxBatteryInfo {
	m.BatteryStatus = v
	return m
}

// SetBatteryIdentifier sets DeviceAuxBatteryInfo value.
func (m *DeviceAuxBatteryInfo) SetBatteryIdentifier(v uint8) *DeviceAuxBatteryInfo {
	m.BatteryIdentifier = v
	return m
}

// SetDeveloperFields DeviceAuxBatteryInfo's DeveloperFields.
func (m *DeviceAuxBatteryInfo) SetDeveloperFields(developerFields ...proto.DeveloperField) *DeviceAuxBatteryInfo {
	m.DeveloperFields = developerFields
	return m
}
