// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	cd                = filepath.Dir(filename)
	testdata          = filepath.Join(cd, "..", "testdata")
)

func TestEncodeRealFiles(t *testing.T) {
	now := time.Date(2023, 9, 15, 6, 0, 0, 0, time.UTC)
	fit := &proto.FIT{
		Messages: []proto.Message{
			{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerBryton),
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("Bryton Active App"),
			}},
			{Num: mesgnum.Activity, Fields: []proto.Field{
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityTypeCycling),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now)),
				factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
			}},
			{Num: mesgnum.Session, Fields: []proto.Field{
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgSpeed).WithValue(uint16(1000)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgCadence).WithValue(uint8(78)),
				factory.CreateField(mesgnum.Session, fieldnum.SessionAvgHeartRate).WithValue(uint8(100)),
			}},
			{Num: mesgnum.Record, Fields: []proto.Field{
				factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(78)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(100)),
			}},
		},
	}

	f := &writeSeekerStub{}
	enc := New(f, WithWriteBufferSize(0))
	if err := enc.EncodeWithContext(context.Background(), fit); err != nil {
		t.Fatal(err)
	}

	testEncode, err := os.ReadFile(filepath.Join(testdata, "TestEncode.fit"))
	if err != nil {
		t.Fatal(err)
	}

	// Ignore profile version and crc checksum since it will change when we update the Profile.xlsx
	testEncode[2], testEncode[3] = f.buf[2], f.buf[3]
	testEncode[12], testEncode[13] = f.buf[12], f.buf[13]

	// Compare with the actual file
	if diff := cmp.Diff(f.buf, testEncode); diff != "" {
		t.Fatal(diff)
	}

	t.Run("Encode to existing file using WriteSeeker: OK", func(t *testing.T) {
		triathlon, err := os.ReadFile(filepath.Join(testdata, "from_garmin_forums", "triathlon_summary_last.fit"))
		if err != nil {
			t.Fatal(err)
		}
		expected := append(triathlon, testEncode...) // Chained FIT File

		f := &writeSeekerStub{buf: triathlon}
		_, err = f.Seek(0, io.SeekEnd)
		if err != nil {
			t.Fatal(err)
		}

		fit.FileHeader.DataSize = 0 // Must reset
		enc.Reset(f, WithWriteBufferSize(0))
		if err := enc.EncodeWithContext(context.Background(), fit); err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(f.buf, expected); diff != "" {
			t.Fatal(diff)
		}
	})

	t.Run("Encode to existing file using WriterAt: Corrupt", func(t *testing.T) {
		triathlon, err := os.ReadFile(filepath.Join(testdata, "from_garmin_forums", "triathlon_summary_last.fit"))
		if err != nil {
			t.Fatal(err)
		}

		f := &writerAtStub{buf: triathlon}

		fit.FileHeader.DataSize = 0 // Must reset
		enc.Reset(f, WithWriteBufferSize(0))

		// This will overwrite the content of the existing file,
		// as updateHeader will write at offset 0, corrupting the file.
		if err := enc.EncodeWithContext(context.Background(), fit); err != nil {
			t.Fatal(err)
		}

		// triathlon's header part is overwritten.
		if diff := cmp.Diff(f.buf[:12], testEncode[:12]); diff != "" {
			t.Fatal(diff)
		}
	})
}

type writeSeekerStub struct {
	buf []byte
	off int64
}

var _ io.Writer = (*writeSeekerStub)(nil)
var _ io.Seeker = (*writeSeekerStub)(nil)

func (f *writeSeekerStub) Write(p []byte) (n int, err error) {
	if len(f.buf[f.off:]) < len(p) {
		buf := make([]byte, f.off+int64(len(p)))
		copy(buf, f.buf)
		f.buf = buf
	}
	n = copy(f.buf[f.off:], p)
	f.off += int64(len(p))
	return n, nil
}

