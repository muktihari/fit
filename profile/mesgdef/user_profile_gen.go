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

// UserProfile is a UserProfile message.
type UserProfile struct {
	MessageIndex               typedef.MessageIndex
	FriendlyName               string // Used for Morning Report greeting
	Gender                     typedef.Gender
	Age                        uint8  // Units: years;
	Height                     uint8  // Scale: 100; Units: m;
	Weight                     uint16 // Scale: 10; Units: kg;
	Language                   typedef.Language
	ElevSetting                typedef.DisplayMeasure
	WeightSetting              typedef.DisplayMeasure
	RestingHeartRate           uint8 // Units: bpm;
	DefaultMaxRunningHeartRate uint8 // Units: bpm;
	DefaultMaxBikingHeartRate  uint8 // Units: bpm;
	DefaultMaxHeartRate        uint8 // Units: bpm;
	HrSetting                  typedef.DisplayHeart
	SpeedSetting               typedef.DisplayMeasure
	DistSetting                typedef.DisplayMeasure
	PowerSetting               typedef.DisplayPower
	ActivityClass              typedef.ActivityClass
	PositionSetting            typedef.DisplayPosition
	TemperatureSetting         typedef.DisplayMeasure
	LocalId                    typedef.UserLocalId
	GlobalId                   []byte                   // Array: [6];
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

// NewUserProfile creates new UserProfile struct based on given mesg. If mesg is nil or mesg.Num is not equal to UserProfile mesg number, it will return nil.
func NewUserProfile(mesg proto.Message) *UserProfile {
	if mesg.Num != typedef.MesgNumUserProfile {
		return nil
	}

	vals := [255]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
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

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts UserProfile into proto.Message.
func (m *UserProfile) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumUserProfile)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FriendlyName != basetype.StringInvalid && m.FriendlyName != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.FriendlyName
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Gender) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToEnum[byte](m.Gender)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Age != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Age
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Height != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Height
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Weight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.Weight
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Language) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = typeconv.ToEnum[byte](m.Language)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.ElevSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = typeconv.ToEnum[byte](m.ElevSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.WeightSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToEnum[byte](m.WeightSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.RestingHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DefaultMaxRunningHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.DefaultMaxRunningHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DefaultMaxBikingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.DefaultMaxBikingHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DefaultMaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.DefaultMaxHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.HrSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = typeconv.ToEnum[byte](m.HrSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.SpeedSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = typeconv.ToEnum[byte](m.SpeedSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.DistSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = typeconv.ToEnum[byte](m.DistSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.PowerSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = typeconv.ToEnum[byte](m.PowerSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.ActivityClass) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = typeconv.ToEnum[byte](m.ActivityClass)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.PositionSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = typeconv.ToEnum[byte](m.PositionSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.TemperatureSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = typeconv.ToEnum[byte](m.TemperatureSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.LocalId) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = typeconv.ToUint16[uint16](m.LocalId)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.GlobalId != nil {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.GlobalId
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32[uint32](m.WakeTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 28)
		field.Value = typeconv.ToUint32[uint32](m.WakeTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32[uint32](m.SleepTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 29)
		field.Value = typeconv.ToUint32[uint32](m.SleepTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.HeightSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 30)
		field.Value = typeconv.ToEnum[byte](m.HeightSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.UserRunningStepLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 31)
		field.Value = m.UserRunningStepLength
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.UserWalkingStepLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = m.UserWalkingStepLength
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.DepthSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 47)
		field.Value = typeconv.ToEnum[byte](m.DepthSetting)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DiveCount != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 49)
		field.Value = m.DiveCount
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of UserProfile's valid fields.
func (m *UserProfile) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if m.FriendlyName != basetype.StringInvalid && m.FriendlyName != "" {
		size++
	}
	if typeconv.ToEnum[byte](m.Gender) != basetype.EnumInvalid {
		size++
	}
	if m.Age != basetype.Uint8Invalid {
		size++
	}
	if m.Height != basetype.Uint8Invalid {
		size++
	}
	if m.Weight != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Language) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.ElevSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.WeightSetting) != basetype.EnumInvalid {
		size++
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.DefaultMaxRunningHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.DefaultMaxBikingHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.DefaultMaxHeartRate != basetype.Uint8Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.HrSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.SpeedSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.DistSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.PowerSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.ActivityClass) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.PositionSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.TemperatureSetting) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.LocalId) != basetype.Uint16Invalid {
		size++
	}
	if m.GlobalId != nil {
		size++
	}
	if typeconv.ToUint32[uint32](m.WakeTime) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToUint32[uint32](m.SleepTime) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.HeightSetting) != basetype.EnumInvalid {
		size++
	}
	if m.UserRunningStepLength != basetype.Uint16Invalid {
		size++
	}
	if m.UserWalkingStepLength != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.DepthSetting) != basetype.EnumInvalid {
		size++
	}
	if m.DiveCount != basetype.Uint32Invalid {
		size++
	}
	return size
}
