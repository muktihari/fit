// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bytes"
	"container/list"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	cd                = filepath.Dir(filename)
	testdata          = filepath.Join(cd, "..", "testdata")
)

func TestEncodeRealFiles(t *testing.T) {
	now := time.Date(2023, 9, 15, 6, 0, 0, 0, time.UTC)
	fit := &proto.Fit{
		Messages: []proto.Message{
			factory.CreateMesgOnly(mesgnum.FileId).WithFields(
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerBryton),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("Bryton Active App"),
			),
			factory.CreateMesgOnly(mesgnum.Activity).WithFields(
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityTypeCycling),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
			),
			factory.CreateMesgOnly(mesgnum.Session).WithFields(
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgSpeed).WithValue(uint16(1000)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgCadence).WithValue(uint8(78)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgHeartRate).WithValue(uint8(100)),
			),
			factory.CreateMesgOnly(mesgnum.Record).WithFields(
				factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(78)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(100)),
			),
		},
	}

	f := new(bytes.Buffer)
	enc := New(f)
	if err := enc.EncodeWithContext(context.Background(), fit); err != nil {
		panic(err)
	}

	testEncodeFit, err := os.Open(filepath.Join(testdata, "TestEncode.fit"))
	if err != nil {
		panic(err)
	}
	defer testEncodeFit.Close()

	b, err := io.ReadAll(testEncodeFit)
	if err != nil {
		t.Fatal(err)
	}

	resultb := f.Bytes()

	// Ignore profile version and crc checksum since it will change when we update the Profile.xlsx
	b[2], b[3] = resultb[2], resultb[3]
	b[12], b[13] = resultb[12], resultb[13]

	// Compare with the actual file
	if diff := cmp.Diff(resultb, b); diff != "" {
		t.Fatal(diff)
	}
}

type fnValidate func(mesg *proto.Message) error

func (f fnValidate) Validate(mesg *proto.Message) error { return f(mesg) }

var (
	fnValidateOK  = fnValidate(func(mesg *proto.Message) error { return nil })
	fnValidateErr = fnValidate(func(mesg *proto.Message) error { return fmt.Errorf("validate error") })
)

