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

// SegmentId is a SegmentId message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SegmentId struct {
	Name                  string                       // Friendly name assigned to segment
	Uuid                  string                       // UUID of the segment
	UserProfilePrimaryKey uint32                       // Primary key of the user that created the segment
	DeviceId              uint32                       // ID of the device that created the segment
	Sport                 typedef.Sport                // Sport associated with the segment
	Enabled               bool                         // Segment enabled for evaluation
	DefaultRaceLeader     uint8                        // Index for the Leader Board entry selected as the default race participant
	DeleteStatus          typedef.SegmentDeleteStatus  // Indicates if any segments should be deleted
	SelectionType         typedef.SegmentSelectionType // Indicates how the segment was selected to be sent to the device

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentId creates new SegmentId struct based on given mesg.
// If mesg is nil, it will return SegmentId with all fields being set to its corresponding invalid value.
func NewSegmentId(mesg *proto.Message) *SegmentId {
	vals := [9]proto.Value{}

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

	return &SegmentId{
		Name:                  vals[0].String(),
		Uuid:                  vals[1].String(),
		Sport:                 typedef.Sport(vals[2].Uint8()),
		Enabled:               vals[3].Bool(),
		UserProfilePrimaryKey: vals[4].Uint32(),
		DeviceId:              vals[5].Uint32(),
		DefaultRaceLeader:     vals[6].Uint8(),
		DeleteStatus:          typedef.SegmentDeleteStatus(vals[7].Uint8()),
		SelectionType:         typedef.SegmentSelectionType(vals[8].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentId into proto.Message. If options is nil, default options will be used.
func (m *SegmentId) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumSegmentId}

	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if m.Uuid != basetype.StringInvalid && m.Uuid != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.Uuid)
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if m.Enabled != false {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if m.UserProfilePrimaryKey != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.UserProfilePrimaryKey)
		fields = append(fields, field)
	}
	if m.DeviceId != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.DeviceId)
		fields = append(fields, field)
	}
	if m.DefaultRaceLeader != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.DefaultRaceLeader)
		fields = append(fields, field)
	}
	if byte(m.DeleteStatus) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.DeleteStatus))
		fields = append(fields, field)
	}
	if byte(m.SelectionType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(byte(m.SelectionType))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetName sets Name value.
//
// Friendly name assigned to segment
func (m *SegmentId) SetName(v string) *SegmentId {
	m.Name = v
	return m
}

// SetUuid sets Uuid value.
//
// UUID of the segment
func (m *SegmentId) SetUuid(v string) *SegmentId {
	m.Uuid = v
	return m
}

// SetSport sets Sport value.
//
// Sport associated with the segment
func (m *SegmentId) SetSport(v typedef.Sport) *SegmentId {
	m.Sport = v
	return m
}

// SetEnabled sets Enabled value.
//
// Segment enabled for evaluation
func (m *SegmentId) SetEnabled(v bool) *SegmentId {
	m.Enabled = v
	return m
}

// SetUserProfilePrimaryKey sets UserProfilePrimaryKey value.
//
// Primary key of the user that created the segment
func (m *SegmentId) SetUserProfilePrimaryKey(v uint32) *SegmentId {
	m.UserProfilePrimaryKey = v
	return m
}

// SetDeviceId sets DeviceId value.
//
// ID of the device that created the segment
func (m *SegmentId) SetDeviceId(v uint32) *SegmentId {
	m.DeviceId = v
	return m
}

// SetDefaultRaceLeader sets DefaultRaceLeader value.
//
// Index for the Leader Board entry selected as the default race participant
func (m *SegmentId) SetDefaultRaceLeader(v uint8) *SegmentId {
	m.DefaultRaceLeader = v
	return m
}

// SetDeleteStatus sets DeleteStatus value.
//
// Indicates if any segments should be deleted
func (m *SegmentId) SetDeleteStatus(v typedef.SegmentDeleteStatus) *SegmentId {
	m.DeleteStatus = v
	return m
}

// SetSelectionType sets SelectionType value.
//
// Indicates how the segment was selected to be sent to the device
func (m *SegmentId) SetSelectionType(v typedef.SegmentSelectionType) *SegmentId {
	m.SelectionType = v
	return m
}

// SetDeveloperFields SegmentId's DeveloperFields.
func (m *SegmentId) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentId {
	m.DeveloperFields = developerFields
	return m
}
