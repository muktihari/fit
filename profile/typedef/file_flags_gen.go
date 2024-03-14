// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type FileFlags uint8

const (
	FileFlagsRead    FileFlags = 0x02
	FileFlagsWrite   FileFlags = 0x04
	FileFlagsErase   FileFlags = 0x08
	FileFlagsInvalid FileFlags = 0x0
)

func (f FileFlags) String() string {
	switch f {
	case FileFlagsRead:
		return "read"
	case FileFlagsWrite:
		return "write"
	case FileFlagsErase:
		return "erase"
	default:
		return "FileFlagsInvalid(" + strconv.FormatUint(uint64(f), 10) + ")"
	}
}

// FromString parse string into FileFlags constant it's represent, return FileFlagsInvalid if not found.
func FileFlagsFromString(s string) FileFlags {
	switch s {
	case "read":
		return FileFlagsRead
	case "write":
		return FileFlagsWrite
	case "erase":
		return FileFlagsErase
	default:
		return FileFlagsInvalid
	}
}

// List returns all constants.
func ListFileFlags() []FileFlags {
	return []FileFlags{
		FileFlagsRead,
		FileFlagsWrite,
		FileFlagsErase,
	}
}
