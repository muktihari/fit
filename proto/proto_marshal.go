// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"encoding"
	"encoding/binary"
	"fmt"
	"sync"
)

const littleEndian = 0

// Marshaler should only do one thing: marshaling to its bytes representation, any validation should be done outside.

// Header + ((max n Fields) * (n value)) + ((max n DeveloperFields) * (n value))
const MaxBytesPerMessage = 1 + (255*255)*2

// Header + Reserved + Architecture + MesgNum (2 bytes) + n Fields + (Max n Fields * 3) + n DevFields + (Max n DevFields * 3).
const MaxBytesPerMessageDefinition = 5 + 1 + (255 * 3) + 1 + (255 * 3)

var pool = sync.Pool{New: func() any { return new([MaxBytesPerMessage]byte) }}

var (
	_ encoding.BinaryMarshaler = &FileHeader{}
	_ encoding.BinaryMarshaler = &MessageDefinition{}
	_ encoding.BinaryMarshaler = &Message{}
)

// MarshalBinary returns the FIT format encoding of FileHeader and nil error.
func (h FileHeader) MarshalBinary() ([]byte, error) {
	arr := pool.Get().(*[MaxBytesPerMessage]byte)
	defer pool.Put(arr)
	b := arr[:0]

	b, _ = h.MarshalAppend(b)

	return append([]byte{}, b...), nil
}

// MarshalAppend appends the FIT format encoding of FileHeader to b, returning the result.
func (h FileHeader) MarshalAppend(b []byte) ([]byte, error) {
	b = append(b, h.Size, h.ProtocolVersion)
	b = binary.LittleEndian.AppendUint16(b, h.ProfileVersion)
	b = binary.LittleEndian.AppendUint32(b, h.DataSize)
	b = append(b, h.DataType[:4]...)
	if h.Size >= 14 {
		b = binary.LittleEndian.AppendUint16(b, h.CRC)
	}
	return b, nil
}

// MarshalBinary returns the FIT format encoding of MessageDefinition and nil error.
func (m MessageDefinition) MarshalBinary() ([]byte, error) {
	arr := pool.Get().(*[MaxBytesPerMessage]byte)
	defer pool.Put(arr)
	b := arr[:0]

	b, _ = m.MarshalAppend(b)

	return append([]byte{}, b...), nil
}

// MarshalAppend appends the FIT format encoding of MessageDefinition to b, returning the result.
func (m MessageDefinition) MarshalAppend(b []byte) ([]byte, error) {
	b = append(b, m.Header)
	b = append(b, m.Reserved)
	b = append(b, m.Architecture)

	if m.Architecture == littleEndian {
		b = binary.LittleEndian.AppendUint16(b, uint16(m.MesgNum))
	} else {
		b = binary.BigEndian.AppendUint16(b, uint16(m.MesgNum))
	}

	b = append(b, byte(len(m.FieldDefinitions)))
	for i := range m.FieldDefinitions {
		b = append(b,
			m.FieldDefinitions[i].Num,
			m.FieldDefinitions[i].Size,
			byte(m.FieldDefinitions[i].BaseType),
		)
	}

	if (m.Header & DevDataMask) == DevDataMask {
		b = append(b, byte(len(m.DeveloperFieldDefinitions)))
		for i := range m.DeveloperFieldDefinitions {
			b = append(b,
				m.DeveloperFieldDefinitions[i].Num,
				m.DeveloperFieldDefinitions[i].Size,
				m.DeveloperFieldDefinitions[i].DeveloperDataIndex,
			)
		}
	}

	return b, nil
}

// MarshalBinary returns the FIT format encoding of Message and any error encountered during marshal.
func (m Message) MarshalBinary() ([]byte, error) {
	arr := pool.Get().(*[MaxBytesPerMessage]byte)
	defer pool.Put(arr)
	b := arr[:0]

	b, err := m.MarshalAppend(b)
	if err != nil {
		return nil, err
	}

	return append([]byte{}, b...), nil
}

// MarshalAppend appends the FIT format encoding of Message to b, returning the result.
func (m Message) MarshalAppend(b []byte) ([]byte, error) {
	b = append(b, m.Header)

	var err error
	for i := range m.Fields {
		b, err = m.Fields[i].Value.MarshalAppend(b, m.Architecture)
		if err != nil {
			return nil, fmt.Errorf("field: [num: %d, value: %v]: %w",
				m.Fields[i].Num, m.Fields[i].Value.Any(), err)
		}
	}

	for i := range m.DeveloperFields {
		b, err = m.DeveloperFields[i].Value.MarshalAppend(b, m.Architecture)
		if err != nil {
			return nil, fmt.Errorf("developer field: [num: %d, value: %v]: %w",
				m.DeveloperFields[i].Num, m.DeveloperFields[i].Value.Any(), err)
		}
	}

	return b, nil
}
