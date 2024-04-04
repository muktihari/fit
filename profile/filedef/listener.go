// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type options struct {
	channelBuffer uint
	fileSets      FileSets
}

func defaultOptions() *options {
	return &options{
		channelBuffer: 1000,
		fileSets:      PredefinedFileSet(),
	}
}

// PredefinedFileSet is a list of default filesets used in listener, it's exported so user can append their own types and register it as an option.
func PredefinedFileSet() FileSets {
	return FileSets{
		typedef.FileDevice:          func() File { return NewDevice() },
		typedef.FileSettings:        func() File { return NewSettings() },
		typedef.FileSport:           func() File { return NewSport() },
		typedef.FileBloodPressure:   func() File { return NewBloodPressure() },
		typedef.FileWeight:          func() File { return NewWeight() },
		typedef.FileWorkout:         func() File { return NewWorkout() },
		typedef.FileActivity:        func() File { return NewActivity() },
		typedef.FileCourse:          func() File { return NewCourse() },
		typedef.FileGoals:           func() File { return NewGoals() },
		typedef.FileTotals:          func() File { return NewTotals() },
		typedef.FileSchedules:       func() File { return NewSchedules() },
		typedef.FileMonitoringA:     func() File { return NewMonitoringAB() },
		typedef.FileMonitoringB:     func() File { return NewMonitoringAB() },
		typedef.FileMonitoringDaily: func() File { return NewMonitoringDaily() },
		typedef.FileSegment:         func() File { return NewSegment() },
		typedef.FileSegmentList:     func() File { return NewSegmentList() },
	}
}

type Option interface{ apply(o *options) }

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

// WithChannelBuffer sets the size of buffered channel, default is 1000.
func WithChannelBuffer(size uint) Option {
	return fnApply(func(o *options) {
		o.channelBuffer = size
	})
}

// FileSets is a set of file type mapped to a function to create that File.
type FileSets = map[typedef.File]func() File

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

// Listener is Message Listener.
type Listener struct {
	options    *options
	activeFile File
	mesgc      chan proto.Message
	done       chan struct{}
}

// NewListener creates mesg listener.
func NewListener(opts ...Option) *Listener {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	l := &Listener{
		options: options,
		mesgc:   make(chan proto.Message, options.channelBuffer),
		done:    make(chan struct{}),
	}

	go l.loop()

	return l
}

func (l *Listener) loop() {
	for mesg := range l.mesgc {
		if mesg.Num == mesgnum.FileId {
			fileType := typedef.File(mesg.FieldValueByNum(fieldnum.FileIdType).Uint8())
			newFile, ok := l.options.fileSets[fileType]
			if !ok {
				continue
			}
			l.activeFile = newFile()
		}
		if l.activeFile == nil {
			continue // No file is created since not defined in fileSets. Skip.
		}
		l.activeFile.Add(mesg)
	}
	close(l.done)
}

func (l *Listener) OnMesg(mesg proto.Message) {
	l.mesgc <- mesg
	// if mesg.Num == mesgnum.FileId {
	// 	fileType := typedef.File(mesg.FieldValueByNum(fieldnum.FileIdType).Uint8())
	// 	newFile, ok := l.options.fileSets[fileType]
	// 	if !ok {
	// 		return
	// 	}
	// 	l.activeFile = newFile()
	// }
	// if l.activeFile == nil {
	// 	return
	// }
	// l.activeFile.Add(mesg)
}

// Close closes channel and wait until all messages is consumed.
func (l *Listener) Close() {
	close(l.mesgc)
	<-l.done
}

// File returns the resulting file after the a single decode process is completed. If we the current decoded result is not listed
// in fileSets, nil will be returned, it's  recommended to use switch type assertion to check. This will reset fields used by listener
// and the listener is ready to be used for next chained FIT file.
func (l *Listener) File() File {
	l.Close()

	file := l.activeFile
	l.reset()

	go l.loop()

	return file
}

func (l *Listener) reset() {
	l.activeFile = nil
	l.mesgc = make(chan proto.Message, l.options.channelBuffer)
	l.done = make(chan struct{})
}
