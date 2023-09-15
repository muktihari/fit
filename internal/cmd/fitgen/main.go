// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
	"github.com/muktihari/fit/internal/cmd/fitgen/factory"
	"github.com/muktihari/fit/internal/cmd/fitgen/generator"
	"github.com/muktihari/fit/internal/cmd/fitgen/parser"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/flagutil"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/strutil"
	"github.com/muktihari/fit/internal/cmd/fitgen/pkg/xlsxlite"
	"github.com/muktihari/fit/internal/cmd/fitgen/profile"
	"github.com/muktihari/fit/internal/cmd/fitgen/profile/mesgdef"
	"github.com/muktihari/fit/internal/cmd/fitgen/profile/typedef"
	"github.com/muktihari/fit/internal/cmd/fitgen/profile/untyped/fieldnum"
	"github.com/muktihari/fit/internal/cmd/fitgen/profile/untyped/mesgnum"

	"github.com/thedatashed/xlsxreader"
)

var aboutFitgen = `
The Fit SDK Generator in Go, also known as "fitgen", is a program designed to create 
several *.go files using Garmin SDK specifications (Profile.xlsx). These generated files
enable this Fit SDK to carry out the decoding and encoding process of Fit files.

The files are organized into distinct packages: 
 - profile: mesgdef, typedef, untyped
 - factory

To define your manufacturer specifications, duplicate the Profile.xlsx file and 
incorporate your specifications within it. Afterward, utilize the provided command-line 
interface (CLI) to generate a customized SDK. When executing the CLI command, specify 
the path to the edited-file such as "Profile-copy.xlsx" using the "--path" option.

Example: 
 - "./fitgen --profile-file Profile-copy.xlsx --path ../../ --builders all -v -y"
 - "./fitgen -f Profile-copy.xlsx -p ../../ -b all -v -y"

Note: The existing Garmin SDK specifications must not be altered, since it might 
result in data that does not align with the terms and conditions of the Fit Protocol.
`

func main() {
	var verbose, verboseHelp = false, "Print list of generated files to stdout"
	flag.BoolVar(&verbose, "v", false, verboseHelp)
	flag.BoolVar(&verbose, "verbose", false, verboseHelp)

	var profileFilePath, profileFilePathHelp = "", "Path/to/Profile.xlsx"
	flag.StringVar(&profileFilePath, "f", "", profileFilePathHelp)
	flag.StringVar(&profileFilePath, "profile-file", "", profileFilePathHelp)

	var generatePath, generatePathHelp = "", "Root path to generate files (e.g. \"../../\")"
	flag.StringVar(&generatePath, "p", "", generatePathHelp)
	flag.StringVar(&generatePath, "path", "", generatePathHelp)

	var whichBuilder string
	whichBuilderHelp := "Which builder to generate (separated by comma): [types, mesgs, profile, factory] or all"
	flag.StringVar(&whichBuilder, "b", "", whichBuilderHelp)
	flag.StringVar(&whichBuilder, "builders", "", whichBuilderHelp)

	var sdkVersion, sdkVersionHelp = "", "Garmin Fit SDK Version (e.g. \"21.115\")"
	flag.StringVar(&sdkVersion, "sdk-version", "", sdkVersionHelp)

	var confirm, confirmHelp = false, "Confirm action"
	flag.BoolVar(&confirm, "y", false, confirmHelp)

	var about bool
	flag.BoolVar(&about, "about", false, "Show about fitgen CLI description then exit program.")

	flag.Usage = flagutil.Usage
	flag.Parse()

	if about {
		fmt.Println(aboutFitgen)
		os.Exit(0)
	}

	if profileFilePath == "" {
		fatalf("missing flag: --profile-file=Path/to/Profile.xlsx\n")
	}

	if generatePath == "" {
		fatalf("missing flag: --path=root/path/to/generate/files\n")
	}

	if sdkVersion == "" {
		fatalf("missing flag: --sdk-version=<version> e.g 21.115\n")
	}

	generatePath = abspath(generatePath)

	if !confirm {
		fmt.Printf("fitgen will generate files relative to given path: %q.\nDo you want to continue? [y/n]: ", generatePath)
		var confimation string
		fmt.Scanf("%s\n", &confimation)
		if strings.ToLower(confimation)[0] != 'y' {
			fatalf("aborted.\n")
		}
	}

	xlsxreader, err := xlsxreader.OpenFile(profileFilePath)
	if err != nil {
		fatalf("could not open Profile.xlsx: %v\n", err)
	}
	defer xlsxreader.Close()

	ps := parser.New(xlsxlite.New(xlsxreader), map[parser.Sheet]string{
		parser.SheetTypes:    "Types",    // maps the actual sheet name in the file
		parser.SheetMessages: "Messages", // to the one that the parser is using.
	})

	parsedtypes, err := ps.ParseTypes()
	if err != nil {
		fatalf(fmt.Sprintf("could no parse types: %v\n", err))
	}
	parsedmesgs, err := ps.ParseMessages()
	if err != nil {
		fatalf(fmt.Sprintf("could no parse message: %v\n", err))
	}

	var (
		typedefb = typedef.NewBuilder(generatePath, sdkVersion, parsedtypes)
		profileb = profile.NewBuilder(generatePath, sdkVersion, parsedtypes)
		factoryb = factory.NewBuilder(generatePath, sdkVersion, parsedtypes, parsedmesgs)
		mesgnumb = mesgnum.NewBuilder(generatePath, sdkVersion, parsedtypes)
		fielnumb = fieldnum.NewBuilder(generatePath, sdkVersion, parsedmesgs, parsedtypes)
		mesgdefb = mesgdef.NewBuilder(generatePath, sdkVersion, parsedmesgs, parsedtypes)
	)

	var builders []builder.Builder
	whichBuilders := strings.Split(strutil.TrimRepeatedChar(whichBuilder, ','), ",")

loop:
	for _, selected := range whichBuilders {
		switch s := strings.TrimSpace(selected); s {
		case "", "all":
			builders = []builder.Builder{typedefb, profileb, factoryb, mesgnumb, fielnumb, mesgdefb}
			break loop
		case "typedef":
			builders = append(builders, typedefb)
		case "profile":
			builders = append(builders, profileb)
		case "factory":
			builders = append(builders, factoryb)
		case "untyped":
			builders = append(builders, fielnumb, mesgnumb)
		case "mesgnum":
			builders = append(builders, mesgnumb)
		case "fieldnum":
			builders = append(builders, fielnumb)
		case "mesgdef":
			builders = append(builders, mesgdefb)
		default:
			fatalf("invalid builder named %q\n", strings.TrimSpace(selected))
		}
	}

	if err := generator.New(verbose).Generate(builders, 0o755); err != nil {
		fatalf("could not generate files: %v\n", err)
	}
}

func fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func abspath(path string) string {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abspath
}
