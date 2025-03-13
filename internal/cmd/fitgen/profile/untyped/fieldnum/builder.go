// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fieldnum

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"slices"

	"github.com/muktihari/fit/internal/cmd/fitgen/generator"
	"github.com/muktihari/fit/internal/cmd/fitgen/lookup"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/internal/cmd/fitgen/shared"
)

type Builder struct {
	template     *template.Template
	templateExec string

	path string // path to generate the file

	lookup   *lookup.Lookup
	messages []parser.Message
	types    []parser.Type
}

var _ generator.Builder = (*Builder)(nil)

func NewBuilder(path string, lookup *lookup.Lookup, message []parser.Message, types []parser.Type) *Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &Builder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(
				filepath.Join(cd, "fieldnum.tmpl"),
				filepath.Join(cd, "..", "..", "..", "shared", "untyped_constant.tmpl"))),
		templateExec: "fieldnum",
		path:         filepath.Join(path, "profile", "untyped", "fieldnum"),
		lookup:       lookup,
		messages:     message,
		types:        types,
	}
}

func (b *Builder) Build() ([]generator.Data, error) {
	messages := slices.Clone(b.messages)
	slices.SortStableFunc(messages, func(x, y parser.Message) int {
		if x.Name < y.Name {
			return -1
		} else if x.Name > y.Name {
			return 1
		}
		return 0
	})

	constants := make([]shared.Constant, 0)
	for _, mesg := range messages {
		for _, field := range mesg.Fields {
			baseType := b.lookup.BaseType(field.Type).String()
			if field.Type == "bool" {
				baseType = fmt.Sprintf("bool | %s", baseType)
			}

			constants = append(constants, shared.Constant{
				Name:    strutil.ToTitle(mesg.Name) + strutil.ToTitle(field.Name),
				Op:      "=",
				Value:   strconv.Itoa(int(field.Num)),
				String:  mesg.Name + ": " + field.Name,
				Comment: createComment(mesg.Name, &field, baseType),
			})
		}
	}

	constants = append(constants, shared.Constant{
		Name:   "Invalid",
		Op:     "=",
		Value:  "255", // max byte
		String: "invalid",
	})

	dataBuilder := generator.Data{
		Template:     b.template,
		TemplateExec: b.templateExec,
		Path:         b.path,
		Filename:     "fieldnum_gen.go",
		Data: shared.ConstantData{
			Package:   "fieldnum",
			Constants: constants,
		},
	}

	return []generator.Data{dataBuilder}, nil
}

func createComment(mesgName string, field *parser.Field, baseType string) string {
	buf := new(strings.Builder)
	buf.WriteString(fmt.Sprintf("[ %s ] [Type: %s, Base: %s",
		strutil.ToTitle(mesgName), strutil.ToTitle(field.Type), baseType))

	if field.Array != "" {
		buf.WriteString(", Array: ")
		buf.WriteString(field.Array)
	}

	scale := scaleOrDefault(field.Scales, 0)
	offset := offsetOrDefault(field.Offsets, 0)
	if scale != 1 || offset != 0 {
		buf.WriteString(", Scale: ")
		buf.WriteString(strconv.FormatFloat(scale, 'g', -1, 64))
		buf.WriteString(", Offset: ")
		buf.WriteString(strconv.FormatFloat(offset, 'g', -1, 64))
	}

	if field.Units != "" {
		buf.WriteString(", Units: ")
		buf.WriteString(field.Units)
	}

	buf.WriteString("]; ")
	buf.WriteString(field.Comment)

	return buf.String()
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
