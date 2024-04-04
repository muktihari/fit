// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// ZonesTarget is a ZonesTarget message.
type ZonesTarget struct {
	FunctionalThresholdPower uint16
	MaxHeartRate             uint8
	ThresholdHeartRate       uint8
	HrCalcType               typedef.HrZoneCalc
	PwrCalcType              typedef.PwrZoneCalc

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewZonesTarget creates new ZonesTarget struct based on given mesg.
// If mesg is nil, it will return ZonesTarget with all fields being set to its corresponding invalid value.
func NewZonesTarget(mesg *proto.Message) *ZonesTarget {
	vals := [8]proto.Value{}

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

	return &ZonesTarget{
		FunctionalThresholdPower: vals[3].Uint16(),
		MaxHeartRate:             vals[1].Uint8(),
		ThresholdHeartRate:       vals[2].Uint8(),
		HrCalcType:               typedef.HrZoneCalc(vals[5].Uint8()),
		PwrCalcType:              typedef.PwrZoneCalc(vals[7].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ZonesTarget into proto.Message. If options is nil, default options will be used.
func (m *ZonesTarget) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumZonesTarget)

	if m.FunctionalThresholdPower != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.FunctionalThresholdPower)
		fields = append(fields, field)
	}
	if m.MaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.MaxHeartRate)
		fields = append(fields, field)
	}
	if m.ThresholdHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.ThresholdHeartRate)
		fields = append(fields, field)
	}
	if byte(m.HrCalcType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(byte(m.HrCalcType))
		fields = append(fields, field)
	}
	if byte(m.PwrCalcType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.PwrCalcType))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetFunctionalThresholdPower sets ZonesTarget value.
func (m *ZonesTarget) SetFunctionalThresholdPower(v uint16) *ZonesTarget {
	m.FunctionalThresholdPower = v
	return m
}

// SetMaxHeartRate sets ZonesTarget value.
func (m *ZonesTarget) SetMaxHeartRate(v uint8) *ZonesTarget {
	m.MaxHeartRate = v
	return m
}

// SetThresholdHeartRate sets ZonesTarget value.
func (m *ZonesTarget) SetThresholdHeartRate(v uint8) *ZonesTarget {
	m.ThresholdHeartRate = v
	return m
}

// SetHrCalcType sets ZonesTarget value.
func (m *ZonesTarget) SetHrCalcType(v typedef.HrZoneCalc) *ZonesTarget {
	m.HrCalcType = v
	return m
}

// SetPwrCalcType sets ZonesTarget value.
func (m *ZonesTarget) SetPwrCalcType(v typedef.PwrZoneCalc) *ZonesTarget {
	m.PwrCalcType = v
	return m
}

// SetDeveloperFields ZonesTarget's DeveloperFields.
func (m *ZonesTarget) SetDeveloperFields(developerFields ...proto.DeveloperField) *ZonesTarget {
	m.DeveloperFields = developerFields
	return m
}
