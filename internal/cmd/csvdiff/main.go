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
	cols, _ := reader.Read()

	ss := make([]string, 0, len(cols)/3)
	for i := 0; i < len(cols); i += 3 {
		strbuf.Reset()
		strbuf.WriteString(cols[i])
		strbuf.WriteByte(',')
		strbuf.WriteString(cols[i+1])
		strbuf.WriteByte(',')
		strbuf.WriteString(cols[i+2])
		strbuf.WriteByte(',')
		ss = append(ss, strbuf.String()) // no need to copy, strings is immutable.
	}
	return ss
}
