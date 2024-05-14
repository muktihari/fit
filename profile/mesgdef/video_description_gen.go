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

// VideoDescription is a VideoDescription message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type VideoDescription struct {
	Text         string
	MessageIndex typedef.MessageIndex // Long descriptions will be split into multiple parts
	MessageCount uint16               // Total number of description parts

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideoDescription creates new VideoDescription struct based on given mesg.
// If mesg is nil, it will return VideoDescription with all fields being set to its corresponding invalid value.
func NewVideoDescription(mesg *proto.Message) *VideoDescription {
	vals := [255]proto.Value{}

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

	return &VideoDescription{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		MessageCount: vals[0].Uint16(),
		Text:         vals[1].String(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts VideoDescription into proto.Message. If options is nil, default options will be used.
func (m *VideoDescription) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumVideoDescription}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.MessageCount != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.MessageCount)
		fields = append(fields, field)
	}
	if m.Text != basetype.StringInvalid && m.Text != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.Text)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
//
// Long descriptions will be split into multiple parts
func (m *VideoDescription) SetMessageIndex(v typedef.MessageIndex) *VideoDescription {
	m.MessageIndex = v
	return m
}

// SetMessageCount sets MessageCount value.
//
// Total number of description parts
func (m *VideoDescription) SetMessageCount(v uint16) *VideoDescription {
	m.MessageCount = v
	return m
}

// SetText sets Text value.
func (m *VideoDescription) SetText(v string) *VideoDescription {
	m.Text = v
	return m
}

// SetDeveloperFields VideoDescription's DeveloperFields.
func (m *VideoDescription) SetDeveloperFields(developerFields ...proto.DeveloperField) *VideoDescription {
	m.DeveloperFields = developerFields
	return m
}
