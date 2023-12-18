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

// OhrSettings is a OhrSettings message.
type OhrSettings struct {
	Timestamp typedef.DateTime // Units: s;
	Enabled   typedef.Switch

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewOhrSettings creates new OhrSettings struct based on given mesg. If mesg is nil or mesg.Num is not equal to OhrSettings mesg number, it will return nil.
func NewOhrSettings(mesg proto.Message) *OhrSettings {
	if mesg.Num != typedef.MesgNumOhrSettings {
		return nil
	}

	vals := [254]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &OhrSettings{
		Timestamp: typeconv.ToUint32[typedef.DateTime](vals[253]),
		Enabled:   typeconv.ToEnum[typedef.Switch](vals[0]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts OhrSettings into proto.Message.
func (m *OhrSettings) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumOhrSettings)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = typeconv.ToUint32[uint32](m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Enabled) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.Enabled)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of OhrSettings's valid fields.
func (m *OhrSettings) size() byte {
	var size byte
	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Enabled) != basetype.EnumInvalid {
		size++
	}
	return size
}
