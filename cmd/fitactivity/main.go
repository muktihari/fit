// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/muktihari/fit/cmd/fitactivity/combiner"
	"github.com/muktihari/fit/cmd/fitactivity/concealer"
	"github.com/muktihari/fit/cmd/fitactivity/opener"
	"github.com/muktihari/fit/cmd/fitactivity/reducer"
	"github.com/muktihari/fit/cmd/fitactivity/remover"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	errBadArgument       = errorString("bad argument")
	errPrintUsageAndExit = errorString("print usage and exit")
)

// NOTE: This variables can be changed during build using -ldflags.
// e.g. go build -ldflags="-X 'main.version=$(git describe --tags)'" -o fitactivity main.go
var (
	cli     = "fitactivity"
	version = "dev"
)

const (
	combineDesc = "combine multiple activities into one continuous activity"
	concealDesc = "conceal first or last x meters GPS positions for privacy"
	reduceDesc  = `reduce the size of record messages, available methods:
             1. Based on GPS points using RDP [Ramer-Douglas-Peucker]
             2. Based on distance interval in meters
             3. Based on time interval in seconds`
	removeDesc = "remove messages based on given message numbers and other parameters"

	perm = 0o644
)

var mainUsage = `About:
  ` + cli + ` is a program to manage FIT activity files based on provided command.
	
Usage:
  ` + cli + ` [command] 

Available Commands:
  combine    ` + combineDesc + `
  conceal    ` + concealDesc + `
  reduce     ` + reduceDesc + `
  remove     ` + removeDesc + `

Flags:
  -h, --help       Print help
  -v, --version    Print version
`

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fs := flag.NewFlagSet("main", flag.ExitOnError)
	fs.Usage = func() { fmt.Fprint(os.Stderr, mainUsage) }

	var v bool
	fs.BoolVar(&v, "v", false, "Print version")
	fs.BoolVar(&v, "version", false, "Print version")
	fs.Parse(os.Args[1:])

	if v {
		fmt.Println(version)
		return
	}

	args := fs.Args()
	if len(args) == 0 {
		fs.Usage()
		return
	}

	var begin = time.Now()

	var command = args[0]
	switch command {
	case "combine":
		fs := flag.NewFlagSet(command, flag.ExitOnError)
		printerror(fs, command, combine(ctx, fs, args[1:]))
	case "conceal":
		fs := flag.NewFlagSet(command, flag.ExitOnError)
		printerror(fs, command, conceal(ctx, fs, args[1:]))
	case "reduce":
		fs := flag.NewFlagSet(command, flag.ExitOnError)
		printerror(fs, command, reduce(ctx, fs, args[1:]))
	case "remove":
		fs := flag.NewFlagSet(command, flag.ExitOnError)
		printerror(fs, command, remove(ctx, fs, args[1:]))
	default:
		printerror(fs, command, fmt.Errorf("command provided but not defined: %s", command))
	}

	fmt.Fprintf(os.Stderr, "~ DONE! in %s\n", time.Since(begin))
}

// printerror prints error and exit if err != nil.
func printerror(fs *flag.FlagSet, command string, err error) {
	if err == nil {
		return
	}
	if errors.Is(err, errBadArgument) {
		fmt.Fprintln(os.Stderr, strings.TrimSuffix(err.Error(), ": "+errBadArgument.Error()))
		fs.Usage()
	} else if errors.Is(err, errPrintUsageAndExit) {
		fs.Usage()
	} else if errors.Is(err, context.Canceled) {
		fmt.Fprintf(os.Stderr, "~ %s: %v\n", command, err)
	} else {
		fmt.Fprintf(os.Stderr, "~ %s: error: %v\n", command, err)
	}
	os.Exit(1)
}

