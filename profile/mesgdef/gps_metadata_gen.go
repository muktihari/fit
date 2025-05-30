// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// GpsMetadata is a GpsMetadata message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type GpsMetadata struct {
	Timestamp        time.Time // Units: s; Whole second part of the timestamp.
	UtcTimestamp     time.Time // Units: s; Used to correlate UTC to system time if the timestamp of the message is in system time. This UTC time is derived from the GPS data.
	PositionLat      int32     // Units: semicircles
	PositionLong     int32     // Units: semicircles
	EnhancedAltitude uint32    // Scale: 5; Offset: 500; Units: m
	EnhancedSpeed    uint32    // Scale: 1000; Units: m/s
	Velocity         [3]int16  // Array: [3]; Scale: 100; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.
	TimestampMs      uint16    // Units: ms; Millisecond part of the timestamp.
	Heading          uint16    // Scale: 100; Units: degrees

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewGpsMetadata creates new GpsMetadata struct based on given mesg.
// If mesg is nil, it will return GpsMetadata with all fields being set to its corresponding invalid value.
func NewGpsMetadata(mesg *proto.Message) *GpsMetadata {
	m := new(GpsMetadata)
	m.Reset(mesg)
	return m
}

// Reset resets all GpsMetadata's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *GpsMetadata) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

	if mesg != nil {
		var n int
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				n++
			}
		}
		unknownFields = make([]proto.Field, 0, n)
		for i := range mesg.Fields {
			if mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 254 {
				vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
			}
		}
		developerFields = mesg.DeveloperFields
	}

	*m = GpsMetadata{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		TimestampMs:      vals[0].Uint16(),
		PositionLat:      vals[1].Int32(),
		PositionLong:     vals[2].Int32(),
		EnhancedAltitude: vals[3].Uint32(),
		EnhancedSpeed:    vals[4].Uint32(),
		Heading:          vals[5].Uint16(),
		UtcTimestamp:     datetime.ToTime(vals[6].Uint32()),
		Velocity: func() (arr [3]int16) {
			arr = [3]int16{
				basetype.Sint16Invalid,
				basetype.Sint16Invalid,
				basetype.Sint16Invalid,
			}
			copy(arr[:], vals[7].SliceInt16())
			return arr
		}(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts GpsMetadata into proto.Message. If options is nil, default options will be used.
func (m *GpsMetadata) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fields := make([]proto.Field, 0, 9)
	mesg := proto.Message{Num: typedef.MesgNumGpsMetadata}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Int32(m.PositionLat)
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Int32(m.PositionLong)
		fields = append(fields, field)
	}
	if m.EnhancedAltitude != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.EnhancedAltitude)
		fields = append(fields, field)
	}
	if m.EnhancedSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.EnhancedSpeed)
		fields = append(fields, field)
	}
	if m.Heading != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.Heading)
		fields = append(fields, field)
	}
	if !m.UtcTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(uint32(m.UtcTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Velocity != [3]int16{
		basetype.Sint16Invalid,
		basetype.Sint16Invalid,
		basetype.Sint16Invalid,
	} {
		field := fac.CreateField(mesg.Num, 7)
		copied := m.Velocity
		field.Value = proto.SliceInt16(copied[:])
		fields = append(fields, field)
	}

	n := len(fields)
	mesg.Fields = make([]proto.Field, n+len(m.UnknownFields))
	copy(mesg.Fields[:n], fields)
	copy(mesg.Fields[n:], m.UnknownFields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *GpsMetadata) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// UtcTimestampUint32 returns UtcTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *GpsMetadata) UtcTimestampUint32() uint32 { return datetime.ToUint32(m.UtcTimestamp) }

// EnhancedAltitudeScaled return EnhancedAltitude in its scaled value.
// If EnhancedAltitude value is invalid, float64 invalid value will be returned.
//
// Scale: 5; Offset: 500; Units: m
func (m *GpsMetadata) EnhancedAltitudeScaled() float64 {
	if m.EnhancedAltitude == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.EnhancedAltitude)/5 - 500
}

// EnhancedSpeedScaled return EnhancedSpeed in its scaled value.
// If EnhancedSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *GpsMetadata) EnhancedSpeedScaled() float64 {
	if m.EnhancedSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.EnhancedSpeed)/1000 - 0
}

// HeadingScaled return Heading in its scaled value.
// If Heading value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: degrees
func (m *GpsMetadata) HeadingScaled() float64 {
	if m.Heading == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Heading)/100 - 0
}

// VelocityScaled return Velocity in its scaled value.
// If Velocity value is invalid, nil will be returned.
//
// Array: [3]; Scale: 100; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.
func (m *GpsMetadata) VelocityScaled() [3]float64 {
	if m.Velocity == [3]int16{
		basetype.Sint16Invalid,
		basetype.Sint16Invalid,
		basetype.Sint16Invalid,
	} {
		return [3]float64{
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
			math.Float64frombits(basetype.Float64Invalid),
		}
	}
	var vals [3]float64
	for i := range m.Velocity {
		if m.Velocity[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.Velocity[i])/100 - 0
	}
	return vals
}

