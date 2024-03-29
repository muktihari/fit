// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// Monitoring is a Monitoring message.
type Monitoring struct {
	Timestamp                    time.Time           // Units: s; Must align to logging interval, for example, time must be 00:00:00 for daily log.
	LocalTimestamp               time.Time           // Must align to logging interval, for example, time must be 00:00:00 for daily log.
	ActivityTime                 []uint16            // Array: [8]; Units: minutes; Indexed using minute_activity_level enum
	Distance                     uint32              // Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
	Cycles                       uint32              // Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
	ActiveTime                   uint32              // Scale: 1000; Units: s
	Duration                     uint32              // Units: s
	Ascent                       uint32              // Scale: 1000; Units: m
	Descent                      uint32              // Scale: 1000; Units: m
	Calories                     uint16              // Units: kcal; Accumulated total calories. Maintained by MonitoringReader for each activity_type. See SDK documentation
	Distance16                   uint16              // Units: 100 * m
	Cycles16                     uint16              // Units: 2 * cycles (steps)
	ActiveTime16                 uint16              // Units: s
	Temperature                  int16               // Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
	TemperatureMin               int16               // Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
	TemperatureMax               int16               // Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
	ActiveCalories               uint16              // Units: kcal
	Timestamp16                  uint16              // Units: s
	DurationMin                  uint16              // Units: min
	ModerateActivityMinutes      uint16              // Units: minutes
	VigorousActivityMinutes      uint16              // Units: minutes
	DeviceIndex                  typedef.DeviceIndex // Associates this data to device_info message. Not required for file with single device (sensor).
	ActivityType                 typedef.ActivityType
	ActivitySubtype              typedef.ActivitySubtype
	ActivityLevel                typedef.ActivityLevel
	CurrentActivityTypeIntensity byte  // Indicates single type / intensity for duration since last monitoring message.
	TimestampMin8                uint8 // Units: min
	HeartRate                    uint8 // Units: bpm
	Intensity                    uint8 // Scale: 10

	IsExpandedFields [29]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMonitoring creates new Monitoring struct based on given mesg.
// If mesg is nil, it will return Monitoring with all fields being set to its corresponding invalid value.
func NewMonitoring(mesg *proto.Message) *Monitoring {
	vals := [254]any{}
	isExpandedFields := [29]bool{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < byte(len(isExpandedFields)) {
				isExpandedFields[mesg.Fields[i].Num] = mesg.Fields[i].IsExpandedField
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Monitoring{
		Timestamp:                    datetime.ToTime(vals[253]),
		LocalTimestamp:               datetime.ToTime(vals[11]),
		ActivityTime:                 typeconv.ToSliceUint16[uint16](vals[16]),
		Distance:                     typeconv.ToUint32[uint32](vals[2]),
		Cycles:                       typeconv.ToUint32[uint32](vals[3]),
		ActiveTime:                   typeconv.ToUint32[uint32](vals[4]),
		Duration:                     typeconv.ToUint32[uint32](vals[30]),
		Ascent:                       typeconv.ToUint32[uint32](vals[31]),
		Descent:                      typeconv.ToUint32[uint32](vals[32]),
		Calories:                     typeconv.ToUint16[uint16](vals[1]),
		Distance16:                   typeconv.ToUint16[uint16](vals[8]),
		Cycles16:                     typeconv.ToUint16[uint16](vals[9]),
		ActiveTime16:                 typeconv.ToUint16[uint16](vals[10]),
		Temperature:                  typeconv.ToSint16[int16](vals[12]),
		TemperatureMin:               typeconv.ToSint16[int16](vals[14]),
		TemperatureMax:               typeconv.ToSint16[int16](vals[15]),
		ActiveCalories:               typeconv.ToUint16[uint16](vals[19]),
		Timestamp16:                  typeconv.ToUint16[uint16](vals[26]),
		DurationMin:                  typeconv.ToUint16[uint16](vals[29]),
		ModerateActivityMinutes:      typeconv.ToUint16[uint16](vals[33]),
		VigorousActivityMinutes:      typeconv.ToUint16[uint16](vals[34]),
		DeviceIndex:                  typeconv.ToUint8[typedef.DeviceIndex](vals[0]),
		ActivityType:                 typeconv.ToEnum[typedef.ActivityType](vals[5]),
		ActivitySubtype:              typeconv.ToEnum[typedef.ActivitySubtype](vals[6]),
		ActivityLevel:                typeconv.ToEnum[typedef.ActivityLevel](vals[7]),
		CurrentActivityTypeIntensity: typeconv.ToByte[byte](vals[24]),
		TimestampMin8:                typeconv.ToUint8[uint8](vals[25]),
		HeartRate:                    typeconv.ToUint8[uint8](vals[27]),
		Intensity:                    typeconv.ToUint8[uint8](vals[28]),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Monitoring into proto.Message.
func (m *Monitoring) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumMonitoring)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.LocalTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = datetime.ToUint32(m.LocalTimestamp)
		fields = append(fields, field)
	}
	if m.ActivityTime != nil {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = m.ActivityTime
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Distance
		fields = append(fields, field)
	}
	if m.Cycles != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Cycles
		fields = append(fields, field)
	}
	if m.ActiveTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.ActiveTime
		fields = append(fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 30)
		field.Value = m.Duration
		fields = append(fields, field)
	}
	if m.Ascent != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 31)
		field.Value = m.Ascent
		fields = append(fields, field)
	}
	if m.Descent != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = m.Descent
		fields = append(fields, field)
	}
	if m.Calories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Calories
		fields = append(fields, field)
	}
	if m.Distance16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.Distance16
		fields = append(fields, field)
	}
	if m.Cycles16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.Cycles16
		fields = append(fields, field)
	}
	if m.ActiveTime16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.ActiveTime16
		fields = append(fields, field)
	}
	if m.Temperature != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = m.Temperature
		fields = append(fields, field)
	}
	if m.TemperatureMin != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.TemperatureMin
		fields = append(fields, field)
	}
	if m.TemperatureMax != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = m.TemperatureMax
		fields = append(fields, field)
	}
	if m.ActiveCalories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = m.ActiveCalories
		fields = append(fields, field)
	}
	if m.Timestamp16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 26)
		field.Value = m.Timestamp16
		fields = append(fields, field)
	}
	if m.DurationMin != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 29)
		field.Value = m.DurationMin
		fields = append(fields, field)
	}
	if m.ModerateActivityMinutes != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 33)
		field.Value = m.ModerateActivityMinutes
		fields = append(fields, field)
	}
	if m.VigorousActivityMinutes != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 34)
		field.Value = m.VigorousActivityMinutes
		fields = append(fields, field)
	}
	if uint8(m.DeviceIndex) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = uint8(m.DeviceIndex)
		fields = append(fields, field)
	}
	if byte(m.ActivityType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = byte(m.ActivityType)
		field.IsExpandedField = m.IsExpandedFields[5]
		fields = append(fields, field)
	}
	if byte(m.ActivitySubtype) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = byte(m.ActivitySubtype)
		fields = append(fields, field)
	}
	if byte(m.ActivityLevel) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = byte(m.ActivityLevel)
		fields = append(fields, field)
	}
	if m.CurrentActivityTypeIntensity != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = m.CurrentActivityTypeIntensity
		fields = append(fields, field)
	}
	if m.TimestampMin8 != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = m.TimestampMin8
		fields = append(fields, field)
	}
	if m.HeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = m.HeartRate
		fields = append(fields, field)
	}
	if m.Intensity != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 28)
		field.Value = m.Intensity
		field.IsExpandedField = m.IsExpandedFields[28]
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// DistanceScaled return Distance in its scaled value [Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.].
//
// If Distance value is invalid, float64 invalid value will be returned.
func (m *Monitoring) DistanceScaled() float64 {
	if m.Distance == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Distance, 100, 0)
}

