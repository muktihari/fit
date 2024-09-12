// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
)

func TestNewMessageDefinition(t *testing.T) {
	tt := []struct {
		name    string
		mesg    *Message
		mesgDef *MessageDefinition
		err     error
	}{
		{name: "nil mesg", err: errNilMesg},
		{
			name: "field value exceed max 255",
			mesg: &Message{Num: mesgnum.FileId, Fields: []Field{
				{
					FieldBase: &FieldBase{Num: fieldnum.FileIdProductName, BaseType: basetype.String},
					Value:     String(string(make([]byte, 256))),
				},
			}},
			err: errValueSizeExceed255,
		},
		{
			name: "developerField value exceed max 255",
			mesg: &Message{Num: mesgnum.FileId, Fields: []Field{},
				DeveloperFields: []DeveloperField{
					{Value: String(string(make([]byte, 256)))},
				},
			},
			err: errValueSizeExceed255,
		},
		{
			name: "fields only with non-array values",
			mesg: &Message{Num: mesgnum.FileId, Fields: []Field{
				{FieldBase: &FieldBase{Num: fieldnum.FileIdType, BaseType: basetype.Enum}, Value: Uint8(typedef.FileActivity.Byte())},
			}},
			mesgDef: &MessageDefinition{
				Header:  MesgDefinitionMask,
				MesgNum: mesgnum.FileId,
				FieldDefinitions: []FieldDefinition{
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
			mesg: &Message{Num: mesgnum.FileId, Fields: []Field{
				{FieldBase: &FieldBase{Num: fieldnum.FileIdProductName, BaseType: basetype.String}, Value: String("FIT SDK Go")},
			}},
			mesgDef: &MessageDefinition{
				Header:  MesgDefinitionMask,
				MesgNum: mesgnum.FileId,
				FieldDefinitions: []FieldDefinition{
					{
						Num:      fieldnum.FileIdProductName,
						Size:     1 * 11, // len("FIT SDK Go") == 10 + '0x00'
						BaseType: basetype.String,
					},
				},
			},
		},
		{
			name: "fields only with array of byte",
			mesg: &Message{Num: mesgnum.UserProfile, Fields: []Field{
				{FieldBase: &FieldBase{Num: fieldnum.UserProfileGlobalId, BaseType: basetype.Byte}, Value: SliceUint8([]byte{2, 9})},
			}},
			mesgDef: &MessageDefinition{
				Header:  MesgDefinitionMask,
				MesgNum: mesgnum.UserProfile,
				FieldDefinitions: []FieldDefinition{
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
			mesg: &Message{Num: mesgnum.UserProfile,
				Fields: []Field{
					{FieldBase: &FieldBase{Num: fieldnum.UserProfileGlobalId, BaseType: basetype.Byte}, Value: SliceUint8([]byte{2, 9})},
				},
				DeveloperFields: []DeveloperField{
					{Num: 0, DeveloperDataIndex: 0, Value: Uint8(1)},
				}},
			mesgDef: &MessageDefinition{
				Header:  MesgDefinitionMask | DevDataMask,
				MesgNum: mesgnum.UserProfile,
				FieldDefinitions: []FieldDefinition{
					{
						Num:      fieldnum.UserProfileGlobalId,
						Size:     2,
						BaseType: basetype.Byte,
					},
				},
				DeveloperFieldDefinitions: []DeveloperFieldDefinition{
					{
						Num: 0, Size: 1, DeveloperDataIndex: 0,
					},
				},
			},
		},
		{
			name: "developer fields with string value \"FIT SDK Go\", size should be 11",
			mesg: &Message{Num: mesgnum.UserProfile,
				Fields: []Field{
					{FieldBase: &FieldBase{Num: fieldnum.UserProfileGlobalId, BaseType: basetype.Byte}, Value: SliceUint8([]byte{2, 9})},
				},
				DeveloperFields: []DeveloperField{
					{
						Num: 0, DeveloperDataIndex: 0, Value: String("FIT SDK Go"),
					},
				}},
			mesgDef: &MessageDefinition{
				Header:  MesgDefinitionMask | DevDataMask,
				MesgNum: mesgnum.UserProfile,
				FieldDefinitions: []FieldDefinition{
					{
						Num:      fieldnum.UserProfileGlobalId,
						Size:     2,
						BaseType: basetype.Byte,
					},
				},
				DeveloperFieldDefinitions: []DeveloperFieldDefinition{
					{
						Num: 0, Size: 11, DeveloperDataIndex: 0,
					},
				},
			},
		},
		{
			name: "developer fields with value []uint16{1,2,3}, size should be 3*2 = 6",
			mesg: &Message{Num: mesgnum.UserProfile,
				Fields: []Field{
					{FieldBase: &FieldBase{Num: fieldnum.UserProfileGlobalId, BaseType: basetype.Byte}, Value: SliceUint8([]byte{2, 9})},
				},
				DeveloperFields: []DeveloperField{
					{Num: 0, DeveloperDataIndex: 0, Value: SliceUint16([]uint16{1, 2, 3})},
				}},
			mesgDef: &MessageDefinition{
				Header:  MesgDefinitionMask | DevDataMask,
				MesgNum: mesgnum.UserProfile,
				FieldDefinitions: []FieldDefinition{
					{
						Num:      fieldnum.UserProfileGlobalId,
						Size:     2,
						BaseType: basetype.Byte,
					},
				},
				DeveloperFieldDefinitions: []DeveloperFieldDefinition{
					{
						Num: 0, Size: 6, DeveloperDataIndex: 0,
					},
				},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			mesgDef, err := NewMessageDefinition(tc.mesg)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(mesgDef, tc.mesgDef); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestIsValueEqualTo(t *testing.T) {
	tt := []struct {
		field Field
		value int64
		eq    bool
	}{
		{
			field: Field{Value: Uint8(89)},
			value: 89,
			eq:    true,
		},
		{
			field: Field{Value: String("fit")},
			value: 89,
			eq:    false,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v, %t", tc.value, tc.eq), func(t *testing.T) {
			if eq := tc.field.isValueEqualTo(tc.value); eq != tc.eq {
				t.Fatalf("expected: %t, got: %t", tc.eq, eq)
			}
		})
	}
}

func TestConvertToInt64(t *testing.T) {
	tt := []struct {
		value    Value
		expected int64
		ok       bool
	}{
		{value: Value{}, expected: 0, ok: false},
		{value: Int8(10), expected: 10, ok: true},
		{value: Uint8(10), expected: 10, ok: true},
		{value: Int16(10), expected: 10, ok: true},
		{value: Uint16(10), expected: 10, ok: true},
		{value: Int32(10), expected: 10, ok: true},
		{value: Uint32(10), expected: 10, ok: true},
		{value: Int64(10), expected: 10, ok: true},
		{value: Uint64(10), expected: 10, ok: true},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %v", i, tc.value.Any()), func(t *testing.T) {
			val, ok := convertToInt64(tc.value)
			if ok != tc.ok {
				t.Fatalf("expected: %v, got: %v", tc.ok, ok)
			}
			if val != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, val)
			}
		})
	}
}
