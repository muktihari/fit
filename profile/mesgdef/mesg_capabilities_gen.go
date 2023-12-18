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

// MesgCapabilities is a MesgCapabilities message.
type MesgCapabilities struct {
	MessageIndex typedef.MessageIndex
	File         typedef.File
	MesgNum      typedef.MesgNum
	CountType    typedef.MesgCount
	Count        uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMesgCapabilities creates new MesgCapabilities struct based on given mesg. If mesg is nil or mesg.Num is not equal to MesgCapabilities mesg number, it will return nil.
func NewMesgCapabilities(mesg proto.Message) *MesgCapabilities {
	if mesg.Num != typedef.MesgNumMesgCapabilities {
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

	return &MesgCapabilities{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		File:         typeconv.ToEnum[typedef.File](vals[0]),
		MesgNum:      typeconv.ToUint16[typedef.MesgNum](vals[1]),
		CountType:    typeconv.ToEnum[typedef.MesgCount](vals[2]),
		Count:        typeconv.ToUint16[uint16](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts MesgCapabilities into proto.Message.
func (m *MesgCapabilities) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumMesgCapabilities)
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
	if typeconv.ToEnum[byte](m.CountType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = typeconv.ToEnum[byte](m.CountType)
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

// size returns size of MesgCapabilities's valid fields.
func (m *MesgCapabilities) size() byte {
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
	if typeconv.ToEnum[byte](m.CountType) != basetype.EnumInvalid {
		size++
	}
	if m.Count != basetype.Uint16Invalid {
		size++
	}
	return size
}