const (
	defaultInterleave     = 15
	interleaveDesc        = "max interleave for message definition [valid: 0-15, default: 15]"
	compressDesc          = "compress timestamp into message header [default: false; this overrides interleave]"
	combineOutDesc        = "combine output file"
	concealFirstDesc      = "conceal distance: first x meters"
	concealLastDesc       = "conceal distance: last x meters"
	reduceByPointsRdpDesc = "reduce method: RDP [Ramer-Douglas-Peucker] based on GPS points, epsilon > 0"
	reduceByDistanceDesc  = "reduce method: distance interval in meters"
	reduceByTimeDesc      = "reduce method: time interval in seconds"
	removeUnknownDesc     = "remove unknown messages"
	removeMesgNumsDesc    = "remove message numbers (value separated by comma)"
	removeDevDataDesc     = "remove developer data"
)

var combineUsage = `About:
  ` + combineDesc + `

Usage:
  ` + cli + ` combine [subcommands] [flags] [files]

Available Subcommands (optional):
  conceal    ` + concealDesc + `
  reduce     ` + reduceDesc + `
  remove     ` + removeDesc + `

Flags:
  (required):
  -o, --out  string    ` + combineOutDesc + `

  (optional):
  -i, --interleave  uint8    ` + interleaveDesc + `
  -c, --compress    bool     ` + compressDesc + `

Subcommand Flags (only if subcommand is provided):
  conceal: (select at least one)
   --first  uint32    ` + concealFirstDesc + `
   --last   uint32    ` + concealLastDesc + `

  reduce: (select only one)
   --rdp       float64    ` + reduceByPointsRdpDesc + `
   --distance  float64    ` + reduceByDistanceDesc + `
   --time      uint32     ` + reduceByTimeDesc + `

  remove: (select at least one)
   --unknown   bool       ` + removeUnknownDesc + `
   --nums      string     ` + removeMesgNumsDesc + `
   --devdata   bool       ` + removeDevDataDesc + `

Examples:
  ` + cli + ` combine -o result.fit part1.fit part2.fit
  ` + cli + ` combine reduce -o result.fit --rdp 0.0001 part1.fit part2.fit
  ` + cli + ` combine conceal -o result.fit --first 1000 part1.fit part2.fit
  ` + cli + ` combine remove -o result.fit --unknown --nums 160,164 part1.fit part2.fit
  ` + cli + ` combine conceal reduce -o result.fit --last 1000 --time 5 part1.fit part2.fit
`