// PositionLatDegrees returns PositionLat in degrees instead of semicircles.
// If PositionLat value is invalid, float64 invalid value will be returned.
func (m *GpsMetadata) PositionLatDegrees() float64 {
	return semicircles.ToDegrees(m.PositionLat)
}

// PositionLongDegrees returns PositionLong in degrees instead of semicircles.
// If PositionLong value is invalid, float64 invalid value will be returned.
func (m *GpsMetadata) PositionLongDegrees() float64 {
	return semicircles.ToDegrees(m.PositionLong)
}

// SetTimestamp sets Timestamp value.
//
// Units: s; Whole second part of the timestamp.
func (m *GpsMetadata) SetTimestamp(v time.Time) *GpsMetadata {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *GpsMetadata) SetTimestampMs(v uint16) *GpsMetadata {
	m.TimestampMs = v
	return m
}

// SetPositionLat sets PositionLat value.
//
// Units: semicircles
func (m *GpsMetadata) SetPositionLat(v int32) *GpsMetadata {
	m.PositionLat = v
	return m
}

// SetPositionLatDegrees is similar to SetPositionLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *GpsMetadata) SetPositionLatDegrees(degrees float64) *GpsMetadata {
	m.PositionLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetPositionLong sets PositionLong value.
//
// Units: semicircles
func (m *GpsMetadata) SetPositionLong(v int32) *GpsMetadata {
	m.PositionLong = v
	return m
}

// SetPositionLongDegrees is similar to SetPositionLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *GpsMetadata) SetPositionLongDegrees(degrees float64) *GpsMetadata {
	m.PositionLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetEnhancedAltitude sets EnhancedAltitude value.
//
// Scale: 5; Offset: 500; Units: m
func (m *GpsMetadata) SetEnhancedAltitude(v uint32) *GpsMetadata {
	m.EnhancedAltitude = v
	return m
}

// SetEnhancedAltitudeScaled is similar to SetEnhancedAltitude except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 5; Offset: 500; Units: m
func (m *GpsMetadata) SetEnhancedAltitudeScaled(v float64) *GpsMetadata {
	unscaled := (v + 500) * 5
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.EnhancedAltitude = uint32(basetype.Uint32Invalid)
		return m
	}
	m.EnhancedAltitude = uint32(unscaled)
	return m
}

// SetEnhancedSpeed sets EnhancedSpeed value.
//
// Scale: 1000; Units: m/s
func (m *GpsMetadata) SetEnhancedSpeed(v uint32) *GpsMetadata {
	m.EnhancedSpeed = v
	return m
}

// SetEnhancedSpeedScaled is similar to SetEnhancedSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *GpsMetadata) SetEnhancedSpeedScaled(v float64) *GpsMetadata {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.EnhancedSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.EnhancedSpeed = uint32(unscaled)
	return m
}

// SetHeading sets Heading value.
//
// Scale: 100; Units: degrees
func (m *GpsMetadata) SetHeading(v uint16) *GpsMetadata {
	m.Heading = v
	return m
}

// SetHeadingScaled is similar to SetHeading except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: degrees
func (m *GpsMetadata) SetHeadingScaled(v float64) *GpsMetadata {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Heading = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Heading = uint16(unscaled)
	return m
}

// SetUtcTimestamp sets UtcTimestamp value.
//
// Units: s; Used to correlate UTC to system time if the timestamp of the message is in system time. This UTC time is derived from the GPS data.
func (m *GpsMetadata) SetUtcTimestamp(v time.Time) *GpsMetadata {
	m.UtcTimestamp = v
	return m
}

// SetVelocity sets Velocity value.
//
// Array: [3]; Scale: 100; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.
func (m *GpsMetadata) SetVelocity(v [3]int16) *GpsMetadata {
	m.Velocity = v
	return m
}

// SetVelocityScaled is similar to SetVelocity except it accepts a scaled value.
// This method automatically converts the given value to its [3]int16 form, discarding any applied scale and offset.
//
// Array: [3]; Scale: 100; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.
func (m *GpsMetadata) SetVelocityScaled(vs [3]float64) *GpsMetadata {
	m.Velocity = [3]int16{
		basetype.Sint16Invalid,
		basetype.Sint16Invalid,
		basetype.Sint16Invalid,
	}
	for i := range vs {
		unscaled := (vs[i] + 0) * 100
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			continue
		}
		m.Velocity[i] = int16(unscaled)
	}
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *GpsMetadata) SetUnknownFields(unknownFields ...proto.Field) *GpsMetadata {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *GpsMetadata) SetDeveloperFields(developerFields ...proto.DeveloperField) *GpsMetadata {
	m.DeveloperFields = developerFields
	return m
}
