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
	"time"
	"unsafe"
)

// DeviceSettings is a DeviceSettings message.
type DeviceSettings struct {
	TimeOffset                          []uint32                   // Array: [N]; Units: s; Offset from system time.
	TimeMode                            []typedef.TimeMode         // Array: [N]; Display mode for the time
	TimeZoneOffset                      []int8                     // Array: [N]; Scale: 4; Units: hr; timezone offset in 1/4 hour increments
	ClockTime                           time.Time                  // UTC timestamp used to set the devices clock and date
	PagesEnabled                        []uint16                   // Array: [N]; Bitfield to configure enabled screens for each supported loop
	DefaultPage                         []uint16                   // Array: [N]; Bitfield to indicate one page as default for each supported loop
	UtcOffset                           uint32                     // Offset from system time. Required to convert timestamp from system time to UTC.
	AutoActivityDetect                  typedef.AutoActivityDetect // Allows setting specific activities auto-activity detect enabled/disabled settings
	AutosyncMinSteps                    uint16                     // Units: steps; Minimum steps before an autosync can occur
	AutosyncMinTime                     uint16                     // Units: minutes; Minimum minutes before an autosync can occur
	ActiveTimeZone                      uint8                      // Index into time zone arrays.
	BacklightMode                       typedef.BacklightMode      // Mode for backlight
	DateMode                            typedef.DateMode           // Display mode for the date
	DisplayOrientation                  typedef.DisplayOrientation
	MountingSide                        typedef.Side
	AutoSyncFrequency                   typedef.AutoSyncFrequency  // Helps to conserve battery by changing modes
	NumberOfScreens                     uint8                      // Number of screens configured to display
	SmartNotificationDisplayOrientation typedef.DisplayOrientation // Smart Notification display orientation
	TapInterface                        typedef.Switch
	TapSensitivity                      typedef.TapSensitivity // Used to hold the tap threshold setting
	ActivityTrackerEnabled              bool                   // Enabled state of the activity tracker functionality
	MoveAlertEnabled                    bool                   // Enabled state of the move alert
	LactateThresholdAutodetectEnabled   bool                   // Enable auto-detect setting for the lactate threshold feature.
	BleAutoUploadEnabled                bool                   // Automatically upload using BLE

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDeviceSettings creates new DeviceSettings struct based on given mesg.
// If mesg is nil, it will return DeviceSettings with all fields being set to its corresponding invalid value.
func NewDeviceSettings(mesg *proto.Message) *DeviceSettings {
	vals := [175]proto.Value{}

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

	return &DeviceSettings{
		TimeOffset: vals[2].SliceUint32(),
		TimeMode: func() []typedef.TimeMode {
			sliceValue := vals[4].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.TimeMode)(ptr), len(sliceValue))
		}(),
		TimeZoneOffset:                      vals[5].SliceInt8(),
		ClockTime:                           datetime.ToTime(vals[39].Uint32()),
		PagesEnabled:                        vals[40].SliceUint16(),
		DefaultPage:                         vals[57].SliceUint16(),
		UtcOffset:                           vals[1].Uint32(),
		AutoActivityDetect:                  typedef.AutoActivityDetect(vals[90].Uint32()),
		AutosyncMinSteps:                    vals[58].Uint16(),
		AutosyncMinTime:                     vals[59].Uint16(),
		ActiveTimeZone:                      vals[0].Uint8(),
		BacklightMode:                       typedef.BacklightMode(vals[12].Uint8()),
		DateMode:                            typedef.DateMode(vals[47].Uint8()),
		DisplayOrientation:                  typedef.DisplayOrientation(vals[55].Uint8()),
		MountingSide:                        typedef.Side(vals[56].Uint8()),
		AutoSyncFrequency:                   typedef.AutoSyncFrequency(vals[89].Uint8()),
		NumberOfScreens:                     vals[94].Uint8(),
		SmartNotificationDisplayOrientation: typedef.DisplayOrientation(vals[95].Uint8()),
		TapInterface:                        typedef.Switch(vals[134].Uint8()),
		TapSensitivity:                      typedef.TapSensitivity(vals[174].Uint8()),
		ActivityTrackerEnabled:              vals[36].Bool(),
		MoveAlertEnabled:                    vals[46].Bool(),
		LactateThresholdAutodetectEnabled:   vals[80].Bool(),
		BleAutoUploadEnabled:                vals[86].Bool(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts DeviceSettings into proto.Message. If options is nil, default options will be used.
func (m *DeviceSettings) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumDeviceSettings)

	if m.TimeOffset != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint32(m.TimeOffset)
		fields = append(fields, field)
	}
	if m.TimeMode != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint8(unsafe.Slice((*byte)(unsafe.SliceData(m.TimeMode)), len(m.TimeMode)))
		fields = append(fields, field)
	}
	if m.TimeZoneOffset != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceInt8(m.TimeZoneOffset)
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.ClockTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 39)
		field.Value = proto.Uint32(datetime.ToUint32(m.ClockTime))
		fields = append(fields, field)
	}
	if m.PagesEnabled != nil {
		field := fac.CreateField(mesg.Num, 40)
		field.Value = proto.SliceUint16(m.PagesEnabled)
		fields = append(fields, field)
	}
	if m.DefaultPage != nil {
		field := fac.CreateField(mesg.Num, 57)
		field.Value = proto.SliceUint16(m.DefaultPage)
		fields = append(fields, field)
	}
	if m.UtcOffset != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.UtcOffset)
		fields = append(fields, field)
	}
	if uint32(m.AutoActivityDetect) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 90)
		field.Value = proto.Uint32(uint32(m.AutoActivityDetect))
		fields = append(fields, field)
	}
	if m.AutosyncMinSteps != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 58)
		field.Value = proto.Uint16(m.AutosyncMinSteps)
		fields = append(fields, field)
	}
	if m.AutosyncMinTime != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 59)
		field.Value = proto.Uint16(m.AutosyncMinTime)
		fields = append(fields, field)
	}
	if m.ActiveTimeZone != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.ActiveTimeZone)
		fields = append(fields, field)
	}
	if byte(m.BacklightMode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(byte(m.BacklightMode))
		fields = append(fields, field)
	}
	if byte(m.DateMode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 47)
		field.Value = proto.Uint8(byte(m.DateMode))
		fields = append(fields, field)
	}
	if byte(m.DisplayOrientation) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 55)
		field.Value = proto.Uint8(byte(m.DisplayOrientation))
		fields = append(fields, field)
	}
	if byte(m.MountingSide) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 56)
		field.Value = proto.Uint8(byte(m.MountingSide))
		fields = append(fields, field)
	}
	if byte(m.AutoSyncFrequency) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 89)
		field.Value = proto.Uint8(byte(m.AutoSyncFrequency))
		fields = append(fields, field)
	}
	if m.NumberOfScreens != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 94)
		field.Value = proto.Uint8(m.NumberOfScreens)
		fields = append(fields, field)
	}
	if byte(m.SmartNotificationDisplayOrientation) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 95)
		field.Value = proto.Uint8(byte(m.SmartNotificationDisplayOrientation))
		fields = append(fields, field)
	}
	if byte(m.TapInterface) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 134)
		field.Value = proto.Uint8(byte(m.TapInterface))
		fields = append(fields, field)
	}
	if byte(m.TapSensitivity) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 174)
		field.Value = proto.Uint8(byte(m.TapSensitivity))
		fields = append(fields, field)
	}
	if m.ActivityTrackerEnabled != false {
		field := fac.CreateField(mesg.Num, 36)
		field.Value = proto.Bool(m.ActivityTrackerEnabled)
		fields = append(fields, field)
	}
	if m.MoveAlertEnabled != false {
		field := fac.CreateField(mesg.Num, 46)
		field.Value = proto.Bool(m.MoveAlertEnabled)
		fields = append(fields, field)
	}
	if m.LactateThresholdAutodetectEnabled != false {
		field := fac.CreateField(mesg.Num, 80)
		field.Value = proto.Bool(m.LactateThresholdAutodetectEnabled)
		fields = append(fields, field)
	}
	if m.BleAutoUploadEnabled != false {
		field := fac.CreateField(mesg.Num, 86)
		field.Value = proto.Bool(m.BleAutoUploadEnabled)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimeZoneOffsetScaled return TimeZoneOffset in its scaled value [Array: [N]; Scale: 4; Units: hr; timezone offset in 1/4 hour increments].
//
// If TimeZoneOffset value is invalid, nil will be returned.
func (m *DeviceSettings) TimeZoneOffsetScaled() []float64 {
	if m.TimeZoneOffset == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.TimeZoneOffset, 4, 0)
}

