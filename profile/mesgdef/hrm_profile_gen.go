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

// HrmProfile is a HrmProfile message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HrmProfile struct {
	MessageIndex      typedef.MessageIndex
	HrmAntId          uint16
	Enabled           bool
	LogHrv            bool
	HrmAntIdTransType uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrmProfile creates new HrmProfile struct based on given mesg.
// If mesg is nil, it will return HrmProfile with all fields being set to its corresponding invalid value.
func NewHrmProfile(mesg *proto.Message) *HrmProfile {
	vals := [255]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &HrmProfile{
		MessageIndex:      typedef.MessageIndex(vals[254].Uint16()),
		Enabled:           vals[0].Bool(),
		HrmAntId:          vals[1].Uint16z(),
		LogHrv:            vals[2].Bool(),
		HrmAntIdTransType: vals[3].Uint8z(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HrmProfile into proto.Message. If options is nil, default options will be used.
func (m *HrmProfile) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHrmProfile}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Enabled != false {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if uint16(m.HrmAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.HrmAntId)
		fields = append(fields, field)
	}
	if m.LogHrv != false {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Bool(m.LogHrv)
		fields = append(fields, field)
	}
	if uint8(m.HrmAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.HrmAntIdTransType)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
func (m *HrmProfile) SetMessageIndex(v typedef.MessageIndex) *HrmProfile {
	m.MessageIndex = v
	return m
}

// SetEnabled sets Enabled value.
func (m *HrmProfile) SetEnabled(v bool) *HrmProfile {
	m.Enabled = v
	return m
}

// SetHrmAntId sets HrmAntId value.
func (m *HrmProfile) SetHrmAntId(v uint16) *HrmProfile {
	m.HrmAntId = v
	return m
}

// SetLogHrv sets LogHrv value.
func (m *HrmProfile) SetLogHrv(v bool) *HrmProfile {
	m.LogHrv = v
	return m
}

// SetHrmAntIdTransType sets HrmAntIdTransType value.
func (m *HrmProfile) SetHrmAntIdTransType(v uint8) *HrmProfile {
	m.HrmAntIdTransType = v
	return m
}

// SetDeveloperFields HrmProfile's DeveloperFields.
func (m *HrmProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *HrmProfile {
	m.DeveloperFields = developerFields
	return m
}