func (f *writeSeekerStub) Seek(off int64, whence int) (int64, error) {
	switch whence {
	case io.SeekCurrent:
		l := int64(len(f.buf))
		l2 := f.off + off
		if l2 < 0 {
			return 0, os.ErrInvalid
		}
		if l2 > l {
			buf := make([]byte, l2)
			copy(buf, f.buf)
			f.buf = buf
		}
		f.off = l2
	case io.SeekStart:
		l := int64(len(f.buf))
		if off < 0 {
			return 0, os.ErrInvalid
		}
		if off > l {
			buf := make([]byte, off)
			copy(buf, f.buf)
			f.buf = buf
		}
		f.off = off
	case io.SeekEnd:
		l := int64(len(f.buf))
		l2 := l + off
		if l2 < 0 {
			return 0, os.ErrInvalid
		}
		if l2 > l {
			buf := make([]byte, l2)
			copy(buf, f.buf)
			f.buf = buf
		}
		f.off = l2
	default:
		return 0, os.ErrInvalid
	}
	return off, nil
}

type writerAtStub struct{ buf []byte }

var _ io.Writer = (*writerAtStub)(nil)
var _ io.WriterAt = (*writerAtStub)(nil)

func (w *writerAtStub) Write(p []byte) (n int, err error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}

func (w *writerAtStub) WriteAt(p []byte, pos int64) (int, error) {
	if pos < 0 {
		return 0, os.ErrInvalid
	}
	l := int64(len(w.buf))
	l2 := int64(len(p)) + pos
	if l2 > l {
		buf := make([]byte, l2)
		copy(buf, w.buf)
		w.buf = buf
	}
	n := copy(w.buf[pos:], p)
	return n, nil
}

type fnValidate func(mesg *proto.Message) error

func (f fnValidate) Validate(mesg *proto.Message) error { return f(mesg) }
func (f fnValidate) Reset()                             {}

var (
	fnValidateOK  = fnValidate(func(mesg *proto.Message) error { return nil })
	fnValidateErr = fnValidate(func(mesg *proto.Message) error { return fmt.Errorf("validate error") })
)

