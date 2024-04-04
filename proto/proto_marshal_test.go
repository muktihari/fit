// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"errors"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestHeaderMarshaler(t *testing.T) {
	tt := []struct {
		name       string
		fileHeader *proto.FileHeader
		b          []byte
		err        error
	}{
		{
			name: "correct header",
			fileHeader: &proto.FileHeader{
				Size:            14,
				ProtocolVersion: 32,
				ProfileVersion:  2132,
				DataSize:        642262,
				DataType:        ".FIT",
				CRC:             12856,
			},
			b: []byte{
				14,
				32,
				84, 8,
				214, 204, 9, 0,
				46, 70, 73, 84,
				56, 50,
			},
		},
		{
			name: "correct header size 12",
			fileHeader: &proto.FileHeader{
				Size:            12,
				ProtocolVersion: 32,
				ProfileVersion:  2132,
				DataSize:        642262,
				DataType:        ".FIT",
			},
			b: []byte{
				12,
				32,
				84, 8,
				214, 204, 9, 0,
				46, 70, 73, 84,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.fileHeader.MarshalBinary()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if diff := cmp.Diff(b, tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageDefinitionMarshaler(t *testing.T) {
	tt := []struct {
		name    string
		mesgdef *proto.MessageDefinition
		b       []byte
	}{
		{
			name: "mesg def fields only",
			mesgdef: &proto.MessageDefinition{
				Header:       64,
				Reserved:     1,
				Architecture: 0,
				MesgNum:      typedef.MesgNumFileId,
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Enum},
					{Num: 1, Size: 2, BaseType: basetype.Uint16},
					{Num: 2, Size: 2, BaseType: basetype.Uint16},
					{Num: 3, Size: 4, BaseType: basetype.Uint32z},
					{Num: 8, Size: 13, BaseType: basetype.String},
					{Num: 5, Size: 2, BaseType: basetype.Uint16},
				},
				DeveloperFieldDefinitions: nil},
			b: []byte{
				64,   // Header
				1,    // Reserved
				0,    // Architecture
				0, 0, // MesgNum
				6, // len(FieldDefinitions)
				0, 1, 0,
				1, 2, 132,
				2, 2, 132,
				3, 4, 140,
				8, 13, 7,
				5, 2, 132,
			},
		},
		{
			name: "mesg def fields and developer fields",
			mesgdef: &proto.MessageDefinition{
				Header:       64 | 32,
				Architecture: 0,
				MesgNum:      typedef.MesgNumFileId,
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Enum},
					{Num: 1, Size: 2, BaseType: basetype.Uint16},
					{Num: 2, Size: 2, BaseType: basetype.Uint16},
					{Num: 3, Size: 4, BaseType: basetype.Uint32z},
					{Num: 8, Size: 13, BaseType: basetype.String},
					{Num: 5, Size: 2, BaseType: basetype.Uint16},
				},
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{Num: 0, Size: 1, DeveloperDataIndex: 0},
				}},
			b: []byte{
				64 | 32, // Header
				0,       // Reserved
				0,       // Architecture
				0, 0,    // MesgNum
				6, // len(FieldDefinitions)
				0, 1, 0,
				1, 2, 132,
				2, 2, 132,
				3, 4, 140,
				8, 13, 7,
				5, 2, 132,
				1, // len(DeveloperFieldDefinitions)
				0, 1, 0,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, _ := tc.mesgdef.MarshalBinary()
			if diff := cmp.Diff(b, tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageMarshaler(t *testing.T) {
	tt := []struct {
		name string
		mesg *proto.Message
		b    []byte
		err  error
	}{
		{
			name: "file_id mesg",
			mesg: &proto.Message{
				Header: 0,
				Num:    typedef.MesgNumFileId,
				Fields: []proto.Field{
					{Value: proto.Uint8(4)},
					{Value: proto.Uint16(292)},
					{Value: proto.Uint16(100)},
					{Value: proto.Uint32(120188)},
					{Value: proto.String("XOSS iOS APP")},
					{Value: proto.Uint16(1873)},
				},
				DeveloperFields: nil,
			},
			b: []byte{
				0, // Header
				4, // Field[0] ...
				36, 1,
				100, 0,
				124, 213, 1, 0,
				88, 79, 83, 83, 32, 105, 79, 83, 32, 65, 80, 80, 00,
				81, 7,
			},
		},
		{
			name: "record mesg with developer fields",
			mesg: &proto.Message{
				Header: 0,
				Num:    typedef.MesgNumRecord,
				Fields: []proto.Field{
					{Value: proto.Uint8(4)},
					{Value: proto.Uint16(292)},
					{Value: proto.Uint16(100)},
					{Value: proto.Uint32(120188)},
					{Value: proto.String("XOSS iOS APP")},
					{Value: proto.Uint16(1873)},
				},
				DeveloperFields: []proto.DeveloperField{
					{Value: proto.Uint8(10)},
				},
			},
			b: []byte{
				0, // Header
				4, // Field[0] ...
				36, 1,
				100, 0,
				124, 213, 1, 0,
				88, 79, 83, 83, 32, 105, 79, 83, 32, 65, 80, 80, 00,
				81, 7,
				10, // DeveloperField[0]
			},
		},
		{
			name: "marshal fields return error",
			mesg: &proto.Message{
				Header: 0,
				Num:    typedef.MesgNumFileId,
				Fields: []proto.Field{
					{FieldBase: &proto.FieldBase{Num: 99, Name: "error"}, Value: proto.Value{}}, // TODO: check later
				},
				DeveloperFields: nil,
			},
			err: proto.ErrTypeNotSupported,
		},
		{
			name: "marshal fields return error",
			mesg: &proto.Message{
				Header: 0,
				Num:    typedef.MesgNumFileId,
				DeveloperFields: []proto.DeveloperField{
					{Num: 0, DeveloperDataIndex: 0, Value: proto.Value{}},
				},
			},
			err: proto.ErrTypeNotSupported,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.mesg.MarshalBinary()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if diff := cmp.Diff(b, tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func BenchmarkHeaderMarshalBinary(b *testing.B) {
	b.StopTimer()
	header := proto.FileHeader{
		Size:            14,
		ProtocolVersion: 32,
		ProfileVersion:  2132,
		DataSize:        642262,
		DataType:        ".FIT",
		CRC:             12856,
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = header.MarshalBinary()
	}
}

func BenchmarkMessageDefinitionMarshalBinary(b *testing.B) {
	b.StopTimer()
	mesg := factory.CreateMesg(mesgnum.Record)
	mesgDef := proto.CreateMessageDefinition(&mesg)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = mesgDef.MarshalBinary()
	}
}

func BenchmarkMessageDefinitionWriteTo(b *testing.B) {
	b.StopTimer()
	mesg := factory.CreateMesg(mesgnum.Record)
	mesgDef := proto.CreateMessageDefinition(&mesg)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = mesgDef.WriteTo(io.Discard)
	}
}

func BenchmarkMessageMarshalBinary(b *testing.B) {
	b.StopTimer()
	mesg := factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
		fieldnum.RecordPositionLat:  proto.Int32(1000),
		fieldnum.RecordPositionLong: proto.Int32(1000),
		fieldnum.RecordSpeed:        proto.Uint16(1000),
	})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = mesg.MarshalBinary()
	}
}

func BenchmarkMessageWriterTo(b *testing.B) {
	b.StopTimer()
	mesg := factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
		fieldnum.RecordPositionLat:  proto.Int32(1000),
		fieldnum.RecordPositionLong: proto.Int32(1000),
		fieldnum.RecordSpeed:        proto.Uint16(1000),
	})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = mesg.WriteTo(io.Discard)
	}
}
