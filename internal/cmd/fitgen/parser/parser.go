// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/client9/misspell"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/xlsxlite"
)

type Sheet byte

const (
	SheetTypes Sheet = iota
	SheetMessages
)

// Parser is Profile.xlsx parser
type Parser struct {
	xl               *xlsxlite.XlsxLite
	sheets           map[Sheet]string
	misspellReplacer *misspell.Replacer
}

// New creates new Parser.
func New(xl *xlsxlite.XlsxLite, sheetNames map[Sheet]string) *Parser {
	return &Parser{xl: xl, sheets: sheetNames, misspellReplacer: misspell.New()}
}

// ParseTypes parse sheet "Types" in Profile.xlsx
func (p *Parser) ParseTypes() ([]Type, error) {
	var ts []Type
	var cur = -1
	row := p.xl.RowIterator(p.sheets[SheetTypes])

	var diff []misspell.Diff
	for row.Next() {
		if row.Index() == 1 { // header
			continue
		}
		if row.Cell("A") != "" {
			var (
				name     = row.Cell("A")
				baseType = row.Cell("B")
			)
			name, diff = p.misspellReplacer.Replace(name)
			if len(diff) > 0 {
				fmt.Printf("parser: misspell: sheet: %q, row index: %d: type: type name: %v\n",
					"Types", row.Index(), formatMisspellDiff(diff))
			}
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
		name, diff = p.misspellReplacer.Replace(name)
		if len(diff) > 0 {
			fmt.Printf("parser: misspell: sheet: %q, row index: %d: type: value name: %v\n",
				"Types", row.Index(), formatMisspellDiff(diff))
		}
		comment, diff = p.misspellReplacer.Replace(comment)
		if len(diff) > 0 {
			fmt.Printf("parser: misspell: sheet: %q, row index: %d: type: comment: %v\n",
				"Types", row.Index(), formatMisspellDiff(diff))
		}

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

	var misspellDiff []misspell.Diff
	for row.Next() {
		if row.Index() == 1 { // header
			continue
		}

		if row.Cell("A") != "" {
			var (
				name    = row.Cell("A")
				comment = row.Cell("N")
			)
			name, misspellDiff = p.misspellReplacer.Replace(name)
			if len(misspellDiff) > 0 {
				fmt.Printf("parser: misspell: sheet: %q, row index: %d: mesg.Name: %v\n",
					"Messages", row.Index(), formatMisspellDiff(misspellDiff))
			}
			comment, misspellDiff = p.misspellReplacer.Replace(comment)
			if len(misspellDiff) > 0 {
				fmt.Printf("parser: misspell: sheet: %q, row index: %d: mesg.Comment: %v\n",
					"Messages", row.Index(), formatMisspellDiff(misspellDiff))
			}
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
		fieldName, misspellDiff = p.misspellReplacer.Replace(fieldName)
		if len(misspellDiff) > 0 {
			fmt.Printf("parser: misspell: sheet: %q, row index: %d: mesg.Name: %s: field.Name: %v\n",
				"Messages", row.Index(), ms[cur].Name, formatMisspellDiff(misspellDiff))
		}
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
		comment, misspellDiff = p.misspellReplacer.Replace(comment)
		if len(misspellDiff) > 0 {
			fmt.Printf("parser: misspell: sheet: %q, row index: %d: mesg.Name: %s: field.Comment: %v\n",
				"Messages", row.Index(), ms[cur].Name, formatMisspellDiff(misspellDiff))
		}

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
	cells := row.Cells()
	if len(cells) == 1 && cells[0].Name == "D" {
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

func formatMisspellDiff(diff []misspell.Diff) string {
	strbuf := new(strings.Builder)
	fmt.Fprintf(strbuf, "[\n")
	for i := range diff {
		fullLine := "full text: <omitted: too long> -> "
		if len(diff[i].FullLine) <= 50 {
			fullLine = fmt.Sprintf("full text: %q -> ", diff[i].FullLine)
		}
		fmt.Fprintf(strbuf, "- %s%q is a misspelling of %q\n", fullLine, diff[i].Original, diff[i].Corrected)
	}
	fmt.Fprintf(strbuf, "]")
	return strbuf.String()
}
