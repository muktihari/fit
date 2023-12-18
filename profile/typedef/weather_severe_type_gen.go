// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WeatherSevereType byte

const (
	WeatherSevereTypeUnspecified             WeatherSevereType = 0
	WeatherSevereTypeTornado                 WeatherSevereType = 1
	WeatherSevereTypeTsunami                 WeatherSevereType = 2
	WeatherSevereTypeHurricane               WeatherSevereType = 3
	WeatherSevereTypeExtremeWind             WeatherSevereType = 4
	WeatherSevereTypeTyphoon                 WeatherSevereType = 5
	WeatherSevereTypeInlandHurricane         WeatherSevereType = 6
	WeatherSevereTypeHurricaneForceWind      WeatherSevereType = 7
	WeatherSevereTypeWaterspout              WeatherSevereType = 8
	WeatherSevereTypeSevereThunderstorm      WeatherSevereType = 9
	WeatherSevereTypeWreckhouseWinds         WeatherSevereType = 10
	WeatherSevereTypeLesSuetesWind           WeatherSevereType = 11
	WeatherSevereTypeAvalanche               WeatherSevereType = 12
	WeatherSevereTypeFlashFlood              WeatherSevereType = 13
	WeatherSevereTypeTropicalStorm           WeatherSevereType = 14
	WeatherSevereTypeInlandTropicalStorm     WeatherSevereType = 15
	WeatherSevereTypeBlizzard                WeatherSevereType = 16
	WeatherSevereTypeIceStorm                WeatherSevereType = 17
	WeatherSevereTypeFreezingRain            WeatherSevereType = 18
	WeatherSevereTypeDebrisFlow              WeatherSevereType = 19
	WeatherSevereTypeFlashFreeze             WeatherSevereType = 20
	WeatherSevereTypeDustStorm               WeatherSevereType = 21
	WeatherSevereTypeHighWind                WeatherSevereType = 22
	WeatherSevereTypeWinterStorm             WeatherSevereType = 23
	WeatherSevereTypeHeavyFreezingSpray      WeatherSevereType = 24
	WeatherSevereTypeExtremeCold             WeatherSevereType = 25
	WeatherSevereTypeWindChill               WeatherSevereType = 26
	WeatherSevereTypeColdWave                WeatherSevereType = 27
	WeatherSevereTypeHeavySnowAlert          WeatherSevereType = 28
	WeatherSevereTypeLakeEffectBlowingSnow   WeatherSevereType = 29
	WeatherSevereTypeSnowSquall              WeatherSevereType = 30
	WeatherSevereTypeLakeEffectSnow          WeatherSevereType = 31
	WeatherSevereTypeWinterWeather           WeatherSevereType = 32
	WeatherSevereTypeSleet                   WeatherSevereType = 33
	WeatherSevereTypeSnowfall                WeatherSevereType = 34
	WeatherSevereTypeSnowAndBlowingSnow      WeatherSevereType = 35
	WeatherSevereTypeBlowingSnow             WeatherSevereType = 36
	WeatherSevereTypeSnowAlert               WeatherSevereType = 37
	WeatherSevereTypeArcticOutflow           WeatherSevereType = 38
	WeatherSevereTypeFreezingDrizzle         WeatherSevereType = 39
	WeatherSevereTypeStorm                   WeatherSevereType = 40
	WeatherSevereTypeStormSurge              WeatherSevereType = 41
	WeatherSevereTypeRainfall                WeatherSevereType = 42
	WeatherSevereTypeArealFlood              WeatherSevereType = 43
	WeatherSevereTypeCoastalFlood            WeatherSevereType = 44
	WeatherSevereTypeLakeshoreFlood          WeatherSevereType = 45
	WeatherSevereTypeExcessiveHeat           WeatherSevereType = 46
	WeatherSevereTypeHeat                    WeatherSevereType = 47
	WeatherSevereTypeWeather                 WeatherSevereType = 48
	WeatherSevereTypeHighHeatAndHumidity     WeatherSevereType = 49
	WeatherSevereTypeHumidexAndHealth        WeatherSevereType = 50
	WeatherSevereTypeHumidex                 WeatherSevereType = 51
	WeatherSevereTypeGale                    WeatherSevereType = 52
	WeatherSevereTypeFreezingSpray           WeatherSevereType = 53
	WeatherSevereTypeSpecialMarine           WeatherSevereType = 54
	WeatherSevereTypeSquall                  WeatherSevereType = 55
	WeatherSevereTypeStrongWind              WeatherSevereType = 56
	WeatherSevereTypeLakeWind                WeatherSevereType = 57
	WeatherSevereTypeMarineWeather           WeatherSevereType = 58
	WeatherSevereTypeWind                    WeatherSevereType = 59
	WeatherSevereTypeSmallCraftHazardousSeas WeatherSevereType = 60
	WeatherSevereTypeHazardousSeas           WeatherSevereType = 61
	WeatherSevereTypeSmallCraft              WeatherSevereType = 62
	WeatherSevereTypeSmallCraftWinds         WeatherSevereType = 63
	WeatherSevereTypeSmallCraftRoughBar      WeatherSevereType = 64
	WeatherSevereTypeHighWaterLevel          WeatherSevereType = 65
	WeatherSevereTypeAshfall                 WeatherSevereType = 66
	WeatherSevereTypeFreezingFog             WeatherSevereType = 67
	WeatherSevereTypeDenseFog                WeatherSevereType = 68
	WeatherSevereTypeDenseSmoke              WeatherSevereType = 69
	WeatherSevereTypeBlowingDust             WeatherSevereType = 70
	WeatherSevereTypeHardFreeze              WeatherSevereType = 71
	WeatherSevereTypeFreeze                  WeatherSevereType = 72
	WeatherSevereTypeFrost                   WeatherSevereType = 73
	WeatherSevereTypeFireWeather             WeatherSevereType = 74
	WeatherSevereTypeFlood                   WeatherSevereType = 75
	WeatherSevereTypeRipTide                 WeatherSevereType = 76
	WeatherSevereTypeHighSurf                WeatherSevereType = 77
	WeatherSevereTypeSmog                    WeatherSevereType = 78
	WeatherSevereTypeAirQuality              WeatherSevereType = 79
	WeatherSevereTypeBriskWind               WeatherSevereType = 80
	WeatherSevereTypeAirStagnation           WeatherSevereType = 81
	WeatherSevereTypeLowWater                WeatherSevereType = 82
	WeatherSevereTypeHydrological            WeatherSevereType = 83
	WeatherSevereTypeSpecialWeather          WeatherSevereType = 84
	WeatherSevereTypeInvalid                 WeatherSevereType = 0xFF
)

