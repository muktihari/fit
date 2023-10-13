// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/listener"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
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
	fromOfficialSDK   = filepath.Join(testdata, "from_official_sdk")
)

func TestDecodeRealFiles(t *testing.T) {
	t.Run("testdata/from_official_sdk", func(t *testing.T) {
		_ = filepath.Walk(fromOfficialSDK, func(path string, info fs.FileInfo, _ error) error {
			if info.IsDir() {
				return nil
			}

			ext := filepath.Ext(info.Name())
			if strings.ToLower(ext) != ".fit" {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				t.Errorf("filename: %s: %v", info.Name(), err)
				return nil
			}
			defer f.Close()

			dec := New(bufio.NewReader(f))

			_, err = dec.DecodeWithContext(context.Background())
			if err != nil {
				// NOTE: Doubts exist regarding the integrity of these files.
				if info.Name() == "Settings.fit" || info.Name() == "WeightScaleMultiUser.fit" {
					if errors.Is(err, ErrCRCChecksumMismatch) {
						return nil
					}
				}

				t.Errorf("filename: %s: %v", info.Name(), err)

				return nil
			}

			return nil
		})
	})
}

type fnMesgListener func(mesg proto.Message)

func (f fnMesgListener) OnMesg(mesg proto.Message) { f(mesg) }

type fnMesgDefListener func(mesgDef proto.MessageDefinition)

func (f fnMesgDefListener) OnMesgDef(mesgDef proto.MessageDefinition) { f(mesgDef) }

