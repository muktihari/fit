// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

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
	TimerTriggerInvalid          TimerTrigger = 0xFF // INVALID
)

var timertriggertostrs = map[TimerTrigger]string{
	TimerTriggerManual:           "manual",
	TimerTriggerAuto:             "auto",
	TimerTriggerFitnessEquipment: "fitness_equipment",
	TimerTriggerInvalid:          "invalid",
}

func (t TimerTrigger) String() string {
	val, ok := timertriggertostrs[t]
	if !ok {
		return strconv.Itoa(int(t))
	}
	return val
}

var strtotimertrigger = func() map[string]TimerTrigger {
	m := make(map[string]TimerTrigger)
	for t, str := range timertriggertostrs {
		m[str] = TimerTrigger(t)
	}
	return m
}()

// FromString parse string into TimerTrigger constant it's represent, return TimerTriggerInvalid if not found.
func TimerTriggerFromString(s string) TimerTrigger {
	val, ok := strtotimertrigger[s]
	if !ok {
		return strtotimertrigger["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListTimerTrigger() []TimerTrigger {
	vs := make([]TimerTrigger, 0, len(timertriggertostrs))
	for i := range timertriggertostrs {
		vs = append(vs, TimerTrigger(i))
	}
	return vs
}
