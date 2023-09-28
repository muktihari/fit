// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

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

// Software is a Software message.
type Software struct {
	MessageIndex typedef.MessageIndex
	Version      uint16 // Scale: 100;
	PartNumber   string

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSoftware creates new Software struct based on given mesg. If mesg is nil or mesg.Num is not equal to Software mesg number, it will return nil.
func NewSoftware(mesg proto.Message) *Software {
	if mesg.Num != typedef.MesgNumSoftware {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		3:   basetype.Uint16Invalid, /* Version */
		5:   basetype.StringInvalid, /* PartNumber */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Software{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Version:      typeconv.ToUint16[uint16](vals[3]),
		PartNumber:   typeconv.ToString[string](vals[5]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Software mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSoftware)
func (m Software) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSoftware {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		3:   m.Version,
		5:   m.PartNumber,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
