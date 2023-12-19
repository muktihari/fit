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

	goTypesByBaseTypes       map[string]string // (k -> v) enum -> byte
	goTypesByProfileTypes    map[string]string // (k -> v) typedef.DateTime -> uint32
	baseTypeMapByProfileType map[string]string // (k -> v) enum -> enum, typedef.DateTime -> uint32
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
		goTypesByBaseTypes:       make(map[string]string),
		goTypesByProfileTypes:    make(map[string]string),
		baseTypeMapByProfileType: make(map[string]string),
	}
}

func (b *mesgdefBuilder) populateLookupData() {
	b.goTypesByBaseTypes["bool"] = "bool"
	b.goTypesByBaseTypes["fit_base_type"] = "basetype.BaseType"
	for _, v := range basetype.List() {
		b.goTypesByBaseTypes[v.String()] = v.GoType()
		b.goTypesByProfileTypes[v.String()] = v.GoType()
		b.baseTypeMapByProfileType[v.String()] = v.String()
	}

	// additional profile type which is not defined in basetype.
	b.types = append(b.types, parser.Type{Name: "bool", BaseType: "bool"})

	for _, _type := range b.types {
		b.goTypesByProfileTypes[_type.Name] = b.goTypesByBaseTypes[_type.BaseType]
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

		var maxFieldNum byte
		fields := make([]Field, 0, len(mesg.Fields))
		for i := range mesg.Fields {
			field := &mesg.Fields[i]

			if field.Num > maxFieldNum {
				maxFieldNum = field.Num
			}

			f := Field{
				Num:            field.Num,
				Name:           strutil.ToTitle(mesg.Fields[i].Name),
				String:         field.Name,
				Type:           b.transformType(field.Type, field.Array),
				TypedValue:     b.transformTypedValue(field.Num, field.Type, field.Array),
				PrimitiveValue: b.transformPrimitiveValue(strutil.ToTitle(field.Name), field.Type, field.Array),
				InvalidValue:   b.invalidValueOf(field.Type, field.Array),
				Comment:        field.Comment,
			}

			f.ComparableValue = b.transformComparableValue(field.Type, field.Array, f.PrimitiveValue)

			if field.Array == "" && b.baseTypeMapByProfileType[field.Type] == "string" {
				f.InvalidValue += fmt.Sprintf("&& %s != \"\"", f.ComparableValue)
			}

			if field.Units != "" {
				f.Comment = fmt.Sprintf("Units: %s; %s", field.Units, field.Comment)
			}

			offset := offsetOrDefault(field.Offsets, 0)
			if offset != 0 {
				f.Comment = fmt.Sprintf("Offset: %g; %s", offset, f.Comment)
			}

			scale := scaleOrDefault(field.Scales, 0)
			if scale != 1 {
				f.Comment = fmt.Sprintf("Scale: %g; %s", scale, f.Comment)
			}

			if strings.HasPrefix(f.Type, "[]") {
				f.Comment = fmt.Sprintf("Array: %s; %s", field.Array, f.Comment)
			}

			fields = append(fields, f)
			if strings.HasPrefix(f.Type, "basetype") || strings.HasPrefix(f.InvalidValue, "basetype") {
				imports["github.com/muktihari/fit/profile/basetype"] = struct{}{}
			}
			if isTypeTime(field.Type) {
				imports["time"] = struct{}{}
				imports["github.com/muktihari/fit/kit/datetime"] = struct{}{}
			}
		}

		data := Data{
			SDKVersion:  b.sdkVersion,
			Package:     "mesgdef",
			Imports:     []string{},
			Name:        strutil.ToTitle(mesg.Name),
			Fields:      fields,
			MaxFieldNum: maxFieldNum + 1,
		}

		for k := range imports {
			data.Imports = append(data.Imports, k)
		}

		dataBuilder.Data = data
		dataBuilders = append(dataBuilders, dataBuilder)
	}

	return dataBuilders, nil
}

func (b *mesgdefBuilder) transformType(fieldType, fieldArray string) string {
	if isTypeTime(fieldType) {
		return "time.Time"
	}

	var _type string
	if v, ok := b.goTypesByBaseTypes[fieldType]; ok {
		_type = v
	} else {
		_type = fmt.Sprintf("typedef.%s", strutil.ToTitle(fieldType))
	}

	if fieldArray == "" {
		return _type
	}

	return fmt.Sprintf("[]%s", _type)
}

func isTypeTime(fieldType string) bool {
	typeName := strutil.ToTitle(fieldType)
	switch typeName {
	case "DateTime", "LocalDateTime":
		return true
	}
	return false
}

func (b *mesgdefBuilder) transformPrimitiveValue(fieldName, fieldType, array string) string {
	if isTypeTime(fieldType) {
		return fmt.Sprintf("datetime.ToUint32(m.%s)", fieldName)
	}

	if !strings.HasSuffix(fieldType, "z") && b.baseTypeMapByProfileType[fieldType] == fieldType {
		return fmt.Sprintf("m.%s", fieldName) // only for primitive go types.
	}

	slicePrefix := ""
	if array != "" {
		slicePrefix = "Slice"
	}

	return fmt.Sprintf("typeconv.To%s%s[%s](m.%s)",
		slicePrefix, strutil.ToTitle(b.baseTypeMapByProfileType[fieldType]), b.goTypesByProfileTypes[fieldType], fieldName)
}

func (b *mesgdefBuilder) transformTypedValue(num byte, fieldType, array string) string {
	if fieldType == "fit_base_type" {
		return fmt.Sprintf("typeconv.ToUint8[basetype.BaseType](vals[%d])", num)
	}

	if isTypeTime(fieldType) {
		return fmt.Sprintf("datetime.ToTime(vals[%d])", num)
	}

	var _type string
	if v, ok := b.goTypesByBaseTypes[fieldType]; ok {
		_type = v
	} else {
		_type = fmt.Sprintf("typedef.%s", strutil.ToTitle(fieldType))
	}

	slicePrefix := ""
	if array != "" {
		slicePrefix = "Slice"
	}

	return fmt.Sprintf("typeconv.To%s%s[%s](vals[%d])",
		slicePrefix, strutil.ToTitle(b.baseTypeMapByProfileType[fieldType]), _type, num)
}

func (b *mesgdefBuilder) transformComparableValue(fieldType, array, primitiveValue string) string {
	if array == "" {
		switch b.baseTypeMapByProfileType[fieldType] {
		case "float32":
			return fmt.Sprintf("typeconv.ToUint32[uint32](%s)", primitiveValue)
		case "float64":
			return fmt.Sprintf("typeconv.ToUint64[uint64](%s)", primitiveValue)
		}
	}

	return primitiveValue
}

func (b *mesgdefBuilder) invalidValueOf(fieldType, array string) string {
	if fieldType == "bool" {
		return "false"
	}

	if array != "" {
		return "nil"
	}

	if b.baseTypeMapByProfileType[fieldType] == "float32" {
		return "basetype.Uint32Invalid"
	}

	if b.baseTypeMapByProfileType[fieldType] == "float64" {
		return "basetype.Uint64Invalid"
	}

	bt := basetype.FromString(b.baseTypeMapByProfileType[fieldType]).String()
	return fmt.Sprintf("basetype.%sInvalid", strutil.ToTitle(bt))
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
