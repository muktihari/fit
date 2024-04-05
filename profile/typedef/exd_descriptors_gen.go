// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ExdDescriptors byte

const (
	ExdDescriptorsBikeLightBatteryStatus           ExdDescriptors = 0
	ExdDescriptorsBeamAngleStatus                  ExdDescriptors = 1
	ExdDescriptorsBateryLevel                      ExdDescriptors = 2
	ExdDescriptorsLightNetworkMode                 ExdDescriptors = 3
	ExdDescriptorsNumberLightsConnected            ExdDescriptors = 4
	ExdDescriptorsCadence                          ExdDescriptors = 5
	ExdDescriptorsDistance                         ExdDescriptors = 6
	ExdDescriptorsEstimatedTimeOfArrival           ExdDescriptors = 7
	ExdDescriptorsHeading                          ExdDescriptors = 8
	ExdDescriptorsTime                             ExdDescriptors = 9
	ExdDescriptorsBatteryLevel                     ExdDescriptors = 10
	ExdDescriptorsTrainerResistance                ExdDescriptors = 11
	ExdDescriptorsTrainerTargetPower               ExdDescriptors = 12
	ExdDescriptorsTimeSeated                       ExdDescriptors = 13
	ExdDescriptorsTimeStanding                     ExdDescriptors = 14
	ExdDescriptorsElevation                        ExdDescriptors = 15
	ExdDescriptorsGrade                            ExdDescriptors = 16
	ExdDescriptorsAscent                           ExdDescriptors = 17
	ExdDescriptorsDescent                          ExdDescriptors = 18
	ExdDescriptorsVerticalSpeed                    ExdDescriptors = 19
	ExdDescriptorsDi2BatteryLevel                  ExdDescriptors = 20
	ExdDescriptorsFrontGear                        ExdDescriptors = 21
	ExdDescriptorsRearGear                         ExdDescriptors = 22
	ExdDescriptorsGearRatio                        ExdDescriptors = 23
	ExdDescriptorsHeartRate                        ExdDescriptors = 24
	ExdDescriptorsHeartRateZone                    ExdDescriptors = 25
	ExdDescriptorsTimeInHeartRateZone              ExdDescriptors = 26
	ExdDescriptorsHeartRateReserve                 ExdDescriptors = 27
	ExdDescriptorsCalories                         ExdDescriptors = 28
	ExdDescriptorsGpsAccuracy                      ExdDescriptors = 29
	ExdDescriptorsGpsSignalStrength                ExdDescriptors = 30
	ExdDescriptorsTemperature                      ExdDescriptors = 31
	ExdDescriptorsTimeOfDay                        ExdDescriptors = 32
	ExdDescriptorsBalance                          ExdDescriptors = 33
	ExdDescriptorsPedalSmoothness                  ExdDescriptors = 34
	ExdDescriptorsPower                            ExdDescriptors = 35
	ExdDescriptorsFunctionalThresholdPower         ExdDescriptors = 36
	ExdDescriptorsIntensityFactor                  ExdDescriptors = 37
	ExdDescriptorsWork                             ExdDescriptors = 38
	ExdDescriptorsPowerRatio                       ExdDescriptors = 39
	ExdDescriptorsNormalizedPower                  ExdDescriptors = 40
	ExdDescriptorsTrainingStressScore              ExdDescriptors = 41
	ExdDescriptorsTimeOnZone                       ExdDescriptors = 42
	ExdDescriptorsSpeed                            ExdDescriptors = 43
	ExdDescriptorsLaps                             ExdDescriptors = 44
	ExdDescriptorsReps                             ExdDescriptors = 45
	ExdDescriptorsWorkoutStep                      ExdDescriptors = 46
	ExdDescriptorsCourseDistance                   ExdDescriptors = 47
	ExdDescriptorsNavigationDistance               ExdDescriptors = 48
	ExdDescriptorsCourseEstimatedTimeOfArrival     ExdDescriptors = 49
	ExdDescriptorsNavigationEstimatedTimeOfArrival ExdDescriptors = 50
	ExdDescriptorsCourseTime                       ExdDescriptors = 51
	ExdDescriptorsNavigationTime                   ExdDescriptors = 52
	ExdDescriptorsCourseHeading                    ExdDescriptors = 53
	ExdDescriptorsNavigationHeading                ExdDescriptors = 54
	ExdDescriptorsPowerZone                        ExdDescriptors = 55
	ExdDescriptorsTorqueEffectiveness              ExdDescriptors = 56
	ExdDescriptorsTimerTime                        ExdDescriptors = 57
	ExdDescriptorsPowerWeightRatio                 ExdDescriptors = 58
	ExdDescriptorsLeftPlatformCenterOffset         ExdDescriptors = 59
	ExdDescriptorsRightPlatformCenterOffset        ExdDescriptors = 60
	ExdDescriptorsLeftPowerPhaseStartAngle         ExdDescriptors = 61
	ExdDescriptorsRightPowerPhaseStartAngle        ExdDescriptors = 62
	ExdDescriptorsLeftPowerPhaseFinishAngle        ExdDescriptors = 63
	ExdDescriptorsRightPowerPhaseFinishAngle       ExdDescriptors = 64
	ExdDescriptorsGears                            ExdDescriptors = 65 // Combined gear information
	ExdDescriptorsPace                             ExdDescriptors = 66
	ExdDescriptorsTrainingEffect                   ExdDescriptors = 67
	ExdDescriptorsVerticalOscillation              ExdDescriptors = 68
	ExdDescriptorsVerticalRatio                    ExdDescriptors = 69
	ExdDescriptorsGroundContactTime                ExdDescriptors = 70
	ExdDescriptorsLeftGroundContactTimeBalance     ExdDescriptors = 71
	ExdDescriptorsRightGroundContactTimeBalance    ExdDescriptors = 72
	ExdDescriptorsStrideLength                     ExdDescriptors = 73
	ExdDescriptorsRunningCadence                   ExdDescriptors = 74
	ExdDescriptorsPerformanceCondition             ExdDescriptors = 75
	ExdDescriptorsCourseType                       ExdDescriptors = 76
	ExdDescriptorsTimeInPowerZone                  ExdDescriptors = 77
	ExdDescriptorsNavigationTurn                   ExdDescriptors = 78
	ExdDescriptorsCourseLocation                   ExdDescriptors = 79
	ExdDescriptorsNavigationLocation               ExdDescriptors = 80
	ExdDescriptorsCompass                          ExdDescriptors = 81
	ExdDescriptorsGearCombo                        ExdDescriptors = 82
	ExdDescriptorsMuscleOxygen                     ExdDescriptors = 83
	ExdDescriptorsIcon                             ExdDescriptors = 84
	ExdDescriptorsCompassHeading                   ExdDescriptors = 85
	ExdDescriptorsGpsHeading                       ExdDescriptors = 86
	ExdDescriptorsGpsElevation                     ExdDescriptors = 87
	ExdDescriptorsAnaerobicTrainingEffect          ExdDescriptors = 88
	ExdDescriptorsCourse                           ExdDescriptors = 89
	ExdDescriptorsOffCourse                        ExdDescriptors = 90
	ExdDescriptorsGlideRatio                       ExdDescriptors = 91
	ExdDescriptorsVerticalDistance                 ExdDescriptors = 92
	ExdDescriptorsVmg                              ExdDescriptors = 93
	ExdDescriptorsAmbientPressure                  ExdDescriptors = 94
	ExdDescriptorsPressure                         ExdDescriptors = 95
	ExdDescriptorsVam                              ExdDescriptors = 96
	ExdDescriptorsInvalid                          ExdDescriptors = 0xFF
)

