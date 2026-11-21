// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/xlsxlite"
)

type Sheet byte

const (
	SheetTypes Sheet = iota
	SheetMessages
)

// Parser is Profile.xlsx parser
type Parser struct {
	xl     *xlsxlite.XlsxLite
	sheets map[Sheet]string
}

// New creates new Parser.
func New(xl *xlsxlite.XlsxLite, sheetNames map[Sheet]string) *Parser {
	return &Parser{xl: xl, sheets: sheetNames}
}

// ParseTypes parse sheet "Types" in Profile.xlsx
func (p *Parser) ParseTypes() ([]Type, error) {
	var ts []Type
	var cur = -1
	row := p.xl.RowIterator(p.sheets[SheetTypes])

	for row.Next() {
		if row.Index() == 1 { // header
			continue
		}
		if row.Cell("A") != "" {
			var (
				name     = row.Cell("A")
				baseType = row.Cell("B")
			)
			name = fixTypos(name, "Types", row.Index(), "type: type name")
			ts = append(ts, Type{
				Name:     name,
				BaseType: baseType,
			})
			cur++
			continue
		}

		var (
			name    = row.Cell("C")
			value   = row.Cell("D")
			comment = row.Cell("E")
		)
		name = fixTypos(name, "Types", row.Index(), "type: value name")
		comment = fixTypos(comment, "Types", row.Index(), "type: comment")

		ts[cur].Values = append(ts[cur].Values, Value{
			Name:    name,
			Value:   value,
			Comment: comment,
		})
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return ts, nil
}

// ParseTypes parse sheet "Messages" in Profile.xlsx
func (p *Parser) ParseMessages() ([]Message, error) {
	var ms []Message
	var cur = -1 // message cursor
	row := p.xl.RowIterator(p.sheets[SheetMessages])

	for row.Next() {
		if row.Index() == 1 { // header
			continue
		}

		if row.Cell("A") != "" {
			var (
				name    = row.Cell("A")
				comment = row.Cell("N")
			)
			name = fixTypos(name, "Messages", row.Index(), "mesg.Name")
			comment = fixTypos(comment, "Messages", row.Index(), "mesg.Comment")
			ms = append(ms, Message{
				Name:    name,
				Comment: comment,
			})
			cur++
			continue
		}

		if skipMessagesSection(row) {
			continue
		}

		fieldName := row.Cell("C")
		fieldName = fixTypos(fieldName, "Messages", row.Index(), "mesg.Name: "+ms[cur].Name+": field.Name")
		fieldType := row.Cell("D")
		array := row.Cell("E")
		components := splitStringOrNil(row.Cell("F"))

		scales, err := splitFloatsOrNil(row.Cell("G"))
		if err != nil {
			return nil, err
		}

		offsets, err := splitFloatsOrNil(row.Cell("H"))
		if err != nil {
			return nil, err
		}

		units := row.Cell("I")

		bits, err := splitBytesOrNil(row.Cell("J"))
		if err != nil {
			return nil, err
		}

		accumulates, err := splitBoolsOrNil(row.Cell("K"))
		if err != nil {
			return nil, err
		}

		comment := row.Cell("N")
		comment = fixTypos(comment, "Messages", row.Index(), "mesg.Name: "+ms[cur].Name+": field.Comment")

		product := row.Cell("O")
		example := row.Cell("P")

		if row.Cell("B") == "" { // sub-fields, doesn't have fieldNum
			field := ms[cur].Fields[len(ms[cur].Fields)-1] // field must exist before subfield, no need to check len.
			field.SubFields = append(field.SubFields, SubField{
				FieldNum:      field.Num,
				Name:          fieldName,
				Type:          fieldType,
				Components:    components,
				Scales:        scales,
				Offsets:       offsets,
				Units:         units,
				Bits:          bits,
				Accumulates:   accumulates,
				RefFieldNames: splitStringOrNil(row.Cell("L")),
				RefFieldValue: splitStringOrNil(row.Cell("M")),
				Comment:       comment,
				Product:       product,
				Example:       example,
			})
			ms[cur].Fields[len(ms[cur].Fields)-1] = field
			continue
		}

		// only field has fieldNum, sub-field doesn't.
		fieldNum, err := strconv.ParseUint(row.Cell("B"), 0, 8)
		if err != nil {
			return nil, err
		}

		ms[cur].Fields = append(ms[cur].Fields, Field{
			Num:        byte(fieldNum),
			Name:       fieldName,
			Type:       fieldType,
			Array:      array,
			Components: components,
			Scales:     scales,
			Offsets:    offsets,
			Units:      units,
			Bits:       bits,
			Accumulate: accumulates,
			Comment:    comment,
			Product:    product,
			Example:    example,
		})
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

func skipMessagesSection(row *xlsxlite.RowIterator) bool {
	// Previously on profile <= 21.188, it has len(cells) == 1 and cells[0].Name is "D".
	// Starting from profile 21.195, it has len(cells) == 16, cells[3].Name is "D", and only this cell has a value.
	cells := row.Cells()
	count := 0
	for _, v := range cells {
		if v.Value != "" {
			count++
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func splitStringOrNil(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

func splitFloatsOrNil(s string) ([]float64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	strVals := strings.Split(s, ",")
	floats := make([]float64, 0, len(strVals))
	for _, s := range strVals {
		scale, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, scale)
	}

	return floats, nil
}

func splitBytesOrNil(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	strVals := strings.Split(s, ",")
	bs := make([]byte, 0, len(strVals))
	for _, s := range strVals {
		b, err := strconv.ParseUint(s, 0, 8)
		if err != nil {
			return nil, err
		}
		bs = append(bs, byte(b))
	}

	return bs, nil
}

func splitBoolsOrNil(s string) ([]bool, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	strVals := strings.Split(s, ",")
	bs := make([]bool, 0, len(strVals))
	for _, s := range strVals {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}

	return bs, nil
}

// typoReplacer fixes known typos in Profile.xlsx so the generated identifiers and comments are spelled correctly.
//
//nolint:misspell // intentional typos
var typoReplacer = strings.NewReplacer(
	"accross", "across",
	"bondary", "boundary",
	"compatability", "compatibility",
	"corrosponding", "corresponding",
	"decribing", "describing",
	"farenheit", "fahrenheit",
	"managment", "management",
)

func fixTypos(s, sheet string, rowIndex int, location string) string {
	fixed := typoReplacer.Replace(s)
	if fixed != s {
		fmt.Printf("parser: typo: sheet: %q, row index: %d: %s: %q -> %q\n", sheet, rowIndex, location, s, fixed)
	}
	return fixed
}
