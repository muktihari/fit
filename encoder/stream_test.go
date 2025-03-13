// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bytes"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestStreamEncoderOneSequenceHappyFlow(t *testing.T) {
	w := &writerAtStub{}
	enc := New(w)
	streamEnc, err := enc.StreamEncoder()
	if err != nil {
		t.Fatal(err)
	}

	fileIdMesg := proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
		factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
	}}
	recordMesg := proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
		factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1000)),
	}}

	err = streamEnc.WriteMessage(&fileIdMesg)
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
		return
	}
	err = streamEnc.WriteMessage(&recordMesg)
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
		return
	}

	err = streamEnc.SequenceCompleted()
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
		return
	}

	expected := bytes.NewBuffer(nil)
	expectedEnc := New(expected)
	expectedEnc.Encode(&proto.FIT{
		Messages: []proto.Message{fileIdMesg, recordMesg},
	})

	if diff := cmp.Diff(w.buf, expected.Bytes()); diff != "" {
		t.Fatal(diff)
	}

	if streamEnc.fileHeaderWritten != false {
		t.Fatalf("expected fileHeaderWritten is %t, got: %t",
			false, streamEnc.fileHeaderWritten)
	}
}

func TestStreamEncoderUnhappyFlow(t *testing.T) {
	// Decode Header Return Error
	enc := New(mockWriterAt{Writer: fnWriteErr}, WithWriteBufferSize(0))
	streamEnc, _ := enc.StreamEncoder()

	mesg := proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
		factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
	}}
	err := streamEnc.WriteMessage(&mesg)
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected err: %v, got: %v", io.EOF, err)
	}

	// Valid bytes to checking the correctness
	valid := bytes.NewBuffer(nil)
	validEnc := New(valid)
	validEnc.Encode(&proto.FIT{
		Messages: []proto.Message{mesg},
	})

	// SequenceCompleted error on encode crc
	writeCounter := 0
	w := fnWriter(func(b []byte) (n int, err error) {
		if writeCounter == valid.Len()-2 { // minus crc
			return 0, io.ErrShortWrite
		}
		writeCounter += len(b)
		return len(b), nil
	})

	enc = New(mockWriterAt{Writer: w, WriterAt: fnWriteAtErr},
		WithWriteBufferSize(0),
	)
	streamEnc, _ = enc.StreamEncoder()
	err = streamEnc.WriteMessage(&mesg)
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
	}

	err = streamEnc.SequenceCompleted()
	if !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("expected err: %v, got: %v", io.ErrShortWrite, err)
	}

	// SequenceCompleted error on update header
	writeCounter = 0

	wa := fnWriterAt(func(p []byte, offset int64) (n int, err error) {
		return 0, io.ErrUnexpectedEOF
	})
	enc = New(mockWriterAt{Writer: fnWriteOK, WriterAt: wa})
	streamEnc, _ = enc.StreamEncoder()
	err = streamEnc.WriteMessage(&mesg)
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
	}

	err = streamEnc.SequenceCompleted()
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("expected err: %v, got: %v", io.ErrShortWrite, err)
	}

	// Encode message return error
	ws := []io.Writer{fnWriteOK, fnWriteErr}
	cur := 0
	w = fnWriter(func(b []byte) (n int, err error) {
		cur++
		return ws[cur-1].Write(b)
	})
	enc = New(mockWriterAt{Writer: w, WriterAt: wa},
		WithWriteBufferSize(0),
	)
	streamEnc, _ = enc.StreamEncoder()

	err = streamEnc.WriteMessage(&mesg)
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected err: %v, got: %v", io.EOF, err)
	}

	// Protocol validation error
	streamEnc, _ = New(mockWriterAt{}).StreamEncoder()
	streamEnc.enc.protocolValidator.ProtocolVersion = proto.V1
	err = streamEnc.WriteMessage(&proto.Message{Fields: []proto.Field{
		factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed1S).WithValue(make([]uint8, 256)),
	}, DeveloperFields: []proto.DeveloperField{{}}})
	if !errors.Is(err, proto.ErrProtocolViolation) {
		t.Fatalf("expected err: %v, got: %v", proto.ErrProtocolViolation, err)
	}

	// Message validation error
	streamEnc, _ = New(mockWriterAt{}).StreamEncoder()
	err = streamEnc.WriteMessage(&proto.Message{Fields: []proto.Field{
		factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed1S).WithValue(make([]uint8, 256))}},
	)
	if !errors.Is(err, errExceedMaxAllowed) {
		t.Fatalf("expected err: %v, got: %v", errExceedMaxAllowed, err)
	}
}

func TestStreamEncoderWithoutWriteBuffer(t *testing.T) {
	w := &writerAtStub{}
	enc := New(w, WithWriteBufferSize(0))
	streamEnc, err := enc.StreamEncoder()
	if err != nil {
		t.Fatal(err)
	}

	fileIdMesg := proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
		factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
	}}

	err = streamEnc.WriteMessage(&fileIdMesg)
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
		return
	}
	err = streamEnc.SequenceCompleted()
	if err != nil {
		t.Fatalf("expected err: %v, got: %v", nil, err)
		return
	}
}

func TestStreamEncoderReset(t *testing.T) {
	tt := []struct {
		name string
		w1   io.Writer
		w2   io.Writer
		err  error
	}{
		{
			name: "io.WriteSeeker reset with io.WriteSeeker",
			w1:   mockWriteSeeker{fnWriteOK, fnSeekOK},
			w2:   mockWriteSeeker{fnWriteOK, fnSeekOK},
			err:  nil,
		},
		{
			name: "io.WriteSeeker reset with io.Writer",
			w1:   mockWriteSeeker{fnWriteOK, fnSeekOK},
			w2:   fnWriteOK,
			err:  ErrInvalidWriter,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			streamEnc, _ := New(tc.w1).StreamEncoder()
			err := streamEnc.Reset(tc.w2)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
		})
	}
}
