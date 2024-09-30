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
)

// Event is a Event message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Event struct {
	Timestamp                   time.Time // Units: s
	StartTimestamp              time.Time // Units: s; Timestamp of when the event started
	Data                        uint32
	Data16                      uint16
	Score                       uint16 // Do not populate directly. Autogenerated by decoder for sport_point subfield components
	OpponentScore               uint16 // Do not populate directly. Autogenerated by decoder for sport_point subfield components
	Event                       typedef.Event
	EventType                   typedef.EventType
	EventGroup                  uint8
	FrontGearNum                uint8 // Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Front gear number. 1 is innermost.
	FrontGear                   uint8 // Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of front teeth.
	RearGearNum                 uint8 // Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Rear gear number. 1 is innermost.
	RearGear                    uint8 // Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of rear teeth.
	DeviceIndex                 typedef.DeviceIndex
	ActivityType                typedef.ActivityType         // Activity Type associated with an auto_activity_detect event
	RadarThreatLevelMax         typedef.RadarThreatLevelType // Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
	RadarThreatCount            uint8                        // Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
	RadarThreatAvgApproachSpeed uint8                        // Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
	RadarThreatMaxApproachSpeed uint8                        // Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components

	state [4]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewEvent creates new Event struct based on given mesg.
// If mesg is nil, it will return Event with all fields being set to its corresponding invalid value.
func NewEvent(mesg *proto.Message) *Event {
	vals := [254]proto.Value{}

	var state [4]uint8
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
			if mesg.Fields[i].Num < 25 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Event{
		Timestamp:                   datetime.ToTime(vals[253].Uint32()),
		Event:                       typedef.Event(vals[0].Uint8()),
		EventType:                   typedef.EventType(vals[1].Uint8()),
		Data16:                      vals[2].Uint16(),
		Data:                        vals[3].Uint32(),
		EventGroup:                  vals[4].Uint8(),
		Score:                       vals[7].Uint16(),
		OpponentScore:               vals[8].Uint16(),
		FrontGearNum:                vals[9].Uint8z(),
		FrontGear:                   vals[10].Uint8z(),
		RearGearNum:                 vals[11].Uint8z(),
		RearGear:                    vals[12].Uint8z(),
		DeviceIndex:                 typedef.DeviceIndex(vals[13].Uint8()),
		ActivityType:                typedef.ActivityType(vals[14].Uint8()),
		StartTimestamp:              datetime.ToTime(vals[15].Uint32()),
		RadarThreatLevelMax:         typedef.RadarThreatLevelType(vals[21].Uint8()),
		RadarThreatCount:            vals[22].Uint8(),
		RadarThreatAvgApproachSpeed: vals[23].Uint8(),
		RadarThreatMaxApproachSpeed: vals[24].Uint8(),

		state: state,

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Event into proto.Message. If options is nil, default options will be used.
func (m *Event) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumEvent}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Event != typedef.EventInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Event))
		fields = append(fields, field)
	}
	if m.EventType != typedef.EventTypeInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.EventType))
		fields = append(fields, field)
	}
	if m.Data16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.Data16)
		fields = append(fields, field)
	}
	if m.Data != basetype.Uint32Invalid {
		if expanded := m.IsExpandedField(3); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 3)
			field.Value = proto.Uint32(m.Data)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(m.EventGroup)
		fields = append(fields, field)
	}
	if m.Score != basetype.Uint16Invalid {
		if expanded := m.IsExpandedField(7); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 7)
			field.Value = proto.Uint16(m.Score)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		if expanded := m.IsExpandedField(8); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 8)
			field.Value = proto.Uint16(m.OpponentScore)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.FrontGearNum != basetype.Uint8zInvalid {
		if expanded := m.IsExpandedField(9); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 9)
			field.Value = proto.Uint8(m.FrontGearNum)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.FrontGear != basetype.Uint8zInvalid {
		if expanded := m.IsExpandedField(10); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 10)
			field.Value = proto.Uint8(m.FrontGear)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.RearGearNum != basetype.Uint8zInvalid {
		if expanded := m.IsExpandedField(11); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 11)
			field.Value = proto.Uint8(m.RearGearNum)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.RearGear != basetype.Uint8zInvalid {
		if expanded := m.IsExpandedField(12); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 12)
			field.Value = proto.Uint8(m.RearGear)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.DeviceIndex != typedef.DeviceIndexInvalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
	}
	if m.ActivityType != typedef.ActivityTypeInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(byte(m.ActivityType))
		fields = append(fields, field)
	}
	if !m.StartTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint32(uint32(m.StartTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.RadarThreatLevelMax != typedef.RadarThreatLevelTypeInvalid {
		if expanded := m.IsExpandedField(21); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 21)
			field.Value = proto.Uint8(byte(m.RadarThreatLevelMax))
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.RadarThreatCount != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(22); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 22)
			field.Value = proto.Uint8(m.RadarThreatCount)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.RadarThreatAvgApproachSpeed != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(23); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 23)
			field.Value = proto.Uint8(m.RadarThreatAvgApproachSpeed)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.RadarThreatMaxApproachSpeed != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(24); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 24)
			field.Value = proto.Uint8(m.RadarThreatMaxApproachSpeed)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetData returns Dynamic Field interpretation of Data. Otherwise, returns the original value of Data.
