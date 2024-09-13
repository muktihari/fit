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

// ExdScreenConfiguration is a ExdScreenConfiguration message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ExdScreenConfiguration struct {
	ScreenIndex   uint8
	FieldCount    uint8 // number of fields in screen
	Layout        typedef.ExdLayout
	ScreenEnabled typedef.Bool

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewExdScreenConfiguration creates new ExdScreenConfiguration struct based on given mesg.
// If mesg is nil, it will return ExdScreenConfiguration with all fields being set to its corresponding invalid value.
func NewExdScreenConfiguration(mesg *proto.Message) *ExdScreenConfiguration {
	vals := [4]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 3 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &ExdScreenConfiguration{
		ScreenIndex:   vals[0].Uint8(),
		FieldCount:    vals[1].Uint8(),
		Layout:        typedef.ExdLayout(vals[2].Uint8()),
		ScreenEnabled: vals[3].Bool(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ExdScreenConfiguration into proto.Message. If options is nil, default options will be used.
func (m *ExdScreenConfiguration) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumExdScreenConfiguration}

	if m.ScreenIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.ScreenIndex)
		fields = append(fields, field)
	}
	if m.FieldCount != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.FieldCount)
		fields = append(fields, field)
	}
	if m.Layout != typedef.ExdLayoutInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Layout))
		fields = append(fields, field)
	}
	if m.ScreenEnabled < 2 {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Bool(m.ScreenEnabled)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetScreenIndex sets ScreenIndex value.
func (m *ExdScreenConfiguration) SetScreenIndex(v uint8) *ExdScreenConfiguration {
	m.ScreenIndex = v
	return m
}

// SetFieldCount sets FieldCount value.
//
// number of fields in screen
func (m *ExdScreenConfiguration) SetFieldCount(v uint8) *ExdScreenConfiguration {
	m.FieldCount = v
	return m
}

// SetLayout sets Layout value.
func (m *ExdScreenConfiguration) SetLayout(v typedef.ExdLayout) *ExdScreenConfiguration {
	m.Layout = v
	return m
}

// SetScreenEnabled sets ScreenEnabled value.
func (m *ExdScreenConfiguration) SetScreenEnabled(v typedef.Bool) *ExdScreenConfiguration {
	m.ScreenEnabled = v
	return m
}

// SetDeveloperFields ExdScreenConfiguration's DeveloperFields.
func (m *ExdScreenConfiguration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExdScreenConfiguration {
	m.DeveloperFields = developerFields
	return m
}
