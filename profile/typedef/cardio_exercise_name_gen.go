// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CardioExerciseName uint16

const (
	CardioExerciseNameBobAndWeaveCircle         CardioExerciseName = 0
	CardioExerciseNameWeightedBobAndWeaveCircle CardioExerciseName = 1
	CardioExerciseNameCardioCoreCrawl           CardioExerciseName = 2
	CardioExerciseNameWeightedCardioCoreCrawl   CardioExerciseName = 3
	CardioExerciseNameDoubleUnder               CardioExerciseName = 4
	CardioExerciseNameWeightedDoubleUnder       CardioExerciseName = 5
	CardioExerciseNameJumpRope                  CardioExerciseName = 6
	CardioExerciseNameWeightedJumpRope          CardioExerciseName = 7
	CardioExerciseNameJumpRopeCrossover         CardioExerciseName = 8
	CardioExerciseNameWeightedJumpRopeCrossover CardioExerciseName = 9
	CardioExerciseNameJumpRopeJog               CardioExerciseName = 10
	CardioExerciseNameWeightedJumpRopeJog       CardioExerciseName = 11
	CardioExerciseNameJumpingJacks              CardioExerciseName = 12
	CardioExerciseNameWeightedJumpingJacks      CardioExerciseName = 13
	CardioExerciseNameSkiMoguls                 CardioExerciseName = 14
	CardioExerciseNameWeightedSkiMoguls         CardioExerciseName = 15
	CardioExerciseNameSplitJacks                CardioExerciseName = 16
	CardioExerciseNameWeightedSplitJacks        CardioExerciseName = 17
	CardioExerciseNameSquatJacks                CardioExerciseName = 18
	CardioExerciseNameWeightedSquatJacks        CardioExerciseName = 19
	CardioExerciseNameTripleUnder               CardioExerciseName = 20
	CardioExerciseNameWeightedTripleUnder       CardioExerciseName = 21
	CardioExerciseNameInvalid                   CardioExerciseName = 0xFFFF // INVALID
)

var cardioexercisenametostrs = map[CardioExerciseName]string{
	CardioExerciseNameBobAndWeaveCircle:         "bob_and_weave_circle",
	CardioExerciseNameWeightedBobAndWeaveCircle: "weighted_bob_and_weave_circle",
	CardioExerciseNameCardioCoreCrawl:           "cardio_core_crawl",
	CardioExerciseNameWeightedCardioCoreCrawl:   "weighted_cardio_core_crawl",
	CardioExerciseNameDoubleUnder:               "double_under",
	CardioExerciseNameWeightedDoubleUnder:       "weighted_double_under",
	CardioExerciseNameJumpRope:                  "jump_rope",
	CardioExerciseNameWeightedJumpRope:          "weighted_jump_rope",
	CardioExerciseNameJumpRopeCrossover:         "jump_rope_crossover",
	CardioExerciseNameWeightedJumpRopeCrossover: "weighted_jump_rope_crossover",
	CardioExerciseNameJumpRopeJog:               "jump_rope_jog",
	CardioExerciseNameWeightedJumpRopeJog:       "weighted_jump_rope_jog",
	CardioExerciseNameJumpingJacks:              "jumping_jacks",
	CardioExerciseNameWeightedJumpingJacks:      "weighted_jumping_jacks",
	CardioExerciseNameSkiMoguls:                 "ski_moguls",
	CardioExerciseNameWeightedSkiMoguls:         "weighted_ski_moguls",
	CardioExerciseNameSplitJacks:                "split_jacks",
	CardioExerciseNameWeightedSplitJacks:        "weighted_split_jacks",
	CardioExerciseNameSquatJacks:                "squat_jacks",
	CardioExerciseNameWeightedSquatJacks:        "weighted_squat_jacks",
	CardioExerciseNameTripleUnder:               "triple_under",
	CardioExerciseNameWeightedTripleUnder:       "weighted_triple_under",
	CardioExerciseNameInvalid:                   "invalid",
}

func (c CardioExerciseName) String() string {
	val, ok := cardioexercisenametostrs[c]
	if !ok {
		return strconv.FormatUint(uint64(c), 10)
	}
	return val
}

var strtocardioexercisename = func() map[string]CardioExerciseName {
	m := make(map[string]CardioExerciseName)
	for t, str := range cardioexercisenametostrs {
		m[str] = CardioExerciseName(t)
	}
	return m
}()

// FromString parse string into CardioExerciseName constant it's represent, return CardioExerciseNameInvalid if not found.
func CardioExerciseNameFromString(s string) CardioExerciseName {
	val, ok := strtocardioexercisename[s]
	if !ok {
		return strtocardioexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListCardioExerciseName() []CardioExerciseName {
	vs := make([]CardioExerciseName, 0, len(cardioexercisenametostrs))
	for i := range cardioexercisenametostrs {
		vs = append(vs, CardioExerciseName(i))
	}
	return vs
}
