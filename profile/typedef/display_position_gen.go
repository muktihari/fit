// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DisplayPosition byte

const (
	DisplayPositionDegree               DisplayPosition = 0    // dd.dddddd
	DisplayPositionDegreeMinute         DisplayPosition = 1    // dddmm.mmm
	DisplayPositionDegreeMinuteSecond   DisplayPosition = 2    // dddmmss
	DisplayPositionAustrianGrid         DisplayPosition = 3    // Austrian Grid (BMN)
	DisplayPositionBritishGrid          DisplayPosition = 4    // British National Grid
	DisplayPositionDutchGrid            DisplayPosition = 5    // Dutch grid system
	DisplayPositionHungarianGrid        DisplayPosition = 6    // Hungarian grid system
	DisplayPositionFinnishGrid          DisplayPosition = 7    // Finnish grid system Zone3 KKJ27
	DisplayPositionGermanGrid           DisplayPosition = 8    // Gausss Krueger (German)
	DisplayPositionIcelandicGrid        DisplayPosition = 9    // Icelandic Grid
	DisplayPositionIndonesianEquatorial DisplayPosition = 10   // Indonesian Equatorial LCO
	DisplayPositionIndonesianIrian      DisplayPosition = 11   // Indonesian Irian LCO
	DisplayPositionIndonesianSouthern   DisplayPosition = 12   // Indonesian Southern LCO
	DisplayPositionIndiaZone0           DisplayPosition = 13   // India zone 0
	DisplayPositionIndiaZoneIa          DisplayPosition = 14   // India zone IA
	DisplayPositionIndiaZoneIb          DisplayPosition = 15   // India zone IB
	DisplayPositionIndiaZoneIia         DisplayPosition = 16   // India zone IIA
	DisplayPositionIndiaZoneIib         DisplayPosition = 17   // India zone IIB
	DisplayPositionIndiaZoneIiia        DisplayPosition = 18   // India zone IIIA
	DisplayPositionIndiaZoneIiib        DisplayPosition = 19   // India zone IIIB
	DisplayPositionIndiaZoneIva         DisplayPosition = 20   // India zone IVA
	DisplayPositionIndiaZoneIvb         DisplayPosition = 21   // India zone IVB
	DisplayPositionIrishTransverse      DisplayPosition = 22   // Irish Transverse Mercator
	DisplayPositionIrishGrid            DisplayPosition = 23   // Irish Grid
	DisplayPositionLoran                DisplayPosition = 24   // Loran TD
	DisplayPositionMaidenheadGrid       DisplayPosition = 25   // Maidenhead grid system
	DisplayPositionMgrsGrid             DisplayPosition = 26   // MGRS grid system
	DisplayPositionNewZealandGrid       DisplayPosition = 27   // New Zealand grid system
	DisplayPositionNewZealandTransverse DisplayPosition = 28   // New Zealand Transverse Mercator
	DisplayPositionQatarGrid            DisplayPosition = 29   // Qatar National Grid
	DisplayPositionModifiedSwedishGrid  DisplayPosition = 30   // Modified RT-90 (Sweden)
	DisplayPositionSwedishGrid          DisplayPosition = 31   // RT-90 (Sweden)
	DisplayPositionSouthAfricanGrid     DisplayPosition = 32   // South African Grid
	DisplayPositionSwissGrid            DisplayPosition = 33   // Swiss CH-1903 grid
	DisplayPositionTaiwanGrid           DisplayPosition = 34   // Taiwan Grid
	DisplayPositionUnitedStatesGrid     DisplayPosition = 35   // United States National Grid
	DisplayPositionUtmUpsGrid           DisplayPosition = 36   // UTM/UPS grid system
	DisplayPositionWestMalayan          DisplayPosition = 37   // West Malayan RSO
	DisplayPositionBorneoRso            DisplayPosition = 38   // Borneo RSO
	DisplayPositionEstonianGrid         DisplayPosition = 39   // Estonian grid system
	DisplayPositionLatvianGrid          DisplayPosition = 40   // Latvian Transverse Mercator
	DisplayPositionSwedishRef99Grid     DisplayPosition = 41   // Reference Grid 99 TM (Swedish)
	DisplayPositionInvalid              DisplayPosition = 0xFF // INVALID
)

