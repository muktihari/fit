// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SessionTrigger byte

const (
	SessionTriggerActivityEnd      SessionTrigger = 0
	SessionTriggerManual           SessionTrigger = 1    // User changed sport.
	SessionTriggerAutoMultiSport   SessionTrigger = 2    // Auto multi-sport feature is enabled and user pressed lap button to advance session.
	SessionTriggerFitnessEquipment SessionTrigger = 3    // Auto sport change caused by user linking to fitness equipment.
	SessionTriggerInvalid          SessionTrigger = 0xFF // INVALID
)

var sessiontriggertostrs = map[SessionTrigger]string{
	SessionTriggerActivityEnd:      "activity_end",
	SessionTriggerManual:           "manual",
	SessionTriggerAutoMultiSport:   "auto_multi_sport",
	SessionTriggerFitnessEquipment: "fitness_equipment",
	SessionTriggerInvalid:          "invalid",
}

func (s SessionTrigger) String() string {
	val, ok := sessiontriggertostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosessiontrigger = func() map[string]SessionTrigger {
	m := make(map[string]SessionTrigger)
	for t, str := range sessiontriggertostrs {
		m[str] = SessionTrigger(t)
	}
	return m
}()

// FromString parse string into SessionTrigger constant it's represent, return SessionTriggerInvalid if not found.
func SessionTriggerFromString(s string) SessionTrigger {
	val, ok := strtosessiontrigger[s]
	if !ok {
		return strtosessiontrigger["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSessionTrigger() []SessionTrigger {
	vs := make([]SessionTrigger, 0, len(sessiontriggertostrs))
	for i := range sessiontriggertostrs {
		vs = append(vs, SessionTrigger(i))
	}
	return vs
}
