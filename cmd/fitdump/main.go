// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/proto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing arguments\n")
		os.Exit(2)
	}
	for i, arg := range os.Args[1:] {
		if err := textdump(arg); err != nil {
			fmt.Printf("could not dump args[%d] %q: %v\n", i, arg, err)
		}
	}
}

func textdump(path string) error {
	ext := filepath.Ext(path)
	if strings.ToLower(ext) != ".fit" {
		return fmt.Errorf("expected ext: *.fit, got: %s", ext)
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	base := filepath.Base(path)
	dir := filepath.Dir(path)

	name := base
	if len(ext) < len(base) {
		name = base[:len(base)-len(ext)]
	}
	name = name + ".txt"

	resultPath := filepath.Join(dir, name)
	out, err := os.OpenFile(resultPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	defer w.Flush()

	dec := decoder.NewRaw()

	fmt.Fprintf(w, "%s %16s %7s %8s\n", "SEGMENT", "|LOCAL NUM|", "HEADER", "BYTES")
	_, err = dec.Decode(bufio.NewReader(f), func(flag decoder.RawFlag, b []byte) error {
		if flag == decoder.RawFlagMesgDef || flag == decoder.RawFlagMesgData {
			fmt.Fprintf(w, "%-19s |%2d|  %08b  %v\n", flag, proto.LocalMesgNum(b[0]), b[0], b)
		} else {
			fmt.Fprintf(w, "%-19s %15s %v\n", flag, "", b)
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Printf("FIT dumped: %q -> %q\n", path, resultPath)

	return nil
}