var displaypositiontostrs = map[DisplayPosition]string{
	DisplayPositionDegree:               "degree",
	DisplayPositionDegreeMinute:         "degree_minute",
	DisplayPositionDegreeMinuteSecond:   "degree_minute_second",
	DisplayPositionAustrianGrid:         "austrian_grid",
	DisplayPositionBritishGrid:          "british_grid",
	DisplayPositionDutchGrid:            "dutch_grid",
	DisplayPositionHungarianGrid:        "hungarian_grid",
	DisplayPositionFinnishGrid:          "finnish_grid",
	DisplayPositionGermanGrid:           "german_grid",
	DisplayPositionIcelandicGrid:        "icelandic_grid",
	DisplayPositionIndonesianEquatorial: "indonesian_equatorial",
	DisplayPositionIndonesianIrian:      "indonesian_irian",
	DisplayPositionIndonesianSouthern:   "indonesian_southern",
	DisplayPositionIndiaZone0:           "india_zone_0",
	DisplayPositionIndiaZoneIa:          "india_zone_IA",
	DisplayPositionIndiaZoneIb:          "india_zone_IB",
	DisplayPositionIndiaZoneIia:         "india_zone_IIA",
	DisplayPositionIndiaZoneIib:         "india_zone_IIB",
	DisplayPositionIndiaZoneIiia:        "india_zone_IIIA",
	DisplayPositionIndiaZoneIiib:        "india_zone_IIIB",
	DisplayPositionIndiaZoneIva:         "india_zone_IVA",
	DisplayPositionIndiaZoneIvb:         "india_zone_IVB",
	DisplayPositionIrishTransverse:      "irish_transverse",
	DisplayPositionIrishGrid:            "irish_grid",
	DisplayPositionLoran:                "loran",
	DisplayPositionMaidenheadGrid:       "maidenhead_grid",
	DisplayPositionMgrsGrid:             "mgrs_grid",
	DisplayPositionNewZealandGrid:       "new_zealand_grid",
	DisplayPositionNewZealandTransverse: "new_zealand_transverse",
	DisplayPositionQatarGrid:            "qatar_grid",
	DisplayPositionModifiedSwedishGrid:  "modified_swedish_grid",
	DisplayPositionSwedishGrid:          "swedish_grid",
	DisplayPositionSouthAfricanGrid:     "south_african_grid",
	DisplayPositionSwissGrid:            "swiss_grid",
	DisplayPositionTaiwanGrid:           "taiwan_grid",
	DisplayPositionUnitedStatesGrid:     "united_states_grid",
	DisplayPositionUtmUpsGrid:           "utm_ups_grid",
	DisplayPositionWestMalayan:          "west_malayan",
	DisplayPositionBorneoRso:            "borneo_rso",
	DisplayPositionEstonianGrid:         "estonian_grid",
	DisplayPositionLatvianGrid:          "latvian_grid",
	DisplayPositionSwedishRef99Grid:     "swedish_ref_99_grid",
	DisplayPositionInvalid:              "invalid",
}

func (d DisplayPosition) String() string {
	val, ok := displaypositiontostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodisplayposition = func() map[string]DisplayPosition {
	m := make(map[string]DisplayPosition)
	for t, str := range displaypositiontostrs {
		m[str] = DisplayPosition(t)
	}
	return m
}()

// FromString parse string into DisplayPosition constant it's represent, return DisplayPositionInvalid if not found.
func DisplayPositionFromString(s string) DisplayPosition {
	val, ok := strtodisplayposition[s]
	if !ok {
		return strtodisplayposition["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDisplayPosition() []DisplayPosition {
	vs := make([]DisplayPosition, 0, len(displaypositiontostrs))
	for i := range displaypositiontostrs {
		vs = append(vs, DisplayPosition(i))
	}
	return vs
}
