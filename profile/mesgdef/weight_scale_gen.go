// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

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

// WeightScale is a WeightScale message.
type WeightScale struct {
	Timestamp         typedef.DateTime // Units: s;
	Weight            typedef.Weight   // Scale: 100; Units: kg;
	PercentFat        uint16           // Scale: 100; Units: %;
	PercentHydration  uint16           // Scale: 100; Units: %;
	VisceralFatMass   uint16           // Scale: 100; Units: kg;
	BoneMass          uint16           // Scale: 100; Units: kg;
	MuscleMass        uint16           // Scale: 100; Units: kg;
	BasalMet          uint16           // Scale: 4; Units: kcal/day;
	PhysiqueRating    uint8
	ActiveMet         uint16 // Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal
	MetabolicAge      uint8  // Units: years;
	VisceralFatRating uint8
	UserProfileIndex  typedef.MessageIndex // Associates this weight scale message to a user. This corresponds to the index of the user profile message in the weight scale file.
	Bmi               uint16               // Scale: 10; Units: kg/m^2;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWeightScale creates new WeightScale struct based on given mesg. If mesg is nil or mesg.Num is not equal to WeightScale mesg number, it will return nil.
func NewWeightScale(mesg proto.Message) *WeightScale {
	if mesg.Num != typedef.MesgNumWeightScale {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* Weight */
		1:   basetype.Uint16Invalid, /* PercentFat */
		2:   basetype.Uint16Invalid, /* PercentHydration */
		3:   basetype.Uint16Invalid, /* VisceralFatMass */
		4:   basetype.Uint16Invalid, /* BoneMass */
		5:   basetype.Uint16Invalid, /* MuscleMass */
		7:   basetype.Uint16Invalid, /* BasalMet */
		8:   basetype.Uint8Invalid,  /* PhysiqueRating */
		9:   basetype.Uint16Invalid, /* ActiveMet */
		10:  basetype.Uint8Invalid,  /* MetabolicAge */
		11:  basetype.Uint8Invalid,  /* VisceralFatRating */
		12:  basetype.Uint16Invalid, /* UserProfileIndex */
		13:  basetype.Uint16Invalid, /* Bmi */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &WeightScale{
		Timestamp:         typeconv.ToUint32[typedef.DateTime](vals[253]),
		Weight:            typeconv.ToUint16[typedef.Weight](vals[0]),
		PercentFat:        typeconv.ToUint16[uint16](vals[1]),
		PercentHydration:  typeconv.ToUint16[uint16](vals[2]),
		VisceralFatMass:   typeconv.ToUint16[uint16](vals[3]),
		BoneMass:          typeconv.ToUint16[uint16](vals[4]),
		MuscleMass:        typeconv.ToUint16[uint16](vals[5]),
		BasalMet:          typeconv.ToUint16[uint16](vals[7]),
		PhysiqueRating:    typeconv.ToUint8[uint8](vals[8]),
		ActiveMet:         typeconv.ToUint16[uint16](vals[9]),
		MetabolicAge:      typeconv.ToUint8[uint8](vals[10]),
		VisceralFatRating: typeconv.ToUint8[uint8](vals[11]),
		UserProfileIndex:  typeconv.ToUint16[typedef.MessageIndex](vals[12]),
		Bmi:               typeconv.ToUint16[uint16](vals[13]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to WeightScale mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumWeightScale)
func (m WeightScale) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumWeightScale {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.Weight,
		1:   m.PercentFat,
		2:   m.PercentHydration,
		3:   m.VisceralFatMass,
		4:   m.BoneMass,
		5:   m.MuscleMass,
		7:   m.BasalMet,
		8:   m.PhysiqueRating,
		9:   m.ActiveMet,
		10:  m.MetabolicAge,
		11:  m.VisceralFatRating,
		12:  m.UserProfileIndex,
		13:  m.Bmi,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
