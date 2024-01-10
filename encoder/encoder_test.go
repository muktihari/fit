// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bytes"
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
				endianness:               0,
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
				endianness:               1,
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
	fnWriteOK    = fnWriter(func(p []byte) (n int, err error) { return len(p), nil })
	fnWriteErr   = fnWriter(func(p []byte) (n int, err error) { return 0, io.EOF })
	fnWriteAtOK  = fnWriterAt(func(p []byte, offset int64) (n int, err error) { return len(p), nil })
	fnWriteAtErr = fnWriterAt(func(p []byte, offset int64) (n int, err error) { return 0, io.EOF })
	fnSeekOK     = fnSeeker(func(offset int64, whence int) (n int64, err error) { return 0, nil })
	fnSeekErr    = fnSeeker(func(offset int64, whence int) (n int64, err error) { return 0, io.EOF })
)

func TestEncode(t *testing.T) {
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

	// Test same logic for EncodeWithContext
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			_ = enc.EncodeWithContext(context.Background(), &proto.Fit{})
		})
	}
}

type encodeWithDirectUpdateTestCase struct {
	name string
	fit  *proto.Fit
	w    io.Writer
}

func makeEncodeWithDirectUpdateStrategyTableTest() []encodeWithDirectUpdateTestCase {
	return []encodeWithDirectUpdateTestCase{
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
}

func TestEncodeWithDirectUpdateStrategy(t *testing.T) {
	tt := makeEncodeWithDirectUpdateStrategyTableTest()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			_ = enc.Encode(tc.fit)
		})
	}

	// Test same logic for EncodeWithContext
	tt2 := makeEncodeWithDirectUpdateStrategyTableTest()

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			_ = enc.EncodeWithContext(context.Background(), tc.fit)
		})
	}
}

type encodeWithEarlyCheckStrategyTestCase struct {
	name string
	fit  *proto.Fit
	w    io.Writer
	err  error
}

func makeEncodeWithEarlyCheckStrategy() []encodeWithEarlyCheckStrategyTestCase {
	return []encodeWithEarlyCheckStrategyTestCase{
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
				fnInstances := []io.Writer{fnWriteOK, fnWriteOK, fnWriteOK, fnWriteErr}
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
}

func TestEncodeWithEarlyCheckStrategy(t *testing.T) {
	tt := makeEncodeWithEarlyCheckStrategy()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w, WithMessageValidator(fnValidateOK))
			err := enc.Encode(tc.fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
		})
	}

	// Test same logic for EncodeWithContext
	tt2 := makeEncodeWithEarlyCheckStrategy()

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w, WithMessageValidator(fnValidateOK))
			err := enc.EncodeWithContext(context.Background(), tc.fit)
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
		name       string
		mesg       proto.Message
		opts       []Option
		w          io.Writer
		endianness byte
		err        error
	}{
		{
			name: "encode message with default header option happy flow",
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w: fnWriteOK,
		},
		{
			name: "encode message with big-endian",
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w:          fnWriteOK,
			opts:       []Option{WithBigEndian()},
			endianness: bigEndian,
		},
		{
			name: "encode message with header normal multiple local message type happy flow",
			opts: []Option{
				WithNormalHeader(2),
			},
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w: fnWriteOK,
		},
		{
			name: "encode message with compressed timestamp header happy flow",
			opts: []Option{
				WithCompressedTimestampHeader(),
			},
			mesg: factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
				fieldnum.FileIdType: typedef.FileActivity,
			}),
			w: fnWriteOK,
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

			if tc.mesg.Architecture != tc.endianness {
				t.Fatalf("expected endianness: %d, got: %d", tc.endianness, tc.mesg.Architecture)
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

		enc := New(nil, WithNormalHeader(2))
		for i, mesg := range mesgs {
			w := new(bytes.Buffer)
			err := enc.encodeMessage(w, &mesg)
			if err != nil {
				t.Fatal(err)
			}

			mesgDefHeader := w.Bytes()
			expectedHeader := (mesgDefHeader[0] &^ proto.LocalMesgNumMask) | byte(i)
			if mesgDefHeader[0] != expectedHeader {
				t.Fatalf("[%d] expected 0b%08b, got: 0b%08b", i, expectedHeader, mesgDefHeader[0])
			}
		}

		// add 4th mesg, header should be 0, reset.
		mesg := factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(now),
		})
		w := new(bytes.Buffer)
		if err := enc.encodeMessage(w, &mesg); err != nil {
			t.Fatal(err)
		}
		mesgDefHeader := w.Bytes()
		expectedHeader := byte(0)
		if mesgDefHeader[0] != expectedHeader {
			t.Fatalf("expected 0b%08b, got: 0b%08b", expectedHeader, mesgDefHeader[0])
		}
	})
}

