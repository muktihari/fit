// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flagutil

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func Usage() {
	fmt.Printf("Usage of %s [options]:\nOptions:\n", os.Args[0])

	var group = make(map[string][]string)
	flag.VisitAll(func(f *flag.Flag) {
		valueType, _ := flag.UnquoteUsage(f)
		key := fmt.Sprintf("%s|%s|%s", f.DefValue, valueType, f.Usage)
		group[key] = append(group[key], f.Name)
	})

	var items = make([]string, len(group))
	for key, flagNames := range group {
		vs := strings.Split(key, "|")
		valueType := vs[1]
		var defaultValue string
		if !(vs[0] == "" || valueType == "") {
			defaultValue = fmt.Sprintf("(default: %s)", vs[0])
		}

		items = append(items, fmt.Sprintf("  -%s %s \n\t %s %s\n",
			strings.Join(flagNames, ", --"), valueType, vs[2], defaultValue))
	}

	slices.Sort(items)
	fmt.Print(strings.Join(items, ""))
}
