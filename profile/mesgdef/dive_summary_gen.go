// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// DiveSummary is a DiveSummary message.
type DiveSummary struct {
	Timestamp       time.Time // Units: s;
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

// NewDiveSummary creates new DiveSummary struct based on given mesg.
// If mesg is nil, it will return DiveSummary with all fields being set to its corresponding invalid value.
func NewDiveSummary(mesg *proto.Message) *DiveSummary {
	vals := [254]any{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &DiveSummary{
		Timestamp:       datetime.ToTime(vals[253]),
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

		DeveloperFields: developerFields,
	}
}

// ToMesg converts DiveSummary into proto.Message.
func (m *DiveSummary) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumDiveSummary)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.ReferenceMesg) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToUint16[uint16](m.ReferenceMesg)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16[uint16](m.ReferenceIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToUint16[uint16](m.ReferenceIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgDepth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.AvgDepth
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxDepth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.MaxDepth
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.SurfaceInterval != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.SurfaceInterval
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.StartCns != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.StartCns
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EndCns != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.EndCns
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.StartN2 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.StartN2
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EndN2 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.EndN2
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.O2Toxicity != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.O2Toxicity
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DiveNumber != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.DiveNumber
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.BottomTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.BottomTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgPressureSac != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = m.AvgPressureSac
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgVolumeSac != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = m.AvgVolumeSac
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgRmv != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.AvgRmv
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.DescentTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = m.DescentTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AscentTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = m.AscentTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgAscentRate != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = m.AvgAscentRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgDescentRate != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = m.AvgDescentRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxAscentRate != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.MaxAscentRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxDescentRate != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = m.MaxDescentRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.HangTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = m.HangTime
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of DiveSummary's valid fields.
func (m *DiveSummary) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.ReferenceMesg) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToUint16[uint16](m.ReferenceIndex) != basetype.Uint16Invalid {
		size++
	}
	if m.AvgDepth != basetype.Uint32Invalid {
		size++
	}
	if m.MaxDepth != basetype.Uint32Invalid {
		size++
	}
	if m.SurfaceInterval != basetype.Uint32Invalid {
		size++
	}
	if m.StartCns != basetype.Uint8Invalid {
		size++
	}
	if m.EndCns != basetype.Uint8Invalid {
		size++
	}
	if m.StartN2 != basetype.Uint16Invalid {
		size++
	}
	if m.EndN2 != basetype.Uint16Invalid {
		size++
	}
	if m.O2Toxicity != basetype.Uint16Invalid {
		size++
	}
	if m.DiveNumber != basetype.Uint32Invalid {
		size++
	}
	if m.BottomTime != basetype.Uint32Invalid {
		size++
	}
	if m.AvgPressureSac != basetype.Uint16Invalid {
		size++
	}
	if m.AvgVolumeSac != basetype.Uint16Invalid {
		size++
	}
	if m.AvgRmv != basetype.Uint16Invalid {
		size++
	}
	if m.DescentTime != basetype.Uint32Invalid {
		size++
	}
	if m.AscentTime != basetype.Uint32Invalid {
		size++
	}
	if m.AvgAscentRate != basetype.Sint32Invalid {
		size++
	}
	if m.AvgDescentRate != basetype.Uint32Invalid {
		size++
	}
	if m.MaxAscentRate != basetype.Uint32Invalid {
		size++
	}
	if m.MaxDescentRate != basetype.Uint32Invalid {
		size++
	}
	if m.HangTime != basetype.Uint32Invalid {
		size++
	}
	return size
}

// SetTimestamp sets DiveSummary value.
//
// Units: s;
func (m *DiveSummary) SetTimestamp(v time.Time) *DiveSummary {
	m.Timestamp = v
	return m
}

// SetReferenceMesg sets DiveSummary value.
func (m *DiveSummary) SetReferenceMesg(v typedef.MesgNum) *DiveSummary {
	m.ReferenceMesg = v
	return m
}

// SetReferenceIndex sets DiveSummary value.
func (m *DiveSummary) SetReferenceIndex(v typedef.MessageIndex) *DiveSummary {
	m.ReferenceIndex = v
	return m
}

// SetAvgDepth sets DiveSummary value.
//
// Scale: 1000; Units: m; 0 if above water
func (m *DiveSummary) SetAvgDepth(v uint32) *DiveSummary {
	m.AvgDepth = v
	return m
}

