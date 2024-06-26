// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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
	BacklightModeInvalid                             BacklightMode = 0xFF
)

func (b BacklightMode) Byte() byte { return byte(b) }

func (b BacklightMode) String() string {
	switch b {
	case BacklightModeOff:
		return "off"
	case BacklightModeManual:
		return "manual"
	case BacklightModeKeyAndMessages:
		return "key_and_messages"
	case BacklightModeAutoBrightness:
		return "auto_brightness"
	case BacklightModeSmartNotifications:
		return "smart_notifications"
	case BacklightModeKeyAndMessagesNight:
		return "key_and_messages_night"
	case BacklightModeKeyAndMessagesAndSmartNotifications:
		return "key_and_messages_and_smart_notifications"
	default:
		return "BacklightModeInvalid(" + strconv.Itoa(int(b)) + ")"
	}
}

// FromString parse string into BacklightMode constant it's represent, return BacklightModeInvalid if not found.
func BacklightModeFromString(s string) BacklightMode {
	switch s {
	case "off":
		return BacklightModeOff
	case "manual":
		return BacklightModeManual
	case "key_and_messages":
		return BacklightModeKeyAndMessages
	case "auto_brightness":
		return BacklightModeAutoBrightness
	case "smart_notifications":
		return BacklightModeSmartNotifications
	case "key_and_messages_night":
		return BacklightModeKeyAndMessagesNight
	case "key_and_messages_and_smart_notifications":
		return BacklightModeKeyAndMessagesAndSmartNotifications
	default:
		return BacklightModeInvalid
	}
}

// List returns all constants.
func ListBacklightMode() []BacklightMode {
	return []BacklightMode{
		BacklightModeOff,
		BacklightModeManual,
		BacklightModeKeyAndMessages,
		BacklightModeAutoBrightness,
		BacklightModeSmartNotifications,
		BacklightModeKeyAndMessagesNight,
		BacklightModeKeyAndMessagesAndSmartNotifications,
	}
}
