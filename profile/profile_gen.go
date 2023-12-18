// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

import (
	"strconv"

	"github.com/muktihari/fit/profile/basetype"
)

type ProfileType uint16

const (
	Enum ProfileType = iota
	Sint8
	Uint8
	Sint16
	Uint16
	Sint32
	Uint32
	String
	Float32
	Float64
	Uint8z
	Uint16z
	Uint32z
	Byte
	Sint64
	Uint64
	Uint64z
	Bool
	File
	MesgNum
	Checksum
	FileFlags
	MesgCount
	DateTime
	LocalDateTime
	MessageIndex
	DeviceIndex
	Gender
	Language
	LanguageBits0
	LanguageBits1
	LanguageBits2
	LanguageBits3
	LanguageBits4
	TimeZone
	DisplayMeasure
	DisplayHeart
	DisplayPower
	DisplayPosition
	Switch
	Sport
	SportBits0
	SportBits1
	SportBits2
	SportBits3
	SportBits4
	SportBits5
	SportBits6
	SubSport
	SportEvent
	Activity
	Intensity
	SessionTrigger
	AutolapTrigger
	LapTrigger
	TimeMode
	BacklightMode
	DateMode
	BacklightTimeout
	Event
	EventType
	TimerTrigger
	FitnessEquipmentState
	Tone
	Autoscroll
	ActivityClass
	HrZoneCalc
	PwrZoneCalc
	WktStepDuration
	WktStepTarget
	Goal
	GoalRecurrence
	GoalSource
	Schedule
	CoursePoint
	Manufacturer
	GarminProduct
	AntplusDeviceType
	AntNetwork
	WorkoutCapabilities
	BatteryStatus
	HrType
	CourseCapabilities
	Weight
	WorkoutHr
	WorkoutPower
	BpStatus
	UserLocalId
	SwimStroke
	ActivityType
	ActivitySubtype
	ActivityLevel
	Side
	LeftRightBalance
	LeftRightBalance100
	LengthType
	DayOfWeek
	ConnectivityCapabilities
	WeatherReport
	WeatherStatus
	WeatherSeverity
	WeatherSevereType
	TimeIntoDay
	LocaltimeIntoDay
	StrokeType
	BodyLocation
	SegmentLapStatus
	SegmentLeaderboardType
	SegmentDeleteStatus
	SegmentSelectionType
	SourceType
	LocalDeviceType
	BleDeviceType
	AntChannelId
	DisplayOrientation
	WorkoutEquipment
	WatchfaceMode
	DigitalWatchfaceLayout
	AnalogWatchfaceLayout
	RiderPositionType
	PowerPhaseType
	CameraEventType
	SensorType
	BikeLightNetworkConfigType
	CommTimeoutType
	CameraOrientationType
	AttitudeStage
	AttitudeValidity
	AutoSyncFrequency
	ExdLayout
	ExdDisplayType
	ExdDataUnits
	ExdQualifiers
	ExdDescriptors
	AutoActivityDetect
	SupportedExdScreenLayouts
	FitBaseType
	TurnType
	BikeLightBeamAngleMode
	FitBaseUnit
	SetType
	MaxMetCategory
	ExerciseCategory
	BenchPressExerciseName
	CalfRaiseExerciseName
	CardioExerciseName
	CarryExerciseName
	ChopExerciseName
	CoreExerciseName
	CrunchExerciseName
	CurlExerciseName
	DeadliftExerciseName
	FlyeExerciseName
	HipRaiseExerciseName
	HipStabilityExerciseName
	HipSwingExerciseName
	HyperextensionExerciseName
	LateralRaiseExerciseName
	LegCurlExerciseName
	LegRaiseExerciseName
	LungeExerciseName
	OlympicLiftExerciseName
	PlankExerciseName
	PlyoExerciseName
	PullUpExerciseName
	PushUpExerciseName
	RowExerciseName
	ShoulderPressExerciseName
	ShoulderStabilityExerciseName
	ShrugExerciseName
	SitUpExerciseName
	SquatExerciseName
	TotalBodyExerciseName
	TricepsExtensionExerciseName
	WarmUpExerciseName
	RunExerciseName
	WaterType
	TissueModelType
	DiveGasStatus
	DiveAlert
	DiveAlarmType
	DiveBacklightMode
	SleepLevel
	Spo2MeasurementType
	CcrSetpointSwitchMode
	DiveGasMode
	FaveroProduct
	SplitType
	ClimbProEvent
	GasConsumptionRateType
	TapSensitivity
	RadarThreatLevelType
	MaxMetSpeedSource
	MaxMetHeartRateSource
	HrvStatus
	NoFlyTimeMode
	Invalid
)

