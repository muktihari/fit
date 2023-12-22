// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"unicode"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/builder/shared"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/profile/basetype"
)

const (
	ProfileType string = "ProfileType"
	FitBaseType string = "fit_base_type"
)

type profilebuilder struct {
	template *template.Template

	path       string        // path to generate the file
	sdkVersion string        // Fit SDK Version
	types      []parser.Type // type parsed from profile.xlsx
}

func NewBuilder(path, sdkVersion string, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &profilebuilder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(
				filepath.Join(cd, "profile.tmpl"),
				filepath.Join(cd, "..", "builder", "shared", "constant.tmpl"))),
		path:       filepath.Join(path, "profile"),
		sdkVersion: sdkVersion,
		types:      types,
	}
}

func (b *profilebuilder) Build() ([]builder.Data, error) {
	profileDataBuilder := b.buildProfile()
	versionDataBuilder := b.buildVersion()

	return []builder.Data{profileDataBuilder, versionDataBuilder}, nil
}

func (b *profilebuilder) buildProfile() builder.Data {
	constants := make([]shared.Constant, 0, len(b.types))
	mappingToBaseTypes := make([]MappingBaseType, 0, len(b.types))

	for _, t := range b.types {
		if t.Name == FitBaseType { // special types to be included, mapping to itself (profile.Uint8 == basetype.Uint8)
			for _, v := range t.Values {
				constantName := strutil.ToTitle(v.Name)
				baseType := transformBaseType(strutil.ToTitle(v.Name))
				constants = append(constants, shared.Constant{
					Name:   constantName,
					String: v.Name,
				})
				mappingToBaseTypes = append(mappingToBaseTypes, MappingBaseType{
					ConstantName: constantName,
					BaseType:     baseType,
				})
			}

			constants = append(constants, shared.Constant{Name: "Bool", String: "bool"})
			mappingToBaseTypes = append(mappingToBaseTypes, MappingBaseType{ConstantName: "Bool", BaseType: transformBaseType("Enum")})
		}
	}

	for _, t := range b.types {
		constants = append(constants, shared.Constant{
			Name:   strutil.ToTitle(t.Name),
			String: t.Name,
		})
		mappingToBaseTypes = append(mappingToBaseTypes, MappingBaseType{
			ConstantName: strutil.ToTitle(t.Name),
			BaseType:     transformBaseType(strutil.ToTitle(t.BaseType)),
		})
	}

	if len(constants) > 0 {
		constants[0].Type = ProfileType
		constants[0].Op = "="
		constants[0].Value = "iota"
	}

	data := shared.ConstantData{
		Package: "profile",
		Imports: []string{
			"strconv",
			"github.com/muktihari/fit/profile/basetype",
		},
		SDKVersion: b.sdkVersion,
		Type:       ProfileType,
		Base:       "uint16",
		Constants:  constants,
	}

	data.Invalid = shared.Constant{
		Name:    "Invalid",
		String:  fmt.Sprintf("%sInvalid(%d)", ProfileType, basetype.FromString(data.Base).Invalid()),
		Comment: "INVALID",
	}

	return builder.Data{
		Template:     b.template,
		TemplateExec: "profile",
		Path:         b.path,
		Filename:     "profile_gen.go",
		Data: ProfileData{
			ConstantData:     data,
			MappingBaseTypes: mappingToBaseTypes,
		},
	}
}

func (b *profilebuilder) buildVersion() builder.Data {
	return builder.Data{
		Template:     b.template,
		TemplateExec: "version",
		Path:         b.path,
		Filename:     "version_gen.go",
		Data: VersionData{
			SDKVersion:     b.sdkVersion,
			Package:        "profile",
			ProfileVersion: toProfileVersion(b.sdkVersion),
		},
	}
}

func transformBaseType(s string) string {
	return "basetype." + s
}

func toProfileVersion(s string) string {
	s = strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, s)
	return s
}
