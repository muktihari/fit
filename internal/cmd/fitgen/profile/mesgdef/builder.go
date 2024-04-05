// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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
	"golang.org/x/exp/slices"
)

type mesgdefBuilder struct {
	template     *template.Template
	templateExec string

	path string // path to generate the file

	once sync.Once

	messages []parser.Message
	types    []parser.Type

	goTypesByBaseTypes       map[string]string // (k -> v) enum -> byte
	goTypesByProfileTypes    map[string]string // (k -> v) typedef.DateTime -> uint32
	baseTypeMapByProfileType map[string]string // (k -> v) enum -> enum, typedef.DateTime -> uint32

	fieldByMesgNameByFieldName map[string]map[string]parser.Field
}

func NewBuilder(path string, message []parser.Message, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &mesgdefBuilder{
		template: template.Must(template.New("main").
			ParseFiles(filepath.Join(cd, "mesgdef.tmpl"))),
		templateExec:               "mesgdef",
		path:                       filepath.Join(path, "profile", "mesgdef"),
		messages:                   message,
		types:                      types,
		goTypesByBaseTypes:         make(map[string]string),
		goTypesByProfileTypes:      make(map[string]string),
		baseTypeMapByProfileType:   make(map[string]string),
		fieldByMesgNameByFieldName: make(map[string]map[string]parser.Field),
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

	for _, mesg := range b.messages {
		b.fieldByMesgNameByFieldName[mesg.Name] = make(map[string]parser.Field)
		for _, field := range mesg.Fields {
			b.fieldByMesgNameByFieldName[mesg.Name][field.Name] = field
		}
	}
}

func (b *mesgdefBuilder) Build() ([]builder.Data, error) {
	b.once.Do(func() { b.populateLookupData() })

	dataBuilders := make([]builder.Data, 0, len(b.messages))
	for _, mesg := range b.messages {
		canExpandMap := map[string]byte{}
		var maxFieldExpandNum byte
		for _, field := range mesg.Fields {
			for _, component := range field.Components {
				ref := b.fieldByMesgNameByFieldName[mesg.Name][component]
				canExpandMap[ref.Name] = ref.Num
				if ref.Num > maxFieldExpandNum {
					maxFieldExpandNum = ref.Num
				}
			}
			for _, subfield := range field.SubFields {
				for _, component := range subfield.Components {
					ref := b.fieldByMesgNameByFieldName[mesg.Name][component]
					canExpandMap[ref.Name] = ref.Num
					if ref.Num > maxFieldExpandNum {
						maxFieldExpandNum = ref.Num
					}
				}
			}
		}

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
				ProfileType:    field.Type,
				BaseType:       b.baseTypeMapByProfileType[field.Type],
				Type:           b.transformType(field.Type, field.Array),
				TypedValue:     b.transformTypedValue(field.Num, field.Type, field.Array),
				PrimitiveValue: b.transformPrimitiveValue(strutil.ToTitle(field.Name), field.Type, field.Array),
				ProtoValue:     b.transformToProtoValue(strutil.ToTitle(field.Name), field.Type, field.Array),
				InvalidValue:   b.invalidValueOf(field.Type, field.Array),
				Comment:        field.Comment,
				Array:          field.Array != "",
			}

			if _, ok := canExpandMap[field.Name]; ok {
				f.CanExpand = true
			}

			f.ComparableValue = b.transformComparableValue(field.Type, field.Array, f.PrimitiveValue)

			if field.Array == "" && b.baseTypeMapByProfileType[field.Type] == "string" {
				f.InvalidValue += fmt.Sprintf("&& %s != \"\"", f.ComparableValue)
			}

			if field.Units != "" {
				f.Comment = fmt.Sprintf("Units: %s; %s", field.Units, field.Comment)
			}

			if len(field.Offsets) > 1 { // Multiple offsets only for components
				f.Offset = 0
			} else {
				f.Offset = offsetOrDefault(field.Offsets, 0)
				if f.Offset != 0 {
					f.Comment = fmt.Sprintf("Offset: %g; %s", f.Offset, f.Comment)
				}
			}

			if len(field.Scales) > 1 { // Multiple scales only for components
				f.Scale = 1
			} else {
				f.Scale = scaleOrDefault(field.Scales, 0)
				if f.Scale != 1 {
					f.Comment = fmt.Sprintf("Scale: %g; %s", f.Scale, f.Comment)
				}
			}

			if !(f.Scale == 1 && f.Offset == 0) {
				imports["github.com/muktihari/fit/kit/scaleoffset"] = struct{}{}
			}

			if strings.HasPrefix(f.Type, "[]") {
				f.Comment = fmt.Sprintf("Array: %s; %s", field.Array, f.Comment)
			}

			f.Comment = strings.Trim(f.Comment, "; ")

			fields = append(fields, f)
			if strings.HasPrefix(f.Type, "basetype") || strings.HasPrefix(f.InvalidValue, "basetype") {
				imports["github.com/muktihari/fit/profile/basetype"] = struct{}{}
			}
			if isTypeTime(field.Type) {
				imports["time"] = struct{}{}
				imports["github.com/muktihari/fit/kit/datetime"] = struct{}{}
			}

			if !(f.Scale == 1 && f.Offset == 0) && f.Array != true {
				imports["math"] = struct{}{}
			}

			if strings.HasPrefix(f.ComparableValue, "math") {
				imports["math"] = struct{}{}
			}
			if strings.Contains(f.ProtoValue, "unsafe") ||
				strings.Contains(f.TypedValue, "unsafe") ||
				strings.Contains(f.ComparableValue, "unsafe") {
				imports["unsafe"] = struct{}{}
			}
		}

		// Optimize memory usage by aligning the memory in the struct.
		b.simpleMemoryAlignment(fields)

		data := Data{
			Package:           "mesgdef",
			Imports:           []string{},
			Name:              strutil.ToTitle(mesg.Name),
			Fields:            fields,
			MaxFieldNum:       maxFieldNum + 1,
			MaxFieldExpandNum: maxFieldExpandNum + 1,
		}

		for k := range imports {
			data.Imports = append(data.Imports, k)
		}

		dataBuilder.Data = data
		dataBuilders = append(dataBuilders, dataBuilder)
	}

	return dataBuilders, nil
}

