// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/muktihari/fit/cmd/fitconv/csv"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/kit/bufferedwriter"
	"github.com/pkg/profile"
)

func main() {
	var opt string
	flag.StringVar(&opt, "opt", "", "")

	var path string
	flag.StringVar(&path, "path", "", "path/to/file.fit")
	flag.Parse()

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

	if path == "" {
		fmt.Println("missing path to file, e.g. ./fitconv --path Activity.fit")
		os.Exit(1)
	}

	ff, err := os.Open(path)
	if err != nil {
		fmt.Printf("could not open file: %s: %v\n", path, err)
		return
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

	cf, err := os.OpenFile(filepath.Join(dir, namecsv), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Printf("could not open file: %s: %v\n", filepath.Join(dir, namecsv), err)
		return
	}
	defer cf.Close()

	bw := bufferedwriter.NewSize(cf, 1000<<10)
	csvconv := csv.NewConverter(bw)
	dec := decoder.New(bufio.NewReaderSize(ff, 1000<<10),
		decoder.WithMesgDefListener(csvconv),
		decoder.WithMesgListener(csvconv),
		decoder.WithBroadcastOnly(),
	)

	_, err = dec.Decode(context.Background())
	if err != nil {
		fmt.Printf("decode failed: %v\n", err)
		return
	}

	csvconv.Wait()
	bw.Flush()

	fmt.Printf("Converted! %s\n", filepath.Join(dir, namecsv))
}