// SetTimeOffset sets DeviceSettings value.
//
// Array: [N]; Units: s; Offset from system time.
func (m *DeviceSettings) SetTimeOffset(v []uint32) *DeviceSettings {
	m.TimeOffset = v
	return m
}

// SetTimeMode sets DeviceSettings value.
//
// Array: [N]; Display mode for the time
func (m *DeviceSettings) SetTimeMode(v []typedef.TimeMode) *DeviceSettings {
	m.TimeMode = v
	return m
}

// SetTimeZoneOffset sets DeviceSettings value.
//
// Array: [N]; Scale: 4; Units: hr; timezone offset in 1/4 hour increments
func (m *DeviceSettings) SetTimeZoneOffset(v []int8) *DeviceSettings {
	m.TimeZoneOffset = v
	return m
}

// SetClockTime sets DeviceSettings value.
//
// UTC timestamp used to set the devices clock and date
func (m *DeviceSettings) SetClockTime(v time.Time) *DeviceSettings {
	m.ClockTime = v
	return m
}

// SetPagesEnabled sets DeviceSettings value.
//
// Array: [N]; Bitfield to configure enabled screens for each supported loop
func (m *DeviceSettings) SetPagesEnabled(v []uint16) *DeviceSettings {
	m.PagesEnabled = v
	return m
}

