// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SegmentDeleteStatus byte

const (
	SegmentDeleteStatusDoNotDelete SegmentDeleteStatus = 0
	SegmentDeleteStatusDeleteOne   SegmentDeleteStatus = 1
	SegmentDeleteStatusDeleteAll   SegmentDeleteStatus = 2
	SegmentDeleteStatusInvalid     SegmentDeleteStatus = 0xFF
)

func (s SegmentDeleteStatus) String() string {
	switch s {
	case SegmentDeleteStatusDoNotDelete:
		return "do_not_delete"
	case SegmentDeleteStatusDeleteOne:
		return "delete_one"
	case SegmentDeleteStatusDeleteAll:
		return "delete_all"
	default:
		return "SegmentDeleteStatusInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SegmentDeleteStatus constant it's represent, return SegmentDeleteStatusInvalid if not found.
func SegmentDeleteStatusFromString(s string) SegmentDeleteStatus {
	switch s {
	case "do_not_delete":
		return SegmentDeleteStatusDoNotDelete
	case "delete_one":
		return SegmentDeleteStatusDeleteOne
	case "delete_all":
		return SegmentDeleteStatusDeleteAll
	default:
		return SegmentDeleteStatusInvalid
	}
}

// List returns all constants.
func ListSegmentDeleteStatus() []SegmentDeleteStatus {
	return []SegmentDeleteStatus{
		SegmentDeleteStatusDoNotDelete,
		SegmentDeleteStatusDeleteOne,
		SegmentDeleteStatusDeleteAll,
	}
}
