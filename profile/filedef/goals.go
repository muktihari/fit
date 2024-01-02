package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Goals files allow a user to communicate their exercise/health goals.
type Goals struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Goals             []*mesgdef.Goal
	UnrelatedMessages []proto.Message
}

var _ File = &Goals{}

func NewGoals(mesgs ...proto.Message) *Goals {
	f := &Goals{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Goals) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Goal:
		f.Goals = append(f.Goals, mesgdef.NewGoal(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Goals) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 1 // non slice fields

	size += len(f.Goals) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, fac, mesgnum.Goal, f.Goals)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