func (p ProfileType) String() string {
	switch p {
	case Enum:
		return "enum"
	case Sint8:
		return "sint8"
	case Uint8:
		return "uint8"
	case Sint16:
		return "sint16"
	case Uint16:
		return "uint16"
	case Sint32:
		return "sint32"
	case Uint32:
		return "uint32"
	case String:
		return "string"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Uint8z:
		return "uint8z"
	case Uint16z:
		return "uint16z"
	case Uint32z:
		return "uint32z"
	case Byte:
		return "byte"
	case Sint64:
		return "sint64"
	case Uint64:
		return "uint64"
	case Uint64z:
		return "uint64z"
	case Bool:
		return "bool"
	case File:
		return "file"
	case MesgNum:
		return "mesg_num"
	case Checksum:
		return "checksum"
	case FileFlags:
		return "file_flags"
	case MesgCount:
		return "mesg_count"
	case DateTime:
		return "date_time"
	case LocalDateTime:
		return "local_date_time"
	case MessageIndex:
		return "message_index"
	case DeviceIndex:
		return "device_index"
	case Gender:
		return "gender"
	case Language:
		return "language"
	case LanguageBits0:
		return "language_bits_0"
	case LanguageBits1:
		return "language_bits_1"
	case LanguageBits2:
		return "language_bits_2"
	case LanguageBits3:
		return "language_bits_3"
	case LanguageBits4:
		return "language_bits_4"
	case TimeZone:
		return "time_zone"
	case DisplayMeasure:
		return "display_measure"
	case DisplayHeart:
		return "display_heart"
	case DisplayPower:
		return "display_power"
	case DisplayPosition:
		return "display_position"
	case Switch:
		return "switch"
	case Sport:
		return "sport"
	case SportBits0:
		return "sport_bits_0"
	case SportBits1:
		return "sport_bits_1"
	case SportBits2:
		return "sport_bits_2"
	case SportBits3:
		return "sport_bits_3"
	case SportBits4:
		return "sport_bits_4"
	case SportBits5:
		return "sport_bits_5"
	case SportBits6:
		return "sport_bits_6"
	case SubSport:
		return "sub_sport"
	case SportEvent:
		return "sport_event"
	case Activity:
		return "activity"
	case Intensity:
		return "intensity"
	case SessionTrigger:
		return "session_trigger"
	case AutolapTrigger:
		return "autolap_trigger"
	case LapTrigger:
		return "lap_trigger"
	case TimeMode:
		return "time_mode"
	case BacklightMode:
		return "backlight_mode"
	case DateMode:
		return "date_mode"
	case BacklightTimeout:
		return "backlight_timeout"
	case Event:
		return "event"
	case EventType:
		return "event_type"
	case TimerTrigger:
		return "timer_trigger"
	case FitnessEquipmentState:
		return "fitness_equipment_state"
	case Tone:
		return "tone"
	case Autoscroll:
		return "autoscroll"
	case ActivityClass:
		return "activity_class"
	case HrZoneCalc:
		return "hr_zone_calc"
	case PwrZoneCalc:
		return "pwr_zone_calc"
	case WktStepDuration:
		return "wkt_step_duration"
	case WktStepTarget:
		return "wkt_step_target"
	case Goal:
		return "goal"
	case GoalRecurrence:
		return "goal_recurrence"
	case GoalSource:
		return "goal_source"
	case Schedule:
		return "schedule"
	case CoursePoint:
		return "course_point"
	case Manufacturer:
		return "manufacturer"
	case GarminProduct:
		return "garmin_product"
	case AntplusDeviceType:
		return "antplus_device_type"
	case AntNetwork:
		return "ant_network"
	case WorkoutCapabilities:
		return "workout_capabilities"
	case BatteryStatus:
		return "battery_status"
	case HrType:
		return "hr_type"
	case CourseCapabilities:
		return "course_capabilities"
	case Weight:
		return "weight"
	case WorkoutHr:
		return "workout_hr"
	case WorkoutPower:
		return "workout_power"
	case BpStatus:
		return "bp_status"
	case UserLocalId:
		return "user_local_id"
	case SwimStroke:
		return "swim_stroke"
	case ActivityType:
		return "activity_type"
	case ActivitySubtype:
		return "activity_subtype"
	case ActivityLevel:
		return "activity_level"
	case Side:
		return "side"
	case LeftRightBalance:
		return "left_right_balance"
	case LeftRightBalance100:
		return "left_right_balance_100"
	case LengthType:
		return "length_type"
	case DayOfWeek:
		return "day_of_week"
	case ConnectivityCapabilities:
		return "connectivity_capabilities"
	case WeatherReport:
		return "weather_report"
	case WeatherStatus:
		return "weather_status"
	case WeatherSeverity:
		return "weather_severity"
	case WeatherSevereType:
		return "weather_severe_type"
	case TimeIntoDay:
		return "time_into_day"
	case LocaltimeIntoDay:
		return "localtime_into_day"
	case StrokeType:
		return "stroke_type"
	case BodyLocation:
		return "body_location"
	case SegmentLapStatus:
		return "segment_lap_status"
	case SegmentLeaderboardType:
		return "segment_leaderboard_type"
	case SegmentDeleteStatus:
		return "segment_delete_status"
	case SegmentSelectionType:
		return "segment_selection_type"
	case SourceType:
		return "source_type"
	case LocalDeviceType:
		return "local_device_type"
	case BleDeviceType:
		return "ble_device_type"
	case AntChannelId:
		return "ant_channel_id"
	case DisplayOrientation:
		return "display_orientation"
	case WorkoutEquipment:
		return "workout_equipment"
	case WatchfaceMode:
		return "watchface_mode"
	case DigitalWatchfaceLayout:
		return "digital_watchface_layout"
	case AnalogWatchfaceLayout:
		return "analog_watchface_layout"
	case RiderPositionType:
		return "rider_position_type"
	case PowerPhaseType:
		return "power_phase_type"
	case CameraEventType:
		return "camera_event_type"
	case SensorType:
		return "sensor_type"
	case BikeLightNetworkConfigType:
		return "bike_light_network_config_type"
	case CommTimeoutType:
		return "comm_timeout_type"
	case CameraOrientationType:
		return "camera_orientation_type"
	case AttitudeStage:
		return "attitude_stage"
	case AttitudeValidity:
		return "attitude_validity"
	case AutoSyncFrequency:
		return "auto_sync_frequency"
	case ExdLayout:
		return "exd_layout"
	case ExdDisplayType:
		return "exd_display_type"
	case ExdDataUnits:
		return "exd_data_units"
	case ExdQualifiers:
		return "exd_qualifiers"
	case ExdDescriptors:
		return "exd_descriptors"
	case AutoActivityDetect:
		return "auto_activity_detect"
	case SupportedExdScreenLayouts:
		return "supported_exd_screen_layouts"
	case FitBaseType:
		return "fit_base_type"
	case TurnType:
		return "turn_type"
	case BikeLightBeamAngleMode:
		return "bike_light_beam_angle_mode"
	case FitBaseUnit:
		return "fit_base_unit"
	case SetType:
		return "set_type"
	case MaxMetCategory:
		return "max_met_category"
	case ExerciseCategory:
		return "exercise_category"
	case BenchPressExerciseName:
		return "bench_press_exercise_name"
	case CalfRaiseExerciseName:
		return "calf_raise_exercise_name"
	case CardioExerciseName:
		return "cardio_exercise_name"
	case CarryExerciseName:
		return "carry_exercise_name"
	case ChopExerciseName:
		return "chop_exercise_name"
	case CoreExerciseName:
		return "core_exercise_name"
	case CrunchExerciseName:
		return "crunch_exercise_name"
	case CurlExerciseName:
		return "curl_exercise_name"
	case DeadliftExerciseName:
		return "deadlift_exercise_name"
	case FlyeExerciseName:
		return "flye_exercise_name"
	case HipRaiseExerciseName:
		return "hip_raise_exercise_name"
	case HipStabilityExerciseName:
		return "hip_stability_exercise_name"
	case HipSwingExerciseName:
		return "hip_swing_exercise_name"
	case HyperextensionExerciseName:
		return "hyperextension_exercise_name"
	case LateralRaiseExerciseName:
		return "lateral_raise_exercise_name"
	case LegCurlExerciseName:
		return "leg_curl_exercise_name"
	case LegRaiseExerciseName:
		return "leg_raise_exercise_name"
	case LungeExerciseName:
		return "lunge_exercise_name"
	case OlympicLiftExerciseName:
		return "olympic_lift_exercise_name"
	case PlankExerciseName:
		return "plank_exercise_name"
	case PlyoExerciseName:
		return "plyo_exercise_name"
	case PullUpExerciseName:
		return "pull_up_exercise_name"
	case PushUpExerciseName:
		return "push_up_exercise_name"
	case RowExerciseName:
		return "row_exercise_name"
	case ShoulderPressExerciseName:
		return "shoulder_press_exercise_name"
	case ShoulderStabilityExerciseName:
		return "shoulder_stability_exercise_name"
	case ShrugExerciseName:
		return "shrug_exercise_name"
	case SitUpExerciseName:
		return "sit_up_exercise_name"
	case SquatExerciseName:
		return "squat_exercise_name"
	case TotalBodyExerciseName:
		return "total_body_exercise_name"
	case TricepsExtensionExerciseName:
		return "triceps_extension_exercise_name"
	case WarmUpExerciseName:
		return "warm_up_exercise_name"
	case RunExerciseName:
		return "run_exercise_name"
	case WaterType:
		return "water_type"
	case TissueModelType:
		return "tissue_model_type"
	case DiveGasStatus:
		return "dive_gas_status"
	case DiveAlert:
		return "dive_alert"
	case DiveAlarmType:
		return "dive_alarm_type"
	case DiveBacklightMode:
		return "dive_backlight_mode"
	case SleepLevel:
		return "sleep_level"
	case Spo2MeasurementType:
		return "spo2_measurement_type"
	case CcrSetpointSwitchMode:
		return "ccr_setpoint_switch_mode"
	case DiveGasMode:
		return "dive_gas_mode"
	case FaveroProduct:
		return "favero_product"
	case SplitType:
		return "split_type"
	case ClimbProEvent:
		return "climb_pro_event"
	case GasConsumptionRateType:
		return "gas_consumption_rate_type"
	case TapSensitivity:
		return "tap_sensitivity"
	case RadarThreatLevelType:
		return "radar_threat_level_type"
	case MaxMetSpeedSource:
		return "max_met_speed_source"
	case MaxMetHeartRateSource:
		return "max_met_heart_rate_source"
	case HrvStatus:
		return "hrv_status"
	case NoFlyTimeMode:
		return "no_fly_time_mode"
	default:
		return "ProfileTypeInvalid(" + strconv.FormatUint(uint64(p), 10) + ")"
	}
}

