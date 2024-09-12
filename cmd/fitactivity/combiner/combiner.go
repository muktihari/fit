// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package combiner

import (
	"fmt"
	"time"

	"slices"

	"github.com/muktihari/fit/cmd/fitactivity/aggregator"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	errNoSessionFound = errorString("no session found")
	errSportMismatch  = errorString("sport mismatch")
)

// Combine combines multiple FIT activities into one continuous activity.
func Combine(fits []*proto.FIT) (result *proto.FIT, err error) {
	for i := 0; i < len(fits); i++ {
		if len(fits[i].Messages) == 0 {
			fits = append(fits[:i], fits[i+1:]...)
			i--
		}
	}

	slices.SortStableFunc(fits, func(f1, f2 *proto.FIT) int {
		timeCreated1 := f1.Messages[0].FieldValueByNum(fieldnum.FileIdTimeCreated).Uint32()
		timeCreated2 := f2.Messages[0].FieldValueByNum(fieldnum.FileIdTimeCreated).Uint32()
		if timeCreated1 < timeCreated2 {
			return -1
		} else if timeCreated1 > timeCreated2 {
			return 1
		}
		return 0
	})

	var (
		splitSummaries  []*mesgdef.SplitSummary
		sessionsByIndex = make([][]*mesgdef.Session, len(fits))
		activities      = make([]*mesgdef.Activity, 0, len(fits))
	)

	result = fits[0]

	for i, fit := range fits {
		var valid int
		for j, mesg := range fit.Messages {
			switch mesg.Num {
			case mesgnum.Session:
				sessionsByIndex[i] = append(sessionsByIndex[i], mesgdef.NewSession(&mesg))
				continue
			case mesgnum.SplitSummary:
				m := mesgdef.NewSplitSummary(&mesg)
				var ok bool
				for _, v := range splitSummaries {
					if v.SplitType == m.SplitType {
						aggregator.Aggregate(v, m)
						ok = true
						break
					}
				}
				if !ok {
					splitSummaries = append(splitSummaries, m)
				}
				continue
			case mesgnum.Activity:
				activities = append(activities, mesgdef.NewActivity(&mesg))
				continue
			}
			if j != valid {
				fit.Messages[valid], fit.Messages[j] = fit.Messages[j], fit.Messages[valid]
			}
			valid++
		}
		if len(sessionsByIndex[i]) == 0 {
			return nil, fmt.Errorf("fits[%d]: %w", i, errNoSessionFound)
		}
		fit.Messages = fit.Messages[:valid]
	}

	sessions := sessionsByIndex[0]
	lastPrevDistance := getLastDistanceOrZero(result.Messages)
	for i := 1; i < len(fits); i++ {
		var (
			nextFitSessions = sessionsByIndex[i]
			ses             = sessions[len(sessions)-1]
			nextSes         = nextFitSessions[0]
		)
		if ses.Sport != nextSes.Sport {
			return nil, fmt.Errorf("fits[%d] %q != %q: %w",
				i, ses.Sport, nextSes.Sport, errSportMismatch)
		}

		var lastDistance uint32
		for _, mesg := range fits[i].Messages {
			switch mesg.Num {
			case mesgnum.FileId, mesgnum.FileCreator:
				continue // skip
			case mesgnum.Record:
				// Accumulate distance
				field := mesg.FieldByNum(fieldnum.RecordDistance)
				if field != nil && field.Value.Uint32() != basetype.Uint32Invalid {
					lastDistance = field.Value.Uint32() + lastPrevDistance
					field.Value = proto.Uint32(lastDistance)
				}
				fallthrough
			default:
				result.Messages = append(result.Messages, mesg)
			}
		}
		lastPrevDistance = lastDistance

		// Add time gap
		endTime := ses.StartTime.Add(time.Duration(ses.TotalElapsedTime/1000) * time.Second)
		gap := uint32(nextSes.StartTime.Sub(endTime).Seconds() * 1000)
		ses.TotalElapsedTime += gap
		ses.TotalTimerTime += gap
		aggregator.Aggregate(ses, nextSes)

		if len(nextFitSessions) > 1 { // append the rest of the sessions
			sessions = append(sessions, nextFitSessions[1:]...)
		}
	}

	// Summarize

	firstTimestamp := getFirstTimestamp(result.Messages)
	lastTimestamp := getLastTimestamp(result.Messages)

	for _, v := range splitSummaries {
		mesg := v.ToMesg(nil)
		// Split Summary does not have timestamp, but we found a case where it may contains
		// timestamp on FIT files produced by Garmin Devices, so let's create one.
		mesg.Fields = append([]proto.Field{
			factory.CreateField(mesgnum.Session, proto.FieldNumTimestamp).WithValue(lastTimestamp),
		}, mesg.Fields...)
		result.Messages = append(result.Messages, mesg)
	}

	for _, v := range sessions {
		v.Timestamp = datetime.ToTime(lastTimestamp)
		result.Messages = append(result.Messages, v.ToMesg(nil))
	}

	timezone := getTimezone(activities)

	var activity *mesgdef.Activity
	if len(activities) > 0 {
		activity = activities[0]
	} else {
		activity = mesgdef.NewActivity(nil)
	}

	if activity.Type == typedef.Activity(basetype.EnumInvalid) {
		activity.Type = typedef.ActivityAutoMultiSport
	}

	activity.Timestamp = datetime.ToTime(lastTimestamp)
	activity.LocalTimestamp = activity.Timestamp.Add(time.Duration(timezone) * time.Hour)
	activity.TotalTimerTime = uint32((lastTimestamp - firstTimestamp) * 1000) // Scale: 1000
	activity.NumSessions = uint16(len(sessions))

	result.Messages = append(result.Messages, activity.ToMesg(nil))

	return result, nil
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

func getFirstTimestamp(mesgs []proto.Message) uint32 {
	for i := range mesgs {
		timestamp := mesgs[i].FieldValueByNum(proto.FieldNumTimestamp).Uint32()
		if timestamp != basetype.Uint32Invalid {
			return timestamp
		}
	}
	return basetype.Uint32Invalid
}

func getLastTimestamp(mesgs []proto.Message) uint32 {
	for i := len(mesgs) - 1; i >= 0; i-- {
		timestamp := mesgs[i].FieldValueByNum(proto.FieldNumTimestamp).Uint32()
		if timestamp != basetype.Uint32Invalid {
			return timestamp
		}
	}
	return basetype.Uint32Invalid
}

func getTimezone(activities []*mesgdef.Activity) (timezone int) {
	for _, activity := range activities {
		if activity.Timestamp.IsZero() || activity.LocalTimestamp.IsZero() {
			continue
		}
		timezone = datetime.TzOffsetHours(activity.LocalTimestamp, activity.Timestamp)
		if timezone != 0 {
			return timezone
		}
	}
	return timezone
}
