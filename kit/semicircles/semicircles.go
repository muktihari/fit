// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semicircles

const (
	piRadians        = 1 << 31 // 2^31; 31 bit representation
	conversionFactor = 180.0 / piRadians
)

func ToDegrees(semicircles int32) float64 {
	return float64(semicircles) * conversionFactor
}

func ToSemicircles(degrees float64) int32 {
	return int32(degrees / conversionFactor)
}
