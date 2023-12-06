// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// Lap is a Lap message.
type Lap struct {
	MessageIndex                  typedef.MessageIndex
	Timestamp                     typedef.DateTime // Units: s; Lap end time.
	Event                         typedef.Event
	EventType                     typedef.EventType
	StartTime                     typedef.DateTime
	StartPositionLat              int32  // Units: semicircles;
	StartPositionLong             int32  // Units: semicircles;
	EndPositionLat                int32  // Units: semicircles;
	EndPositionLong               int32  // Units: semicircles;
	TotalElapsedTime              uint32 // Scale: 1000; Units: s; Time (includes pauses)
	TotalTimerTime                uint32 // Scale: 1000; Units: s; Timer Time (excludes pauses)
	TotalDistance                 uint32 // Scale: 100; Units: m;
	TotalCycles                   uint32 // Units: cycles;
	TotalCalories                 uint16 // Units: kcal;
	TotalFatCalories              uint16 // Units: kcal; If New Leaf
	AvgSpeed                      uint16 // Scale: 1000; Units: m/s;
	MaxSpeed                      uint16 // Scale: 1000; Units: m/s;
	AvgHeartRate                  uint8  // Units: bpm;
	MaxHeartRate                  uint8  // Units: bpm;
	AvgCadence                    uint8  // Units: rpm; total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time
	MaxCadence                    uint8  // Units: rpm;
	AvgPower                      uint16 // Units: watts; total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time
	MaxPower                      uint16 // Units: watts;
	TotalAscent                   uint16 // Units: m;
	TotalDescent                  uint16 // Units: m;
	Intensity                     typedef.Intensity
	LapTrigger                    typedef.LapTrigger
	Sport                         typedef.Sport
	EventGroup                    uint8
	NumLengths                    uint16 // Units: lengths; # of lengths of swim pool
	NormalizedPower               uint16 // Units: watts;
	LeftRightBalance              typedef.LeftRightBalance100
	FirstLengthIndex              uint16
	AvgStrokeDistance             uint16 // Scale: 100; Units: m;
	SwimStroke                    typedef.SwimStroke
	SubSport                      typedef.SubSport
	NumActiveLengths              uint16   // Units: lengths; # of active lengths of swim pool
	TotalWork                     uint32   // Units: J;
	AvgAltitude                   uint16   // Scale: 5; Offset: 500; Units: m;
	MaxAltitude                   uint16   // Scale: 5; Offset: 500; Units: m;
	GpsAccuracy                   uint8    // Units: m;
	AvgGrade                      int16    // Scale: 100; Units: %;
	AvgPosGrade                   int16    // Scale: 100; Units: %;
	AvgNegGrade                   int16    // Scale: 100; Units: %;
	MaxPosGrade                   int16    // Scale: 100; Units: %;
	MaxNegGrade                   int16    // Scale: 100; Units: %;
	AvgTemperature                int8     // Units: C;
	MaxTemperature                int8     // Units: C;
	TotalMovingTime               uint32   // Scale: 1000; Units: s;
	AvgPosVerticalSpeed           int16    // Scale: 1000; Units: m/s;
	AvgNegVerticalSpeed           int16    // Scale: 1000; Units: m/s;
	MaxPosVerticalSpeed           int16    // Scale: 1000; Units: m/s;
	MaxNegVerticalSpeed           int16    // Scale: 1000; Units: m/s;
	TimeInHrZone                  []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInSpeedZone               []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInCadenceZone             []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInPowerZone               []uint32 // Scale: 1000; Array: [N]; Units: s;
	RepetitionNum                 uint16
	MinAltitude                   uint16 // Scale: 5; Offset: 500; Units: m;
	MinHeartRate                  uint8  // Units: bpm;
	WktStepIndex                  typedef.MessageIndex
	OpponentScore                 uint16
	StrokeCount                   []uint16 // Array: [N]; Units: counts; stroke_type enum used as the index
	ZoneCount                     []uint16 // Array: [N]; Units: counts; zone number used as the index
	AvgVerticalOscillation        uint16   // Scale: 10; Units: mm;
	AvgStanceTimePercent          uint16   // Scale: 100; Units: percent;
	AvgStanceTime                 uint16   // Scale: 10; Units: ms;
	AvgFractionalCadence          uint8    // Scale: 128; Units: rpm; fractional part of the avg_cadence
	MaxFractionalCadence          uint8    // Scale: 128; Units: rpm; fractional part of the max_cadence
	TotalFractionalCycles         uint8    // Scale: 128; Units: cycles; fractional part of the total_cycles
	PlayerScore                   uint16
	AvgTotalHemoglobinConc        []uint16 // Scale: 100; Array: [N]; Units: g/dL; Avg saturated and unsaturated hemoglobin
	MinTotalHemoglobinConc        []uint16 // Scale: 100; Array: [N]; Units: g/dL; Min saturated and unsaturated hemoglobin
	MaxTotalHemoglobinConc        []uint16 // Scale: 100; Array: [N]; Units: g/dL; Max saturated and unsaturated hemoglobin
	AvgSaturatedHemoglobinPercent []uint16 // Scale: 10; Array: [N]; Units: %; Avg percentage of hemoglobin saturated with oxygen
	MinSaturatedHemoglobinPercent []uint16 // Scale: 10; Array: [N]; Units: %; Min percentage of hemoglobin saturated with oxygen
	MaxSaturatedHemoglobinPercent []uint16 // Scale: 10; Array: [N]; Units: %; Max percentage of hemoglobin saturated with oxygen
	AvgLeftTorqueEffectiveness    uint8    // Scale: 2; Units: percent;
	AvgRightTorqueEffectiveness   uint8    // Scale: 2; Units: percent;
	AvgLeftPedalSmoothness        uint8    // Scale: 2; Units: percent;
	AvgRightPedalSmoothness       uint8    // Scale: 2; Units: percent;
	AvgCombinedPedalSmoothness    uint8    // Scale: 2; Units: percent;
	TimeStanding                  uint32   // Scale: 1000; Units: s; Total time spent in the standing position
	StandCount                    uint16   // Number of transitions to the standing state
	AvgLeftPco                    int8     // Units: mm; Average left platform center offset
	AvgRightPco                   int8     // Units: mm; Average right platform center offset
	AvgLeftPowerPhase             []uint8  // Scale: 0.7111111; Array: [N]; Units: degrees; Average left power phase angles. Data value indexes defined by power_phase_type.
	AvgLeftPowerPhasePeak         []uint8  // Scale: 0.7111111; Array: [N]; Units: degrees; Average left power phase peak angles. Data value indexes defined by power_phase_type.
	AvgRightPowerPhase            []uint8  // Scale: 0.7111111; Array: [N]; Units: degrees; Average right power phase angles. Data value indexes defined by power_phase_type.
	AvgRightPowerPhasePeak        []uint8  // Scale: 0.7111111; Array: [N]; Units: degrees; Average right power phase peak angles. Data value indexes defined by power_phase_type.
	AvgPowerPosition              []uint16 // Array: [N]; Units: watts; Average power by position. Data value indexes defined by rider_position_type.
	MaxPowerPosition              []uint16 // Array: [N]; Units: watts; Maximum power by position. Data value indexes defined by rider_position_type.
	AvgCadencePosition            []uint8  // Array: [N]; Units: rpm; Average cadence by position. Data value indexes defined by rider_position_type.
	MaxCadencePosition            []uint8  // Array: [N]; Units: rpm; Maximum cadence by position. Data value indexes defined by rider_position_type.
	EnhancedAvgSpeed              uint32   // Scale: 1000; Units: m/s;
	EnhancedMaxSpeed              uint32   // Scale: 1000; Units: m/s;
	EnhancedAvgAltitude           uint32   // Scale: 5; Offset: 500; Units: m;
	EnhancedMinAltitude           uint32   // Scale: 5; Offset: 500; Units: m;
	EnhancedMaxAltitude           uint32   // Scale: 5; Offset: 500; Units: m;
	AvgLevMotorPower              uint16   // Units: watts; lev average motor power during lap
	MaxLevMotorPower              uint16   // Units: watts; lev maximum motor power during lap
	LevBatteryConsumption         uint8    // Scale: 2; Units: percent; lev battery consumption during lap
	AvgVerticalRatio              uint16   // Scale: 100; Units: percent;
	AvgStanceTimeBalance          uint16   // Scale: 100; Units: percent;
	AvgStepLength                 uint16   // Scale: 10; Units: mm;
	AvgVam                        uint16   // Scale: 1000; Units: m/s;
	AvgDepth                      uint32   // Scale: 1000; Units: m; 0 if above water
	MaxDepth                      uint32   // Scale: 1000; Units: m; 0 if above water
	MinTemperature                int8     // Units: C;
	EnhancedAvgRespirationRate    uint16   // Scale: 100; Units: Breaths/min;
	EnhancedMaxRespirationRate    uint16   // Scale: 100; Units: Breaths/min;
	AvgRespirationRate            uint8
	MaxRespirationRate            uint8
	TotalGrit                     float32 // Units: kGrit; The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.
	TotalFlow                     float32 // Units: Flow; The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.
	JumpCount                     uint16
	AvgGrit                       float32 // Units: kGrit; The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.
	AvgFlow                       float32 // Units: Flow; The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.
	TotalFractionalAscent         uint8   // Scale: 100; Units: m; fractional part of total_ascent
	TotalFractionalDescent        uint8   // Scale: 100; Units: m; fractional part of total_descent
	AvgCoreTemperature            uint16  // Scale: 100; Units: C;
	MinCoreTemperature            uint16  // Scale: 100; Units: C;
	MaxCoreTemperature            uint16  // Scale: 100; Units: C;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewLap creates new Lap struct based on given mesg. If mesg is nil or mesg.Num is not equal to Lap mesg number, it will return nil.
func NewLap(mesg proto.Message) *Lap {
	if mesg.Num != typedef.MesgNumLap {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid,                        /* MessageIndex */
		253: basetype.Uint32Invalid,                        /* Timestamp */
		0:   basetype.EnumInvalid,                          /* Event */
		1:   basetype.EnumInvalid,                          /* EventType */
		2:   basetype.Uint32Invalid,                        /* StartTime */
		3:   basetype.Sint32Invalid,                        /* StartPositionLat */
		4:   basetype.Sint32Invalid,                        /* StartPositionLong */
		5:   basetype.Sint32Invalid,                        /* EndPositionLat */
		6:   basetype.Sint32Invalid,                        /* EndPositionLong */
		7:   basetype.Uint32Invalid,                        /* TotalElapsedTime */
		8:   basetype.Uint32Invalid,                        /* TotalTimerTime */
		9:   basetype.Uint32Invalid,                        /* TotalDistance */
		10:  basetype.Uint32Invalid,                        /* TotalCycles */
		11:  basetype.Uint16Invalid,                        /* TotalCalories */
		12:  basetype.Uint16Invalid,                        /* TotalFatCalories */
		13:  basetype.Uint16Invalid,                        /* AvgSpeed */
		14:  basetype.Uint16Invalid,                        /* MaxSpeed */
		15:  basetype.Uint8Invalid,                         /* AvgHeartRate */
		16:  basetype.Uint8Invalid,                         /* MaxHeartRate */
		17:  basetype.Uint8Invalid,                         /* AvgCadence */
		18:  basetype.Uint8Invalid,                         /* MaxCadence */
		19:  basetype.Uint16Invalid,                        /* AvgPower */
		20:  basetype.Uint16Invalid,                        /* MaxPower */
		21:  basetype.Uint16Invalid,                        /* TotalAscent */
		22:  basetype.Uint16Invalid,                        /* TotalDescent */
		23:  basetype.EnumInvalid,                          /* Intensity */
		24:  basetype.EnumInvalid,                          /* LapTrigger */
		25:  basetype.EnumInvalid,                          /* Sport */
		26:  basetype.Uint8Invalid,                         /* EventGroup */
		32:  basetype.Uint16Invalid,                        /* NumLengths */
		33:  basetype.Uint16Invalid,                        /* NormalizedPower */
		34:  basetype.Uint16Invalid,                        /* LeftRightBalance */
		35:  basetype.Uint16Invalid,                        /* FirstLengthIndex */
		37:  basetype.Uint16Invalid,                        /* AvgStrokeDistance */
		38:  basetype.EnumInvalid,                          /* SwimStroke */
		39:  basetype.EnumInvalid,                          /* SubSport */
		40:  basetype.Uint16Invalid,                        /* NumActiveLengths */
		41:  basetype.Uint32Invalid,                        /* TotalWork */
		42:  basetype.Uint16Invalid,                        /* AvgAltitude */
		43:  basetype.Uint16Invalid,                        /* MaxAltitude */
		44:  basetype.Uint8Invalid,                         /* GpsAccuracy */
		45:  basetype.Sint16Invalid,                        /* AvgGrade */
		46:  basetype.Sint16Invalid,                        /* AvgPosGrade */
		47:  basetype.Sint16Invalid,                        /* AvgNegGrade */
		48:  basetype.Sint16Invalid,                        /* MaxPosGrade */
		49:  basetype.Sint16Invalid,                        /* MaxNegGrade */
		50:  basetype.Sint8Invalid,                         /* AvgTemperature */
		51:  basetype.Sint8Invalid,                         /* MaxTemperature */
		52:  basetype.Uint32Invalid,                        /* TotalMovingTime */
		53:  basetype.Sint16Invalid,                        /* AvgPosVerticalSpeed */
		54:  basetype.Sint16Invalid,                        /* AvgNegVerticalSpeed */
		55:  basetype.Sint16Invalid,                        /* MaxPosVerticalSpeed */
		56:  basetype.Sint16Invalid,                        /* MaxNegVerticalSpeed */
		57:  nil,                                           /* TimeInHrZone */
		58:  nil,                                           /* TimeInSpeedZone */
		59:  nil,                                           /* TimeInCadenceZone */
		60:  nil,                                           /* TimeInPowerZone */
		61:  basetype.Uint16Invalid,                        /* RepetitionNum */
		62:  basetype.Uint16Invalid,                        /* MinAltitude */
		63:  basetype.Uint8Invalid,                         /* MinHeartRate */
		71:  basetype.Uint16Invalid,                        /* WktStepIndex */
		74:  basetype.Uint16Invalid,                        /* OpponentScore */
		75:  nil,                                           /* StrokeCount */
		76:  nil,                                           /* ZoneCount */
		77:  basetype.Uint16Invalid,                        /* AvgVerticalOscillation */
		78:  basetype.Uint16Invalid,                        /* AvgStanceTimePercent */
		79:  basetype.Uint16Invalid,                        /* AvgStanceTime */
		80:  basetype.Uint8Invalid,                         /* AvgFractionalCadence */
		81:  basetype.Uint8Invalid,                         /* MaxFractionalCadence */
		82:  basetype.Uint8Invalid,                         /* TotalFractionalCycles */
		83:  basetype.Uint16Invalid,                        /* PlayerScore */
		84:  nil,                                           /* AvgTotalHemoglobinConc */
		85:  nil,                                           /* MinTotalHemoglobinConc */
		86:  nil,                                           /* MaxTotalHemoglobinConc */
		87:  nil,                                           /* AvgSaturatedHemoglobinPercent */
		88:  nil,                                           /* MinSaturatedHemoglobinPercent */
		89:  nil,                                           /* MaxSaturatedHemoglobinPercent */
		91:  basetype.Uint8Invalid,                         /* AvgLeftTorqueEffectiveness */
		92:  basetype.Uint8Invalid,                         /* AvgRightTorqueEffectiveness */
		93:  basetype.Uint8Invalid,                         /* AvgLeftPedalSmoothness */
		94:  basetype.Uint8Invalid,                         /* AvgRightPedalSmoothness */
		95:  basetype.Uint8Invalid,                         /* AvgCombinedPedalSmoothness */
		98:  basetype.Uint32Invalid,                        /* TimeStanding */
		99:  basetype.Uint16Invalid,                        /* StandCount */
		100: basetype.Sint8Invalid,                         /* AvgLeftPco */
		101: basetype.Sint8Invalid,                         /* AvgRightPco */
		102: nil,                                           /* AvgLeftPowerPhase */
		103: nil,                                           /* AvgLeftPowerPhasePeak */
		104: nil,                                           /* AvgRightPowerPhase */
		105: nil,                                           /* AvgRightPowerPhasePeak */
		106: nil,                                           /* AvgPowerPosition */
		107: nil,                                           /* MaxPowerPosition */
		108: nil,                                           /* AvgCadencePosition */
		109: nil,                                           /* MaxCadencePosition */
		110: basetype.Uint32Invalid,                        /* EnhancedAvgSpeed */
		111: basetype.Uint32Invalid,                        /* EnhancedMaxSpeed */
		112: basetype.Uint32Invalid,                        /* EnhancedAvgAltitude */
		113: basetype.Uint32Invalid,                        /* EnhancedMinAltitude */
		114: basetype.Uint32Invalid,                        /* EnhancedMaxAltitude */
		115: basetype.Uint16Invalid,                        /* AvgLevMotorPower */
		116: basetype.Uint16Invalid,                        /* MaxLevMotorPower */
		117: basetype.Uint8Invalid,                         /* LevBatteryConsumption */
		118: basetype.Uint16Invalid,                        /* AvgVerticalRatio */
		119: basetype.Uint16Invalid,                        /* AvgStanceTimeBalance */
		120: basetype.Uint16Invalid,                        /* AvgStepLength */
		121: basetype.Uint16Invalid,                        /* AvgVam */
		122: basetype.Uint32Invalid,                        /* AvgDepth */
		123: basetype.Uint32Invalid,                        /* MaxDepth */
		124: basetype.Sint8Invalid,                         /* MinTemperature */
		136: basetype.Uint16Invalid,                        /* EnhancedAvgRespirationRate */
		137: basetype.Uint16Invalid,                        /* EnhancedMaxRespirationRate */
		147: basetype.Uint8Invalid,                         /* AvgRespirationRate */
		148: basetype.Uint8Invalid,                         /* MaxRespirationRate */
		149: math.Float32frombits(basetype.Float32Invalid), /* TotalGrit */
		150: math.Float32frombits(basetype.Float32Invalid), /* TotalFlow */
		151: basetype.Uint16Invalid,                        /* JumpCount */
		153: math.Float32frombits(basetype.Float32Invalid), /* AvgGrit */
		154: math.Float32frombits(basetype.Float32Invalid), /* AvgFlow */
		156: basetype.Uint8Invalid,                         /* TotalFractionalAscent */
		157: basetype.Uint8Invalid,                         /* TotalFractionalDescent */
		158: basetype.Uint16Invalid,                        /* AvgCoreTemperature */
		159: basetype.Uint16Invalid,                        /* MinCoreTemperature */
		160: basetype.Uint16Invalid,                        /* MaxCoreTemperature */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Lap{
		MessageIndex:                  typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Timestamp:                     typeconv.ToUint32[typedef.DateTime](vals[253]),
		Event:                         typeconv.ToEnum[typedef.Event](vals[0]),
		EventType:                     typeconv.ToEnum[typedef.EventType](vals[1]),
		StartTime:                     typeconv.ToUint32[typedef.DateTime](vals[2]),
		StartPositionLat:              typeconv.ToSint32[int32](vals[3]),
		StartPositionLong:             typeconv.ToSint32[int32](vals[4]),
		EndPositionLat:                typeconv.ToSint32[int32](vals[5]),
		EndPositionLong:               typeconv.ToSint32[int32](vals[6]),
		TotalElapsedTime:              typeconv.ToUint32[uint32](vals[7]),
		TotalTimerTime:                typeconv.ToUint32[uint32](vals[8]),
		TotalDistance:                 typeconv.ToUint32[uint32](vals[9]),
		TotalCycles:                   typeconv.ToUint32[uint32](vals[10]),
		TotalCalories:                 typeconv.ToUint16[uint16](vals[11]),
		TotalFatCalories:              typeconv.ToUint16[uint16](vals[12]),
		AvgSpeed:                      typeconv.ToUint16[uint16](vals[13]),
		MaxSpeed:                      typeconv.ToUint16[uint16](vals[14]),
		AvgHeartRate:                  typeconv.ToUint8[uint8](vals[15]),
		MaxHeartRate:                  typeconv.ToUint8[uint8](vals[16]),
		AvgCadence:                    typeconv.ToUint8[uint8](vals[17]),
		MaxCadence:                    typeconv.ToUint8[uint8](vals[18]),
		AvgPower:                      typeconv.ToUint16[uint16](vals[19]),
		MaxPower:                      typeconv.ToUint16[uint16](vals[20]),
		TotalAscent:                   typeconv.ToUint16[uint16](vals[21]),
		TotalDescent:                  typeconv.ToUint16[uint16](vals[22]),
		Intensity:                     typeconv.ToEnum[typedef.Intensity](vals[23]),
		LapTrigger:                    typeconv.ToEnum[typedef.LapTrigger](vals[24]),
		Sport:                         typeconv.ToEnum[typedef.Sport](vals[25]),
		EventGroup:                    typeconv.ToUint8[uint8](vals[26]),
		NumLengths:                    typeconv.ToUint16[uint16](vals[32]),
		NormalizedPower:               typeconv.ToUint16[uint16](vals[33]),
		LeftRightBalance:              typeconv.ToUint16[typedef.LeftRightBalance100](vals[34]),
		FirstLengthIndex:              typeconv.ToUint16[uint16](vals[35]),
		AvgStrokeDistance:             typeconv.ToUint16[uint16](vals[37]),
		SwimStroke:                    typeconv.ToEnum[typedef.SwimStroke](vals[38]),
		SubSport:                      typeconv.ToEnum[typedef.SubSport](vals[39]),
		NumActiveLengths:              typeconv.ToUint16[uint16](vals[40]),
		TotalWork:                     typeconv.ToUint32[uint32](vals[41]),
		AvgAltitude:                   typeconv.ToUint16[uint16](vals[42]),
		MaxAltitude:                   typeconv.ToUint16[uint16](vals[43]),
		GpsAccuracy:                   typeconv.ToUint8[uint8](vals[44]),
		AvgGrade:                      typeconv.ToSint16[int16](vals[45]),
		AvgPosGrade:                   typeconv.ToSint16[int16](vals[46]),
		AvgNegGrade:                   typeconv.ToSint16[int16](vals[47]),
		MaxPosGrade:                   typeconv.ToSint16[int16](vals[48]),
		MaxNegGrade:                   typeconv.ToSint16[int16](vals[49]),
		AvgTemperature:                typeconv.ToSint8[int8](vals[50]),
		MaxTemperature:                typeconv.ToSint8[int8](vals[51]),
		TotalMovingTime:               typeconv.ToUint32[uint32](vals[52]),
		AvgPosVerticalSpeed:           typeconv.ToSint16[int16](vals[53]),
		AvgNegVerticalSpeed:           typeconv.ToSint16[int16](vals[54]),
		MaxPosVerticalSpeed:           typeconv.ToSint16[int16](vals[55]),
		MaxNegVerticalSpeed:           typeconv.ToSint16[int16](vals[56]),
		TimeInHrZone:                  typeconv.ToSliceUint32[uint32](vals[57]),
		TimeInSpeedZone:               typeconv.ToSliceUint32[uint32](vals[58]),
		TimeInCadenceZone:             typeconv.ToSliceUint32[uint32](vals[59]),
		TimeInPowerZone:               typeconv.ToSliceUint32[uint32](vals[60]),
		RepetitionNum:                 typeconv.ToUint16[uint16](vals[61]),
		MinAltitude:                   typeconv.ToUint16[uint16](vals[62]),
		MinHeartRate:                  typeconv.ToUint8[uint8](vals[63]),
		WktStepIndex:                  typeconv.ToUint16[typedef.MessageIndex](vals[71]),
		OpponentScore:                 typeconv.ToUint16[uint16](vals[74]),
		StrokeCount:                   typeconv.ToSliceUint16[uint16](vals[75]),
		ZoneCount:                     typeconv.ToSliceUint16[uint16](vals[76]),
		AvgVerticalOscillation:        typeconv.ToUint16[uint16](vals[77]),
		AvgStanceTimePercent:          typeconv.ToUint16[uint16](vals[78]),
		AvgStanceTime:                 typeconv.ToUint16[uint16](vals[79]),
		AvgFractionalCadence:          typeconv.ToUint8[uint8](vals[80]),
		MaxFractionalCadence:          typeconv.ToUint8[uint8](vals[81]),
		TotalFractionalCycles:         typeconv.ToUint8[uint8](vals[82]),
		PlayerScore:                   typeconv.ToUint16[uint16](vals[83]),
		AvgTotalHemoglobinConc:        typeconv.ToSliceUint16[uint16](vals[84]),
		MinTotalHemoglobinConc:        typeconv.ToSliceUint16[uint16](vals[85]),
		MaxTotalHemoglobinConc:        typeconv.ToSliceUint16[uint16](vals[86]),
		AvgSaturatedHemoglobinPercent: typeconv.ToSliceUint16[uint16](vals[87]),
		MinSaturatedHemoglobinPercent: typeconv.ToSliceUint16[uint16](vals[88]),
		MaxSaturatedHemoglobinPercent: typeconv.ToSliceUint16[uint16](vals[89]),
		AvgLeftTorqueEffectiveness:    typeconv.ToUint8[uint8](vals[91]),
		AvgRightTorqueEffectiveness:   typeconv.ToUint8[uint8](vals[92]),
		AvgLeftPedalSmoothness:        typeconv.ToUint8[uint8](vals[93]),
		AvgRightPedalSmoothness:       typeconv.ToUint8[uint8](vals[94]),
		AvgCombinedPedalSmoothness:    typeconv.ToUint8[uint8](vals[95]),
		TimeStanding:                  typeconv.ToUint32[uint32](vals[98]),
		StandCount:                    typeconv.ToUint16[uint16](vals[99]),
		AvgLeftPco:                    typeconv.ToSint8[int8](vals[100]),
		AvgRightPco:                   typeconv.ToSint8[int8](vals[101]),
		AvgLeftPowerPhase:             typeconv.ToSliceUint8[uint8](vals[102]),
		AvgLeftPowerPhasePeak:         typeconv.ToSliceUint8[uint8](vals[103]),
		AvgRightPowerPhase:            typeconv.ToSliceUint8[uint8](vals[104]),
		AvgRightPowerPhasePeak:        typeconv.ToSliceUint8[uint8](vals[105]),
		AvgPowerPosition:              typeconv.ToSliceUint16[uint16](vals[106]),
		MaxPowerPosition:              typeconv.ToSliceUint16[uint16](vals[107]),
		AvgCadencePosition:            typeconv.ToSliceUint8[uint8](vals[108]),
		MaxCadencePosition:            typeconv.ToSliceUint8[uint8](vals[109]),
		EnhancedAvgSpeed:              typeconv.ToUint32[uint32](vals[110]),
		EnhancedMaxSpeed:              typeconv.ToUint32[uint32](vals[111]),
		EnhancedAvgAltitude:           typeconv.ToUint32[uint32](vals[112]),
		EnhancedMinAltitude:           typeconv.ToUint32[uint32](vals[113]),
		EnhancedMaxAltitude:           typeconv.ToUint32[uint32](vals[114]),
		AvgLevMotorPower:              typeconv.ToUint16[uint16](vals[115]),
		MaxLevMotorPower:              typeconv.ToUint16[uint16](vals[116]),
		LevBatteryConsumption:         typeconv.ToUint8[uint8](vals[117]),
		AvgVerticalRatio:              typeconv.ToUint16[uint16](vals[118]),
		AvgStanceTimeBalance:          typeconv.ToUint16[uint16](vals[119]),
		AvgStepLength:                 typeconv.ToUint16[uint16](vals[120]),
		AvgVam:                        typeconv.ToUint16[uint16](vals[121]),
		AvgDepth:                      typeconv.ToUint32[uint32](vals[122]),
		MaxDepth:                      typeconv.ToUint32[uint32](vals[123]),
		MinTemperature:                typeconv.ToSint8[int8](vals[124]),
		EnhancedAvgRespirationRate:    typeconv.ToUint16[uint16](vals[136]),
		EnhancedMaxRespirationRate:    typeconv.ToUint16[uint16](vals[137]),
		AvgRespirationRate:            typeconv.ToUint8[uint8](vals[147]),
		MaxRespirationRate:            typeconv.ToUint8[uint8](vals[148]),
		TotalGrit:                     typeconv.ToFloat32[float32](vals[149]),
		TotalFlow:                     typeconv.ToFloat32[float32](vals[150]),
		JumpCount:                     typeconv.ToUint16[uint16](vals[151]),
		AvgGrit:                       typeconv.ToFloat32[float32](vals[153]),
		AvgFlow:                       typeconv.ToFloat32[float32](vals[154]),
		TotalFractionalAscent:         typeconv.ToUint8[uint8](vals[156]),
		TotalFractionalDescent:        typeconv.ToUint8[uint8](vals[157]),
		AvgCoreTemperature:            typeconv.ToUint16[uint16](vals[158]),
		MinCoreTemperature:            typeconv.ToUint16[uint16](vals[159]),
		MaxCoreTemperature:            typeconv.ToUint16[uint16](vals[160]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Lap mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumLap)
func (m Lap) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumLap {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		253: m.Timestamp,
		0:   m.Event,
		1:   m.EventType,
		2:   m.StartTime,
		3:   m.StartPositionLat,
		4:   m.StartPositionLong,
		5:   m.EndPositionLat,
		6:   m.EndPositionLong,
		7:   m.TotalElapsedTime,
		8:   m.TotalTimerTime,
		9:   m.TotalDistance,
		10:  m.TotalCycles,
		11:  m.TotalCalories,
		12:  m.TotalFatCalories,
		13:  m.AvgSpeed,
		14:  m.MaxSpeed,
		15:  m.AvgHeartRate,
		16:  m.MaxHeartRate,
		17:  m.AvgCadence,
		18:  m.MaxCadence,
		19:  m.AvgPower,
		20:  m.MaxPower,
		21:  m.TotalAscent,
		22:  m.TotalDescent,
		23:  m.Intensity,
		24:  m.LapTrigger,
		25:  m.Sport,
		26:  m.EventGroup,
		32:  m.NumLengths,
		33:  m.NormalizedPower,
		34:  m.LeftRightBalance,
		35:  m.FirstLengthIndex,
		37:  m.AvgStrokeDistance,
		38:  m.SwimStroke,
		39:  m.SubSport,
		40:  m.NumActiveLengths,
		41:  m.TotalWork,
		42:  m.AvgAltitude,
		43:  m.MaxAltitude,
		44:  m.GpsAccuracy,
		45:  m.AvgGrade,
		46:  m.AvgPosGrade,
		47:  m.AvgNegGrade,
		48:  m.MaxPosGrade,
		49:  m.MaxNegGrade,
		50:  m.AvgTemperature,
		51:  m.MaxTemperature,
		52:  m.TotalMovingTime,
		53:  m.AvgPosVerticalSpeed,
		54:  m.AvgNegVerticalSpeed,
		55:  m.MaxPosVerticalSpeed,
		56:  m.MaxNegVerticalSpeed,
		57:  m.TimeInHrZone,
		58:  m.TimeInSpeedZone,
		59:  m.TimeInCadenceZone,
		60:  m.TimeInPowerZone,
		61:  m.RepetitionNum,
		62:  m.MinAltitude,
		63:  m.MinHeartRate,
		71:  m.WktStepIndex,
		74:  m.OpponentScore,
		75:  m.StrokeCount,
		76:  m.ZoneCount,
		77:  m.AvgVerticalOscillation,
		78:  m.AvgStanceTimePercent,
		79:  m.AvgStanceTime,
		80:  m.AvgFractionalCadence,
		81:  m.MaxFractionalCadence,
		82:  m.TotalFractionalCycles,
		83:  m.PlayerScore,
		84:  m.AvgTotalHemoglobinConc,
		85:  m.MinTotalHemoglobinConc,
		86:  m.MaxTotalHemoglobinConc,
		87:  m.AvgSaturatedHemoglobinPercent,
		88:  m.MinSaturatedHemoglobinPercent,
		89:  m.MaxSaturatedHemoglobinPercent,
		91:  m.AvgLeftTorqueEffectiveness,
		92:  m.AvgRightTorqueEffectiveness,
		93:  m.AvgLeftPedalSmoothness,
		94:  m.AvgRightPedalSmoothness,
		95:  m.AvgCombinedPedalSmoothness,
		98:  m.TimeStanding,
		99:  m.StandCount,
		100: m.AvgLeftPco,
		101: m.AvgRightPco,
		102: m.AvgLeftPowerPhase,
		103: m.AvgLeftPowerPhasePeak,
		104: m.AvgRightPowerPhase,
		105: m.AvgRightPowerPhasePeak,
		106: m.AvgPowerPosition,
		107: m.MaxPowerPosition,
		108: m.AvgCadencePosition,
		109: m.MaxCadencePosition,
		110: m.EnhancedAvgSpeed,
		111: m.EnhancedMaxSpeed,
		112: m.EnhancedAvgAltitude,
		113: m.EnhancedMinAltitude,
		114: m.EnhancedMaxAltitude,
		115: m.AvgLevMotorPower,
		116: m.MaxLevMotorPower,
		117: m.LevBatteryConsumption,
		118: m.AvgVerticalRatio,
		119: m.AvgStanceTimeBalance,
		120: m.AvgStepLength,
		121: m.AvgVam,
		122: m.AvgDepth,
		123: m.MaxDepth,
		124: m.MinTemperature,
		136: m.EnhancedAvgRespirationRate,
		137: m.EnhancedMaxRespirationRate,
		147: m.AvgRespirationRate,
		148: m.MaxRespirationRate,
		149: m.TotalGrit,
		150: m.TotalFlow,
		151: m.JumpCount,
		153: m.AvgGrit,
		154: m.AvgFlow,
		156: m.TotalFractionalAscent,
		157: m.TotalFractionalDescent,
		158: m.AvgCoreTemperature,
		159: m.MinCoreTemperature,
		160: m.MaxCoreTemperature,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
