// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// WeatherConditions is a WeatherConditions message.
type WeatherConditions struct {
	Timestamp                typedef.DateTime      // time of update for current conditions, else forecast time
	WeatherReport            typedef.WeatherReport // Current or forecast
	Temperature              int8                  // Units: C;
	Condition                typedef.WeatherStatus // Corresponds to GSC Response weatherIcon field
	WindDirection            uint16                // Units: degrees;
	WindSpeed                uint16                // Scale: 1000; Units: m/s;
	PrecipitationProbability uint8                 // range 0-100
	TemperatureFeelsLike     int8                  // Units: C; Heat Index if GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F
	RelativeHumidity         uint8
	Location                 string // string corresponding to GCS response location string
	ObservedAtTime           typedef.DateTime
	ObservedLocationLat      int32 // Units: semicircles;
	ObservedLocationLong     int32 // Units: semicircles;
	DayOfWeek                typedef.DayOfWeek
	HighTemperature          int8 // Units: C;
	LowTemperature           int8 // Units: C;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWeatherConditions creates new WeatherConditions struct based on given mesg. If mesg is nil or mesg.Num is not equal to WeatherConditions mesg number, it will return nil.
func NewWeatherConditions(mesg proto.Message) *WeatherConditions {
	if mesg.Num != typedef.MesgNumWeatherConditions {
		return nil
	}

	vals := [254]any{}
	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &WeatherConditions{
		Timestamp:                typeconv.ToUint32[typedef.DateTime](vals[253]),
		WeatherReport:            typeconv.ToEnum[typedef.WeatherReport](vals[0]),
		Temperature:              typeconv.ToSint8[int8](vals[1]),
		Condition:                typeconv.ToEnum[typedef.WeatherStatus](vals[2]),
		WindDirection:            typeconv.ToUint16[uint16](vals[3]),
		WindSpeed:                typeconv.ToUint16[uint16](vals[4]),
		PrecipitationProbability: typeconv.ToUint8[uint8](vals[5]),
		TemperatureFeelsLike:     typeconv.ToSint8[int8](vals[6]),
		RelativeHumidity:         typeconv.ToUint8[uint8](vals[7]),
		Location:                 typeconv.ToString[string](vals[8]),
		ObservedAtTime:           typeconv.ToUint32[typedef.DateTime](vals[9]),
		ObservedLocationLat:      typeconv.ToSint32[int32](vals[10]),
		ObservedLocationLong:     typeconv.ToSint32[int32](vals[11]),
		DayOfWeek:                typeconv.ToEnum[typedef.DayOfWeek](vals[12]),
		HighTemperature:          typeconv.ToSint8[int8](vals[13]),
		LowTemperature:           typeconv.ToSint8[int8](vals[14]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// ToMesg converts WeatherConditions into proto.Message.
func (m *WeatherConditions) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumWeatherConditions)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = typeconv.ToUint32[uint32](m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.WeatherReport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.WeatherReport)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Temperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Temperature
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Condition) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = typeconv.ToEnum[byte](m.Condition)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.WindDirection != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.WindDirection
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.WindSpeed != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.WindSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.PrecipitationProbability != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.PrecipitationProbability
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TemperatureFeelsLike != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.TemperatureFeelsLike
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.RelativeHumidity != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.RelativeHumidity
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Location != basetype.StringInvalid && m.Location != "" {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.Location
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint32[uint32](m.ObservedAtTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToUint32[uint32](m.ObservedAtTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ObservedLocationLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.ObservedLocationLat
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ObservedLocationLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.ObservedLocationLong
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.DayOfWeek) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = typeconv.ToEnum[byte](m.DayOfWeek)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.HighTemperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = m.HighTemperature
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.LowTemperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.LowTemperature
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of WeatherConditions's valid fields.
func (m *WeatherConditions) size() byte {
	var size byte
	if typeconv.ToUint32[uint32](m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.WeatherReport) != basetype.EnumInvalid {
		size++
	}
	if m.Temperature != basetype.Sint8Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.Condition) != basetype.EnumInvalid {
		size++
	}
	if m.WindDirection != basetype.Uint16Invalid {
		size++
	}
	if m.WindSpeed != basetype.Uint16Invalid {
		size++
	}
	if m.PrecipitationProbability != basetype.Uint8Invalid {
		size++
	}
	if m.TemperatureFeelsLike != basetype.Sint8Invalid {
		size++
	}
	if m.RelativeHumidity != basetype.Uint8Invalid {
		size++
	}
	if m.Location != basetype.StringInvalid && m.Location != "" {
		size++
	}
	if typeconv.ToUint32[uint32](m.ObservedAtTime) != basetype.Uint32Invalid {
		size++
	}
	if m.ObservedLocationLat != basetype.Sint32Invalid {
		size++
	}
	if m.ObservedLocationLong != basetype.Sint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.DayOfWeek) != basetype.EnumInvalid {
		size++
	}
	if m.HighTemperature != basetype.Sint8Invalid {
		size++
	}
	if m.LowTemperature != basetype.Sint8Invalid {
		size++
	}
	return size
}
