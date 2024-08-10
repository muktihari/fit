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
)

// DiveSettings is a DiveSettings message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type DiveSettings struct {
	Timestamp                 time.Time
	Name                      string
	WaterDensity              float32 // Units: kg/m^3; Fresh water is usually 1000; salt water is usually 1025
	BottomDepth               float32
	BottomTime                uint32
	ApneaCountdownTime        uint32
	CcrLowSetpointDepth       uint32 // Scale: 1000; Units: m; Depth to switch to low setpoint in automatic mode
	CcrHighSetpointDepth      uint32 // Scale: 1000; Units: m; Depth to switch to high setpoint in automatic mode
	MessageIndex              typedef.MessageIndex
	RepeatDiveInterval        uint16               // Units: s; Time between surfacing and ending the activity
	SafetyStopTime            uint16               // Units: s; Time at safety stop (if enabled)
	TravelGas                 typedef.MessageIndex // Index of travel dive_gas message
	Model                     typedef.TissueModelType
	GfLow                     uint8 // Units: percent
	GfHigh                    uint8 // Units: percent
	WaterType                 typedef.WaterType
	Po2Warn                   uint8 // Scale: 100; Units: percent; Typically 1.40
	Po2Critical               uint8 // Scale: 100; Units: percent; Typically 1.60
	Po2Deco                   uint8 // Scale: 100; Units: percent
	SafetyStopEnabled         bool
	ApneaCountdownEnabled     bool
	BacklightMode             typedef.DiveBacklightMode
	BacklightBrightness       uint8
	BacklightTimeout          typedef.BacklightTimeout
	HeartRateSourceType       typedef.SourceType
	HeartRateSource           uint8
	CcrLowSetpointSwitchMode  typedef.CcrSetpointSwitchMode  // If low PO2 should be switched to automatically
	CcrLowSetpoint            uint8                          // Scale: 100; Units: percent; Target PO2 when using low setpoint
	CcrHighSetpointSwitchMode typedef.CcrSetpointSwitchMode  // If high PO2 should be switched to automatically
	CcrHighSetpoint           uint8                          // Scale: 100; Units: percent; Target PO2 when using high setpoint
	GasConsumptionDisplay     typedef.GasConsumptionRateType // Type of gas consumption rate to display. Some values are only valid if tank volume is known.
	UpKeyEnabled              bool                           // Indicates whether the up key is enabled during dives
	DiveSounds                typedef.Tone                   // Sounds and vibration enabled or disabled in-dive
	LastStopMultiple          uint8                          // Scale: 10; Usually 1.0/1.5/2.0 representing 3/4.5/6m or 10/15/20ft
	NoFlyTimeMode             typedef.NoFlyTimeMode          // Indicates which guidelines to use for no-fly surface interval.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDiveSettings creates new DiveSettings struct based on given mesg.
