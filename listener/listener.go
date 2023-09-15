// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package listener

import "github.com/muktihari/fit/proto"

// MesgListener is an interface to listen to message decoded event.
type MesgListener interface {
	// OnMesg receives message from emitter.
	OnMesg(mesg proto.Message)
}

// MesgDefListener is an interface to listen to message definition decoded event.
type MesgDefListener interface {
	// OnMesgDef receives message definition from emitter.
	OnMesgDef(mesgDef proto.MessageDefinition)
}
