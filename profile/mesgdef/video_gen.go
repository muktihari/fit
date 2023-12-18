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

// Video is a Video message.
type Video struct {
	Url             string
	HostingProvider string
	Duration        uint32 // Units: ms; Playback time of video

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideo creates new Video struct based on given mesg. If mesg is nil or mesg.Num is not equal to Video mesg number, it will return nil.
func NewVideo(mesg proto.Message) *Video {
	if mesg.Num != typedef.MesgNumVideo {
		return nil
	}

	vals := [3]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &Video{
		Url:             typeconv.ToString[string](vals[0]),
		HostingProvider: typeconv.ToString[string](vals[1]),
		Duration:        typeconv.ToUint32[uint32](vals[2]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts Video into proto.Message.
func (m *Video) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumVideo)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if m.Url != basetype.StringInvalid && m.Url != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Url
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.HostingProvider != basetype.StringInvalid && m.HostingProvider != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.HostingProvider
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Duration
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Video's valid fields.
func (m *Video) size() byte {
	var size byte
	if m.Url != basetype.StringInvalid && m.Url != "" {
		size++
	}
	if m.HostingProvider != basetype.StringInvalid && m.HostingProvider != "" {
		size++
	}
	if m.Duration != basetype.Uint32Invalid {
		size++
	}
	return size
}
