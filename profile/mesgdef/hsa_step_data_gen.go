// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// HsaStepData is a HsaStepData message.
type HsaStepData struct {
	Timestamp          time.Time // Units: s
	Steps              []uint32  // Array: [N]; Units: steps; Total step sum
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaStepData creates new HsaStepData struct based on given mesg.
// If mesg is nil, it will return HsaStepData with all fields being set to its corresponding invalid value.
func NewHsaStepData(mesg *proto.Message) *HsaStepData {
	vals := [254]proto.Value{}

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

	return &HsaStepData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		Steps:              vals[1].SliceUint32(),
		ProcessingInterval: vals[0].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaStepData into proto.Message. If options is nil, default options will be used.
func (m *HsaStepData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHsaStepData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.Steps != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint32(m.Steps)
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ProcessingInterval)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets HsaStepData value.
//
// Units: s
func (m *HsaStepData) SetTimestamp(v time.Time) *HsaStepData {
	m.Timestamp = v
	return m
}

// SetSteps sets HsaStepData value.
//
// Array: [N]; Units: steps; Total step sum
func (m *HsaStepData) SetSteps(v []uint32) *HsaStepData {
	m.Steps = v
	return m
}

// SetProcessingInterval sets HsaStepData value.
//
// Units: s; Processing interval length in seconds
func (m *HsaStepData) SetProcessingInterval(v uint16) *HsaStepData {
	m.ProcessingInterval = v
	return m
}

// SetDeveloperFields HsaStepData's DeveloperFields.
func (m *HsaStepData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaStepData {
	m.DeveloperFields = developerFields
	return m
}
