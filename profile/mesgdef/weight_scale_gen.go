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
)

// WeightScale is a WeightScale message.
type WeightScale struct {
	Timestamp         time.Time            // Units: s
	Weight            typedef.Weight       // Scale: 100; Units: kg
	PercentFat        uint16               // Scale: 100; Units: %
	PercentHydration  uint16               // Scale: 100; Units: %
	VisceralFatMass   uint16               // Scale: 100; Units: kg
	BoneMass          uint16               // Scale: 100; Units: kg
	MuscleMass        uint16               // Scale: 100; Units: kg
	BasalMet          uint16               // Scale: 4; Units: kcal/day
	ActiveMet         uint16               // Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal
	UserProfileIndex  typedef.MessageIndex // Associates this weight scale message to a user. This corresponds to the index of the user profile message in the weight scale file.
	Bmi               uint16               // Scale: 10; Units: kg/m^2
	PhysiqueRating    uint8
	MetabolicAge      uint8 // Units: years
	VisceralFatRating uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWeightScale creates new WeightScale struct based on given mesg.
// If mesg is nil, it will return WeightScale with all fields being set to its corresponding invalid value.
func NewWeightScale(mesg *proto.Message) *WeightScale {
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

	return &WeightScale{
		Timestamp:         datetime.ToTime(vals[253].Uint32()),
		Weight:            typedef.Weight(vals[0].Uint16()),
		PercentFat:        vals[1].Uint16(),
		PercentHydration:  vals[2].Uint16(),
		VisceralFatMass:   vals[3].Uint16(),
		BoneMass:          vals[4].Uint16(),
		MuscleMass:        vals[5].Uint16(),
		BasalMet:          vals[7].Uint16(),
		ActiveMet:         vals[9].Uint16(),
		UserProfileIndex:  typedef.MessageIndex(vals[12].Uint16()),
		Bmi:               vals[13].Uint16(),
		PhysiqueRating:    vals[8].Uint8(),
		MetabolicAge:      vals[10].Uint8(),
		VisceralFatRating: vals[11].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WeightScale into proto.Message. If options is nil, default options will be used.
func (m *WeightScale) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumWeightScale)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if uint16(m.Weight) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.Weight))
		fields = append(fields, field)
	}
	if m.PercentFat != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.PercentFat)
		fields = append(fields, field)
	}
	if m.PercentHydration != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.PercentHydration)
		fields = append(fields, field)
	}
	if m.VisceralFatMass != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.VisceralFatMass)
		fields = append(fields, field)
	}
	if m.BoneMass != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.BoneMass)
		fields = append(fields, field)
	}
	if m.MuscleMass != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.MuscleMass)
		fields = append(fields, field)
	}
	if m.BasalMet != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint16(m.BasalMet)
		fields = append(fields, field)
	}
	if m.ActiveMet != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(m.ActiveMet)
		fields = append(fields, field)
	}
	if uint16(m.UserProfileIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint16(uint16(m.UserProfileIndex))
		fields = append(fields, field)
	}
	if m.Bmi != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint16(m.Bmi)
		fields = append(fields, field)
	}
	if m.PhysiqueRating != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(m.PhysiqueRating)
		fields = append(fields, field)
	}
	if m.MetabolicAge != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(m.MetabolicAge)
		fields = append(fields, field)
	}
	if m.VisceralFatRating != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.VisceralFatRating)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// WeightScaled return Weight in its scaled value [Scale: 100; Units: kg].
//
// If Weight value is invalid, float64 invalid value will be returned.
func (m *WeightScale) WeightScaled() float64 {
	if uint16(m.Weight) == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Weight, 100, 0)
}

// PercentFatScaled return PercentFat in its scaled value [Scale: 100; Units: %].
//
// If PercentFat value is invalid, float64 invalid value will be returned.
func (m *WeightScale) PercentFatScaled() float64 {
	if m.PercentFat == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.PercentFat, 100, 0)
}

// PercentHydrationScaled return PercentHydration in its scaled value [Scale: 100; Units: %].
//
// If PercentHydration value is invalid, float64 invalid value will be returned.
func (m *WeightScale) PercentHydrationScaled() float64 {
	if m.PercentHydration == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.PercentHydration, 100, 0)
}

// VisceralFatMassScaled return VisceralFatMass in its scaled value [Scale: 100; Units: kg].
//
// If VisceralFatMass value is invalid, float64 invalid value will be returned.
func (m *WeightScale) VisceralFatMassScaled() float64 {
	if m.VisceralFatMass == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.VisceralFatMass, 100, 0)
}

