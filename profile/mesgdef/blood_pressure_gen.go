// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// BloodPressure is a BloodPressure message.
type BloodPressure struct {
	Timestamp            time.Time // Units: s;
	SystolicPressure     uint16    // Units: mmHg;
	DiastolicPressure    uint16    // Units: mmHg;
	MeanArterialPressure uint16    // Units: mmHg;
	Map3SampleMean       uint16    // Units: mmHg;
	MapMorningValues     uint16    // Units: mmHg;
	MapEveningValues     uint16    // Units: mmHg;
	HeartRate            uint8     // Units: bpm;
	HeartRateType        typedef.HrType
	Status               typedef.BpStatus
	UserProfileIndex     typedef.MessageIndex // Associates this blood pressure message to a user. This corresponds to the index of the user profile message in the blood pressure file.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewBloodPressure creates new BloodPressure struct based on given mesg.
// If mesg is nil, it will return BloodPressure with all fields being set to its corresponding invalid value.
func NewBloodPressure(mesg *proto.Message) *BloodPressure {
	vals := [254]any{}

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

	return &BloodPressure{
		Timestamp:            datetime.ToTime(vals[253]),
		SystolicPressure:     typeconv.ToUint16[uint16](vals[0]),
		DiastolicPressure:    typeconv.ToUint16[uint16](vals[1]),
		MeanArterialPressure: typeconv.ToUint16[uint16](vals[2]),
		Map3SampleMean:       typeconv.ToUint16[uint16](vals[3]),
		MapMorningValues:     typeconv.ToUint16[uint16](vals[4]),
		MapEveningValues:     typeconv.ToUint16[uint16](vals[5]),
		HeartRate:            typeconv.ToUint8[uint8](vals[6]),
		HeartRateType:        typeconv.ToEnum[typedef.HrType](vals[7]),
		Status:               typeconv.ToEnum[typedef.BpStatus](vals[8]),
		UserProfileIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[9]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts BloodPressure into proto.Message.
func (m *BloodPressure) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumBloodPressure)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.SystolicPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.SystolicPressure
		fields = append(fields, field)
	}
	if m.DiastolicPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.DiastolicPressure
		fields = append(fields, field)
	}
	if m.MeanArterialPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.MeanArterialPressure
		fields = append(fields, field)
	}
	if m.Map3SampleMean != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Map3SampleMean
		fields = append(fields, field)
	}
	if m.MapMorningValues != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.MapMorningValues
		fields = append(fields, field)
	}
	if m.MapEveningValues != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.MapEveningValues
		fields = append(fields, field)
	}
	if m.HeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.HeartRate
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.HeartRateType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToEnum[byte](m.HeartRateType)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.Status) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = typeconv.ToEnum[byte](m.Status)
		fields = append(fields, field)
	}
	if typeconv.ToUint16[uint16](m.UserProfileIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToUint16[uint16](m.UserProfileIndex)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets BloodPressure value.
//
// Units: s;
func (m *BloodPressure) SetTimestamp(v time.Time) *BloodPressure {
	m.Timestamp = v
	return m
}

// SetSystolicPressure sets BloodPressure value.
//
// Units: mmHg;
func (m *BloodPressure) SetSystolicPressure(v uint16) *BloodPressure {
	m.SystolicPressure = v
	return m
}

// SetDiastolicPressure sets BloodPressure value.
//
// Units: mmHg;
func (m *BloodPressure) SetDiastolicPressure(v uint16) *BloodPressure {
	m.DiastolicPressure = v
	return m
}

// SetMeanArterialPressure sets BloodPressure value.
//
// Units: mmHg;
func (m *BloodPressure) SetMeanArterialPressure(v uint16) *BloodPressure {
	m.MeanArterialPressure = v
	return m
}

// SetMap3SampleMean sets BloodPressure value.
//
// Units: mmHg;
func (m *BloodPressure) SetMap3SampleMean(v uint16) *BloodPressure {
	m.Map3SampleMean = v
	return m
}

// SetMapMorningValues sets BloodPressure value.
//
// Units: mmHg;
func (m *BloodPressure) SetMapMorningValues(v uint16) *BloodPressure {
	m.MapMorningValues = v
	return m
}

// SetMapEveningValues sets BloodPressure value.
//
// Units: mmHg;
func (m *BloodPressure) SetMapEveningValues(v uint16) *BloodPressure {
	m.MapEveningValues = v
	return m
}

// SetHeartRate sets BloodPressure value.
//
// Units: bpm;
func (m *BloodPressure) SetHeartRate(v uint8) *BloodPressure {
	m.HeartRate = v
	return m
}

// SetHeartRateType sets BloodPressure value.
func (m *BloodPressure) SetHeartRateType(v typedef.HrType) *BloodPressure {
	m.HeartRateType = v
	return m
}

// SetStatus sets BloodPressure value.
func (m *BloodPressure) SetStatus(v typedef.BpStatus) *BloodPressure {
	m.Status = v
	return m
}

// SetUserProfileIndex sets BloodPressure value.
//
// Associates this blood pressure message to a user. This corresponds to the index of the user profile message in the blood pressure file.
func (m *BloodPressure) SetUserProfileIndex(v typedef.MessageIndex) *BloodPressure {
	m.UserProfileIndex = v
	return m
}

// SetDeveloperFields BloodPressure's DeveloperFields.
func (m *BloodPressure) SetDeveloperFields(developerFields ...proto.DeveloperField) *BloodPressure {
	m.DeveloperFields = developerFields
	return m
}
