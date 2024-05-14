// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufferedwriter

import (
	"bufio"
	"io"
)

const defaultBufferSize = 4 << 10 // 4096 KB

// BufferedWriter is a writter with buffer, it supports io.Writer, io.WriterAt and io.WriteSeeker.
type BufferedWriter interface {
	io.Writer
	// Flush writes any buffered data to the underlying io.Writer.
	Flush() error
}

// New is shorthand for NewSize(w, 4096), a 4KB Buffer. See NewSize() for details.
func New(w io.Writer) BufferedWriter {
	return NewSize(w, defaultBufferSize)
}

// NewSize creates a buffered writer with a specified size while taking into account the underlying
// capabilities of the writer w, which might implement either [io.WriterAt] or [io.WriteSeeker].
// This allows the buffered writer to maintain the ability to write at a specific byte position.
//
// Use-case scenario:
//   - An *os.File may be passed to a function receiving an [io.Writer], while that function may
//     assert the original ability of the value to write at a specific byte position if possible to
//     enable a faster processing path. Nonetheless, directly working with *os.File for frequent writing
//     small byte segments can affect performance due to numerous syscalls. To alleviate this issue,
//     incorporating a buffered writer for the *os.File becomes essential. Unlike bufio.Writer that
//     encapsulates everything as an [io.Writer], this approach preserves the inherent
//     capabilities of [io.WriterAt] and [io.WriteSeeker].
//
// Just like any other buffered writer, the Flush() method should be called after the process is
// completed to write the unwritten buffered data.
func NewSize(w io.Writer, size int) BufferedWriter {
	bw := bufio.NewWriterSize(w, size)
	switch wr := w.(type) {
	case io.WriterAt:
		return &WriterAt{bw, wr}
	case io.WriteSeeker:
		return &WriteSeeker{bw, wr}
	default:
		return bw
	}
}

// WriterAt is an io.WriterAt buffer wrapper.
type WriterAt struct {
	*bufio.Writer
	wa io.WriterAt
}

func (w *WriterAt) WriteAt(p []byte, offset int64) (n int, err error) {
	// Flush any buffered data from buffered writer in case we are writing on unflushed data.
	if err = w.Writer.Flush(); err != nil {
		return
	}
	return w.wa.WriteAt(p, offset)
}

// WriteSeeker is an io.WriteSeeker buffer wrapper.
type WriteSeeker struct {
	*bufio.Writer
	sk io.Seeker
}

func (w *WriteSeeker) Seek(offset int64, whence int) (n int64, err error) {
	// Flush any buffered data from buffered writer in case we are seeking on unflushed data.
	if err = w.Writer.Flush(); err != nil {
		return
	}
	return w.sk.Seek(offset, whence)
}
