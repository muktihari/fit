// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/profile/basetype"
)

type mesgdefBuilder struct {
	template     *template.Template
	templateExec string

	path       string // path to generate the file
	sdkVersion string // Fit SDK Version

	once sync.Once

	messages []parser.Message
	types    []parser.Type

	baseTypeMapByProfileType map[string]string
}

func NewBuilder(path, version string, message []parser.Message, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &mesgdefBuilder{
		template: template.Must(template.New("main").
			ParseFiles(filepath.Join(cd, "mesgdef.tmpl"))),
		templateExec:             "mesgdef",
		path:                     filepath.Join(path, "profile", "mesgdef"),
		sdkVersion:               version,
		messages:                 message,
		types:                    types,
		baseTypeMapByProfileType: make(map[string]string),
	}
}

func (b *mesgdefBuilder) populateLookupData() {
	for _, _basetype := range basetype.List() { // map to itself
		b.baseTypeMapByProfileType[_basetype.String()] = _basetype.String()
	}
	b.baseTypeMapByProfileType["bool"] = "bool" // additional profile type which is not defined in basetype.

	for _, _type := range b.types {
		b.baseTypeMapByProfileType[_type.Name] = _type.BaseType
	}
}

func (b *mesgdefBuilder) Build() ([]builder.Data, error) {
	b.once.Do(func() { b.populateLookupData() })

	dataBuilders := make([]builder.Data, 0, len(b.messages))
	for _, mesg := range b.messages {
		dataBuilder := builder.Data{
			Template:     b.template,
			TemplateExec: b.templateExec,
			Path:         b.path,
			Filename:     strutil.ToSnake(mesg.Name) + "_gen.go",
		}

		imports := map[string]struct{}{}
		fields := make([]Field, 0, len(mesg.Fields))
		for i := range mesg.Fields {
			field := &mesg.Fields[i]
			f := Field{
				Num:           field.Num,
				Name:          strutil.ToTitle(mesg.Fields[i].Name),
				String:        field.Name,
				Type:          transformType(field.Type, field.Array),
				AssignedValue: b.transformValue(field.Num, field.Type, field.Array),
				Comment:       field.Comment,
			}

			if field.Units != "" {
				f.Comment = fmt.Sprintf("Units: %s; %s", field.Units, field.Comment)
			}

			if strings.HasPrefix(f.Type, "[]") {
				f.Comment = fmt.Sprintf("Array: %s; %s", field.Array, f.Comment)
			}

			offset := offsetOrDefault(field.Offsets, 0)
			if offset != 0 {
				f.Comment = fmt.Sprintf("Offset: %g; %s", offset, f.Comment)
			}

			scale := scaleOrDefault(field.Scales, 0)
			if scale != 1 {
				f.Comment = fmt.Sprintf("Scale: %g; %s", scale, f.Comment)

			}

			fields = append(fields, f)
			if strings.HasPrefix(f.Type, "basetype") {
				imports["github.com/muktihari/fit/profile/basetype"] = struct{}{}
			}
		}

		data := Data{
			SDKVersion: b.sdkVersion,
			Package:    "mesgdef",
			Imports:    []string{},
			Name:       strutil.ToTitle(mesg.Name),
			Fields:     fields,
		}

		for k := range imports {
			data.Imports = append(data.Imports, k)
		}

		dataBuilder.Data = data
		dataBuilders = append(dataBuilders, dataBuilder)
	}

	return dataBuilders, nil
}

func transformType(name, array string) string {
	isBaseType := make(map[string]string)
	for _, bt := range basetype.List() {
		isBaseType[bt.String()] = bt.GoType()
	}
	isBaseType["bool"] = "bool"
	isBaseType["fit_base_type"] = "basetype.BaseType"

	var _type string
	if v, ok := isBaseType[name]; ok {
		_type = v
	} else {
		_type = fmt.Sprintf("typedef.%s", strutil.ToTitle(name))
	}

	if array == "" || _type == "string" {
		return _type
	}

	return fmt.Sprintf("[]%s", _type)
}

func (b *mesgdefBuilder) transformValue(num byte, name, array string) string {
	isBaseType := make(map[string]string)
	for _, bt := range basetype.List() {
		isBaseType[bt.String()] = bt.GoType()
	}
	isBaseType["bool"] = "bool"
	isBaseType["fit_base_type"] = "basetype.BaseType"

	if name == "fit_base_type" {
		return fmt.Sprintf("typeconv.ToUint8[basetype.BaseType](vals[%d])", num)
	}

	var _type string
	if v, ok := isBaseType[name]; ok {
		_type = v
	} else {
		_type = fmt.Sprintf("typedef.%s", strutil.ToTitle(name))
	}

	if array == "" || _type == "string" {
		return fmt.Sprintf("typeconv.To%s[%s](vals[%d])", strutil.ToTitle(b.baseTypeMapByProfileType[name]), _type, num)
	}

	return fmt.Sprintf("typeconv.ToSlice%s[%s](vals[%d])", strutil.ToTitle(b.baseTypeMapByProfileType[name]), _type, num)
}

// Profile.xlsx says unless otherwise specified, scale of 1 is assumed.
func scaleOrDefault(scales []float64, index int) float64 {
	if index < len(scales) {
		return scales[index]
	}
	return 1.0
}

// Profile.xlsx says unless otherwise specified, offset of 0 is assumed.
func offsetOrDefault(offsets []float64, index int) float64 {
	if index < len(offsets) {
		return offsets[index]
	}
	return 0.0
}
