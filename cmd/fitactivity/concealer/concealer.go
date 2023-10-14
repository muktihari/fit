// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concealer

import (
	"github.com/muktihari/fit/cmd/fitactivity/finder"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
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

	sessionInfo := finder.FindFirstSessionInfo(fit)
	for i := sessionInfo.RecordFirstIndex; i <= sessionInfo.RecordLastIndex; i++ {
		if fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLat) == nil ||
			fit.Messages[i].FieldValueByNum(fieldnum.RecordPositionLong) == nil {
			continue
		}

		distance, ok := fit.Messages[i].FieldValueByNum(fieldnum.RecordDistance).(uint32)
		if !ok {
			continue
		}
		if distance > concealDistance {
			break
		}

		fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLat)
		fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLong)
	}

	return nil
}

// ConcealPositionEnd removes coordinates (lat, long) as far as end distance from the given fit file.
// If concealDistance == 0, it will not do anything, nil will be returned.
func ConcealPositionEnd(fit *proto.Fit, concealDistance uint32) error {
	if concealDistance == 0 {
		return nil
	}

	sessionInfo := finder.FindLastSessionInfo(fit)
	var lastDistance uint32
	for i := sessionInfo.RecordLastIndex; i >= sessionInfo.RecordFirstIndex; i-- {
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
			break
		}

		fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLat)
		fit.Messages[i].RemoveFieldByNum(fieldnum.RecordPositionLong)
	}

	return nil
}