// SetMaxDepth sets DiveSummary value.
//
// Scale: 1000; Units: m; 0 if above water
func (m *DiveSummary) SetMaxDepth(v uint32) *DiveSummary {
	m.MaxDepth = v
	return m
}

// SetSurfaceInterval sets DiveSummary value.
//
// Units: s; Time since end of last dive
func (m *DiveSummary) SetSurfaceInterval(v uint32) *DiveSummary {
	m.SurfaceInterval = v
	return m
}

// SetStartCns sets DiveSummary value.
//
// Units: percent;
func (m *DiveSummary) SetStartCns(v uint8) *DiveSummary {
	m.StartCns = v
	return m
}

// SetEndCns sets DiveSummary value.
//
// Units: percent;
func (m *DiveSummary) SetEndCns(v uint8) *DiveSummary {
	m.EndCns = v
	return m
}

// SetStartN2 sets DiveSummary value.
//
// Units: percent;
func (m *DiveSummary) SetStartN2(v uint16) *DiveSummary {
	m.StartN2 = v
	return m
}

// SetEndN2 sets DiveSummary value.
//
// Units: percent;
func (m *DiveSummary) SetEndN2(v uint16) *DiveSummary {
	m.EndN2 = v
	return m
}

// SetO2Toxicity sets DiveSummary value.
//
// Units: OTUs;
func (m *DiveSummary) SetO2Toxicity(v uint16) *DiveSummary {
	m.O2Toxicity = v
	return m
}

// SetDiveNumber sets DiveSummary value.
func (m *DiveSummary) SetDiveNumber(v uint32) *DiveSummary {
	m.DiveNumber = v
	return m
}

// SetBottomTime sets DiveSummary value.
//
// Scale: 1000; Units: s;
func (m *DiveSummary) SetBottomTime(v uint32) *DiveSummary {
	m.BottomTime = v
	return m
}

// SetAvgPressureSac sets DiveSummary value.
//
// Scale: 100; Units: bar/min; Average pressure-based surface air consumption
func (m *DiveSummary) SetAvgPressureSac(v uint16) *DiveSummary {
	m.AvgPressureSac = v
	return m
}

// SetAvgVolumeSac sets DiveSummary value.
//
// Scale: 100; Units: L/min; Average volumetric surface air consumption
func (m *DiveSummary) SetAvgVolumeSac(v uint16) *DiveSummary {
	m.AvgVolumeSac = v
	return m
}

// SetAvgRmv sets DiveSummary value.
//
// Scale: 100; Units: L/min; Average respiratory minute volume
func (m *DiveSummary) SetAvgRmv(v uint16) *DiveSummary {
	m.AvgRmv = v
	return m
}

// SetDescentTime sets DiveSummary value.
//
// Scale: 1000; Units: s; Time to reach deepest level stop
func (m *DiveSummary) SetDescentTime(v uint32) *DiveSummary {
	m.DescentTime = v
	return m
}

// SetAscentTime sets DiveSummary value.
//
// Scale: 1000; Units: s; Time after leaving bottom until reaching surface
func (m *DiveSummary) SetAscentTime(v uint32) *DiveSummary {
	m.AscentTime = v
	return m
}

// SetAvgAscentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Average ascent rate, not including descents or stops
func (m *DiveSummary) SetAvgAscentRate(v int32) *DiveSummary {
	m.AvgAscentRate = v
	return m
}

// SetAvgDescentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Average descent rate, not including ascents or stops
func (m *DiveSummary) SetAvgDescentRate(v uint32) *DiveSummary {
	m.AvgDescentRate = v
	return m
}

// SetMaxAscentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Maximum ascent rate
func (m *DiveSummary) SetMaxAscentRate(v uint32) *DiveSummary {
	m.MaxAscentRate = v
	return m
}

// SetMaxDescentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Maximum descent rate
func (m *DiveSummary) SetMaxDescentRate(v uint32) *DiveSummary {
	m.MaxDescentRate = v
	return m
}

// SetHangTime sets DiveSummary value.
//
// Scale: 1000; Units: s; Time spent neither ascending nor descending
func (m *DiveSummary) SetHangTime(v uint32) *DiveSummary {
	m.HangTime = v
	return m
}

// SetDeveloperFields DiveSummary's DeveloperFields.
func (m *DiveSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *DiveSummary {
	m.DeveloperFields = developerFields
	return m
}
