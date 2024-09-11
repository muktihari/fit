// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestLocalMesgNum(t *testing.T) {
	tt := []struct {
		header       byte
		localMesgNum byte
	}{
		{
			header:       proto.MesgCompressedHeaderMask | 0b00111111,
			localMesgNum: 1,
		},
		{
			header:       proto.MesgNormalHeaderMask | 2,
			localMesgNum: 2,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("0b%08b", tc.header), func(t *testing.T) {
			localMesgNum := proto.LocalMesgNum(tc.header)
			if localMesgNum != tc.localMesgNum {
				t.Fatalf("expected: %d, got: %d", tc.localMesgNum, localMesgNum)
			}
		})
	}
}

func TestFitWithMessages(t *testing.T) {
	tt := []struct {
		name     string
		messages []proto.Message
	}{
		{
			name: "withMessages",
			messages: []proto.Message{
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed),
					factory.CreateField(mesgnum.Record, fieldnum.RecordCadence),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed),
					factory.CreateField(mesgnum.Record, fieldnum.RecordCadence),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate),
				}},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			fit := &proto.FIT{Messages: tc.messages}
			if diff := cmp.Diff(fit.Messages, tc.messages,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
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

func TestMessageFieldByNum(t *testing.T) {
	sharedField := factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart)

	tt := []struct {
		name     string
		mesg     proto.Message
		fieldNum byte
		field    *proto.Field
	}{
		{
			name: "FieldByNum found",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				sharedField,
			}},
			fieldNum: fieldnum.EventEventType,
			field:    &sharedField,
		},
		{
			name: "FieldByNum not found",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				sharedField,
			}},
			fieldNum: fieldnum.EventData,
			field:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			field := tc.mesg.FieldByNum(tc.fieldNum)
			if diff := cmp.Diff(tc.field, field,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageFieldValueByNum(t *testing.T) {
	tt := []struct {
		name     string
		mesg     proto.Message
		fieldNum byte
		value    proto.Value
	}{
		{
			name: "FieldValueByNum found",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			}},
			fieldNum: fieldnum.EventEventType,
			value:    proto.Uint8(uint8(typedef.EventTypeStart)),
		},
		{
			name: "FieldValueByNum not found",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			}},
			fieldNum: fieldnum.EventData,
			value:    proto.Value{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.mesg.FieldValueByNum(tc.fieldNum)
			if value.Any() != tc.value.Any() {
				t.Fatalf("expected value: %T(%v), got: %T(%v)", tc.value, tc.value, value, value)
			}
		})
	}
}

func TestMessageRemoveFieldByNum(t *testing.T) {
	tt := []struct {
		name     string
		mesg     proto.Message
		fieldNum byte
		field    *proto.Field
		size     int
	}{
		{
			name: "remove existing field",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			}},
			fieldNum: fieldnum.EventEventType,
			field:    nil,
			size:     0,
		},
		{
			name: "remove field that is not exist",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(typedef.EventTypeStart),
			}},
			fieldNum: fieldnum.EventData,
			field:    nil,
			size:     1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.mesg.RemoveFieldByNum(tc.fieldNum)
			field := tc.mesg.FieldByNum(tc.fieldNum)
			if diff := cmp.Diff(tc.field, field); diff != "" {
				t.Fatal(diff)
			}
			if len(tc.mesg.Fields) != tc.size {
				t.Fatalf("expected len after removal: %d, got: %d", tc.size, len(tc.mesg.Fields))
			}
		})
	}
}

func TestMessageClone(t *testing.T) {
	mesg := proto.Message{Num: mesgnum.Session, Fields: []proto.Field{
		factory.CreateField(mesgnum.Session, fieldnum.SessionAvgAltitude).WithValue(uint16(1000)),
		factory.CreateField(mesgnum.Session, fieldnum.SessionAvgSpeed).WithValue(uint16(1000)),
	},
		DeveloperFields: []proto.DeveloperField{
			{
				Num:                0,
				DeveloperDataIndex: 0,
				Value:              proto.Uint8(1),
			},
			{},
		}}

	cloned := mesg.Clone()
	cloned.Fields[0].Num = 100
	cloned.DeveloperFields[0].Num = 100

	if diff := cmp.Diff(mesg, cloned,
		cmp.Transformer("Value", func(v proto.Value) any {
			return v.Any()
		}),
	); diff == "" {
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
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(10)),
			}},
			field:        factory.CreateField(mesgnum.Event, fieldnum.EventData),
			subfieldName: "course_point_index",
			ok:           true,
		},
		{
			name: "SubFieldSubtitution not ok, can't interpret main field.",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(100)),
			}},
			field: factory.CreateField(mesgnum.Event, fieldnum.EventData),
			ok:    false,
		},
		{
			name: "SubFieldSubtitution field reference not found",
			mesg: proto.Message{Num: mesgnum.Event, Fields: []proto.Field{
				factory.CreateField(mesgnum.Event, fieldnum.EventActivityType).WithValue(uint8(10)),
			}},
			field: factory.CreateField(mesgnum.Event, fieldnum.EventData),
			ok:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			subfield := tc.field.SubFieldSubtitution(&tc.mesg)
			if subfield != nil != tc.ok {
				t.Fatalf("expected: %t, got: %t", tc.ok, subfield != nil)
			}
			if subfield == nil {
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

	if diff := cmp.Diff(field, cloned,
		cmp.Transformer("Value", func(v proto.Value) any {
			return v.Any()
		}),
	); diff == "" {
		t.Fatalf("expected deep cloned, but some data still being referenced.")
	}

	field = factory.CreateField(mesgnum.Session, fieldnum.SessionTotalCycles)
	cloned = field.Clone()
	field.SubFields[0].Name = "FIT SDK for Go"

	if diff := cmp.Diff(field, cloned,
		cmp.Transformer("Value", func(v proto.Value) any {
			return v.Any()
		}),
	); diff == "" {
		t.Fatalf("should not changed")
	}

	field = proto.Field{}
	cloned = field.Clone()

	if diff := cmp.Diff(field, cloned,
		cmp.Transformer("Value", func(v proto.Value) any {
			return v.Any()
		}),
	); diff != "" {
		t.Fatalf("empty field base, field should be returned as is: %v", diff)
	}
}
