package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// MonitoringDaily files follow the same format as monitoring files, however data is logged at 24 hour time intervals.
type MonitoringDaily struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number, time_created, number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	MonitoringInfo *mesgdef.MonitoringInfo
	Monitorings    []*mesgdef.Monitoring // required fields: timestamp
	DeviceInfos    []*mesgdef.DeviceInfo

	UnrelatedMessages []proto.Message
}

var _ File = &MonitoringDaily{}

func NewMonitoringDaily(mesgs ...proto.Message) *MonitoringDaily {
	f := &MonitoringDaily{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *MonitoringDaily) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.MonitoringInfo:
		f.MonitoringInfo = mesgdef.NewMonitoringInfo(&mesg)
	case mesgnum.Monitoring:
		f.Monitorings = append(f.Monitorings, mesgdef.NewMonitoring(&mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *MonitoringDaily) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 2 // non slice fields

	size += len(f.Monitorings) + len(f.DeviceInfos) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.MonitoringInfo != nil {
		fit.Messages = append(fit.Messages, f.MonitoringInfo.ToMesg(fac))
	}

	ToMesgs(&fit.Messages, fac, mesgnum.Monitoring, f.Monitorings)
	ToMesgs(&fit.Messages, fac, mesgnum.DeviceInfo, f.DeviceInfos)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
