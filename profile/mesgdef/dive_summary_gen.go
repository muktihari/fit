// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// DiveSummary is a DiveSummary message.
type DiveSummary struct {
	Timestamp       typedef.DateTime // Units: s;
	ReferenceMesg   typedef.MesgNum
	ReferenceIndex  typedef.MessageIndex
	AvgDepth        uint32 // Scale: 1000; Units: m; 0 if above water
	MaxDepth        uint32 // Scale: 1000; Units: m; 0 if above water
	SurfaceInterval uint32 // Units: s; Time since end of last dive
	StartCns        uint8  // Units: percent;
	EndCns          uint8  // Units: percent;
	StartN2         uint16 // Units: percent;
	EndN2           uint16 // Units: percent;
	O2Toxicity      uint16 // Units: OTUs;
	DiveNumber      uint32
	BottomTime      uint32 // Scale: 1000; Units: s;
	AvgPressureSac  uint16 // Scale: 100; Units: bar/min; Average pressure-based surface air consumption
	AvgVolumeSac    uint16 // Scale: 100; Units: L/min; Average volumetric surface air consumption
	AvgRmv          uint16 // Scale: 100; Units: L/min; Average respiratory minute volume
	DescentTime     uint32 // Scale: 1000; Units: s; Time to reach deepest level stop
	AscentTime      uint32 // Scale: 1000; Units: s; Time after leaving bottom until reaching surface
	AvgAscentRate   int32  // Scale: 1000; Units: m/s; Average ascent rate, not including descents or stops
	AvgDescentRate  uint32 // Scale: 1000; Units: m/s; Average descent rate, not including ascents or stops
	MaxAscentRate   uint32 // Scale: 1000; Units: m/s; Maximum ascent rate
	MaxDescentRate  uint32 // Scale: 1000; Units: m/s; Maximum descent rate
	HangTime        uint32 // Scale: 1000; Units: s; Time spent neither ascending nor descending

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDiveSummary creates new DiveSummary struct based on given mesg. If mesg is nil or mesg.Num is not equal to DiveSummary mesg number, it will return nil.
func NewDiveSummary(mesg proto.Message) *DiveSummary {
	if mesg.Num != typedef.MesgNumDiveSummary {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* ReferenceMesg */
		1:   basetype.Uint16Invalid, /* ReferenceIndex */
		2:   basetype.Uint32Invalid, /* AvgDepth */
		3:   basetype.Uint32Invalid, /* MaxDepth */
		4:   basetype.Uint32Invalid, /* SurfaceInterval */
		5:   basetype.Uint8Invalid,  /* StartCns */
		6:   basetype.Uint8Invalid,  /* EndCns */
		7:   basetype.Uint16Invalid, /* StartN2 */
		8:   basetype.Uint16Invalid, /* EndN2 */
		9:   basetype.Uint16Invalid, /* O2Toxicity */
		10:  basetype.Uint32Invalid, /* DiveNumber */
		11:  basetype.Uint32Invalid, /* BottomTime */
		12:  basetype.Uint16Invalid, /* AvgPressureSac */
		13:  basetype.Uint16Invalid, /* AvgVolumeSac */
		14:  basetype.Uint16Invalid, /* AvgRmv */
		15:  basetype.Uint32Invalid, /* DescentTime */
		16:  basetype.Uint32Invalid, /* AscentTime */
		17:  basetype.Sint32Invalid, /* AvgAscentRate */
		22:  basetype.Uint32Invalid, /* AvgDescentRate */
		23:  basetype.Uint32Invalid, /* MaxAscentRate */
		24:  basetype.Uint32Invalid, /* MaxDescentRate */
		25:  basetype.Uint32Invalid, /* HangTime */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &DiveSummary{
		Timestamp:       typeconv.ToUint32[typedef.DateTime](vals[253]),
		ReferenceMesg:   typeconv.ToUint16[typedef.MesgNum](vals[0]),
		ReferenceIndex:  typeconv.ToUint16[typedef.MessageIndex](vals[1]),
		AvgDepth:        typeconv.ToUint32[uint32](vals[2]),
		MaxDepth:        typeconv.ToUint32[uint32](vals[3]),
		SurfaceInterval: typeconv.ToUint32[uint32](vals[4]),
		StartCns:        typeconv.ToUint8[uint8](vals[5]),
		EndCns:          typeconv.ToUint8[uint8](vals[6]),
		StartN2:         typeconv.ToUint16[uint16](vals[7]),
		EndN2:           typeconv.ToUint16[uint16](vals[8]),
		O2Toxicity:      typeconv.ToUint16[uint16](vals[9]),
		DiveNumber:      typeconv.ToUint32[uint32](vals[10]),
		BottomTime:      typeconv.ToUint32[uint32](vals[11]),
		AvgPressureSac:  typeconv.ToUint16[uint16](vals[12]),
		AvgVolumeSac:    typeconv.ToUint16[uint16](vals[13]),
		AvgRmv:          typeconv.ToUint16[uint16](vals[14]),
		DescentTime:     typeconv.ToUint32[uint32](vals[15]),
		AscentTime:      typeconv.ToUint32[uint32](vals[16]),
		AvgAscentRate:   typeconv.ToSint32[int32](vals[17]),
		AvgDescentRate:  typeconv.ToUint32[uint32](vals[22]),
		MaxAscentRate:   typeconv.ToUint32[uint32](vals[23]),
		MaxDescentRate:  typeconv.ToUint32[uint32](vals[24]),
		HangTime:        typeconv.ToUint32[uint32](vals[25]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DiveSummary mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDiveSummary)
func (m DiveSummary) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDiveSummary {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.ReferenceMesg,
		1:   m.ReferenceIndex,
		2:   m.AvgDepth,
		3:   m.MaxDepth,
		4:   m.SurfaceInterval,
		5:   m.StartCns,
		6:   m.EndCns,
		7:   m.StartN2,
		8:   m.EndN2,
		9:   m.O2Toxicity,
		10:  m.DiveNumber,
		11:  m.BottomTime,
		12:  m.AvgPressureSac,
		13:  m.AvgVolumeSac,
		14:  m.AvgRmv,
		15:  m.DescentTime,
		16:  m.AscentTime,
		17:  m.AvgAscentRate,
		22:  m.AvgDescentRate,
		23:  m.MaxAscentRate,
		24:  m.MaxDescentRate,
		25:  m.HangTime,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
