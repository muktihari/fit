// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// WorkoutStep is a WorkoutStep message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type WorkoutStep struct {
	WktStepName                    string
	Notes                          string
	DurationValue                  uint32
	TargetValue                    uint32
	CustomTargetValueLow           uint32
	CustomTargetValueHigh          uint32
	SecondaryTargetValue           uint32
	SecondaryCustomTargetValueLow  uint32
	SecondaryCustomTargetValueHigh uint32
	MessageIndex                   typedef.MessageIndex
	ExerciseCategory               typedef.ExerciseCategory
	ExerciseName                   uint16
	ExerciseWeight                 uint16 // Scale: 100; Units: kg
	WeightDisplayUnit              typedef.FitBaseUnit
	DurationType                   typedef.WktStepDuration
	TargetType                     typedef.WktStepTarget
	Intensity                      typedef.Intensity
	Equipment                      typedef.WorkoutEquipment
	SecondaryTargetType            typedef.WktStepTarget

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWorkoutStep creates new WorkoutStep struct based on given mesg.
// If mesg is nil, it will return WorkoutStep with all fields being set to its corresponding invalid value.
func NewWorkoutStep(mesg *proto.Message) *WorkoutStep {
	vals := [255]proto.Value{}

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

	return &WorkoutStep{
		MessageIndex:                   typedef.MessageIndex(vals[254].Uint16()),
		WktStepName:                    vals[0].String(),
		DurationType:                   typedef.WktStepDuration(vals[1].Uint8()),
		DurationValue:                  vals[2].Uint32(),
		TargetType:                     typedef.WktStepTarget(vals[3].Uint8()),
		TargetValue:                    vals[4].Uint32(),
		CustomTargetValueLow:           vals[5].Uint32(),
		CustomTargetValueHigh:          vals[6].Uint32(),
		Intensity:                      typedef.Intensity(vals[7].Uint8()),
		Notes:                          vals[8].String(),
		Equipment:                      typedef.WorkoutEquipment(vals[9].Uint8()),
		ExerciseCategory:               typedef.ExerciseCategory(vals[10].Uint16()),
		ExerciseName:                   vals[11].Uint16(),
		ExerciseWeight:                 vals[12].Uint16(),
		WeightDisplayUnit:              typedef.FitBaseUnit(vals[13].Uint16()),
		SecondaryTargetType:            typedef.WktStepTarget(vals[19].Uint8()),
		SecondaryTargetValue:           vals[20].Uint32(),
		SecondaryCustomTargetValueLow:  vals[21].Uint32(),
		SecondaryCustomTargetValueHigh: vals[22].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WorkoutStep into proto.Message. If options is nil, default options will be used.
func (m *WorkoutStep) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumWorkoutStep}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.WktStepName != basetype.StringInvalid && m.WktStepName != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.WktStepName)
		fields = append(fields, field)
	}
	if byte(m.DurationType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.DurationType))
		fields = append(fields, field)
	}
	if m.DurationValue != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.DurationValue)
		fields = append(fields, field)
	}
	if byte(m.TargetType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.TargetType))
		fields = append(fields, field)
	}
	if m.TargetValue != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.TargetValue)
		fields = append(fields, field)
	}
	if m.CustomTargetValueLow != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.CustomTargetValueLow)
		fields = append(fields, field)
	}
	if m.CustomTargetValueHigh != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.CustomTargetValueHigh)
		fields = append(fields, field)
	}
	if byte(m.Intensity) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.Intensity))
		fields = append(fields, field)
	}
	if m.Notes != basetype.StringInvalid && m.Notes != "" {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.String(m.Notes)
		fields = append(fields, field)
	}
	if byte(m.Equipment) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(byte(m.Equipment))
		fields = append(fields, field)
	}
	if uint16(m.ExerciseCategory) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint16(uint16(m.ExerciseCategory))
		fields = append(fields, field)
	}
	if m.ExerciseName != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint16(m.ExerciseName)
		fields = append(fields, field)
	}
	if m.ExerciseWeight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint16(m.ExerciseWeight)
		fields = append(fields, field)
	}
	if uint16(m.WeightDisplayUnit) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint16(uint16(m.WeightDisplayUnit))
		fields = append(fields, field)
	}
	if byte(m.SecondaryTargetType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.Uint8(byte(m.SecondaryTargetType))
		fields = append(fields, field)
	}
	if m.SecondaryTargetValue != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = proto.Uint32(m.SecondaryTargetValue)
		fields = append(fields, field)
	}
	if m.SecondaryCustomTargetValueLow != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint32(m.SecondaryCustomTargetValueLow)
		fields = append(fields, field)
	}
	if m.SecondaryCustomTargetValueHigh != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint32(m.SecondaryCustomTargetValueHigh)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetDurationValue returns Dynamic Field interpretation of DurationValue. Otherwise, returns the original value of DurationValue.
