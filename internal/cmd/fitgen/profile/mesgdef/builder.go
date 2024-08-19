// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/lookup"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/profile/basetype"
	"golang.org/x/exp/slices"
)

type mesgdefBuilder struct {
	template     *template.Template
	templateExec string

	packageName string

	path string // path to generate the file

	lookup   *lookup.Lookup
	messages []parser.Message
	types    []parser.Type
}

func NewBuilder(path string, lookup *lookup.Lookup, message []parser.Message, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &mesgdefBuilder{
		template: template.Must(template.New("main").
			Funcs(template.FuncMap{
				"stringsJoin":   strings.Join,
				"stringReplace": strings.Replace,
				"byteDiv":       func(a, b byte) byte { return a / b },
				"byteAdd":       func(a, b byte) byte { return a + b },
				"byteSub":       func(a, b byte) byte { return a - b },
				"extractExactlyType": func(s string) string {
					if !strings.HasPrefix(s, "[") {
						return s
					}
					return strings.TrimSuffix(strings.Split(s, "]")[1], "]")
				},
			}).
			ParseFiles(filepath.Join(cd, "mesgdef.tmpl"))),
		templateExec: "mesgdef",
		packageName:  "mesgdef",
		path:         filepath.Join(path, "profile", "mesgdef"),
		lookup:       lookup,
		messages:     message,
		types:        types,
	}
}

func (b *mesgdefBuilder) Build() ([]builder.Data, error) {
	dataBuilders := make([]builder.Data, 0, len(b.messages))
	var maxLenFields int
	for _, mesg := range b.messages {
		canExpand, maxFieldExpandNum := b.componentExpansionAbility(&mesg)

		if len(mesg.Fields) > maxLenFields {
			maxLenFields = len(mesg.Fields)
		}
		var (
			maxFieldNum   byte
			dynamicFields []DynamicField
			fields        = make([]Field, 0, len(mesg.Fields))
			imports       = make(map[string]struct{})
		)
		for _, parserField := range mesg.Fields {
			if parserField.Num > maxFieldNum {
				maxFieldNum = parserField.Num
			}

			var fixedArraySize uint8
			if len(parserField.Array) > 1 && parserField.Array[1] != 'N' {
				n := strings.TrimFunc(parserField.Array, func(r rune) bool {
					return r == '[' || r == ']'
				})
				v, err := strconv.ParseInt(n, 10, 8)
				if err != nil {
					return nil, fmt.Errorf("could not parse array size: %w", err)
				}
				fixedArraySize = uint8(v)
			}

			baseType := b.lookup.BaseType(parserField.Type)

			field := Field{
				Num:             parserField.Num,
				Name:            strutil.ToTitle(parserField.Name),
				String:          parserField.Name,
				ProfileType:     parserField.Type,
				BaseType:        baseType.String(),
				BaseTypeInvalid: fmt.Sprintf("basetype.%sInvalid", strutil.ToTitle(baseType.String())),
				Type:            b.transformType(parserField.Type, parserField.Array, fixedArraySize),
				TypedValue:      b.transformTypedValue(parserField.Num, parserField.Type, parserField.Array, fixedArraySize),
				PrimitiveValue:  b.transformPrimitiveValue(strutil.ToTitle(parserField.Name), parserField.Type, parserField.Array),
				ProtoValue:      b.transformToProtoValue(strutil.ToTitle(parserField.Name), parserField.Type, parserField.Array),
				InvalidValue:    b.invalidValueOf(parserField.Type, parserField.Array, fixedArraySize),
				Comment:         parserField.Comment,
				Scale:           1,
				Offset:          0,
				Array:           parserField.Array != "",
				FixedArraySize:  fixedArraySize,
				Units:           parserField.Units,
			}
			field.ComparableValue = b.transformComparableValue(parserField.Type, parserField.Array, field.PrimitiveValue)

			// Special case
			if isTypeTime(parserField.Type) {
				field.IsValidValue = fmt.Sprintf("!m.%s.Before(datetime.Epoch())", field.Name)
			} else {
				field.IsValidValue = fmt.Sprintf("%s != %s", field.ComparableValue, field.InvalidValue)
			}

			if _, ok := canExpand[parserField.Name]; ok {
				field.CanExpand = true
			}

			if parserField.Array == "" && field.BaseType == "string" {
				field.InvalidValue += fmt.Sprintf("&& %s != \"\"", field.ComparableValue)
			}

			if len(parserField.Scales) <= 1 { // Multiple scales only for components
				field.Scale = scaleOrDefault(parserField.Scales, 0)
			}

			if len(parserField.Offsets) <= 1 { // Multiple offsets only for components
				field.Offset = offsetOrDefault(parserField.Offsets, 0)
			}

			if field.FixedArraySize > 0 && (field.Scale != 1 || field.Offset != 0) {
				field.InvalidArrayValueScaled = b.invalidArrayValueScaled(field.FixedArraySize)
			}

			field.Comment = createComment(&field, parserField.Array)

			fields = append(fields, field)

			imports = appendImports(imports, &field, parserField.Type)

			if len(parserField.SubFields) == 0 {
				continue
			}

			dynamicFields = append(dynamicFields, b.createDynamicField(mesg.Name, &field, &parserField))
		}

		// Optimize memory usage by aligning the memory in the struct.
		optimizedFields := append(fields[:0:0], fields...)
		b.simpleMemoryAlignment(optimizedFields)

		data := Data{
			Package:           b.packageName,
			Imports:           []string{},
			Name:              strutil.ToTitle(mesg.Name),
			Fields:            fields,
			OptimizedFields:   optimizedFields,
			DynamicFields:     dynamicFields,
			MaxFieldNum:       maxFieldNum + 1,
			MaxFieldExpandNum: maxFieldExpandNum + 1,
		}

		for k := range imports {
			data.Imports = append(data.Imports, k)
		}

		dataBuilders = append(dataBuilders,
			builder.Data{
				Template:     b.template,
				TemplateExec: b.templateExec,
				Path:         b.path,
				Filename:     strutil.ToSnake(mesg.Name) + "_gen.go",
				Data:         data,
			},
		)
	}

	dataBuilders = append(dataBuilders, builder.Data{
		Template:     b.template,
		TemplateExec: "util",
		Path:         b.path,
		Filename:     fmt.Sprintf("%s_util_gen.go", b.packageName),
		Data: UtilData{
			Package:      b.packageName,
			MaxLenFields: byte(maxLenFields),
		},
	})

	return dataBuilders, nil
}

