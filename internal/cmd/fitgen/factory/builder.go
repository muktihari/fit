// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/lookup"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
)

type ( // type aliasing for better code reading.
	MessageName = string
	FieldName   = string
)

type factoryBuilder struct {
	template *template.Template

	mesgnumPackageName string
	profilePackageName string

	path string // path to generate the file

	lookup   *lookup.Lookup
	messages []parser.Message // messages parsed from profile.xlsx
	types    []parser.Type
}

func NewBuilder(path string, lookup *lookup.Lookup, types []parser.Type, messages []parser.Message) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	f := &factoryBuilder{
		template:           template.Must(template.New("main").ParseFiles(filepath.Join(cd, "factory.tmpl"))),
		path:               filepath.Join(path, "factory"),
		mesgnumPackageName: "typedef",
		profilePackageName: "profile",
		types:              types,
		messages:           messages,
		lookup:             lookup,
	}
	f.preproccessMessageField()
	return f
}

func (b *factoryBuilder) preproccessMessageField() {
	// Prepare lookup table for field indexes
	fieldIndexMapByMessageNameByFieldName := make(map[MessageName]map[FieldName]int)
	for _, message := range b.messages {
		fieldIndexMapByMessageNameByFieldName[message.Name] = make(map[FieldName]int)
		for i, field := range message.Fields {
			fieldIndexMapByMessageNameByFieldName[message.Name][field.Name] = i
		}
	}

	// NOTE: This is only a deduction since I can't find the proper explanation in the official documentation anywhere.
	// However, based on the example provided in the Official SDK, this seems to be the most sensible approach.
	//
	// Updating field's accumulate based on component ref:
	// When a field is being referred by components, accumulate value of that field is updated according to that component accumulate value.
	// For example "event_timestamp" accumulate value is false but it's is being referred as a component of "event_timestamp_12"
	// and that component accumulate is true so "event_timestamp" accumulate becomes true.
	for messageIndex, message := range b.messages {
		for _, field := range message.Fields {
			for i, fieldNameRef := range field.Components {
				indexFieldRef := fieldIndexMapByMessageNameByFieldName[message.Name][fieldNameRef]
				b.messages[messageIndex].Fields[indexFieldRef].Accumulate = []bool{accumulateOrDefault(field.Accumulate, i)}
			}
		}
	}
}

func (b *factoryBuilder) Build() ([]builder.Data, error) {
	// Create structure of []proto.Message as string using strings.Builder{},
	// This way, we don't depend on generated value such as types and profile package to be able to generate factory.
	// And also we don't need to process the data in the template which is a bit painful for complex data structure.

	strbuf := new(strings.Builder)
	strbuf.WriteString("[...]message{\n")
	for _, message := range b.messages {
		strbuf.WriteString(b.transformMesgnum(message.Name) + ": {\n") // indexed to create fixed-size slice.
		strbuf.WriteString(fmt.Sprintf("Num: %s, /* %s */\n", b.transformMesgnum(message.Name), message.Name))
		strbuf.WriteString(fmt.Sprintf("Fields: %s,\n", b.makeFields(message)))
		strbuf.WriteString("},\n")
	}

	strbuf.WriteString("}")

	mesgnums := make([]string, 0, len(b.messages))
	for i := range b.messages {
		mesgnums = append(mesgnums, b.transformMesgnum(b.messages[i].Name))
	}

	return []builder.Data{
		{
			Template:     b.template,
			TemplateExec: "factory",
			Path:         b.path,
			Filename:     "factory_gen.go",
			Data: Data{
				Package:  "factory",
				Messages: strbuf.String(),
				Mesgnums: mesgnums,
			},
		},
		{
			Template:     b.template,
			TemplateExec: "exported",
			Path:         b.path,
			Filename:     "exported_gen.go",
			Data: Data{
				Package: "factory",
			},
		},
	}, nil
}

