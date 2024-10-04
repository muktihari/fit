// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
	"unsafe"
)

// MonitoringInfo is a MonitoringInfo message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type MonitoringInfo struct {
	Timestamp            time.Time              // Units: s
	LocalTimestamp       time.Time              // Units: s; Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
	ActivityType         []typedef.ActivityType // Array: [N]
	CyclesToDistance     []uint16               // Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type
	CyclesToCalories     []uint16               // Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type
	RestingMetabolicRate uint16                 // Units: kcal / day

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewMonitoringInfo creates new MonitoringInfo struct based on given mesg.
// If mesg is nil, it will return MonitoringInfo with all fields being set to its corresponding invalid value.
func NewMonitoringInfo(mesg *proto.Message) *MonitoringInfo {
	vals := [254]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &MonitoringInfo{
		Timestamp:      datetime.ToTime(vals[253].Uint32()),
		LocalTimestamp: datetime.ToTime(vals[0].Uint32()),
		ActivityType: func() []typedef.ActivityType {
			sliceValue := vals[1].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.ActivityType)(ptr), len(sliceValue))
		}(),
		CyclesToDistance:     vals[3].SliceUint16(),
		CyclesToCalories:     vals[4].SliceUint16(),
		RestingMetabolicRate: vals[5].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts MonitoringInfo into proto.Message. If options is nil, default options will be used.
func (m *MonitoringInfo) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumMonitoringInfo}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if !m.LocalTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(uint32(m.LocalTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.ActivityType != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint8(m.ActivityType)
		fields = append(fields, field)
	}
	if m.CyclesToDistance != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceUint16(m.CyclesToDistance)
		fields = append(fields, field)
	}
	if m.CyclesToCalories != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint16(m.CyclesToCalories)
		fields = append(fields, field)
	}
	if m.RestingMetabolicRate != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.RestingMetabolicRate)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *MonitoringInfo) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// LocalTimestampUint32 returns LocalTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *MonitoringInfo) LocalTimestampUint32() uint32 { return datetime.ToUint32(m.LocalTimestamp) }

// CyclesToDistanceScaled return CyclesToDistance in its scaled value.
// If CyclesToDistance value is invalid, nil will be returned.
//
// Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type
func (m *MonitoringInfo) CyclesToDistanceScaled() []float64 {
	if m.CyclesToDistance == nil {
		return nil
	}
	var vals = make([]float64, len(m.CyclesToDistance))
	for i := range m.CyclesToDistance {
		if m.CyclesToDistance[i] == basetype.Uint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.CyclesToDistance[i])/5000 - 0
	}
	return vals
}

// CyclesToCaloriesScaled return CyclesToCalories in its scaled value.
// If CyclesToCalories value is invalid, nil will be returned.
//
// Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type
func (m *MonitoringInfo) CyclesToCaloriesScaled() []float64 {
	if m.CyclesToCalories == nil {
		return nil
	}
	var vals = make([]float64, len(m.CyclesToCalories))
	for i := range m.CyclesToCalories {
		if m.CyclesToCalories[i] == basetype.Uint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.CyclesToCalories[i])/5000 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *MonitoringInfo) SetTimestamp(v time.Time) *MonitoringInfo {
	m.Timestamp = v
	return m
}

// SetLocalTimestamp sets LocalTimestamp value.
//
// Units: s; Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
func (m *MonitoringInfo) SetLocalTimestamp(v time.Time) *MonitoringInfo {
	m.LocalTimestamp = v
	return m
}

// SetActivityType sets ActivityType value.
//
// Array: [N]
func (m *MonitoringInfo) SetActivityType(v []typedef.ActivityType) *MonitoringInfo {
	m.ActivityType = v
	return m
}

// SetCyclesToDistance sets CyclesToDistance value.
//
// Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type
func (m *MonitoringInfo) SetCyclesToDistance(v []uint16) *MonitoringInfo {
	m.CyclesToDistance = v
	return m
}

// SetCyclesToDistanceScaled is similar to SetCyclesToDistance except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type
func (m *MonitoringInfo) SetCyclesToDistanceScaled(vs []float64) *MonitoringInfo {
	if vs == nil {
		m.CyclesToDistance = nil
		return m
	}
	m.CyclesToDistance = make([]uint16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 5000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
			m.CyclesToDistance[i] = uint16(basetype.Uint16Invalid)
			continue
		}
		m.CyclesToDistance[i] = uint16(unscaled)
	}
	return m
}

// SetCyclesToCalories sets CyclesToCalories value.
//
// Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type
func (m *MonitoringInfo) SetCyclesToCalories(v []uint16) *MonitoringInfo {
	m.CyclesToCalories = v
	return m
}

// SetCyclesToCaloriesScaled is similar to SetCyclesToCalories except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type
func (m *MonitoringInfo) SetCyclesToCaloriesScaled(vs []float64) *MonitoringInfo {
	if vs == nil {
		m.CyclesToCalories = nil
		return m
	}
	m.CyclesToCalories = make([]uint16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 5000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
			m.CyclesToCalories[i] = uint16(basetype.Uint16Invalid)
			continue
		}
		m.CyclesToCalories[i] = uint16(unscaled)
	}
	return m
}

// SetRestingMetabolicRate sets RestingMetabolicRate value.
//
// Units: kcal / day
func (m *MonitoringInfo) SetRestingMetabolicRate(v uint16) *MonitoringInfo {
	m.RestingMetabolicRate = v
	return m
}

// SetUnknownFields MonitoringInfo's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *MonitoringInfo) SetUnknownFields(unknownFields ...proto.Field) *MonitoringInfo {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields MonitoringInfo's DeveloperFields.
func (m *MonitoringInfo) SetDeveloperFields(developerFields ...proto.DeveloperField) *MonitoringInfo {
	m.DeveloperFields = developerFields
	return m
}