//
// Based on m.Event:
//   - name: "timer_trigger", value: typedef.TimerTrigger(m.Data)
//   - name: "course_point_index", value: typedef.MessageIndex(m.Data)
//   - name: "battery_level", units: "V" , value: (float64(m.Data) * 1000) - 0
//   - name: "virtual_partner_speed", units: "m/s" , value: (float64(m.Data) * 1000) - 0
//   - name: "hr_high_alert", units: "bpm" , value: uint8(m.Data)
//   - name: "hr_low_alert", units: "bpm" , value: uint8(m.Data)
//   - name: "speed_high_alert", units: "m/s" , value: (float64(m.Data) * 1000) - 0
//   - name: "speed_low_alert", units: "m/s" , value: (float64(m.Data) * 1000) - 0
//   - name: "cad_high_alert", units: "rpm" , value: uint16(m.Data)
//   - name: "cad_low_alert", units: "rpm" , value: uint16(m.Data)
//   - name: "power_high_alert", units: "watts" , value: uint16(m.Data)
//   - name: "power_low_alert", units: "watts" , value: uint16(m.Data)
//   - name: "time_duration_alert", units: "s" , value: (float64(m.Data) * 1000) - 0
//   - name: "distance_duration_alert", units: "m" , value: (float64(m.Data) * 100) - 0
//   - name: "calorie_duration_alert", units: "calories" , value: uint32(m.Data)
//   - name: "fitness_equipment_state", value: typedef.FitnessEquipmentState(m.Data)
//   - name: "sport_point", value: uint32(m.Data)
//   - name: "gear_change_data", value: uint32(m.Data)
//   - name: "rider_position", value: typedef.RiderPositionType(m.Data)
//   - name: "comm_timeout", value: typedef.CommTimeoutType(m.Data)
//   - name: "dive_alert", value: typedef.DiveAlert(m.Data)
//   - name: "auto_activity_detect_duration", units: "min" , value: uint16(m.Data)
//   - name: "radar_threat_alert", value: uint32(m.Data)
//
// Otherwise:
//   - name: "data", value: m.Data
func (m *Event) GetData() (name string, value any) {
	switch m.Event {
	case typedef.EventTimer:
		return "timer_trigger", typedef.TimerTrigger(m.Data)
	case typedef.EventCoursePoint:
		return "course_point_index", typedef.MessageIndex(m.Data)
	case typedef.EventBattery:
		return "battery_level", (float64(m.Data) * 1000) - 0
	case typedef.EventVirtualPartnerPace:
		return "virtual_partner_speed", (float64(m.Data) * 1000) - 0
	case typedef.EventHrHighAlert:
		return "hr_high_alert", uint8(m.Data)
	case typedef.EventHrLowAlert:
		return "hr_low_alert", uint8(m.Data)
	case typedef.EventSpeedHighAlert:
		return "speed_high_alert", (float64(m.Data) * 1000) - 0
	case typedef.EventSpeedLowAlert:
		return "speed_low_alert", (float64(m.Data) * 1000) - 0
	case typedef.EventCadHighAlert:
		return "cad_high_alert", uint16(m.Data)
	case typedef.EventCadLowAlert:
		return "cad_low_alert", uint16(m.Data)
	case typedef.EventPowerHighAlert:
		return "power_high_alert", uint16(m.Data)
	case typedef.EventPowerLowAlert:
		return "power_low_alert", uint16(m.Data)
	case typedef.EventTimeDurationAlert:
		return "time_duration_alert", (float64(m.Data) * 1000) - 0
	case typedef.EventDistanceDurationAlert:
		return "distance_duration_alert", (float64(m.Data) * 100) - 0
	case typedef.EventCalorieDurationAlert:
		return "calorie_duration_alert", uint32(m.Data)
	case typedef.EventFitnessEquipment:
		return "fitness_equipment_state", typedef.FitnessEquipmentState(m.Data)
	case typedef.EventSportPoint:
		return "sport_point", uint32(m.Data)
	case typedef.EventFrontGearChange, typedef.EventRearGearChange:
		return "gear_change_data", uint32(m.Data)
	case typedef.EventRiderPositionChange:
		return "rider_position", typedef.RiderPositionType(m.Data)
	case typedef.EventCommTimeout:
		return "comm_timeout", typedef.CommTimeoutType(m.Data)
	case typedef.EventDiveAlert:
		return "dive_alert", typedef.DiveAlert(m.Data)
	case typedef.EventAutoActivityDetect:
		return "auto_activity_detect_duration", uint16(m.Data)
	case typedef.EventRadarThreatAlert:
		return "radar_threat_alert", uint32(m.Data)
	}
	return "data", m.Data
}

