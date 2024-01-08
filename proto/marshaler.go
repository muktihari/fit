// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/muktihari/fit/kit/byteorder"
	"github.com/muktihari/fit/profile/typedef"
)

// Marshaler should only do one thing: marshaling to its bytes representation, any validation should be done outside.

// m.Header + ((max cap of m.Fields) * (n value)) + ((max cap of m.DeveloperFields) * (n value))
const MaxBytesPerMessage = 1 + (255*255)*2

var arrayPool = sync.Pool{
	New: func() any {
		b := [MaxBytesPerMessage]byte{}
		return &b
	},
}

var bufPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, MaxBytesPerMessage))
	},
}

var (
	// Zero alloc marshaler for efficient marshaling.
	// We don't need to implement io.WriterTo for FileHeader, as writing FileHeader is much less frequently
	// compared to writing MessageDefinition or Message.
	_ io.WriterTo = &Message{}
	_ io.WriterTo = &MessageDefinition{}

	_ encoding.BinaryMarshaler = &FileHeader{}
	_ encoding.BinaryMarshaler = &MessageDefinition{}
	_ encoding.BinaryMarshaler = &Message{}
)

func (h *FileHeader) MarshalBinary() ([]byte, error) {
	b := make([]byte, h.Size)

	b[0] = h.Size
	b[1] = h.ProtocolVersion

	binary.LittleEndian.PutUint16(b[2:4], h.ProfileVersion)
	binary.LittleEndian.PutUint32(b[4:8], h.DataSize)

	copy(b[8:12], []byte(h.DataType))

	if h.Size < 14 {
		return b, nil
	}

	binary.LittleEndian.PutUint16(b[12:14], h.CRC)

	return b, nil
}

func (m *MessageDefinition) MarshalBinary() ([]byte, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()

	_, _ = m.WriteTo(buf)

	b := make([]byte, buf.Len())
	copy(b, buf.Bytes())

	return b, nil
}

// WriteTo zero alloc marshal then copy it to w.
func (m *MessageDefinition) WriteTo(w io.Writer) (n int64, err error) {
	arr := arrayPool.Get().(*[MaxBytesPerMessage]byte)
	defer arrayPool.Put(arr)
	b := (*arr)[:0]

	b = append(b, m.Header)
	b = append(b, m.Reserved)
	b = append(b, m.Architecture)

	b = append(b, 0, 0)
	byteorder.Select(m.Architecture).PutUint16(b[len(b)-2:], uint16(m.MesgNum))

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

	nn, err := w.Write(b)
	return int64(nn), err
}

func (m *Message) MarshalBinary() ([]byte, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()

	_, err := m.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	b := make([]byte, buf.Len())
	copy(b, buf.Bytes())

	return b, nil
}

// WriteTo zero alloc marshal then copy it to w.
func (m *Message) WriteTo(w io.Writer) (n int64, err error) {
	arr := arrayPool.Get().(*[MaxBytesPerMessage]byte)
	defer arrayPool.Put(arr)
	b := (*arr)[:0]

	b = append(b, m.Header)

	for i := range m.Fields {
		field := &m.Fields[i]
		err = typedef.MarshalTo(&b, field.Value, byteorder.Select(m.Architecture))
		if err != nil {
			return 0, fmt.Errorf("field: [num: %d, value: %v]: %w", field.Num, field.Value, err)
		}
	}

	for i := range m.DeveloperFields {
		developerField := &m.DeveloperFields[i]
		err = typedef.MarshalTo(&b, developerField.Value, byteorder.Select(m.Architecture))
		if err != nil {
			return 0, fmt.Errorf("developer field: [num: %d, value: %v]: %w", developerField.Num, developerField.Value, err)
		}
	}

	nn, err := w.Write(b)
	return int64(nn), err
}
