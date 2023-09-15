// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shared

import (
	"fmt"
	"strings"
	"text/template"
)

type StringerMode byte

const (
	StringerMap StringerMode = iota // default
	StringerArray
	NoStringer
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"ToLower": strings.ToLower,
		"sprintf": fmt.Sprintf,
	}
}

// ConstantData is data representation of constant definitions template.
// This map one to one with file that will be generated. 1 data == 1 proto.
type ConstantData struct {
	SDKVersion    string
	PackageDoc    string
	Package       string
	Imports       []string
	StringerMode  StringerMode
	AllowRegister bool
	Type          string
	Base          string
	Constants     []Constant
}

// Constant represent declared constants within proto.
type Constant struct {
	Name    string
	Type    string
	Op      string
	Value   string
	String  string
	Comment string
}