// GetStartTimestamp returns Dynamic Field interpretation of StartTimestamp. Otherwise, returns the original value of StartTimestamp.
//
// Based on m.Event:
//   - name: "auto_activity_detect_start_timestamp", units: "s" , value: time.Time(m.StartTimestamp)
//
// Otherwise:
//   - name: "start_timestamp", units: "s" , value: m.StartTimestamp
func (m *Event) GetStartTimestamp() (name string, value any) {
	switch m.Event {
	case typedef.EventAutoActivityDetect:
		return "auto_activity_detect_start_timestamp", time.Time(m.StartTimestamp)
	}
	return "start_timestamp", m.StartTimestamp
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Event) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// StartTimestampUint32 returns StartTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Event) StartTimestampUint32() uint32 { return datetime.ToUint32(m.StartTimestamp) }

// RadarThreatAvgApproachSpeedScaled return RadarThreatAvgApproachSpeed in its scaled value.
// If RadarThreatAvgApproachSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) RadarThreatAvgApproachSpeedScaled() float64 {
	if m.RadarThreatAvgApproachSpeed == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.RadarThreatAvgApproachSpeed)/10 - 0
}

// RadarThreatMaxApproachSpeedScaled return RadarThreatMaxApproachSpeed in its scaled value.
// If RadarThreatMaxApproachSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) RadarThreatMaxApproachSpeedScaled() float64 {
	if m.RadarThreatMaxApproachSpeed == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.RadarThreatMaxApproachSpeed)/10 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *Event) SetTimestamp(v time.Time) *Event {
	m.Timestamp = v
	return m
}

// SetEvent sets Event value.
func (m *Event) SetEvent(v typedef.Event) *Event {
	m.Event = v
	return m
}

// SetEventType sets EventType value.
func (m *Event) SetEventType(v typedef.EventType) *Event {
	m.EventType = v
	return m
}

// SetData16 sets Data16 value.
func (m *Event) SetData16(v uint16) *Event {
	m.Data16 = v
	return m
}

// SetData sets Data value.
func (m *Event) SetData(v uint32) *Event {
	m.Data = v
	return m
}

// SetEventGroup sets EventGroup value.
func (m *Event) SetEventGroup(v uint8) *Event {
	m.EventGroup = v
	return m
}

// SetScore sets Score value.
//
// Do not populate directly. Autogenerated by decoder for sport_point subfield components
func (m *Event) SetScore(v uint16) *Event {
	m.Score = v
	return m
}

