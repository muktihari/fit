// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// WatchfaceSettings is a WatchfaceSettings message.
type WatchfaceSettings struct {
	MessageIndex typedef.MessageIndex
	Mode         typedef.WatchfaceMode
	Layout       byte

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWatchfaceSettings creates new WatchfaceSettings struct based on given mesg. If mesg is nil or mesg.Num is not equal to WatchfaceSettings mesg number, it will return nil.
func NewWatchfaceSettings(mesg proto.Message) *WatchfaceSettings {
	if mesg.Num != typedef.MesgNumWatchfaceSettings {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		0:   basetype.EnumInvalid,   /* Mode */
		1:   basetype.ByteInvalid,   /* Layout */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &WatchfaceSettings{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Mode:         typeconv.ToEnum[typedef.WatchfaceMode](vals[0]),
		Layout:       typeconv.ToByte[byte](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to WatchfaceSettings mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumWatchfaceSettings)
func (m WatchfaceSettings) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumWatchfaceSettings {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.Mode,
		1:   m.Layout,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