//
// Based on m.DurationType:
//   - name: "duration_time", units: "s" , value: (float64(m.DurationValue) * 1000) - 0
//   - name: "duration_distance", units: "m" , value: (float64(m.DurationValue) * 100) - 0
//   - name: "duration_hr", units: "% or bpm" , value: typedef.WorkoutHr(m.DurationValue)
//   - name: "duration_calories", units: "calories" , value: uint32(m.DurationValue)
//   - name: "duration_step", value: uint32(m.DurationValue)
//   - name: "duration_power", units: "% or watts" , value: typedef.WorkoutPower(m.DurationValue)
//   - name: "duration_reps", value: uint32(m.DurationValue)
//
// Otherwise:
//   - name: "duration_value", value: m.DurationValue
func (m *WorkoutStep) GetDurationValue() (name string, value any) {
	switch m.DurationType {
	case typedef.WktStepDurationTime, typedef.WktStepDurationRepetitionTime:
		return "duration_time", (float64(m.DurationValue) * 1000) - 0
	case typedef.WktStepDurationDistance:
		return "duration_distance", (float64(m.DurationValue) * 100) - 0
	case typedef.WktStepDurationHrLessThan, typedef.WktStepDurationHrGreaterThan:
		return "duration_hr", typedef.WorkoutHr(m.DurationValue)
	case typedef.WktStepDurationCalories:
		return "duration_calories", uint32(m.DurationValue)
	case typedef.WktStepDurationRepeatUntilStepsCmplt, typedef.WktStepDurationRepeatUntilTime, typedef.WktStepDurationRepeatUntilDistance, typedef.WktStepDurationRepeatUntilCalories, typedef.WktStepDurationRepeatUntilHrLessThan, typedef.WktStepDurationRepeatUntilHrGreaterThan, typedef.WktStepDurationRepeatUntilPowerLessThan, typedef.WktStepDurationRepeatUntilPowerGreaterThan:
		return "duration_step", uint32(m.DurationValue)
	case typedef.WktStepDurationPowerLessThan, typedef.WktStepDurationPowerGreaterThan:
		return "duration_power", typedef.WorkoutPower(m.DurationValue)
	case typedef.WktStepDurationReps:
		return "duration_reps", uint32(m.DurationValue)
	}
	return "duration_value", m.DurationValue
}

