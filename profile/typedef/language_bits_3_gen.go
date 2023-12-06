// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LanguageBits3 uint8

const (
	LanguageBits3Bulgarian LanguageBits3 = 0x01
	LanguageBits3Romanian  LanguageBits3 = 0x02
	LanguageBits3Chinese   LanguageBits3 = 0x04
	LanguageBits3Japanese  LanguageBits3 = 0x08
	LanguageBits3Korean    LanguageBits3 = 0x10
	LanguageBits3Taiwanese LanguageBits3 = 0x20
	LanguageBits3Thai      LanguageBits3 = 0x40
	LanguageBits3Hebrew    LanguageBits3 = 0x80
	LanguageBits3Invalid   LanguageBits3 = 0x0 // INVALID
)

var languagebits3tostrs = map[LanguageBits3]string{
	LanguageBits3Bulgarian: "bulgarian",
	LanguageBits3Romanian:  "romanian",
	LanguageBits3Chinese:   "chinese",
	LanguageBits3Japanese:  "japanese",
	LanguageBits3Korean:    "korean",
	LanguageBits3Taiwanese: "taiwanese",
	LanguageBits3Thai:      "thai",
	LanguageBits3Hebrew:    "hebrew",
	LanguageBits3Invalid:   "invalid",
}

func (l LanguageBits3) String() string {
	val, ok := languagebits3tostrs[l]
	if !ok {
		return strconv.FormatUint(uint64(l), 10)
	}
	return val
}

var strtolanguagebits3 = func() map[string]LanguageBits3 {
	m := make(map[string]LanguageBits3)
	for t, str := range languagebits3tostrs {
		m[str] = LanguageBits3(t)
	}
	return m
}()

// FromString parse string into LanguageBits3 constant it's represent, return LanguageBits3Invalid if not found.
func LanguageBits3FromString(s string) LanguageBits3 {
	val, ok := strtolanguagebits3[s]
	if !ok {
		return strtolanguagebits3["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLanguageBits3() []LanguageBits3 {
	vs := make([]LanguageBits3, 0, len(languagebits3tostrs))
	for i := range languagebits3tostrs {
		vs = append(vs, LanguageBits3(i))
	}
	return vs
}
