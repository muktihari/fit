// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SegmentFile is a SegmentFile message.
type SegmentFile struct {
	MessageIndex           typedef.MessageIndex
	FileUuid               string                           // UUID of the segment file
	Enabled                bool                             // Enabled state of the segment file
	UserProfilePrimaryKey  uint32                           // Primary key of the user that created the segment file
	LeaderType             []typedef.SegmentLeaderboardType // Array: [N]; Leader type of each leader in the segment file
	LeaderGroupPrimaryKey  []uint32                         // Array: [N]; Group primary key of each leader in the segment file
	LeaderActivityId       []uint32                         // Array: [N]; Activity ID of each leader in the segment file
	LeaderActivityIdString string                           // String version of the activity ID of each leader in the segment file. 21 characters long for each ID, express in decimal
	DefaultRaceLeader      uint8                            // Index for the Leader Board entry selected as the default race participant

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentFile creates new SegmentFile struct based on given mesg. If mesg is nil or mesg.Num is not equal to SegmentFile mesg number, it will return nil.
func NewSegmentFile(mesg proto.Message) *SegmentFile {
	if mesg.Num != typedef.MesgNumSegmentFile {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		254: nil, /* MessageIndex */
		1:   nil, /* FileUuid */
		3:   nil, /* Enabled */
		4:   nil, /* UserProfilePrimaryKey */
		7:   nil, /* LeaderType */
		8:   nil, /* LeaderGroupPrimaryKey */
		9:   nil, /* LeaderActivityId */
		10:  nil, /* LeaderActivityIdString */
		11:  nil, /* DefaultRaceLeader */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &SegmentFile{
		MessageIndex:           typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		FileUuid:               typeconv.ToString[string](vals[1]),
		Enabled:                typeconv.ToBool[bool](vals[3]),
		UserProfilePrimaryKey:  typeconv.ToUint32[uint32](vals[4]),
		LeaderType:             typeconv.ToSliceEnum[typedef.SegmentLeaderboardType](vals[7]),
		LeaderGroupPrimaryKey:  typeconv.ToSliceUint32[uint32](vals[8]),
		LeaderActivityId:       typeconv.ToSliceUint32[uint32](vals[9]),
		LeaderActivityIdString: typeconv.ToString[string](vals[10]),
		DefaultRaceLeader:      typeconv.ToUint8[uint8](vals[11]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SegmentFile mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSegmentFile)
func (m *SegmentFile) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSegmentFile {
		return
	}

	vals := [...]any{
		254: typeconv.ToUint16[uint16](m.MessageIndex),
		1:   m.FileUuid,
		3:   m.Enabled,
		4:   m.UserProfilePrimaryKey,
		7:   typeconv.ToSliceEnum[byte](m.LeaderType),
		8:   m.LeaderGroupPrimaryKey,
		9:   m.LeaderActivityId,
		10:  m.LeaderActivityIdString,
		11:  m.DefaultRaceLeader,
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		field.Value = vals[field.Num]
	}

	mesg.DeveloperFields = m.DeveloperFields
}
