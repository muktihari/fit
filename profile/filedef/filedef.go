// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

type File interface {
	Add(mesg proto.Message)
}

type Factory interface {
	CreateMesg(num typedef.MesgNum) proto.Message
}

type PutMessage interface {
	PutMessage(mesg *proto.Message)
}

func PutMessages[S []E, E PutMessage](factory Factory, messages *[]proto.Message, mesgNum typedef.MesgNum, s S) {
	for i := range s {
		mesg := factory.CreateMesg(mesgNum)
		s[i].PutMessage(&mesg)
		*messages = append(*messages, mesg)
	}
}