// componentExpansionAbility checks whether fields or subfields have components that can be expanded.
// If they do, retrieve the largest field's number.
func (b *mesgdefBuilder) componentExpansionAbility(mesg *parser.Message) (canExpand map[string]byte, maxFieldExpandNum byte) {
	canExpand = make(map[string]byte)
	for _, field := range mesg.Fields {
		for _, component := range field.Components {
			ref := b.lookup.FieldByName(mesg.Name, component)
			canExpand[ref.Name] = ref.Num
			if ref.Num > maxFieldExpandNum {
				maxFieldExpandNum = ref.Num
			}
		}
		for _, subfield := range field.SubFields {
			for _, component := range subfield.Components {
				ref := b.lookup.FieldByName(mesg.Name, component)
				canExpand[ref.Name] = ref.Num
				if ref.Num > maxFieldExpandNum {
					maxFieldExpandNum = ref.Num
				}
			}
		}
	}
	return
}

func createComment(field *Field, array string) string {
	buf := new(strings.Builder)

	if strings.HasPrefix(field.Type, "[") {
		buf.WriteString("Array: ")
		buf.WriteString(array)
		buf.WriteString("; ")
	}

	if field.Scale != 1 {
		buf.WriteString("Scale: ")
		buf.WriteString(strconv.FormatFloat(field.Scale, 'g', -1, 64))
		buf.WriteString("; ")
	}

	if field.Offset != 0 {
		buf.WriteString("Offset: ")
		buf.WriteString(strconv.FormatFloat(field.Offset, 'g', -1, 64))
		buf.WriteString("; ")
	}

	if field.Units != "" {
		buf.WriteString("Units: ")
		buf.WriteString(field.Units)
		buf.WriteString("; ")
	}

	buf.WriteString(field.Comment)

	return strings.TrimSuffix(buf.String(), "; ")
}

