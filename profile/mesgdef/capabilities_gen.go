// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

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

// Capabilities is a Capabilities message.
type Capabilities struct {
	Languages             []uint8              // Array: [N]; Use language_bits_x types where x is index of array.
	Sports                []typedef.SportBits0 // Array: [N]; Use sport_bits_x types where x is index of array.
	WorkoutsSupported     typedef.WorkoutCapabilities
	ConnectivitySupported typedef.ConnectivityCapabilities

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewCapabilities creates new Capabilities struct based on given mesg. If mesg is nil or mesg.Num is not equal to Capabilities mesg number, it will return nil.
func NewCapabilities(mesg proto.Message) *Capabilities {
	if mesg.Num != typedef.MesgNumCapabilities {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0:  nil,                     /* Languages */
		1:  nil,                     /* Sports */
		21: basetype.Uint32zInvalid, /* WorkoutsSupported */
		23: basetype.Uint32zInvalid, /* ConnectivitySupported */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Capabilities{
		Languages:             typeconv.ToSliceUint8z[uint8](vals[0]),
		Sports:                typeconv.ToSliceUint8z[typedef.SportBits0](vals[1]),
		WorkoutsSupported:     typeconv.ToUint32z[typedef.WorkoutCapabilities](vals[21]),
		ConnectivitySupported: typeconv.ToUint32z[typedef.ConnectivityCapabilities](vals[23]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Capabilities mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumCapabilities)
func (m Capabilities) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumCapabilities {
		return
	}

	vals := [256]any{
		0:  m.Languages,
		1:  m.Sports,
		21: m.WorkoutsSupported,
		23: m.ConnectivitySupported,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