func combine(ctx context.Context, fs *flag.FlagSet, args []string) (err error) {
	fs.Usage = func() { fmt.Fprint(os.Stderr, combineUsage) }

	const (
		subcommandConceal = "conceal"
		subcommandReduce  = "reduce"
		subcommandRemove  = "remove"
	)

	var subcommands []string
	var i int
loop:
	for i = range args {
		switch args[i] { // Subcommands
		case subcommandConceal, subcommandReduce, subcommandRemove:
			subcommands = append(subcommands, args[i])
		default:
			if !strings.HasPrefix(args[i], "-") {
				return fmt.Errorf("subcommand provided but not defined: %s: %w", args[i], errBadArgument)
			}
			break loop
		}
	}

	var out string
	fs.StringVar(&out, "o", "", combineOutDesc)
	fs.StringVar(&out, "out", "", combineOutDesc)

	var interleave uint
	fs.UintVar(&interleave, "i", defaultInterleave, interleaveDesc)
	fs.UintVar(&interleave, "interleave", defaultInterleave, interleaveDesc)

	var compress bool
	fs.BoolVar(&compress, "c", false, compressDesc)
	fs.BoolVar(&compress, "compress", false, compressDesc)

	const flagNameRdp = "rdp"
	var reduceByRdp float64
	fs.Float64Var(&reduceByRdp, flagNameRdp, 0, reduceByPointsRdpDesc)

	const flagNameDistance = "distance"
	var reduceByDistance float64
	fs.Float64Var(&reduceByDistance, flagNameDistance, 0, reduceByDistanceDesc)

	const flagNameTime = "time"
	var reduceByTime uint
	fs.UintVar(&reduceByTime, flagNameTime, 0, reduceByTimeDesc)

	const flagNameFirst = "first"
	var concealFirst uint
	fs.UintVar(&concealFirst, flagNameFirst, 0, concealFirstDesc)

	const flagNameLast = "last"
	var concealLast uint
	fs.UintVar(&concealLast, flagNameLast, 0, concealLastDesc)

	const flagNameRemoveUnknown = "unknown"
	var removeUnknown bool
	fs.BoolVar(&removeUnknown, flagNameRemoveUnknown, false, removeUnknownDesc)

	const flagNameRemoveMesgNums = "nums"
	var removeNums string
	fs.StringVar(&removeNums, flagNameRemoveMesgNums, "", removeMesgNumsDesc)

	const flagNameRemoveDevData = "devdata"
	var removeDevData bool
	fs.BoolVar(&removeDevData, flagNameRemoveDevData, false, removeDevDataDesc)

	fs.Parse(args[i:])

	var flagSpecified bool
	fs.Visit(func(f *flag.Flag) { flagSpecified = true })
	if !flagSpecified {
		return errPrintUsageAndExit
	}

	if out == "" {
		return fmt.Errorf("flag is required but not provided: -o or --out: %w", errBadArgument)
	}

	if interleave > 15 {
		return fmt.Errorf("interleave: valid value is between 0 to 15, got: %d: %w", interleave, errBadArgument)
	}

	if subcommandProvided(subcommands, subcommandReduce) {
		if countSelectedFlag(fs, flagNameRdp, flagNameDistance, flagNameTime) != 1 {
			return fmt.Errorf("reduce: please select (only) one method: %w", errBadArgument)
		}
		if reduceByRdp == 0 && reduceByDistance == 0 && reduceByTime == 0 {
			return fmt.Errorf("reduce: input value could not be zero: %w", errBadArgument)
		}
	}

	if subcommandProvided(subcommands, subcommandConceal) && countSelectedFlag(fs, flagNameFirst, flagNameLast) == 0 {
		return fmt.Errorf("conceal: no distance is provided: %w", errBadArgument)
	}

	removeMesgNums := make(map[typedef.MesgNum]struct{})
	if subcommandProvided(subcommands, subcommandRemove) {
		if countSelectedFlag(fs, flagNameRemoveUnknown, flagNameRemoveMesgNums, flagNameRemoveDevData) == 0 {
			return fmt.Errorf("remove: argument is provided: %w", errBadArgument)
		}
		parts := strings.Split(removeNums, ",")
		for _, part := range parts {
			u16, err := strconv.ParseUint(part, 10, 16)
			if err != nil {
				return err
			}
			removeMesgNums[typedef.MesgNum(u16)] = struct{}{}
		}
	}

	files := fs.Args()
	if len(files) < 2 {
		return fmt.Errorf("provide at least 2 valid FIT files to combine: %w", errBadArgument)
	}

	if err = expectdotfit(files); err != nil {
		return err
	}

	var fits []*proto.FIT
	verboserun(fmt.Sprintf("Decoding %d files", len(files)), func() {
		fits, err = opener.Open(ctx, files)
	})
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	var fit *proto.FIT
	verboserun(fmt.Sprintf("Combining %d files", len(files)), func() {
		fit, err = combiner.Combine(fits)
	})
	if err != nil {
		return err
	}

	fit.FileHeader.ProtocolVersion = latestProtocolVersion(fits)
	fit.FileHeader.ProfileVersion = latestProfileVersion(fits)

	for _, subcommand := range subcommands {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		switch subcommand {
		case subcommandConceal:
			msg := fmt.Sprintf("Concealing [start: %sm, end: %sm]",
				formatThousand(int(concealFirst)), formatThousand(int(concealLast)))

			verboserun(msg, func() {
				concealer.Conceal(fit.Messages, uint32(concealFirst)*100, uint32(concealLast)*100)
			})
			if err != nil {
				return err
			}
		case subcommandReduce:
			var option reducer.Option

			var param string
			switch {
			case reduceByRdp > 0:
				option = reducer.WithRDP(reduceByRdp)
				param = fmt.Sprintf("rdp: %g ε", reduceByRdp)
			case reduceByDistance > 0:
				option = reducer.WithDistanceInterval(reduceByDistance)
				param = fmt.Sprintf("distance: %.2fm", reduceByDistance)
			case reduceByTime > 0:
				option = reducer.WithTimeInterval(uint32(reduceByTime))
				param = fmt.Sprintf("time: %ds", reduceByTime)
			}
			prevLen := len(fit.Messages)

			msg := fmt.Sprintf("Reducing [%s]", param)
			verboserun(msg, func() {
				err = reducer.Reduce(fit, option)
			})

			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "  # messages are reduced from %s into %s\n",
				formatThousand(prevLen), formatThousand(len(fit.Messages)))
		case subcommandRemove:
			var opts []remover.Option
			if removeUnknown {
				opts = append(opts, remover.WithRemoveUnknown())
			}
			if len(removeMesgNums) > 0 {
				opts = append(opts, remover.WithRemoveMesgNums(removeMesgNums))
			}
			if removeDevData {
				opts = append(opts, remover.WithRemoveDeveloperData())
			}

			prevLen := len(fit.Messages)

			msg := fmt.Sprintf("Removing [unknown: %t, nums: %s, devdata: %t]",
				removeUnknown, removeNums, removeDevData)
			verboserun(msg, func() {
				remover.Remove(fit, opts...)
			})

			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "  # messages are removed from %s into %s\n",
				formatThousand(prevLen), formatThousand(len(fit.Messages)))
		}
	}

	headerInfo := fmt.Sprintf("interleave: %d", interleave)
	headerOption := encoder.WithNormalHeader(byte(interleave))
	if compress {
		headerInfo = "compress"
		headerOption = encoder.WithCompressedTimestampHeader()
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	msg := fmt.Sprintf("Encoding [%s]", headerInfo)
	verboserun(msg, func() {
		var f *os.File
		f, err = os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
		if err != nil {
			return
		}
		defer f.Close()
		enc := encoder.New(f, headerOption)
		err = enc.EncodeWithContext(ctx, fit)
		if err != nil {
			_ = os.Remove(out)
		}
	})

	return err
}

