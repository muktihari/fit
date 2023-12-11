// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// HrmProfile is a HrmProfile message.
type HrmProfile struct {
	MessageIndex      typedef.MessageIndex
	Enabled           bool
	HrmAntId          uint16
	LogHrv            bool
	HrmAntIdTransType uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrmProfile creates new HrmProfile struct based on given mesg. If mesg is nil or mesg.Num is not equal to HrmProfile mesg number, it will return nil.
func NewHrmProfile(mesg proto.Message) *HrmProfile {
	if mesg.Num != typedef.MesgNumHrmProfile {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		254: nil, /* MessageIndex */
		0:   nil, /* Enabled */
		1:   nil, /* HrmAntId */
		2:   nil, /* LogHrv */
		3:   nil, /* HrmAntIdTransType */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &HrmProfile{
		MessageIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Enabled:           typeconv.ToBool[bool](vals[0]),
		HrmAntId:          typeconv.ToUint16z[uint16](vals[1]),
		LogHrv:            typeconv.ToBool[bool](vals[2]),
		HrmAntIdTransType: typeconv.ToUint8z[uint8](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to HrmProfile mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumHrmProfile)
func (m HrmProfile) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumHrmProfile {
		return
	}

	vals := [...]any{
		254: m.MessageIndex,
		0:   m.Enabled,
		1:   m.HrmAntId,
		2:   m.LogHrv,
		3:   m.HrmAntIdTransType,
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		field.Value = vals[field.Num]
	}

	mesg.DeveloperFields = m.DeveloperFields
}
