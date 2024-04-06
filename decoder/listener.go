// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import "github.com/muktihari/fit/proto"

// MesgListener is an interface for listening to message decoded events.
type MesgListener interface {
	// OnMesg receives message from Decoder.
	OnMesg(mesg proto.Message)
}

// MesgDefListener is an interface for listening to message definition decoded event.
type MesgDefListener interface {
	// OnMesgDef receives message definition from Decoder.
	OnMesgDef(mesgDef proto.MessageDefinition)
}
