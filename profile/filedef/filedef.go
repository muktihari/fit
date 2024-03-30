// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

// File is an interface for defining common type file, any defined common file type should implement
// the following methods to be able to work with Listener (and other building block in filedef package).
type File interface {
	// Add adds message into file structure.
	Add(mesg proto.Message)
	// ToFit converts file back to proto.Fit structure.
	ToFit(options *mesgdef.Options) proto.Fit
}

// ToMesgs bulks convert mesgdef into proto.Message and append it to messages
func ToMesgs[S []E, E ToMesg](messages *[]proto.Message, options *mesgdef.Options, mesgNum typedef.MesgNum, s S) {
	for i := range s {
		*messages = append(*messages, s[i].ToMesg(options))
	}
}

// ToMesg is a type constraint to retrieve all mesgdef structures which implement ToMesg method.
type ToMesg interface {
	ToMesg(options *mesgdef.Options) proto.Message
}

// SortMessagesByTimestamp sorts messages by timestamp only if the message has timestamp field.
// When a message has no timestamp field, its order will not be changed.
func SortMessagesByTimestamp(messages []proto.Message) {
	slices.SortStableFunc(messages, func(m1, m2 proto.Message) int {
		value1 := m1.FieldValueByNum(proto.FieldNumTimestamp)
		if value1 == nil {
			return 0
		}

		value2 := m2.FieldValueByNum(proto.FieldNumTimestamp)
		if value2 == nil {
			return 0
		}

		timestamp1 := typeconv.ToUint32[uint32](value1)
		timestamp2 := typeconv.ToUint32[uint32](value2)

		if timestamp1 <= timestamp2 {
			return -1
		}

		return 1
	})
}
