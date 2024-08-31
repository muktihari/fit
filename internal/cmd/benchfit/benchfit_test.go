// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package benchfit_test

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/tormoder/fit"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	cd                = filepath.Dir(filename)
	testdata          = filepath.Join(cd, "..", "..", "..", "testdata")
	big_activity      = filepath.Join(testdata, "big_activity.fit")
)

func BenchmarkDecode(b *testing.B) {
	b.Run("muktihari/fit raw", func(b *testing.B) {
		b.StopTimer()
		// NOTE: We use os.ReadFile to remove syscall process in our decoding process. So we have pure decoding performance.
		f, err := os.ReadFile(big_activity)
		if err != nil {
			b.Fatalf("open file: %v", err)
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			dec := decoder.New(bytes.NewReader(f))
			_, err = dec.Decode()
			if err != nil {
				b.Fatalf("decode error: %v", err)
			}
		}
	})
	b.Run("muktihari/fit", func(b *testing.B) {
		b.StopTimer()
		f, err := os.ReadFile(big_activity)
		if err != nil {
			b.Fatalf("open file: %v", err)
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			al := filedef.NewListener()
			dec := decoder.New(bytes.NewReader(f),
				decoder.WithMesgListener(al),
				decoder.WithBroadcastOnly(),
			)
			_, err = dec.Decode()
			if err != nil {
				b.Fatalf("decode error: %v", err)
			}
			_ = al.File()
			al.Close()
		}
	})
	b.Run("tormoder/fit", func(b *testing.B) {
		b.StopTimer()
		f, err := os.ReadFile(big_activity)
		if err != nil {
			b.Fatalf("open file: %v", err)
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			_, err = fit.Decode(bytes.NewReader(f))
			if err != nil {
				b.Fatalf("decode error: %v", err)
			}
		}
	})
}

var discard = discardAt{}

type discardAt struct{}

var _ io.Writer = (*discardAt)(nil)
var _ io.WriterAt = (*discardAt)(nil)

func (discardAt) Write(p []byte) (int, error)                    { return len(p), nil }
func (discardAt) WriteAt(p []byte, off int64) (n int, err error) { return len(p), nil }

func BenchmarkEncode(b *testing.B) {
	b.Run("muktihari/fit raw", func(b *testing.B) {
		b.StopTimer()
		f, err := os.ReadFile(big_activity)
		if err != nil {
			b.Fatalf("open file: %v", err)
		}

		dec := decoder.New(bytes.NewReader(f))

		fit, err := dec.Decode()
		if err != nil {
			b.Fatalf("decode error: %v", err)
		}

		b.StartTimer()

		for i := 0; i < b.N; i++ {
			enc := encoder.New(discard) // include the encoder creation.
			if err := enc.Encode(fit); err != nil {
				b.Fatalf("encode error: %v", err)
			}
		}
	})
	b.Run("muktihari/fit", func(b *testing.B) {
		b.StopTimer()
		f, err := os.ReadFile(big_activity)
		if err != nil {
			b.Fatalf("open file: %v", err)
		}

		al := filedef.NewListener()
		defer al.Close()

		dec := decoder.New(bytes.NewReader(f),
			decoder.WithMesgListener(al),
			decoder.WithBroadcastOnly(),
		)

		_, err = dec.Decode()
		if err != nil {
			b.Fatalf("decode error: %v", err)
		}

		activity := al.File().(*filedef.Activity)

		b.StartTimer()

		for i := 0; i < b.N; i++ {
			fit := activity.ToFIT(nil)
			enc := encoder.New(discard)
			if err := enc.Encode(&fit); err != nil {
				b.Fatalf("encode error: %v", err)
			}
		}
	})
	b.Run("tormoder/fit", func(b *testing.B) {
		b.StopTimer()
		f, err := os.ReadFile(big_activity)
		if err != nil {
			b.Fatalf("open file: %v", err)
		}

		fitFile, err := fit.Decode(bytes.NewReader(f))
		if err != nil {
			b.Fatalf("decode error: %v", err)
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			if err := fit.Encode(discard, fitFile, binary.LittleEndian); err != nil {
				b.Fatalf("encode error: %v", err)
			}
		}
	})
}
