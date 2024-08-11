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
	Flags        typedef.FileFlags

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewFileCapabilities creates new FileCapabilities struct based on given mesg.
// If mesg is nil, it will return FileCapabilities with all fields being set to its corresponding invalid value.
func NewFileCapabilities(mesg *proto.Message) *FileCapabilities {
	vals := [255]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &FileCapabilities{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Type:         typedef.File(vals[0].Uint8()),
		Flags:        typedef.FileFlags(vals[1].Uint8z()),
		Directory:    vals[2].String(),
		MaxCount:     vals[3].Uint16(),
		MaxSize:      vals[4].Uint32(),

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

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumFileCapabilities}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if uint8(m.Flags) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(uint8(m.Flags))
		fields = append(fields, field)
	}
	if m.Directory != basetype.StringInvalid && m.Directory != "" {
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

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

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

// SetDeveloperFields FileCapabilities's DeveloperFields.
func (m *FileCapabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *FileCapabilities {
	m.DeveloperFields = developerFields
	return m
}
