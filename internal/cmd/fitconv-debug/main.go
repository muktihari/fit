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
	"time"

	"github.com/muktihari/fit/cmd/fitconv/csv"
	"github.com/muktihari/fit/decoder"
	"github.com/pkg/profile"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage Example: \n" +
			" $ fitconv file1.fit\n" +
			" $ fitconv file1.fit file2.fit\n",
		)
	}

	var opt string
	flag.StringVar(&opt, "opt", "", "")

	flag.Parse()
	paths := flag.Args()

	if len(paths) == 0 {
		fatalf("missing file argument, e.g.: fitconv Activity.fit\n")
	}

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

	for _, path := range paths {
		if err := convertCSV(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not convert %q to csv: %v\n", path, err)
		}
	}
}

func convertCSV(path string) error {
	ff, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", path, err)
	}
	defer ff.Close()

	base := filepath.Base(path)
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)

	if ext == ".csv" {
		return fmt.Errorf("expected *.fit, got *.csv")
	}

	name := base
	if len(ext) < len(base) {
		name = base[:len(base)-len(ext)]
	}

	namecsv := name + ".csv"
	pathcsv := filepath.Join(dir, namecsv)

	sequenceCompleted := 0
	defer func() {
		if sequenceCompleted == 0 {
			_ = os.Remove(pathcsv)
		}
	}()

	cf, err := os.OpenFile(pathcsv, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("could not open file: %s: %w", pathcsv, err)
	}
	defer cf.Close()

	const bsize = 1000 << 10 // 1 MB
	bw := bufio.NewWriterSize(cf, bsize)
	csvconv := csv.NewConverter(bw)

	dec := decoder.New(bufio.NewReaderSize(ff, bsize),
		decoder.WithMesgDefListener(csvconv),
		decoder.WithMesgListener(csvconv),
		decoder.WithBroadcastOnly(),
	)

	for {
		_, err = dec.Decode()
		if err != nil {
			return fmt.Errorf("decode failed: %w", err)
		}
		sequenceCompleted++
		if !dec.Next() {
			break
		}
	}

	csvconv.Wait()

	if err := bw.Flush(); err != nil {
		return fmt.Errorf("could not flush buffered data: %w", err)
	}

	fmt.Printf("Converted! %s\n", filepath.Join(dir, namecsv))

	return nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
