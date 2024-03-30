// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// ActivitySummary is a compact version of the activity file and contain only activity, session and lap messages
type ActivitySummary struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number, time_created

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Activity *mesgdef.Activity
	Sessions []*mesgdef.Session
	Laps     []*mesgdef.Lap

	UnrelatedMessages []proto.Message
}

var _ File = &ActivitySummary{}

// NewActivitySummary creates new ActivitySummary File.
func NewActivitySummary(mesgs ...proto.Message) *ActivitySummary {
	f := &ActivitySummary{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

// Add adds mesg to the ActivitySummary.
func (f *ActivitySummary) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Activity:
		f.Activity = mesgdef.NewActivity(&mesg)
	case mesgnum.Session:
		f.Sessions = append(f.Sessions, mesgdef.NewSession(&mesg))
	case mesgnum.Lap:
		f.Laps = append(f.Laps, mesgdef.NewLap(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFit converts ActivitySummary to proto.Fit. If options is nil, default options will be used.
func (f *ActivitySummary) ToFit(options *mesgdef.Options) proto.Fit {
	var size = 2 // non slice fields

	size += len(f.Sessions) + len(f.Laps) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	ToMesgs(&fit.Messages, options, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, options, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.Activity != nil {
		fit.Messages = append(fit.Messages, f.Activity.ToMesg(options))
	}

	ToMesgs(&fit.Messages, options, mesgnum.Session, f.Sessions)
	ToMesgs(&fit.Messages, options, mesgnum.Lap, f.Laps)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
