package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Totals files are used to summarize a user’s activities and may contain multiple totals messages each representing
// summaries of a different activity type/sport
type Totals struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Totals []*mesgdef.Totals

	UnrelatedMessages []proto.Message
}

var _ File = &Totals{}

func NewTotals(mesgs ...proto.Message) *Totals {
	f := &Totals{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Totals) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Totals:
		f.Totals = append(f.Totals, mesgdef.NewTotals(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Totals) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 3 // non slice fields

	size += len(f.Totals) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	ToMesgs(&fit.Messages, fac, mesgnum.Totals, f.Totals)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages)

	return fit
}
