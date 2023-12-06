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

// AntRx is a AntRx message.
type AntRx struct {
	Timestamp           typedef.DateTime // Units: s;
	FractionalTimestamp uint16           // Scale: 32768; Units: s;
	MesgId              byte
	MesgData            []byte // Array: [N];
	ChannelNumber       uint8
	Data                []byte // Array: [N];

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAntRx creates new AntRx struct based on given mesg. If mesg is nil or mesg.Num is not equal to AntRx mesg number, it will return nil.
func NewAntRx(mesg proto.Message) *AntRx {
	if mesg.Num != typedef.MesgNumAntRx {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* FractionalTimestamp */
		1:   basetype.ByteInvalid,   /* MesgId */
		2:   nil,                    /* MesgData */
		3:   basetype.Uint8Invalid,  /* ChannelNumber */
		4:   nil,                    /* Data */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &AntRx{
		Timestamp:           typeconv.ToUint32[typedef.DateTime](vals[253]),
		FractionalTimestamp: typeconv.ToUint16[uint16](vals[0]),
		MesgId:              typeconv.ToByte[byte](vals[1]),
		MesgData:            typeconv.ToSliceByte[byte](vals[2]),
		ChannelNumber:       typeconv.ToUint8[uint8](vals[3]),
		Data:                typeconv.ToSliceByte[byte](vals[4]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to AntRx mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumAntRx)
func (m AntRx) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumAntRx {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.FractionalTimestamp,
		1:   m.MesgId,
		2:   m.MesgData,
		3:   m.ChannelNumber,
		4:   m.Data,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
