// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// SegmentList files maintain a list of available segments on the device.
type SegmentList struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	FileCreator  *mesgdef.FileCreator
	SegmentFiles []*mesgdef.SegmentFile

	UnrelatedMessages []proto.Message
}

var _ File = &SegmentList{}

// NewSegmentList creates new SegmentList File.
func NewSegmentList(mesgs ...proto.Message) *SegmentList {
	f := &SegmentList{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

// Add adds mesg to the SegmentList.
func (f *SegmentList) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.FileCreator:
		f.FileCreator = mesgdef.NewFileCreator(&mesg)
	case mesgnum.SegmentFile:
		f.SegmentFiles = append(f.SegmentFiles, mesgdef.NewSegmentFile(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFit converts SegmentList to proto.Fit. If options is nil, default options will be used.
func (f *SegmentList) ToFit(options *mesgdef.Options) proto.Fit {
	var size = 2 // non slice fields

	size += len(f.SegmentFiles) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	ToMesgs(&fit.Messages, options, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, options, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.FileCreator != nil {
		fit.Messages = append(fit.Messages, f.FileCreator.ToMesg(options))
	}

	ToMesgs(&fit.Messages, options, mesgnum.SegmentFile, f.SegmentFiles)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
