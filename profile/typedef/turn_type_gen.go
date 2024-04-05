// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TurnType byte

const (
	TurnTypeArrivingIdx             TurnType = 0
	TurnTypeArrivingLeftIdx         TurnType = 1
	TurnTypeArrivingRightIdx        TurnType = 2
	TurnTypeArrivingViaIdx          TurnType = 3
	TurnTypeArrivingViaLeftIdx      TurnType = 4
	TurnTypeArrivingViaRightIdx     TurnType = 5
	TurnTypeBearKeepLeftIdx         TurnType = 6
	TurnTypeBearKeepRightIdx        TurnType = 7
	TurnTypeContinueIdx             TurnType = 8
	TurnTypeExitLeftIdx             TurnType = 9
	TurnTypeExitRightIdx            TurnType = 10
	TurnTypeFerryIdx                TurnType = 11
	TurnTypeRoundabout45Idx         TurnType = 12
	TurnTypeRoundabout90Idx         TurnType = 13
	TurnTypeRoundabout135Idx        TurnType = 14
	TurnTypeRoundabout180Idx        TurnType = 15
	TurnTypeRoundabout225Idx        TurnType = 16
	TurnTypeRoundabout270Idx        TurnType = 17
	TurnTypeRoundabout315Idx        TurnType = 18
	TurnTypeRoundabout360Idx        TurnType = 19
	TurnTypeRoundaboutNeg45Idx      TurnType = 20
	TurnTypeRoundaboutNeg90Idx      TurnType = 21
	TurnTypeRoundaboutNeg135Idx     TurnType = 22
	TurnTypeRoundaboutNeg180Idx     TurnType = 23
	TurnTypeRoundaboutNeg225Idx     TurnType = 24
	TurnTypeRoundaboutNeg270Idx     TurnType = 25
	TurnTypeRoundaboutNeg315Idx     TurnType = 26
	TurnTypeRoundaboutNeg360Idx     TurnType = 27
	TurnTypeRoundaboutGenericIdx    TurnType = 28
	TurnTypeRoundaboutNegGenericIdx TurnType = 29
	TurnTypeSharpTurnLeftIdx        TurnType = 30
	TurnTypeSharpTurnRightIdx       TurnType = 31
	TurnTypeTurnLeftIdx             TurnType = 32
	TurnTypeTurnRightIdx            TurnType = 33
	TurnTypeUturnLeftIdx            TurnType = 34
	TurnTypeUturnRightIdx           TurnType = 35
	TurnTypeIconInvIdx              TurnType = 36
	TurnTypeIconIdxCnt              TurnType = 37
	TurnTypeInvalid                 TurnType = 0xFF
)

func (t TurnType) Byte() byte { return byte(t) }

func (t TurnType) String() string {
	switch t {
	case TurnTypeArrivingIdx:
		return "arriving_idx"
	case TurnTypeArrivingLeftIdx:
		return "arriving_left_idx"
	case TurnTypeArrivingRightIdx:
		return "arriving_right_idx"
	case TurnTypeArrivingViaIdx:
		return "arriving_via_idx"
	case TurnTypeArrivingViaLeftIdx:
		return "arriving_via_left_idx"
	case TurnTypeArrivingViaRightIdx:
		return "arriving_via_right_idx"
	case TurnTypeBearKeepLeftIdx:
		return "bear_keep_left_idx"
	case TurnTypeBearKeepRightIdx:
		return "bear_keep_right_idx"
	case TurnTypeContinueIdx:
		return "continue_idx"
	case TurnTypeExitLeftIdx:
		return "exit_left_idx"
	case TurnTypeExitRightIdx:
		return "exit_right_idx"
	case TurnTypeFerryIdx:
		return "ferry_idx"
	case TurnTypeRoundabout45Idx:
		return "roundabout_45_idx"
	case TurnTypeRoundabout90Idx:
		return "roundabout_90_idx"
	case TurnTypeRoundabout135Idx:
		return "roundabout_135_idx"
	case TurnTypeRoundabout180Idx:
		return "roundabout_180_idx"
	case TurnTypeRoundabout225Idx:
		return "roundabout_225_idx"
	case TurnTypeRoundabout270Idx:
		return "roundabout_270_idx"
	case TurnTypeRoundabout315Idx:
		return "roundabout_315_idx"
	case TurnTypeRoundabout360Idx:
		return "roundabout_360_idx"
	case TurnTypeRoundaboutNeg45Idx:
		return "roundabout_neg_45_idx"
	case TurnTypeRoundaboutNeg90Idx:
		return "roundabout_neg_90_idx"
	case TurnTypeRoundaboutNeg135Idx:
		return "roundabout_neg_135_idx"
	case TurnTypeRoundaboutNeg180Idx:
		return "roundabout_neg_180_idx"
	case TurnTypeRoundaboutNeg225Idx:
		return "roundabout_neg_225_idx"
	case TurnTypeRoundaboutNeg270Idx:
		return "roundabout_neg_270_idx"
	case TurnTypeRoundaboutNeg315Idx:
		return "roundabout_neg_315_idx"
	case TurnTypeRoundaboutNeg360Idx:
		return "roundabout_neg_360_idx"
	case TurnTypeRoundaboutGenericIdx:
		return "roundabout_generic_idx"
	case TurnTypeRoundaboutNegGenericIdx:
		return "roundabout_neg_generic_idx"
	case TurnTypeSharpTurnLeftIdx:
		return "sharp_turn_left_idx"
	case TurnTypeSharpTurnRightIdx:
		return "sharp_turn_right_idx"
	case TurnTypeTurnLeftIdx:
		return "turn_left_idx"
	case TurnTypeTurnRightIdx:
		return "turn_right_idx"
	case TurnTypeUturnLeftIdx:
		return "uturn_left_idx"
	case TurnTypeUturnRightIdx:
		return "uturn_right_idx"
	case TurnTypeIconInvIdx:
		return "icon_inv_idx"
	case TurnTypeIconIdxCnt:
		return "icon_idx_cnt"
	default:
		return "TurnTypeInvalid(" + strconv.Itoa(int(t)) + ")"
	}
}

