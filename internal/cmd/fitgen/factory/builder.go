// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

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

type ( // type aliasing for better code reading.
	// types
	ValueName = string
	BaseType  = string

	// messages
	MessageName = string
	FieldName   = string
	FieldNum    = byte

	// profile
	ProfileType = string
)

type factoryBuilder struct {
	template *template.Template

	mesgnumPackageName string
	profilePackageName string

	path       string           // path to generate the file
	sdkVersion string           // Fit SDK Version
	messages   []parser.Message // messages parsed from profile.xlsx
	types      []parser.Type

	once sync.Once

	// types lookup
	goTypesByProfileTypes            map[ProfileType]string                     // (k -> v) typedef.DateTime -> uint32
	baseTypeMapByProfileType         map[ProfileType]BaseType                   // e.g. mesg_num 				 -> uint16 ,  data_time -> uint32
	valueMapByProfileTypeByValueName map[ProfileType]map[ValueName]parser.Value // e.g. map[mesg_num][file_id] -> Value{}

	// message-field lookup
	fieldMapByMessageNameByFieldNum  map[MessageName]map[FieldNum]parser.Field  // e.g. map[file_id][0] 					 -> Field{}
	fieldMapByMessageNameByFieldName map[MessageName]map[FieldName]parser.Field // e.g. map[file_createor][software_version] -> Field{}
}

func NewBuilder(path, sdkVersion string, types []parser.Type, messages []parser.Message) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &factoryBuilder{
		template:                         template.Must(template.New("main").ParseFiles(filepath.Join(cd, "factory.tmpl"))),
		path:                             filepath.Join(path, "factory"),
		mesgnumPackageName:               "typedef",
		profilePackageName:               "profile",
		sdkVersion:                       sdkVersion,
		types:                            types,
		messages:                         messages,
		goTypesByProfileTypes:            make(map[ProfileType]string),
		baseTypeMapByProfileType:         make(map[ProfileType]BaseType),
		valueMapByProfileTypeByValueName: make(map[ProfileType]map[ValueName]parser.Value),
	}
}

func (b *factoryBuilder) populateLookupData() {
	goTypesByBaseTypes := map[BaseType]string{
		"bool":          "bool",
		"fit_base_type": "basetype.BaseType",
	}

	for _, v := range basetype.List() { // map to itself
		goTypesByBaseTypes[v.String()] = v.GoType()
		b.goTypesByProfileTypes[v.String()] = v.GoType()
		b.baseTypeMapByProfileType[v.String()] = v.String()
	}

	// additional profile type which is not defined in basetype.
	b.types = append(b.types, parser.Type{Name: "bool", BaseType: "bool"})

	for _, _type := range b.types {
		b.goTypesByProfileTypes[_type.Name] = goTypesByBaseTypes[_type.BaseType]
		b.baseTypeMapByProfileType[_type.Name] = _type.BaseType
		b.valueMapByProfileTypeByValueName[_type.Name] = make(map[ValueName]parser.Value)
		for _, value := range _type.Values {
			b.valueMapByProfileTypeByValueName[_type.Name][value.Name] = value
		}
	}

	b.fieldMapByMessageNameByFieldNum = make(map[MessageName]map[FieldNum]parser.Field)
	b.fieldMapByMessageNameByFieldName = make(map[MessageName]map[FieldName]parser.Field)
	fieldIndexMapByMessageNameByFieldName := make(map[MessageName]map[FieldName]int)

	for _, message := range b.messages {
		b.fieldMapByMessageNameByFieldNum[message.Name] = make(map[FieldNum]parser.Field)
		b.fieldMapByMessageNameByFieldName[message.Name] = make(map[FieldName]parser.Field)

		fieldIndexMapByMessageNameByFieldName[message.Name] = make(map[FieldName]int)
		for i, field := range message.Fields {
			b.fieldMapByMessageNameByFieldNum[message.Name][field.Num] = field
			b.fieldMapByMessageNameByFieldName[message.Name][field.Name] = field
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
	b.once.Do(func() { b.populateLookupData() })

	// Create structure of []proto.Message as string using strings.Builder{},
	// This way, we don't depend on generated value such as types and profile package to be able to generate factory.
	// And also we don't need to process the data in the template which is a bit painful for complex data structure.

	strbuf := new(strings.Builder)
	strbuf.WriteString("[...]proto.Message{\n")
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
				Package:    "factory",
				SDKVersion: b.sdkVersion,
				Messages:   strbuf.String(),
				Mesgnums:   mesgnums,
			},
		},
		{
			Template:     b.template,
			TemplateExec: "exported",
			Path:         b.path,
			Filename:     "exported_gen.go",
			Data: Data{
				Package:    "factory",
				SDKVersion: b.sdkVersion,
			},
		},
	}, nil
}

