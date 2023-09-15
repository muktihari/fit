// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/proto"
)

// Listener is Message Listener.
type Listener[T File] struct {
	file  T
	mesgc chan proto.Message
	done  chan struct{}
}

// NewListener creates mesg listener for given file T.
func NewListener[T File](file T, opts ...Option) *Listener[T] {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	l := &Listener[T]{
		file:  file,
		mesgc: make(chan proto.Message, options.channelBuffer),
		done:  make(chan struct{}),
	}

	go l.loop()

	return l
}

func (l *Listener[T]) loop() {
	for mesg := range l.mesgc {
		l.file.Add(mesg)
	}
	close(l.done)
}

func (l *Listener[T]) OnMesg(mesg proto.Message) { l.mesgc <- mesg }

// File returns the resulting file after the decode process is completed.
func (l *Listener[T]) File() T {
	close(l.mesgc)
	<-l.done
	return l.file
}

type options struct {
	channelBuffer uint
}

func defaultOptions() *options {
	return &options{
		channelBuffer: 1000,
	}
}

type Option interface{ apply(o *options) }

type fnApply func(o *options)

func (f fnApply) apply(o *options) { f(o) }

func WithChannelBuffer(size uint) Option {
	return fnApply(func(o *options) {
		if size > 0 {
			o.channelBuffer = size
		}
	})
}
