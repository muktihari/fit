// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// Hr is a Hr message.
type Hr struct {
	Timestamp           time.Time
	FractionalTimestamp uint16   // Scale: 32768; Units: s;
	Time256             uint8    // Scale: 256; Units: s;
	FilteredBpm         []uint8  // Array: [N]; Units: bpm;
	EventTimestamp      []uint32 // Array: [N]; Scale: 1024; Units: s;
	EventTimestamp12    []byte   // Array: [N]; Scale: 1024; Units: s;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHr creates new Hr struct based on given mesg.
// If mesg is nil, it will return Hr with all fields being set to its corresponding invalid value.
func NewHr(mesg *proto.Message) *Hr {
	vals := [254]any{}

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

	return &Hr{
		Timestamp:           datetime.ToTime(vals[253]),
		FractionalTimestamp: typeconv.ToUint16[uint16](vals[0]),
		Time256:             typeconv.ToUint8[uint8](vals[1]),
		FilteredBpm:         typeconv.ToSliceUint8[uint8](vals[6]),
		EventTimestamp:      typeconv.ToSliceUint32[uint32](vals[9]),
		EventTimestamp12:    typeconv.ToSliceByte[byte](vals[10]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Hr into proto.Message.
func (m *Hr) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumHr)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.FractionalTimestamp
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Time256 != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Time256
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FilteredBpm != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.FilteredBpm
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EventTimestamp != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.EventTimestamp
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EventTimestamp12 != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.EventTimestamp12
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Hr's valid fields.
func (m *Hr) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		size++
	}
	if m.Time256 != basetype.Uint8Invalid {
		size++
	}
	if m.FilteredBpm != nil {
		size++
	}
	if m.EventTimestamp != nil {
		size++
	}
	if m.EventTimestamp12 != nil {
		size++
	}
	return size
}

// SetTimestamp sets Hr value.
func (m *Hr) SetTimestamp(v time.Time) *Hr {
	m.Timestamp = v
	return m
}

// SetFractionalTimestamp sets Hr value.
//
// Scale: 32768; Units: s;
func (m *Hr) SetFractionalTimestamp(v uint16) *Hr {
	m.FractionalTimestamp = v
	return m
}

// SetTime256 sets Hr value.
//
// Scale: 256; Units: s;
func (m *Hr) SetTime256(v uint8) *Hr {
	m.Time256 = v
	return m
}

// SetFilteredBpm sets Hr value.
//
// Array: [N]; Units: bpm;
func (m *Hr) SetFilteredBpm(v []uint8) *Hr {
	m.FilteredBpm = v
	return m
}

// SetEventTimestamp sets Hr value.
//
// Array: [N]; Scale: 1024; Units: s;
func (m *Hr) SetEventTimestamp(v []uint32) *Hr {
	m.EventTimestamp = v
	return m
}

// SetEventTimestamp12 sets Hr value.
//
// Array: [N]; Scale: 1024; Units: s;
func (m *Hr) SetEventTimestamp12(v []byte) *Hr {
	m.EventTimestamp12 = v
	return m
}

// SetDeveloperFields Hr's DeveloperFields.
func (m *Hr) SetDeveloperFields(developerFields ...proto.DeveloperField) *Hr {
	m.DeveloperFields = developerFields
	return m
}
