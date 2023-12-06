// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ClimbProEvent byte

const (
	ClimbProEventApproach ClimbProEvent = 0
	ClimbProEventStart    ClimbProEvent = 1
	ClimbProEventComplete ClimbProEvent = 2
	ClimbProEventInvalid  ClimbProEvent = 0xFF // INVALID
)

var climbproeventtostrs = map[ClimbProEvent]string{
	ClimbProEventApproach: "approach",
	ClimbProEventStart:    "start",
	ClimbProEventComplete: "complete",
	ClimbProEventInvalid:  "invalid",
}

func (c ClimbProEvent) String() string {
	val, ok := climbproeventtostrs[c]
	if !ok {
		return strconv.Itoa(int(c))
	}
	return val
}

var strtoclimbproevent = func() map[string]ClimbProEvent {
	m := make(map[string]ClimbProEvent)
	for t, str := range climbproeventtostrs {
		m[str] = ClimbProEvent(t)
	}
	return m
}()

// FromString parse string into ClimbProEvent constant it's represent, return ClimbProEventInvalid if not found.
func ClimbProEventFromString(s string) ClimbProEvent {
	val, ok := strtoclimbproevent[s]
	if !ok {
		return strtoclimbproevent["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListClimbProEvent() []ClimbProEvent {
	vs := make([]ClimbProEvent, 0, len(climbproeventtostrs))
	for i := range climbproeventtostrs {
		vs = append(vs, ClimbProEvent(i))
	}
	return vs
}
