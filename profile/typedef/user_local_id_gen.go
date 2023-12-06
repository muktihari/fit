// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.128

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type UserLocalId uint16

const (
	UserLocalIdLocalMin      UserLocalId = 0x0000
	UserLocalIdLocalMax      UserLocalId = 0x000F
	UserLocalIdStationaryMin UserLocalId = 0x0010
	UserLocalIdStationaryMax UserLocalId = 0x00FF
	UserLocalIdPortableMin   UserLocalId = 0x0100
	UserLocalIdPortableMax   UserLocalId = 0xFFFE
	UserLocalIdInvalid       UserLocalId = 0xFFFF // INVALID
)

var userlocalidtostrs = map[UserLocalId]string{
	UserLocalIdLocalMin:      "local_min",
	UserLocalIdLocalMax:      "local_max",
	UserLocalIdStationaryMin: "stationary_min",
	UserLocalIdStationaryMax: "stationary_max",
	UserLocalIdPortableMin:   "portable_min",
	UserLocalIdPortableMax:   "portable_max",
	UserLocalIdInvalid:       "invalid",
}

func (u UserLocalId) String() string {
	val, ok := userlocalidtostrs[u]
	if !ok {
		return strconv.FormatUint(uint64(u), 10)
	}
	return val
}

var strtouserlocalid = func() map[string]UserLocalId {
	m := make(map[string]UserLocalId)
	for t, str := range userlocalidtostrs {
		m[str] = UserLocalId(t)
	}
	return m
}()

// FromString parse string into UserLocalId constant it's represent, return UserLocalIdInvalid if not found.
func UserLocalIdFromString(s string) UserLocalId {
	val, ok := strtouserlocalid[s]
	if !ok {
		return strtouserlocalid["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListUserLocalId() []UserLocalId {
	vs := make([]UserLocalId, 0, len(userlocalidtostrs))
	for i := range userlocalidtostrs {
		vs = append(vs, UserLocalId(i))
	}
	return vs
}
