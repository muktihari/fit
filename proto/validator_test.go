// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"errors"
	"testing"

	"github.com/muktihari/fit/profile/basetype"
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
			validator := proto.NewValidator(tc.protocolVersion)
			err := validator.ValidateMessageDefinition(&tc.mesgDef)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, git: %v", tc.err, err)
			}
		})
	}
}
