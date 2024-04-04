// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shared

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"ToTitle": strutil.ToTitle,
		"ToLower": strings.ToLower,
		"sprintf": fmt.Sprintf,
	}
}

// ConstantData is data representation of constant definitions template.
// This map one to one with file that will be generated. 1 data == 1 proto.
type ConstantData struct {
	PackageDoc    string
	Package       string
	Imports       []string
	AllowRegister bool
	Type          string
	Base          string
	Constants     []Constant
	Invalid       Constant
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