func TestOptions(t *testing.T) {
	tt := []struct {
		name     string
		opts     []Option
		expected options
	}{
		{
			name: "defaultOptions",
			opts: nil,
			expected: options{
				multipleLocalMessageType: 0,
				endianness:               0,
				messageValidator:         NewMessageValidator(),
				writeBufferSize:          defaultWriteBufferSize,
			},
		},
		{
			name: "with options: normal header",
			opts: []Option{
				WithBigEndian(),
				WithNormalHeader(20),
				WithProtocolVersion(proto.V2),
				WithMessageValidator(fnValidateOK),
				WithWriteBufferSize(8192),
			},
			expected: options{
				multipleLocalMessageType: 15,
				endianness:               1,
				protocolVersion:          proto.V2,
				messageValidator:         fnValidateOK,
				headerOption:             headerOptionNormal,
				writeBufferSize:          8192,
			},
		},
		{
			name: "with options: compressed timestamp header",
			opts: []Option{
				WithBigEndian(),
				WithProtocolVersion(proto.V2),
				WithCompressedTimestampHeader(),
				WithMessageValidator(fnValidateOK),
			},
			expected: options{
				multipleLocalMessageType: 0,
				endianness:               1,
				protocolVersion:          proto.V2,
				messageValidator:         fnValidateOK,
				headerOption:             headerOptionCompressedTimestamp,
				writeBufferSize:          defaultWriteBufferSize,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil, tc.opts...)

			cmpOpts := []cmp.Option{
				cmp.AllowUnexported(options{}),
			}

			if tc.opts == nil { // defaultOptions
				cmpOpts = append(cmpOpts,
					cmp.Transformer("MessageValidator", func(mv MessageValidator) string {
						return fmt.Sprintf("%T", mv) // compare type only
					}),
				)
			} else {
				cmpOpts = append(cmpOpts,
					cmp.Transformer("MessageValidator", func(v MessageValidator) uintptr {
						return reflect.ValueOf(v).Pointer() // should reference the same object
					}),
				)
			}

			if diff := cmp.Diff(enc.options, tc.expected, cmpOpts...); diff != "" {
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
		err  error
	}{
		{name: "encode with nil", w: nil, err: ErrNilWriter},
		{name: "encode with writer", w: fnWriteOK},
		{name: "encode with writerAt", w: mockWriterAt{fnWriteOK, fnWriteAtOK}},
		{name: "encode with writeSeeker", w: mockWriteSeeker{fnWriteOK, fnSeekOK}},
	}

	fit := proto.FIT{
		Messages: []proto.Message{
			{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity.Byte()),
			}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			err := enc.Encode(&fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}

	// Test same logic for EncodeWithContext
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w)
			err := enc.EncodeWithContext(context.Background(), &fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}
}

type encodeWithDirectUpdateTestCase struct {
	name string
	fit  *proto.FIT
	w    io.Writer
	err  error
}

func makeEncodeWithDirectUpdateStrategyTableTest() []encodeWithDirectUpdateTestCase {
	return []encodeWithDirectUpdateTestCase{
		{
			name: "happy flow coverage",
			fit: &proto.FIT{Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				}},
			}},
			w: mockWriterAt{fnWriteOK, fnWriteAtOK},
		},
		{
			name: "encode header error",
			fit:  &proto.FIT{},
			w:    mockWriterAt{fnWriteErr, fnWriteAtErr},
			err:  io.EOF,
		},
		{
			name: "encode messages error",
			fit:  &proto.FIT{},
			w:    mockWriterAt{fnWriteOK, fnWriteAtErr},
			err:  ErrEmptyMessages,
		},
		{
			name: "encode crc error",
			fit: &proto.FIT{Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				}},
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
			err: io.EOF,
		},
		{
			name: "update error",
			fit: &proto.FIT{FileHeader: proto.FileHeader{Size: 14, DataSize: 100, DataType: proto.DataTypeFIT}, Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
				}},
			}},
			w:   mockWriterAt{fnWriteOK, fnWriteAtErr},
			err: io.EOF,
		},
	}
}

func TestEncodeWithDirectUpdateStrategy(t *testing.T) {
	tt := makeEncodeWithDirectUpdateStrategyTableTest()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w, WithWriteBufferSize(0))
			err := enc.Encode(tc.fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}

	// Test same logic for EncodeWithContext
	tt2 := makeEncodeWithDirectUpdateStrategyTableTest()

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w, WithWriteBufferSize(0))
			err := enc.EncodeWithContext(context.Background(), tc.fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
		})
	}
}

type encodeWithEarlyCheckStrategyTestCase struct {
	name string
	fit  *proto.FIT
	w    io.Writer
	err  error
}

