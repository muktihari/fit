// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Hrv is a Hrv message.
type Hrv struct {
	Time []uint16 // Array: [N]; Scale: 1000; Units: s; Time between beats

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrv creates new Hrv struct based on given mesg.
// If mesg is nil, it will return Hrv with all fields being set to its corresponding invalid value.
func NewHrv(mesg *proto.Message) *Hrv {
	vals := [1]any{}

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

	return &Hrv{
		Time: typeconv.ToSliceUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Hrv into proto.Message. If options is nil, default options will be used.
func (m *Hrv) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumHrv)

	if m.Time != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Time
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimeScaled return Time in its scaled value [Array: [N]; Scale: 1000; Units: s; Time between beats].
//
// If Time value is invalid, nil will be returned.
func (m *Hrv) TimeScaled() []float64 {
	if m.Time == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.Time, 1000, 0)
}

// SetTime sets Hrv value.
//
// Array: [N]; Scale: 1000; Units: s; Time between beats
func (m *Hrv) SetTime(v []uint16) *Hrv {
	m.Time = v
	return m
}

// SetDeveloperFields Hrv's DeveloperFields.
func (m *Hrv) SetDeveloperFields(developerFields ...proto.DeveloperField) *Hrv {
	m.DeveloperFields = developerFields
	return m
}
