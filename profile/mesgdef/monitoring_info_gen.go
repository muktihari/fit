// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// MonitoringInfo is a MonitoringInfo message.
type MonitoringInfo struct {
	Timestamp            time.Time              // Units: s
	LocalTimestamp       time.Time              // Units: s; Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
	ActivityType         []typedef.ActivityType // Array: [N]
	CyclesToDistance     []uint16               // Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type
	CyclesToCalories     []uint16               // Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type
	RestingMetabolicRate uint16                 // Units: kcal / day

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMonitoringInfo creates new MonitoringInfo struct based on given mesg.
// If mesg is nil, it will return MonitoringInfo with all fields being set to its corresponding invalid value.
func NewMonitoringInfo(mesg *proto.Message) *MonitoringInfo {
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

	return &MonitoringInfo{
		Timestamp:            datetime.ToTime(vals[253]),
		LocalTimestamp:       datetime.ToTime(vals[0]),
		ActivityType:         typeconv.ToSliceEnum[typedef.ActivityType](vals[1]),
		CyclesToDistance:     typeconv.ToSliceUint16[uint16](vals[3]),
		CyclesToCalories:     typeconv.ToSliceUint16[uint16](vals[4]),
		RestingMetabolicRate: typeconv.ToUint16[uint16](vals[5]),

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

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumMonitoringInfo)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.LocalTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = datetime.ToUint32(m.LocalTimestamp)
		fields = append(fields, field)
	}
	if typeconv.ToSliceEnum[byte](m.ActivityType) != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToSliceEnum[byte](m.ActivityType)
		fields = append(fields, field)
	}
	if m.CyclesToDistance != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.CyclesToDistance
		fields = append(fields, field)
	}
	if m.CyclesToCalories != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.CyclesToCalories
		fields = append(fields, field)
	}
	if m.RestingMetabolicRate != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.RestingMetabolicRate
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// CyclesToDistanceScaled return CyclesToDistance in its scaled value [Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type].
//
// If CyclesToDistance value is invalid, nil will be returned.
func (m *MonitoringInfo) CyclesToDistanceScaled() []float64 {
	if m.CyclesToDistance == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.CyclesToDistance, 5000, 0)
}

// CyclesToCaloriesScaled return CyclesToCalories in its scaled value [Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type].
//
// If CyclesToCalories value is invalid, nil will be returned.
func (m *MonitoringInfo) CyclesToCaloriesScaled() []float64 {
	if m.CyclesToCalories == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.CyclesToCalories, 5000, 0)
}

// SetTimestamp sets MonitoringInfo value.
//
// Units: s
func (m *MonitoringInfo) SetTimestamp(v time.Time) *MonitoringInfo {
	m.Timestamp = v
	return m
}

// SetLocalTimestamp sets MonitoringInfo value.
//
// Units: s; Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
func (m *MonitoringInfo) SetLocalTimestamp(v time.Time) *MonitoringInfo {
	m.LocalTimestamp = v
	return m
}

// SetActivityType sets MonitoringInfo value.
//
// Array: [N]
func (m *MonitoringInfo) SetActivityType(v []typedef.ActivityType) *MonitoringInfo {
	m.ActivityType = v
	return m
}

// SetCyclesToDistance sets MonitoringInfo value.
//
// Array: [N]; Scale: 5000; Units: m/cycle; Indexed by activity_type
func (m *MonitoringInfo) SetCyclesToDistance(v []uint16) *MonitoringInfo {
	m.CyclesToDistance = v
	return m
}

// SetCyclesToCalories sets MonitoringInfo value.
//
// Array: [N]; Scale: 5000; Units: kcal/cycle; Indexed by activity_type
func (m *MonitoringInfo) SetCyclesToCalories(v []uint16) *MonitoringInfo {
	m.CyclesToCalories = v
	return m
}

// SetRestingMetabolicRate sets MonitoringInfo value.
//
// Units: kcal / day
func (m *MonitoringInfo) SetRestingMetabolicRate(v uint16) *MonitoringInfo {
	m.RestingMetabolicRate = v
	return m
}

// SetDeveloperFields MonitoringInfo's DeveloperFields.
func (m *MonitoringInfo) SetDeveloperFields(developerFields ...proto.DeveloperField) *MonitoringInfo {
	m.DeveloperFields = developerFields
	return m
}
