// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// WeightScale is a WeightScale message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
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
			if mesg.Fields[i].Num > 253 {
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
		PhysiqueRating:    vals[8].Uint8(),
		ActiveMet:         vals[9].Uint16(),
		MetabolicAge:      vals[10].Uint8(),
		VisceralFatRating: vals[11].Uint8(),
		UserProfileIndex:  typedef.MessageIndex(vals[12].Uint16()),
		Bmi:               vals[13].Uint16(),

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

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumWeightScale}

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
	if m.PhysiqueRating != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(m.PhysiqueRating)
		fields = append(fields, field)
	}
	if m.ActiveMet != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(m.ActiveMet)
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

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *WeightScale) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// WeightScaled return Weight in its scaled value.
// If Weight value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: kg
func (m *WeightScale) WeightScaled() float64 {
	if uint16(m.Weight) == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Weight)/100 - 0
}

// PercentFatScaled return PercentFat in its scaled value.
// If PercentFat value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: %
func (m *WeightScale) PercentFatScaled() float64 {
	if m.PercentFat == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.PercentFat)/100 - 0
}

// PercentHydrationScaled return PercentHydration in its scaled value.
// If PercentHydration value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: %
func (m *WeightScale) PercentHydrationScaled() float64 {
	if m.PercentHydration == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.PercentHydration)/100 - 0
}

// VisceralFatMassScaled return VisceralFatMass in its scaled value.
// If VisceralFatMass value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: kg
func (m *WeightScale) VisceralFatMassScaled() float64 {
	if m.VisceralFatMass == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.VisceralFatMass)/100 - 0
}

// BoneMassScaled return BoneMass in its scaled value.
// If BoneMass value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: kg
func (m *WeightScale) BoneMassScaled() float64 {
	if m.BoneMass == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BoneMass)/100 - 0
}

// MuscleMassScaled return MuscleMass in its scaled value.
// If MuscleMass value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: kg
func (m *WeightScale) MuscleMassScaled() float64 {
	if m.MuscleMass == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.MuscleMass)/100 - 0
}

// BasalMetScaled return BasalMet in its scaled value.
// If BasalMet value is invalid, float64 invalid value will be returned.
//
// Scale: 4; Units: kcal/day
func (m *WeightScale) BasalMetScaled() float64 {
	if m.BasalMet == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BasalMet)/4 - 0
}

// ActiveMetScaled return ActiveMet in its scaled value.
// If ActiveMet value is invalid, float64 invalid value will be returned.
//
// Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal
func (m *WeightScale) ActiveMetScaled() float64 {
	if m.ActiveMet == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.ActiveMet)/4 - 0
}

// BmiScaled return Bmi in its scaled value.
// If Bmi value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: kg/m^2
func (m *WeightScale) BmiScaled() float64 {
	if m.Bmi == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Bmi)/10 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *WeightScale) SetTimestamp(v time.Time) *WeightScale {
	m.Timestamp = v
	return m
}

// SetWeight sets Weight value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetWeight(v typedef.Weight) *WeightScale {
	m.Weight = v
	return m
}

// SetWeightScaled is similar to SetWeight except it accepts a scaled value.
// This method automatically converts the given value to its typedef.Weight form, discarding any applied scale and offset.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetWeightScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Weight = typedef.Weight(basetype.Uint16Invalid)
		return m
	}
	m.Weight = typedef.Weight(unscaled)
	return m
}

// SetPercentFat sets PercentFat value.
//
// Scale: 100; Units: %
func (m *WeightScale) SetPercentFat(v uint16) *WeightScale {
	m.PercentFat = v
	return m
}

// SetPercentFatScaled is similar to SetPercentFat except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: %
func (m *WeightScale) SetPercentFatScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.PercentFat = uint16(basetype.Uint16Invalid)
		return m
	}
	m.PercentFat = uint16(unscaled)
	return m
}

// SetPercentHydration sets PercentHydration value.
//
// Scale: 100; Units: %
func (m *WeightScale) SetPercentHydration(v uint16) *WeightScale {
	m.PercentHydration = v
	return m
}

