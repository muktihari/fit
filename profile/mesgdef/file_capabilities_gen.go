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
)

// FileCapabilities is a FileCapabilities message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type FileCapabilities struct {
	Directory    string
	MaxSize      uint32 // Units: bytes
	MessageIndex typedef.MessageIndex
	MaxCount     uint16
	Type         typedef.File
	Flags        typedef.FileFlags // Base: uint8z

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewFileCapabilities creates new FileCapabilities struct based on given mesg.
// If mesg is nil, it will return FileCapabilities with all fields being set to its corresponding invalid value.
func NewFileCapabilities(mesg *proto.Message) *FileCapabilities {
	m := new(FileCapabilities)
	m.Reset(mesg)
	return m
}

// Reset resets all FileCapabilities's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *FileCapabilities) Reset(mesg *proto.Message) {
	var (
		vals            [255]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

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
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = FileCapabilities{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Type:         typedef.File(vals[0].Uint8()),
		Flags:        typedef.FileFlags(vals[1].Uint8z()),
		Directory:    vals[2].String(),
		MaxCount:     vals[3].Uint16(),
		MaxSize:      vals[4].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts FileCapabilities into proto.Message. If options is nil, default options will be used.
func (m *FileCapabilities) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumFileCapabilities}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Type != typedef.FileInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if m.Flags != typedef.FileFlagsInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(uint8(m.Flags))
		fields = append(fields, field)
	}
	if m.Directory != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.String(m.Directory)
		fields = append(fields, field)
	}
	if m.MaxCount != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.MaxCount)
		fields = append(fields, field)
	}
	if m.MaxSize != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.MaxSize)
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

// SetMessageIndex sets MessageIndex value.
func (m *FileCapabilities) SetMessageIndex(v typedef.MessageIndex) *FileCapabilities {
	m.MessageIndex = v
	return m
}

// SetType sets Type value.
func (m *FileCapabilities) SetType(v typedef.File) *FileCapabilities {
	m.Type = v
	return m
}

// SetFlags sets Flags value.
//
// Base: uint8z
func (m *FileCapabilities) SetFlags(v typedef.FileFlags) *FileCapabilities {
	m.Flags = v
	return m
}

// SetDirectory sets Directory value.
func (m *FileCapabilities) SetDirectory(v string) *FileCapabilities {
	m.Directory = v
	return m
}

// SetMaxCount sets MaxCount value.
func (m *FileCapabilities) SetMaxCount(v uint16) *FileCapabilities {
	m.MaxCount = v
	return m
}

// SetMaxSize sets MaxSize value.
//
// Units: bytes
func (m *FileCapabilities) SetMaxSize(v uint32) *FileCapabilities {
	m.MaxSize = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *FileCapabilities) SetUnknownFields(unknownFields ...proto.Field) *FileCapabilities {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *FileCapabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *FileCapabilities {
	m.DeveloperFields = developerFields
	return m
}
