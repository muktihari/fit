// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

type Type struct {
	Name     string
	BaseType string
	Values   []Value
}

type Value struct {
	Name    string
	Value   string
	Comment string
}