func (b *factoryBuilder) makeFields(message parser.Message) string {
	if len(message.Fields) == 0 {
		return "nil"
	}

	strbuf := new(strings.Builder)
	strbuf.WriteString("[256]proto.Field{\n")
	for _, field := range message.Fields {
		strbuf.WriteString(fmt.Sprintf("%d: {\n", field.Num))
		strbuf.WriteString("FieldBase: &proto.FieldBase{\n")
		strbuf.WriteString(fmt.Sprintf("Name: %q,\n", field.Name))
		strbuf.WriteString(fmt.Sprintf("Num: %d,\n", field.Num))
		strbuf.WriteString(fmt.Sprintf("Type: %s,\n", b.transformProfileType(field.Type)))
		strbuf.WriteString(fmt.Sprintf("BaseType: %s, /* (size: %d) */\n",
			b.transformBaseType(field.Type),
			b.lookup.BaseType(field.Type).Size(),
		))
		strbuf.WriteString(fmt.Sprintf("Array: %t, %s\n", field.Array != "", makeArrayComment(field.Array)))
		strbuf.WriteString(fmt.Sprintf("Components: %s,\n", b.makeComponents(field, message.Name)))
		strbuf.WriteString(fmt.Sprintf("Scale: %g,\n", scaleOrDefault(field.Scales, 0)))    // first index or default
		strbuf.WriteString(fmt.Sprintf("Offset: %g,\n", offsetOrDefault(field.Offsets, 0))) // first index or default
		strbuf.WriteString(fmt.Sprintf("Units: %q,\n", field.Units))
		strbuf.WriteString(fmt.Sprintf("Accumulate: %t,\n", accumulateOrDefault(field.Accumulate, 0)))
		strbuf.WriteString(fmt.Sprintf("SubFields: %s,\n", b.makeSubFields(field, message.Name)))
		strbuf.WriteString("},\n")
		strbuf.WriteString(fmt.Sprintf("Value: %s, /* Default Value: Invalid */\n", b.invalidValueOf(field.Type, field.Array)))
		strbuf.WriteString("},\n")
	}
	strbuf.WriteString("}")

	return strbuf.String()
}

func (b *factoryBuilder) makeComponents(compField parser.ComponentField, messageName string) string {
	if len(compField.GetComponents()) == 0 {
		return "nil"
	}

	strbuf := new(strings.Builder)
	strbuf.WriteString("[]proto.Component{\n")
	for i, fieldNameRef := range compField.GetComponents() {
		fieldRef := b.lookup.FieldByName(messageName, fieldNameRef)
		strbuf.WriteString("{")
		strbuf.WriteString(fmt.Sprintf("FieldNum: %d, /* %s */", fieldRef.Num, fieldRef.Name))
		strbuf.WriteString(fmt.Sprintf("Scale: %g,", scaleOrDefault(compField.GetScales(), i)))               // component index or default
		strbuf.WriteString(fmt.Sprintf("Offset: %g,", offsetOrDefault(compField.GetOffsets(), i)))            // component index or default
		strbuf.WriteString(fmt.Sprintf("Accumulate: %t,", accumulateOrDefault(compField.GetAccumulate(), i))) // component index or default
		strbuf.WriteString(fmt.Sprintf("Bits: %d,", bitsOrDefault(compField.GetBits(), i)))                   // component index or default
		strbuf.WriteString("},\n")
	}
	strbuf.WriteString("}")

	return strbuf.String()
}

func (b *factoryBuilder) makeSubFields(field parser.Field, messageName string) string {
	if len(field.SubFields) == 0 {
		return "nil"
	}

	strbuf := new(strings.Builder)
	strbuf.WriteString("[]proto.SubField{\n")
	for _, subField := range field.SubFields {
		strbuf.WriteString("{")
		strbuf.WriteString(fmt.Sprintf("Name: %q,", subField.Name))
		strbuf.WriteString(fmt.Sprintf("Type: %s, /* %s */", b.transformProfileType(subField.Type), b.transformBaseType(subField.Type)))
		strbuf.WriteString(fmt.Sprintf("Scale: %g,", scaleOrDefault(subField.Scales, 0)))    // first index or default
		strbuf.WriteString(fmt.Sprintf("Offset: %g,", offsetOrDefault(subField.Offsets, 0))) // first index or default
		strbuf.WriteString(fmt.Sprintf("Units: %q,\n", subField.Units))
		strbuf.WriteString(fmt.Sprintf("Components: %s,\n", b.makeComponents(subField, messageName)))
		strbuf.WriteString(fmt.Sprintf("Maps: %s,\n", b.makeSubFieldMaps(subField, messageName)))
		strbuf.WriteString("},\n")
	}
	strbuf.WriteString("}")

	return strbuf.String()
}