func appendImports(imports map[string]struct{}, field *Field, profileType string) map[string]struct{} {
	if strings.HasPrefix(field.Type, "basetype") || strings.HasPrefix(field.InvalidValue, "basetype") {
		imports["github.com/muktihari/fit/profile/basetype"] = struct{}{}
	}
	if isTypeTime(profileType) {
		imports["time"] = struct{}{}
		imports["github.com/muktihari/fit/kit/datetime"] = struct{}{}
	}

	if field.Scale != 1 || field.Offset != 0 {
		imports["math"] = struct{}{}
		imports["github.com/muktihari/fit/profile/basetype"] = struct{}{}
	}

	if strings.HasPrefix(field.ComparableValue, "math") {
		imports["math"] = struct{}{}
	}
	if strings.Contains(field.ProtoValue, "unsafe") ||
		strings.Contains(field.TypedValue, "unsafe") ||
		strings.Contains(field.ComparableValue, "unsafe") {
		imports["unsafe"] = struct{}{}
	}
	if field.Units == "semicircles" {
		imports["github.com/muktihari/fit/kit/semicircles"] = struct{}{}
	}
	return imports
}

func (b *mesgdefBuilder) createDynamicField(mesgName string, field *Field, parserField *parser.Field) DynamicField {
	var (
		rawSwitchCases      = make(map[string][]CondValue)
		rawSwitchCasesOrder = make(map[string]int)
		valuesOrder         = make(map[string]map[ReturnValue]int)
	)
	for _, subField := range parserField.SubFields {
		condValue := CondValue{
			ReturnValue: ReturnValue{
				Name:  subField.Name,
				Units: subField.Units,
			},
		}

		scale := scaleOrDefault(subField.Scales, 0)
		offset := offsetOrDefault(subField.Offsets, 0)
		if scale != 1 || offset != 0 {
			condValue.ReturnValue.Value = fmt.Sprintf("(float64(m.%s) * %g) - %g", field.Name, scale, offset)
		} else {
			condValue.ReturnValue.Value = fmt.Sprintf("%s(m.%s)", b.transformType(subField.Type, "", field.FixedArraySize), field.Name)
		}

		for i, refValueName := range subField.RefFieldNames {
			fieldRef := b.lookup.FieldByName(mesgName, refValueName)

			_, ok := rawSwitchCases[fieldRef.Name]
			if !ok {
				rawSwitchCasesOrder[fieldRef.Name] = len(rawSwitchCasesOrder)
				valuesOrder[fieldRef.Name] = make(map[ReturnValue]int)
			}

			valOrder, ok := valuesOrder[fieldRef.Name][condValue.ReturnValue]
			if !ok {
				valOrder = len(rawSwitchCases[fieldRef.Name])
				valuesOrder[fieldRef.Name][condValue.ReturnValue] = valOrder
				rawSwitchCases[fieldRef.Name] = append(rawSwitchCases[fieldRef.Name], condValue)
			}

			condValue = rawSwitchCases[fieldRef.Name][valOrder]
			condValue.Conds = append(condValue.Conds,
				fmt.Sprintf("%s%s",
					b.transformType(fieldRef.Type, fieldRef.Array, field.FixedArraySize), strutil.ToTitle(subField.RefFieldValue[i])))

			rawSwitchCases[fieldRef.Name][valOrder] = condValue
		}
	}

	switchCases := make([]SwitchCase, len(rawSwitchCases))
	for fieldNameRef, i := range rawSwitchCasesOrder {
		switchCases[i] = SwitchCase{
			Name:       fmt.Sprintf("m.%s", strutil.ToTitle(fieldNameRef)),
			CondValues: rawSwitchCases[fieldNameRef],
		}
	}

	return DynamicField{
		Name:        field.Name,
		SwitchCases: switchCases,
		Default: ReturnValue{
			Name:  parserField.Name,
			Units: parserField.Units,
			Value: fmt.Sprintf("m.%s", field.Name),
		},
	}
}

