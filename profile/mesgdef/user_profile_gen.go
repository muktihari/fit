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

// UserProfile is a UserProfile message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type UserProfile struct {
	FriendlyName               string                   // Used for Morning Report greeting
	WakeTime                   typedef.LocaltimeIntoDay // Typical wake time
	SleepTime                  typedef.LocaltimeIntoDay // Typical bed time
	DiveCount                  uint32
	GlobalId                   [6]byte
	MessageIndex               typedef.MessageIndex
	Weight                     uint16 // Scale: 10; Units: kg
	LocalId                    typedef.UserLocalId
	UserRunningStepLength      uint16 // Scale: 1000; Units: m; User defined running step length set to 0 for auto length
	UserWalkingStepLength      uint16 // Scale: 1000; Units: m; User defined walking step length set to 0 for auto length
	Gender                     typedef.Gender
	Age                        uint8 // Units: years
	Height                     uint8 // Scale: 100; Units: m
	Language                   typedef.Language
	ElevSetting                typedef.DisplayMeasure
	WeightSetting              typedef.DisplayMeasure
	RestingHeartRate           uint8 // Units: bpm
	DefaultMaxRunningHeartRate uint8 // Units: bpm
	DefaultMaxBikingHeartRate  uint8 // Units: bpm
	DefaultMaxHeartRate        uint8 // Units: bpm
	HrSetting                  typedef.DisplayHeart
	SpeedSetting               typedef.DisplayMeasure
	DistSetting                typedef.DisplayMeasure
	PowerSetting               typedef.DisplayPower
	ActivityClass              typedef.ActivityClass
	PositionSetting            typedef.DisplayPosition
	TemperatureSetting         typedef.DisplayMeasure
	HeightSetting              typedef.DisplayMeasure
	DepthSetting               typedef.DisplayMeasure

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewUserProfile creates new UserProfile struct based on given mesg.
// If mesg is nil, it will return UserProfile with all fields being set to its corresponding invalid value.
func NewUserProfile(mesg *proto.Message) *UserProfile {
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

	return &UserProfile{
		MessageIndex:               typedef.MessageIndex(vals[254].Uint16()),
		FriendlyName:               vals[0].String(),
		Gender:                     typedef.Gender(vals[1].Uint8()),
		Age:                        vals[2].Uint8(),
		Height:                     vals[3].Uint8(),
		Weight:                     vals[4].Uint16(),
		Language:                   typedef.Language(vals[5].Uint8()),
		ElevSetting:                typedef.DisplayMeasure(vals[6].Uint8()),
		WeightSetting:              typedef.DisplayMeasure(vals[7].Uint8()),
		RestingHeartRate:           vals[8].Uint8(),
		DefaultMaxRunningHeartRate: vals[9].Uint8(),
		DefaultMaxBikingHeartRate:  vals[10].Uint8(),
		DefaultMaxHeartRate:        vals[11].Uint8(),
		HrSetting:                  typedef.DisplayHeart(vals[12].Uint8()),
		SpeedSetting:               typedef.DisplayMeasure(vals[13].Uint8()),
		DistSetting:                typedef.DisplayMeasure(vals[14].Uint8()),
		PowerSetting:               typedef.DisplayPower(vals[16].Uint8()),
		ActivityClass:              typedef.ActivityClass(vals[17].Uint8()),
		PositionSetting:            typedef.DisplayPosition(vals[18].Uint8()),
		TemperatureSetting:         typedef.DisplayMeasure(vals[21].Uint8()),
		LocalId:                    typedef.UserLocalId(vals[22].Uint16()),
		GlobalId: func() (arr [6]uint8) {
			arr = [6]uint8{
				basetype.ByteInvalid,
				basetype.ByteInvalid,
				basetype.ByteInvalid,
				basetype.ByteInvalid,
				basetype.ByteInvalid,
				basetype.ByteInvalid,
			}
			copy(arr[:], vals[23].SliceUint8())
			return arr
		}(),
		WakeTime:              typedef.LocaltimeIntoDay(vals[28].Uint32()),
		SleepTime:             typedef.LocaltimeIntoDay(vals[29].Uint32()),
		HeightSetting:         typedef.DisplayMeasure(vals[30].Uint8()),
		UserRunningStepLength: vals[31].Uint16(),
		UserWalkingStepLength: vals[32].Uint16(),
		DepthSetting:          typedef.DisplayMeasure(vals[47].Uint8()),
		DiveCount:             vals[49].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts UserProfile into proto.Message. If options is nil, default options will be used.
func (m *UserProfile) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumUserProfile}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.FriendlyName != basetype.StringInvalid && m.FriendlyName != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.FriendlyName)
		fields = append(fields, field)
	}
	if byte(m.Gender) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.Gender))
		fields = append(fields, field)
	}
	if m.Age != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.Age)
		fields = append(fields, field)
	}
	if m.Height != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.Height)
		fields = append(fields, field)
	}
	if m.Weight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.Weight)
		fields = append(fields, field)
	}
	if byte(m.Language) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(byte(m.Language))
		fields = append(fields, field)
	}
	if byte(m.ElevSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(byte(m.ElevSetting))
		fields = append(fields, field)
	}
	if byte(m.WeightSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.WeightSetting))
		fields = append(fields, field)
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(m.RestingHeartRate)
		fields = append(fields, field)
	}
	if m.DefaultMaxRunningHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(m.DefaultMaxRunningHeartRate)
		fields = append(fields, field)
	}
	if m.DefaultMaxBikingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(m.DefaultMaxBikingHeartRate)
		fields = append(fields, field)
	}
	if m.DefaultMaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.DefaultMaxHeartRate)
		fields = append(fields, field)
	}
	if byte(m.HrSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(byte(m.HrSetting))
		fields = append(fields, field)
	}
	if byte(m.SpeedSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint8(byte(m.SpeedSetting))
		fields = append(fields, field)
	}
	if byte(m.DistSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(byte(m.DistSetting))
		fields = append(fields, field)
	}
	if byte(m.PowerSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = proto.Uint8(byte(m.PowerSetting))
		fields = append(fields, field)
	}
	if byte(m.ActivityClass) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = proto.Uint8(byte(m.ActivityClass))
		fields = append(fields, field)
	}
	if byte(m.PositionSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = proto.Uint8(byte(m.PositionSetting))
		fields = append(fields, field)
	}
	if byte(m.TemperatureSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint8(byte(m.TemperatureSetting))
		fields = append(fields, field)
	}
	if uint16(m.LocalId) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint16(uint16(m.LocalId))
		fields = append(fields, field)
	}
	if m.GlobalId != [6]uint8{
		basetype.ByteInvalid,
		basetype.ByteInvalid,
		basetype.ByteInvalid,
		basetype.ByteInvalid,
		basetype.ByteInvalid,
		basetype.ByteInvalid,
	} {
		field := fac.CreateField(mesg.Num, 23)
		copied := m.GlobalId
		field.Value = proto.SliceUint8(copied[:])
		fields = append(fields, field)
	}
	if uint32(m.WakeTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 28)
		field.Value = proto.Uint32(uint32(m.WakeTime))
		fields = append(fields, field)
	}
	if uint32(m.SleepTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 29)
		field.Value = proto.Uint32(uint32(m.SleepTime))
		fields = append(fields, field)
	}
	if byte(m.HeightSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 30)
		field.Value = proto.Uint8(byte(m.HeightSetting))
		fields = append(fields, field)
	}
	if m.UserRunningStepLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 31)
		field.Value = proto.Uint16(m.UserRunningStepLength)
		fields = append(fields, field)
	}
	if m.UserWalkingStepLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = proto.Uint16(m.UserWalkingStepLength)
		fields = append(fields, field)
	}
	if byte(m.DepthSetting) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 47)
		field.Value = proto.Uint8(byte(m.DepthSetting))
		fields = append(fields, field)
	}
	if m.DiveCount != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 49)
		field.Value = proto.Uint32(m.DiveCount)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// HeightScaled return Height in its scaled value.