// SetPercentHydrationScaled is similar to SetPercentHydration except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: %
func (m *WeightScale) SetPercentHydrationScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.PercentHydration = uint16(basetype.Uint16Invalid)
		return m
	}
	m.PercentHydration = uint16(unscaled)
	return m
}

// SetVisceralFatMass sets VisceralFatMass value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetVisceralFatMass(v uint16) *WeightScale {
	m.VisceralFatMass = v
	return m
}

// SetVisceralFatMassScaled is similar to SetVisceralFatMass except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetVisceralFatMassScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.VisceralFatMass = uint16(basetype.Uint16Invalid)
		return m
	}
	m.VisceralFatMass = uint16(unscaled)
	return m
}

// SetBoneMass sets BoneMass value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetBoneMass(v uint16) *WeightScale {
	m.BoneMass = v
	return m
}

// SetBoneMassScaled is similar to SetBoneMass except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetBoneMassScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BoneMass = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BoneMass = uint16(unscaled)
	return m
}

// SetMuscleMass sets MuscleMass value.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetMuscleMass(v uint16) *WeightScale {
	m.MuscleMass = v
	return m
}

// SetMuscleMassScaled is similar to SetMuscleMass except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: kg
func (m *WeightScale) SetMuscleMassScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.MuscleMass = uint16(basetype.Uint16Invalid)
		return m
	}
	m.MuscleMass = uint16(unscaled)
	return m
}

// SetBasalMet sets BasalMet value.
//
// Scale: 4; Units: kcal/day
func (m *WeightScale) SetBasalMet(v uint16) *WeightScale {
	m.BasalMet = v
	return m
}

// SetBasalMetScaled is similar to SetBasalMet except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 4; Units: kcal/day
func (m *WeightScale) SetBasalMetScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 4
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BasalMet = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BasalMet = uint16(unscaled)
	return m
}

// SetPhysiqueRating sets PhysiqueRating value.
func (m *WeightScale) SetPhysiqueRating(v uint8) *WeightScale {
	m.PhysiqueRating = v
	return m
}

// SetActiveMet sets ActiveMet value.
//
// Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal
func (m *WeightScale) SetActiveMet(v uint16) *WeightScale {
	m.ActiveMet = v
	return m
}

// SetActiveMetScaled is similar to SetActiveMet except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 4; Units: kcal/day; ~4kJ per kcal, 0.25 allows max 16384 kcal
func (m *WeightScale) SetActiveMetScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 4
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.ActiveMet = uint16(basetype.Uint16Invalid)
		return m
	}
	m.ActiveMet = uint16(unscaled)
	return m
}

// SetMetabolicAge sets MetabolicAge value.
//
// Units: years
func (m *WeightScale) SetMetabolicAge(v uint8) *WeightScale {
	m.MetabolicAge = v
	return m
}

// SetVisceralFatRating sets VisceralFatRating value.
func (m *WeightScale) SetVisceralFatRating(v uint8) *WeightScale {
	m.VisceralFatRating = v
	return m
}

// SetUserProfileIndex sets UserProfileIndex value.
//
// Associates this weight scale message to a user. This corresponds to the index of the user profile message in the weight scale file.
func (m *WeightScale) SetUserProfileIndex(v typedef.MessageIndex) *WeightScale {
	m.UserProfileIndex = v
	return m
}

// SetBmi sets Bmi value.
//
// Scale: 10; Units: kg/m^2
func (m *WeightScale) SetBmi(v uint16) *WeightScale {
	m.Bmi = v
	return m
}

// SetBmiScaled is similar to SetBmi except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: kg/m^2
func (m *WeightScale) SetBmiScaled(v float64) *WeightScale {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Bmi = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Bmi = uint16(unscaled)
	return m
}

// SetDeveloperFields WeightScale's DeveloperFields.
func (m *WeightScale) SetDeveloperFields(developerFields ...proto.DeveloperField) *WeightScale {
	m.DeveloperFields = developerFields
	return m
}
