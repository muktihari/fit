// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"unsafe"
)

// SegmentFile is a SegmentFile message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SegmentFile struct {
	FileUuid               string                           // UUID of the segment file
	LeaderType             []typedef.SegmentLeaderboardType // Array: [N]; Leader type of each leader in the segment file
	LeaderGroupPrimaryKey  []uint32                         // Array: [N]; Group primary key of each leader in the segment file
	LeaderActivityId       []uint32                         // Array: [N]; Activity ID of each leader in the segment file
	LeaderActivityIdString []string                         // Array: [N]; String version of the activity ID of each leader in the segment file. 21 characters long for each ID, express in decimal
	UserProfilePrimaryKey  uint32                           // Primary key of the user that created the segment file
	MessageIndex           typedef.MessageIndex
	Enabled                typedef.Bool // Enabled state of the segment file
	DefaultRaceLeader      uint8        // Index for the Leader Board entry selected as the default race participant

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSegmentFile creates new SegmentFile struct based on given mesg.
// If mesg is nil, it will return SegmentFile with all fields being set to its corresponding invalid value.
func NewSegmentFile(mesg *proto.Message) *SegmentFile {
	m := new(SegmentFile)
	m.Reset(mesg)
	return m
}

// Reset resets all SegmentFile's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *SegmentFile) Reset(mesg *proto.Message) {
	var (
		vals            [255]proto.Value
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
			if mesg.Fields[i].Num < 255 {
				vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
			}
		}
		developerFields = mesg.DeveloperFields
	}

	*m = SegmentFile{
		MessageIndex:          typedef.MessageIndex(vals[254].Uint16()),
		FileUuid:              vals[1].String(),
		Enabled:               vals[3].Bool(),
		UserProfilePrimaryKey: vals[4].Uint32(),
		LeaderType: func() []typedef.SegmentLeaderboardType {
			sliceValue := vals[7].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.SegmentLeaderboardType)(ptr), len(sliceValue))
		}(),
		LeaderGroupPrimaryKey:  vals[8].SliceUint32(),
		LeaderActivityId:       vals[9].SliceUint32(),
		LeaderActivityIdString: vals[10].SliceString(),
		DefaultRaceLeader:      vals[11].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentFile into proto.Message. If options is nil, default options will be used.
func (m *SegmentFile) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fields := make([]proto.Field, 0, 9)
	mesg := proto.Message{Num: typedef.MesgNumSegmentFile}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.FileUuid != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.FileUuid)
		fields = append(fields, field)
	}
	if m.Enabled < 2 {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if m.UserProfilePrimaryKey != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.UserProfilePrimaryKey)
		fields = append(fields, field)
	}
	if m.LeaderType != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceUint8(m.LeaderType)
		fields = append(fields, field)
	}
	if m.LeaderGroupPrimaryKey != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceUint32(m.LeaderGroupPrimaryKey)
		fields = append(fields, field)
	}
	if m.LeaderActivityId != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.SliceUint32(m.LeaderActivityId)
		fields = append(fields, field)
	}
	if m.LeaderActivityIdString != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.SliceString(m.LeaderActivityIdString)
		fields = append(fields, field)
	}
	if m.DefaultRaceLeader != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.DefaultRaceLeader)
		fields = append(fields, field)
	}

	n := len(fields)
	mesg.Fields = make([]proto.Field, n+len(m.UnknownFields))
	copy(mesg.Fields[:n], fields)
	copy(mesg.Fields[n:], m.UnknownFields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
func (m *SegmentFile) SetMessageIndex(v typedef.MessageIndex) *SegmentFile {
	m.MessageIndex = v
	return m
}

// SetFileUuid sets FileUuid value.
//
// UUID of the segment file
func (m *SegmentFile) SetFileUuid(v string) *SegmentFile {
	m.FileUuid = v
	return m
}

// SetEnabled sets Enabled value.
//
// Enabled state of the segment file
func (m *SegmentFile) SetEnabled(v typedef.Bool) *SegmentFile {
	m.Enabled = v
	return m
}

// SetUserProfilePrimaryKey sets UserProfilePrimaryKey value.
//
// Primary key of the user that created the segment file
func (m *SegmentFile) SetUserProfilePrimaryKey(v uint32) *SegmentFile {
	m.UserProfilePrimaryKey = v
	return m
}

// SetLeaderType sets LeaderType value.
//
// Array: [N]; Leader type of each leader in the segment file
func (m *SegmentFile) SetLeaderType(v []typedef.SegmentLeaderboardType) *SegmentFile {
	m.LeaderType = v
	return m
}

// SetLeaderGroupPrimaryKey sets LeaderGroupPrimaryKey value.
//
// Array: [N]; Group primary key of each leader in the segment file
func (m *SegmentFile) SetLeaderGroupPrimaryKey(v []uint32) *SegmentFile {
	m.LeaderGroupPrimaryKey = v
	return m
}

// SetLeaderActivityId sets LeaderActivityId value.
//
// Array: [N]; Activity ID of each leader in the segment file
func (m *SegmentFile) SetLeaderActivityId(v []uint32) *SegmentFile {
	m.LeaderActivityId = v
	return m
}

// SetLeaderActivityIdString sets LeaderActivityIdString value.
//
// Array: [N]; String version of the activity ID of each leader in the segment file. 21 characters long for each ID, express in decimal
func (m *SegmentFile) SetLeaderActivityIdString(v []string) *SegmentFile {
	m.LeaderActivityIdString = v
	return m
}

// SetDefaultRaceLeader sets DefaultRaceLeader value.
//
// Index for the Leader Board entry selected as the default race participant
func (m *SegmentFile) SetDefaultRaceLeader(v uint8) *SegmentFile {
	m.DefaultRaceLeader = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SegmentFile) SetUnknownFields(unknownFields ...proto.Field) *SegmentFile {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *SegmentFile) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentFile {
	m.DeveloperFields = developerFields
	return m
}