func (w WeatherSevereType) String() string {
	switch w {
	case WeatherSevereTypeUnspecified:
		return "unspecified"
	case WeatherSevereTypeTornado:
		return "tornado"
	case WeatherSevereTypeTsunami:
		return "tsunami"
	case WeatherSevereTypeHurricane:
		return "hurricane"
	case WeatherSevereTypeExtremeWind:
		return "extreme_wind"
	case WeatherSevereTypeTyphoon:
		return "typhoon"
	case WeatherSevereTypeInlandHurricane:
		return "inland_hurricane"
	case WeatherSevereTypeHurricaneForceWind:
		return "hurricane_force_wind"
	case WeatherSevereTypeWaterspout:
		return "waterspout"
	case WeatherSevereTypeSevereThunderstorm:
		return "severe_thunderstorm"
	case WeatherSevereTypeWreckhouseWinds:
		return "wreckhouse_winds"
	case WeatherSevereTypeLesSuetesWind:
		return "les_suetes_wind"
	case WeatherSevereTypeAvalanche:
		return "avalanche"
	case WeatherSevereTypeFlashFlood:
		return "flash_flood"
	case WeatherSevereTypeTropicalStorm:
		return "tropical_storm"
	case WeatherSevereTypeInlandTropicalStorm:
		return "inland_tropical_storm"
	case WeatherSevereTypeBlizzard:
		return "blizzard"
	case WeatherSevereTypeIceStorm:
		return "ice_storm"
	case WeatherSevereTypeFreezingRain:
		return "freezing_rain"
	case WeatherSevereTypeDebrisFlow:
		return "debris_flow"
	case WeatherSevereTypeFlashFreeze:
		return "flash_freeze"
	case WeatherSevereTypeDustStorm:
		return "dust_storm"
	case WeatherSevereTypeHighWind:
		return "high_wind"
	case WeatherSevereTypeWinterStorm:
		return "winter_storm"
	case WeatherSevereTypeHeavyFreezingSpray:
		return "heavy_freezing_spray"
	case WeatherSevereTypeExtremeCold:
		return "extreme_cold"
	case WeatherSevereTypeWindChill:
		return "wind_chill"
	case WeatherSevereTypeColdWave:
		return "cold_wave"
	case WeatherSevereTypeHeavySnowAlert:
		return "heavy_snow_alert"
	case WeatherSevereTypeLakeEffectBlowingSnow:
		return "lake_effect_blowing_snow"
	case WeatherSevereTypeSnowSquall:
		return "snow_squall"
	case WeatherSevereTypeLakeEffectSnow:
		return "lake_effect_snow"
	case WeatherSevereTypeWinterWeather:
		return "winter_weather"
	case WeatherSevereTypeSleet:
		return "sleet"
	case WeatherSevereTypeSnowfall:
		return "snowfall"
	case WeatherSevereTypeSnowAndBlowingSnow:
		return "snow_and_blowing_snow"
	case WeatherSevereTypeBlowingSnow:
		return "blowing_snow"
	case WeatherSevereTypeSnowAlert:
		return "snow_alert"
	case WeatherSevereTypeArcticOutflow:
		return "arctic_outflow"
	case WeatherSevereTypeFreezingDrizzle:
		return "freezing_drizzle"
	case WeatherSevereTypeStorm:
		return "storm"
	case WeatherSevereTypeStormSurge:
		return "storm_surge"
	case WeatherSevereTypeRainfall:
		return "rainfall"
	case WeatherSevereTypeArealFlood:
		return "areal_flood"
	case WeatherSevereTypeCoastalFlood:
		return "coastal_flood"
	case WeatherSevereTypeLakeshoreFlood:
		return "lakeshore_flood"
	case WeatherSevereTypeExcessiveHeat:
		return "excessive_heat"
	case WeatherSevereTypeHeat:
		return "heat"
	case WeatherSevereTypeWeather:
		return "weather"
	case WeatherSevereTypeHighHeatAndHumidity:
		return "high_heat_and_humidity"
	case WeatherSevereTypeHumidexAndHealth:
		return "humidex_and_health"
	case WeatherSevereTypeHumidex:
		return "humidex"
	case WeatherSevereTypeGale:
		return "gale"
	case WeatherSevereTypeFreezingSpray:
		return "freezing_spray"
	case WeatherSevereTypeSpecialMarine:
		return "special_marine"
	case WeatherSevereTypeSquall:
		return "squall"
	case WeatherSevereTypeStrongWind:
		return "strong_wind"
	case WeatherSevereTypeLakeWind:
		return "lake_wind"
	case WeatherSevereTypeMarineWeather:
		return "marine_weather"
	case WeatherSevereTypeWind:
		return "wind"
	case WeatherSevereTypeSmallCraftHazardousSeas:
		return "small_craft_hazardous_seas"
	case WeatherSevereTypeHazardousSeas:
		return "hazardous_seas"
	case WeatherSevereTypeSmallCraft:
		return "small_craft"
	case WeatherSevereTypeSmallCraftWinds:
		return "small_craft_winds"
	case WeatherSevereTypeSmallCraftRoughBar:
		return "small_craft_rough_bar"
	case WeatherSevereTypeHighWaterLevel:
		return "high_water_level"
	case WeatherSevereTypeAshfall:
		return "ashfall"
	case WeatherSevereTypeFreezingFog:
		return "freezing_fog"
	case WeatherSevereTypeDenseFog:
		return "dense_fog"
	case WeatherSevereTypeDenseSmoke:
		return "dense_smoke"
	case WeatherSevereTypeBlowingDust:
		return "blowing_dust"
	case WeatherSevereTypeHardFreeze:
		return "hard_freeze"
	case WeatherSevereTypeFreeze:
		return "freeze"
	case WeatherSevereTypeFrost:
		return "frost"
	case WeatherSevereTypeFireWeather:
		return "fire_weather"
	case WeatherSevereTypeFlood:
		return "flood"
	case WeatherSevereTypeRipTide:
		return "rip_tide"
	case WeatherSevereTypeHighSurf:
		return "high_surf"
	case WeatherSevereTypeSmog:
		return "smog"
	case WeatherSevereTypeAirQuality:
		return "air_quality"
	case WeatherSevereTypeBriskWind:
		return "brisk_wind"
	case WeatherSevereTypeAirStagnation:
		return "air_stagnation"
	case WeatherSevereTypeLowWater:
		return "low_water"
	case WeatherSevereTypeHydrological:
		return "hydrological"
	case WeatherSevereTypeSpecialWeather:
		return "special_weather"
	default:
		return "WeatherSevereTypeInvalid(" + strconv.Itoa(int(w)) + ")"
	}
}