func (b *mesgdefBuilder) simpleMemoryAlignment(fields []Field) {
	// In 64 bits machine, the layout is per 8 bytes.
	// If the size is a multply of 8, set 8.
	for i := range fields {
		if strings.HasPrefix(fields[i].Type, "[]") { // Slice
			fields[i].Size = 8 // 24 bytes (pointer to array (8 bytes), len (8 bytes), cap (8 bytes))
		} else if isTypeTime(fields[i].ProfileType) { // time.Time
			fields[i].Size = 8 // 24 bytes (wall (8 bytes), ext (8 bytes), *loc (8 bytes))
		} else if fields[i].BaseType == "string" {
			fields[i].Size = 8 // 16 bytes (pointer to array (8 bytes) + len (8 bytes))
		} else { // Everything else, 8 bytes or lower.
			fields[i].Size = int(basetype.FromString(fields[i].BaseType).Size())
		}
	}
	// Order by the size desc.
	slices.SortStableFunc(fields, func(a, b Field) int {
		if a.Size > b.Size {
			return -1
		} else if a.Size < b.Size {
			return 1
		}
		return 0
	})
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

func (b *mesgdefBuilder) transformToProtoValue(fieldName, fieldType, array string) string {
	if isTypeTime(fieldType) {
		return fmt.Sprintf("proto.Uint32(datetime.ToUint32(m.%s))", fieldName)
	}

	baseType := strutil.ToTitle(b.baseTypeMapByProfileType[fieldType])
	baseType = baseTypeReplacer.Replace(baseType)

	baseType = strings.TrimSuffix(baseType, "z")

	goType := b.goTypesByProfileTypes[fieldType]

	val := fmt.Sprintf("m.%s", fieldName)
	if b.baseTypeMapByProfileType[fieldType] != fieldType {
		if array == "" {
			val = fmt.Sprintf("%s(%s)", goType, val)
		}
	}

	slicePrefix := ""
	if array != "" {
		slicePrefix = "Slice"
	}

	return fmt.Sprintf("proto.%s%s(%s)", slicePrefix, baseType, val)
}

func (b *mesgdefBuilder) transformPrimitiveValue(fieldName, fieldType, array string) string {
	if isTypeTime(fieldType) {
		return fmt.Sprintf("datetime.ToUint32(m.%s)", fieldName)
	}

	if !strings.HasSuffix(fieldType, "z") && b.baseTypeMapByProfileType[fieldType] == fieldType {
		return fmt.Sprintf("m.%s", fieldName) // only for primitive go types.
	}

	goType := b.goTypesByProfileTypes[fieldType]
	if array == "" && goType != "" {
		return fmt.Sprintf("%s(m.%s)", goType, fieldName)
	}

	if array != "" {
		return fmt.Sprintf("m.%s", fieldName)
	}

	return fmt.Sprintf("%s(m.%s)", goType, fieldName)
}

var baseTypeReplacer = strings.NewReplacer(
	"Enum", "Uint8",
	"Sint", "Int",
	"Byte", "Uint8",
)

func (b *mesgdefBuilder) transformTypedValue(num byte, fieldType, array string) string {
	baseType := strutil.ToTitle(b.baseTypeMapByProfileType[fieldType])
	baseType = baseTypeReplacer.Replace(baseType)
	if array != "" && strings.HasSuffix(baseType, "z") {
		baseType = strings.TrimSuffix(baseType, "z")
	}

	if fieldType == "fit_base_type" {
		return fmt.Sprintf("basetype.BaseType((vals[%d]).%s())", num, baseType)
	}

	if isTypeTime(fieldType) {
		return fmt.Sprintf("datetime.ToTime(vals[%d].Uint32())", num)
	}

	var _type string
	if _, ok := b.goTypesByBaseTypes[fieldType]; !ok {
		_type = fmt.Sprintf("typedef.%s", strutil.ToTitle(fieldType))
	}

	slicePrefix := ""
	if array != "" {
		slicePrefix = "Slice"
	}

	res := fmt.Sprintf("vals[%d].%s%s()",
		num,
		slicePrefix,
		baseType,
	)

	if _type != "" {
		if array != "" {
			res = fmt.Sprintf(`func() []%s {
				sliceValue := %s
				ptr := unsafe.SliceData(sliceValue)
				return unsafe.Slice((*%s)(ptr), len(sliceValue))
			}()`, _type, res, _type)
		} else {
			res = fmt.Sprintf("%s(%s)", _type, res)
		}
	}

	return res
}

func (b *mesgdefBuilder) transformComparableValue(fieldType, array, primitiveValue string) string {
	if array == "" {
		switch b.baseTypeMapByProfileType[fieldType] {
		case "float32":
			return fmt.Sprintf("math.Float32bits(%s)", primitiveValue)
		case "float64":
			return fmt.Sprintf("math.Float32bits(%s)", primitiveValue)
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
