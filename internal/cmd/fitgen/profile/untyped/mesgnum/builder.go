// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgnum

import (
	"fmt"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/builder/shared"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/profile/basetype"
)

type mesgnumbuilder struct {
	template     *template.Template
	templateExec string

	path  string        // path to generate the file
	types []parser.Type // type parsed from profile.xlsx
}

func NewBuilder(path string, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &mesgnumbuilder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(
				filepath.Join(cd, "mesgnum.tmpl"),
				filepath.Join(cd, "..", "..", "..", "builder", "shared", "untyped_constant.tmpl"))),
		templateExec: "mesgnum",
		path:         filepath.Join(path, "profile", "untyped", "mesgnum"),
		types:        types,
	}
}

func (b *mesgnumbuilder) Build() ([]builder.Data, error) {
	dataBuilders := make([]builder.Data, 0, len(b.types))
	for _, t := range b.types {
		if t.Name != "mesg_num" {
			continue
		}

		typeName := strutil.ToTitle(t.Name)
		constants := make([]shared.Constant, 0, len(t.Values)+1) // +Invalid
		for _, v := range t.Values {
			constants = append(constants, shared.Constant{
				Name:    strutil.ToLetterPrefix(strutil.ToTitle(v.Name)),
				Type:    typeName,
				Op:      "=",
				Value:   v.Value,
				String:  v.Name,
				Comment: v.Comment,
			})
		}

		constants = append(constants, shared.Constant{
			Name:    strutil.ToLetterPrefix("Invalid"),
			Type:    typeName,
			Op:      "=",
			Value:   fmt.Sprintf("%#X", basetype.FromString(t.BaseType).Invalid()),
			String:  "invalid",
			Comment: "INVALID",
		})

		dataBuilders = append(dataBuilders, builder.Data{
			Template:     b.template,
			TemplateExec: b.templateExec,
			Path:         b.path,
			Filename:     "mesgnum_gen.go",
			Data: shared.ConstantData{
				Package:   "mesgnum",
				Constants: constants,
			},
		})

		break
	}

	return dataBuilders, nil
}