func (e ExdDescriptors) Byte() byte { return byte(e) }

func (e ExdDescriptors) String() string {
	switch e {
	case ExdDescriptorsBikeLightBatteryStatus:
		return "bike_light_battery_status"
	case ExdDescriptorsBeamAngleStatus:
		return "beam_angle_status"
	case ExdDescriptorsBateryLevel:
		return "batery_level"
	case ExdDescriptorsLightNetworkMode:
		return "light_network_mode"
	case ExdDescriptorsNumberLightsConnected:
		return "number_lights_connected"
	case ExdDescriptorsCadence:
		return "cadence"
	case ExdDescriptorsDistance:
		return "distance"
	case ExdDescriptorsEstimatedTimeOfArrival:
		return "estimated_time_of_arrival"
	case ExdDescriptorsHeading:
		return "heading"
	case ExdDescriptorsTime:
		return "time"
	case ExdDescriptorsBatteryLevel:
		return "battery_level"
	case ExdDescriptorsTrainerResistance:
		return "trainer_resistance"
	case ExdDescriptorsTrainerTargetPower:
		return "trainer_target_power"
	case ExdDescriptorsTimeSeated:
		return "time_seated"
	case ExdDescriptorsTimeStanding:
		return "time_standing"
	case ExdDescriptorsElevation:
		return "elevation"
	case ExdDescriptorsGrade:
		return "grade"
	case ExdDescriptorsAscent:
		return "ascent"
	case ExdDescriptorsDescent:
		return "descent"
	case ExdDescriptorsVerticalSpeed:
		return "vertical_speed"
	case ExdDescriptorsDi2BatteryLevel:
		return "di2_battery_level"
	case ExdDescriptorsFrontGear:
		return "front_gear"
	case ExdDescriptorsRearGear:
		return "rear_gear"
	case ExdDescriptorsGearRatio:
		return "gear_ratio"
	case ExdDescriptorsHeartRate:
		return "heart_rate"
	case ExdDescriptorsHeartRateZone:
		return "heart_rate_zone"
	case ExdDescriptorsTimeInHeartRateZone:
		return "time_in_heart_rate_zone"
	case ExdDescriptorsHeartRateReserve:
		return "heart_rate_reserve"
	case ExdDescriptorsCalories:
		return "calories"
	case ExdDescriptorsGpsAccuracy:
		return "gps_accuracy"
	case ExdDescriptorsGpsSignalStrength:
		return "gps_signal_strength"
	case ExdDescriptorsTemperature:
		return "temperature"
	case ExdDescriptorsTimeOfDay:
		return "time_of_day"
	case ExdDescriptorsBalance:
		return "balance"
	case ExdDescriptorsPedalSmoothness:
		return "pedal_smoothness"
	case ExdDescriptorsPower:
		return "power"
	case ExdDescriptorsFunctionalThresholdPower:
		return "functional_threshold_power"
	case ExdDescriptorsIntensityFactor:
		return "intensity_factor"
	case ExdDescriptorsWork:
		return "work"
	case ExdDescriptorsPowerRatio:
		return "power_ratio"
	case ExdDescriptorsNormalizedPower:
		return "normalized_power"
	case ExdDescriptorsTrainingStressScore:
		return "training_stress_Score"
	case ExdDescriptorsTimeOnZone:
		return "time_on_zone"
	case ExdDescriptorsSpeed:
		return "speed"
	case ExdDescriptorsLaps:
		return "laps"
	case ExdDescriptorsReps:
		return "reps"
	case ExdDescriptorsWorkoutStep:
		return "workout_step"
	case ExdDescriptorsCourseDistance:
		return "course_distance"
	case ExdDescriptorsNavigationDistance:
		return "navigation_distance"
	case ExdDescriptorsCourseEstimatedTimeOfArrival:
		return "course_estimated_time_of_arrival"
	case ExdDescriptorsNavigationEstimatedTimeOfArrival:
		return "navigation_estimated_time_of_arrival"
	case ExdDescriptorsCourseTime:
		return "course_time"
	case ExdDescriptorsNavigationTime:
		return "navigation_time"
	case ExdDescriptorsCourseHeading:
		return "course_heading"
	case ExdDescriptorsNavigationHeading:
		return "navigation_heading"
	case ExdDescriptorsPowerZone:
		return "power_zone"
	case ExdDescriptorsTorqueEffectiveness:
		return "torque_effectiveness"
	case ExdDescriptorsTimerTime:
		return "timer_time"
	case ExdDescriptorsPowerWeightRatio:
		return "power_weight_ratio"
	case ExdDescriptorsLeftPlatformCenterOffset:
		return "left_platform_center_offset"
	case ExdDescriptorsRightPlatformCenterOffset:
		return "right_platform_center_offset"
	case ExdDescriptorsLeftPowerPhaseStartAngle:
		return "left_power_phase_start_angle"
	case ExdDescriptorsRightPowerPhaseStartAngle:
		return "right_power_phase_start_angle"
	case ExdDescriptorsLeftPowerPhaseFinishAngle:
		return "left_power_phase_finish_angle"
	case ExdDescriptorsRightPowerPhaseFinishAngle:
		return "right_power_phase_finish_angle"
	case ExdDescriptorsGears:
		return "gears"
	case ExdDescriptorsPace:
		return "pace"
	case ExdDescriptorsTrainingEffect:
		return "training_effect"
	case ExdDescriptorsVerticalOscillation:
		return "vertical_oscillation"
	case ExdDescriptorsVerticalRatio:
		return "vertical_ratio"
	case ExdDescriptorsGroundContactTime:
		return "ground_contact_time"
	case ExdDescriptorsLeftGroundContactTimeBalance:
		return "left_ground_contact_time_balance"
	case ExdDescriptorsRightGroundContactTimeBalance:
		return "right_ground_contact_time_balance"
	case ExdDescriptorsStrideLength:
		return "stride_length"
	case ExdDescriptorsRunningCadence:
		return "running_cadence"
	case ExdDescriptorsPerformanceCondition:
		return "performance_condition"
	case ExdDescriptorsCourseType:
		return "course_type"
	case ExdDescriptorsTimeInPowerZone:
		return "time_in_power_zone"
	case ExdDescriptorsNavigationTurn:
		return "navigation_turn"
	case ExdDescriptorsCourseLocation:
		return "course_location"
	case ExdDescriptorsNavigationLocation:
		return "navigation_location"
	case ExdDescriptorsCompass:
		return "compass"
	case ExdDescriptorsGearCombo:
		return "gear_combo"
	case ExdDescriptorsMuscleOxygen:
		return "muscle_oxygen"
	case ExdDescriptorsIcon:
		return "icon"
	case ExdDescriptorsCompassHeading:
		return "compass_heading"
	case ExdDescriptorsGpsHeading:
		return "gps_heading"
	case ExdDescriptorsGpsElevation:
		return "gps_elevation"
	case ExdDescriptorsAnaerobicTrainingEffect:
		return "anaerobic_training_effect"
	case ExdDescriptorsCourse:
		return "course"
	case ExdDescriptorsOffCourse:
		return "off_course"
	case ExdDescriptorsGlideRatio:
		return "glide_ratio"
	case ExdDescriptorsVerticalDistance:
		return "vertical_distance"
	case ExdDescriptorsVmg:
		return "vmg"
	case ExdDescriptorsAmbientPressure:
		return "ambient_pressure"
	case ExdDescriptorsPressure:
		return "pressure"
	case ExdDescriptorsVam:
		return "vam"
	default:
		return "ExdDescriptorsInvalid(" + strconv.Itoa(int(e)) + ")"
	}
}

