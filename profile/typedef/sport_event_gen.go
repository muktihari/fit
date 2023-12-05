// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.118

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportEvent byte

const (
	SportEventUncategorized  SportEvent = 0
	SportEventGeocaching     SportEvent = 1
	SportEventFitness        SportEvent = 2
	SportEventRecreation     SportEvent = 3
	SportEventRace           SportEvent = 4
	SportEventSpecialEvent   SportEvent = 5
	SportEventTraining       SportEvent = 6
	SportEventTransportation SportEvent = 7
	SportEventTouring        SportEvent = 8
	SportEventInvalid        SportEvent = 0xFF // INVALID
)

var sporteventtostrs = map[SportEvent]string{
	SportEventUncategorized:  "uncategorized",
	SportEventGeocaching:     "geocaching",
	SportEventFitness:        "fitness",
	SportEventRecreation:     "recreation",
	SportEventRace:           "race",
	SportEventSpecialEvent:   "special_event",
	SportEventTraining:       "training",
	SportEventTransportation: "transportation",
	SportEventTouring:        "touring",
	SportEventInvalid:        "invalid",
}

func (s SportEvent) String() string {
	val, ok := sporteventtostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosportevent = func() map[string]SportEvent {
	m := make(map[string]SportEvent)
	for t, str := range sporteventtostrs {
		m[str] = SportEvent(t)
	}
	return m
}()

// FromString parse string into SportEvent constant it's represent, return SportEventInvalid if not found.
func SportEventFromString(s string) SportEvent {
	val, ok := strtosportevent[s]
	if !ok {
		return strtosportevent["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportEvent() []SportEvent {
	vs := make([]SportEvent, 0, len(sporteventtostrs))
	for i := range sporteventtostrs {
		vs = append(vs, SportEvent(i))
	}
	return vs
}