// If Height value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *UserProfile) HeightScaled() float64 {
	if m.Height == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Height)/100 - 0
}

// WeightScaled return Weight in its scaled value.
// If Weight value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: kg
func (m *UserProfile) WeightScaled() float64 {
	if m.Weight == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Weight)/10 - 0
}

// UserRunningStepLengthScaled return UserRunningStepLength in its scaled value.
// If UserRunningStepLength value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m; User defined running step length set to 0 for auto length
func (m *UserProfile) UserRunningStepLengthScaled() float64 {
	if m.UserRunningStepLength == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.UserRunningStepLength)/1000 - 0
}

// UserWalkingStepLengthScaled return UserWalkingStepLength in its scaled value.
// If UserWalkingStepLength value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m; User defined walking step length set to 0 for auto length
func (m *UserProfile) UserWalkingStepLengthScaled() float64 {
	if m.UserWalkingStepLength == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.UserWalkingStepLength)/1000 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *UserProfile) SetMessageIndex(v typedef.MessageIndex) *UserProfile {
	m.MessageIndex = v
	return m
}

// SetFriendlyName sets FriendlyName value.
//
// Used for Morning Report greeting
func (m *UserProfile) SetFriendlyName(v string) *UserProfile {
	m.FriendlyName = v
	return m
}

