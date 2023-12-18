// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

type Data struct {
	SDKVersion  string
	Package     string
	Imports     []string
	Name        string
	Fields      []Field
	MaxFieldNum byte
}

type Field struct {
	Num             byte
	Name            string
	String          string
	Type            string
	TypedValue      string
	PrimitiveValue  string
	ComparableValue string
	InvalidValue    string
	Comment         string
}