type encodeMessagesTestCase struct {
	name          string
	mesgValidator MessageValidator
	mesgs         []proto.Message
	err           error
}

func makeEncodeMessagesTableTest() []encodeMessagesTestCase {
	return []encodeMessagesTestCase{
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
}

func TestEncodeMessages(t *testing.T) {
	tt := makeEncodeMessagesTableTest()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			err := enc.encodeMessages(io.Discard, tc.mesgs)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
		})
	}

	// Test same logic for encodeMessagesWithContext
	tt2 := makeEncodeMessagesTableTest()

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			err := enc.encodeMessagesWithContext(context.Background(), io.Discard, tc.mesgs)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
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

func TestEncodeMessagesWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mesgs := []proto.Message{
		factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivity)),
		),
	}
	enc := New(nil)
	err := enc.encodeMessagesWithContext(ctx, nil, mesgs)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected: %v, got: %v", context.Canceled, err)
	}
}

func TestStreamEncoder(t *testing.T) {
	tt := []struct {
		name string
		w    io.Writer
		err  error
	}{
		{
			name: "writer is io.WriterAt",
			w:    mockWriterAt{},
		},
		{
			name: "writer is io.WriteSeeker",
			w:    mockWriteSeeker{},
		},
		{
			name: "writer is pure io.Writer",
			w:    fnWriteOK,
			err:  ErrWriterAtOrWriteSeekerIsExpected,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := New(tc.w).StreamEncoder()
			if !errors.Is(err, tc.err) {
				t.Errorf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}

		})
	}
}

