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
			ParseFiles(filepath.Join(cd, "profile.tmpl"), "builder/shared/constant.tmpl")),
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
	mappingToBaseTypes := make([]MappingBaseType, 0, len(b.types))
	constants := make([]shared.Constant, 0, len(b.types))

	for _, t := range b.types {
		if t.Name == FitBaseType { // special types to be included
			ms := make([]MappingBaseType, len(mappingToBaseTypes)+len(t.Values)+1, len(t.Values)+cap(mappingToBaseTypes)+1)
			cs := make([]shared.Constant, len(constants)+len(t.Values)+1, len(t.Values)+cap(constants)+1)
			for i, v := range t.Values {
				ms[i] = MappingBaseType{
					ConstantName: strutil.ToTitle(v.Name),
					BaseType:     transformBaseType(strutil.ToTitle(v.Name)), // mapping to itself
				}
				cs[i] = shared.Constant{
					Name:   strutil.ToTitle(v.Name),
					String: v.Name,
				}
			}

			cs[len(t.Values)] = shared.Constant{Name: "Bool", String: "bool"} // +1
			copy(cs[len(t.Values)+1:], constants)
			constants = cs

			ms[len(t.Values)] = MappingBaseType{ConstantName: "Bool", BaseType: transformBaseType("Enum")} // +1
			copy(ms[len(t.Values)+1:], mappingToBaseTypes)
			mappingToBaseTypes = ms
		}

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

	constants = append(constants, shared.Constant{
		Name:    "Invalid",
		String:  "invalid",
		Comment: "INVALID",
	})

	mappingToBaseTypes = append(mappingToBaseTypes, MappingBaseType{
		ConstantName: "Invalid",
		BaseType:     fmt.Sprint(basetype.Byte.Invalid()),
	})

	return builder.Data{
		Template:     b.template,
		TemplateExec: "profile",
		Path:         b.path,
		Filename:     "profile_gen.go",
		Data: ProfileData{
			ConstantData: shared.ConstantData{
				Package: "profile",
				Imports: []string{
					"strconv",
					"github.com/muktihari/fit/profile/basetype",
				},
				StringerMode: shared.StringerArray,
				SDKVersion:   b.sdkVersion,
				Type:         ProfileType,
				Base:         "uint16",
				Constants:    constants,
			},
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
