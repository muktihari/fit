// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/muktihari/fit/internal/cmd/fitgen/builder"
)

// New return a generator
func New(verbose bool) *Generator { return &Generator{verbose: verbose} }

type Generator struct {
	verbose bool
}

func (g *Generator) Generate(builders []builder.Builder, perm os.FileMode) error {
	for _, builder := range builders {
		dataBuilders, err := builder.Build()
		if err != nil {
			return err
		}

		paths := make(map[string]struct{})
		for _, dataBuilder := range dataBuilders {
			paths[dataBuilder.Path] = struct{}{}
			if err := g.GenerateTemplateData(dataBuilder, perm); err != nil {
				return err
			}
		}

		for path := range paths { // format all files in the path
			path = abspath(path)
			stderr := new(bytes.Buffer)
			cmd := exec.Command("gofmt", "-s", "-w", path)
			cmd.Stderr = stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("gofmt -s -w %q, stderr: \n%s\n", path, stderr.String())
				return fmt.Errorf("cmd err: %w", err)
			}
		}
	}

	return nil
}

func (g *Generator) GenerateTemplateData(dataBuilder builder.Data, perm os.FileMode) error {
	if err := os.MkdirAll(dataBuilder.Path, 0755); err != nil {
		return fmt.Errorf("mkdir %q: %w", dataBuilder.Path, err)
	}
	name := filepath.Join(dataBuilder.Path, dataBuilder.Filename)
	name = abspath(name)

	buf := new(bytes.Buffer)
	if err := dataBuilder.Template.ExecuteTemplate(buf, dataBuilder.TemplateExec, dataBuilder.Data); err != nil {
		return fmt.Errorf("template exec templ %q: %w", name, err)
	}

	if err := os.WriteFile(name, buf.Bytes(), perm); err != nil {
		return fmt.Errorf("write file %q: %w", name, err)
	}

	if g.verbose { // Print the list of written files.
		fmt.Println(name)
	}

	return nil
}

func abspath(path string) string {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abspath
}
