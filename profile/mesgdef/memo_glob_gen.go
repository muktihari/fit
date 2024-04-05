// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// MemoGlob is a MemoGlob message.
type MemoGlob struct {
	Memo        []byte               // Array: [N]; Deprecated. Use data field.
	Data        []uint8              // Array: [N]; Block of utf8 bytes. Note, mutltibyte characters may be split across adjoining memo_glob messages.
	PartIndex   uint32               // Sequence number of memo blocks
	MesgNum     typedef.MesgNum      // Message Number of the parent message
	ParentIndex typedef.MessageIndex // Index of mesg that this glob is associated with.
	FieldNum    uint8                // Field within the parent that this glob is associated with

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMemoGlob creates new MemoGlob struct based on given mesg.
// If mesg is nil, it will return MemoGlob with all fields being set to its corresponding invalid value.
func NewMemoGlob(mesg *proto.Message) *MemoGlob {
	vals := [251]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &MemoGlob{
		Memo:        vals[0].SliceUint8(),
		Data:        vals[4].SliceUint8(),
		PartIndex:   vals[250].Uint32(),
		MesgNum:     typedef.MesgNum(vals[1].Uint16()),
		ParentIndex: typedef.MessageIndex(vals[2].Uint16()),
		FieldNum:    vals[3].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts MemoGlob into proto.Message. If options is nil, default options will be used.
func (m *MemoGlob) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumMemoGlob}

	if m.Memo != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.SliceUint8(m.Memo)
		fields = append(fields, field)
	}
	if m.Data != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint8(m.Data)
		fields = append(fields, field)
	}
	if m.PartIndex != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 250)
		field.Value = proto.Uint32(m.PartIndex)
		fields = append(fields, field)
	}
	if uint16(m.MesgNum) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.MesgNum))
		fields = append(fields, field)
	}
	if uint16(m.ParentIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(uint16(m.ParentIndex))
		fields = append(fields, field)
	}
	if m.FieldNum != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.FieldNum)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMemo sets MemoGlob value.
//
// Array: [N]; Deprecated. Use data field.
func (m *MemoGlob) SetMemo(v []byte) *MemoGlob {
	m.Memo = v
	return m
}

// SetData sets MemoGlob value.
//
// Array: [N]; Block of utf8 bytes. Note, mutltibyte characters may be split across adjoining memo_glob messages.
func (m *MemoGlob) SetData(v []uint8) *MemoGlob {
	m.Data = v
	return m
}

// SetPartIndex sets MemoGlob value.
//
// Sequence number of memo blocks
func (m *MemoGlob) SetPartIndex(v uint32) *MemoGlob {
	m.PartIndex = v
	return m
}

// SetMesgNum sets MemoGlob value.
//
// Message Number of the parent message
func (m *MemoGlob) SetMesgNum(v typedef.MesgNum) *MemoGlob {
	m.MesgNum = v
	return m
}

// SetParentIndex sets MemoGlob value.
//
// Index of mesg that this glob is associated with.
func (m *MemoGlob) SetParentIndex(v typedef.MessageIndex) *MemoGlob {
	m.ParentIndex = v
	return m
}

// SetFieldNum sets MemoGlob value.
//
// Field within the parent that this glob is associated with
func (m *MemoGlob) SetFieldNum(v uint8) *MemoGlob {
	m.FieldNum = v
	return m
}

// SetDeveloperFields MemoGlob's DeveloperFields.
func (m *MemoGlob) SetDeveloperFields(developerFields ...proto.DeveloperField) *MemoGlob {
	m.DeveloperFields = developerFields
	return m
}
