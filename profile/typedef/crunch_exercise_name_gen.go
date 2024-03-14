// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CrunchExerciseName uint16

const (
	CrunchExerciseNameBicycleCrunch                           CrunchExerciseName = 0
	CrunchExerciseNameCableCrunch                             CrunchExerciseName = 1
	CrunchExerciseNameCircularArmCrunch                       CrunchExerciseName = 2
	CrunchExerciseNameCrossedArmsCrunch                       CrunchExerciseName = 3
	CrunchExerciseNameWeightedCrossedArmsCrunch               CrunchExerciseName = 4
	CrunchExerciseNameCrossLegReverseCrunch                   CrunchExerciseName = 5
	CrunchExerciseNameWeightedCrossLegReverseCrunch           CrunchExerciseName = 6
	CrunchExerciseNameCrunchChop                              CrunchExerciseName = 7
	CrunchExerciseNameWeightedCrunchChop                      CrunchExerciseName = 8
	CrunchExerciseNameDoubleCrunch                            CrunchExerciseName = 9
	CrunchExerciseNameWeightedDoubleCrunch                    CrunchExerciseName = 10
	CrunchExerciseNameElbowToKneeCrunch                       CrunchExerciseName = 11
	CrunchExerciseNameWeightedElbowToKneeCrunch               CrunchExerciseName = 12
	CrunchExerciseNameFlutterKicks                            CrunchExerciseName = 13
	CrunchExerciseNameWeightedFlutterKicks                    CrunchExerciseName = 14
	CrunchExerciseNameFoamRollerReverseCrunchOnBench          CrunchExerciseName = 15
	CrunchExerciseNameWeightedFoamRollerReverseCrunchOnBench  CrunchExerciseName = 16
	CrunchExerciseNameFoamRollerReverseCrunchWithDumbbell     CrunchExerciseName = 17
	CrunchExerciseNameFoamRollerReverseCrunchWithMedicineBall CrunchExerciseName = 18
	CrunchExerciseNameFrogPress                               CrunchExerciseName = 19
	CrunchExerciseNameHangingKneeRaiseObliqueCrunch           CrunchExerciseName = 20
	CrunchExerciseNameWeightedHangingKneeRaiseObliqueCrunch   CrunchExerciseName = 21
	CrunchExerciseNameHipCrossover                            CrunchExerciseName = 22
	CrunchExerciseNameWeightedHipCrossover                    CrunchExerciseName = 23
	CrunchExerciseNameHollowRock                              CrunchExerciseName = 24
	CrunchExerciseNameWeightedHollowRock                      CrunchExerciseName = 25
	CrunchExerciseNameInclineReverseCrunch                    CrunchExerciseName = 26
	CrunchExerciseNameWeightedInclineReverseCrunch            CrunchExerciseName = 27
	CrunchExerciseNameKneelingCableCrunch                     CrunchExerciseName = 28
	CrunchExerciseNameKneelingCrossCrunch                     CrunchExerciseName = 29
	CrunchExerciseNameWeightedKneelingCrossCrunch             CrunchExerciseName = 30
	CrunchExerciseNameKneelingObliqueCableCrunch              CrunchExerciseName = 31
	CrunchExerciseNameKneesToElbow                            CrunchExerciseName = 32
	CrunchExerciseNameLegExtensions                           CrunchExerciseName = 33
	CrunchExerciseNameWeightedLegExtensions                   CrunchExerciseName = 34
	CrunchExerciseNameLegLevers                               CrunchExerciseName = 35
	CrunchExerciseNameMcgillCurlUp                            CrunchExerciseName = 36
	CrunchExerciseNameWeightedMcgillCurlUp                    CrunchExerciseName = 37
	CrunchExerciseNameModifiedPilatesRollUpWithBall           CrunchExerciseName = 38
	CrunchExerciseNameWeightedModifiedPilatesRollUpWithBall   CrunchExerciseName = 39
	CrunchExerciseNamePilatesCrunch                           CrunchExerciseName = 40
	CrunchExerciseNameWeightedPilatesCrunch                   CrunchExerciseName = 41
	CrunchExerciseNamePilatesRollUpWithBall                   CrunchExerciseName = 42
	CrunchExerciseNameWeightedPilatesRollUpWithBall           CrunchExerciseName = 43
	CrunchExerciseNameRaisedLegsCrunch                        CrunchExerciseName = 44
	CrunchExerciseNameWeightedRaisedLegsCrunch                CrunchExerciseName = 45
	CrunchExerciseNameReverseCrunch                           CrunchExerciseName = 46
	CrunchExerciseNameWeightedReverseCrunch                   CrunchExerciseName = 47
	CrunchExerciseNameReverseCrunchOnABench                   CrunchExerciseName = 48
	CrunchExerciseNameWeightedReverseCrunchOnABench           CrunchExerciseName = 49
	CrunchExerciseNameReverseCurlAndLift                      CrunchExerciseName = 50
	CrunchExerciseNameWeightedReverseCurlAndLift              CrunchExerciseName = 51
	CrunchExerciseNameRotationalLift                          CrunchExerciseName = 52
	CrunchExerciseNameWeightedRotationalLift                  CrunchExerciseName = 53
	CrunchExerciseNameSeatedAlternatingReverseCrunch          CrunchExerciseName = 54
	CrunchExerciseNameWeightedSeatedAlternatingReverseCrunch  CrunchExerciseName = 55
	CrunchExerciseNameSeatedLegU                              CrunchExerciseName = 56
	CrunchExerciseNameWeightedSeatedLegU                      CrunchExerciseName = 57
	CrunchExerciseNameSideToSideCrunchAndWeave                CrunchExerciseName = 58
	CrunchExerciseNameWeightedSideToSideCrunchAndWeave        CrunchExerciseName = 59
	CrunchExerciseNameSingleLegReverseCrunch                  CrunchExerciseName = 60
	CrunchExerciseNameWeightedSingleLegReverseCrunch          CrunchExerciseName = 61
	CrunchExerciseNameSkaterCrunchCross                       CrunchExerciseName = 62
	CrunchExerciseNameWeightedSkaterCrunchCross               CrunchExerciseName = 63
	CrunchExerciseNameStandingCableCrunch                     CrunchExerciseName = 64
	CrunchExerciseNameStandingSideCrunch                      CrunchExerciseName = 65
	CrunchExerciseNameStepClimb                               CrunchExerciseName = 66
	CrunchExerciseNameWeightedStepClimb                       CrunchExerciseName = 67
	CrunchExerciseNameSwissBallCrunch                         CrunchExerciseName = 68
	CrunchExerciseNameSwissBallReverseCrunch                  CrunchExerciseName = 69
	CrunchExerciseNameWeightedSwissBallReverseCrunch          CrunchExerciseName = 70
	CrunchExerciseNameSwissBallRussianTwist                   CrunchExerciseName = 71
	CrunchExerciseNameWeightedSwissBallRussianTwist           CrunchExerciseName = 72
	CrunchExerciseNameSwissBallSideCrunch                     CrunchExerciseName = 73
	CrunchExerciseNameWeightedSwissBallSideCrunch             CrunchExerciseName = 74
	CrunchExerciseNameThoracicCrunchesOnFoamRoller            CrunchExerciseName = 75
	CrunchExerciseNameWeightedThoracicCrunchesOnFoamRoller    CrunchExerciseName = 76
	CrunchExerciseNameTricepsCrunch                           CrunchExerciseName = 77
	CrunchExerciseNameWeightedBicycleCrunch                   CrunchExerciseName = 78
	CrunchExerciseNameWeightedCrunch                          CrunchExerciseName = 79
	CrunchExerciseNameWeightedSwissBallCrunch                 CrunchExerciseName = 80
	CrunchExerciseNameToesToBar                               CrunchExerciseName = 81
	CrunchExerciseNameWeightedToesToBar                       CrunchExerciseName = 82
	CrunchExerciseNameCrunch                                  CrunchExerciseName = 83
	CrunchExerciseNameStraightLegCrunchWithBall               CrunchExerciseName = 84
	CrunchExerciseNameInvalid                                 CrunchExerciseName = 0xFFFF
)

