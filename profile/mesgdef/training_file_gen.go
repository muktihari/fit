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

// TrainingFile is a TrainingFile message.
type TrainingFile struct {
	Timestamp    typedef.DateTime
	Type         typedef.File
	Manufacturer typedef.Manufacturer
	Product      uint16
	SerialNumber uint32
	TimeCreated  typedef.DateTime

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTrainingFile creates new TrainingFile struct based on given mesg. If mesg is nil or mesg.Num is not equal to TrainingFile mesg number, it will return nil.
func NewTrainingFile(mesg proto.Message) *TrainingFile {
	if mesg.Num != typedef.MesgNumTrainingFile {
		return nil
	}

	vals := [254]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &TrainingFile{
		Timestamp:    typeconv.ToUint32[typedef.DateTime](vals[253]),
		Type:         typeconv.ToEnum[typedef.File](vals[0]),
		Manufacturer: typeconv.ToUint16[typedef.Manufacturer](vals[1]),
		Product:      typeconv.ToUint16[uint16](vals[2]),
		SerialNumber: typeconv.ToUint32z[uint32](vals[3]),
		TimeCreated:  typeconv.ToUint32[typedef.DateTime](vals[4]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts TrainingFile into proto.Message.
func (m *TrainingFile) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumTrainingFile)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = typeconv.ToUint32[uint32](m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.Type)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.Manufacturer) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToUint16[uint16](m.Manufacturer)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Product
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32z[uint32](m.SerialNumber) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = typeconv.ToUint32z[uint32](m.SerialNumber)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32[uint32](m.TimeCreated) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = typeconv.ToUint32[uint32](m.TimeCreated)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of TrainingFile's valid fields.
func (m *TrainingFile) size() byte {
	var size byte
	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Type) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.Manufacturer) != basetype.Uint16Invalid {
		size++
	}
	if m.Product != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint32z[uint32](m.SerialNumber) != basetype.Uint32zInvalid {
		size++
	}
	if typeconv.ToUint32[uint32](m.TimeCreated) != basetype.Uint32Invalid {
		size++
	}
	return size
}
