// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TimerTrigger byte

const (
	TimerTriggerManual           TimerTrigger = 0
	TimerTriggerAuto             TimerTrigger = 1
	TimerTriggerFitnessEquipment TimerTrigger = 2
	TimerTriggerInvalid          TimerTrigger = 0xFF
)

func (t TimerTrigger) String() string {
	switch t {
	case TimerTriggerManual:
		return "manual"
	case TimerTriggerAuto:
		return "auto"
	case TimerTriggerFitnessEquipment:
		return "fitness_equipment"
	default:
		return "TimerTriggerInvalid(" + strconv.Itoa(int(t)) + ")"
	}
}

// FromString parse string into TimerTrigger constant it's represent, return TimerTriggerInvalid if not found.
func TimerTriggerFromString(s string) TimerTrigger {
	switch s {
	case "manual":
		return TimerTriggerManual
	case "auto":
		return TimerTriggerAuto
	case "fitness_equipment":
		return TimerTriggerFitnessEquipment
	default:
		return TimerTriggerInvalid
	}
}

// List returns all constants.
func ListTimerTrigger() []TimerTrigger {
	return []TimerTrigger{
		TimerTriggerManual,
		TimerTriggerAuto,
		TimerTriggerFitnessEquipment,
	}
}
