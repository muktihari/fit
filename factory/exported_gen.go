// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

var std = New()

// StandardFactory returns standard factory.
func StandardFactory() *Factory { return std }

// CreateMesg creates new message based on defined messages in the factory. If not found, it returns new message with "unknown" name.
//
// This will create a shallow copy of the Fields, so changing any value declared in Field's FieldBase is prohibited (except fot the unknown field).
// If you want a deep copy of the mesg, create it by calling mesg.Clone().
func CreateMesg(num typedef.MesgNum) proto.Message {
	return std.CreateMesg(num)
}

// CreateMesgOnly is similar to CreateMesg, but it sets Fields to nil. This is useful when we plan to fill these values ourselves
// to avoid unnecessary malloc when cloning them, as they will be removed anyway. For example, the decoding process will populate them with decoded data.
func CreateMesgOnly(num typedef.MesgNum) proto.Message {
	return std.CreateMesgOnly(num)
}

// CreateField creates new field based on defined messages in the factory. If not found, it returns new field with "unknown" name.
//
// Field's FieldBase is a pointer struct embedded, and this will only create a shallow copy of the field, so changing any value declared in
// FieldBase is prohibited (except fot the unknown field) since it still referencing the same struct. If you want a deep copy of the Field,
// create it by calling field.Clone().
func CreateField(mesgNum typedef.MesgNum, num byte) proto.Field {
	return std.CreateField(mesgNum, num)
}

// RegisterMesg registers a new message that is not defined in the profile.xlsx.
// You can not edit or replace existing predefined messages in the factory, you can only edit the messages you have registered.
// However, we don't create a lock for efficiency, since this is intended to be used on instantiation. If you want to
// change something without triggering data race, you can create a new instance of Factory using New().
//
// By registering, any Fit file containing these messages can be recognized instead of returning "unknown" message.
func RegisterMesg(mesg proto.Message) error {
	return std.RegisterMesg(mesg)
}
