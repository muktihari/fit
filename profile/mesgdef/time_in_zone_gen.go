// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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

// TimeInZone is a TimeInZone message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type TimeInZone struct {
	Timestamp                time.Time // Units: s
	TimeInHrZone             []uint32  // Array: [N]; Scale: 1000; Units: s
	TimeInSpeedZone          []uint32  // Array: [N]; Scale: 1000; Units: s
	TimeInCadenceZone        []uint32  // Array: [N]; Scale: 1000; Units: s
	TimeInPowerZone          []uint32  // Array: [N]; Scale: 1000; Units: s
	HrZoneHighBoundary       []uint8   // Array: [N]; Units: bpm
	SpeedZoneHighBoundary    []uint16  // Array: [N]; Scale: 1000; Units: m/s
	CadenceZoneHighBoundary  []uint8   // Array: [N]; Units: rpm
	PowerZoneHighBoundary    []uint16  // Array: [N]; Units: watts
	ReferenceMesg            typedef.MesgNum
	ReferenceIndex           typedef.MessageIndex
	FunctionalThresholdPower uint16
	HrCalcType               typedef.HrZoneCalc
	MaxHeartRate             uint8
	RestingHeartRate         uint8
	ThresholdHeartRate       uint8
	PwrCalcType              typedef.PwrZoneCalc

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTimeInZone creates new TimeInZone struct based on given mesg.
// If mesg is nil, it will return TimeInZone with all fields being set to its corresponding invalid value.
func NewTimeInZone(mesg *proto.Message) *TimeInZone {
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

	return &TimeInZone{
		Timestamp:                datetime.ToTime(vals[253].Uint32()),
		ReferenceMesg:            typedef.MesgNum(vals[0].Uint16()),
		ReferenceIndex:           typedef.MessageIndex(vals[1].Uint16()),
		TimeInHrZone:             vals[2].SliceUint32(),
		TimeInSpeedZone:          vals[3].SliceUint32(),
		TimeInCadenceZone:        vals[4].SliceUint32(),
		TimeInPowerZone:          vals[5].SliceUint32(),
		HrZoneHighBoundary:       vals[6].SliceUint8(),
		SpeedZoneHighBoundary:    vals[7].SliceUint16(),
		CadenceZoneHighBoundary:  vals[8].SliceUint8(),
		PowerZoneHighBoundary:    vals[9].SliceUint16(),
		HrCalcType:               typedef.HrZoneCalc(vals[10].Uint8()),
		MaxHeartRate:             vals[11].Uint8(),
		RestingHeartRate:         vals[12].Uint8(),
		ThresholdHeartRate:       vals[13].Uint8(),
		PwrCalcType:              typedef.PwrZoneCalc(vals[14].Uint8()),
		FunctionalThresholdPower: vals[15].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts TimeInZone into proto.Message. If options is nil, default options will be used.
func (m *TimeInZone) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumTimeInZone}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if uint16(m.ReferenceMesg) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.ReferenceMesg))
		fields = append(fields, field)
	}
	if uint16(m.ReferenceIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.ReferenceIndex))
		fields = append(fields, field)
	}
	if m.TimeInHrZone != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint32(m.TimeInHrZone)
		fields = append(fields, field)
	}
	if m.TimeInSpeedZone != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceUint32(m.TimeInSpeedZone)
		fields = append(fields, field)
	}
	if m.TimeInCadenceZone != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint32(m.TimeInCadenceZone)
		fields = append(fields, field)
	}
	if m.TimeInPowerZone != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceUint32(m.TimeInPowerZone)
		fields = append(fields, field)
	}
	if m.HrZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceUint8(m.HrZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.SpeedZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceUint16(m.SpeedZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.CadenceZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceUint8(m.CadenceZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.PowerZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.SliceUint16(m.PowerZoneHighBoundary)
		fields = append(fields, field)
	}
	if byte(m.HrCalcType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(byte(m.HrCalcType))
		fields = append(fields, field)
	}
	if m.MaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.MaxHeartRate)
		fields = append(fields, field)
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(m.RestingHeartRate)
		fields = append(fields, field)
	}
	if m.ThresholdHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint8(m.ThresholdHeartRate)
		fields = append(fields, field)
	}
	if byte(m.PwrCalcType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(byte(m.PwrCalcType))
		fields = append(fields, field)
	}
	if m.FunctionalThresholdPower != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint16(m.FunctionalThresholdPower)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TimeInZone) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// TimeInHrZoneScaled return TimeInHrZone in its scaled value.
// If TimeInHrZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInHrZoneScaled() []float64 {
	if m.TimeInHrZone == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.TimeInHrZone, 1000, 0)
}

// TimeInSpeedZoneScaled return TimeInSpeedZone in its scaled value.
// If TimeInSpeedZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInSpeedZoneScaled() []float64 {
	if m.TimeInSpeedZone == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.TimeInSpeedZone, 1000, 0)
}

// TimeInCadenceZoneScaled return TimeInCadenceZone in its scaled value.
// If TimeInCadenceZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInCadenceZoneScaled() []float64 {
	if m.TimeInCadenceZone == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.TimeInCadenceZone, 1000, 0)
}

// TimeInPowerZoneScaled return TimeInPowerZone in its scaled value.
// If TimeInPowerZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInPowerZoneScaled() []float64 {
	if m.TimeInPowerZone == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.TimeInPowerZone, 1000, 0)
}