func makeEncodeWithEarlyCheckStrategy() []encodeWithEarlyCheckStrategyTestCase {
	return []encodeWithEarlyCheckStrategyTestCase{
		{
			name: "happy flow coverage",
			fit:  &proto.FIT{Messages: []proto.Message{{}}},
			w:    fnWriteOK,
		},
		{
			name: "calculate data size error",
			fit:  &proto.FIT{Messages: []proto.Message{}},
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
			fit: &proto.FIT{Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
				}},
			}},
			w:   fnWriteErr,
			err: io.EOF,
		},
		{
			name: "encode messages error",
			fit: &proto.FIT{Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
				}},
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
			fit:  &proto.FIT{Messages: []proto.Message{{}}},
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

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			enc := New(tc.w,
				WithMessageValidator(fnValidateOK),
				WithWriteBufferSize(0),
			)
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
			enc := New(tc.w,
				WithMessageValidator(fnValidateOK),
				WithWriteBufferSize(0),
			)
			err := enc.EncodeWithContext(context.Background(), tc.fit)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestUpdateHeader(t *testing.T) {
	tt := []struct {
		name      string
		w         io.Writer
		n         int64
		dataSize  uint32
		header    proto.FileHeader
		headerPos int64
		expect    []byte
		err       error
	}{
		{
			name:     "data size not changed",
			dataSize: 1000,
			header:   proto.FileHeader{DataSize: 1000},
			w:        nil,
		},
		{
			name:     "writerAt flow",
			header:   proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT},
			w:        mockWriterAt{fnWriteOK, fnWriteAtOK},
			dataSize: 1,
		},
		{
			name: "writeSeeker using stub",
			header: proto.FileHeader{
				Size:            12,
				ProtocolVersion: byte(proto.V1),
				ProfileVersion:  profile.Version,
				DataType:        proto.DataTypeFIT,
			},
			n:         12 + 4,
			headerPos: 4,
			dataSize:  2,
			w: &writeSeekerStub{
				buf: make([]byte, 12+4), // n
				off: 12 + 4,             // n
			},
			expect: func() []byte {
				h := proto.FileHeader{
					Size:            12,
					ProtocolVersion: byte(proto.V1),
					ProfileVersion:  profile.Version,
					DataType:        proto.DataTypeFIT,
					DataSize:        2, // updated
				}
				b, _ := h.MarshalAppend(make([]byte, 4)) // headerPos = 4
				return b
			}(),
		},
		{
			name: "writerAt using stub",
			header: proto.FileHeader{
				Size:            12,
				ProtocolVersion: byte(proto.V1),
				ProfileVersion:  profile.Version,
				DataType:        proto.DataTypeFIT,
			},
			headerPos: 2,
			dataSize:  2,
			w:         &writerAtStub{},
			expect: func() []byte {
				h := proto.FileHeader{
					Size:            12,
					ProtocolVersion: byte(proto.V1),
					ProfileVersion:  profile.Version,
					DataType:        proto.DataTypeFIT,
					DataSize:        2, // updated
				}
				b, _ := h.MarshalAppend(make([]byte, 2)) // headerPos = 2
				return b
			}(),
		},
		{
			name:     "writeSeeker happy flow",
			header:   proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT},
			w:        mockWriteSeeker{fnWriteOK, fnSeekOK},
			dataSize: 1,
		},
		{
			name:     "writeSeeker error on seek",
			header:   proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT},
			w:        mockWriteSeeker{fnWriteOK, fnSeekErr},
			dataSize: 1,
			err:      io.EOF,
		},
		{
			name:     "writeSeeker error on write",
			header:   proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT},
			w:        mockWriteSeeker{fnWriteErr, fnSeekOK},
			dataSize: 1,
			err:      io.EOF,
		},
		{
			name:   "writeSeeker error on seek for resetting offset",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT},
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
			dataSize: 1,
			err:      io.EOF,
		},
		{
			name:   "encoder internal error caused by bad implementation",
			header: proto.FileHeader{Size: 14, DataType: proto.DataTypeFIT, DataSize: 1},
			w:      nil,
			err:    errInternal,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			enc := New(tc.w, WithWriteBufferSize(0))
			enc.n = tc.n
			enc.dataSize = tc.dataSize
			enc.lastFileHeaderPos = tc.headerPos

			err := enc.updateFileHeader(&tc.header)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
			switch w := tc.w.(type) {
			case *writeSeekerStub:
				if string(w.buf) != string(tc.expect) {
					t.Fatalf("\n%v\n%v", w.buf, tc.expect)
				}
			case *writerAtStub:
				if string(w.buf) != string(tc.expect) {
					t.Fatalf("\n%v\n%v", w.buf, tc.expect)
				}
			}
		})
	}
}