// CyclesScaled return Cycles in its scaled value [Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.].
//
// If Cycles value is invalid, float64 invalid value will be returned.
func (m *Monitoring) CyclesScaled() float64 {
	if m.Cycles == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Cycles, 2, 0)
}

// ActiveTimeScaled return ActiveTime in its scaled value [Scale: 1000; Units: s].
//
// If ActiveTime value is invalid, float64 invalid value will be returned.
func (m *Monitoring) ActiveTimeScaled() float64 {
	if m.ActiveTime == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.ActiveTime, 1000, 0)
}

// AscentScaled return Ascent in its scaled value [Scale: 1000; Units: m].
//
// If Ascent value is invalid, float64 invalid value will be returned.
func (m *Monitoring) AscentScaled() float64 {
	if m.Ascent == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Ascent, 1000, 0)
}

// DescentScaled return Descent in its scaled value [Scale: 1000; Units: m].
//
// If Descent value is invalid, float64 invalid value will be returned.
func (m *Monitoring) DescentScaled() float64 {
	if m.Descent == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Descent, 1000, 0)
}

// TemperatureScaled return Temperature in its scaled value [Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp].
//
// If Temperature value is invalid, float64 invalid value will be returned.
func (m *Monitoring) TemperatureScaled() float64 {
	if m.Temperature == basetype.Sint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Temperature, 100, 0)
}

