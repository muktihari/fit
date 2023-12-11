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

// SegmentLeaderboardEntry is a SegmentLeaderboardEntry message.
type SegmentLeaderboardEntry struct {
	MessageIndex     typedef.MessageIndex
	Name             string                         // Friendly name assigned to leader
	Type             typedef.SegmentLeaderboardType // Leader classification
	GroupPrimaryKey  uint32                         // Primary user ID of this leader
	ActivityId       uint32                         // ID of the activity associated with this leader time
	SegmentTime      uint32                         // Scale: 1000; Units: s; Segment Time (includes pauses)
	ActivityIdString string                         // String version of the activity_id. 21 characters long, express in decimal

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentLeaderboardEntry creates new SegmentLeaderboardEntry struct based on given mesg. If mesg is nil or mesg.Num is not equal to SegmentLeaderboardEntry mesg number, it will return nil.
func NewSegmentLeaderboardEntry(mesg proto.Message) *SegmentLeaderboardEntry {
	if mesg.Num != typedef.MesgNumSegmentLeaderboardEntry {
		return nil
	}

	vals := [...]any{ // nil value will be converted to its corresponding invalid value by typeconv.
		254: nil, /* MessageIndex */
		0:   nil, /* Name */
		1:   nil, /* Type */
		2:   nil, /* GroupPrimaryKey */
		3:   nil, /* ActivityId */
		4:   nil, /* SegmentTime */
		5:   nil, /* ActivityIdString */
	}

	for i := range mesg.Fields {
		field := &mesg.Fields[i]
		if field.Num >= byte(len(vals)) {
			continue
		}
		vals[field.Num] = field.Value
	}

	return &SegmentLeaderboardEntry{
		MessageIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Name:             typeconv.ToString[string](vals[0]),
		Type:             typeconv.ToEnum[typedef.SegmentLeaderboardType](vals[1]),
		GroupPrimaryKey:  typeconv.ToUint32[uint32](vals[2]),
		ActivityId:       typeconv.ToUint32[uint32](vals[3]),
		SegmentTime:      typeconv.ToUint32[uint32](vals[4]),
		ActivityIdString: typeconv.ToString[string](vals[5]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SegmentLeaderboardEntry mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSegmentLeaderboardEntry)
func (m SegmentLeaderboardEntry) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSegmentLeaderboardEntry {
		return
	}

	vals := [...]any{
		254: m.MessageIndex,
		0:   m.Name,
		1:   m.Type,
		2:   m.GroupPrimaryKey,
		3:   m.ActivityId,
		4:   m.SegmentTime,
		5:   m.ActivityIdString,
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