// If mesg is nil, it will return DiveSettings with all fields being set to its corresponding invalid value.
func NewDiveSettings(mesg *proto.Message) *DiveSettings {
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

	return &DiveSettings{
		Timestamp:                 datetime.ToTime(vals[253].Uint32()),
		MessageIndex:              typedef.MessageIndex(vals[254].Uint16()),
		Name:                      vals[0].String(),
		Model:                     typedef.TissueModelType(vals[1].Uint8()),
		GfLow:                     vals[2].Uint8(),
		GfHigh:                    vals[3].Uint8(),
		WaterType:                 typedef.WaterType(vals[4].Uint8()),
		WaterDensity:              vals[5].Float32(),
		Po2Warn:                   vals[6].Uint8(),
		Po2Critical:               vals[7].Uint8(),
		Po2Deco:                   vals[8].Uint8(),
		SafetyStopEnabled:         vals[9].Bool(),
		BottomDepth:               vals[10].Float32(),
		BottomTime:                vals[11].Uint32(),
		ApneaCountdownEnabled:     vals[12].Bool(),
		ApneaCountdownTime:        vals[13].Uint32(),
		BacklightMode:             typedef.DiveBacklightMode(vals[14].Uint8()),
		BacklightBrightness:       vals[15].Uint8(),
		BacklightTimeout:          typedef.BacklightTimeout(vals[16].Uint8()),
		RepeatDiveInterval:        vals[17].Uint16(),
		SafetyStopTime:            vals[18].Uint16(),
		HeartRateSourceType:       typedef.SourceType(vals[19].Uint8()),
		HeartRateSource:           vals[20].Uint8(),
		TravelGas:                 typedef.MessageIndex(vals[21].Uint16()),
		CcrLowSetpointSwitchMode:  typedef.CcrSetpointSwitchMode(vals[22].Uint8()),
		CcrLowSetpoint:            vals[23].Uint8(),
		CcrLowSetpointDepth:       vals[24].Uint32(),
		CcrHighSetpointSwitchMode: typedef.CcrSetpointSwitchMode(vals[25].Uint8()),
		CcrHighSetpoint:           vals[26].Uint8(),
		CcrHighSetpointDepth:      vals[27].Uint32(),
		GasConsumptionDisplay:     typedef.GasConsumptionRateType(vals[29].Uint8()),
		UpKeyEnabled:              vals[30].Bool(),
		DiveSounds:                typedef.Tone(vals[35].Uint8()),
		LastStopMultiple:          vals[36].Uint8(),
		NoFlyTimeMode:             typedef.NoFlyTimeMode(vals[37].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts DiveSettings into proto.Message. If options is nil, default options will be used.
func (m *DiveSettings) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumDiveSettings}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if byte(m.Model) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.Model))
		fields = append(fields, field)
	}
	if m.GfLow != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.GfLow)
		fields = append(fields, field)
	}
	if m.GfHigh != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.GfHigh)
		fields = append(fields, field)
	}
	if byte(m.WaterType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.WaterType))
		fields = append(fields, field)
	}
	if math.Float32bits(m.WaterDensity) != basetype.Float32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Float32(m.WaterDensity)
		fields = append(fields, field)
	}
	if m.Po2Warn != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.Po2Warn)
		fields = append(fields, field)
	}
	if m.Po2Critical != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(m.Po2Critical)
		fields = append(fields, field)
	}
	if m.Po2Deco != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(m.Po2Deco)
		fields = append(fields, field)
	}
	if m.SafetyStopEnabled != false {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Bool(m.SafetyStopEnabled)
		fields = append(fields, field)
	}
	if math.Float32bits(m.BottomDepth) != basetype.Float32Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Float32(m.BottomDepth)
		fields = append(fields, field)
	}
	if m.BottomTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint32(m.BottomTime)
		fields = append(fields, field)
	}
	if m.ApneaCountdownEnabled != false {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Bool(m.ApneaCountdownEnabled)
		fields = append(fields, field)
	}
	if m.ApneaCountdownTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint32(m.ApneaCountdownTime)
		fields = append(fields, field)
	}
	if byte(m.BacklightMode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(byte(m.BacklightMode))
		fields = append(fields, field)
	}
	if m.BacklightBrightness != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint8(m.BacklightBrightness)
		fields = append(fields, field)
	}
	if uint8(m.BacklightTimeout) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = proto.Uint8(uint8(m.BacklightTimeout))
		fields = append(fields, field)
	}
	if m.RepeatDiveInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = proto.Uint16(m.RepeatDiveInterval)
		fields = append(fields, field)
	}
	if m.SafetyStopTime != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = proto.Uint16(m.SafetyStopTime)
		fields = append(fields, field)
	}
	if byte(m.HeartRateSourceType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.Uint8(byte(m.HeartRateSourceType))
		fields = append(fields, field)
	}
	if m.HeartRateSource != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = proto.Uint8(m.HeartRateSource)
		fields = append(fields, field)
	}
	if uint16(m.TravelGas) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint16(uint16(m.TravelGas))
		fields = append(fields, field)
	}
	if byte(m.CcrLowSetpointSwitchMode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint8(byte(m.CcrLowSetpointSwitchMode))
		fields = append(fields, field)
	}
	if m.CcrLowSetpoint != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Uint8(m.CcrLowSetpoint)
		fields = append(fields, field)
	}
	if m.CcrLowSetpointDepth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = proto.Uint32(m.CcrLowSetpointDepth)
		fields = append(fields, field)
	}
	if byte(m.CcrHighSetpointSwitchMode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint8(byte(m.CcrHighSetpointSwitchMode))
		fields = append(fields, field)
	}
	if m.CcrHighSetpoint != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 26)
		field.Value = proto.Uint8(m.CcrHighSetpoint)
		fields = append(fields, field)
	}
	if m.CcrHighSetpointDepth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = proto.Uint32(m.CcrHighSetpointDepth)
		fields = append(fields, field)
	}
	if byte(m.GasConsumptionDisplay) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 29)
		field.Value = proto.Uint8(byte(m.GasConsumptionDisplay))
		fields = append(fields, field)
	}
	if m.UpKeyEnabled != false {
		field := fac.CreateField(mesg.Num, 30)
		field.Value = proto.Bool(m.UpKeyEnabled)
		fields = append(fields, field)
	}
	if byte(m.DiveSounds) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 35)
		field.Value = proto.Uint8(byte(m.DiveSounds))
		fields = append(fields, field)
	}
	if m.LastStopMultiple != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 36)
		field.Value = proto.Uint8(m.LastStopMultiple)
		fields = append(fields, field)
	}
	if byte(m.NoFlyTimeMode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 37)
		field.Value = proto.Uint8(byte(m.NoFlyTimeMode))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetHeartRateSource returns Dynamic Field interpretation of HeartRateSource. Otherwise, returns the original value of HeartRateSource.
