// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type writerAtStub struct {
	b *[]byte
}

var _ io.Writer = &writerAtStub{}

func (w *writerAtStub) Write(p []byte) (n int, err error) {
	*w.b = append(*w.b, p...)
	return len(p), nil
}

var _ io.WriterAt = &writerAtStub{}

func (w *writerAtStub) WriteAt(p []byte, pos int64) (int, error) {
	for i := 0; i < len(p); i++ {
		(*w.b)[pos] = p[i]
		pos++
	}
	return len(p), nil
}

func TestStreamEncoderOneSequenceHappyFlow(t *testing.T) {
	var result []byte
	enc := New(&writerAtStub{&result})
	streamEnc, _ := enc.StreamEncoder()

	fileIdMesg := factory.CreateMesgOnly(mesgnum.FileId).WithFields(
		factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
	)
	recordMesg := factory.CreateMesgOnly(mesgnum.Record).WithFields(
		factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(1000)),
	)

	var err error
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
	expectedEnc.Encode(&proto.Fit{
		Messages: []proto.Message{fileIdMesg, recordMesg},
	})

	if diff := cmp.Diff(result, expected.Bytes()); diff != "" {
		t.Fatal(diff)
	}

	if streamEnc.fileHeaderWritten != false {
		t.Fatalf("expected fileHeaderWritten is %t, got: %t",
			false, streamEnc.fileHeaderWritten)
	}
}

func TestStreamEncoderUnhappyFlow(t *testing.T) {
	// Decode Header Return Error
	enc := New(mockWriterAt{Writer: fnWriteErr})
	streamEnc, _ := enc.StreamEncoder()

	mesg := factory.CreateMesgOnly(mesgnum.FileId).WithFields(
		factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
	)
	err := streamEnc.WriteMessage(&mesg)
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected err: %v, got: %v", io.EOF, err)
	}

	// Valid bytes to checking the correctness
	valid := bytes.NewBuffer(nil)
	validEnc := New(valid)
	validEnc.Encode(&proto.Fit{
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

	enc = New(mockWriterAt{Writer: w, WriterAt: fnWriteAtErr})
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
}
