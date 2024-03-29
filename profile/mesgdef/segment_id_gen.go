// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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

// SegmentId is a SegmentId message.
type SegmentId struct {
	Name                  string                       // Friendly name assigned to segment
	Uuid                  string                       // UUID of the segment
	UserProfilePrimaryKey uint32                       // Primary key of the user that created the segment
	DeviceId              uint32                       // ID of the device that created the segment
	Sport                 typedef.Sport                // Sport associated with the segment
	DefaultRaceLeader     uint8                        // Index for the Leader Board entry selected as the default race participant
	DeleteStatus          typedef.SegmentDeleteStatus  // Indicates if any segments should be deleted
	SelectionType         typedef.SegmentSelectionType // Indicates how the segment was selected to be sent to the device
	Enabled               bool                         // Segment enabled for evaluation

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentId creates new SegmentId struct based on given mesg.
// If mesg is nil, it will return SegmentId with all fields being set to its corresponding invalid value.
func NewSegmentId(mesg *proto.Message) *SegmentId {
	vals := [9]any{}

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
		Name:                  typeconv.ToString[string](vals[0]),
		Uuid:                  typeconv.ToString[string](vals[1]),
		UserProfilePrimaryKey: typeconv.ToUint32[uint32](vals[4]),
		DeviceId:              typeconv.ToUint32[uint32](vals[5]),
		Sport:                 typeconv.ToEnum[typedef.Sport](vals[2]),
		DefaultRaceLeader:     typeconv.ToUint8[uint8](vals[6]),
		DeleteStatus:          typeconv.ToEnum[typedef.SegmentDeleteStatus](vals[7]),
		SelectionType:         typeconv.ToEnum[typedef.SegmentSelectionType](vals[8]),
		Enabled:               typeconv.ToBool[bool](vals[3]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentId into proto.Message.
func (m *SegmentId) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSegmentId)

	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Name
		fields = append(fields, field)
	}
	if m.Uuid != basetype.StringInvalid && m.Uuid != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Uuid
		fields = append(fields, field)
	}
	if m.UserProfilePrimaryKey != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.UserProfilePrimaryKey
		fields = append(fields, field)
	}
	if m.DeviceId != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.DeviceId
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = byte(m.Sport)
		fields = append(fields, field)
	}
	if m.DefaultRaceLeader != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.DefaultRaceLeader
		fields = append(fields, field)
	}
	if byte(m.DeleteStatus) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = byte(m.DeleteStatus)
		fields = append(fields, field)
	}
	if byte(m.SelectionType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = byte(m.SelectionType)
		fields = append(fields, field)
	}
	if m.Enabled != false {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Enabled
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetName sets SegmentId value.
//
// Friendly name assigned to segment
func (m *SegmentId) SetName(v string) *SegmentId {
	m.Name = v
	return m
}

// SetUuid sets SegmentId value.
//
// UUID of the segment
func (m *SegmentId) SetUuid(v string) *SegmentId {
	m.Uuid = v
	return m
}

// SetUserProfilePrimaryKey sets SegmentId value.
//
// Primary key of the user that created the segment
func (m *SegmentId) SetUserProfilePrimaryKey(v uint32) *SegmentId {
	m.UserProfilePrimaryKey = v
	return m
}

// SetDeviceId sets SegmentId value.
//
// ID of the device that created the segment
func (m *SegmentId) SetDeviceId(v uint32) *SegmentId {
	m.DeviceId = v
	return m
}

// SetSport sets SegmentId value.
//
// Sport associated with the segment
func (m *SegmentId) SetSport(v typedef.Sport) *SegmentId {
	m.Sport = v
	return m
}

// SetDefaultRaceLeader sets SegmentId value.
//
// Index for the Leader Board entry selected as the default race participant
func (m *SegmentId) SetDefaultRaceLeader(v uint8) *SegmentId {
	m.DefaultRaceLeader = v
	return m
}

// SetDeleteStatus sets SegmentId value.
//
// Indicates if any segments should be deleted
func (m *SegmentId) SetDeleteStatus(v typedef.SegmentDeleteStatus) *SegmentId {
	m.DeleteStatus = v
	return m
}

// SetSelectionType sets SegmentId value.
//
// Indicates how the segment was selected to be sent to the device
func (m *SegmentId) SetSelectionType(v typedef.SegmentSelectionType) *SegmentId {
	m.SelectionType = v
	return m
}

// SetEnabled sets SegmentId value.
//
// Segment enabled for evaluation
func (m *SegmentId) SetEnabled(v bool) *SegmentId {
	m.Enabled = v
	return m
}

// SetDeveloperFields SegmentId's DeveloperFields.
func (m *SegmentId) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentId {
	m.DeveloperFields = developerFields
	return m
}
