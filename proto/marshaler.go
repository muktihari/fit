// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"fmt"

	"github.com/muktihari/fit/kit/byteorder"
	"github.com/muktihari/fit/profile/typedef"
)

// Marshaler should only do one thing: marshaling to its bytes representation, any validation should be done outside.

const (
	DevDataMask = 0b00100000 // Mask for determining if a message contains developer fields.
)

var (
	_ encoding.BinaryMarshaler = &FileHeader{}
	_ encoding.BinaryMarshaler = &MessageDefinition{}
	_ encoding.BinaryMarshaler = &FieldDefinition{}
	_ encoding.BinaryMarshaler = &DeveloperFieldDefinition{}
	_ encoding.BinaryMarshaler = &Message{}
)

func (h *FileHeader) MarshalBinary() ([]byte, error) {
	b := make([]byte, 0, h.Size)

	var profileVersion = make([]byte, 2)
	binary.LittleEndian.PutUint16(profileVersion, h.ProfileVersion)
	var dataSize = make([]byte, 4)
	binary.LittleEndian.PutUint32(dataSize, h.DataSize)

	// Ensure the size of the DataType is fixed even if h.DataType is empty or even exceeding 4 bytes.
	var dataType = make([]byte, 4)
	copy(dataType[:4], []byte(h.DataType))

	b = append(b, h.Size)
	b = append(b, h.ProtocolVersion)
	b = append(b, profileVersion...)
	b = append(b, dataSize...)
	b = append(b, dataType...)

	var crc = make([]byte, 2)
	binary.LittleEndian.PutUint16(crc, h.CRC)
	b = append(b, crc...)

	return b, nil
}

func (m *MessageDefinition) MarshalBinary() ([]byte, error) {
	// 6 is the size of non-slice m's fields, and 3 is the amount of byte per field.
	size := 6 + (len(m.FieldDefinitions) * 3) + (len(m.DeveloperFieldDefinitions) * 3)
	b := make([]byte, 0, size)

	b = append(b, m.Header)
	b = append(b, m.Reserved)
	b = append(b, m.Architecture)

	globalMesgNum := make([]byte, 2)
	byteorder.Select(m.Architecture).PutUint16(globalMesgNum, uint16(m.MesgNum))
	b = append(b, globalMesgNum...)

	b = append(b, byte(len(m.FieldDefinitions)))

	for i := range m.FieldDefinitions {
		bs, _ := m.FieldDefinitions[i].MarshalBinary()
		b = append(b, bs...)
	}

	if (m.Header & DevDataMask) != DevDataMask {
		return b, nil
	}

	b = append(b, byte(len(m.DeveloperFieldDefinitions)))
	for i := range m.DeveloperFieldDefinitions {
		bs, _ := m.DeveloperFieldDefinitions[i].MarshalBinary()
		b = append(b, bs...)
	}

	return b, nil
}

func (f *FieldDefinition) MarshalBinary() ([]byte, error) {
	return []byte{f.Num, f.Size, byte(f.BaseType)}, nil
}

func (f *DeveloperFieldDefinition) MarshalBinary() ([]byte, error) {
	return []byte{f.Num, f.Size, f.DeveloperDataIndex}, nil
}

func (m *Message) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteByte(m.Header)

	for i := range m.Fields {
		field := &m.Fields[i]
		b, err := typedef.Marshal(field.Value, byteorder.Select(m.Architecture))
		if err != nil {
			return nil, fmt.Errorf("field: [num: %d, value: %v]: %w", field.Num, field.Value, err)
		}
		buf.Write(b)
	}

	for i := range m.DeveloperFields {
		developerField := &m.DeveloperFields[i]
		b, err := typedef.Marshal(developerField.Value, byteorder.Select(m.Architecture))
		if err != nil {
			return nil, fmt.Errorf("developer field: [num: %d, value: %v]: %w", developerField.Num, developerField.Value, err)
		}
		buf.Write(b)
	}

	return buf.Bytes(), nil
}
