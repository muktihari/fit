// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestValidateMessageDefinition(t *testing.T) {
	tt := []struct {
		name            string
		mesgDef         proto.MessageDefinition
		protocolVersion proto.Version
		err             error
	}{
		{
			name:            "validate using v1.0 happy flow",
			protocolVersion: proto.V1,
			mesgDef: proto.MessageDefinition{
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Uint16},
				},
			},
		},
		{
			name:            "validate using v2.0 happy flow",
			protocolVersion: proto.V2,
			mesgDef: proto.MessageDefinition{
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Uint32z},
				},
			},
		},
		{
			name:            "validate using v1.0 contains developer fields",
			protocolVersion: proto.V1,
			mesgDef: proto.MessageDefinition{
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Uint16},
				},
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{Num: 0, DeveloperDataIndex: 0, Size: 1},
				},
			},
			err: proto.ErrProtocolViolation,
		},
		{
			name:            "validate using v1.0 contains unsupported type",
			protocolVersion: proto.V1,
			mesgDef: proto.MessageDefinition{
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Uint64},
				},
			},
			err: proto.ErrProtocolViolation,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := proto.Validator.ValidateMessageDefinition(&tc.mesgDef, tc.protocolVersion)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, git: %v", tc.err, err)
			}
		})
	}
}

func TestValidateMessage(t *testing.T) {
	tt := []struct {
		name            string
		mesg            proto.Message
		protocolVersion proto.Version
		err             error
	}{
		{
			name:            "happy flow v1",
			protocolVersion: proto.V1,
			mesg: proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
				factory.CreateField(mesgnum.Record, fieldnum.RecordDistance),
			}},
		},
		{
			name:            "happy flow v2",
			protocolVersion: proto.V2,
			mesg: proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
				factory.CreateField(mesgnum.Record, fieldnum.RecordDistance),
			}},
		},
		{
			name:            "v1 is selected but mesg has developer data",
			protocolVersion: proto.V1,
			mesg: proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
				factory.CreateField(mesgnum.Record, fieldnum.RecordDistance),
			}, DeveloperFields: []proto.DeveloperField{{}}},
			err: proto.ErrProtocolViolation,
		},
		{
			name:            "v1 is selected but mesg has type Int64",
			protocolVersion: proto.V1,
			mesg: proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
				{FieldBase: &proto.FieldBase{BaseType: basetype.Sint64}},
			}},
			err: proto.ErrProtocolViolation,
		},
		{
			name:            "v2 is selected it's ok that mesg has developer data",
			protocolVersion: proto.V2,
			mesg: proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
				factory.CreateField(mesgnum.Record, fieldnum.RecordDistance),
			}, DeveloperFields: []proto.DeveloperField{{}}},
		},
		{
			name:            "v2 is selected it's ok that mesg has type Int64",
			protocolVersion: proto.V2,
			mesg: proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
				{FieldBase: &proto.FieldBase{BaseType: basetype.Sint64}},
			}},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			err := proto.Validator.ValidateMessage(&tc.mesg, tc.protocolVersion)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}
}