func TestOptions(t *testing.T) {
	// predefined
	decoderFactory := factory.New()
	mesglis := fnMesgListener(func(mesg proto.Message) {})
	mesgDefLis := fnMesgDefListener(func(mesgDef proto.MessageDefinition) {})

	tt := []struct {
		name    string
		opts    []Option
		options *options
	}{
		{
			name: "defaultOptions",
			options: &options{
				factory:               factory.StandardFactory(),
				shouldChecksum:        true,
				broadcastOnly:         false,
				shouldExpandComponent: true,
			},
		},
		{
			name: "with options",
			opts: []Option{
				WithFactory(decoderFactory),
				WithIgnoreChecksum(),
				WithMesgListener(mesglis), WithMesgListener(mesglis),
				WithMesgDefListener(mesgDefLis), WithMesgDefListener(mesgDefLis),
				WithBroadcastOnly(),
				WithNoComponentExpansion(),
			},
			options: &options{
				factory:               decoderFactory,
				shouldChecksum:        false,
				mesgListeners:         []listener.MesgListener{mesglis, mesglis},
				mesgDefListeners:      []listener.MesgDefListener{mesgDefLis, mesgDefLis},
				broadcastOnly:         true,
				shouldExpandComponent: false,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(nil, tc.opts...)

			if diff := cmp.Diff(dec.options, tc.options,
				cmp.AllowUnexported(options{}),
				cmp.Transformer("factory", func(t Factory) uintptr {
					return reflect.ValueOf(t).Pointer()
				}),
				cmp.Comparer(func(a, b []listener.MesgListener) bool {
					if len(a) != len(b) {
						return false
					}
					for i := range a {
						if reflect.ValueOf(a[i]).Pointer() != reflect.ValueOf(b[i]).Pointer() {
							return false
						}
					}
					return true
				}),
				cmp.Comparer(func(a, b []listener.MesgDefListener) bool {
					if len(a) != len(b) {
						return false
					}
					for i := range a {
						if reflect.ValueOf(a[i]).Pointer() != reflect.ValueOf(b[i]).Pointer() {
							return false
						}
					}
					return true
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

type fnReader func(b []byte) (n int, err error)

func (f fnReader) Read(b []byte) (n int, err error) { return f(b) }

var (
	fnReaderOK  = fnReader(func(b []byte) (n int, err error) { return len(b), nil })
	fnReaderErr = fnReader(func(b []byte) (n int, err error) { return 0, io.EOF })
)

func TestDecodeHeaderOnce(t *testing.T) {
	var r io.Reader = func() io.Reader {
		fnInstances := []io.Reader{
			fnReader(func(b []byte) (n int, err error) {
				copy(b, []byte{14}) // header size: 14
				return 1, nil
			}),
			fnReader(func(b []byte) (n int, err error) {
				copy(b, []byte{
					32,    // protocol version: 32
					84, 8, // profile version: 2132
					214, 204, 9, 0, // data size: 642262
					46, 70, 73, 84, // data type: .FIT
					56, 50, // crc: 12856
				})
				return 13, nil
			}),
		}
		index := 0
		return fnReader(func(b []byte) (n int, err error) {
			f := fnInstances[index]
			index++
			return f.Read(b)
		})
	}()

	dec := New(r)
	err1 := dec.decodeHeaderOnce()
	err2 := dec.decodeHeaderOnce()
	if err1 != nil || !errors.Is(err1, err2) {
		t.Fatalf("expected %v: err1 == err2, got: %v == %v", nil, err1, err2)
	}

	dec = New(fnReaderErr)
	err1 = dec.decodeHeaderOnce()
	err2 = dec.decodeHeaderOnce()
	if !errors.Is(err1, io.EOF) || !errors.Is(err1, err2) {
		t.Fatalf("expected %v: err1 == err2, got: %v == %v", io.EOF, err1, err2)
	}
}

func TestPeekFileId(t *testing.T) {
	fit, buf := createFitForTest()

	tt := []struct {
		name   string
		r      io.Reader
		fileId *mesgdef.FileId
		err    error
	}{
		{
			name: "peek file id happy flow",
			r: func() io.Reader {
				buf, cur := slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fileId: mesgdef.NewFileId(fit.Messages[0]),
		},
		{
			name: "peek file id decode header return error",
			r: func() io.Reader {
				return fnReader(func(b []byte) (n int, err error) {
					return 0, io.EOF
				})
			}(),
			fileId: mesgdef.NewFileId(fit.Messages[0]),
			err:    io.EOF,
		},
		{
			name: "peek file id decode message return error",
			r: func() io.Reader {
				buf, cur := slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 14 { // only decode header
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fileId: mesgdef.NewFileId(fit.Messages[0]),
			err:    io.EOF,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			fileId, err := dec.PeekFileId()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(fileId, tc.fileId); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func headerForTest() proto.FileHeader {
	return proto.FileHeader{
		Size:            14,
		ProtocolVersion: 32,
		ProfileVersion:  2132,
		DataSize:        642262,
		DataType:        proto.DataTypeFIT,
		CRC:             12856,
	}
}

func createFitForTest() (proto.Fit, []byte) {
	fit := proto.Fit{
		FileHeader: headerForTest(),
		Messages: []proto.Message{
			factory.CreateMesgOnly(mesgnum.FileId).WithFields(
				factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(uint8(typedef.FileActivity)),
			),
			factory.CreateMesgOnly(mesgnum.DeveloperDataId).WithFields(
				factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdDeveloperDataIndex).WithValue(uint8(0)),
				factory.CreateField(mesgnum.DeveloperDataId, fieldnum.DeveloperDataIdApplicationId).WithValue([]byte{0, 1, 2, 3}),
			),
			factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
			),
			factory.CreateMesgOnly(mesgnum.Record).WithFields(
				factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
			),
			factory.CreateMesgOnly(mesgnum.Record).
				WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(77)),
				).
				WithDeveloperFields(
					proto.DeveloperField{
						DeveloperDataIndex: 0,
						Num:                0,
						Size:               1,
						Name:               "Heart Rate",
						NativeMesgNum:      mesgnum.Record,
						NativeFieldNum:     fieldnum.RecordHeartRate,
						Type:               basetype.Uint8,
						Value:              uint8(100),
					},
				),
		},
	}

	bytesbuffer := new(bytes.Buffer)
	b, _ := fit.FileHeader.MarshalBinary()
	bytesbuffer.Write(b)

	// Marshal and calculate data size and crc checksum
	crc16checker := crc16.New(crc16.MakeFitTable())
	for i := range fit.Messages {
		mesg := fit.Messages[i]
		mesgDef := proto.CreateMessageDefinition(&mesg)
		b, _ := mesgDef.MarshalBinary()
		bytesbuffer.Write(b)
		crc16checker.Write(b)

		b, err := mesg.MarshalBinary()
		if err != nil {
			panic(err)
		}
		bytesbuffer.Write(b)
		crc16checker.Write(b)
	}

	// From here onward we'll use []byte instead of bytesbuffer.
	b = bytesbuffer.Bytes()

	// Calculate messages data size and update the file header
	fit.FileHeader.DataSize = uint32(bytesbuffer.Len() - 14)
	var dataSize = make([]byte, 4)
	binary.LittleEndian.PutUint32(dataSize, fit.FileHeader.DataSize)

	// Update file header data size in []byte form as well
	copy(b[4:8], dataSize)

	// Update Fit File CRC
	fit.CRC = crc16checker.Sum16()
	crc16checker.Reset()
	var crc = make([]byte, 2)
	binary.LittleEndian.PutUint16(crc, fit.CRC)
	b = append(b, crc...) // append crc to the []byte form.

	// Calculate FileHeader CRC checksum and update it
	crc16checker.Write(b[:12])
	fit.FileHeader.CRC = crc16checker.Sum16()
	crc16checker.Reset()

	fileHeaderCRC := make([]byte, 2)
	binary.LittleEndian.PutUint16(fileHeaderCRC, fit.FileHeader.CRC)
	copy(b[12:14], fileHeaderCRC) // update FileHeader CRC

	return fit, b
}

func TestNext(t *testing.T) {
	// NOTE: Testing next should include at least one complete Decode process without fail,
	// 		 ensuring we are really got next chained proto.

	// Setup
	_, buf := createFitForTest()

	// New header of the next chained Fit sequences.
	header := headerForTest()
	b, _ := header.MarshalBinary()
	buf = append(buf, b...)

	r := func() io.Reader {
		bbbuf := buf
		buf, cur := make([]byte, len(bbbuf)), 0
		copy(buf, bbbuf)
		return fnReader(func(b []byte) (n int, err error) {
			if cur == len(buf) {
				return 0, io.EOF
			}
			cur += copy(b, buf[cur:cur+len(b)])
			return len(b), nil
		})
	}()

	// Test Begin
	dec := New(r)
	prevAccumulator := dec.accumulator

	_, err := dec.DecodeWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Check whether after decode, fields are reset and next sequence is retrieved.

	if !dec.Next() {
		t.Fatalf("should have next, return false")
	}

	if prevAccumulator == dec.accumulator {
		t.Fatalf("expected new accumulator got same")
	}

	if dec.crc16.Sum16() != 0 { // not necessary since reset every decode header anyway, but let's just add it
		t.Fatalf("crc16 should reset")
	}

	if dec.fileHeader != header {
		t.Fatalf("header should be replaced with new decoded header")
	}

	for i := range dec.localMessageDefinitions {
		if dec.localMessageDefinitions[i] != nil {
			t.Errorf("message definition index %d should be nil", i)
		}
	}

	if len(dec.messages) != 0 {
		t.Fatalf("messages should be empty")
	}

	if dec.crc != 0 {
		t.Fatalf("crc should be zero")
	}

	if dec.cur != 0 {
		t.Fatalf("cur should be zero")
	}

	if dec.timestamp != 0 {
		t.Fatalf("timestamp should be zero")
	}

	if dec.lastTimeOffset != 0 {
		t.Fatalf("lastTimeOffset should be zero")
	}
}

func TestDecodeExported(t *testing.T) {
	tt := []struct {
		name string
		r    io.Reader
		ctx  context.Context
		err  error
	}{
		{
			name: "context nil",
			r:    fnReaderErr,
			ctx:  nil,
			err:  io.EOF,
		},
		{
			name: "context DeadlineExceeded",
			r: fnReader(func(b []byte) (n int, err error) {
				time.Sleep(1 * time.Second) // Let's make our process take longer, 1s per reader Read call.
				return len(b), nil
			}),
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
				_ = cancel // it's okay to discard this since only for testing.
				return ctx
			}(),
			err: context.DeadlineExceeded,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			_, err := dec.DecodeWithContext(tc.ctx)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestDecodeUnexported(t *testing.T) {
	fit, buf := createFitForTest()

	tt := []struct {
		name string
		r    io.Reader
		fit  proto.Fit
		err  error
	}{
		{
			name: "decode happy flow",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fit: fit,
		},
		{
			name: "decode messages return error",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 14 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fit: fit,
			err: io.EOF,
		},
		{
			name: "decode crc return error",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf)-2 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fit: fit,
			err: io.EOF,
		},
		{
			name: "decode crc checksum mismatch",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					ln := len(buf) - 2
					if cur == ln {
						copy(b, []byte{0, 0}) // zeroing crc
						return 2, nil
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fit: fit,
			err: ErrCRCChecksumMismatch,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			fit, err := dec.DecodeWithContext(context.Background())
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(*fit, tc.fit); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDecodeHeader(t *testing.T) {
	fit, buf := createFitForTest()

	tt := []struct {
		name   string
		r      io.Reader
		header proto.FileHeader
		err    error
	}{
		{
			name: "decode header happy flow",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			header: fit.FileHeader,
		},
		{
			name: "decode header invalid size",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				buf[0] = 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: ErrNotAFitFile,
		},
		{
			name: "decode header invalid size",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				buf = buf[:1] // trimmed
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: io.EOF,
		},
		{
			name: "decode invalid protocol",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				buf[1] = 100 // invalid protocol
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: proto.ErrProtocolVersionNotSupported,
		},
		{
			name: "decode data type not `.FIT`",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				copy(buf[5:9], []byte("F.IT"))
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: ErrNotAFitFile,
		},
		{
			name: "decode crc == 0x000",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				buf[12], buf[13] = 0, 0

				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			header: func() proto.FileHeader {
				header := fit.FileHeader
				header.CRC = 0
				return header
			}(),
		},
		{
			name: "decode crc mismatch",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				buf[12], buf[13] = 0, 1

				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: ErrCRCChecksumMismatch,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			err := dec.decodeHeader()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(dec.fileHeader, tc.header); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDecodeMessageDefinition(t *testing.T) {
	fit, buf := createFitForTest()

	tt := []struct {
		name    string
		r       io.Reader
		opts    []Option
		header  byte
		mesgDef proto.MessageDefinition
		err     error
	}{
		{
			name: "decode message definition happy flow",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf[15:]), 0 // trim header
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			opts: []Option{
				WithMesgDefListener(fnMesgDefListener(func(mesgDef proto.MessageDefinition) {})),
			},
			header:  proto.MesgDefinitionMask,
			mesgDef: proto.CreateMessageDefinition(&fit.Messages[0]), // file_id
		},
		{
			name: "decode read return io.EOF when retrieving init data",
			r:    fnReaderErr,
			err:  io.EOF,
		},
		{
			name: "decode read return io.EOF when retrieving field data",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf[15:]), 0 // trim header
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 5 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: io.EOF,
		},
		{
			name: "decode read return io.EOF when retrieving developer field size",
			r: func() io.Reader {
				buf := []byte{0, 0, 0, 0, 1 /* n fields */, 0, 0, 0}
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			header: proto.MesgDefinitionMask | proto.DevDataMask,
			err:    io.EOF,
		},
		{
			name: "decode read return io.EOF when retrieving developer field size",
			r: func() io.Reader {
				buf := []byte{0, 0, 0, 0, 1 /* n fields */, 0, 0, 0, 1 /* dev fields */}
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			header: proto.MesgDefinitionMask | proto.DevDataMask,
			err:    io.EOF,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r, tc.opts...)
			err := dec.decodeMessageDefinition(tc.header)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			mesgDef := *dec.localMessageDefinitions[proto.MesgDefinitionMask&proto.LocalMesgNumMask]
			if diff := cmp.Diff(mesgDef, tc.mesgDef); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDecodeMessageData(t *testing.T) {
	tt := []struct {
		name             string
		r                io.Reader
		header           byte
		mesgdef          *proto.MessageDefinition
		fieldDescription *mesgdef.FieldDescription
		mesg             proto.Message
		err              error
	}{
		{
			name:   "decode message data normal header happy flow",
			r:      fnReaderOK,
			header: 0,
			mesgdef: &proto.MessageDefinition{
				Header: proto.MesgDefinitionMask,
				// LocalMesgNum: 0,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordTimestamp,
						Size:     4,
						BaseType: basetype.Uint32,
					},
				},
			},
		},
		{
			name:   "decode message data compressed header happy flow",
			r:      fnReaderOK,
			header: proto.MesgCompressedHeaderMask,
			mesgdef: &proto.MessageDefinition{
				Header: proto.MesgDefinitionMask,
				// LocalMesgNum: 0,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordTimestamp,
						Size:     4,
						BaseType: basetype.Uint32,
					},
				},
			},
		},
		{
			name:    "decode message data normal header missing mesg definition",
			r:       fnReaderOK,
			header:  0,
			mesgdef: nil,
			err:     ErrMesgDefMissing,
		},
		{
			name:    "decode message data compressed header missing mesg definition",
			r:       fnReaderOK,
			header:  proto.MesgCompressedHeaderMask,
			mesgdef: nil,
			err:     ErrMesgDefMissing,
		},
		{
			name:   "decode message data decode fields return error",
			r:      fnReaderErr,
			header: proto.MesgCompressedHeaderMask,
			mesgdef: &proto.MessageDefinition{
				Header: proto.MesgDefinitionMask,
				// LocalMesgNum: 0,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordTimestamp,
						Size:     4,
						BaseType: basetype.Uint32,
					},
				},
			},
			err: io.EOF,
		},
		{
			name: "decode message data decode developer fields return error",
			r: func() io.Reader {
				fnIntances, cur := []io.Reader{fnReaderOK, fnReaderErr}, 0
				return fnReader(func(b []byte) (n int, err error) {
					f := fnIntances[cur]
					cur++
					return f.Read(b)
				})
			}(),
			header: proto.MesgNormalHeaderMask,
			mesgdef: &proto.MessageDefinition{
				Header: proto.MesgDefinitionMask | proto.DevDataMask,
				// LocalMesgNum: 0,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordTimestamp,
						Size:     4,
						BaseType: basetype.Uint32,
					},
				},
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{
						Num:                0,
						Size:               4,
						DeveloperDataIndex: 0,
					},
				},
			},
			fieldDescription: &mesgdef.FieldDescription{
				DeveloperDataIndex:    0,
				FieldDefinitionNumber: 0,
			},
			err: io.EOF,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r, WithMesgListener(fnMesgListener(func(mesg proto.Message) {})))
			if tc.mesgdef != nil {
				if (tc.mesgdef.Header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask {
					dec.localMessageDefinitions[(tc.mesgdef.Header&proto.CompressedLocalMesgNumMask)>>proto.CompressedBitShift] = tc.mesgdef
				} else {
					dec.localMessageDefinitions[tc.mesgdef.Header&proto.LocalMesgNumMask] = tc.mesgdef

				}

			}
			if tc.fieldDescription != nil {
				dec.fieldDescriptions = append(dec.fieldDescriptions, tc.fieldDescription)
			}
			err := dec.decodeMessageData(tc.header)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestDecodeFields(t *testing.T) {
	tt := []struct {
		name    string
		r       io.Reader
		opts    []Option
		mesgdef *proto.MessageDefinition
		mesg    proto.Message
		err     error
	}{
		{
			name: "decode fields happy flow",
			r:    fnReaderOK,
			mesgdef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordCadence,
						Size:     1,
						BaseType: basetype.Uint8,
					},
				},
			},
		},
		{
			name: "decode fields unknown field",
			r:    fnReaderOK,
			mesgdef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.FileId,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      68,
						Size:     1,
						BaseType: basetype.Uint8,
					},
				},
			},
		},
		{
			name: "decode fields timestamp not uint32",
			r:    fnReaderOK,
			mesgdef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: 68,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldNumTimestamp,
						Size:     1,
						BaseType: basetype.Uint8,
					},
				},
			},
			err: ErrFieldValueTypeMismatch,
		},
		{
			name: "decode fields accumulate distance",
			r:    fnReaderOK,
			mesgdef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordDistance,
						Size:     4,
						BaseType: basetype.Uint32,
					},
				},
			},
		},
		{
			name: "subfield subtitution",
			r:    fnReaderOK,
			mesgdef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Event,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.EventEvent,
						Size:     1,
						BaseType: basetype.Enum,
					},
					{
						Num:      fieldnum.EventData,
						Size:     4,
						BaseType: basetype.Uint32,
					},
				},
			},
		},
		{
			name: "decode fields without component expansion",
			r:    fnReaderOK,
			opts: []Option{WithNoComponentExpansion()},
			mesgdef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				FieldDefinitions: []proto.FieldDefinition{
					{
						Num:      fieldnum.RecordCadence,
						Size:     1,
						BaseType: basetype.Uint8,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r, tc.opts...)
			err := dec.decodeFields(tc.mesgdef, &tc.mesg)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestExpandComponents(t *testing.T) {
	tt := []struct {
		name                 string
		mesg                 proto.Message
		containingField      proto.Field
		components           []proto.Component
		nFieldAfterExpansion int
	}{
		{
			name: "expand components single happy flow",
			mesg: factory.CreateMesgOnly(mesgnum.Record).WithFields(
				factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
			),
			containingField:      factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
			components:           factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).Components,
			nFieldAfterExpansion: 2, // 1 for speed, +1 expand field enhanced_speed
		},
		{
			name: "expand components multiple happy flow",
			mesg: factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventFrontGearChange)),
			),
			containingField: factory.CreateField(mesgnum.Event, fieldnum.EventData).WithValue(uint32(0x27010E08)),
			components: func() []proto.Component {
				subfields := factory.CreateField(mesgnum.Event, fieldnum.EventData).SubFields
				for _, subfield := range subfields {
					if subfield.Name == "gear_change_data" {
						return subfield.Components
					}
				}
				return nil
			}(),
			nFieldAfterExpansion: 5, // 1 for Event, 4 for expansion fields (rear_gear_num, rear_gear, front_gear_num. front_gear)
		},
		{
			name: "expand components run out bits for the last component",
			mesg: factory.CreateMesgOnly(mesgnum.Event).WithFields(
				factory.CreateField(mesgnum.Event, fieldnum.EventEvent).WithValue(uint8(typedef.EventFrontGearChange)),
			),
			containingField: factory.CreateField(mesgnum.Event, fieldnum.EventData).WithValue(uint32(0x00010E08)),
			components: func() []proto.Component {
				subfields := factory.CreateField(mesgnum.Event, fieldnum.EventData).SubFields
				for _, subfield := range subfields {
					if subfield.Name == "gear_change_data" {
						return subfield.Components
					}
				}
				return nil
			}(),
			nFieldAfterExpansion: 4, // 1 for Event, 3 for expansion fields (rear_gear_num, rear_gear, front_gear_num)
		},
		{
			name: "expand components containing field value mismatch",
			mesg: factory.CreateMesgOnly(mesgnum.Record).WithFields(
				factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue("invalid value"),
			),
			containingField:      factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue("invalid value"),
			components:           factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).Components,
			nFieldAfterExpansion: 1,
		},
		{
			name: "expand components accumulate",
			mesg: factory.CreateMesgOnly(mesgnum.Hr).WithFields(
				factory.CreateField(mesgnum.Hr, fieldnum.HrEventTimestamp).WithValue(uint8(10)),
			),
			containingField:      factory.CreateField(mesgnum.Hr, fieldnum.HrEventTimestamp12).WithValue(uint8(10)),
			components:           factory.CreateField(mesgnum.Hr, fieldnum.HrEventTimestamp12).Components,
			nFieldAfterExpansion: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(nil)
			dec.expandComponents(&tc.mesg, &tc.containingField, tc.components)
			if len(tc.mesg.Fields) != tc.nFieldAfterExpansion {
				t.Fatalf("expected n fields: %d, got: %d", tc.nFieldAfterExpansion, len(tc.mesg.Fields))
			}
		})
	}
}

