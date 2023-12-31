// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Weight contains time-stamped discrete measurement data of weight.
type Weight struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	UserProfile  *mesgdef.UserProfile
	WeightScales []*mesgdef.WeightScale
	DeviceInfos  []*mesgdef.DeviceInfo

	UnrelatedMessages []proto.Message
}

var _ File = &Weight{}

func NewWeight(mesgs ...proto.Message) *Weight {
	f := &Weight{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Weight) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(&mesg)
	case mesgnum.WeightScale:
		f.WeightScales = append(f.WeightScales, mesgdef.NewWeightScale(&mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Weight) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 2 // non slice fields

	size += len(f.WeightScales) + len(f.DeviceInfos) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.UserProfile != nil {
		fit.Messages = append(fit.Messages, f.UserProfile.ToMesg(fac))
	}

	ToMesgs(&fit.Messages, fac, mesgnum.WeightScale, f.WeightScales)
	ToMesgs(&fit.Messages, fac, mesgnum.DeviceInfo, f.DeviceInfos)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