//
// Based on m.HeartRateSourceType:
//   - name: "heart_rate_antplus_device_type", value: typedef.AntplusDeviceType(m.HeartRateSource)
//   - name: "heart_rate_local_device_type", value: typedef.LocalDeviceType(m.HeartRateSource)
//
// Otherwise:
//   - name: "heart_rate_source", value: m.HeartRateSource
func (m *DiveSettings) GetHeartRateSource() (name string, value any) {
	switch m.HeartRateSourceType {
	case typedef.SourceTypeAntplus:
		return "heart_rate_antplus_device_type", typedef.AntplusDeviceType(m.HeartRateSource)
	case typedef.SourceTypeLocal:
		return "heart_rate_local_device_type", typedef.LocalDeviceType(m.HeartRateSource)
	}
	return "heart_rate_source", m.HeartRateSource
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *DiveSettings) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// Po2WarnScaled return Po2Warn in its scaled value.
// If Po2Warn value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: percent; Typically 1.40
func (m *DiveSettings) Po2WarnScaled() float64 {
	if m.Po2Warn == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Po2Warn)/100 - 0
}

// Po2CriticalScaled return Po2Critical in its scaled value.
// If Po2Critical value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: percent; Typically 1.60
func (m *DiveSettings) Po2CriticalScaled() float64 {
	if m.Po2Critical == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Po2Critical)/100 - 0
}

// Po2DecoScaled return Po2Deco in its scaled value.
// If Po2Deco value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: percent
func (m *DiveSettings) Po2DecoScaled() float64 {
	if m.Po2Deco == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Po2Deco)/100 - 0
}

// CcrLowSetpointScaled return CcrLowSetpoint in its scaled value.
// If CcrLowSetpoint value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: percent; Target PO2 when using low setpoint
func (m *DiveSettings) CcrLowSetpointScaled() float64 {
	if m.CcrLowSetpoint == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.CcrLowSetpoint)/100 - 0
}

// CcrLowSetpointDepthScaled return CcrLowSetpointDepth in its scaled value.
// If CcrLowSetpointDepth value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m; Depth to switch to low setpoint in automatic mode
func (m *DiveSettings) CcrLowSetpointDepthScaled() float64 {
	if m.CcrLowSetpointDepth == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.CcrLowSetpointDepth)/1000 - 0
}

// CcrHighSetpointScaled return CcrHighSetpoint in its scaled value.
// If CcrHighSetpoint value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: percent; Target PO2 when using high setpoint
func (m *DiveSettings) CcrHighSetpointScaled() float64 {
	if m.CcrHighSetpoint == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.CcrHighSetpoint)/100 - 0
}

