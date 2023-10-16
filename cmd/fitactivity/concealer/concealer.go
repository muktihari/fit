// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concealer

import (
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// ConcealPosition removes coordinates (lat, long) as far as start distance and end distance from the given fit file.
// If startDistance and endDistance == 0, it will not do anything, nil will be returned.
func ConcealPosition(fit *proto.Fit, startDistance, endDistance uint32) error {
	if err := ConcealPositionStart(fit, startDistance); err != nil {
		return err
	}
	return ConcealPositionEnd(fit, endDistance)
}

// ConcealPositionStart removes coordinates (lat, long) as far as start distance from the given fit file.
// If concealDistance == 0, it will not do anything, nil will be returned.
func ConcealPositionStart(fit *proto.Fit, concealDistance uint32) error {
	if concealDistance == 0 {
		return nil
	}

	var sessionIndex = -1
	var newStartRecordIndex = -1
	for i := range fit.Messages {
		switch fit.Messages[i].Num {
		case mesgnum.Session:
			if sessionIndex == -1 {
				sessionIndex = i
			}
		case mesgnum.Record:
			if fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat) == nil ||
				fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong) == nil {
				continue
			}

			distance, ok := fit.Messages[i].FieldValueByNum(fieldnum.RecordDistance).(uint32)
			if !ok {
				continue
			}
			if distance > concealDistance {
				newStartRecordIndex = i
				break
			}

			fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLong)
		}
	}

	if sessionIndex == -1 { // no session found during first iteration
		for i := newStartRecordIndex + 1; i < len(fit.Messages); i++ {
			if fit.Messages[i].Num == mesgnum.Session {
				sessionIndex = i
				break
			}
		}
	}

	if newStartRecordIndex == -1 { // all record are concealed
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLat)
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLong)
	} else {
		newLat, newLong := basetype.Sint32Invalid, basetype.Sint32Invalid
		for i := newStartRecordIndex; i < len(fit.Messages); i++ { // find new start record that has lat and long
			lat := fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat)
			long := fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong)
			if lat != nil && long != nil {
				newLat, _ = lat.(int32)
				newLong, _ = long.(int32)
				break
			}
		}

		fieldEndPositionLat := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionEndPositionLat)
		if fieldEndPositionLat != nil {
			fieldEndPositionLat.Value = newLat
		}

		fieldEndPositionLong := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionEndPositionLong)
		if fieldEndPositionLong != nil {
			fieldEndPositionLong.Value = newLong
		}
	}

	return nil
}

// ConcealPositionEnd removes coordinates (lat, long) as far as end distance from the given fit file.
// If concealDistance == 0, it will not do anything, nil will be returned.
func ConcealPositionEnd(fit *proto.Fit, concealDistance uint32) error {
	if concealDistance == 0 {
		return nil
	}

	var sessionIndex = -1
	var newEndRecordIndex = -1
	var lastDistance uint32
	for i := len(fit.Messages) - 1; i >= 0; i-- {
		if fit.Messages[i].Num != mesgnum.Record {
			continue
		}
		switch fit.Messages[i].Num {
		case mesgnum.Session:
			if sessionIndex == -1 {
				sessionIndex = i
			}
		case mesgnum.Record:
			if fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat) == nil ||
				fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong) == nil {
				continue
			}
			distance, ok := fit.Messages[i].FieldValueByNum(fieldnum.RecordDistance).(uint32)
			if !ok {
				continue
			}

			if lastDistance == 0 { // first valid last distance
				lastDistance = distance
			}

			if lastDistance-distance > concealDistance {
				newEndRecordIndex = i
				break
			}

			fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLong)
		}
	}

	if sessionIndex == -1 { // no session found during first iteration
		for i := newEndRecordIndex - 1; i > 0; i-- { // find session to update
			if fit.Messages[i].Num == mesgnum.Session {
				sessionIndex = i
				break
			}
		}
	}

	if newEndRecordIndex == -1 { // all record concelead
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLat)
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLong)
	} else {
		newLat, newLong := basetype.Sint32Invalid, basetype.Sint32Invalid
		for i := newEndRecordIndex; i > 0; i++ { // find new end record that has lat and long
			lat := fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat)
			long := fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong)
			if lat != nil && long != nil {
				newLat, _ = lat.(int32)
				newLong, _ = long.(int32)
				break
			}
		}

		fieldEndPositionLat := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionEndPositionLat)
		if fieldEndPositionLat != nil {
			fieldEndPositionLat.Value = newLat
		}

		fieldEndPositionLong := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionEndPositionLong)
		if fieldEndPositionLong != nil {
			fieldEndPositionLong.Value = newLong
		}
	}

	return nil
}

// ConcealLapStartAndEndPosition removes StartPositionLat, StartPositionLong, EndPositionLat and EndPositionLong from any lap messages.
func ConcealLapStartAndEndPosition(fit *proto.Fit) {
	for i := range fit.Messages {
		if fit.Messages[i].Num == mesgnum.Lap {
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapStartPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapStartPositionLong)
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapEndPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapEndPositionLong)
		}
	}
}
