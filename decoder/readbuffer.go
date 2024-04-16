// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"io"
	"math"
)

const (
	minReadBufferSize     = 16
	maxReadBufferSize     = math.MaxUint32
	defaultReadBufferSize = 4096

	// reservedbuf is the maximum bytes that will be requested by the Decoder in one read.
	// The value is obtained from the maximum n field definition in a mesg is 255 and
	// we need 3 byte per field. So 255 * 3 = 765.
	reservedbuf = 765
)

// readBuffer is a custom buffered reader. See newReadBuffer() for details.
type readBuffer struct {
	r io.Reader // reader provided by the client

	// buf is a bytes buffer to read from io.Reader.
	// This has unique memory layout:
	// [reserved section] + [resizable section]
	// [0, 1,..., 765,    + 766, 767,..., size]
	//
	// reserved section is used to memmove remaining bytes when remaining < n.
	// resizable section is the space for reading from io.Reader.
	//
	// This way, fragmented remaining bytes is handled and it will always try to
	// read exactly x size bytes from io.Reader.
	buf []byte

	cur, last int // cur and last of buf positions
}

// newReadBuffer creates a new reader that will automatically handle buffering,
// allowing us to read bytes directly from the buffer without extra copying.
// This is unlike *bufio.Reader, which requires us to copy the bytes on every Read() method call.
func newReadBuffer(rd io.Reader, size int) *readBuffer {
	if size < minReadBufferSize {
		size = minReadBufferSize
	} else if size > maxReadBufferSize {
		size = maxReadBufferSize
	}
	r := new(readBuffer)
	r.Reset(rd, size)
	return r
}

// ReadN reads bytes from the buffer and return exactly n bytes.
// If the remaining bytes in the buffer is less than n bytes requested, it will automatically fill the buffer.
// And if in the process it got less than n, an error will be returned.
func (b *readBuffer) ReadN(n int) ([]byte, error) {
	remaining := b.last - b.cur
	if remaining == 0 {
		b.cur = reservedbuf
		b.last = reservedbuf
	} else if n > remaining {
		b.cur = reservedbuf - remaining               // cursor is now pointing at index on 'reserved section'
		copy(b.buf[b.cur:], b.buf[b.last-remaining:]) // memmove remaining bytes to 'reserved section'.
	}

	if n > remaining { // fill buf
		nr, err := io.ReadAtLeast(b.r, b.buf[reservedbuf:], n-remaining)
		if err != nil {
			return nil, err
		}
		b.last = reservedbuf + nr
	}

	buf := b.buf[b.cur : b.cur+n]
	b.cur += n
	return buf, nil
}

// Reset resets buf reader.
func (b *readBuffer) Reset(rd io.Reader, size int) {
	oldsize := cap(b.buf) - reservedbuf
	if size != oldsize {
		b.buf = make([]byte, reservedbuf+size)
	}
	b.buf = b.buf[:reservedbuf+size]

	b.r = rd
	b.cur = reservedbuf
	b.last = reservedbuf
}

// Size return the len of resizeable section of buf.
func (b *readBuffer) Size() int { return len(b.buf) - reservedbuf }