var concealUsage = `About:
  ` + concealDesc + `

Usage:
  ` + cli + ` conceal [flags] [files]

Flags: 
  (select at least one):
    --start         uint32     ` + concealFirstDesc + `
    --end           uint32     ` + concealLastDesc + `

  (optional):
  -i, --interleave  uint8      ` + interleaveDesc + `
  -c, --compress    bool       ` + compressDesc + `

Examples:
  ` + cli + ` conceal --first 1000 a.fit b.fit
  ` + cli + ` conceal --first 1000 --last 1000 a.fit b.fit
`

func conceal(ctx context.Context, fs *flag.FlagSet, args []string) (err error) {
	fs.Usage = func() { fmt.Fprint(os.Stderr, concealUsage) }

	var interleave int
	fs.IntVar(&interleave, "i", defaultInterleave, interleaveDesc)
	fs.IntVar(&interleave, "interleave", defaultInterleave, interleaveDesc)

	var compress bool
	fs.BoolVar(&compress, "c", false, compressDesc)
	fs.BoolVar(&compress, "compress", false, compressDesc)

	const flagNameFirst = "first"
	var first uint
	fs.UintVar(&first, flagNameFirst, 0, concealFirstDesc)

	const flagNameLast = "last"
	var last uint
	fs.UintVar(&last, flagNameLast, 0, concealLastDesc)

	fs.Parse(args)

	var flagSpecified bool
	fs.Visit(func(f *flag.Flag) { flagSpecified = true })
	if !flagSpecified {
		return errPrintUsageAndExit
	}

	if first == 0 && last == 0 {
		return fmt.Errorf("conceal: no distance is specified: %w", errBadArgument)
	}

	if interleave < 0 || interleave > 15 {
		return fmt.Errorf("interleave: valid value is between 0 to 15, got: %d: %w", interleave, errBadArgument)
	}

	files := fs.Args()
	if len(files) == 0 {
		return fmt.Errorf("provide at least 1 FIT files to conceal: %w", errBadArgument)
	}

	if err := expectdotfit(files); err != nil {
		return err
	}

	headerInfo := fmt.Sprintf("interleave: %d", interleave)
	headerOption := encoder.WithNormalHeader(byte(interleave))
	if compress {
		headerInfo = "compress"
		headerOption = encoder.WithCompressedTimestampHeader()
	}

	fmt.Fprintf(os.Stderr, "- Concealing %d file(s) [first: %s m; last: %s m]\n",
		len(files), formatThousand(int(first)), formatThousand(int(last)))

	var dec = decoder.New(nil)
	var enc = encoder.New(nil)
	for i, path := range files {
		var fits []*proto.FIT
		verboserun(fmt.Sprintf("[%d] Decoding", i), func() {
			var f *os.File
			f, err = os.Open(path)
			if err != nil {
				return
			}
			defer f.Close()

			dec.Reset(f)

			var fit *proto.FIT
			for dec.Next() {
				fit, err = dec.DecodeWithContext(ctx)
				if err != nil {
					return
				}
				fits = append(fits, fit)
			}
		})
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		verboserun(fmt.Sprintf("[%d] Concealing", i), func() {
			for _, fit := range fits {
				concealer.Conceal(fit.Messages, uint32(first)*100, uint32(last)*100)
			}
		})
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		verboserun(fmt.Sprintf("[%d] Encoding [%s]", i, headerInfo), func() {
			name := fmt.Sprintf("%s_concealed_%d_%d.fit",
				strings.TrimSuffix(path, filepath.Ext(path)), first, last)

			var f *os.File
			f, err = os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
			if err != nil {
				return
			}
			defer f.Close()

			enc.Reset(f, headerOption)

			for _, fit := range fits {
				if err = enc.EncodeWithContext(ctx, fit); err != nil {
					_ = os.Remove(name)
					return
				}
			}
		})
		if err != nil {
			return fmt.Errorf("could not conceal %q: %v", path, err)
		}
	}
	return nil
}