// GetTargetValue returns Dynamic Field interpretation of TargetValue. Otherwise, returns the original value of TargetValue.
//
// Based on m.TargetType:
//   - name: "target_speed_zone", value: uint32(m.TargetValue)
//   - name: "target_hr_zone", value: uint32(m.TargetValue)
//   - name: "target_cadence_zone", value: uint32(m.TargetValue)
//   - name: "target_power_zone", value: uint32(m.TargetValue)
//   - name: "target_stroke_type", value: typedef.SwimStroke(m.TargetValue)
//
// Based on m.DurationType:
//   - name: "repeat_steps", value: uint32(m.TargetValue)
//   - name: "repeat_time", units: "s" , value: (float64(m.TargetValue) * 1000) - 0
//   - name: "repeat_distance", units: "m" , value: (float64(m.TargetValue) * 100) - 0
//   - name: "repeat_calories", units: "calories" , value: uint32(m.TargetValue)
//   - name: "repeat_hr", units: "% or bpm" , value: typedef.WorkoutHr(m.TargetValue)
//   - name: "repeat_power", units: "% or watts" , value: typedef.WorkoutPower(m.TargetValue)
//
// Otherwise:
//   - name: "target_value", value: m.TargetValue
func (m *WorkoutStep) GetTargetValue() (name string, value any) {
	switch m.TargetType {
	case typedef.WktStepTargetSpeed:
		return "target_speed_zone", uint32(m.TargetValue)
	case typedef.WktStepTargetHeartRate:
		return "target_hr_zone", uint32(m.TargetValue)
	case typedef.WktStepTargetCadence:
		return "target_cadence_zone", uint32(m.TargetValue)
	case typedef.WktStepTargetPower:
		return "target_power_zone", uint32(m.TargetValue)
	case typedef.WktStepTargetSwimStroke:
		return "target_stroke_type", typedef.SwimStroke(m.TargetValue)
	}
	switch m.DurationType {
	case typedef.WktStepDurationRepeatUntilStepsCmplt:
		return "repeat_steps", uint32(m.TargetValue)
	case typedef.WktStepDurationRepeatUntilTime:
		return "repeat_time", (float64(m.TargetValue) * 1000) - 0
	case typedef.WktStepDurationRepeatUntilDistance:
		return "repeat_distance", (float64(m.TargetValue) * 100) - 0
	case typedef.WktStepDurationRepeatUntilCalories:
		return "repeat_calories", uint32(m.TargetValue)
	case typedef.WktStepDurationRepeatUntilHrLessThan, typedef.WktStepDurationRepeatUntilHrGreaterThan:
		return "repeat_hr", typedef.WorkoutHr(m.TargetValue)
	case typedef.WktStepDurationRepeatUntilPowerLessThan, typedef.WktStepDurationRepeatUntilPowerGreaterThan:
		return "repeat_power", typedef.WorkoutPower(m.TargetValue)
	}
	return "target_value", m.TargetValue
}

// GetCustomTargetValueLow returns Dynamic Field interpretation of CustomTargetValueLow. Otherwise, returns the original value of CustomTargetValueLow.
//
// Based on m.TargetType:
//   - name: "custom_target_speed_low", units: "m/s" , value: (float64(m.CustomTargetValueLow) * 1000) - 0
//   - name: "custom_target_heart_rate_low", units: "% or bpm" , value: typedef.WorkoutHr(m.CustomTargetValueLow)
//   - name: "custom_target_cadence_low", units: "rpm" , value: uint32(m.CustomTargetValueLow)
//   - name: "custom_target_power_low", units: "% or watts" , value: typedef.WorkoutPower(m.CustomTargetValueLow)
//
// Otherwise:
//   - name: "custom_target_value_low", value: m.CustomTargetValueLow
func (m *WorkoutStep) GetCustomTargetValueLow() (name string, value any) {
	switch m.TargetType {
	case typedef.WktStepTargetSpeed:
		return "custom_target_speed_low", (float64(m.CustomTargetValueLow) * 1000) - 0
	case typedef.WktStepTargetHeartRate:
		return "custom_target_heart_rate_low", typedef.WorkoutHr(m.CustomTargetValueLow)
	case typedef.WktStepTargetCadence:
		return "custom_target_cadence_low", uint32(m.CustomTargetValueLow)
	case typedef.WktStepTargetPower:
		return "custom_target_power_low", typedef.WorkoutPower(m.CustomTargetValueLow)
	}
	return "custom_target_value_low", m.CustomTargetValueLow
}