// FromString parse string into ProfileType constant it's represent, return ProfileTypeInvalid if not found.
func ProfileTypeFromString(s string) ProfileType {
	switch s {
	case "enum":
		return Enum
	case "sint8":
		return Sint8
	case "uint8":
		return Uint8
	case "sint16":
		return Sint16
	case "uint16":
		return Uint16
	case "sint32":
		return Sint32
	case "uint32":
		return Uint32
	case "string":
		return String
	case "float32":
		return Float32
	case "float64":
		return Float64
	case "uint8z":
		return Uint8z
	case "uint16z":
		return Uint16z
	case "uint32z":
		return Uint32z
	case "byte":
		return Byte
	case "sint64":
		return Sint64
	case "uint64":
		return Uint64
	case "uint64z":
		return Uint64z
	case "bool":
		return Bool
	case "file":
		return File
	case "mesg_num":
		return MesgNum
	case "checksum":
		return Checksum
	case "file_flags":
		return FileFlags
	case "mesg_count":
		return MesgCount
	case "date_time":
		return DateTime
	case "local_date_time":
		return LocalDateTime
	case "message_index":
		return MessageIndex
	case "device_index":
		return DeviceIndex
	case "gender":
		return Gender
	case "language":
		return Language
	case "language_bits_0":
		return LanguageBits0
	case "language_bits_1":
		return LanguageBits1
	case "language_bits_2":
		return LanguageBits2
	case "language_bits_3":
		return LanguageBits3
	case "language_bits_4":
		return LanguageBits4
	case "time_zone":
		return TimeZone
	case "display_measure":
		return DisplayMeasure
	case "display_heart":
		return DisplayHeart
	case "display_power":
		return DisplayPower
	case "display_position":
		return DisplayPosition
	case "switch":
		return Switch
	case "sport":
		return Sport
	case "sport_bits_0":
		return SportBits0
	case "sport_bits_1":
		return SportBits1
	case "sport_bits_2":
		return SportBits2
	case "sport_bits_3":
		return SportBits3
	case "sport_bits_4":
		return SportBits4
	case "sport_bits_5":
		return SportBits5
	case "sport_bits_6":
		return SportBits6
	case "sub_sport":
		return SubSport
	case "sport_event":
		return SportEvent
	case "activity":
		return Activity
	case "intensity":
		return Intensity
	case "session_trigger":
		return SessionTrigger
	case "autolap_trigger":
		return AutolapTrigger
	case "lap_trigger":
		return LapTrigger
	case "time_mode":
		return TimeMode
	case "backlight_mode":
		return BacklightMode
	case "date_mode":
		return DateMode
	case "backlight_timeout":
		return BacklightTimeout
	case "event":
		return Event
	case "event_type":
		return EventType
	case "timer_trigger":
		return TimerTrigger
	case "fitness_equipment_state":
		return FitnessEquipmentState
	case "tone":
		return Tone
	case "autoscroll":
		return Autoscroll
	case "activity_class":
		return ActivityClass
	case "hr_zone_calc":
		return HrZoneCalc
	case "pwr_zone_calc":
		return PwrZoneCalc
	case "wkt_step_duration":
		return WktStepDuration
	case "wkt_step_target":
		return WktStepTarget
	case "goal":
		return Goal
	case "goal_recurrence":
		return GoalRecurrence
	case "goal_source":
		return GoalSource
	case "schedule":
		return Schedule
	case "course_point":
		return CoursePoint
	case "manufacturer":
		return Manufacturer
	case "garmin_product":
		return GarminProduct
	case "antplus_device_type":
		return AntplusDeviceType
	case "ant_network":
		return AntNetwork
	case "workout_capabilities":
		return WorkoutCapabilities
	case "battery_status":
		return BatteryStatus
	case "hr_type":
		return HrType
	case "course_capabilities":
		return CourseCapabilities
	case "weight":
		return Weight
	case "workout_hr":
		return WorkoutHr
	case "workout_power":
		return WorkoutPower
	case "bp_status":
		return BpStatus
	case "user_local_id":
		return UserLocalId
	case "swim_stroke":
		return SwimStroke
	case "activity_type":
		return ActivityType
	case "activity_subtype":
		return ActivitySubtype
	case "activity_level":
		return ActivityLevel
	case "side":
		return Side
	case "left_right_balance":
		return LeftRightBalance
	case "left_right_balance_100":
		return LeftRightBalance100
	case "length_type":
		return LengthType
	case "day_of_week":
		return DayOfWeek
	case "connectivity_capabilities":
		return ConnectivityCapabilities
	case "weather_report":
		return WeatherReport
	case "weather_status":
		return WeatherStatus
	case "weather_severity":
		return WeatherSeverity
	case "weather_severe_type":
		return WeatherSevereType
	case "time_into_day":
		return TimeIntoDay
	case "localtime_into_day":
		return LocaltimeIntoDay
	case "stroke_type":
		return StrokeType
	case "body_location":
		return BodyLocation
	case "segment_lap_status":
		return SegmentLapStatus
	case "segment_leaderboard_type":
		return SegmentLeaderboardType
	case "segment_delete_status":
		return SegmentDeleteStatus
	case "segment_selection_type":
		return SegmentSelectionType
	case "source_type":
		return SourceType
	case "local_device_type":
		return LocalDeviceType
	case "ble_device_type":
		return BleDeviceType
	case "ant_channel_id":
		return AntChannelId
	case "display_orientation":
		return DisplayOrientation
	case "workout_equipment":
		return WorkoutEquipment
	case "watchface_mode":
		return WatchfaceMode
	case "digital_watchface_layout":
		return DigitalWatchfaceLayout
	case "analog_watchface_layout":
		return AnalogWatchfaceLayout
	case "rider_position_type":
		return RiderPositionType
	case "power_phase_type":
		return PowerPhaseType
	case "camera_event_type":
		return CameraEventType
	case "sensor_type":
		return SensorType
	case "bike_light_network_config_type":
		return BikeLightNetworkConfigType
	case "comm_timeout_type":
		return CommTimeoutType
	case "camera_orientation_type":
		return CameraOrientationType
	case "attitude_stage":
		return AttitudeStage
	case "attitude_validity":
		return AttitudeValidity
	case "auto_sync_frequency":
		return AutoSyncFrequency
	case "exd_layout":
		return ExdLayout
	case "exd_display_type":
		return ExdDisplayType
	case "exd_data_units":
		return ExdDataUnits
	case "exd_qualifiers":
		return ExdQualifiers
	case "exd_descriptors":
		return ExdDescriptors
	case "auto_activity_detect":
		return AutoActivityDetect
	case "supported_exd_screen_layouts":
		return SupportedExdScreenLayouts
	case "fit_base_type":
		return FitBaseType
	case "turn_type":
		return TurnType
	case "bike_light_beam_angle_mode":
		return BikeLightBeamAngleMode
	case "fit_base_unit":
		return FitBaseUnit
	case "set_type":
		return SetType
	case "max_met_category":
		return MaxMetCategory
	case "exercise_category":
		return ExerciseCategory
	case "bench_press_exercise_name":
		return BenchPressExerciseName
	case "calf_raise_exercise_name":
		return CalfRaiseExerciseName
	case "cardio_exercise_name":
		return CardioExerciseName
	case "carry_exercise_name":
		return CarryExerciseName
	case "chop_exercise_name":
		return ChopExerciseName
	case "core_exercise_name":
		return CoreExerciseName
	case "crunch_exercise_name":
		return CrunchExerciseName
	case "curl_exercise_name":
		return CurlExerciseName
	case "deadlift_exercise_name":
		return DeadliftExerciseName
	case "flye_exercise_name":
		return FlyeExerciseName
	case "hip_raise_exercise_name":
		return HipRaiseExerciseName
	case "hip_stability_exercise_name":
		return HipStabilityExerciseName
	case "hip_swing_exercise_name":
		return HipSwingExerciseName
	case "hyperextension_exercise_name":
		return HyperextensionExerciseName
	case "lateral_raise_exercise_name":
		return LateralRaiseExerciseName
	case "leg_curl_exercise_name":
		return LegCurlExerciseName
	case "leg_raise_exercise_name":
		return LegRaiseExerciseName
	case "lunge_exercise_name":
		return LungeExerciseName
	case "olympic_lift_exercise_name":
		return OlympicLiftExerciseName
	case "plank_exercise_name":
		return PlankExerciseName
	case "plyo_exercise_name":
		return PlyoExerciseName
	case "pull_up_exercise_name":
		return PullUpExerciseName
	case "push_up_exercise_name":
		return PushUpExerciseName
	case "row_exercise_name":
		return RowExerciseName
	case "shoulder_press_exercise_name":
		return ShoulderPressExerciseName
	case "shoulder_stability_exercise_name":
		return ShoulderStabilityExerciseName
	case "shrug_exercise_name":
		return ShrugExerciseName
	case "sit_up_exercise_name":
		return SitUpExerciseName
	case "squat_exercise_name":
		return SquatExerciseName
	case "total_body_exercise_name":
		return TotalBodyExerciseName
	case "triceps_extension_exercise_name":
		return TricepsExtensionExerciseName
	case "warm_up_exercise_name":
		return WarmUpExerciseName
	case "run_exercise_name":
		return RunExerciseName
	case "water_type":
		return WaterType
	case "tissue_model_type":
		return TissueModelType
	case "dive_gas_status":
		return DiveGasStatus
	case "dive_alert":
		return DiveAlert
	case "dive_alarm_type":
		return DiveAlarmType
	case "dive_backlight_mode":
		return DiveBacklightMode
	case "sleep_level":
		return SleepLevel
	case "spo2_measurement_type":
		return Spo2MeasurementType
	case "ccr_setpoint_switch_mode":
		return CcrSetpointSwitchMode
	case "dive_gas_mode":
		return DiveGasMode
	case "favero_product":
		return FaveroProduct
	case "split_type":
		return SplitType
	case "climb_pro_event":
		return ClimbProEvent
	case "gas_consumption_rate_type":
		return GasConsumptionRateType
	case "tap_sensitivity":
		return TapSensitivity
	case "radar_threat_level_type":
		return RadarThreatLevelType
	case "max_met_speed_source":
		return MaxMetSpeedSource
	case "max_met_heart_rate_source":
		return MaxMetHeartRateSource
	case "hrv_status":
		return HrvStatus
	case "no_fly_time_mode":
		return NoFlyTimeMode
	default:
		return Invalid
	}
}

