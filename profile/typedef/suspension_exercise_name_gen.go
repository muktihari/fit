// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SuspensionExerciseName uint16

const (
	SuspensionExerciseNameChestFly               SuspensionExerciseName = 0
	SuspensionExerciseNameChestPress             SuspensionExerciseName = 1
	SuspensionExerciseNameCrunch                 SuspensionExerciseName = 2
	SuspensionExerciseNameCurl                   SuspensionExerciseName = 3
	SuspensionExerciseNameDip                    SuspensionExerciseName = 4
	SuspensionExerciseNameFacePull               SuspensionExerciseName = 5
	SuspensionExerciseNameGluteBridge            SuspensionExerciseName = 6
	SuspensionExerciseNameHamstringCurl          SuspensionExerciseName = 7
	SuspensionExerciseNameHipDrop                SuspensionExerciseName = 8
	SuspensionExerciseNameInvertedRow            SuspensionExerciseName = 9
	SuspensionExerciseNameKneeDriveJump          SuspensionExerciseName = 10
	SuspensionExerciseNameKneeToChest            SuspensionExerciseName = 11
	SuspensionExerciseNameLatPullover            SuspensionExerciseName = 12
	SuspensionExerciseNameLunge                  SuspensionExerciseName = 13
	SuspensionExerciseNameMountainClimber        SuspensionExerciseName = 14
	SuspensionExerciseNamePendulum               SuspensionExerciseName = 15
	SuspensionExerciseNamePike                   SuspensionExerciseName = 16
	SuspensionExerciseNamePlank                  SuspensionExerciseName = 17
	SuspensionExerciseNamePowerPull              SuspensionExerciseName = 18
	SuspensionExerciseNamePullUp                 SuspensionExerciseName = 19
	SuspensionExerciseNamePushUp                 SuspensionExerciseName = 20
	SuspensionExerciseNameReverseMountainClimber SuspensionExerciseName = 21
	SuspensionExerciseNameReversePlank           SuspensionExerciseName = 22
	SuspensionExerciseNameRollout                SuspensionExerciseName = 23
	SuspensionExerciseNameRow                    SuspensionExerciseName = 24
	SuspensionExerciseNameSideLunge              SuspensionExerciseName = 25
	SuspensionExerciseNameSidePlank              SuspensionExerciseName = 26
	SuspensionExerciseNameSingleLegDeadlift      SuspensionExerciseName = 27
	SuspensionExerciseNameSingleLegSquat         SuspensionExerciseName = 28
	SuspensionExerciseNameSitUp                  SuspensionExerciseName = 29
	SuspensionExerciseNameSplit                  SuspensionExerciseName = 30
	SuspensionExerciseNameSquat                  SuspensionExerciseName = 31
	SuspensionExerciseNameSquatJump              SuspensionExerciseName = 32
	SuspensionExerciseNameTricepPress            SuspensionExerciseName = 33
	SuspensionExerciseNameYFly                   SuspensionExerciseName = 34
	SuspensionExerciseNameInvalid                SuspensionExerciseName = 0xFFFF
)

func (s SuspensionExerciseName) Uint16() uint16 { return uint16(s) }

func (s SuspensionExerciseName) String() string {
	switch s {
	case SuspensionExerciseNameChestFly:
		return "chest_fly"
	case SuspensionExerciseNameChestPress:
		return "chest_press"
	case SuspensionExerciseNameCrunch:
		return "crunch"
	case SuspensionExerciseNameCurl:
		return "curl"
	case SuspensionExerciseNameDip:
		return "dip"
	case SuspensionExerciseNameFacePull:
		return "face_pull"
	case SuspensionExerciseNameGluteBridge:
		return "glute_bridge"
	case SuspensionExerciseNameHamstringCurl:
		return "hamstring_curl"
	case SuspensionExerciseNameHipDrop:
		return "hip_drop"
	case SuspensionExerciseNameInvertedRow:
		return "inverted_row"
	case SuspensionExerciseNameKneeDriveJump:
		return "knee_drive_jump"
	case SuspensionExerciseNameKneeToChest:
		return "knee_to_chest"
	case SuspensionExerciseNameLatPullover:
		return "lat_pullover"
	case SuspensionExerciseNameLunge:
		return "lunge"
	case SuspensionExerciseNameMountainClimber:
		return "mountain_climber"
	case SuspensionExerciseNamePendulum:
		return "pendulum"
	case SuspensionExerciseNamePike:
		return "pike"
	case SuspensionExerciseNamePlank:
		return "plank"
	case SuspensionExerciseNamePowerPull:
		return "power_pull"
	case SuspensionExerciseNamePullUp:
		return "pull_up"
	case SuspensionExerciseNamePushUp:
		return "push_up"
	case SuspensionExerciseNameReverseMountainClimber:
		return "reverse_mountain_climber"
	case SuspensionExerciseNameReversePlank:
		return "reverse_plank"
	case SuspensionExerciseNameRollout:
		return "rollout"
	case SuspensionExerciseNameRow:
		return "row"
	case SuspensionExerciseNameSideLunge:
		return "side_lunge"
	case SuspensionExerciseNameSidePlank:
		return "side_plank"
	case SuspensionExerciseNameSingleLegDeadlift:
		return "single_leg_deadlift"
	case SuspensionExerciseNameSingleLegSquat:
		return "single_leg_squat"
	case SuspensionExerciseNameSitUp:
		return "sit_up"
	case SuspensionExerciseNameSplit:
		return "split"
	case SuspensionExerciseNameSquat:
		return "squat"
	case SuspensionExerciseNameSquatJump:
		return "squat_jump"
	case SuspensionExerciseNameTricepPress:
		return "tricep_press"
	case SuspensionExerciseNameYFly:
		return "y_fly"
	default:
		return "SuspensionExerciseNameInvalid(" + strconv.FormatUint(uint64(s), 10) + ")"
	}
}