func TestDecodeDeveloperFields(t *testing.T) {
	tt := []struct {
		name             string
		r                io.Reader
		fieldDescription *mesgdef.FieldDescription
		mesgDef          *proto.MessageDefinition
		mesg             *proto.Message
		err              error
	}{
		{
			name: "decode developer fields happy flow",
			r:    fnReaderOK,
			fieldDescription: mesgdef.NewFieldDescription(
				factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				),
			),
			mesgDef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{
						Num:                0,
						DeveloperDataIndex: 0,
						Size:               1,
					},
				},
			},
			mesg: &proto.Message{},
		},
		{
			name: "decode developer fields missing developer data index 1",
			r:    fnReaderOK,
			fieldDescription: mesgdef.NewFieldDescription(
				factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				),
			),
			mesgDef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{
						Num:                0,
						DeveloperDataIndex: 1,
						Size:               1,
					},
				},
			},
			mesg: &proto.Message{},
		},
		{
			name: "decode developer fields missing field description number",
			r:    fnReaderOK,
			fieldDescription: mesgdef.NewFieldDescription(
				factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				),
			),
			mesgDef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{
						Num:                1,
						DeveloperDataIndex: 0,
						Size:               1,
					},
				},
			},
			mesg: &proto.Message{},
		},
		{
			name: "decode developer fields got io.EOF",
			r:    fnReaderErr,
			fieldDescription: mesgdef.NewFieldDescription(
				factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				),
			),
			mesgDef: &proto.MessageDefinition{
				Header:  proto.MesgDefinitionMask,
				MesgNum: mesgnum.Record,
				DeveloperFieldDefinitions: []proto.DeveloperFieldDefinition{
					{
						Num:                0,
						DeveloperDataIndex: 0,
						Size:               1,
					},
				},
			},
			mesg: &proto.Message{},
			err:  io.EOF,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			dec.fieldDescriptions = append(dec.fieldDescriptions, tc.fieldDescription)
			err := dec.decodeDeveloperFields(tc.mesgDef, tc.mesg)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestReadValue(t *testing.T) {
	tt := []struct {
		name     string
		r        io.Reader
		size     byte
		baseType basetype.BaseType
		arch     byte
		result   any
		err      error
	}{
		{
			name:     "readValue happy flow",
			r:        fnReaderOK, // will produce 0
			size:     1,
			baseType: basetype.Sint8,
			arch:     0,
			result:   int8(0),
		},
		{
			name:     "readValue happy flow",
			r:        fnReaderOK, // will produce 0
			size:     1,
			baseType: basetype.BaseType(100), // invalid basetype.
			arch:     0,
			err:      typedef.ErrTypeNotSupported,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			res, err := dec.readValue(tc.size, tc.baseType, tc.arch)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(res, tc.result); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
