// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SetType uint8

const (
	SetTypeRest    SetType = 0
	SetTypeActive  SetType = 1
	SetTypeInvalid SetType = 0xFF // INVALID
)

var settypetostrs = map[SetType]string{
	SetTypeRest:    "rest",
	SetTypeActive:  "active",
	SetTypeInvalid: "invalid",
}

func (s SetType) String() string {
	val, ok := settypetostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosettype = func() map[string]SetType {
	m := make(map[string]SetType)
	for t, str := range settypetostrs {
		m[str] = SetType(t)
	}
	return m
}()

// FromString parse string into SetType constant it's represent, return SetTypeInvalid if not found.
func SetTypeFromString(s string) SetType {
	val, ok := strtosettype[s]
	if !ok {
		return strtosettype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSetType() []SetType {
	vs := make([]SetType, 0, len(settypetostrs))
	for i := range settypetostrs {
		vs = append(vs, SetType(i))
	}
	return vs
}