func trimSpaceAddPaddingPerLine(s string, paddingLeft int) string {
	var b strings.Builder
	for _, v := range strings.Split(s, "\n") {
		b.WriteString(strings.Repeat(" ", paddingLeft))
		b.WriteString(strings.TrimSpace(v))
		b.WriteString("\n")
	}
	return b.String()
}

var reduceUsage = `About:
` + trimSpaceAddPaddingPerLine(reduceDesc, 2) + `

Usage:
  ` + cli + ` reduce [flags] [files]

Flags:
  (select only one):
    --rdp           float64    ` + reduceByPointsRdpDesc + `
    --distance      float64    ` + reduceByDistanceDesc + `
    --time          uint32     ` + reduceByTimeDesc + `

  (optional):
  -i, --interleave  uint8      ` + interleaveDesc + `
  -c, --compress    bool       ` + compressDesc + `


Examples:
  ` + cli + ` reduce --rdp 0.0001 a.fit b.fit
  ` + cli + ` reduce --distance 0.5 a.fit b.fit
  ` + cli + ` reduce --time 5 a.fit b.fit
`

func reduce(ctx context.Context, fs *flag.FlagSet, args []string) (err error) {
	fs.Usage = func() { fmt.Fprint(os.Stderr, reduceUsage) }

	var interleave int
	fs.IntVar(&interleave, "i", defaultInterleave, interleaveDesc)
	fs.IntVar(&interleave, "interleave", defaultInterleave, interleaveDesc)

	var compress bool
	fs.BoolVar(&compress, "c", false, compressDesc)
	fs.BoolVar(&compress, "compress", false, compressDesc)

	const flagNameRdp = "rdp"
	var reduceByRdp float64
	fs.Float64Var(&reduceByRdp, flagNameRdp, 0, reduceByPointsRdpDesc)

	const flagNameDistance = "distance"
	var reduceByDistance float64
	fs.Float64Var(&reduceByDistance, flagNameDistance, 0, reduceByDistanceDesc)

	const flagNameTime = "time"
	var reduceByTime uint
	fs.UintVar(&reduceByTime, flagNameTime, 0, reduceByTimeDesc)

	fs.Parse(args)

	var flagSpecified bool
	fs.Visit(func(f *flag.Flag) { flagSpecified = true })
	if !flagSpecified {
		return errPrintUsageAndExit
	}

	if countSelectedFlag(fs, flagNameRdp, flagNameDistance, flagNameTime) != 1 {
		return fmt.Errorf("please select (only) one method: %w", errBadArgument)
	}

	if reduceByRdp == 0 && reduceByDistance == 0 && reduceByTime == 0 {
		return fmt.Errorf("input value could not be zero: %w", errBadArgument)
	}

	if interleave < 0 || interleave > 15 {
		return fmt.Errorf("interleave: valid value is between 0 to 15, got: %d: %w", interleave, errBadArgument)
	}

	files := fs.Args()
	if len(files) == 0 {
		return fmt.Errorf("provide at least 1 FIT files to conceal: %w", errBadArgument)
	}

	if err = expectdotfit(files); err != nil {
		return err
	}

	var option reducer.Option
	var nameSuffix string
	var methodInfo string
	switch {
	case reduceByRdp > 0:
		option = reducer.WithRDP(reduceByRdp)
		methodInfo = fmt.Sprintf("rdp ε: %g", reduceByRdp)
		nameSuffix = fmt.Sprintf("rdp_epsilon_%g", reduceByRdp)
	case reduceByDistance > 0:
		option = reducer.WithDistanceInterval(reduceByDistance)
		methodInfo = fmt.Sprintf("distance interval: %sm", formatThousand(int(reduceByDistance)))
		nameSuffix = fmt.Sprintf("distance_interval_%.2fm", reduceByDistance)
	case reduceByTime > 0:
		option = reducer.WithTimeInterval(uint32(reduceByTime))
		methodInfo = fmt.Sprintf("time interval: %ss", formatThousand(int(reduceByTime)))
		nameSuffix = fmt.Sprintf("time_interval_%ds", reduceByTime)
	}

	headerInfo := fmt.Sprintf("interleave: %d", interleave)
	headerOption := encoder.WithNormalHeader(byte(interleave))
	if compress {
		headerInfo = "compress"
		headerOption = encoder.WithCompressedTimestampHeader()
	}

	fmt.Fprintf(os.Stderr, "- Reducing %d file(s) [%s]\n",
		len(files), methodInfo)

	var dec = decoder.New(nil)
	var enc = encoder.New(nil)
	for i, path := range files {
		var fits []*proto.FIT
		verboserun(fmt.Sprintf("[%d] Decoding", i), func() {
			var f *os.File
			f, err = os.Open(path)
			if err != nil {
				return
			}
			defer f.Close()

			dec.Reset(f)

			var fit *proto.FIT
			for dec.Next() {
				fit, err = dec.DecodeWithContext(ctx)
				if err != nil {
					return
				}
				fits = append(fits, fit)
			}
		})
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		var msgs []string
		verboserun(fmt.Sprintf("[%d] Reducing", i), func() {
			for _, fit := range fits {
				prevLen := len(fit.Messages)
				if err = reducer.Reduce(fit, option); err != nil {
					return
				}
				msgs = append(msgs, fmt.Sprintf("# messages are reduced from %s into %s",
					formatThousand(prevLen), formatThousand(len(fit.Messages))))
			}
		})
		if err != nil {
			return err
		}
		for _, msg := range msgs {
			fmt.Fprintf(os.Stderr, "      %s\n", msg)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		verboserun(fmt.Sprintf("[%d] Encoding [%s]", i, headerInfo), func() {
			name := fmt.Sprintf("%s_reduced_%s.fit",
				strings.TrimSuffix(path, filepath.Ext(path)), nameSuffix)

			var f *os.File
			f, err = os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
			if err != nil {
				return
			}
			defer f.Close()

			enc.Reset(f, headerOption)

			for _, fit := range fits {
				if err = enc.EncodeWithContext(ctx, fit); err != nil {
					_ = os.Remove(name)
					return
				}
			}
		})
		if err != nil {
			return err
		}
	}

	return nil
}

var removeUsage = `About:
  ` + removeDesc + `

Usage:
  ` + cli + ` remove [flags] [files]

Flags:
  (select at least one):
    --unknown   bool       ` + removeUnknownDesc + `
    --nums      string     ` + removeMesgNumsDesc + `
    --devdata   bool       ` + removeDevDataDesc + `

  (optional):
  -i, --interleave  uint8      ` + interleaveDesc + `
  -c, --compress    bool       ` + compressDesc + `


Examples:
  ` + cli + ` remove --unknown a.fit b.fit
  ` + cli + ` remove --nums 160,162 a.fit b.fit
  ` + cli + ` remove --devdata a.fit b.fit
  ` + cli + ` remove --unknown --nums 160,162 --devdata a.fit b.fit
`

func remove(ctx context.Context, fs *flag.FlagSet, args []string) (err error) {
	fs.Usage = func() { fmt.Fprint(os.Stderr, removeUsage) }

	var interleave int
	fs.IntVar(&interleave, "i", defaultInterleave, interleaveDesc)
	fs.IntVar(&interleave, "interleave", defaultInterleave, interleaveDesc)

	var compress bool
	fs.BoolVar(&compress, "c", false, compressDesc)
	fs.BoolVar(&compress, "compress", false, compressDesc)

	const flagNameRemoveUnknown = "unknown"
	var removeUnknown bool
	fs.BoolVar(&removeUnknown, flagNameRemoveUnknown, false, removeUnknownDesc)

	const flagNameRemoveMesgNums = "nums"
	var mesgNums string
	fs.StringVar(&mesgNums, flagNameRemoveMesgNums, "", removeMesgNumsDesc)

	const flagNameRemoveDevData = "devdata"
	var removeDevData bool
	fs.BoolVar(&removeDevData, flagNameRemoveDevData, false, removeDevDataDesc)

	fs.Parse(args)

	var flagSpecified bool
	fs.Visit(func(f *flag.Flag) { flagSpecified = true })
	if !flagSpecified {
		return errPrintUsageAndExit
	}

	if countSelectedFlag(fs, flagNameRemoveUnknown, flagNameRemoveMesgNums) == 0 {
		return fmt.Errorf("please select (only) one method: %w", errBadArgument)
	}

	parts := strings.Split(mesgNums, ",")
	removeMesgNums := make(map[typedef.MesgNum]struct{})
	for _, part := range parts {
		u16, err := strconv.ParseUint(part, 10, 16)
		if err != nil {
			return err
		}
		removeMesgNums[typedef.MesgNum(u16)] = struct{}{}
	}

	if interleave < 0 || interleave > 15 {
		return fmt.Errorf("interleave: valid value is between 0 to 15, got: %d: %w", interleave, errBadArgument)
	}

	files := fs.Args()
	if len(files) == 0 {
		return fmt.Errorf("provide at least 1 FIT files to conceal: %w", errBadArgument)
	}

	if err = expectdotfit(files); err != nil {
		return err
	}

	headerInfo := fmt.Sprintf("interleave: %d", interleave)
	headerOption := encoder.WithNormalHeader(byte(interleave))
	if compress {
		headerInfo = "compress"
		headerOption = encoder.WithCompressedTimestampHeader()
	}

	var nameSuffix string
	if removeUnknown {
		nameSuffix = "unknown"
	}
	if mesgNums != "" {
		nameSuffix = fmt.Sprintf("%s_%s",
			nameSuffix, strings.ReplaceAll(mesgNums, ",", "_"))
	}

	fmt.Fprintf(os.Stderr, "- Removing %d file(s) [unknown: %t, nums: %s, devdata: %t]\n",
		len(files), removeUnknown, mesgNums, removeDevData)

	var dec = decoder.New(nil)
	var enc = encoder.New(nil)
	for i, path := range files {
		var fits []*proto.FIT
		verboserun(fmt.Sprintf("[%d] Decoding", i), func() {
			var f *os.File
			f, err = os.Open(path)
			if err != nil {
				return
			}
			defer f.Close()

			dec.Reset(f)

			var fit *proto.FIT
			for dec.Next() {
				fit, err = dec.DecodeWithContext(ctx)
				if err != nil {
					return
				}
				fits = append(fits, fit)
			}
		})
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		var msgs []string
		verboserun(fmt.Sprintf("[%d] Removing", i), func() {
			for _, fit := range fits {
				var opts []remover.Option
				if removeUnknown {
					opts = append(opts, remover.WithRemoveUnknown())
				}
				if len(removeMesgNums) > 0 {
					opts = append(opts, remover.WithRemoveMesgNums(removeMesgNums))
				}
				if removeDevData {
					opts = append(opts, remover.WithRemoveDeveloperData())
				}

				prevLen := len(fit.Messages)

				remover.Remove(fit, opts...)

				msgs = append(msgs, fmt.Sprintf("# messages are removed from %s into %s",
					formatThousand(prevLen), formatThousand(len(fit.Messages))))
			}
		})
		if err != nil {
			return err
		}
		for _, msg := range msgs {
			fmt.Fprintf(os.Stderr, "      %s\n", msg)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		verboserun(fmt.Sprintf("[%d] Encoding [%s]", i, headerInfo), func() {
			name := fmt.Sprintf("%s_removed_%s.fit",
				strings.TrimSuffix(path, filepath.Ext(path)), nameSuffix)

			var f *os.File
			f, err = os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
			if err != nil {
				return
			}
			defer f.Close()

			enc.Reset(f, headerOption)

			for _, fit := range fits {
				if err = enc.EncodeWithContext(ctx, fit); err != nil {
					_ = os.Remove(name)
					return
				}
			}
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// verboserun wraps and runs fn with printing msg and elapsed time of the running process.
func verboserun(msg string, fn func()) {
	begin := time.Now()
	fmt.Fprintf(os.Stderr, "- %-52s ", msg)
	fn()
	fmt.Fprintf(os.Stderr, "[took: %s]\n", time.Since(begin))
}

func countSelectedFlag(fs *flag.FlagSet, flagNames ...string) (count int) {
	for _, name := range flagNames {
		fs.Visit(func(f *flag.Flag) {
			if f.Name == name {
				count++
			}
		})
	}
	return
}

func subcommandProvided(subcommands []string, target string) bool {
	for _, subcommand := range subcommands {
		if target == subcommand {
			return true
		}
	}
	return false
}

func expectdotfit(paths []string) error {
	for _, path := range paths {
		ext := filepath.Ext(path)
		if strings.ToLower(ext) != ".fit" {
			return fmt.Errorf("%q expected file extension: .fit, got: %s: %w", path, ext, errBadArgument)
		}
	}
	return nil
}

func formatThousand(v int) string {
	s := strconv.Itoa(v)
	var result strings.Builder

	n := len(s)
	for i, digit := range s {
		if (n-i)%3 == 0 && i != 0 {
			result.WriteRune('.')
		}
		result.WriteRune(digit)
	}

	return result.String()
}

func latestProtocolVersion(fits []*proto.FIT) proto.Version {
	var version = proto.V1
	for i := range fits {
		if fits[i].FileHeader.ProtocolVersion > version {
			version = fits[i].FileHeader.ProtocolVersion
		}
	}
	return version
}

func latestProfileVersion(fits []*proto.FIT) uint16 {
	var version uint16
	for i := range fits {
		if fits[i].FileHeader.ProfileVersion > version {
			version = fits[i].FileHeader.ProfileVersion
		}
	}
	if version == 0 {
		return profile.Version
	}
	return version
}
