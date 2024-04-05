// Copyright 2024 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Settings files contain user and device information in the form of profiles.
type Settings struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	UserProfiles   []*mesgdef.UserProfile
	HrmProfiles    []*mesgdef.HrmProfile
	SdmProfiles    []*mesgdef.SdmProfile
	BikeProfiles   []*mesgdef.BikeProfile
	DeviceSettings []*mesgdef.DeviceSettings

	UnrelatedMessages []proto.Message
}

var _ File = &Settings{}

// NewSettings creates new Settings File.
func NewSettings(mesgs ...proto.Message) *Settings {
	f := &Settings{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

// Add adds mesg to the Settings.
func (f *Settings) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.UserProfile:
		f.UserProfiles = append(f.UserProfiles, mesgdef.NewUserProfile(&mesg))
	case mesgnum.HrmProfile:
		f.HrmProfiles = append(f.HrmProfiles, mesgdef.NewHrmProfile(&mesg))
	case mesgnum.SdmProfile:
		f.SdmProfiles = append(f.SdmProfiles, mesgdef.NewSdmProfile(&mesg))
	case mesgnum.BikeProfile:
		f.BikeProfiles = append(f.BikeProfiles, mesgdef.NewBikeProfile(&mesg))
	case mesgnum.DeviceSettings:
		f.DeviceSettings = append(f.DeviceSettings, mesgdef.NewDeviceSettings(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFit converts Settings to proto.Fit. If options is nil, default options will be used.
func (f *Settings) ToFit(options *mesgdef.Options) proto.FIT {
	var size = 1 // non slice fields

	size += len(f.UserProfiles) + len(f.HrmProfiles) + len(f.SdmProfiles) +
		len(f.BikeProfiles) + len(f.DeviceSettings) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.FIT{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	ToMesgs(&fit.Messages, options, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, options, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, options, mesgnum.UserProfile, f.UserProfiles)
	ToMesgs(&fit.Messages, options, mesgnum.HrmProfile, f.HrmProfiles)
	ToMesgs(&fit.Messages, options, mesgnum.SdmProfile, f.SdmProfiles)
	ToMesgs(&fit.Messages, options, mesgnum.BikeProfile, f.BikeProfiles)
	ToMesgs(&fit.Messages, options, mesgnum.DeviceSettings, f.DeviceSettings)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
