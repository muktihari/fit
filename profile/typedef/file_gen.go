// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"fmt"
	"strconv"
)

type File byte

const (
	FileDevice           File = 1  // Read only, single file. Must be in root directory.
	FileSettings         File = 2  // Read/write, single file. Directory=Settings
	FileSport            File = 3  // Read/write, multiple files, file number = sport type. Directory=Sports
	FileActivity         File = 4  // Read/erase, multiple files. Directory=Activities
	FileWorkout          File = 5  // Read/write/erase, multiple files. Directory=Workouts
	FileCourse           File = 6  // Read/write/erase, multiple files. Directory=Courses
	FileSchedules        File = 7  // Read/write, single file. Directory=Schedules
	FileWeight           File = 9  // Read only, single file. Circular buffer. All message definitions at start of file. Directory=Weight
	FileTotals           File = 10 // Read only, single file. Directory=Totals
	FileGoals            File = 11 // Read/write, single file. Directory=Goals
	FileBloodPressure    File = 14 // Read only. Directory=Blood Pressure
	FileMonitoringA      File = 15 // Read only. Directory=Monitoring. File number=sub type.
	FileActivitySummary  File = 20 // Read/erase, multiple files. Directory=Activities
	FileMonitoringDaily  File = 28
	FileMonitoringB      File = 32   // Read only. Directory=Monitoring. File number=identifier
	FileSegment          File = 34   // Read/write/erase. Multiple Files. Directory=Segments
	FileSegmentList      File = 35   // Read/write/erase. Single File. Directory=Segments
	FileExdConfiguration File = 40   // Read/write/erase. Single File. Directory=Settings
	FileMfgRangeMin      File = 0xF7 // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileMfgRangeMax      File = 0xFE // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileInvalid          File = 0xFF // INVALID
)

var filetostrs = map[File]string{
	FileDevice:           "device",
	FileSettings:         "settings",
	FileSport:            "sport",
	FileActivity:         "activity",
	FileWorkout:          "workout",
	FileCourse:           "course",
	FileSchedules:        "schedules",
	FileWeight:           "weight",
	FileTotals:           "totals",
	FileGoals:            "goals",
	FileBloodPressure:    "blood_pressure",
	FileMonitoringA:      "monitoring_a",
	FileActivitySummary:  "activity_summary",
	FileMonitoringDaily:  "monitoring_daily",
	FileMonitoringB:      "monitoring_b",
	FileSegment:          "segment",
	FileSegmentList:      "segment_list",
	FileExdConfiguration: "exd_configuration",
	FileMfgRangeMin:      "mfg_range_min",
	FileMfgRangeMax:      "mfg_range_max",
	FileInvalid:          "invalid",
}

func (f File) String() string {
	val, ok := filetostrs[f]
	if !ok {
		return strconv.Itoa(int(f))
	}
	return val
}

var strtofile = func() map[string]File {
	m := make(map[string]File)
	for t, str := range filetostrs {
		m[str] = File(t)
	}
	return m
}()

// FromString parse string into File constant it's represent, return FileInvalid if not found.
func FileFromString(s string) File {
	val, ok := strtofile[s]
	if !ok {
		return strtofile["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListFile() []File {
	vs := make([]File, 0, len(filetostrs))
	for i := range filetostrs {
		vs = append(vs, File(i))
	}
	return vs
}

// FileRegister registers a manufacturer specific File so that the value can be recognized.
// It is recommended to define the constants somewhere else to track your own specifications.
//
// This is intended for those who prefer using this SDK as it is without the need to generate custom SDK using cmd/fitgen.
func FileRegister(v File, s string) error {
	if v >= FileInvalid {
		return fmt.Errorf("could not register outside max range: %d", FileInvalid)
	}

	if str, ok := filetostrs[v]; ok {
		return fmt.Errorf("could not register to an existing File: %d (%s)", v, str)
	}

	filetostrs[v] = s
	strtofile[s] = v

	return nil
}