// GetCustomTargetValueHigh returns Dynamic Field interpretation of CustomTargetValueHigh. Otherwise, returns the original value of CustomTargetValueHigh.
//
// Based on m.TargetType:
//   - name: "custom_target_speed_high", units: "m/s" , value: (float64(m.CustomTargetValueHigh) * 1000) - 0
//   - name: "custom_target_heart_rate_high", units: "% or bpm" , value: typedef.WorkoutHr(m.CustomTargetValueHigh)
//   - name: "custom_target_cadence_high", units: "rpm" , value: uint32(m.CustomTargetValueHigh)
//   - name: "custom_target_power_high", units: "% or watts" , value: typedef.WorkoutPower(m.CustomTargetValueHigh)
//
// Otherwise:
//   - name: "custom_target_value_high", value: m.CustomTargetValueHigh
func (m *WorkoutStep) GetCustomTargetValueHigh() (name string, value any) {
	switch m.TargetType {
	case typedef.WktStepTargetSpeed:
		return "custom_target_speed_high", (float64(m.CustomTargetValueHigh) * 1000) - 0
	case typedef.WktStepTargetHeartRate:
		return "custom_target_heart_rate_high", typedef.WorkoutHr(m.CustomTargetValueHigh)
	case typedef.WktStepTargetCadence:
		return "custom_target_cadence_high", uint32(m.CustomTargetValueHigh)
	case typedef.WktStepTargetPower:
		return "custom_target_power_high", typedef.WorkoutPower(m.CustomTargetValueHigh)
	}
	return "custom_target_value_high", m.CustomTargetValueHigh
}

// GetSecondaryTargetValue returns Dynamic Field interpretation of SecondaryTargetValue. Otherwise, returns the original value of SecondaryTargetValue.
//
// Based on m.SecondaryTargetType:
//   - name: "secondary_target_speed_zone", value: uint32(m.SecondaryTargetValue)
//   - name: "secondary_target_hr_zone", value: uint32(m.SecondaryTargetValue)
//   - name: "secondary_target_cadence_zone", value: uint32(m.SecondaryTargetValue)
//   - name: "secondary_target_power_zone", value: uint32(m.SecondaryTargetValue)
//   - name: "secondary_target_stroke_type", value: typedef.SwimStroke(m.SecondaryTargetValue)
//
// Otherwise:
//   - name: "secondary_target_value", value: m.SecondaryTargetValue
func (m *WorkoutStep) GetSecondaryTargetValue() (name string, value any) {
	switch m.SecondaryTargetType {
	case typedef.WktStepTargetSpeed:
		return "secondary_target_speed_zone", uint32(m.SecondaryTargetValue)
	case typedef.WktStepTargetHeartRate:
		return "secondary_target_hr_zone", uint32(m.SecondaryTargetValue)
	case typedef.WktStepTargetCadence:
		return "secondary_target_cadence_zone", uint32(m.SecondaryTargetValue)
	case typedef.WktStepTargetPower:
		return "secondary_target_power_zone", uint32(m.SecondaryTargetValue)
	case typedef.WktStepTargetSwimStroke:
		return "secondary_target_stroke_type", typedef.SwimStroke(m.SecondaryTargetValue)
	}
	return "secondary_target_value", m.SecondaryTargetValue
}

