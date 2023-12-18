// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DisplayPosition byte

const (
	DisplayPositionDegree               DisplayPosition = 0  // dd.dddddd
	DisplayPositionDegreeMinute         DisplayPosition = 1  // dddmm.mmm
	DisplayPositionDegreeMinuteSecond   DisplayPosition = 2  // dddmmss
	DisplayPositionAustrianGrid         DisplayPosition = 3  // Austrian Grid (BMN)
	DisplayPositionBritishGrid          DisplayPosition = 4  // British National Grid
	DisplayPositionDutchGrid            DisplayPosition = 5  // Dutch grid system
	DisplayPositionHungarianGrid        DisplayPosition = 6  // Hungarian grid system
	DisplayPositionFinnishGrid          DisplayPosition = 7  // Finnish grid system Zone3 KKJ27
	DisplayPositionGermanGrid           DisplayPosition = 8  // Gausss Krueger (German)
	DisplayPositionIcelandicGrid        DisplayPosition = 9  // Icelandic Grid
	DisplayPositionIndonesianEquatorial DisplayPosition = 10 // Indonesian Equatorial LCO
	DisplayPositionIndonesianIrian      DisplayPosition = 11 // Indonesian Irian LCO
	DisplayPositionIndonesianSouthern   DisplayPosition = 12 // Indonesian Southern LCO
	DisplayPositionIndiaZone0           DisplayPosition = 13 // India zone 0
	DisplayPositionIndiaZoneIa          DisplayPosition = 14 // India zone IA
	DisplayPositionIndiaZoneIb          DisplayPosition = 15 // India zone IB
	DisplayPositionIndiaZoneIia         DisplayPosition = 16 // India zone IIA
	DisplayPositionIndiaZoneIib         DisplayPosition = 17 // India zone IIB
	DisplayPositionIndiaZoneIiia        DisplayPosition = 18 // India zone IIIA
	DisplayPositionIndiaZoneIiib        DisplayPosition = 19 // India zone IIIB
	DisplayPositionIndiaZoneIva         DisplayPosition = 20 // India zone IVA
	DisplayPositionIndiaZoneIvb         DisplayPosition = 21 // India zone IVB
	DisplayPositionIrishTransverse      DisplayPosition = 22 // Irish Transverse Mercator
	DisplayPositionIrishGrid            DisplayPosition = 23 // Irish Grid
	DisplayPositionLoran                DisplayPosition = 24 // Loran TD
	DisplayPositionMaidenheadGrid       DisplayPosition = 25 // Maidenhead grid system
	DisplayPositionMgrsGrid             DisplayPosition = 26 // MGRS grid system
	DisplayPositionNewZealandGrid       DisplayPosition = 27 // New Zealand grid system
	DisplayPositionNewZealandTransverse DisplayPosition = 28 // New Zealand Transverse Mercator
	DisplayPositionQatarGrid            DisplayPosition = 29 // Qatar National Grid
	DisplayPositionModifiedSwedishGrid  DisplayPosition = 30 // Modified RT-90 (Sweden)
	DisplayPositionSwedishGrid          DisplayPosition = 31 // RT-90 (Sweden)
	DisplayPositionSouthAfricanGrid     DisplayPosition = 32 // South African Grid
	DisplayPositionSwissGrid            DisplayPosition = 33 // Swiss CH-1903 grid
	DisplayPositionTaiwanGrid           DisplayPosition = 34 // Taiwan Grid
	DisplayPositionUnitedStatesGrid     DisplayPosition = 35 // United States National Grid
	DisplayPositionUtmUpsGrid           DisplayPosition = 36 // UTM/UPS grid system
	DisplayPositionWestMalayan          DisplayPosition = 37 // West Malayan RSO
	DisplayPositionBorneoRso            DisplayPosition = 38 // Borneo RSO
	DisplayPositionEstonianGrid         DisplayPosition = 39 // Estonian grid system
	DisplayPositionLatvianGrid          DisplayPosition = 40 // Latvian Transverse Mercator
	DisplayPositionSwedishRef99Grid     DisplayPosition = 41 // Reference Grid 99 TM (Swedish)
	DisplayPositionInvalid              DisplayPosition = 0xFF
)

