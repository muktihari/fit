// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

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

// SegmentFile is a SegmentFile message.
type SegmentFile struct {
	FileUuid               string                           // UUID of the segment file
	LeaderType             []typedef.SegmentLeaderboardType // Array: [N]; Leader type of each leader in the segment file
	LeaderGroupPrimaryKey  []uint32                         // Array: [N]; Group primary key of each leader in the segment file
	LeaderActivityId       []uint32                         // Array: [N]; Activity ID of each leader in the segment file
	LeaderActivityIdString []string                         // Array: [N]; String version of the activity ID of each leader in the segment file. 21 characters long for each ID, express in decimal
	UserProfilePrimaryKey  uint32                           // Primary key of the user that created the segment file
	MessageIndex           typedef.MessageIndex
	DefaultRaceLeader      uint8 // Index for the Leader Board entry selected as the default race participant
	Enabled                bool  // Enabled state of the segment file

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentFile creates new SegmentFile struct based on given mesg.
// If mesg is nil, it will return SegmentFile with all fields being set to its corresponding invalid value.
func NewSegmentFile(mesg *proto.Message) *SegmentFile {
	vals := [255]any{}

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

	return &SegmentFile{
		FileUuid:               typeconv.ToString[string](vals[1]),
		LeaderType:             typeconv.ToSliceEnum[typedef.SegmentLeaderboardType](vals[7]),
		LeaderGroupPrimaryKey:  typeconv.ToSliceUint32[uint32](vals[8]),
		LeaderActivityId:       typeconv.ToSliceUint32[uint32](vals[9]),
		LeaderActivityIdString: typeconv.ToSliceString[string](vals[10]),
		UserProfilePrimaryKey:  typeconv.ToUint32[uint32](vals[4]),
		MessageIndex:           typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		DefaultRaceLeader:      typeconv.ToUint8[uint8](vals[11]),
		Enabled:                typeconv.ToBool[bool](vals[3]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentFile into proto.Message.
func (m *SegmentFile) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSegmentFile)

	if m.FileUuid != basetype.StringInvalid && m.FileUuid != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.FileUuid
		fields = append(fields, field)
	}
	if typeconv.ToSliceEnum[byte](m.LeaderType) != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToSliceEnum[byte](m.LeaderType)
		fields = append(fields, field)
	}
	if m.LeaderGroupPrimaryKey != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.LeaderGroupPrimaryKey
		fields = append(fields, field)
	}
	if m.LeaderActivityId != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.LeaderActivityId
		fields = append(fields, field)
	}
	if m.LeaderActivityIdString != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.LeaderActivityIdString
		fields = append(fields, field)
	}
	if m.UserProfilePrimaryKey != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.UserProfilePrimaryKey
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if m.DefaultRaceLeader != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.DefaultRaceLeader
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

// SetFileUuid sets SegmentFile value.
//
// UUID of the segment file
func (m *SegmentFile) SetFileUuid(v string) *SegmentFile {
	m.FileUuid = v
	return m
}

// SetLeaderType sets SegmentFile value.
//
// Array: [N]; Leader type of each leader in the segment file
func (m *SegmentFile) SetLeaderType(v []typedef.SegmentLeaderboardType) *SegmentFile {
	m.LeaderType = v
	return m
}

// SetLeaderGroupPrimaryKey sets SegmentFile value.
//
// Array: [N]; Group primary key of each leader in the segment file
func (m *SegmentFile) SetLeaderGroupPrimaryKey(v []uint32) *SegmentFile {
	m.LeaderGroupPrimaryKey = v
	return m
}

// SetLeaderActivityId sets SegmentFile value.
//
// Array: [N]; Activity ID of each leader in the segment file
func (m *SegmentFile) SetLeaderActivityId(v []uint32) *SegmentFile {
	m.LeaderActivityId = v
	return m
}

// SetLeaderActivityIdString sets SegmentFile value.
//
// Array: [N]; String version of the activity ID of each leader in the segment file. 21 characters long for each ID, express in decimal
func (m *SegmentFile) SetLeaderActivityIdString(v []string) *SegmentFile {
	m.LeaderActivityIdString = v
	return m
}

// SetUserProfilePrimaryKey sets SegmentFile value.
//
// Primary key of the user that created the segment file
func (m *SegmentFile) SetUserProfilePrimaryKey(v uint32) *SegmentFile {
	m.UserProfilePrimaryKey = v
	return m
}

// SetMessageIndex sets SegmentFile value.
func (m *SegmentFile) SetMessageIndex(v typedef.MessageIndex) *SegmentFile {
	m.MessageIndex = v
	return m
}

// SetDefaultRaceLeader sets SegmentFile value.
//
// Index for the Leader Board entry selected as the default race participant
func (m *SegmentFile) SetDefaultRaceLeader(v uint8) *SegmentFile {
	m.DefaultRaceLeader = v
	return m
}

// SetEnabled sets SegmentFile value.
//
// Enabled state of the segment file
func (m *SegmentFile) SetEnabled(v bool) *SegmentFile {
	m.Enabled = v
	return m
}

// SetDeveloperFields SegmentFile's DeveloperFields.
func (m *SegmentFile) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentFile {
	m.DeveloperFields = developerFields
	return m
}