// SetDefaultPage sets DeviceSettings value.
//
// Array: [N]; Bitfield to indicate one page as default for each supported loop
func (m *DeviceSettings) SetDefaultPage(v []uint16) *DeviceSettings {
	m.DefaultPage = v
	return m
}

// SetUtcOffset sets DeviceSettings value.
//
// Offset from system time. Required to convert timestamp from system time to UTC.
func (m *DeviceSettings) SetUtcOffset(v uint32) *DeviceSettings {
	m.UtcOffset = v
	return m
}

// SetAutoActivityDetect sets DeviceSettings value.
//
// Allows setting specific activities auto-activity detect enabled/disabled settings
func (m *DeviceSettings) SetAutoActivityDetect(v typedef.AutoActivityDetect) *DeviceSettings {
	m.AutoActivityDetect = v
	return m
}

// SetAutosyncMinSteps sets DeviceSettings value.
//
// Units: steps; Minimum steps before an autosync can occur
func (m *DeviceSettings) SetAutosyncMinSteps(v uint16) *DeviceSettings {
	m.AutosyncMinSteps = v
	return m
}

// SetAutosyncMinTime sets DeviceSettings value.
//
// Units: minutes; Minimum minutes before an autosync can occur
func (m *DeviceSettings) SetAutosyncMinTime(v uint16) *DeviceSettings {
	m.AutosyncMinTime = v
	return m
}

// SetActiveTimeZone sets DeviceSettings value.
//
// Index into time zone arrays.
func (m *DeviceSettings) SetActiveTimeZone(v uint8) *DeviceSettings {
	m.ActiveTimeZone = v
	return m
}

// SetBacklightMode sets DeviceSettings value.
//
// Mode for backlight
func (m *DeviceSettings) SetBacklightMode(v typedef.BacklightMode) *DeviceSettings {
	m.BacklightMode = v
	return m
}

// SetDateMode sets DeviceSettings value.
//
// Display mode for the date
func (m *DeviceSettings) SetDateMode(v typedef.DateMode) *DeviceSettings {
	m.DateMode = v
	return m
}

// SetDisplayOrientation sets DeviceSettings value.
func (m *DeviceSettings) SetDisplayOrientation(v typedef.DisplayOrientation) *DeviceSettings {
	m.DisplayOrientation = v
	return m
}

// SetMountingSide sets DeviceSettings value.
func (m *DeviceSettings) SetMountingSide(v typedef.Side) *DeviceSettings {
	m.MountingSide = v
	return m
}

// SetAutoSyncFrequency sets DeviceSettings value.
//
// Helps to conserve battery by changing modes
func (m *DeviceSettings) SetAutoSyncFrequency(v typedef.AutoSyncFrequency) *DeviceSettings {
	m.AutoSyncFrequency = v
	return m
}

// SetNumberOfScreens sets DeviceSettings value.
//
// Number of screens configured to display
func (m *DeviceSettings) SetNumberOfScreens(v uint8) *DeviceSettings {
	m.NumberOfScreens = v
	return m
}

// SetSmartNotificationDisplayOrientation sets DeviceSettings value.
//
// Smart Notification display orientation
func (m *DeviceSettings) SetSmartNotificationDisplayOrientation(v typedef.DisplayOrientation) *DeviceSettings {
	m.SmartNotificationDisplayOrientation = v
	return m
}

// SetTapInterface sets DeviceSettings value.
func (m *DeviceSettings) SetTapInterface(v typedef.Switch) *DeviceSettings {
	m.TapInterface = v
	return m
}

// SetTapSensitivity sets DeviceSettings value.
//
// Used to hold the tap threshold setting
func (m *DeviceSettings) SetTapSensitivity(v typedef.TapSensitivity) *DeviceSettings {
	m.TapSensitivity = v
	return m
}

// SetActivityTrackerEnabled sets DeviceSettings value.
//
// Enabled state of the activity tracker functionality
func (m *DeviceSettings) SetActivityTrackerEnabled(v bool) *DeviceSettings {
	m.ActivityTrackerEnabled = v
	return m
}

// SetMoveAlertEnabled sets DeviceSettings value.
//
// Enabled state of the move alert
func (m *DeviceSettings) SetMoveAlertEnabled(v bool) *DeviceSettings {
	m.MoveAlertEnabled = v
	return m
}

// SetLactateThresholdAutodetectEnabled sets DeviceSettings value.
//
// Enable auto-detect setting for the lactate threshold feature.
func (m *DeviceSettings) SetLactateThresholdAutodetectEnabled(v bool) *DeviceSettings {
	m.LactateThresholdAutodetectEnabled = v
	return m
}

// SetBleAutoUploadEnabled sets DeviceSettings value.
//
// Automatically upload using BLE
func (m *DeviceSettings) SetBleAutoUploadEnabled(v bool) *DeviceSettings {
	m.BleAutoUploadEnabled = v
	return m
}

// SetDeveloperFields DeviceSettings's DeveloperFields.
func (m *DeviceSettings) SetDeveloperFields(developerFields ...proto.DeveloperField) *DeviceSettings {
	m.DeveloperFields = developerFields
	return m
}
