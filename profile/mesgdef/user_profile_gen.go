// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// UserProfile is a UserProfile message.
type UserProfile struct {
	MessageIndex               typedef.MessageIndex
	FriendlyName               string // Used for Morning Report greeting
	Gender                     typedef.Gender
	Age                        uint8  // Units: years
	Height                     uint8  // Scale: 100; Units: m
	Weight                     uint16 // Scale: 10; Units: kg
	Language                   typedef.Language
	ElevSetting                typedef.DisplayMeasure
	WeightSetting              typedef.DisplayMeasure
	RestingHeartRate           uint8 // Units: bpm
	DefaultMaxRunningHeartRate uint8 // Units: bpm
	DefaultMaxBikingHeartRate  uint8 // Units: bpm
	DefaultMaxHeartRate        uint8 // Units: bpm
	HrSetting                  typedef.DisplayHeart
	SpeedSetting               typedef.DisplayMeasure
	DistSetting                typedef.DisplayMeasure
	PowerSetting               typedef.DisplayPower
	ActivityClass              typedef.ActivityClass
	PositionSetting            typedef.DisplayPosition
	TemperatureSetting         typedef.DisplayMeasure
	LocalId                    typedef.UserLocalId
	GlobalId                   []byte                   // Array: [6]
	WakeTime                   typedef.LocaltimeIntoDay // Typical wake time
	SleepTime                  typedef.LocaltimeIntoDay // Typical bed time
	HeightSetting              typedef.DisplayMeasure
	UserRunningStepLength      uint16 // Scale: 1000; Units: m; User defined running step length set to 0 for auto length
	UserWalkingStepLength      uint16 // Scale: 1000; Units: m; User defined walking step length set to 0 for auto length
	DepthSetting               typedef.DisplayMeasure
	DiveCount                  uint32

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewUserProfile creates new UserProfile struct based on given mesg.
// If mesg is nil, it will return UserProfile with all fields being set to its corresponding invalid value.
func NewUserProfile(mesg *proto.Message) *UserProfile {
	vals := [255]any{}

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

	return &UserProfile{
		MessageIndex:               typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		FriendlyName:               typeconv.ToString[string](vals[0]),
		Gender:                     typeconv.ToEnum[typedef.Gender](vals[1]),
		Age:                        typeconv.ToUint8[uint8](vals[2]),
		Height:                     typeconv.ToUint8[uint8](vals[3]),
		Weight:                     typeconv.ToUint16[uint16](vals[4]),
		Language:                   typeconv.ToEnum[typedef.Language](vals[5]),
		ElevSetting:                typeconv.ToEnum[typedef.DisplayMeasure](vals[6]),
		WeightSetting:              typeconv.ToEnum[typedef.DisplayMeasure](vals[7]),
		RestingHeartRate:           typeconv.ToUint8[uint8](vals[8]),
		DefaultMaxRunningHeartRate: typeconv.ToUint8[uint8](vals[9]),
		DefaultMaxBikingHeartRate:  typeconv.ToUint8[uint8](vals[10]),
		DefaultMaxHeartRate:        typeconv.ToUint8[uint8](vals[11]),
		HrSetting:                  typeconv.ToEnum[typedef.DisplayHeart](vals[12]),
		SpeedSetting:               typeconv.ToEnum[typedef.DisplayMeasure](vals[13]),
		DistSetting:                typeconv.ToEnum[typedef.DisplayMeasure](vals[14]),
		PowerSetting:               typeconv.ToEnum[typedef.DisplayPower](vals[16]),
		ActivityClass:              typeconv.ToEnum[typedef.ActivityClass](vals[17]),
		PositionSetting:            typeconv.ToEnum[typedef.DisplayPosition](vals[18]),
		TemperatureSetting:         typeconv.ToEnum[typedef.DisplayMeasure](vals[21]),
		LocalId:                    typeconv.ToUint16[typedef.UserLocalId](vals[22]),
		GlobalId:                   typeconv.ToSliceByte[byte](vals[23]),
		WakeTime:                   typeconv.ToUint32[typedef.LocaltimeIntoDay](vals[28]),
		SleepTime:                  typeconv.ToUint32[typedef.LocaltimeIntoDay](vals[29]),
		HeightSetting:              typeconv.ToEnum[typedef.DisplayMeasure](vals[30]),
		UserRunningStepLength:      typeconv.ToUint16[uint16](vals[31]),
		UserWalkingStepLength:      typeconv.ToUint16[uint16](vals[32]),
		DepthSetting:               typeconv.ToEnum[typedef.DisplayMeasure](vals[47]),
		DiveCount:                  typeconv.ToUint32[uint32](vals[49]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts UserProfile into proto.Message.
func (m *UserProfile) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumUserProfile)

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if m.FriendlyName != basetype.StringInvalid && m.FriendlyName != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.FriendlyName
		fields = append(fields, field)
	}
	if byte(m.Gender) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = byte(m.Gender)
		fields = append(fields, field)
	}
	if m.Age != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Age
		fields = append(fields, field)
	}
	if m.Height != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Height
		fields = append(fields, field)
	}
	if m.Weight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.Weight
		fields = append(fields, field)
	}
	if byte(m.Language) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = byte(m.Language)
		fields = append(fields, field)
	}
	if byte(m.ElevSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = byte(m.ElevSetting)
		fields = append(fields, field)
	}
	if byte(m.WeightSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = byte(m.WeightSetting)
		fields = append(fields, field)
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.RestingHeartRate
		fields = append(fields, field)
	}
	if m.DefaultMaxRunningHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.DefaultMaxRunningHeartRate
		fields = append(fields, field)
	}
	if m.DefaultMaxBikingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.DefaultMaxBikingHeartRate
		fields = append(fields, field)
	}
	if m.DefaultMaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.DefaultMaxHeartRate
		fields = append(fields, field)
	}
	if byte(m.HrSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = byte(m.HrSetting)
		fields = append(fields, field)
	}
	if byte(m.SpeedSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = byte(m.SpeedSetting)
		fields = append(fields, field)
	}
	if byte(m.DistSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = byte(m.DistSetting)
		fields = append(fields, field)
	}
	if byte(m.PowerSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = byte(m.PowerSetting)
		fields = append(fields, field)
	}
	if byte(m.ActivityClass) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = byte(m.ActivityClass)
		fields = append(fields, field)
	}
	if byte(m.PositionSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = byte(m.PositionSetting)
		fields = append(fields, field)
	}
	if byte(m.TemperatureSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = byte(m.TemperatureSetting)
		fields = append(fields, field)
	}
	if uint16(m.LocalId) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = uint16(m.LocalId)
		fields = append(fields, field)
	}
	if m.GlobalId != nil {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.GlobalId
		fields = append(fields, field)
	}
	if uint32(m.WakeTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 28)
		field.Value = uint32(m.WakeTime)
		fields = append(fields, field)
	}
	if uint32(m.SleepTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 29)
		field.Value = uint32(m.SleepTime)
		fields = append(fields, field)
	}
	if byte(m.HeightSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 30)
		field.Value = byte(m.HeightSetting)
		fields = append(fields, field)
	}
	if m.UserRunningStepLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 31)
		field.Value = m.UserRunningStepLength
		fields = append(fields, field)
	}
	if m.UserWalkingStepLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = m.UserWalkingStepLength
		fields = append(fields, field)
	}
	if byte(m.DepthSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 47)
		field.Value = byte(m.DepthSetting)
		fields = append(fields, field)
	}
	if m.DiveCount != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 49)
		field.Value = m.DiveCount
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// HeightScaled return Height in its scaled value [Scale: 100; Units: m].
//
// If Height value is invalid, float64 invalid value will be returned.
func (m *UserProfile) HeightScaled() float64 {
	if m.Height == basetype.Uint8Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Height, 100, 0)
}