// FromString parse string into ExdDescriptors constant it's represent, return ExdDescriptorsInvalid if not found.
func ExdDescriptorsFromString(s string) ExdDescriptors {
	switch s {
	case "bike_light_battery_status":
		return ExdDescriptorsBikeLightBatteryStatus
	case "beam_angle_status":
		return ExdDescriptorsBeamAngleStatus
	case "batery_level":
		return ExdDescriptorsBateryLevel
	case "light_network_mode":
		return ExdDescriptorsLightNetworkMode
	case "number_lights_connected":
		return ExdDescriptorsNumberLightsConnected
	case "cadence":
		return ExdDescriptorsCadence
	case "distance":
		return ExdDescriptorsDistance
	case "estimated_time_of_arrival":
		return ExdDescriptorsEstimatedTimeOfArrival
	case "heading":
		return ExdDescriptorsHeading
	case "time":
		return ExdDescriptorsTime
	case "battery_level":
		return ExdDescriptorsBatteryLevel
	case "trainer_resistance":
		return ExdDescriptorsTrainerResistance
	case "trainer_target_power":
		return ExdDescriptorsTrainerTargetPower
	case "time_seated":
		return ExdDescriptorsTimeSeated
	case "time_standing":
		return ExdDescriptorsTimeStanding
	case "elevation":
		return ExdDescriptorsElevation
	case "grade":
		return ExdDescriptorsGrade
	case "ascent":
		return ExdDescriptorsAscent
	case "descent":
		return ExdDescriptorsDescent
	case "vertical_speed":
		return ExdDescriptorsVerticalSpeed
	case "di2_battery_level":
		return ExdDescriptorsDi2BatteryLevel
	case "front_gear":
		return ExdDescriptorsFrontGear
	case "rear_gear":
		return ExdDescriptorsRearGear
	case "gear_ratio":
		return ExdDescriptorsGearRatio
	case "heart_rate":
		return ExdDescriptorsHeartRate
	case "heart_rate_zone":
		return ExdDescriptorsHeartRateZone
	case "time_in_heart_rate_zone":
		return ExdDescriptorsTimeInHeartRateZone
	case "heart_rate_reserve":
		return ExdDescriptorsHeartRateReserve
	case "calories":
		return ExdDescriptorsCalories
	case "gps_accuracy":
		return ExdDescriptorsGpsAccuracy
	case "gps_signal_strength":
		return ExdDescriptorsGpsSignalStrength
	case "temperature":
		return ExdDescriptorsTemperature
	case "time_of_day":
		return ExdDescriptorsTimeOfDay
	case "balance":
		return ExdDescriptorsBalance
	case "pedal_smoothness":
		return ExdDescriptorsPedalSmoothness
	case "power":
		return ExdDescriptorsPower
	case "functional_threshold_power":
		return ExdDescriptorsFunctionalThresholdPower
	case "intensity_factor":
		return ExdDescriptorsIntensityFactor
	case "work":
		return ExdDescriptorsWork
	case "power_ratio":
		return ExdDescriptorsPowerRatio
	case "normalized_power":
		return ExdDescriptorsNormalizedPower
	case "training_stress_Score":
		return ExdDescriptorsTrainingStressScore
	case "time_on_zone":
		return ExdDescriptorsTimeOnZone
	case "speed":
		return ExdDescriptorsSpeed
	case "laps":
		return ExdDescriptorsLaps
	case "reps":
		return ExdDescriptorsReps
	case "workout_step":
		return ExdDescriptorsWorkoutStep
	case "course_distance":
		return ExdDescriptorsCourseDistance
	case "navigation_distance":
		return ExdDescriptorsNavigationDistance
	case "course_estimated_time_of_arrival":
		return ExdDescriptorsCourseEstimatedTimeOfArrival
	case "navigation_estimated_time_of_arrival":
		return ExdDescriptorsNavigationEstimatedTimeOfArrival
	case "course_time":
		return ExdDescriptorsCourseTime
	case "navigation_time":
		return ExdDescriptorsNavigationTime
	case "course_heading":
		return ExdDescriptorsCourseHeading
	case "navigation_heading":
		return ExdDescriptorsNavigationHeading
	case "power_zone":
		return ExdDescriptorsPowerZone
	case "torque_effectiveness":
		return ExdDescriptorsTorqueEffectiveness
	case "timer_time":
		return ExdDescriptorsTimerTime
	case "power_weight_ratio":
		return ExdDescriptorsPowerWeightRatio
	case "left_platform_center_offset":
		return ExdDescriptorsLeftPlatformCenterOffset
	case "right_platform_center_offset":
		return ExdDescriptorsRightPlatformCenterOffset
	case "left_power_phase_start_angle":
		return ExdDescriptorsLeftPowerPhaseStartAngle
	case "right_power_phase_start_angle":
		return ExdDescriptorsRightPowerPhaseStartAngle
	case "left_power_phase_finish_angle":
		return ExdDescriptorsLeftPowerPhaseFinishAngle
	case "right_power_phase_finish_angle":
		return ExdDescriptorsRightPowerPhaseFinishAngle
	case "gears":
		return ExdDescriptorsGears
	case "pace":
		return ExdDescriptorsPace
	case "training_effect":
		return ExdDescriptorsTrainingEffect
	case "vertical_oscillation":
		return ExdDescriptorsVerticalOscillation
	case "vertical_ratio":
		return ExdDescriptorsVerticalRatio
	case "ground_contact_time":
		return ExdDescriptorsGroundContactTime
	case "left_ground_contact_time_balance":
		return ExdDescriptorsLeftGroundContactTimeBalance
	case "right_ground_contact_time_balance":
		return ExdDescriptorsRightGroundContactTimeBalance
	case "stride_length":
		return ExdDescriptorsStrideLength
	case "running_cadence":
		return ExdDescriptorsRunningCadence
	case "performance_condition":
		return ExdDescriptorsPerformanceCondition
	case "course_type":
		return ExdDescriptorsCourseType
	case "time_in_power_zone":
		return ExdDescriptorsTimeInPowerZone
	case "navigation_turn":
		return ExdDescriptorsNavigationTurn
	case "course_location":
		return ExdDescriptorsCourseLocation
	case "navigation_location":
		return ExdDescriptorsNavigationLocation
	case "compass":
		return ExdDescriptorsCompass
	case "gear_combo":
		return ExdDescriptorsGearCombo
	case "muscle_oxygen":
		return ExdDescriptorsMuscleOxygen
	case "icon":
		return ExdDescriptorsIcon
	case "compass_heading":
		return ExdDescriptorsCompassHeading
	case "gps_heading":
		return ExdDescriptorsGpsHeading
	case "gps_elevation":
		return ExdDescriptorsGpsElevation
	case "anaerobic_training_effect":
		return ExdDescriptorsAnaerobicTrainingEffect
	case "course":
		return ExdDescriptorsCourse
	case "off_course":
		return ExdDescriptorsOffCourse
	case "glide_ratio":
		return ExdDescriptorsGlideRatio
	case "vertical_distance":
		return ExdDescriptorsVerticalDistance
	case "vmg":
		return ExdDescriptorsVmg
	case "ambient_pressure":
		return ExdDescriptorsAmbientPressure
	case "pressure":
		return ExdDescriptorsPressure
	case "vam":
		return ExdDescriptorsVam
	default:
		return ExdDescriptorsInvalid
	}
}

