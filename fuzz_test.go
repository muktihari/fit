// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fit_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
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
	var i int
	f.Log("seed corpus:")
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

		f.Logf("[%d] filename: %s", i, info.Name())
		f.Add(buf.Bytes()) // add seed corpus
		i++

		return nil
	}); err != nil {
		f.Fatalf("could not generate seed corpus: %v", err)
	}
}

func FuzzDecodeEncodeRoundTrip(f *testing.F) {
	generateSeedCorpus(f)

	f.Fuzz(func(t *testing.T, in []byte) {
		r := bytes.NewReader(in)
		dec := decoder.New(r, decoder.WithIgnoreChecksum())

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
				t.Skipf("could not be decoded: %v", err)
			}
			if len(fit.Messages) == 0 {
				t.Skipf("no messages")
			}
			if fit.Messages[0].Num != mesgnum.FileId {
				t.Skipf("missing file_id mesg")
			}
			if err := enc.Encode(fit); err != nil {
				if errors.Is(err, encoder.ErrMissingDeveloperDataId) {
					// Currently, the decoder is does not strictly verify whether
					// the DeveloperDataId message exist prior to decoding developer data,
					// as long as FieldDefinition messages are present to define the data.
					t.Skipf("missing developer data id message")
				}
				t.Fatal(err)
			}
		}

		encoded := buf.Bytes()

		// encoded bytes should be able to be decoded again.
		dec.Reset(bytes.NewReader(encoded))
		for dec.Next() {
			_, err := dec.Decode()
			if err != nil {
				t.Fatalf("resulting encoded could not be decoded: %v", err)
			}
		}

		if !comparable(in) {
			return
		}

		sanitizedIn := sanitizeOutput(in)
		sanitizedEncoded := sanitizeOutput(encoded)
		if diff := cmp.Diff(sanitizedIn, sanitizedEncoded); diff != "" {
			fitdump("in", t, sanitizedIn)
			fitdump("encoded", t, sanitizedEncoded)
			t.Fatal(diff)
		}
	})
}

// comparable checks whether the given bytes can be compared to encoded bytes
// due to our testing limitations.
func comparable(in []byte) bool {
	// uncomparable is an error flag when decoder allow some processes but
	// the resulting FIT is not reflecting with the actual bytes. In such cases,
	// encoder will produce different bytes from the actual input bytes.
	var uncomparable = fmt.Errorf("uncomparable")

	localMessageDefinitionUsageCounts := map[byte]int{}
	_, err := decoder.NewRaw().Decode(bytes.NewReader(in), func(flag decoder.RawFlag, b []byte) error {
		switch flag {
		case decoder.RawFlagFileHeader:
		case decoder.RawFlagMesgDef:
			if b[2] != 0 {
				// our encoder used for test encode using little-endian byte order.
				return uncomparable
			}

			localMesgNum := b[0] & proto.LocalMesgNumMask
			count, ok := localMessageDefinitionUsageCounts[localMesgNum]
			if ok && count == 0 {
				return uncomparable
			}
			localMessageDefinitionUsageCounts[localMesgNum] = 0

			num := typedef.MesgNum(binary.LittleEndian.Uint16(b[3:5]))
			for b = b[6:]; len(b) >= 3; b = b[3:] {
				fieldDef := proto.FieldDefinition{
					Num:      b[0],
					Size:     b[1],
					BaseType: basetype.BaseType(b[2]),
				}
				if fieldDef.BaseType == basetype.String {
					return uncomparable
				}
				field := factory.CreateField(num, fieldDef.Num)
				if field.Name == factory.NameUnknown {
					field.BaseType = fieldDef.BaseType
					field.Array = fieldDef.Size > field.BaseType.Size() && fieldDef.Size%field.BaseType.Size() == 0
				}
				if fieldDef.BaseType != field.BaseType {
					return uncomparable
				}
				if fieldDef.Size < field.BaseType.Size() {
					return uncomparable
				}
				if fieldDef.Size > field.BaseType.Size() && !field.Array {
					return uncomparable
				}
			}
		case decoder.RawFlagMesgData:
			localMessageDefinitionUsageCounts[proto.LocalMesgNum(b[0])]++

			if b[0]&proto.MesgCompressedHeaderMask == proto.MesgCompressedHeaderMask {
				return uncomparable
			}
		}
		return nil
	})

	// Bad encoded FIT files may contains dummy message definition
	// that are not being used by any of message.
	for _, count := range localMessageDefinitionUsageCounts {
		if count == 0 {
			return false
		}
	}
	return err == nil
}

// sanitizeOutput clears some bytes before comparison.
func sanitizeOutput(in []byte) []byte {
	out := make([]byte, 0, len(in))
	decoder.NewRaw().Decode(bytes.NewReader(in), func(flag decoder.RawFlag, b []byte) error {
		switch flag {
		case decoder.RawFlagFileHeader:
			// ignore protocol version, the encoder used for test
			// will always encode to proto.V2
			b[1] = byte(proto.V2)
			// ignore profileVersion when its value is zero
			// our encoder will replace it with current profile.Version
			if b[2] == 0 && b[3] == 0 {
				binary.LittleEndian.PutUint16(b[2:4], profile.Version)
			}
			// ignore data size, we allow decoding bytes stream with
			// its content's size is not equal with file header's dataSize
			// as long as we could retrieve the content without error.
			binary.LittleEndian.PutUint16(b[4:8], 0)
			if len(b) > 12 {
				binary.LittleEndian.PutUint16(b[12:14], 0) // ignore crc
			}
		case decoder.RawFlagMesgDef:
			// the encoder used for test always use localMesgNum 0
			b[0] = proto.MesgDefinitionMask
			b[1] = 0 // Reserved
		case decoder.RawFlagMesgData:
			// the encoder used for test always use localMesgNum 0
			b[0] = 0
		case decoder.RawFlagCRC:
			binary.LittleEndian.PutUint16(b, 0) // ignore crc
		}
		out = append(out, b...)
		return nil
	})
	return out
}

func fitdump(name string, t *testing.T, in []byte) {
	var i int
	t.Logf("%s:", name)
	decoder.NewRaw().Decode(bytes.NewReader(in), func(flag decoder.RawFlag, b []byte) error {
		t.Logf("%s: %v (%d-%d)\n", flag, b, i, i+len(b)-1)
		i += len(b)
		return nil
	})
}

// bufferAt wraps bytes.Buffer to enable WriteAt for faster encoding.
type bufferAt struct{ *bytes.Buffer }

func (b *bufferAt) WriteAt(p []byte, off int64) (n int, err error) {
	return copy(b.Bytes()[off:], p), nil
}
