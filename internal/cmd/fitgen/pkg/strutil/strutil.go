// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strutil is used only for naming package, variable, etc inside the go template.
package strutil

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToTitle tranforms s into TitleCase and remove non-alphanumeric characters.
func ToTitle(s string) string {
	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return ' '
		}
		return r
	}, s)
	s = cases.Title(language.English).String(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}

// ToCamel tranforms s into camelCase and remove non-alphanumeric characters.
func ToCamel(s string) string {
	s = ToTitle(s)
	s = strings.ToLower(string(s[0])) + s[1:]
	return s
}

// ToLower tranforms s into lowercase and remove non-alphanumeric characters.
func ToLower(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, s)
	return s
}

// ToSnake tranforms s into snake_case and replace non-alphanumeric characters with underscore '_'
func ToSnake(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return '_'
		}
		return r
	}, s)
	return s
}

// ToNonGoIdent transform s when s is equal with go identifier, e.g. switch -> switch_z
func ToNonGoIdent(s string) string {
	if token.Lookup(s).IsKeyword() {
		return s + "_z"
	}
	return s
}

// ToLetterPrefix transform s when s is prefixed with non-letter character, e.g. 1Partcarbon -> Z_1Partcarbon
func ToLetterPrefix(s string) string {
	if !unicode.IsLetter(rune(s[0])) {
		return "Z_" + s
	}
	return s
}

// TrimRepeatedChar trim repeated given char in any position in the s. e.g. (s: "a,,b,,,c", char: ',') -> a,b,c
func TrimRepeatedChar(s string, char rune) string {
	if s == "" {
		return s
	}
	var last byte
	var strbuf = new(strings.Builder)
	for i := range s {
		c := s[i]
		if c != last || rune(c) != char {
			last = c
			strbuf.WriteByte(c)
		}
	}
	return strbuf.String()
}
