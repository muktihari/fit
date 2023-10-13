// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type FaveroProduct uint16

const (
	FaveroProductAssiomaUno FaveroProduct = 10
	FaveroProductAssiomaDuo FaveroProduct = 12
	FaveroProductInvalid    FaveroProduct = 0xFFFF // INVALID
)

var faveroproducttostrs = map[FaveroProduct]string{
	FaveroProductAssiomaUno: "assioma_uno",
	FaveroProductAssiomaDuo: "assioma_duo",
	FaveroProductInvalid:    "invalid",
}

func (f FaveroProduct) String() string {
	val, ok := faveroproducttostrs[f]
	if !ok {
		return strconv.FormatUint(uint64(f), 10)
	}
	return val
}

var strtofaveroproduct = func() map[string]FaveroProduct {
	m := make(map[string]FaveroProduct)
	for t, str := range faveroproducttostrs {
		m[str] = FaveroProduct(t)
	}
	return m
}()

// FromString parse string into FaveroProduct constant it's represent, return FaveroProductInvalid if not found.
func FaveroProductFromString(s string) FaveroProduct {
	val, ok := strtofaveroproduct[s]
	if !ok {
		return strtofaveroproduct["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListFaveroProduct() []FaveroProduct {
	vs := make([]FaveroProduct, 0, len(faveroproducttostrs))
	for i := range faveroproducttostrs {
		vs = append(vs, FaveroProduct(i))
	}
	return vs
}