// FromString parse string into WeatherSevereType constant it's represent, return WeatherSevereTypeInvalid if not found.
func WeatherSevereTypeFromString(s string) WeatherSevereType {
	switch s {
	case "unspecified":
		return WeatherSevereTypeUnspecified
	case "tornado":
		return WeatherSevereTypeTornado
	case "tsunami":
		return WeatherSevereTypeTsunami
	case "hurricane":
		return WeatherSevereTypeHurricane
	case "extreme_wind":
		return WeatherSevereTypeExtremeWind
	case "typhoon":
		return WeatherSevereTypeTyphoon
	case "inland_hurricane":
		return WeatherSevereTypeInlandHurricane
	case "hurricane_force_wind":
		return WeatherSevereTypeHurricaneForceWind
	case "waterspout":
		return WeatherSevereTypeWaterspout
	case "severe_thunderstorm":
		return WeatherSevereTypeSevereThunderstorm
	case "wreckhouse_winds":
		return WeatherSevereTypeWreckhouseWinds
	case "les_suetes_wind":
		return WeatherSevereTypeLesSuetesWind
	case "avalanche":
		return WeatherSevereTypeAvalanche
	case "flash_flood":
		return WeatherSevereTypeFlashFlood
	case "tropical_storm":
		return WeatherSevereTypeTropicalStorm
	case "inland_tropical_storm":
		return WeatherSevereTypeInlandTropicalStorm
	case "blizzard":
		return WeatherSevereTypeBlizzard
	case "ice_storm":
		return WeatherSevereTypeIceStorm
	case "freezing_rain":
		return WeatherSevereTypeFreezingRain
	case "debris_flow":
		return WeatherSevereTypeDebrisFlow
	case "flash_freeze":
		return WeatherSevereTypeFlashFreeze
	case "dust_storm":
		return WeatherSevereTypeDustStorm
	case "high_wind":
		return WeatherSevereTypeHighWind
	case "winter_storm":
		return WeatherSevereTypeWinterStorm
	case "heavy_freezing_spray":
		return WeatherSevereTypeHeavyFreezingSpray
	case "extreme_cold":
		return WeatherSevereTypeExtremeCold
	case "wind_chill":
		return WeatherSevereTypeWindChill
	case "cold_wave":
		return WeatherSevereTypeColdWave
	case "heavy_snow_alert":
		return WeatherSevereTypeHeavySnowAlert
	case "lake_effect_blowing_snow":
		return WeatherSevereTypeLakeEffectBlowingSnow
	case "snow_squall":
		return WeatherSevereTypeSnowSquall
	case "lake_effect_snow":
		return WeatherSevereTypeLakeEffectSnow
	case "winter_weather":
		return WeatherSevereTypeWinterWeather
	case "sleet":
		return WeatherSevereTypeSleet
	case "snowfall":
		return WeatherSevereTypeSnowfall
	case "snow_and_blowing_snow":
		return WeatherSevereTypeSnowAndBlowingSnow
	case "blowing_snow":
		return WeatherSevereTypeBlowingSnow
	case "snow_alert":
		return WeatherSevereTypeSnowAlert
	case "arctic_outflow":
		return WeatherSevereTypeArcticOutflow
	case "freezing_drizzle":
		return WeatherSevereTypeFreezingDrizzle
	case "storm":
		return WeatherSevereTypeStorm
	case "storm_surge":
		return WeatherSevereTypeStormSurge
	case "rainfall":
		return WeatherSevereTypeRainfall
	case "areal_flood":
		return WeatherSevereTypeArealFlood
	case "coastal_flood":
		return WeatherSevereTypeCoastalFlood
	case "lakeshore_flood":
		return WeatherSevereTypeLakeshoreFlood
	case "excessive_heat":
		return WeatherSevereTypeExcessiveHeat
	case "heat":
		return WeatherSevereTypeHeat
	case "weather":
		return WeatherSevereTypeWeather
	case "high_heat_and_humidity":
		return WeatherSevereTypeHighHeatAndHumidity
	case "humidex_and_health":
		return WeatherSevereTypeHumidexAndHealth
	case "humidex":
		return WeatherSevereTypeHumidex
	case "gale":
		return WeatherSevereTypeGale
	case "freezing_spray":
		return WeatherSevereTypeFreezingSpray
	case "special_marine":
		return WeatherSevereTypeSpecialMarine
	case "squall":
		return WeatherSevereTypeSquall
	case "strong_wind":
		return WeatherSevereTypeStrongWind
	case "lake_wind":
		return WeatherSevereTypeLakeWind
	case "marine_weather":
		return WeatherSevereTypeMarineWeather
	case "wind":
		return WeatherSevereTypeWind
	case "small_craft_hazardous_seas":
		return WeatherSevereTypeSmallCraftHazardousSeas
	case "hazardous_seas":
		return WeatherSevereTypeHazardousSeas
	case "small_craft":
		return WeatherSevereTypeSmallCraft
	case "small_craft_winds":
		return WeatherSevereTypeSmallCraftWinds
	case "small_craft_rough_bar":
		return WeatherSevereTypeSmallCraftRoughBar
	case "high_water_level":
		return WeatherSevereTypeHighWaterLevel
	case "ashfall":
		return WeatherSevereTypeAshfall
	case "freezing_fog":
		return WeatherSevereTypeFreezingFog
	case "dense_fog":
		return WeatherSevereTypeDenseFog
	case "dense_smoke":
		return WeatherSevereTypeDenseSmoke
	case "blowing_dust":
		return WeatherSevereTypeBlowingDust
	case "hard_freeze":
		return WeatherSevereTypeHardFreeze
	case "freeze":
		return WeatherSevereTypeFreeze
	case "frost":
		return WeatherSevereTypeFrost
	case "fire_weather":
		return WeatherSevereTypeFireWeather
	case "flood":
		return WeatherSevereTypeFlood
	case "rip_tide":
		return WeatherSevereTypeRipTide
	case "high_surf":
		return WeatherSevereTypeHighSurf
	case "smog":
		return WeatherSevereTypeSmog
	case "air_quality":
		return WeatherSevereTypeAirQuality
	case "brisk_wind":
		return WeatherSevereTypeBriskWind
	case "air_stagnation":
		return WeatherSevereTypeAirStagnation
	case "low_water":
		return WeatherSevereTypeLowWater
	case "hydrological":
		return WeatherSevereTypeHydrological
	case "special_weather":
		return WeatherSevereTypeSpecialWeather
	default:
		return WeatherSevereTypeInvalid
	}
}