func (c CrunchExerciseName) String() string {
	switch c {
	case CrunchExerciseNameBicycleCrunch:
		return "bicycle_crunch"
	case CrunchExerciseNameCableCrunch:
		return "cable_crunch"
	case CrunchExerciseNameCircularArmCrunch:
		return "circular_arm_crunch"
	case CrunchExerciseNameCrossedArmsCrunch:
		return "crossed_arms_crunch"
	case CrunchExerciseNameWeightedCrossedArmsCrunch:
		return "weighted_crossed_arms_crunch"
	case CrunchExerciseNameCrossLegReverseCrunch:
		return "cross_leg_reverse_crunch"
	case CrunchExerciseNameWeightedCrossLegReverseCrunch:
		return "weighted_cross_leg_reverse_crunch"
	case CrunchExerciseNameCrunchChop:
		return "crunch_chop"
	case CrunchExerciseNameWeightedCrunchChop:
		return "weighted_crunch_chop"
	case CrunchExerciseNameDoubleCrunch:
		return "double_crunch"
	case CrunchExerciseNameWeightedDoubleCrunch:
		return "weighted_double_crunch"
	case CrunchExerciseNameElbowToKneeCrunch:
		return "elbow_to_knee_crunch"
	case CrunchExerciseNameWeightedElbowToKneeCrunch:
		return "weighted_elbow_to_knee_crunch"
	case CrunchExerciseNameFlutterKicks:
		return "flutter_kicks"
	case CrunchExerciseNameWeightedFlutterKicks:
		return "weighted_flutter_kicks"
	case CrunchExerciseNameFoamRollerReverseCrunchOnBench:
		return "foam_roller_reverse_crunch_on_bench"
	case CrunchExerciseNameWeightedFoamRollerReverseCrunchOnBench:
		return "weighted_foam_roller_reverse_crunch_on_bench"
	case CrunchExerciseNameFoamRollerReverseCrunchWithDumbbell:
		return "foam_roller_reverse_crunch_with_dumbbell"
	case CrunchExerciseNameFoamRollerReverseCrunchWithMedicineBall:
		return "foam_roller_reverse_crunch_with_medicine_ball"
	case CrunchExerciseNameFrogPress:
		return "frog_press"
	case CrunchExerciseNameHangingKneeRaiseObliqueCrunch:
		return "hanging_knee_raise_oblique_crunch"
	case CrunchExerciseNameWeightedHangingKneeRaiseObliqueCrunch:
		return "weighted_hanging_knee_raise_oblique_crunch"
	case CrunchExerciseNameHipCrossover:
		return "hip_crossover"
	case CrunchExerciseNameWeightedHipCrossover:
		return "weighted_hip_crossover"
	case CrunchExerciseNameHollowRock:
		return "hollow_rock"
	case CrunchExerciseNameWeightedHollowRock:
		return "weighted_hollow_rock"
	case CrunchExerciseNameInclineReverseCrunch:
		return "incline_reverse_crunch"
	case CrunchExerciseNameWeightedInclineReverseCrunch:
		return "weighted_incline_reverse_crunch"
	case CrunchExerciseNameKneelingCableCrunch:
		return "kneeling_cable_crunch"
	case CrunchExerciseNameKneelingCrossCrunch:
		return "kneeling_cross_crunch"
	case CrunchExerciseNameWeightedKneelingCrossCrunch:
		return "weighted_kneeling_cross_crunch"
	case CrunchExerciseNameKneelingObliqueCableCrunch:
		return "kneeling_oblique_cable_crunch"
	case CrunchExerciseNameKneesToElbow:
		return "knees_to_elbow"
	case CrunchExerciseNameLegExtensions:
		return "leg_extensions"
	case CrunchExerciseNameWeightedLegExtensions:
		return "weighted_leg_extensions"
	case CrunchExerciseNameLegLevers:
		return "leg_levers"
	case CrunchExerciseNameMcgillCurlUp:
		return "mcgill_curl_up"
	case CrunchExerciseNameWeightedMcgillCurlUp:
		return "weighted_mcgill_curl_up"
	case CrunchExerciseNameModifiedPilatesRollUpWithBall:
		return "modified_pilates_roll_up_with_ball"
	case CrunchExerciseNameWeightedModifiedPilatesRollUpWithBall:
		return "weighted_modified_pilates_roll_up_with_ball"
	case CrunchExerciseNamePilatesCrunch:
		return "pilates_crunch"
	case CrunchExerciseNameWeightedPilatesCrunch:
		return "weighted_pilates_crunch"
	case CrunchExerciseNamePilatesRollUpWithBall:
		return "pilates_roll_up_with_ball"
	case CrunchExerciseNameWeightedPilatesRollUpWithBall:
		return "weighted_pilates_roll_up_with_ball"
	case CrunchExerciseNameRaisedLegsCrunch:
		return "raised_legs_crunch"
	case CrunchExerciseNameWeightedRaisedLegsCrunch:
		return "weighted_raised_legs_crunch"
	case CrunchExerciseNameReverseCrunch:
		return "reverse_crunch"
	case CrunchExerciseNameWeightedReverseCrunch:
		return "weighted_reverse_crunch"
	case CrunchExerciseNameReverseCrunchOnABench:
		return "reverse_crunch_on_a_bench"
	case CrunchExerciseNameWeightedReverseCrunchOnABench:
		return "weighted_reverse_crunch_on_a_bench"
	case CrunchExerciseNameReverseCurlAndLift:
		return "reverse_curl_and_lift"
	case CrunchExerciseNameWeightedReverseCurlAndLift:
		return "weighted_reverse_curl_and_lift"
	case CrunchExerciseNameRotationalLift:
		return "rotational_lift"
	case CrunchExerciseNameWeightedRotationalLift:
		return "weighted_rotational_lift"
	case CrunchExerciseNameSeatedAlternatingReverseCrunch:
		return "seated_alternating_reverse_crunch"
	case CrunchExerciseNameWeightedSeatedAlternatingReverseCrunch:
		return "weighted_seated_alternating_reverse_crunch"
	case CrunchExerciseNameSeatedLegU:
		return "seated_leg_u"
	case CrunchExerciseNameWeightedSeatedLegU:
		return "weighted_seated_leg_u"
	case CrunchExerciseNameSideToSideCrunchAndWeave:
		return "side_to_side_crunch_and_weave"
	case CrunchExerciseNameWeightedSideToSideCrunchAndWeave:
		return "weighted_side_to_side_crunch_and_weave"
	case CrunchExerciseNameSingleLegReverseCrunch:
		return "single_leg_reverse_crunch"
	case CrunchExerciseNameWeightedSingleLegReverseCrunch:
		return "weighted_single_leg_reverse_crunch"
	case CrunchExerciseNameSkaterCrunchCross:
		return "skater_crunch_cross"
	case CrunchExerciseNameWeightedSkaterCrunchCross:
		return "weighted_skater_crunch_cross"
	case CrunchExerciseNameStandingCableCrunch:
		return "standing_cable_crunch"
	case CrunchExerciseNameStandingSideCrunch:
		return "standing_side_crunch"
	case CrunchExerciseNameStepClimb:
		return "step_climb"
	case CrunchExerciseNameWeightedStepClimb:
		return "weighted_step_climb"
	case CrunchExerciseNameSwissBallCrunch:
		return "swiss_ball_crunch"
	case CrunchExerciseNameSwissBallReverseCrunch:
		return "swiss_ball_reverse_crunch"
	case CrunchExerciseNameWeightedSwissBallReverseCrunch:
		return "weighted_swiss_ball_reverse_crunch"
	case CrunchExerciseNameSwissBallRussianTwist:
		return "swiss_ball_russian_twist"
	case CrunchExerciseNameWeightedSwissBallRussianTwist:
		return "weighted_swiss_ball_russian_twist"
	case CrunchExerciseNameSwissBallSideCrunch:
		return "swiss_ball_side_crunch"
	case CrunchExerciseNameWeightedSwissBallSideCrunch:
		return "weighted_swiss_ball_side_crunch"
	case CrunchExerciseNameThoracicCrunchesOnFoamRoller:
		return "thoracic_crunches_on_foam_roller"
	case CrunchExerciseNameWeightedThoracicCrunchesOnFoamRoller:
		return "weighted_thoracic_crunches_on_foam_roller"
	case CrunchExerciseNameTricepsCrunch:
		return "triceps_crunch"
	case CrunchExerciseNameWeightedBicycleCrunch:
		return "weighted_bicycle_crunch"
	case CrunchExerciseNameWeightedCrunch:
		return "weighted_crunch"
	case CrunchExerciseNameWeightedSwissBallCrunch:
		return "weighted_swiss_ball_crunch"
	case CrunchExerciseNameToesToBar:
		return "toes_to_bar"
	case CrunchExerciseNameWeightedToesToBar:
		return "weighted_toes_to_bar"
	case CrunchExerciseNameCrunch:
		return "crunch"
	case CrunchExerciseNameStraightLegCrunchWithBall:
		return "straight_leg_crunch_with_ball"
	default:
		return "CrunchExerciseNameInvalid(" + strconv.FormatUint(uint64(c), 10) + ")"
	}
}

