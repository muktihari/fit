// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.127

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type PwrZoneCalc byte

const (
	PwrZoneCalcCustom     PwrZoneCalc = 0
	PwrZoneCalcPercentFtp PwrZoneCalc = 1
	PwrZoneCalcInvalid    PwrZoneCalc = 0xFF // INVALID
)

var pwrzonecalctostrs = map[PwrZoneCalc]string{
	PwrZoneCalcCustom:     "custom",
	PwrZoneCalcPercentFtp: "percent_ftp",
	PwrZoneCalcInvalid:    "invalid",
}

func (p PwrZoneCalc) String() string {
	val, ok := pwrzonecalctostrs[p]
	if !ok {
		return strconv.Itoa(int(p))
	}
	return val
}

var strtopwrzonecalc = func() map[string]PwrZoneCalc {
	m := make(map[string]PwrZoneCalc)
	for t, str := range pwrzonecalctostrs {
		m[str] = PwrZoneCalc(t)
	}
	return m
}()

// FromString parse string into PwrZoneCalc constant it's represent, return PwrZoneCalcInvalid if not found.
func PwrZoneCalcFromString(s string) PwrZoneCalc {
	val, ok := strtopwrzonecalc[s]
	if !ok {
		return strtopwrzonecalc["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListPwrZoneCalc() []PwrZoneCalc {
	vs := make([]PwrZoneCalc, 0, len(pwrzonecalctostrs))
	for i := range pwrzonecalctostrs {
		vs = append(vs, PwrZoneCalc(i))
	}
	return vs
}
