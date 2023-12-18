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

// NewFieldCapabilities creates new FieldCapabilities struct based on given mesg. If mesg is nil or mesg.Num is not equal to FieldCapabilities mesg number, it will return nil.
func NewFieldCapabilities(mesg proto.Message) *FieldCapabilities {
	if mesg.Num != typedef.MesgNumFieldCapabilities {
		return nil
	}

	vals := [255]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &FieldCapabilities{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		File:         typeconv.ToEnum[typedef.File](vals[0]),
		MesgNum:      typeconv.ToUint16[typedef.MesgNum](vals[1]),
		FieldNum:     typeconv.ToUint8[uint8](vals[2]),
		Count:        typeconv.ToUint16[uint16](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts FieldCapabilities into proto.Message.
func (m *FieldCapabilities) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumFieldCapabilities)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.File) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.File)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.MesgNum) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToUint16[uint16](m.MesgNum)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FieldNum != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.FieldNum
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Count != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Count
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of FieldCapabilities's valid fields.
func (m *FieldCapabilities) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.File) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.MesgNum) != basetype.Uint16Invalid {
		size++
	}
	if m.FieldNum != basetype.Uint8Invalid {
		size++
	}
	if m.Count != basetype.Uint16Invalid {
		size++
	}
	return size
}
