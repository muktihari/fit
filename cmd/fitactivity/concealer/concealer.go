// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concealer

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Conceal removes coordinates (lat, long) as far as start distance and end distance from the given FIT file.
// If startDistance and endDistance == 0, it will not do anything, nil will be returned.
// This will also remove StartPositionLat, StartPositionLong, EndPositionLat and EndPositionLong from any lap messages.
func Conceal(fit *proto.FIT, startDistance, endDistance uint32) error {
	if err := concealPositionStart(fit, startDistance); err != nil {
		return err
	}
	if err := concealPositionEnd(fit, endDistance); err != nil {
		return err
	}
	return nil
}

// concealPositionStart removes coordinates (lat, long) as far as start distance from the given FIT file.
// If concealDistance == 0, it will not do anything, nil will be returned.
func concealPositionStart(fit *proto.FIT, concealDistance uint32) error {
	if concealDistance == 0 {
		return nil
	}

	var sessionIndex = -1
	var newStartRecordIndex = -1

loop:
	for i := range fit.Messages {
		switch fit.Messages[i].Num {
		case mesgnum.Lap:
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapStartPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapStartPositionLong)
		case mesgnum.Session:
			if sessionIndex == -1 {
				sessionIndex = i
			}
		case mesgnum.Record:
			if fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat).Int32() == basetype.Sint32Invalid ||
				fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong).Int32() == basetype.Sint32Invalid {
				continue loop
			}

			distance := fit.Messages[i].FieldValueByNum(fieldnum.RecordDistance).Uint32()
			if distance != basetype.Uint32Invalid {
				continue loop
			}
			if distance > concealDistance {
				newStartRecordIndex = i
				break loop
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
		if sessionIndex == -1 {
			return nil // no session found to update
		}
	}

	if newStartRecordIndex == -1 { // all record are concealed
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLat)
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLong)
	} else {
		newLat, newLong := basetype.Sint32Invalid, basetype.Sint32Invalid
		for i := newStartRecordIndex; i < len(fit.Messages); i++ { // find new start record that has newLat and long
			newLat = fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat).Int32()
			newLong = fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong).Int32()
			if newLat != basetype.Sint32Invalid && newLong != basetype.Sint32Invalid {
				break
			}
		}

		fieldStartPositionLat := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionStartPositionLat)
		if fieldStartPositionLat == nil {
			fit.Messages[sessionIndex].Fields = append(fit.Messages[sessionIndex].Fields,
				factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLat),
			)
			lastIndex := len(fit.Messages[sessionIndex].Fields) - 1
			fieldStartPositionLat = &fit.Messages[sessionIndex].Fields[lastIndex]
		}
		fieldStartPositionLat.Value = proto.Int32(newLat)

		fieldStartPositionLong := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionStartPositionLong)
		if fieldStartPositionLong == nil {
			fit.Messages[sessionIndex].Fields = append(fit.Messages[sessionIndex].Fields,
				factory.CreateField(mesgnum.Session, fieldnum.SessionStartPositionLong),
			)
			lastIndex := len(fit.Messages[sessionIndex].Fields) - 1
			fieldStartPositionLong = &fit.Messages[sessionIndex].Fields[lastIndex]
		}
		fieldStartPositionLong.Value = proto.Int32(newLong)
	}

	return nil
}

// concealPositionEnd removes coordinates (lat, long) as far as end distance from the given FIT file.
// If concealDistance == 0, it will not do anything, nil will be returned.
func concealPositionEnd(fit *proto.FIT, concealDistance uint32) error {
	if concealDistance == 0 {
		return nil
	}

	var sessionIndex = -1
	var newEndRecordIndex = -1
	var lastDistance uint32

loop:
	for i := len(fit.Messages) - 1; i >= 0; i-- {
		switch fit.Messages[i].Num {
		case mesgnum.Lap:
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapEndPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.LapEndPositionLong)
		case mesgnum.Session:
			if sessionIndex == -1 {
				sessionIndex = i
			}
		case mesgnum.Record:
			if fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat).Int32() == basetype.Sint32Invalid ||
				fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong).Int32() == basetype.Sint32Invalid {
				continue loop
			}
			distance := fit.Messages[i].FieldValueByNum(fieldnum.RecordDistance).Uint32()
			if distance != basetype.Uint32Invalid {
				continue loop
			}

			if lastDistance == 0 { // first valid last distance
				lastDistance = distance
			}

			if lastDistance-distance > concealDistance {
				newEndRecordIndex = i
				break loop
			}

			fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLat)
			fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLong)
		}
	}

	if sessionIndex == -1 { // no session found during first iteration
		for i := newEndRecordIndex - 1; i >= 0; i-- { // find session to update
			if fit.Messages[i].Num == mesgnum.Session {
				sessionIndex = i
				break
			}
		}
		if sessionIndex == -1 {
			return nil // no session found to update
		}
	}

	if newEndRecordIndex == -1 { // all record are concealed
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLat)
		fit.Messages[sessionIndex].RemoveFieldByNum(fieldnum.SessionEndPositionLong)
	} else {
		newLat, newLong := basetype.Sint32Invalid, basetype.Sint32Invalid
		for i := newEndRecordIndex; i > 0; i++ { // find new end record that has newLat and long
			newLat = fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat).Int32()
			newLong = fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong).Int32()
			if newLat != basetype.Sint32Invalid && newLong != basetype.Sint32Invalid {
				break
			}
		}

		fieldEndPositionLat := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionEndPositionLat)
		if fieldEndPositionLat == nil {
			fit.Messages[sessionIndex].Fields = append(fit.Messages[sessionIndex].Fields,
				factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLat),
			)
			lastIndex := len(fit.Messages[sessionIndex].Fields) - 1
			fieldEndPositionLat = &fit.Messages[sessionIndex].Fields[lastIndex]
		}
		fieldEndPositionLat.Value = proto.Int32(newLat)

		fieldEndPositionLong := fit.Messages[sessionIndex].FieldByNum(fieldnum.SessionEndPositionLong)
		if fieldEndPositionLong == nil {
			fit.Messages[sessionIndex].Fields = append(fit.Messages[sessionIndex].Fields,
				factory.CreateField(mesgnum.Session, fieldnum.SessionEndPositionLong),
			)
			lastIndex := len(fit.Messages[sessionIndex].Fields) - 1
			fieldEndPositionLong = &fit.Messages[sessionIndex].Fields[lastIndex]
		}
		fieldEndPositionLong.Value = proto.Int32(newLong)
	}

	return nil
}
