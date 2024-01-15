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

// DeveloperDataId is a DeveloperDataId message.
type DeveloperDataId struct {
	DeveloperId        []byte // Array: [N]
	ApplicationId      []byte // Array: [N]
	ApplicationVersion uint32
	ManufacturerId     typedef.Manufacturer
	DeveloperDataIndex uint8
}

// NewDeveloperDataId creates new DeveloperDataId struct based on given mesg.
// If mesg is nil, it will return DeveloperDataId with all fields being set to its corresponding invalid value.
func NewDeveloperDataId(mesg *proto.Message) *DeveloperDataId {
	vals := [5]any{}

	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
	}

	return &DeveloperDataId{
		DeveloperId:        typeconv.ToSliceByte[byte](vals[0]),
		ApplicationId:      typeconv.ToSliceByte[byte](vals[1]),
		ApplicationVersion: typeconv.ToUint32[uint32](vals[4]),
		ManufacturerId:     typeconv.ToUint16[typedef.Manufacturer](vals[2]),
		DeveloperDataIndex: typeconv.ToUint8[uint8](vals[3]),
	}
}

// ToMesg converts DeveloperDataId into proto.Message.
func (m *DeveloperDataId) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumDeveloperDataId)

	if m.DeveloperId != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.DeveloperId
		fields = append(fields, field)
	}
	if m.ApplicationId != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.ApplicationId
		fields = append(fields, field)
	}
	if m.ApplicationVersion != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.ApplicationVersion
		fields = append(fields, field)
	}
	if uint16(m.ManufacturerId) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = uint16(m.ManufacturerId)
		fields = append(fields, field)
	}
	if m.DeveloperDataIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.DeveloperDataIndex
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	return mesg
}

// SetDeveloperId sets DeveloperDataId value.
//
// Array: [N]
func (m *DeveloperDataId) SetDeveloperId(v []byte) *DeveloperDataId {
	m.DeveloperId = v
	return m
}

// SetApplicationId sets DeveloperDataId value.
//
// Array: [N]
func (m *DeveloperDataId) SetApplicationId(v []byte) *DeveloperDataId {
	m.ApplicationId = v
	return m
}

// SetApplicationVersion sets DeveloperDataId value.
func (m *DeveloperDataId) SetApplicationVersion(v uint32) *DeveloperDataId {
	m.ApplicationVersion = v
	return m
}

// SetManufacturerId sets DeveloperDataId value.
func (m *DeveloperDataId) SetManufacturerId(v typedef.Manufacturer) *DeveloperDataId {
	m.ManufacturerId = v
	return m
}

// SetDeveloperDataIndex sets DeveloperDataId value.
func (m *DeveloperDataId) SetDeveloperDataIndex(v uint8) *DeveloperDataId {
	m.DeveloperDataIndex = v
	return m
}
