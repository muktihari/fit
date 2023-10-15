// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
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

	"github.com/muktihari/fit/cmd/fitactivity/combiner"
	"github.com/muktihari/fit/cmd/fitactivity/concealer"
	"github.com/muktihari/fit/cmd/fitactivity/opener"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/kit/bufferedwriter"
	"github.com/muktihari/fit/proto"
)

func main() {
	var flagOut string
	flag.StringVar(&flagOut, "o", "", "output of combined files: result.fit")

	var flagCombine bool
	flag.BoolVar(&flagCombine, "combine", false, "[option] Combine multiple fit activity files into one continoues fit activity. Example:\n"+
		" 1. fitactivity --combine -o result.fit part1.fit part2.fit\n"+
		" 2. fitactivity --combine part1.fit part2.fit > result.fit",
	)

	var flagConcealStart, flagConcealEnd uint
	flag.UintVar(&flagConcealStart, "conceal-start", 0, "[option] Amount of distance to conceal from the start (in meters). Example:\n"+
		" 1. fitactivity --conceal-start 1000 part1.fit part2.fit\n"+
		" 2. fitactivity --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit\n"+
		" 3. fitactivity --combine --conceal-start 1000 part1.fit part2.fit > result.fit",
	)
	flag.UintVar(&flagConcealEnd, "conceal-end", 0, "[option] Amount of distance to conceal from the end (in meters). Example:\n"+
		" 1. fitactivity --conceal-end 1000 part1.fit part2.fit\n"+
		" 2. fitactivity --combine --conceal-end 1000 part1.fit part2.fit > result.fit\n"+
		" 3. fitactivity --combine --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit > result.fit",
	)

	flag.Parse()

	paths := flag.Args()

	if !flagCombine && flagConcealStart == 0 && flagConcealEnd == 0 {
		fatalf("no options selected. see --help\n")
	}

	// Convert meters into record's distance scaled value (scale: 100, offset: 0).
	flagConcealStart, flagConcealEnd = flagConcealStart*100, flagConcealEnd*100

	if flagCombine { // Combine (and Conceal Position if specified)
		if err := combineAndConcealPosition(paths, flagOut, uint32(flagConcealStart), uint32(flagConcealEnd)); err != nil {
			fatalf("could not combine: %v\n", err)
		}
		return
	}

	// Conceal Position Only
	for _, path := range paths {
		if err := openAndConcealPosition(path, uint32(flagConcealStart), uint32(flagConcealEnd)); err != nil {
			fatalf("could not openAndConcealPosition [path %q]: %v\n", path, err)
		}
	}
}

func combineAndConcealPosition(paths []string, out string, concealStart, concealEnd uint32) error {
	fits, err := opener.Open(paths...)
	if err != nil {
		return fmt.Errorf("could not bulk open: %v", err)
	}

	fit, err := combiner.Combine(fits...)
	if err != nil {
		return fmt.Errorf("could not combine: %v", err)
	}

	if err := concealer.ConcealPosition(fit, uint32(concealStart), uint32(concealEnd)); err != nil {
		return fmt.Errorf("could not conceal: %v", err)
	}

	concealer.ConcealLapStartAndEndPosition(fit)

	var fout = os.Stdout // default output to stdout if not specified.
	if out != "" {
		fout, err = os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
		if err != nil {
			return fmt.Errorf("could not open output file: %v", err)
		}
		defer fout.Close()
	}

	bw := bufferedwriter.New(fout)
	enc := encoder.New(bw,
		encoder.WithProtocolVersion(proto.V2),
	)

	if err := enc.Encode(fit); err != nil {
		return fmt.Errorf("could not encode: %v", err)
	}

	if err := bw.Flush(); err != nil {
		return fmt.Errorf("could not flush: %v", err)
	}

	return nil
}

func openAndConcealPosition(path string, concealStart, concealEnd uint32) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	var fits []*proto.Fit
	dec := decoder.New(bufio.NewReader(f),
		decoder.WithNoComponentExpansion(),
	)

	for dec.Next() {
		fit, err := dec.Decode()
		if err != nil {
			return fmt.Errorf("could not decode: %v", err)
		}
		fits = append(fits, fit)
	}

	for i := range fits {
		if err := concealer.ConcealPosition(fits[i], concealStart, concealEnd); err != nil {
			return fmt.Errorf("could not conceal fits[%d of %d]: %v", i+1, len(fits), err)
		}
		concealer.ConcealLapStartAndEndPosition(fits[i])
	}

	ext := filepath.Ext(path)
	name := fmt.Sprintf("%s_concealed_%d_%d", strings.TrimSuffix(path, ext), concealStart, concealEnd)

	fout, err := os.OpenFile(name+ext, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		return fmt.Errorf("could not open output file: %v", err)
	}
	defer fout.Close()

	bw := bufferedwriter.New(fout)
	enc := encoder.New(bw,
		encoder.WithProtocolVersion(proto.V2),
	)

	for _, fit := range fits {
		if err := enc.Encode(fit); err != nil {
			return fmt.Errorf("could not encode: %v", err)
		}
	}

	return bw.Flush()
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
