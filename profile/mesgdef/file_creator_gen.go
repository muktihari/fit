// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// FileCreator is a FileCreator message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type FileCreator struct {
	SoftwareVersion uint16
	HardwareVersion uint8

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewFileCreator creates new FileCreator struct based on given mesg.
// If mesg is nil, it will return FileCreator with all fields being set to its corresponding invalid value.
func NewFileCreator(mesg *proto.Message) *FileCreator {
	m := new(FileCreator)
	m.Reset(mesg)
	return m
}

// Reset resets all FileCreator's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *FileCreator) Reset(mesg *proto.Message) {
	var (
		vals            [2]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 1 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = FileCreator{
		SoftwareVersion: vals[0].Uint16(),
		HardwareVersion: vals[1].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts FileCreator into proto.Message. If options is nil, default options will be used.
func (m *FileCreator) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumFileCreator}

	if m.SoftwareVersion != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.SoftwareVersion)
		fields = append(fields, field)
	}
	if m.HardwareVersion != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.HardwareVersion)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetSoftwareVersion sets SoftwareVersion value.
func (m *FileCreator) SetSoftwareVersion(v uint16) *FileCreator {
	m.SoftwareVersion = v
	return m
}

// SetHardwareVersion sets HardwareVersion value.
func (m *FileCreator) SetHardwareVersion(v uint8) *FileCreator {
	m.HardwareVersion = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *FileCreator) SetUnknownFields(unknownFields ...proto.Field) *FileCreator {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *FileCreator) SetDeveloperFields(developerFields ...proto.DeveloperField) *FileCreator {
	m.DeveloperFields = developerFields
	return m
}