func TestEncodeHeader(t *testing.T) {
	tt := []struct {
		name            string
		opts            []Option
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

				binary.LittleEndian.PutUint16(b[2:4], profile.Version)

				crc := crc16.New(nil)
				crc.Write(b[:12])
				binary.LittleEndian.PutUint16(b[12:14], crc.Sum16())

				return b
			}(),
			protocolVersion: proto.V1,
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
		{
			name:            "force use protocol version from Option",
			protocolVersion: proto.V2,
			opts: []Option{
				WithProtocolVersion(proto.V2),
			},
			header: proto.FileHeader{
				Size:            14,
				ProtocolVersion: byte(proto.V1),
				ProfileVersion:  2135,
				DataSize:        136830,
				DataType:        ".FIT",
				CRC:             21830,
			},
			b: []byte{
				14,
				32, // Previously 16
				87, 8,
				126, 22, 2, 0,
				46, 70, 73, 84,
				185, 85, // Previously was 70, 85,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bytebuf := new(bytes.Buffer)
			enc := New(bytebuf, append(tc.opts, WithWriteBufferSize(0))...)
			_ = enc.encodeFileHeader(&tc.header)

			if diff := cmp.Diff(bytebuf.Bytes(), tc.b); diff != "" {
				t.Fatal(diff)
			}

			if enc.protocolValidator.ProtocolVersion() != tc.protocolVersion {
				t.Fatalf("expected protocol version: %v, got: %v",
					tc.protocolVersion, enc.options.protocolVersion)
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
			mesg: proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
			}},
			w: fnWriteOK,
		},
		{
			name: "encode message with big-endian",
			mesg: proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
			}},
			w:          fnWriteOK,
			opts:       []Option{WithBigEndian()},
			endianness: bigEndian,
		},
		{
			name: "encode message with header normal multiple local message type happy flow",
			opts: []Option{
				WithNormalHeader(2),
			},
			mesg: proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
			}},
			w: fnWriteOK,
		},
		{
			name: "encode message with compressed timestamp header happy flow",
			opts: []Option{
				WithCompressedTimestampHeader(),
			},
			mesg: proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
			}},
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
						Name:     factory.NameUnknown,
						Type:     profile.Sint64, // int64 type is ilegal for protocol v1.0
						BaseType: profile.Sint64.BaseType(),
					},
					Value: proto.Int64(1234),
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
						Name:     factory.NameUnknown,
						Type:     profile.Sint64, // int64 type is ilegal for protocol v1.0
						BaseType: profile.Sint64.BaseType(),
					},
					Value: proto.Int64(1234),
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
			err: proto.ErrTypeNotSupported,
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
			tc.opts = append(tc.opts, WithWriteBufferSize(0))
			enc := New(tc.w, tc.opts...)
			// Protocol Version now is set on encodeFileHeader as we allow dynamic protocol version
			// based on FileHeader. This by pass it since we don't encode file header.
			enc.protocolValidator.SetProtocolVersion(enc.options.protocolVersion)
			err := enc.encodeMessage(&tc.mesg)
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

	// Tests that does not fit in test table:
	t.Run("encode message with early check must place timestamp field back to its original index", func(t *testing.T) {
		pivotTime := time.Now()
		mesg := proto.Message{
			Fields: []proto.Field{
				factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(80)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(pivotTime)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordAltitude).WithValue(uint16((166.0 + 500.0) * 5.0)),
			},
		}
		expected := mesg.Clone()

		enc := New(io.Discard,
			WithCompressedTimestampHeader(),
			WithMessageValidator(fnValidateOK),
			WithWriteBufferSize(0), // Direct write
		)
		enc.timestampReference = datetime.ToUint32(pivotTime)

		err := enc.encodeMessage(&mesg)
		if err != nil {
			t.Fatalf("expected err: nil, got: %v", err)
		}

		if diff := cmp.Diff(mesg, expected,
			cmp.Transformer("Message", func(m proto.Message) proto.Message {
				m.Header = 0 // Clear
				return m
			}),
			cmp.Transformer("Value", func(v proto.Value) any { return v.Any() }),
		); diff != "" {
			t.Fatal(diff)
		}
	})
}

