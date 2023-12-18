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

// ExdDataConceptConfiguration is a ExdDataConceptConfiguration message.
type ExdDataConceptConfiguration struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptIndex uint8
	DataPage     uint8
	ConceptKey   uint8
	Scaling      uint8
	DataUnits    typedef.ExdDataUnits
	Qualifier    typedef.ExdQualifiers
	Descriptor   typedef.ExdDescriptors
	IsSigned     bool

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewExdDataConceptConfiguration creates new ExdDataConceptConfiguration struct based on given mesg. If mesg is nil or mesg.Num is not equal to ExdDataConceptConfiguration mesg number, it will return nil.
func NewExdDataConceptConfiguration(mesg proto.Message) *ExdDataConceptConfiguration {
	if mesg.Num != typedef.MesgNumExdDataConceptConfiguration {
		return nil
	}

	vals := [12]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &ExdDataConceptConfiguration{
		ScreenIndex:  typeconv.ToUint8[uint8](vals[0]),
		ConceptField: typeconv.ToByte[byte](vals[1]),
		FieldId:      typeconv.ToUint8[uint8](vals[2]),
		ConceptIndex: typeconv.ToUint8[uint8](vals[3]),
		DataPage:     typeconv.ToUint8[uint8](vals[4]),
		ConceptKey:   typeconv.ToUint8[uint8](vals[5]),
		Scaling:      typeconv.ToUint8[uint8](vals[6]),
		DataUnits:    typeconv.ToEnum[typedef.ExdDataUnits](vals[8]),
		Qualifier:    typeconv.ToEnum[typedef.ExdQualifiers](vals[9]),
		Descriptor:   typeconv.ToEnum[typedef.ExdDescriptors](vals[10]),
		IsSigned:     typeconv.ToBool[bool](vals[11]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts ExdDataConceptConfiguration into proto.Message.
func (m *ExdDataConceptConfiguration) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumExdDataConceptConfiguration)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if m.ScreenIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ScreenIndex
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ConceptField != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.ConceptField
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FieldId != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.FieldId
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ConceptIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.ConceptIndex
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DataPage != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.DataPage
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ConceptKey != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.ConceptKey
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Scaling != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.Scaling
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.DataUnits) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = typeconv.ToEnum[byte](m.DataUnits)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Qualifier) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToEnum[byte](m.Qualifier)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Descriptor) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = typeconv.ToEnum[byte](m.Descriptor)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.IsSigned != false {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.IsSigned
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of ExdDataConceptConfiguration's valid fields.
func (m *ExdDataConceptConfiguration) size() byte {
	var size byte
	if m.ScreenIndex != basetype.Uint8Invalid {
		size++
	}
	if m.ConceptField != basetype.ByteInvalid {
		size++
	}
	if m.FieldId != basetype.Uint8Invalid {
		size++
	}
	if m.ConceptIndex != basetype.Uint8Invalid {
		size++
	}
	if m.DataPage != basetype.Uint8Invalid {
		size++
	}
	if m.ConceptKey != basetype.Uint8Invalid {
		size++
	}
	if m.Scaling != basetype.Uint8Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.DataUnits) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Qualifier) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Descriptor) != basetype.EnumInvalid {
		size++
	}
	if m.IsSigned != false {
		size++
	}
	return size
}
