// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xlsxlite

import (
	"github.com/thedatashed/xlsxreader" // choosing this library due to its lightweight nature, with fewer features compared to libraries like tealeg or excelize.
)

// XlsxLite is a wrapper for any implementation used to process the xlsx file format, designed to be as minimal as possible.
// In case the underlying library or implementation changes, modifications will only be necessary within this package.
type XlsxLite struct {
	reader *xlsxreader.XlsxFileCloser
}

// New creates new XlsxLite
func New(reader *xlsxreader.XlsxFileCloser) *XlsxLite {
	return &XlsxLite{reader: reader}
}

// Sheets returns all sheet names.
func (x *XlsxLite) Sheets() []string { return x.reader.Sheets }

// RowIterator returns row's iterator.
func (x *XlsxLite) RowIterator(sheet string) *RowIterator {
	return &RowIterator{rowch: x.reader.ReadRows(sheet)}
}

// RowIterator for iterating through rows.
type RowIterator struct {
	rowch <-chan xlsxreader.Row
	row   xlsxreader.Row
	ok    bool
}

// Next returns whether row has next iteration.
func (ri *RowIterator) Next() bool {
	if ri.row.Error != nil {
		return false
	}
	ri.row, ri.ok = <-ri.rowch
	return ri.ok
}

func (ri *RowIterator) Err() error {
	return ri.row.Error
}

// Index returns current row's index
func (ri *RowIterator) Index() int {
	return ri.row.Index
}

// Cell return value as string, return empty string if column not found.
func (ri *RowIterator) Cell(column string) string {
	for _, cell := range ri.row.Cells {
		if cell.Column == column {
			return cell.Value
		}
	}
	return ""
}

// Cells returns only all cell that have values
func (ri *RowIterator) Cells() []Cell {
	cells := make([]Cell, 0, len(ri.row.Cells))
	for _, cell := range ri.row.Cells {
		cells = append(cells, Cell{
			Name:  cell.Column,
			Value: cell.Value,
		})
	}
	return cells
}

// Cell is minimal representation of xlsx's cell
type Cell struct {
	Name  string
	Value string
}
