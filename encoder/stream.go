// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"fmt"

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
		e.fileHeader = e.enc.defaultFileHeader
		if err := e.enc.encodeHeader(&e.fileHeader); err != nil {
			return fmt.Errorf("could not encode file header: %w", err)
		}
		e.fileHeaderWritten = true
	}
	if err := e.enc.encodeMessage(e.enc.w, mesg); err != nil {
		return fmt.Errorf("coould not encode mesg: mesgNum: %d (%q): %w", mesg.Num, mesg.Num, err)
	}
	return nil
}

// SequenceCompleted finalises the FIT File by updating its FileHeader's DataSize & CRC, as well as the File's CRC.
// This will also reset variables so that the StreamEncoder can be used for the next sequence of FIT file.
func (e *StreamEncoder) SequenceCompleted() error {
	if err := e.enc.encodeCRC(); err != nil {
		return fmt.Errorf("could not encode crc: %w", err)
	}
	if err := e.enc.updateHeader(&e.fileHeader); err != nil {
		return fmt.Errorf("could not update header: %w", err)
	}
	e.fileHeaderWritten = false
	e.enc.reset()
	return nil
}
