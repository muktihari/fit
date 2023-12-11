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

// FileId is a FileId message.
type FileId struct {
	Type         typedef.File
	Manufacturer typedef.Manufacturer
	Product      uint16
	SerialNumber uint32
	TimeCreated  typedef.DateTime // Only set for files that are can be created/erased.
	Number       uint16           // Only set for files that are not created/erased.
	ProductName  string           // Optional free form string to indicate the devices name or model
}

// NewFileId creates new FileId struct based on given mesg. If mesg is nil or mesg.Num is not equal to FileId mesg number, it will return nil.
func NewFileId(mesg proto.Message) *FileId {
	if mesg.Num != typedef.MesgNumFileId {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		0: nil, /* Type */
		1: nil, /* Manufacturer */
		2: nil, /* Product */
		3: nil, /* SerialNumber */
		4: nil, /* TimeCreated */
		5: nil, /* Number */
		8: nil, /* ProductName */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &FileId{
		Type:         typeconv.ToEnum[typedef.File](vals[0]),
		Manufacturer: typeconv.ToUint16[typedef.Manufacturer](vals[1]),
		Product:      typeconv.ToUint16[uint16](vals[2]),
		SerialNumber: typeconv.ToUint32z[uint32](vals[3]),
		TimeCreated:  typeconv.ToUint32[typedef.DateTime](vals[4]),
		Number:       typeconv.ToUint16[uint16](vals[5]),
		ProductName:  typeconv.ToString[string](vals[8]),
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to FileId mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumFileId)
func (m FileId) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumFileId {
		return
	}

	vals := [...]any{
		0: m.Type,
		1: m.Manufacturer,
		2: m.Product,
		3: m.SerialNumber,
		4: m.TimeCreated,
		5: m.Number,
		8: m.ProductName,
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		field.Value = vals[field.Num]
	}

}
