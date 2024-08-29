// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

import "github.com/muktihari/fit/internal/cmd/fitgen/shared"

// ProfileData is data representative of profile.tmpl
type ProfileData struct {
	ConstantData                  shared.ConstantData
	MappingProfileTypeToBaseTypes []ProfileTypeBaseType
	MappingBaseTypeToProfileTypes []ProfileTypeBaseType
}

// ProfileTypeBaseType is mapping struct from ProfileType to its BaseType or vise versa.
type ProfileTypeBaseType struct {
	ProfileType string
	BaseType    string
}

// VersionData is data representative of version.tmpl
type VersionData struct {
	ProfileVersion string
	Package        string
	Major          uint16
	Minor          uint16
	Version        uint16
}
