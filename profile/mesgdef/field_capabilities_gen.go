// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// FieldCapabilities is a FieldCapabilities message.
type FieldCapabilities struct {
	MessageIndex typedef.MessageIndex
	File         typedef.File
	MesgNum      typedef.MesgNum
	FieldNum     uint8
	Count        uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewFieldCapabilities creates new FieldCapabilities struct based on given mesg.
// If mesg is nil, it will return FieldCapabilities with all fields being set to its corresponding invalid value.
func NewFieldCapabilities(mesg *proto.Message) *FieldCapabilities {
	vals := [255]any{}

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

	return &FieldCapabilities{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		File:         typeconv.ToEnum[typedef.File](vals[0]),
		MesgNum:      typeconv.ToUint16[typedef.MesgNum](vals[1]),
		FieldNum:     typeconv.ToUint8[uint8](vals[2]),
		Count:        typeconv.ToUint16[uint16](vals[3]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts FieldCapabilities into proto.Message.
func (m *FieldCapabilities) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumFieldCapabilities)

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if byte(m.File) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.File)
		fields = append(fields, field)
	}
	if uint16(m.MesgNum) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = uint16(m.MesgNum)
		fields = append(fields, field)
	}
	if m.FieldNum != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.FieldNum
		fields = append(fields, field)
	}
	if m.Count != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Count
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets FieldCapabilities value.
func (m *FieldCapabilities) SetMessageIndex(v typedef.MessageIndex) *FieldCapabilities {
	m.MessageIndex = v
	return m
}

// SetFile sets FieldCapabilities value.
func (m *FieldCapabilities) SetFile(v typedef.File) *FieldCapabilities {
	m.File = v
	return m
}

// SetMesgNum sets FieldCapabilities value.
func (m *FieldCapabilities) SetMesgNum(v typedef.MesgNum) *FieldCapabilities {
	m.MesgNum = v
	return m
}

// SetFieldNum sets FieldCapabilities value.
func (m *FieldCapabilities) SetFieldNum(v uint8) *FieldCapabilities {
	m.FieldNum = v
	return m
}

// SetCount sets FieldCapabilities value.
func (m *FieldCapabilities) SetCount(v uint16) *FieldCapabilities {
	m.Count = v
	return m
}

// SetDeveloperFields FieldCapabilities's DeveloperFields.
func (m *FieldCapabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *FieldCapabilities {
	m.DeveloperFields = developerFields
	return m
}