// WeightScaled return Weight in its scaled value [Scale: 10; Units: kg].
//
// If Weight value is invalid, float64 invalid value will be returned.
func (m *UserProfile) WeightScaled() float64 {
	if m.Weight == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Weight, 10, 0)
}

// UserRunningStepLengthScaled return UserRunningStepLength in its scaled value [Scale: 1000; Units: m; User defined running step length set to 0 for auto length].
//
// If UserRunningStepLength value is invalid, float64 invalid value will be returned.
func (m *UserProfile) UserRunningStepLengthScaled() float64 {
	if m.UserRunningStepLength == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.UserRunningStepLength, 1000, 0)
}

// UserWalkingStepLengthScaled return UserWalkingStepLength in its scaled value [Scale: 1000; Units: m; User defined walking step length set to 0 for auto length].
//
// If UserWalkingStepLength value is invalid, float64 invalid value will be returned.
func (m *UserProfile) UserWalkingStepLengthScaled() float64 {
	if m.UserWalkingStepLength == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.UserWalkingStepLength, 1000, 0)
}

// SetMessageIndex sets UserProfile value.
func (m *UserProfile) SetMessageIndex(v typedef.MessageIndex) *UserProfile {
	m.MessageIndex = v
	return m
}

// SetFriendlyName sets UserProfile value.
//
// Used for Morning Report greeting
func (m *UserProfile) SetFriendlyName(v string) *UserProfile {
	m.FriendlyName = v
	return m
}

// SetGender sets UserProfile value.
func (m *UserProfile) SetGender(v typedef.Gender) *UserProfile {
	m.Gender = v
	return m
}

// SetAge sets UserProfile value.
//
// Units: years
func (m *UserProfile) SetAge(v uint8) *UserProfile {
	m.Age = v
	return m
}

// SetHeight sets UserProfile value.
//
// Scale: 100; Units: m
func (m *UserProfile) SetHeight(v uint8) *UserProfile {
	m.Height = v
	return m
}

// SetWeight sets UserProfile value.
//
// Scale: 10; Units: kg
func (m *UserProfile) SetWeight(v uint16) *UserProfile {
	m.Weight = v
	return m
}

