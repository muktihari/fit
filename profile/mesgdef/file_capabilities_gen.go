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

// FileCapabilities is a FileCapabilities message.
type FileCapabilities struct {
	MessageIndex typedef.MessageIndex
	Type         typedef.File
	Flags        typedef.FileFlags
	Directory    string
	MaxCount     uint16
	MaxSize      uint32 // Units: bytes;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewFileCapabilities creates new FileCapabilities struct based on given mesg. If mesg is nil or mesg.Num is not equal to FileCapabilities mesg number, it will return nil.
func NewFileCapabilities(mesg proto.Message) *FileCapabilities {
	if mesg.Num != typedef.MesgNumFileCapabilities {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		254: nil, /* MessageIndex */
		0:   nil, /* Type */
		1:   nil, /* Flags */
		2:   nil, /* Directory */
		3:   nil, /* MaxCount */
		4:   nil, /* MaxSize */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &FileCapabilities{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Type:         typeconv.ToEnum[typedef.File](vals[0]),
		Flags:        typeconv.ToUint8z[typedef.FileFlags](vals[1]),
		Directory:    typeconv.ToString[string](vals[2]),
		MaxCount:     typeconv.ToUint16[uint16](vals[3]),
		MaxSize:      typeconv.ToUint32[uint32](vals[4]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to FileCapabilities mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumFileCapabilities)
func (m FileCapabilities) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumFileCapabilities {
		return
	}

	vals := [...]any{
		254: m.MessageIndex,
		0:   m.Type,
		1:   m.Flags,
		2:   m.Directory,
		3:   m.MaxCount,
		4:   m.MaxSize,
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