// CcrHighSetpointDepthScaled return CcrHighSetpointDepth in its scaled value.
// If CcrHighSetpointDepth value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m; Depth to switch to high setpoint in automatic mode
func (m *DiveSettings) CcrHighSetpointDepthScaled() float64 {
	if m.CcrHighSetpointDepth == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.CcrHighSetpointDepth)/1000 - 0
}

// LastStopMultipleScaled return LastStopMultiple in its scaled value.
// If LastStopMultiple value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Usually 1.0/1.5/2.0 representing 3/4.5/6m or 10/15/20ft
func (m *DiveSettings) LastStopMultipleScaled() float64 {
	if m.LastStopMultiple == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.LastStopMultiple)/10 - 0
}

// SetTimestamp sets Timestamp value.
func (m *DiveSettings) SetTimestamp(v time.Time) *DiveSettings {
	m.Timestamp = v
	return m
}

// SetMessageIndex sets MessageIndex value.
func (m *DiveSettings) SetMessageIndex(v typedef.MessageIndex) *DiveSettings {
	m.MessageIndex = v
	return m
}

// SetName sets Name value.
func (m *DiveSettings) SetName(v string) *DiveSettings {
	m.Name = v
	return m
}

// SetModel sets Model value.
func (m *DiveSettings) SetModel(v typedef.TissueModelType) *DiveSettings {
	m.Model = v
	return m
}

// SetGfLow sets GfLow value.
//
// Units: percent
func (m *DiveSettings) SetGfLow(v uint8) *DiveSettings {
	m.GfLow = v
	return m
}

// SetGfHigh sets GfHigh value.
//
// Units: percent
func (m *DiveSettings) SetGfHigh(v uint8) *DiveSettings {
	m.GfHigh = v
	return m
}

// SetWaterType sets WaterType value.
func (m *DiveSettings) SetWaterType(v typedef.WaterType) *DiveSettings {
	m.WaterType = v
	return m
}

// SetWaterDensity sets WaterDensity value.
//
// Units: kg/m^3; Fresh water is usually 1000; salt water is usually 1025
func (m *DiveSettings) SetWaterDensity(v float32) *DiveSettings {
	m.WaterDensity = v
	return m
}

// SetPo2Warn sets Po2Warn value.
//
// Scale: 100; Units: percent; Typically 1.40
func (m *DiveSettings) SetPo2Warn(v uint8) *DiveSettings {
	m.Po2Warn = v
	return m
}

// SetPo2WarnScaled is similar to SetPo2Warn except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 100; Units: percent; Typically 1.40
func (m *DiveSettings) SetPo2WarnScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.Po2Warn = uint8(basetype.Uint8Invalid)
		return m
	}
	m.Po2Warn = uint8(unscaled)
	return m
}

// SetPo2Critical sets Po2Critical value.
//
// Scale: 100; Units: percent; Typically 1.60
func (m *DiveSettings) SetPo2Critical(v uint8) *DiveSettings {
	m.Po2Critical = v
	return m
}

// SetPo2CriticalScaled is similar to SetPo2Critical except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 100; Units: percent; Typically 1.60
func (m *DiveSettings) SetPo2CriticalScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.Po2Critical = uint8(basetype.Uint8Invalid)
		return m
	}
	m.Po2Critical = uint8(unscaled)
	return m
}

// SetPo2Deco sets Po2Deco value.
//
// Scale: 100; Units: percent
func (m *DiveSettings) SetPo2Deco(v uint8) *DiveSettings {
	m.Po2Deco = v
	return m
}

// SetPo2DecoScaled is similar to SetPo2Deco except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 100; Units: percent
func (m *DiveSettings) SetPo2DecoScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.Po2Deco = uint8(basetype.Uint8Invalid)
		return m
	}
	m.Po2Deco = uint8(unscaled)
	return m
}

// SetSafetyStopEnabled sets SafetyStopEnabled value.
func (m *DiveSettings) SetSafetyStopEnabled(v bool) *DiveSettings {
	m.SafetyStopEnabled = v
	return m
}