func (b *mesgdefBuilder) simpleMemoryAlignment(fields []Field) {
	// In 64 bits machine, the layout is per 8 bytes.
	// If the size is a multply of 8, set 8.
	for i := range fields {
		if strings.HasPrefix(fields[i].Type, "[]") { // Slice
			fields[i].Size = 24 // pointer to array (8 bytes), len (8 bytes), cap (8 bytes)
		} else if strings.HasPrefix(fields[i].Type, "[") { // Array
			size := basetype.FromString(fields[i].BaseType).Size() * fields[i].FixedArraySize
			rem := size % 8
			if rem == 0 {
				fields[i].Size = size
			} else if rem%4 == 0 {
				fields[i].Size = 4
			} else if rem%2 == 0 {
				fields[i].Size = 2
			} else {
				fields[i].Size = 1
			}
		} else if isTypeTime(fields[i].ProfileType) { // time.Time
			fields[i].Size = 24 // wall (8 bytes), ext (8 bytes), *loc (8 bytes)
		} else if fields[i].BaseType == "string" {
			fields[i].Size = 16 // pointer to array (8 bytes) + len (8 bytes)
		} else { // Everything else, 8 bytes or lower.
			fields[i].Size = basetype.FromString(fields[i].BaseType).Size()
		}
	}
	// Order by the size desc.
	slices.SortStableFunc(fields, func(a, b Field) int {
		if a.Size%8 == 0 && b.Size%8 == 0 {
			return 0
		}
		if a.Size > b.Size {
			return -1
		} else if a.Size < b.Size {
			return 1
		}
		if a.Array && b.Array {
			return 0
		} else if a.Array {
			return -1
		} else if b.Array {
			return 1
		}
		return 0
	})
}

func (b *mesgdefBuilder) transformType(fieldType, fieldArray string, fixedArraySize byte) string {
	if isTypeTime(fieldType) {
		return "time.Time"
	}

	if fieldType == "fit_base_type" {
		return "basetype.BaseType"
	}

	var typ string
	if v := b.lookup.BaseType(fieldType).String(); v == fieldType || fieldType == "bool" {
		typ = b.lookup.GoType(fieldType)
	} else {
		typ = fmt.Sprintf("typedef.%s", strutil.ToTitle(fieldType))
	}

	if fieldArray == "" {
		return typ
	}

	if fixedArraySize > 0 {
		return fmt.Sprintf("[%d]%s", fixedArraySize, typ)
	}

	return fmt.Sprintf("[]%s", typ)
}

func isTypeTime(fieldType string) bool {
	typeName := strutil.ToTitle(fieldType)
	switch typeName {
	case "DateTime", "LocalDateTime":
		return true
	}
	return false
}

var baseTypeReplacer = strings.NewReplacer(
	"Enum", "Uint8",
	"Sint", "Int",
	"Byte", "Uint8",
)

func (b *mesgdefBuilder) transformToProtoValue(fieldName, fieldType, array string) string {
	if isTypeTime(fieldType) {
		return fmt.Sprintf("proto.Uint32(uint32(m.%s.Sub(datetime.Epoch()).Seconds()))", fieldName)
	}

	baseType := b.lookup.BaseType(fieldType).String()
	goType := b.lookup.GoType(fieldType)

	typ := strutil.ToTitle(goType)
	typ = baseTypeReplacer.Replace(typ)
	typ = strings.TrimSuffix(typ, "z")

	if array != "" {
		return fmt.Sprintf("proto.Slice%s(%s)", typ, fmt.Sprintf("m.%s", fieldName))
	}

	val := fmt.Sprintf("m.%s", fieldName)

	if fieldType != "bool" && baseType != fieldType {
		val = fmt.Sprintf("%s(%s)", b.lookup.GoType(fieldType), val)
	}

	return fmt.Sprintf("proto.%s(%s)", typ, val)
}

func (b *mesgdefBuilder) transformPrimitiveValue(fieldName, fieldType, array string) string {
	if isTypeTime(fieldType) {
		return fmt.Sprintf("datetime.ToUint32(m.%s)", fieldName)
	}

	if !strings.HasSuffix(fieldType, "z") &&
		b.lookup.BaseType(fieldType).String() == fieldType {
		return fmt.Sprintf("m.%s", fieldName) // only for primitive go types.
	}

	goType := b.lookup.GoType(fieldType)
	if goType == "bool" {
		return fmt.Sprintf("m.%s", fieldName)
	}

	if array == "" {
		return fmt.Sprintf("%s(m.%s)", goType, fieldName)
	}

	return fmt.Sprintf("m.%s", fieldName)

}

