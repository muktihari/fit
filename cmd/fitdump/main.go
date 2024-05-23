// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/proto"
)

func main() {
	var opt options
	flag.BoolVar(&opt.hex, "hex", false, "print bytes in hexadecimal")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "missing arguments\n")
		os.Exit(2)
	}
	for i, arg := range args {
		if err := textdump(arg, &opt); err != nil {
			fmt.Printf("could not dump args[%d] %q: %v\n", i, arg, err)
		}
	}
}

type options struct {
	hex bool
}

func textdump(path string, opt *options) error {
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
	name = name + "-fitdump.txt"

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
			fmt.Fprintf(w, "%-19s |%s|  %08b  %s\n",
				flag, formatByte(proto.LocalMesgNum(b[0]), opt.hex), b[0], formatBytes(b, opt.hex))
		} else {
			fmt.Fprintf(w, "%-19s %15s %s\n", flag, "", formatBytes(b, opt.hex))
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Printf("FIT dumped: %q -> %q\n", path, resultPath)

	return nil
}

func formatByte(b byte, hexadecimal bool) string {
	if !hexadecimal {
		return fmt.Sprintf("%.2d", b)
	}
	return fmt.Sprintf("%.2x", b)
}

func formatBytes(b []byte, hexadecimal bool) string {
	if !hexadecimal {
		return fmt.Sprintf("%v", b)
	}

	var buf strings.Builder
	buf.WriteByte('[')
	for ; len(b) > 0; b = b[1:] {
		buf.WriteString(fmt.Sprintf("%.2x", b[0]))
		if len(b) > 1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte(']')

	return buf.String()
}
