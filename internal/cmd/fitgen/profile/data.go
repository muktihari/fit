// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

import "github.com/muktihari/fit/internal/cmd/fitgen/builder/shared"

// ProfileData is data representative of profile.tmpl
type ProfileData struct {
	ConstantData     shared.ConstantData
	MappingBaseTypes []MappingBaseType
}

// MappingBaseType is mapping struct from ProfileType to its BaseType
type MappingBaseType struct {
	ConstantName string
	BaseType     string
}

// VersionData is data representative of version.tmpl
type VersionData struct {
	SDKVersion     string
	Package        string
	ProfileVersion string
}