// SetGender sets Gender value.
func (m *UserProfile) SetGender(v typedef.Gender) *UserProfile {
	m.Gender = v
	return m
}

// SetAge sets Age value.
//
// Units: years
func (m *UserProfile) SetAge(v uint8) *UserProfile {
	m.Age = v
	return m
}

// SetHeight sets Height value.
//
// Scale: 100; Units: m
func (m *UserProfile) SetHeight(v uint8) *UserProfile {
	m.Height = v
	return m
}

// SetHeightScaled is similar to SetHeight except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *UserProfile) SetHeightScaled(v float64) *UserProfile {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.Height = uint8(basetype.Uint8Invalid)
		return m
	}
	m.Height = uint8(unscaled)
	return m
}

// SetWeight sets Weight value.
//
// Scale: 10; Units: kg
func (m *UserProfile) SetWeight(v uint16) *UserProfile {
	m.Weight = v
	return m
}

// SetWeightScaled is similar to SetWeight except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: kg
func (m *UserProfile) SetWeightScaled(v float64) *UserProfile {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Weight = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Weight = uint16(unscaled)
	return m
}

// SetLanguage sets Language value.
func (m *UserProfile) SetLanguage(v typedef.Language) *UserProfile {
	m.Language = v
	return m
}

// SetElevSetting sets ElevSetting value.
func (m *UserProfile) SetElevSetting(v typedef.DisplayMeasure) *UserProfile {
	m.ElevSetting = v
	return m
}

// SetWeightSetting sets WeightSetting value.
func (m *UserProfile) SetWeightSetting(v typedef.DisplayMeasure) *UserProfile {
	m.WeightSetting = v
	return m
}

// SetRestingHeartRate sets RestingHeartRate value.
//
// Units: bpm
func (m *UserProfile) SetRestingHeartRate(v uint8) *UserProfile {
	m.RestingHeartRate = v
	return m
}

// SetDefaultMaxRunningHeartRate sets DefaultMaxRunningHeartRate value.
//
// Units: bpm
func (m *UserProfile) SetDefaultMaxRunningHeartRate(v uint8) *UserProfile {
	m.DefaultMaxRunningHeartRate = v
	return m
}

// SetDefaultMaxBikingHeartRate sets DefaultMaxBikingHeartRate value.
//
// Units: bpm
func (m *UserProfile) SetDefaultMaxBikingHeartRate(v uint8) *UserProfile {
	m.DefaultMaxBikingHeartRate = v
	return m
}

// SetDefaultMaxHeartRate sets DefaultMaxHeartRate value.
//
// Units: bpm
func (m *UserProfile) SetDefaultMaxHeartRate(v uint8) *UserProfile {
	m.DefaultMaxHeartRate = v
	return m
}

// SetHrSetting sets HrSetting value.
func (m *UserProfile) SetHrSetting(v typedef.DisplayHeart) *UserProfile {
	m.HrSetting = v
	return m
}

