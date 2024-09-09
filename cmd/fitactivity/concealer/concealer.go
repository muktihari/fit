// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concealer

import (
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Conceal removes Latitude and Longitude data of first and last <inputs> meters.
// Affected Laps and Sessions will be updated accordingly.
func Conceal(mesgs []proto.Message, first, last uint32) {
	var lapIndices, sessionIndices []int
	for i := range mesgs {
		switch mesgs[i].Num {
		case mesgnum.Lap:
			lapIndices = append(lapIndices, i)
		case mesgnum.Session:
			sessionIndices = append(sessionIndices, i)
		}
	}
	lastIndex := concealStartPosition(mesgs, lapIndices, sessionIndices, first)
	concealEndPosition(mesgs, lapIndices, sessionIndices, lastIndex, last)
}

// placeholder for replacing Lap and Session field numbers since both share
// the same field name but have different field numbers.
type placeholder struct {
	startTime         byte
	totalTimerTime    byte
	startPositionLat  byte
	startPositionLong byte
	endPositionLat    byte
	endPositionLong   byte
}

var lapPlaceholder = placeholder{
	startTime:         fieldnum.LapStartTime,
	totalTimerTime:    fieldnum.LapTotalTimerTime,
	startPositionLat:  fieldnum.LapStartPositionLat,
	startPositionLong: fieldnum.LapStartPositionLong,
	endPositionLat:    fieldnum.LapEndPositionLat,
	endPositionLong:   fieldnum.LapEndPositionLong,
}

var sessionPlaceholder = placeholder{
	startTime:         fieldnum.SessionStartTime,
	totalTimerTime:    fieldnum.SessionTotalTimerTime,
	startPositionLat:  fieldnum.SessionStartPositionLat,
	startPositionLong: fieldnum.SessionStartPositionLong,
	endPositionLat:    fieldnum.SessionEndPositionLat,
	endPositionLong:   fieldnum.SessionEndPositionLong,
}

func concealStartPosition(mesgs []proto.Message, lapIndices, sessionIndices []int, threshold uint32) (lastConcealStartIndex int) {
	if threshold == 0 {
		return 0
	}

	lastConcealStartIndex = -1
	for i := range mesgs {
		mesg := &mesgs[i]
		if mesg.Num == mesgnum.Record {
			d := mesg.FieldValueByNum(fieldnum.RecordDistance).Uint32()
			if d < threshold {
				mesg.RemoveFieldByNum(fieldnum.RecordPositionLat)
				mesg.RemoveFieldByNum(fieldnum.RecordPositionLong)
				continue
			}
			lastConcealStartIndex = i
			break
		}
	}

	updateStartPosition(mesgs, lapIndices, lapPlaceholder, lastConcealStartIndex)         // Update Laps
	updateStartPosition(mesgs, sessionIndices, sessionPlaceholder, lastConcealStartIndex) // Update Sessions

	return lastConcealStartIndex
}

// updateStartPosition update start position of Laps or Sessions.
func updateStartPosition(mesgs []proto.Message, indices []int, ph placeholder, recordIndex int) {
	var rec proto.Message
	if recordIndex != -1 {
		rec = mesgs[recordIndex]
	}

	var (
		recTimestamp = rec.FieldValueByNum(fieldnum.RecordTimestamp).Uint32()
		recLat       = rec.FieldValueByNum(fieldnum.RecordPositionLat).Int32()
		recLong      = rec.FieldValueByNum(fieldnum.RecordPositionLong).Int32()
	)
	for i := range indices {
		var (
			mesg           = &mesgs[indices[i]]
			startTime      = mesg.FieldValueByNum(ph.startTime).Uint32()
			totalTimerTime = mesg.FieldValueByNum(ph.totalTimerTime).Uint32()
		)
		if startTime == basetype.Uint32Invalid || totalTimerTime == basetype.Uint32Invalid ||
			startTime+totalTimerTime < recTimestamp {
			mesg.RemoveFieldByNum(ph.startPositionLat)
			mesg.RemoveFieldByNum(ph.startPositionLong)
			mesg.RemoveFieldByNum(ph.endPositionLat)
			mesg.RemoveFieldByNum(ph.endPositionLong)
			continue
		}

		if recLat == basetype.Sint32Invalid {
			mesg.RemoveFieldByNum(ph.startPositionLat)
		} else if field := mesg.FieldByNum(ph.startPositionLat); field != nil {
			field.Value = proto.Int32(recLat)
		}

		if recLong == basetype.Sint32Invalid {
			mesg.RemoveFieldByNum(ph.startPositionLong)
		} else if field := mesg.FieldByNum(ph.startPositionLong); field != nil {
			field.Value = proto.Int32(recLong)
		}
		break
	}
}

func concealEndPosition(mesgs []proto.Message, lapIndices, sessionIndices []int, lastConcealStartIndex int, threshold uint32) {
	if threshold == 0 {
		return
	}

	var (
		lastConcealEndIndex = -1
		lastRecDist         = basetype.Uint32Invalid
	)
	for i := len(mesgs) - 1; i >= 0; i-- {
		mesg := &mesgs[i]
		if mesg.Num == mesgnum.Record {
			d := mesg.FieldValueByNum(fieldnum.RecordDistance).Uint32()
			if lastRecDist == basetype.Uint32Invalid {
				lastRecDist = d
			}
			if lastRecDist-d < threshold {
				mesg.RemoveFieldByNum(fieldnum.RecordPositionLat)
				mesg.RemoveFieldByNum(fieldnum.RecordPositionLong)
				continue
			}
			lastConcealEndIndex = i
			break
		}
	}

	updateEndPosition(mesgs, lapIndices, lapPlaceholder, lastConcealStartIndex, lastConcealEndIndex)         // Update Laps
	updateEndPosition(mesgs, sessionIndices, sessionPlaceholder, lastConcealStartIndex, lastConcealEndIndex) // Update Sessions
}

// updateEndPosition update end position of Laps or Sessions.
func updateEndPosition(mesgs []proto.Message, indices []int, ph placeholder, lastConcealStartIndex, recordIndex int) {
	var rec proto.Message
	if recordIndex != -1 {
		rec = mesgs[recordIndex]
	}

	var (
		recTimestamp = rec.FieldValueByNum(fieldnum.RecordTimestamp).Uint32()
		recLat       = rec.FieldValueByNum(fieldnum.RecordPositionLat).Int32()
		recLong      = rec.FieldValueByNum(fieldnum.RecordPositionLong).Int32()
	)
	for i := len(indices) - 1; i >= 0; i-- {
		var (
			mesg      = &mesgs[indices[i]]
			startTime = mesg.FieldValueByNum(ph.startTime).Uint32()
		)
		if startTime == basetype.Uint32Invalid || startTime > recTimestamp {
			mesg.RemoveFieldByNum(ph.startPositionLat)
			mesg.RemoveFieldByNum(ph.startPositionLong)
			mesg.RemoveFieldByNum(ph.endPositionLat)
			mesg.RemoveFieldByNum(ph.endPositionLong)
			continue
		}

		if lastConcealStartIndex > recordIndex { // Overlap
			mesg.RemoveFieldByNum(ph.startPositionLat)
			mesg.RemoveFieldByNum(ph.startPositionLong)
		}

		if recLat == basetype.Sint32Invalid {
			mesg.RemoveFieldByNum(ph.endPositionLat)
		} else if field := mesg.FieldByNum(ph.endPositionLat); field != nil {
			field.Value = proto.Int32(recLat)
		}

		if recLong == basetype.Sint32Invalid {
			mesg.RemoveFieldByNum(ph.endPositionLong)
		} else if field := mesg.FieldByNum(ph.endPositionLong); field != nil {
			field.Value = proto.Int32(recLong)
		}
		break
	}
}
