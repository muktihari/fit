// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fit_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/proto"
)

// generateSeedCorpus generates seed corpus using all FIT files within testdata
// including files in sub-directories. The files will be re-encoded for simplicity
// since each file may contain different encoding options such as encoded using big-endian,
// compressed timestamp header, multiple local message types, chained files, etc.
// Such cases can be tricky to handle and compare, often leading to messy implementation details.
// And some cases might also be impossible to reproduce due to different algorithm usage such as
// when generating local message numbers for multiple local message types.
func generateSeedCorpus(f *testing.F) {
	if err := filepath.Walk("testdata", func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(info.Name())
		if strings.ToLower(ext) != ".fit" {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("filename: %s: %w", info.Name(), err)
		}
		defer file.Close()

		dec := decoder.New(file, decoder.WithIgnoreChecksum())

		mv := encoder.NewMessageValidator(
			encoder.ValidatorWithPreserveInvalidValues())

		buf := &bufferAt{new(bytes.Buffer)}
		enc := encoder.New(buf,
			encoder.WithMessageValidator(mv),
			encoder.WithProtocolVersion(proto.V2),
		)

		for dec.Next() {
			fit, err := dec.Decode()
			if err != nil {
				return fmt.Errorf("filename: %s: %w", info.Name(), err)
			}
			if err := enc.Encode(fit); err != nil {
				return fmt.Errorf("filename: %s: %w", info.Name(), err)
			}
		}

		f.Add(buf.Bytes()) // add seed corpus

		return nil
	}); err != nil {
		f.Fatalf("could not generate seed corpus: %v", err)
	}
}

func FuzzDecodeEncodeRoundTrip(f *testing.F) {
	generateSeedCorpus(f)

	f.Fuzz(func(t *testing.T, in []byte) {
		r := bytes.NewReader(in)
		dec := decoder.New(r)

		mv := encoder.NewMessageValidator(
			encoder.ValidatorWithPreserveInvalidValues())

		buf := &bufferAt{new(bytes.Buffer)}
		enc := encoder.New(buf,
			encoder.WithMessageValidator(mv),
			encoder.WithProtocolVersion(proto.V2),
		)

		for dec.Next() {
			fit, err := dec.Decode()
			if err != nil {
				return
			}
			if len(fit.Messages) == 0 {
				return
			}
			if err := enc.Encode(fit); err != nil {
				t.Fatal(err)
			}
		}

		encoded := buf.Bytes()

		ignoreSpecificBytes(in, encoded)

		if diff := cmp.Diff(in[:len(encoded)], encoded); diff != "" {
			t.Fatal(diff)
		}
	})
}

// ignoreSpecificBytes ignores some bytes due to our testing limitation.
func ignoreSpecificBytes(in, encoded []byte) {
	in = in[:len(encoded)] // we only care in bytes as much as encoded bytes.

	// ignore for all FIT sequences.
	for len(in) > 0 {
		// ignore protocol version, the encoder used for test
		// will always encode to proto.V2
		in[1] = encoded[1]

		if len(in) < 8 {
			break
		}
		// ignore data size, we allow decoding bytes stream with
		// its content's size is not equal with file header's dataSize
		// as long as we could retrieve the content without error.
		copy(in[4:8], encoded[4:8])

		dataSize := binary.LittleEndian.Uint32(in[4:8])
		fileHeaderSize := uint32(in[0])
		if len(in) < int(fileHeaderSize+dataSize+2) {
			break
		}
		in = in[fileHeaderSize+dataSize+2:]
		encoded = encoded[fileHeaderSize+dataSize+2:]
	}
}

// bufferAt wraps bytes.Buffer to enable WriteAt for faster encoding.
type bufferAt struct{ *bytes.Buffer }

func (b *bufferAt) WriteAt(p []byte, off int64) (n int, err error) {
	return copy(b.Bytes()[off:], p), nil
}