// TemperatureMinScaled return TemperatureMin in its scaled value [Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp].
//
// If TemperatureMin value is invalid, float64 invalid value will be returned.
func (m *Monitoring) TemperatureMinScaled() float64 {
	if m.TemperatureMin == basetype.Sint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.TemperatureMin, 100, 0)
}

// TemperatureMaxScaled return TemperatureMax in its scaled value [Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp].
//
// If TemperatureMax value is invalid, float64 invalid value will be returned.
func (m *Monitoring) TemperatureMaxScaled() float64 {
	if m.TemperatureMax == basetype.Sint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.TemperatureMax, 100, 0)
}

// IntensityScaled return Intensity in its scaled value [Scale: 10].
//
// If Intensity value is invalid, float64 invalid value will be returned.
func (m *Monitoring) IntensityScaled() float64 {
	if m.Intensity == basetype.Uint8Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Intensity, 10, 0)
}

// SetTimestamp sets Monitoring value.
//
// Units: s; Must align to logging interval, for example, time must be 00:00:00 for daily log.
func (m *Monitoring) SetTimestamp(v time.Time) *Monitoring {
	m.Timestamp = v
	return m
}

// SetLocalTimestamp sets Monitoring value.
//
// Must align to logging interval, for example, time must be 00:00:00 for daily log.
func (m *Monitoring) SetLocalTimestamp(v time.Time) *Monitoring {
	m.LocalTimestamp = v
	return m
}

// SetActivityTime sets Monitoring value.
//
// Array: [8]; Units: minutes; Indexed using minute_activity_level enum
func (m *Monitoring) SetActivityTime(v []uint16) *Monitoring {
	m.ActivityTime = v
	return m
}

// SetDistance sets Monitoring value.
//
// Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) SetDistance(v uint32) *Monitoring {
	m.Distance = v
	return m
}

// SetCycles sets Monitoring value.
//
// Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) SetCycles(v uint32) *Monitoring {
	m.Cycles = v
	return m
}

// SetActiveTime sets Monitoring value.
//
// Scale: 1000; Units: s
func (m *Monitoring) SetActiveTime(v uint32) *Monitoring {
	m.ActiveTime = v
	return m
}

// SetDuration sets Monitoring value.
//
// Units: s
func (m *Monitoring) SetDuration(v uint32) *Monitoring {
	m.Duration = v
	return m
}

// SetAscent sets Monitoring value.
//
// Scale: 1000; Units: m
func (m *Monitoring) SetAscent(v uint32) *Monitoring {
	m.Ascent = v
	return m
}

// SetDescent sets Monitoring value.
//
// Scale: 1000; Units: m
func (m *Monitoring) SetDescent(v uint32) *Monitoring {
	m.Descent = v
	return m
}

// SetCalories sets Monitoring value.
//
// Units: kcal; Accumulated total calories. Maintained by MonitoringReader for each activity_type. See SDK documentation
func (m *Monitoring) SetCalories(v uint16) *Monitoring {
	m.Calories = v
	return m
}

