// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lookup

import (
	"fmt"
	"math"
	"strconv"

	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/profile/basetype"
)

type Lookup struct {
	mesgNumByName    map[string]uint16
	fieldsByMesgName map[string][]parser.Field // fields is small, slice should be sufficient.
	profiles         map[string]profile
}

type profile struct {
	goType       string
	baseType     basetype.BaseType
	valuesByName map[string]string
}

func New(types []parser.Type, messages []parser.Message) *Lookup {
	l := &Lookup{
		mesgNumByName:    make(map[string]uint16),
		fieldsByMesgName: make(map[string][]parser.Field),
		profiles:         make(map[string]profile),
	}
	l.populateLookupData(types, messages)
	return l
}

func (l *Lookup) populateLookupData(types []parser.Type, messages []parser.Message) {
	l.populateMesgNumLookup(types)
	l.populateProfileLookup(types)
	l.populateMessageLookup(messages)
}

func (l *Lookup) populateMesgNumLookup(types []parser.Type) {
	for _, typ := range types {
		if typ.Name == "mesg_num" {
			for _, v := range typ.Values {
				x, err := strconv.ParseUint(v.Value, 0, 16)
				if err != nil {
					panic(fmt.Sprintf("could not parse mesg_num %q with value %q: %v", v.Name, v.Value, err))
				}
				l.mesgNumByName[v.Name] = uint16(x)
			}
			break
		}
	}
}

func (l *Lookup) populateProfileLookup(types []parser.Type) {
	// additional profile type, mark it as enum
	l.profiles["bool"] = profile{
		baseType: basetype.Enum,
		goType:   "bool",
	}

	// map fit_base_type to our defined type "basetype.BaseType"
	l.profiles["fit_base_type"] = profile{
		baseType: basetype.Uint8,
		goType:   "basetype.BaseType",
	}

	basetypes := make(map[string]basetype.BaseType)

	// map basetype as profile type
	for _, bt := range basetype.List() {
		l.profiles[bt.String()] = profile{
			goType:   bt.GoType(),
			baseType: bt,
		}
		basetypes[bt.String()] = bt
	}

	for _, typ := range types {
		p := profile{
			goType:       basetypes[typ.BaseType].GoType(),
			baseType:     basetype.FromString(typ.BaseType),
			valuesByName: make(map[string]string),
		}

		for _, val := range typ.Values {
			p.valuesByName[val.Name] = val.Value
		}

		l.profiles[typ.Name] = p
	}
}

func (l *Lookup) populateMessageLookup(messages []parser.Message) {
	for _, mesg := range messages {
		l.fieldsByMesgName[mesg.Name] = mesg.Fields
	}
}

func (l *Lookup) MesgNumByName(name string) uint16 {
	num, ok := l.mesgNumByName[name]
	if !ok {
		return math.MaxUint16
	}
	return num
}

func (l *Lookup) BaseType(profileType string) basetype.BaseType {
	return l.profiles[profileType].baseType
}

func (l *Lookup) GoType(profileType string) string {
	return l.profiles[profileType].goType
}

func (l *Lookup) TypeValue(profileType string, valueName string) string {
	return l.profiles[profileType].valuesByName[valueName]
}

func (l *Lookup) FieldByNum(mesgName string, fieldNum byte) parser.Field {
	for _, field := range l.fieldsByMesgName[mesgName] {
		if field.Num == fieldNum {
			return field
		}
	}
	return parser.Field{Num: 255, Name: "unknown"}
}

func (l *Lookup) FieldByName(mesgName string, fieldName string) parser.Field {
	for _, field := range l.fieldsByMesgName[mesgName] {
		if field.Name == fieldName {
			return field
		}
	}
	return parser.Field{Num: 255, Name: "unknown"}
}
