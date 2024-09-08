// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package remover

import (
	"sync"

	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type options struct {
	removeNums          map[typedef.MesgNum]struct{}
	removeUnknown       bool
	removeDeveloperData bool
}

// Option is remover options.
type Option func(o *options)

// WithRemoveUnknown directs Remove to removes unknown messages not defined in Global Profile (Profile.xlsx).
func WithRemoveUnknown() Option {
	return func(o *options) { o.removeUnknown = true }
}

// WithRemoveMesgNums directs Remove to removes messages based on given message numbers.
func WithRemoveMesgNums(nums map[typedef.MesgNum]struct{}) Option {
	return func(o *options) { o.removeNums = nums }
}

// WithRemoveDeveloperData directs Remove to removes all developer data: DeveloperId, FieldDescription
// and all DeveloperFields in messages.
func WithRemoveDeveloperData() Option {
	return func(o *options) { o.removeDeveloperData = true }
}

// Remove removes messages based on given options.
func Remove(fit *proto.FIT, opts ...Option) {
	var o options
	for i := range opts {
		opts[i](&o)
	}

	if o.removeUnknown {
		removeUnknownMessages(fit)
	}

	if len(o.removeNums) != 0 {
		removeNums(fit, o.removeNums)
	}

	if o.removeDeveloperData {
		removeDeveloperData(fit)
	}
}

// only populate when used.
var knownNums map[typedef.MesgNum]struct{}
var once sync.Once

func removeUnknownMessages(fit *proto.FIT) {
	once.Do(func() {
		knownNums = make(map[typedef.MesgNum]struct{})
		nums := typedef.ListMesgNum()
		for _, num := range nums {
			if num == typedef.MesgNumMfgRangeMin {
				continue
			}
			if num == typedef.MesgNumMfgRangeMax {
				continue
			}
			knownNums[num] = struct{}{}
		}
	})
	var valid int
	for i, mesg := range fit.Messages {
		if _, ok := knownNums[mesg.Num]; !ok {
			continue
		}
		if i != valid {
			fit.Messages[i], fit.Messages[valid] = fit.Messages[valid], fit.Messages[i]
		}
		valid++
	}
	fit.Messages = fit.Messages[:valid]
}

func removeNums(fit *proto.FIT, removeNums map[typedef.MesgNum]struct{}) {
	var valid int
	for i, mesg := range fit.Messages {
		if _, ok := removeNums[mesg.Num]; ok {
			continue
		}
		if i != valid {
			fit.Messages[i], fit.Messages[valid] = fit.Messages[valid], fit.Messages[i]
		}
		valid++
	}
	fit.Messages = fit.Messages[:valid]
}

func removeDeveloperData(fit *proto.FIT) {
	var valid int
	for i, mesg := range fit.Messages {
		if mesg.Num == mesgnum.DeveloperDataId || mesg.Num == mesgnum.FieldDescription {
			continue
		}
		fit.Messages[i].DeveloperFields = nil
		if i != valid {
			fit.Messages[i], fit.Messages[valid] = fit.Messages[valid], fit.Messages[i]
		}
		valid++
	}
	fit.Messages = fit.Messages[:valid]
}