// SetBottomDepth sets BottomDepth value.
func (m *DiveSettings) SetBottomDepth(v float32) *DiveSettings {
	m.BottomDepth = v
	return m
}

// SetBottomTime sets BottomTime value.
func (m *DiveSettings) SetBottomTime(v uint32) *DiveSettings {
	m.BottomTime = v
	return m
}

// SetApneaCountdownEnabled sets ApneaCountdownEnabled value.
func (m *DiveSettings) SetApneaCountdownEnabled(v bool) *DiveSettings {
	m.ApneaCountdownEnabled = v
	return m
}

// SetApneaCountdownTime sets ApneaCountdownTime value.
func (m *DiveSettings) SetApneaCountdownTime(v uint32) *DiveSettings {
	m.ApneaCountdownTime = v
	return m
}

// SetBacklightMode sets BacklightMode value.
func (m *DiveSettings) SetBacklightMode(v typedef.DiveBacklightMode) *DiveSettings {
	m.BacklightMode = v
	return m
}

// SetBacklightBrightness sets BacklightBrightness value.
func (m *DiveSettings) SetBacklightBrightness(v uint8) *DiveSettings {
	m.BacklightBrightness = v
	return m
}

// SetBacklightTimeout sets BacklightTimeout value.
func (m *DiveSettings) SetBacklightTimeout(v typedef.BacklightTimeout) *DiveSettings {
	m.BacklightTimeout = v
	return m
}

// SetRepeatDiveInterval sets RepeatDiveInterval value.
//
// Units: s; Time between surfacing and ending the activity
func (m *DiveSettings) SetRepeatDiveInterval(v uint16) *DiveSettings {
	m.RepeatDiveInterval = v
	return m
}

// SetSafetyStopTime sets SafetyStopTime value.
//
// Units: s; Time at safety stop (if enabled)
func (m *DiveSettings) SetSafetyStopTime(v uint16) *DiveSettings {
	m.SafetyStopTime = v
	return m
}

// SetHeartRateSourceType sets HeartRateSourceType value.
func (m *DiveSettings) SetHeartRateSourceType(v typedef.SourceType) *DiveSettings {
	m.HeartRateSourceType = v
	return m
}

// SetHeartRateSource sets HeartRateSource value.
func (m *DiveSettings) SetHeartRateSource(v uint8) *DiveSettings {
	m.HeartRateSource = v
	return m
}

// SetTravelGas sets TravelGas value.
//
// Index of travel dive_gas message
func (m *DiveSettings) SetTravelGas(v typedef.MessageIndex) *DiveSettings {
	m.TravelGas = v
	return m
}

// SetCcrLowSetpointSwitchMode sets CcrLowSetpointSwitchMode value.
//
// If low PO2 should be switched to automatically
func (m *DiveSettings) SetCcrLowSetpointSwitchMode(v typedef.CcrSetpointSwitchMode) *DiveSettings {
	m.CcrLowSetpointSwitchMode = v
	return m
}

// SetCcrLowSetpoint sets CcrLowSetpoint value.
//
// Scale: 100; Units: percent; Target PO2 when using low setpoint
func (m *DiveSettings) SetCcrLowSetpoint(v uint8) *DiveSettings {
	m.CcrLowSetpoint = v
	return m
}

// SetCcrLowSetpointScaled is similar to SetCcrLowSetpoint except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 100; Units: percent; Target PO2 when using low setpoint
func (m *DiveSettings) SetCcrLowSetpointScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.CcrLowSetpoint = uint8(basetype.Uint8Invalid)
		return m
	}
	m.CcrLowSetpoint = uint8(unscaled)
	return m
}

// SetCcrLowSetpointDepth sets CcrLowSetpointDepth value.
//
// Scale: 1000; Units: m; Depth to switch to low setpoint in automatic mode
func (m *DiveSettings) SetCcrLowSetpointDepth(v uint32) *DiveSettings {
	m.CcrLowSetpointDepth = v
	return m
}