func (b *factoryBuilder) makeSubFieldMaps(subfield parser.SubField, messageName string) string {
	if len(subfield.RefFieldNames) == 0 {
		return "nil"
	}

	strbuf := new(strings.Builder)
	strbuf.WriteString("[]proto.SubFieldMap{\n")
	for i, refValueName := range subfield.RefFieldNames {
		fieldRef := b.lookup.FieldByName(messageName, refValueName)
		strbuf.WriteString("{")
		strbuf.WriteString(fmt.Sprintf("RefFieldNum: %d /* %s */,", fieldRef.Num, fieldRef.Name))

		typeValue := b.lookup.TypeValue(fieldRef.Type, subfield.RefFieldValue[i])
		strbuf.WriteString(fmt.Sprintf("RefFieldValue: %s /* %s */,", typeValue, subfield.RefFieldValue[i]))
		strbuf.WriteString("},\n")
	}
	strbuf.WriteString("}")
	return strbuf.String()
}

func (b *factoryBuilder) transformProfileType(fieldType string) string {
	return b.profilePackageName + "." + strutil.ToTitle(fieldType) // profile.Uint8
}

func (b *factoryBuilder) transformBaseType(fieldType string) string {
	baseType := b.lookup.BaseType(fieldType)
	return "basetype." + strutil.ToTitle(baseType.String()) // basetype.Uint16z
}

func (b *factoryBuilder) transformMesgnum(s string) string {
	return b.mesgnumPackageName + ".MesgNum" + strutil.ToTitle(s) // types.MesgNumFileId
}

var baseTypeReplacer = strings.NewReplacer(
	"Enum", "Uint8",
	"Sint", "Int",
	"Byte", "Uint8",
)

func (b *factoryBuilder) invalidValueOf(fieldType, array string) string {
	if fieldType == "bool" { // Special type, bool does not have basetype
		if array != "" {
			return "proto.SliceBool([]bool(nil))"
		}
		return "proto.Bool(false)"
	}

	var (
		baseType      = strutil.ToTitle(b.lookup.BaseType(fieldType).String())
		protoFuncName = strings.TrimSuffix(baseTypeReplacer.Replace(baseType), "z")
		goType        = b.lookup.GoType(fieldType)
	)

	if array != "" {
		return fmt.Sprintf("proto.Slice%s([]%s(nil))", protoFuncName, goType)
	}

	// Float is a special case since NaN is not comparable, so for example, basetype.Float32Invalid, is not a float,
	// but its representation in integer form. This way we can compare it in its integer form later.
	if baseType == "Float32" {
		return "proto.Float32(math.Float32frombits(basetype.Float32Invalid))" // same as `math.Float32frombits(basetype.Float32Invalid)`
	}

	if baseType == "Float64" {
		return "proto.Float64(math.Float64frombits(basetype.Float64Invalid))" // same as `math.Float64frombits(basetype.Float64Invalid)`
	}

	return fmt.Sprintf("proto.%s(basetype.%sInvalid)", protoFuncName, baseType)

}

func bitsOrDefault(bits []byte, index int) byte {
	if index < len(bits) {
		return bits[index]
	}
	return 0
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

func accumulateOrDefault(accumulates []bool, index int) bool {
	if index < len(accumulates) {
		return accumulates[index]
	}
	return false
}

func makeArrayComment(arr string) string {
	if arr == "" {
		return ""
	}
	return fmt.Sprintf("/* Array Size %s */", arr)
}
