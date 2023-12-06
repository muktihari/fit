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

// BikeProfile is a BikeProfile message.
type BikeProfile struct {
	MessageIndex             typedef.MessageIndex
	Name                     string
	Sport                    typedef.Sport
	SubSport                 typedef.SubSport
	Odometer                 uint32 // Scale: 100; Units: m;
	BikeSpdAntId             uint16
	BikeCadAntId             uint16
	BikeSpdcadAntId          uint16
	BikePowerAntId           uint16
	CustomWheelsize          uint16 // Scale: 1000; Units: m;
	AutoWheelsize            uint16 // Scale: 1000; Units: m;
	BikeWeight               uint16 // Scale: 10; Units: kg;
	PowerCalFactor           uint16 // Scale: 10; Units: %;
	AutoWheelCal             bool
	AutoPowerZero            bool
	Id                       uint8
	SpdEnabled               bool
	CadEnabled               bool
	SpdcadEnabled            bool
	PowerEnabled             bool
	CrankLength              uint8 // Scale: 2; Offset: -110; Units: mm;
	Enabled                  bool
	BikeSpdAntIdTransType    uint8
	BikeCadAntIdTransType    uint8
	BikeSpdcadAntIdTransType uint8
	BikePowerAntIdTransType  uint8
	OdometerRollover         uint8   // Rollover counter that can be used to extend the odometer
	FrontGearNum             uint8   // Number of front gears
	FrontGear                []uint8 // Array: [N]; Number of teeth on each gear 0 is innermost
	RearGearNum              uint8   // Number of rear gears
	RearGear                 []uint8 // Array: [N]; Number of teeth on each gear 0 is innermost
	ShimanoDi2Enabled        bool

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewBikeProfile creates new BikeProfile struct based on given mesg. If mesg is nil or mesg.Num is not equal to BikeProfile mesg number, it will return nil.
func NewBikeProfile(mesg proto.Message) *BikeProfile {
	if mesg.Num != typedef.MesgNumBikeProfile {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid,  /* MessageIndex */
		0:   basetype.StringInvalid,  /* Name */
		1:   basetype.EnumInvalid,    /* Sport */
		2:   basetype.EnumInvalid,    /* SubSport */
		3:   basetype.Uint32Invalid,  /* Odometer */
		4:   basetype.Uint16zInvalid, /* BikeSpdAntId */
		5:   basetype.Uint16zInvalid, /* BikeCadAntId */
		6:   basetype.Uint16zInvalid, /* BikeSpdcadAntId */
		7:   basetype.Uint16zInvalid, /* BikePowerAntId */
		8:   basetype.Uint16Invalid,  /* CustomWheelsize */
		9:   basetype.Uint16Invalid,  /* AutoWheelsize */
		10:  basetype.Uint16Invalid,  /* BikeWeight */
		11:  basetype.Uint16Invalid,  /* PowerCalFactor */
		12:  false,                   /* AutoWheelCal */
		13:  false,                   /* AutoPowerZero */
		14:  basetype.Uint8Invalid,   /* Id */
		15:  false,                   /* SpdEnabled */
		16:  false,                   /* CadEnabled */
		17:  false,                   /* SpdcadEnabled */
		18:  false,                   /* PowerEnabled */
		19:  basetype.Uint8Invalid,   /* CrankLength */
		20:  false,                   /* Enabled */
		21:  basetype.Uint8zInvalid,  /* BikeSpdAntIdTransType */
		22:  basetype.Uint8zInvalid,  /* BikeCadAntIdTransType */
		23:  basetype.Uint8zInvalid,  /* BikeSpdcadAntIdTransType */
		24:  basetype.Uint8zInvalid,  /* BikePowerAntIdTransType */
		37:  basetype.Uint8Invalid,   /* OdometerRollover */
		38:  basetype.Uint8zInvalid,  /* FrontGearNum */
		39:  nil,                     /* FrontGear */
		40:  basetype.Uint8zInvalid,  /* RearGearNum */
		41:  nil,                     /* RearGear */
		44:  false,                   /* ShimanoDi2Enabled */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &BikeProfile{
		MessageIndex:             typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Name:                     typeconv.ToString[string](vals[0]),
		Sport:                    typeconv.ToEnum[typedef.Sport](vals[1]),
		SubSport:                 typeconv.ToEnum[typedef.SubSport](vals[2]),
		Odometer:                 typeconv.ToUint32[uint32](vals[3]),
		BikeSpdAntId:             typeconv.ToUint16z[uint16](vals[4]),
		BikeCadAntId:             typeconv.ToUint16z[uint16](vals[5]),
		BikeSpdcadAntId:          typeconv.ToUint16z[uint16](vals[6]),
		BikePowerAntId:           typeconv.ToUint16z[uint16](vals[7]),
		CustomWheelsize:          typeconv.ToUint16[uint16](vals[8]),
		AutoWheelsize:            typeconv.ToUint16[uint16](vals[9]),
		BikeWeight:               typeconv.ToUint16[uint16](vals[10]),
		PowerCalFactor:           typeconv.ToUint16[uint16](vals[11]),
		AutoWheelCal:             typeconv.ToBool[bool](vals[12]),
		AutoPowerZero:            typeconv.ToBool[bool](vals[13]),
		Id:                       typeconv.ToUint8[uint8](vals[14]),
		SpdEnabled:               typeconv.ToBool[bool](vals[15]),
		CadEnabled:               typeconv.ToBool[bool](vals[16]),
		SpdcadEnabled:            typeconv.ToBool[bool](vals[17]),
		PowerEnabled:             typeconv.ToBool[bool](vals[18]),
		CrankLength:              typeconv.ToUint8[uint8](vals[19]),
		Enabled:                  typeconv.ToBool[bool](vals[20]),
		BikeSpdAntIdTransType:    typeconv.ToUint8z[uint8](vals[21]),
		BikeCadAntIdTransType:    typeconv.ToUint8z[uint8](vals[22]),
		BikeSpdcadAntIdTransType: typeconv.ToUint8z[uint8](vals[23]),
		BikePowerAntIdTransType:  typeconv.ToUint8z[uint8](vals[24]),
		OdometerRollover:         typeconv.ToUint8[uint8](vals[37]),
		FrontGearNum:             typeconv.ToUint8z[uint8](vals[38]),
		FrontGear:                typeconv.ToSliceUint8z[uint8](vals[39]),
		RearGearNum:              typeconv.ToUint8z[uint8](vals[40]),
		RearGear:                 typeconv.ToSliceUint8z[uint8](vals[41]),
		ShimanoDi2Enabled:        typeconv.ToBool[bool](vals[44]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to BikeProfile mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumBikeProfile)
func (m BikeProfile) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumBikeProfile {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.Name,
		1:   m.Sport,
		2:   m.SubSport,
		3:   m.Odometer,
		4:   m.BikeSpdAntId,
		5:   m.BikeCadAntId,
		6:   m.BikeSpdcadAntId,
		7:   m.BikePowerAntId,
		8:   m.CustomWheelsize,
		9:   m.AutoWheelsize,
		10:  m.BikeWeight,
		11:  m.PowerCalFactor,
		12:  m.AutoWheelCal,
		13:  m.AutoPowerZero,
		14:  m.Id,
		15:  m.SpdEnabled,
		16:  m.CadEnabled,
		17:  m.SpdcadEnabled,
		18:  m.PowerEnabled,
		19:  m.CrankLength,
		20:  m.Enabled,
		21:  m.BikeSpdAntIdTransType,
		22:  m.BikeCadAntIdTransType,
		23:  m.BikeSpdcadAntIdTransType,
		24:  m.BikePowerAntIdTransType,
		37:  m.OdometerRollover,
		38:  m.FrontGearNum,
		39:  m.FrontGear,
		40:  m.RearGearNum,
		41:  m.RearGear,
		44:  m.ShimanoDi2Enabled,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
