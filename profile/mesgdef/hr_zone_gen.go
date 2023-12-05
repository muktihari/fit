// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

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

// HrZone is a HrZone message.
type HrZone struct {
	MessageIndex typedef.MessageIndex
	HighBpm      uint8 // Units: bpm;
	Name         string

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrZone creates new HrZone struct based on given mesg. If mesg is nil or mesg.Num is not equal to HrZone mesg number, it will return nil.
func NewHrZone(mesg proto.Message) *HrZone {
	if mesg.Num != typedef.MesgNumHrZone {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		1:   basetype.Uint8Invalid,  /* HighBpm */
		2:   basetype.StringInvalid, /* Name */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &HrZone{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		HighBpm:      typeconv.ToUint8[uint8](vals[1]),
		Name:         typeconv.ToString[string](vals[2]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to HrZone mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumHrZone)
func (m HrZone) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumHrZone {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		1:   m.HighBpm,
		2:   m.Name,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
