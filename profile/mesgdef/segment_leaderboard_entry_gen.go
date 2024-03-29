// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SegmentLeaderboardEntry is a SegmentLeaderboardEntry message.
type SegmentLeaderboardEntry struct {
	Name             string // Friendly name assigned to leader
	ActivityIdString string // String version of the activity_id. 21 characters long, express in decimal
	GroupPrimaryKey  uint32 // Primary user ID of this leader
	ActivityId       uint32 // ID of the activity associated with this leader time
	SegmentTime      uint32 // Scale: 1000; Units: s; Segment Time (includes pauses)
	MessageIndex     typedef.MessageIndex
	Type             typedef.SegmentLeaderboardType // Leader classification

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentLeaderboardEntry creates new SegmentLeaderboardEntry struct based on given mesg.
// If mesg is nil, it will return SegmentLeaderboardEntry with all fields being set to its corresponding invalid value.
func NewSegmentLeaderboardEntry(mesg *proto.Message) *SegmentLeaderboardEntry {
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

	return &SegmentLeaderboardEntry{
		Name:             typeconv.ToString[string](vals[0]),
		ActivityIdString: typeconv.ToString[string](vals[5]),
		GroupPrimaryKey:  typeconv.ToUint32[uint32](vals[2]),
		ActivityId:       typeconv.ToUint32[uint32](vals[3]),
		SegmentTime:      typeconv.ToUint32[uint32](vals[4]),
		MessageIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Type:             typeconv.ToEnum[typedef.SegmentLeaderboardType](vals[1]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentLeaderboardEntry into proto.Message.
func (m *SegmentLeaderboardEntry) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSegmentLeaderboardEntry)

	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Name
		fields = append(fields, field)
	}
	if m.ActivityIdString != basetype.StringInvalid && m.ActivityIdString != "" {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.ActivityIdString
		fields = append(fields, field)
	}
	if m.GroupPrimaryKey != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.GroupPrimaryKey
		fields = append(fields, field)
	}
	if m.ActivityId != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.ActivityId
		fields = append(fields, field)
	}
	if m.SegmentTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.SegmentTime
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = byte(m.Type)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SegmentTimeScaled return SegmentTime in its scaled value [Scale: 1000; Units: s; Segment Time (includes pauses)].
//
// If SegmentTime value is invalid, float64 invalid value will be returned.
func (m *SegmentLeaderboardEntry) SegmentTimeScaled() float64 {
	if m.SegmentTime == basetype.Uint32Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.SegmentTime, 1000, 0)
}

// SetName sets SegmentLeaderboardEntry value.
//
// Friendly name assigned to leader
func (m *SegmentLeaderboardEntry) SetName(v string) *SegmentLeaderboardEntry {
	m.Name = v
	return m
}

// SetActivityIdString sets SegmentLeaderboardEntry value.
//
// String version of the activity_id. 21 characters long, express in decimal
func (m *SegmentLeaderboardEntry) SetActivityIdString(v string) *SegmentLeaderboardEntry {
	m.ActivityIdString = v
	return m
}

// SetGroupPrimaryKey sets SegmentLeaderboardEntry value.
//
// Primary user ID of this leader
func (m *SegmentLeaderboardEntry) SetGroupPrimaryKey(v uint32) *SegmentLeaderboardEntry {
	m.GroupPrimaryKey = v
	return m
}

// SetActivityId sets SegmentLeaderboardEntry value.
//
// ID of the activity associated with this leader time
func (m *SegmentLeaderboardEntry) SetActivityId(v uint32) *SegmentLeaderboardEntry {
	m.ActivityId = v
	return m
}

// SetSegmentTime sets SegmentLeaderboardEntry value.
//
// Scale: 1000; Units: s; Segment Time (includes pauses)
func (m *SegmentLeaderboardEntry) SetSegmentTime(v uint32) *SegmentLeaderboardEntry {
	m.SegmentTime = v
	return m
}

// SetMessageIndex sets SegmentLeaderboardEntry value.
func (m *SegmentLeaderboardEntry) SetMessageIndex(v typedef.MessageIndex) *SegmentLeaderboardEntry {
	m.MessageIndex = v
	return m
}

// SetType sets SegmentLeaderboardEntry value.
//
// Leader classification
func (m *SegmentLeaderboardEntry) SetType(v typedef.SegmentLeaderboardType) *SegmentLeaderboardEntry {
	m.Type = v
	return m
}

// SetDeveloperFields SegmentLeaderboardEntry's DeveloperFields.
func (m *SegmentLeaderboardEntry) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentLeaderboardEntry {
	m.DeveloperFields = developerFields
	return m
}
