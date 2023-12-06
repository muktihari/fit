// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LanguageBits2 uint8

const (
	LanguageBits2Slovenian LanguageBits2 = 0x01
	LanguageBits2Swedish   LanguageBits2 = 0x02
	LanguageBits2Russian   LanguageBits2 = 0x04
	LanguageBits2Turkish   LanguageBits2 = 0x08
	LanguageBits2Latvian   LanguageBits2 = 0x10
	LanguageBits2Ukrainian LanguageBits2 = 0x20
	LanguageBits2Arabic    LanguageBits2 = 0x40
	LanguageBits2Farsi     LanguageBits2 = 0x80
	LanguageBits2Invalid   LanguageBits2 = 0x0 // INVALID
)

var languagebits2tostrs = map[LanguageBits2]string{
	LanguageBits2Slovenian: "slovenian",
	LanguageBits2Swedish:   "swedish",
	LanguageBits2Russian:   "russian",
	LanguageBits2Turkish:   "turkish",
	LanguageBits2Latvian:   "latvian",
	LanguageBits2Ukrainian: "ukrainian",
	LanguageBits2Arabic:    "arabic",
	LanguageBits2Farsi:     "farsi",
	LanguageBits2Invalid:   "invalid",
}

func (l LanguageBits2) String() string {
	val, ok := languagebits2tostrs[l]
	if !ok {
		return strconv.FormatUint(uint64(l), 10)
	}
	return val
}

var strtolanguagebits2 = func() map[string]LanguageBits2 {
	m := make(map[string]LanguageBits2)
	for t, str := range languagebits2tostrs {
		m[str] = LanguageBits2(t)
	}
	return m
}()

// FromString parse string into LanguageBits2 constant it's represent, return LanguageBits2Invalid if not found.
func LanguageBits2FromString(s string) LanguageBits2 {
	val, ok := strtolanguagebits2[s]
	if !ok {
		return strtolanguagebits2["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLanguageBits2() []LanguageBits2 {
	vs := make([]LanguageBits2, 0, len(languagebits2tostrs))
	for i := range languagebits2tostrs {
		vs = append(vs, LanguageBits2(i))
	}
	return vs
}