func (d DisplayPosition) String() string {
	switch d {
	case DisplayPositionDegree:
		return "degree"
	case DisplayPositionDegreeMinute:
		return "degree_minute"
	case DisplayPositionDegreeMinuteSecond:
		return "degree_minute_second"
	case DisplayPositionAustrianGrid:
		return "austrian_grid"
	case DisplayPositionBritishGrid:
		return "british_grid"
	case DisplayPositionDutchGrid:
		return "dutch_grid"
	case DisplayPositionHungarianGrid:
		return "hungarian_grid"
	case DisplayPositionFinnishGrid:
		return "finnish_grid"
	case DisplayPositionGermanGrid:
		return "german_grid"
	case DisplayPositionIcelandicGrid:
		return "icelandic_grid"
	case DisplayPositionIndonesianEquatorial:
		return "indonesian_equatorial"
	case DisplayPositionIndonesianIrian:
		return "indonesian_irian"
	case DisplayPositionIndonesianSouthern:
		return "indonesian_southern"
	case DisplayPositionIndiaZone0:
		return "india_zone_0"
	case DisplayPositionIndiaZoneIa:
		return "india_zone_IA"
	case DisplayPositionIndiaZoneIb:
		return "india_zone_IB"
	case DisplayPositionIndiaZoneIia:
		return "india_zone_IIA"
	case DisplayPositionIndiaZoneIib:
		return "india_zone_IIB"
	case DisplayPositionIndiaZoneIiia:
		return "india_zone_IIIA"
	case DisplayPositionIndiaZoneIiib:
		return "india_zone_IIIB"
	case DisplayPositionIndiaZoneIva:
		return "india_zone_IVA"
	case DisplayPositionIndiaZoneIvb:
		return "india_zone_IVB"
	case DisplayPositionIrishTransverse:
		return "irish_transverse"
	case DisplayPositionIrishGrid:
		return "irish_grid"
	case DisplayPositionLoran:
		return "loran"
	case DisplayPositionMaidenheadGrid:
		return "maidenhead_grid"
	case DisplayPositionMgrsGrid:
		return "mgrs_grid"
	case DisplayPositionNewZealandGrid:
		return "new_zealand_grid"
	case DisplayPositionNewZealandTransverse:
		return "new_zealand_transverse"
	case DisplayPositionQatarGrid:
		return "qatar_grid"
	case DisplayPositionModifiedSwedishGrid:
		return "modified_swedish_grid"
	case DisplayPositionSwedishGrid:
		return "swedish_grid"
	case DisplayPositionSouthAfricanGrid:
		return "south_african_grid"
	case DisplayPositionSwissGrid:
		return "swiss_grid"
	case DisplayPositionTaiwanGrid:
		return "taiwan_grid"
	case DisplayPositionUnitedStatesGrid:
		return "united_states_grid"
	case DisplayPositionUtmUpsGrid:
		return "utm_ups_grid"
	case DisplayPositionWestMalayan:
		return "west_malayan"
	case DisplayPositionBorneoRso:
		return "borneo_rso"
	case DisplayPositionEstonianGrid:
		return "estonian_grid"
	case DisplayPositionLatvianGrid:
		return "latvian_grid"
	case DisplayPositionSwedishRef99Grid:
		return "swedish_ref_99_grid"
	default:
		return "DisplayPositionInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DisplayPosition constant it's represent, return DisplayPositionInvalid if not found.
func DisplayPositionFromString(s string) DisplayPosition {
	switch s {
	case "degree":
		return DisplayPositionDegree
	case "degree_minute":
		return DisplayPositionDegreeMinute
	case "degree_minute_second":
		return DisplayPositionDegreeMinuteSecond
	case "austrian_grid":
		return DisplayPositionAustrianGrid
	case "british_grid":
		return DisplayPositionBritishGrid
	case "dutch_grid":
		return DisplayPositionDutchGrid
	case "hungarian_grid":
		return DisplayPositionHungarianGrid
	case "finnish_grid":
		return DisplayPositionFinnishGrid
	case "german_grid":
		return DisplayPositionGermanGrid
	case "icelandic_grid":
		return DisplayPositionIcelandicGrid
	case "indonesian_equatorial":
		return DisplayPositionIndonesianEquatorial
	case "indonesian_irian":
		return DisplayPositionIndonesianIrian
	case "indonesian_southern":
		return DisplayPositionIndonesianSouthern
	case "india_zone_0":
		return DisplayPositionIndiaZone0
	case "india_zone_IA":
		return DisplayPositionIndiaZoneIa
	case "india_zone_IB":
		return DisplayPositionIndiaZoneIb
	case "india_zone_IIA":
		return DisplayPositionIndiaZoneIia
	case "india_zone_IIB":
		return DisplayPositionIndiaZoneIib
	case "india_zone_IIIA":
		return DisplayPositionIndiaZoneIiia
	case "india_zone_IIIB":
		return DisplayPositionIndiaZoneIiib
	case "india_zone_IVA":
		return DisplayPositionIndiaZoneIva
	case "india_zone_IVB":
		return DisplayPositionIndiaZoneIvb
	case "irish_transverse":
		return DisplayPositionIrishTransverse
	case "irish_grid":
		return DisplayPositionIrishGrid
	case "loran":
		return DisplayPositionLoran
	case "maidenhead_grid":
		return DisplayPositionMaidenheadGrid
	case "mgrs_grid":
		return DisplayPositionMgrsGrid
	case "new_zealand_grid":
		return DisplayPositionNewZealandGrid
	case "new_zealand_transverse":
		return DisplayPositionNewZealandTransverse
	case "qatar_grid":
		return DisplayPositionQatarGrid
	case "modified_swedish_grid":
		return DisplayPositionModifiedSwedishGrid
	case "swedish_grid":
		return DisplayPositionSwedishGrid
	case "south_african_grid":
		return DisplayPositionSouthAfricanGrid
	case "swiss_grid":
		return DisplayPositionSwissGrid
	case "taiwan_grid":
		return DisplayPositionTaiwanGrid
	case "united_states_grid":
		return DisplayPositionUnitedStatesGrid
	case "utm_ups_grid":
		return DisplayPositionUtmUpsGrid
	case "west_malayan":
		return DisplayPositionWestMalayan
	case "borneo_rso":
		return DisplayPositionBorneoRso
	case "estonian_grid":
		return DisplayPositionEstonianGrid
	case "latvian_grid":
		return DisplayPositionLatvianGrid
	case "swedish_ref_99_grid":
		return DisplayPositionSwedishRef99Grid
	default:
		return DisplayPositionInvalid
	}
}

// List returns all constants.
func ListDisplayPosition() []DisplayPosition {
	return []DisplayPosition{
		DisplayPositionDegree,
		DisplayPositionDegreeMinute,
		DisplayPositionDegreeMinuteSecond,
		DisplayPositionAustrianGrid,
		DisplayPositionBritishGrid,
		DisplayPositionDutchGrid,
		DisplayPositionHungarianGrid,
		DisplayPositionFinnishGrid,
		DisplayPositionGermanGrid,
		DisplayPositionIcelandicGrid,
		DisplayPositionIndonesianEquatorial,
		DisplayPositionIndonesianIrian,
		DisplayPositionIndonesianSouthern,
		DisplayPositionIndiaZone0,
		DisplayPositionIndiaZoneIa,
		DisplayPositionIndiaZoneIb,
		DisplayPositionIndiaZoneIia,
		DisplayPositionIndiaZoneIib,
		DisplayPositionIndiaZoneIiia,
		DisplayPositionIndiaZoneIiib,
		DisplayPositionIndiaZoneIva,
		DisplayPositionIndiaZoneIvb,
		DisplayPositionIrishTransverse,
		DisplayPositionIrishGrid,
		DisplayPositionLoran,
		DisplayPositionMaidenheadGrid,
		DisplayPositionMgrsGrid,
		DisplayPositionNewZealandGrid,
		DisplayPositionNewZealandTransverse,
		DisplayPositionQatarGrid,
		DisplayPositionModifiedSwedishGrid,
		DisplayPositionSwedishGrid,
		DisplayPositionSouthAfricanGrid,
		DisplayPositionSwissGrid,
		DisplayPositionTaiwanGrid,
		DisplayPositionUnitedStatesGrid,
		DisplayPositionUtmUpsGrid,
		DisplayPositionWestMalayan,
		DisplayPositionBorneoRso,
		DisplayPositionEstonianGrid,
		DisplayPositionLatvianGrid,
		DisplayPositionSwedishRef99Grid,
	}
}
