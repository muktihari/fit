// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

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

// RespirationRate is a RespirationRate message.
type RespirationRate struct {
	Timestamp       typedef.DateTime
	RespirationRate int16 // Scale: 100; Units: breaths/min; Breaths * 100 /min, -300 indicates invalid, -200 indicates large motion, -100 indicates off wrist

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewRespirationRate creates new RespirationRate struct based on given mesg. If mesg is nil or mesg.Num is not equal to RespirationRate mesg number, it will return nil.
func NewRespirationRate(mesg proto.Message) *RespirationRate {
	if mesg.Num != typedef.MesgNumRespirationRate {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Sint16Invalid, /* RespirationRate */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &RespirationRate{
		Timestamp:       typeconv.ToUint32[typedef.DateTime](vals[253]),
		RespirationRate: typeconv.ToSint16[int16](vals[0]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to RespirationRate mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumRespirationRate)
func (m RespirationRate) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumRespirationRate {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.RespirationRate,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
