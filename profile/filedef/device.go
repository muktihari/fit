package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Device files contain information about a device’s file structure/capabilities.
type Device struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Softwares         []*mesgdef.Software
	Capabilities      []*mesgdef.Capabilities
	FileCapabilities  []*mesgdef.FileCapabilities
	MesgCapabilities  []*mesgdef.MesgCapabilities
	FieldCapabilities []*mesgdef.FieldCapabilities

	UnrelatedMessages []proto.Message
}

var _ File = &Device{}

func NewDevice(mesgs ...proto.Message) *Device {
	f := &Device{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Device) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Software:
		f.Softwares = append(f.Softwares, mesgdef.NewSoftware(&mesg))
	case mesgnum.Capabilities:
		f.Capabilities = append(f.Capabilities, mesgdef.NewCapabilities(&mesg))
	case mesgnum.FileCapabilities:
		f.FileCapabilities = append(f.FileCapabilities, mesgdef.NewFileCapabilities(&mesg))
	case mesgnum.MesgCapabilities:
		f.MesgCapabilities = append(f.MesgCapabilities, mesgdef.NewMesgCapabilities(&mesg))
	case mesgnum.FieldCapabilities:
		f.FieldCapabilities = append(f.FieldCapabilities, mesgdef.NewFieldCapabilities(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Device) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 1 // non slice fields

	size += len(f.Softwares) + len(f.Capabilities) + len(f.FileCapabilities) +
		len(f.MesgCapabilities) + len(f.FieldCapabilities) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, fac, mesgnum.Software, f.Softwares)
	ToMesgs(&fit.Messages, fac, mesgnum.Capabilities, f.Capabilities)
	ToMesgs(&fit.Messages, fac, mesgnum.FileCapabilities, f.FileCapabilities)
	ToMesgs(&fit.Messages, fac, mesgnum.MesgCapabilities, f.MesgCapabilities)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldCapabilities, f.FieldCapabilities)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