func TestOptions(t *testing.T) {
	tt := []struct {
		name     string
		opts     []Option
		expected *options
	}{
		{
			name: "defaultOptions",
			opts: nil,
			expected: &options{
				multipleLocalMessageType: 0,
				endianess:                0,
				protocolVersion:          proto.V1,
				messageValidator:         NewMessageValidator(),
			},
		},
		{
			name: "with options",
			opts: []Option{
				WithBigEndian(),
				WithNormalHeader(20),
				WithProtocolVersion(proto.V2),
				WithCompressedTimestampHeader(),
				WithMessageValidator(fnValidateOK),
			},
			expected: &options{
				multipleLocalMessageType: 15,
				endianess:                1,
				protocolVersion:          proto.V2,
				messageValidator:         fnValidateOK,
				headerOption:             headerOptionCompressedTimestamp,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil, tc.opts...)
			if diff := cmp.Diff(enc.options, tc.expected,
				cmp.AllowUnexported(options{}),
				cmp.FilterValues(func(x, y MessageValidator) bool {
					return true
				}, cmp.Ignore()),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestEncodeExported(t *testing.T) {
	tt := []struct {
		name string
		ctx  context.Context
		fit  *proto.Fit
		err  error
	}{
		{
			name: "context nil", ctx: nil, err: ErrNilWriter,
		},
		{
			name: "context canceled",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			err: context.Canceled,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			err := enc.EncodeWithContext(tc.ctx, tc.fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

type fnWriter func(b []byte) (n int, err error)

func (f fnWriter) Write(b []byte) (n int, err error) { return f(b) }

type fnWriterAt func(p []byte, offset int64) (n int, err error)

func (f fnWriterAt) WriteAt(p []byte, offset int64) (n int, err error) { return f(p, offset) }

type mockWriterAt struct {
	io.Writer
	io.WriterAt
}

type fnSeeker func(offset int64, whence int) (n int64, err error)

func (f fnSeeker) Seek(offset int64, whence int) (n int64, err error) { return f(offset, whence) }

type mockWriteSeeker struct {
	io.Writer
	io.Seeker
}

var (
	fnWriteOK    = fnWriter(func(b []byte) (n int, err error) { return 0, nil })
	fnWriteErr   = fnWriter(func(b []byte) (n int, err error) { return 0, io.EOF })
	fnWriteAtOK  = fnWriterAt(func(p []byte, offset int64) (n int, err error) { return 0, nil })
	fnWriteAtErr = fnWriterAt(func(p []byte, offset int64) (n int, err error) { return 0, io.EOF })
	fnSeekOK     = fnSeeker(func(offset int64, whence int) (n int64, err error) { return 0, nil })
	fnSeekErr    = fnSeeker(func(offset int64, whence int) (n int64, err error) { return 0, io.EOF })
)

func TestEncodeUnexported(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
	}{
		{name: "encode with nil", w: nil},
		{name: "encode with writer", w: fnWriteOK},
		{name: "encode with writerAt", w: mockWriterAt{fnWriteOK, fnWriteAtOK}},
		{name: "encode with writeSeeker", w: mockWriteSeeker{fnWriteOK, fnSeekOK}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			_ = enc.Encode(&proto.Fit{})
		})
	}
}

func TestEncodeWithDirectUpdateStrategy(t *testing.T) {
	tt := []struct {
		name string
		opts []Option
		fit  *proto.Fit
		w    io.Writer
	}{
		{
			name: "happy flow coverage",
			fit: &proto.Fit{Messages: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				),
			}},
			w: mockWriterAt{fnWriteOK, fnWriteAtOK},
		},
		{
			name: "encode header error",
			fit:  &proto.Fit{},
			w:    mockWriterAt{fnWriteErr, fnWriteAtErr},
		},
		{
			name: "encode messages error",
			fit:  &proto.Fit{Messages: []proto.Message{{}}},
			w:    mockWriterAt{fnWriteErr, fnWriteAtErr},
		},
		{
			name: "encode crc error",
			fit: &proto.Fit{Messages: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				),
			}},
			w: func() io.Writer {
				fnWrites := []io.Writer{fnWriteOK, fnWriteOK, fnWriteOK, fnWriteErr}
				index := 0

				return mockWriterAt{
					fnWriter(func(b []byte) (n int, err error) {
						fn := fnWrites[index]
						index++
						return fn.Write(b)
					}),
					fnWriteAtOK,
				}
			}(),
		},
		{
			name: "update error",
			fit: &proto.Fit{FileHeader: proto.FileHeader{Size: 14, DataSize: 100, DataType: proto.DataTypeFIT}, Messages: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				),
			}},
			w: mockWriterAt{fnWriteOK, fnWriteAtErr},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			_ = enc.Encode(tc.fit)
		})
	}
}

