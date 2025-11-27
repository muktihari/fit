// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bufio"
	"io"
)

const defaultWriteBufferSize = 4096

type flusher interface{ Flush() error }

// newWriteBuffer creates new writer that will automatically handle buffering with default buffer size 4096.
// This maintains the capability to write at specific bytes when the underlying writer implements either
// io.WriterAt or io.Seeker. When size <= 0, the original writer will be returned.
func newWriteBuffer(w io.Writer, size int) io.Writer {
	if size <= 0 || w == nil {
		return w
	}
	bw := bufio.NewWriterSize(w, size)
	switch wr := w.(type) {
	case io.WriteSeeker:
		return &writeSeeker{Writer: bw, Seeker: wr}
	case io.WriterAt:
		return &writerAt{Writer: bw, WriterAt: wr}
	default:
		return bw
	}
}

type writerAt struct {
	*bufio.Writer
	io.WriterAt
}

func (w *writerAt) WriteAt(p []byte, offset int64) (n int, err error) {
	// Flush any buffered data from buffered writer in case we are writing on unflushed data.
	if err = w.Writer.Flush(); err != nil {
		return
	}
	return w.WriterAt.WriteAt(p, offset)
}

type writeSeeker struct {
	*bufio.Writer
	io.Seeker
}

func (w *writeSeeker) Seek(offset int64, whence int) (n int64, err error) {
	// Flush any buffered data from buffered writer in case we are seeking on unflushed data.
	if err = w.Writer.Flush(); err != nil {
		return
	}
	return w.Seeker.Seek(offset, whence)
}
