// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package combiner

import (
	"fmt"
	"time"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	ErrMinimalCombine = errorString("minimal combine")
	ErrNoSessionFound = errorString("no session found")
	ErrSportMismatch  = errorString("sport mismatch")
)

func Combine(fits ...proto.FIT) (*proto.FIT, error) {
	if len(fits) < 2 {
		return nil, fmt.Errorf("provide at least 2 fits to combine: %w", ErrMinimalCombine)
	}

	slices.SortStableFunc(fits, func(f1, f2 proto.FIT) int {
		if len(f1.Messages) == 0 || len(f2.Messages) == 0 {
			return 0
		}
		timeCreated1 := f1.Messages[0].FieldValueByNum(fieldnum.FileIdTimeCreated).Uint32()
		timeCreated2 := f2.Messages[0].FieldValueByNum(fieldnum.FileIdTimeCreated).Uint32()
		if timeCreated1 < timeCreated2 {
			return -1
		}
		return 1
	})

	var sessionMesgs []proto.Message
	var activityMesg proto.Message

	sessionsByFIT := make([][]proto.Message, len(fits))
	for i := range fits {
		fit := &fits[i]
		for j := 0; j < len(fit.Messages); j++ {
			mesg := fit.Messages[j]
			switch mesg.Num {
			case mesgnum.Session:
				if i == 0 { // First FIT as base
					sessionMesgs = append(sessionMesgs, mesg)
					fit.Messages = append(fit.Messages[:j], fit.Messages[j+1:]...) // remove all sessions from result
					j--
				}
				sessionsByFIT[i] = append(sessionsByFIT[i], mesg)
			case mesgnum.Activity:
				if activityMesg.Num != mesgnum.Activity {
					activityMesg = mesg
					fit.Messages = append(fit.Messages[:j], fit.Messages[j+1:]...) // remove activity from result
					j--
				}
			}
		}
	}

	fitResult := &fits[0]

	for i := range sessionsByFIT {
		if len(sessionsByFIT[i]) == 0 {
			return nil, fmt.Errorf("fits[i]: %w", ErrNoSessionFound)
		}
	}

	lastDistance := getLastDistanceOrZero(fitResult.Messages)
	for i := 1; i < len(fits); i++ {
		var (
			nextFitSessions = sessionsByFIT[i]
			curSes          = mesgdef.NewSession(&sessionMesgs[len(sessionMesgs)-1])
			nextSes         = mesgdef.NewSession(&nextFitSessions[0])
		)
		if curSes.Sport != nextSes.Sport {
			return nil, fmt.Errorf("fits[%d] %q != %q: %w",
				i, curSes.Sport, nextSes.Sport, ErrSportMismatch)
		}

		var lastDist uint32
		for _, mesg := range fits[i].Messages {
			switch mesg.Num {
			case mesgnum.FileId, mesgnum.FileCreator, mesgnum.Activity, mesgnum.Session:
				continue // skip
			case mesgnum.SplitSummary:
				continue // TODO: Still failed to upload to Garmin Connect if we include this message.
			case mesgnum.Record:
				// Accumulate distance
				field := mesg.FieldByNum(fieldnum.RecordDistance)
				if field != nil && field.Value.Uint32() != basetype.Uint32Invalid {
					lastDist = field.Value.Uint32() + lastDistance
					field.Value = proto.Uint32(lastDist)
				}
				fallthrough
			default:
				fitResult.Messages = append(fitResult.Messages, mesg)
			}
		}
		lastDistance = lastDist

		combineSession(curSes, nextSes)
		sessionMesgs[len(sessionMesgs)-1] = curSes.ToMesg(nil) // Update Session

		if len(nextFitSessions) > 1 { // append the rest of the sessions
			sessionMesgs = append(sessionMesgs, nextFitSessions[1:]...)
		}
	}

	// Now that all messages has been appended, let's update the session messages and
	// activity message and place it at the end of the resulting FIT (Sessions Last Summary).

	firstTimestamp := basetype.Uint32Invalid
	for _, mesg := range fitResult.Messages {
		if firstTimestamp == basetype.Uint32Invalid {
			timestamp := mesg.FieldValueByNum(proto.FieldNumTimestamp).Uint32()
			if timestamp != basetype.Uint32Invalid {
				firstTimestamp = timestamp
				break
			}
		}
	}

	lastTimestamp := basetype.Uint32Invalid
	for i := len(fitResult.Messages) - 1; i > 0; i-- {
		timestamp := fitResult.Messages[i].FieldValueByNum(proto.FieldNumTimestamp).Uint32()
		if timestamp != basetype.Uint32Invalid {
			lastTimestamp = timestamp
			break
		}
	}

	for _, sesMesg := range sessionMesgs {
		field := sesMesg.FieldByNum(proto.FieldNumTimestamp)
		if field == nil {
			sesMesg.Fields = append(sesMesg.Fields, factory.CreateField(mesgnum.Session, proto.FieldNumTimestamp))
			field = &sesMesg.Fields[len(sesMesg.Fields)-1]
		}
		field.Value = proto.Uint32(lastTimestamp) // Update session timestamp
		fitResult.Messages = append(fitResult.Messages, sesMesg)
	}

	// Update activity.Timestamp & activity.LocalTimestamp
	timestampField := activityMesg.FieldByNum(proto.FieldNumTimestamp)
	if timestampField == nil {
		activityMesg.Fields = append(activityMesg.Fields, factory.CreateField(mesgnum.Activity, proto.FieldNumTimestamp))
		timestampField = &activityMesg.Fields[len(activityMesg.Fields)-1]
	}

	localTimestampField := activityMesg.FieldByNum(fieldnum.ActivityLocalTimestamp)
	if localTimestampField == nil {
		activityMesg.Fields = append(activityMesg.Fields, factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp))
		localTimestampField = &activityMesg.Fields[len(activityMesg.Fields)-1]
	}

	tzOffsetHour := datetime.TzOffsetHours(
		datetime.ToTime(localTimestampField.Value.Uint32()),
		datetime.ToTime(timestampField.Value.Uint32()),
	)

	timestampField.Value = proto.Uint32(lastTimestamp)
	localTimestampField.Value = proto.Uint32(uint32(int64(lastTimestamp) + int64(tzOffsetHour*3600)))

	// Update activity.TotalTimerTime
	timestampField = activityMesg.FieldByNum(fieldnum.ActivityTotalTimerTime)
	if timestampField == nil {
		activityMesg.Fields = append(activityMesg.Fields, factory.CreateField(mesgnum.Activity, fieldnum.ActivityTotalTimerTime))
		timestampField = &activityMesg.Fields[len(activityMesg.Fields)-1]
	}
	timestampField.Value = proto.Uint32((lastTimestamp - firstTimestamp) * 1000) // Scale: 1000, Offset: 0

	// Update activity.NumSessions
	numSessions := activityMesg.FieldByNum(fieldnum.ActivityNumSessions)
	if numSessions == nil {
		activityMesg.Fields = append(activityMesg.Fields, factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions))
		numSessions = &activityMesg.Fields[len(activityMesg.Fields)-1]
	}
	numSessions.Value = proto.Uint16(uint16(len(sessionMesgs)))

	fitResult.Messages = append(fitResult.Messages, activityMesg)

	return fitResult, nil
}

func getLastDistanceOrZero(mesgs []proto.Message) uint32 {
	for i := len(mesgs) - 1; i >= 0; i-- {
		if mesgs[i].Num != mesgnum.Record {
			continue
		}
		v := mesgs[i].FieldValueByNum(fieldnum.RecordDistance).Uint32()
		if v == basetype.Uint32Invalid {
			continue
		}
		return v
	}
	return 0
}

// combineSession combines s2 into s1.
func combineSession(s1, s2 *mesgdef.Session) {
	s1EndTime := s1.StartTime.Add(time.Duration(s1.TotalElapsedTime/1000) * time.Second)
	gap := s2.StartTime.Sub(s1EndTime).Seconds() * 1000
	s1.TotalElapsedTime += uint32(gap)
	s1.TotalTimerTime += uint32(gap)

	s1.TotalElapsedTime += s2.TotalElapsedTime
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
