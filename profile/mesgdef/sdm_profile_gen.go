// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// SdmProfile is a SdmProfile message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SdmProfile struct {
	Odometer          uint32 // Scale: 100; Units: m
	MessageIndex      typedef.MessageIndex
	SdmAntId          uint16 // Base: uint16z
	SdmCalFactor      uint16 // Scale: 10; Units: %
	Enabled           typedef.Bool
	SpeedSource       typedef.Bool // Use footpod for speed source instead of GPS
	SdmAntIdTransType uint8        // Base: uint8z
	OdometerRollover  uint8        // Rollover counter that can be used to extend the odometer

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSdmProfile creates new SdmProfile struct based on given mesg.
// If mesg is nil, it will return SdmProfile with all fields being set to its corresponding invalid value.
func NewSdmProfile(mesg *proto.Message) *SdmProfile {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
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

	return &SdmProfile{
		MessageIndex:      typedef.MessageIndex(vals[254].Uint16()),
		Enabled:           vals[0].Bool(),
		SdmAntId:          vals[1].Uint16z(),
		SdmCalFactor:      vals[2].Uint16(),
		Odometer:          vals[3].Uint32(),
		SpeedSource:       vals[4].Bool(),
		SdmAntIdTransType: vals[5].Uint8z(),
		OdometerRollover:  vals[7].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SdmProfile into proto.Message. If options is nil, default options will be used.
func (m *SdmProfile) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSdmProfile}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Enabled < 2 {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if m.SdmAntId != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.SdmAntId)
		fields = append(fields, field)
	}
	if m.SdmCalFactor != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.SdmCalFactor)
		fields = append(fields, field)
	}
	if m.Odometer != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.Odometer)
		fields = append(fields, field)
	}
	if m.SpeedSource < 2 {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Bool(m.SpeedSource)
		fields = append(fields, field)
	}
	if m.SdmAntIdTransType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(m.SdmAntIdTransType)
		fields = append(fields, field)
	}
	if m.OdometerRollover != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(m.OdometerRollover)
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

// SdmCalFactorScaled return SdmCalFactor in its scaled value.
// If SdmCalFactor value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: %
func (m *SdmProfile) SdmCalFactorScaled() float64 {
	if m.SdmCalFactor == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.SdmCalFactor)/10 - 0
}

// OdometerScaled return Odometer in its scaled value.
// If Odometer value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *SdmProfile) OdometerScaled() float64 {
	if m.Odometer == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Odometer)/100 - 0
}

// SetMessageIndex sets MessageIndex value.
func (m *SdmProfile) SetMessageIndex(v typedef.MessageIndex) *SdmProfile {
	m.MessageIndex = v
	return m
}

// SetEnabled sets Enabled value.
func (m *SdmProfile) SetEnabled(v typedef.Bool) *SdmProfile {
	m.Enabled = v
	return m
}

// SetSdmAntId sets SdmAntId value.
//
// Base: uint16z
func (m *SdmProfile) SetSdmAntId(v uint16) *SdmProfile {
	m.SdmAntId = v
	return m
}

// SetSdmCalFactor sets SdmCalFactor value.
//
// Scale: 10; Units: %
func (m *SdmProfile) SetSdmCalFactor(v uint16) *SdmProfile {
	m.SdmCalFactor = v
	return m
}

// SetSdmCalFactorScaled is similar to SetSdmCalFactor except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: %
func (m *SdmProfile) SetSdmCalFactorScaled(v float64) *SdmProfile {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.SdmCalFactor = uint16(basetype.Uint16Invalid)
		return m
	}
	m.SdmCalFactor = uint16(unscaled)
	return m
}

// SetOdometer sets Odometer value.
//
// Scale: 100; Units: m
func (m *SdmProfile) SetOdometer(v uint32) *SdmProfile {
	m.Odometer = v
	return m
}

// SetOdometerScaled is similar to SetOdometer except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *SdmProfile) SetOdometerScaled(v float64) *SdmProfile {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Odometer = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Odometer = uint32(unscaled)
	return m
}

// SetSpeedSource sets SpeedSource value.
//
// Use footpod for speed source instead of GPS
func (m *SdmProfile) SetSpeedSource(v typedef.Bool) *SdmProfile {
	m.SpeedSource = v
	return m
}

// SetSdmAntIdTransType sets SdmAntIdTransType value.
//
// Base: uint8z
func (m *SdmProfile) SetSdmAntIdTransType(v uint8) *SdmProfile {
	m.SdmAntIdTransType = v
	return m
}

// SetOdometerRollover sets OdometerRollover value.
//
// Rollover counter that can be used to extend the odometer
func (m *SdmProfile) SetOdometerRollover(v uint8) *SdmProfile {
	m.OdometerRollover = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SdmProfile) SetUnknownFields(unknownFields ...proto.Field) *SdmProfile {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *SdmProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *SdmProfile {
	m.DeveloperFields = developerFields
	return m
}
