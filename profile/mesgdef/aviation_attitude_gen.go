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
	"unsafe"
)

// AviationAttitude is a AviationAttitude message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type AviationAttitude struct {
	Timestamp             time.Time                  // Units: s; Timestamp message was output
	SystemTime            []uint32                   // Array: [N]; Units: ms; System time associated with sample expressed in ms.
	Pitch                 []int16                    // Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2
	Roll                  []int16                    // Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI
	AccelLateral          []int16                    // Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	AccelNormal           []int16                    // Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	TurnRate              []int16                    // Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
	Stage                 []typedef.AttitudeStage    // Array: [N]
	AttitudeStageComplete []uint8                    // Array: [N]; Units: %; The percent complete of the current attitude stage. Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it. Range - 100
	Track                 []uint16                   // Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi
	Validity              []typedef.AttitudeValidity // Array: [N]
	TimestampMs           uint16                     // Units: ms; Fractional part of timestamp, added to timestamp

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAviationAttitude creates new AviationAttitude struct based on given mesg.
// If mesg is nil, it will return AviationAttitude with all fields being set to its corresponding invalid value.
func NewAviationAttitude(mesg *proto.Message) *AviationAttitude {
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

	return &AviationAttitude{
		Timestamp:    datetime.ToTime(vals[253].Uint32()),
		TimestampMs:  vals[0].Uint16(),
		SystemTime:   vals[1].SliceUint32(),
		Pitch:        vals[2].SliceInt16(),
		Roll:         vals[3].SliceInt16(),
		AccelLateral: vals[4].SliceInt16(),
		AccelNormal:  vals[5].SliceInt16(),
		TurnRate:     vals[6].SliceInt16(),
		Stage: func() []typedef.AttitudeStage {
			sliceValue := vals[7].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.AttitudeStage)(ptr), len(sliceValue))
		}(),
		AttitudeStageComplete: vals[8].SliceUint8(),
		Track:                 vals[9].SliceUint16(),
		Validity: func() []typedef.AttitudeValidity {
			sliceValue := vals[10].SliceUint16()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.AttitudeValidity)(ptr), len(sliceValue))
		}(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts AviationAttitude into proto.Message. If options is nil, default options will be used.
func (m *AviationAttitude) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumAviationAttitude}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.SystemTime != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint32(m.SystemTime)
		fields = append(fields, field)
	}
	if m.Pitch != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceInt16(m.Pitch)
		fields = append(fields, field)
	}
	if m.Roll != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceInt16(m.Roll)
		fields = append(fields, field)
	}
	if m.AccelLateral != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceInt16(m.AccelLateral)
		fields = append(fields, field)
	}
	if m.AccelNormal != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceInt16(m.AccelNormal)
		fields = append(fields, field)
	}
	if m.TurnRate != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceInt16(m.TurnRate)
		fields = append(fields, field)
	}
	if m.Stage != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceUint8(m.Stage)
		fields = append(fields, field)
	}
	if m.AttitudeStageComplete != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceUint8(m.AttitudeStageComplete)
		fields = append(fields, field)
	}
	if m.Track != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.SliceUint16(m.Track)
		fields = append(fields, field)
	}
	if m.Validity != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.SliceUint16(m.Validity)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *AviationAttitude) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// PitchScaled return Pitch in its scaled value.
// If Pitch value is invalid, nil will be returned.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2
func (m *AviationAttitude) PitchScaled() []float64 {
	if m.Pitch == nil {
		return nil
	}
	var vals = make([]float64, len(m.Pitch))
	for i := range m.Pitch {
		if m.Pitch[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.Pitch[i])/10430.38 - 0
	}
	return vals
}

// RollScaled return Roll in its scaled value.
// If Roll value is invalid, nil will be returned.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI
func (m *AviationAttitude) RollScaled() []float64 {
	if m.Roll == nil {
		return nil
	}
	var vals = make([]float64, len(m.Roll))
	for i := range m.Roll {
		if m.Roll[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.Roll[i])/10430.38 - 0
	}
	return vals
}

// AccelLateralScaled return AccelLateral in its scaled value.
// If AccelLateral value is invalid, nil will be returned.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) AccelLateralScaled() []float64 {
	if m.AccelLateral == nil {
		return nil
	}
	var vals = make([]float64, len(m.AccelLateral))
	for i := range m.AccelLateral {
		if m.AccelLateral[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.AccelLateral[i])/100 - 0
	}
	return vals
}

// AccelNormalScaled return AccelNormal in its scaled value.
// If AccelNormal value is invalid, nil will be returned.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) AccelNormalScaled() []float64 {
	if m.AccelNormal == nil {
		return nil
	}
	var vals = make([]float64, len(m.AccelNormal))
	for i := range m.AccelNormal {
		if m.AccelNormal[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.AccelNormal[i])/100 - 0
	}
	return vals
}

// TurnRateScaled return TurnRate in its scaled value.
// If TurnRate value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
func (m *AviationAttitude) TurnRateScaled() []float64 {
	if m.TurnRate == nil {
		return nil
	}
	var vals = make([]float64, len(m.TurnRate))
	for i := range m.TurnRate {
		if m.TurnRate[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.TurnRate[i])/1024 - 0
	}
	return vals
}

