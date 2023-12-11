// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/proto"
)

// Listener is Message Listener.
type Listener[F any, T FilePtr[F]] struct {
	file    F
	options *options
	mesgc   chan proto.Message
	done    chan struct{}
}

// FilePtr is a type constraint for pointer of File.
type FilePtr[T any] interface {
	*T
	File
}

// NewListener creates mesg listener for given file T.
func NewListener[F any, T FilePtr[F]](opts ...Option) *Listener[F, T] {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	l := &Listener[F, T]{
		file:    *new(F),
		options: options,
		mesgc:   make(chan proto.Message, options.channelBuffer),
		done:    make(chan struct{}),
	}

	go l.loop()

	return l
}

func (l *Listener[F, T]) loop() {
	for mesg := range l.mesgc {
		T(&l.file).Add(mesg)
	}
	close(l.done)
}

func (l *Listener[F, T]) OnMesg(mesg proto.Message) { l.mesgc <- mesg }

// Close closes channel and wait until all messages is consumed.
func (l *Listener[F, T]) Close() {
	close(l.mesgc)
	<-l.done
}

// File returns the resulting file after the a single decode process is completed. This will reset fields used by listener
// and the listener is ready to be used for next chained FIT file.
func (l *Listener[F, T]) File() T {
	l.Close()

	file := l.file
	l.reset()

	go l.loop()

	return T(&file)
}

func (l *Listener[F, T]) reset() {
	l.file = *new(F)
	l.mesgc = make(chan proto.Message, l.options.channelBuffer)
	l.done = make(chan struct{})
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