// FromString parse string into CrunchExerciseName constant it's represent, return CrunchExerciseNameInvalid if not found.
func CrunchExerciseNameFromString(s string) CrunchExerciseName {
	switch s {
	case "bicycle_crunch":
		return CrunchExerciseNameBicycleCrunch
	case "cable_crunch":
		return CrunchExerciseNameCableCrunch
	case "circular_arm_crunch":
		return CrunchExerciseNameCircularArmCrunch
	case "crossed_arms_crunch":
		return CrunchExerciseNameCrossedArmsCrunch
	case "weighted_crossed_arms_crunch":
		return CrunchExerciseNameWeightedCrossedArmsCrunch
	case "cross_leg_reverse_crunch":
		return CrunchExerciseNameCrossLegReverseCrunch
	case "weighted_cross_leg_reverse_crunch":
		return CrunchExerciseNameWeightedCrossLegReverseCrunch
	case "crunch_chop":
		return CrunchExerciseNameCrunchChop
	case "weighted_crunch_chop":
		return CrunchExerciseNameWeightedCrunchChop
	case "double_crunch":
		return CrunchExerciseNameDoubleCrunch
	case "weighted_double_crunch":
		return CrunchExerciseNameWeightedDoubleCrunch
	case "elbow_to_knee_crunch":
		return CrunchExerciseNameElbowToKneeCrunch
	case "weighted_elbow_to_knee_crunch":
		return CrunchExerciseNameWeightedElbowToKneeCrunch
	case "flutter_kicks":
		return CrunchExerciseNameFlutterKicks
	case "weighted_flutter_kicks":
		return CrunchExerciseNameWeightedFlutterKicks
	case "foam_roller_reverse_crunch_on_bench":
		return CrunchExerciseNameFoamRollerReverseCrunchOnBench
	case "weighted_foam_roller_reverse_crunch_on_bench":
		return CrunchExerciseNameWeightedFoamRollerReverseCrunchOnBench
	case "foam_roller_reverse_crunch_with_dumbbell":
		return CrunchExerciseNameFoamRollerReverseCrunchWithDumbbell
	case "foam_roller_reverse_crunch_with_medicine_ball":
		return CrunchExerciseNameFoamRollerReverseCrunchWithMedicineBall
	case "frog_press":
		return CrunchExerciseNameFrogPress
	case "hanging_knee_raise_oblique_crunch":
		return CrunchExerciseNameHangingKneeRaiseObliqueCrunch
	case "weighted_hanging_knee_raise_oblique_crunch":
		return CrunchExerciseNameWeightedHangingKneeRaiseObliqueCrunch
	case "hip_crossover":
		return CrunchExerciseNameHipCrossover
	case "weighted_hip_crossover":
		return CrunchExerciseNameWeightedHipCrossover
	case "hollow_rock":
		return CrunchExerciseNameHollowRock
	case "weighted_hollow_rock":
		return CrunchExerciseNameWeightedHollowRock
	case "incline_reverse_crunch":
		return CrunchExerciseNameInclineReverseCrunch
	case "weighted_incline_reverse_crunch":
		return CrunchExerciseNameWeightedInclineReverseCrunch
	case "kneeling_cable_crunch":
		return CrunchExerciseNameKneelingCableCrunch
	case "kneeling_cross_crunch":
		return CrunchExerciseNameKneelingCrossCrunch
	case "weighted_kneeling_cross_crunch":
		return CrunchExerciseNameWeightedKneelingCrossCrunch
	case "kneeling_oblique_cable_crunch":
		return CrunchExerciseNameKneelingObliqueCableCrunch
	case "knees_to_elbow":
		return CrunchExerciseNameKneesToElbow
	case "leg_extensions":
		return CrunchExerciseNameLegExtensions
	case "weighted_leg_extensions":
		return CrunchExerciseNameWeightedLegExtensions
	case "leg_levers":
		return CrunchExerciseNameLegLevers
	case "mcgill_curl_up":
		return CrunchExerciseNameMcgillCurlUp
	case "weighted_mcgill_curl_up":
		return CrunchExerciseNameWeightedMcgillCurlUp
	case "modified_pilates_roll_up_with_ball":
		return CrunchExerciseNameModifiedPilatesRollUpWithBall
	case "weighted_modified_pilates_roll_up_with_ball":
		return CrunchExerciseNameWeightedModifiedPilatesRollUpWithBall
	case "pilates_crunch":
		return CrunchExerciseNamePilatesCrunch
	case "weighted_pilates_crunch":
		return CrunchExerciseNameWeightedPilatesCrunch
	case "pilates_roll_up_with_ball":
		return CrunchExerciseNamePilatesRollUpWithBall
	case "weighted_pilates_roll_up_with_ball":
		return CrunchExerciseNameWeightedPilatesRollUpWithBall
	case "raised_legs_crunch":
		return CrunchExerciseNameRaisedLegsCrunch
	case "weighted_raised_legs_crunch":
		return CrunchExerciseNameWeightedRaisedLegsCrunch
	case "reverse_crunch":
		return CrunchExerciseNameReverseCrunch
	case "weighted_reverse_crunch":
		return CrunchExerciseNameWeightedReverseCrunch
	case "reverse_crunch_on_a_bench":
		return CrunchExerciseNameReverseCrunchOnABench
	case "weighted_reverse_crunch_on_a_bench":
		return CrunchExerciseNameWeightedReverseCrunchOnABench
	case "reverse_curl_and_lift":
		return CrunchExerciseNameReverseCurlAndLift
	case "weighted_reverse_curl_and_lift":
		return CrunchExerciseNameWeightedReverseCurlAndLift
	case "rotational_lift":
		return CrunchExerciseNameRotationalLift
	case "weighted_rotational_lift":
		return CrunchExerciseNameWeightedRotationalLift
	case "seated_alternating_reverse_crunch":
		return CrunchExerciseNameSeatedAlternatingReverseCrunch
	case "weighted_seated_alternating_reverse_crunch":
		return CrunchExerciseNameWeightedSeatedAlternatingReverseCrunch
	case "seated_leg_u":
		return CrunchExerciseNameSeatedLegU
	case "weighted_seated_leg_u":
		return CrunchExerciseNameWeightedSeatedLegU
	case "side_to_side_crunch_and_weave":
		return CrunchExerciseNameSideToSideCrunchAndWeave
	case "weighted_side_to_side_crunch_and_weave":
		return CrunchExerciseNameWeightedSideToSideCrunchAndWeave
	case "single_leg_reverse_crunch":
		return CrunchExerciseNameSingleLegReverseCrunch
	case "weighted_single_leg_reverse_crunch":
		return CrunchExerciseNameWeightedSingleLegReverseCrunch
	case "skater_crunch_cross":
		return CrunchExerciseNameSkaterCrunchCross
	case "weighted_skater_crunch_cross":
		return CrunchExerciseNameWeightedSkaterCrunchCross
	case "standing_cable_crunch":
		return CrunchExerciseNameStandingCableCrunch
	case "standing_side_crunch":
		return CrunchExerciseNameStandingSideCrunch
	case "step_climb":
		return CrunchExerciseNameStepClimb
	case "weighted_step_climb":
		return CrunchExerciseNameWeightedStepClimb
	case "swiss_ball_crunch":
		return CrunchExerciseNameSwissBallCrunch
	case "swiss_ball_reverse_crunch":
		return CrunchExerciseNameSwissBallReverseCrunch
	case "weighted_swiss_ball_reverse_crunch":
		return CrunchExerciseNameWeightedSwissBallReverseCrunch
	case "swiss_ball_russian_twist":
		return CrunchExerciseNameSwissBallRussianTwist
	case "weighted_swiss_ball_russian_twist":
		return CrunchExerciseNameWeightedSwissBallRussianTwist
	case "swiss_ball_side_crunch":
		return CrunchExerciseNameSwissBallSideCrunch
	case "weighted_swiss_ball_side_crunch":
		return CrunchExerciseNameWeightedSwissBallSideCrunch
	case "thoracic_crunches_on_foam_roller":
		return CrunchExerciseNameThoracicCrunchesOnFoamRoller
	case "weighted_thoracic_crunches_on_foam_roller":
		return CrunchExerciseNameWeightedThoracicCrunchesOnFoamRoller
	case "triceps_crunch":
		return CrunchExerciseNameTricepsCrunch
	case "weighted_bicycle_crunch":
		return CrunchExerciseNameWeightedBicycleCrunch
	case "weighted_crunch":
		return CrunchExerciseNameWeightedCrunch
	case "weighted_swiss_ball_crunch":
		return CrunchExerciseNameWeightedSwissBallCrunch
	case "toes_to_bar":
		return CrunchExerciseNameToesToBar
	case "weighted_toes_to_bar":
		return CrunchExerciseNameWeightedToesToBar
	case "crunch":
		return CrunchExerciseNameCrunch
	case "straight_leg_crunch_with_ball":
		return CrunchExerciseNameStraightLegCrunchWithBall
	default:
		return CrunchExerciseNameInvalid
	}
}

