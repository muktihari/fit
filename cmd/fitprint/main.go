// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run printer/typedef.go

package main

import (
	"fmt"
	"os"

	"github.com/muktihari/fit/cmd/fitprint/printer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing arguments\n")
		os.Exit(2)
	}
	for i, path := range os.Args[1:] {
		if err := printer.Print(path); err != nil {
			fmt.Fprintf(os.Stderr, "path[%d] %q: %v\n", i, path, err)
		}
	}
}
