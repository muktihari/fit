// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"fmt"
	"io"

	"github.com/muktihari/fit/proto"
)

// StreamEncoder is one layer above Encoder to enable encoding in streaming fashion.
// This will only valid when the Writer given to the Encoder should either implement io.WriterAt or io.WriteSeeker.
// This can only be created using (*Encoder).StreamEncoder() method.
type StreamEncoder struct {
	enc               *Encoder
	fileHeader        proto.FileHeader
	fileHeaderWritten bool
}

// WriteMessage writes message to the writer, it will auto write FileHeader when
//   - This method is invoked on the first time of use.
//   - This method is called right after SequenceCompleted method has been called.
func (e *StreamEncoder) WriteMessage(mesg *proto.Message) error {
	if !e.fileHeaderWritten {
		if err := e.enc.encodeFileHeader(&e.fileHeader); err != nil {
			return fmt.Errorf("could not encode file header: %w", err)
		}
		e.fileHeaderWritten = true
	}
	if err := e.enc.options.messageValidator.Validate(mesg); err != nil {
		return fmt.Errorf("message validation failed: mesgNum: %d (%s): %w", mesg.Num, mesg.Num, err)
	}
	if err := e.enc.encodeMessage(mesg); err != nil {
		return fmt.Errorf("could not encode mesg: mesgNum: %d (%q): %w", mesg.Num, mesg.Num, err)
	}
	return nil
}

// SequenceCompleted finalises the FIT File by updating its FileHeader's DataSize & CRC, as well as the File's CRC.
// This will also reset variables so that the StreamEncoder can be used for the next sequence of FIT file.
func (e *StreamEncoder) SequenceCompleted() error {
	if err := e.enc.encodeCRC(); err != nil {
		return fmt.Errorf("could not encode crc: %w", err)
	}
	if err := e.enc.updateFileHeader(&e.fileHeader); err != nil {
		return fmt.Errorf("could not update header: %w", err)
	}
	e.fileHeaderWritten = false
	e.enc.reset()
	if f, ok := e.enc.w.(flusher); ok {
		return f.Flush()
	}
	return nil
}

// Reset resets the Stream Encoder and the underlying Encoder to write its output to w
// and reset previous options to default options so any options needs to be inputed again.
// It is similar to New() but it retains the underlying storage for use by future encode to reduce memory allocs.
// If w does not implement io.WriterAt or io.WriteSeeker, error will be returned.
func (e *StreamEncoder) Reset(w io.Writer, opts ...Option) error {
	switch w.(type) {
	case io.WriterAt, io.WriteSeeker:
		e.enc.Reset(w, opts...)
		e.fileHeader = proto.FileHeader{}
		e.fileHeaderWritten = false
		return nil
	}
	return fmt.Errorf("could not reset: %w", ErrWriterAtOrWriteSeekerIsExpected)
}
