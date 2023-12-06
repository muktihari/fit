// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Checksum uint8

const (
	ChecksumClear   Checksum = 0    // Allows clear of checksum for flash memory where can only write 1 to 0 without erasing sector.
	ChecksumOk      Checksum = 1    // Set to mark checksum as valid if computes to invalid values 0 or 0xFF. Checksum can also be set to ok to save encoding computation time.
	ChecksumInvalid Checksum = 0xFF // INVALID
)

var checksumtostrs = map[Checksum]string{
	ChecksumClear:   "clear",
	ChecksumOk:      "ok",
	ChecksumInvalid: "invalid",
}

func (c Checksum) String() string {
	val, ok := checksumtostrs[c]
	if !ok {
		return strconv.FormatUint(uint64(c), 10)
	}
	return val
}

var strtochecksum = func() map[string]Checksum {
	m := make(map[string]Checksum)
	for t, str := range checksumtostrs {
		m[str] = Checksum(t)
	}
	return m
}()

// FromString parse string into Checksum constant it's represent, return ChecksumInvalid if not found.
func ChecksumFromString(s string) Checksum {
	val, ok := strtochecksum[s]
	if !ok {
		return strtochecksum["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListChecksum() []Checksum {
	vs := make([]Checksum, 0, len(checksumtostrs))
	for i := range checksumtostrs {
		vs = append(vs, Checksum(i))
	}
	return vs
}
