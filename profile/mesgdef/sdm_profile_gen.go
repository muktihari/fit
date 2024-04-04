// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SdmProfile is a SdmProfile message.
type SdmProfile struct {
	Odometer          uint32 // Scale: 100; Units: m
	MessageIndex      typedef.MessageIndex
	SdmAntId          uint16
	SdmCalFactor      uint16 // Scale: 10; Units: %
	SdmAntIdTransType uint8
	OdometerRollover  uint8 // Rollover counter that can be used to extend the odometer
	Enabled           bool
	SpeedSource       bool // Use footpod for speed source instead of GPS

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSdmProfile creates new SdmProfile struct based on given mesg.
// If mesg is nil, it will return SdmProfile with all fields being set to its corresponding invalid value.
func NewSdmProfile(mesg *proto.Message) *SdmProfile {
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

	return &SdmProfile{
		Odometer:          vals[3].Uint32(),
		MessageIndex:      typedef.MessageIndex(vals[254].Uint16()),
		SdmAntId:          vals[1].Uint16z(),
		SdmCalFactor:      vals[2].Uint16(),
		SdmAntIdTransType: vals[5].Uint8z(),
		OdometerRollover:  vals[7].Uint8(),
		Enabled:           vals[0].Bool(),
		SpeedSource:       vals[4].Bool(),

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

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumSdmProfile}

	if m.Odometer != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.Odometer)
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if uint16(m.SdmAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.SdmAntId)
		fields = append(fields, field)
	}
	if m.SdmCalFactor != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.SdmCalFactor)
		fields = append(fields, field)
	}
	if uint8(m.SdmAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(m.SdmAntIdTransType)
		fields = append(fields, field)
	}
	if m.OdometerRollover != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(m.OdometerRollover)
		fields = append(fields, field)
	}
	if m.Enabled != false {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if m.SpeedSource != false {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Bool(m.SpeedSource)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// OdometerScaled return Odometer in its scaled value [Scale: 100; Units: m].
//
// If Odometer value is invalid, float64 invalid value will be returned.
func (m *SdmProfile) OdometerScaled() float64 {
	if m.Odometer == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.Odometer, 100, 0)
}

// SdmCalFactorScaled return SdmCalFactor in its scaled value [Scale: 10; Units: %].
//
// If SdmCalFactor value is invalid, float64 invalid value will be returned.
func (m *SdmProfile) SdmCalFactorScaled() float64 {
	if m.SdmCalFactor == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.SdmCalFactor, 10, 0)
}

// SetOdometer sets SdmProfile value.
//
// Scale: 100; Units: m
func (m *SdmProfile) SetOdometer(v uint32) *SdmProfile {
	m.Odometer = v
	return m
}

// SetMessageIndex sets SdmProfile value.
func (m *SdmProfile) SetMessageIndex(v typedef.MessageIndex) *SdmProfile {
	m.MessageIndex = v
	return m
}

// SetSdmAntId sets SdmProfile value.
func (m *SdmProfile) SetSdmAntId(v uint16) *SdmProfile {
	m.SdmAntId = v
	return m
}

// SetSdmCalFactor sets SdmProfile value.
//
// Scale: 10; Units: %
func (m *SdmProfile) SetSdmCalFactor(v uint16) *SdmProfile {
	m.SdmCalFactor = v
	return m
}

// SetSdmAntIdTransType sets SdmProfile value.
func (m *SdmProfile) SetSdmAntIdTransType(v uint8) *SdmProfile {
	m.SdmAntIdTransType = v
	return m
}

// SetOdometerRollover sets SdmProfile value.
//
// Rollover counter that can be used to extend the odometer
func (m *SdmProfile) SetOdometerRollover(v uint8) *SdmProfile {
	m.OdometerRollover = v
	return m
}

// SetEnabled sets SdmProfile value.
func (m *SdmProfile) SetEnabled(v bool) *SdmProfile {
	m.Enabled = v
	return m
}

// SetSpeedSource sets SdmProfile value.
//
// Use footpod for speed source instead of GPS
func (m *SdmProfile) SetSpeedSource(v bool) *SdmProfile {
	m.SpeedSource = v
	return m
}

// SetDeveloperFields SdmProfile's DeveloperFields.
func (m *SdmProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *SdmProfile {
	m.DeveloperFields = developerFields
	return m
}
