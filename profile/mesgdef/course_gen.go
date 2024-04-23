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

// Course is a Course message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Course struct {
	Name         string
	Capabilities typedef.CourseCapabilities
	Sport        typedef.Sport
	SubSport     typedef.SubSport

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewCourse creates new Course struct based on given mesg.
// If mesg is nil, it will return Course with all fields being set to its corresponding invalid value.
func NewCourse(mesg *proto.Message) *Course {
	vals := [8]proto.Value{}

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

	return &Course{
		Sport:        typedef.Sport(vals[4].Uint8()),
		Name:         vals[5].String(),
		Capabilities: typedef.CourseCapabilities(vals[6].Uint32z()),
		SubSport:     typedef.SubSport(vals[7].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Course into proto.Message. If options is nil, default options will be used.
func (m *Course) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumCourse}

	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if uint32(m.Capabilities) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(uint32(m.Capabilities))
		fields = append(fields, field)
	}
	if byte(m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.SubSport))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetSport sets Course value.
func (m *Course) SetSport(v typedef.Sport) *Course {
	m.Sport = v
	return m
}

// SetName sets Course value.
func (m *Course) SetName(v string) *Course {
	m.Name = v
	return m
}

// SetCapabilities sets Course value.
func (m *Course) SetCapabilities(v typedef.CourseCapabilities) *Course {
	m.Capabilities = v
	return m
}

// SetSubSport sets Course value.
func (m *Course) SetSubSport(v typedef.SubSport) *Course {
	m.SubSport = v
	return m
}

// SetDeveloperFields Course's DeveloperFields.
func (m *Course) SetDeveloperFields(developerFields ...proto.DeveloperField) *Course {
	m.DeveloperFields = developerFields
	return m
}
