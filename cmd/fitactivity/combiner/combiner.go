// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package combiner

import (
	"errors"
	"fmt"

	"github.com/muktihari/fit/cmd/fitactivity/finder"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

var (
	ErrMinimalCombine = errors.New("minimal combine")
	ErrNoSessionFound = errors.New("no session found")
	ErrSportMismatch  = errors.New("sport mismatch")
)

const (
	fieldNumTimestamp = 253 // all message have the same field number for timestamp.
)

func Combine(fits ...proto.Fit) (*proto.Fit, error) {
	if len(fits) < 2 {
		return nil, fmt.Errorf("provide at least 2 fits to combine: %w", ErrMinimalCombine)
	}

	slices.SortStableFunc(fits, func(f1, f2 proto.Fit) int {
		if len(f1.Messages) == 0 || len(f2.Messages) == 0 {
			return 0
		}
		timeCreated1, _ := f1.Messages[0].FieldValueByNum(fieldnum.FileIdTimeCreated).(uint32)
		timeCreated2, _ := f2.Messages[0].FieldValueByNum(fieldnum.FileIdTimeCreated).(uint32)
		if timeCreated1 < timeCreated2 {
			return -1
		}
		return 1
	})

	result := &fits[0]

	for i := 1; i < len(fits); i++ {
		currentFIT := &fits[i-1]
		nextFIT := &fits[i]

		sessionInfo := finder.FindLastSessionInfo(currentFIT)
		if sessionInfo.SessionIndex == -1 {
			return nil, fmt.Errorf("could not find current last session message, fit index: %d: %w", i-1, ErrNoSessionFound)
		}
		sessionMesg := currentFIT.Messages[sessionInfo.SessionIndex]
		session := mesgdef.NewSession(&sessionMesg)

		if i-1 == 0 {
			// remove target session from result, session will be added later depend on sequence order.
			result.Messages = append(result.Messages[:sessionInfo.SessionIndex], result.Messages[sessionInfo.SessionIndex+1:]...)
		}

		nextSessionInfo := finder.FindFirstSessionInfo(nextFIT)
		if nextSessionInfo.SessionIndex == -1 {
			return nil, fmt.Errorf("could not find next first session message, fit index: %d: %w", i, ErrNoSessionFound)
		}
		nextSession := mesgdef.NewSession(&nextFIT.Messages[nextSessionInfo.SessionIndex])

		if session.Sport != nextSession.Sport {
			return nil, fmt.Errorf("last session's sport: %s, next first session's sport %s, fit index: %d: %w",
				session.Sport, nextSession.Sport, i, ErrSportMismatch)
		}

	loop:
		for j := nextSessionInfo.RecordFirstIndex; j <= nextSessionInfo.RecordLastIndex; j++ {
			mesg := nextFIT.Messages[j]
			switch mesg.Num {
			case mesgnum.FileId, mesgnum.FileCreator:
				continue loop // skip these messages
			case mesgnum.Record:
				fieldDistance := mesg.FieldByNum(fieldnum.RecordDistance)
				if fieldDistance == nil {
					continue loop
				}
				distance, ok := fieldDistance.Value.(uint32)
				if !ok {
					continue loop
				}
				distance += session.TotalDistance
				fieldDistance.Value = distance
			}
			result.Messages = append(result.Messages, mesg)
		}

		combineSession(session, nextSession)
		session.TotalElapsedTime = calculateSessionTotalElapsedTime(currentFIT, nextFIT, &sessionInfo, &nextSessionInfo)

		sessionMesg = session.ToMesg(nil)

		// Let's make "summary last sequence" order by updating session's timestamp with last lastRecord's timestamp
		lastRecord := nextFIT.Messages[nextSessionInfo.RecordLastIndex]
		newSessionTimestamp, ok := lastRecord.FieldValueByNum(fieldnum.RecordTimestamp).(uint32)
		if !ok {
			return nil, fmt.Errorf("timestamp is a required field but not present in record")
		}

		fieldTimestamp := sessionMesg.FieldByNum(fieldnum.SessionTimestamp)
		if fieldTimestamp == nil {
			return nil, fmt.Errorf("timestamp is a required field but not present in session")
		}

		fieldTimestamp.Value = newSessionTimestamp

		result.Messages = append(result.Messages, sessionMesg)
	}

	slices.SortStableFunc(result.Messages, func(mesg1, mesg2 proto.Message) int {
		timestamp1, ok := mesg1.FieldValueByNum(fieldNumTimestamp).(uint32)
		if !ok {
			return 0
		}
		timestamp2, ok := mesg2.FieldValueByNum(fieldNumTimestamp).(uint32)
		if !ok {
			return 0
		}
		if timestamp1 < timestamp2 {
			return -1
		}
		return 1
	})

	return result, nil
}

// combineSession combines s2 into s1. // s1.TotalElapsedTime will be calculated separately.
func combineSession(s1, s2 *mesgdef.Session) {
	s1.TotalTimerTime += s2.TotalTimerTime
	s1.TotalDistance += s2.TotalDistance
	s1.EndPositionLat = s2.EndPositionLat
	s1.EndPositionLong = s2.EndPositionLong

	if s1.TotalAscent != basetype.Uint16Invalid && s2.TotalAscent != basetype.Uint16Invalid {
		s1.TotalAscent += s2.TotalAscent
	} else if s1.TotalAscent == basetype.Uint16Invalid {
		s1.TotalAscent = s2.TotalAscent
	}

	if s1.TotalDescent != basetype.Uint16Invalid && s2.TotalDescent != basetype.Uint16Invalid {
		s1.TotalDescent += s2.TotalDescent
	} else if s1.TotalDescent == basetype.Uint16Invalid {
		s1.TotalDescent = s2.TotalDescent
	}

	if s1.TotalCycles != basetype.Uint32Invalid && s2.TotalCycles != basetype.Uint32Invalid {
		s1.TotalCycles += s2.TotalCycles
	} else if s1.TotalCycles == basetype.Uint32Invalid {
		s1.TotalCycles = s2.TotalCycles
	}

	if s1.TotalCalories != basetype.Uint16Invalid && s2.TotalCalories != basetype.Uint16Invalid {
		s1.TotalCalories += s2.TotalCalories
	} else if s1.TotalCalories == basetype.Uint16Invalid {
		s1.TotalCalories = s2.TotalCalories
	}

	if s1.AvgSpeed != basetype.Uint16Invalid && s2.AvgSpeed != basetype.Uint16Invalid {
		s1.AvgSpeed = (s1.AvgSpeed + s2.AvgSpeed) / 2
	} else if s1.AvgSpeed == basetype.Uint16Invalid {
		s1.AvgSpeed = s2.AvgSpeed
	}

	if s1.MaxSpeed != basetype.Uint16Invalid && s2.MaxSpeed != basetype.Uint16Invalid {
		if s1.MaxSpeed < s2.MaxSpeed {
			s1.MaxSpeed = s2.MaxSpeed
		}
	} else if s1.MaxSpeed == basetype.Uint16Invalid {
		s1.MaxSpeed = s2.MaxSpeed
	}

	if s1.AvgHeartRate != basetype.Uint8Invalid && s2.AvgHeartRate != basetype.Uint8Invalid {
		s1.AvgHeartRate = (s1.AvgHeartRate + s2.AvgHeartRate) / 2
	} else if s1.AvgHeartRate == basetype.Uint8Invalid {
		s1.AvgHeartRate = s2.AvgHeartRate
	}

	if s1.MaxHeartRate != basetype.Uint8Invalid && s2.MaxHeartRate != basetype.Uint8Invalid {
		if s1.MaxHeartRate < s2.MaxHeartRate {
			s1.MaxHeartRate = s2.MaxHeartRate
		}
	} else if s1.MaxHeartRate == basetype.Uint8Invalid {
		s1.MaxHeartRate = s2.MaxHeartRate
	}

	if s1.AvgCadence != basetype.Uint8Invalid && s2.AvgCadence != basetype.Uint8Invalid {
		s1.AvgCadence = (s1.AvgCadence + s2.AvgCadence) / 2
	} else if s1.AvgCadence == basetype.Uint8Invalid {
		s1.AvgCadence = s2.AvgCadence
	}

	if s1.MaxCadence != basetype.Uint8Invalid && s2.MaxCadence != basetype.Uint8Invalid {
		if s1.MaxCadence < s2.MaxCadence {
			s1.MaxCadence = s2.MaxCadence
		}
	} else if s1.MaxCadence == basetype.Uint8Invalid {
		s1.MaxCadence = s2.MaxCadence
	}

	if s1.AvgPower != basetype.Uint16Invalid && s2.AvgPower != basetype.Uint16Invalid {
		s1.AvgPower = (s1.AvgPower + s2.AvgPower) / 2
	} else if s1.AvgPower == basetype.Uint16Invalid {
		s1.AvgPower = s2.AvgPower
	}

	if s1.MaxPower != basetype.Uint16Invalid && s2.MaxPower != basetype.Uint16Invalid {
		if s1.MaxPower < s2.MaxPower {
			s1.MaxPower = s2.MaxPower
		}
	} else if s1.MaxPower == basetype.Uint16Invalid {
		s1.MaxPower = s2.MaxPower
	}

	if s1.AvgTemperature != basetype.Sint8Invalid && s2.AvgTemperature != basetype.Sint8Invalid {
		s1.AvgTemperature = (s1.AvgTemperature + s2.AvgTemperature) / 2
	} else if s1.AvgTemperature == basetype.Sint8Invalid {
		s1.AvgTemperature = s2.AvgTemperature
	}

	if s1.MaxTemperature != basetype.Sint8Invalid && s2.MaxTemperature != basetype.Sint8Invalid {
		if s1.MaxTemperature < s2.MaxTemperature {
			s1.MaxTemperature = s2.MaxTemperature
		}
	} else if s1.MaxTemperature == basetype.Sint8Invalid {
		s1.MaxTemperature = s2.MaxTemperature
	}

	if s1.AvgAltitude != basetype.Uint16Invalid && s2.AvgAltitude != basetype.Uint16Invalid {
		s1.AvgAltitude = (s1.AvgAltitude + s2.AvgAltitude) / 2
	} else if s1.AvgAltitude == basetype.Uint16Invalid {
		s1.AvgAltitude = s2.AvgAltitude
	}

	if s1.MaxAltitude != basetype.Uint16Invalid && s2.MaxAltitude != basetype.Uint16Invalid {
		if s1.MaxAltitude < s2.MaxAltitude {
			s1.MaxAltitude = s2.MaxAltitude
		}
	} else if s1.MaxAltitude == basetype.Uint16Invalid {
		s1.MaxAltitude = s2.MaxAltitude
	}
}

func calculateSessionTotalElapsedTime(currentFIT, nextFIT *proto.Fit,
	sessionInfo, nextSessionInfo *finder.SessionInfo) (totalElapsed uint32) {
	var firstSessionRecordTimestamp uint32
	var ok bool

	// Find first record of the current session that has timestamp
	for i := sessionInfo.RecordFirstIndex; i <= sessionInfo.RecordLastIndex; i++ {
		firstSessionRecordTimestamp, ok = currentFIT.Messages[i].FieldValueByNum(proto.FieldNumTimestamp).(uint32)
		if ok {
			break
		}
	}

	// Find last record of the next session that has timestamp
	var nextSessionLastRecordTimestamp uint32
	for i := nextSessionInfo.RecordLastIndex; i >= nextSessionInfo.RecordFirstIndex; i-- {
		nextSessionLastRecordTimestamp, ok = nextFIT.Messages[i].FieldValueByNum(proto.FieldNumTimestamp).(uint32)
		if ok {
			break
		}
	}

	const sessionTotalElapedTimeScale = 1000 // & session.TotalElapedTime's offset is 0
	totalElapsed = (nextSessionLastRecordTimestamp - firstSessionRecordTimestamp) * sessionTotalElapedTimeScale

	return totalElapsed
}