func createFitForBenchmark(recodSize int) *proto.Fit {
	now := time.Now()
	fit := new(proto.Fit)
	fit.Messages = make([]proto.Message, 0, recodSize)
	fit.Messages = append(fit.Messages,
		factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
			fieldnum.FileIdType:         typedef.FileActivity,
			fieldnum.FileIdManufacturer: typedef.ManufacturerBryton,
			fieldnum.FileIdProductName:  "1901",
			fieldnum.FileIdNumber:       uint16(0),
			fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
			fieldnum.FileIdSerialNumber: uint32(5122),
		}),
		factory.CreateMesg(mesgnum.Sport).WithFieldValues(map[byte]any{
			fieldnum.SportSport:    typedef.SportCycling,
			fieldnum.SportSubSport: typedef.SubSportRoad,
		}),
		factory.CreateMesg(mesgnum.Activity).WithFieldValues(map[byte]any{
			fieldnum.ActivityTimestamp:      datetime.ToUint32(now),
			fieldnum.ActivityType:           typedef.ActivityTypeCycling,
			fieldnum.ActivityTotalTimerTime: uint32(30877.0 * 1000),
			fieldnum.ActivityNumSessions:    uint16(1),
			fieldnum.ActivityEvent:          typedef.EventActivity,
		}),
		factory.CreateMesg(mesgnum.Session).WithFieldValues(map[byte]any{
			fieldnum.SessionTimestamp:        datetime.ToUint32(now),
			fieldnum.SessionStartTime:        datetime.ToUint32(now),
			fieldnum.SessionTotalElapsedTime: uint32(30877.0 * 1000),
			fieldnum.SessionTotalDistance:    uint32(32172.05 * 100),
			fieldnum.SessionSport:            typedef.SportCycling,
			fieldnum.SessionSubSport:         typedef.SubSportRoad,
			fieldnum.SessionTotalMovingTime:  uint32(22079.0 * 1000),
			fieldnum.SessionTotalCalories:    uint16(12824),
			fieldnum.SessionAvgSpeed:         uint16(5.98 * 1000),
			fieldnum.SessionMaxSpeed:         uint16(13.05 * 1000),
			fieldnum.SessionMaxAltitude:      uint16((504.0 + 500) * 5),
			fieldnum.SessionTotalAscent:      uint16(909),
			fieldnum.SessionTotalDescent:     uint16(901),
			fieldnum.SessionSwcLat:           int32(0),
			fieldnum.SessionSwcLong:          int32(0),
			fieldnum.SessionNecLat:           int32(0),
			fieldnum.SessionNecLong:          int32(0),
		}),
	)

	for i := 0; i < recodSize-len(fit.Messages); i++ {
		now = now.Add(time.Second) // only time is moving forward
		if i%100 == 0 {            // add event every 100 message
			fit.Messages = append(fit.Messages, factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStop)),
			))
			now = now.Add(10 * time.Second) // gap
			fit.Messages = append(fit.Messages, factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventActivity)),
				factory.CreateField(mesgnum.Event, fieldnum.EventEventType).WithValue(uint8(typedef.EventTypeStart)),
			))
			now = now.Add(time.Second) // gap
		}

		record := factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp:    datetime.ToUint32(now),
			fieldnum.RecordPositionLat:  int32(-90481372),
			fieldnum.RecordPositionLong: int32(1323227263),
			fieldnum.RecordSpeed:        uint16(8.33 * 1000),
			fieldnum.RecordDistance:     uint32(405.81 * 100),
			fieldnum.RecordHeartRate:    uint8(110),
			fieldnum.RecordCadence:      uint8(85),
			fieldnum.RecordAltitude:     uint16((166.0 + 500.0) * 5.0),
			fieldnum.RecordTemperature:  int8(32),
		})

		if i%200 == 0 { // assume every 200 record hr sensor is not sending any data
			record.RemoveFieldByNum(fieldnum.RecordHeartRate)
		}

		fit.Messages = append(fit.Messages, record)
	}

	return fit
}

func BenchmarkEncode(b *testing.B) {
	b.StopTimer()
	fit := createFitForBenchmark(100_000)
	b.StartTimer()

	b.Run("normal header zero", func(b *testing.B) {
		b.StopTimer()
		enc := New(io.Discard)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.reset()
		}
	})
	b.Run("normal header 15", func(b *testing.B) {
		b.StopTimer()
		enc := New(io.Discard, WithNormalHeader(15))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.reset()
		}
	})
	b.Run("compressed timestamp header", func(b *testing.B) {
		b.StopTimer()
		enc := New(io.Discard, WithCompressedTimestampHeader())
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.reset()
		}
	})
}

var discard = discardAt{}

type discardAt struct{}

var _ io.Writer = discardAt{}
var _ io.WriterAt = discardAt{}

func (discardAt) Write(p []byte) (int, error) {
	return len(p), nil
}

func (discardAt) WriteAt(p []byte, off int64) (n int, err error) {
	return len(p), nil
}

func BenchmarkEncodeWriterAt(b *testing.B) {
	b.StopTimer()
	fit := createFitForBenchmark(100_000)
	b.StartTimer()

	b.Run("normal header zero", func(b *testing.B) {
		b.StopTimer()
		enc := New(discard)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.reset()
		}
	})
	b.Run("normal header 15", func(b *testing.B) {
		b.StopTimer()
		enc := New(discard, WithNormalHeader(15))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.reset()
		}
	})
	b.Run("compressed timestamp header", func(b *testing.B) {
		b.StopTimer()
		enc := New(discard, WithCompressedTimestampHeader())
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(fit)
			enc.reset()
		}
	})
}