func TestEncodeWithEarlyCheckStrategy(t *testing.T) {
	tt := []struct {
		name string
		fit  *proto.Fit
		w    io.Writer
		err  error
	}{
		{
			name: "happy flow coverage",
			fit:  &proto.Fit{Messages: []proto.Message{{}}},
			w:    fnWriteOK,
		},
		{
			name: "calculate data size error",
			fit:  &proto.Fit{Messages: []proto.Message{}},
			w: func() io.Writer {
				fnInstances := []io.Writer{fnWriteErr}
				index := 0

				return fnWriter(func(b []byte) (n int, err error) {
					f := fnInstances[index]
					index++
					return f.Write(b)
				})
			}(),
			err: ErrEmptyMessages,
		},
		{
			name: "encode header error",
			fit: &proto.Fit{Messages: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
				),
			}},
			w:   fnWriteErr,
			err: io.EOF,
		},
		{
			name: "encode messages error",
			fit: &proto.Fit{Messages: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
				),
			}},
			w: func() io.Writer {
				fnInstances := []io.Writer{fnWriteOK, fnWriteErr}
				index := 0

				return fnWriter(func(b []byte) (n int, err error) {
					f := fnInstances[index]
					index++
					return f.Write(b)
				})
			}(),
			err: io.EOF, // since fnWriteErr produce io.EOF
		},
		{
			name: "encode crc error",
			fit:  &proto.Fit{Messages: []proto.Message{{}}},
			w: func() io.Writer {
				fnInstances := []io.Writer{fnWriteOK, fnWriteOK, fnWriteErr}
				index := 0

				return fnWriter(func(b []byte) (n int, err error) {
					f := fnInstances[index%len(fnInstances)]
					index++
					return f.Write(b)
				})
			}(),
			err: io.EOF, // since fnWriteErr produce io.EOF
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w, WithMessageValidator(fnValidateOK))
			err := enc.Encode(tc.fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestUpdateHeader(t *testing.T) {
	tt := []struct {
		name   string
		header proto.FileHeader
		w      io.Writer
	}{
		{
			name: "data size not changed",
			w:    nil,
		},
		{
			name:   "writerAt flow",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w:      mockWriterAt{fnWriteOK, fnWriteAtOK},
		},
		{
			name:   "writeSeeker happy flow",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w:      mockWriteSeeker{fnWriteOK, fnSeekOK},
		},
		{
			name:   "writeSeeker error on seek",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w:      mockWriteSeeker{fnWriteOK, fnSeekErr},
		},
		{
			name:   "writeSeeker error on write",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w:      mockWriteSeeker{fnWriteErr, fnSeekOK},
		},
		{
			name:   "writeSeeker error on seek for resetting offset",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w: func() io.Writer {
				fnSeeks := []io.Seeker{fnSeekOK, fnSeekErr}
				index := 0

				return mockWriteSeeker{
					fnWriteOK,
					fnSeeker(func(offset int64, whence int) (n int64, err error) {
						fn := fnSeeks[index]
						index++
						return fn.Seek(offset, whence)
					}),
				}
			}(),
		},
		{
			name:   "encoder internal error caused by bad implementation",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w:      nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			_ = enc.updateHeader(&tc.header)
		})
	}
}

