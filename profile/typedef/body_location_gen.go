// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BodyLocation byte

const (
	BodyLocationLeftLeg               BodyLocation = 0
	BodyLocationLeftCalf              BodyLocation = 1
	BodyLocationLeftShin              BodyLocation = 2
	BodyLocationLeftHamstring         BodyLocation = 3
	BodyLocationLeftQuad              BodyLocation = 4
	BodyLocationLeftGlute             BodyLocation = 5
	BodyLocationRightLeg              BodyLocation = 6
	BodyLocationRightCalf             BodyLocation = 7
	BodyLocationRightShin             BodyLocation = 8
	BodyLocationRightHamstring        BodyLocation = 9
	BodyLocationRightQuad             BodyLocation = 10
	BodyLocationRightGlute            BodyLocation = 11
	BodyLocationTorsoBack             BodyLocation = 12
	BodyLocationLeftLowerBack         BodyLocation = 13
	BodyLocationLeftUpperBack         BodyLocation = 14
	BodyLocationRightLowerBack        BodyLocation = 15
	BodyLocationRightUpperBack        BodyLocation = 16
	BodyLocationTorsoFront            BodyLocation = 17
	BodyLocationLeftAbdomen           BodyLocation = 18
	BodyLocationLeftChest             BodyLocation = 19
	BodyLocationRightAbdomen          BodyLocation = 20
	BodyLocationRightChest            BodyLocation = 21
	BodyLocationLeftArm               BodyLocation = 22
	BodyLocationLeftShoulder          BodyLocation = 23
	BodyLocationLeftBicep             BodyLocation = 24
	BodyLocationLeftTricep            BodyLocation = 25
	BodyLocationLeftBrachioradialis   BodyLocation = 26 // Left anterior forearm
	BodyLocationLeftForearmExtensors  BodyLocation = 27 // Left posterior forearm
	BodyLocationRightArm              BodyLocation = 28
	BodyLocationRightShoulder         BodyLocation = 29
	BodyLocationRightBicep            BodyLocation = 30
	BodyLocationRightTricep           BodyLocation = 31
	BodyLocationRightBrachioradialis  BodyLocation = 32 // Right anterior forearm
	BodyLocationRightForearmExtensors BodyLocation = 33 // Right posterior forearm
	BodyLocationNeck                  BodyLocation = 34
	BodyLocationThroat                BodyLocation = 35
	BodyLocationWaistMidBack          BodyLocation = 36
	BodyLocationWaistFront            BodyLocation = 37
	BodyLocationWaistLeft             BodyLocation = 38
	BodyLocationWaistRight            BodyLocation = 39
	BodyLocationInvalid               BodyLocation = 0xFF
)

func (b BodyLocation) Byte() byte { return byte(b) }

func (b BodyLocation) String() string {
	switch b {
	case BodyLocationLeftLeg:
		return "left_leg"
	case BodyLocationLeftCalf:
		return "left_calf"
	case BodyLocationLeftShin:
		return "left_shin"
	case BodyLocationLeftHamstring:
		return "left_hamstring"
	case BodyLocationLeftQuad:
		return "left_quad"
	case BodyLocationLeftGlute:
		return "left_glute"
	case BodyLocationRightLeg:
		return "right_leg"
	case BodyLocationRightCalf:
		return "right_calf"
	case BodyLocationRightShin:
		return "right_shin"
	case BodyLocationRightHamstring:
		return "right_hamstring"
	case BodyLocationRightQuad:
		return "right_quad"
	case BodyLocationRightGlute:
		return "right_glute"
	case BodyLocationTorsoBack:
		return "torso_back"
	case BodyLocationLeftLowerBack:
		return "left_lower_back"
	case BodyLocationLeftUpperBack:
		return "left_upper_back"
	case BodyLocationRightLowerBack:
		return "right_lower_back"
	case BodyLocationRightUpperBack:
		return "right_upper_back"
	case BodyLocationTorsoFront:
		return "torso_front"
	case BodyLocationLeftAbdomen:
		return "left_abdomen"
	case BodyLocationLeftChest:
		return "left_chest"
	case BodyLocationRightAbdomen:
		return "right_abdomen"
	case BodyLocationRightChest:
		return "right_chest"
	case BodyLocationLeftArm:
		return "left_arm"
	case BodyLocationLeftShoulder:
		return "left_shoulder"
	case BodyLocationLeftBicep:
		return "left_bicep"
	case BodyLocationLeftTricep:
		return "left_tricep"
	case BodyLocationLeftBrachioradialis:
		return "left_brachioradialis"
	case BodyLocationLeftForearmExtensors:
		return "left_forearm_extensors"
	case BodyLocationRightArm:
		return "right_arm"
	case BodyLocationRightShoulder:
		return "right_shoulder"
	case BodyLocationRightBicep:
		return "right_bicep"
	case BodyLocationRightTricep:
		return "right_tricep"
	case BodyLocationRightBrachioradialis:
		return "right_brachioradialis"
	case BodyLocationRightForearmExtensors:
		return "right_forearm_extensors"
	case BodyLocationNeck:
		return "neck"
	case BodyLocationThroat:
		return "throat"
	case BodyLocationWaistMidBack:
		return "waist_mid_back"
	case BodyLocationWaistFront:
		return "waist_front"
	case BodyLocationWaistLeft:
		return "waist_left"
	case BodyLocationWaistRight:
		return "waist_right"
	default:
		return "BodyLocationInvalid(" + strconv.Itoa(int(b)) + ")"
	}
}

