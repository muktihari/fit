package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Sport files contain information about the user’s desired target zones.
type Sport struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	ZonesTargets []*mesgdef.ZonesTarget
	Sport        *mesgdef.Sport
	HrZones      []*mesgdef.HrZone
	PowerZones   []*mesgdef.PowerZone
	MetZones     []*mesgdef.MetZone
	SpeedZones   []*mesgdef.SpeedZone
	CadenceZones []*mesgdef.CadenceZone

	UnrelatedMessages []proto.Message
}

var _ File = &Sport{}

func NewSport(mesgs ...proto.Message) *Sport {
	f := &Sport{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Sport) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.ZonesTarget:
		f.ZonesTargets = append(f.ZonesTargets, mesgdef.NewZonesTarget(&mesg))
	case mesgnum.Sport:
		f.Sport = mesgdef.NewSport(&mesg)
	case mesgnum.HrZone:
		f.HrZones = append(f.HrZones, mesgdef.NewHrZone(&mesg))
	case mesgnum.PowerZone:
		f.PowerZones = append(f.PowerZones, mesgdef.NewPowerZone(&mesg))
	case mesgnum.MetZone:
		f.MetZones = append(f.MetZones, mesgdef.NewMetZone(&mesg))
	case mesgnum.SpeedZone:
		f.SpeedZones = append(f.SpeedZones, mesgdef.NewSpeedZone(&mesg))
	case mesgnum.CadenceZone:
		f.CadenceZones = append(f.CadenceZones, mesgdef.NewCadenceZone(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Sport) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 2 // non slice fields

	size += len(f.ZonesTargets) + len(f.HrZones) + len(f.PowerZones) +
		len(f.MetZones) + len(f.SpeedZones) + len(f.CadenceZones) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, fac, mesgnum.ZonesTarget, f.ZonesTargets)

	if f.Sport != nil {
		fit.Messages = append(fit.Messages, f.Sport.ToMesg(fac))
	}

	ToMesgs(&fit.Messages, fac, mesgnum.HrZone, f.HrZones)
	ToMesgs(&fit.Messages, fac, mesgnum.PowerZone, f.PowerZones)
	ToMesgs(&fit.Messages, fac, mesgnum.MetZone, f.MetZones)
	ToMesgs(&fit.Messages, fac, mesgnum.SpeedZone, f.SpeedZones)
	ToMesgs(&fit.Messages, fac, mesgnum.CadenceZone, f.CadenceZones)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
