// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.116

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WeatherReport byte

const (
	WeatherReportCurrent WeatherReport = 0
	// WeatherReportForecast WeatherReport = 1  // [DUPLICATE!] Deprecated use hourly_forecast instead
	WeatherReportHourlyForecast WeatherReport = 1
	WeatherReportDailyForecast  WeatherReport = 2
	WeatherReportInvalid        WeatherReport = 0xFF // INVALID
)

var weatherreporttostrs = map[WeatherReport]string{
	WeatherReportCurrent: "current",
	// WeatherReportForecast: "forecast",
	WeatherReportHourlyForecast: "hourly_forecast",
	WeatherReportDailyForecast:  "daily_forecast",
	WeatherReportInvalid:        "invalid",
}

func (w WeatherReport) String() string {
	val, ok := weatherreporttostrs[w]
	if !ok {
		return strconv.Itoa(int(w))
	}
	return val
}

var strtoweatherreport = func() map[string]WeatherReport {
	m := make(map[string]WeatherReport)
	for t, str := range weatherreporttostrs {
		m[str] = WeatherReport(t)
	}
	return m
}()

// FromString parse string into WeatherReport constant it's represent, return WeatherReportInvalid if not found.
func WeatherReportFromString(s string) WeatherReport {
	val, ok := strtoweatherreport[s]
	if !ok {
		return strtoweatherreport["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListWeatherReport() []WeatherReport {
	vs := make([]WeatherReport, 0, len(weatherreporttostrs))
	for i := range weatherreporttostrs {
		vs = append(vs, WeatherReport(i))
	}
	return vs
}
