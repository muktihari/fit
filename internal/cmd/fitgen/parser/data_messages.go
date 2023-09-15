// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

type Message struct {
	Name    string
	Comment string
	Fields  []Field
}

type ComponentField interface {
	GetComponents() []string
	GetScales() []float64
	GetOffsets() []float64
	GetAccumulate() []bool
	GetBits() []byte
}

var _ ComponentField = Field{}

type Field struct {
	Num        byte
	Name       string
	Type       string
	Array      string
	Components []string
	Scales     []float64
	Offsets    []float64
	Units      string
	Bits       []byte
	Accumulate []bool
	SubFields  []SubField
	Comment    string
	Product    string
	Example    string
}

func (f Field) GetComponents() []string { return f.Components }
func (f Field) GetScales() []float64    { return f.Scales }
func (f Field) GetOffsets() []float64   { return f.Offsets }
func (f Field) GetAccumulate() []bool   { return f.Accumulate }
func (f Field) GetBits() []byte         { return f.Bits }

var _ ComponentField = SubField{}

type SubField struct {
	FieldNum      byte
	Name          string
	Type          string
	Array         string
	Components    []string
	Scales        []float64
	Offsets       []float64
	Units         string
	Bits          []byte
	Accumulates   []bool
	RefFieldNames []string
	RefFieldValue []string
	Comment       string
	Product       string
	Example       string
}

func (f SubField) GetComponents() []string { return f.Components }
func (f SubField) GetScales() []float64    { return f.Scales }
func (f SubField) GetOffsets() []float64   { return f.Offsets }
func (f SubField) GetAccumulate() []bool   { return f.Accumulates }
func (f SubField) GetBits() []byte         { return f.Bits }
