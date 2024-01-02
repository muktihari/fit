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

// BloodPressure files contain time-stamped discrete measurement data of blood pressure.
type BloodPressure struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	UserProfile    *mesgdef.UserProfile
	BloodPressures []*mesgdef.BloodPressure
	DeviceInfos    []*mesgdef.DeviceInfo

	UnrelatedMessages []proto.Message
}

var _ File = &BloodPressure{}

func NewBloodPressure(mesgs ...proto.Message) *BloodPressure {
	f := &BloodPressure{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *BloodPressure) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(&mesg)
	case mesgnum.BloodPressure:
		f.BloodPressures = append(f.BloodPressures, mesgdef.NewBloodPressure(&mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *BloodPressure) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 2 // non slice fields

	size += len(f.BloodPressures) + len(f.DeviceInfos) +
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

	ToMesgs(&fit.Messages, fac, mesgnum.BloodPressure, f.BloodPressures)
	ToMesgs(&fit.Messages, fac, mesgnum.DeviceInfo, f.DeviceInfos)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