// SetDistance16 sets Monitoring value.
//
// Units: 100 * m
func (m *Monitoring) SetDistance16(v uint16) *Monitoring {
	m.Distance16 = v
	return m
}

// SetCycles16 sets Monitoring value.
//
// Units: 2 * cycles (steps)
func (m *Monitoring) SetCycles16(v uint16) *Monitoring {
	m.Cycles16 = v
	return m
}

// SetActiveTime16 sets Monitoring value.
//
// Units: s
func (m *Monitoring) SetActiveTime16(v uint16) *Monitoring {
	m.ActiveTime16 = v
	return m
}

// SetTemperature sets Monitoring value.
//
// Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperature(v int16) *Monitoring {
	m.Temperature = v
	return m
}

// SetTemperatureMin sets Monitoring value.
//
// Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureMin(v int16) *Monitoring {
	m.TemperatureMin = v
	return m
}

// SetTemperatureMax sets Monitoring value.
//
// Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureMax(v int16) *Monitoring {
	m.TemperatureMax = v
	return m
}

// SetActiveCalories sets Monitoring value.
//
// Units: kcal
func (m *Monitoring) SetActiveCalories(v uint16) *Monitoring {
	m.ActiveCalories = v
	return m
}

// SetTimestamp16 sets Monitoring value.
//
// Units: s
func (m *Monitoring) SetTimestamp16(v uint16) *Monitoring {
	m.Timestamp16 = v
	return m
}

// SetDurationMin sets Monitoring value.
//
// Units: min
func (m *Monitoring) SetDurationMin(v uint16) *Monitoring {
	m.DurationMin = v
	return m
}

// SetModerateActivityMinutes sets Monitoring value.
//
// Units: minutes
func (m *Monitoring) SetModerateActivityMinutes(v uint16) *Monitoring {
	m.ModerateActivityMinutes = v
	return m
}

// SetVigorousActivityMinutes sets Monitoring value.
//
// Units: minutes
func (m *Monitoring) SetVigorousActivityMinutes(v uint16) *Monitoring {
	m.VigorousActivityMinutes = v
	return m
}

// SetDeviceIndex sets Monitoring value.
//
// Associates this data to device_info message. Not required for file with single device (sensor).
func (m *Monitoring) SetDeviceIndex(v typedef.DeviceIndex) *Monitoring {
	m.DeviceIndex = v
	return m
}

// SetActivityType sets Monitoring value.
func (m *Monitoring) SetActivityType(v typedef.ActivityType) *Monitoring {
	m.ActivityType = v
	return m
}

// SetActivitySubtype sets Monitoring value.
func (m *Monitoring) SetActivitySubtype(v typedef.ActivitySubtype) *Monitoring {
	m.ActivitySubtype = v
	return m
}

// SetActivityLevel sets Monitoring value.
func (m *Monitoring) SetActivityLevel(v typedef.ActivityLevel) *Monitoring {
	m.ActivityLevel = v
	return m
}

// SetCurrentActivityTypeIntensity sets Monitoring value.
//
// Indicates single type / intensity for duration since last monitoring message.
func (m *Monitoring) SetCurrentActivityTypeIntensity(v byte) *Monitoring {
	m.CurrentActivityTypeIntensity = v
	return m
}

// SetTimestampMin8 sets Monitoring value.
//
// Units: min
func (m *Monitoring) SetTimestampMin8(v uint8) *Monitoring {
	m.TimestampMin8 = v
	return m
}

// SetHeartRate sets Monitoring value.
//
// Units: bpm
func (m *Monitoring) SetHeartRate(v uint8) *Monitoring {
	m.HeartRate = v
	return m
}

// SetIntensity sets Monitoring value.
//
// Scale: 10
func (m *Monitoring) SetIntensity(v uint8) *Monitoring {
	m.Intensity = v
	return m
}

// SetDeveloperFields Monitoring's DeveloperFields.
func (m *Monitoring) SetDeveloperFields(developerFields ...proto.DeveloperField) *Monitoring {
	m.DeveloperFields = developerFields
	return m
}
