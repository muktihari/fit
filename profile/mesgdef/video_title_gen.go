// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// VideoTitle is a VideoTitle message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type VideoTitle struct {
	Text         string
	MessageIndex typedef.MessageIndex // Long titles will be split into multiple parts
	MessageCount uint16               // Total number of title parts

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewVideoTitle creates new VideoTitle struct based on given mesg.
// If mesg is nil, it will return VideoTitle with all fields being set to its corresponding invalid value.
func NewVideoTitle(mesg *proto.Message) *VideoTitle {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &VideoTitle{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		MessageCount: vals[0].Uint16(),
		Text:         vals[1].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts VideoTitle into proto.Message. If options is nil, default options will be used.
func (m *VideoTitle) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumVideoTitle}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.MessageCount != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.MessageCount)
		fields = append(fields, field)
	}
	if m.Text != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.Text)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
//
// Long titles will be split into multiple parts
func (m *VideoTitle) SetMessageIndex(v typedef.MessageIndex) *VideoTitle {
	m.MessageIndex = v
	return m
}

// SetMessageCount sets MessageCount value.
//
// Total number of title parts
func (m *VideoTitle) SetMessageCount(v uint16) *VideoTitle {
	m.MessageCount = v
	return m
}

// SetText sets Text value.
func (m *VideoTitle) SetText(v string) *VideoTitle {
	m.Text = v
	return m
}

// SetUnknownFields VideoTitle's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *VideoTitle) SetUnknownFields(unknownFields ...proto.Field) *VideoTitle {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields VideoTitle's DeveloperFields.
func (m *VideoTitle) SetDeveloperFields(developerFields ...proto.DeveloperField) *VideoTitle {
	m.DeveloperFields = developerFields
	return m
}