func (b *mesgdefBuilder) transformTypedValue(num byte, fieldType, array string, fixedArraySize uint8) string {
	if isTypeTime(fieldType) {
		return fmt.Sprintf("datetime.ToTime(vals[%d].Uint32())", num)
	}

	baseType := b.lookup.BaseType(fieldType).String()
	baseTypeTitleCase := strutil.ToTitle(baseType)
	typ := baseTypeReplacer.Replace(baseTypeTitleCase)

	if fieldType == "fit_base_type" {
		return fmt.Sprintf("basetype.BaseType((vals[%d]).%s())", num, typ)
	}

	if array != "" && strings.HasSuffix(typ, "z") {
		typ = strings.TrimSuffix(typ, "z")
	}

	if fieldType == "bool" {
		typ = "Bool"
	}

	var value string
	if array == "" {
		value = fmt.Sprintf("vals[%d].%s()", num, typ)
	} else if fixedArraySize > 0 { // Array
		goTyp := strings.ToLower(strings.TrimSuffix(typ, "z"))
		value = fmt.Sprintf(`func() (arr [%d]%s) {
			arr = [%d]%s{
				%s
			}
			copy(arr[:], %s)
			return arr
		}()`, fixedArraySize, goTyp,
			fixedArraySize, goTyp,
			strings.Repeat(
				fmt.Sprintf("basetype.%sInvalid,\n", baseTypeTitleCase),
				int(fixedArraySize),
			),
			fmt.Sprintf("vals[%d].Slice%s()", num, typ),
		)
	} else {
		value = fmt.Sprintf("vals[%d].Slice%s()", num, typ)
	}

	if baseType == fieldType || fieldType == "bool" { // primitive-types
		return value
	}

	typdef := fmt.Sprintf("typedef.%s", strutil.ToTitle(fieldType))

	if array == "" {
		return fmt.Sprintf("%s(%s)", typdef, value)
	}

	return fmt.Sprintf(`func() []%s {
			sliceValue := %s
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*%s)(ptr), len(sliceValue))
		}()`, typdef, value, typdef)
}

func (b *mesgdefBuilder) transformComparableValue(fieldType, array, primitiveValue string) string {
	if array == "" {
		switch b.lookup.BaseType(fieldType).String() {
		case "float32":
			return fmt.Sprintf("math.Float32bits(%s)", primitiveValue)
		case "float64":
			return fmt.Sprintf("math.Float32bits(%s)", primitiveValue)
		}
	}

	return primitiveValue
}

func (b *mesgdefBuilder) invalidValueOf(fieldType, array string, fixedArraySize byte) string {
	if fieldType == "bool" {
		return "false"
	}

	if array != "" {
		if fixedArraySize == 0 { // Slice
			return "nil"
		}

		baseType := b.lookup.BaseType(fieldType).String()
		baseTypeTitleCase := strutil.ToTitle(baseType)
		typ := baseTypeReplacer.Replace(baseTypeTitleCase)
		typ = strings.TrimSuffix(typ, "z")
		goTyp := strings.ToLower(typ)

		return fmt.Sprintf(`[%d]%s{
			%s
		}`, fixedArraySize, goTyp,
			strings.Repeat(
				fmt.Sprintf("basetype.%sInvalid,\n", baseTypeTitleCase),
				int(fixedArraySize),
			),
		)
	}

	return fmt.Sprintf("basetype.%sInvalid",
		strutil.ToTitle(b.lookup.BaseType(fieldType).String()))
}

func (b *mesgdefBuilder) invalidArrayValueScaled(fixedArraySize byte) string {
	return fmt.Sprintf(`[%d]float64{
		%s
	}`,
		fixedArraySize,
		strings.Repeat(
			"math.Float64frombits(basetype.Float64Invalid),\n",
			int(fixedArraySize),
		),
	)
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
