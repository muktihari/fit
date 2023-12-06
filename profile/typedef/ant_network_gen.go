// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.117

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AntNetwork byte

const (
	AntNetworkPublic  AntNetwork = 0
	AntNetworkAntplus AntNetwork = 1
	AntNetworkAntfs   AntNetwork = 2
	AntNetworkPrivate AntNetwork = 3
	AntNetworkInvalid AntNetwork = 0xFF // INVALID
)

var antnetworktostrs = map[AntNetwork]string{
	AntNetworkPublic:  "public",
	AntNetworkAntplus: "antplus",
	AntNetworkAntfs:   "antfs",
	AntNetworkPrivate: "private",
	AntNetworkInvalid: "invalid",
}

func (a AntNetwork) String() string {
	val, ok := antnetworktostrs[a]
	if !ok {
		return strconv.Itoa(int(a))
	}
	return val
}

var strtoantnetwork = func() map[string]AntNetwork {
	m := make(map[string]AntNetwork)
	for t, str := range antnetworktostrs {
		m[str] = AntNetwork(t)
	}
	return m
}()

// FromString parse string into AntNetwork constant it's represent, return AntNetworkInvalid if not found.
func AntNetworkFromString(s string) AntNetwork {
	val, ok := strtoantnetwork[s]
	if !ok {
		return strtoantnetwork["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListAntNetwork() []AntNetwork {
	vs := make([]AntNetwork, 0, len(antnetworktostrs))
	for i := range antnetworktostrs {
		vs = append(vs, AntNetwork(i))
	}
	return vs
}
