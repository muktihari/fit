// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

var version = "dev"

type options struct {
	combineOutputPath string

	// list of commands
	concealStart uint32
	concealEnd   uint32
	combine      bool

	encoderHeaderOption encoder.Option
}

func main() {
	var opts options
	outputPathHelpText := "[option] Output of combined files: result.fit"
	flag.StringVar(&opts.combineOutputPath, "o", "", outputPathHelpText)
	flag.StringVar(&opts.combineOutputPath, "out", "", outputPathHelpText)

	flag.BoolVar(&opts.combine, "combine", false, "[command] Combine multiple fit activity files into one continoues fit activity. Example:\n"+
		" 1. fitactivity --combine -o result.fit part1.fit part2.fit\n"+
		" 2. fitactivity --combine part1.fit part2.fit > result.fit",
	)

	var flagConcealStart uint
	flag.UintVar(&flagConcealStart, "conceal-start", 0, "[command] Amount of distance to conceal from the start (in meters). Example:\n"+
		" 1. fitactivity --conceal-start 1000 part1.fit part2.fit\n"+
		" 2. fitactivity --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit\n"+
		" 3. fitactivity --combine --conceal-start 1000 part1.fit part2.fit > result.fit",
	)

	var flagConcealEnd uint
	flag.UintVar(&flagConcealEnd, "conceal-end", 0, "[command] Amount of distance to conceal from the end (in meters). Example:\n"+
		" 1. fitactivity --conceal-end 1000 part1.fit part2.fit\n"+
		" 2. fitactivity --combine --conceal-end 1000 part1.fit part2.fit > result.fit\n"+
		" 3. fitactivity --combine --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit > result.fit",
	)

	var flagInterleave uint
	flag.UintVar(&flagInterleave, "interleave", 15, "[option] Max interleave allowed to reduce writing the same message definition on encoding process. Valid value: 0-15. Example:\n"+
		" 1. fitactivity --combine -o result.fit --interleave 0 part1.fit part2.fit\n *",
	)

	var flagCompress bool
	flag.BoolVar(&flagCompress, "compress", false, "[option] Compress timestamp in message header. Save 7 bytes per message up for every message written in 31s interval. Example:\n"+
		" 1. fitactivity --combine -o result.fit --compress part1.fit part2.fit",
	)

	var flagVersion bool
	flagVersionHelpText := "[command] Show version"
	flag.BoolVar(&flagVersion, "v", false, flagVersionHelpText)
	flag.BoolVar(&flagVersion, "version", false, flagVersionHelpText)

	flag.Parse()

	if flagVersion {
		fmt.Println(version)
		return
	}

	if flagConcealStart > math.MaxUint32 {
		fatalf("max conceal-start value is %d.", math.MaxUint32)
	}
	if flagConcealEnd > math.MaxUint32 {
		fatalf("max conceal-end value is %d.", math.MaxUint32)
	}
	opts.concealStart = uint32(flagConcealStart)
	opts.concealEnd = uint32(flagConcealEnd)

	if !opts.combine && opts.concealStart == 0 && opts.concealEnd == 0 {
		fatalf("no options selected. see --help\n")
	}

	if flagInterleave > 15 {
		fatalf("max message definition interleave is 15, got: %d.", flagInterleave)
	}

	if flagCompress {
		opts.encoderHeaderOption = encoder.WithCompressedTimestampHeader()
	} else {
		opts.encoderHeaderOption = encoder.WithNormalHeader(byte(flagInterleave))
	}

	// Convert meters into record's distance scaled value (scale: 100, offset: 0).
	opts.concealStart, opts.concealEnd = opts.concealStart*100, opts.concealEnd*100

	paths := flag.Args()

	if opts.combine { // Combine (and Conceal Position if specified)
		if err := combineAndConcealPosition(paths, &opts); err != nil {
			fatalf("could not combine: %v\n", err)
		}
		return
	}

	// Conceal Position Only
	for _, path := range paths {
		if err := openAndConcealPosition(path, &opts); err != nil {
			fatalf("could not openAndConcealPosition [path %q]: %v\n", path, err)
		}
	}
}

func combineAndConcealPosition(paths []string, opts *options) error {
	fits, err := opener.Open(paths...)
	if err != nil {
		return fmt.Errorf("could not bulk open: %v", err)
	}

	fit, err := combiner.Combine(fits...)
	if err != nil {
		return fmt.Errorf("could not combine: %v", err)
	}

	if err := concealer.ConcealPosition(fit, opts.concealStart, opts.concealEnd); err != nil {
		return fmt.Errorf("could not conceal: %v", err)
	}

	concealer.ConcealLapStartAndEndPosition(fit)

	var fout = os.Stdout // default output to stdout if not specified.
	if opts.combineOutputPath != "" {
		fout, err = os.OpenFile(opts.combineOutputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
		if err != nil {
			return fmt.Errorf("could not open output file: %v", err)
		}
		defer fout.Close()
	}

	bw := bufferedwriter.New(fout)
	enc := encoder.New(bw,
		encoder.WithProtocolVersion(proto.V2),
		opts.encoderHeaderOption,
	)

	if err := enc.Encode(fit); err != nil {
		return fmt.Errorf("could not encode: %v", err)
	}

	if err := bw.Flush(); err != nil {
		return fmt.Errorf("could not flush: %v", err)
	}

	return nil
}

func openAndConcealPosition(path string, opts *options) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	var fits []*proto.FIT
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
		if err := concealer.ConcealPosition(fits[i], opts.concealStart, opts.concealEnd); err != nil {
			return fmt.Errorf("could not conceal fits[%d of %d]: %v", i+1, len(fits), err)
		}
		concealer.ConcealLapStartAndEndPosition(fits[i])
	}

	ext := filepath.Ext(path)
	name := fmt.Sprintf("%s_concealed_%d_%d", strings.TrimSuffix(path, ext), opts.concealStart, opts.concealEnd)

	fout, err := os.OpenFile(name+ext, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		return fmt.Errorf("could not open output file: %v", err)
	}
	defer fout.Close()

	bw := bufferedwriter.New(fout)
	enc := encoder.New(bw,
		encoder.WithProtocolVersion(proto.V2),
		opts.encoderHeaderOption,
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
