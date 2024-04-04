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
	"github.com/muktihari/fit/kit"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/listener"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/filedef"
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
	fromGarminForums  = filepath.Join(testdata, "from_garmin_forums")
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

	t.Run("testdata/from_garmin_forums", func(t *testing.T) {
		_ = filepath.Walk(fromGarminForums, func(path string, info fs.FileInfo, _ error) error {
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
				WithMesgListener(mesglis, mesglis),
				WithMesgDefListener(mesgDefLis, mesgDefLis),
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

func TestReinvocationOfExportedMethodsWhenDecoderHasExistingError(t *testing.T) {
	dec := New(nil)
	dec.err = errors.New("intentional error")

	if _, err := dec.CheckIntegrity(); !errors.Is(err, dec.err) {
		t.Fatalf("expected err: %v, got: %v", dec.err, err)
	}
	if _, err := dec.PeekFileId(); !errors.Is(err, dec.err) {
		t.Fatalf("expected err: %v, got: %v", dec.err, err)
	}
	if err := dec.Discard(); !errors.Is(err, dec.err) {
		t.Fatalf("expected err: %v, got: %v", dec.err, err)
	}
	if _, err := dec.Decode(); !errors.Is(err, dec.err) {
		t.Fatalf("expected err: %v, got: %v", dec.err, err)
	}
	if _, err := dec.DecodeWithContext(context.Background()); !errors.Is(err, dec.err) {
		t.Fatalf("expected err: %v, got: %v", dec.err, err)
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
			fileId: mesgdef.NewFileId(&fit.Messages[0]),
		},
		{
			name: "peek file id decode header return error",
			r: func() io.Reader {
				return fnReader(func(b []byte) (n int, err error) {
					return 0, io.EOF
				})
			}(),
			fileId: mesgdef.NewFileId(&fit.Messages[0]),
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
			fileId: mesgdef.NewFileId(&fit.Messages[0]),
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

func TestCheckIntegrity(t *testing.T) {
	_, b := createFitForTest()

	tt := []struct {
		name string
		r    io.Reader
		n    int
		err  error
	}{
		{
			name: "happy flow",
			r: func() io.Reader {
				// Chained FIT File
				b := slices.Clone(b)
				nextb := slices.Clone(b)
				b = append(b, nextb...)
				return bytes.NewReader(b)
			}(),
			n:   2,
			err: nil,
		},
		{
			name: "decode header return io.EOF on first sequence",
			r:    fnReaderErr,
			n:    0,
			err:  io.EOF,
		},
		{
			name: "file header's DataSize == 0",
			r: func() io.Reader {
				h := proto.FileHeader{
					Size:            14,
					ProtocolVersion: byte(proto.V2),
					ProfileVersion:  profile.Version(),
					DataSize:        0,
					DataType:        proto.DataTypeFIT,
				}
				b, _ := h.MarshalBinary()
				crc := crc16.New(crc16.MakeFitTable())
				crc.Write(b[:12])
				binary.LittleEndian.PutUint16(b[12:14], crc.Sum16())
				return bytes.NewReader(b)
			}(),
			n:   0,
			err: ErrDataSizeZero,
		},
		{
			name: "read message return error",
			r: func() io.Reader {
				bb := slices.Clone(b)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 14 {
						return 0, io.EOF
					}
					cur += copy(b, bb[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			n:   0,
			err: io.EOF,
		},
		{
			name: "decode crc return error",
			r: func() io.Reader {
				bb := slices.Clone(b)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(bb)-2 {
						return 0, io.EOF
					}
					cur += copy(b, bb[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			n:   0,
			err: io.EOF,
		},
		{
			name: "crc checksum mismatch",
			r: func() io.Reader {
				bb := slices.Clone(b)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(bb)-2 {
						cur += copy(b, []byte{255, 255}) // crc intentionally altered
						return len(b), nil
					}
					cur += copy(b, bb[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			n:   0,
			err: ErrCRCChecksumMismatch,
		},
		{
			name: "second sequence of FIT File return error",
			r: func() io.Reader {
				// Chained FIT File but with next sequence header is
				b := slices.Clone(b)
				h := headerForTest()
				nextb, _ := h.MarshalBinary()
				nextb[0] = 100 // alter FileHeader's Size
				b = append(b, nextb...)
				return bytes.NewReader(b)
			}(),
			n:   1,
			err: ErrNotAFitFile,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			n, err := dec.CheckIntegrity()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
			if n != tc.n {
				t.Fatalf("expected n sequence of FIT: %d, got: %d", tc.n, n)
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
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue([]string{"Heart Rate"}),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
				factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
			),
			factory.CreateMesgOnly(mesgnum.Record).WithFields(
				factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
			),
			factory.CreateMesgOnly(mesgnum.Record).
				WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(0)),
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
						BaseType:           basetype.Uint8,
						Value:              proto.Uint8(100),
					},
				),
		},
	}

	for i := 0; i < 100; i++ {
		fit.Messages = append(fit.Messages,
			factory.CreateMesgOnly(mesgnum.Record).WithFields(
				factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32((i+1)*1000)),
				factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
			),
		)
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

func TestDiscard(t *testing.T) {
	_, buf := createFitForTest()

	tt := []struct {
		name string
		r    io.Reader
		err  error
	}{
		{
			name: "discard happy flow",
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
			err: nil,
		},
		{
			name: "discard error on decode header",
			r: func() io.Reader {
				return fnReaderErr
			}(),
			err: io.EOF,
		},
		{
			name: "discard error when reading data",
			r: func() io.Reader {
				var buf, cur = slices.Clone(buf), 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur > 14 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			err: io.EOF,
		},
		{
			name: "discard error when reading crc",
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
			err: io.EOF,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			err := dec.Discard()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				if dec.sequenceCompleted {
					t.Fatalf("sequenceCompleted should be false")
				}
				return
			}
			if !dec.sequenceCompleted {
				t.Fatalf("sequenceCompleted should be true")
			}
		})
	}
}

func TestDiscardChained(t *testing.T) {
	activityFile, err := os.ReadFile(filepath.Join(fromOfficialSDK, "Activity.fit"))
	if err != nil {
		t.Fatal(err)
	}
	monitoringFile, err := os.ReadFile(filepath.Join(fromOfficialSDK, "MonitoringFile.fit"))
	if err != nil {
		t.Fatal(err)
	}

	// make chained files
	b := make([]byte, 0, len(activityFile)+len(monitoringFile)+len(activityFile))
	b = append(b, activityFile...)
	b = append(b, monitoringFile...)
	b = append(b, activityFile...)

	r := bytes.NewBuffer(b)

	fits := make([]*proto.Fit, 0, 2)
	dec := New(r)
	for dec.Next() {
		fileId, err := dec.PeekFileId()
		if err != nil {
			t.Fatal(err)
		}

		if fileId.Type != typedef.FileActivity {
			if err := dec.Discard(); err != nil {
				t.Fatal(err)
			}
			continue
		}

		fit, err := dec.Decode()
		if err != nil {
			t.Fatal(err)
		}
		fits = append(fits, fit)
	}

	if len(fits) != 2 {
		t.Fatalf("expected activities is 2, got: %d", len(fits))
	}
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

	if !dec.Next() {
		t.Fatalf("should have next, return false")
	}

	_, err := dec.DecodeWithContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Check whether after decode, fields are reset and next sequence is retrieved.

	if !dec.Next() {
		t.Fatalf("should have next, return false")
	}

	if len(dec.accumulator.AccumulatedValues) != 0 {
		t.Fatalf("expected accumulator's AccumulatedValues is 0, got: %d", len(dec.accumulator.AccumulatedValues))
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

	if _, err := dec.PeekFileId(); err != io.EOF {
		t.Fatalf("expected EOF got %v", err)
	}

	if dec.Next() {
		t.Fatalf("should be false, got true")
	}
}

type decodeTestCase struct {
	name string
	r    io.Reader
	fit  proto.Fit
	err  error
}

func makeDecodeTableTest() []decodeTestCase {
	fit, buf := createFitForTest()
	return []decodeTestCase{
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
			name: "decode header return error",
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
}

func TestDecode(t *testing.T) {
	tt := makeDecodeTableTest()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			fit, err := dec.Decode()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(*fit, tc.fit,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
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
		opts             []Option
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
			name:   "decode message data normal header happy flow without component expansions",
			r:      fnReaderOK,
			opts:   []Option{WithNoComponentExpansion()},
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
			opts := append(tc.opts, WithMesgListener(fnMesgListener(func(mesg proto.Message) {})))
			dec := New(tc.r, opts...)
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
						Num:      proto.FieldNumTimestamp,
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

func TestExpandMutipleComponents(t *testing.T) {
	// Expand componentField.Components that require expansion
	compressedSepeedDistanceField := factory.CreateField(mesgnum.Record, fieldnum.RecordCompressedSpeedDistance).
		WithValue([]byte{0, 4, 1})

	mesg := factory.CreateMesgOnly(mesgnum.Record).WithFields(compressedSepeedDistanceField)
	dec := New(nil)
	dec.expandComponents(&mesg, &compressedSepeedDistanceField, compressedSepeedDistanceField.Components)

	if len(mesg.Fields) != 4 {
		t.Errorf("expected n fields after expansion: %d, got: %d", 4, len(mesg.Fields))
	}

	if diff := cmp.Diff(
		mesg.FieldValueByNum(fieldnum.RecordCompressedSpeedDistance).Any(),
		[]byte{0, 4, 1},
	); diff != "" {
		t.Errorf("compressed_speed_distance: %s", diff)
	}

	// Formula: value = (value / component_speed_scale) * destination_field_scale

	if diff := cmp.Diff(
		mesg.FieldValueByNum(fieldnum.RecordSpeed).Any(),
		uint16(10240), // (1024 / 100) * 1000
	); diff != "" {
		t.Errorf("speed: %s", diff)
	}

	if diff := cmp.Diff(
		mesg.FieldValueByNum(fieldnum.RecordDistance).Any(),
		uint32(100), // (1600 / 16) * 1
	); diff != "" {
		t.Errorf("distance: %s", diff)
	}

	if diff := cmp.Diff(
		mesg.FieldValueByNum(fieldnum.RecordEnhancedSpeed).Any(),
		uint32(10240), // (1024 / 1000) * 1000
	); diff != "" {
		t.Errorf("enhanced_speed: %s", diff)
	}
}

func TestExpandMutipleComponentsDynamicField(t *testing.T) {
	// Test expand component's components that have dynamic field
	// Since we don't have real world example, message from Profile.xlsx doesn't not have this but it is possible,
	// Let's create our own custom message.

	const customMesgNum = 65280
	fac := factory.New()
	fac.RegisterMesg(
		proto.Message{
			Num: customMesgNum,
			Fields: []proto.Field{
				{
					FieldBase: &proto.FieldBase{
						Num:        0,
						Name:       "event",
						Type:       profile.Event, /* basetype.Enum (size: 1) */
						BaseType:   profile.Event.BaseType(),
						Components: nil,
						Scale:      1, Offset: 0,
					},
					Value: proto.Uint8(basetype.EnumInvalid),
				},
				{
					FieldBase: &proto.FieldBase{
						Num:      1,
						Name:     "data",
						Type:     profile.Uint32,
						BaseType: basetype.Uint32,
						Scale:    1, Offset: 0,
						SubFields: []proto.SubField{
							{Name: "timer_trigger", Type: profile.TimerTrigger /* basetype.Enum */, Scale: 1, Offset: 0,
								Components: nil,
								Maps: []proto.SubFieldMap{
									{RefFieldNum: 0 /* event */, RefFieldValue: 0 /* timer */},
								},
							},
							{Name: "course_point_index", Type: profile.MessageIndex /* basetype.Uint16 */, Scale: 1, Offset: 0,
								Components: nil,
								Maps: []proto.SubFieldMap{
									{RefFieldNum: 0 /* event */, RefFieldValue: 10 /* course_point */},
								},
							},
						},
					},
				},
				{
					FieldBase: &proto.FieldBase{
						Num:      2,
						Name:     "compressed_data",
						Type:     profile.Uint32,
						BaseType: profile.Uint32.BaseType(),
						Components: []proto.Component{
							{FieldNum: 1 /* data */, Scale: 1, Offset: 0, Accumulate: false, Bits: 8},
						},
						Scale: 1, Offset: 0,
					},
					Value: proto.Uint32(basetype.Uint32Invalid),
				},
			},
		},
	)

	mesg := fac.CreateMesgOnly(customMesgNum).WithFields(
		fac.CreateField(customMesgNum, 0).WithValue(uint8(10)),  // event
		fac.CreateField(customMesgNum, 2).WithValue(uint32(10)), // compressed_data
	)

	dec := New(nil, WithFactory(fac))
	fieldToExpand := mesg.FieldByNum(2)
	dec.expandComponents(&mesg, fieldToExpand, fieldToExpand.Components)

	if len(mesg.Fields) != 3 {
		t.Errorf("expected n fields: %d, got %d", 3, len(mesg.Fields))
	}

	if diff := cmp.Diff(
		mesg.FieldValueByNum(1).Any(),
		uint32(10),
	); diff != "" {
		t.Errorf("data: %s", diff)
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
				kit.Ptr(factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				)),
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
				kit.Ptr(factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				)),
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
				kit.Ptr(factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				)),
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
			name: "decode developer fields missing field description number but unable to read acquired bytes",
			r:    fnReaderErr,
			fieldDescription: mesgdef.NewFieldDescription(
				kit.Ptr(factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				)),
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
			err:  io.EOF,
		},
		{
			name: "decode developer fields got io.EOF",
			r:    fnReaderErr,
			fieldDescription: mesgdef.NewFieldDescription(
				kit.Ptr(factory.CreateMesgOnly(mesgnum.FieldDescription).WithFields(
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionDeveloperDataIndex).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldDefinitionNumber).WithValue(uint8(0)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFieldName).WithValue("Heart Rate"),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeMesgNum).WithValue(uint16(mesgnum.Record)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionNativeFieldNum).WithValue(uint8(fieldnum.RecordHeartRate)),
					factory.CreateField(mesgnum.FieldDescription, fieldnum.FieldDescriptionFitBaseTypeId).WithValue(uint8(basetype.Uint8)),
				)),
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
		result   proto.Value
		err      error
	}{
		{
			name:     "readValue happy flow",
			r:        fnReaderOK, // will produce 0
			size:     1,
			baseType: basetype.Sint8,
			arch:     0,
			result:   proto.Int8(0),
		},
		{
			name:     "readValue happy flow",
			r:        fnReaderOK, // will produce 0
			size:     1,
			baseType: basetype.BaseType(100), // invalid basetype.
			arch:     0,
			err:      proto.ErrTypeNotSupported,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			res, err := dec.readValue(tc.size, tc.baseType, false, tc.arch)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(res, tc.result,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestDecodeWithContext(t *testing.T) {
	tt := makeDecodeTableTest()
	var ctx context.Context

	// Testing logic same as Decode()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			fit, err := dec.DecodeWithContext(ctx)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(*fit, tc.fit,
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
				}),
			); diff != "" {
				t.Fatal(diff)
			}
		})
	}

	type strct struct {
		name string
		ctx  context.Context
		r    io.Reader
		err  error
	}

	// Test context related logic
	tt2 := []strct{
		{
			name: "context canceled before decode header",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			err: context.Canceled,
		},
		func() strct {
			ctx, cancel := context.WithCancel(context.Background())
			_, buffer := createFitForTest()
			buf, cur := slices.Clone(buffer), 0
			r := fnReader(func(b []byte) (n int, err error) {
				cur += copy(b, buf[cur:cur+len(b)])
				if cur == len(buf)-2 {
					cancel() // cancel right after completing decode messages
				}
				return len(b), nil
			})

			return strct{
				name: "context canceled before decode crc",
				r:    r,
				ctx:  ctx,
				err:  context.Canceled,
			}
		}(),
	}

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			dec := New(tc.r)
			_, err := dec.DecodeWithContext(tc.ctx)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestDecodeMessagesWithContext(t *testing.T) {
	tt := []struct {
		name string
		r    io.Reader
		ctx  context.Context
		err  error
	}{
		{
			name: "context DeadlineExceeded",
			r: fnReader(func(b []byte) (n int, err error) {
				time.Sleep(1 * time.Second) // Let's make our process take longer, 1s per reader Read call.
				return len(b), nil
			}),
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
			dec := New(tc.r)
			dec.fileHeader.DataSize = 1024
			err := dec.decodeMessagesWithContext(tc.ctx)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}

func TestReset(t *testing.T) {
	// predefined
	decoderFactory := factory.New()
	mesglis := fnMesgListener(func(mesg proto.Message) {})
	mesgDefLis := fnMesgDefListener(func(mesgDef proto.MessageDefinition) {})

	tt := []struct {
		name string
		opts []Option
		dec  *Decoder
	}{
		{
			name: "reset with options",
			opts: []Option{
				WithFactory(decoderFactory),
				WithIgnoreChecksum(),
				WithMesgListener(mesglis, mesglis),
				WithMesgDefListener(mesgDefLis, mesgDefLis),
				WithBroadcastOnly(),
				WithNoComponentExpansion(),
			},
			dec: New(nil,
				WithFactory(decoderFactory),
				WithIgnoreChecksum(),
				WithMesgListener(mesglis, mesglis),
				WithMesgDefListener(mesgDefLis, mesgDefLis),
				WithBroadcastOnly(),
				WithNoComponentExpansion(),
			),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			dec := New(buf, tc.opts...)

			dec.Reset(buf, tc.opts...)

			if diff := cmp.Diff(dec, tc.dec,
				cmp.AllowUnexported(options{}),
				cmp.AllowUnexported(Decoder{}),
				cmp.FilterValues(func(x, y io.Reader) bool { return true }, cmp.Ignore()),
				cmp.FilterValues(func(x, y hash.Hash16) bool { return true }, cmp.Ignore()),
				cmp.FilterValues(func(x, y func() error) bool { return true }, cmp.Ignore()),
				cmp.Transformer("factory", func(t Factory) uintptr {
					return reflect.ValueOf(t).Pointer()
				}),
				cmp.Transformer("Value", func(v proto.Value) any {
					return v.Any()
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

func BenchmarkDecodeMessageData(b *testing.B) {
	b.StopTimer()
	mesg := factory.CreateMesg(typedef.MesgNumRecord)
	mesgDef := proto.CreateMessageDefinition(&mesg)
	mesgb, err := mesg.MarshalBinary()
	if err != nil {
		b.Fatalf("marshal binary: %v", err)
	}
	buf := bytes.NewBuffer(mesgb)
	dec := New(buf, WithIgnoreChecksum(), WithNoComponentExpansion(), WithBroadcastOnly())
	dec.localMessageDefinitions[0] = &mesgDef
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(mesgb)
		err := dec.decodeMessageData(0)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	b.StopTimer()
	// This is not a typical FIT in term of file size (2.3M) and the messages it contains (200.000 messages)
	// But since it's big, it's should be good to benchmark.
	f, err := os.Open("../testdata/big_activity.fit")
	// f, err := os.Open("../testdata/from_official_sdk/activity_lowbattery.fit")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	all, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(all)
	dec := New(buf)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(all)
		_, err = dec.Decode()
		if err != nil {
			b.Fatal(err)
		}
		dec.reset()
	}
}

func BenchmarkDecodeWithFiledef(b *testing.B) {
	b.StopTimer()
	// This is not a typical FIT in term of file size (2.3M) and the messages it contains (200.000 messages)
	// But since it's big, it's should be good to benchmark.
	f, err := os.Open("../testdata/big_activity.fit")
	// f, err := os.Open("../testdata/from_official_sdk/activity_lowbattery.fit")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	all, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	al := filedef.NewListener()
	buf := bytes.NewBuffer(all)
	dec := New(buf, WithMesgListener(al), WithBroadcastOnly())
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(all)
		_, err = dec.Decode()
		if err != nil {
			b.Fatal(err)
		}
		_ = al.File()
		dec.reset()
	}
}

func BenchmarkCheckIntegrity(b *testing.B) {
	b.StopTimer()
	// This is not a typical FIT in term of file size (2.3M) and the messages it contains (200.000 messages)
	// But since it's big, it's should be good to benchmark.
	f, err := os.Open("../testdata/big_activity.fit")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	all, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	b.StartTimer()

	buf := bytes.NewBuffer(all)
	dec := New(buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(all)
		_, err = dec.CheckIntegrity()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReset(b *testing.B) {
	b.Run("benchmark New()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = New(nil)
		}
	})
	b.Run("benchmark Reset()", func(b *testing.B) {
		b.StopTimer()
		dec := New(nil)
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			dec.Reset(nil)
		}
	})
}
