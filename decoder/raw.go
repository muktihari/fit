// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"encoding/binary"
	"fmt"
	"io"
	"strconv"

	"github.com/muktihari/fit/proto"
)

// RawFlag is the kind of the incomming bytes, the size of the incomming bytes is vary but the the size is guaranteed
// by the corresponding RawFlag.
type RawFlag byte

const (
	// RawFlagFileHeader is guaranteed to have either 12 or 14 bytes (all in little-endian byte order):
	// Size + ProtocolVersion + ProfileVersion (2 bytes) +  DataSize (4 bytes) + DataType (4 bytes) + (only if Size is 14) CRC (2 bytes)
	RawFlagFileHeader RawFlag = iota

	// RawFlagMesgDef is guaranteed to have:
	// Header + Reserved + Architecture + MesgNum (2 bytes) + n FieldDefinitions + (n FieldDefinitions * 3) +
	// (only if Header & 0b00100000 == 0b00100000) n DeveloperFieldDefinitions + (n DeveloperFieldDefinitions * 3)
	RawFlagMesgDef

	// RawFlagMesgData is guaranteed to have:
	// Header + Fields' value represented by its Message Definition + (only if it has developer fields) Developer Fields' value.
	RawFlagMesgData

	// RawFlagCRC is guaranteed to have:
	// 2 bytes (in little-endian byte order) as the checksum of the messages.
	RawFlagCRC
)

func (f RawFlag) String() string {
	switch f {
	case RawFlagFileHeader:
		return "file_header"
	case RawFlagMesgDef:
		return "message_definition"
	case RawFlagMesgData:
		return "message_data"
	case RawFlagCRC:
		return "crc"
	}
	return "unknown(" + strconv.Itoa(int(f)) + ")"
}

// RawDecoder is a sequence of FIT bytes decoder. See NewRaw() for details.
type RawDecoder struct {
	bytesArray [255 * 255 * 2]byte // [MesgDef: 6 + 255 * 3 = 771] < [Mesg: (255 * 255) * 2 = 130050]. Use bigger capacity.
	lenMesgs   [proto.LocalMesgNumMask + 1]uint32
}

// NewRaw creates new RawDecoder which provides low-level building block to work with FIT bytes for the maximum performance gain.
// RawDecoder will split bytes by its corresponding RawFlag (FileHeader, MessageDefinition, MessageData and CRC) for scoping the operation.
//
// However, this is still considered unsafe operation since we work with bytes directly and the responsibility for validation now placed on the user-space. The only thing that this validates is
// the reader should be a FIT (FileHeader: has valid Size, bytes 8-12 is ".FIT", ProtocolVersion is supported and DataSize > 0).
//
// This is only intended to be used for performance and memory critical situation where every computation or memory usage is constrained
// (RawDecoder itself is using constant memory < 131 KB and the Decode method has zero heap alloc (except errors) while it may use additional small stack memory).
// The implementation of the callback function is also expected to have minimal overhead.
//
// For general purpose usage, use Decoder instead.
func NewRaw() *RawDecoder {
	return &RawDecoder{}
}

