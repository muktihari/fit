// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/listener"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/proto"
)

const usage = "./fitprint activity.fit"

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("> missing filepath in the argument. e.g. %q\n", usage)
		os.Exit(1)
	}

	defer func(begin time.Time) {
		fmt.Printf("\ntook: %s\n", time.Since(begin))
	}(time.Now())

	name := os.Args[1]

	f, err := os.Open(name)
	if err != nil {
		fatalf("could not open file %s: %v\n", name, err)
	}
	defer f.Close()

	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()

	p := newPrinter(bw)
	defer p.Wait()

	dec := decoder.New(bufio.NewReader(f),
		decoder.WithBroadcastOnly(),
		decoder.WithMesgListener(p),
		decoder.WithIgnoreChecksum(),
	)

	_, err = dec.Decode()
	if err != nil {
		fatalf("could not decode: %v\n", err)
	}
}

func fatalf(format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

var _ listener.MesgListener = &printer{}

type printer struct {
	w       io.Writer
	eventch chan proto.Message
	done    chan struct{}
}

func newPrinter(w io.Writer) *printer {
	p := &printer{
		w:       w,
		eventch: make(chan proto.Message, 10), // 10 should be enough buffer
		done:    make(chan struct{}),
	}
	go p.handleEvent()
	return p
}

func (p *printer) handleEvent() {
	for mesg := range p.eventch {
		p.writeMesg(mesg)
	}
	close(p.done)
}

func (p *printer) Wait() {
	close(p.eventch)
	<-p.done
}

func (p *printer) OnMesg(mesg proto.Message) { p.eventch <- mesg }

func (p *printer) writeMesg(mesg proto.Message) {
	fmt.Fprintf(p.w, "%s:\n", mesg.Num.String())
	for i := range mesg.Fields {
		fmt.Fprintf(p.w, "    %s: %v\n", mesg.Fields[i].Name, formatValue(&mesg.Fields[i]))
	}
}

func formatValue(field *proto.Field) string {
	if field.Type == profile.DateTime {
		return datetime.ToTime(field.Value).String()
	}
	return fmt.Sprintf("%v %s",
		scaleoffset.ApplyAny(field.Value, field.Scale, field.Offset), field.Units)
}