// SpeedZoneHighBoundaryScaled return SpeedZoneHighBoundary in its scaled value.
// If SpeedZoneHighBoundary value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: m/s
func (m *TimeInZone) SpeedZoneHighBoundaryScaled() []float64 {
	if m.SpeedZoneHighBoundary == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.SpeedZoneHighBoundary, 1000, 0)
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *TimeInZone) SetTimestamp(v time.Time) *TimeInZone {
	m.Timestamp = v
	return m
}

// SetReferenceMesg sets ReferenceMesg value.
func (m *TimeInZone) SetReferenceMesg(v typedef.MesgNum) *TimeInZone {
	m.ReferenceMesg = v
	return m
}

// SetReferenceIndex sets ReferenceIndex value.
func (m *TimeInZone) SetReferenceIndex(v typedef.MessageIndex) *TimeInZone {
	m.ReferenceIndex = v
	return m
}

// SetTimeInHrZone sets TimeInHrZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInHrZone(v []uint32) *TimeInZone {
	m.TimeInHrZone = v
	return m
}

// SetTimeInHrZoneScaled is similar to SetTimeInHrZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInHrZoneScaled(vs []float64) *TimeInZone {
	m.TimeInHrZone = scaleoffset.DiscardSlice[uint32](vs, 1000, 0)
	return m
}

// SetTimeInSpeedZone sets TimeInSpeedZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInSpeedZone(v []uint32) *TimeInZone {
	m.TimeInSpeedZone = v
	return m
}

// SetTimeInSpeedZoneScaled is similar to SetTimeInSpeedZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInSpeedZoneScaled(vs []float64) *TimeInZone {
	m.TimeInSpeedZone = scaleoffset.DiscardSlice[uint32](vs, 1000, 0)
	return m
}

// SetTimeInCadenceZone sets TimeInCadenceZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInCadenceZone(v []uint32) *TimeInZone {
	m.TimeInCadenceZone = v
	return m
}

// SetTimeInCadenceZoneScaled is similar to SetTimeInCadenceZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInCadenceZoneScaled(vs []float64) *TimeInZone {
	m.TimeInCadenceZone = scaleoffset.DiscardSlice[uint32](vs, 1000, 0)
	return m
}

// SetTimeInPowerZone sets TimeInPowerZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInPowerZone(v []uint32) *TimeInZone {
	m.TimeInPowerZone = v
	return m
}

// SetTimeInPowerZoneScaled is similar to SetTimeInPowerZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInPowerZoneScaled(vs []float64) *TimeInZone {
	m.TimeInPowerZone = scaleoffset.DiscardSlice[uint32](vs, 1000, 0)
	return m
}

// SetHrZoneHighBoundary sets HrZoneHighBoundary value.
//
// Array: [N]; Units: bpm
func (m *TimeInZone) SetHrZoneHighBoundary(v []uint8) *TimeInZone {
	m.HrZoneHighBoundary = v
	return m
}

// SetSpeedZoneHighBoundary sets SpeedZoneHighBoundary value.
//
// Array: [N]; Scale: 1000; Units: m/s
func (m *TimeInZone) SetSpeedZoneHighBoundary(v []uint16) *TimeInZone {
	m.SpeedZoneHighBoundary = v
	return m
}

// SetSpeedZoneHighBoundaryScaled is similar to SetSpeedZoneHighBoundary except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: m/s
func (m *TimeInZone) SetSpeedZoneHighBoundaryScaled(vs []float64) *TimeInZone {
	m.SpeedZoneHighBoundary = scaleoffset.DiscardSlice[uint16](vs, 1000, 0)
	return m
}

// SetCadenceZoneHighBoundary sets CadenceZoneHighBoundary value.
//
// Array: [N]; Units: rpm
func (m *TimeInZone) SetCadenceZoneHighBoundary(v []uint8) *TimeInZone {
	m.CadenceZoneHighBoundary = v
	return m
}

// SetPowerZoneHighBoundary sets PowerZoneHighBoundary value.
//
// Array: [N]; Units: watts
func (m *TimeInZone) SetPowerZoneHighBoundary(v []uint16) *TimeInZone {
	m.PowerZoneHighBoundary = v
	return m
}

// SetHrCalcType sets HrCalcType value.
func (m *TimeInZone) SetHrCalcType(v typedef.HrZoneCalc) *TimeInZone {
	m.HrCalcType = v
	return m
}

// SetMaxHeartRate sets MaxHeartRate value.
func (m *TimeInZone) SetMaxHeartRate(v uint8) *TimeInZone {
	m.MaxHeartRate = v
	return m
}

// SetRestingHeartRate sets RestingHeartRate value.
func (m *TimeInZone) SetRestingHeartRate(v uint8) *TimeInZone {
	m.RestingHeartRate = v
	return m
}

// SetThresholdHeartRate sets ThresholdHeartRate value.
func (m *TimeInZone) SetThresholdHeartRate(v uint8) *TimeInZone {
	m.ThresholdHeartRate = v
	return m
}

// SetPwrCalcType sets PwrCalcType value.
func (m *TimeInZone) SetPwrCalcType(v typedef.PwrZoneCalc) *TimeInZone {
	m.PwrCalcType = v
	return m
}

// SetFunctionalThresholdPower sets FunctionalThresholdPower value.
func (m *TimeInZone) SetFunctionalThresholdPower(v uint16) *TimeInZone {
	m.FunctionalThresholdPower = v
	return m
}

// SetDeveloperFields TimeInZone's DeveloperFields.
func (m *TimeInZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *TimeInZone {
	m.DeveloperFields = developerFields
	return m
}
