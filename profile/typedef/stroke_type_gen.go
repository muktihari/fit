// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type StrokeType byte

const (
	StrokeTypeNoEvent  StrokeType = 0
	StrokeTypeOther    StrokeType = 1 // stroke was detected but cannot be identified
	StrokeTypeServe    StrokeType = 2
	StrokeTypeForehand StrokeType = 3
	StrokeTypeBackhand StrokeType = 4
	StrokeTypeSmash    StrokeType = 5
	StrokeTypeInvalid  StrokeType = 0xFF
)

func (s StrokeType) String() string {
	switch s {
	case StrokeTypeNoEvent:
		return "no_event"
	case StrokeTypeOther:
		return "other"
	case StrokeTypeServe:
		return "serve"
	case StrokeTypeForehand:
		return "forehand"
	case StrokeTypeBackhand:
		return "backhand"
	case StrokeTypeSmash:
		return "smash"
	default:
		return "StrokeTypeInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into StrokeType constant it's represent, return StrokeTypeInvalid if not found.
func StrokeTypeFromString(s string) StrokeType {
	switch s {
	case "no_event":
		return StrokeTypeNoEvent
	case "other":
		return StrokeTypeOther
	case "serve":
		return StrokeTypeServe
	case "forehand":
		return StrokeTypeForehand
	case "backhand":
		return StrokeTypeBackhand
	case "smash":
		return StrokeTypeSmash
	default:
		return StrokeTypeInvalid
	}
}

// List returns all constants.
func ListStrokeType() []StrokeType {
	return []StrokeType{
		StrokeTypeNoEvent,
		StrokeTypeOther,
		StrokeTypeServe,
		StrokeTypeForehand,
		StrokeTypeBackhand,
		StrokeTypeSmash,
	}
}
