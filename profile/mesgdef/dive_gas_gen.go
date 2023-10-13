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

// DiveGas is a DiveGas message.
type DiveGas struct {
	MessageIndex  typedef.MessageIndex
	HeliumContent uint8 // Units: percent;
	OxygenContent uint8 // Units: percent;
	Status        typedef.DiveGasStatus
	Mode          typedef.DiveGasMode

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDiveGas creates new DiveGas struct based on given mesg. If mesg is nil or mesg.Num is not equal to DiveGas mesg number, it will return nil.
func NewDiveGas(mesg proto.Message) *DiveGas {
	if mesg.Num != typedef.MesgNumDiveGas {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		0:   basetype.Uint8Invalid,  /* HeliumContent */
		1:   basetype.Uint8Invalid,  /* OxygenContent */
		2:   basetype.EnumInvalid,   /* Status */
		3:   basetype.EnumInvalid,   /* Mode */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &DiveGas{
		MessageIndex:  typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		HeliumContent: typeconv.ToUint8[uint8](vals[0]),
		OxygenContent: typeconv.ToUint8[uint8](vals[1]),
		Status:        typeconv.ToEnum[typedef.DiveGasStatus](vals[2]),
		Mode:          typeconv.ToEnum[typedef.DiveGasMode](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DiveGas mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDiveGas)
func (m DiveGas) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDiveGas {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.HeliumContent,
		1:   m.OxygenContent,
		2:   m.Status,
		3:   m.Mode,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}