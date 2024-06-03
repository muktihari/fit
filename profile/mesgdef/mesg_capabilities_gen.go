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

// MesgCapabilities is a MesgCapabilities message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type MesgCapabilities struct {
	MessageIndex typedef.MessageIndex
	MesgNum      typedef.MesgNum
	Count        uint16
	File         typedef.File
	CountType    typedef.MesgCount

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMesgCapabilities creates new MesgCapabilities struct based on given mesg.
// If mesg is nil, it will return MesgCapabilities with all fields being set to its corresponding invalid value.
func NewMesgCapabilities(mesg *proto.Message) *MesgCapabilities {
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

	return &MesgCapabilities{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		File:         typedef.File(vals[0].Uint8()),
		MesgNum:      typedef.MesgNum(vals[1].Uint16()),
		CountType:    typedef.MesgCount(vals[2].Uint8()),
		Count:        vals[3].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts MesgCapabilities into proto.Message. If options is nil, default options will be used.
func (m *MesgCapabilities) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumMesgCapabilities}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if byte(m.File) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.File))
		fields = append(fields, field)
	}
	if uint16(m.MesgNum) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.MesgNum))
		fields = append(fields, field)
	}
	if byte(m.CountType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.CountType))
		fields = append(fields, field)
	}
	if m.Count != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.Count)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetCount returns Dynamic Field interpretation of Count. Otherwise, returns the original value of Count.
//
// Based on m.CountType:
//   - name: "num_per_file", value: uint16(m.Count)
//   - name: "max_per_file", value: uint16(m.Count)
//   - name: "max_per_file_type", value: uint16(m.Count)
//
// Otherwise:
//   - name: "count", value: m.Count
func (m *MesgCapabilities) GetCount() (name string, value any) {
	switch m.CountType {
	case typedef.MesgCountNumPerFile:
		return "num_per_file", uint16(m.Count)
	case typedef.MesgCountMaxPerFile:
		return "max_per_file", uint16(m.Count)
	case typedef.MesgCountMaxPerFileType:
		return "max_per_file_type", uint16(m.Count)
	}
	return "count", m.Count
}

// SetMessageIndex sets MessageIndex value.
func (m *MesgCapabilities) SetMessageIndex(v typedef.MessageIndex) *MesgCapabilities {
	m.MessageIndex = v
	return m
}

// SetFile sets File value.
func (m *MesgCapabilities) SetFile(v typedef.File) *MesgCapabilities {
	m.File = v
	return m
}

// SetMesgNum sets MesgNum value.
func (m *MesgCapabilities) SetMesgNum(v typedef.MesgNum) *MesgCapabilities {
	m.MesgNum = v
	return m
}

// SetCountType sets CountType value.
func (m *MesgCapabilities) SetCountType(v typedef.MesgCount) *MesgCapabilities {
	m.CountType = v
	return m
}

// SetCount sets Count value.
func (m *MesgCapabilities) SetCount(v uint16) *MesgCapabilities {
	m.Count = v
	return m
}

// SetDeveloperFields MesgCapabilities's DeveloperFields.
func (m *MesgCapabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *MesgCapabilities {
	m.DeveloperFields = developerFields
	return m
}
