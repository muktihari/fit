// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

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
	FileFlagsInvalid FileFlags = 0x0 // INVALID
)

var fileflagstostrs = map[FileFlags]string{
	FileFlagsRead:    "read",
	FileFlagsWrite:   "write",
	FileFlagsErase:   "erase",
	FileFlagsInvalid: "invalid",
}

func (f FileFlags) String() string {
	val, ok := fileflagstostrs[f]
	if !ok {
		return strconv.FormatUint(uint64(f), 10)
	}
	return val
}

var strtofileflags = func() map[string]FileFlags {
	m := make(map[string]FileFlags)
	for t, str := range fileflagstostrs {
		m[str] = FileFlags(t)
	}
	return m
}()

// FromString parse string into FileFlags constant it's represent, return FileFlagsInvalid if not found.
func FileFlagsFromString(s string) FileFlags {
	val, ok := strtofileflags[s]
	if !ok {
		return strtofileflags["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListFileFlags() []FileFlags {
	vs := make([]FileFlags, 0, len(fileflagstostrs))
	for i := range fileflagstostrs {
		vs = append(vs, FileFlags(i))
	}
	return vs
}
