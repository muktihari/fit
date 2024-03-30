// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

var fnDecodeRawOK = func(flag RawFlag, b []byte) error {
	return nil
}

var fnDecodeRawErr = func(flag RawFlag, b []byte) error {
	return io.EOF
}

func TestRawDecoderDecode(t *testing.T) {
	_, buf := createFitForTest()
	hash16 := crc16.New(crc16.MakeFitTable())

	tt := []struct {
		name string
		fn   func(flag RawFlag, b []byte) error
		r    io.Reader
		b    []byte
		err  error
	}{
		{
			name: "happy flow",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn: func(flag RawFlag, b []byte) error {
				// Test CRC Checksum
				switch flag {
				case RawFlagFileHeader:
					if b[0] == 14 {
						hash16.Write(b[:12])
						if binary.LittleEndian.Uint16(b[12:14]) != hash16.Sum16() {
							return ErrCRCChecksumMismatch
						}
						hash16.Reset()
					}
				case RawFlagMesgDef, RawFlagMesgData:
					hash16.Write(b)
				case RawFlagCRC:
					if binary.LittleEndian.Uint16(b[:2]) != hash16.Sum16() {
						return ErrCRCChecksumMismatch
					}
					hash16.Reset()
				}
				return nil
			},
			b: func() []byte {
				return buf
			}(),
			err: nil,
		},
		{
			name: "decode header 1st sequence return io.EOF",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 0 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn: func(flag RawFlag, b []byte) error {
				if flag == RawFlagCRC {
					return io.EOF
				}
				return nil
			},
			err: io.EOF,
		},
		{
			name: "invalid FileHeader's Size",
			r: func() io.Reader {
				buf := slices.Clone(buf)
				buf[0] = 100
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: ErrNotAFitFile,
		},
		{
			name: "unexpected EOF when decode header",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 1 {
						return 0, io.ErrUnexpectedEOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.ErrUnexpectedEOF,
		},
		{
			name: "bytes 8-12 is not .FIT",
			r: func() io.Reader {
				buf := slices.Clone(buf)
				copy(buf[8:12], []byte(".FTT"))
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: ErrNotAFitFile,
		},
		{
			name: "fn FileHeader returns io.EOF",
			r: func() io.Reader {
				buf := slices.Clone(buf)
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawErr,
			err: io.EOF,
		},
		{
			name: "decode mesgDef header return io.EOF",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 14 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode mesgDef bytes 1-5 return io.EOF",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == 15 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode mesgDef fields return io.EOF",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgDef := proto.CreateMessageDefinition(&mesg)
				mesgDefb, _ := mesgDef.MarshalBinary()
				buf = append(buf, mesgDefb...)

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf)-3 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode mesgDef n developer fields return io.EOF",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				).WithDeveloperFields(
					proto.DeveloperField{
						DeveloperDataIndex: 0,
						Num:                0,
						Size:               1,
						Name:               "Heart Rate",
						NativeMesgNum:      mesgnum.Record,
						NativeFieldNum:     fieldnum.RecordHeartRate,
						BaseType:           basetype.Uint8,
						Value:              uint8(100),
					},
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgDef := proto.CreateMessageDefinition(&mesg)
				mesgDefb, _ := mesgDef.MarshalBinary()
				buf = append(buf, mesgDefb...)

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf)-4 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode mesgDef developer fields return io.EOF",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				).WithDeveloperFields(
					proto.DeveloperField{
						DeveloperDataIndex: 0,
						Num:                0,
						Size:               1,
						Name:               "Heart Rate",
						NativeMesgNum:      mesgnum.Record,
						NativeFieldNum:     fieldnum.RecordHeartRate,
						BaseType:           basetype.Uint8,
						Value:              uint8(100),
					},
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgDef := proto.CreateMessageDefinition(&mesg)
				mesgDefb, _ := mesgDef.MarshalBinary()
				buf = append(buf, mesgDefb...)

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf)-3 {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode mesgDef fn return io.EOF",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				).WithDeveloperFields(
					proto.DeveloperField{
						DeveloperDataIndex: 0,
						Num:                0,
						Size:               1,
						Name:               "Heart Rate",
						NativeMesgNum:      mesgnum.Record,
						NativeFieldNum:     fieldnum.RecordHeartRate,
						BaseType:           basetype.Uint8,
						Value:              uint8(100),
					},
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgDef := proto.CreateMessageDefinition(&mesg)
				mesgDefb, _ := mesgDef.MarshalBinary()
				buf = append(buf, mesgDefb...)

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn: func(flag RawFlag, b []byte) error {
				if flag == RawFlagMesgDef {
					return io.EOF
				}
				return nil
			},
			err: io.EOF,
		},
		{
			name: "decode mesg, mesgDef not found",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgb, _ := mesg.MarshalBinary()
				buf = append(buf, mesgb...)

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: ErrMesgDefMissing,
		},
		{
			name: "decode mesg, read return io.EOF",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgDef := proto.CreateMessageDefinition(&mesg)
				mesgDefb, _ := mesgDef.MarshalBinary()
				buf = append(buf, mesgDefb...)
				buf = append(buf, mesgDefb[0]&proto.LocalMesgNumMask)

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode mesg fn return io.EOF",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn: func(flag RawFlag, b []byte) error {
				if flag == RawFlagMesgData {
					return io.EOF
				}
				return nil
			},
			err: io.EOF,
		},
		{
			name: "decode crc return io.EOF",
			r: func() io.Reader {
				mesg := factory.CreateMesg(mesgnum.Record).WithFields(
					factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(time.Now())),
				)
				h := headerForTest()
				buf, _ := h.MarshalBinary()
				mesgDef := proto.CreateMessageDefinition(&mesg)
				mesgDefb, _ := mesgDef.MarshalBinary()
				buf = append(buf, mesgDefb...)
				mesgb, _ := mesg.MarshalBinary()
				buf = append(buf, mesgb...)

				binary.LittleEndian.PutUint32(buf[4:8], uint32(len(buf)-14))

				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn:  fnDecodeRawOK,
			err: io.EOF,
		},
		{
			name: "decode crc return io.EOF",
			r: func() io.Reader {
				cur := 0
				return fnReader(func(b []byte) (n int, err error) {
					if cur == len(buf) {
						return 0, io.EOF
					}
					cur += copy(b, buf[cur:cur+len(b)])
					return len(b), nil
				})
			}(),
			fn: func(flag RawFlag, b []byte) error {
				if flag == RawFlagCRC {
					return io.EOF
				}
				return nil
			},
			err: io.EOF,
		},
	}

	dec := NewRaw()
	result := new(bytes.Buffer)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			n, err := dec.Decode(tc.r, func(flag RawFlag, b []byte) error {
				result.Write(b) // Inject to test correctness of bytes
				return tc.fn(flag, b)
			})

			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}

			if err != nil {
				return
			}

			if n != int64(result.Len()) {
				t.Fatalf("expecteed n: %d, got: %d", result.Len(), n)
			}

			if diff := cmp.Diff(tc.b, result.Bytes()); diff != "" {
				t.Fatal(diff)
			}

			result.Reset()
			dec.reset()
			hash16.Reset()
		})
	}
}

func BenchmarkRawDecoderDecode(b *testing.B) {
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

	buf := bytes.NewBuffer(all)
	dec := NewRaw()
	fmt.Println(unsafe.Sizeof(*dec))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(all)
		_, err = dec.Decode(buf, func(flag RawFlag, b []byte) error { return nil })
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestRawFlagString(t *testing.T) {
	tt := []struct {
		str string
		f   RawFlag
	}{
		{
			str: "file_header",
			f:   RawFlagFileHeader,
		},
		{
			str: "message_definition",
			f:   RawFlagMesgDef,
		},
		{
			str: "message_data",
			f:   RawFlagMesgData,
		},
		{
			str: "crc",
			f:   RawFlagCRC,
		},
		{
			str: "unknown(255)",
			f:   255,
		},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			if diff := cmp.Diff(tc.f.String(), tc.str); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