// SetOpponentScore sets OpponentScore value.
//
// Do not populate directly. Autogenerated by decoder for sport_point subfield components
func (m *Event) SetOpponentScore(v uint16) *Event {
	m.OpponentScore = v
	return m
}

// SetFrontGearNum sets FrontGearNum value.
//
// Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Front gear number. 1 is innermost.
func (m *Event) SetFrontGearNum(v uint8) *Event {
	m.FrontGearNum = v
	return m
}

// SetFrontGear sets FrontGear value.
//
// Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of front teeth.
func (m *Event) SetFrontGear(v uint8) *Event {
	m.FrontGear = v
	return m
}

// SetRearGearNum sets RearGearNum value.
//
// Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Rear gear number. 1 is innermost.
func (m *Event) SetRearGearNum(v uint8) *Event {
	m.RearGearNum = v
	return m
}

// SetRearGear sets RearGear value.
//
// Base: uint8z; Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of rear teeth.
func (m *Event) SetRearGear(v uint8) *Event {
	m.RearGear = v
	return m
}

// SetDeviceIndex sets DeviceIndex value.
func (m *Event) SetDeviceIndex(v typedef.DeviceIndex) *Event {
	m.DeviceIndex = v
	return m
}

// SetActivityType sets ActivityType value.
//
// Activity Type associated with an auto_activity_detect event
func (m *Event) SetActivityType(v typedef.ActivityType) *Event {
	m.ActivityType = v
	return m
}

// SetStartTimestamp sets StartTimestamp value.
//
// Units: s; Timestamp of when the event started
func (m *Event) SetStartTimestamp(v time.Time) *Event {
	m.StartTimestamp = v
	return m
}

// SetRadarThreatLevelMax sets RadarThreatLevelMax value.
//
// Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
func (m *Event) SetRadarThreatLevelMax(v typedef.RadarThreatLevelType) *Event {
	m.RadarThreatLevelMax = v
	return m
}

// SetRadarThreatCount sets RadarThreatCount value.
//
// Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
func (m *Event) SetRadarThreatCount(v uint8) *Event {
	m.RadarThreatCount = v
	return m
}

// SetRadarThreatAvgApproachSpeed sets RadarThreatAvgApproachSpeed value.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) SetRadarThreatAvgApproachSpeed(v uint8) *Event {
	m.RadarThreatAvgApproachSpeed = v
	return m
}

// SetRadarThreatAvgApproachSpeedScaled is similar to SetRadarThreatAvgApproachSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) SetRadarThreatAvgApproachSpeedScaled(v float64) *Event {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.RadarThreatAvgApproachSpeed = uint8(basetype.Uint8Invalid)
		return m
	}
	m.RadarThreatAvgApproachSpeed = uint8(unscaled)
	return m
}

// SetRadarThreatMaxApproachSpeed sets RadarThreatMaxApproachSpeed value.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) SetRadarThreatMaxApproachSpeed(v uint8) *Event {
	m.RadarThreatMaxApproachSpeed = v
	return m
}

// SetRadarThreatMaxApproachSpeedScaled is similar to SetRadarThreatMaxApproachSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) SetRadarThreatMaxApproachSpeedScaled(v float64) *Event {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.RadarThreatMaxApproachSpeed = uint8(basetype.Uint8Invalid)
		return m
	}
	m.RadarThreatMaxApproachSpeed = uint8(unscaled)
	return m
}

// SetDeveloperFields Event's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Event) SetUnknownFields(unknownFields ...proto.Field) *Event {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Event's DeveloperFields.
func (m *Event) SetDeveloperFields(developerFields ...proto.DeveloperField) *Event {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 3, 7, 8, 9, 10, 11, 12, 21, 22, 23, 24.
func (m *Event) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 3, 7, 8, 9, 10, 11, 12, 21, 22, 23, 24:
	default:
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	m.state[pos] &^= bit
	if flag {
		m.state[pos] |= bit
	}
	return true
}

// IsExpandedField checks whether given fieldNum is a field generated through
// a component expansion. Eligible for field number: 3, 7, 8, 9, 10, 11, 12, 21, 22, 23, 24.
func (m *Event) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 25 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}
