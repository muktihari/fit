// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
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

func TestFieldFormat(t *testing.T) {
	field := factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(7))

	tt := []struct {
		format   string
		expected string
	}{
		{
			format: "%f",
			expected: fmt.Sprintf("%%!f(proto.Field={(%p)(%v) %v %t})",
				field.FieldBase, field.FieldBase, field.Value, field.IsExpandedField),
		},
		{
			format: "%s",
			expected: fmt.Sprintf("%%!s(proto.Field={(%p)(%v) %v %t})",
				field.FieldBase, field.FieldBase, field.Value, field.IsExpandedField),
		},
		{
			format: "%v",
			expected: fmt.Sprintf("{(%p)(%v) %v %t}",
				field.FieldBase, field.FieldBase, field.Value, field.IsExpandedField),
		},
		{
			format: "%+v",
			expected: fmt.Sprintf("{FieldBase:(%p)(%+v) Value:%+v IsExpandedField:%t}",
				field.FieldBase, field.FieldBase, field.Value, field.IsExpandedField),
		},
		{
			format: "%#v",
			expected: fmt.Sprintf("{FieldBase:(%p)(%#v), Value:%#v, IsExpandedField:%t}",
				field.FieldBase, field.FieldBase, field.Value, field.IsExpandedField),
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.format), func(t *testing.T) {
			var buf strings.Builder
			fmt.Fprintf(&buf, tc.format, field)
			if str := buf.String(); str != tc.expected {
				t.Fatalf("expected:\n%q,\ngot:\n%q", tc.expected, str)
			}
		})
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