// FromString parse string into BodyLocation constant it's represent, return BodyLocationInvalid if not found.
func BodyLocationFromString(s string) BodyLocation {
	switch s {
	case "left_leg":
		return BodyLocationLeftLeg
	case "left_calf":
		return BodyLocationLeftCalf
	case "left_shin":
		return BodyLocationLeftShin
	case "left_hamstring":
		return BodyLocationLeftHamstring
	case "left_quad":
		return BodyLocationLeftQuad
	case "left_glute":
		return BodyLocationLeftGlute
	case "right_leg":
		return BodyLocationRightLeg
	case "right_calf":
		return BodyLocationRightCalf
	case "right_shin":
		return BodyLocationRightShin
	case "right_hamstring":
		return BodyLocationRightHamstring
	case "right_quad":
		return BodyLocationRightQuad
	case "right_glute":
		return BodyLocationRightGlute
	case "torso_back":
		return BodyLocationTorsoBack
	case "left_lower_back":
		return BodyLocationLeftLowerBack
	case "left_upper_back":
		return BodyLocationLeftUpperBack
	case "right_lower_back":
		return BodyLocationRightLowerBack
	case "right_upper_back":
		return BodyLocationRightUpperBack
	case "torso_front":
		return BodyLocationTorsoFront
	case "left_abdomen":
		return BodyLocationLeftAbdomen
	case "left_chest":
		return BodyLocationLeftChest
	case "right_abdomen":
		return BodyLocationRightAbdomen
	case "right_chest":
		return BodyLocationRightChest
	case "left_arm":
		return BodyLocationLeftArm
	case "left_shoulder":
		return BodyLocationLeftShoulder
	case "left_bicep":
		return BodyLocationLeftBicep
	case "left_tricep":
		return BodyLocationLeftTricep
	case "left_brachioradialis":
		return BodyLocationLeftBrachioradialis
	case "left_forearm_extensors":
		return BodyLocationLeftForearmExtensors
	case "right_arm":
		return BodyLocationRightArm
	case "right_shoulder":
		return BodyLocationRightShoulder
	case "right_bicep":
		return BodyLocationRightBicep
	case "right_tricep":
		return BodyLocationRightTricep
	case "right_brachioradialis":
		return BodyLocationRightBrachioradialis
	case "right_forearm_extensors":
		return BodyLocationRightForearmExtensors
	case "neck":
		return BodyLocationNeck
	case "throat":
		return BodyLocationThroat
	case "waist_mid_back":
		return BodyLocationWaistMidBack
	case "waist_front":
		return BodyLocationWaistFront
	case "waist_left":
		return BodyLocationWaistLeft
	case "waist_right":
		return BodyLocationWaistRight
	default:
		return BodyLocationInvalid
	}
}

// List returns all constants.
func ListBodyLocation() []BodyLocation {
	return []BodyLocation{
		BodyLocationLeftLeg,
		BodyLocationLeftCalf,
		BodyLocationLeftShin,
		BodyLocationLeftHamstring,
		BodyLocationLeftQuad,
		BodyLocationLeftGlute,
		BodyLocationRightLeg,
		BodyLocationRightCalf,
		BodyLocationRightShin,
		BodyLocationRightHamstring,
		BodyLocationRightQuad,
		BodyLocationRightGlute,
		BodyLocationTorsoBack,
		BodyLocationLeftLowerBack,
		BodyLocationLeftUpperBack,
		BodyLocationRightLowerBack,
		BodyLocationRightUpperBack,
		BodyLocationTorsoFront,
		BodyLocationLeftAbdomen,
		BodyLocationLeftChest,
		BodyLocationRightAbdomen,
		BodyLocationRightChest,
		BodyLocationLeftArm,
		BodyLocationLeftShoulder,
		BodyLocationLeftBicep,
		BodyLocationLeftTricep,
		BodyLocationLeftBrachioradialis,
		BodyLocationLeftForearmExtensors,
		BodyLocationRightArm,
		BodyLocationRightShoulder,
		BodyLocationRightBicep,
		BodyLocationRightTricep,
		BodyLocationRightBrachioradialis,
		BodyLocationRightForearmExtensors,
		BodyLocationNeck,
		BodyLocationThroat,
		BodyLocationWaistMidBack,
		BodyLocationWaistFront,
		BodyLocationWaistLeft,
		BodyLocationWaistRight,
	}
}
