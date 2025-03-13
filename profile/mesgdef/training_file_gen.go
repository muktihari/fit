// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// TrainingFile is a TrainingFile message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type TrainingFile struct {
	Timestamp    time.Time
	TimeCreated  time.Time
	SerialNumber uint32 // Base: uint32z
	Manufacturer typedef.Manufacturer
	Product      uint16
	Type         typedef.File

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewTrainingFile creates new TrainingFile struct based on given mesg.
// If mesg is nil, it will return TrainingFile with all fields being set to its corresponding invalid value.
func NewTrainingFile(mesg *proto.Message) *TrainingFile {
	m := new(TrainingFile)
	m.Reset(mesg)
	return m
}

// Reset resets all TrainingFile's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *TrainingFile) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = TrainingFile{
		Timestamp:    datetime.ToTime(vals[253].Uint32()),
		Type:         typedef.File(vals[0].Uint8()),
		Manufacturer: typedef.Manufacturer(vals[1].Uint16()),
		Product:      vals[2].Uint16(),
		SerialNumber: vals[3].Uint32z(),
		TimeCreated:  datetime.ToTime(vals[4].Uint32()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts TrainingFile into proto.Message. If options is nil, default options will be used.
func (m *TrainingFile) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumTrainingFile}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Type != typedef.FileInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if m.Manufacturer != typedef.ManufacturerInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.Manufacturer))
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.Product)
		fields = append(fields, field)
	}
	if m.SerialNumber != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.SerialNumber)
		fields = append(fields, field)
	}
	if !m.TimeCreated.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(uint32(m.TimeCreated.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
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
func (m *TrainingFile) GetProduct() (name string, value any) {
	switch m.Manufacturer {
	case typedef.ManufacturerFaveroElectronics:
		return "favero_product", typedef.FaveroProduct(m.Product)
	case typedef.ManufacturerGarmin, typedef.ManufacturerDynastream, typedef.ManufacturerDynastreamOem, typedef.ManufacturerTacx:
		return "garmin_product", typedef.GarminProduct(m.Product)
	}
	return "product", m.Product
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TrainingFile) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// TimeCreatedUint32 returns TimeCreated in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TrainingFile) TimeCreatedUint32() uint32 { return datetime.ToUint32(m.TimeCreated) }

// SetTimestamp sets Timestamp value.
func (m *TrainingFile) SetTimestamp(v time.Time) *TrainingFile {
	m.Timestamp = v
	return m
}

// SetType sets Type value.
func (m *TrainingFile) SetType(v typedef.File) *TrainingFile {
	m.Type = v
	return m
}

// SetManufacturer sets Manufacturer value.
func (m *TrainingFile) SetManufacturer(v typedef.Manufacturer) *TrainingFile {
	m.Manufacturer = v
	return m
}

// SetProduct sets Product value.
func (m *TrainingFile) SetProduct(v uint16) *TrainingFile {
	m.Product = v
	return m
}

// SetSerialNumber sets SerialNumber value.
//
// Base: uint32z
func (m *TrainingFile) SetSerialNumber(v uint32) *TrainingFile {
	m.SerialNumber = v
	return m
}

// SetTimeCreated sets TimeCreated value.
func (m *TrainingFile) SetTimeCreated(v time.Time) *TrainingFile {
	m.TimeCreated = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *TrainingFile) SetUnknownFields(unknownFields ...proto.Field) *TrainingFile {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *TrainingFile) SetDeveloperFields(developerFields ...proto.DeveloperField) *TrainingFile {
	m.DeveloperFields = developerFields
	return m
}
