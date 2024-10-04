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

// Connectivity is a Connectivity message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Connectivity struct {
	Name                        string
	BluetoothEnabled            typedef.Bool // Use Bluetooth for connectivity features
	BluetoothLeEnabled          typedef.Bool // Use Bluetooth Low Energy for connectivity features
	AntEnabled                  typedef.Bool // Use ANT for connectivity features
	LiveTrackingEnabled         typedef.Bool
	WeatherConditionsEnabled    typedef.Bool
	WeatherAlertsEnabled        typedef.Bool
	AutoActivityUploadEnabled   typedef.Bool
	CourseDownloadEnabled       typedef.Bool
	WorkoutDownloadEnabled      typedef.Bool
	GpsEphemerisDownloadEnabled typedef.Bool
	IncidentDetectionEnabled    typedef.Bool
	GrouptrackEnabled           typedef.Bool

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewConnectivity creates new Connectivity struct based on given mesg.
// If mesg is nil, it will return Connectivity with all fields being set to its corresponding invalid value.
func NewConnectivity(mesg *proto.Message) *Connectivity {
	vals := [13]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 12 || mesg.Fields[i].Name == factory.NameUnknown {
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

	return &Connectivity{
		BluetoothEnabled:            vals[0].Bool(),
		BluetoothLeEnabled:          vals[1].Bool(),
		AntEnabled:                  vals[2].Bool(),
		Name:                        vals[3].String(),
		LiveTrackingEnabled:         vals[4].Bool(),
		WeatherConditionsEnabled:    vals[5].Bool(),
		WeatherAlertsEnabled:        vals[6].Bool(),
		AutoActivityUploadEnabled:   vals[7].Bool(),
		CourseDownloadEnabled:       vals[8].Bool(),
		WorkoutDownloadEnabled:      vals[9].Bool(),
		GpsEphemerisDownloadEnabled: vals[10].Bool(),
		IncidentDetectionEnabled:    vals[11].Bool(),
		GrouptrackEnabled:           vals[12].Bool(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Connectivity into proto.Message. If options is nil, default options will be used.
func (m *Connectivity) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumConnectivity}

	if m.BluetoothEnabled < 2 {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Bool(m.BluetoothEnabled)
		fields = append(fields, field)
	}
	if m.BluetoothLeEnabled < 2 {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Bool(m.BluetoothLeEnabled)
		fields = append(fields, field)
	}
	if m.AntEnabled < 2 {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Bool(m.AntEnabled)
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if m.LiveTrackingEnabled < 2 {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Bool(m.LiveTrackingEnabled)
		fields = append(fields, field)
	}
	if m.WeatherConditionsEnabled < 2 {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Bool(m.WeatherConditionsEnabled)
		fields = append(fields, field)
	}
	if m.WeatherAlertsEnabled < 2 {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Bool(m.WeatherAlertsEnabled)
		fields = append(fields, field)
	}
	if m.AutoActivityUploadEnabled < 2 {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Bool(m.AutoActivityUploadEnabled)
		fields = append(fields, field)
	}
	if m.CourseDownloadEnabled < 2 {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Bool(m.CourseDownloadEnabled)
		fields = append(fields, field)
	}
	if m.WorkoutDownloadEnabled < 2 {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Bool(m.WorkoutDownloadEnabled)
		fields = append(fields, field)
	}
	if m.GpsEphemerisDownloadEnabled < 2 {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Bool(m.GpsEphemerisDownloadEnabled)
		fields = append(fields, field)
	}
	if m.IncidentDetectionEnabled < 2 {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Bool(m.IncidentDetectionEnabled)
		fields = append(fields, field)
	}
	if m.GrouptrackEnabled < 2 {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Bool(m.GrouptrackEnabled)
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

// SetBluetoothEnabled sets BluetoothEnabled value.
//
// Use Bluetooth for connectivity features
func (m *Connectivity) SetBluetoothEnabled(v typedef.Bool) *Connectivity {
	m.BluetoothEnabled = v
	return m
}

// SetBluetoothLeEnabled sets BluetoothLeEnabled value.
//
// Use Bluetooth Low Energy for connectivity features
func (m *Connectivity) SetBluetoothLeEnabled(v typedef.Bool) *Connectivity {
	m.BluetoothLeEnabled = v
	return m
}

// SetAntEnabled sets AntEnabled value.
//
// Use ANT for connectivity features
func (m *Connectivity) SetAntEnabled(v typedef.Bool) *Connectivity {
	m.AntEnabled = v
	return m
}

// SetName sets Name value.
func (m *Connectivity) SetName(v string) *Connectivity {
	m.Name = v
	return m
}

// SetLiveTrackingEnabled sets LiveTrackingEnabled value.
func (m *Connectivity) SetLiveTrackingEnabled(v typedef.Bool) *Connectivity {
	m.LiveTrackingEnabled = v
	return m
}

// SetWeatherConditionsEnabled sets WeatherConditionsEnabled value.
func (m *Connectivity) SetWeatherConditionsEnabled(v typedef.Bool) *Connectivity {
	m.WeatherConditionsEnabled = v
	return m
}

// SetWeatherAlertsEnabled sets WeatherAlertsEnabled value.
func (m *Connectivity) SetWeatherAlertsEnabled(v typedef.Bool) *Connectivity {
	m.WeatherAlertsEnabled = v
	return m
}

// SetAutoActivityUploadEnabled sets AutoActivityUploadEnabled value.
func (m *Connectivity) SetAutoActivityUploadEnabled(v typedef.Bool) *Connectivity {
	m.AutoActivityUploadEnabled = v
	return m
}

// SetCourseDownloadEnabled sets CourseDownloadEnabled value.
func (m *Connectivity) SetCourseDownloadEnabled(v typedef.Bool) *Connectivity {
	m.CourseDownloadEnabled = v
	return m
}

// SetWorkoutDownloadEnabled sets WorkoutDownloadEnabled value.
func (m *Connectivity) SetWorkoutDownloadEnabled(v typedef.Bool) *Connectivity {
	m.WorkoutDownloadEnabled = v
	return m
}

// SetGpsEphemerisDownloadEnabled sets GpsEphemerisDownloadEnabled value.
func (m *Connectivity) SetGpsEphemerisDownloadEnabled(v typedef.Bool) *Connectivity {
	m.GpsEphemerisDownloadEnabled = v
	return m
}

// SetIncidentDetectionEnabled sets IncidentDetectionEnabled value.
func (m *Connectivity) SetIncidentDetectionEnabled(v typedef.Bool) *Connectivity {
	m.IncidentDetectionEnabled = v
	return m
}

// SetGrouptrackEnabled sets GrouptrackEnabled value.
func (m *Connectivity) SetGrouptrackEnabled(v typedef.Bool) *Connectivity {
	m.GrouptrackEnabled = v
	return m
}

// SetUnknownFields Connectivity's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Connectivity) SetUnknownFields(unknownFields ...proto.Field) *Connectivity {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Connectivity's DeveloperFields.
func (m *Connectivity) SetDeveloperFields(developerFields ...proto.DeveloperField) *Connectivity {
	m.DeveloperFields = developerFields
	return m
}
