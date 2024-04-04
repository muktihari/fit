// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ExdDisplayType byte

const (
	ExdDisplayTypeNumerical         ExdDisplayType = 0
	ExdDisplayTypeSimple            ExdDisplayType = 1
	ExdDisplayTypeGraph             ExdDisplayType = 2
	ExdDisplayTypeBar               ExdDisplayType = 3
	ExdDisplayTypeCircleGraph       ExdDisplayType = 4
	ExdDisplayTypeVirtualPartner    ExdDisplayType = 5
	ExdDisplayTypeBalance           ExdDisplayType = 6
	ExdDisplayTypeStringList        ExdDisplayType = 7
	ExdDisplayTypeString            ExdDisplayType = 8
	ExdDisplayTypeSimpleDynamicIcon ExdDisplayType = 9
	ExdDisplayTypeGauge             ExdDisplayType = 10
	ExdDisplayTypeInvalid           ExdDisplayType = 0xFF
)

func (e ExdDisplayType) Byte() byte { return byte(e) }

func (e ExdDisplayType) String() string {
	switch e {
	case ExdDisplayTypeNumerical:
		return "numerical"
	case ExdDisplayTypeSimple:
		return "simple"
	case ExdDisplayTypeGraph:
		return "graph"
	case ExdDisplayTypeBar:
		return "bar"
	case ExdDisplayTypeCircleGraph:
		return "circle_graph"
	case ExdDisplayTypeVirtualPartner:
		return "virtual_partner"
	case ExdDisplayTypeBalance:
		return "balance"
	case ExdDisplayTypeStringList:
		return "string_list"
	case ExdDisplayTypeString:
		return "string"
	case ExdDisplayTypeSimpleDynamicIcon:
		return "simple_dynamic_icon"
	case ExdDisplayTypeGauge:
		return "gauge"
	default:
		return "ExdDisplayTypeInvalid(" + strconv.Itoa(int(e)) + ")"
	}
}

// FromString parse string into ExdDisplayType constant it's represent, return ExdDisplayTypeInvalid if not found.
func ExdDisplayTypeFromString(s string) ExdDisplayType {
	switch s {
	case "numerical":
		return ExdDisplayTypeNumerical
	case "simple":
		return ExdDisplayTypeSimple
	case "graph":
		return ExdDisplayTypeGraph
	case "bar":
		return ExdDisplayTypeBar
	case "circle_graph":
		return ExdDisplayTypeCircleGraph
	case "virtual_partner":
		return ExdDisplayTypeVirtualPartner
	case "balance":
		return ExdDisplayTypeBalance
	case "string_list":
		return ExdDisplayTypeStringList
	case "string":
		return ExdDisplayTypeString
	case "simple_dynamic_icon":
		return ExdDisplayTypeSimpleDynamicIcon
	case "gauge":
		return ExdDisplayTypeGauge
	default:
		return ExdDisplayTypeInvalid
	}
}

// List returns all constants.
func ListExdDisplayType() []ExdDisplayType {
	return []ExdDisplayType{
		ExdDisplayTypeNumerical,
		ExdDisplayTypeSimple,
		ExdDisplayTypeGraph,
		ExdDisplayTypeBar,
		ExdDisplayTypeCircleGraph,
		ExdDisplayTypeVirtualPartner,
		ExdDisplayTypeBalance,
		ExdDisplayTypeStringList,
		ExdDisplayTypeString,
		ExdDisplayTypeSimpleDynamicIcon,
		ExdDisplayTypeGauge,
	}
}