func TestEncodeMessageWithMultipleLocalMessageType(t *testing.T) {
	now := time.Now()
	mesgs := []proto.Message{
		{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(time.Second))),
			factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(70)),
		}},
		{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
			factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
		}},
	}

	t.Run("multiple local mesg type", func(t *testing.T) {
		// We have 3 messages with differents field definitions,
		// this should produces different localMesgNum in header.

		mesgs := append(mesgs[:0:0], mesgs...)
		for i := range mesgs {
			mesgs[i] = mesgs[i].Clone()
		}

		buf := new(bytes.Buffer)
		enc := New(buf, WithNormalHeader(2), WithWriteBufferSize(0))
		for i, mesg := range mesgs {
			buf.Reset()
			err := enc.encodeMessage(&mesg)
			if err != nil {
				t.Fatal(err)
			}

			mesgDefHeader := buf.Bytes()
			expectedHeader := (mesgDefHeader[0] &^ proto.LocalMesgNumMask) | byte(i)
			if mesgDefHeader[0] != expectedHeader {
				t.Fatalf("[%d] expected 0b%08b, got: 0b%08b", i, expectedHeader, mesgDefHeader[0])
			}
		}

		// add 4th mesg, header should be 0, reset.
		mesg := proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
		}}
		buf.Reset()
		if err := enc.encodeMessage(&mesg); err != nil {
			t.Fatal(err)
		}
		mesgDefHeader := buf.Bytes()
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
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(uint16(typedef.ManufacturerGarmin)),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdProduct).WithValue(uint16(typedef.GarminProductEdge1030)),
				}},
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
			mesgs: []proto.Message{{Num: mesgnum.Record}},
			err:   ErrMissingFileId,
		},
	}
}

func TestEncodeMessages(t *testing.T) {
	tt := makeEncodeMessagesTableTest()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(io.Discard)
			err := enc.encodeMessages(tc.mesgs)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
		})
	}

	// Test same logic for encodeMessagesWithContext
	tt2 := makeEncodeMessagesTableTest()

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(io.Discard)
			err := enc.encodeMessagesWithContext(context.Background(), tc.mesgs)
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
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerGarmin),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(time.Second))), // +1),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))), // +2),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(32 * time.Second))), // +32 rollove),
				}},
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
			name: "compress timestamp in header happy flow: roll over occurred exactly after 32 seconds",
			mesgs: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerGarmin),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(32 * time.Second))),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(33 * time.Second))),
				}},
			},
			headers: []byte{
				proto.MesgNormalHeaderMask, // file_id: has no timestamp
				proto.MesgNormalHeaderMask, // record: the message containing timestamp reference prior to the use of compressed header.
				proto.MesgNormalHeaderMask, // record: roll over has occurred, the timestamp is used new timestamp reference.
				proto.MesgCompressedHeaderMask | (offset+1)&proto.CompressedTimeMask,
			},
		},
		{
			name: "timestamp less than DateTimeMin",
			mesgs: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerGarmin),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(uint32(1234)),
				}},
			},
			headers: []byte{
				proto.MesgNormalHeaderMask,
				proto.MesgNormalHeaderMask,
			},
		},
		{
			name: "timestamp wrong type not uint32 or typedef.DateTime",
			mesgs: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerGarmin),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(typedef.DateTime(datetime.ToUint32(now))),
				}},
			},
			headers: []byte{
				proto.MesgNormalHeaderMask,
				proto.MesgNormalHeaderMask,
			},
		},
		{
			name: "timestamp wrong type not uint32 or typedef.DateTime",
			mesgs: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerGarmin),
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(now), // time.Time{),
				}},
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

// bufferAt wraps bytes.Buffer to enable WriteAt for faster encoding.
type bufferAt struct{ *bytes.Buffer }

func (b *bufferAt) WriteAt(p []byte, off int64) (n int, err error) {
	return copy(b.Bytes()[off:], p), nil
}

