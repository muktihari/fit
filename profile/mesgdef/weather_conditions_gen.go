// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// WeatherConditions is a WeatherConditions message.
type WeatherConditions struct {
	Timestamp                time.Time // time of update for current conditions, else forecast time
	Location                 string    // string corresponding to GCS response location string
	ObservedAtTime           time.Time
	ObservedLocationLat      int32                 // Units: semicircles
	ObservedLocationLong     int32                 // Units: semicircles
	WindDirection            uint16                // Units: degrees
	WindSpeed                uint16                // Scale: 1000; Units: m/s
	WeatherReport            typedef.WeatherReport // Current or forecast
	Temperature              int8                  // Units: C
	Condition                typedef.WeatherStatus // Corresponds to GSC Response weatherIcon field
	PrecipitationProbability uint8                 // range 0-100
	TemperatureFeelsLike     int8                  // Units: C; Heat Index if GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F
	RelativeHumidity         uint8
	DayOfWeek                typedef.DayOfWeek
	HighTemperature          int8 // Units: C
	LowTemperature           int8 // Units: C

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWeatherConditions creates new WeatherConditions struct based on given mesg.
// If mesg is nil, it will return WeatherConditions with all fields being set to its corresponding invalid value.
func NewWeatherConditions(mesg *proto.Message) *WeatherConditions {
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

	return &WeatherConditions{
		Timestamp:                datetime.ToTime(vals[253]),
		Location:                 typeconv.ToString[string](vals[8]),
		ObservedAtTime:           datetime.ToTime(vals[9]),
		ObservedLocationLat:      typeconv.ToSint32[int32](vals[10]),
		ObservedLocationLong:     typeconv.ToSint32[int32](vals[11]),
		WindDirection:            typeconv.ToUint16[uint16](vals[3]),
		WindSpeed:                typeconv.ToUint16[uint16](vals[4]),
		WeatherReport:            typeconv.ToEnum[typedef.WeatherReport](vals[0]),
		Temperature:              typeconv.ToSint8[int8](vals[1]),
		Condition:                typeconv.ToEnum[typedef.WeatherStatus](vals[2]),
		PrecipitationProbability: typeconv.ToUint8[uint8](vals[5]),
		TemperatureFeelsLike:     typeconv.ToSint8[int8](vals[6]),
		RelativeHumidity:         typeconv.ToUint8[uint8](vals[7]),
		DayOfWeek:                typeconv.ToEnum[typedef.DayOfWeek](vals[12]),
		HighTemperature:          typeconv.ToSint8[int8](vals[13]),
		LowTemperature:           typeconv.ToSint8[int8](vals[14]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WeatherConditions into proto.Message.
func (m *WeatherConditions) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumWeatherConditions)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.Location != basetype.StringInvalid && m.Location != "" {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.Location
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.ObservedAtTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = datetime.ToUint32(m.ObservedAtTime)
		fields = append(fields, field)
	}
	if m.ObservedLocationLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.ObservedLocationLat
		fields = append(fields, field)
	}
	if m.ObservedLocationLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.ObservedLocationLong
		fields = append(fields, field)
	}
	if m.WindDirection != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.WindDirection
		fields = append(fields, field)
	}
	if m.WindSpeed != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.WindSpeed
		fields = append(fields, field)
	}
	if byte(m.WeatherReport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.WeatherReport)
		fields = append(fields, field)
	}
	if m.Temperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Temperature
		fields = append(fields, field)
	}
	if byte(m.Condition) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = byte(m.Condition)
		fields = append(fields, field)
	}
	if m.PrecipitationProbability != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.PrecipitationProbability
		fields = append(fields, field)
	}
	if m.TemperatureFeelsLike != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.TemperatureFeelsLike
		fields = append(fields, field)
	}
	if m.RelativeHumidity != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.RelativeHumidity
		fields = append(fields, field)
	}
	if byte(m.DayOfWeek) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = byte(m.DayOfWeek)
		fields = append(fields, field)
	}
	if m.HighTemperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = m.HighTemperature
		fields = append(fields, field)
	}
	if m.LowTemperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.LowTemperature
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// WindSpeedScaled return WindSpeed in its scaled value [Scale: 1000; Units: m/s].
//
// If WindSpeed value is invalid, float64 invalid value will be returned.
func (m *WeatherConditions) WindSpeedScaled() float64 {
	if m.WindSpeed == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.WindSpeed, 1000, 0)
}

