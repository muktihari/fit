// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("USAGE: csvcmp file1.csv file2.csv")
		return
	}

	f1, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	s1 := bufio.NewScanner(f1)
	s2 := bufio.NewScanner(f2)

	var diffCounter int
	var lineNumber int
	for s1.Scan() {
		input1 := s1.Text()
		s2.Scan()
		input2 := s2.Text()

		input1 = strings.TrimRight(input1, ",")
		input2 = strings.TrimRight(input2, ",")

		// Remove UTF-8 BOM if exist
		input1 = strings.Trim(input1, "\xef\xbb\xbf")
		input2 = strings.Trim(input2, "\xef\xbb\xbf")

		if input1 != input2 {
			fmt.Printf("[%d] Line Number: %d\n", diffCounter+1, lineNumber+1)
			fmt.Printf("input1: %s\n", input1)
			fmt.Printf("input2: %s\n", input2)

			diff := cmp.Diff(splitByField(input1), splitByField(input2))
			diff = strings.Replace(diff, "[]string{", "", 1)
			diff = strings.Replace(diff, "}", "", 1)
			fmt.Printf("[diff]:%s", diff)
			diffCounter++
		}
		lineNumber++
	}

	if diffCounter != 0 {
		fmt.Println()
	}

	fmt.Printf("[Total diff: %d]\n", diffCounter)
}

var strbuf = new(strings.Builder)
var buf = new(bytes.Buffer)

func splitByField(s string) []string {
	buf.Reset()
	buf.WriteString(s)

	reader := csv.NewReader(buf)
	reader.FieldsPerRecord = -1
	cols, _ := reader.Read()

	if len(cols) < 2 {
		return nil
	}

	ss := make([]string, 0, ((len(cols)-2)/3)+2)

	strbuf.Reset()
	strbuf.WriteString(cols[0]) // Data/Definition
	strbuf.WriteByte(',')
	strbuf.WriteString(cols[1]) // LocalNum
	ss = append(ss, strbuf.String())

	for i := 2; i < len(cols)-2; i += 3 {
		strbuf.Reset()
		strbuf.WriteString(cols[i])
		strbuf.WriteByte(',')
		strbuf.WriteString(cols[i+1])
		strbuf.WriteByte(',')
		strbuf.WriteString(cols[i+2])
		strbuf.WriteByte(',')
		ss = append(ss, strbuf.String())
	}
	return ss
}
