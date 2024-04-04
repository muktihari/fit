// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// HsaRespirationData is a HsaRespirationData message.
type HsaRespirationData struct {
	Timestamp          time.Time // Units: s
	RespirationRate    []int16   // Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaRespirationData creates new HsaRespirationData struct based on given mesg.
// If mesg is nil, it will return HsaRespirationData with all fields being set to its corresponding invalid value.
func NewHsaRespirationData(mesg *proto.Message) *HsaRespirationData {
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

	return &HsaRespirationData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		RespirationRate:    vals[1].SliceInt16(),
		ProcessingInterval: vals[0].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaRespirationData into proto.Message. If options is nil, default options will be used.
func (m *HsaRespirationData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumHsaRespirationData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.RespirationRate != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceInt16(m.RespirationRate)
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

// RespirationRateScaled return RespirationRate in its scaled value [Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist].
//
// If RespirationRate value is invalid, nil will be returned.
func (m *HsaRespirationData) RespirationRateScaled() []float64 {
	if m.RespirationRate == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.RespirationRate, 100, 0)
}

// SetTimestamp sets HsaRespirationData value.
//
// Units: s
func (m *HsaRespirationData) SetTimestamp(v time.Time) *HsaRespirationData {
	m.Timestamp = v
	return m
}

// SetRespirationRate sets HsaRespirationData value.
//
// Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist
func (m *HsaRespirationData) SetRespirationRate(v []int16) *HsaRespirationData {
	m.RespirationRate = v
	return m
}

// SetProcessingInterval sets HsaRespirationData value.
//
// Units: s; Processing interval length in seconds
func (m *HsaRespirationData) SetProcessingInterval(v uint16) *HsaRespirationData {
	m.ProcessingInterval = v
	return m
}

// SetDeveloperFields HsaRespirationData's DeveloperFields.
func (m *HsaRespirationData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaRespirationData {
	m.DeveloperFields = developerFields
	return m
}
