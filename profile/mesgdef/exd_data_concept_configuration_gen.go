// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

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

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0:  basetype.Uint8Invalid, /* ScreenIndex */
		1:  basetype.ByteInvalid,  /* ConceptField */
		2:  basetype.Uint8Invalid, /* FieldId */
		3:  basetype.Uint8Invalid, /* ConceptIndex */
		4:  basetype.Uint8Invalid, /* DataPage */
		5:  basetype.Uint8Invalid, /* ConceptKey */
		6:  basetype.Uint8Invalid, /* Scaling */
		8:  basetype.EnumInvalid,  /* DataUnits */
		9:  basetype.EnumInvalid,  /* Qualifier */
		10: basetype.EnumInvalid,  /* Descriptor */
		11: false,                 /* IsSigned */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
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

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to ExdDataConceptConfiguration mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumExdDataConceptConfiguration)
func (m ExdDataConceptConfiguration) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumExdDataConceptConfiguration {
		return
	}

	vals := [256]any{
		0:  m.ScreenIndex,
		1:  m.ConceptField,
		2:  m.FieldId,
		3:  m.ConceptIndex,
		4:  m.DataPage,
		5:  m.ConceptKey,
		6:  m.Scaling,
		8:  m.DataUnits,
		9:  m.Qualifier,
		10: m.Descriptor,
		11: m.IsSigned,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