// TrackScaled return Track in its scaled value.
// If Track value is invalid, nil will be returned.
//
// Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi
func (m *AviationAttitude) TrackScaled() []float64 {
	if m.Track == nil {
		return nil
	}
	var vals = make([]float64, len(m.Track))
	for i := range m.Track {
		if m.Track[i] == basetype.Uint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.Track[i])/10430.38 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s; Timestamp message was output
func (m *AviationAttitude) SetTimestamp(v time.Time) *AviationAttitude {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Fractional part of timestamp, added to timestamp
func (m *AviationAttitude) SetTimestampMs(v uint16) *AviationAttitude {
	m.TimestampMs = v
	return m
}

// SetSystemTime sets SystemTime value.
//
// Array: [N]; Units: ms; System time associated with sample expressed in ms.
func (m *AviationAttitude) SetSystemTime(v []uint32) *AviationAttitude {
	m.SystemTime = v
	return m
}

// SetPitch sets Pitch value.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2
func (m *AviationAttitude) SetPitch(v []int16) *AviationAttitude {
	m.Pitch = v
	return m
}

// SetPitchScaled is similar to SetPitch except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2
func (m *AviationAttitude) SetPitchScaled(vs []float64) *AviationAttitude {
	if vs == nil {
		m.Pitch = nil
		return m
	}
	m.Pitch = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 10430.38
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.Pitch[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.Pitch[i] = int16(unscaled)
	}
	return m
}

// SetRoll sets Roll value.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI
func (m *AviationAttitude) SetRoll(v []int16) *AviationAttitude {
	m.Roll = v
	return m
}

// SetRollScaled is similar to SetRoll except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI
func (m *AviationAttitude) SetRollScaled(vs []float64) *AviationAttitude {
	if vs == nil {
		m.Roll = nil
		return m
	}
	m.Roll = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 10430.38
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.Roll[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.Roll[i] = int16(unscaled)
	}
	return m
}

// SetAccelLateral sets AccelLateral value.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) SetAccelLateral(v []int16) *AviationAttitude {
	m.AccelLateral = v
	return m
}

// SetAccelLateralScaled is similar to SetAccelLateral except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) SetAccelLateralScaled(vs []float64) *AviationAttitude {
	if vs == nil {
		m.AccelLateral = nil
		return m
	}
	m.AccelLateral = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 100
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.AccelLateral[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.AccelLateral[i] = int16(unscaled)
	}
	return m
}

// SetAccelNormal sets AccelNormal value.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) SetAccelNormal(v []int16) *AviationAttitude {
	m.AccelNormal = v
	return m
}

// SetAccelNormalScaled is similar to SetAccelNormal except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) SetAccelNormalScaled(vs []float64) *AviationAttitude {
	if vs == nil {
		m.AccelNormal = nil
		return m
	}
	m.AccelNormal = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 100
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.AccelNormal[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.AccelNormal[i] = int16(unscaled)
	}
	return m
}

// SetTurnRate sets TurnRate value.
//
// Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
func (m *AviationAttitude) SetTurnRate(v []int16) *AviationAttitude {
	m.TurnRate = v
	return m
}

// SetTurnRateScaled is similar to SetTurnRate except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
func (m *AviationAttitude) SetTurnRateScaled(vs []float64) *AviationAttitude {
	if vs == nil {
		m.TurnRate = nil
		return m
	}
	m.TurnRate = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1024
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.TurnRate[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.TurnRate[i] = int16(unscaled)
	}
	return m
}

// SetStage sets Stage value.
//
// Array: [N]
func (m *AviationAttitude) SetStage(v []typedef.AttitudeStage) *AviationAttitude {
	m.Stage = v
	return m
}

// SetAttitudeStageComplete sets AttitudeStageComplete value.
//
// Array: [N]; Units: %; The percent complete of the current attitude stage. Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it. Range - 100
func (m *AviationAttitude) SetAttitudeStageComplete(v []uint8) *AviationAttitude {
	m.AttitudeStageComplete = v
	return m
}

// SetTrack sets Track value.
//
// Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi
func (m *AviationAttitude) SetTrack(v []uint16) *AviationAttitude {
	m.Track = v
	return m
}

// SetTrackScaled is similar to SetTrack except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi
func (m *AviationAttitude) SetTrackScaled(vs []float64) *AviationAttitude {
	if vs == nil {
		m.Track = nil
		return m
	}
	m.Track = make([]uint16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 10430.38
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
			m.Track[i] = uint16(basetype.Uint16Invalid)
			continue
		}
		m.Track[i] = uint16(unscaled)
	}
	return m
}

// SetValidity sets Validity value.
//
// Array: [N]
func (m *AviationAttitude) SetValidity(v []typedef.AttitudeValidity) *AviationAttitude {
	m.Validity = v
	return m
}

// SetDeveloperFields AviationAttitude's DeveloperFields.
func (m *AviationAttitude) SetDeveloperFields(developerFields ...proto.DeveloperField) *AviationAttitude {
	m.DeveloperFields = developerFields
	return m
}