// BoneMassScaled return BoneMass in its scaled value [Scale: 100; Units: kg].
//
// If BoneMass value is invalid, float64 invalid value will be returned.
func (m *WeightScale) BoneMassScaled() float64 {
	if m.BoneMass == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.BoneMass, 100, 0)
}

// MuscleMassScaled return MuscleMass in its scaled value [Scale: 100; Units: kg].
//
// If MuscleMass value is invalid, float64 invalid value will be returned.
func (m *WeightScale) MuscleMassScaled() float64 {
	if m.MuscleMass == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.MuscleMass, 100, 0)
}

// BasalMetScaled return BasalMet in its scaled value [Scale: 4; Units: kcal/day].
//
// If BasalMet value is invalid, float64 invalid value will be returned.
func (m *WeightScale) BasalMetScaled() float64 {
	if m.BasalMet == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.BasalMet, 4, 0)
}

// ActiveMetScaled return ActiveMet in its scaled value [Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal].
//
// If ActiveMet value is invalid, float64 invalid value will be returned.
func (m *WeightScale) ActiveMetScaled() float64 {
	if m.ActiveMet == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.ActiveMet, 4, 0)
}

// BmiScaled return Bmi in its scaled value [Scale: 10; Units: kg/m^2].
//
// If Bmi value is invalid, float64 invalid value will be returned.
func (m *WeightScale) BmiScaled() float64 {
	if m.Bmi == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Bmi, 10, 0)
}

// SetTimestamp sets WeightScale value.
//
// Units: s
func (m *WeightScale) SetTimestamp(v time.Time) *WeightScale {
	m.Timestamp = v
	return m
}

// SetWeight sets WeightScale value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetWeight(v typedef.Weight) *WeightScale {
	m.Weight = v
	return m
}

// SetPercentFat sets WeightScale value.
//
// Scale: 100; Units: %
func (m *WeightScale) SetPercentFat(v uint16) *WeightScale {
	m.PercentFat = v
	return m
}

// SetPercentHydration sets WeightScale value.
//
// Scale: 100; Units: %
func (m *WeightScale) SetPercentHydration(v uint16) *WeightScale {
	m.PercentHydration = v
	return m
}

// SetVisceralFatMass sets WeightScale value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetVisceralFatMass(v uint16) *WeightScale {
	m.VisceralFatMass = v
	return m
}

// SetBoneMass sets WeightScale value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetBoneMass(v uint16) *WeightScale {
	m.BoneMass = v
	return m
}

// SetMuscleMass sets WeightScale value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetMuscleMass(v uint16) *WeightScale {
	m.MuscleMass = v
	return m
}

// SetBasalMet sets WeightScale value.
//
// Scale: 4; Units: kcal/day
func (m *WeightScale) SetBasalMet(v uint16) *WeightScale {
	m.BasalMet = v
	return m
}

// SetActiveMet sets WeightScale value.
//
// Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal
func (m *WeightScale) SetActiveMet(v uint16) *WeightScale {
	m.ActiveMet = v
	return m
}

// SetUserProfileIndex sets WeightScale value.
//
// Associates this weight scale message to a user. This corresponds to the index of the user profile message in the weight scale file.
func (m *WeightScale) SetUserProfileIndex(v typedef.MessageIndex) *WeightScale {
	m.UserProfileIndex = v
	return m
}

// SetBmi sets WeightScale value.
//
// Scale: 10; Units: kg/m^2
func (m *WeightScale) SetBmi(v uint16) *WeightScale {
	m.Bmi = v
	return m
}

// SetPhysiqueRating sets WeightScale value.
func (m *WeightScale) SetPhysiqueRating(v uint8) *WeightScale {
	m.PhysiqueRating = v
	return m
}

// SetMetabolicAge sets WeightScale value.
//
// Units: years
func (m *WeightScale) SetMetabolicAge(v uint8) *WeightScale {
	m.MetabolicAge = v
	return m
}

// SetVisceralFatRating sets WeightScale value.
func (m *WeightScale) SetVisceralFatRating(v uint8) *WeightScale {
	m.VisceralFatRating = v
	return m
}

// SetDeveloperFields WeightScale's DeveloperFields.
func (m *WeightScale) SetDeveloperFields(developerFields ...proto.DeveloperField) *WeightScale {
	m.DeveloperFields = developerFields
	return m
}
