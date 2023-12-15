// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fieldnum

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"text/template"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/builder/shared"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/profile/basetype"
	"golang.org/x/exp/slices"
)

type (
	ProfileType = string
	BaseType    = string
)

type fieldnumBuilder struct {
	template     *template.Template
	templateExec string

	path       string // path to generate the file
	sdkVersion string // Fit SDK Version
	messages   []parser.Message
	types      []parser.Type

	once sync.Once

	baseTypeMapByProfileType  map[ProfileType]BaseType
	constantsMapByProfileType map[ProfileType][]string
}

func NewBuilder(path, version string, message []parser.Message, types []parser.Type) builder.Builder {
	_, filename, _, _ := runtime.Caller(0)
	cd := filepath.Dir(filename)
	return &fieldnumBuilder{
		template: template.Must(template.New("main").
			Funcs(shared.FuncMap()).
			ParseFiles(
				filepath.Join(cd, "fieldnum.tmpl"),
				filepath.Join(cd, "..", "..", "..", "builder", "shared", "untyped_constant.tmpl"))),
		templateExec: "fieldnum",
		path:         filepath.Join(path, "profile", "untyped", "fieldnum"),
		sdkVersion:   version,
		messages:     message,
		types:        types,
	}
}

func (b *fieldnumBuilder) prepopulateLookup() {
	b.baseTypeMapByProfileType = make(map[ProfileType]BaseType)
	for _, t := range basetype.List() { // map to itself
		b.baseTypeMapByProfileType[t.String()] = t.String()
	}

	b.constantsMapByProfileType = make(map[ProfileType][]string)
	for _, _type := range b.types {
		b.baseTypeMapByProfileType[_type.Name] = _type.BaseType
		for _, value := range _type.Values {
			b.constantsMapByProfileType[_type.Name] = append(b.constantsMapByProfileType[_type.Name], value.Name, value.Value)
		}
	}
	b.baseTypeMapByProfileType["bool"] = "bool"
}

func (b *fieldnumBuilder) Build() ([]builder.Data, error) {
	b.once.Do(func() { b.prepopulateLookup() })

	_type := "FieldNum"
	constants := make([]shared.Constant, 0)
	for _, mesg := range b.messages {
		for _, field := range mesg.Fields {
			var scaleOffset string
			scale := scaleOrDefault(field.Scales, 0)
			offset := offsetOrDefault(field.Offsets, 0)
			if scale != 1 || offset != 0 {
				scaleOffset = fmt.Sprintf(", Scale: %g, Offset: %g", scale, offset)
			}

			var units string
			if field.Units != "" {
				units = fmt.Sprintf(", Units: %s", field.Units)
			}
			var array string
			if field.Array != "" {
				array = fmt.Sprintf(", Array: %s", field.Array)
			}

			c := shared.Constant{
				Name:   strutil.ToTitle(mesg.Name) + strutil.ToTitle(field.Name),
				Type:   _type,
				Op:     "=",
				Value:  strconv.Itoa(int(field.Num)),
				String: mesg.Name + ": " + field.Name,
				Comment: fmt.Sprintf("[Type: %s, Base: %s%s%s%s]; %s",
					strutil.ToTitle(field.Type),
					b.baseTypeMapByProfileType[field.Type],
					array,
					scaleOffset,
					units,
					field.Comment,
				),
			}
			constants = append(constants, c)
		}
	}
	slices.SortStableFunc(constants, func(x, y shared.Constant) int {
		if x.Name < y.Name {
			return -1
		} else if x.Name > y.Name {
			return 1
		}
		return 0
	})

	constants = append(constants, shared.Constant{
		Name:   "Invalid",
		Type:   _type,
		Op:     "=",
		Value:  "255", // max byte
		String: "invalid",
	})

	dataBuilder := builder.Data{
		Template:     b.template,
		TemplateExec: b.templateExec,
		Path:         b.path,
		Filename:     "fieldnum_gen.go",
		Data: shared.ConstantData{
			Package:    "fieldnum",
			SDKVersion: b.sdkVersion,
			Type:       _type,
			Base:       "byte",
			Constants:  constants,
		},
	}

	return []builder.Data{dataBuilder}, nil
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