func TestEncodeHeader(t *testing.T) {
	tt := []struct {
		name            string
		protocolVersion proto.Version
		header          proto.FileHeader
		b               []byte
	}{
		{
			name:   "no header",
			header: proto.FileHeader{},
			b: func() []byte {
				b := []byte{
					14,
					16,
					0, 0, // profile version will be updated
					0, 0, 0, 0, // dataSize zero
					46, 70, 73, 84, // .FIT
					0, 0, // crc checksum will be calculated
				}

				binary.LittleEndian.PutUint16(b[2:4], profile.Version())

				crc := crc16.New(crc16.MakeFitTable())
				crc.Write(b[:12])
				binary.LittleEndian.PutUint16(b[12:14], crc.Sum16())

				return b
			}(),
		},
		{
			name:            "header 12 legacy",
			protocolVersion: proto.V2,
			header: proto.FileHeader{
				Size:            12,
				ProtocolVersion: 32,
				ProfileVersion:  2132,
				DataSize:        642262,
				DataType:        ".FIT",
				CRC:             12856,
			},
			b: []byte{
				12,
				32,
				84, 8,
				214, 204, 9, 0,
				46, 70, 73, 84, // .FIT
			},
		},
		{
			name:            "header 14",
			protocolVersion: proto.V2,
			header: proto.FileHeader{
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
				46, 70, 73, 84, // .FIT
				56, 50,
			},
		},
		{
			name:            "header 14 crc mismatch",
			protocolVersion: proto.Version(31),
			header: proto.FileHeader{
				Size:            14,
				ProtocolVersion: 31, // this is changed, crc should change too
				ProfileVersion:  2132,
				DataSize:        642262,
				DataType:        ".FIT",
				CRC:             12856,
			},
			b: []byte{
				14,
				31, // this is changed, crc should change too
				84, 8,
				214, 204, 9, 0,
				46, 70, 73, 84,
				247, 38,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bytebuf := new(bytes.Buffer)
			enc := New(bytebuf)
			if tc.protocolVersion != 0 {
				enc.options.protocolVersion = tc.protocolVersion
			}
			_ = enc.encodeHeader(&tc.header)

			if diff := cmp.Diff(bytebuf.Bytes(), tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestEncodeMessage(t *testing.T) {
	tt := []struct {
		name string
		mesg proto.Message
		opts []Option
		w    io.Writer
		err  error
	}{
		{
			name: "encode message with default header option happy flow",
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w: fnWriter(func(b []byte) (n int, err error) {
				return 0, nil
			}),
		},
		{
			name: "encode message with header normal multiple local message type happy flow",
			opts: []Option{
				WithNormalHeader(2),
			},
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w: fnWriter(func(b []byte) (n int, err error) {
				return 0, nil
			}),
		},
		{
			name: "encode message with compressed timestamp header happy flow",
			opts: []Option{
				WithCompressedTimestampHeader(),
			},
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w: fnWriter(func(b []byte) (n int, err error) {
				return 0, nil
			}),
		},
		{
			name: "message validator's validate return error",
			mesg: proto.Message{},
			w:    nil,
			err:  ErrNoFields,
		},
		{
			name: "normal header: protocol validator's validate message definition return error",
			opts: []Option{
				WithProtocolVersion(proto.V1),
			},
			mesg: proto.Message{Fields: []proto.Field{
				{
					FieldBase: &proto.FieldBase{
						Name: factory.NameUnknown,
						Type: profile.Sint64, // int64 type is ilegal for protocol v1.0
					},
					Value: int64(1234),
				},
			}},
			w:   nil,
			err: proto.ErrProtocolViolation,
		},
		{
			name: "compressed timestamp header: protocol validator's validate message definition return error",
			opts: []Option{
				WithProtocolVersion(proto.V1),
				WithCompressedTimestampHeader(),
			},
			mesg: proto.Message{Fields: []proto.Field{
				{
					FieldBase: &proto.FieldBase{
						Name: factory.NameUnknown,
						Type: profile.Sint64, // int64 type is ilegal for protocol v1.0
					},
					Value: int64(1234),
				},
			}},
			w:   nil,
			err: proto.ErrProtocolViolation,
		},
		{
			name: "write message definition return error",
			opts: []Option{
				WithMessageValidator(fnValidateOK),
			},
			mesg: proto.Message{},
			w:    fnWriteErr,
			err:  io.EOF,
		},
		{
			name: "write marshal message return error",
			opts: []Option{
				WithMessageValidator(fnValidateOK),
			},
			mesg: proto.Message{Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(nil),
			}},
			w:   fnWriteOK,
			err: typedef.ErrNilValue,
		},
		{
			name: "write message return error",
			opts: []Option{
				WithMessageValidator(fnValidateOK),
			},
			mesg: proto.Message{},
			w: func() io.Writer {
				fnInstances := []io.Writer{fnWriteOK, fnWriteErr}
				index := 0

				return fnWriter(func(b []byte) (n int, err error) {
					f := fnInstances[index]
					index++
					return f.Write(b)
				})
			}(),
			err: io.EOF,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w, tc.opts...)
			err := enc.encodeMessage(tc.w, &tc.mesg)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
			if (tc.mesg.Header & proto.DevDataMask) == proto.DevDataMask {
				t.Fatalf("message header should not contain Developer Data Flag")
			}
		})
	}
}

func TestEncodeMessageWithMultipleLocalMessageType(t *testing.T) {
	now := time.Now()
	mesgs := []proto.Message{
		factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(now),
		}),
		factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(now.Add(time.Second)),
			fieldnum.RecordHeartRate: uint8(70),
		}),
		factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(now.Add(2 * time.Second)),
			fieldnum.RecordSpeed:     uint16(1000),
		}),
	}

	t.Run("multiple local mesg type", func(t *testing.T) {
		// We have 3 messages with differents field definitions,
		// this should produces different localMesgNum in header.

		mesgs := slices.Clone(mesgs)
		for i := range mesgs {
			mesgs[i] = mesgs[i].Clone()
		}

		enc := New(nil, WithNormalHeader(3))
		for i, mesg := range mesgs {
			w := new(bytes.Buffer)
			err := enc.encodeMessage(w, &mesg)
			if err != nil {
				t.Fatal(err)
			}

			mesgDefHeader := w.Bytes()
			expectedHeader := (mesgDefHeader[0] &^ proto.LocalMesgNumMask) | byte(i)
			if mesgDefHeader[0] != expectedHeader {
				t.Fatalf("expected 0b%08b, got: 0b%08b", expectedHeader, mesgDefHeader[0])
			}
		}
	})
}

