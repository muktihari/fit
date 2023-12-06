// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

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

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0: basetype.StringInvalid, /* Url */
		1: basetype.StringInvalid, /* HostingProvider */
		2: basetype.Uint32Invalid, /* Duration */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Video{
		Url:             typeconv.ToString[string](vals[0]),
		HostingProvider: typeconv.ToString[string](vals[1]),
		Duration:        typeconv.ToUint32[uint32](vals[2]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Video mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumVideo)
func (m Video) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumVideo {
		return
	}

	vals := [256]any{
		0: m.Url,
		1: m.HostingProvider,
		2: m.Duration,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
