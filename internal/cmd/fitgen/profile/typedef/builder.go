// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/builder/shared"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/profile/basetype"
)

const (
	FitBaseType string = "fit_base_type"
)

type typebuilder struct {
	template     *template.Template
	templateExec string

	path       string        // path to generate the file
	sdkVersion string        // Fit SDK Version
	types      []parser.Type // type parsed from profile.xlsx
}

func NewBuilder(path, sdkVersion string, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &typebuilder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(filepath.Join(cd, "typedef.tmpl"), "builder/shared/constant.tmpl")),
		templateExec: "typedef",
		path:         filepath.Join(path, "profile", "typedef"),
		sdkVersion:   sdkVersion,
		types:        types,
	}
}

func (b *typebuilder) Build() ([]builder.Data, error) {
	dataBuilders := make([]builder.Data, 0, len(b.types))
	for _, t := range b.types {
		if t.Name == FitBaseType {
			continue // ignore, manual creation, see [typedefs/basetype] package.
		}
		var hasMfgRangeMin, hashasMfgRangeMax bool
		data := shared.ConstantData{
			Package:      "typedef",
			SDKVersion:   b.sdkVersion,
			Imports:      []string{"strconv"},
			Type:         strutil.ToTitle(t.Name),
			Base:         basetype.FromString(t.BaseType).GoType(),
			StringerMode: shared.StringerMap,
		}

		duplicates := make(map[string][]shared.Constant)
		data.Constants = make([]shared.Constant, 0, len(t.Values))
		for _, v := range t.Values {
			c := shared.Constant{
				Name:    strutil.ToLetterPrefix(strutil.ToTitle(t.Name) + strutil.ToTitle(v.Name)),
				Type:    data.Type,
				Op:      "=",
				Value:   v.Value,
				String:  v.Name,
				Comment: v.Comment,
			}

			if v.Name == "mfg_range_min" {
				hasMfgRangeMin = true
			}

			if v.Name == "mfg_range_max" {
				hashasMfgRangeMax = true
			}

			duplicates[c.Value] = append(duplicates[c.Value], c)
			data.Constants = append(data.Constants, c)

			if hasMfgRangeMin && hashasMfgRangeMax {
				data.AllowRegister = true
			}
		}

		data.Constants = append(data.Constants, shared.Constant{
			Name:    strutil.ToLetterPrefix(strutil.ToTitle(t.Name) + "Invalid"),
			Type:    data.Type,
			Op:      "=",
			Value:   fmt.Sprintf("%#X", basetype.FromString(t.BaseType).Invalid()),
			String:  "invalid",
			Comment: "INVALID",
		})

		// handling duplicate values caused by deprecated
		for cvalue, constant := range duplicates {
			if len(constant) == 1 {
				continue
			}

			for i := range data.Constants {
				if data.Constants[i].Value != cvalue {
					continue
				}

				comment := strings.ToLower(data.Constants[i].Comment)
				if strings.Contains(comment, "deprecated") {
					data.Constants[i].Name = "// " + data.Constants[i].Name
					data.Constants[i].Comment = "[DUPLICATE!] " + data.Constants[i].Comment
				}
			}
		}

		dataBuilders = append(dataBuilders, builder.Data{
			Template:     b.template,
			TemplateExec: b.templateExec,
			Path:         b.path,
			Filename:     strutil.ToSnake(t.Name) + "_gen.go",
			Data:         data,
		})
	}

	return dataBuilders, nil
}
