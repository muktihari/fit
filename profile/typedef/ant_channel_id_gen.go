// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AntChannelId uint32

const (
	AntChannelIdAntExtendedDeviceNumberUpperNibble AntChannelId = 0xF0000000
	AntChannelIdAntTransmissionTypeLowerNibble     AntChannelId = 0x0F000000
	AntChannelIdAntDeviceType                      AntChannelId = 0x00FF0000
	AntChannelIdAntDeviceNumber                    AntChannelId = 0x0000FFFF
	AntChannelIdInvalid                            AntChannelId = 0x0
)

func (a AntChannelId) String() string {
	switch a {
	case AntChannelIdAntExtendedDeviceNumberUpperNibble:
		return "ant_extended_device_number_upper_nibble"
	case AntChannelIdAntTransmissionTypeLowerNibble:
		return "ant_transmission_type_lower_nibble"
	case AntChannelIdAntDeviceType:
		return "ant_device_type"
	case AntChannelIdAntDeviceNumber:
		return "ant_device_number"
	default:
		return "AntChannelIdInvalid(" + strconv.FormatUint(uint64(a), 10) + ")"
	}
}

// FromString parse string into AntChannelId constant it's represent, return AntChannelIdInvalid if not found.
func AntChannelIdFromString(s string) AntChannelId {
	switch s {
	case "ant_extended_device_number_upper_nibble":
		return AntChannelIdAntExtendedDeviceNumberUpperNibble
	case "ant_transmission_type_lower_nibble":
		return AntChannelIdAntTransmissionTypeLowerNibble
	case "ant_device_type":
		return AntChannelIdAntDeviceType
	case "ant_device_number":
		return AntChannelIdAntDeviceNumber
	default:
		return AntChannelIdInvalid
	}
}

// List returns all constants.
func ListAntChannelId() []AntChannelId {
	return []AntChannelId{
		AntChannelIdAntExtendedDeviceNumberUpperNibble,
		AntChannelIdAntTransmissionTypeLowerNibble,
		AntChannelIdAntDeviceType,
		AntChannelIdAntDeviceNumber,
	}
}
