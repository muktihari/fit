// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

type Data struct {
	Package           string
	Imports           []string
	Name              string
	Fields            []Field
	MaxFieldNum       byte
	MaxFieldExpandNum byte
}

type Field struct {
	Num             byte
	Name            string
	String          string
	ProfileType     string
	BaseType        string
	Size            int
	Type            string
	TypedValue      string
	PrimitiveValue  string
	ComparableValue string
	InvalidValue    string
	Comment         string
	Scale           float64
	Offset          float64
	Array           bool
	CanExpand       bool
}
