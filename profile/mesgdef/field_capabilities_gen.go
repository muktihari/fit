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

// FieldCapabilities is a FieldCapabilities message.
type FieldCapabilities struct {
	MessageIndex typedef.MessageIndex
	File         typedef.File
	MesgNum      typedef.MesgNum
	FieldNum     uint8
	Count        uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewFieldCapabilities creates new FieldCapabilities struct based on given mesg. If mesg is nil or mesg.Num is not equal to FieldCapabilities mesg number, it will return nil.
func NewFieldCapabilities(mesg proto.Message) *FieldCapabilities {
	if mesg.Num != typedef.MesgNumFieldCapabilities {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		254: nil, /* MessageIndex */
		0:   nil, /* File */
		1:   nil, /* MesgNum */
		2:   nil, /* FieldNum */
		3:   nil, /* Count */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &FieldCapabilities{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		File:         typeconv.ToEnum[typedef.File](vals[0]),
		MesgNum:      typeconv.ToUint16[typedef.MesgNum](vals[1]),
		FieldNum:     typeconv.ToUint8[uint8](vals[2]),
		Count:        typeconv.ToUint16[uint16](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to FieldCapabilities mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumFieldCapabilities)
func (m *FieldCapabilities) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumFieldCapabilities {
		return
	}

	vals := [...]any{
		254: typeconv.ToUint16[uint16](m.MessageIndex),
		0:   typeconv.ToEnum[byte](m.File),
		1:   typeconv.ToUint16[uint16](m.MesgNum),
		2:   m.FieldNum,
		3:   m.Count,
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
