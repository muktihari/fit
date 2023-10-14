// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package finder

import (
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

const (
	fieldNumTimestamp = 253 // all message have the same field number for timestamp.
)

type SessionInfo struct {
	SessionIndex     int
	RecordFirstIndex int
	RecordLastIndex  int
}

// FindFirstSessionInfo finds the session and records info of the first session found.
func FindFirstSessionInfo(fit *proto.Fit) SessionInfo {
	res := SessionInfo{
		SessionIndex:     -1,
		RecordFirstIndex: -1,
		RecordLastIndex:  -1,
	}

	var startTime, endTime *uint32
	for i := range fit.Messages {
		if fit.Messages[i].Num == mesgnum.Session {
			sessionStartTime, _ := fit.Messages[i].FieldValueByNum(fieldnum.SessionStartTime).(uint32)
			if startTime == nil { // first session
				startTime = &sessionStartTime
				res.SessionIndex = i
			} else if endTime == nil { // session next to the first session
				endTime = &sessionStartTime
				break
			}
		}
	}

	// Find records info of the corresponding session: between startTime and endTime
	for i := range fit.Messages {
		switch fit.Messages[i].Num {
		case mesgnum.Record, mesgnum.Event, mesgnum.Lap:
			timestamp, _ := fit.Messages[i].FieldValueByNum(fieldNumTimestamp).(uint32)
			if timestamp < *startTime {
				continue
			}
			if endTime != nil && timestamp > *endTime {
				break
			}

			if res.RecordFirstIndex == -1 {
				res.RecordFirstIndex = i
			}
			res.RecordLastIndex = i
		}
	}

	return res
}

// FindLastSessionInfo finds the session and records info of the last session found.
func FindLastSessionInfo(fit *proto.Fit) SessionInfo {
	res := SessionInfo{
		SessionIndex:     -1,
		RecordFirstIndex: -1,
		RecordLastIndex:  -1,
	}

	var startTime *uint32
	for i := len(fit.Messages) - 1; i > 0; i-- {
		if fit.Messages[i].Num == mesgnum.Session {
			sessionStartTime, _ := fit.Messages[i].FieldValueByNum(fieldnum.SessionStartTime).(uint32)
			if startTime == nil { // last session
				startTime = &sessionStartTime
				res.SessionIndex = i
				break
			}
		}
	}

	// Find records info of the corresponding session's startTime to the end of file since it's the last session.
	for i := range fit.Messages {
		switch fit.Messages[i].Num {
		case mesgnum.Record, mesgnum.Event, mesgnum.Lap:
			timestamp, _ := fit.Messages[i].FieldValueByNum(fieldNumTimestamp).(uint32)
			if timestamp < *startTime {
				continue
			}

			if res.RecordFirstIndex == -1 {
				res.RecordFirstIndex = i
			}
			res.RecordLastIndex = i
		}
	}

	return res
}
