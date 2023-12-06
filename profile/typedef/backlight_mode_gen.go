// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BacklightMode byte

const (
	BacklightModeOff                                 BacklightMode = 0
	BacklightModeManual                              BacklightMode = 1
	BacklightModeKeyAndMessages                      BacklightMode = 2
	BacklightModeAutoBrightness                      BacklightMode = 3
	BacklightModeSmartNotifications                  BacklightMode = 4
	BacklightModeKeyAndMessagesNight                 BacklightMode = 5
	BacklightModeKeyAndMessagesAndSmartNotifications BacklightMode = 6
	BacklightModeInvalid                             BacklightMode = 0xFF // INVALID
)

var backlightmodetostrs = map[BacklightMode]string{
	BacklightModeOff:                                 "off",
	BacklightModeManual:                              "manual",
	BacklightModeKeyAndMessages:                      "key_and_messages",
	BacklightModeAutoBrightness:                      "auto_brightness",
	BacklightModeSmartNotifications:                  "smart_notifications",
	BacklightModeKeyAndMessagesNight:                 "key_and_messages_night",
	BacklightModeKeyAndMessagesAndSmartNotifications: "key_and_messages_and_smart_notifications",
	BacklightModeInvalid:                             "invalid",
}

func (b BacklightMode) String() string {
	val, ok := backlightmodetostrs[b]
	if !ok {
		return strconv.Itoa(int(b))
	}
	return val
}

var strtobacklightmode = func() map[string]BacklightMode {
	m := make(map[string]BacklightMode)
	for t, str := range backlightmodetostrs {
		m[str] = BacklightMode(t)
	}
	return m
}()

// FromString parse string into BacklightMode constant it's represent, return BacklightModeInvalid if not found.
func BacklightModeFromString(s string) BacklightMode {
	val, ok := strtobacklightmode[s]
	if !ok {
		return strtobacklightmode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListBacklightMode() []BacklightMode {
	vs := make([]BacklightMode, 0, len(backlightmodetostrs))
	for i := range backlightmodetostrs {
		vs = append(vs, BacklightMode(i))
	}
	return vs
}