// List returns all constants.
func ListExdDescriptors() []ExdDescriptors {
	return []ExdDescriptors{
		ExdDescriptorsBikeLightBatteryStatus,
		ExdDescriptorsBeamAngleStatus,
		ExdDescriptorsBateryLevel,
		ExdDescriptorsLightNetworkMode,
		ExdDescriptorsNumberLightsConnected,
		ExdDescriptorsCadence,
		ExdDescriptorsDistance,
		ExdDescriptorsEstimatedTimeOfArrival,
		ExdDescriptorsHeading,
		ExdDescriptorsTime,
		ExdDescriptorsBatteryLevel,
		ExdDescriptorsTrainerResistance,
		ExdDescriptorsTrainerTargetPower,
		ExdDescriptorsTimeSeated,
		ExdDescriptorsTimeStanding,
		ExdDescriptorsElevation,
		ExdDescriptorsGrade,
		ExdDescriptorsAscent,
		ExdDescriptorsDescent,
		ExdDescriptorsVerticalSpeed,
		ExdDescriptorsDi2BatteryLevel,
		ExdDescriptorsFrontGear,
		ExdDescriptorsRearGear,
		ExdDescriptorsGearRatio,
		ExdDescriptorsHeartRate,
		ExdDescriptorsHeartRateZone,
		ExdDescriptorsTimeInHeartRateZone,
		ExdDescriptorsHeartRateReserve,
		ExdDescriptorsCalories,
		ExdDescriptorsGpsAccuracy,
		ExdDescriptorsGpsSignalStrength,
		ExdDescriptorsTemperature,
		ExdDescriptorsTimeOfDay,
		ExdDescriptorsBalance,
		ExdDescriptorsPedalSmoothness,
		ExdDescriptorsPower,
		ExdDescriptorsFunctionalThresholdPower,
		ExdDescriptorsIntensityFactor,
		ExdDescriptorsWork,
		ExdDescriptorsPowerRatio,
		ExdDescriptorsNormalizedPower,
		ExdDescriptorsTrainingStressScore,
		ExdDescriptorsTimeOnZone,
		ExdDescriptorsSpeed,
		ExdDescriptorsLaps,
		ExdDescriptorsReps,
		ExdDescriptorsWorkoutStep,
		ExdDescriptorsCourseDistance,
		ExdDescriptorsNavigationDistance,
		ExdDescriptorsCourseEstimatedTimeOfArrival,
		ExdDescriptorsNavigationEstimatedTimeOfArrival,
		ExdDescriptorsCourseTime,
		ExdDescriptorsNavigationTime,
		ExdDescriptorsCourseHeading,
		ExdDescriptorsNavigationHeading,
		ExdDescriptorsPowerZone,
		ExdDescriptorsTorqueEffectiveness,
		ExdDescriptorsTimerTime,
		ExdDescriptorsPowerWeightRatio,
		ExdDescriptorsLeftPlatformCenterOffset,
		ExdDescriptorsRightPlatformCenterOffset,
		ExdDescriptorsLeftPowerPhaseStartAngle,
		ExdDescriptorsRightPowerPhaseStartAngle,
		ExdDescriptorsLeftPowerPhaseFinishAngle,
		ExdDescriptorsRightPowerPhaseFinishAngle,
		ExdDescriptorsGears,
		ExdDescriptorsPace,
		ExdDescriptorsTrainingEffect,
		ExdDescriptorsVerticalOscillation,
		ExdDescriptorsVerticalRatio,
		ExdDescriptorsGroundContactTime,
		ExdDescriptorsLeftGroundContactTimeBalance,
		ExdDescriptorsRightGroundContactTimeBalance,
		ExdDescriptorsStrideLength,
		ExdDescriptorsRunningCadence,
		ExdDescriptorsPerformanceCondition,
		ExdDescriptorsCourseType,
		ExdDescriptorsTimeInPowerZone,
		ExdDescriptorsNavigationTurn,
		ExdDescriptorsCourseLocation,
		ExdDescriptorsNavigationLocation,
		ExdDescriptorsCompass,
		ExdDescriptorsGearCombo,
		ExdDescriptorsMuscleOxygen,
		ExdDescriptorsIcon,
		ExdDescriptorsCompassHeading,
		ExdDescriptorsGpsHeading,
		ExdDescriptorsGpsElevation,
		ExdDescriptorsAnaerobicTrainingEffect,
		ExdDescriptorsCourse,
		ExdDescriptorsOffCourse,
		ExdDescriptorsGlideRatio,
		ExdDescriptorsVerticalDistance,
		ExdDescriptorsVmg,
		ExdDescriptorsAmbientPressure,
		ExdDescriptorsPressure,
		ExdDescriptorsVam,
	}
}
