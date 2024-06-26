// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type EventType byte

const (
	EventTypeStart                  EventType = 0
	EventTypeStop                   EventType = 1
	EventTypeConsecutiveDepreciated EventType = 2
	EventTypeMarker                 EventType = 3
	EventTypeStopAll                EventType = 4
	EventTypeBeginDepreciated       EventType = 5
	EventTypeEndDepreciated         EventType = 6
	EventTypeEndAllDepreciated      EventType = 7
	EventTypeStopDisable            EventType = 8
	EventTypeStopDisableAll         EventType = 9
	EventTypeInvalid                EventType = 0xFF
)

func (e EventType) Byte() byte { return byte(e) }

func (e EventType) String() string {
	switch e {
	case EventTypeStart:
		return "start"
	case EventTypeStop:
		return "stop"
	case EventTypeConsecutiveDepreciated:
		return "consecutive_depreciated"
	case EventTypeMarker:
		return "marker"
	case EventTypeStopAll:
		return "stop_all"
	case EventTypeBeginDepreciated:
		return "begin_depreciated"
	case EventTypeEndDepreciated:
		return "end_depreciated"
	case EventTypeEndAllDepreciated:
		return "end_all_depreciated"
	case EventTypeStopDisable:
		return "stop_disable"
	case EventTypeStopDisableAll:
		return "stop_disable_all"
	default:
		return "EventTypeInvalid(" + strconv.Itoa(int(e)) + ")"
	}
}

// FromString parse string into EventType constant it's represent, return EventTypeInvalid if not found.
func EventTypeFromString(s string) EventType {
	switch s {
	case "start":
		return EventTypeStart
	case "stop":
		return EventTypeStop
	case "consecutive_depreciated":
		return EventTypeConsecutiveDepreciated
	case "marker":
		return EventTypeMarker
	case "stop_all":
		return EventTypeStopAll
	case "begin_depreciated":
		return EventTypeBeginDepreciated
	case "end_depreciated":
		return EventTypeEndDepreciated
	case "end_all_depreciated":
		return EventTypeEndAllDepreciated
	case "stop_disable":
		return EventTypeStopDisable
	case "stop_disable_all":
		return EventTypeStopDisableAll
	default:
		return EventTypeInvalid
	}
}

// List returns all constants.
func ListEventType() []EventType {
	return []EventType{
		EventTypeStart,
		EventTypeStop,
		EventTypeConsecutiveDepreciated,
		EventTypeMarker,
		EventTypeStopAll,
		EventTypeBeginDepreciated,
		EventTypeEndDepreciated,
		EventTypeEndAllDepreciated,
		EventTypeStopDisable,
		EventTypeStopDisableAll,
	}
}
