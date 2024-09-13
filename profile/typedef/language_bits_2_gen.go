// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LanguageBits2 uint8 // Base: uint8z

const (
	LanguageBits2Slovenian LanguageBits2 = 0x01
	LanguageBits2Swedish   LanguageBits2 = 0x02
	LanguageBits2Russian   LanguageBits2 = 0x04
	LanguageBits2Turkish   LanguageBits2 = 0x08
	LanguageBits2Latvian   LanguageBits2 = 0x10
	LanguageBits2Ukrainian LanguageBits2 = 0x20
	LanguageBits2Arabic    LanguageBits2 = 0x40
	LanguageBits2Farsi     LanguageBits2 = 0x80
	LanguageBits2Invalid   LanguageBits2 = 0x0
)

func (l LanguageBits2) Uint8() uint8 { return uint8(l) }

func (l LanguageBits2) String() string {
	switch l {
	case LanguageBits2Slovenian:
		return "slovenian"
	case LanguageBits2Swedish:
		return "swedish"
	case LanguageBits2Russian:
		return "russian"
	case LanguageBits2Turkish:
		return "turkish"
	case LanguageBits2Latvian:
		return "latvian"
	case LanguageBits2Ukrainian:
		return "ukrainian"
	case LanguageBits2Arabic:
		return "arabic"
	case LanguageBits2Farsi:
		return "farsi"
	default:
		return "LanguageBits2Invalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LanguageBits2 constant it's represent, return LanguageBits2Invalid if not found.
func LanguageBits2FromString(s string) LanguageBits2 {
	switch s {
	case "slovenian":
		return LanguageBits2Slovenian
	case "swedish":
		return LanguageBits2Swedish
	case "russian":
		return LanguageBits2Russian
	case "turkish":
		return LanguageBits2Turkish
	case "latvian":
		return LanguageBits2Latvian
	case "ukrainian":
		return LanguageBits2Ukrainian
	case "arabic":
		return LanguageBits2Arabic
	case "farsi":
		return LanguageBits2Farsi
	default:
		return LanguageBits2Invalid
	}
}

// List returns all constants.
func ListLanguageBits2() []LanguageBits2 {
	return []LanguageBits2{
		LanguageBits2Slovenian,
		LanguageBits2Swedish,
		LanguageBits2Russian,
		LanguageBits2Turkish,
		LanguageBits2Latvian,
		LanguageBits2Ukrainian,
		LanguageBits2Arabic,
		LanguageBits2Farsi,
	}
}