func (b *factoryBuilder) makeFields(message parser.Message) string {
	if len(message.Fields) == 0 {
		return "nil"
	}

	strbuf := new(strings.Builder)
	strbuf.WriteString("[]proto.Field{\n")
	for _, field := range message.Fields {
		strbuf.WriteString("{\n")
		strbuf.WriteString("FieldBase: &proto.FieldBase{\n")
		strbuf.WriteString(fmt.Sprintf("Name: %q,\n", field.Name))
		strbuf.WriteString(fmt.Sprintf("Num: %d,\n", field.Num))
		strbuf.WriteString(fmt.Sprintf("Type: %s, /* %s */\n",
			b.transformProfileType(field.Type),
			b.transformBaseType(field.Type),
		))
		strbuf.WriteString(fmt.Sprintf("Array: %t, %s\n", field.Array != "", makeArrayComment(field.Array)))
		strbuf.WriteString(fmt.Sprintf("Size: %d,\n",
			basetype.FromString(b.baseTypeMapByProfileType[field.Type]).Size(),
		))
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
		fieldRef := b.fieldMapByMessageNameByFieldName[messageName][fieldNameRef]
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
		fieldRef := b.fieldMapByMessageNameByFieldName[messageName][refValueName]
		strbuf.WriteString("{")
		strbuf.WriteString(fmt.Sprintf("RefFieldNum: %d /* %s */,", fieldRef.Num, fieldRef.Name))

		typeRef := b.valueMapByProfileTypeByValueName[fieldRef.Type][subfield.RefFieldValue[i]]
		strbuf.WriteString(fmt.Sprintf("RefFieldValue: %s /* %s */,", typeRef.Value, typeRef.Name))
		strbuf.WriteString("},\n")
	}
	strbuf.WriteString("}")
	return strbuf.String()
}

func (b *factoryBuilder) transformProfileType(fieldType string) string {
	return b.profilePackageName + "." + strutil.ToTitle(fieldType) // profile.Uint8
}

func (b *factoryBuilder) transformBaseType(fieldType string) string {
	return "basetype." + strutil.ToTitle(b.baseTypeMapByProfileType[fieldType]) // basetype.Uint16z
}

func (b *factoryBuilder) transformMesgnum(s string) string {
	return b.mesgnumPackageName + ".MesgNum" + strutil.ToTitle(s) // types.MesgNumFileId
}

func (b *factoryBuilder) invalidValueOf(fieldType, array string) string {
	if fieldType == "bool" {
		return "false"
	}

	if array != "" {
		goType := b.goTypesByProfileTypes[fieldType]
		return fmt.Sprintf("[]%s(nil)", goType)
	}

	// Float is a special case since NaN is not comparable, so for example, basetype.Float32Invalid, is not a float,
	// but its representation in integer form. This way we can compare it in its integer form later.
	if b.baseTypeMapByProfileType[fieldType] == "float32" {
		return "basetype.Float32.Invalid()" // same as `math.Float32frombits(basetype.Float32Invalid)`
	}

	if b.baseTypeMapByProfileType[fieldType] == "float64" {
		return "basetype.Float64.Invalid()" // same as `math.Float64frombits(basetype.Float64Invalid)`
	}

	bt := basetype.FromString(b.baseTypeMapByProfileType[fieldType]).String()
	return fmt.Sprintf("basetype.%sInvalid", strutil.ToTitle(bt))
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
