// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

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

// DeviceSettings is a DeviceSettings message.
type DeviceSettings struct {
	ActiveTimeZone                      uint8                 // Index into time zone arrays.
	UtcOffset                           uint32                // Offset from system time. Required to convert timestamp from system time to UTC.
	TimeOffset                          []uint32              // Array: [N]; Units: s; Offset from system time.
	TimeMode                            []typedef.TimeMode    // Array: [N]; Display mode for the time
	TimeZoneOffset                      []int8                // Scale: 4; Array: [N]; Units: hr; timezone offset in 1/4 hour increments
	BacklightMode                       typedef.BacklightMode // Mode for backlight
	ActivityTrackerEnabled              bool                  // Enabled state of the activity tracker functionality
	ClockTime                           typedef.DateTime      // UTC timestamp used to set the devices clock and date
	PagesEnabled                        []uint16              // Array: [N]; Bitfield to configure enabled screens for each supported loop
	MoveAlertEnabled                    bool                  // Enabled state of the move alert
	DateMode                            typedef.DateMode      // Display mode for the date
	DisplayOrientation                  typedef.DisplayOrientation
	MountingSide                        typedef.Side
	DefaultPage                         []uint16                   // Array: [N]; Bitfield to indicate one page as default for each supported loop
	AutosyncMinSteps                    uint16                     // Units: steps; Minimum steps before an autosync can occur
	AutosyncMinTime                     uint16                     // Units: minutes; Minimum minutes before an autosync can occur
	LactateThresholdAutodetectEnabled   bool                       // Enable auto-detect setting for the lactate threshold feature.
	BleAutoUploadEnabled                bool                       // Automatically upload using BLE
	AutoSyncFrequency                   typedef.AutoSyncFrequency  // Helps to conserve battery by changing modes
	AutoActivityDetect                  typedef.AutoActivityDetect // Allows setting specific activities auto-activity detect enabled/disabled settings
	NumberOfScreens                     uint8                      // Number of screens configured to display
	SmartNotificationDisplayOrientation typedef.DisplayOrientation // Smart Notification display orientation
	TapInterface                        typedef.Switch
	TapSensitivity                      typedef.TapSensitivity // Used to hold the tap threshold setting

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDeviceSettings creates new DeviceSettings struct based on given mesg. If mesg is nil or mesg.Num is not equal to DeviceSettings mesg number, it will return nil.
func NewDeviceSettings(mesg proto.Message) *DeviceSettings {
	if mesg.Num != typedef.MesgNumDeviceSettings {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0:   basetype.Uint8Invalid,  /* ActiveTimeZone */
		1:   basetype.Uint32Invalid, /* UtcOffset */
		2:   nil,                    /* TimeOffset */
		4:   nil,                    /* TimeMode */
		5:   nil,                    /* TimeZoneOffset */
		12:  basetype.EnumInvalid,   /* BacklightMode */
		36:  false,                  /* ActivityTrackerEnabled */
		39:  basetype.Uint32Invalid, /* ClockTime */
		40:  nil,                    /* PagesEnabled */
		46:  false,                  /* MoveAlertEnabled */
		47:  basetype.EnumInvalid,   /* DateMode */
		55:  basetype.EnumInvalid,   /* DisplayOrientation */
		56:  basetype.EnumInvalid,   /* MountingSide */
		57:  nil,                    /* DefaultPage */
		58:  basetype.Uint16Invalid, /* AutosyncMinSteps */
		59:  basetype.Uint16Invalid, /* AutosyncMinTime */
		80:  false,                  /* LactateThresholdAutodetectEnabled */
		86:  false,                  /* BleAutoUploadEnabled */
		89:  basetype.EnumInvalid,   /* AutoSyncFrequency */
		90:  basetype.Uint32Invalid, /* AutoActivityDetect */
		94:  basetype.Uint8Invalid,  /* NumberOfScreens */
		95:  basetype.EnumInvalid,   /* SmartNotificationDisplayOrientation */
		134: basetype.EnumInvalid,   /* TapInterface */
		174: basetype.EnumInvalid,   /* TapSensitivity */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &DeviceSettings{
		ActiveTimeZone:                      typeconv.ToUint8[uint8](vals[0]),
		UtcOffset:                           typeconv.ToUint32[uint32](vals[1]),
		TimeOffset:                          typeconv.ToSliceUint32[uint32](vals[2]),
		TimeMode:                            typeconv.ToSliceEnum[typedef.TimeMode](vals[4]),
		TimeZoneOffset:                      typeconv.ToSliceSint8[int8](vals[5]),
		BacklightMode:                       typeconv.ToEnum[typedef.BacklightMode](vals[12]),
		ActivityTrackerEnabled:              typeconv.ToBool[bool](vals[36]),
		ClockTime:                           typeconv.ToUint32[typedef.DateTime](vals[39]),
		PagesEnabled:                        typeconv.ToSliceUint16[uint16](vals[40]),
		MoveAlertEnabled:                    typeconv.ToBool[bool](vals[46]),
		DateMode:                            typeconv.ToEnum[typedef.DateMode](vals[47]),
		DisplayOrientation:                  typeconv.ToEnum[typedef.DisplayOrientation](vals[55]),
		MountingSide:                        typeconv.ToEnum[typedef.Side](vals[56]),
		DefaultPage:                         typeconv.ToSliceUint16[uint16](vals[57]),
		AutosyncMinSteps:                    typeconv.ToUint16[uint16](vals[58]),
		AutosyncMinTime:                     typeconv.ToUint16[uint16](vals[59]),
		LactateThresholdAutodetectEnabled:   typeconv.ToBool[bool](vals[80]),
		BleAutoUploadEnabled:                typeconv.ToBool[bool](vals[86]),
		AutoSyncFrequency:                   typeconv.ToEnum[typedef.AutoSyncFrequency](vals[89]),
		AutoActivityDetect:                  typeconv.ToUint32[typedef.AutoActivityDetect](vals[90]),
		NumberOfScreens:                     typeconv.ToUint8[uint8](vals[94]),
		SmartNotificationDisplayOrientation: typeconv.ToEnum[typedef.DisplayOrientation](vals[95]),
		TapInterface:                        typeconv.ToEnum[typedef.Switch](vals[134]),
		TapSensitivity:                      typeconv.ToEnum[typedef.TapSensitivity](vals[174]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DeviceSettings mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDeviceSettings)
func (m DeviceSettings) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDeviceSettings {
		return
	}

	vals := [256]any{
		0:   m.ActiveTimeZone,
		1:   m.UtcOffset,
		2:   m.TimeOffset,
		4:   m.TimeMode,
		5:   m.TimeZoneOffset,
		12:  m.BacklightMode,
		36:  m.ActivityTrackerEnabled,
		39:  m.ClockTime,
		40:  m.PagesEnabled,
		46:  m.MoveAlertEnabled,
		47:  m.DateMode,
		55:  m.DisplayOrientation,
		56:  m.MountingSide,
		57:  m.DefaultPage,
		58:  m.AutosyncMinSteps,
		59:  m.AutosyncMinTime,
		80:  m.LactateThresholdAutodetectEnabled,
		86:  m.BleAutoUploadEnabled,
		89:  m.AutoSyncFrequency,
		90:  m.AutoActivityDetect,
		94:  m.NumberOfScreens,
		95:  m.SmartNotificationDisplayOrientation,
		134: m.TapInterface,
		174: m.TapSensitivity,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