// GetSecondaryCustomTargetValueLow returns Dynamic Field interpretation of SecondaryCustomTargetValueLow. Otherwise, returns the original value of SecondaryCustomTargetValueLow.
//
// Based on m.SecondaryTargetType:
//   - name: "secondary_custom_target_speed_low", units: "m/s" , value: (float64(m.SecondaryCustomTargetValueLow) * 1000) - 0
//   - name: "secondary_custom_target_heart_rate_low", units: "% or bpm" , value: typedef.WorkoutHr(m.SecondaryCustomTargetValueLow)
//   - name: "secondary_custom_target_cadence_low", units: "rpm" , value: uint32(m.SecondaryCustomTargetValueLow)
//   - name: "secondary_custom_target_power_low", units: "% or watts" , value: typedef.WorkoutPower(m.SecondaryCustomTargetValueLow)
//
// Otherwise:
//   - name: "secondary_custom_target_value_low", value: m.SecondaryCustomTargetValueLow
func (m *WorkoutStep) GetSecondaryCustomTargetValueLow() (name string, value any) {
	switch m.SecondaryTargetType {
	case typedef.WktStepTargetSpeed:
		return "secondary_custom_target_speed_low", (float64(m.SecondaryCustomTargetValueLow) * 1000) - 0
	case typedef.WktStepTargetHeartRate:
		return "secondary_custom_target_heart_rate_low", typedef.WorkoutHr(m.SecondaryCustomTargetValueLow)
	case typedef.WktStepTargetCadence:
		return "secondary_custom_target_cadence_low", uint32(m.SecondaryCustomTargetValueLow)
	case typedef.WktStepTargetPower:
		return "secondary_custom_target_power_low", typedef.WorkoutPower(m.SecondaryCustomTargetValueLow)
	}
	return "secondary_custom_target_value_low", m.SecondaryCustomTargetValueLow
}

// GetSecondaryCustomTargetValueHigh returns Dynamic Field interpretation of SecondaryCustomTargetValueHigh. Otherwise, returns the original value of SecondaryCustomTargetValueHigh.
//
// Based on m.SecondaryTargetType:
//   - name: "secondary_custom_target_speed_high", units: "m/s" , value: (float64(m.SecondaryCustomTargetValueHigh) * 1000) - 0
//   - name: "secondary_custom_target_heart_rate_high", units: "% or bpm" , value: typedef.WorkoutHr(m.SecondaryCustomTargetValueHigh)
//   - name: "secondary_custom_target_cadence_high", units: "rpm" , value: uint32(m.SecondaryCustomTargetValueHigh)
//   - name: "secondary_custom_target_power_high", units: "% or watts" , value: typedef.WorkoutPower(m.SecondaryCustomTargetValueHigh)
//
// Otherwise:
//   - name: "secondary_custom_target_value_high", value: m.SecondaryCustomTargetValueHigh
func (m *WorkoutStep) GetSecondaryCustomTargetValueHigh() (name string, value any) {
	switch m.SecondaryTargetType {
	case typedef.WktStepTargetSpeed:
		return "secondary_custom_target_speed_high", (float64(m.SecondaryCustomTargetValueHigh) * 1000) - 0
	case typedef.WktStepTargetHeartRate:
		return "secondary_custom_target_heart_rate_high", typedef.WorkoutHr(m.SecondaryCustomTargetValueHigh)
	case typedef.WktStepTargetCadence:
		return "secondary_custom_target_cadence_high", uint32(m.SecondaryCustomTargetValueHigh)
	case typedef.WktStepTargetPower:
		return "secondary_custom_target_power_high", typedef.WorkoutPower(m.SecondaryCustomTargetValueHigh)
	}
	return "secondary_custom_target_value_high", m.SecondaryCustomTargetValueHigh
}

// ExerciseWeightScaled return ExerciseWeight in its scaled value.
// If ExerciseWeight value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: kg
func (m *WorkoutStep) ExerciseWeightScaled() float64 {
	if m.ExerciseWeight == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.ExerciseWeight)/100 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *WorkoutStep) SetMessageIndex(v typedef.MessageIndex) *WorkoutStep {
	m.MessageIndex = v
	return m
}

// SetWktStepName sets WktStepName value.
func (m *WorkoutStep) SetWktStepName(v string) *WorkoutStep {
	m.WktStepName = v
	return m
}

// SetDurationType sets DurationType value.
func (m *WorkoutStep) SetDurationType(v typedef.WktStepDuration) *WorkoutStep {
	m.DurationType = v
	return m
}

// SetDurationValue sets DurationValue value.
func (m *WorkoutStep) SetDurationValue(v uint32) *WorkoutStep {
	m.DurationValue = v
	return m
}

