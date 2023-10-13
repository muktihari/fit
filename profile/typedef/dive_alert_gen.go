// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DiveAlert byte

const (
	DiveAlertNdlReached                DiveAlert = 0
	DiveAlertGasSwitchPrompted         DiveAlert = 1
	DiveAlertNearSurface               DiveAlert = 2
	DiveAlertApproachingNdl            DiveAlert = 3
	DiveAlertPo2Warn                   DiveAlert = 4
	DiveAlertPo2CritHigh               DiveAlert = 5
	DiveAlertPo2CritLow                DiveAlert = 6
	DiveAlertTimeAlert                 DiveAlert = 7
	DiveAlertDepthAlert                DiveAlert = 8
	DiveAlertDecoCeilingBroken         DiveAlert = 9
	DiveAlertDecoComplete              DiveAlert = 10
	DiveAlertSafetyStopBroken          DiveAlert = 11
	DiveAlertSafetyStopComplete        DiveAlert = 12
	DiveAlertCnsWarning                DiveAlert = 13
	DiveAlertCnsCritical               DiveAlert = 14
	DiveAlertOtuWarning                DiveAlert = 15
	DiveAlertOtuCritical               DiveAlert = 16
	DiveAlertAscentCritical            DiveAlert = 17
	DiveAlertAlertDismissedByKey       DiveAlert = 18
	DiveAlertAlertDismissedByTimeout   DiveAlert = 19
	DiveAlertBatteryLow                DiveAlert = 20
	DiveAlertBatteryCritical           DiveAlert = 21
	DiveAlertSafetyStopStarted         DiveAlert = 22
	DiveAlertApproachingFirstDecoStop  DiveAlert = 23
	DiveAlertSetpointSwitchAutoLow     DiveAlert = 24
	DiveAlertSetpointSwitchAutoHigh    DiveAlert = 25
	DiveAlertSetpointSwitchManualLow   DiveAlert = 26
	DiveAlertSetpointSwitchManualHigh  DiveAlert = 27
	DiveAlertAutoSetpointSwitchIgnored DiveAlert = 28
	DiveAlertSwitchedToOpenCircuit     DiveAlert = 29
	DiveAlertSwitchedToClosedCircuit   DiveAlert = 30
	DiveAlertTankBatteryLow            DiveAlert = 32
	DiveAlertPo2CcrDilLow              DiveAlert = 33   // ccr diluent has low po2
	DiveAlertDecoStopCleared           DiveAlert = 34   // a deco stop has been cleared
	DiveAlertApneaNeutralBuoyancy      DiveAlert = 35   // Target Depth Apnea Alarm triggered
	DiveAlertApneaTargetDepth          DiveAlert = 36   // Neutral Buoyance Apnea Alarm triggered
	DiveAlertApneaSurface              DiveAlert = 37   // Surface Apnea Alarm triggered
	DiveAlertApneaHighSpeed            DiveAlert = 38   // High Speed Apnea Alarm triggered
	DiveAlertApneaLowSpeed             DiveAlert = 39   // Low Speed Apnea Alarm triggered
	DiveAlertInvalid                   DiveAlert = 0xFF // INVALID
)

var divealerttostrs = map[DiveAlert]string{
	DiveAlertNdlReached:                "ndl_reached",
	DiveAlertGasSwitchPrompted:         "gas_switch_prompted",
	DiveAlertNearSurface:               "near_surface",
	DiveAlertApproachingNdl:            "approaching_ndl",
	DiveAlertPo2Warn:                   "po2_warn",
	DiveAlertPo2CritHigh:               "po2_crit_high",
	DiveAlertPo2CritLow:                "po2_crit_low",
	DiveAlertTimeAlert:                 "time_alert",
	DiveAlertDepthAlert:                "depth_alert",
	DiveAlertDecoCeilingBroken:         "deco_ceiling_broken",
	DiveAlertDecoComplete:              "deco_complete",
	DiveAlertSafetyStopBroken:          "safety_stop_broken",
	DiveAlertSafetyStopComplete:        "safety_stop_complete",
	DiveAlertCnsWarning:                "cns_warning",
	DiveAlertCnsCritical:               "cns_critical",
	DiveAlertOtuWarning:                "otu_warning",
	DiveAlertOtuCritical:               "otu_critical",
	DiveAlertAscentCritical:            "ascent_critical",
	DiveAlertAlertDismissedByKey:       "alert_dismissed_by_key",
	DiveAlertAlertDismissedByTimeout:   "alert_dismissed_by_timeout",
	DiveAlertBatteryLow:                "battery_low",
	DiveAlertBatteryCritical:           "battery_critical",
	DiveAlertSafetyStopStarted:         "safety_stop_started",
	DiveAlertApproachingFirstDecoStop:  "approaching_first_deco_stop",
	DiveAlertSetpointSwitchAutoLow:     "setpoint_switch_auto_low",
	DiveAlertSetpointSwitchAutoHigh:    "setpoint_switch_auto_high",
	DiveAlertSetpointSwitchManualLow:   "setpoint_switch_manual_low",
	DiveAlertSetpointSwitchManualHigh:  "setpoint_switch_manual_high",
	DiveAlertAutoSetpointSwitchIgnored: "auto_setpoint_switch_ignored",
	DiveAlertSwitchedToOpenCircuit:     "switched_to_open_circuit",
	DiveAlertSwitchedToClosedCircuit:   "switched_to_closed_circuit",
	DiveAlertTankBatteryLow:            "tank_battery_low",
	DiveAlertPo2CcrDilLow:              "po2_ccr_dil_low",
	DiveAlertDecoStopCleared:           "deco_stop_cleared",
	DiveAlertApneaNeutralBuoyancy:      "apnea_neutral_buoyancy",
	DiveAlertApneaTargetDepth:          "apnea_target_depth",
	DiveAlertApneaSurface:              "apnea_surface",
	DiveAlertApneaHighSpeed:            "apnea_high_speed",
	DiveAlertApneaLowSpeed:             "apnea_low_speed",
	DiveAlertInvalid:                   "invalid",
}

func (d DiveAlert) String() string {
	val, ok := divealerttostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodivealert = func() map[string]DiveAlert {
	m := make(map[string]DiveAlert)
	for t, str := range divealerttostrs {
		m[str] = DiveAlert(t)
	}
	return m
}()

// FromString parse string into DiveAlert constant it's represent, return DiveAlertInvalid if not found.
func DiveAlertFromString(s string) DiveAlert {
	val, ok := strtodivealert[s]
	if !ok {
		return strtodivealert["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDiveAlert() []DiveAlert {
	vs := make([]DiveAlert, 0, len(divealerttostrs))
	for i := range divealerttostrs {
		vs = append(vs, DiveAlert(i))
	}
	return vs
}