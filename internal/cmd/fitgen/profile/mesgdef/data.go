// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

type Data struct {
	Package           string
	Imports           []string
	Name              string
	Fields            []Field
	OptimizedFields   []Field
	DynamicFields     []DynamicField
	MaxFieldNum       byte
	MaxFieldExpandNum byte
}

type Field struct {
	Num             byte
	Name            string
	String          string
	ProfileType     string
	BaseType        string
	Size            byte
	Type            string
	TypedValue      string
	PrimitiveValue  string
	ProtoValue      string
	ComparableValue string
	InvalidValue    string
	Comment         string
	Units           string
	Scale           float64
	Offset          float64
	Array           bool
	CanExpand       bool
}

type DynamicField struct {
	Name        string
	SwitchCases []SwitchCase
	Default     ReturnValue
}

type SwitchCase struct {
	Name       string
	CondValues []CondValue
}

type CondValue struct {
	Conds       []string
	ReturnValue ReturnValue
}

type ReturnValue struct {
	Name  string
	Units string
	Value string
}
