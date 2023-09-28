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

// UserProfile is a UserProfile message.
type UserProfile struct {
	MessageIndex               typedef.MessageIndex
	FriendlyName               string
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

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		0:   basetype.StringInvalid, /* FriendlyName */
		1:   basetype.EnumInvalid,   /* Gender */
		2:   basetype.Uint8Invalid,  /* Age */
		3:   basetype.Uint8Invalid,  /* Height */
		4:   basetype.Uint16Invalid, /* Weight */
		5:   basetype.EnumInvalid,   /* Language */
		6:   basetype.EnumInvalid,   /* ElevSetting */
		7:   basetype.EnumInvalid,   /* WeightSetting */
		8:   basetype.Uint8Invalid,  /* RestingHeartRate */
		9:   basetype.Uint8Invalid,  /* DefaultMaxRunningHeartRate */
		10:  basetype.Uint8Invalid,  /* DefaultMaxBikingHeartRate */
		11:  basetype.Uint8Invalid,  /* DefaultMaxHeartRate */
		12:  basetype.EnumInvalid,   /* HrSetting */
		13:  basetype.EnumInvalid,   /* SpeedSetting */
		14:  basetype.EnumInvalid,   /* DistSetting */
		16:  basetype.EnumInvalid,   /* PowerSetting */
		17:  basetype.EnumInvalid,   /* ActivityClass */
		18:  basetype.EnumInvalid,   /* PositionSetting */
		21:  basetype.EnumInvalid,   /* TemperatureSetting */
		22:  basetype.Uint16Invalid, /* LocalId */
		23:  nil,                    /* GlobalId */
		28:  basetype.Uint32Invalid, /* WakeTime */
		29:  basetype.Uint32Invalid, /* SleepTime */
		30:  basetype.EnumInvalid,   /* HeightSetting */
		31:  basetype.Uint16Invalid, /* UserRunningStepLength */
		32:  basetype.Uint16Invalid, /* UserWalkingStepLength */
		47:  basetype.EnumInvalid,   /* DepthSetting */
		49:  basetype.Uint32Invalid, /* DiveCount */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
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

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to UserProfile mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumUserProfile)
func (m UserProfile) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumUserProfile {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.FriendlyName,
		1:   m.Gender,
		2:   m.Age,
		3:   m.Height,
		4:   m.Weight,
		5:   m.Language,
		6:   m.ElevSetting,
		7:   m.WeightSetting,
		8:   m.RestingHeartRate,
		9:   m.DefaultMaxRunningHeartRate,
		10:  m.DefaultMaxBikingHeartRate,
		11:  m.DefaultMaxHeartRate,
		12:  m.HrSetting,
		13:  m.SpeedSetting,
		14:  m.DistSetting,
		16:  m.PowerSetting,
		17:  m.ActivityClass,
		18:  m.PositionSetting,
		21:  m.TemperatureSetting,
		22:  m.LocalId,
		23:  m.GlobalId,
		28:  m.WakeTime,
		29:  m.SleepTime,
		30:  m.HeightSetting,
		31:  m.UserRunningStepLength,
		32:  m.UserWalkingStepLength,
		47:  m.DepthSetting,
		49:  m.DiveCount,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
