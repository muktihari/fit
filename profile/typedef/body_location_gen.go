// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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
	BodyLocationInvalid               BodyLocation = 0xFF // INVALID
)

var bodylocationtostrs = map[BodyLocation]string{
	BodyLocationLeftLeg:               "left_leg",
	BodyLocationLeftCalf:              "left_calf",
	BodyLocationLeftShin:              "left_shin",
	BodyLocationLeftHamstring:         "left_hamstring",
	BodyLocationLeftQuad:              "left_quad",
	BodyLocationLeftGlute:             "left_glute",
	BodyLocationRightLeg:              "right_leg",
	BodyLocationRightCalf:             "right_calf",
	BodyLocationRightShin:             "right_shin",
	BodyLocationRightHamstring:        "right_hamstring",
	BodyLocationRightQuad:             "right_quad",
	BodyLocationRightGlute:            "right_glute",
	BodyLocationTorsoBack:             "torso_back",
	BodyLocationLeftLowerBack:         "left_lower_back",
	BodyLocationLeftUpperBack:         "left_upper_back",
	BodyLocationRightLowerBack:        "right_lower_back",
	BodyLocationRightUpperBack:        "right_upper_back",
	BodyLocationTorsoFront:            "torso_front",
	BodyLocationLeftAbdomen:           "left_abdomen",
	BodyLocationLeftChest:             "left_chest",
	BodyLocationRightAbdomen:          "right_abdomen",
	BodyLocationRightChest:            "right_chest",
	BodyLocationLeftArm:               "left_arm",
	BodyLocationLeftShoulder:          "left_shoulder",
	BodyLocationLeftBicep:             "left_bicep",
	BodyLocationLeftTricep:            "left_tricep",
	BodyLocationLeftBrachioradialis:   "left_brachioradialis",
	BodyLocationLeftForearmExtensors:  "left_forearm_extensors",
	BodyLocationRightArm:              "right_arm",
	BodyLocationRightShoulder:         "right_shoulder",
	BodyLocationRightBicep:            "right_bicep",
	BodyLocationRightTricep:           "right_tricep",
	BodyLocationRightBrachioradialis:  "right_brachioradialis",
	BodyLocationRightForearmExtensors: "right_forearm_extensors",
	BodyLocationNeck:                  "neck",
	BodyLocationThroat:                "throat",
	BodyLocationWaistMidBack:          "waist_mid_back",
	BodyLocationWaistFront:            "waist_front",
	BodyLocationWaistLeft:             "waist_left",
	BodyLocationWaistRight:            "waist_right",
	BodyLocationInvalid:               "invalid",
}

func (b BodyLocation) String() string {
	val, ok := bodylocationtostrs[b]
	if !ok {
		return strconv.Itoa(int(b))
	}
	return val
}

var strtobodylocation = func() map[string]BodyLocation {
	m := make(map[string]BodyLocation)
	for t, str := range bodylocationtostrs {
		m[str] = BodyLocation(t)
	}
	return m
}()

// FromString parse string into BodyLocation constant it's represent, return BodyLocationInvalid if not found.
func BodyLocationFromString(s string) BodyLocation {
	val, ok := strtobodylocation[s]
	if !ok {
		return strtobodylocation["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListBodyLocation() []BodyLocation {
	vs := make([]BodyLocation, 0, len(bodylocationtostrs))
	for i := range bodylocationtostrs {
		vs = append(vs, BodyLocation(i))
	}
	return vs
}