// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LocalDateTime uint32

const (
	LocalDateTimeMin     LocalDateTime = 0x10000000 // if date_time is < 0x10000000 then it is system time (seconds from device power on)
	LocalDateTimeInvalid LocalDateTime = 0xFFFFFFFF
)

func (l LocalDateTime) Uint32() uint32 { return uint32(l) }

func (l LocalDateTime) String() string {
	switch l {
	case LocalDateTimeMin:
		return "min"
	default:
		return "LocalDateTimeInvalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LocalDateTime constant it's represent, return LocalDateTimeInvalid if not found.
func LocalDateTimeFromString(s string) LocalDateTime {
	switch s {
	case "min":
		return LocalDateTimeMin
	default:
		return LocalDateTimeInvalid
	}
}

// List returns all constants.
func ListLocalDateTime() []LocalDateTime {
	return []LocalDateTime{
		LocalDateTimeMin,
	}
}
