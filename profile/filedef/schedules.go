package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Schedule files are used to schedule a user’s workouts and may contain multiple schedule messages each representing the start time of a workout.
type Schedule struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Schedules []*mesgdef.Schedule

	UnrelatedMessages []proto.Message
}

var _ File = &Schedule{}

func NewSchedules(mesgs ...proto.Message) *Schedule {
	f := &Schedule{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Schedule) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Schedule:
		f.Schedules = append(f.Schedules, mesgdef.NewSchedule(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Schedule) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 1 // non slice fields

	size += len(f.Schedules) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, fac, mesgnum.Schedule, f.Schedules)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
