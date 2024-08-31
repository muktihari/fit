// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder_test

import (
	"io"
	"os"
	"testing"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/filedef"
)

type fnReader func(b []byte) (n int, err error)

func (f fnReader) Read(b []byte) (n int, err error) { return f(b) }

func BenchmarkDecode(b *testing.B) {
	b.StopTimer()
	// This is not a typical FIT in term of file size (2.3M) and the messages size.
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

	cur := 0
	r := fnReader(func(b []byte) (n int, err error) {
		if cur == len(all) {
			return 0, io.EOF
		}
		n = copy(b, all[cur:])
		cur += n
		return
	})

	dec := decoder.New(r)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err = dec.Decode()
		if err != nil {
			b.Fatal(err)
		}
		dec.Reset(r)
		cur = 0 // reset reader
	}
}

func BenchmarkDecodeWithFiledef(b *testing.B) {
	b.StopTimer()
	// This is not a typical FIT in term of file size (2.3M) and the messages size.
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

	lis := filedef.NewListener()
	defer lis.Close()

	cur := 0
	r := fnReader(func(b []byte) (n int, err error) {
		if cur == len(all) {
			return 0, io.EOF
		}
		n = copy(b, all[cur:])
		cur += n
		return
	})

	opts := []decoder.Option{
		decoder.WithMesgListener(lis),
		decoder.WithBroadcastOnly(),
	}
	dec := decoder.New(r, opts...)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err = dec.Decode()
		if err != nil {
			b.Fatal(err)
		}
		_ = lis.File()
		dec.Reset(r, opts...)
		cur = 0 // reset reader
	}
}

func BenchmarkCheckIntegrity(b *testing.B) {
	b.StopTimer()
	// This is not a typical FIT in term of file size (2.3M) and the messages size.
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

	cur := 0
	r := fnReader(func(b []byte) (n int, err error) {
		if cur == len(all) {
			return 0, io.EOF
		}
		n = copy(b, all[cur:])
		cur += n
		return
	})

	dec := decoder.New(r)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err = dec.CheckIntegrity()
		if err != nil {
			b.Fatal(err)
		}
		cur = 0 // reset reader
	}
}

func BenchmarkReset(b *testing.B) {
	b.Run("benchmark New()", func(b *testing.B) {
		b.StopTimer()
		lis := filedef.NewListener()
		defer lis.Close()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			_ = decoder.New(nil, decoder.WithMesgListener(lis))
		}
	})
	b.Run("benchmark Reset()", func(b *testing.B) {
		b.StopTimer()
		lis := filedef.NewListener()
		defer lis.Close()

		dec := decoder.New(nil, decoder.WithMesgListener(lis))
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			dec.Reset(nil, decoder.WithMesgListener(lis))
		}
	})
}