// List returns all constants.
func ListCrunchExerciseName() []CrunchExerciseName {
	return []CrunchExerciseName{
		CrunchExerciseNameBicycleCrunch,
		CrunchExerciseNameCableCrunch,
		CrunchExerciseNameCircularArmCrunch,
		CrunchExerciseNameCrossedArmsCrunch,
		CrunchExerciseNameWeightedCrossedArmsCrunch,
		CrunchExerciseNameCrossLegReverseCrunch,
		CrunchExerciseNameWeightedCrossLegReverseCrunch,
		CrunchExerciseNameCrunchChop,
		CrunchExerciseNameWeightedCrunchChop,
		CrunchExerciseNameDoubleCrunch,
		CrunchExerciseNameWeightedDoubleCrunch,
		CrunchExerciseNameElbowToKneeCrunch,
		CrunchExerciseNameWeightedElbowToKneeCrunch,
		CrunchExerciseNameFlutterKicks,
		CrunchExerciseNameWeightedFlutterKicks,
		CrunchExerciseNameFoamRollerReverseCrunchOnBench,
		CrunchExerciseNameWeightedFoamRollerReverseCrunchOnBench,
		CrunchExerciseNameFoamRollerReverseCrunchWithDumbbell,
		CrunchExerciseNameFoamRollerReverseCrunchWithMedicineBall,
		CrunchExerciseNameFrogPress,
		CrunchExerciseNameHangingKneeRaiseObliqueCrunch,
		CrunchExerciseNameWeightedHangingKneeRaiseObliqueCrunch,
		CrunchExerciseNameHipCrossover,
		CrunchExerciseNameWeightedHipCrossover,
		CrunchExerciseNameHollowRock,
		CrunchExerciseNameWeightedHollowRock,
		CrunchExerciseNameInclineReverseCrunch,
		CrunchExerciseNameWeightedInclineReverseCrunch,
		CrunchExerciseNameKneelingCableCrunch,
		CrunchExerciseNameKneelingCrossCrunch,
		CrunchExerciseNameWeightedKneelingCrossCrunch,
		CrunchExerciseNameKneelingObliqueCableCrunch,
		CrunchExerciseNameKneesToElbow,
		CrunchExerciseNameLegExtensions,
		CrunchExerciseNameWeightedLegExtensions,
		CrunchExerciseNameLegLevers,
		CrunchExerciseNameMcgillCurlUp,
		CrunchExerciseNameWeightedMcgillCurlUp,
		CrunchExerciseNameModifiedPilatesRollUpWithBall,
		CrunchExerciseNameWeightedModifiedPilatesRollUpWithBall,
		CrunchExerciseNamePilatesCrunch,
		CrunchExerciseNameWeightedPilatesCrunch,
		CrunchExerciseNamePilatesRollUpWithBall,
		CrunchExerciseNameWeightedPilatesRollUpWithBall,
		CrunchExerciseNameRaisedLegsCrunch,
		CrunchExerciseNameWeightedRaisedLegsCrunch,
		CrunchExerciseNameReverseCrunch,
		CrunchExerciseNameWeightedReverseCrunch,
		CrunchExerciseNameReverseCrunchOnABench,
		CrunchExerciseNameWeightedReverseCrunchOnABench,
		CrunchExerciseNameReverseCurlAndLift,
		CrunchExerciseNameWeightedReverseCurlAndLift,
		CrunchExerciseNameRotationalLift,
		CrunchExerciseNameWeightedRotationalLift,
		CrunchExerciseNameSeatedAlternatingReverseCrunch,
		CrunchExerciseNameWeightedSeatedAlternatingReverseCrunch,
		CrunchExerciseNameSeatedLegU,
		CrunchExerciseNameWeightedSeatedLegU,
		CrunchExerciseNameSideToSideCrunchAndWeave,
		CrunchExerciseNameWeightedSideToSideCrunchAndWeave,
		CrunchExerciseNameSingleLegReverseCrunch,
		CrunchExerciseNameWeightedSingleLegReverseCrunch,
		CrunchExerciseNameSkaterCrunchCross,
		CrunchExerciseNameWeightedSkaterCrunchCross,
		CrunchExerciseNameStandingCableCrunch,
		CrunchExerciseNameStandingSideCrunch,
		CrunchExerciseNameStepClimb,
		CrunchExerciseNameWeightedStepClimb,
		CrunchExerciseNameSwissBallCrunch,
		CrunchExerciseNameSwissBallReverseCrunch,
		CrunchExerciseNameWeightedSwissBallReverseCrunch,
		CrunchExerciseNameSwissBallRussianTwist,
		CrunchExerciseNameWeightedSwissBallRussianTwist,
		CrunchExerciseNameSwissBallSideCrunch,
		CrunchExerciseNameWeightedSwissBallSideCrunch,
		CrunchExerciseNameThoracicCrunchesOnFoamRoller,
		CrunchExerciseNameWeightedThoracicCrunchesOnFoamRoller,
		CrunchExerciseNameTricepsCrunch,
		CrunchExerciseNameWeightedBicycleCrunch,
		CrunchExerciseNameWeightedCrunch,
		CrunchExerciseNameWeightedSwissBallCrunch,
		CrunchExerciseNameToesToBar,
		CrunchExerciseNameWeightedToesToBar,
		CrunchExerciseNameCrunch,
		CrunchExerciseNameStraightLegCrunchWithBall,
	}
}
