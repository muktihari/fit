// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WeatherStatus byte

const (
	WeatherStatusClear                  WeatherStatus = 0
	WeatherStatusPartlyCloudy           WeatherStatus = 1
	WeatherStatusMostlyCloudy           WeatherStatus = 2
	WeatherStatusRain                   WeatherStatus = 3
	WeatherStatusSnow                   WeatherStatus = 4
	WeatherStatusWindy                  WeatherStatus = 5
	WeatherStatusThunderstorms          WeatherStatus = 6
	WeatherStatusWintryMix              WeatherStatus = 7
	WeatherStatusFog                    WeatherStatus = 8
	WeatherStatusHazy                   WeatherStatus = 11
	WeatherStatusHail                   WeatherStatus = 12
	WeatherStatusScatteredShowers       WeatherStatus = 13
	WeatherStatusScatteredThunderstorms WeatherStatus = 14
	WeatherStatusUnknownPrecipitation   WeatherStatus = 15
	WeatherStatusLightRain              WeatherStatus = 16
	WeatherStatusHeavyRain              WeatherStatus = 17
	WeatherStatusLightSnow              WeatherStatus = 18
	WeatherStatusHeavySnow              WeatherStatus = 19
	WeatherStatusLightRainSnow          WeatherStatus = 20
	WeatherStatusHeavyRainSnow          WeatherStatus = 21
	WeatherStatusCloudy                 WeatherStatus = 22
	WeatherStatusInvalid                WeatherStatus = 0xFF
)

func (w WeatherStatus) Byte() byte { return byte(w) }

func (w WeatherStatus) String() string {
	switch w {
	case WeatherStatusClear:
		return "clear"
	case WeatherStatusPartlyCloudy:
		return "partly_cloudy"
	case WeatherStatusMostlyCloudy:
		return "mostly_cloudy"
	case WeatherStatusRain:
		return "rain"
	case WeatherStatusSnow:
		return "snow"
	case WeatherStatusWindy:
		return "windy"
	case WeatherStatusThunderstorms:
		return "thunderstorms"
	case WeatherStatusWintryMix:
		return "wintry_mix"
	case WeatherStatusFog:
		return "fog"
	case WeatherStatusHazy:
		return "hazy"
	case WeatherStatusHail:
		return "hail"
	case WeatherStatusScatteredShowers:
		return "scattered_showers"
	case WeatherStatusScatteredThunderstorms:
		return "scattered_thunderstorms"
	case WeatherStatusUnknownPrecipitation:
		return "unknown_precipitation"
	case WeatherStatusLightRain:
		return "light_rain"
	case WeatherStatusHeavyRain:
		return "heavy_rain"
	case WeatherStatusLightSnow:
		return "light_snow"
	case WeatherStatusHeavySnow:
		return "heavy_snow"
	case WeatherStatusLightRainSnow:
		return "light_rain_snow"
	case WeatherStatusHeavyRainSnow:
		return "heavy_rain_snow"
	case WeatherStatusCloudy:
		return "cloudy"
	default:
		return "WeatherStatusInvalid(" + strconv.Itoa(int(w)) + ")"
	}
}

// FromString parse string into WeatherStatus constant it's represent, return WeatherStatusInvalid if not found.
func WeatherStatusFromString(s string) WeatherStatus {
	switch s {
	case "clear":
		return WeatherStatusClear
	case "partly_cloudy":
		return WeatherStatusPartlyCloudy
	case "mostly_cloudy":
		return WeatherStatusMostlyCloudy
	case "rain":
		return WeatherStatusRain
	case "snow":
		return WeatherStatusSnow
	case "windy":
		return WeatherStatusWindy
	case "thunderstorms":
		return WeatherStatusThunderstorms
	case "wintry_mix":
		return WeatherStatusWintryMix
	case "fog":
		return WeatherStatusFog
	case "hazy":
		return WeatherStatusHazy
	case "hail":
		return WeatherStatusHail
	case "scattered_showers":
		return WeatherStatusScatteredShowers
	case "scattered_thunderstorms":
		return WeatherStatusScatteredThunderstorms
	case "unknown_precipitation":
		return WeatherStatusUnknownPrecipitation
	case "light_rain":
		return WeatherStatusLightRain
	case "heavy_rain":
		return WeatherStatusHeavyRain
	case "light_snow":
		return WeatherStatusLightSnow
	case "heavy_snow":
		return WeatherStatusHeavySnow
	case "light_rain_snow":
		return WeatherStatusLightRainSnow
	case "heavy_rain_snow":
		return WeatherStatusHeavyRainSnow
	case "cloudy":
		return WeatherStatusCloudy
	default:
		return WeatherStatusInvalid
	}
}

// List returns all constants.
func ListWeatherStatus() []WeatherStatus {
	return []WeatherStatus{
		WeatherStatusClear,
		WeatherStatusPartlyCloudy,
		WeatherStatusMostlyCloudy,
		WeatherStatusRain,
		WeatherStatusSnow,
		WeatherStatusWindy,
		WeatherStatusThunderstorms,
		WeatherStatusWintryMix,
		WeatherStatusFog,
		WeatherStatusHazy,
		WeatherStatusHail,
		WeatherStatusScatteredShowers,
		WeatherStatusScatteredThunderstorms,
		WeatherStatusUnknownPrecipitation,
		WeatherStatusLightRain,
		WeatherStatusHeavyRain,
		WeatherStatusLightSnow,
		WeatherStatusHeavySnow,
		WeatherStatusLightRainSnow,
		WeatherStatusHeavyRainSnow,
		WeatherStatusCloudy,
	}
}
