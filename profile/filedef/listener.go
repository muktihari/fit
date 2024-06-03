// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
	"golang.org/x/exp/slices"
)

// Listener is Message Listener.
type Listener struct {
	options options
	file    File
	poolc   chan []proto.Field // pool of reusable objects to minimalize slice allocations. do not close this channel.
	mesgc   chan proto.Message // queue messages to be processed concurrently.
	done    chan struct{}
	active  bool
}

// FileSets is a set of file type mapped to a function to create that File.
// This alias is created for documentation purpose.
type FileSets = map[typedef.File]func() File

type options struct {
	fileSets      FileSets
	channelBuffer uint
}

func defaultOptions() options {
	return options{
		fileSets:      PredefinedFileSet(),
		channelBuffer: 128,
	}
}

// PredefinedFileSet is a list of default filesets used in listener, it's exported so user can append their own types and register it as an option.
func PredefinedFileSet() FileSets {
	return FileSets{
		typedef.FileActivity:        func() File { return NewActivity() },
		typedef.FileActivitySummary: func() File { return NewActivitySummary() },
		typedef.FileBloodPressure:   func() File { return NewBloodPressure() },
		typedef.FileCourse:          func() File { return NewCourse() },
		typedef.FileDevice:          func() File { return NewDevice() },
		typedef.FileGoals:           func() File { return NewGoals() },
		typedef.FileMonitoringA:     func() File { return NewMonitoringAB() },
		typedef.FileMonitoringB:     func() File { return NewMonitoringAB() },
		typedef.FileMonitoringDaily: func() File { return NewMonitoringDaily() },
		typedef.FileSchedules:       func() File { return NewSchedules() },
		typedef.FileSegment:         func() File { return NewSegment() },
		typedef.FileSegmentList:     func() File { return NewSegmentList() },
		typedef.FileSettings:        func() File { return NewSettings() },
		typedef.FileSport:           func() File { return NewSport() },
		typedef.FileTotals:          func() File { return NewTotals() },
		typedef.FileWeight:          func() File { return NewWeight() },
		typedef.FileWorkout:         func() File { return NewWorkout() },
	}
}

type Option interface{ apply(o *options) }

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

// WithChannelBuffer sets the size of buffered channel, default is 128.
func WithChannelBuffer(size uint) Option {
	return fnApply(func(o *options) { o.channelBuffer = size })
}

// WithFileSets sets what kind of file listener should listen to, when we encounter a file type that is not listed in fileset,
// that file type will be skipped. This will replace the default filesets registered in listener, if you intend to append your own
// file types, please call PredefinedFileSet() and add your file types.
func WithFileSets(fileSets FileSets) Option {
	return fnApply(func(o *options) {
		if o.fileSets != nil {
			o.fileSets = fileSets
		}
	})
}

var _ decoder.MesgListener = (*Listener)(nil)

// NewListener creates mesg listener.
func NewListener(opts ...Option) *Listener {
	l := &Listener{
		options: defaultOptions(),
		active:  true,
	}
	for i := range opts {
		opts[i].apply(&l.options)
	}

	l.reset()

	l.poolc = make(chan []proto.Field, l.options.channelBuffer)
	for i := uint(0); i < l.options.channelBuffer; i++ {
		l.poolc <- nil // fill pool with nil slice, alloc as needed.
	}

	go l.loop()

	return l
}

func (l *Listener) loop() {
	for mesg := range l.mesgc {
		l.processMesg(mesg)
		l.poolc <- mesg.Fields // put the slice back to the pool to be recycled.
	}
	close(l.done)
}

func (l *Listener) processMesg(mesg proto.Message) {
	if mesg.Num == mesgnum.FileId {
		fileType := typedef.File(mesg.FieldValueByNum(fieldnum.FileIdType).Uint8())
		fnNew, ok := l.options.fileSets[fileType]
		if !ok {
			return
		}
		l.file = fnNew()
	}
	if l.file == nil {
		return // No file is created since not defined in fileSets. Skip.
	}
	l.file.Add(mesg)
}

func (l *Listener) OnMesg(mesg proto.Message) {
	if !l.active {
		l.reset()
		go l.loop()
		l.active = true
	}
	l.mesgc <- l.prep(mesg)
}

func (l *Listener) prep(mesg proto.Message) proto.Message {
	fields := <-l.poolc

	if cap(fields) < len(mesg.Fields) {
		fields = make([]proto.Field, len(mesg.Fields))
	}
	fields = fields[:len(mesg.Fields)]
	copy(fields, mesg.Fields)
	mesg.Fields = fields

	// Must clone DeveloperFields since it is being referenced in mesgdef's structs.
	mesg.DeveloperFields = slices.Clone(mesg.DeveloperFields)

	return mesg
}

// Close closes channel and wait until all messages is consumed.
func (l *Listener) Close() {
	if !l.active {
		return
	}
	close(l.mesgc)
	<-l.done
	l.active = false
}

// File returns the resulting file after the a single decode process is completed. If we the current decoded result is not listed
// in fileSets, nil will be returned, it's recommended to use switch type assertion to check. This will reset fields used by listener
// and the listener is ready to be used for next chained FIT file.
func (l *Listener) File() File {
	l.Close()
	return l.file
}

func (l *Listener) reset() {
	l.file = nil
	l.mesgc = make(chan proto.Message, l.options.channelBuffer)
	l.done = make(chan struct{})
}