// SetLanguage sets UserProfile value.
func (m *UserProfile) SetLanguage(v typedef.Language) *UserProfile {
	m.Language = v
	return m
}

// SetElevSetting sets UserProfile value.
func (m *UserProfile) SetElevSetting(v typedef.DisplayMeasure) *UserProfile {
	m.ElevSetting = v
	return m
}

// SetWeightSetting sets UserProfile value.
func (m *UserProfile) SetWeightSetting(v typedef.DisplayMeasure) *UserProfile {
	m.WeightSetting = v
	return m
}

// SetRestingHeartRate sets UserProfile value.
//
// Units: bpm
func (m *UserProfile) SetRestingHeartRate(v uint8) *UserProfile {
	m.RestingHeartRate = v
	return m
}

// SetDefaultMaxRunningHeartRate sets UserProfile value.
//
// Units: bpm
func (m *UserProfile) SetDefaultMaxRunningHeartRate(v uint8) *UserProfile {
	m.DefaultMaxRunningHeartRate = v
	return m
}

// SetDefaultMaxBikingHeartRate sets UserProfile value.
//
// Units: bpm
func (m *UserProfile) SetDefaultMaxBikingHeartRate(v uint8) *UserProfile {
	m.DefaultMaxBikingHeartRate = v
	return m
}

// SetDefaultMaxHeartRate sets UserProfile value.
//
// Units: bpm
func (m *UserProfile) SetDefaultMaxHeartRate(v uint8) *UserProfile {
	m.DefaultMaxHeartRate = v
	return m
}

// SetHrSetting sets UserProfile value.
func (m *UserProfile) SetHrSetting(v typedef.DisplayHeart) *UserProfile {
	m.HrSetting = v
	return m
}

// SetSpeedSetting sets UserProfile value.
func (m *UserProfile) SetSpeedSetting(v typedef.DisplayMeasure) *UserProfile {
	m.SpeedSetting = v
	return m
}

// SetDistSetting sets UserProfile value.
func (m *UserProfile) SetDistSetting(v typedef.DisplayMeasure) *UserProfile {
	m.DistSetting = v
	return m
}

// SetPowerSetting sets UserProfile value.
func (m *UserProfile) SetPowerSetting(v typedef.DisplayPower) *UserProfile {
	m.PowerSetting = v
	return m
}

// SetActivityClass sets UserProfile value.
func (m *UserProfile) SetActivityClass(v typedef.ActivityClass) *UserProfile {
	m.ActivityClass = v
	return m
}

// SetPositionSetting sets UserProfile value.
func (m *UserProfile) SetPositionSetting(v typedef.DisplayPosition) *UserProfile {
	m.PositionSetting = v
	return m
}

// SetTemperatureSetting sets UserProfile value.
func (m *UserProfile) SetTemperatureSetting(v typedef.DisplayMeasure) *UserProfile {
	m.TemperatureSetting = v
	return m
}

// SetLocalId sets UserProfile value.
func (m *UserProfile) SetLocalId(v typedef.UserLocalId) *UserProfile {
	m.LocalId = v
	return m
}

// SetGlobalId sets UserProfile value.
//
// Array: [6]
func (m *UserProfile) SetGlobalId(v []byte) *UserProfile {
	m.GlobalId = v
	return m
}

// SetWakeTime sets UserProfile value.
//
// Typical wake time
func (m *UserProfile) SetWakeTime(v typedef.LocaltimeIntoDay) *UserProfile {
	m.WakeTime = v
	return m
}

// SetSleepTime sets UserProfile value.
//
// Typical bed time
func (m *UserProfile) SetSleepTime(v typedef.LocaltimeIntoDay) *UserProfile {
	m.SleepTime = v
	return m
}

// SetHeightSetting sets UserProfile value.
func (m *UserProfile) SetHeightSetting(v typedef.DisplayMeasure) *UserProfile {
	m.HeightSetting = v
	return m
}

// SetUserRunningStepLength sets UserProfile value.
//
// Scale: 1000; Units: m; User defined running step length set to 0 for auto length
func (m *UserProfile) SetUserRunningStepLength(v uint16) *UserProfile {
	m.UserRunningStepLength = v
	return m
}

// SetUserWalkingStepLength sets UserProfile value.
//
// Scale: 1000; Units: m; User defined walking step length set to 0 for auto length
func (m *UserProfile) SetUserWalkingStepLength(v uint16) *UserProfile {
	m.UserWalkingStepLength = v
	return m
}

// SetDepthSetting sets UserProfile value.
func (m *UserProfile) SetDepthSetting(v typedef.DisplayMeasure) *UserProfile {
	m.DepthSetting = v
	return m
}

// SetDiveCount sets UserProfile value.
func (m *UserProfile) SetDiveCount(v uint32) *UserProfile {
	m.DiveCount = v
	return m
}

// SetDeveloperFields UserProfile's DeveloperFields.
func (m *UserProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *UserProfile {
	m.DeveloperFields = developerFields
	return m
}
