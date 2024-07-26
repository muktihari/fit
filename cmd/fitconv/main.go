// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run fitcsv/lookup.go

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/muktihari/fit/cmd/fitconv/fitcsv"
	"github.com/muktihari/fit/decoder"
)

var version = "dev"

const blockSize = 8 << 10

func main() {
	var opt string
	if version != "dev" {
		flag.StringVar(&opt, "opt", "", "")
	}

	var flagVersion bool
	flag.BoolVar(&flagVersion, "v", false, "Show version")

	var flagUseDisk bool
	flag.BoolVar(&flagUseDisk, "disk", false, "Use disk instead of load everything in memory")

	var flagPrintUnknownMesgNum bool
	flag.BoolVar(&flagPrintUnknownMesgNum, "unknown", false, "Print unknown mesg num e.g. 'unknown(68)' instead of 'unknown'")

	var flagPrintOnlyValidValue bool
	flag.BoolVar(&flagPrintOnlyValidValue, "valid", false, "Print only valid value")

	var flagPrintRawValue bool
	flag.BoolVar(&flagPrintRawValue, "raw", false, "Use raw value instead of scaled value")

	var flagPrintDegrees bool
	flag.BoolVar(&flagPrintDegrees, "deg", false, "Print GPS position (Lat & Long) in degrees instead of semicircles")

	var flagTrimTrailingCommas bool
	flag.BoolVar(&flagTrimTrailingCommas, "trim", false, "Trim trailing commas in every line")

	var flagNoExpandComponents bool
	flag.BoolVar(&flagNoExpandComponents, "no-expand", false, "[Decode Option] Do not expand components")

	var flagNoChecksum bool
	flag.BoolVar(&flagNoChecksum, "no-checksum", false, "[Decode Option] should not do crc checksum")

	flag.Parse()

	if flagVersion {
		fmt.Println(version)
		return
	}

	/*
		// For Debugging
		switch opt {
		case "cpu":
			defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
		case "mem":
			defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
		case "clock":
			defer profile.Start(profile.ClockProfile, profile.ProfilePath(".")).Stop()
		case "trace":
			defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
		case "threadcreation":
			defer profile.Start(profile.ThreadcreationProfile, profile.ProfilePath(".")).Stop()
		case "block":
			defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
		case "took":
			defer func(begin time.Time) {
				fmt.Printf("took: %s\n", time.Since(begin))
			}(time.Now())
		}
	*/

	var fitToCsvOptions []fitcsv.Option
	if flagUseDisk {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithUseDisk(blockSize))
	}
	if flagPrintRawValue {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithPrintRawValue())
	}
	if flagPrintUnknownMesgNum {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithPrintUnknownMesgNum())
	}
	if flagPrintDegrees {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithPrintGPSPositionInDegrees())
	}
	if flagPrintOnlyValidValue {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithPrintOnlyValidValue())
	}
	if flagTrimTrailingCommas {
		fitToCsvOptions = append(fitToCsvOptions, fitcsv.WithTrimTrailingCommas())
	}

	var decoderOptions []decoder.Option
	if flagNoExpandComponents {
		decoderOptions = append(decoderOptions, decoder.WithNoComponentExpansion())
	}
	if flagNoChecksum {
		decoderOptions = append(decoderOptions, decoder.WithIgnoreChecksum())
	}

	paths := flag.Args()

	if len(paths) == 0 {
		fatalf("missing file argument, e.g.: fitconv Activity.fit\n")
	}

	for _, path := range paths {
		ext := filepath.Ext(path)
		switch ext {
		case ".fit":
			if err := fitToCsv(path, decoderOptions, fitToCsvOptions...); err != nil {
				fmt.Fprintf(os.Stderr, "could not convert %q to csv: %v\n", path, err)
			}
		case ".csv":
			if err := csvToFit(path); err != nil {
				fmt.Fprintf(os.Stderr, "could not convert %q to fit: %v\n", path, err)
			}
		default:
			fmt.Fprintf(os.Stderr, "unrecognized format: %s\n", ext)
		}
	}
}

func fitToCsv(path string, decoderOptions []decoder.Option, opts ...fitcsv.Option) error {
	ff, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", path, err)
	}
	defer ff.Close()

	base := filepath.Base(path)
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)

	name := base
	if len(ext) < len(base) {
		name = base[:len(base)-len(ext)]
	}

	namecsv := name + ".csv"
	pathcsv := filepath.Join(dir, namecsv)

	cf, err := os.OpenFile(pathcsv, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", pathcsv, err)
	}
	defer cf.Close()

	bw := bufio.NewWriterSize(cf, blockSize)
	defer bw.Flush()

	conv := fitcsv.NewFITToCSVConv(bw, opts...)

	options := []decoder.Option{
		decoder.WithMesgDefListener(conv),
		decoder.WithMesgListener(conv),
		decoder.WithBroadcastOnly(),
		decoder.WithBroadcastMesgCopy(),
	}
	options = append(options, decoderOptions...)
	dec := decoder.New(ff, options...)

	for dec.Next() {
		_, err = dec.Decode()
		if err != nil {
			break
		}
	}

	conv.Wait()

	if err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	if err := conv.Err(); err != nil {
		return fmt.Errorf("convert done with error: %v", err)
	}

	fmt.Printf("📄 %q -> %q\n", filepath.Join(dir, path), filepath.Join(dir, namecsv))

	return nil
}

func csvToFit(path string) error {
	cf, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", path, err)
	}
	defer cf.Close()

	base := filepath.Base(path)
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)

	name := base
	if len(ext) < len(base) {
		name = base[:len(base)-len(ext)]
	}

	namefit := name + ".fit"
	pathfit := filepath.Join(dir, namefit)

	ff, err := os.OpenFile(pathfit, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer ff.Close()

	conv := fitcsv.NewCSVToFITConv(ff, bufio.NewReaderSize(cf, blockSize))
	if err := conv.Convert(); err != nil {
		return err
	}

	resultInfo := conv.ResultInfo()
	skipped := []string{}

	if resultInfo.UnknownMesg != 0 {
		skipped = append(skipped, fmt.Sprintf("%d unknown messages", resultInfo.UnknownMesg))
	}
	if resultInfo.UnknownField != 0 {
		skipped = append(skipped, fmt.Sprintf("%d unknown fields", resultInfo.UnknownField))
	}
	if resultInfo.UnknownDynamicField != 0 {
		skipped = append(skipped, fmt.Sprintf("%d unknown dynamic fields", resultInfo.UnknownDynamicField))
	}

	var info string
	if len(skipped) > 0 {
		info = fmt.Sprintf(" [Info: %s are skipped]", strings.Join(skipped, ", "))
	}

	fmt.Printf("🚀 %q -> %q.%s\n", filepath.Join(dir, path), filepath.Join(dir, namefit), info)

	return nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
