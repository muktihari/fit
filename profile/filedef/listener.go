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
)

// Listener is a common file types listener that implement decoder.MesgListener
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
	fileSets      [256]func() File
	channelBuffer uint
}

func defaultFileSets() [256]func() File {
	return [256]func() File{
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

func defaultOptions() options {
	return options{
		fileSets:      defaultFileSets(),
		channelBuffer: 128,
	}
}

// PredefinedFileSet is a list of default filesets used in listener, it's exported so user can append their own types and register it as an option.
func PredefinedFileSet() FileSets {
	m := make(map[typedef.File]func() File)
	for i, v := range defaultFileSets() {
		m[typedef.File(i)] = v
	}
	return m
}

// Option is Listener's option.
type Option func(o *options)

// WithChannelBuffer sets the size of buffered channel, default is 128.
func WithChannelBuffer(size uint) Option {
	return func(o *options) { o.channelBuffer = size }
}

// WithFileSets sets what kind of file listener should listen to, when we encounter a file type that is not listed in fileset,
// that file type will be skipped. This will replace the default listener's filesets, if you intend to append your own
// file types, please call PredefinedFileSet() and add your file type before using this option; or use WithFileFunc instead.
func WithFileSets(fileSets FileSets) Option {
	return func(o *options) {
		o.fileSets = [256]func() File{} // Clear all.
		for file, fn := range fileSets {
			o.fileSets[file] = fn
		}
	}
}

// WithFileFunc sets File with its File creator function. It overrides the default options.
func WithFileFunc(f typedef.File, fn func() File) Option {
	return func(o *options) { o.fileSets[f] = fn }
}

var _ decoder.MesgListener = (*Listener)(nil)

// NewListener creates new common file types listener that implement decoder.MesgListener.
// This will handle message conversion from proto.Message received from Decoder into
// mesgdef's structure and group it by its correspoding defined file types.
func NewListener(opts ...Option) *Listener {
	l := new(Listener)
	l.Reset(opts...)
	return l
}

// Reset resets the Listener for reuse. It resets options to default options so any
// options needs to be inputed again. It is similar to NewListener() but it retains
// the underlying storage for use by future decode to reduce memory allocs.
func (l *Listener) Reset(opts ...Option) {
	l.Close()
	prevChannelBuffer := l.options.channelBuffer
	l.options = defaultOptions()
	for i := range opts {
		opts[i](&l.options)
	}

	if prevChannelBuffer != l.options.channelBuffer {
		prevPoolc := l.poolc
		l.poolc = make(chan []proto.Field, l.options.channelBuffer)
		for i := uint(0); i < l.options.channelBuffer; i++ {
			select {
			case v := <-prevPoolc:
				l.poolc <- v // fill with previously allocated slice.
			default:
				l.poolc <- nil // fill pool with nil slice, alloc as needed.
			}
		}
	}
	l.reset()
}

func (l *Listener) reset() {
	l.file = nil
	l.mesgc = make(chan proto.Message, l.options.channelBuffer)
	l.done = make(chan struct{})
	l.active = true

	go l.loop()
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
		fileType := mesg.FieldValueByNum(fieldnum.FileIdType).Uint8()
		fn := l.options.fileSets[fileType]
		if fn == nil {
			return
		}
		l.file = fn()
	}
	if l.file == nil {
		return // No file is created since not defined in fileSets. Skip.
	}
	l.file.Add(mesg)
}

func (l *Listener) OnMesg(mesg proto.Message) {
	if !l.active {
		l.reset()
	}

	mesg.Fields = append((<-l.poolc)[:0], mesg.Fields...)
	// Must clone DeveloperFields since it is being referenced in mesgdef's structs.
	mesg.DeveloperFields = append(mesg.DeveloperFields[:0:0], mesg.DeveloperFields...)

	l.mesgc <- mesg
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
