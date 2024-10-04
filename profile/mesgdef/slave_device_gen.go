// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SlaveDevice is a SlaveDevice message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SlaveDevice struct {
	Manufacturer typedef.Manufacturer
	Product      uint16

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSlaveDevice creates new SlaveDevice struct based on given mesg.
// If mesg is nil, it will return SlaveDevice with all fields being set to its corresponding invalid value.
func NewSlaveDevice(mesg *proto.Message) *SlaveDevice {
	vals := [2]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 1 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &SlaveDevice{
		Manufacturer: typedef.Manufacturer(vals[0].Uint16()),
		Product:      vals[1].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SlaveDevice into proto.Message. If options is nil, default options will be used.
func (m *SlaveDevice) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSlaveDevice}

	if m.Manufacturer != typedef.ManufacturerInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.Manufacturer))
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.Product)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetProduct returns Dynamic Field interpretation of Product. Otherwise, returns the original value of Product.
//
// Based on m.Manufacturer:
//   - name: "favero_product", value: typedef.FaveroProduct(m.Product)
//   - name: "garmin_product", value: typedef.GarminProduct(m.Product)
//
// Otherwise:
//   - name: "product", value: m.Product
func (m *SlaveDevice) GetProduct() (name string, value any) {
	switch m.Manufacturer {
	case typedef.ManufacturerFaveroElectronics:
		return "favero_product", typedef.FaveroProduct(m.Product)
	case typedef.ManufacturerGarmin, typedef.ManufacturerDynastream, typedef.ManufacturerDynastreamOem, typedef.ManufacturerTacx:
		return "garmin_product", typedef.GarminProduct(m.Product)
	}
	return "product", m.Product
}

// SetManufacturer sets Manufacturer value.
func (m *SlaveDevice) SetManufacturer(v typedef.Manufacturer) *SlaveDevice {
	m.Manufacturer = v
	return m
}

// SetProduct sets Product value.
func (m *SlaveDevice) SetProduct(v uint16) *SlaveDevice {
	m.Product = v
	return m
}

// SetUnknownFields SlaveDevice's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SlaveDevice) SetUnknownFields(unknownFields ...proto.Field) *SlaveDevice {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields SlaveDevice's DeveloperFields.
func (m *SlaveDevice) SetDeveloperFields(developerFields ...proto.DeveloperField) *SlaveDevice {
	m.DeveloperFields = developerFields
	return m
}
