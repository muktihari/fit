// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reducer

import (
	"fmt"
	"math"
	"strconv"

	"github.com/muktihari/carto/rdp"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

type errorString string

func (e errorString) Error() string { return string(e) }

const (
	errBadArgument = errorString("bad arguments")
)

type method byte

const (
	methodUnknown method = iota
	methodRDP
	methodDistanceInterval
	methodTimeInterval
)

func (m method) String() string {
	switch m {
	case methodRDP:
		return "method_rdp"
	case methodDistanceInterval:
		return "method_distance"
	case methodTimeInterval:
		return "method_time"
	default:
		return "unknown(" + strconv.Itoa(int(m)) + ")"
	}
}

type options struct {
	method     method  // reduce method
	rdpEpsilon float64 // tolerance should be > 0
	interval   uint32  // meters when by distance, seconds when by time.
}

// Options is reducer options.
type Option func(*options)

// WithRDP directs Reduce to use Ramer-Douglas-Peucker algorithm to simplify by GPS points.
func WithRDP(epsilon float64) Option {
	return func(o *options) {
		o.method = methodRDP
		o.rdpEpsilon = epsilon
	}
}

// WithDistanceInterval directs Reduce to by specific distance interval.
func WithDistanceInterval(interval float64) Option {
	return func(o *options) {
		o.method = methodDistanceInterval
		o.interval = uint32(interval * 100)
	}
}

// WithTimeInterval directs Reduce to by specific time interval.
func WithTimeInterval(interval uint32) Option {
	return func(o *options) {
		o.method = methodTimeInterval
		o.interval = interval
	}
}

// Reduce reduces message size by simplifying a curve of points using Ramer-Douglas-Peucker
// algorithm to a similar curve with fewer points, then only keep messages having those points.
func Reduce(fit *proto.FIT, opts ...Option) error {
	var o options
	for i := range opts {
		opts[i](&o)
	}

	switch o.method {
	case methodRDP:
		if o.rdpEpsilon == 0 {
			return fmt.Errorf("epsilon should be > 0, got: %f: %v", o.rdpEpsilon, errBadArgument)
		}
		return reduceByRDP(fit, o.rdpEpsilon)
	case methodDistanceInterval:
		if o.interval == 0 {
			return fmt.Errorf("interval should be > 0, got: %d: %v", o.interval, errBadArgument)
		}
		return reduceByDistanceInterval(fit, o.interval)
	case methodTimeInterval:
		if o.interval == 0 {
			return fmt.Errorf("interval should be > 0, got: %d: %v", o.interval, errBadArgument)
		}
		return reduceByTimeInterval(fit, o.interval)
	default:
		return fmt.Errorf("method is undefined: %s: %v", o.method, errBadArgument)
	}
}

func reduceByRDP(fit *proto.FIT, epsilon float64) error {
	epsilon /= 1000

	var n int
	for _, mesg := range fit.Messages {
		if mesg.Num == mesgnum.Record {
			n++
		}
	}

	recordIndexes := make([]int, 0, n)
	points := make([]rdp.Point, 0, n)
	for i, mesg := range fit.Messages {
		if mesg.Num != mesgnum.Record {
			continue
		}

		recordIndexes = append(recordIndexes, i)

		x := semicircles.ToDegrees(mesg.FieldValueByNum(fieldnum.RecordPositionLat).Int32())
		y := semicircles.ToDegrees(mesg.FieldValueByNum(fieldnum.RecordPositionLong).Int32())
		if math.IsNaN(x) || math.IsNaN(y) {
			continue
		}

		points = append(points, rdp.Point{X: x, Y: y, Index: i})
	}

	if len(points) == 0 {
		return fmt.Errorf("zero valid points")
	}

	points = rdp.Simplify(points, epsilon)

	fragments := findFragments(recordIndexes, points)
	fit.Messages = defragment(fit.Messages, fragments)

	return nil
}

// findFragments finds omitted records based on RDP result and returning the index as fragments.
func findFragments(recordIndexes []int, points []rdp.Point) (fragments []int) {
	var i, j int
	// Reslicing is safe as index will always move forward
	fragments = recordIndexes[:0]
	for i < len(recordIndexes) && j < len(points) {
		if recordIndexes[i] < points[j].Index {
			fragments = append(fragments, recordIndexes[i])
			i++
		} else if recordIndexes[i] == points[j].Index {
			i, j = i+1, j+1
		} else {
			j++
		}
	}
	fragments = append(fragments, recordIndexes[i:]...)
	return fragments
}

// defragment defrags mesgs by filling every index occupied by the fragmented data.
func defragment(mesgs []proto.Message, fragments []int) []proto.Message {
	var i, valid, cur int
	for i = 0; i < len(mesgs); i++ {
		if cur == len(fragments) {
			break
		}
		if i == fragments[cur] {
			cur++
			continue
		}
		if i != valid {
			mesgs[i], mesgs[valid] = mesgs[valid], mesgs[i]
		}
		valid++
	}
	return append(mesgs[:valid], mesgs[i:]...)
}

func reduceByDistanceInterval(fit *proto.FIT, threshold uint32) error {
	var (
		last               uint32
		valid              int
		firstRecordReached bool
	)
	for i := 0; i < len(fit.Messages); i++ {
		mesg := &fit.Messages[i]
		if mesg.Num == mesgnum.Record {
			d := mesg.FieldValueByNum(fieldnum.RecordDistance).Uint32()
			if !firstRecordReached {
				firstRecordReached = true
				if d != basetype.Uint32Invalid {
					last = d
				}
				valid++ // Preserve first record
				continue
			}
			if d == basetype.Uint32Invalid {
				continue
			}
			if d-last < threshold {
				continue
			}
			last = d
		}
		if i != valid {
			fit.Messages[i], fit.Messages[valid] = fit.Messages[valid], fit.Messages[i]
		}
		valid++
	}
	fit.Messages = fit.Messages[:valid]
	return nil
}

func reduceByTimeInterval(fit *proto.FIT, threshold uint32) error {
	var (
		last               uint32
		valid              int
		firstRecordReached bool
	)
	for i := 0; i < len(fit.Messages); i++ {
		mesg := &fit.Messages[i]
		if mesg.Num == mesgnum.Record {
			t := mesg.FieldValueByNum(fieldnum.RecordTimestamp).Uint32()
			if !firstRecordReached {
				firstRecordReached = true
				if t != basetype.Uint32Invalid {
					last = t
				}
				valid++ // Preserve first record
				continue
			}
			if t == basetype.Uint32Invalid {
				continue
			}
			if t-last < threshold {
				continue
			}
			last = t
		}
		if i != valid {
			fit.Messages[i], fit.Messages[valid] = fit.Messages[valid], fit.Messages[i]
		}
		valid++
	}
	fit.Messages = fit.Messages[:valid]
	return nil
}