// SetCcrLowSetpointDepthScaled is similar to SetCcrLowSetpointDepth except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m; Depth to switch to low setpoint in automatic mode
func (m *DiveSettings) SetCcrLowSetpointDepthScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.CcrLowSetpointDepth = uint32(basetype.Uint32Invalid)
		return m
	}
	m.CcrLowSetpointDepth = uint32(unscaled)
	return m
}

// SetCcrHighSetpointSwitchMode sets CcrHighSetpointSwitchMode value.
//
// If high PO2 should be switched to automatically
func (m *DiveSettings) SetCcrHighSetpointSwitchMode(v typedef.CcrSetpointSwitchMode) *DiveSettings {
	m.CcrHighSetpointSwitchMode = v
	return m
}

// SetCcrHighSetpoint sets CcrHighSetpoint value.
//
// Scale: 100; Units: percent; Target PO2 when using high setpoint
func (m *DiveSettings) SetCcrHighSetpoint(v uint8) *DiveSettings {
	m.CcrHighSetpoint = v
	return m
}

// SetCcrHighSetpointScaled is similar to SetCcrHighSetpoint except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 100; Units: percent; Target PO2 when using high setpoint
func (m *DiveSettings) SetCcrHighSetpointScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.CcrHighSetpoint = uint8(basetype.Uint8Invalid)
		return m
	}
	m.CcrHighSetpoint = uint8(unscaled)
	return m
}

// SetCcrHighSetpointDepth sets CcrHighSetpointDepth value.
//
// Scale: 1000; Units: m; Depth to switch to high setpoint in automatic mode
func (m *DiveSettings) SetCcrHighSetpointDepth(v uint32) *DiveSettings {
	m.CcrHighSetpointDepth = v
	return m
}

// SetCcrHighSetpointDepthScaled is similar to SetCcrHighSetpointDepth except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m; Depth to switch to high setpoint in automatic mode
func (m *DiveSettings) SetCcrHighSetpointDepthScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.CcrHighSetpointDepth = uint32(basetype.Uint32Invalid)
		return m
	}
	m.CcrHighSetpointDepth = uint32(unscaled)
	return m
}

// SetGasConsumptionDisplay sets GasConsumptionDisplay value.
//
// Type of gas consumption rate to display. Some values are only valid if tank volume is known.
func (m *DiveSettings) SetGasConsumptionDisplay(v typedef.GasConsumptionRateType) *DiveSettings {
	m.GasConsumptionDisplay = v
	return m
}

// SetUpKeyEnabled sets UpKeyEnabled value.
//
// Indicates whether the up key is enabled during dives
func (m *DiveSettings) SetUpKeyEnabled(v bool) *DiveSettings {
	m.UpKeyEnabled = v
	return m
}

// SetDiveSounds sets DiveSounds value.
//
// Sounds and vibration enabled or disabled in-dive
func (m *DiveSettings) SetDiveSounds(v typedef.Tone) *DiveSettings {
	m.DiveSounds = v
	return m
}

// SetLastStopMultiple sets LastStopMultiple value.
//
// Scale: 10; Usually 1.0/1.5/2.0 representing 3/4.5/6m or 10/15/20ft
func (m *DiveSettings) SetLastStopMultiple(v uint8) *DiveSettings {
	m.LastStopMultiple = v
	return m
}

// SetLastStopMultipleScaled is similar to SetLastStopMultiple except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 10; Usually 1.0/1.5/2.0 representing 3/4.5/6m or 10/15/20ft
func (m *DiveSettings) SetLastStopMultipleScaled(v float64) *DiveSettings {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.LastStopMultiple = uint8(basetype.Uint8Invalid)
		return m
	}
	m.LastStopMultiple = uint8(unscaled)
	return m
}

// SetNoFlyTimeMode sets NoFlyTimeMode value.
//
// Indicates which guidelines to use for no-fly surface interval.
func (m *DiveSettings) SetNoFlyTimeMode(v typedef.NoFlyTimeMode) *DiveSettings {
	m.NoFlyTimeMode = v
	return m
}

// SetDeveloperFields DiveSettings's DeveloperFields.
func (m *DiveSettings) SetDeveloperFields(developerFields ...proto.DeveloperField) *DiveSettings {
	m.DeveloperFields = developerFields
	return m
}