// List returns all constants.
func ListWeatherSevereType() []WeatherSevereType {
	return []WeatherSevereType{
		WeatherSevereTypeUnspecified,
		WeatherSevereTypeTornado,
		WeatherSevereTypeTsunami,
		WeatherSevereTypeHurricane,
		WeatherSevereTypeExtremeWind,
		WeatherSevereTypeTyphoon,
		WeatherSevereTypeInlandHurricane,
		WeatherSevereTypeHurricaneForceWind,
		WeatherSevereTypeWaterspout,
		WeatherSevereTypeSevereThunderstorm,
		WeatherSevereTypeWreckhouseWinds,
		WeatherSevereTypeLesSuetesWind,
		WeatherSevereTypeAvalanche,
		WeatherSevereTypeFlashFlood,
		WeatherSevereTypeTropicalStorm,
		WeatherSevereTypeInlandTropicalStorm,
		WeatherSevereTypeBlizzard,
		WeatherSevereTypeIceStorm,
		WeatherSevereTypeFreezingRain,
		WeatherSevereTypeDebrisFlow,
		WeatherSevereTypeFlashFreeze,
		WeatherSevereTypeDustStorm,
		WeatherSevereTypeHighWind,
		WeatherSevereTypeWinterStorm,
		WeatherSevereTypeHeavyFreezingSpray,
		WeatherSevereTypeExtremeCold,
		WeatherSevereTypeWindChill,
		WeatherSevereTypeColdWave,
		WeatherSevereTypeHeavySnowAlert,
		WeatherSevereTypeLakeEffectBlowingSnow,
		WeatherSevereTypeSnowSquall,
		WeatherSevereTypeLakeEffectSnow,
		WeatherSevereTypeWinterWeather,
		WeatherSevereTypeSleet,
		WeatherSevereTypeSnowfall,
		WeatherSevereTypeSnowAndBlowingSnow,
		WeatherSevereTypeBlowingSnow,
		WeatherSevereTypeSnowAlert,
		WeatherSevereTypeArcticOutflow,
		WeatherSevereTypeFreezingDrizzle,
		WeatherSevereTypeStorm,
		WeatherSevereTypeStormSurge,
		WeatherSevereTypeRainfall,
		WeatherSevereTypeArealFlood,
		WeatherSevereTypeCoastalFlood,
		WeatherSevereTypeLakeshoreFlood,
		WeatherSevereTypeExcessiveHeat,
		WeatherSevereTypeHeat,
		WeatherSevereTypeWeather,
		WeatherSevereTypeHighHeatAndHumidity,
		WeatherSevereTypeHumidexAndHealth,
		WeatherSevereTypeHumidex,
		WeatherSevereTypeGale,
		WeatherSevereTypeFreezingSpray,
		WeatherSevereTypeSpecialMarine,
		WeatherSevereTypeSquall,
		WeatherSevereTypeStrongWind,
		WeatherSevereTypeLakeWind,
		WeatherSevereTypeMarineWeather,
		WeatherSevereTypeWind,
		WeatherSevereTypeSmallCraftHazardousSeas,
		WeatherSevereTypeHazardousSeas,
		WeatherSevereTypeSmallCraft,
		WeatherSevereTypeSmallCraftWinds,
		WeatherSevereTypeSmallCraftRoughBar,
		WeatherSevereTypeHighWaterLevel,
		WeatherSevereTypeAshfall,
		WeatherSevereTypeFreezingFog,
		WeatherSevereTypeDenseFog,
		WeatherSevereTypeDenseSmoke,
		WeatherSevereTypeBlowingDust,
		WeatherSevereTypeHardFreeze,
		WeatherSevereTypeFreeze,
		WeatherSevereTypeFrost,
		WeatherSevereTypeFireWeather,
		WeatherSevereTypeFlood,
		WeatherSevereTypeRipTide,
		WeatherSevereTypeHighSurf,
		WeatherSevereTypeSmog,
		WeatherSevereTypeAirQuality,
		WeatherSevereTypeBriskWind,
		WeatherSevereTypeAirStagnation,
		WeatherSevereTypeLowWater,
		WeatherSevereTypeHydrological,
		WeatherSevereTypeSpecialWeather,
	}
}