// List returns all constants.
func ListProfileType() []ProfileType {
	return []ProfileType{
		Enum,
		Sint8,
		Uint8,
		Sint16,
		Uint16,
		Sint32,
		Uint32,
		String,
		Float32,
		Float64,
		Uint8z,
		Uint16z,
		Uint32z,
		Byte,
		Sint64,
		Uint64,
		Uint64z,
		Bool,
		File,
		MesgNum,
		Checksum,
		FileFlags,
		MesgCount,
		DateTime,
		LocalDateTime,
		MessageIndex,
		DeviceIndex,
		Gender,
		Language,
		LanguageBits0,
		LanguageBits1,
		LanguageBits2,
		LanguageBits3,
		LanguageBits4,
		TimeZone,
		DisplayMeasure,
		DisplayHeart,
		DisplayPower,
		DisplayPosition,
		Switch,
		Sport,
		SportBits0,
		SportBits1,
		SportBits2,
		SportBits3,
		SportBits4,
		SportBits5,
		SportBits6,
		SubSport,
		SportEvent,
		Activity,
		Intensity,
		SessionTrigger,
		AutolapTrigger,
		LapTrigger,
		TimeMode,
		BacklightMode,
		DateMode,
		BacklightTimeout,
		Event,
		EventType,
		TimerTrigger,
		FitnessEquipmentState,
		Tone,
		Autoscroll,
		ActivityClass,
		HrZoneCalc,
		PwrZoneCalc,
		WktStepDuration,
		WktStepTarget,
		Goal,
		GoalRecurrence,
		GoalSource,
		Schedule,
		CoursePoint,
		Manufacturer,
		GarminProduct,
		AntplusDeviceType,
		AntNetwork,
		WorkoutCapabilities,
		BatteryStatus,
		HrType,
		CourseCapabilities,
		Weight,
		WorkoutHr,
		WorkoutPower,
		BpStatus,
		UserLocalId,
		SwimStroke,
		ActivityType,
		ActivitySubtype,
		ActivityLevel,
		Side,
		LeftRightBalance,
		LeftRightBalance100,
		LengthType,
		DayOfWeek,
		ConnectivityCapabilities,
		WeatherReport,
		WeatherStatus,
		WeatherSeverity,
		WeatherSevereType,
		TimeIntoDay,
		LocaltimeIntoDay,
		StrokeType,
		BodyLocation,
		SegmentLapStatus,
		SegmentLeaderboardType,
		SegmentDeleteStatus,
		SegmentSelectionType,
		SourceType,
		LocalDeviceType,
		BleDeviceType,
		AntChannelId,
		DisplayOrientation,
		WorkoutEquipment,
		WatchfaceMode,
		DigitalWatchfaceLayout,
		AnalogWatchfaceLayout,
		RiderPositionType,
		PowerPhaseType,
		CameraEventType,
		SensorType,
		BikeLightNetworkConfigType,
		CommTimeoutType,
		CameraOrientationType,
		AttitudeStage,
		AttitudeValidity,
		AutoSyncFrequency,
		ExdLayout,
		ExdDisplayType,
		ExdDataUnits,
		ExdQualifiers,
		ExdDescriptors,
		AutoActivityDetect,
		SupportedExdScreenLayouts,
		FitBaseType,
		TurnType,
		BikeLightBeamAngleMode,
		FitBaseUnit,
		SetType,
		MaxMetCategory,
		ExerciseCategory,
		BenchPressExerciseName,
		CalfRaiseExerciseName,
		CardioExerciseName,
		CarryExerciseName,
		ChopExerciseName,
		CoreExerciseName,
		CrunchExerciseName,
		CurlExerciseName,
		DeadliftExerciseName,
		FlyeExerciseName,
		HipRaiseExerciseName,
		HipStabilityExerciseName,
		HipSwingExerciseName,
		HyperextensionExerciseName,
		LateralRaiseExerciseName,
		LegCurlExerciseName,
		LegRaiseExerciseName,
		LungeExerciseName,
		OlympicLiftExerciseName,
		PlankExerciseName,
		PlyoExerciseName,
		PullUpExerciseName,
		PushUpExerciseName,
		RowExerciseName,
		ShoulderPressExerciseName,
		ShoulderStabilityExerciseName,
		ShrugExerciseName,
		SitUpExerciseName,
		SquatExerciseName,
		TotalBodyExerciseName,
		TricepsExtensionExerciseName,
		WarmUpExerciseName,
		RunExerciseName,
		WaterType,
		TissueModelType,
		DiveGasStatus,
		DiveAlert,
		DiveAlarmType,
		DiveBacklightMode,
		SleepLevel,
		Spo2MeasurementType,
		CcrSetpointSwitchMode,
		DiveGasMode,
		FaveroProduct,
		SplitType,
		ClimbProEvent,
		GasConsumptionRateType,
		TapSensitivity,
		RadarThreatLevelType,
		MaxMetSpeedSource,
		MaxMetHeartRateSource,
		HrvStatus,
		NoFlyTimeMode,
	}
}