// SetSpeedSetting sets SpeedSetting value.
func (m *UserProfile) SetSpeedSetting(v typedef.DisplayMeasure) *UserProfile {
	m.SpeedSetting = v
	return m
}

// SetDistSetting sets DistSetting value.
func (m *UserProfile) SetDistSetting(v typedef.DisplayMeasure) *UserProfile {
	m.DistSetting = v
	return m
}

// SetPowerSetting sets PowerSetting value.
func (m *UserProfile) SetPowerSetting(v typedef.DisplayPower) *UserProfile {
	m.PowerSetting = v
	return m
}

// SetActivityClass sets ActivityClass value.
func (m *UserProfile) SetActivityClass(v typedef.ActivityClass) *UserProfile {
	m.ActivityClass = v
	return m
}

// SetPositionSetting sets PositionSetting value.
func (m *UserProfile) SetPositionSetting(v typedef.DisplayPosition) *UserProfile {
	m.PositionSetting = v
	return m
}

// SetTemperatureSetting sets TemperatureSetting value.
func (m *UserProfile) SetTemperatureSetting(v typedef.DisplayMeasure) *UserProfile {
	m.TemperatureSetting = v
	return m
}

// SetLocalId sets LocalId value.
func (m *UserProfile) SetLocalId(v typedef.UserLocalId) *UserProfile {
	m.LocalId = v
	return m
}

// SetGlobalId sets GlobalId value.
func (m *UserProfile) SetGlobalId(v [6]byte) *UserProfile {
	m.GlobalId = v
	return m
}

// SetWakeTime sets WakeTime value.
//
// Typical wake time
func (m *UserProfile) SetWakeTime(v typedef.LocaltimeIntoDay) *UserProfile {
	m.WakeTime = v
	return m
}

// SetSleepTime sets SleepTime value.
//
// Typical bed time
func (m *UserProfile) SetSleepTime(v typedef.LocaltimeIntoDay) *UserProfile {
	m.SleepTime = v
	return m
}

// SetHeightSetting sets HeightSetting value.
func (m *UserProfile) SetHeightSetting(v typedef.DisplayMeasure) *UserProfile {
	m.HeightSetting = v
	return m
}

// SetUserRunningStepLength sets UserRunningStepLength value.
//
// Scale: 1000; Units: m; User defined running step length set to 0 for auto length
func (m *UserProfile) SetUserRunningStepLength(v uint16) *UserProfile {
	m.UserRunningStepLength = v
	return m
}

// SetUserRunningStepLengthScaled is similar to SetUserRunningStepLength except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m; User defined running step length set to 0 for auto length
func (m *UserProfile) SetUserRunningStepLengthScaled(v float64) *UserProfile {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.UserRunningStepLength = uint16(basetype.Uint16Invalid)
		return m
	}
	m.UserRunningStepLength = uint16(unscaled)
	return m
}

// SetUserWalkingStepLength sets UserWalkingStepLength value.
//
// Scale: 1000; Units: m; User defined walking step length set to 0 for auto length
func (m *UserProfile) SetUserWalkingStepLength(v uint16) *UserProfile {
	m.UserWalkingStepLength = v
	return m
}

// SetUserWalkingStepLengthScaled is similar to SetUserWalkingStepLength except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m; User defined walking step length set to 0 for auto length
func (m *UserProfile) SetUserWalkingStepLengthScaled(v float64) *UserProfile {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.UserWalkingStepLength = uint16(basetype.Uint16Invalid)
		return m
	}
	m.UserWalkingStepLength = uint16(unscaled)
	return m
}

// SetDepthSetting sets DepthSetting value.
func (m *UserProfile) SetDepthSetting(v typedef.DisplayMeasure) *UserProfile {
	m.DepthSetting = v
	return m
}

// SetDiveCount sets DiveCount value.
func (m *UserProfile) SetDiveCount(v uint32) *UserProfile {
	m.DiveCount = v
	return m
}

// SetDeveloperFields UserProfile's DeveloperFields.
func (m *UserProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *UserProfile {
	m.DeveloperFields = developerFields
	return m
}
