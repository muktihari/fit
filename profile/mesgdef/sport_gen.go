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

// Sport is a Sport message.
type Sport struct {
	Name     string
	Sport    typedef.Sport
	SubSport typedef.SubSport

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSport creates new Sport struct based on given mesg.
// If mesg is nil, it will return Sport with all fields being set to its corresponding invalid value.
func NewSport(mesg *proto.Message) *Sport {
	vals := [4]any{}

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

	return &Sport{
		Name:     typeconv.ToString[string](vals[3]),
		Sport:    typeconv.ToEnum[typedef.Sport](vals[0]),
		SubSport: typeconv.ToEnum[typedef.SubSport](vals[1]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Sport into proto.Message.
func (m *Sport) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSport)

	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Name
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.Sport)
		fields = append(fields, field)
	}
	if byte(m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = byte(m.SubSport)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetName sets Sport value.
func (m *Sport) SetName(v string) *Sport {
	m.Name = v
	return m
}

// SetSport sets Sport value.
func (m *Sport) SetSport(v typedef.Sport) *Sport {
	m.Sport = v
	return m
}

// SetSubSport sets Sport value.
func (m *Sport) SetSubSport(v typedef.SubSport) *Sport {
	m.SubSport = v
	return m
}

// SetDeveloperFields Sport's DeveloperFields.
func (m *Sport) SetDeveloperFields(developerFields ...proto.DeveloperField) *Sport {
	m.DeveloperFields = developerFields
	return m
}
