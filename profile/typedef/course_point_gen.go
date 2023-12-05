// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CoursePoint byte

const (
	CoursePointGeneric         CoursePoint = 0
	CoursePointSummit          CoursePoint = 1
	CoursePointValley          CoursePoint = 2
	CoursePointWater           CoursePoint = 3
	CoursePointFood            CoursePoint = 4
	CoursePointDanger          CoursePoint = 5
	CoursePointLeft            CoursePoint = 6
	CoursePointRight           CoursePoint = 7
	CoursePointStraight        CoursePoint = 8
	CoursePointFirstAid        CoursePoint = 9
	CoursePointFourthCategory  CoursePoint = 10
	CoursePointThirdCategory   CoursePoint = 11
	CoursePointSecondCategory  CoursePoint = 12
	CoursePointFirstCategory   CoursePoint = 13
	CoursePointHorsCategory    CoursePoint = 14
	CoursePointSprint          CoursePoint = 15
	CoursePointLeftFork        CoursePoint = 16
	CoursePointRightFork       CoursePoint = 17
	CoursePointMiddleFork      CoursePoint = 18
	CoursePointSlightLeft      CoursePoint = 19
	CoursePointSharpLeft       CoursePoint = 20
	CoursePointSlightRight     CoursePoint = 21
	CoursePointSharpRight      CoursePoint = 22
	CoursePointUTurn           CoursePoint = 23
	CoursePointSegmentStart    CoursePoint = 24
	CoursePointSegmentEnd      CoursePoint = 25
	CoursePointCampsite        CoursePoint = 27
	CoursePointAidStation      CoursePoint = 28
	CoursePointRestArea        CoursePoint = 29
	CoursePointGeneralDistance CoursePoint = 30 // Used with UpAhead
	CoursePointService         CoursePoint = 31
	CoursePointEnergyGel       CoursePoint = 32
	CoursePointSportsDrink     CoursePoint = 33
	CoursePointMileMarker      CoursePoint = 34
	CoursePointCheckpoint      CoursePoint = 35
	CoursePointShelter         CoursePoint = 36
	CoursePointMeetingSpot     CoursePoint = 37
	CoursePointOverlook        CoursePoint = 38
	CoursePointToilet          CoursePoint = 39
	CoursePointShower          CoursePoint = 40
	CoursePointGear            CoursePoint = 41
	CoursePointSharpCurve      CoursePoint = 42
	CoursePointSteepIncline    CoursePoint = 43
	CoursePointTunnel          CoursePoint = 44
	CoursePointBridge          CoursePoint = 45
	CoursePointObstacle        CoursePoint = 46
	CoursePointCrossing        CoursePoint = 47
	CoursePointStore           CoursePoint = 48
	CoursePointTransition      CoursePoint = 49
	CoursePointNavaid          CoursePoint = 50
	CoursePointTransport       CoursePoint = 51
	CoursePointAlert           CoursePoint = 52
	CoursePointInfo            CoursePoint = 53
	CoursePointInvalid         CoursePoint = 0xFF // INVALID
)

var coursepointtostrs = map[CoursePoint]string{
	CoursePointGeneric:         "generic",
	CoursePointSummit:          "summit",
	CoursePointValley:          "valley",
	CoursePointWater:           "water",
	CoursePointFood:            "food",
	CoursePointDanger:          "danger",
	CoursePointLeft:            "left",
	CoursePointRight:           "right",
	CoursePointStraight:        "straight",
	CoursePointFirstAid:        "first_aid",
	CoursePointFourthCategory:  "fourth_category",
	CoursePointThirdCategory:   "third_category",
	CoursePointSecondCategory:  "second_category",
	CoursePointFirstCategory:   "first_category",
	CoursePointHorsCategory:    "hors_category",
	CoursePointSprint:          "sprint",
	CoursePointLeftFork:        "left_fork",
	CoursePointRightFork:       "right_fork",
	CoursePointMiddleFork:      "middle_fork",
	CoursePointSlightLeft:      "slight_left",
	CoursePointSharpLeft:       "sharp_left",
	CoursePointSlightRight:     "slight_right",
	CoursePointSharpRight:      "sharp_right",
	CoursePointUTurn:           "u_turn",
	CoursePointSegmentStart:    "segment_start",
	CoursePointSegmentEnd:      "segment_end",
	CoursePointCampsite:        "campsite",
	CoursePointAidStation:      "aid_station",
	CoursePointRestArea:        "rest_area",
	CoursePointGeneralDistance: "general_distance",
	CoursePointService:         "service",
	CoursePointEnergyGel:       "energy_gel",
	CoursePointSportsDrink:     "sports_drink",
	CoursePointMileMarker:      "mile_marker",
	CoursePointCheckpoint:      "checkpoint",
	CoursePointShelter:         "shelter",
	CoursePointMeetingSpot:     "meeting_spot",
	CoursePointOverlook:        "overlook",
	CoursePointToilet:          "toilet",
	CoursePointShower:          "shower",
	CoursePointGear:            "gear",
	CoursePointSharpCurve:      "sharp_curve",
	CoursePointSteepIncline:    "steep_incline",
	CoursePointTunnel:          "tunnel",
	CoursePointBridge:          "bridge",
	CoursePointObstacle:        "obstacle",
	CoursePointCrossing:        "crossing",
	CoursePointStore:           "store",
	CoursePointTransition:      "transition",
	CoursePointNavaid:          "navaid",
	CoursePointTransport:       "transport",
	CoursePointAlert:           "alert",
	CoursePointInfo:            "info",
	CoursePointInvalid:         "invalid",
}

func (c CoursePoint) String() string {
	val, ok := coursepointtostrs[c]
	if !ok {
		return strconv.Itoa(int(c))
	}
	return val
}

var strtocoursepoint = func() map[string]CoursePoint {
	m := make(map[string]CoursePoint)
	for t, str := range coursepointtostrs {
		m[str] = CoursePoint(t)
	}
	return m
}()

// FromString parse string into CoursePoint constant it's represent, return CoursePointInvalid if not found.
func CoursePointFromString(s string) CoursePoint {
	val, ok := strtocoursepoint[s]
	if !ok {
		return strtocoursepoint["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListCoursePoint() []CoursePoint {
	vs := make([]CoursePoint, 0, len(coursepointtostrs))
	for i := range coursepointtostrs {
		vs = append(vs, CoursePoint(i))
	}
	return vs
}
