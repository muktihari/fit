// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

import (
	"fmt"
	"math"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

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

	path           string        // path to generate the file
	profileVersion string        // FIT SDK Profile Version
	types          []parser.Type // type parsed from profile.xlsx
}

func NewBuilder(path, profileVersion string, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &profilebuilder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(
				filepath.Join(cd, "profile.tmpl"),
				filepath.Join(cd, "..", "builder", "shared", "constant.tmpl"))),
		path:           filepath.Join(path, "profile"),
		profileVersion: profileVersion,
		types:          types,
	}
}

func (b *profilebuilder) Build() ([]builder.Data, error) {
	profileDataBuilder := b.buildProfile()
	versionDataBuilder := b.buildVersion()

	return []builder.Data{profileDataBuilder, versionDataBuilder}, nil
}

func (b *profilebuilder) buildProfile() builder.Data {
	constants := make([]shared.Constant, 0, len(b.types))
	mappingProfileTypeToBaseTypes := make([]ProfileTypeBaseType, 0, len(b.types))
	var mappingBaseTypeToProfileTypes []ProfileTypeBaseType

	for _, t := range b.types {
		if t.Name == FitBaseType { // special types to be included, mapping to itself (profile.Uint8 == basetype.Uint8)
			mappingBaseTypeToProfileTypes = make([]ProfileTypeBaseType, 0, len(t.Values))
			for _, v := range t.Values {
				constantName := strutil.ToTitle(v.Name)
				baseType := transformBaseType(strutil.ToTitle(v.Name))
				constants = append(constants, shared.Constant{
					Name:   constantName,
					String: v.Name,
				})
				mappingProfileTypeToBaseTypes = append(mappingProfileTypeToBaseTypes, ProfileTypeBaseType{
					ProfileType: constantName,
					BaseType:    baseType,
				})
				mappingBaseTypeToProfileTypes = append(mappingBaseTypeToProfileTypes, ProfileTypeBaseType{
					BaseType:    baseType,
					ProfileType: constantName,
				})
			}

			constants = append(constants, shared.Constant{Name: "Bool", String: "bool"})
			mappingProfileTypeToBaseTypes = append(mappingProfileTypeToBaseTypes, ProfileTypeBaseType{ProfileType: "Bool", BaseType: transformBaseType("Enum")})
		}
	}

	for _, t := range b.types {
		constants = append(constants, shared.Constant{
			Name:   strutil.ToTitle(t.Name),
			String: t.Name,
		})
		mappingProfileTypeToBaseTypes = append(mappingProfileTypeToBaseTypes, ProfileTypeBaseType{
			ProfileType: strutil.ToTitle(t.Name),
			BaseType:    transformBaseType(strutil.ToTitle(t.BaseType)),
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
		Type:      ProfileType,
		Base:      "uint16",
		Constants: constants,
	}

	data.Invalid = shared.Constant{
		Name:    "Invalid",
		Op:      "=",
		Type:    ProfileType,
		Value:   strconv.Itoa(math.MaxUint16),
		String:  fmt.Sprintf("%sInvalid(%d)", ProfileType, basetype.FromString(data.Base).Invalid()),
		Comment: "INVALID",
	}

	return builder.Data{
		Template:     b.template,
		TemplateExec: "profile",
		Path:         b.path,
		Filename:     "profile_gen.go",
		Data: ProfileData{
			ConstantData:                  data,
			MappingProfileTypeToBaseTypes: mappingProfileTypeToBaseTypes,
			MappingBaseTypeToProfileTypes: mappingBaseTypeToProfileTypes,
		},
	}
}

func (b *profilebuilder) buildVersion() builder.Data {
	return builder.Data{
		Template:     b.template,
		TemplateExec: "version",
		Path:         b.path,
		Filename:     "version_gen.go",
		Data:         createVersionData(b.profileVersion),
	}
}

func transformBaseType(s string) string {
	return "basetype." + s
}

func createVersionData(profileVersion string) VersionData {
	// On error, use panic so we can get stack trace, should not generate when version is invalid.
	parts := strings.Split(profileVersion, ".")
	if len(parts) < 2 {
		panic(fmt.Errorf("malformed profile version, should in the form of <major>.<minor>, got: %s", profileVersion))
	}
	var (
		majorPart = parts[0]
		minorPart = parts[1]
	)

	major, err := strconv.ParseUint(majorPart, 10, 16)
	if err != nil {
		panic(fmt.Errorf("invalid major version: %w", err))
	}
	minor, err := strconv.ParseUint(minorPart, 10, 16)
	if err != nil {
		panic(fmt.Errorf("invalid minor version: %w", err))
	}
	version, err := strconv.ParseUint(majorPart+minorPart, 10, 16)
	if err != nil {
		panic(fmt.Errorf("invalid version: %w", err))
	}

	return VersionData{
		Package:        "profile",
		ProfileVersion: profileVersion,
		Major:          uint16(major),
		Minor:          uint16(minor),
		Version:        uint16(version),
	}
}
