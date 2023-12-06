// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LapTrigger byte

const (
	LapTriggerManual           LapTrigger = 0
	LapTriggerTime             LapTrigger = 1
	LapTriggerDistance         LapTrigger = 2
	LapTriggerPositionStart    LapTrigger = 3
	LapTriggerPositionLap      LapTrigger = 4
	LapTriggerPositionWaypoint LapTrigger = 5
	LapTriggerPositionMarked   LapTrigger = 6
	LapTriggerSessionEnd       LapTrigger = 7
	LapTriggerFitnessEquipment LapTrigger = 8
	LapTriggerInvalid          LapTrigger = 0xFF // INVALID
)

var laptriggertostrs = map[LapTrigger]string{
	LapTriggerManual:           "manual",
	LapTriggerTime:             "time",
	LapTriggerDistance:         "distance",
	LapTriggerPositionStart:    "position_start",
	LapTriggerPositionLap:      "position_lap",
	LapTriggerPositionWaypoint: "position_waypoint",
	LapTriggerPositionMarked:   "position_marked",
	LapTriggerSessionEnd:       "session_end",
	LapTriggerFitnessEquipment: "fitness_equipment",
	LapTriggerInvalid:          "invalid",
}

func (l LapTrigger) String() string {
	val, ok := laptriggertostrs[l]
	if !ok {
		return strconv.Itoa(int(l))
	}
	return val
}

var strtolaptrigger = func() map[string]LapTrigger {
	m := make(map[string]LapTrigger)
	for t, str := range laptriggertostrs {
		m[str] = LapTrigger(t)
	}
	return m
}()

// FromString parse string into LapTrigger constant it's represent, return LapTriggerInvalid if not found.
func LapTriggerFromString(s string) LapTrigger {
	val, ok := strtolaptrigger[s]
	if !ok {
		return strtolaptrigger["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLapTrigger() []LapTrigger {
	vs := make([]LapTrigger, 0, len(laptriggertostrs))
	for i := range laptriggertostrs {
		vs = append(vs, LapTrigger(i))
	}
	return vs
}
