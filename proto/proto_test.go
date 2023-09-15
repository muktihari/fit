// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestFitWithMessages(t *testing.T) {
	tt := []struct {
		name     string
		messages []proto.Message
	}{
		{
			name: "withMessages",
			messages: []proto.Message{
				factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed),
					factory.CreateField(mesgnum.Record, fieldnum.RecordCadence),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate),
				),
				factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed),
					factory.CreateField(mesgnum.Record, fieldnum.RecordCadence),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate),
				),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			fit := new(proto.Fit).WithMessages(tc.messages...)
			if diff := cmp.Diff(fit.Messages, tc.messages); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageDefinitionClone(t *testing.T) {
	mesgDef := proto.MessageDefinition{
		FieldDefinitions: []proto.FieldDefinition{
			{Num: fieldnum.RecordCadence, Size: 1, BaseType: basetype.Uint8},
			{Num: fieldnum.RecordHeartRate, Size: 1, BaseType: basetype.Uint8},
		},
		DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
			{Num: 0, DeveloperDataIndex: 0, Size: 1},
		},
	}

	cloned := mesgDef.Clone()
	cloned.FieldDefinitions[0].Num = 100
	cloned.DeveloperFieldDefinitions[0].Num = 100

	if diff := cmp.Diff(mesgDef, cloned); diff == "" {
		t.Fatalf("expected deep cloned, but some data still being referenced.")
	}
}

func TestMessageWithFieldValues(t *testing.T) {
	tt := []struct {
		name        string
		fieldValues map[byte]any
	}{
		{
			name: "withFieldValues",
			fieldValues: map[byte]any{
				fieldnum.RecordSpeed:   uint16(1000),
				fieldnum.RecordCadence: uint16(100),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mesg := factory.CreateMesg(mesgnum.Record)
			mesg.WithFieldValues(tc.fieldValues)
			for i := range mesg.Fields {
				if value, ok := tc.fieldValues[mesg.Fields[i].Num]; ok {
					if mesg.Fields[i].Value != value {
						t.Errorf("expected %T(%v), got: %T(%v)", value, value, mesg.Fields[i].Value, mesg.Fields[i].Value)
					}
				}
			}
		})
	}
}

func TestMessageFieldByNum(t *testing.T) {
	tt := []struct {
		name     string
		mesg     proto.Message
		fieldNum byte
		ok       bool
	}{
		{
			name: "FieldByNum found",
			mesg: factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			),
			fieldNum: fieldnum.EventEventType,
			ok:       true,
		},
		{
			name: "FieldByNum not found",
			mesg: factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			),
			fieldNum: fieldnum.EventData,
			ok:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := tc.mesg.FieldByNum(tc.fieldNum)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
		})
	}
}

func TestMessageRemoveByNum(t *testing.T) {
	tt := []struct {
		name     string
		mesg     proto.Message
		fieldNum byte
		ok       bool
	}{
		{
			name: "RemoveByNum found",
			mesg: factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			),
			fieldNum: fieldnum.EventEventType,
			ok:       true,
		},
		{
			name: "FieldByNum not found",
			mesg: factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			),
			fieldNum: fieldnum.EventData,
			ok:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.mesg.RemoveFieldByNum(tc.fieldNum)
			_, ok := tc.mesg.FieldByNum(tc.fieldNum)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
		})
	}
}

func TestMessageClone(t *testing.T) {
	mesg := factory.CreateMesg(mesgnum.Session).WithFieldValues(map[byte]any{
		fieldnum.SessionAvgAltitude: uint16(1000),
		fieldnum.SessionAvgSpeed:    uint16(1000),
	}).WithDeveloperFields(
		proto.DeveloperField{
			Num:                0,
			DeveloperDataIndex: 0,
			Size:               1,
			Type:               basetype.Uint8,
			Value:              uint8(1),
		},
		proto.DeveloperField{},
	)

	cloned := mesg.Clone()
	cloned.Fields[0].Num = 100
	cloned.DeveloperFields[0].Num = 100

	if diff := cmp.Diff(mesg, cloned); diff == "" {
		t.Fatalf("expected deep cloned, but some data still being referenced.")
	}
}

