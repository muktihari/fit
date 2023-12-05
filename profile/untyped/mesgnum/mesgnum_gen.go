// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mesgnum contains untyped constants for greater flexibility, intended to simplify code typing when creating messages and fields
// using the factory and reduce human error.
//
// This should not be confused with typed constants in typedef. The value of these untyped constants remains the same as declared in typedef.
// The only difference is that untyped constants can take on many forms (byte, int, types.DateTime, types.File, etc.),
//
// For more information about Go constants, visit: https://go.dev/blog/constants.
package mesgnum

const (
	FileId                      = 0
	Capabilities                = 1
	DeviceSettings              = 2
	UserProfile                 = 3
	HrmProfile                  = 4
	SdmProfile                  = 5
	BikeProfile                 = 6
	ZonesTarget                 = 7
	HrZone                      = 8
	PowerZone                   = 9
	MetZone                     = 10
	Sport                       = 12
	Goal                        = 15
	Session                     = 18
	Lap                         = 19
	Record                      = 20
	Event                       = 21
	DeviceInfo                  = 23
	Workout                     = 26
	WorkoutStep                 = 27
	Schedule                    = 28
	WeightScale                 = 30
	Course                      = 31
	CoursePoint                 = 32
	Totals                      = 33
	Activity                    = 34
	Software                    = 35
	FileCapabilities            = 37
	MesgCapabilities            = 38
	FieldCapabilities           = 39
	FileCreator                 = 49
	BloodPressure               = 51
	SpeedZone                   = 53
	Monitoring                  = 55
	TrainingFile                = 72
	Hrv                         = 78
	AntRx                       = 80
	AntTx                       = 81
	AntChannelId                = 82
	Length                      = 101
	MonitoringInfo              = 103
	Pad                         = 105
	SlaveDevice                 = 106
	Connectivity                = 127
	WeatherConditions           = 128
	WeatherAlert                = 129
	CadenceZone                 = 131
	Hr                          = 132
	SegmentLap                  = 142
	MemoGlob                    = 145
	SegmentId                   = 148
	SegmentLeaderboardEntry     = 149
	SegmentPoint                = 150
	SegmentFile                 = 151
	WorkoutSession              = 158
	WatchfaceSettings           = 159
	GpsMetadata                 = 160
	CameraEvent                 = 161
	TimestampCorrelation        = 162
	GyroscopeData               = 164
	AccelerometerData           = 165
	ThreeDSensorCalibration     = 167
	VideoFrame                  = 169
	ObdiiData                   = 174
	NmeaSentence                = 177
	AviationAttitude            = 178
	Video                       = 184
	VideoTitle                  = 185
	VideoDescription            = 186
	VideoClip                   = 187
	OhrSettings                 = 188
	ExdScreenConfiguration      = 200
	ExdDataFieldConfiguration   = 201
	ExdDataConceptConfiguration = 202
	FieldDescription            = 206
	DeveloperDataId             = 207
	MagnetometerData            = 208
	BarometerData               = 209
	OneDSensorCalibration       = 210
	MonitoringHrData            = 211
	TimeInZone                  = 216
	Set                         = 225
	StressLevel                 = 227
	MaxMetData                  = 229
	DiveSettings                = 258
	DiveGas                     = 259
	DiveAlarm                   = 262
	ExerciseTitle               = 264
	DiveSummary                 = 268
	Spo2Data                    = 269
	SleepLevel                  = 275
	Jump                        = 285
	BeatIntervals               = 290
	RespirationRate             = 297
	Split                       = 312
	ClimbPro                    = 317
	TankUpdate                  = 319
	TankSummary                 = 323
	SleepAssessment             = 346
	HrvStatusSummary            = 370
	HrvValue                    = 371
	DeviceAuxBatteryInfo        = 375
	DiveApneaAlarm              = 393
	MfgRangeMin                 = 0xFF00 // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MfgRangeMax                 = 0xFFFE // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	Invalid                     = 0xFFFF // INVALID
)
