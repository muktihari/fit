// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ProjectileType byte

const (
	ProjectileTypeArrow           ProjectileType = 0 // Arrow projectile type
	ProjectileTypeRifleCartridge  ProjectileType = 1 // Rifle cartridge projectile type
	ProjectileTypePistolCartridge ProjectileType = 2 // Pistol cartridge projectile type
	ProjectileTypeShotshell       ProjectileType = 3 // Shotshell projectile type
	ProjectileTypeAirRiflePellet  ProjectileType = 4 // Air rifle pellet projectile type
	ProjectileTypeOther           ProjectileType = 5 // Other projectile type
	ProjectileTypeInvalid         ProjectileType = 0xFF
)

func (p ProjectileType) Byte() byte { return byte(p) }

func (p ProjectileType) String() string {
	switch p {
	case ProjectileTypeArrow:
		return "arrow"
	case ProjectileTypeRifleCartridge:
		return "rifle_cartridge"
	case ProjectileTypePistolCartridge:
		return "pistol_cartridge"
	case ProjectileTypeShotshell:
		return "shotshell"
	case ProjectileTypeAirRiflePellet:
		return "air_rifle_pellet"
	case ProjectileTypeOther:
		return "other"
	default:
		return "ProjectileTypeInvalid(" + strconv.Itoa(int(p)) + ")"
	}
}

// FromString parse string into ProjectileType constant it's represent, return ProjectileTypeInvalid if not found.
func ProjectileTypeFromString(s string) ProjectileType {
	switch s {
	case "arrow":
		return ProjectileTypeArrow
	case "rifle_cartridge":
		return ProjectileTypeRifleCartridge
	case "pistol_cartridge":
		return ProjectileTypePistolCartridge
	case "shotshell":
		return ProjectileTypeShotshell
	case "air_rifle_pellet":
		return ProjectileTypeAirRiflePellet
	case "other":
		return ProjectileTypeOther
	default:
		return ProjectileTypeInvalid
	}
}

// List returns all constants.
func ListProjectileType() []ProjectileType {
	return []ProjectileType{
		ProjectileTypeArrow,
		ProjectileTypeRifleCartridge,
		ProjectileTypePistolCartridge,
		ProjectileTypeShotshell,
		ProjectileTypeAirRiflePellet,
		ProjectileTypeOther,
	}
}
