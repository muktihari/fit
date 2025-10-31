// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/generator"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/internal/cmd/fitgen/shared"
	"github.com/muktihari/fit/profile/basetype"
)

const (
	FitBaseType string = "fit_base_type"
)

type Builder struct {
	template     *template.Template
	templateExec string

	path  string        // path to generate the file
	types []parser.Type // type parsed from profile.xlsx
}

var _ generator.Builder = (*Builder)(nil)

func NewBuilder(path string, types []parser.Type) *Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &Builder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(
				filepath.Join(cd, "typedef.tmpl"),
				filepath.Join(cd, "..", "..", "shared", "constant.tmpl"))),
		templateExec: "typedef",
		path:         filepath.Join(path, "profile", "typedef"),
		types:        types,
	}
}

func (b *Builder) Build() ([]generator.Data, error) {
	dataBuilders := make([]generator.Data, 0, len(b.types))
	for _, t := range b.types {
		if t.Name == FitBaseType {
			continue // ignore, manual creation, see [profile/basetype] package.
		}

		var (
			hasMfgRangeMin bool
			hasMfgRangeMax bool
		)

		typeName := strutil.ToTitle(t.Name)

		duplicates := make(map[string]int)
		constants := make([]shared.Constant, 0, len(t.Values))
		for _, v := range t.Values {
			if v.Name == "mfg_range_min" {
				hasMfgRangeMin = true
			}
			if v.Name == "mfg_range_max" {
				hasMfgRangeMax = true
			}

			duplicates[v.Value]++
			constants = append(constants, shared.Constant{
				Name:    strutil.ToLetterPrefix(strutil.ToTitle(t.Name) + strutil.ToTitle(v.Name)),
				Type:    typeName,
				Op:      "=",
				Value:   v.Value,
				String:  v.Name,
				Comment: v.Comment,
			})
		}

		// handling duplicate values caused by deprecated
		for value, count := range duplicates {
			if count == 1 {
				continue
			}
			for i := range constants {
				if constants[i].Value != value {
					continue
				}
				comment := strings.ToLower(constants[i].Comment)
				if strings.Contains(comment, "deprecated") {
					constants[i].Name = "// " + constants[i].Name
					constants[i].Comment = "[DUPLICATE!] " + constants[i].Comment
				}
			}
		}

		var typeComment string
		if strings.HasSuffix(t.BaseType, "z") {
			typeComment = fmt.Sprintf("// Base: %s", t.BaseType)
		}

		dataBuilders = append(dataBuilders, generator.Data{
			Template:     b.template,
			TemplateExec: b.templateExec,
			Path:         b.path,
			Filename:     strutil.ToSnake(t.Name) + "_gen.go",
			Data: shared.ConstantData{
				Package:       "typedef",
				Imports:       []string{"strconv"},
				Type:          typeName,
				Base:          basetype.FromString(t.BaseType).GoType(),
				Comment:       typeComment,
				Constants:     constants,
				AllowRegister: hasMfgRangeMin && hasMfgRangeMax,
				Invalid: shared.Constant{
					Name:  strutil.ToLetterPrefix(strutil.ToTitle(t.Name) + "Invalid"),
					Type:  typeName,
					Op:    "=",
					Value: fmt.Sprintf("%#X", basetype.FromString(t.BaseType).Invalid()),
				},
			},
		})
	}

	return dataBuilders, nil
}
