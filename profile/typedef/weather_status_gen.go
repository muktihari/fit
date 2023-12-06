// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

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
	WeatherStatusInvalid                WeatherStatus = 0xFF // INVALID
)

var weatherstatustostrs = map[WeatherStatus]string{
	WeatherStatusClear:                  "clear",
	WeatherStatusPartlyCloudy:           "partly_cloudy",
	WeatherStatusMostlyCloudy:           "mostly_cloudy",
	WeatherStatusRain:                   "rain",
	WeatherStatusSnow:                   "snow",
	WeatherStatusWindy:                  "windy",
	WeatherStatusThunderstorms:          "thunderstorms",
	WeatherStatusWintryMix:              "wintry_mix",
	WeatherStatusFog:                    "fog",
	WeatherStatusHazy:                   "hazy",
	WeatherStatusHail:                   "hail",
	WeatherStatusScatteredShowers:       "scattered_showers",
	WeatherStatusScatteredThunderstorms: "scattered_thunderstorms",
	WeatherStatusUnknownPrecipitation:   "unknown_precipitation",
	WeatherStatusLightRain:              "light_rain",
	WeatherStatusHeavyRain:              "heavy_rain",
	WeatherStatusLightSnow:              "light_snow",
	WeatherStatusHeavySnow:              "heavy_snow",
	WeatherStatusLightRainSnow:          "light_rain_snow",
	WeatherStatusHeavyRainSnow:          "heavy_rain_snow",
	WeatherStatusCloudy:                 "cloudy",
	WeatherStatusInvalid:                "invalid",
}

func (w WeatherStatus) String() string {
	val, ok := weatherstatustostrs[w]
	if !ok {
		return strconv.Itoa(int(w))
	}
	return val
}

var strtoweatherstatus = func() map[string]WeatherStatus {
	m := make(map[string]WeatherStatus)
	for t, str := range weatherstatustostrs {
		m[str] = WeatherStatus(t)
	}
	return m
}()

// FromString parse string into WeatherStatus constant it's represent, return WeatherStatusInvalid if not found.
func WeatherStatusFromString(s string) WeatherStatus {
	val, ok := strtoweatherstatus[s]
	if !ok {
		return strtoweatherstatus["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWeatherStatus() []WeatherStatus {
	vs := make([]WeatherStatus, 0, len(weatherstatustostrs))
	for i := range weatherstatustostrs {
		vs = append(vs, WeatherStatus(i))
	}
	return vs
}