func TestEncodeMessages(t *testing.T) {
	tt := []struct {
		name          string
		mesgValidator MessageValidator
		mesgs         []proto.Message
		err           error
	}{
		{
			name:          "encode messages happy flow",
			mesgValidator: fnValidateOK,
			mesgs: []proto.Message{
				factory.CreateMesgOnly(mesgnum.FileId).WithFields(
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdProduct).WithValue(uint16(typedef.GarminProductEdge1030)),
				),
			},
		},
		{
			name:  "encode messages return empty messages error",
			mesgs: []proto.Message{},
			err:   ErrEmptyMessages,
		},
		{
			name:          "encode messages return error",
			mesgValidator: fnValidateErr,
			mesgs:         []proto.Message{{}},
			err:           ErrNoFields, // Validator error since the first mesg is invalid.
		},
		{
			name:  "missing file_id mesg",
			mesgs: []proto.Message{factory.CreateMesg(mesgnum.Record)},
			err:   ErrMissingFileId,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			err := enc.encodeMessages(io.Discard, tc.mesgs)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestRedefineLocalMesgNum(t *testing.T) {
	type _struct struct {
		name         string
		b            []byte
		listCapacity byte
		list         *list.List
		lru          *list.List
		num          byte
		writtable    bool
	}

	tt := []_struct{
		{
			name:      "init value",
			b:         []byte{0, 1},
			list:      list.New(),
			lru:       list.New(),
			num:       0,
			writtable: true,
		},
		{
			name: "eq with 1st index",
			b:    []byte{0, 1},
			list: func() *list.List {
				l := list.New()
				l.PushFront([]byte{0, 1})
				return l
			}(),
			lru:       list.New(),
			num:       0,
			writtable: false,
		},
		{
			name: "eq with 2st index",
			b:    []byte{0, 2},
			list: func() *list.List {
				l := list.New()
				l.PushBack([]byte{0, 1})
				l.PushBack([]byte{0, 2})
				return l
			}(),
			lru:       list.New(),
			num:       1,
			writtable: false,
		},
		func() _struct {
			ls := list.New()
			lru := list.New()
			lru.PushBack(ls.PushBack([]byte{0, 1}))
			lru.PushBack(ls.PushBack([]byte{0, 2}))

			return _struct{
				name:         "full, replace LRU item on the first index",
				b:            []byte{0, 3},
				listCapacity: 1,
				list:         ls,
				lru:          lru,
				num:          0,
				writtable:    true,
			}
		}(),
		func() _struct {
			ls, lru := list.New(), list.New()

			lru.PushBack(ls.PushBack([]byte{0, 1}))
			lru.PushBack(ls.PushBack([]byte{0, 2}))
			lru.PushBack(ls.PushBack([]byte{0, 3}))

			lru.MoveToBack(lru.Front())

			return _struct{
				name:         "full, replace LRU item on 2nd index",
				b:            []byte{0, 4},
				listCapacity: 2,
				list:         ls,
				lru:          lru,
				num:          1,
				writtable:    true,
			}
		}(),
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			enc.options.multipleLocalMessageType = tc.listCapacity
			enc.localMesgDefinitions = tc.list
			enc.localMesgDefinitionsLRU = tc.lru

			num, writtable := enc.redefineLocalMesgNum(tc.b)
			if num != tc.num {
				t.Fatalf("expected: %d, got: %d", tc.num, num)
			}
			if writtable != tc.writtable {
				t.Fatalf("expected: %t, got: %t", tc.writtable, writtable)
			}

		})
	}
}

