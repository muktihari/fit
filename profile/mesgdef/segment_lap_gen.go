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

// SegmentLap is a SegmentLap message.
type SegmentLap struct {
	MessageIndex                typedef.MessageIndex
	Timestamp                   typedef.DateTime // Units: s; Lap end time.
	Event                       typedef.Event
	EventType                   typedef.EventType
	StartTime                   typedef.DateTime
	StartPositionLat            int32  // Units: semicircles;
	StartPositionLong           int32  // Units: semicircles;
	EndPositionLat              int32  // Units: semicircles;
	EndPositionLong             int32  // Units: semicircles;
	TotalElapsedTime            uint32 // Scale: 1000; Units: s; Time (includes pauses)
	TotalTimerTime              uint32 // Scale: 1000; Units: s; Timer Time (excludes pauses)
	TotalDistance               uint32 // Scale: 100; Units: m;
	TotalCycles                 uint32 // Units: cycles;
	TotalCalories               uint16 // Units: kcal;
	TotalFatCalories            uint16 // Units: kcal; If New Leaf
	AvgSpeed                    uint16 // Scale: 1000; Units: m/s;
	MaxSpeed                    uint16 // Scale: 1000; Units: m/s;
	AvgHeartRate                uint8  // Units: bpm;
	MaxHeartRate                uint8  // Units: bpm;
	AvgCadence                  uint8  // Units: rpm; total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time
	MaxCadence                  uint8  // Units: rpm;
	AvgPower                    uint16 // Units: watts; total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time
	MaxPower                    uint16 // Units: watts;
	TotalAscent                 uint16 // Units: m;
	TotalDescent                uint16 // Units: m;
	Sport                       typedef.Sport
	EventGroup                  uint8
	NecLat                      int32 // Units: semicircles; North east corner latitude.
	NecLong                     int32 // Units: semicircles; North east corner longitude.
	SwcLat                      int32 // Units: semicircles; South west corner latitude.
	SwcLong                     int32 // Units: semicircles; South west corner latitude.
	Name                        string
	NormalizedPower             uint16 // Units: watts;
	LeftRightBalance            typedef.LeftRightBalance100
	SubSport                    typedef.SubSport
	TotalWork                   uint32   // Units: J;
	AvgAltitude                 uint16   // Scale: 5; Offset: 500; Units: m;
	MaxAltitude                 uint16   // Scale: 5; Offset: 500; Units: m;
	GpsAccuracy                 uint8    // Units: m;
	AvgGrade                    int16    // Scale: 100; Units: %;
	AvgPosGrade                 int16    // Scale: 100; Units: %;
	AvgNegGrade                 int16    // Scale: 100; Units: %;
	MaxPosGrade                 int16    // Scale: 100; Units: %;
	MaxNegGrade                 int16    // Scale: 100; Units: %;
	AvgTemperature              int8     // Units: C;
	MaxTemperature              int8     // Units: C;
	TotalMovingTime             uint32   // Scale: 1000; Units: s;
	AvgPosVerticalSpeed         int16    // Scale: 1000; Units: m/s;
	AvgNegVerticalSpeed         int16    // Scale: 1000; Units: m/s;
	MaxPosVerticalSpeed         int16    // Scale: 1000; Units: m/s;
	MaxNegVerticalSpeed         int16    // Scale: 1000; Units: m/s;
	TimeInHrZone                []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInSpeedZone             []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInCadenceZone           []uint32 // Scale: 1000; Array: [N]; Units: s;
	TimeInPowerZone             []uint32 // Scale: 1000; Array: [N]; Units: s;
	RepetitionNum               uint16
	MinAltitude                 uint16 // Scale: 5; Offset: 500; Units: m;
	MinHeartRate                uint8  // Units: bpm;
	ActiveTime                  uint32 // Scale: 1000; Units: s;
	WktStepIndex                typedef.MessageIndex
	SportEvent                  typedef.SportEvent
	AvgLeftTorqueEffectiveness  uint8 // Scale: 2; Units: percent;
	AvgRightTorqueEffectiveness uint8 // Scale: 2; Units: percent;
	AvgLeftPedalSmoothness      uint8 // Scale: 2; Units: percent;
	AvgRightPedalSmoothness     uint8 // Scale: 2; Units: percent;
	AvgCombinedPedalSmoothness  uint8 // Scale: 2; Units: percent;
	Status                      typedef.SegmentLapStatus
	Uuid                        string
	AvgFractionalCadence        uint8 // Scale: 128; Units: rpm; fractional part of the avg_cadence
	MaxFractionalCadence        uint8 // Scale: 128; Units: rpm; fractional part of the max_cadence
	TotalFractionalCycles       uint8 // Scale: 128; Units: cycles; fractional part of the total_cycles
	FrontGearShiftCount         uint16
	RearGearShiftCount          uint16
	TimeStanding                uint32               // Scale: 1000; Units: s; Total time spent in the standing position
	StandCount                  uint16               // Number of transitions to the standing state
	AvgLeftPco                  int8                 // Units: mm; Average left platform center offset
	AvgRightPco                 int8                 // Units: mm; Average right platform center offset
	AvgLeftPowerPhase           []uint8              // Scale: 0.7111111; Array: [N]; Units: degrees; Average left power phase angles. Data value indexes defined by power_phase_type.
	AvgLeftPowerPhasePeak       []uint8              // Scale: 0.7111111; Array: [N]; Units: degrees; Average left power phase peak angles. Data value indexes defined by power_phase_type.
	AvgRightPowerPhase          []uint8              // Scale: 0.7111111; Array: [N]; Units: degrees; Average right power phase angles. Data value indexes defined by power_phase_type.
	AvgRightPowerPhasePeak      []uint8              // Scale: 0.7111111; Array: [N]; Units: degrees; Average right power phase peak angles. Data value indexes defined by power_phase_type.
	AvgPowerPosition            []uint16             // Array: [N]; Units: watts; Average power by position. Data value indexes defined by rider_position_type.
	MaxPowerPosition            []uint16             // Array: [N]; Units: watts; Maximum power by position. Data value indexes defined by rider_position_type.
	AvgCadencePosition          []uint8              // Array: [N]; Units: rpm; Average cadence by position. Data value indexes defined by rider_position_type.
	MaxCadencePosition          []uint8              // Array: [N]; Units: rpm; Maximum cadence by position. Data value indexes defined by rider_position_type.
	Manufacturer                typedef.Manufacturer // Manufacturer that produced the segment
	TotalGrit                   float32              // Units: kGrit; The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.
	TotalFlow                   float32              // Units: Flow; The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.
	AvgGrit                     float32              // Units: kGrit; The grit score estimates how challenging a route could be for a cyclist in terms of time spent going over sharp turns or large grade slopes.
	AvgFlow                     float32              // Units: Flow; The flow score estimates how long distance wise a cyclist deaccelerates over intervals where deacceleration is unnecessary such as smooth turns or small grade angle intervals.
	TotalFractionalAscent       uint8                // Scale: 100; Units: m; fractional part of total_ascent
	TotalFractionalDescent      uint8                // Scale: 100; Units: m; fractional part of total_descent
	EnhancedAvgAltitude         uint32               // Scale: 5; Offset: 500; Units: m;
	EnhancedMaxAltitude         uint32               // Scale: 5; Offset: 500; Units: m;
	EnhancedMinAltitude         uint32               // Scale: 5; Offset: 500; Units: m;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentLap creates new SegmentLap struct based on given mesg. If mesg is nil or mesg.Num is not equal to SegmentLap mesg number, it will return nil.
func NewSegmentLap(mesg proto.Message) *SegmentLap {
	if mesg.Num != typedef.MesgNumSegmentLap {
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
		23:  basetype.EnumInvalid,                          /* Sport */
		24:  basetype.Uint8Invalid,                         /* EventGroup */
		25:  basetype.Sint32Invalid,                        /* NecLat */
		26:  basetype.Sint32Invalid,                        /* NecLong */
		27:  basetype.Sint32Invalid,                        /* SwcLat */
		28:  basetype.Sint32Invalid,                        /* SwcLong */
		29:  basetype.StringInvalid,                        /* Name */
		30:  basetype.Uint16Invalid,                        /* NormalizedPower */
		31:  basetype.Uint16Invalid,                        /* LeftRightBalance */
		32:  basetype.EnumInvalid,                          /* SubSport */
		33:  basetype.Uint32Invalid,                        /* TotalWork */
		34:  basetype.Uint16Invalid,                        /* AvgAltitude */
		35:  basetype.Uint16Invalid,                        /* MaxAltitude */
		36:  basetype.Uint8Invalid,                         /* GpsAccuracy */
		37:  basetype.Sint16Invalid,                        /* AvgGrade */
		38:  basetype.Sint16Invalid,                        /* AvgPosGrade */
		39:  basetype.Sint16Invalid,                        /* AvgNegGrade */
		40:  basetype.Sint16Invalid,                        /* MaxPosGrade */
		41:  basetype.Sint16Invalid,                        /* MaxNegGrade */
		42:  basetype.Sint8Invalid,                         /* AvgTemperature */
		43:  basetype.Sint8Invalid,                         /* MaxTemperature */
		44:  basetype.Uint32Invalid,                        /* TotalMovingTime */
		45:  basetype.Sint16Invalid,                        /* AvgPosVerticalSpeed */
		46:  basetype.Sint16Invalid,                        /* AvgNegVerticalSpeed */
		47:  basetype.Sint16Invalid,                        /* MaxPosVerticalSpeed */
		48:  basetype.Sint16Invalid,                        /* MaxNegVerticalSpeed */
		49:  nil,                                           /* TimeInHrZone */
		50:  nil,                                           /* TimeInSpeedZone */
		51:  nil,                                           /* TimeInCadenceZone */
		52:  nil,                                           /* TimeInPowerZone */
		53:  basetype.Uint16Invalid,                        /* RepetitionNum */
		54:  basetype.Uint16Invalid,                        /* MinAltitude */
		55:  basetype.Uint8Invalid,                         /* MinHeartRate */
		56:  basetype.Uint32Invalid,                        /* ActiveTime */
		57:  basetype.Uint16Invalid,                        /* WktStepIndex */
		58:  basetype.EnumInvalid,                          /* SportEvent */
		59:  basetype.Uint8Invalid,                         /* AvgLeftTorqueEffectiveness */
		60:  basetype.Uint8Invalid,                         /* AvgRightTorqueEffectiveness */
		61:  basetype.Uint8Invalid,                         /* AvgLeftPedalSmoothness */
		62:  basetype.Uint8Invalid,                         /* AvgRightPedalSmoothness */
		63:  basetype.Uint8Invalid,                         /* AvgCombinedPedalSmoothness */
		64:  basetype.EnumInvalid,                          /* Status */
		65:  basetype.StringInvalid,                        /* Uuid */
		66:  basetype.Uint8Invalid,                         /* AvgFractionalCadence */
		67:  basetype.Uint8Invalid,                         /* MaxFractionalCadence */
		68:  basetype.Uint8Invalid,                         /* TotalFractionalCycles */
		69:  basetype.Uint16Invalid,                        /* FrontGearShiftCount */
		70:  basetype.Uint16Invalid,                        /* RearGearShiftCount */
		71:  basetype.Uint32Invalid,                        /* TimeStanding */
		72:  basetype.Uint16Invalid,                        /* StandCount */
		73:  basetype.Sint8Invalid,                         /* AvgLeftPco */
		74:  basetype.Sint8Invalid,                         /* AvgRightPco */
		75:  nil,                                           /* AvgLeftPowerPhase */
		76:  nil,                                           /* AvgLeftPowerPhasePeak */
		77:  nil,                                           /* AvgRightPowerPhase */
		78:  nil,                                           /* AvgRightPowerPhasePeak */
		79:  nil,                                           /* AvgPowerPosition */
		80:  nil,                                           /* MaxPowerPosition */
		81:  nil,                                           /* AvgCadencePosition */
		82:  nil,                                           /* MaxCadencePosition */
		83:  basetype.Uint16Invalid,                        /* Manufacturer */
		84:  math.Float32frombits(basetype.Float32Invalid), /* TotalGrit */
		85:  math.Float32frombits(basetype.Float32Invalid), /* TotalFlow */
		86:  math.Float32frombits(basetype.Float32Invalid), /* AvgGrit */
		87:  math.Float32frombits(basetype.Float32Invalid), /* AvgFlow */
		89:  basetype.Uint8Invalid,                         /* TotalFractionalAscent */
		90:  basetype.Uint8Invalid,                         /* TotalFractionalDescent */
		91:  basetype.Uint32Invalid,                        /* EnhancedAvgAltitude */
		92:  basetype.Uint32Invalid,                        /* EnhancedMaxAltitude */
		93:  basetype.Uint32Invalid,                        /* EnhancedMinAltitude */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &SegmentLap{
		MessageIndex:                typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Timestamp:                   typeconv.ToUint32[typedef.DateTime](vals[253]),
		Event:                       typeconv.ToEnum[typedef.Event](vals[0]),
		EventType:                   typeconv.ToEnum[typedef.EventType](vals[1]),
		StartTime:                   typeconv.ToUint32[typedef.DateTime](vals[2]),
		StartPositionLat:            typeconv.ToSint32[int32](vals[3]),
		StartPositionLong:           typeconv.ToSint32[int32](vals[4]),
		EndPositionLat:              typeconv.ToSint32[int32](vals[5]),
		EndPositionLong:             typeconv.ToSint32[int32](vals[6]),
		TotalElapsedTime:            typeconv.ToUint32[uint32](vals[7]),
		TotalTimerTime:              typeconv.ToUint32[uint32](vals[8]),
		TotalDistance:               typeconv.ToUint32[uint32](vals[9]),
		TotalCycles:                 typeconv.ToUint32[uint32](vals[10]),
		TotalCalories:               typeconv.ToUint16[uint16](vals[11]),
		TotalFatCalories:            typeconv.ToUint16[uint16](vals[12]),
		AvgSpeed:                    typeconv.ToUint16[uint16](vals[13]),
		MaxSpeed:                    typeconv.ToUint16[uint16](vals[14]),
		AvgHeartRate:                typeconv.ToUint8[uint8](vals[15]),
		MaxHeartRate:                typeconv.ToUint8[uint8](vals[16]),
		AvgCadence:                  typeconv.ToUint8[uint8](vals[17]),
		MaxCadence:                  typeconv.ToUint8[uint8](vals[18]),
		AvgPower:                    typeconv.ToUint16[uint16](vals[19]),
		MaxPower:                    typeconv.ToUint16[uint16](vals[20]),
		TotalAscent:                 typeconv.ToUint16[uint16](vals[21]),
		TotalDescent:                typeconv.ToUint16[uint16](vals[22]),
		Sport:                       typeconv.ToEnum[typedef.Sport](vals[23]),
		EventGroup:                  typeconv.ToUint8[uint8](vals[24]),
		NecLat:                      typeconv.ToSint32[int32](vals[25]),
		NecLong:                     typeconv.ToSint32[int32](vals[26]),
		SwcLat:                      typeconv.ToSint32[int32](vals[27]),
		SwcLong:                     typeconv.ToSint32[int32](vals[28]),
		Name:                        typeconv.ToString[string](vals[29]),
		NormalizedPower:             typeconv.ToUint16[uint16](vals[30]),
		LeftRightBalance:            typeconv.ToUint16[typedef.LeftRightBalance100](vals[31]),
		SubSport:                    typeconv.ToEnum[typedef.SubSport](vals[32]),
		TotalWork:                   typeconv.ToUint32[uint32](vals[33]),
		AvgAltitude:                 typeconv.ToUint16[uint16](vals[34]),
		MaxAltitude:                 typeconv.ToUint16[uint16](vals[35]),
		GpsAccuracy:                 typeconv.ToUint8[uint8](vals[36]),
		AvgGrade:                    typeconv.ToSint16[int16](vals[37]),
		AvgPosGrade:                 typeconv.ToSint16[int16](vals[38]),
		AvgNegGrade:                 typeconv.ToSint16[int16](vals[39]),
		MaxPosGrade:                 typeconv.ToSint16[int16](vals[40]),
		MaxNegGrade:                 typeconv.ToSint16[int16](vals[41]),
		AvgTemperature:              typeconv.ToSint8[int8](vals[42]),
		MaxTemperature:              typeconv.ToSint8[int8](vals[43]),
		TotalMovingTime:             typeconv.ToUint32[uint32](vals[44]),
		AvgPosVerticalSpeed:         typeconv.ToSint16[int16](vals[45]),
		AvgNegVerticalSpeed:         typeconv.ToSint16[int16](vals[46]),
		MaxPosVerticalSpeed:         typeconv.ToSint16[int16](vals[47]),
		MaxNegVerticalSpeed:         typeconv.ToSint16[int16](vals[48]),
		TimeInHrZone:                typeconv.ToSliceUint32[uint32](vals[49]),
		TimeInSpeedZone:             typeconv.ToSliceUint32[uint32](vals[50]),
		TimeInCadenceZone:           typeconv.ToSliceUint32[uint32](vals[51]),
		TimeInPowerZone:             typeconv.ToSliceUint32[uint32](vals[52]),
		RepetitionNum:               typeconv.ToUint16[uint16](vals[53]),
		MinAltitude:                 typeconv.ToUint16[uint16](vals[54]),
		MinHeartRate:                typeconv.ToUint8[uint8](vals[55]),
		ActiveTime:                  typeconv.ToUint32[uint32](vals[56]),
		WktStepIndex:                typeconv.ToUint16[typedef.MessageIndex](vals[57]),
		SportEvent:                  typeconv.ToEnum[typedef.SportEvent](vals[58]),
		AvgLeftTorqueEffectiveness:  typeconv.ToUint8[uint8](vals[59]),
		AvgRightTorqueEffectiveness: typeconv.ToUint8[uint8](vals[60]),
		AvgLeftPedalSmoothness:      typeconv.ToUint8[uint8](vals[61]),
		AvgRightPedalSmoothness:     typeconv.ToUint8[uint8](vals[62]),
		AvgCombinedPedalSmoothness:  typeconv.ToUint8[uint8](vals[63]),
		Status:                      typeconv.ToEnum[typedef.SegmentLapStatus](vals[64]),
		Uuid:                        typeconv.ToString[string](vals[65]),
		AvgFractionalCadence:        typeconv.ToUint8[uint8](vals[66]),
		MaxFractionalCadence:        typeconv.ToUint8[uint8](vals[67]),
		TotalFractionalCycles:       typeconv.ToUint8[uint8](vals[68]),
		FrontGearShiftCount:         typeconv.ToUint16[uint16](vals[69]),
		RearGearShiftCount:          typeconv.ToUint16[uint16](vals[70]),
		TimeStanding:                typeconv.ToUint32[uint32](vals[71]),
		StandCount:                  typeconv.ToUint16[uint16](vals[72]),
		AvgLeftPco:                  typeconv.ToSint8[int8](vals[73]),
		AvgRightPco:                 typeconv.ToSint8[int8](vals[74]),
		AvgLeftPowerPhase:           typeconv.ToSliceUint8[uint8](vals[75]),
		AvgLeftPowerPhasePeak:       typeconv.ToSliceUint8[uint8](vals[76]),
		AvgRightPowerPhase:          typeconv.ToSliceUint8[uint8](vals[77]),
		AvgRightPowerPhasePeak:      typeconv.ToSliceUint8[uint8](vals[78]),
		AvgPowerPosition:            typeconv.ToSliceUint16[uint16](vals[79]),
		MaxPowerPosition:            typeconv.ToSliceUint16[uint16](vals[80]),
		AvgCadencePosition:          typeconv.ToSliceUint8[uint8](vals[81]),
		MaxCadencePosition:          typeconv.ToSliceUint8[uint8](vals[82]),
		Manufacturer:                typeconv.ToUint16[typedef.Manufacturer](vals[83]),
		TotalGrit:                   typeconv.ToFloat32[float32](vals[84]),
		TotalFlow:                   typeconv.ToFloat32[float32](vals[85]),
		AvgGrit:                     typeconv.ToFloat32[float32](vals[86]),
		AvgFlow:                     typeconv.ToFloat32[float32](vals[87]),
		TotalFractionalAscent:       typeconv.ToUint8[uint8](vals[89]),
		TotalFractionalDescent:      typeconv.ToUint8[uint8](vals[90]),
		EnhancedAvgAltitude:         typeconv.ToUint32[uint32](vals[91]),
		EnhancedMaxAltitude:         typeconv.ToUint32[uint32](vals[92]),
		EnhancedMinAltitude:         typeconv.ToUint32[uint32](vals[93]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SegmentLap mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSegmentLap)
func (m SegmentLap) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSegmentLap {
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
		23:  m.Sport,
		24:  m.EventGroup,
		25:  m.NecLat,
		26:  m.NecLong,
		27:  m.SwcLat,
		28:  m.SwcLong,
		29:  m.Name,
		30:  m.NormalizedPower,
		31:  m.LeftRightBalance,
		32:  m.SubSport,
		33:  m.TotalWork,
		34:  m.AvgAltitude,
		35:  m.MaxAltitude,
		36:  m.GpsAccuracy,
		37:  m.AvgGrade,
		38:  m.AvgPosGrade,
		39:  m.AvgNegGrade,
		40:  m.MaxPosGrade,
		41:  m.MaxNegGrade,
		42:  m.AvgTemperature,
		43:  m.MaxTemperature,
		44:  m.TotalMovingTime,
		45:  m.AvgPosVerticalSpeed,
		46:  m.AvgNegVerticalSpeed,
		47:  m.MaxPosVerticalSpeed,
		48:  m.MaxNegVerticalSpeed,
		49:  m.TimeInHrZone,
		50:  m.TimeInSpeedZone,
		51:  m.TimeInCadenceZone,
		52:  m.TimeInPowerZone,
		53:  m.RepetitionNum,
		54:  m.MinAltitude,
		55:  m.MinHeartRate,
		56:  m.ActiveTime,
		57:  m.WktStepIndex,
		58:  m.SportEvent,
		59:  m.AvgLeftTorqueEffectiveness,
		60:  m.AvgRightTorqueEffectiveness,
		61:  m.AvgLeftPedalSmoothness,
		62:  m.AvgRightPedalSmoothness,
		63:  m.AvgCombinedPedalSmoothness,
		64:  m.Status,
		65:  m.Uuid,
		66:  m.AvgFractionalCadence,
		67:  m.MaxFractionalCadence,
		68:  m.TotalFractionalCycles,
		69:  m.FrontGearShiftCount,
		70:  m.RearGearShiftCount,
		71:  m.TimeStanding,
		72:  m.StandCount,
		73:  m.AvgLeftPco,
		74:  m.AvgRightPco,
		75:  m.AvgLeftPowerPhase,
		76:  m.AvgLeftPowerPhasePeak,
		77:  m.AvgRightPowerPhase,
		78:  m.AvgRightPowerPhasePeak,
		79:  m.AvgPowerPosition,
		80:  m.MaxPowerPosition,
		81:  m.AvgCadencePosition,
		82:  m.MaxCadencePosition,
		83:  m.Manufacturer,
		84:  m.TotalGrit,
		85:  m.TotalFlow,
		86:  m.AvgGrit,
		87:  m.AvgFlow,
		89:  m.TotalFractionalAscent,
		90:  m.TotalFractionalDescent,
		91:  m.EnhancedAvgAltitude,
		92:  m.EnhancedMaxAltitude,
		93:  m.EnhancedMinAltitude,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