func TestFieldSubFieldSubtitution(t *testing.T) {
	tt := []struct {
		name         string
		mesg         proto.Message
		field        proto.Field
		subfieldName string
		ok           bool
	}{
		{
			name: "SubFieldSubtitution ok, main field can be interpreted.",
			mesg: factory.CreateMesg(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(10)),
			),
			field:        factory.CreateField(mesgnum.Event, fieldnum.EventData),
			subfieldName: "course_point_index",
			ok:           true,
		},
		{
			name: "SubFieldSubtitution not ok, can't interpret main field.",
			mesg: factory.CreateMesg(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(100)),
			),
			field: factory.CreateField(mesgnum.Event, fieldnum.EventData),
			ok:    false,
		},
		{
			name: "SubFieldSubtitution field reference not found",
			mesg: factory.CreateMesg(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventActivityType).WithValue(uint8(10)),
			),
			field: factory.CreateField(mesgnum.Event, fieldnum.EventData),
			ok:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			subfield, ok := tc.field.SubFieldSubtitution(&tc.mesg)
			if ok != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, ok)
			}
			if !ok {
				return
			}
			if subfield.Name != tc.subfieldName {
				t.Fatalf("expected: %s, got: %s", tc.subfieldName, subfield.Name)
			}
		})
	}
}

func TestFieldClone(t *testing.T) {
	field := factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed)

	cloned := field.Clone()
	cloned.Components[0].Scale = 777

	if diff := cmp.Diff(field, cloned); diff == "" {
		t.Fatalf("expected deep cloned, but some data still being referenced.")
	}

	field = proto.Field{}
	cloned = field.Clone()

	if diff := cmp.Diff(field, cloned); diff != "" {
		t.Fatalf("should not changed")
	}
}

func TestCreateMessageDefinition(t *testing.T) {
	tt := []struct {
		name    string
		mesg    proto.Message
		mesgDef proto.MessageDefinition
	}{
		{
			name: "fields only with non-array values",
			mesg: factory.CreateMesgOnly(mesgnum.FileId).WithFields(
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
			),
			mesgDef: proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.FileId,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.FileIdType,
						Size:     1,
						BaseType: basetype.Enum,
					},
				},
			},
		},
		{
			name: "fields only with string value",
			mesg: factory.CreateMesgOnly(mesgnum.FileId).WithFields(
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("Fit SDK Go"),
			),
			mesgDef: proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.FileId,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.FileIdProductName,
						Size:     1 * 11, // len("Fit SDK Go") == 10 + '0x00'
						BaseType: basetype.String,
					},
				},
			},
		},
		{
			name: "fields only with array of byte",
			mesg: factory.CreateMesgOnly(mesgnum.UserProfile).WithFields(
				factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileGlobalId).WithValue([]byte{2, 9}),
			),
			mesgDef: proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.UserProfile,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.UserProfileGlobalId,
						Size:     2,
						BaseType: basetype.Byte,
					},
				},
			},
		},

		{
			name: "developer fields",
			mesg: factory.CreateMesgOnly(mesgnum.UserProfile).
				WithFields(
					factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileGlobalId).WithValue([]any{byte(2), byte(9)})).
				WithDeveloperFields(
					proto.DeveloperField{
						Num: 0, Name: "Fit SDK Go", Type: basetype.Byte, DeveloperDataIndex: 0, Value: byte(1),
					},
				),
			mesgDef: proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask | proto.DevDataMask,
				MesgNum: mesgnum.UserProfile,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.UserProfileGlobalId,
						Size:     2,
						BaseType: basetype.Byte,
					},
				},
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{
						Num: 0, Size: 1, DeveloperDataIndex: 0,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mesgDef := proto.CreateMessageDefinition(&tc.mesg)
			if diff := cmp.Diff(mesgDef, tc.mesgDef); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