func TestIsMesgDefinitionWriteable(t *testing.T) {
	tt := []struct {
		name      string
		b         []byte
		list      *list.List
		writeable bool
	}{
		{
			name:      "init element value",
			b:         []byte{1, 1},
			list:      list.New(),
			writeable: true,
		},
		{
			name: "eq with 1st element value",
			b:    []byte{1, 1},
			list: func() *list.List {
				ls := list.New()
				ls.PushFront([]byte{1, 1})
				return ls
			}(),
			writeable: false,
		},
		{
			name: "not eq, replace existing",
			b:    []byte{1, 1},
			list: func() *list.List {
				ls := list.New()
				ls.PushFront([]byte{0, 1})
				return ls
			}(),
			writeable: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			enc.localMesgDefinitions = tc.list

			writeable := enc.isMesgDefinitionWriteable(tc.b)
			if writeable != tc.writeable {
				t.Fatalf("expected: %t, got: %t", tc.writeable, writeable)
			}
		})
	}
}

func TestCompressTimestampInHeader(t *testing.T) {
	now := time.Now()
	offset := byte(datetime.ToUint32(now) & proto.CompressedTimeMask)
	tt := []struct {
		name    string
		mesgs   []proto.Message
		headers []byte
	}{
		{
			name: "compress timestamp in header happy flow",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
					fieldnum.FileIdManufacturer: typedef.ManufacturerGarmin,
					fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: datetime.ToUint32(now),
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: datetime.ToUint32(now.Add(time.Second)), // +1s
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: datetime.ToUint32(now.Add(2 * time.Second)), // +2s
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: datetime.ToUint32(now.Add(32 * time.Second)), // +32 rollover
				}),
			},
			headers: []byte{
				proto.MesgNormalHeaderMask, // file_id: has no timestamp
				proto.MesgNormalHeaderMask, // record: the message containing timestamp reference prior to the use of compressed header.
				proto.MesgCompressedHeaderMask | (offset+1)&proto.CompressedTimeMask,
				proto.MesgCompressedHeaderMask | (offset+2)&proto.CompressedTimeMask,
				proto.MesgCompressedHeaderMask | (offset+32)&proto.CompressedTimeMask,
			},
		},
		{
			name: "timestamp less than DateTimeMin",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
					fieldnum.FileIdManufacturer: typedef.ManufacturerGarmin,
					fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: uint32(1234),
				}),
			},
			headers: []byte{
				proto.MesgNormalHeaderMask,
				proto.MesgNormalHeaderMask,
			},
		},
		{
			name: "timestamp wrong type not uint32 or typedef.DateTime",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
					fieldnum.FileIdManufacturer: typedef.ManufacturerGarmin,
					fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: typedef.DateTime(datetime.ToUint32(now)),
				}),
			},
			headers: []byte{
				proto.MesgNormalHeaderMask,
				proto.MesgNormalHeaderMask,
			},
		},
		{
			name: "timestamp wrong type not uint32 or typedef.DateTime",
			mesgs: []proto.Message{
				factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
					fieldnum.FileIdManufacturer: typedef.ManufacturerGarmin,
					fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
				}),
				factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
					fieldnum.RecordTimestamp: now, // time.Time{}
				}),
			},
			headers: []byte{
				proto.MesgNormalHeaderMask,
				proto.MesgNormalHeaderMask,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			for i := range tc.mesgs {
				enc.compressTimestampIntoHeader(&tc.mesgs[i])
			}
			// Now that all message have been processed let's check the header
			for i := range tc.mesgs {
				if diff := cmp.Diff(tc.mesgs[i].Header, tc.headers[i]); diff != "" {
					t.Errorf("index: %d: %s", i, diff)
				}
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	tt := []struct {
		name   string
		prev   []byte
		target []byte
		eq     bool
	}{
		{name: "same as prev", prev: []byte{1, 2}, target: []byte{1, 2}, eq: true},
		{name: "diff len", prev: []byte{1, 2}, target: []byte{1}, eq: false},
		{name: "diff byte", prev: []byte{1, 2}, target: []byte{2, 1}, eq: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			eq := isEqual(tc.prev, tc.target)
			if eq != tc.eq {
				t.Fatalf("expected: %t, got: %t", tc.eq, eq)
			}
		})
	}
}
