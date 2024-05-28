// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

// File is an interface for defining common type file, any defined common file type should implement
// the following methods to be able to work with Listener (and other building block in filedef package).
type File interface {
	// Add adds message into file structure.
	Add(mesg proto.Message)
	// ToFIT converts file back to proto.FIT structure.
	ToFIT(options *mesgdef.Options) proto.FIT
}

// SortMessagesByTimestamp sorts messages by timestamp. The following rules will apply:
//   - Any message without timestamp field will be placed to the beginning of the slice
//     to enable these messages to be retrieved early such as UserProfile.
//   - Any message with invalid timestamp will be places at the end of the slices.
func SortMessagesByTimestamp(messages []proto.Message) {
	slices.SortStableFunc(messages, func(m1, m2 proto.Message) int {
		v1 := m1.FieldByNum(proto.FieldNumTimestamp)
		v2 := m2.FieldByNum(proto.FieldNumTimestamp)

		// Place message which does not have a timestamp at the beginning of the slice.
		if v1 == nil && v2 == nil {
			return 0
		} else if v1 == nil {
			return -1
		} else if v2 == nil {
			return 1
		}

		// Sort timestamps regardless of whether any of the values are invalid.
		// Any invalid value will be placed at the end of the slice.
		t1 := v1.Value.Uint32()
		t2 := v2.Value.Uint32()
		if t1 < t2 {
			return -1
		}
		if t1 > t2 {
			return 1
		}
		return 0
	})
}