func (p ProfileType) BaseType() basetype.BaseType {
	switch p {
	case Enum:
		return basetype.Enum
	case Sint8:
		return basetype.Sint8
	case Uint8:
		return basetype.Uint8
	case Sint16:
		return basetype.Sint16
	case Uint16:
		return basetype.Uint16
	case Sint32:
		return basetype.Sint32
	case Uint32:
		return basetype.Uint32
	case String:
		return basetype.String
	case Float32:
		return basetype.Float32
	case Float64:
		return basetype.Float64
	case Uint8z:
		return basetype.Uint8z
	case Uint16z:
		return basetype.Uint16z
	case Uint32z:
		return basetype.Uint32z
	case Byte:
		return basetype.Byte
	case Sint64:
		return basetype.Sint64
	case Uint64:
		return basetype.Uint64
	case Uint64z:
		return basetype.Uint64z
	case Bool:
		return basetype.Enum
	case File:
		return basetype.Enum
	case MesgNum:
		return basetype.Uint16
	case Checksum:
		return basetype.Uint8
	case FileFlags:
		return basetype.Uint8z
	case MesgCount:
		return basetype.Enum
	case DateTime:
		return basetype.Uint32
	case LocalDateTime:
		return basetype.Uint32
	case MessageIndex:
		return basetype.Uint16
	case DeviceIndex:
		return basetype.Uint8
	case Gender:
		return basetype.Enum
	case Language:
		return basetype.Enum
	case LanguageBits0:
		return basetype.Uint8z
	case LanguageBits1:
		return basetype.Uint8z
	case LanguageBits2:
		return basetype.Uint8z
	case LanguageBits3:
		return basetype.Uint8z
	case LanguageBits4:
		return basetype.Uint8z
	case TimeZone:
		return basetype.Enum
	case DisplayMeasure:
		return basetype.Enum
	case DisplayHeart:
		return basetype.Enum
	case DisplayPower:
		return basetype.Enum
	case DisplayPosition:
		return basetype.Enum
	case Switch:
		return basetype.Enum
	case Sport:
		return basetype.Enum
	case SportBits0:
		return basetype.Uint8z
	case SportBits1:
		return basetype.Uint8z
	case SportBits2:
		return basetype.Uint8z
	case SportBits3:
		return basetype.Uint8z
	case SportBits4:
		return basetype.Uint8z
	case SportBits5:
		return basetype.Uint8z
	case SportBits6:
		return basetype.Uint8z
	case SubSport:
		return basetype.Enum
	case SportEvent:
		return basetype.Enum
	case Activity:
		return basetype.Enum
	case Intensity:
		return basetype.Enum
	case SessionTrigger:
		return basetype.Enum
	case AutolapTrigger:
		return basetype.Enum
	case LapTrigger:
		return basetype.Enum
	case TimeMode:
		return basetype.Enum
	case BacklightMode:
		return basetype.Enum
	case DateMode:
		return basetype.Enum
	case BacklightTimeout:
		return basetype.Uint8
	case Event:
		return basetype.Enum
	case EventType:
		return basetype.Enum
	case TimerTrigger:
		return basetype.Enum
	case FitnessEquipmentState:
		return basetype.Enum
	case Tone:
		return basetype.Enum
	case Autoscroll:
		return basetype.Enum
	case ActivityClass:
		return basetype.Enum
	case HrZoneCalc:
		return basetype.Enum
	case PwrZoneCalc:
		return basetype.Enum
	case WktStepDuration:
		return basetype.Enum
	case WktStepTarget:
		return basetype.Enum
	case Goal:
		return basetype.Enum
	case GoalRecurrence:
		return basetype.Enum
	case GoalSource:
		return basetype.Enum
	case Schedule:
		return basetype.Enum
	case CoursePoint:
		return basetype.Enum
	case Manufacturer:
		return basetype.Uint16
	case GarminProduct:
		return basetype.Uint16
	case AntplusDeviceType:
		return basetype.Uint8
	case AntNetwork:
		return basetype.Enum
	case WorkoutCapabilities:
		return basetype.Uint32z
	case BatteryStatus:
		return basetype.Uint8
	case HrType:
		return basetype.Enum
	case CourseCapabilities:
		return basetype.Uint32z
	case Weight:
		return basetype.Uint16
	case WorkoutHr:
		return basetype.Uint32
	case WorkoutPower:
		return basetype.Uint32
	case BpStatus:
		return basetype.Enum
	case UserLocalId:
		return basetype.Uint16
	case SwimStroke:
		return basetype.Enum
	case ActivityType:
		return basetype.Enum
	case ActivitySubtype:
		return basetype.Enum
	case ActivityLevel:
		return basetype.Enum
	case Side:
		return basetype.Enum
	case LeftRightBalance:
		return basetype.Uint8
	case LeftRightBalance100:
		return basetype.Uint16
	case LengthType:
		return basetype.Enum
	case DayOfWeek:
		return basetype.Enum
	case ConnectivityCapabilities:
		return basetype.Uint32z
	case WeatherReport:
		return basetype.Enum
	case WeatherStatus:
		return basetype.Enum
	case WeatherSeverity:
		return basetype.Enum
	case WeatherSevereType:
		return basetype.Enum
	case TimeIntoDay:
		return basetype.Uint32
	case LocaltimeIntoDay:
		return basetype.Uint32
	case StrokeType:
		return basetype.Enum
	case BodyLocation:
		return basetype.Enum
	case SegmentLapStatus:
		return basetype.Enum
	case SegmentLeaderboardType:
		return basetype.Enum
	case SegmentDeleteStatus:
		return basetype.Enum
	case SegmentSelectionType:
		return basetype.Enum
	case SourceType:
		return basetype.Enum
	case LocalDeviceType:
		return basetype.Uint8
	case BleDeviceType:
		return basetype.Uint8
	case AntChannelId:
		return basetype.Uint32z
	case DisplayOrientation:
		return basetype.Enum
	case WorkoutEquipment:
		return basetype.Enum
	case WatchfaceMode:
		return basetype.Enum
	case DigitalWatchfaceLayout:
		return basetype.Enum
	case AnalogWatchfaceLayout:
		return basetype.Enum
	case RiderPositionType:
		return basetype.Enum
	case PowerPhaseType:
		return basetype.Enum
	case CameraEventType:
		return basetype.Enum
	case SensorType:
		return basetype.Enum
	case BikeLightNetworkConfigType:
		return basetype.Enum
	case CommTimeoutType:
		return basetype.Uint16
	case CameraOrientationType:
		return basetype.Enum
	case AttitudeStage:
		return basetype.Enum
	case AttitudeValidity:
		return basetype.Uint16
	case AutoSyncFrequency:
		return basetype.Enum
	case ExdLayout:
		return basetype.Enum
	case ExdDisplayType:
		return basetype.Enum
	case ExdDataUnits:
		return basetype.Enum
	case ExdQualifiers:
		return basetype.Enum
	case ExdDescriptors:
		return basetype.Enum
	case AutoActivityDetect:
		return basetype.Uint32
	case SupportedExdScreenLayouts:
		return basetype.Uint32z
	case FitBaseType:
		return basetype.Uint8
	case TurnType:
		return basetype.Enum
	case BikeLightBeamAngleMode:
		return basetype.Uint8
	case FitBaseUnit:
		return basetype.Uint16
	case SetType:
		return basetype.Uint8
	case MaxMetCategory:
		return basetype.Enum
	case ExerciseCategory:
		return basetype.Uint16
	case BenchPressExerciseName:
		return basetype.Uint16
	case CalfRaiseExerciseName:
		return basetype.Uint16
	case CardioExerciseName:
		return basetype.Uint16
	case CarryExerciseName:
		return basetype.Uint16
	case ChopExerciseName:
		return basetype.Uint16
	case CoreExerciseName:
		return basetype.Uint16
	case CrunchExerciseName:
		return basetype.Uint16
	case CurlExerciseName:
		return basetype.Uint16
	case DeadliftExerciseName:
		return basetype.Uint16
	case FlyeExerciseName:
		return basetype.Uint16
	case HipRaiseExerciseName:
		return basetype.Uint16
	case HipStabilityExerciseName:
		return basetype.Uint16
	case HipSwingExerciseName:
		return basetype.Uint16
	case HyperextensionExerciseName:
		return basetype.Uint16
	case LateralRaiseExerciseName:
		return basetype.Uint16
	case LegCurlExerciseName:
		return basetype.Uint16
	case LegRaiseExerciseName:
		return basetype.Uint16
	case LungeExerciseName:
		return basetype.Uint16
	case OlympicLiftExerciseName:
		return basetype.Uint16
	case PlankExerciseName:
		return basetype.Uint16
	case PlyoExerciseName:
		return basetype.Uint16
	case PullUpExerciseName:
		return basetype.Uint16
	case PushUpExerciseName:
		return basetype.Uint16
	case RowExerciseName:
		return basetype.Uint16
	case ShoulderPressExerciseName:
		return basetype.Uint16
	case ShoulderStabilityExerciseName:
		return basetype.Uint16
	case ShrugExerciseName:
		return basetype.Uint16
	case SitUpExerciseName:
		return basetype.Uint16
	case SquatExerciseName:
		return basetype.Uint16
	case TotalBodyExerciseName:
		return basetype.Uint16
	case TricepsExtensionExerciseName:
		return basetype.Uint16
	case WarmUpExerciseName:
		return basetype.Uint16
	case RunExerciseName:
		return basetype.Uint16
	case WaterType:
		return basetype.Enum
	case TissueModelType:
		return basetype.Enum
	case DiveGasStatus:
		return basetype.Enum
	case DiveAlert:
		return basetype.Enum
	case DiveAlarmType:
		return basetype.Enum
	case DiveBacklightMode:
		return basetype.Enum
	case SleepLevel:
		return basetype.Enum
	case Spo2MeasurementType:
		return basetype.Enum
	case CcrSetpointSwitchMode:
		return basetype.Enum
	case DiveGasMode:
		return basetype.Enum
	case FaveroProduct:
		return basetype.Uint16
	case SplitType:
		return basetype.Enum
	case ClimbProEvent:
		return basetype.Enum
	case GasConsumptionRateType:
		return basetype.Enum
	case TapSensitivity:
		return basetype.Enum
	case RadarThreatLevelType:
		return basetype.Enum
	case MaxMetSpeedSource:
		return basetype.Enum
	case MaxMetHeartRateSource:
		return basetype.Enum
	case HrvStatus:
		return basetype.Enum
	case NoFlyTimeMode:
		return basetype.Enum
	}

	return 255
}
