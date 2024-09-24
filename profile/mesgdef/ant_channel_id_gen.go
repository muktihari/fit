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
)

// AntChannelId is a AntChannelId message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type AntChannelId struct {
	DeviceNumber     uint16 // Base: uint16z
	ChannelNumber    uint8
	DeviceType       uint8 // Base: uint8z
	TransmissionType uint8 // Base: uint8z
	DeviceIndex      typedef.DeviceIndex

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewAntChannelId creates new AntChannelId struct based on given mesg.
// If mesg is nil, it will return AntChannelId with all fields being set to its corresponding invalid value.
func NewAntChannelId(mesg *proto.Message) *AntChannelId {
	vals := [5]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 4 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &AntChannelId{
		ChannelNumber:    vals[0].Uint8(),
		DeviceType:       vals[1].Uint8z(),
		DeviceNumber:     vals[2].Uint16z(),
		TransmissionType: vals[3].Uint8z(),
		DeviceIndex:      typedef.DeviceIndex(vals[4].Uint8()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts AntChannelId into proto.Message. If options is nil, default options will be used.
func (m *AntChannelId) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumAntChannelId}

	if m.ChannelNumber != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.ChannelNumber)
		fields = append(fields, field)
	}
	if m.DeviceType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.DeviceType)
		fields = append(fields, field)
	}
	if m.DeviceNumber != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.DeviceNumber)
		fields = append(fields, field)
	}
	if m.TransmissionType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.TransmissionType)
		fields = append(fields, field)
	}
	if m.DeviceIndex != typedef.DeviceIndexInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
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

// SetChannelNumber sets ChannelNumber value.
func (m *AntChannelId) SetChannelNumber(v uint8) *AntChannelId {
	m.ChannelNumber = v
	return m
}

// SetDeviceType sets DeviceType value.
//
// Base: uint8z
func (m *AntChannelId) SetDeviceType(v uint8) *AntChannelId {
	m.DeviceType = v
	return m
}

// SetDeviceNumber sets DeviceNumber value.
//
// Base: uint16z
func (m *AntChannelId) SetDeviceNumber(v uint16) *AntChannelId {
	m.DeviceNumber = v
	return m
}

// SetTransmissionType sets TransmissionType value.
//
// Base: uint8z
func (m *AntChannelId) SetTransmissionType(v uint8) *AntChannelId {
	m.TransmissionType = v
	return m
}

// SetDeviceIndex sets DeviceIndex value.
func (m *AntChannelId) SetDeviceIndex(v typedef.DeviceIndex) *AntChannelId {
	m.DeviceIndex = v
	return m
}

// SetDeveloperFields AntChannelId's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *AntChannelId) SetUnknownFields(unknownFields ...proto.Field) *AntChannelId {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields AntChannelId's DeveloperFields.
func (m *AntChannelId) SetDeveloperFields(developerFields ...proto.DeveloperField) *AntChannelId {
	m.DeveloperFields = developerFields
	return m
}
