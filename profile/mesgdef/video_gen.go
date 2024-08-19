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

// Video is a Video message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Video struct {
	Url             string
	HostingProvider string
	Duration        uint32 // Units: ms; Playback time of video

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideo creates new Video struct based on given mesg.
// If mesg is nil, it will return Video with all fields being set to its corresponding invalid value.
func NewVideo(mesg *proto.Message) *Video {
	vals := [3]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 2 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Video{
		Url:             vals[0].String(),
		HostingProvider: vals[1].String(),
		Duration:        vals[2].Uint32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Video into proto.Message. If options is nil, default options will be used.
func (m *Video) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumVideo}

	if m.Url != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.Url)
		fields = append(fields, field)
	}
	if m.HostingProvider != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.HostingProvider)
		fields = append(fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.Duration)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetUrl sets Url value.
func (m *Video) SetUrl(v string) *Video {
	m.Url = v
	return m
}

// SetHostingProvider sets HostingProvider value.
func (m *Video) SetHostingProvider(v string) *Video {
	m.HostingProvider = v
	return m
}

// SetDuration sets Duration value.
//
// Units: ms; Playback time of video
func (m *Video) SetDuration(v uint32) *Video {
	m.Duration = v
	return m
}

// SetDeveloperFields Video's DeveloperFields.
func (m *Video) SetDeveloperFields(developerFields ...proto.DeveloperField) *Video {
	m.DeveloperFields = developerFields
	return m
}