// SetTargetType sets TargetType value.
func (m *WorkoutStep) SetTargetType(v typedef.WktStepTarget) *WorkoutStep {
	m.TargetType = v
	return m
}

// SetTargetValue sets TargetValue value.
func (m *WorkoutStep) SetTargetValue(v uint32) *WorkoutStep {
	m.TargetValue = v
	return m
}

// SetCustomTargetValueLow sets CustomTargetValueLow value.
func (m *WorkoutStep) SetCustomTargetValueLow(v uint32) *WorkoutStep {
	m.CustomTargetValueLow = v
	return m
}

// SetCustomTargetValueHigh sets CustomTargetValueHigh value.
func (m *WorkoutStep) SetCustomTargetValueHigh(v uint32) *WorkoutStep {
	m.CustomTargetValueHigh = v
	return m
}

// SetIntensity sets Intensity value.
func (m *WorkoutStep) SetIntensity(v typedef.Intensity) *WorkoutStep {
	m.Intensity = v
	return m
}

// SetNotes sets Notes value.
func (m *WorkoutStep) SetNotes(v string) *WorkoutStep {
	m.Notes = v
	return m
}

// SetEquipment sets Equipment value.
func (m *WorkoutStep) SetEquipment(v typedef.WorkoutEquipment) *WorkoutStep {
	m.Equipment = v
	return m
}

// SetExerciseCategory sets ExerciseCategory value.
func (m *WorkoutStep) SetExerciseCategory(v typedef.ExerciseCategory) *WorkoutStep {
	m.ExerciseCategory = v
	return m
}

// SetExerciseName sets ExerciseName value.
func (m *WorkoutStep) SetExerciseName(v uint16) *WorkoutStep {
	m.ExerciseName = v
	return m
}

// SetExerciseWeight sets ExerciseWeight value.
//
// Scale: 100; Units: kg
func (m *WorkoutStep) SetExerciseWeight(v uint16) *WorkoutStep {
	m.ExerciseWeight = v
	return m
}

// SetExerciseWeightScaled is similar to SetExerciseWeight except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: kg
func (m *WorkoutStep) SetExerciseWeightScaled(v float64) *WorkoutStep {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.ExerciseWeight = uint16(basetype.Uint16Invalid)
		return m
	}
	m.ExerciseWeight = uint16(unscaled)
	return m
}

// SetWeightDisplayUnit sets WeightDisplayUnit value.
func (m *WorkoutStep) SetWeightDisplayUnit(v typedef.FitBaseUnit) *WorkoutStep {
	m.WeightDisplayUnit = v
	return m
}

// SetSecondaryTargetType sets SecondaryTargetType value.
func (m *WorkoutStep) SetSecondaryTargetType(v typedef.WktStepTarget) *WorkoutStep {
	m.SecondaryTargetType = v
	return m
}

// SetSecondaryTargetValue sets SecondaryTargetValue value.
func (m *WorkoutStep) SetSecondaryTargetValue(v uint32) *WorkoutStep {
	m.SecondaryTargetValue = v
	return m
}

// SetSecondaryCustomTargetValueLow sets SecondaryCustomTargetValueLow value.
func (m *WorkoutStep) SetSecondaryCustomTargetValueLow(v uint32) *WorkoutStep {
	m.SecondaryCustomTargetValueLow = v
	return m
}

// SetSecondaryCustomTargetValueHigh sets SecondaryCustomTargetValueHigh value.
func (m *WorkoutStep) SetSecondaryCustomTargetValueHigh(v uint32) *WorkoutStep {
	m.SecondaryCustomTargetValueHigh = v
	return m
}

// SetDeveloperFields WorkoutStep's DeveloperFields.
func (m *WorkoutStep) SetDeveloperFields(developerFields ...proto.DeveloperField) *WorkoutStep {
	m.DeveloperFields = developerFields
	return m
}
