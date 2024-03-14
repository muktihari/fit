// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Tone byte

const (
	ToneOff            Tone = 0
	ToneTone           Tone = 1
	ToneVibrate        Tone = 2
	ToneToneAndVibrate Tone = 3
	ToneInvalid        Tone = 0xFF
)

func (t Tone) String() string {
	switch t {
	case ToneOff:
		return "off"
	case ToneTone:
		return "tone"
	case ToneVibrate:
		return "vibrate"
	case ToneToneAndVibrate:
		return "tone_and_vibrate"
	default:
		return "ToneInvalid(" + strconv.Itoa(int(t)) + ")"
	}
}

// FromString parse string into Tone constant it's represent, return ToneInvalid if not found.
func ToneFromString(s string) Tone {
	switch s {
	case "off":
		return ToneOff
	case "tone":
		return ToneTone
	case "vibrate":
		return ToneVibrate
	case "tone_and_vibrate":
		return ToneToneAndVibrate
	default:
		return ToneInvalid
	}
}

// List returns all constants.
func ListTone() []Tone {
	return []Tone{
		ToneOff,
		ToneTone,
		ToneVibrate,
		ToneToneAndVibrate,
	}
}