func TestEncodeWithCompressedTimestampHeader(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // uint32(1073001600) -> []byte{128, 180, 244, 63}
	fnCreateFIT := func() proto.FIT {
		return proto.FIT{
			Messages: []proto.Message{
				{Num: mesgnum.FileId, Fields: []proto.Field{
					factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerDevelopment.Uint16()),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(60)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(1 * time.Second))),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(60)),
				}},
				{Num: mesgnum.Record, Fields: []proto.Field{
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
					factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(60)),
				}},
			},
		}
	}

	expected := []byte{ // Records only
		proto.MesgDefinitionMask, 0, 0, 0, 0, 1, fieldnum.FileIdManufacturer, basetype.Uint16.Size(), byte(basetype.Uint16),
		/* FileId */ 0, 255, 0,
		proto.MesgDefinitionMask, 0, 0, 20, 0, 2,
		/* ~ */ fieldnum.RecordTimestamp, basetype.Uint32.Size(), byte(basetype.Uint32),
		/* ~ */ fieldnum.RecordHeartRate, basetype.Uint8.Size(), byte(basetype.Uint8),
		/* Record 0 */ 0, 128, 180, 244, 63, 60, // Timestamp + HeartRate
		proto.MesgDefinitionMask, 0, 0, 20, 0, 1, fieldnum.RecordHeartRate, basetype.Uint8.Size(), byte(basetype.Uint8),
		/* Record 1 */ 0 | proto.MesgCompressedHeaderMask | 1, 60, // HeartRate only
		/* Record 2 */ 0 | proto.MesgCompressedHeaderMask | 2, 60, // HeartRate only
	}

	tt := []struct {
		name string
		w    interface {
			io.Writer
			Bytes() []byte
		}
		fit proto.FIT
	}{
		{
			name: "early check strategy",
			w:    new(bytes.Buffer),
			fit:  fnCreateFIT(),
		},
		{
			name: "direct update strategy",
			w:    &bufferAt{new(bytes.Buffer)},
			fit:  fnCreateFIT(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(tc.w,
				WithCompressedTimestampHeader(),
				WithWriteBufferSize(0),
			)
			if err := enc.Encode(&tc.fit); err != nil {
				t.Fatalf("expected error nil, got: %v", err)
			}

			b := tc.w.(interface{ Bytes() []byte }).Bytes()
			b = b[14:]       // omit FileHeader
			b = b[:len(b)-2] // omit CRC

			if diff := cmp.Diff(expected, b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestEncodeMessagesWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mesgs := []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivity)),
		}},
	}
	enc := New(nil, WithWriteBufferSize(0))
	err := enc.encodeMessagesWithContext(ctx, mesgs)
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
			_, err := New(tc.w, WithWriteBufferSize(0)).StreamEncoder()
			if !errors.Is(err, tc.err) {
				t.Errorf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}

		})
	}
}

func TestReset(t *testing.T) {
	tt := []struct {
		name     string
		w        io.Writer
		opts     []Option
		expected *Encoder
	}{
		{
			name: "reset with options",
			w:    io.Discard,
			opts: []Option{
				WithBigEndian(),
				WithProtocolVersion(proto.V2),
				WithNormalHeader(15),
				WithMessageValidator(fnValidateOK),
			},
			expected: New(io.Discard,
				WithBigEndian(),
				WithProtocolVersion(proto.V2),
				WithNormalHeader(15),
				WithMessageValidator(fnValidateOK)),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			enc := New(nil)
			enc.Reset(tc.w, tc.opts...)
			if diff := cmp.Diff(enc, tc.expected,
				cmp.AllowUnexported(options{}),
				cmp.AllowUnexported(Encoder{}),
				cmpopts.IgnoreUnexported(bufio.Writer{}),
				cmpopts.IgnoreUnexported(writerAt{}),
				cmpopts.IgnoreUnexported(writeSeeker{}),
				cmp.FilterValues(func(x, y MessageValidator) bool { return true }, cmp.Ignore()),
				cmp.FilterValues(func(x, y hash.Hash16) bool { return true }, cmp.Ignore()),
				cmp.FilterValues(func(x, y *proto.Validator) bool { return true }, cmp.Ignore()),
				cmp.FilterValues(func(x, y *lru) bool { return true }, cmp.Ignore()),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