// FromString parse string into TurnType constant it's represent, return TurnTypeInvalid if not found.
func TurnTypeFromString(s string) TurnType {
	switch s {
	case "arriving_idx":
		return TurnTypeArrivingIdx
	case "arriving_left_idx":
		return TurnTypeArrivingLeftIdx
	case "arriving_right_idx":
		return TurnTypeArrivingRightIdx
	case "arriving_via_idx":
		return TurnTypeArrivingViaIdx
	case "arriving_via_left_idx":
		return TurnTypeArrivingViaLeftIdx
	case "arriving_via_right_idx":
		return TurnTypeArrivingViaRightIdx
	case "bear_keep_left_idx":
		return TurnTypeBearKeepLeftIdx
	case "bear_keep_right_idx":
		return TurnTypeBearKeepRightIdx
	case "continue_idx":
		return TurnTypeContinueIdx
	case "exit_left_idx":
		return TurnTypeExitLeftIdx
	case "exit_right_idx":
		return TurnTypeExitRightIdx
	case "ferry_idx":
		return TurnTypeFerryIdx
	case "roundabout_45_idx":
		return TurnTypeRoundabout45Idx
	case "roundabout_90_idx":
		return TurnTypeRoundabout90Idx
	case "roundabout_135_idx":
		return TurnTypeRoundabout135Idx
	case "roundabout_180_idx":
		return TurnTypeRoundabout180Idx
	case "roundabout_225_idx":
		return TurnTypeRoundabout225Idx
	case "roundabout_270_idx":
		return TurnTypeRoundabout270Idx
	case "roundabout_315_idx":
		return TurnTypeRoundabout315Idx
	case "roundabout_360_idx":
		return TurnTypeRoundabout360Idx
	case "roundabout_neg_45_idx":
		return TurnTypeRoundaboutNeg45Idx
	case "roundabout_neg_90_idx":
		return TurnTypeRoundaboutNeg90Idx
	case "roundabout_neg_135_idx":
		return TurnTypeRoundaboutNeg135Idx
	case "roundabout_neg_180_idx":
		return TurnTypeRoundaboutNeg180Idx
	case "roundabout_neg_225_idx":
		return TurnTypeRoundaboutNeg225Idx
	case "roundabout_neg_270_idx":
		return TurnTypeRoundaboutNeg270Idx
	case "roundabout_neg_315_idx":
		return TurnTypeRoundaboutNeg315Idx
	case "roundabout_neg_360_idx":
		return TurnTypeRoundaboutNeg360Idx
	case "roundabout_generic_idx":
		return TurnTypeRoundaboutGenericIdx
	case "roundabout_neg_generic_idx":
		return TurnTypeRoundaboutNegGenericIdx
	case "sharp_turn_left_idx":
		return TurnTypeSharpTurnLeftIdx
	case "sharp_turn_right_idx":
		return TurnTypeSharpTurnRightIdx
	case "turn_left_idx":
		return TurnTypeTurnLeftIdx
	case "turn_right_idx":
		return TurnTypeTurnRightIdx
	case "uturn_left_idx":
		return TurnTypeUturnLeftIdx
	case "uturn_right_idx":
		return TurnTypeUturnRightIdx
	case "icon_inv_idx":
		return TurnTypeIconInvIdx
	case "icon_idx_cnt":
		return TurnTypeIconIdxCnt
	default:
		return TurnTypeInvalid
	}
}

// List returns all constants.
func ListTurnType() []TurnType {
	return []TurnType{
		TurnTypeArrivingIdx,
		TurnTypeArrivingLeftIdx,
		TurnTypeArrivingRightIdx,
		TurnTypeArrivingViaIdx,
		TurnTypeArrivingViaLeftIdx,
		TurnTypeArrivingViaRightIdx,
		TurnTypeBearKeepLeftIdx,
		TurnTypeBearKeepRightIdx,
		TurnTypeContinueIdx,
		TurnTypeExitLeftIdx,
		TurnTypeExitRightIdx,
		TurnTypeFerryIdx,
		TurnTypeRoundabout45Idx,
		TurnTypeRoundabout90Idx,
		TurnTypeRoundabout135Idx,
		TurnTypeRoundabout180Idx,
		TurnTypeRoundabout225Idx,
		TurnTypeRoundabout270Idx,
		TurnTypeRoundabout315Idx,
		TurnTypeRoundabout360Idx,
		TurnTypeRoundaboutNeg45Idx,
		TurnTypeRoundaboutNeg90Idx,
		TurnTypeRoundaboutNeg135Idx,
		TurnTypeRoundaboutNeg180Idx,
		TurnTypeRoundaboutNeg225Idx,
		TurnTypeRoundaboutNeg270Idx,
		TurnTypeRoundaboutNeg315Idx,
		TurnTypeRoundaboutNeg360Idx,
		TurnTypeRoundaboutGenericIdx,
		TurnTypeRoundaboutNegGenericIdx,
		TurnTypeSharpTurnLeftIdx,
		TurnTypeSharpTurnRightIdx,
		TurnTypeTurnLeftIdx,
		TurnTypeTurnRightIdx,
		TurnTypeUturnLeftIdx,
		TurnTypeUturnRightIdx,
		TurnTypeIconInvIdx,
		TurnTypeIconIdxCnt,
	}
}
