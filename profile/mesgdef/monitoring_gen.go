// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Monitoring is a Monitoring message.
type Monitoring struct {
	Timestamp                    typedef.DateTime    // Units: s; Must align to logging interval, for example, time must be 00:00:00 for daily log.
	DeviceIndex                  typedef.DeviceIndex // Associates this data to device_info message. Not required for file with single device (sensor).
	Calories                     uint16              // Units: kcal; Accumulated total calories. Maintained by MonitoringReader for each activity_type. See SDK documentation
	Distance                     uint32              // Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
	Cycles                       uint32              // Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
	ActiveTime                   uint32              // Scale: 1000; Units: s;
	ActivityType                 typedef.ActivityType
	ActivitySubtype              typedef.ActivitySubtype
	ActivityLevel                typedef.ActivityLevel
	Distance16                   uint16                // Units: 100 * m;
	Cycles16                     uint16                // Units: 2 * cycles (steps);
	ActiveTime16                 uint16                // Units: s;
	LocalTimestamp               typedef.LocalDateTime // Must align to logging interval, for example, time must be 00:00:00 for daily log.
	Temperature                  int16                 // Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
	TemperatureMin               int16                 // Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
	TemperatureMax               int16                 // Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
	ActivityTime                 []uint16              // Array: [8]; Units: minutes; Indexed using minute_activity_level enum
	ActiveCalories               uint16                // Units: kcal;
	CurrentActivityTypeIntensity byte                  // Indicates single type / intensity for duration since last monitoring message.
	TimestampMin8                uint8                 // Units: min;
	Timestamp16                  uint16                // Units: s;
	HeartRate                    uint8                 // Units: bpm;
	Intensity                    uint8                 // Scale: 10;
	DurationMin                  uint16                // Units: min;
	Duration                     uint32                // Units: s;
	Ascent                       uint32                // Scale: 1000; Units: m;
	Descent                      uint32                // Scale: 1000; Units: m;
	ModerateActivityMinutes      uint16                // Units: minutes;
	VigorousActivityMinutes      uint16                // Units: minutes;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMonitoring creates new Monitoring struct based on given mesg. If mesg is nil or mesg.Num is not equal to Monitoring mesg number, it will return nil.
func NewMonitoring(mesg proto.Message) *Monitoring {
	if mesg.Num != typedef.MesgNumMonitoring {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		253: nil, /* Timestamp */
		0:   nil, /* DeviceIndex */
		1:   nil, /* Calories */
		2:   nil, /* Distance */
		3:   nil, /* Cycles */
		4:   nil, /* ActiveTime */
		5:   nil, /* ActivityType */
		6:   nil, /* ActivitySubtype */
		7:   nil, /* ActivityLevel */
		8:   nil, /* Distance16 */
		9:   nil, /* Cycles16 */
		10:  nil, /* ActiveTime16 */
		11:  nil, /* LocalTimestamp */
		12:  nil, /* Temperature */
		14:  nil, /* TemperatureMin */
		15:  nil, /* TemperatureMax */
		16:  nil, /* ActivityTime */
		19:  nil, /* ActiveCalories */
		24:  nil, /* CurrentActivityTypeIntensity */
		25:  nil, /* TimestampMin8 */
		26:  nil, /* Timestamp16 */
		27:  nil, /* HeartRate */
		28:  nil, /* Intensity */
		29:  nil, /* DurationMin */
		30:  nil, /* Duration */
		31:  nil, /* Ascent */
		32:  nil, /* Descent */
		33:  nil, /* ModerateActivityMinutes */
		34:  nil, /* VigorousActivityMinutes */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &Monitoring{
		Timestamp:                    typeconv.ToUint32[typedef.DateTime](vals[253]),
		DeviceIndex:                  typeconv.ToUint8[typedef.DeviceIndex](vals[0]),
		Calories:                     typeconv.ToUint16[uint16](vals[1]),
		Distance:                     typeconv.ToUint32[uint32](vals[2]),
		Cycles:                       typeconv.ToUint32[uint32](vals[3]),
		ActiveTime:                   typeconv.ToUint32[uint32](vals[4]),
		ActivityType:                 typeconv.ToEnum[typedef.ActivityType](vals[5]),
		ActivitySubtype:              typeconv.ToEnum[typedef.ActivitySubtype](vals[6]),
		ActivityLevel:                typeconv.ToEnum[typedef.ActivityLevel](vals[7]),
		Distance16:                   typeconv.ToUint16[uint16](vals[8]),
		Cycles16:                     typeconv.ToUint16[uint16](vals[9]),
		ActiveTime16:                 typeconv.ToUint16[uint16](vals[10]),
		LocalTimestamp:               typeconv.ToUint32[typedef.LocalDateTime](vals[11]),
		Temperature:                  typeconv.ToSint16[int16](vals[12]),
		TemperatureMin:               typeconv.ToSint16[int16](vals[14]),
		TemperatureMax:               typeconv.ToSint16[int16](vals[15]),
		ActivityTime:                 typeconv.ToSliceUint16[uint16](vals[16]),
		ActiveCalories:               typeconv.ToUint16[uint16](vals[19]),
		CurrentActivityTypeIntensity: typeconv.ToByte[byte](vals[24]),
		TimestampMin8:                typeconv.ToUint8[uint8](vals[25]),
		Timestamp16:                  typeconv.ToUint16[uint16](vals[26]),
		HeartRate:                    typeconv.ToUint8[uint8](vals[27]),
		Intensity:                    typeconv.ToUint8[uint8](vals[28]),
		DurationMin:                  typeconv.ToUint16[uint16](vals[29]),
		Duration:                     typeconv.ToUint32[uint32](vals[30]),
		Ascent:                       typeconv.ToUint32[uint32](vals[31]),
		Descent:                      typeconv.ToUint32[uint32](vals[32]),
		ModerateActivityMinutes:      typeconv.ToUint16[uint16](vals[33]),
		VigorousActivityMinutes:      typeconv.ToUint16[uint16](vals[34]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Monitoring mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumMonitoring)
func (m Monitoring) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumMonitoring {
		return
	}

	vals := [...]any{
		253: m.Timestamp,
		0:   m.DeviceIndex,
		1:   m.Calories,
		2:   m.Distance,
		3:   m.Cycles,
		4:   m.ActiveTime,
		5:   m.ActivityType,
		6:   m.ActivitySubtype,
		7:   m.ActivityLevel,
		8:   m.Distance16,
		9:   m.Cycles16,
		10:  m.ActiveTime16,
		11:  m.LocalTimestamp,
		12:  m.Temperature,
		14:  m.TemperatureMin,
		15:  m.TemperatureMax,
		16:  m.ActivityTime,
		19:  m.ActiveCalories,
		24:  m.CurrentActivityTypeIntensity,
		25:  m.TimestampMin8,
		26:  m.Timestamp16,
		27:  m.HeartRate,
		28:  m.Intensity,
		29:  m.DurationMin,
		30:  m.Duration,
		31:  m.Ascent,
		32:  m.Descent,
		33:  m.ModerateActivityMinutes,
		34:  m.VigorousActivityMinutes,
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		field.Value = vals[field.Num]
	}

	mesg.DeveloperFields = m.DeveloperFields
}