// Decode decodes r reader into sequence of FIT bytes splitted by its corresponding RawFlag (FileHeader, MessageDefinition, MessageData and CRC)
// for every FIT sequences in the reader, until it reaches EOF. It returns the number of bytes read and any error encountered.
// When fn returns an error, Decode will immediately return the error.
//
// For performance, the b is not copied and the underlying array's values will be replaced each fn call. If you need
// to work with b in its slice form later on, it should be cloned.
//
// Note: We encourage wrapping r into a buffered reader such as bufio.NewReader(r),
// decode process requires byte by byte reading and having frequent read on non-buffered reader might impact performance,
// especially if it involves syscall such as reading a file.
func (d *RawDecoder) Decode(r io.Reader, fn func(flag RawFlag, b []byte) error) (n int64, err error) {
	defer d.reset() // Must reset before return, so we can invoke Decode again for the next reader.

	var seq int
	for {
		// 1. Decode File Header
		nr, err := io.ReadFull(r, d.bytesArray[:1])
		n += int64(nr)
		if seq != 0 && err == io.EOF {
			return n, nil // Reach desirable EOF.
		}
		if err != nil {
			return n, err
		}

		headerSize := d.bytesArray[0]
		if headerSize != 12 && headerSize != 14 {
			return n, fmt.Errorf("header size [%d]: %w", headerSize, ErrNotAFitFile)
		}

		nr, err = io.ReadFull(r, d.bytesArray[1:headerSize])
		n += int64(nr)
		if err != nil {
			return n, err
		}

		if string(d.bytesArray[8:12]) != proto.DataTypeFIT {
			return n, ErrNotAFitFile
		}

		if err := proto.Validate(d.bytesArray[1]); err != nil {
			return n, err
		}

		dataSize := binary.LittleEndian.Uint32(d.bytesArray[4:8])
		if dataSize == 0 {
			return n, ErrDataSizeZero
		}

		if err := fn(RawFlagFileHeader, d.bytesArray[:headerSize]); err != nil {
			return n, err
		}

		// 2. Decode Messages
		var pos = int64(n)
		for uint32(n-pos) < dataSize {
			nr, err = io.ReadFull(r, d.bytesArray[:1])
			n += int64(nr)
			if err != nil {
				return n, fmt.Errorf("mesg's header: %w", err)
			}

			// 2. a. Decode Message Definition
			if (d.bytesArray[0] & proto.MesgDefinitionMask) == proto.MesgDefinitionMask {
				const fixedSize = uint16(6) //  Header + Reserved + Architecture + MesgNum (2 bytes) + n Fields
				nr, err = io.ReadFull(r, d.bytesArray[1:fixedSize])
				n += int64(nr)
				if err != nil {
					return n, fmt.Errorf("mesgDef bytes 1-5: %w", err)
				}
				lenMesgDef := fixedSize

				nFields := uint16(d.bytesArray[5])
				nr, err = io.ReadFull(r, d.bytesArray[lenMesgDef:lenMesgDef+nFields*3])
				n += int64(nr)
				if err != nil {
					return n, fmt.Errorf("fieldDefs: %w", err)
				}
				lenMesgDef += nFields * 3 // 3 bytes per field

				// Calculate the Message Data's size as we read the Field and DeveloperField definitions.
				lenMesg := uint32(1) // Header
				const fieldFirstIndex = fixedSize
				for i := uint16(0); i < nFields*3; i += 3 {
					lenMesg += uint32(d.bytesArray[fieldFirstIndex+i+1]) // // [0, |1|, 2] -> [Num, |Size|, Type]
				}

				if (d.bytesArray[0] & proto.DevDataMask) == proto.DevDataMask {
					nr, err = io.ReadFull(r, d.bytesArray[lenMesgDef:lenMesgDef+1])
					n += int64(nr)
					if err != nil {
						return n, fmt.Errorf("nDevFieldDef: %w", err)
					}

					nDevFields := uint16(d.bytesArray[lenMesgDef])
					lenMesgDef += 1
					devFieldFirstIndex := lenMesgDef
					nr, err = io.ReadFull(r, d.bytesArray[devFieldFirstIndex:devFieldFirstIndex+nDevFields*3])
					n += int64(nr)
					if err != nil {
						return n, fmt.Errorf("devFieldDefs: %w", err)
					}
					lenMesgDef += nDevFields * 3 // 3 bytes per field

					for i := uint16(0); i < nDevFields*3; i += 3 {
						lenMesg += uint32(d.bytesArray[devFieldFirstIndex+i+1]) // [0, |1|, 2] -> [Num, |Size|, Type]
					}
				}

				localMesgNum := d.bytesArray[0] & proto.LocalMesgNumMask
				d.lenMesgs[localMesgNum] = lenMesg

				if err := fn(RawFlagMesgDef, d.bytesArray[:lenMesgDef]); err != nil {
					return n, err
				}

				continue
			}

			// 2. b. Decode Message Data
			localMesgNum := proto.LocalMesgNum(d.bytesArray[0])
			lenMesg := d.lenMesgs[localMesgNum]
			if lenMesg == 0 {
				return n, fmt.Errorf("localMesgNum: %d: %w", localMesgNum, ErrMesgDefMissing)
			}

			nr, err = io.ReadFull(r, d.bytesArray[1:lenMesg])
			n += int64(nr)
			if err != nil {
				return n, fmt.Errorf("mesg: %w", err)
			}

			if err = fn(RawFlagMesgData, d.bytesArray[:lenMesg]); err != nil {
				return n, err
			}
		}

		// 3. Decode File CRC
		nr, err = io.ReadFull(r, d.bytesArray[:2])
		n += int64(nr)
		if err != nil {
			return n, err
		}

		if err = fn(RawFlagCRC, d.bytesArray[:2]); err != nil {
			return n, err
		}

		seq++
	}
}

func (d *RawDecoder) reset() {
	for i := range d.lenMesgs {
		d.lenMesgs[i] = 0
	}
}
