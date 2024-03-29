// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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

// HsaEvent is a HsaEvent message.
type HsaEvent struct {
	Timestamp time.Time // Units: s
	EventId   uint8     // Event ID

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaEvent creates new HsaEvent struct based on given mesg.
// If mesg is nil, it will return HsaEvent with all fields being set to its corresponding invalid value.
func NewHsaEvent(mesg *proto.Message) *HsaEvent {
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

	return &HsaEvent{
		Timestamp: datetime.ToTime(vals[253]),
		EventId:   typeconv.ToUint8[uint8](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaEvent into proto.Message.
func (m *HsaEvent) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumHsaEvent)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.EventId != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.EventId
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets HsaEvent value.
//
// Units: s
func (m *HsaEvent) SetTimestamp(v time.Time) *HsaEvent {
	m.Timestamp = v
	return m
}

// SetEventId sets HsaEvent value.
//
// Event ID
func (m *HsaEvent) SetEventId(v uint8) *HsaEvent {
	m.EventId = v
	return m
}

// SetDeveloperFields HsaEvent's DeveloperFields.
func (m *HsaEvent) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaEvent {
	m.DeveloperFields = developerFields
	return m
}