// FromString parse string into SuspensionExerciseName constant it's represent, return SuspensionExerciseNameInvalid if not found.
func SuspensionExerciseNameFromString(s string) SuspensionExerciseName {
	switch s {
	case "chest_fly":
		return SuspensionExerciseNameChestFly
	case "chest_press":
		return SuspensionExerciseNameChestPress
	case "crunch":
		return SuspensionExerciseNameCrunch
	case "curl":
		return SuspensionExerciseNameCurl
	case "dip":
		return SuspensionExerciseNameDip
	case "face_pull":
		return SuspensionExerciseNameFacePull
	case "glute_bridge":
		return SuspensionExerciseNameGluteBridge
	case "hamstring_curl":
		return SuspensionExerciseNameHamstringCurl
	case "hip_drop":
		return SuspensionExerciseNameHipDrop
	case "inverted_row":
		return SuspensionExerciseNameInvertedRow
	case "knee_drive_jump":
		return SuspensionExerciseNameKneeDriveJump
	case "knee_to_chest":
		return SuspensionExerciseNameKneeToChest
	case "lat_pullover":
		return SuspensionExerciseNameLatPullover
	case "lunge":
		return SuspensionExerciseNameLunge
	case "mountain_climber":
		return SuspensionExerciseNameMountainClimber
	case "pendulum":
		return SuspensionExerciseNamePendulum
	case "pike":
		return SuspensionExerciseNamePike
	case "plank":
		return SuspensionExerciseNamePlank
	case "power_pull":
		return SuspensionExerciseNamePowerPull
	case "pull_up":
		return SuspensionExerciseNamePullUp
	case "push_up":
		return SuspensionExerciseNamePushUp
	case "reverse_mountain_climber":
		return SuspensionExerciseNameReverseMountainClimber
	case "reverse_plank":
		return SuspensionExerciseNameReversePlank
	case "rollout":
		return SuspensionExerciseNameRollout
	case "row":
		return SuspensionExerciseNameRow
	case "side_lunge":
		return SuspensionExerciseNameSideLunge
	case "side_plank":
		return SuspensionExerciseNameSidePlank
	case "single_leg_deadlift":
		return SuspensionExerciseNameSingleLegDeadlift
	case "single_leg_squat":
		return SuspensionExerciseNameSingleLegSquat
	case "sit_up":
		return SuspensionExerciseNameSitUp
	case "split":
		return SuspensionExerciseNameSplit
	case "squat":
		return SuspensionExerciseNameSquat
	case "squat_jump":
		return SuspensionExerciseNameSquatJump
	case "tricep_press":
		return SuspensionExerciseNameTricepPress
	case "y_fly":
		return SuspensionExerciseNameYFly
	default:
		return SuspensionExerciseNameInvalid
	}
}

// List returns all constants.
func ListSuspensionExerciseName() []SuspensionExerciseName {
	return []SuspensionExerciseName{
		SuspensionExerciseNameChestFly,
		SuspensionExerciseNameChestPress,
		SuspensionExerciseNameCrunch,
		SuspensionExerciseNameCurl,
		SuspensionExerciseNameDip,
		SuspensionExerciseNameFacePull,
		SuspensionExerciseNameGluteBridge,
		SuspensionExerciseNameHamstringCurl,
		SuspensionExerciseNameHipDrop,
		SuspensionExerciseNameInvertedRow,
		SuspensionExerciseNameKneeDriveJump,
		SuspensionExerciseNameKneeToChest,
		SuspensionExerciseNameLatPullover,
		SuspensionExerciseNameLunge,
		SuspensionExerciseNameMountainClimber,
		SuspensionExerciseNamePendulum,
		SuspensionExerciseNamePike,
		SuspensionExerciseNamePlank,
		SuspensionExerciseNamePowerPull,
		SuspensionExerciseNamePullUp,
		SuspensionExerciseNamePushUp,
		SuspensionExerciseNameReverseMountainClimber,
		SuspensionExerciseNameReversePlank,
		SuspensionExerciseNameRollout,
		SuspensionExerciseNameRow,
		SuspensionExerciseNameSideLunge,
		SuspensionExerciseNameSidePlank,
		SuspensionExerciseNameSingleLegDeadlift,
		SuspensionExerciseNameSingleLegSquat,
		SuspensionExerciseNameSitUp,
		SuspensionExerciseNameSplit,
		SuspensionExerciseNameSquat,
		SuspensionExerciseNameSquatJump,
		SuspensionExerciseNameTricepPress,
		SuspensionExerciseNameYFly,
	}
}