// SetTimestamp sets WeatherConditions value.
//
// time of update for current conditions, else forecast time
func (m *WeatherConditions) SetTimestamp(v time.Time) *WeatherConditions {
	m.Timestamp = v
	return m
}

// SetLocation sets WeatherConditions value.
//
// string corresponding to GCS response location string
func (m *WeatherConditions) SetLocation(v string) *WeatherConditions {
	m.Location = v
	return m
}

// SetObservedAtTime sets WeatherConditions value.
func (m *WeatherConditions) SetObservedAtTime(v time.Time) *WeatherConditions {
	m.ObservedAtTime = v
	return m
}

// SetObservedLocationLat sets WeatherConditions value.
//
// Units: semicircles
func (m *WeatherConditions) SetObservedLocationLat(v int32) *WeatherConditions {
	m.ObservedLocationLat = v
	return m
}

// SetObservedLocationLong sets WeatherConditions value.
//
// Units: semicircles
func (m *WeatherConditions) SetObservedLocationLong(v int32) *WeatherConditions {
	m.ObservedLocationLong = v
	return m
}

// SetWindDirection sets WeatherConditions value.
//
// Units: degrees
func (m *WeatherConditions) SetWindDirection(v uint16) *WeatherConditions {
	m.WindDirection = v
	return m
}

// SetWindSpeed sets WeatherConditions value.
//
// Scale: 1000; Units: m/s
func (m *WeatherConditions) SetWindSpeed(v uint16) *WeatherConditions {
	m.WindSpeed = v
	return m
}

// SetWeatherReport sets WeatherConditions value.
//
// Current or forecast
func (m *WeatherConditions) SetWeatherReport(v typedef.WeatherReport) *WeatherConditions {
	m.WeatherReport = v
	return m
}

// SetTemperature sets WeatherConditions value.
//
// Units: C
func (m *WeatherConditions) SetTemperature(v int8) *WeatherConditions {
	m.Temperature = v
	return m
}

// SetCondition sets WeatherConditions value.
//
// Corresponds to GSC Response weatherIcon field
func (m *WeatherConditions) SetCondition(v typedef.WeatherStatus) *WeatherConditions {
	m.Condition = v
	return m
}

// SetPrecipitationProbability sets WeatherConditions value.
//
// range 0-100
func (m *WeatherConditions) SetPrecipitationProbability(v uint8) *WeatherConditions {
	m.PrecipitationProbability = v
	return m
}

// SetTemperatureFeelsLike sets WeatherConditions value.
//
// Units: C; Heat Index if GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F
func (m *WeatherConditions) SetTemperatureFeelsLike(v int8) *WeatherConditions {
	m.TemperatureFeelsLike = v
	return m
}

// SetRelativeHumidity sets WeatherConditions value.
func (m *WeatherConditions) SetRelativeHumidity(v uint8) *WeatherConditions {
	m.RelativeHumidity = v
	return m
}

// SetDayOfWeek sets WeatherConditions value.
func (m *WeatherConditions) SetDayOfWeek(v typedef.DayOfWeek) *WeatherConditions {
	m.DayOfWeek = v
	return m
}

// SetHighTemperature sets WeatherConditions value.
//
// Units: C
func (m *WeatherConditions) SetHighTemperature(v int8) *WeatherConditions {
	m.HighTemperature = v
	return m
}

// SetLowTemperature sets WeatherConditions value.
//
// Units: C
func (m *WeatherConditions) SetLowTemperature(v int8) *WeatherConditions {
	m.LowTemperature = v
	return m
}

// SetDeveloperFields WeatherConditions's DeveloperFields.
func (m *WeatherConditions) SetDeveloperFields(developerFields ...proto.DeveloperField) *WeatherConditions {
	m.DeveloperFields = developerFields
	return m
}
